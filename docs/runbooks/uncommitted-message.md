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
  fRoleDON: {type: integer, required: false, description: "optional. If known, expected live oracle count for destChain's DON is 3*fRoleDON+1 and the consensus threshold is 2*fRoleDON+1. Omit if uncertain (e.g. local devenvs with a possible bootstrap-node offset) -- scenario3's H0 check reports the raw live oracle count either way, it just can't grade it against a threshold without this."}
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

1. **[`uncommitted-message.yaml`](uncommitted-message.yaml)** is the authoritative control flow —
   the single source of truth this doc and the `cmd/runbook` tool both execute. It lives next to
   this file so the two can't silently drift: change the YAML and the tool picks it up. Each step
   has a `query` (fill in `$sourceChain`/`$destChain`/`$seqNum` from the Inputs above) and a
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

The authoritative control flow is [`uncommitted-message.yaml`](uncommitted-message.yaml) — the
single source of truth this doc and the `cmd/runbook` tool both execute. It defines each step's
`queries`, its `condition`, and the `STOP` / `CONTINUE:<id>` / `REPORT:<owner>` / `AGENT`
outcomes in YAML. The prose in [Steps](#steps) below explains *why* each step exists; the YAML is
what you (or an agent) run step-for-step.

To execute it deterministically against your datasource:

```sh
runbook run uncommitted-message -D destChain=<value> -D sourceChain=<value> -D seqNum=<value>
```

The bulk of the machine-readable steps are `automatable`; the two chain-explorer / TXM
hypotheses in [Scenario 2](#scenario-2--report-transmission-stuck-step5) are not (no query here),
as annotated in the YAML.

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
