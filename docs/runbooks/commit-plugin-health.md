---
name: commit-plugin-health
description: Point-in-time health check for a commit plugin instance and its merkle root processor. No incident or specific message required.
trigger: "ad hoc / scheduled health check, or as step0 before docs/runbooks/uncommitted-message.md"
severity: informational
owner: ccip-commit-oncall
inputs:
  destChain: {type: string, description: "chainID/chain_id label value of the node's destination chain"}
  sourceChains: {type: string, description: "source_network_name regex filter; defaults to all lanes into destChain", default: ".*"}
related:
  - docs/runbooks/uncommitted-message.md   # incident triage for one specific stuck message
  - docs/metrics/commit-metrics.md         # design rationale + full metric-by-metric findings
  - devenv/dashboards/commit-plugin.json   # Grafana rendering of this runbook (Health section)
status: living
---

- [Runbook: commit plugin health](#runbook-commit-plugin-health)
   * [For agents](#for-agents)
      + [Label key cheat sheet](#label-key-cheat-sheet)
      + [Empty result sets: the rule that actually matters](#empty-result-sets-the-rule-that-actually-matters)
   * [Checklist](#checklist)
   * [Groups](#groups)
      + [Liveness](#liveness)
      + [Lane throughput & backlog](#lane-throughput--backlog)
      + [Cursing & consensus](#cursing--consensus)
      + [Report transmission](#report-transmission)
   * [Aggregation rule](#aggregation-rule)
   * [Output contract](#output-contract)

# Runbook: commit plugin health

Unlike [`uncommitted-message.md`](uncommitted-message.md), this isn't triggered by a specific
stuck message — it's "is this commit plugin instance (and its merkle root processor) healthy
right now." Run it periodically, on demand, or as the first thing you do before diving into a
specific-message incident (its `step0` is a subset of this runbook's `heartbeat_observation` /
`heartbeat_outcome` checks).

## For agents

Every check below is independent — run all of them, don't stop early. Each has a fixed
severity mapping, not a branch.

### Label key cheat sheet

The destination chain is filtered by **two different label names depending on the metric** —
this is a real inconsistency in `commit/metrics/prom.go`, not a typo to normalize away. Copy the
label key from each check's own query; don't assume one spelling applies file-wide:

| label key | checks that use it |
|---|---|
| `chainID` | `heartbeat_observation`, `heartbeat_outcome`, `processor_errors`, `processor_latency_p95`, `fchain_read_errors`, `report_validation_rejected`, `report_transmission_attempts_p95` |
| `chain_id` | `config_digest_mismatch`, `rmn_curse_active`, `consensus_observation_failed` |
| `source_network_name` | every check under `lane_throughput`, plus `source_chain_cursed`, `report_transmission_gave_up` |

### Empty result sets: the rule that actually matters

An empty Prometheus/PromQL result (zero series returned) is **not** the same as "all series are
0" — `ok_if: "all series == 0"` is vacuously true over an empty set, and resolving it that way
silently turns "we have no idea" into "confirmed fine." Getting this wrong is the single easiest
way to make this checklist wrong in practice, so the rule is:

1. Every check below is tagged `always_emitted: true` or `always_emitted: false`.
2. **`always_emitted: true`** means the underlying code records this metric unconditionally,
   every round, regardless of outcome (heartbeats, gauges reported every round, the latency
   histogram). An empty result here means the telemetry pipeline itself isn't delivering data for
   this series — status is `UNKNOWN`, full stop, regardless of what the check's own
   `ok_if`/`warn_if`/`crit_if` says.
3. **`always_emitted: false`** means the underlying metric is an event counter that Prometheus/
   OTel only creates a time series for *after* it increments at least once — these are almost
   always the error/failure counters. **An empty result here is the expected, healthy state**
   (the event has never happened) — treat it as `value == 0` and evaluate `ok_if`/`warn_if`/
   `crit_if` against that value normally, do not report `UNKNOWN`.
4. The one exception: if step 3 would resolve a check to `OK`-via-empty, but any
   `always_emitted: true` check in the same run is itself `UNKNOWN` (pipeline may be broken),
   don't trust the distinction — treat the `always_emitted: false` check as `UNKNOWN` too, and
   say so in `concerns`. You cannot tell "never happened" apart from "telemetry is broken" once
   you've already lost confidence in the pipeline.

For metric types, label sets, and the `_total`/`_bucket` suffix caveat, see
[`uncommitted-message.md#metric-reference`](uncommitted-message.md#metric-reference).

Produce the [output contract](#output-contract) at the end — don't just narrate findings inline.

## Checklist

```yaml
checks:
  # --- liveness ---
  - id: heartbeat_observation
    group: liveness
    always_emitted: true
    query: 'sum(rate(ccip_commit_plugin_heartbeat_total{chainID=~"$destChain", phase="observation"}[5m]))'
    severity: {crit_if: "result == 0", ok_if: "result > 0"}
    owner: ccip-commit-oncall
    note: "unconditional per-round liveness signal for Observation(); 0 or empty means every other check below may simply be stale, not bad"

  - id: heartbeat_outcome
    group: liveness
    always_emitted: true
    query: 'sum(rate(ccip_commit_plugin_heartbeat_total{chainID=~"$destChain", phase="outcome"}[5m]))'
    severity: {crit_if: "result == 0", ok_if: "result > 0"}
    owner: ccip-commit-oncall
    note: "same signal for Outcome(); flat while observation is healthy means OCR is failing to schedule/deliver to the outcome stage"

  - id: config_digest_mismatch
    group: liveness
    always_emitted: true
    query: 'max(ccip_commit_config_digest_mismatch{chain_id=~"$destChain"})'
    severity: {crit_if: "result == 1", ok_if: "result == 0"}
    owner: ccip-commit-oncall
    note: "home-chain config digest differs from the offramp's; root cause for a wide range of downstream rejections"

  - id: processor_errors
    group: liveness
    always_emitted: false
    query: 'sum(rate(ccip_commit_processor_errors_total{chainID=~"$destChain"}[15m])) by (processor, method)'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: ccip-commit-oncall
    note: "generic TrackedProcessor error counter; empty means zero errors, not unknown. Does not include swallowed per-goroutine failures (see merkleroot_observation_errors in uncommitted-message.md)"

  - id: processor_latency_p95
    group: liveness
    always_emitted: true
    query: 'histogram_quantile(0.95, sum(rate(ccip_commit_processor_latency_bucket{chainID=~"$destChain"}[15m])) by (le, processor, method))'
    severity: {info: true}
    owner: ccip-commit-oncall
    note: "report raw value only, this is a single point-in-time sample with nothing to trend against -- do not infer a trend from one run. If a value lands exactly on a bucket boundary (e.g. pinned at the histogram's max bucket), report that fact explicitly as it likely means real latency exceeds the highest defined bucket; still do not elevate this to WARN/CRIT, there is no defined SLO"

  # --- lane throughput & backlog (per source chain) ---
  - id: pending_messages
    group: lane_throughput
    always_emitted: true
    query: 'max by (source_network_name) (ccip_commit_pending_messages{source_network_name=~"$sourceChains"})'
    severity: {info: true}
    owner: ccip-commit-oncall
    note: "onRampMaxSeqNum - offRampNextSeqNum; no universal healthy value, watch for sustained growth not a single sample. Empty result here (no lanes matched) is UNKNOWN, not zero backlog -- this metric is always emitted per active lane"

  - id: range_truncated
    group: lane_throughput
    always_emitted: false
    query: 'sum(rate(ccip_commit_range_truncated_total{source_network_name=~"$sourceChains"}[15m])) by (source_network_name)'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: ccip-commit-oncall
    note: "chain-throughput ceiling (MaxMerkleTreeSize) being hit"

  - id: seqnum_invariant_violation
    group: lane_throughput
    always_emitted: false
    query: 'sum(rate(ccip_commit_seqnum_invariant_violation_total{source_network_name=~"$sourceChains"}[15m])) by (source_network_name, type)'
    severity: {crit_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: ccip-commit-oncall
    note: "offramp_ahead_of_onramp / onramp_max_zero / offramp_seqnum_regression -- documented in-code as stall signals, not routine noise"

  - id: offramp_consensus_insufficient
    group: lane_throughput
    always_emitted: false
    query: 'sum(rate(ccip_commit_offramp_consensus_insufficient_total{source_network_name=~"$sourceChains"}[15m])) by (source_network_name)'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: ccip-commit-oncall
    note: "DON couldn't agree on OffRampNextSeqNum for this chain; looks identical to \"no new messages\" without this"

  - id: offramp_lane_status
    group: lane_throughput
    always_emitted: true
    query: 'max by (source_network_name, status) (ccip_commit_offramp_lane_status{source_network_name=~"$sourceChains"})'
    severity: {crit_if: 'status="rmn_misconfigured" and result == 1', info_if: 'status=~"skipped_.*" and result == 1', ok_if: 'status="live" and result == 1'}
    owner: ccip-commit-oncall
    note: "rmn_misconfigured flags real config drift (a lane still expecting RMN blessing while RMN is globally off); skipped_* are expected states, not unhealthy. Reported for all four statuses every round by design, so empty result is UNKNOWN, not \"no active lanes\""

  # --- cursing & consensus ---
  - id: rmn_curse_active
    group: cursing_consensus
    always_emitted: true
    query: 'max(ccip_commit_rmn_curse_active{chain_id=~"$destChain", curse_type=~"global|destination"}) by (curse_type)'
    severity: {warn_if: "any series == 1", ok_if: "all series == 0"}
    owner: curse-owner
    note: "halts ALL reporting for the dest chain while active; often intentional/incident-flagged (see uncommitted-message.md step4) -- WARN not CRIT because an active curse is frequently expected, not a plugin bug"

  - id: source_chain_cursed
    group: cursing_consensus
    always_emitted: true
    query: 'max by (source_network_name) (ccip_commit_source_chain_cursed{source_network_name=~"$sourceChains"})'
    severity: {warn_if: "any series == 1", ok_if: "all series == 0"}
    owner: curse-owner
    note: "per-lane curse; same intentional-vs-bug ambiguity as rmn_curse_active"

  - id: consensus_observation_failed
    group: cursing_consensus
    always_emitted: false
    query: 'sum(rate(ccip_commit_consensus_observation_failed_total{chain_id=~"$destChain"}[5m]))'
    severity: {crit_if: "result > 0", ok_if: "result == 0 (including empty result)"}
    owner: ccip-commit-oncall
    note: "the only way a whole round fails: DON can't reach 2*fRoleDON+1 agreement on FChain for the dest chain. Destination-chain-wide by construction, not lane-specific"

  - id: fchain_read_errors
    group: cursing_consensus
    always_emitted: false
    query: 'sum(rate(ccip_commit_fchain_read_errors_total{chainID=~"$destChain"}[5m]))'
    severity: {warn_if: "result > 0", ok_if: "result == 0 (including empty result)"}
    owner: home-chain-infra-oncall
    note: "if this AND consensus_observation_failed are both firing, treat the combination as one CRIT finding, not two (see aggregation rule) -- likely H1, home-chain RPC/read outage"

  # --- report transmission ---
  - id: report_validation_rejected
    group: report_transmission
    always_emitted: false
    query: 'sum(rate(ccip_commit_report_validation_rejected_total{chainID=~"$destChain"}[15m])) by (phase, reason)'
    severity: {warn_if: 'any series with reason!="stale" > 0', info_if: 'only reason="stale" > 0', ok_if: "empty result, or all series == 0"}
    owner: ccip-commit-oncall
    note: "reason=\"stale\" is usually benign (an overlapping report already landed); other reasons are worth a look"

  - id: report_transmission_gave_up
    group: report_transmission
    always_emitted: false
    query: 'sum(rate(ccip_commit_report_transmission_gave_up_total{source_network_name=~"$sourceChains"}[15m])) by (source_network_name)'
    severity: {crit_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: ccip-commit-oncall
    note: "one of the most common on-call pages historically; previously a Warnw with no counter at all. If CRIT, uncommitted-message.md Scenario 2 may name a different owner (chain-b-txm-oncall) once root-caused -- this check alone doesn't know which"

  - id: report_transmission_attempts_p95
    group: report_transmission
    always_emitted: false
    query: 'histogram_quantile(0.95, sum(rate(ccip_commit_report_transmission_attempts_bucket{chainID=~"$destChain"}[15m])) by (le, success))'
    severity: {info: true}
    owner: ccip-commit-oncall
    note: "recorded once per transmission-check attempt; empty result means zero transmission cycles in the window, which is itself informational (report as INFO with value \"no data\"), not UNKNOWN. Rising attempts alongside report_transmission_gave_up firing is the actionable combination, not this alone"
```

## Groups

### Liveness

`heartbeat_observation` / `heartbeat_outcome` first — if either is `CRIT` or `UNKNOWN`, every
other check's `OK` in this run is suspect (a wedged plugin can leave gauges parked at their
last-good value, and an `always_emitted: false` check resolving empty-to-OK assumes the pipeline
itself is trustworthy). `config_digest_mismatch` and `processor_errors` are the other two checks
here with an unambiguous bad state; `processor_latency_p95` is context, not a verdict.

### Lane throughput & backlog

Run per source chain (`$sourceChains` regex, default all lanes into `$destChain`).
`seqnum_invariant_violation` is the one `CRIT` here — it's explicitly documented in-code as a
stall signal, not routine noise. `offramp_lane_status="rmn_misconfigured"` is the other `CRIT`:
config drift, not a transient condition. Both `pending_messages` and `offramp_lane_status` are
`always_emitted: true` — don't let an empty result for either read as "no lanes, nothing to
worry about."

### Cursing & consensus

Cursing is `WARN` not `CRIT` deliberately: an active curse is frequently an intentional,
already-known state (see `uncommitted-message.md` step4), and a health check that pages on
expected state trains people to ignore it. `consensus_observation_failed` is unambiguous `CRIT`
when it fires — by construction it can only mean the DON-wide FChain agreement failed for this
destination chain, which cannot happen for benign reasons — but it's `always_emitted: false`, so
absence is the expected/healthy state, not a gap in coverage.

### Report transmission

`report_transmission_gave_up` is `CRIT`. `report_validation_rejected` is split by reason:
`stale` alone is `INFO`, anything else present is `WARN`, and — being `always_emitted: false` —
a fully empty result is `OK`, not `UNKNOWN`.

## Aggregation rule

Overall status is the worst individual status, `CRIT` > `WARN` > `UNKNOWN` > `INFO`/`OK`, with
one explicit combination rule: **`fchain_read_errors` (`WARN`) firing at the same time as
`consensus_observation_failed` (`CRIT`) doesn't change the overall verdict (already `CRIT`) but
should be called out together in `concerns` as one finding, not two** — they're the same
incident (H1 in `uncommitted-message.md`'s Scenario 3), and reporting them separately reads as
two problems instead of one with corroborating evidence.

`UNKNOWN` is never silently dropped to `OK` — but per the [empty-result rule](#empty-result-sets-the-rule-that-actually-matters),
most checks should legitimately resolve `UNKNOWN` only when an `always_emitted: true` check comes
back empty, which in a genuinely healthy, actively-running system should be rare. If you're
seeing `UNKNOWN` on an `always_emitted: false` check, re-check that you applied the
empty-means-zero rule for that check before reporting it as a pipeline problem.

## Output contract

```yaml
health_report:
  destChain: string
  sourceChains: string        # the regex actually used
  timestamp: string           # ISO8601, when checks were run
  overall: OK | WARN | CRIT | UNKNOWN
  checks:
    - id: string               # matches an id in the Checklist above
      status: OK | WARN | CRIT | UNKNOWN | INFO
      value: string            # scalar checks: the raw number. Multi-series (`by (...)`) checks,
                                # regardless of how many grouping labels the query has: for each
                                # series, join ALL of its grouping labels into one
                                # "label1=v1,label2=v2" key, then ":", then the numeric result;
                                # separate series with "; ". E.g. a `by (source_network_name,
                                # status)` result becomes
                                # "source_network_name=chain-a,status=live:1; source_network_name=chain-a,status=rmn_misconfigured:0; ...".
                                # This is the canonical non-paraphrased form, not free-form
                                # summarization -- do not drop labels or collapse series to save space.
      note: string             # optional, only if it adds something the checklist note doesn't
  concerns:                    # every non-OK/non-INFO finding, worst first, plus any UNKNOWN.
                                # Use an empty list (`concerns: []`) when there are none -- never
                                # omit the key.
    - checks: [string]         # one or more related check ids (see the fchain/consensus combination rule)
      severity: WARN | CRIT | UNKNOWN
      summary: string
      owner: string             # copy from the matching check's `owner` field in the Checklist.
                                 # If a concern spans checks with different owners, list the
                                 # primary one first and name the others in `summary`.
```

Do not page, escalate, or act on any `concerns` entry — this contract's job is to hand a human
(or the next runbook) a report, the same non-negotiable boundary as `uncommitted-message.md`'s
`REPORT:<owner>`.
