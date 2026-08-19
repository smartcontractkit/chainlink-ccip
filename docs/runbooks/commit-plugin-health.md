---
name: commit-plugin-health
description: Point-in-time health check for a commit plugin instance and its merkle root processor. No incident or specific message required.
trigger: "ad hoc / scheduled health check, or as step0 before docs/runbooks/uncommitted-message.md"
severity: informational
owner: ccip-commit-oncall
inputs:
  destChain: {type: string, description: "destChainID label value of the node's destination chain (destChainName/destChainSelector also work as an equivalent filter -- see the label key cheat sheet). May be handed to you as a chain ID, chain selector, or name -- see docs/runbooks/chain-identifiers.md to translate if you need a form other than what you were given"}
  sourceChains: {type: string, description: "sourceChainName regex filter; defaults to all lanes into destChain", default: ".*"}
  fRoleDON: {type: integer, required: false, description: "optional. If known, expected live oracle count for destChain's DON is 3*fRoleDON+1 and the consensus threshold is 2*fRoleDON+1. Omit if uncertain (e.g. local devenvs with a possible bootstrap-node offset) -- live_oracle_count reports the raw count either way, it just can't grade it against a threshold without this."}
related:
  - docs/runbooks/uncommitted-message.md   # incident triage for one specific stuck message
  - docs/metrics/commit-metrics.md         # design rationale + full metric-by-metric findings
  - docs/metrics/reader-metrics.md         # data-source (reader + config poller) metrics & gaps
  - devenv/dashboards/commit-plugin.json   # Grafana rendering of this runbook (Health section)
  - docs/runbooks/chain-identifiers.md     # translating chain ID / chain selector / chain name
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
specific-message incident (its `step0` is a subset of this runbook's `discovery_state` /
`heartbeat_observation` / `heartbeat_outcome` checks).

## For agents

Every check below is independent — run all of them, don't stop early. Each has a fixed
severity mapping, not a branch.

### Label key cheat sheet

**As of the `commit/metrics/prom.go` label overhaul, every commit-plugin metric uses one uniform
chain-identity label set — this replaced the old per-metric `chainID`/`chain_id`/`source_network_name`
inconsistency this section used to document, and it's a breaking rename, not an addition.** Every
check below now carries four destination-chain labels from `destChainAttrs()`:

| label key | meaning |
|---|---|
| `destChainID` | destination chain ID (what `$destChain` filters against) |
| `destChainFamily` | destination chain family (evm, solana, ...) |
| `destChainName` | destination chain name, e.g. `ethereum-mainnet` |
| `destChainSelector` | destination CCIP chain selector, as a string |

...and every per-lane check (everything under `lane_throughput`, plus `source_chain_cursed`,
`report_transmission_gave_up`, and `consensus_dropped` when it carries a source chain) also
carries the source-side equivalent from `sourceChainAttrs()`:

| label key | meaning |
|---|---|
| `sourceChainID` | source chain ID |
| `sourceChainFamily` | source chain family |
| `sourceChainName` | source chain name (this replaces the old `source_network_name`) |
| `sourceChainSelector` | source CCIP chain selector, as a string |

**The old labels are gone, not aliased.** `chainID`, `chain_id`, `chainFamily`, `chain_family`,
`source_network_name`, and `dest_network_name` no longer appear on any Beholder-emitted series in
this file — every query below now filters on `destChainID=~"$destChain"` and, where relevant,
`sourceChainName=~"$sourceChains"`. If you have an old saved query or alert on this metric family,
it now silently returns empty, not an error — re-check anything written before this rename before
trusting a `UNKNOWN`/blank result from it. (The old names do still exist on a separate,
`promauto`-only direct-scrape path for a handful of metrics in `commit/metrics/legacy_prom.go` —
that path is unrelated to what this checklist queries and is flagged deprecated in-code; don't
mix the two.)

Since `destChainName`/`sourceChainName`/`destChainSelector`/`sourceChainSelector` are now
first-class labels (not just something you compute via [`chain-identifiers.md`](chain-identifiers.md)),
you can equally well filter `$destChain`/`$sourceChains` against whichever representation you were
actually handed — `destChainSelector=~"..."` works exactly as well as `destChainID=~"..."` if a
selector is all you have. Translating up front is still worth doing if you'll need multiple forms
over the course of an investigation, but it's no longer strictly required just to run a query.

`csa_public_key` is a different kind of label — it's not used to filter by chain, it's a
per-node identity present on *every* commit metric series (confirmed against a live devenv, not
just inferred). `live_oracle_count` is the only check that groups by it; nothing else in this
checklist needs to.

The `data_source` group (reader + config poller) is a different package (shared with **execute**)
and has its own, separate label overhaul in `pkg/reader/rcmetrics` (a private `chainAttrs()`
helper, same idea as `destChainAttrs()` above but with no dest/source prefix, since each of these
metrics only ever describes one chain, never a lane). Every metric in this group now carries
`chainID`/`chainFamily`/`chainName`/`chainSelector` — this replaced the old bare numeric-selector
`chain` label (`chainSelector` now carries that same numeric value, as a string, if you need it for
`chain-identifiers.md`-style translation). Every series still also carries `node_id` +
`csa_public_key` inherited from the beholder client like every commit series.

**Every metric in this group also carries `destChainID`/`destChainFamily`/`destChainName`/
`destChainSelector`** — a second, dest-scoped `destChainAttrs()` in the same file, resolved once
per `configPollerV2`/`ccipChainReader` instance. This is required, not cosmetic: a node runs one
config-poller/reader instance **per destination chain**, and more than one instance can
independently track the same source chain (confirmed on staging — two destination-chain lanes both
reading Solana as a source chain collided into one series per `(chainID, node_id)` without this,
`max()`/`sum()` across them silently mixing a healthy lane's fast-refreshing cache with a
different, broken lane's stuck one). Filter/group by `destChainID=~"$destChain"` on every
`data_source` check below, the same way you'd filter a commit-plugin check.

**One exception inside the `data_source` group itself:** `ccip_reader_read_outcome` is recorded at
the `observedCCIPReader` wrapper level (every reader call, not just the per-source-chain ones), so
its `chainID` (and the rest of its chain-identity labels) match the **destination** chain, not
"whichever chain this specific read is about" like every other metric in this group. Before this
turn's rename, that distinction was visible in the label *name itself* (`chainID` vs. the rest of
the group's bare `chain`); now that every metric in this group shares identical label keys, the
distinction is purely semantic and easy to miss — don't assume `reader_read_outcome_error`'s
`chainID` means the same thing `reader_read_empty`'s or `reader_chain_gap`'s does. Filter
`reader_read_outcome_error` with `chainID=~"$destChain"`, same as the destination-chain filter used
throughout the rest of this doc. (Its `destChainID` label carries the exact same value as `chainID`
here — always the same destination chain for this one metric — so either filter works; the other
`data_source` checks need `destChainID` specifically because their `chainID` means something else.)

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
  - id: discovery_state
    group: liveness
    always_emitted: true
    query: 'max by (destChainID) (ccip_commit_discovery_state{destChainID=~"$destChain"})'
    severity: {crit_if: "result == 0", ok_if: "result == 1"}
    owner: ccip-commit-oncall
    note: "per-round readiness gauge, recorded from `commit/plugin.go`'s Observation() via `contractsInitialized.Load()` and reported by `TrackDiscoveryState` in `commit/metrics`. 0 = the plugin is still in its contract-discovery phase, 1 = discovery complete. This is checked FIRST (it gates the whole run): a plugin that has not discovered its contracts is not healthy and cannot produce any observations, no matter how many rounds its heartbeat reports — the classic manifestation is a single unhealthy RPC leaving every committed RPC stuck in discovery while heartbeats keep ticking (see the commit that introduced the metric). Crit is `== 0` across every node: use max() so that only when NO node has exited discovery does this fire as not-ready; if max()==0, treat it as the root cause before reading any observation/outcome metric below, whose values are then suspect because the plugin was never ready to emit them. Aggregate rather than reporting per-node: if one node is stuck in discovery while its peers are not, that is a single-node finding (its observations can't reflect the lane), captured via `by (csa_public_key)` — not a DON-wide verdict"

  - id: heartbeat_observation
    group: liveness
    always_emitted: true
    query: 'sum(rate(ccip_commit_plugin_heartbeat_total{destChainID=~"$destChain", phase="observation"}[1m]))'
    severity: {crit_if: "result == 0", ok_if: "result > 0"}
    owner: ccip-commit-oncall
    note: "unconditional per-round liveness signal for Observation(); 0 or empty means every other check below may simply be stale, not bad. Window is deliberately 1m, not 5m -- rate() extrapolates across its whole window, so a plugin that died N minutes ago still reads as alive for up to [window] more minutes; confirmed empirically (see live_oracle_count's note) that a 5m window delays detecting a real outage by up to 5 minutes. Don't widen this window without re-checking that tradeoff. The MAGNITUDE of a nonzero result is not meaningful -- this is rate(counter[1m]) on a counter incremented once per OCR round, so its steady-state value is a function of this DON's round cadence (itself a function of config, chain, and load), not a fixed target. A value like 0.5 is not 'half alive'; it just means roughly one round every 2s over the window. Don't compare this number across chains/environments or expect a specific magnitude -- the only threshold that means anything is exactly 0 (dead) vs strictly >0 (alive). Use it to answer 'is it running at all', and use pending_messages/lane-specific checks below to answer 'is it keeping up'"

  - id: heartbeat_outcome
    group: liveness
    always_emitted: true
    query: 'sum(rate(ccip_commit_plugin_heartbeat_total{destChainID=~"$destChain", phase="outcome"}[1m]))'
    severity: {crit_if: "result == 0", ok_if: "result > 0"}
    owner: ccip-commit-oncall
    note: "same signal for Outcome(); flat while observation is healthy means OCR is failing to schedule/deliver to the outcome stage. Same 1m-window rationale as heartbeat_observation, and the same 'magnitude isn't meaningful, only zero-vs-nonzero is' caveat -- see heartbeat_observation's note"

  - id: config_digest_mismatch
    group: liveness
    always_emitted: true
    query: 'max(ccip_commit_config_digest_mismatch{destChainID=~"$destChain"})'
    severity: {crit_if: "result == 1", ok_if: "result == 0"}
    owner: ccip-commit-oncall
    note: "home-chain config digest differs from the offramp's; root cause for a wide range of downstream rejections. If $destChain can match more than one chain (a multi-value/'All' selection), do NOT collapse with the bare max() above for reporting -- that only tells you at least one chain mismatches, not which. Use 'max by (destChainID) (ccip_commit_config_digest_mismatch{destChainID=~\"$destChain\"}) == 1' instead (or the equivalent table view) to list only the mismatched destChainID(s); report 'NO MISMATCH' rather than a blank/zero table when that query returns no rows, since an empty result here is the OK state, not UNKNOWN (always_emitted: true means the underlying gauge exists per chain, but a query filtered to ==1 legitimately returns nothing when nothing mismatches)"

  - id: processor_errors
    group: liveness
    always_emitted: false
    query: 'sum(rate(ccip_commit_processor_errors_total{destChainID=~"$destChain"}[15m])) by (processor, method)'
    fallback_query: 'sum(rate(ccip_commit_processor_errors_total{chainID=~"$destChain"}[15m])) by (processor, method)'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: ccip-commit-oncall
    note: "generic TrackedProcessor error counter; empty means zero errors, not unknown -- EXCEPT this metric is a live straggler bug, confirmed by reading `commit/metrics/prom.go`: it's registered as a Beholder Int64Counter (`bhProcessorErrors`) in NewPromReporter, but TrackProcessorLatency's error branch only ever calls the promauto counter (`p.processorErrors...Inc()`) -- `p.bhProcessorErrors` is never invoked anywhere in the file, despite the metric being labeled `dual` in the metric reference. On a datasource fed purely from the Beholder/OTel pipeline (this metric never leaves the node that way), the primary query above will ALWAYS return empty and read as OK regardless of real errors -- a silent false negative, not a benign empty result like the other checks here. If your datasource also scrapes the promauto endpoint directly, use fallback_query instead (old labels: `chainFamily,chainID`, not `destChainID`). Otherwise treat this check as unreliable until `bhProcessorErrors` actually gets called in code, and say so explicitly rather than reporting a clean bill of health. Does not include swallowed per-goroutine failures (see merkleroot_observation_errors in uncommitted-message.md)"

  - id: processor_latency_p95
    group: liveness
    always_emitted: true
    query: 'histogram_quantile(0.95, sum(rate(ccip_commit_processor_latency_bucket{destChainID=~"$destChain"}[15m])) by (le, processor, method))'
    severity: {info: true}
    owner: ccip-commit-oncall
    note: "report raw value only, this is a single point-in-time sample with nothing to trend against -- do not infer a trend from one run. If a value lands exactly on a bucket boundary (e.g. pinned at the histogram's max bucket), report that fact explicitly as it likely means real latency exceeds the highest defined bucket; still do not elevate this to WARN/CRIT, there is no defined SLO. UNIT WARNING, confirmed by live testing on staging: this histogram's buckets and recorded values are in NANOSECONDS (`commit/metrics/prom.go`'s promProcessorLatencyHistogram buckets are literal time.Duration constants, e.g. 20*time.Second == 2e10, and TrackProcessorLatency calls .Observe(float64(latency)) directly on the nanosecond-valued time.Duration) -- the histogram itself is internally consistent, but if you build a Grafana panel (or any display) with unit='s' on this query's raw result, a real value like 2e10 (20 real seconds) renders as 2e10 SECONDS, i.e. roughly 634 years. Either divide the query result by 1e9 to report actual seconds, or use a unit-aware display set to nanoseconds ('ns' in Grafana, which auto-scales to us/ms/s) -- do not display the raw number with a seconds unit"

  # --- lane throughput & backlog (per source chain) ---
  - id: pending_messages
    group: lane_throughput
    always_emitted: true
    query: 'max by (sourceChainName) (ccip_commit_pending_messages{sourceChainName=~"$sourceChains", destChainID=~"$destChain"})'
    severity: {info: true}
    owner: ccip-commit-oncall
    note: "onRampMaxSeqNum - offRampNextSeqNum; no universal healthy value, watch for sustained growth not a single sample. Empty result here (no lanes matched) is UNKNOWN, not zero backlog -- this metric is always emitted per active lane"

  - id: range_truncated
    group: lane_throughput
    always_emitted: false
    query: 'sum(rate(ccip_commit_range_truncated_total{sourceChainName=~"$sourceChains", destChainID=~"$destChain"}[15m])) by (sourceChainName)'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: ccip-commit-oncall
    note: "chain-throughput ceiling (MaxMerkleTreeSize) being hit"

  - id: seqnum_invariant_violation
    group: lane_throughput
    always_emitted: false
    query: 'sum(rate(ccip_commit_seqnum_invariant_violation_total{sourceChainName=~"$sourceChains", destChainID=~"$destChain"}[15m])) by (sourceChainName, type)'
    severity: {crit_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: ccip-commit-oncall
    note: "offramp_ahead_of_onramp / onramp_max_zero / offramp_seqnum_regression -- documented in-code as stall signals, not routine noise"

  - id: offramp_consensus_insufficient
    group: lane_throughput
    always_emitted: false
    query: 'sum(rate(ccip_commit_offramp_consensus_insufficient_total{sourceChainName=~"$sourceChains", destChainID=~"$destChain"}[15m])) by (sourceChainName)'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: ccip-commit-oncall
    note: "DON couldn't agree on OffRampNextSeqNum for this chain; looks identical to \"no new messages\" without this"

  - id: offramp_lane_status
    group: lane_throughput
    always_emitted: true
    query: 'max by (sourceChainName, status) (ccip_commit_offramp_lane_status{sourceChainName=~"$sourceChains", destChainID=~"$destChain"})'
    severity: {crit_if: 'status="rmn_misconfigured" and result == 1', info_if: 'status=~"skipped_.*" and result == 1', ok_if: 'status="live" and result == 1'}
    owner: ccip-commit-oncall
    note: "rmn_misconfigured flags real config drift (a lane still expecting RMN blessing while RMN is globally off); skipped_* are expected states, not unhealthy. Reported for all four statuses every round by design, so empty result is UNKNOWN, not \"no active lanes\""

  # --- cursing & consensus ---
  - id: rmn_curse_active
    group: cursing_consensus
    always_emitted: true
    query: 'max(ccip_commit_rmn_curse_active{destChainID=~"$destChain", curse_type=~"global|destination"}) by (curse_type)'
    severity: {warn_if: "any series == 1", ok_if: "all series == 0"}
    owner: curse-owner
    note: "halts ALL reporting for the dest chain while active; often intentional/incident-flagged (see uncommitted-message.md step4) -- WARN not CRIT because an active curse is frequently expected, not a plugin bug"

  - id: source_chain_cursed
    group: cursing_consensus
    always_emitted: true
    query: 'max by (sourceChainName) (ccip_commit_source_chain_cursed{sourceChainName=~"$sourceChains", destChainID=~"$destChain"})'
    severity: {warn_if: "any series == 1", ok_if: "all series == 0"}
    owner: curse-owner
    note: "per-lane curse; same intentional-vs-bug ambiguity as rmn_curse_active"

  - id: consensus_observation_failed
    group: cursing_consensus
    always_emitted: false
    query: 'sum(rate(ccip_commit_consensus_observation_failed_total{destChainID=~"$destChain"}[5m]))'
    severity: {crit_if: "result > 0", ok_if: "result == 0 (including empty result)"}
    owner: ccip-commit-oncall
    note: "the only way a whole round fails: DON can't reach 2*fRoleDON+1 agreement on FChain for the dest chain. Destination-chain-wide by construction, not lane-specific"

  - id: consensus_dropped
    group: cursing_consensus
    always_emitted: false
    query: 'sum(rate(ccip_commit_consensus_dropped_total{destChainID=~"$destChain", objectName!="RMNRemoteConfig"}[5m])) by (objectName, reason)'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: ccip-commit-oncall
    note: "per-key consensus drop breakdown. reason='split' on objectName='fChain' is H2 (oracles disagree). reason='insufficient_agreement' on objectName='fChain' is H1 when fchain_read_errors is also spiking (too few oracles could report). reason='threshold_not_defined' is a config/data mismatch. For chain-keyed objectNames (MerkleRoot, OnRampMaxSeqNums, etc.) the metric also carries sourceChainName, so you can drill down to the affected lane. objectName='RMNRemoteConfig' is excluded from the query entirely (not just downgraded): with RMN disabled in prod, its remote config is permanently empty, so reason='insufficient_agreement' for it fires constantly and is pure expected noise, never a real finding -- if you need to confirm RMN's config state directly for some other reason, query ccip_commit_consensus_dropped_total{objectName=\"RMNRemoteConfig\"} without this exclusion. Empty result on the excluded-noise-free query is OK, not UNKNOWN"

  - id: live_oracle_count
    group: cursing_consensus
    always_emitted: true
    query: 'count(count by (csa_public_key) (timestamp(ccip_commit_plugin_heartbeat_total{destChainID=~"$destChain", phase="observation"}) > time() - 60))'
    severity: {crit_if: "$fRoleDON known and result < 2*$fRoleDON+1", warn_if: "$fRoleDON known and 2*$fRoleDON+1 <= result < 3*$fRoleDON+1", ok_if: "$fRoleDON known and result >= 3*$fRoleDON+1", info: "$fRoleDON not supplied -- report the raw count, no verdict"}
    owner: ccip-commit-oncall
    note: "csa_public_key uniquely identifies each node in every commit metric series (confirmed empirically against a live devenv, not just inferred from source). This deliberately uses timestamp()-vs-time(), not rate()>0: confirmed by live testing that rate([5m])>0 stays true for up to 5 minutes after a node actually stops, because rate() extrapolates across its whole window using the oldest and newest samples it finds -- a node that died 3 minutes ago still reports a positive rate at the 5m window size, making the headcount silently wrong for the first several minutes of a real outage. timestamp() answers 'did this series get a sample in the last 60s' directly, with no extrapolation lag. This is the ONLY check in this file that can detect 'not enough oracles are even running' -- fchain_read_errors below only reports a *surviving* oracle's own read failures, it has no visibility into oracles that never started. If this comes back CRIT, it likely explains consensus_observation_failed on its own; don't let fchain_read_errors being flat push you toward blaming a split vote instead (see the aggregation rule)"

  - id: fchain_read_errors
    group: cursing_consensus
    always_emitted: false
    query: 'sum(rate(ccip_commit_fchain_read_errors_total{destChainID=~"$destChain"}[5m]))'
    severity: {warn_if: "result > 0", ok_if: "result == 0 (including empty result)"}
    owner: home-chain-infra-oncall
    note: "if this AND consensus_observation_failed are both firing, treat the combination as one CRIT finding, not two (see aggregation rule) -- likely H1, home-chain RPC/read outage"

  # --- report transmission ---
  - id: report_validation_rejected
    group: report_transmission
    always_emitted: false
    query: 'sum(rate(ccip_commit_report_validation_rejected_total{destChainID=~"$destChain"}[15m])) by (phase, reason)'
    severity: {warn_if: 'any series with reason!="stale" > 0', info_if: 'only reason="stale" > 0', ok_if: "empty result, or all series == 0"}
    owner: ccip-commit-oncall
    note: "reason=\"stale\" is usually benign (an overlapping report already landed); other reasons are worth a look"

  - id: report_transmission_gave_up
    group: report_transmission
    always_emitted: false
    query: 'sum(rate(ccip_commit_report_transmission_gave_up_total{sourceChainName=~"$sourceChains", destChainID=~"$destChain"}[15m])) by (sourceChainName)'
    severity: {crit_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: ccip-commit-oncall
    note: "one of the most common on-call pages historically; previously a Warnw with no counter at all. If CRIT, uncommitted-message.md Scenario 2 may name a different owner (chain-b-txm-oncall) once root-caused -- this check alone doesn't know which"

  - id: report_transmission_attempts_p95
    group: report_transmission
    always_emitted: false
    query: 'histogram_quantile(0.95, sum(rate(ccip_commit_report_transmission_attempts_bucket{destChainID=~"$destChain"}[15m])) by (le, success))'
    severity: {info: true}
    owner: ccip-commit-oncall
    note: "recorded once per transmission-check attempt; empty result means zero transmission cycles in the window, which is itself informational (report as INFO with value \"no data\"), not UNKNOWN. Rising attempts alongside report_transmission_gave_up firing is the actionable combination, not this alone"

  # --- data source (reader + config poller) ---
  # These are shared with execute and live in the reader/config-poller layers (see
  # docs/metrics/reader-metrics.md). As of pkg/reader/rcmetrics's chainAttrs() rollout, every
  # metric below carries chainID/chainFamily/chainName/chainSelector (same 4-label pattern as
  # commit/metrics/prom.go's destChainAttrs(), just with no dest/source prefix -- each of these
  # metrics only ever describes one chain, never a lane). This replaced the old bare numeric-
  # selector `chain` label; chainSelector now carries that same numeric value if you need it.
  # EVERY metric below ALSO now carries destChainID/destChainFamily/destChainName/destChainSelector
  # (a second, dest-scoped chainAttrs() resolved once per configPollerV2/ccipChainReader instance)
  # -- confirmed necessary by a real staging incident: a node runs one config-poller instance PER
  # DESTINATION CHAIN, and two lanes sharing the same source chain (e.g. two dest chains both
  # reading Solana) collide into one series per (chainID, node_id) without it -- max()/sum() across
  # them silently mixed a healthy lane's fast-refreshing cache with a different, broken lane's
  # stuck one. Filter every query below by destChainID=~"$destChain" too, not just node/chain.
  # Every series still also carries `node_id` + `csa_public_key` inherited from the beholder client.
  # They answer "is the plugin fine but being fed empty/partial/stale data?" -- upstream of,
  # and a root cause for, several of the outcome checks above.
  - id: reader_read_outcome_error
    group: data_source
    always_emitted: false
    query: 'sum by (query) (rate(ccip_reader_read_outcome_total{chainID=~"$destChain", outcome="error"}[5m]))'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: chain-infra-oncall
    note: "the consolidated ok/empty/error classification recorded at the observedCCIPReader wrapper for every reader call. This is the ERROR leg specifically -- a query that returned a hard error, not just an empty/partial result (that's reader_read_empty/reader_chain_gap below). Empty result here is OK (never happened). chainID here is the DESTINATION chain (this reader instance's) -- for THIS one metric chainID and destChainID are always the same value (redundant but harmless), so filtering by chainID=~\"$destChain\" alone is already lane-scoped correctly; you don't need destChainID here the way the checks below do. Split node-local vs DON-wide by csa_public_key before escalating"

  - id: reader_read_empty
    group: data_source
    always_emitted: false
    query: 'sum by (query, chainID) (rate(ccip_reader_read_empty_total{destChainID=~"$destChain"}[5m]))'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: chain-infra-oncall
    note: "a read returned nothing with no error -- the false-idle primitive. Empty result here is OK (the event never happened). chainID here is whichever chain that specific read was about (source or dest, depending on the query) -- NOT the destination-chain convention reader_read_outcome_error uses above; destChainID=~\"$destChain\" is what actually scopes this to the lane under investigation. chainSelector (also on this series) carries the raw numeric selector if you need it for chain-identifiers.md-style translation. Split node-local vs DON-wide by csa_public_key"

  - id: reader_chain_gap
    group: data_source
    always_emitted: false
    query: 'sum by (query, chainID, state) (rate(ccip_reader_chain_gap_total{destChainID=~"$destChain"}[5m]))'
    severity: {warn_if: 'any series with state!="returned" > 0', ok_if: 'all series with state="returned" or empty result'}
    owner: chain-infra-oncall
    note: "per-chain outcome of a chain read -- serves as BOTH the subset flag and its reason. Actionable (non-returned) states: not_found|disabled|misconfigured|error|invalid|missing|stale|no_accessor|config_error|no_native_token|count_mismatch. count_mismatch = a message read came back with a count that doesn't match the requested range (partial/incomplete data). A read returning a subset = any requested chain in a non-returned state. (Consolidated: the separate ccip_reader_read_partial counter was removed; subset detection is chain_gap{state!=\"returned\"}.) chainID is whichever chain the read was about, same convention as reader_read_empty; destChainID scopes to the lane under investigation."

  - id: config_cache_stale
    group: data_source
    always_emitted: true
    query: 'max by (chainID, kind) (ccip_reader_config_cache_age_seconds{destChainID=~"$destChain"})'
    severity: {crit_if: 'kind="chain" and result > 90', ok_if: 'all kind="chain" result <= 90', info: 'kind="source" values (report, no verdict)'}
    owner: chain-infra-oncall
    note: "how stale the config-poller cache is. This is the SUSTAINED signal for config-refresh failure: because the code refuses to advance the refresh timestamp on an empty snapshot, the age only climbs across MULTIPLE consecutive empty/failed polls -- a single intermittent empty (flip-flop) is absorbed by the next good poll and keeps age low. CRIT at >90s (3x the 30s refresh period) means the poller genuinely cannot refresh, so every cached read (GetRMNRemoteConfig, GetRmnCurseInfo, config digest, router address) is running on hollow/stale data. Tie to config_cache_overwritten_empty / config_poll_failure to name why. destChainID=~\"$destChain\" is REQUIRED here, not optional: confirmed on staging that omitting it lets a chain shared across multiple destination-chain lanes (e.g. Solana as a source into two different dest chains) silently max() together one healthy lane's fast-refreshing cache with a different, broken lane's cache stuck for hours -- a query without destChainID can look internally inconsistent against config_poll_failure below purely from this aggregation collapse, not because either metric is wrong"

  - id: config_poller_liveness
    group: data_source
    always_emitted: true
    query: 'min by (chainID) (ccip_reader_config_poller_last_success_timestamp{destChainID=~"$destChain"})'
    severity: {crit_if: "time() - result > 90", ok_if: "time() - result <= 90"}
    owner: chain-infra-oncall
    note: "last successful background config poll per chain; a wedged/missing poller goroutine makes every cache silently go stale. Uses last-success staleness, not rate(), for the same detection-lag reason as live_oracle_count. destChainID scopes this to the lane under investigation, same reasoning as config_cache_stale above"

  - id: config_poll_failure
    group: data_source
    always_emitted: false
    query: 'sum by (chainID, reason) (rate(ccip_reader_config_poll_failure_total{destChainID=~"$destChain"}[5m]))'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: chain-infra-oncall
    note: "background config refresh failures -- exactly the per-chain accounting that a warn-only log used to hide (which chain, not just 'somewhere'). reason=no_chain_accessor|batch_fetch_failed|no_chain_cache names WHY without needing to correlate against logs: no_chain_accessor = misconfiguration (chain has no accessor wired up at all); batch_fetch_failed = the accessor's GetAllConfigsLegacy call itself errored (RPC/contract-binding problem -- this is the one that fired during the Solana 'no live nodes available' staging incident, see config_digest_mismatch above); no_chain_cache = internal cache-creation failure, should basically never fire. Recorded at the single choke point (batchRefreshChainAndSourceConfigs) so on-demand cache-miss fetches count the same as background-refresh ones, not just the background ticker's failures. destChainID=~\"$destChain\" is REQUIRED here too: without it, a rate summed across every destination-chain lane sharing this source chain can look deceptively moderate even when ONE specific lane's poller is failing nearly every single cycle -- always scope this to the lane you're actually investigating before judging the magnitude against config_cache_stale"

  - id: config_cache_overwritten_empty
    group: data_source
    always_emitted: false
    query: 'sum by (chainID, kind) (rate(ccip_reader_config_cache_overwritten_empty_total{destChainID=~"$destChain"}[5m]))'
    severity: {warn_if: "any series > 0", ok_if: "all series == 0 (including empty result)"}
    owner: chain-infra-oncall
    note: "accessor returned an EMPTY chain-config snapshot (no-bindings / empty-batch). Per-event this is NOT a page: a single occurrence is benign and the guard refuses to clobber a good cache for it. Only actionable when SUSTAINED -- which is captured by config_cache_stale (age > 90s), not by this counter alone. Use this to explain WHY the cache went stale (empty snapshot) and which chains -- the chainID label was added alongside chainAttrs(); before that this metric had no chain identity at all, even though the chain being refreshed was always known at the call site. If this returns empty while config_cache_stale is climbing for the same destChainID/chainID, that RULES OUT the empty-snapshot path as the cause -- look at config_poll_failure (a real accessor error) instead. In dev/no-binding environments this can fire chronically as expected noise (unbound RMN/optional contracts); treat it as a signal only when accompanied by config_cache_stale climbing. WARN + see config_cache_stale for severity"
```

## Groups

### Liveness

`discovery_state` first — it gates the whole run: if the plugin has *not* discovered its
contracts it is not healthy and cannot produce observations, so a `CRIT` here is not just a
symptom but the reason every other observation/outcome check in this run is suspect. Then
`heartbeat_observation` / `heartbeat_outcome` — if either is `CRIT` or `UNKNOWN`, every
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
mismatch. The query excludes `objectName="RMNRemoteConfig"` outright: on an **RMN-disabled**
deployment its remote config is permanently empty, so `reason="insufficient_agreement"` for it
fires constantly and is pure noise, never a real finding — it's filtered out of the check rather
than merely downgraded, since it isn't a signal worth surfacing at all while RMN stays disabled.

### Data source (reader + config poller)

This group is **upstream of the whole checklist**: it tests whether the commit plugin is healthy
*but being fed empty/partial/stale data*. `reader_read_outcome_error` is the consolidated
error-leg signal (a query failed outright) and is complementary to `reader_read_empty`/
`reader_chain_gap` below it (which cover the empty/partial leg) — a query can show up in one
without the other, check both. The `CRIT` here is **`config_cache_stale`** (sustained:
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
  `objectName="RMNRemoteConfig"` never appears here at all — the check's query excludes it, since
  with RMN disabled it's chronic expected noise, not a finding (if you ever need its raw state,
  see the note on `consensus_dropped`'s query for the unfiltered version).
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
                                # separate series with "; ". E.g. a `by (sourceChainName,
                                # status)` result becomes
                                # "sourceChainName=chain-a,status=live:1; sourceChainName=chain-a,status=rmn_misconfigured:0; ...".
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
