- [Metrics coverage assessment — execute plugin](#metrics-coverage-assessment--execute-plugin)
   * [What exists today (`execute/metrics/prom.go`, `execute/tracked.go`, tokendata)](#what-exists-today-executemetricspromgo-executetrackedgo-tokendata)
   * [Cross-cutting themes](#cross-cutting-themes)
   * [Proposed metrics at a glance — High](#proposed-metrics-at-a-glance--high)
      + [Top-level plugin orchestration and liveness](#top-level-plugin-orchestration-and-liveness)
      + [Consensus aggregation](#consensus-aggregation)
      + [Report acceptance / transmission validation](#report-acceptance--transmission-validation)
      + [Message pipeline (skips, build failures, read errors)](#message-pipeline-skips-build-failures-read-errors)
      + [Tokendata — High](#tokendata--high)
   * [Proposed metrics at a glance — Medium](#proposed-metrics-at-a-glance--medium)
   * [Full findings](#full-findings)
      + [Top-level plugin orchestration (`execute/plugin.go`, `execute/tracked.go`, `execute/outcome.go`)](#top-level-plugin-orchestration-executepplugingo-executetrackedgo-executeoutcomego)
      + [Observation phase (`execute/observation.go`)](#observation-phase-executeobservationgo)
      + [Outcome phase (`execute/outcome.go`)](#outcome-phase-executeoutcomego)
      + [Consensus computation (`execute/plugin_functions.go`, shared `internal/plugincommon/consensus`)](#consensus-computation-executepplugin_functionsgo-shared-internalplugincommonconsensus)
      + [Report building (`execute/report/`)](#report-building-executereport)
      + [Caches (`execute/internal/cache/`)](#caches-executeinternalcache)
      + [Tokendata (`execute/tokendata/`)](#tokendata-executetokendata)
   * [Cardinality guardrails](#cardinality-guardrails)

# Metrics coverage assessment — execute plugin

**Ground rule, same as the [commit assessment](commit-metrics.md): the presence of a log line — even a "stable identifier" log-analysis tooling already keys on — is never a reason to exclude a metric.** Logs are the noisy, slow-to-query thing motivating this work; several findings below are *only* visible at `Debug` level, which is worse than a "stable" log, not better. Nothing was cut because it was already logged.

Execute is a state machine over messages (not a fan-out of independent processors like commit), so the analysis is organized by pipeline stage: round orchestration → observation → consensus → outcome → report building → transmission, plus the tokendata subsystem that gates message readiness.

## What exists today (`execute/metrics/prom.go`, `execute/tracked.go`, tokendata)

- Generic per-method instrumentation via `TrackedPlugin` (`tracked.go`): latency histogram + error counter for **only** `Query`/`Observation`/`Outcome`. `Reports()`, `ShouldAcceptAttestedReport`, `ShouldTransmitAcceptedReport`, and `ValidateObservation` have **zero** metric coverage.
- Critically, `ccip_exec_errors` fires **only when the wrapped method returns an error**. Execute has a deliberate `return nil, nil` swallow idiom (see cross-cutting theme 1) that makes persistent total failures invisible to it.
- `ccip_exec_output_sizes` — driven by each type's `Stats()`, so it only counts what *made it into* the outcome (`messages`, `commitReports`, `tokenReady`/`tokenWaiting`, `nonces`). Everything that fell out *before* the outcome is invisible.
- `ccip_exec_max_sequence_number` — fed from the messages/commit reports present in each round's observation/outcome; goes stale exactly when a lane stops producing observations, with no way to tell "no traffic" from "broken".
- `ccip_exec_latest_round_id`, `ccip_exec_config_digest_mismatch`, `ccip_exec_loopp_ccip_provider_supported`, processor-level latency/errors.
- **Tokendata has partial, uneven coverage**: `ccip_attestation_request_duration` and `ccip_time_to_attestation` (`tokendata/observed.go`) wrap **only the LBTC client** — USDC v1 is unwired (`usdc.go:67` constructs the client directly), so USDC v1 attestation fetching has zero request-level metrics. The CCTP v2 observer has its own well-labeled family (`ccip_exec_tokendata_cctpv2_*`).

## Cross-cutting themes

1. **The `return nil, nil` swallow idiom defeats `ccip_exec_errors`.** Three phase entry points catch an error, log it, and return `nil, nil` so the next round resets cleanly: `observation.go:133-137` (`getMessagesObservation`), `observation.go:140-144` (`getFilterObservation`), `outcome.go:93-99` (`getFilterOutcome`). Plus partial swallows that degrade to empty values and continue: curse-read failure (`observation.go:231-237`), nonce-read failure (`observation.go:551-555`), discovery-outcome failure (`outcome.go:74-80`), config-digest check failure (`observation.go:62-75`). A DON where *every* round fails inside one of these shows **zero** on `ccip_exec_errors`, frozen gauges everywhere, and only `Errorw` logs.
2. **No plugin-level heartbeat or state-visibility signal.** `ccip_exec_latest_round_id` is only set inside `TrackObservation`/`TrackOutcome` — and both are skipped on the boring paths: `TrackObservation` is skipped entirely when a phase returns `nil, nil`, and `TrackOutcome` is skipped whenever the outcome is empty (`outcome.go:106-112`). Healthy idle and wedged produce identical metric output. There is also no gauge for *which* state of the state machine the plugin is in (`Initialized → GetCommitReports → GetMessages → Filter`), so "stuck in GetMessages for an hour" is unanswerable from metrics.
3. **Consensus drops are silent.** `getConsensusFChain` (`plugin_functions.go:924`) discards the `dropReasons` return value of `GetConsensusMap` — the source carries a literal `// TODO: report metrics (execute equivalent of ccip_commit_consensus_dropped)` at `plugin_functions.go:923`. The inline consensus routines are worse: several drops are `Debug`-only, and one (`plugin_functions.go:554-558`, conflicting merkle-root metadata) is dropped with **no log at all**. A chain missing consensus `fChain` silently falls out of *all four* downstream consensus computations (messages, roots, hashes, token data).
4. **Computed-then-discarded values everywhere.** `numPendingReports` (the execute backlog) is computed precisely and reduced to an `Infow` field; the age of the oldest pending commit report (the single best stuck-lane signal, from `CommitData.Timestamp`) is carried end-to-end and never gauged; report size/gas are computed then discarded on rejection; tokendata background queue depth and cache size are computed every round and logged at `Debug`; rate-limit cooldown remaining time is computed and dropped.
5. **Per-message failure logging with no aggregate.** A chain with 500 pending messages and a nonce gap emits ~500 `Errorw`/`Warnw` lines per round (`report.go` `CheckNonces`) and zero metrics. Every skip reason (`already executed`, `pseudo-deleted`, `token data not ready`, nonce issues, inflight, size-evicted) is log-only, and one — the inflight skip at `observation.go:452-454` — is a bare `continue` with **no log at all**.
6. **Cache liveness gaps.** The commit-report cache refresh failure degrades silently to a wider query window; the root cache's snooze/executed state and the inflight cache (whose `Delete()` is **never called in production code** — grep-verified, only tests) are entirely unobservable. A stuck cache looks identical to "no new work".
7. **No invariant-violation signal.** Execute has at least two "reached an impossible state" sites: `removeLastExecReport` on an empty builder (`builder.go:341-349`, `Error` then continue with corrupted state) and the empty-report cleanup in `getMessagesOutcome` (`outcome.go:170-174`) whose `RemoveIthElement` call is a silent no-op on the not-yet-appended element (`internal/utils.go:13-19` returns the original slice out-of-range). The commit assessment treats this subclass as independently fatal; execute has nothing.

---

## Proposed metrics at a glance — High

All High metrics are now **✅ Implemented** (beholder-only registration, same pattern as commit — see the naming note under [Cardinality guardrails](#cardinality-guardrails)). Names below are as registered in code; the `_total` suffix from the original proposals was dropped per the repo convention, and the report-builder rejections were folded into the report-validation family with `phase="build"`. Full rationale and file:line evidence live in [Full findings](#full-findings).

### Top-level plugin orchestration and liveness

| Metric | Status | Why |
|---|---|---|
| `ccip_exec_plugin_heartbeat{phase="observation"\|"outcome"}` | ✅ Implemented | Unconditional per-round liveness counter, the execute analogue of `ccip_commit_plugin_heartbeat` — the only way to tell "process wedged" from "valid but stale" for every other metric here. |
| `ccip_exec_current_state{state="unknown"\|"initialized"\|"get_commit_reports"\|"get_messages"\|"filter"}` | ✅ Implemented | The state machine is the plugin; today a plugin wedged in one state is indistinguishable from idle. Combined with the heartbeat and the age gauge below, "stuck" vs "no traffic" separates cleanly. |
| `ccip_exec_phase_errors{phase, reason}` | ✅ Implemented | The `nil, nil` swallows (cross-cutting 1) never trip `ccip_exec_errors`; this counts per-phase failures with coarse reasons (`get_messages`, `get_filter`, `filter`, `discovery`, `curse_read`, `nonce_read`, `config_digest_check`, `commit_report_cache_refresh`, `hashes_validation`, `token_data_validation`). |
| `ccip_exec_oldest_pending_commit_age_seconds{source_network_name}` | ✅ Implemented | Age of the oldest pending commit report per source chain — the best single "is this lane starving" gauge; `CommitData.Timestamp` is carried end-to-end and was discarded every round. |
| `ccip_exec_pending_reports{source_network_name}` | ✅ Implemented | `numPendingReports` was already computed in `selectReports` and reduced to an Info log field — the execute backlog number, gauged per round for every chain seen. |
| `ccip_exec_last_executed_seq_num{source_network_name}` | ✅ Implemented | The execute twin of commit's `offramp_next_seq_num` gauges: what exec believes is executed per lane, derived from `ExecutedMessages` already flowing through commit data. Closes the "already executed vs never committed" ambiguity that an unexecuted-message runbook hits at its first and last steps (see the dry run in `docs/runbooks/unexecuted-message.md`). |

### Consensus aggregation

| Metric | Status | Why |
|---|---|---|
| `ccip_exec_consensus_dropped{objectName="fChain"\|"merkle_root"\|"executed_messages"\|"onramp_address", key, reason, source_network_name}` | ✅ Implemented | The `TODO: metrics` at `plugin_functions.go:923`, plus the inline sites — a dropped fChain silently removes the lane from *all* downstream consensus; root-metadata conflicts previously dropped with no log at all. Mirrors the implemented `ccip_commit_consensus_dropped`. |
| `ccip_exec_message_consensus_conflicts{source_network_name, kind="multi_message"\|"over_two"\|"none"\|"hash"}` | ✅ Implemented | Two different messages/hashes reaching consensus for one seqNum drops the message from execution **entirely** — if persistent, it's stuck forever and blocks ordered execution behind it. Today `Errorw`-only. |

### Report acceptance / transmission validation

| Metric | Status | Why |
|---|---|---|
| `ccip_exec_report_validation_rejected{phase="should_accept"\|"should_transmit"\|"build", reason="nil_report"\|"decode_report"\|"empty_report"\|"dest_not_supported"\|"dest_support_check_error"\|"config_digest_mismatch"\|"config_digest_check_error"\|"already_executed"\|"validation_error"\|"cursed"\|"cursed_check_error"\|"skipped_nonce"\|"size_limit"\|"gas_limit"}` | ✅ Implemented | Direct analogue of the implemented `ccip_commit_report_validation_rejected`. `ShouldAcceptAttestedReport`/`ShouldTransmitAcceptedReport` aren't wrapped by `TrackedPlugin`, and `ErrInvalidReport` collapses four distinct causes into one `Infow`. The builder's `verifyReport` rejections ride the same family with `phase="build"`. "Rejected at accept" vs "rejected at transmit" vs "transmitted but reverted" are different remediation paths. |

### Message pipeline (skips, build failures, read errors)

| Metric | Status | Why |
|---|---|---|
| `ccip_exec_messages_skipped{reason="already_executed"\|"message_pseudo_deleted"\|"token_data_not_ready"\|"missing_nonce"\|"invalid_nonce"\|"missing_nonces_for_chain"\|"already_inflight", source_network_name}` | ✅ Implemented | Unifies the per-message log firehose into one low-cardinality family, tracked at `checkMessage` (the status enum maps 1:1 to reasons) plus the observation-stage inflight/executed skips. The `already_inflight` reason was previously invisible even at trace level (bare `continue`, `observation.go:452-454`). This is the metric backbone an "unexecuted message" runbook needs. |
| `ccip_exec_report_build_errors{reason="token_data_length_mismatch"\|"merkle_tree_construction"\|"merkle_root_mismatch"\|"prove_failed"\|"encode_failed"\|"empty_report", source_network_name}` | ✅ Implemented | `builder.Add` errors funnelled to one `Errorw` + `continue` in `selectReports` — a commit report failing repeatedly (e.g. token-data attach) starves its lane with no counter. `merkle_tree_construction` covers the malformed-report rejections in `report/roots.go` that existed only in error strings; `empty_report` is the previously fully-silent ready-messages-but-nothing-built signal. |
| `ccip_exec_message_read_errors{source_network_name, reason="rpc"\|"range_mismatch"}` | ✅ Implemented | One flaky source-chain read starves the lane across **all oracles** — the canonical false-idle driver for execute (`observation.go:403-411`, `:321-327`). |
| `ccip_exec_nonce_read_errors` | ✅ Implemented | A single destination-chain nonce read failure (`observation.go:551-555`) produces an observation with empty `Nonces`, which cascades into every ordered message skipping with `MissingNoncesForChain` next state — a total-execution-stop mode detectable today only by chaining log lines across two packages. |
| `ccip_exec_last_executed_seq_num{source_network_name}` | ✅ Implemented | The execute twin of commit's `offramp_next_seq_num` gauges: what exec believes is executed per lane, derivable from `ExecutedMessages` already flowing through commit data. Closes the "already executed vs never committed" ambiguity that an unexecuted-message runbook hits at its first and last steps (see the dry run in `docs/runbooks/unexecuted-message.md`). |
| `ccip_exec_commit_report_cache_last_refresh_age_seconds` + `ccip_exec_commit_report_cache_refresh_errors` | ✅ Implemented | Refresh failure degrades to a wider query window silently; a stuck cache looks identical to "no new reports" (false idle). Age-since-last-successful-refresh is the missing liveness signal. Promoted from Medium after the unexecuted-message dry run: it is load-bearing on the runbook's step1 main path (root landed on the offramp but exec never saw it). |

### Tokendata — High

All ✅ Implemented (see the naming note: registered without `_total`).

| Metric | Status | Why |
|---|---|---|
| `ccip_exec_tokendata_http_request{api="usdc"\|"lbtc"\|"cctpv2", outcome="ok"\|"not_ready_404"\|"rate_limited"\|"timeout"\|"http_error"\|"unknown_status"\|"parse_error"}` | ✅ Implemented | Today 404 (healthy "attestation pending") and 5xx (outage) collapse into the same wrapped error; non-200s had no log **or** metric at the HTTP layer at all. This one counter separates "Circle hasn't attested yet" from "Circle is down". |
| `ccip_exec_tokendata_http_cooldown_active{api}` (gauge) | ✅ Implemented | A Circle 429 triggers a ~5-minute hard cooldown during which every fetch fails; today that state is a per-request `Errorw` firehose. |
| `ccip_exec_tokendata_observer_error{observer="usdc_v1"\|"usdc_v2"\|"lbtc", stage="observe"\|"merge"}` | ✅ Implemented | The composite observer swallows per-client errors (`Observe` **always returns nil error**), so a dead USDC/LBTC API degrades every message to not-ready and never trips any error counter. |
| `ccip_exec_tokendata_background_queue_depth{observer}` / `ccip_exec_tokendata_background_cache_size{observer}` (gauges) + `ccip_exec_tokendata_background_observe_failure{observer, reason="observe_error"\|"missing_chain"\|"missing_seq"\|"not_ready"}` | ✅ Implemented | Queue depth/cache size were already computed every round and logged at `Debug`. Queue depth rising + cache flat = workers dead/slow — the canonical "stuck background worker looks healthy" gap. The failure counter covers the dequeue-without-requeue error paths. |
| `ccip_exec_tokendata_lbtc_batch_fetch{outcome="ok"\|"http_error"\|"api_error"\|"parse_error"}` / `ccip_exec_tokendata_usdc_attestation_fetch{outcome="ok"\|"not_ready"\|"rate_limited"\|"timeout"\|"parse_error"\|"http_error"\|"data_missing"}` | ✅ Implemented | LBTC absorbed a whole-batch HTTP failure into per-hash statuses and returned `nil` error (even the LBTC latency histograms recorded it as "successful"); USDC v1 had zero request-level metrics (it's not wrapped by the observed-client at all). |

---

## Proposed metrics at a glance — Medium

| Metric | Status | Why |
|---|---|---|
| `ccip_exec_invariant_violations_total{component="report_builder"\|"outcome_empty_report_cleanup"\|"encoded_size"}` | Proposed | "Reached an impossible state" subclass (cross-cutting 7); independently fatal in the commit taxonomy. |
| `ccip_exec_report_size_bytes{source_network_name}` / `ccip_exec_report_estimated_gas{source_network_name}` (gauges at finalize) | Proposed | Size/gas ceilings are the throughput throttles; the values are computed then discarded on rejection (`report.go:457-479`), and the per-message eviction loop is `Debug`-only. |
| `ccip_exec_observation_truncated_total` + `ccip_exec_observation_size_bytes` | Proposed | Truncation means the round executes a fraction of pending work; sustained truncation is a throughput-ceiling incident, today a single `Infow` per stop. |
| ~~`ccip_exec_commit_report_cache_last_refresh_age_seconds` + `ccip_exec_commit_report_cache_refresh_errors_total`~~ | ✅ Promoted to High | See the High table; promoted after the unexecuted-message dry run (`docs/runbooks/unexecuted-message.md`). |
| `ccip_exec_inflight_cache_size` + `ccip_exec_inflight_cache_marks_total{source_network_name}` | Proposed | The inflight cache is 100% unobservable and `Delete()` is never called in production — orphaned entries silently exclude messages for a full TTL. Marks-per-round is also a leading "we committed to transmitting" progress indicator. |
| `ccip_exec_commit_roots_snoozed_total` / `ccip_exec_commit_roots_marked_executed_total{source_network_name}` | Proposed | Snooze = "transmitted but not finalized"; a root stuck snoozed across rounds is a stuck lane visible only via repeated root-hash log lines. |
| `ccip_exec_cursed{scope="global"\|"dest"\|"source", source_network_name}` + `ccip_exec_curse_read_errors_total` | Proposed | Curse halts execution and today is a `Warnw` per round; the curse-*read* failure makes an oracle observe an empty commit-report set while curse-blind (silent divergence from peers). Mirrors the commit curse gauges. |
| `ccip_exec_contracts_initialized` (gauge) + `ccip_exec_discovery_errors_total{method}` | Proposed | Discovery-outcome failure leaves `contractsInitialized` false forever → every subsequent outcome is empty, log-only. A full-plugin stall visible only in logs. |
| `ccip_exec_config_digest_check_errors_total` | Proposed | When the digest *check* fails, the existing `ccip_exec_config_digest_mismatch` gauge isn't updated — a frozen `0` is ambiguous between "match" and "check broken". |
| `ccip_exec_observation_validation_rejected_total{reason="decode"\|"chains"\|"state"\|"eligibility"}` | Proposed | `ValidateObservation` rejections are counted by OCR but invisible to execute metrics; a persistently-invalid observer is a liveness/divergence signal. Freeform error strings must be bucketed into coarse enums first. |
| `ccip_exec_tokendata_attestation_client_error_total{token}` / `ccip_exec_tokendata_time_to_attestation_expired_total{token}` / `ccip_exec_tokendata_attestations_waiting{token}` | Proposed | The time-to-attestation tracking cache silently expires at 70 min (expired drops are indistinguishable from success); cache size ≈ "messages waiting for attestation" — the best attestation-backlog gauge. |
| `ccip_exec_tokendata_cctpv2_message_not_ready_total{...}` (split from the error counter) | Proposed | The existing cctpv2 error counter conflates the healthy `ErrNotReady` pending path with hex-decode/encode corruption. |

---

## Full findings

Severity is about production-triage value, independent of whether something is logged today. All file:line references verified against the tree at time of writing.

### Top-level plugin orchestration (`execute/plugin.go`, `execute/tracked.go`, `execute/outcome.go`)

**High**

- **No heartbeat / state visibility.** `tracked.go:41-94` wraps only `Query`/`Observation`/`Outcome`; `ccip_exec_latest_round_id` is set inside `TrackObservation`/`TrackOutcome` (`metrics/prom.go:212-231`) — but `TrackObservation` is skipped when a phase returns `nil, nil`, and `TrackOutcome` is skipped whenever the outcome is empty (`outcome.go:106-112`, `Infow(EmptyOutcome)` then early return). Meanwhile the empty-outcome path *encodes and returns* an `Initialized`-state outcome when contracts are initialized, so rounds do cycle — you just can't see it. A wedged pipeline (perpetual curse, contracts never initializing, always-erroring phases) freezes every gauge instead of showing a bad value. → **Proposed**: `ccip_exec_plugin_heartbeat{phase}` counter, incremented unconditionally at the top of `Observation()`/`Outcome()` before any decode or early return (exact pattern of the implemented `ccip_commit_plugin_heartbeat`), plus `ccip_exec_current_state{state}` gauged once per `Outcome()` from `previousOutcome.State.Next()` (`exectypes/outcome.go:29-47`). The three signals (heartbeat advancing / state cycling / oldest-pending age growing) separate "no traffic" from "stuck" cleanly; today both look identical.
- **`nil, nil` swallows and partial degradations.** See cross-cutting 1 for the full list; the two worst: `getFilterOutcome` failure (`outcome.go:93-99`) discards an entire Filter round's built reports with no error and no tracking; discovery `Outcome()` error (`outcome.go:74-80`) leaves `contractsInitialized` false, so **every** subsequent outcome is empty `Initialized` forever — a total stall whose only trace is an `Errorw` per round. → **Proposed**: `ccip_exec_phase_errors_total{phase, reason}`.
- **Report-acceptance / transmission-validation rejections are unmetriced and the methods are unwrapped.** `validateReport` (`plugin.go:667-731`, called from both `ShouldAcceptAttestedReport` and `ShouldTransmitAcceptedReport`) has distinct rejection paths: nil report (`:673-675`), empty report (`:684-688`), dest chain not supported (`:697-704` — the transmission-schedule misconfig), `errOffRampConfigMismatch` (`:705-709`), `errAlreadyExecuted` (`:716-723`), transient validation errors (`:724-727`), plus curse checks via `internal/plugincommon/curses.go:35,50` (`Warnw`-only). Four of these collapse into one `plugincommon.ErrInvalidReport` and surface to callers as a single `Infow("report not valid, not accepting/transmitting")` (`plugin.go:847-851`, `:887-891`). This sits directly upstream of transmission: rejections at accept vs transmit have zero on-chain trace, whereas "transmitted but reverted" has some — without per-reason counts they're indistinguishable from each other. → **Proposed**: `ccip_exec_report_validation_rejected_total{phase, reason}` mirroring the implemented commit counter.
- **Oldest-pending-commit age never computed.** `CommitData.Timestamp` (`exectypes/commit_data.go:16-17`) is carried end-to-end and used for the lane-starvation-critical `LessThan` ordering, but `TrackOutcome` (`metrics/prom.go:222-231`) extracts only max seq nums. → **Proposed**: `ccip_exec_oldest_pending_commit_age_seconds{source_network_name}` = `now − min(cr.Timestamp)` over `outcome.CommitReports`, computed in `TrackOutcome` where the data is already in hand.

**Medium**

- **`selectReports` backlog and builder failures** (`plugin.go:498-541`): `builder.Add` error → `pendingReports++`, `Errorw(UnableToAddReportToBuilder)`, `continue` — the commit report silently drops from the candidate set and the outer call returns no error. The `pendingReports` total (incomplete observations + builder failures + partially-executed reports) is computed precisely and only logged (`plugin.go:536-539`). There is also a sibling `// TODO: count pending reports during Build()` at `plugin.go:533`: reports discarded *inside* `builder.Build()` (size/gas verification failures) are invisible to both log and metric. → **Proposed**: `ccip_exec_pending_reports{source_network_name}` gauge + `ccip_exec_report_build_errors_total{reason}`.
- **Snoozed/executed root skips** (`plugin.go:179-198`): skipped commit roots collected into `skippedCommitRoots`, `Infow`-only. Covered under the root-cache metrics below.
- **`ValidateObservation` unwrapped** (`plugin.go:352-408`): every rejection returns an error OCR counts internally but execute never metrics. → **Proposed**: `ccip_exec_observation_validation_rejected_total{reason}` with coarse enumerated reasons (freeform `fmt.Errorf` text must not become labels).

### Observation phase (`execute/observation.go`)

**High**

- **Per-report message-read failures** (`:394-412`): `readMessagesForReport` error → `Errorw("unable to read all messages for report")` → `continue`; the commit report drops from the observation, outer call returns nil. Sibling at `:321-327`: `msgsConformToSeqRange` failure (`Errorw("missing messages in range")`) — the log-poller-returned-fewer-messages-than-committed symptom. If all oracles hit this on a lane, the lane silently vanishes from every observation: the canonical false-idle mechanism. → **Proposed**: `ccip_exec_message_read_errors_total{source_network_name, reason="rpc"|"range_mismatch"}`.
- **Nonce-read failure cascade** (`:551-555`): `p.ccipReader.Nonces` error → `Errorw("unable to get nonces")` → observation returned **without** `Nonces`. Downstream, `report.CheckNonces` (`report/report.go:260-266`) marks every sequenced message `MissingNoncesForChain` → the entire Filter outcome produces no report, round after round. Root cause one line, symptom N log lines per message. → **Proposed**: `ccip_exec_nonce_read_errors_total` (no labels); the downstream skip volume lands in `ccip_exec_messages_skipped_total{reason="nonces_missing_for_chain"}` for attribution.

**Medium**

- **Config-digest check failure leaves the gauge stale** (`:62-75`): on `configDigestErr`, `TrackConfigDigestMismatch` is **not called** — `ccip_exec_config_digest_mismatch` freezes at its last value; a frozen `0` is ambiguous between "digests match" and "check broken". → **Proposed**: `ccip_exec_config_digest_check_errors_total`; do not duplicate the existing gauge.
- **Curse-read failure and curse state** (`:230-241`): read failure → `Errorw` then `return observation, nil` — this oracle observes **no commit reports** this round while curse-blind (silent divergence); actual curse state (`GlobalCurse || CursedDestination`) → `Warnw("nothing to observe: rmn curse")` every round. → **Proposed**: `ccip_exec_curse_read_errors_total` + `ccip_exec_cursed{scope, source_network_name}` gauge (commit precedent).
- **Commit-report-cache refresh failure** (`:222-228`): `Errorw("...proceeding with potentially wider query window")` and continue — degraded mode means heavier RPC load and latency spikes; see Caches section for the full writeup.
- **Token-data observation failure poisons the message** (`:285-307`): observer error → `Errorw` + `NewErrorTokenData(err)` — the fetch error *launders into* "waiting" state; by the time `CheckTokenData` runs, "attestation pending" vs "observer broken" is indistinguishable. → **Proposed**: `ccip_exec_token_data_observer_errors_total{source_network_name}` (and see Tokendata section for the client-side split).
- **Inflight/already-executed silent skips and truncation** (`:449-489`): inflight skip is a bare `continue` (`:452-454`, no log); already-executed skip likewise (`:455-457`); size-limit demotion replaces the message with an empty stub and stops (`:471-477`), with one `Infow` summary. Sustained truncation = the round executes a fraction of pending work. Note the `observation.go:480` TODO acknowledges fully-executed inflight entries are never cleaned up — orphaned entries silently exclude messages for the full TTL. → **Proposed**: `ccip_exec_observation_truncated_total`, `ccip_exec_observation_size_bytes`, and `reason="already_inflight"` under `ccip_exec_messages_skipped_total`.

### Outcome phase (`execute/outcome.go`)

**High**

- **Empty-outcome early return skips all tracking** (`:104-112`): healthy idle and wedged-idle both produce zero metric deltas (covered by the heartbeat/state gauge proposals above).
- **Validation resets** (`:146-154`): `validateHashesExist`/`validateTokenDataObservations` failure → `Errorw`, `return exectypes.Outcome{}` — the pipeline silently restarts at GetCommitReports; a persistent failure is an infinite GetCommitReports→GetMessages loop, log-only. → **Proposed**: `ccip_exec_phase_errors_total{phase="outcome_get_messages", reason="hashes"|"token_data"}`.

**Medium — invariant violation**

- **Empty-report cleanup is a silent no-op** (`:170-174`): when `len(report.Messages) == 0` the code calls `commitReports = internal.RemoveIthElement(commitReports, i)` — but element `i` hasn't been appended yet (it is appended unconditionally at `:174`), and `RemoveIthElement` (`internal/utils.go:13-19`) returns the original slice on out-of-range indices. Net effect: the intended removal **never happens**; empty-message commit reports flow into the Filter state and consume builder work. No log, no metric. Related: `EncodedSize` (`internal/utils.go:5-11`) returns `0` on marshal error, silently under-reporting observation sizes. → **Proposed**: `ccip_exec_invariant_violations_total{component}` (and fix the bug — the metric would surface whether this path fires in prod at all).

### Consensus computation (`execute/plugin_functions.go`, shared `internal/plugincommon/consensus`)

**High**

- **The literal TODO: `getConsensusFChain`** (`plugin_functions.go:909-927`, TODO at `:923`). `fChain, _ := consensus.GetConsensusMap(...)` discards the `dropReasons` map. Per `consensus/drop_reason.go:9-20` the three reasons are `threshold_not_defined` (config bug), `insufficient_agreement` (< 2f+1 oracles observed fChain — participating-node loss), `split` (oracles disagree — integrity signal). Blast radius: a chain missing consensus fChain causes **every** downstream consensus routine to skip that chain — `prepareValidatorsForComputeMessageObservationsConsensus` (`:510-514`), `computeCommitObservationsConsensus` (`:544-546`, `:562-564`), `computeMessageHashesConsensus` (`:687-690`), `computeTokenDataObservationsConsensus` (`:751-754`), all `Warnw` "no F defined for chain" per round. The lane's messages, roots, hashes, and token data all silently drop out of the consensus observation. This is the execute equivalent of the commit fChain incidents. → **Proposed**: `ccip_exec_consensus_dropped_total{objectName="fChain", key, reason, source_network_name}`; implementation is trivial — thread the reporter in and loop the already-returned map (commit implemented the identical shape at `commit/plugin.go:605`). The shared-infra fix documented in [commit-metrics.md](commit-metrics.md#shared-consensus-aggregation-internalplugincommonconsensusconsensusgo) applies unchanged; execute has exactly **one** `GetConsensusMap` call site.
- **Inline consensus drops, Debug-only or silent** (`plugin_functions.go:535-605`):
  - `:543-552` — merkle root below f+1 votes: `Debugw` (invisible in prod). A root that never reaches f+1 means the commit DON's work never becomes executable.
  - `:554-558` — **silent drop with no log at all**: the `seenCount[mr.MerkleRoot] == 1` filter removes valid roots whose identical root was observed with conflicting metadata (different OnRampAddress/timestamp/range). A consensus split disappears without a trace.
  - `:569-575` — executed-message votes below f+1: `Debugw`.
  - `:581-585` — base64 onRampAddress decode failure: `Errorw` + `continue`, whole root dropped.
  → **Proposed**: fold into `ccip_exec_consensus_dropped_total{objectName="merkle_root"|"executed_messages"|"onramp_address", ...}` — these sites are inline (not `GetConsensusMap`) so instrumentation is manual, but the vote counts are already computed in place.
- **Message/hash consensus conflicts** (`:444-487`, `:684-717`): two *different* full messages reaching consensus for one seqNum (`:469-473`, `:482-486`) or two different hashes for one message (`:708-714`) → `Errorw` + the message is **dropped from execution entirely**; if persistent, it's stuck forever and its nonce blocks ordered execution behind it. Per-message consensus misses are `Debugw` (`:448-450`). → **Proposed**: `ccip_exec_message_consensus_conflicts_total{source_network_name, kind="multi_message"|"over_two"|"none"|"hash"}`. Never label by seqNum or messageID.

**Medium**

- **Token-data consensus failures** (`:749-786`): chain skipped for insufficient oracle coverage → `Infow` (`:757-760`) — every message on that lane becomes `TokenDataNotReady` in the builder → lane stalls; per-token validator returning ≠1 value → `NotReadyToken()` appended with **no log at all** (`:775-786`). → **Proposed**: `ccip_exec_token_consensus_failures_total{source_network_name, level="chain"|"token"}` (count, don't gauge — token-level events can be numerous). Ready/waiting totals are already covered by `ccip_exec_output_sizes{type=tokenReady|tokenWaiting}`; don't duplicate.
- **Nonce-consensus drops** (`:854-859`): chain/sender pairs below threshold, `Debugw`-only — a bare counter suffices; **never label by sender address**.

### Report building (`execute/report/`)

**High**

- **The `builder.Add` funnel** (`builder.go:189-245` → `plugin.go:517-522`): every per-chain report death (nonce-validation failure `:197-199`, token-data length mismatch `:206-209`, `tryBuildReport` failure `:232-234`) funnels to one `Errorw` + `continue` with the reason surviving only in an error string. `ccip_exec_errors` never fires (Outcome returns no error). → **Proposed**: `ccip_exec_report_build_errors_total{reason}` — reason enum extended with the four malformed-report rejections from `roots.go:19-46` (message-count/range mismatch, hash-count mismatch, seqNr outside range, wrong source chain, empty hash — the last one an invariant per message) which today exist only in error strings.
- **`ErrEmptyReport` after successful checks is fully silent** (`report.go:602-604`): no log at the return site; consumed at `builder.go:258-287` where both the `!multipleReportsEnabled` and max-count-reached branches return silently. Ready messages existed, checks passed, yet nothing could be built — the strongest possible "report builder is wedged" signal emits **nothing**. (Side note: `report/errors.go` `ErrNotReady` is dead code — a separate `tokendata.ErrNotReady` is used instead.) → **Proposed**: `ccip_exec_report_build_errors_total{reason="empty_report"}` + `ccip_exec_report_count_limit_hits_total`.
- **Per-message skip taxonomy** (the runbook backbone): every reason is log-only today, one line per message per round:
  - `already_executed` — `report.go:184-198`, `Infow`;
  - `pseudo_deleted` — `report.go:167-181`, `Infow`;
  - `token_data_not_ready` — `report.go:201-229` (`:210`, `Infow`; note `:204` index-out-of-range is a separate `Errorw` that aborts the *whole chain report*);
  - `missing_nonce` / `invalid_nonce` / `nonces_missing_for_chain` — `report.go:248-307` (`Errorw`/`Errorw`/`Warnw` per message — 500 messages with a nonce gap = 500 log lines per round);
  - `already_inflight` — `observation.go:452-454`, **no log at all**; and note `report.CheckIfInflight` (`report.go:309-324`) is **dead code** — never wired into the builder's `defaultChecks` (`builder.go:85-89`) or the Filter checks (`outcome.go:203`);
  - `deferred` (ready but evicted from the report by the greedy size/gas/nonce-continuity loop) — `report.go:591-599`, **`Debugw`-only**.
  → **Proposed**: `ccip_exec_messages_skipped_total{reason, source_network_name}`.
- **Report-level validation rejections** (`report.go:374-479`): `verifyReportNonceContinuity` rejection → `Infow("invalid report, skipped nonce detected")` → `valid=false, nil` — swallowed, no error (`:435-441`); the "assumption violation: bug in the code" branch (`:413`) is likewise reduced to an Info log. Size-limit rejection (`:457-462`) and gas-limit rejection (`:464-479`) compute `len(encoded)`/`totalGas` and discard them. All constituent messages were individually "ready" — a whole-report rejection with no aggregate. → **Proposed**: `ccip_exec_report_validation_rejections_total{reason="skipped_nonce"|"size_limit"|"gas_limit", source_network_name}` + gauges `ccip_exec_report_size_bytes` / `ccip_exec_report_estimated_gas` at finalize.

**Medium**

- **Max report count reached** (`builder.go:224-229`, `Debugw`; silent truncation in `Build()` at `:365-371`): pending work explicitly load-shed, invisible. → covered by `ccip_exec_report_count_limit_hits_total` above.
- **`removeLastExecReport` invariant violation** (`builder.go:341-349`): `Error("Attempted to remove last exec report, but no reports exist")` then plain `return` — execution continues with corrupted builder state. → `ccip_exec_invariant_violations_total{component="report_builder"}`.

### Caches (`execute/internal/cache/`)

**High**

- **Commit-report cache: refresh failure and staleness are unobservable** (`commit_report_cache.go:146-148`, consumed at `observation.go:222-228`): per-refresh `Errorw` only; no last-success age, no error counter. If refresh fails persistently (RPC outage), the cache stops seeing new commit reports and the plugin looks idle while messages pile up — false idle. The query-timestamp derivation (`:211-247`) computes `now − latestFinalizedReportTimestamp` and logs it at `Debugw` — a direct commit-DON staleness proxy, discarded. → **Proposed**: `ccip_exec_commit_report_cache_last_refresh_age_seconds` (gauge) + `ccip_exec_commit_report_cache_refresh_errors_total`. Also: fetched/added counts are logged only when `addedCount > 0` (`:200-207`), so a reader returning empty forever is invisible → `ccip_exec_commit_report_cache_size` gauge; dedup drops (`:294-324`) and rootless-report/key-gen skips (`:177-190`) → `ccip_exec_commit_report_cache_skipped_total{reason}`.

**Medium**

- **Commit-root cache lifecycle** (`commit_root_cache.go:81-102`): `MarkAsExecuted` and `Snooze` are `Infow`-only (log fields correctly include the merkle root — which must **not** become a label). `CanExecute` returns a bare bool; `plugin.go:179-198` logs skipped roots as a list of root strings, no count, no reason. A root stuck snoozed across many rounds is a stuck lane visible only via repeated log lines. → **Proposed**: `ccip_exec_commit_roots_snoozed_total` / `ccip_exec_commit_roots_marked_executed_total{source_network_name}`, plus `ccip_exec_commit_roots_denied_total{reason="executed"|"snoozed"}` if `CanExecute` is extended to say which cache hit.
- **Inflight message cache: 100% unobservable** (`inflight_message_cache.go`): zero logging in the whole file; `MarkInflight` (`:33-35`), `IsInflight` (`:28-31`), and `Delete` (`:37-39` — **never called in production code**, grep-verified, only tests). A stale/orphaned entry (e.g. crash between mark and transmit) silently excludes messages from observations for the full TTL with zero signal; the `observation.go:259` TODO acknowledges fully-executed entries are never cleaned up. → **Proposed**: `ccip_exec_inflight_cache_size` gauge + `ccip_exec_inflight_cache_marks_total{source_network_name}` (marks-per-round is also a leading "we committed to transmitting" progress indicator); skip volume lands in `ccip_exec_messages_skipped_total{reason="already_inflight"}`.

### Tokendata (`execute/tokendata/`)

**Existing coverage (so we don't re-propose it):**

- `tokendata/observed.go`: `ccip_attestation_request_duration{token}` and `ccip_time_to_attestation{token}` histograms — but the wrapper wraps **only LBTC** (`lbtc/attestation.go:69-75`); USDC v1 is constructed directly (`usdc.go:67`) and has zero request-level metrics. Two defects in the wrapper itself: (a) `observed.go:74-88` records the duration of the **entire batch** `Attestations()` call, once **per attestation in the result** — both latency and sample count are misleading; (b) on delegate error (`:75-78`) it returns early recording nothing.
- `tokendata/cctp/v2/metrics.go`: `ccip_exec_tokendata_cctpv2_attestation_api_latency`, `..._observe_latency`, `..._deposit_hash_calculation_error`, `..._message_to_token_data_error`, `..._assign_token_data_failure` — all with bounded, low-cardinality labels. Best-instrumented corner of the subsystem; residual gaps below.

**High**

- **The composite observer defeats every error counter.** `observer/observer.go:154-166`: per-client `Observe` error → `Error` log → `continue`; `Observe` itself **always returns nil error** (`:167`), so the plugin-level generic instrumentation never fires. A dead USDC/LBTC API degrades every supported token in every message to not-ready and looks identical to pending attestations. → **Proposed**: `ccip_exec_tokendata_observer_error_total{observer, stage="observe"|"merge"}`.
- **HTTP layer: no outcome visibility.** `http/http.go`: rate-limit cooldown drops all requests with a per-request `Errorw` (`:181-187`, set at `:239-262`; remaining cooldown time computed at `:267` and discarded); non-200 responses have **no log and no metric at this layer** (`:228-233` — 404 → `ErrNotReady`, anything else → `ErrUnknownResponse`, silently; the HTTP status is computed at `:221` and discarded by every caller except cctpv2 — USDC v1 literally assigns it to `_` at `usdc/attestation.go:129`); timeouts (`:200-216`) surface only as wrapped attestation errors downstream. The operationally critical distinction — 404 = healthy "Circle hasn't attested yet" vs 5xx = outage vs 429 = rate limited — is unanswerable without log parsing. → **Proposed**: `ccip_exec_tokendata_http_request_total{api, outcome}` + `ccip_exec_tokendata_http_cooldown_active{api}` gauge.
- **LBTC batch failure swallowed into `nil` error.** `lbtc/attestation.go:134-141`: `Post` error → every hash in the batch gets `ErrorAttestationStatus(err)` → **`return attestations, nil`**. `ObservedAttestationClient` sees `err == nil`, so even `ccip_attestation_request_duration` fires with a "successful" latency while every message failed. API-level error codes (`:147-153`) are absorbed the same way. → **Proposed**: `ccip_exec_tokendata_lbtc_batch_fetch_total{outcome="ok"|"http_error"|"api_error"|"parse_error"}`.
- **USDC v1 has zero metrics.** Per-attestation fetch failures are `Errorw` + error absorbed into status (`usdc/attestation.go:129-138`); a missing attestation entry degrades silently to `ErrDataMissing` with **no log whatsoever** (`usdc.go:249-252`); the event-reader RPC failure fails the whole observer (`usdc.go:182-192`) and is swallowed by the composite. → **Proposed**: `ccip_exec_tokendata_usdc_attestation_fetch_total{outcome}` + `ccip_exec_tokendata_usdc_event_read_error_total{source_network_name}`.
- **Background worker: stuck-looks-healthy.** `observer/observer_background.go`: failed messages are dequeued and **not re-enqueued** (`:194-212`, three `Errorw`/`Infow` branches; the `reprocessInterval` field at `:28,61` is dead code — retry happens only when the next plugin round re-observes); queue depth and cache size are computed every round and logged at `Debug` (`:76-80`, sizes at `:328-332`, `:385-389`). Queue depth rising + cache size flat = workers dead/slow. Bonus hazard no metric can currently reveal: `enqueue` sends on an **unbuffered** channel (`:277`) after releasing the lock, so the composite `Observe` can block head-of-line when all workers are mid-`Observe`. The cache-expiry goroutine (`:392-423`) has no panic recovery and Debug-only logging — if it dies, entries leak forever silently. → **Proposed**: `ccip_exec_tokendata_background_queue_depth{observer}` + `ccip_exec_tokendata_background_cache_size{observer}` gauges, `ccip_exec_tokendata_background_observe_failure_total{observer, reason}` counter.

**Medium**

- **Time-to-attestation cache silently expires** (`token_data.go:16` 70-min expiration; `observed.go:96-119`): a message pending >70 min falls out of the cache; when it finally attests the measurement is dropped (`:109`) with no log and no counter. Cache size ≈ "messages waiting for attestation" — the best backlog gauge, computed implicitly and never exported. → **Proposed**: `ccip_exec_tokendata_time_to_attestation_expired_total{token}`, `ccip_exec_tokendata_attestations_waiting{token}`, plus `ccip_exec_tokendata_attestation_client_error_total{token}` for the wrapper's early-return path.
- **cctpv2: `ErrNotReady` inflates the error counter** (`observer.go:289-291` counts any `tokenData.Error != nil`, including the healthy pending path `:310-317`) — pending backlog and hex-decode/encode corruption (`:321-361`) share one series. → **Proposed**: `ccip_exec_tokendata_cctpv2_message_not_ready_total` (or an `outcome` label on the existing counter). Also: silent skips in `getCCTPv2RequestParams` (`:159-198`, no-tx-hash and payload-decode failures — the message can never become ready through this observer) → `ccip_exec_tokendata_cctpv2_request_param_skip_total{reason}`; and the success path is uncounted (`observer.go:424-429`), so a fully-quiet healthy DON is indistinguishable from one with no USDC traffic → `ccip_exec_tokendata_cctpv2_assign_token_data_success_total`.
- **Ongoing pending-attestation age is never computed** — `ccip_time_to_attestation` covers LBTC success paths only; nothing measures how long a hash has been pending. → **Proposed**: `ccip_exec_tokendata_pending_attestation_age_seconds` histogram (observed per round per still-pending hash; `token` label only).

---

## Cardinality guardrails

**Naming note on the `_total` suffix.** Verified empirically against the pinned `prometheus/client_golang` (counter registered as `foo`, `foo_total`, `foo_total_total` and dumped through `promhttp`): the Go client exposes counters under **exactly the registered name** — nothing is appended, so `foo_total` in code can never become `foo_total_total` on the wire. All existing `ccip_exec_*`/`ccip_commit_*` counters are registered without `_total` and are exposed that way on the promhttp path (e.g. `ccip_exec_output_sizes`). Two consequences: (a) the `_total` suffix in the proposed names in this doc is advisory — register new counters without it to match the existing series; (b) the beholder/OTLP path *is* suffixed downstream (the platform-side OTel→Prometheus conversion appends `_total` to counters per spec, without doubling), so the same logical counter can appear as `ccip_exec_errors` on the promhttp endpoint and `ccip_exec_errors_total` on the OTLP path. VictoriaMetrics' MetricsQL `rate()`/`increase()` auto-append `_total` at query time when missing, so rate queries are insensitive to the choice; exact-match queries are the only convention-sensitive ones.

Applies to everything above, and to any future execute metrics:

- **Safe label axes** (all bounded): `phase`, `reason`, `state`, `kind`, `objectName`, `component` (fixed enums); `source_network_name` / chain family+ID (bounded by configured chains); `observer`/`api` adapter type (≤3); `outcome`/`http_status` (fixed enums — note `http.go`'s `callAPI` coerces everything to 400/408, so realistic status values are ~6).
- **Never label** (unbounded; they exist throughout these files and must stay in log fields): `messageId`, `txHash`, `messageHash`/`depositHash`, attestation hex, merkle root, onramp address, `sender`, sequence numbers, and freeform `fmt.Errorf` strings (several rejection paths — `validateReport`, `checkAlreadyExecuted` — need enumerated coarse reasons before metering). Per-message drill-down joins via `(plugin, round/ocrSeqNr)` in logs, exactly as in the commit runbooks.
- **Don't duplicate existing series**: `ccip_exec_output_sizes` already counts what made it into the outcome (`messages`, `commitReports`, `tokenReady`/`tokenWaiting`); `ccip_exec_max_sequence_number`, `ccip_exec_latest_round_id`, `ccip_exec_config_digest_mismatch`, `ccip_exec_loopp_ccip_provider_supported` are all kept as-is — the proposals above either complement them (heartbeat vs `latest_round_id`) or count what fell out before them.

---

## Follow-ups

Mirroring the commit workflow, the natural next steps once the High items are implemented:

1. An **execute-plugin-health** runbook (the analogue of `docs/runbooks/commit-plugin-health.md`): step 0 = heartbeat + state gauge, then phase errors, then per-lane skipped-message breakdown, then tokendata client health.
2. An **unexecuted-message** runbook (the analogue of `docs/runbooks/uncommitted-message.md`), tracing a known `(sourceChain, seqNr)` through commit-finalization → observation → consensus → report building → transmission — the "expected work" v2 idea from the commit doc applies verbatim here.
3. The shared `consensus.go` `(result, dropReasons)` fix documented in [commit-metrics.md](commit-metrics.md) covers execute's one shared-infra call site automatically; the six discovery-processor `// TODO: report metrics` siblings (`internal/plugincommon/discovery/processor.go:298,308,329,347,367,388`) should ride the same change.
