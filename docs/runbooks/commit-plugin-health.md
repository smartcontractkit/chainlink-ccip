---
name: commit-plugin-health
description: Point-in-time health check for a commit plugin instance and its merkle root processor. No incident or specific message required.
trigger: "ad hoc / scheduled health check, or as step0 before docs/runbooks/uncommitted-message.md"
severity: informational
owner: ccip-commit-oncall
inputs:
  destChain: {type: string, description: "chainID/chain_id label value of the node's destination chain"}
  sourceChains: {type: string, description: "source_network_name regex filter; defaults to all lanes into destChain", default: ".*"}
  fRoleDON: {type: integer, required: false, description: "optional. If known, expected live oracle count for destChain's DON is 3*fRoleDON+1 and the consensus threshold is 2*fRoleDON+1. Omit if uncertain (e.g. local devenvs with a possible bootstrap-node offset) -- live_oracle_count reports the raw count either way, it just can't grade it against a threshold without this."}
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
      + [rate() detection lag: found by live testing, not designed in](#rate-detection-lag-found-by-live-testing-not-designed-in)
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

Every check in [`commit-plugin-health.yaml`](commit-plugin-health.yaml) is independent — run all
of them, don't stop early. The YAML is the single source of truth this doc and the `cmd/runbook`
tool both execute, so change the YAML and the tool picks it up. Each has a fixed severity mapping,
not a branch.

### Label key cheat sheet

The destination chain is filtered by **two different label names depending on the metric** —
this is a real inconsistency in `commit/metrics/prom.go`, not a typo to normalize away. Copy the
label key from each check's own query; don't assume one spelling applies file-wide:

| label key | checks that use it |
|---|---|
| `chainID` | `heartbeat_observation`, `heartbeat_outcome`, `processor_errors`, `processor_latency_p95`, `live_oracle_count`, `fchain_read_errors`, `report_validation_rejected`, `report_transmission_attempts_p95` |
| `chain_id` | `config_digest_mismatch`, `rmn_curse_active`, `consensus_observation_failed` |
| `source_network_name` | every check under `lane_throughput`, plus `source_chain_cursed`, `report_transmission_gave_up` |

`csa_public_key` is a different kind of label — it's not used to filter by chain, it's a
per-node identity present on *every* commit metric series (confirmed against a live devenv, not
just inferred). `live_oracle_count` is the only check that groups by it; nothing else in this
checklist needs to.

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

### `rate()` detection lag: found by live testing, not designed in

This was found by actually forcing a failure and running this checklist against it, not by
reading the doc carefully — worth internalizing before you edit any check's window size.
`rate(metric[Nm])` extrapolates across its *whole* window using the oldest and newest samples it
finds inside it. That means a `crit_if: result == 0` check (or a headcount built on
`rate(...) > 0`) does not go bad the instant the underlying process dies — it stays looking
healthy for up to `N` more minutes, because there were still real samples earlier in the window.
Confirmed empirically: with a `[5m]` window, `live_oracle_count` kept reporting the pre-outage
oracle count for several minutes after 3 of 5 oracles were actually stopped; only a `[1m]` window
(or, better, a `timestamp()`-based staleness check — what `live_oracle_count` now uses) reflected
the real state promptly. Every liveness check in this checklist uses a short window (`[1m]`) or a
direct staleness check specifically because of this. If you widen one, you're trading detection
speed for smoothing against transient blips — that's a real tradeoff, not a free cleanup, and any
resulting delay should be [re-tested against a forced failure](README.md#re-testing-a-runbook-after-editing-it),
not assumed correct because the query still reads sensibly.

## Checklist

The authoritative checklist is [`commit-plugin-health.yaml`](commit-plugin-health.yaml) — the
single source of truth this doc and the `cmd/runbook` tool both execute. It defines every check's
`query`, its `always_emitted` set, and its fixed severity rules. The prose under
[Groups](#groups) explains the reasoning; the YAML is what runs.

To execute it deterministically against your datasource:

```sh
runbook run commit-plugin-health -D destChain=<value> [-D sourceChains=<regex>] [-D fRoleDON=<n>]
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
absence is the expected/healthy state, not a gap in coverage. If it does fire, check
`live_oracle_count` before `fchain_read_errors`: an oracle that never started can't emit
`fchain_read_errors` at all (only a *surviving* oracle's own read failures are counted there), so
a low oracle count is a root cause `fchain_read_errors` is structurally unable to surface on its
own. `consensus_dropped` then names the *kind* of failure: `reason="split"` on `objectName="fChain"`
points to H2 (oracles disagree), while `reason="insufficient_agreement"` on `objectName="fChain"`
points to H1 (too few oracles could even report). `reason="threshold_not_defined"` is a config/data
mismatch.

### Report transmission

`report_transmission_gave_up` is `CRIT`. `report_validation_rejected` is split by reason:
`stale` alone is `INFO`, anything else present is `WARN`, and — being `always_emitted: false` —
a fully empty result is `OK`, not `UNKNOWN`.

## Aggregation rule

Overall status is the worst individual status, `CRIT` > `WARN` > `UNKNOWN` > `INFO`/`OK`, with
two explicit combination rules:

- **`fchain_read_errors` (`WARN`) firing at the same time as `consensus_observation_failed`
  (`CRIT`) doesn't change the overall verdict (already `CRIT`) but should be called out together
  in `concerns` as one finding, not two** — they're the same incident (H1 in
  `uncommitted-message.md`'s Scenario 3), and reporting them separately reads as two problems
  instead of one with corroborating evidence.
- **`live_oracle_count` resolving `CRIT` (or `WARN`) takes priority over whatever
  `fchain_read_errors` says, when both are present in the same `concerns` entry.** A low oracle
  count fully explains `consensus_observation_failed` on its own; `fchain_read_errors` being flat
  in that situation is expected (dead oracles can't emit it), not evidence that points elsewhere.
  Don't report "H1 ruled out, likely a split vote" if `live_oracle_count` already found the real
  cause — that conclusion would be actively wrong, not just unconfirmed.
- **`consensus_dropped` (`WARN`) firing at the same time as `consensus_observation_failed`
  (`CRIT`) should be folded into the same `concerns` entry and used to name the failure mode.**
  `reason="split"` on `objectName="fChain"` is H2 (oracles disagree). `reason="insufficient_agreement"`
  on `objectName="fChain"` together with `fchain_read_errors` spiking is H1 (home-chain read outage).
  `reason="threshold_not_defined"` is a config/data mismatch independent of H1/H2.

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
