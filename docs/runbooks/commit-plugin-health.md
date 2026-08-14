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
  - docs/metrics/reader-metrics.md         # data-source (reader + config poller) metrics & gaps
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

Every check below is independent — run all of them, don't stop early. Each has a fixed
severity mapping, not a branch.

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

The `data_source` group (reader + config poller) breaks both conventions above: those metrics are
shared with **execute**, carry **no** dest `chainID` label, and key on `chain` = the **numeric chain
selector** plus `query`/`kind`/`state`. Filter them by the metric's own labels, not by
`source_network_name`/`chainID`. They still inherit `node_id` + `csa_public_key` from the beholder
client (like every commit series), so per-node vs DON-wide grouping works the same way.

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

```yaml
checks:
  # --- liveness ---
  - id: heartbeat_observation
    group: liveness
    always_emitted: true
    query: 'sum(rate(ccip_commit_plugin_heartbeat_total{chainID=~"$destChain", phase="observation"}[1m]))'
    severity: {crit_if: "result == 0", ok_if: "result > 0"}
    owner: ccip-commit-oncall
    note: "unconditional per-round liveness signal for Observation(); 0 or empty means every other check below may simply be stale, not bad. Window is deliberately 1m, not 5m -- rate() extrapolates across its whole window, so a plugin that died N minutes ago still reads as alive for up to [window] more minutes; confirmed empirically (see live_oracle_count's note) that a 5m window delays detecting a real outage by up to 5 minutes. Don't widen this window without re-checking that tradeoff"

  - id: heartbeat_outcome
    group: liveness
    always_emitted: true
    query: 'sum(rate(ccip_commit_plugin_heartbeat_total{chainID=~"$destChain", phase="outcome"}[1m]))'
    severity: {crit_if: "result == 0", ok_if: "result > 0"}
    owner: ccip-commit-oncall
    note: "same signal for Outcome(); flat while observation is healthy means OCR is failing to schedule/deliver to the outcome stage. Same 1m-window rationale as heartbeat_observation"

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

  - id: consensus_dropped
    group: cursing_consensus
    always_emitted: false
    query: 'sum(rate(ccip_commit_consensus_dropped_total{chainID=~"$destChain"}[5m])) by (objectName, reason)'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: ccip-commit-oncall
    note: "per-key consensus drop breakdown. reason='split' on objectName='fChain' is H2 (oracles disagree). reason='insufficient_agreement' on objectName='fChain' is H1 when fchain_read_errors is also spiking (too few oracles could report). reason='threshold_not_defined' is a config/data mismatch. For chain-keyed objectNames (MerkleRoot, OnRampMaxSeqNums, etc.) the metric also carries source_network_name, so you can drill down to the affected lane. Empty result is OK, not UNKNOWN"

  - id: live_oracle_count
    group: cursing_consensus
    always_emitted: true
    query: 'count(count by (csa_public_key) (timestamp(ccip_commit_plugin_heartbeat_total{chainID=~"$destChain", phase="observation"}) > time() - 60))'
    severity: {crit_if: "$fRoleDON known and result < 2*$fRoleDON+1", warn_if: "$fRoleDON known and 2*$fRoleDON+1 <= result < 3*$fRoleDON+1", ok_if: "$fRoleDON known and result >= 3*$fRoleDON+1", info: "$fRoleDON not supplied -- report the raw count, no verdict"}
    owner: ccip-commit-oncall
    note: "csa_public_key uniquely identifies each node in every commit metric series (confirmed empirically against a live devenv, not just inferred from source). This deliberately uses timestamp()-vs-time(), not rate()>0: confirmed by live testing that rate([5m])>0 stays true for up to 5 minutes after a node actually stops, because rate() extrapolates across its whole window using the oldest and newest samples it finds -- a node that died 3 minutes ago still reports a positive rate at the 5m window size, making the headcount silently wrong for the first several minutes of a real outage. timestamp() answers 'did this series get a sample in the last 60s' directly, with no extrapolation lag. This is the ONLY check in this file that can detect 'not enough oracles are even running' -- fchain_read_errors below only reports a *surviving* oracle's own read failures, it has no visibility into oracles that never started. If this comes back CRIT, it likely explains consensus_observation_failed on its own; don't let fchain_read_errors being flat push you toward blaming a split vote instead (see the aggregation rule)"

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

  # --- data source (reader + config poller) ---
  # These are shared with execute and live in the reader/config-poller layers (see
  # docs/metrics/reader-metrics.md). Unlike the commit metrics above, `chain` is the NUMERIC
  # chain selector (not source_network_name) and there's no dest `chainID` label to filter on;
  # every series carries `node_id` + `csa_public_key` inherited from the beholder client.
  # They answer "is the plugin fine but being fed empty/partial/stale data?" -- upstream of,
  # and a root cause for, several of the outcome checks above.
  - id: reader_read_empty
    group: data_source
    always_emitted: false
    query: 'sum by (query, chain) (rate(ccip_reader_read_empty_total[5m]))'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: chain-infra-oncall
    note: "a read returned nothing with no error -- the false-idle primitive. Empty result here is OK (the event never happened). Correlate `chain` (numeric selector) to a lane and split node-local vs DON-wide by csa_public_key"

  - id: reader_chain_gap
    group: data_source
    always_emitted: false
    query: 'sum by (query, chain, state) (rate(ccip_reader_chain_gap_total[5m]))'
    severity: {warn_if: 'any series with state!="returned" > 0', ok_if: 'all series with state="returned" or empty result'}
    owner: chain-infra-oncall
    note: "per-chain outcome of a chain read -- serves as BOTH the subset flag and its reason. Actionable (non-returned) states: not_found|disabled|misconfigured|error|invalid|missing|stale|no_accessor|config_error|no_native_token|count_mismatch. count_mismatch = a message read came back with a count that doesn't match the requested range (partial/incomplete data). A read returning a subset = any requested chain in a non-returned state. (Consolidated: the separate ccip_reader_read_partial counter was removed; subset detection is chain_gap{state!=\"returned\"}.)"

  - id: config_cache_stale
    group: data_source
    always_emitted: true
    query: 'max by (chain, kind) (ccip_reader_config_cache_age_seconds)'
    severity: {crit_if: 'kind="chain" and result > 90', ok_if: 'all kind="chain" result <= 90', info: 'kind="source" values (report, no verdict)'}
    owner: chain-infra-oncall
    note: "how stale the config-poller cache is. This is the SUSTAINED signal for config-refresh failure: because the code refuses to advance the refresh timestamp on an empty snapshot, the age only climbs across MULTIPLE consecutive empty/failed polls -- a single intermittent empty (flip-flop) is absorbed by the next good poll and keeps age low. CRIT at >90s (3x the 30s refresh period) means the poller genuinely cannot refresh, so every cached read (GetRMNRemoteConfig, GetRmnCurseInfo, config digest, router address) is running on hollow/stale data. Tie to config_cache_overwritten_empty / config_poll_failure to name why"

  - id: config_poller_liveness
    group: data_source
    always_emitted: true
    query: 'min by (chain) (ccip_reader_config_poller_last_success_timestamp)'
    severity: {crit_if: "time() - result > 90", ok_if: "time() - result <= 90"}
    owner: chain-infra-oncall
    note: "last successful background config poll per chain; a wedged/missing poller goroutine makes every cache silently go stale. Uses last-success staleness, not rate(), for the same detection-lag reason as live_oracle_count"

  - id: config_poll_failure
    group: data_source
    always_emitted: false
    query: 'sum by (chain) (rate(ccip_reader_config_poll_failure_total[5m]))'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: chain-infra-oncall
    note: "background config refresh failures -- exactly the per-chain accounting that a warn-only log used to hide (which chain, not just 'somewhere')"

  - id: config_cache_overwritten_empty
    group: data_source
    always_emitted: false
    query: 'sum by (kind) (rate(ccip_reader_config_cache_overwritten_empty_total[5m]))'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: chain-infra-oncall
    note: "accessor returned an EMPTY chain-config snapshot (no-bindings / empty-batch). Per-event this is NOT a page: a single occurrence is benign and the guard refuses to clobber a good cache for it. Only actionable when SUSTAINED -- which is captured by config_cache_stale (age > 90s), not by this counter alone. Use this to explain WHY the cache went stale (empty snapshot) and which chains. WARN + see config_cache_stale for severity"
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

### Data source (reader + config poller)

This group is **upstream of the whole checklist**: it tests whether the commit plugin is healthy
*but being fed empty/partial/stale data*. The `CRIT` here is **`config_cache_stale`** (sustained:
cache age > 90s = the poller genuinely cannot refresh), which is the thing that makes every cached
read (curse, RMN config, config digest, router address) run on hollow data. `config_cache_overwritten_empty`
alone is only `WARN` — a single empty snapshot is benign (the guard refuses to clobber a good cache),
and per-event empties only matter when they're *sustained*, which `config_cache_stale` already
captures. Use the empty counter + `config_poll_failure` to explain *why* the cache went stale and
which chain(s). `config_cache_stale` and `config_poller_liveness` are the two
`always_emitted: true` checks in this group, so an empty result on them is `UNKNOWN` (pipeline
broken), and they're the ones to check before trusting the *cached* inputs behind step4
(cursing) and the config-digest mismatch check above.

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
- **A `data_source` finding firing at the same time as an outcome check (`pending_messages`
  climbing, `reader_read_*` alongside `consensus_dropped`, `config_cache_stale` alongside a curse
  or digest-mismatch reading) is the SAME incident, and the data-source check is the **root
  cause**.** Fold it in as the named cause with the outcome check as the symptom — do not report
  them as two unrelated problems (a plugin fed empty/partial/stale data is not broken; the reader
  layer feeding it is).

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
