## Metrics coverage assessment — commit plugin

I read through `commit/metrics`, `commit/merkleroot` (+ `rmn`), `commit/chainfee`, `commit/tokenprice`, and the top-level `commit/plugin.go` / `report.go` / `factory.go`. There's a real gap, and it's structural, not just a few missing gauges.

**Ground rule for everything below: the presence of a log line — even a "stable identifier" some log-analysis tooling already keys on — is never a reason to exclude a metric.** Logs are exactly the noisy, low-signal, un-queryable thing motivating this work; a few items here are genuinely *only* visible at `Debugw`, which is worse than a "stable" log, not better. Nothing was cut from the list below because it was already logged.

### What exists today (`commit/metrics/prom.go`)

- Generic per-processor instrumentation via `TrackedProcessor` (wraps merkleroot/chainfee/tokenprice): latency histogram, error counter, and an **output-size counter** driven by each type's `Stats()` — but `Stats()` only returns `len(map)`-style item counts (e.g. `gasPrices: 3`), never the actual values.
- `ccip_commit_max_sequence_number` gauge — but only fed from `MerkleRootChain.SeqNumsRange.End()` in `TrackObservation`/`TrackOutcome`, i.e. **only when a root is actually being built for a report**.
- `ccip_commit_latest_round_id`, `ccip_commit_config_digest_mismatch`, `ccip_commit_loopp_ccip_provider_supported`.
- RMN-specific: `TrackRmnReport` (build latency/success) and `TrackRmnRequest` (per-node request latency/error) — merkleroot and its `rmn` subpackage are the only parts of the plugin with a dedicated `MetricsReporter` sub-interface.
- **`chainfee` and `tokenprice` have no dedicated metrics interface at all** — they only get the generic latency/error/size instrumentation above. Token prices, gas prices, and their staleness/deviation/failure states are otherwise 100% log-only.

### Your example, confirmed

`merkleroot.Observation` has `OnRampMaxSeqNums` and `OffRampNextSeqNums` fields (`commit/merkleroot/types.go:34-35`) populated **every round**, but no `Reporter` method ever reads them — the only sequence-number gauge (`ccip_commit_max_sequence_number`) comes from `MerkleRootChain.SeqNumsRange.End()`, which only exists once a root is already being built. So if the plugin is stuck reading a chain (RPC lag, stuck in `selectingRangesForReport`), there's no gauge showing it — the one gauge that looks related lags or is absent during exactly the failure mode you'd want visibility into.

### Cross-cutting themes (repeat in every processor)

Merkleroot, chainfee, and tokenprice each independently:
1. Compute a **per-item value** (seq num / fee / price) that's never gauged — only counted.
2. Compute **staleness** (age since last update) to decide on a heartbeat write, then discard the number and keep only the boolean.
3. Swallow **fetch/RPC failures** into empty maps with an `Errorw` — these never increment `ccip_commit_processor_errors` because the processor's top-level `Observation()` still returns `nil` error.
4. Drop entries during **consensus aggregation** silently — `internal/plugincommon/consensus/consensus.go:42,52` literally has `// TODO: metrics` next to the exact code path that discards a chain/token when too few nodes agree. This is shared infra, hit by every processor, and the team already flagged it (reviewed fix shape below).
5. `internal/libs/asynclib/goroutines.go` (used by chainfee's async observer, similar pattern in tokenprice) has **zero logging or metrics** when an operation times out — it just drops the key from the results map and the caller's blind type-assertion silently degrades to an empty value. This is the most severe blind spot found: a hung RPC call produces no log line and no metric anywhere.
6. **No plugin-level heartbeat/liveness signal.** Surfaced by war-gaming an incident end-to-end rather than by reading the code in isolation: if the whole round pipeline wedges or the process crashes, every metric in this doc simply stops updating instead of showing a bad value — and "gauge went flat" reads very differently in monitoring than "gauge shows a bad value," easy to miss unless dashboards specifically alert on staleness/absence rather than thresholds. There's no dedicated "is the plugin alive at all" signal today; see below.

---

## Full findings

Severity is about production-triage value, independent of whether something is logged today.

### Top-level plugin orchestration (`commit/plugin.go`, `commit/report.go`)

These aren't owned by any single processor — they're the outer round loop and the report-acceptance path — but two gaps here turned out to matter as much as anything processor-specific once an actual incident was traced end-to-end.

**High**
- **No heartbeat / liveness signal.** `commit/plugin.go:434` (`Outcome()`). If the whole round pipeline wedges (deadlock, a panic-recovery masking a crash loop, OCR itself failing to schedule rounds), every metric in this doc simply stops updating instead of showing a bad value. A stale gauge and a valid-but-bad gauge look identical unless you specifically alert on absence/staleness — easy to forget when every other alert here is threshold-based. This should be step 0 of any triage tree, and today it isn't answerable from commit-plugin metrics at all. → Gauge/counter, e.g. `ccip_commit_last_successful_round_timestamp`, set unconditionally near the top of `Outcome()` (and ideally `Observation()` too) regardless of what happens downstream.
- **Report-acceptance / transmission-validation rejections have zero metrics.** `commit/report.go:126-247` (`validateReport`, called from both `ShouldAcceptAttestedReport` and `ShouldTransmitAcceptedReport`) has eight distinct rejection paths — empty/invalid merkle root, duplicate RMN signature, insufficient RMN signatures, cursed report, dest chain unsupported, config digest mismatch, stale report, root-blessing mismatch — every one of them `Warnw`/`Errorw`/`Debugf` only. This sits directly upstream of merkleroot's own `ReportTransmissionGaveUp` counter: if every oracle's local `validateReport` call rejects a report before anyone ever calls transmit, "gave up" fires with **zero on-chain trace at all** (no tx ever submitted) — a completely different remediation path than "tx submitted but reverted" or "tx stuck in mempool." Without per-reason counts here, those scenarios are indistinguishable from `ccip_commit_report_transmission_gave_up_total` alone. → Counter `ccip_commit_report_validation_rejected_total{phase="should_accept"|"should_transmit", reason="empty_root|invalid_seqnum_range|duplicate_rmn_sig|insufficient_rmn_sigs|cursed|dest_not_supported|config_digest_mismatch|stale|root_blessing_mismatch"}`.

### Shared: consensus aggregation (`internal/plugincommon/consensus/consensus.go`)

Used by commit (`chainfee`, `merkleroot`, `merkleroot/rmn/controller.go`, `plugin.go`, `report.go`), `execute/plugin_functions.go`, **and** `internal/plugincommon/discovery/processor.go` — this is genuinely shared infra, not commit-specific, which shapes the recommended fix below. Reviewed design, not yet implemented.

**High**
- **`GetConsensusMap` collapses three distinguishable outcomes into one `Debugw`/`Warnw` line, both marked `// TODO: metrics`** (`consensus.go:34-56`). `NewMinObservation.GetValid()` (`min_observation.go:57-66`) buckets observations by identical value and returns every distinct value that independently cleared the threshold, which means there are three cases today, not two:
  1. `minObs.Get(key)` returns `!exists` (`consensus.go:51-54`) — no threshold configured for this key at all. A caller-side data/config mismatch, not a DON-agreement problem — deserves its own reason so it doesn't get chased as an oracle-health issue.
  2. `len(items) == 0` after `GetValid()` — no single value reached the threshold (too few oracles reporting, or votes scattered thin). The routine "insufficient agreement" case.
  3. `len(items) > 1` — two or more *distinct* values each independently cleared the threshold. Given BFT-style thresholds this is the more alarming one: it implies something actively changing mid-round across the DON, or is worth treating as a potential integrity signal rather than routine flakiness.

  Cases 2 and 3 want different on-call responses (see the merkleroot `ConsensusObservationFailed` writeup below for a worked example); case 1 wants a bug ticket, not a page.

  **Recommended shape:** don't emit metrics from inside this package — it's shared with `execute`, and hardcoding a `ccip_commit_*`-prefixed metric here would either be wrong for execute or force a plugin-type branch into shared code. Instead, extend the return signature additively so each caller routes the reason through its own `MetricsReporter`:
  ```go
  type ConsensusDropReason int

  const (
      DropReasonNone ConsensusDropReason = iota
      DropReasonThresholdNotDefined
      DropReasonInsufficientAgreement // 0 values met threshold
      DropReasonSplit                 // >1 values met threshold — treat as higher severity
  )

  func GetConsensusMap[K comparable, T any](
      lggr logger.Logger,
      objectName string,
      itemsByKey map[K][]T,
      minObs MultiThreshold[K],
  ) (map[K]T, map[K]ConsensusDropReason)
  ```
  Callers that don't care about metrics can `_` the second return value — small blast radius, ~6 call sites total. Callers that do care loop the returned map and call their own reporter: `ccip_commit_consensus_dropped_total{objectName, key, reason}` for commit, an analogous `ccip_exec_...` metric for execute.

  **Consistency note:** `GetConsensusMapAggregator` (`consensus.go:62-81`) and `GetOrderedConsensus` (`consensus.go:113-148`) have the same unmetriced pattern but structurally can't produce case 3 — the aggregator combines all values regardless of agreement, and ordered-consensus just needs a count, not matching values. `GetOrderedConsensus` does have its own distinct third case (`consensus.go:127-133`, "found a negative or 0 threshold" — a config-error class, currently silently skipped). If this file is being touched, apply the same additive `(result, reasons)` pattern to all three for one coherent `{objectName, key, reason}` metric family instead of fixing `GetConsensusMap` alone and leaving the other two as later follow-ups.

### Merkle root processor (`commit/merkleroot`, `commit/merkleroot/rmn`)

**High**
- **Raw onramp max seq num / offramp next seq num per source chain, every round** (your example). `observation.go:287-288` (`ObserveLatestOnRampSeqNums`/`ObserveOffRampNextSeqNums`), logged at `Debugw` only. → Gauge `{sourceChain}`, both series, distinct from the existing `ccip_commit_max_sequence_number`.
- **Pending-message backlog per chain** (`onRampMaxSeqNum - offRampNextSeqNum`). Computed inline at `outcome.go:152-179` then discarded. → Gauge `{sourceChain}`.
- **RMN global/destination curse active** — halts all reporting for the dest chain; today `Warnw` every round while cursed. `observation.go:553-565`. → Gauge `{chain_id, curse_type}` (1/0).
- **Per-source-chain curse** (lane silently excluded). `observation.go:562-582`, `Infow` only. → Gauge `{sourceChain}` (1/0).
- **Per-chain onramp/merkle-root read failures**, swallowed per-goroutine (no-bindings / timeout / rpc-error / msg-count-mismatch / hash-error / address-lookup-error) — never reaches `ccip_commit_processor_errors` since the outer `Observation()` call still succeeds. `observation.go:656-777`. → Counter `{sourceChain, reason}`.
- **`GetFChain` failure** — literally has `// TODO: metrics` in the source. `observation.go:862-871`. Degrades to empty FChain map, which then fails consensus and stalls the round. → Counter.
- **Consensus-observation-failed round** (`ConsensusObservationFailed`, entire outcome empty) — `outcome.go:88-92`, gated entirely by `getConsensusObservation`'s upfront requirement that the DON reach 2·fRoleDON+1 agreement on a *single* FChain value for the **destination chain** (`outcome.go:471-480`); every other per-chain consensus computation (roots, onramp/offramp seq nums, RMN config) degrades gracefully per-chain and cannot trigger this path. Two consequences for the runbook: (1) this failure is inherently **destination-chain-wide, not lane-specific** — if it's firing, every source chain reporting into this dest chain should be stalling simultaneously, not just one lane, which is a fast way to rule this branch in or out; (2) the counter alone only tells you *that* consensus failed, not *why*. Pair it with the shared `consensus.go` fix above (`objectName="fChain", key=<destChain>`) to see the drop directly, and with the `GetFChain`-failure counter to distinguish "too few oracles could even read it" (home-chain RPC/read outage) from "oracles disagree" (home-chain config mid-rollout/desynced across the DON — corroborate with `ccip_commit_config_digest_mismatch` flipping around the same time). Already a "stable log identifier" the team log-mines — evidence a metric is overdue, not a reason to skip one. → Counter, `{chain_family, chain_id}` (destination chain, matching labels used elsewhere in this reporter).
- **Report transmission gives up / retry cost.** `outcome.go` (`ReportTransmissionGaveUp`), `Outcome.ReportTransmissionCheckAttempts`. One of the most common on-call pages; today `Warnw` with no counter. → Counter + histogram of attempts.
- **RMN-enabled roots dropped wholesale** (signed-roots-not-subset / parse error) — an entire round's RMN-protected chains get zero inclusion. `outcome.go:270-299`, `Errorw`. → Counter `{reason}`.
- **RMN chain-level quorum result** (zero roots / insufficient votes / success) per source chain — `controller.go:557-603`, `Infow` only, easily lost in volume. → Counter `{sourceChain, result}`.
- **Byzantine root disagreement** (`"more than one valid root for chain"` / `"no valid root for chain"`) — a security/integrity signal currently indistinguishable from a mundane timeout. `controller.go:848-876`. → Counter `{reason}`.

**Medium**
- Truncation due to `MaxMerkleTreeSize` (chain throughput-bound). `outcome.go:174-178`, `Debugf`. → Counter `{sourceChain}`.
- Offramp lane misconfiguration bucket (live / skipped-not-a-lane / skipped-disabled / rmn-misconfigured); `rmnMisconfigured` in particular indicates config drift. `observation.go:53-89,634-646`. → Gauge `{sourceChain, status}`.
- OnRamp/OffRamp invariant violations (offramp-ahead-of-onramp, onramp-max-zero, offramp-seqnum-regression) — explicitly documented in-package as stall signals. `outcome.go:43-44,152-161,421-428`. → Counter `{sourceChain, type}`.
- Insufficient consensus on `OffRampNextSeqNums` for a chain — chain silently stops progressing, looks identical to "no new messages" from outside. `outcome.go:515-531`. → Counter `{sourceChain}`.
- Per-root inclusion/exclusion reason (rmn-enabled-unsigned vs rmn-disabled-unexpectedly-signed) — answers "why didn't chain X's root make the report." `outcome.go:342-378`. → Counter `{sourceChain, reason}`.
- `ValidateRootBlessings` mismatches at report-acceptance time (`commit/report.go:234` → `merkleroot/transmission_checks.go:28-52`) — distinct from the existing config-digest-mismatch gauge. → Counter `{reason}`.
- RMN node-coverage shortfall per chain (dropped from signature request entirely). `controller.go:188-203`. → Counter/gauge `{sourceChain}`.
- RMN internal-error classes currently unlabeled or folded into a generic `invalid_response` (dest==source config bug, sanity-check failure, node-info-not-found / stale registry). `controller.go:395-398,358-367,421-425,908-909,1010-1013`. → Counter `{reason}`.
- Signature-verification failure conflated with generic "invalid response" (crypto mismatch vs. malformed-but-honest data are very different security signals), both in the RMN controller (`crypto.go:36-63`, `controller.go:1052-1089`) and on the plugin/query side (`observation.go:245-248`). → Split the existing label or add a dedicated counter.
- `ErrSignaturesNotProvidedByLeader` recurrence — points to a leader-side RMN problem invisible to followers today. `observation.go:41-42,118-125`. → Counter.
- RMN `ObserveRMNRemoteCfg` failures — blocks RMN controller init and outcome's RMN report building. `observation.go:846-856`, `Errorw`. → Counter.

**Low**
- RMN peer/stream connectivity state (active streams, peer-group connected) — currently only per-message `Infow` (log noise, no aggregate). `rmn/peerclient.go:57-183`. → Gauge.
- RMN controller re-initialization churn (should be rare; frequent churn = config flapping or bug). `observation.go:143-185`. → Counter. Also worth a static `ccip_commit_rmn_enabled{chain_id}` gauge since RMN is now hardcoded off — useful for dashboard filtering / confirming the deprecation.

### Chain fee processor (`commit/chainfee`)

No dedicated `MetricsReporter` exists for this package at all (unlike merkleroot) — everything below is currently 100% log-only.

**High**
- **A single `GetChainConfig` failure for the dest chain aborts gas-price updates for *every* chain that round.** `outcome.go:259-263`, `Errorw`. The subagent that found this called it "likely the single highest-value item to alert on" — it's a total, silent loss of gas-price-update capability for a round. → Counter `ccip_commit_chainfee_update_aborted_total`, plus `{chain_id, scope=dest|source}` for the narrower per-source-chain config errors at `outcome.go:268-272,311-315`.
- **Async op timeout → silently empty fee data**, no log, no metric anywhere (see cross-cutting #5). `observation.go:43-62` blindly type-asserts missing keys to zero value. → Counter `{operation}`.
- **Per-chain USD gas fee (exec + DA) never gauged.** `outcome.go:65,103-108,129`, only `Infow`. → Gauge `{chain_id, component=exec|da}`.
- **Fee staleness discarded after the boolean heartbeat decision.** `outcome.go:276,293-298` computes `consensusTimestamp.Sub(lastUpdate.Timestamp)` then keeps only a bool. → Gauge, seconds, `{chain_id}`.
- **Whole-round "no consensus on any fee components."** `outcome.go:59-63`, `Warnw`. A whole-DON health signal — all chains' fee updates stall that round. → Counter `{chain_id}`.
- **Consensus-map silently drops a chain** (fChain / FeeComponents / NativeTokenPrices / ChainFeeUpdates) when too few nodes agree — shared `consensus.go`, has `// TODO: metrics` in the source. See the "Shared: consensus aggregation" writeup above for the reviewed fix (adds a `reason`: threshold-not-defined vs. insufficient-agreement vs. split, not just a bare counter). → Counter `{objectName, chain_id, reason}` (shared fix benefits tokenprice too).

**Medium**
- Chains missing native token price, dropped from the update entirely (`missingNativeTokenPriceChains`). `outcome.go:66,81-85,126`, `Infow`. → Gauge/counter `{chain_id}`.
- Same distinction at observation time, per-oracle (before consensus) — could isolate a single bad node's fee-quoter reader. `observation.go:65-86`. → Counter, watch cardinality.
- Update-decision reason breakdown (no-previous / heartbeat / deviation / not-needed) — another "stable identifier" the log tooling already keys on; that's a reason to graduate it, not skip it. `outcome.go:21-38`. → Counter `{chain_id, reason}`.
- Deviation magnitude discarded, boolean kept ("exceeded threshold" vs. "by how much"). `outcome.go:317-328`. → Gauge `{chain_id, component}`.
- Async-cache staleness (async mode serves a ticker-refreshed cache with no tracked last-successful-sync time — a stuck background `sync()` looks healthy indefinitely). `observation_async.go:264-302`. → Gauge, last-success timestamp.
- Reader-side error branches (`getSupportedSourceChains`, `SupportsDestChain`, `getEnabledSourceChains` failures) collapse into the same empty map as "legitimately zero chains configured." `observation_async.go:53-97`, `Errorw`. → Counter `{operation, chain_id}`.

**Low**
- Disabled/unsupported chain exclusion counts (config-drift detection across a large DON). `observation_async.go:59-129`, `Debugw`. → Gauge, supported/enabled chain counts.

### Token price processor (`commit/tokenprice`)

Also no dedicated `MetricsReporter`. Note `ValidateObservation` isn't even wrapped by the generic `TrackedProcessor` latency/error instrumentation — it's invisible by construction, not just under-instrumented.

**High**
- **`ValidateObservation` rejections have zero logs *and* zero metrics** — every failure path (non-positive price, unsupported chain, missing token info, bad timestamp) is a bare `fmt.Errorf` with no `lggr` call anywhere in `validate_observation.go:12-88`, and the wrapper isn't tracked. This is exactly the signal for detecting a misbehaving/misconfigured oracle in the DON. → Counter `{oracleID, reason}`.
- **Price/FeeQuoter/FChain fetch failures degrade to empty maps that look identical to "legitimately nothing to report."** `observation.go:73-79,115-142,144-181`, `Errorw`, and `Observation()` still returns `nil` error so `ccip_commit_processor_errors` never fires. → Counter `{source=feed|feequoter|fchain}`.
- **Per-token feed/fee-quoter USD price never gauged.** `outcome.go:107`, `types.go:31,59`. → Gauge `{token, destChain}`.
- **Token price staleness discarded after the boolean heartbeat decision** — same pattern as chainfee. `outcome.go:127-129`. → Gauge, seconds, `{token}`.
- **Consensus-not-reached per token/chain silently drops it from the map** (the closest thing to "outlier rejection" in this codebase) — shared `consensus.go`, also has `// TODO: metrics`. See the "Shared: consensus aggregation" writeup above for the reviewed fix (adds a `reason`: threshold-not-defined vs. insufficient-agreement vs. split, not just a bare counter). → Counter `{objectName, key, reason}`.

**Medium**
- Missing `TokenInfo` config count (feed has a token the offchain config doesn't know about) — `Warnw` per token per round, no aggregate. `outcome.go:120-125`. → Counter `{token}`.
- Update-decision reason breakdown (no-previous / heartbeat / deviation / not-needed). `outcome.go:22-30`. → Counter `{token, reason}`.
- Deviation magnitude discarded — `mathslib.Deviates` computes a ppb diff internally and returns only a bool. `internal/libs/mathslib/calc.go:25-39`. → Would need `Deviates` (or a wrapper) to also return the magnitude; Gauge `{token}`.
- Async observer cache staleness (same risk as chainfee's async cache — stuck `sync()` serves old prices indefinitely, looks healthy). `observation.go:187-330`. → Gauge, last-success timestamp `{source}`.
- Oracle feed/dest-chain support gaps — per-node capability info that explains *why* consensus on `fFeedChain` might fail. `observation.go:121-130,152-159`. → Gauge/counter `{oracleID, chainRole}`.

**Low, but flagged as a miss in my first pass — don't let "it's Debug" read as "low priority"**
- **Token-price reporting silently disabled by config** (`TokenPriceBatchWriteFrequency == 0`) is logged at `Debugw` — *less* visible in prod than a "stable" Info/Warn log, since Debug is typically off entirely. `processor.go:102-106`. This is exactly the pattern the existing `ccip_commit_loopp_ccip_provider_supported` boolean gauge was built for, and it's a cheap, precedented fix. → Gauge `{destChain}` (1/0).

---

## Suggested next step

This is large enough to be its own workstream. I'd sequence it as:
1. **Shared infra first** (benefits all three processors): fix `internal/plugincommon/consensus/consensus.go`'s `// TODO: metrics`, and add a timeout counter in `internal/libs/asynclib/goroutines.go`. Small, high-leverage, no processor-specific design needed.
2. **Merkleroot high-severity list** — directly answers your stated pain point (onramp/offramp liveness, curse state, transmission give-up).
3. **Chainfee + tokenprice high-severity list** — both need a dedicated `MetricsReporter` sub-interface first (mirroring `merkleroot.MetricsReporter`), then the per-value gauges/staleness/failure counters.
4. Medium/low tiers as follow-up passes.

Want me to turn this into an implementation plan (starting with #1), or file it as a ticket first?

---

## War games: incident walkthroughs

Three scenarios worked through against the merkleroot-high-severity metric set above (some referencing not-yet-implemented items, noted inline) to validate the design before building it: does an on-call engineer actually get a straight-line diagnosis, or just more noise? Each starts from the same alert: **message `0xabc...def`, onramp sequence number `X`, chain A (source) → chain B (dest), reported in-flight for 15 minutes.**

### Scenario 1 — full decision tree

1. **Is this even a commit-plugin problem?** Check `ccip_commit_offramp_next_seq_num{source_network_name="chain-a", dest_network_name="chain-b"}`.
   - `> X` → a root covering X already landed on the offramp. Commit plugin's job is done; hand off to exec on-call. Stop.
   - `<= X` → continue.
2. **Has the plugin observed X on-chain yet?** Check `ccip_commit_onramp_max_seq_num{source_network_name="chain-a"}`.
   - `< X` → onramp read is lagging. Check `ccip_commit_merkleroot_observation_errors{sourceChain="chain-a"}` by reason (no_bindings/timeout/rpc_error). Source-chain read/infra issue, not commit logic.
   - `>= X` → plugin has seen it; continue.
3. **Is the round producing outcomes at all?** Check `rate(ccip_commit_consensus_observation_failed[5m])`.
   - Incrementing → destination-chain-wide consensus failure — see Scenario 3 deep dive.
   - Flat → continue.
4. **Is chain A or B cursed?** Check `ccip_commit_rmn_curse_active{chain_id="chain-b", curse_type="global"|"destination"}` and `ccip_commit_source_chain_cursed{sourceChain="chain-a"}`.
   - Any `== 1` → found it; likely intentional/incident-flagged, hand off to whoever owns the curse. Stop.
   - All `0` → continue.
5. **Is a report being built but never landing?** Check `ccip_commit_report_transmission_gave_up{sourceChain="chain-a"}` and `ccip_commit_report_transmission_attempts`.
   - Incrementing/maxed → see Scenario 2 deep dive.
   - Otherwise → likely just backlog size; check `ccip_commit_pending_messages{sourceChain="chain-a"}` and estimate ETA from round cadence.

(RMN-signature-specific branches — signed-roots-dropped, chain quorum result, Byzantine root disagreement — were dropped from the implementation scope: RMN blessing/signing is effectively dead code now that `RMNEnabled` is hardcoded off. Cursing is the one RMN-adjacent signal still live in production, hence Step 4.)

### Scenario 2 — deep dive: report transmission stuck (Step 5)

Three hypotheses, in order of how cheap they are to check:

- **A — never transmitted (rejected pre-send).** `sum by (reason) (rate(ccip_commit_report_validation_rejected{phase="should_transmit"}[15m]))`.
  - `config_digest_mismatch` climbing → cross-check `ccip_commit_config_digest_mismatch`; config sync issue.
  - `stale` climbing → often benign — a different, overlapping report already landed; re-check Step 1's gauge, it may have just moved.
  - `dest_not_supported` → transmission-schedule/config bug.
  - `cursed` → converges with Step 4.
  - All flat → move to B.
- **B — transmitted, reverted/stuck on-chain.** Outside commit-plugin metrics: check chain B's TXM/tx-sender dashboards for the assigned transmitter (nonce gap, underpriced gas, wallet balance, revert reason). A revert here for the same reasons as A but caught on-chain instead of locally usually means a race between the pre-check and the tx landing (config rollout, chain congestion).
- **C — actually succeeded, read is lagging.** `ccip_commit_offramp_next_seq_num` is sourced from the plugin's own chain-B reader; if that's behind, "no change" can persist for a few rounds after a real success. Cross-check chain B's block explorer / chain-reader health directly before escalating further.

### Scenario 3 — deep dive: consensus never completes (Step 3)

Grounded in `getConsensusObservation` (`merkleroot/outcome.go:461-508`): the *only* way the whole round fails is the DON not reaching 2·fRoleDON+1 agreement on a single FChain value for the **destination chain** — every other field degrades per-chain gracefully. Two implications:

- **This is destination-chain-wide, not lane-specific.** Before going further, check whether *other* source chains reporting into chain B are also stalled (their `ccip_commit_pending_messages` also climbing). If only chain A is affected, this branch is the wrong one — back to Scenario 1, Step 4/5.
- **The counter alone can't tell you why.** Two sub-hypotheses, using the not-yet-implemented shared `consensus.go` fix:
  - **H1 — too few oracles could read it.** `ccip_commit_fchain_read_errors` spiking broadly across oracles → home-chain RPC/read outage.
  - **H2 — oracles disagree.** `ccip_commit_consensus_dropped{objectName="fChain", reason="split"}` firing while H1 stays flat → home-chain config (e.g. CCIPHome update, DON membership change) mid-rollout/desynced across the DON. Corroborate with `ccip_commit_config_digest_mismatch` flipping for a subset of oracles around the same time.

### Gaps this exercise surfaced (already folded into the findings above)
- No plugin-level heartbeat — added to "Top-level plugin orchestration."
- No reason breakdown for report-acceptance/transmission-validation rejections — added to "Top-level plugin orchestration."
- `GetConsensusMap`'s two-vs-three-outcome ambiguity — added to "Shared: consensus aggregation" (not yet implemented — reviewed design only).
