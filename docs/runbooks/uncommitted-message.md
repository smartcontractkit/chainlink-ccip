---
name: uncommitted-message
description: Triage a message that has not landed on the destination offramp within the expected time window.
trigger: "message reported in-flight for > 15m: onramp seq num observed, no corresponding offramp commit"
severity: page
owner: ccip-commit-oncall
inputs:
  sourceChain: {type: string, description: "source_network_name label value, e.g. chain-a"}
  destChain: {type: string, description: "chainID/chain_id label value of the destination chain, e.g. chain-b"}
  seqNum: {type: integer, description: "onramp sequence number of the message under investigation (X)"}
  msgID: {type: string, description: "message ID, hex, for cross-referencing logs/explorers"}
related:
  - docs/metrics/commit-metrics.md        # design rationale + full metric-by-metric findings
  - devenv/dashboards/commit-plugin.json  # Grafana rendering of this runbook (Guided Debug section)
status: living
---

- [Runbook: uncommitted message](#runbook-uncommitted-message)
   * [For agents](#for-agents)
      + [A note on counter naming](#a-note-on-counter-naming)
   * [Decision graph](#decision-graph)
   * [Steps](#steps)
      + [step0 — is the plugin alive at all?](#step0-is-the-plugin-alive-at-all)
      + [step1 — is this even a commit-plugin problem?](#step1-is-this-even-a-commit-plugin-problem)
      + [step2 — has the plugin observed it on-chain yet?](#step2-has-the-plugin-observed-it-on-chain-yet)
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

If a query's metric doesn't exist on your datasource (e.g. `ccip_commit_consensus_dropped`,
see [Known gaps](#known-gaps)), treat the result as unknown, not as `0`/flat — say so explicitly
rather than silently treating "no such metric" as "no signal."

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
    query: 'sum(rate(ccip_commit_plugin_heartbeat_total{chainID=~"$destChain"}[5m]))'
    condition: "result == 0"
    if_true:  {action: "REPORT:ccip-commit-oncall", reason: "plugin process is not running rounds at all; every other metric below may simply be stale, not bad"}
    if_false: {action: "CONTINUE:step1"}
    automatable: true

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
      config_digest_mismatch: {action: "REPORT:ccip-commit-oncall", reason: "config sync issue; cross-check ccip_commit_config_digest_mismatch"}
      stale:                  {action: "STOP", reason: "often benign — a different overlapping report already landed; re-check step1, the gauge may have just moved"}
      dest_not_supported:     {action: "REPORT:ccip-commit-oncall", reason: "transmission-schedule/config bug"}
      cursed:                 {action: "CONTINUE:step4", reason: "converges with step4"}
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
    check: h1_vs_h2
    query: 'sum(rate(ccip_commit_fchain_read_errors_total{chainID=~"$destChain"}[5m]))'
    condition: "result spiking broadly across oracles"
    if_true:  {action: "REPORT:home-chain-infra-oncall", reason: "H1 — too few oracles could read FChain; home-chain RPC/read outage"}
    if_false: {action: "REPORT:ccip-commit-oncall", reason: "H2 — oracles likely disagree (split vote). Needs ccip_commit_consensus_dropped{reason=\"split\"}, NOT IMPLEMENTED as of this writing — see Known gaps. Corroborate with ccip_commit_config_digest_mismatch flipping for a subset of oracles around the same time.", followup_query: 'ccip_commit_config_digest_mismatch{chain_id=~"$destChain"}'}
    automatable: true
```

## Steps

### step0 — is the plugin alive at all?

Not in the original design doc's walkthrough, but should be checked first: every metric below
reports "no signal" identically whether the plugin is healthy-and-idle or wedged/crashed. Rule
that out before reading anything else as a bad value.

```promql
sum(rate(ccip_commit_plugin_heartbeat_total{chainID=~"$destChain"}[5m]))
```
`0` → nothing else here is trustworthy; report and stop. `> 0` → continue.

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
- `stale` climbing → often benign — a different, overlapping report already landed; re-check
  step1, it may have just moved.
- `dest_not_supported` → transmission-schedule/config bug.
- `cursed` → converges with step4.
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
- **H1 — too few oracles could read it.**
  ```promql
  sum(rate(ccip_commit_fchain_read_errors_total{chainID=~"$destChain"}[5m]))
  ```
  Spiking broadly across oracles → home-chain RPC/read outage.
- **H2 — oracles disagree (split vote).** *(Needs `ccip_commit_consensus_dropped{objectName="fChain", reason="split"}` — **not implemented**, reviewed design only, see [Known gaps](#known-gaps).)*
  Corroborate with:
  ```promql
  ccip_commit_config_digest_mismatch{chain_id=~"$destChain"}
  ```
  flipping for a subset of oracles around the same time.

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
| `ccip_commit_consensus_dropped` | — | `objectName,key,reason` | **not implemented** (reviewed design only) |

Note the label-casing inconsistency in the source (`chainFamily`/`chainID` vs.
`chain_family`/`chain_id`) — copy the exact casing from this table per metric, it's not
uniform across the file.

## Known gaps

- `ccip_commit_consensus_dropped` (Scenario 3, H1/H2 split) is **not implemented** — reviewed
  design only. Any query against it will return no data; that's expected, not a signal.
- This runbook assumes `RMNEnabled` is hardcoded off (current production state). If that
  changes, the RMN-signature branches dropped from step4 need to be reinstated.
