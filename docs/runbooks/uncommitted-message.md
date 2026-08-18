---
name: uncommitted-message
description: Triage a message that has not landed on the destination offramp within the expected time window.
trigger: "message reported in-flight for > 15m: onramp seq num observed, no corresponding offramp commit"
severity: page
owner: ccip-commit-oncall
inputs:
  sourceChain: {type: string, description: "source_network_name label value, e.g. chain-a. May be handed to you as a chain ID, chain selector, or name -- see docs/runbooks/chain-identifiers.md to translate"}
  destChain: {type: string, description: "chainID/chain_id label value of the destination chain, e.g. chain-b. Same translation note as sourceChain"}
  seqNum: {type: integer, description: "onramp sequence number of the message under investigation (X)"}
  msgID: {type: string, description: "message ID, hex, for cross-referencing logs/explorers"}
  fRoleDON: {type: integer, required: false, description: "optional. If known, expected live oracle count for destChain's DON is 3*fRoleDON+1 and the consensus threshold is 2*fRoleDON+1. Omit if uncertain (e.g. local devenvs with a possible bootstrap-node offset) -- scenario3's H0 check reports the raw live oracle count either way, it just can't grade it against a threshold without this."}
related:
  - docs/metrics/commit-metrics.md        # design rationale + full metric-by-metric findings
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
      + [step0 — is the plugin alive at all?](#step0-is-the-plugin-alive-at-all)
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
  - id: step0
    check: plugin_heartbeat
    query: 'sum(rate(ccip_commit_plugin_heartbeat_total{chainID=~"$destChain"}[1m]))'
    condition: "result == 0"
    if_true:  {action: "REPORT:ccip-commit-oncall", reason: "plugin process is not running rounds at all; every other metric below may simply be stale, not bad. Always run followup_query before reporting -- a bare 'not running' is a real but unsatisfyingly blunt finding on its own (confirmed by live testing: an agent that stopped here reported a correct but generic conclusion, then went and ran this exact followup_query anyway 'out of curiosity' because the plain heartbeat result didn't feel actionable)", followup_query: 'count(count by (csa_public_key) (timestamp(ccip_commit_plugin_heartbeat_total{chainID=~"$destChain", phase="observation"}) > time() - 60))', followup_meaning: "live oracle headcount -- this is a DON-level signal, same trust tier as heartbeat itself, safe to check here even though lane-specific metrics below are not (yet) trustworthy. A low count turns 'plugin not running' into 'N of the DON's oracles are actually alive', which is what an on-call needs next regardless of whether it's below consensus threshold -- grade it against $fRoleDON per scenario3b's rule if known, otherwise report the raw number"}
    if_false: {action: "CONTINUE:step1"}
    automatable: true
    note: "window is deliberately 1m, not 5m -- confirmed by live testing (see scenario3b's note) that rate() extrapolates across its whole window, so a plugin that died N minutes ago still reads as alive for up to [window] more minutes. Don't widen this without re-testing against a forced failure."

  - id: step1
    check: offramp_next_seq_num
    query: 'max by (source_network_name) (ccip_commit_offramp_next_seq_num{source_network_name=~"$sourceChain"})'
    condition: "result > $seqNum"
    if_true:  {action: "STOP", reason: "a root covering X already landed on the offramp; commit plugin's job is done, hand off to exec on-call for delivery status"}
    if_false: {action: "CONTINUE:step2"}
    automatable: true

  - id: step2
    check: onramp_max_seq_num
    query: 'max by (source_network_name) (ccip_commit_onramp_max_seq_num{source_network_name=~"$sourceChain"})'
    condition: "result < $seqNum"
    if_true:  {action: "REPORT:chain-infra-oncall", reason: "onramp read is lagging; check ccip_commit_merkleroot_observation_errors_total by reason (no_bindings/timeout/rpc_error) for the source-chain read/infra issue", followup_query: 'sum(rate(ccip_commit_merkleroot_observation_errors_total{source_network_name=~"$sourceChain"}[5m])) by (reason)'}
    if_false: {action: "CONTINUE:step2b"}
    automatable: true

  - id: step2b
    check: per_chain_consensus_dropped
    queries:
      - 'sum(rate(ccip_commit_consensus_dropped_total{source_network_name=~"$sourceChain", objectName=~"MerkleRoot|OnRampMaxSeqNums"}[5m])) by (objectName, reason)'
      - 'sum(rate(ccip_commit_offramp_consensus_insufficient_total{source_network_name=~"$sourceChain"}[5m]))'
    condition: "any result > 0"
    if_true:  {action: "REPORT:ccip-commit-oncall", reason: "DON could not reach consensus on a per-chain value for this lane. For ccip_commit_consensus_dropped: reason='split' means disagreeing oracles, reason='insufficient_agreement' means too few oracles observed the key, reason='threshold_not_defined' is a config/data mismatch. ccip_commit_offramp_consensus_insufficient means the DON could not agree on OffRampNextSeqNums for this lane. The message cannot advance until the disagreement resolves."}
    if_false: {action: "CONTINUE:step3"}
    automatable: true

  - id: step3
    check: consensus_observation_failed
    query: 'sum(rate(ccip_commit_consensus_observation_failed_total{chain_id=~"$destChain"}[5m]))'
    condition: "result > 0"
    if_true:  {action: "CONTINUE:scenario3", reason: "destination-chain-wide consensus failure — see deep dive"}
    if_false: {action: "CONTINUE:step4"}
    automatable: true

  - id: step4
    check: cursing
    queries:
      - 'max(ccip_commit_rmn_curse_active{chain_id=~"$destChain", curse_type="global"})'
      - 'max(ccip_commit_rmn_curse_active{chain_id=~"$destChain", curse_type="destination"})'
      - 'max(ccip_commit_source_chain_cursed{source_network_name=~"$sourceChain"})'
    condition: "any result == 1"
    if_true:  {action: "REPORT:curse-owner", reason: "chain A or B is cursed; likely intentional/incident-flagged, hand off to whoever owns the curse"}
    if_false: {action: "CONTINUE:step5"}
    automatable: true

  - id: step5
    check: report_transmission
    query: 'sum(rate(ccip_commit_report_transmission_gave_up_total{source_network_name=~"$sourceChain"}[5m]))'
    condition: "result > 0"
    if_true:  {action: "CONTINUE:scenario2", reason: "report is being built but never landing — see deep dive"}
    if_false: {action: "REPORT:ccip-commit-oncall", reason: "no failure signal found; likely just backlog size — report ccip_commit_pending_messages and an ETA estimate from round cadence", followup_query: 'max by (source_network_name) (ccip_commit_pending_messages{source_network_name=~"$sourceChain"})'}
    automatable: true

  - id: scenario2
    check: report_validation_rejected
    query: 'sum(rate(ccip_commit_report_validation_rejected_total{chainID=~"$destChain", phase="should_transmit"}[15m])) by (reason)'
    condition: "any reason count > 0"
    if_true:
      config_digest_mismatch:  {action: "REPORT:ccip-commit-oncall", reason: "config sync issue; cross-check ccip_commit_config_digest_mismatch"}
      config_digest_check_error: {action: "REPORT:chain-infra-oncall", reason: "error *reading* the offramp's config digest (RPC/read failure), distinct from an actual mismatch -- chain-B read/infra issue, not a config sync issue"}
      stale:                   {action: "STOP", reason: "often benign — a different overlapping report already landed; re-check step1, the gauge may have just moved"}
      dest_not_supported:      {action: "REPORT:ccip-commit-oncall", reason: "transmission-schedule/config bug"}
      dest_support_check_error: {action: "REPORT:chain-infra-oncall", reason: "error *checking* dest-chain support (RPC/read failure), distinct from a real config gap"}
      cursed:                  {action: "CONTINUE:step4", reason: "converges with step4"}
      cursed_check_error:      {action: "REPORT:chain-infra-oncall", reason: "error *checking* the curse state (RPC/read failure), not an actual curse -- don't conflate with step4's cursed=1 case"}
      empty_root:              {action: "REPORT:ccip-commit-oncall", reason: "report-builder produced an empty merkle root; likely a bug upstream in the merkleroot processor, not an infra issue"}
      invalid_seqnum_range:    {action: "REPORT:ccip-commit-oncall", reason: "report-builder produced start>end seqnums; likely a bug upstream in the merkleroot processor"}
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
    query: 'max by (source_network_name) (ccip_commit_pending_messages{dest_network_name=~".*"})'
    condition: "only $sourceChain's series is climbing, others into the same destChain are flat"
    if_true:  {action: "CONTINUE:step4", reason: "this branch is wrong if only one lane is affected — back to step4/step5"}
    if_false: {action: "CONTINUE:scenario3b", reason: "confirmed destination-chain-wide"}
    automatable: true

  - id: scenario3b
    check: h0_live_oracle_count
    query: 'count(count by (csa_public_key) (timestamp(ccip_commit_plugin_heartbeat_total{chainID=~"$destChain", phase="observation"}) > time() - 60))'
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
    query: 'sum(rate(ccip_commit_fchain_read_errors_total{chainID=~"$destChain"}[5m]))'
    condition: "result spiking broadly across oracles"
    if_true:  {action: "REPORT:home-chain-infra-oncall", reason: "H1 — too few oracles could read FChain; home-chain RPC/read outage"}
    if_false: {action: "REPORT:ccip-commit-oncall", reason: "H2 — oracles likely disagree (split vote). Confirm with ccip_commit_consensus_dropped{objectName=\"fChain\", reason=\"split\"} and corroborate with ccip_commit_config_digest_mismatch flipping for a subset of oracles around the same time. If scenario3b (H0) wasn't able to rule out insufficient participation (fRoleDON unknown), treat this H2 conclusion as low-confidence, not a firm diagnosis.", followup_query: 'ccip_commit_config_digest_mismatch{chain_id=~\"$destChain\"}'}
    automatable: true
```

## Steps

### step0 — is the plugin alive at all?

Not in the original design doc's walkthrough, but should be checked first: every metric below
reports "no signal" identically whether the plugin is healthy-and-idle or wedged/crashed. Rule
that out before reading anything else as a bad value.

```promql
sum(rate(ccip_commit_plugin_heartbeat_total{chainID=~"$destChain"}[1m]))
```
`0` → before reporting, also run the oracle headcount (same query as scenario3b's H0 —
it's a DON-level signal, safe to check even though lane-specific metrics aren't trustworthy yet):
```promql
count(count by (csa_public_key) (timestamp(ccip_commit_plugin_heartbeat_total{chainID=~"$destChain", phase="observation"}) > time() - 60))
```
Report both numbers together. "Plugin not running rounds" on its own is a real but blunt
finding — confirmed by live testing that it's unsatisfying enough that whoever's investigating
will go run this exact query anyway; do it here instead of making them rediscover it. `> 0` →
continue to step1.

### step1 — is this even a commit-plugin problem?

```promql
max by (source_network_name) (ccip_commit_offramp_next_seq_num{source_network_name=~"$sourceChain"})
```
`> $seqNum` → a root covering it already landed on the offramp. Commit plugin's job is done;
hand off to exec on-call. `<= $seqNum` → continue.

### step2 — has the plugin observed it on-chain yet?

```promql
max by (source_network_name) (ccip_commit_onramp_max_seq_num{source_network_name=~"$sourceChain"})
```
`< $seqNum` → onramp read is lagging. Check the reason breakdown:
```promql
sum(rate(ccip_commit_merkleroot_observation_errors_total{source_network_name=~"$sourceChain"}[5m])) by (reason)
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
sum(rate(ccip_commit_consensus_dropped_total{source_network_name=~"$sourceChain", objectName=~"MerkleRoot|OnRampMaxSeqNums"}[5m])) by (objectName, reason)
```

`ccip_commit_consensus_dropped` `reason` values:
- `split` — two or more distinct values each cleared the threshold (oracles disagree)
- `insufficient_agreement` — no single value reached the threshold (too few oracles observed the key)
- `threshold_not_defined` — the consensus code had no threshold configured for this key (config/data mismatch)

OffRampNextSeqNums consensus failures use a dedicated counter (same path conceptually, but a
custom consensus helper):
```promql
sum(rate(ccip_commit_offramp_consensus_insufficient_total{source_network_name=~"$sourceChain"}[5m]))
```

Any of these `> 0` for the lane → the message cannot advance until the disagreement resolves.
Report and stop. All flat (including empty) → continue.

### step3 — is the round producing outcomes at all?

```promql
sum(rate(ccip_commit_consensus_observation_failed_total{chain_id=~"$destChain"}[5m]))
```
Incrementing → destination-chain-wide consensus failure, jump to [Scenario 3](#scenario-3--consensus-never-completes-step3).
Flat → continue.

### step4 — is chain A or B cursed?

```promql
max(ccip_commit_rmn_curse_active{chain_id=~"$destChain", curse_type="global"})
max(ccip_commit_rmn_curse_active{chain_id=~"$destChain", curse_type="destination"})
max(ccip_commit_source_chain_cursed{source_network_name=~"$sourceChain"})
```
Any `== 1` → found it; likely intentional/incident-flagged, report to whoever owns the curse and
stop. All `0` → continue.

> RMN-signature branches (signed-roots-dropped, chain quorum, Byzantine root disagreement) are
> intentionally absent from this runbook: `RMNEnabled` is hardcoded off in production, so those
> paths are dead code. Cursing (this step) is the one RMN-adjacent signal still live.

### step5 — is a report being built but never landing?

```promql
sum(rate(ccip_commit_report_transmission_gave_up_total{source_network_name=~"$sourceChain"}[5m]))
```
Incrementing/maxed → jump to [Scenario 2](#scenario-2--report-transmission-stuck-step5).
Otherwise → likely just backlog size:
```promql
max by (source_network_name) (ccip_commit_pending_messages{source_network_name=~"$sourceChain"})
```
Report the backlog size and an ETA estimate from round cadence, and stop.

## Deep dives

### Scenario 2 — report transmission stuck (step5)

Three hypotheses, cheapest first.

**A — never transmitted (rejected pre-send).**
```promql
sum(rate(ccip_commit_report_validation_rejected_total{chainID=~"$destChain", phase="should_transmit"}[15m])) by (reason)
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
- `empty_root` / `invalid_seqnum_range` → the report-builder produced an invalid report; likely a
  bug upstream in the merkleroot processor, not an infra issue.
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
  max by (source_network_name) (ccip_commit_pending_messages{dest_network_name=~".*"})
  ```
  If only `$sourceChain` is affected, this branch is wrong — back to step4/step5.
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
  count(count by (csa_public_key) (timestamp(ccip_commit_plugin_heartbeat_total{chainID=~"$destChain", phase="observation"}) > time() - 60))
  ```
  If `$fRoleDON` is known: `< 2*$fRoleDON+1` → this **is** the answer, stop here — don't let a
  flat H1 push you toward H2 below, H1/H2 are checks for a different failure mode and will likely
  look inconclusive here even though they're not the cause. `>= 2*$fRoleDON+1` → oracle count
  isn't the explanation, continue. If `$fRoleDON` is unknown, report the raw count as context and
  continue regardless — but treat whatever H1/H2 concludes below as lower-confidence, since you
  haven't ruled this out.
- **H1 — too few oracles could read it.**
  ```promql
  sum(rate(ccip_commit_fchain_read_errors_total{chainID=~"$destChain"}[5m]))
  ```
  Spiking broadly across oracles → home-chain RPC/read outage.
- **H2 — oracles disagree (split vote).** Check `ccip_commit_consensus_dropped{objectName="fChain", reason="split"}`:
  ```promql
  sum(rate(ccip_commit_consensus_dropped_total{chainID=~"$destChain", objectName="fChain", reason="split"}[5m])) by (key)
  ```
  This is genuine disagreement among oracles that DID submit a value — not the same as H0's "too few submitted at all." Corroborate with `ccip_commit_config_digest_mismatch` flipping for a subset of oracles around the same time:
  ```promql
  ccip_commit_config_digest_mismatch{chain_id=~"$destChain"}
  ```

## Metric reference

`otel_type` is what governs whether a `_total`/`_bucket` suffix gets added by an OTel→Prometheus
pipeline — see [the naming note above](#a-note-on-counter-naming). `dual` means the metric is
*also* directly `promauto`-registered (scrapeable without going through Beholder/OTel at all);
everything else is Beholder-only.

| metric | otel_type | key labels | registration |
|---|---|---|---|
| `ccip_commit_plugin_heartbeat` | Counter | `chainFamily,chainID,phase` | beholder-only |
| `ccip_commit_report_validation_rejected` | Counter | `chainFamily,chainID,phase,reason` | beholder-only |
| `ccip_commit_onramp_max_seq_num` | Gauge | `sourceChainFamily,sourceChainID,source_network_name` | beholder-only |
| `ccip_commit_offramp_next_seq_num` | Gauge | `sourceChainFamily,sourceChainID,source_network_name` | beholder-only |
| `ccip_commit_pending_messages` | Gauge | `sourceChainFamily,sourceChainID,source_network_name` | beholder-only |
| `ccip_commit_rmn_curse_active` | Gauge | `chain_family,chain_id,curse_type` | beholder-only |
| `ccip_commit_source_chain_cursed` | Gauge | `sourceChainFamily,sourceChainID,source_network_name` | beholder-only |
| `ccip_commit_merkleroot_observation_errors` | Counter | `...,reason` | beholder-only |
| `ccip_commit_fchain_read_errors` | Counter | `chainFamily,chainID` | beholder-only |
| `ccip_commit_consensus_observation_failed` | Counter | `chain_family,chain_id` | beholder-only |
| `ccip_commit_report_transmission_gave_up` | Counter | `sourceChainFamily,sourceChainID,source_network_name` | beholder-only |
| `ccip_commit_report_transmission_attempts` | Histogram | `chainFamily,chainID,success` | beholder-only |
| `ccip_commit_range_truncated` | Counter | `...,source_network_name` | beholder-only |
| `ccip_commit_offramp_lane_status` | Gauge | `...,status` | beholder-only |
| `ccip_commit_seqnum_invariant_violation` | Counter | `...,type` | beholder-only |
| `ccip_commit_offramp_consensus_insufficient` | Counter | `...,source_network_name` | beholder-only |
| `ccip_commit_config_digest_mismatch` | Gauge | `chain_family,chain_id` | dual |
| `ccip_commit_max_sequence_number` | Gauge | `chainFamily,chainID,sourceChainFamily,sourceChain,method,source_network_name,dest_network_name` | dual |
| `ccip_commit_latest_round_id` | Gauge | `source_network_name,dest_network_name,contract_address,plugin` | dual |
| `ccip_commit_processor_errors` | Counter | `chainFamily,chainID,processor,method` | dual |
| `ccip_commit_processor_latency` | Histogram | `chainFamily,chainID,processor,method` | dual |
| `ccip_commit_loopp_ccip_provider_supported` | Gauge | `chain_family` | dual |
| `ccip_commit_consensus_dropped` | Counter | `chainFamily,chainID,objectName,key,reason,source_network_name` | beholder-only |

Note the label-casing inconsistency in the source (`chainFamily`/`chainID` vs.
`chain_family`/`chain_id`) — copy the exact casing from this table per metric, it's not
uniform across the file.

## Known gaps

- This runbook assumes `RMNEnabled` is hardcoded off (current production state). If that
  changes, the RMN-signature branches dropped from step4 need to be reinstated.
