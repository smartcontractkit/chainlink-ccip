- [Metrics coverage assessment — commit plugin](#metrics-coverage-assessment--commit-plugin)
   * [What exists today (`commit/metrics/prom.go`)](#what-exists-today-commitmetricspromgo)
   * [Cross-cutting themes (repeat in every processor)](#cross-cutting-themes-repeat-in-every-processor)
   * [Proposed metrics at a glance (merkle root processor)](#proposed-metrics-at-a-glance-merkle-root-processor)
      + [Top-level plugin orchestration](#top-level-plugin-orchestration)
      + [Shared: consensus aggregation](#shared-consensus-aggregation)
      + [Merkle root processor — High](#merkle-root-processor--high)
      + [Merkle root processor — Medium](#merkle-root-processor--medium)
   * [Incident triage](#incident-triage)
   * [Full findings](#full-findings)
      + [Top-level plugin orchestration (`commit/plugin.go`, `commit/report.go`)](#top-level-plugin-orchestration-commitplugingo-commitreportgo)
      + [Shared: consensus aggregation (`internal/plugincommon/consensus/consensus.go`)](#shared-consensus-aggregation-internalplugincommonconsensusconsensusgo)
      + [Merkle root processor (`commit/merkleroot`)](#merkle-root-processor-commitmerkleroot)
   * [(Other Processors) Proposed metrics at a glance](#other-processors-proposed-metrics-at-a-glance)
         - [Chain fee processor — High](#chain-fee-processor--high)
         - [Chain fee processor — Medium / Low](#chain-fee-processor--medium--low)
         - [Token price processor — High](#token-price-processor--high)
         - [Token price processor — Medium / Low](#token-price-processor--medium--low)
      + [Chain fee processor (`commit/chainfee`)](#chain-fee-processor-commitchainfee)
      + [Token price processor (`commit/tokenprice`)](#token-price-processor-committokenprice)

# Metrics coverage assessment — commit plugin

**Ground rule for everything below: the presence of a log line — even a "stable identifier" some log-analysis tooling already keys on — is never a reason to exclude a metric.** Logs are exactly the noisy, low-signal, un-queryable thing motivating this work; a few items here are genuinely *only* visible at `Debugw`, which is worse than a "stable" log, not better. Nothing was cut from the list below because it was already logged.

## What exists today (`commit/metrics/prom.go`)

- Generic per-processor instrumentation via `TrackedProcessor` (wraps merkleroot/chainfee/tokenprice): latency histogram, error counter, and an **output-size counter** driven by each type's `Stats()` — but `Stats()` only returns `len(map)`-style item counts (e.g. `gasPrices: 3`), never the actual values.
- `ccip_commit_max_sequence_number` gauge — but only fed from `MerkleRootChain.SeqNumsRange.End()` in `TrackObservation`/`TrackOutcome`, i.e. **only when a root is actually being built for a report**.
- `ccip_commit_latest_round_id`, `ccip_commit_config_digest_mismatch`, `ccip_commit_loopp_ccip_provider_supported`.
- **`chainfee` and `tokenprice` have no dedicated metrics interface at all** — they only get the generic latency/error/size instrumentation above. Token prices, gas prices, and their staleness/deviation/failure states are otherwise 100% log-only.

## Cross-cutting themes (repeat in every processor)

Merkleroot, chainfee, and tokenprice each independently:
1. Compute a **per-item value** (seq num / fee / price) that's never gauged — only counted.
2. Compute **staleness** (age since last update) to decide on a heartbeat write, then discard the number and keep only the boolean.
3. Swallow **fetch/RPC failures** into empty maps with an `Errorw` — these never increment `ccip_commit_processor_errors` because the processor's top-level `Observation()` still returns `nil` error.
4. Drop entries during **consensus aggregation** silently — `internal/plugincommon/consensus/consensus.go:42,52` literally has `// TODO: metrics` next to the exact code path that discards a chain/token when too few nodes agree. This is shared infra, hit by every processor, and the team already flagged it (reviewed fix shape below).
5. `internal/libs/asynclib/goroutines.go` (used by chainfee's async observer, similar pattern in tokenprice) has **zero logging or metrics** when an operation times out — it just drops the key from the results map and the caller's blind type-assertion silently degrades to an empty value. This is the most severe blind spot found: a hung RPC call produces no log line and no metric anywhere.
6. **No plugin-level heartbeat/liveness signal.** Surfaced by war-gaming an incident end-to-end rather than by reading the code in isolation: if the whole round pipeline wedges or the process crashes, every metric in this doc simply stops updating instead of showing a bad value — and "gauge went flat" reads very differently in monitoring than "gauge shows a bad value," easy to miss unless dashboards specifically alert on staleness/absence rather than thresholds. There's no dedicated "is the plugin alive at all" signal today; see below.

---

## Proposed metrics at a glance (merkle root processor)

Every metric proposed in this doc, with a one/two-line "why" — the full rationale, exact file:line evidence, and severity ranking live in [Full findings](#full-findings) below. ✅ = implemented and merged; everything else is proposed/reviewed only.

### Top-level plugin orchestration

| Metric | Status | Why |
|---|---|---|
| `ccip_commit_plugin_heartbeat{chainFamily,chainID,phase="observation"\|"outcome"}` | ✅ Implemented | Unconditional per-round liveness signal — the only way to tell "process wedged" apart from "valid but stale gauge" for every other metric in this doc. |
| `ccip_commit_report_validation_rejected{phase="should_accept"\|"should_transmit",reason}` | ✅ Implemented | Distinguishes "report never submitted" (rejected locally) from "submitted but stuck/reverted on-chain" — different remediation paths entirely. |

### Shared: consensus aggregation

| Metric | Status | Why |
|---|---|---|
| `ccip_commit_consensus_dropped{objectName,key,reason}` | ✅ Implemented | Splits "no threshold configured" (config bug) from "insufficient agreement" (routine) from "split vote" (integrity signal) — previously collapsed into one `TODO: metrics` log line. |

### Merkle root processor — High

| Metric | Status | Why |
|---|---|---|
| `ccip_commit_onramp_max_seq_num{sourceChain}` / `ccip_commit_offramp_next_seq_num{sourceChain}` | ✅ Implemented | Direct per-lane chain-read liveness — the original motivating gap for this whole assessment. |
| `ccip_commit_pending_messages{sourceChain}` | ✅ Implemented | Quantifies "how far behind" instead of forcing on-call to mentally subtract the two seq-num gauges. |
| `ccip_commit_rmn_curse_active{chain_id,curse_type="global"\|"destination"}` | ✅ Implemented | Curse halts all reporting for the dest chain; needs to be page-visible, not a per-round `Warnw`. |
| `ccip_commit_source_chain_cursed{sourceChain}` | ✅ Implemented | Distinguishes "this lane is cursed" (expected, hand off) from "plugin is wedged" (page). |
| `ccip_commit_merkleroot_observation_errors{sourceChain,reason}` | ✅ Implemented | Surfaces per-goroutine RPC/read failures that are swallowed and never trip the generic processor-error counter. |
| `ccip_commit_fchain_read_errors` | ✅ Implemented | Distinguishes "can't read home chain" from "DON disagrees" — the two root causes of a consensus-failed round. |
| `ccip_commit_consensus_observation_failed{chain_family,chain_id}` | ✅ Implemented | Flags a destination-chain-wide stall immediately instead of looking like N independent lane stalls. |
| `ccip_commit_report_transmission_gave_up{sourceChain}` + `ccip_commit_report_transmission_attempts` (histogram) | ✅ Implemented | One of the most common on-call pages today; previously only a `Warnw` with no counter. |

### Merkle root processor — Medium

| Metric | Status | Why |
|---|---|---|
| `ccip_commit_range_truncated{sourceChain}` (`MaxMerkleTreeSize`) | ✅ Implemented | Chain-throughput ceiling being hit was previously `Debugf`-only. |
| `ccip_commit_offramp_lane_status{sourceChain,status}` | ✅ Implemented | `rmn_misconfigured` in particular now flags real config drift, since a lane still expecting RMN blessing while RMN is globally off is permanently broken. |
| `ccip_commit_seqnum_invariant_violation{sourceChain,type}` | ✅ Implemented | These are already documented in-code as stall signals but were never counted. |
| `ccip_commit_offramp_consensus_insufficient{sourceChain}` | ✅ Implemented | Looks identical to "no new messages" from the outside without it. |

---

## Incident triage

The metrics above were validated against simulated incidents before being built (does an
on-call engineer actually get a straight-line diagnosis, or just more noise?), and that
walkthrough has since been extracted, hardened, and kept current as its own pair of runbooks
rather than living here as a point-in-time exercise:

- [`docs/runbooks/uncommitted-message.md`](../runbooks/uncommitted-message.md) — triage for a
  specific stuck message (the full decision tree and both deep dives that used to be here).
- [`docs/runbooks/commit-plugin-health.md`](../runbooks/commit-plugin-health.md) — point-in-time
  health check, no specific incident required.

Both are written to be followed mechanically by a human or an AI agent (see
[`docs/runbooks/README.md`](../runbooks/README.md)) and have been tested against a live devenv
DON under two independently forced failure modes, not just read for plausibility.

---

## Full findings

Severity is about production-triage value, independent of whether something is logged today. This section is the deep dive backing the brief table above: full rationale, file:line evidence, and severity ranking for each metric.

### Top-level plugin orchestration (`commit/plugin.go`, `commit/report.go`)

These aren't owned by any single processor — they're the outer round loop and the report-acceptance path — but two gaps here turned out to matter as much as anything processor-specific once an actual incident was traced end-to-end.

**High**
- **No heartbeat / liveness signal.** `commit/plugin.go:434` (`Outcome()`). If the whole round pipeline wedges (deadlock, a panic-recovery masking a crash loop, OCR itself failing to schedule rounds), every metric in this doc simply stops updating instead of showing a bad value. A stale gauge and a valid-but-bad gauge look identical unless you specifically alert on absence/staleness — easy to forget when every other alert here is threshold-based. This should be step 0 of any triage tree, and today it isn't answerable from commit-plugin metrics at all. → **Implemented**: counter `ccip_commit_plugin_heartbeat{chainFamily, chainID, phase="observation"|"outcome"}`, incremented unconditionally at the very top of `Observation()`/`Outcome()`, before any decode, state check, or early return.

  **Why this isn't already covered by the pre-existing `ccip_commit_latest_round_id` gauge**, since at a glance both look like "something updates roughly once per round": `ccip_commit_latest_round_id` is set by `trackLatestRoundID`, called from inside `TrackObservation`/`TrackOutcome` — but only while looping over `obs.MerkleRootObs.MerkleRoots` / `outcome.MerkleRootOutcome.RootsToReport`. Per `merkleroot.getObservation` (`merkleroot/observation.go:281-331`), that field is only populated in the `buildingReport` state; it's empty in `selectingRangesForReport` and `waitingForReportTransmission`, i.e. for most of the state machine's time. It's also labeled per **source chain**, not per plugin instance. So it only moves during the fraction of rounds where the processor happens to be actively building a report with at least one root in it — a perfectly healthy processor sitting in `selectingRangesForReport` because a lane has no new messages will let this gauge go stale for entirely benign reasons. You can't tell "no traffic on this lane" apart from "plugin is wedged" by watching it alone. `ccip_commit_plugin_heartbeat` is unconditional and per-instance precisely to remove that ambiguity: the two are complementary, not overlapping — `latest_round_id` tells you about a specific chain's last-reported root, `plugin_heartbeat` tells you the process is alive at all, and you need the second to safely interpret staleness in the first (or in any other gauge in this doc).
- **Report-acceptance / transmission-validation rejections have zero metrics.** `commit/report.go:126-247` (`validateReport`, called from both `ShouldAcceptAttestedReport` and `ShouldTransmitAcceptedReport`) has several distinct rejection paths — decode failures, empty/invalid merkle root, cursed report, dest chain unsupported, config digest mismatch, stale report, root-blessing mismatch — every one of them `Warnw`/`Errorw`/`Debugf` only. (The RMN-signature rejection paths in the same function — duplicate/insufficient RMN signatures — aren't instrumented: RMN blessing/signing is effectively dead code now that `RMNEnabled` is hardcoded off, so they shouldn't occur.) This sits directly upstream of merkleroot's own `ReportTransmissionGaveUp` counter: if every oracle's local `validateReport` call rejects a report before anyone ever calls transmit, "gave up" fires with **zero on-chain trace at all** (no tx ever submitted) — a completely different remediation path than "tx submitted but reverted" or "tx stuck in mempool." Without per-reason counts here, those scenarios are indistinguishable from `ccip_commit_report_transmission_gave_up_total` alone. → **Implemented**: counter `ccip_commit_report_validation_rejected{phase="should_accept"|"should_transmit", reason="decode_report|decode_report_info|empty_root|invalid_seqnum_range|cursed_check_error|cursed|dest_support_check_error|dest_not_supported|config_digest_check_error|config_digest_mismatch|stale|root_blessing_mismatch"}`.

### Shared: consensus aggregation (`internal/plugincommon/consensus/consensus.go`)

Used by commit (`chainfee`, `merkleroot`, `plugin.go`, `report.go`), `execute/plugin_functions.go`, **and** `internal/plugincommon/discovery/processor.go` — this is genuinely shared infra, not commit-specific, which shapes the recommended fix below. ✅ Implemented.

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

### Merkle root processor (`commit/merkleroot`)

**High**
- **Raw onramp max seq num / offramp next seq num per source chain, every round** (your example). `observation.go:287-288` (`ObserveLatestOnRampSeqNums`/`ObserveOffRampNextSeqNums`), logged at `Debugw` only. → Gauge `{sourceChain}`, both series, distinct from the existing `ccip_commit_max_sequence_number`.
- **Pending-message backlog per chain** (`onRampMaxSeqNum - offRampNextSeqNum`). Computed inline at `outcome.go:152-179` then discarded. → Gauge `{sourceChain}`.
- **RMN global/destination curse active** — halts all reporting for the dest chain; today `Warnw` every round while cursed. `observation.go:553-565`. → Gauge `{chain_id, curse_type}` (1/0).
- **Per-source-chain curse** (lane silently excluded). `observation.go:562-582`, `Infow` only. → Gauge `{sourceChain}` (1/0).
- **Per-chain onramp/merkle-root read failures**, swallowed per-goroutine (no-bindings / timeout / rpc-error / msg-count-mismatch / hash-error / address-lookup-error) — never reaches `ccip_commit_processor_errors` since the outer `Observation()` call still succeeds. `observation.go:656-777`. → Counter `{sourceChain, reason}`.
- **`GetFChain` failure** — literally has `// TODO: metrics` in the source. `observation.go:862-871`. Degrades to empty FChain map, which then fails consensus and stalls the round. → Counter.
- **Consensus-observation-failed round** (`ConsensusObservationFailed`, entire outcome empty) — `outcome.go:88-92`, gated entirely by `getConsensusObservation`'s upfront requirement that the DON reach 2·fRoleDON+1 agreement on a *single* FChain value for the **destination chain** (`outcome.go:471-480`); every other per-chain consensus computation (roots, onramp/offramp seq nums, RMN config) degrades gracefully per-chain and cannot trigger this path. Two consequences for the runbook: (1) this failure is inherently **destination-chain-wide, not lane-specific** — if it's firing, every source chain reporting into this dest chain should be stalling simultaneously, not just one lane, which is a fast way to rule this branch in or out; (2) the counter alone only tells you *that* consensus failed, not *why*. Pair it with the shared `consensus.go` fix above (`objectName="fChain", key=<destChain>`) to see the drop directly, and with the `GetFChain`-failure counter to distinguish "too few oracles could even read it" (home-chain RPC/read outage) from "oracles disagree" (home-chain config mid-rollout/desynced across the DON — corroborate with `ccip_commit_config_digest_mismatch` flipping around the same time). Already a "stable log identifier" the team log-mines — evidence a metric is overdue, not a reason to skip one. → Counter, `{chain_family, chain_id}` (destination chain, matching labels used elsewhere in this reporter).
- **Report transmission gives up / retry cost.** `outcome.go` (`ReportTransmissionGaveUp`), `Outcome.ReportTransmissionCheckAttempts`. One of the most common on-call pages; today `Warnw` with no counter. → Counter + histogram of attempts.

**Medium**
- Truncation due to `MaxMerkleTreeSize` (chain throughput-bound). `outcome.go:174-178`, `Debugf`. → **Implemented**: counter `ccip_commit_range_truncated{sourceChain}`.
- Offramp lane misconfiguration bucket (live / skipped-not-a-lane / skipped-disabled / rmn-misconfigured); `rmnMisconfigured` in particular indicates config drift. `observation.go:53-89,634-646`. → **Implemented**: gauge `ccip_commit_offramp_lane_status{sourceChain, status="live"|"skipped_not_a_lane"|"skipped_disabled"|"rmn_misconfigured"}` (1/0), reported for all four statuses every round so a chain moving between buckets can't leave a stale "still active" value behind.
- OnRamp/OffRamp invariant violations (offramp-ahead-of-onramp, onramp-max-zero, offramp-seqnum-regression) — explicitly documented in-package as stall signals. `outcome.go:43-44,152-161,421-428`. → **Implemented**: counter `ccip_commit_seqnum_invariant_violation{sourceChain, type="offramp_ahead_of_onramp"|"onramp_max_zero"|"offramp_seqnum_regression"}`.
- Insufficient consensus on `OffRampNextSeqNums` for a chain — chain silently stops progressing, looks identical to "no new messages" from outside. `outcome.go:515-531`. → **Implemented**: counter `ccip_commit_offramp_consensus_insufficient{sourceChain}`.

(RMN blessing/signing findings — signed-roots-dropped, per-root signed/unsigned filtering, chain quorum result, Byzantine root disagreement, RMN controller node-coverage/internal-errors, signature-verification conflation, `ErrSignaturesNotProvidedByLeader`, `ObserveRMNRemoteCfg` failures, peer/stream connectivity, controller re-init churn, and the standalone `ValidateRootBlessings` proposal (now covered by the implemented `root_blessing_mismatch` reason above) — were all dropped from this doc. All of these sit behind `if rmnEnabled` in `buildMerkleRootsOutcome` or are otherwise part of the RMN controller/signing path, which is effectively dead code now that `RMNEnabled` is hardcoded off and the controller is slated for removal; cursing, covered above, is the one RMN-adjacent signal still live in production.)

---

## (Other Processors) Proposed metrics at a glance

#### Chain fee processor — High

| Metric | Status | Why |
|---|---|---|
| `ccip_commit_chainfee_update_aborted_total{chain_id,scope="dest"\|"source"}` | Proposed | A single dest-chain `GetChainConfig` failure silently aborts gas-price updates for *every* chain that round — flagged as likely the single highest-value item to alert on. |
| Async-op-timeout counter `{operation}` | Proposed | A hung RPC call today produces zero log and zero metric anywhere. |
| Per-chain USD gas fee gauge `{chain_id,component="exec"\|"da"}` | Proposed | The actual fee value is only ever logged (`Infow`), never gauged. |
| Fee staleness gauge (seconds) `{chain_id}` | Proposed | Staleness is computed then discarded down to a bare heartbeat boolean. |
| Whole-round no-consensus-on-any-fee-component counter `{chain_id}` | Proposed | A DON-wide health signal that stalls every chain's fee update that round; currently `Warnw` only. |
| Consensus-map dropped-chain counter `{objectName,chain_id,reason}` | Proposed | Same shared-infra fix as above; benefits tokenprice too. |

#### Chain fee processor — Medium / Low

| Metric | Status | Why |
|---|---|---|
| Missing-native-token-price gauge/counter `{chain_id}` | Proposed | Chains silently dropped from the fee update entirely. |
| Per-oracle missing-price counter (pre-consensus) `{chain_id}` | Proposed | Could isolate a single bad node's fee-quoter reader before it's masked by consensus. |
| Update-decision reason-breakdown counter `{chain_id,reason}` | Proposed | Another "stable log identifier" already log-mined — reason to graduate it, not skip it. |
| Deviation-magnitude gauge `{chain_id,component}` | Proposed | Only a boolean ("exceeded threshold") survives today; the magnitude is discarded. |
| Async-cache last-success-timestamp gauge | Proposed | A stuck background `sync()` currently looks healthy indefinitely. |
| Reader-error-branch counter `{operation,chain_id}` | Proposed | Real errors collapse into the same empty map as "legitimately zero chains configured." |
| Disabled/unsupported chain-count gauge | Proposed (Low) | Config-drift detection across a large DON. |

#### Token price processor — High

| Metric | Status | Why |
|---|---|---|
| `ValidateObservation` rejection counter `{oracleID,reason}` | Proposed | Zero logs *and* zero metrics today for the exact signal that detects a misbehaving/misconfigured oracle. |
| Fetch-failure counter `{source="feed"\|"feequoter"\|"fchain"}` | Proposed | Failures degrade to empty maps indistinguishable from "legitimately nothing to report." |
| Per-token USD price gauge `{token,destChain}` | Proposed | Feed/fee-quoter price is never gauged, only logged. |
| Token price staleness gauge (seconds) `{token}` | Proposed | Same discard-the-number-keep-the-bool pattern as chainfee. |
| Consensus-dropped counter `{objectName,key,reason}` | Proposed | Shared fix with chainfee/merkleroot above — the closest thing to "outlier rejection" in this codebase, currently silent. |

#### Token price processor — Medium / Low

| Metric | Status | Why |
|---|---|---|
| Missing-`TokenInfo` counter `{token}` | Proposed | Logged per-token per-round (`Warnw`) with no aggregate view. |
| Update-decision reason-breakdown counter `{token,reason}` | Proposed | Same "graduate, don't skip" logic as chainfee's version. |
| Deviation-magnitude gauge `{token}` | Proposed | `mathslib.Deviates` computes a ppb diff internally and returns only a bool. |
| Async cache staleness gauge `{source}` | Proposed | Same stuck-`sync()`-looks-healthy risk as chainfee. |
| Oracle feed/dest-chain support gauge `{oracleID,chainRole}` | Proposed | Per-node capability info that explains *why* consensus on `fFeedChain` might fail. |
| Token-price-reporting-disabled gauge `{destChain}` (1/0) | Proposed (Low, but don't deprioritize) | Currently `Debugw`-only — *less* visible than a "stable" Info/Warn log, since Debug is typically off entirely in prod. |

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
