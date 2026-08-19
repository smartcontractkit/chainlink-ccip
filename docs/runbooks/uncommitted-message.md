---
name: uncommitted-message
description: Triage a message that has not landed on the destination offramp within the expected time window.
trigger: "message reported in-flight for > 15m: onramp seq num observed, no corresponding offramp commit"
severity: page
owner: ccip-commit-oncall
inputs:
  sourceChain: {type: string, description: "sourceChainName label value, e.g. chain-a (sourceChainID/sourceChainSelector also work -- see the metric reference). May be handed to you as a chain ID, chain selector, or name -- see docs/runbooks/chain-identifiers.md to translate if you need a form other than what you were given"}
  destChain: {type: string, description: "destChainID label value of the destination chain, e.g. chain-b (destChainName/destChainSelector also work). Same translation note as sourceChain"}
  seqNum: {type: integer, description: "onramp sequence number of the message under investigation (X)"}
  msgID: {type: string, description: "message ID, hex, for cross-referencing logs/explorers"}
  fRoleDON: {type: integer, required: false, description: "optional. If known, expected live oracle count for destChain's DON is 3*fRoleDON+1 and the consensus threshold is 2*fRoleDON+1. Omit if uncertain (e.g. local devenvs with a possible bootstrap-node offset) -- scenario3's H0 check reports the raw live oracle count either way, it just can't grade it against a threshold without this."}
related:
  - docs/metrics/commit-metrics.md        # design rationale + full metric-by-metric findings
  - docs/metrics/reader-metrics.md        # data-source (reader + config poller) metrics & gaps
  - devenv/dashboards/commit-plugin.json  # Grafana rendering of this runbook (Guided Debug section)
  - docs/runbooks/chain-identifiers.md    # translating chain ID / chain selector / chain name
status: living
---

- [Runbook: uncommitted message](#runbook-uncommitted-message)
   * [For agents](#for-agents)
      + [Resolving inputs before you start](#resolving-inputs-before-you-start)
      + [A note on counter naming](#a-note-on-counter-naming)
   * [Decision graph](#decision-graph)
   * [Steps](#steps)
      + [step0.1 — has the plugin discovered its contracts?](#step0.1-has-the-plugin-discovered-its-contracts)
      + [step0.2 - heartbeat — is the plugin alive at all?](#step0.2-heartbeat-is-the-plugin-alive-at-all)
      + [step1 — is this even a commit-plugin problem?](#step1-is-this-even-a-commit-plugin-problem)
       + [step2 — has the plugin observed it on-chain yet?](#step2-has-the-plugin-observed-it-on-chain-yet)
       + [step2b — can the DON agree on this chain's per-chain values?](#step2b-can-the-don-agree-on-this-chains-per-chain-values)
       + [step3 — is the round producing outcomes at all?](#step3-is-the-round-producing-outcomes-at-all)
      + [step4 — is chain A or B cursed?](#step4-is-chain-a-or-b-cursed)
      + [step5 — is a report being built but never landing?](#step5-is-a-report-being-built-but-never-landing)
   * [Deep dives](#deep-dives)
      + [Scenario 2 — report transmission stuck (step5)](#scenario-2-report-transmission-stuck-step5)
      + [Scenario 3 — consensus never completes (step3)](#scenario-3-consensus-never-completes-step3)
   * [Metric reference](#metric-reference)
   * [Known gaps](#known-gaps)

# Runbook: uncommitted message

A message's onramp sequence number has been observed but no report covering it has landed on
the destination offramp within the expected window. This runbook walks source → destination
metrics to find where the pipeline is stuck.

## For agents

This doc is meant to be walked mechanically, not just read. Two things make that possible:

1. **[Decision graph](#decision-graph)** below is the authoritative control flow. Each step has
   a `query` (fill in `$sourceChain`/`$destChain`/`$seqNum` from the Inputs above) and a
   `condition`. Evaluate the query against your Prometheus-compatible datasource, evaluate the
   condition, and follow the matching outcome. The prose in [Steps](#steps) explains *why* each
   step exists — read it if you need rationale, skip it if you don't.
2. **Outcome vocabulary** is fixed and every step uses only these three:
   - `STOP` — investigation is resolved from commit-plugin metrics alone. Nothing further to check.
   - `CONTINUE:<step_id>` — proceed to the named step.
   - `REPORT:<owner>` — a plausible root cause was found, but it names a system, decision, or
     action outside this runbook's scope (another team, an on-chain tx, a config change). **Do
     not page, escalate, or act on `REPORT` yourself** — surface the finding, the evidence
     (the query + result that produced it), and the suggested `<owner>` to the human, and stop.

   Every check in this runbook is a read-only PromQL query (`automatable: true`). A few steps in
   the deep dives point at systems this runbook has no query for (chain explorers, TXM
   dashboards) — those are marked `automatable: false`; treat them the same as `REPORT`: name
   what a human should look at, don't try to fetch it yourself unless you have another tool for
   it.

If a query's metric doesn't exist on your datasource, treat the result as unknown, not as
`0`/flat — say so explicitly rather than silently treating "no such metric" as "no signal."

### Resolving inputs before you start

This runbook's frontmatter `inputs` are `sourceChain`, `destChain`, `seqNum`, `msgID` — but you're rarely
handed exactly those four. You might instead have only a message ID, or only a transaction hash
plus which chain it's on. Resolve the full set *before* starting the decision graph:

- **Have a message ID or a tx hash, missing the rest?** Use `ccip-cli show` — see
  [`README.md`'s message-identification section](README.md#identifying-a-ccip-message) for the
  exact invocation. Its JSON output gives you source chain, dest chain, sequence number, and
  message ID together, which is exactly this runbook's `inputs` block.
- **Chain given as a selector or chain ID, but a query below needs a name (or vice versa)?** See
  [`chain-identifiers.md`](chain-identifiers.md) — don't assume, translate explicitly, and keep
  all three forms handy since different labels in this doc want different ones (see the
  [metric reference](#metric-reference) for which).

### A note on counter naming

Every metric below is named as it's registered in `commit/metrics/prom.go`. Whether you need to
append `_total` depends on how metrics reach your datasource, **not** on the runbook — see the
[metric reference table](#metric-reference) for the per-metric OTel type. Beholder-only metrics
(anything not also `promauto`-registered) only ever leave the node as OTel instruments; if your
pipeline converts those to Prometheus via an OTel collector's `prometheusremotewrite` exporter
(e.g. this repo's `devenv` VictoriaMetrics stack, see `devenv/env.toml` /
`docs/metrics/commit-metrics.md`), every `Counter` gets a `_total` suffix and every `Histogram`
becomes `_bucket`/`_sum`/`_count` — `Gauge`s are untouched. If instead you're scraping a
`promauto`-registered metric directly (only the handful of pre-existing metrics that have both a
`promauto` and a Beholder registration — see the table), no suffix is added. Check one query
against your actual datasource's metric browser before trusting the rest; don't assume.

## Decision graph

```yaml
steps:
  - id: step0.1
    check: discovery_state
    query: 'max by (destChainID) (ccip_commit_discovery_state{destChainID=~"$destChain"})'
    condition: "result == 0"
    if_true:  {action: "REPORT:ccip-commit-oncall", reason: "the commit plugin has NOT discovered its contracts, so it is not healthy and cannot produce any observations -- this is the root cause, not a symptom: no downstream observation/outcome metric will read meaningfully while the plugin is stuck in discovery. Check the destination chain's RPC/reader health (ccip_reader_* / ccip_commit_consensus_observation_failed_total), since a single unhealthy RPC is what most commonly leaves the plugin perpetually in discovery. The discovery gauge is recorded even during discovery rounds, so value 0 here is a real signal, not an empty result", followup_query: 'sum by (destChainID) (rate(ccip_commit_consensus_observation_failed_total{destChainID=~"$destChain"}[5m]))'}
    if_false: {action: "CONTINUE:heartbeat"}
    automatable: true
    note: "gated first because it is the one check that answers 'is the plugin actually ready to produce observations' rather than merely 'is it scheduling rounds'. ccip_commit_discovery_state is a gauge set per-round (from commit/plugin.go's Observation, via TrackDiscoveryState/contractsInitialized.Load()): 1 = discovered, 0 = still in discovery. Use max() so only a DON-wide 'none of the nodes exited discovery' hits the == 0 branch; a single node stuck in discovery (by csa_public_key) is a node-local finding, not a stoppage."

  - id: step0.2
    check: plugin_heartbeat
    query: 'sum(rate(ccip_commit_plugin_heartbeat_total{destChainID=~"$destChain"}[1m]))'
    condition: "result == 0"
    if_true:  {action: "REPORT:ccip-commit-oncall", reason: "plugin process is not running rounds at all; every other metric below may simply be stale, not bad. Always run followup_query before reporting -- a bare 'not running' is a real but unsatisfyingly blunt finding on its own (confirmed by live testing: an agent that stopped here reported a correct but generic conclusion, then went and ran this exact followup_query anyway 'out of curiosity' because the plain heartbeat result didn't feel actionable)", followup_query: 'count(count by (csa_public_key) (timestamp(ccip_commit_plugin_heartbeat_total{destChainID=~"$destChain", phase="observation"}) > time() - 60))', followup_meaning: "live oracle headcount -- this is a DON-level signal, same trust tier as heartbeat itself, safe to check here even though lane-specific metrics below are not (yet) trustworthy. A low count turns 'plugin not running' into 'N of the DON's oracles are actually alive', which is what an on-call needs next regardless of whether it's below consensus threshold -- grade it against $fRoleDON per scenario3b's rule if known, otherwise report the raw number"}
    if_false: {action: "CONTINUE:step1"}
    automatable: true
    note: "window is deliberately 1m, not 5m -- confirmed by live testing (see scenario3b's note) that rate() extrapolates across its whole window, so a plugin that died N minutes ago still reads as alive for up to [window] more minutes. Don't widen this without re-testing against a forced failure."

  - id: step1
    check: offramp_next_seq_num
    query: 'max by (sourceChainName) (ccip_commit_offramp_next_seq_num{sourceChainName=~"$sourceChain", destChainID=~"$destChain"})'
    condition: "result > $seqNum"
    if_true:  {action: "STOP", reason: "a root covering X already landed on the offramp; commit plugin's job is done, hand off to exec on-call for delivery status"}
    if_false: {action: "CONTINUE:step2"}
    automatable: true

  - id: step2
    check: onramp_max_seq_num
    query: 'max by (sourceChainName) (ccip_commit_onramp_max_seq_num{sourceChainName=~"$sourceChain", destChainID=~"$destChain"})'
    condition: "result < $seqNum"
    if_true:  {action: "REPORT:chain-infra-oncall", reason: "onramp read is lagging; check ccip_commit_merkleroot_observation_errors_total by reason (no_bindings/timeout/rpc_error) for the source-chain read/infra issue", followup_query: 'sum(rate(ccip_commit_merkleroot_observation_errors_total{sourceChainName=~"$sourceChain", destChainID=~"$destChain"}[5m])) by (reason)'}
    if_false: {action: "CONTINUE:step2b"}
    automatable: true

  - id: step2b
    check: per_chain_consensus_dropped
    queries:
      - 'sum(rate(ccip_commit_consensus_dropped_total{sourceChainName=~"$sourceChain", destChainID=~"$destChain", objectName=~"MerkleRoot|OnRampMaxSeqNums"}[5m])) by (objectName, reason)'
      - 'sum(rate(ccip_commit_offramp_consensus_insufficient_total{sourceChainName=~"$sourceChain", destChainID=~"$destChain"}[5m]))'
    condition: "any result > 0"
    if_true:  {action: "REPORT:ccip-commit-oncall", reason: "DON could not reach consensus on a per-chain value for this lane. For ccip_commit_consensus_dropped: reason='split' means disagreeing oracles, reason='insufficient_agreement' means too few oracles observed the key, reason='threshold_not_defined' is a config/data mismatch. ccip_commit_offramp_consensus_insufficient means the DON could not agree on OffRampNextSeqNums for this lane. The message cannot advance until the disagreement resolves."}
    if_false: {action: "CONTINUE:step2c"}
    automatable: true

  - id: step2c
    check: reader_data_source
    queries:
      - 'sum by (query, chainID) (rate(ccip_reader_read_empty_total{query=~"NextSeqNum|MsgsBetweenSeqNums|LatestMsgSeqNum", destChainID=~"$destChain"}[5m]))'
      - 'sum by (query, chainID, state) (rate(ccip_reader_chain_gap_total{query=~"NextSeqNum|MsgsBetweenSeqNums|GetChainsFeeComponents|GetChainFeePriceUpdate", state!="returned", destChainID=~"$destChain"}[5m]))'
      - 'max by (chainID, kind) (ccip_reader_config_cache_age_seconds{destChainID=~"$destChain"})'
    condition: "any read_empty series > 0, any chain_gap series with state!=\"returned\" > 0, or any config_cache_age_seconds{kind=\"chain\"} > 90 (3x refresh period = sustained refresh failure, not a single empty poll)"
    if_true:  {action: "REPORT:chain-infra-oncall", reason: "the plugin DATA SOURCE is degrading, which is upstream of -- and distinct from -- onramp-lag/consensus/transmission. Reads are returning empty (nothing) or partial (chains not returned, or a message read whose count doesn't match its range) or the config-poller cache went stale. This is the root cause the onramp-gauge stall looks like, not a commit-plugin bug. For ccip_reader_chain_gap, the actionable state values per (query,chainID) are not_found|disabled|misconfigured|error|invalid|missing|stale|count_mismatch. Split node-local (one csa_public_key) from DON-wide (all series) before escalating.", followup_query: 'count by (csa_public_key) (ccip_reader_chain_gap_total{query="NextSeqNum", destChainID=~"$destChain"})'}
    if_false: {action: "CONTINUE:step3"}
    automatable: true
    note: "ccip_reader_* / config-poller metrics are shared with execute and keyed by chainID/chainFamily/chainName/chainSelector (from pkg/reader/rcmetrics's chainAttrs() helper -- replaced the old bare numeric-selector `chain` label; chainSelector carries that same numeric value if you need it) and by `query`. They also ALL carry destChainID/destChainFamily/destChainName/destChainSelector (a second, dest-scoped chainAttrs() resolved per config-poller/reader instance) -- destChainID=~\"$destChain\" is REQUIRED, not optional: a node runs one config-poller/reader instance per destination chain, and more than one instance can independently track the same source chain (confirmed on staging: two destination-chain lanes both reading Solana collided into one series per (chainID, node_id) without this filter, silently mixing a healthy lane's fast-refreshing cache with a different, broken lane's stuck one). Every series carries node_id + csa_public_key inherited from the beholder client, so group by them to distinguish a single bad node from a DON-wide data/index/RPC failure. See docs/metrics/reader-metrics.md."

  - id: step3
    check: consensus_observation_failed
    query: 'sum(rate(ccip_commit_consensus_observation_failed_total{destChainID=~"$destChain"}[5m]))'
    condition: "result > 0"
    if_true:  {action: "CONTINUE:scenario3", reason: "destination-chain-wide consensus failure — see deep dive"}
    if_false: {action: "CONTINUE:step4"}
    automatable: true

  - id: step4
    check: cursing
    queries:
      - 'max(ccip_commit_rmn_curse_active{destChainID=~"$destChain", curse_type="global"})'
      - 'max(ccip_commit_rmn_curse_active{destChainID=~"$destChain", curse_type="destination"})'
      - 'max(ccip_commit_source_chain_cursed{sourceChainName=~"$sourceChain", destChainID=~"$destChain"})'
    condition: "any result == 1"
    if_true:  {action: "REPORT:curse-owner", reason: "chain A or B is cursed; likely intentional/incident-flagged, hand off to whoever owns the curse. NOTE: curse state is served from the config-poller cache (GetRmnCurseInfo) -- if ccip_reader_config_cache_age_seconds{chain=<destSelector>,kind=\"chain\"} is high, the curse reading itself may be stale (a lifted curse still reporting active, or a fresh one unobserved); say so when reporting", followup_query: 'ccip_reader_config_cache_age_seconds{kind="chain"}'}
    if_false: {action: "CONTINUE:step5"}
    automatable: true

  - id: step5
    check: report_transmission
    query: 'sum(rate(ccip_commit_report_transmission_gave_up_total{sourceChainName=~"$sourceChain", destChainID=~"$destChain"}[5m]))'
    condition: "result > 0"
    if_true:  {action: "CONTINUE:scenario2", reason: "report is being built but never landing — see deep dive"}
    if_false: {action: "REPORT:ccip-commit-oncall", reason: "no failure signal found; likely just backlog size — report ccip_commit_pending_messages and an ETA estimate from round cadence", followup_query: 'max by (sourceChainName) (ccip_commit_pending_messages{sourceChainName=~"$sourceChain", destChainID=~"$destChain"})'}
    automatable: true

  - id: scenario2
    check: report_validation_rejected
    query: 'sum(rate(ccip_commit_report_validation_rejected_total{destChainID=~"$destChain", phase="should_transmit"}[15m])) by (reason)'
    condition: "any reason count > 0"
    if_true:
      config_digest_mismatch:  {action: "REPORT:ccip-commit-oncall", reason: "config sync issue; cross-check ccip_commit_config_digest_mismatch. Before trusting this, verify the digest is not coming from a stale config-poller cache -- ccip_reader_config_cache_age_seconds{kind=\"chain\"} high or ccip_reader_config_poller_last_success_timestamp old would make a mismatch read manufactured by stale data, i.e. a chain-reader issue, not a config-sync one"}
      config_digest_check_error: {action: "REPORT:chain-infra-oncall", reason: "error *reading* the offramp's config digest (RPC/read failure), distinct from an actual mismatch -- chain-B read/infra issue, not a config sync issue"}
      stale:                   {action: "STOP", reason: "often benign — a different overlapping report already landed; re-check step1, the gauge may have just moved"}
      dest_not_supported:      {action: "REPORT:ccip-commit-oncall", reason: "transmission-schedule/config bug"}
      dest_support_check_error: {action: "REPORT:chain-infra-oncall", reason: "error *checking* dest-chain support (RPC/read failure), distinct from a real config gap"}
      cursed:                  {action: "CONTINUE:step4", reason: "converges with step4"}
      cursed_check_error:      {action: "REPORT:chain-infra-oncall", reason: "error *checking* the curse state (RPC/read failure), not an actual curse -- don't conflate with step4's cursed=1 case"}
      empty_root:              {action: "REPORT:ccip-commit-oncall", reason: "report-builder produced an empty merkle root. Before treating this as a merkleroot-processor bug, check whether the reader supplied incomplete data for this lane: ccip_reader_read_empty_total{query=\"MsgsBetweenSeqNums\", destChainID=~\"$destChain\"} or ccip_reader_chain_gap_total{query=\"MsgsBetweenSeqNums\", state!=\"returned\", destChainID=~\"$destChain\"} firing for this chain around the same time means the reader handed the processor a subset (or nothing) to build from -- a reader-layer/chain-infra issue, not a processor bug. Both flat → likely genuinely a bug upstream in the merkleroot processor.", followup_query: 'sum by (chainID, state) (rate(ccip_reader_chain_gap_total{query=\"MsgsBetweenSeqNums\", state!=\"returned\", destChainID=~\"$destChain\"}[15m]))'}
      invalid_seqnum_range:    {action: "REPORT:ccip-commit-oncall", reason: "report-builder produced start>end seqnums. Same reader-layer check as empty_root above: ccip_reader_chain_gap_total{query=\"MsgsBetweenSeqNums\", state=\"count_mismatch\", destChainID=~\"$destChain\"} in particular means the reader returned a message count that didn't match its requested range -- a partial read manufacturing an invalid range, not a processor bug. Rule that out before reporting this as a merkleroot-processor defect.", followup_query: 'sum by (chainID) (rate(ccip_reader_chain_gap_total{query=\"MsgsBetweenSeqNums\", state=\"count_mismatch\", destChainID=~\"$destChain\"}[15m]))'}
      decode_report:           {action: "REPORT:ccip-commit-oncall", reason: "report codec failure; check for a report-codec/chainlink-common version skew across the DON"}
      decode_report_info:      {action: "REPORT:ccip-commit-oncall", reason: "same as decode_report but for the report-info envelope"}
      root_blessing_mismatch:  {action: "REPORT:ccip-commit-oncall", reason: "root blessing validation failed; check on-chain blessing state for the roots in this report. RMN signing itself is dead code (RMNEnabled hardcoded off) but this specific check still runs"}
      default:                 {action: "REPORT:ccip-commit-oncall", reason: "reason string not in this list -- do NOT go source-diving to interpret it as one of the reasons above by guessing. Report the exact reason string, the rate, and say explicitly this runbook doesn't have a mapping for it yet. (This is exactly how a hardcoded test-only fault injection with reason=\"forced_test_failure\" surfaces if a stale binary is still deployed -- treat an unmapped reason as equally likely to be a real new rejection path OR leftover test/debug code, and say you can't tell which from metrics alone.)"}
    if_false: {action: "CONTINUE:scenario2b", reason: "never rejected locally — check on-chain / read-lag hypotheses"}
    automatable: true

  - id: scenario2b
    check: transmitted_but_stuck_or_read_lag
    hypotheses:
      - name: "reverted/stuck on-chain"
        action: "REPORT:chain-b-txm-oncall"
        reason: "outside commit-plugin metrics — check chain B's TXM/tx-sender dashboard for the assigned transmitter (nonce gap, underpriced gas, wallet balance, revert reason)"
        automatable: false
      - name: "actually succeeded, read is lagging"
        action: "REPORT:ccip-commit-oncall"
        reason: "ccip_commit_offramp_next_seq_num is sourced from this plugin's own chain-B reader; cross-check chain B's block explorer / chain-reader health directly before escalating further"
        automatable: false

  - id: scenario3
    check: destination_wide_vs_lane_specific
    query: 'max by (sourceChainName) (ccip_commit_pending_messages{destChainID=~"$destChain"})'
    condition: "only $sourceChain's series is climbing, others into the same destChain are flat"
    if_true:  {action: "CONTINUE:step4", reason: "this branch is wrong if only one lane is affected — back to step4/step5"}
    if_false: {action: "CONTINUE:scenario3b", reason: "confirmed destination-chain-wide"}
    automatable: true

  - id: scenario3b
    check: h0_live_oracle_count
    query: 'count(count by (csa_public_key) (timestamp(ccip_commit_plugin_heartbeat_total{destChainID=~"$destChain", phase="observation"}) > time() - 60))'
    condition: "compare result against 2*fRoleDON+1 (consensus threshold) and 3*fRoleDON+1 (full DON), if fRoleDON was supplied as an input"
    if_true:
      below_threshold:   {action: "REPORT:ccip-commit-oncall", reason: "H0 — result < 2*fRoleDON+1: too few oracles are even running to reach consensus, full stop. This is the answer; don't proceed to H1/H2, they're checks for a *different* failure mode (an oracle that's running but can't read its home chain, or oracles that disagree) and will likely read as flat/inconclusive here even though they're not the cause."}
      degraded_but_over_threshold: {action: "CONTINUE:scenario3c", reason: "some oracles down but still >= 2*fRoleDON+1; consensus should theoretically still be possible, so if it's failing anyway H1/H2 are still the right next check"}
      full_don:          {action: "CONTINUE:scenario3c", reason: "oracle count isn't the explanation; proceed to H1/H2"}
    if_false: {action: "CONTINUE:scenario3c", reason: "fRoleDON wasn't supplied — report the raw count as context (see note) and proceed to H1/H2 regardless; a human/agent should sanity-check the count against what they know the DON size to be before trusting H1/H2's conclusion"}
    automatable: true
    note: "csa_public_key uniquely identifies each node in every commit metric series (confirmed empirically, not just in the source) -- this is a direct headcount, not a proxy. Deliberately uses timestamp()-vs-time(), not rate()>0: confirmed by live testing that rate([5m])>0 stays true for up to 5 minutes after a node actually stops (rate() extrapolates across its whole window using the oldest/newest samples it finds), which silently under-counts a fresh outage. timestamp() answers 'did this series get a sample in the last 60s' directly, no extrapolation lag. Distinguishes 'not enough oracles are online' (a category H1/H2 cannot detect at all, since fchain_read_errors only reports a *surviving* oracle's own read failures and has no way to see oracles that never started) from the H1/H2 failure modes below. This is now closed at the source by ccip_commit_consensus_dropped{reason=\"insufficient_agreement\"}; the heartbeat headcount is still useful as cross-check."

  - id: scenario3c
    check: h1_vs_h2
    queries:
      - 'sum(rate(ccip_commit_fchain_read_errors_total{destChainID=~"$destChain"}[5m]))'
      - 'sum by (query, chainID, csa_public_key) (rate(ccip_reader_read_empty_total{destChainID=~"$destChain"}[5m]))'
      - 'sum by (query, chainID, state, csa_public_key) (rate(ccip_reader_chain_gap_total{state!="returned", destChainID=~"$destChain"}[5m]))'
    condition: "fchain_read_errors spiking broadly across oracles, OR reader read_empty/non-returned chain_gap firing across most/all csa_public_key values for the same query"
    if_true:  {action: "REPORT:home-chain-infra-oncall", reason: "H1 — too few oracles could read FChain; home-chain RPC/read outage. The reader-layer series (grouped by csa_public_key) give the direct data-level observation behind this: if read_empty/chain_gap fire for the SAME query across (nearly) every csa_public_key, that's DON-wide data-source failure, not one node's RPC -- report it as such rather than as a bare fchain_read_errors count. If only one or two csa_public_key values show up, that's a single bad node, not a DON-wide outage; say so explicitly, it changes the remediation."}
    if_false: {action: "REPORT:ccip-commit-oncall", reason: "H2 — oracles likely disagree (split vote). Confirm with ccip_commit_consensus_dropped{objectName=\"fChain\", reason=\"split\"} and corroborate with ccip_commit_config_digest_mismatch flipping for a subset of oracles around the same time. If the reader-layer series above are flat (no read_empty/chain_gap), that corroborates H2 -- oracles are reading fine, they're just disagreeing on the value. If scenario3b (H0) wasn't able to rule out insufficient participation (fRoleDON unknown), treat this H2 conclusion as low-confidence, not a firm diagnosis.", followup_query: 'ccip_commit_config_digest_mismatch{destChainID=~\"$destChain\"}'}
    automatable: true
    note: "the reader-layer queries are shared with execute and keyed by chainID/chainFamily/chainName/chainSelector -- a different package's chainAttrs() helper from the commit-plugin's sourceChainName/destChainID, but the reader layer ALSO carries its own destChainID/destChainFamily/destChainName/destChainSelector now (a second, dest-scoped chainAttrs() resolved per reader instance) -- required here for the same reason it's required on step2c: a node runs one reader instance per destination chain, and more than one instance can independently observe the same source chain. See docs/metrics/reader-metrics.md and the data-layer note on step2c. Grouping by csa_public_key here is what turns 'H1 or H2' from a guess into a query: a bad read on every node's series is DON-wide; a bad read on one node's series is that node's RPC, regardless of what fchain_read_errors alone says."
```

## Steps

### step0.1 — has the plugin discovered its contracts?

Check this before anything else, even the heartbeat: a plugin that has *not* discovered its
contracts is not healthy and cannot produce any observations, regardless of how many rounds it
keeps scheduling. Every observation/outcome metric below reports "no signal" identically whether
the plugin is healthy-and-idle, still in discovery, or wedged/crashed — discovery state is the
one direct signal that separates "still starting / can't even initialise" from the rest.

```promql
max by (destChainID) (ccip_commit_discovery_state{destChainID=~"$destChain"})
```
`0` → the plugin has not discovered its contracts. This is the root cause, not a symptom: report
it (alongside the destination chain's RPC/reader health, since a single unhealthy RPC is what most
commonly leaves the plugin stuck in discovery). `1` (or any `> 0`) → discovery complete, continue to
the heartbeat gate. The gauge is recorded every round (including during discovery, from
`commit/plugin.go`'s `Observation` via `TrackDiscoveryState`/`contractsInitialized.Load()`), so a
`0` here is a real finding, not a benign empty result. Use `max()`: only a DON-wide "none of the
nodes exited discovery" trips the `0` branch; one node stuck in discovery (group by `csa_public_key`)
is a node-local finding, not a stoppage.

### step0.2 - heartbeat — is the plugin alive at all?

Not in the original design doc's walkthrough, but should be checked first: every metric below
reports "no signal" identically whether the plugin is healthy-and-idle or wedged/crashed. Rule
that out before reading anything else as a bad value. (Discovery state from step0 already told you
the plugin is *ready*; this tells you it is *still scheduling rounds*.)

```promql
sum(rate(ccip_commit_plugin_heartbeat_total{destChainID=~"$destChain"}[1m]))
```
`0` → before reporting, also run the oracle headcount (same query as scenario3b's H0 —
it's a DON-level signal, safe to check even though lane-specific metrics aren't trustworthy yet):
```promql
count(count by (csa_public_key) (timestamp(ccip_commit_plugin_heartbeat_total{destChainID=~"$destChain", phase="observation"}) > time() - 60))
```
Report both numbers together. "Plugin not running rounds" on its own is a real but blunt
finding — confirmed by live testing that it's unsatisfying enough that whoever's investigating
will go run this exact query anyway; do it here instead of making them rediscover it. `> 0` →
continue to step1.

### step1 — is this even a commit-plugin problem?

```promql
max by (sourceChainName) (ccip_commit_offramp_next_seq_num{sourceChainName=~"$sourceChain", destChainID=~"$destChain"})
```
`> $seqNum` → a root covering it already landed on the offramp. Commit plugin's job is done;
hand off to exec on-call. `<= $seqNum` → continue. `destChainID=~"$destChain"` is required, not
optional, here and on every per-lane query below it — without it you're reading this lane across
*every* destination chain that reports into the same datasource, not just the one you're
investigating (see the metric reference's note on `destChainAttrs`).

### step2 — has the plugin observed it on-chain yet?

```promql
max by (sourceChainName) (ccip_commit_onramp_max_seq_num{sourceChainName=~"$sourceChain", destChainID=~"$destChain"})
```
`< $seqNum` → onramp read is lagging. Check the reason breakdown:
```promql
sum(rate(ccip_commit_merkleroot_observation_errors_total{sourceChainName=~"$sourceChain", destChainID=~"$destChain"}[5m])) by (reason)
```
(`no_bindings` / `timeout` / `rpc_error` / `msg_count_mismatch` / `hash_error` /
`address_lookup_error`). This is a source-chain read/infra issue, not commit logic — report and
stop. `>= $seqNum` → plugin has seen it; continue.

### step2b — can the DON agree on this chain's per-chain values?

The onramp has been observed, but the DON still has to agree on the source-chain values that go
into the merkle root (onramp max seq nums and merkle roots themselves). If consensus on these
fails for this lane, the message will not advance even though the onramp read is healthy.

Per-chain consensus failures:
```promql
sum(rate(ccip_commit_consensus_dropped_total{sourceChainName=~"$sourceChain", destChainID=~"$destChain", objectName=~"MerkleRoot|OnRampMaxSeqNums"}[5m])) by (objectName, reason)
```

`ccip_commit_consensus_dropped` `reason` values:
- `split` — two or more distinct values each cleared the threshold (oracles disagree)
- `insufficient_agreement` — no single value reached the threshold (too few oracles observed the key)
- `threshold_not_defined` — the consensus code had no threshold configured for this key (config/data mismatch)

OffRampNextSeqNums consensus failures use a dedicated counter (same path conceptually, but a
custom consensus helper):
```promql
sum(rate(ccip_commit_offramp_consensus_insufficient_total{sourceChainName=~"$sourceChain", destChainID=~"$destChain"}[5m]))
```

Any of these `> 0` for the lane → the message cannot advance until the disagreement resolves.
Report and stop. All flat (including empty) → continue.

### step2c — is the problem upstream in the data source instead?

The on/offramp gauges and consensus signals are *outcomes* — they can't tell "plugin is broken" from
"plugin is fine but being fed empty/partial data." This step checks the data source directly:

```promql
sum by (query, chainID) (rate(ccip_reader_read_empty_total{query=~"NextSeqNum|MsgsBetweenSeqNums|LatestMsgSeqNum", destChainID=~"$destChain"}[5m]))
sum by (query, chainID, state) (rate(ccip_reader_chain_gap_total{query=~"NextSeqNum|MsgsBetweenSeqNums|GetChainsFeeComponents|GetChainFeePriceUpdate", state!="returned", destChainID=~"$destChain"}[5m]))
max by (chainID, kind) (ccip_reader_config_cache_age_seconds{destChainID=~"$destChain"})
```

Any `read_empty` or non-`returned` `chain_gap` `> 0` (or any `config_cache_age_seconds{kind="chain"}` above ~90s — sustained, 3x refresh; a
single intermittent empty snap isn't enough) → the **data layer** feeding the
commit plugin is degrading, so this is a `chain-infra`/reader problem, not a commit-plugin defect. Still
split single-node (group by `csa_public_key`) vs DON-wide before escalating. All flat → continue.
`chainID` here (plus `chainFamily`/`chainName`/`chainSelector`, from `pkg/reader/rcmetrics`'s
`chainAttrs()` helper) is whichever chain that specific read/query was about — replacing the old
bare numeric-selector `chain` label; `chainSelector` still carries that same numeric value as a
string if you need it. `destChainID=~"$destChain"` (also on every series, from a second dest-scoped
`chainAttrs()` resolved per reader/config-poller instance) is REQUIRED here, not optional: a node
runs one reader/config-poller instance per destination chain, and more than one instance can
independently track the same source chain — confirmed on staging, omitting this filter let two
destination-chain lanes sharing a source chain collide into one series per `(chainID, node_id)`.
`query` names the read. These `ccip_reader_*`/config-poller metrics are shared with execute — see
`docs/metrics/reader-metrics.md`.

> This is the data-layer *gate*: it belongs before the consensus/transmission tree (steps 3-5), which is
> exactly where a reader that silently returns empty/partial would be misattributed to the plugin.

### step3 — is the round producing outcomes at all?

```promql
sum(rate(ccip_commit_consensus_observation_failed_total{destChainID=~"$destChain"}[5m]))
```
Incrementing → destination-chain-wide consensus failure, jump to [Scenario 3](#scenario-3--consensus-never-completes-step3).
Flat → continue.

### step4 — is chain A or B cursed?

```promql
max(ccip_commit_rmn_curse_active{destChainID=~"$destChain", curse_type="global"})
max(ccip_commit_rmn_curse_active{destChainID=~"$destChain", curse_type="destination"})
max(ccip_commit_source_chain_cursed{sourceChainName=~"$sourceChain", destChainID=~"$destChain"})
```
Any `== 1` → found it; likely intentional/incident-flagged, report to whoever owns the curse and
stop. All `0` → continue.

> RMN-signature branches (signed-roots-dropped, chain quorum, Byzantine root disagreement) are
> intentionally absent from this runbook: `RMNEnabled` is hardcoded off in production, so those
> paths are dead code. Cursing (this step) is the one RMN-adjacent signal still live.

### step5 — is a report being built but never landing?

```promql
sum(rate(ccip_commit_report_transmission_gave_up_total{sourceChainName=~"$sourceChain", destChainID=~"$destChain"}[5m]))
```
Incrementing/maxed → jump to [Scenario 2](#scenario-2--report-transmission-stuck-step5).
Otherwise → likely just backlog size:
```promql
max by (sourceChainName) (ccip_commit_pending_messages{sourceChainName=~"$sourceChain", destChainID=~"$destChain"})
```
Report the backlog size and an ETA estimate from round cadence, and stop.

## Deep dives

### Scenario 2 — report transmission stuck (step5)

Three hypotheses, cheapest first.

**A — never transmitted (rejected pre-send).**
```promql
sum(rate(ccip_commit_report_validation_rejected_total{destChainID=~"$destChain", phase="should_transmit"}[15m])) by (reason)
```
- `config_digest_mismatch` climbing → cross-check `ccip_commit_config_digest_mismatch`; config
  sync issue.
- `config_digest_check_error` → an RPC/read error *checking* the digest, not an actual mismatch —
  chain-B infra issue, don't conflate with the reason above.
- `stale` climbing → often benign — a different, overlapping report already landed; re-check
  step1, it may have just moved.
- `dest_not_supported` → transmission-schedule/config bug.
- `dest_support_check_error` → RPC/read error checking dest support, not a real config gap.
- `cursed` → converges with step4.
- `cursed_check_error` → RPC/read error checking curse state, not an actual curse — don't
  conflate with step4's `cursed=1`.
- `empty_root` / `invalid_seqnum_range` → the report-builder produced an invalid report. Before
  attributing this to the merkleroot processor, rule out a reader-supplied partial read for this
  lane:
  ```promql
  sum by (chainID, state) (rate(ccip_reader_chain_gap_total{query="MsgsBetweenSeqNums", state!="returned", destChainID=~"$destChain"}[15m]))
  sum by (chainID) (rate(ccip_reader_read_empty_total{query="MsgsBetweenSeqNums", destChainID=~"$destChain"}[15m]))
  ```
  `chain_gap{state="count_mismatch"}` in particular means the reader returned a message count
  that didn't match its requested range — that alone can manufacture `invalid_seqnum_range`
  without any bug in the processor. Any of these firing for the lane around the same time → this
  is a reader-layer/chain-infra issue, not a processor bug. All flat → likely a genuine bug
  upstream in the merkleroot processor.
- `decode_report` / `decode_report_info` → report/report-info codec failure; check for a
  report-codec or `chainlink-common` version skew across the DON.
- `root_blessing_mismatch` → on-chain root blessing validation failed for this report. RMN
  signing itself is dead code (`RMNEnabled` hardcoded off) but this specific check still runs.
- **Any other reason string** → do not guess which of the above it's "close enough to." Report
  the exact string and rate and say plainly this runbook has no mapping for it — confirmed by
  live testing that an unmapped reason (`forced_test_failure`, deliberately not a real production
  value) forces exactly this situation. An unmapped reason is equally likely to be a genuinely
  new rejection path in the code or leftover test/debug instrumentation in a stale binary; metrics
  alone can't tell you which, and the doc shouldn't pretend otherwise.
- All flat → move to B.

**B — transmitted, reverted/stuck on-chain.** *(`automatable: false` — outside commit-plugin
metrics.)* Check chain B's TXM/tx-sender dashboards for the assigned transmitter (nonce gap,
underpriced gas, wallet balance, revert reason).

**C — actually succeeded, read is lagging.** *(`automatable: false`.)*
`ccip_commit_offramp_next_seq_num` is sourced from this plugin's own chain-B reader; if that's
behind, "no change" can persist for a few rounds after a real success. Cross-check chain B's
block explorer / chain-reader health directly before escalating further.

### Scenario 3 — consensus never completes (step3)

Grounded in `getConsensusObservation` (`merkleroot/outcome.go`): the *only* way the whole round
fails is the DON not reaching 2·fRoleDON+1 agreement on a single FChain value for the
**destination chain** — every other field degrades gracefully per-chain.

- **This is destination-chain-wide, not lane-specific.** Before going further, check whether
  *other* source chains reporting into this dest chain are also stalled:
  ```promql
  max by (sourceChainName) (ccip_commit_pending_messages{destChainID=~"$destChain"})
  ```
  (Deliberately no `sourceChainName` filter here — the whole point is comparing *every* lane
  into `$destChain` side by side, not just `$sourceChain`'s.) If only `$sourceChain` is affected,
  this branch is wrong — back to step4/step5.
- **H0 — not enough oracles are even running.** *(Check this before H1/H2 — it's a distinct
  failure mode neither of them can detect, not a variant of either.)* `fchain_read_errors` below
  only reports a *surviving* oracle's own read failures; it has no way to see an oracle that
  never started at all. Count oracles that got an Observation-phase heartbeat sample in the last
  60s — `csa_public_key` uniquely identifies each node in every commit metric series. Uses
  `timestamp()`, not `rate()>0`: `rate()` extrapolates across its whole window, so it stays
  `> 0` for up to that window's length after a node actually dies — confirmed by live testing,
  this genuinely produced a stale headcount at `[5m]` that only corrected itself once enough
  time had passed:
  ```promql
  count(count by (csa_public_key) (timestamp(ccip_commit_plugin_heartbeat_total{destChainID=~"$destChain", phase="observation"}) > time() - 60))
  ```
  If `$fRoleDON` is known: `< 2*$fRoleDON+1` → this **is** the answer, stop here — don't let a
  flat H1 push you toward H2 below, H1/H2 are checks for a different failure mode and will likely
  look inconclusive here even though they're not the cause. `>= 2*$fRoleDON+1` → oracle count
  isn't the explanation, continue. If `$fRoleDON` is unknown, report the raw count as context and
  continue regardless — but treat whatever H1/H2 concludes below as lower-confidence, since you
  haven't ruled this out.
- **H1 — too few oracles could read it.**
  ```promql
  sum(rate(ccip_commit_fchain_read_errors_total{destChainID=~"$destChain"}[5m]))
  ```
  Spiking broadly across oracles → home-chain RPC/read outage. Corroborate with the reader layer,
  grouped by node so "DON-wide" vs "one bad RPC" is a query, not a guess:
  ```promql
  sum by (query, chainID, csa_public_key) (rate(ccip_reader_read_empty_total{destChainID=~"$destChain"}[5m]))
  sum by (query, chainID, state, csa_public_key) (rate(ccip_reader_chain_gap_total{state!="returned", destChainID=~"$destChain"}[5m]))
  ```
  Firing across (nearly) every `csa_public_key` for the same query → DON-wide data-source failure,
  not a split vote. Firing for one or two `csa_public_key` values only → that node's RPC, not a
  DON-wide condition — say so explicitly, it changes who fixes it. Both flat → corroborates H2
  below (oracles are reading fine, they're just disagreeing).
- **H2 — oracles disagree (split vote).** Check `ccip_commit_consensus_dropped{objectName="fChain", reason="split"}`:
  ```promql
  sum(rate(ccip_commit_consensus_dropped_total{destChainID=~"$destChain", objectName="fChain", reason="split"}[5m])) by (key)
  ```
  This is genuine disagreement among oracles that DID submit a value — not the same as H0's "too few submitted at all." Corroborate with `ccip_commit_config_digest_mismatch` flipping for a subset of oracles around the same time:
  ```promql
  ccip_commit_config_digest_mismatch{destChainID=~"$destChain"}
  ```

## Metric reference

`otel_type` is what governs whether a `_total`/`_bucket` suffix gets added by an OTel→Prometheus
pipeline — see [the naming note above](#a-note-on-counter-naming). `dual` means the metric is
*also* directly `promauto`-registered (scrapeable without going through Beholder/OTel at all);
everything else is Beholder-only.

Dest labels = `destChainID,destChainFamily,destChainName,destChainSelector` (from `destChainAttrs()`).
Source labels = `sourceChainID,sourceChainFamily,sourceChainName,sourceChainSelector` (from
`sourceChainAttrs()`). Every Beholder-emitted commit metric below is built purely from these two
helpers plus its own extra dimensions (`phase`, `reason`, `status`, `type`, `success`, ...) — there
are no bare/ad-hoc chain labels left anywhere on the Beholder side.

| metric | otel_type | key labels (Beholder) | key labels (promauto, where `dual`) | registration |
|---|---|---|---|---|
| `ccip_commit_plugin_heartbeat` | Counter | dest + `phase` | — | beholder-only |
| `ccip_commit_discovery_state` | Gauge | dest + `discovered` (bool; the 0/1 gauge value mirrors it) | — | beholder-only |
| `ccip_commit_report_validation_rejected` | Counter | dest + `phase,reason` | — | beholder-only |
| `ccip_commit_onramp_max_seq_num` | Gauge | source + dest | — | beholder-only |
| `ccip_commit_offramp_next_seq_num` | Gauge | source + dest | — | beholder-only |
| `ccip_commit_pending_messages` | Gauge | source + dest | — | beholder-only |
| `ccip_commit_rmn_curse_active` | Gauge | dest + `curse_type` (no source labels — global/destination curse, not lane-specific) | — | beholder-only |
| `ccip_commit_source_chain_cursed` | Gauge | source + dest | — | beholder-only |
| `ccip_commit_merkleroot_observation_errors` | Counter | source + dest + `reason` | — | beholder-only |
| `ccip_commit_fchain_read_errors` | Counter | dest only (FChain is a destination-chain-wide value, no source lane) | — | beholder-only |
| `ccip_commit_consensus_observation_failed` | Counter | dest only | — | beholder-only |
| `ccip_commit_report_transmission_gave_up` | Counter | source + dest (see note below) | — | beholder-only |
| `ccip_commit_report_transmission_attempts` | Histogram | dest + `success` (no source — one report, not one lane) | — | beholder-only |
| `ccip_commit_range_truncated` | Counter | source + dest | — | beholder-only |
| `ccip_commit_offramp_lane_status` | Gauge | source + dest + `status` | — | beholder-only |
| `ccip_commit_seqnum_invariant_violation` | Counter | source + dest + `type` | — | beholder-only |
| `ccip_commit_offramp_consensus_insufficient` | Counter | source + dest | — | beholder-only |
| `ccip_commit_config_digest_mismatch` | Gauge | dest only | `chain_family,chain_id` (old scheme, unchanged) | dual |
| `ccip_commit_max_sequence_number` | Gauge | source + dest + `method` | `chainFamily,chainID,sourceChainFamily,sourceChain,method,source_network_name,dest_network_name` (old scheme, unchanged) | dual |
| `ccip_commit_latest_round_id` | Gauge | source + dest + `onrampAddress,method` | `source_network_name,dest_network_name,contract_address,plugin` (old scheme, unchanged — and note the promauto side's 4th value is actually `method`, not a real plugin name; pre-existing, not part of this rename) | dual |
| `ccip_commit_processor_errors` | Counter | **none — see caveat below** | `chainFamily,chainID,processor,method` (old scheme, unchanged) | labeled `dual`, **behaves promauto-only** |
| `ccip_commit_processor_latency` | Histogram | dest + `processor,method` | `chainFamily,chainID,processor,method` (old scheme, unchanged) | dual |
| `ccip_commit_loopp_ccip_provider_supported` | Gauge | dest + `loopChainFamily` (the LOOPP provider's own family, distinct from `destChainFamily`) | `chain_family` (old scheme, unchanged — this is the LOOPP provider family too, not `p.chainFamily`) | dual |
| `ccip_commit_consensus_dropped` | Counter | dest + `objectName,key,reason`, **plus source labels only when the drop is lane-scoped** (`objectName` ∈ `MerkleRoot`/`OnRampMaxSeqNums`); `objectName="fChain"` drops are dest-only, no source labels at all | — | beholder-only |

**Caveat — `ccip_commit_processor_errors` is effectively promauto-only in practice.** It's
registered as a Beholder `Int64Counter` (`bhProcessorErrors`) in `NewPromReporter`, but
`TrackProcessorLatency`'s error branch (`commit/metrics/prom.go`) only ever calls
`p.processorErrors.WithLabelValues(...).Inc()` (the promauto counter) — `p.bhProcessorErrors` is
never invoked anywhere in the file. Don't rely on this metric being visible via Beholder/OTel
(e.g. for a mainnet NOP node that isn't Prometheus-scraped); it's only reliably available on the
`promauto`/scrape path, with the old `chainFamily,chainID` labels, despite being marked `dual`
here for consistency with the table's own registration column. This should probably either get a
real `bhProcessorErrors.Add(...)` call or be re-labeled `promauto-only` in a follow-up.

**Per-lane metrics carry both source and dest labels; destination-scoped metrics carry only
dest.** A metric is per-lane (needs `sourceChainAttrs()` too) exactly when the plugin can observe
it separately per source chain; it's destination-scoped (dest labels only) when the value is a
single fact about the destination-chain round as a whole (heartbeat, FChain reads, consensus
round outcome, report-transmission attempts/latency, the digest/processor/LOOPP-provider health
gauges). Filtering a per-lane metric by `sourceChainName` alone without also filtering
`destChainID=~"$destChain"` will silently mix every destination chain sharing this datasource
that reads the same source chain — always filter both.

`ccip_commit_report_transmission_gave_up` carries *both* full source and dest label sets because
a single report can cover multiple source chains, and each one's lane can resolve independently
before the overall check-attempt budget is exhausted (`pendingSources` in
`merkleroot/outcome.go`) — the source labels say *which lane* is still pending, the dest labels
say *which destination chain's* report gave up on it. Neither set makes the other redundant.

The old bare labels this scheme replaced (`chainID`, `chain_id`, `chainFamily`, `chain_family`,
`source_network_name`, `dest_network_name` on the Beholder side) are **gone, not aliased** — this
was a full rename, not an additive alias. A saved query against any of those old names on a
Beholder-emitted commit metric will now silently return empty, not error; the `promauto`/legacy
path in `commit/metrics/legacy_prom.go` is the one place those old names are still genuinely live
(see its `// TODO: consider removing these metrics entirely in a follow-up` — it's slated for
removal, so don't build new tooling against it).

### Reader / config-poller metrics (shared with execute)

These are the data-source signals used by step2c/step4/scenario2 above, defined in the reader and
config-poller layers rather than the commit plugin. Every series is **beholder-only** and carries
`node_id` + `csa_public_key` **inherited** from the beholder client, plus `chainID`, `chainFamily`,
`chainName`, `chainSelector` from `pkg/reader/rcmetrics`'s private `chainAttrs()` helper (same
family/ID/name/selector pattern as `destChainAttrs()`/`sourceChainAttrs()` above, minus the
dest/source prefix — each of these metrics only ever describes one chain, never a lane). This
replaced the old bare numeric-selector `chain` label; `chainSelector` carries that exact numeric
value as a string if you need it.

**Every metric below also carries `destChainID`/`destChainFamily`/`destChainName`/
`destChainSelector`** — a second, dest-scoped `destChainAttrs()` in the same file, resolved once
per reader/config-poller instance. **Filter by `destChainID=~"$destChain"` alongside `chainID` on
every query in this table** — confirmed necessary by a real staging incident: a node runs one
reader/config-poller instance **per destination chain**, and more than one instance can
independently track the same source chain (two destination-chain lanes both listing Solana as a
source chain collided into one series per `(chainID, node_id)` without this, `max()`/`sum()` across
them silently mixing a healthy lane's fast-refreshing cache with a different, broken lane's stuck
one). `query` names the specific read. Design + gap rationale: `docs/metrics/reader-metrics.md`.

| metric | otel_type | key labels | registration |
|---|---|---|---|
| `ccip_reader_read_outcome` | Counter | chain(dest)+dest + `query,outcome="ok"\|"empty"\|"error"` | beholder-only |¹
| `ccip_reader_read_empty` | Counter | chain+dest + `query` | beholder-only |
| `ccip_reader_chain_gap` | Counter | chain+dest + `query,state` | beholder-only |
| `ccip_reader_msg_dropped` | Counter | dest + `query,reason` (no chain — instrument registered, no call site yet) | beholder-only |
| `ccip_reader_chain_fee_components` | Gauge | chain+dest + `feeType` | beholder-only (promauto side unchanged, bare `chainFamily,chainID`, no dest disambiguation) |
| `ccip_reader_config_cache_age_seconds` | Gauge | chain+dest + `kind="chain"\|"source"` | beholder-only |
| `ccip_reader_config_poll_success` | Counter | chain+dest | beholder-only |
| `ccip_reader_config_poll_failure` | Counter | chain+dest + `reason="no_chain_accessor"\|"batch_fetch_failed"\|"no_chain_cache"` | beholder-only |
| `ccip_reader_config_cache_overwritten_empty` | Counter | chain+dest + `kind` | beholder-only |
| `ccip_reader_config_poller_last_success_timestamp` | Gauge | chain+dest | beholder-only |

("chain" above = `chainID,chainFamily,chainName,chainSelector`; "dest" =
`destChainID,destChainFamily,destChainName,destChainSelector` — both from `chainAttrs()`/
`destChainAttrs()` in `pkg/reader/rcmetrics`.)

¹ `ccip_reader_read_outcome` is the one semantic exception in this table: it's recorded at the
`observedCCIPReader` wrapper level for every reader call, so its chain-identity labels describe
the **destination** chain (this reader instance's) — meaning its `chainID` and `destChainID` are
always the same value for this one metric (redundant but harmless), unlike "whichever chain this
specific read is about" like every other row here. Before this rename, that distinction was visible
in the label *name itself* (`chainID` vs. the rest of the table's bare `chain`); now every row
shares identical label keys, so the distinction is easy to miss — don't assume this row's `chainID`
means what `read_empty`'s or `chain_gap`'s does. Filter it with `chainID=~"$destChain"` (its own `destChainID` label works identically here, since
they're always the same value for this metric — but don't confuse either with the commit-plugin
metrics' `destChainID` above, a different package's `destChainAttrs()`, even though the label name
happens to match).

## Known gaps

- This runbook assumes `RMNEnabled` is hardcoded off (current production state). If that
  changes, the RMN-signature branches dropped from step4 need to be reinstated.
