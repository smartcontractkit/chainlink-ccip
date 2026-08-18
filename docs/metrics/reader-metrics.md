- [Metrics coverage assessment — commit data-source layer (reader + accessor)](#metrics-coverage-assessment--commit-data-source-layer-reader--accessor)
   * [What exists today](#what-exists-today)
   * [Cross-cutting themes (repeat across both layers)](#cross-cutting-themes-repeat-across-both-layers)
   * [Proposed metrics at a glance — reader layer](#proposed-metrics-at-a-glance--reader-layer)
   * [Proposed metrics at a glance — accessor / price layer](#proposed-metrics-at-a-glance--accessor--price-layer)
   * [Proposed metrics at a glance — config poller](#proposed-metrics-at-a-glance--config-poller)
   * [How this extends the runbooks (all additive)](#how-this-extends-the-runbooks-all-additive)
   * [Full findings — CCIP reader (`pkg/reader`)](#full-findings--ccip-reader-pkgreader)
   * [Full findings — default chain accessor & price reader (`pkg/chainaccessor`)](#full-findings--default-chain-accessor--price-reader-pkgchainaccessor)
   * [Full findings — config poller (`pkg/reader/config_poller_v2.go`)](#full-findings--config-poller-pkgreaderconfig_poller_v2go)

# Metrics coverage assessment — commit data-source layer (reader + accessor)

This is the counterpart to `commit-metrics.md` for the data that *goes into* the commit
processors. Where that doc covered what the processors compute and discard, this one covers
the reads that feed them: the CCIP reader (`pkg/reader/ccip.go`, interface
`pkg/reader/ccip_interface.go`) and the chain accessor it sits on
(`pkg/chainaccessor/default_accessor.go`, `config_processors.go`, `default_price_reader.go`).

**Ground rule (same as `commit-metrics.md`):** the presence of a log line — even a "stable
identifier" some log-analysis tooling keyed on — is never a reason to exclude a signal. Worse,
much of what matters here is only at `Debugw`, or recorded only as an `Infow` on an *empty*
result, i.e. effectively invisible in production. Most of the failure paths found below convert
an error into an empty/partial result with a `nil` error, so the read looks "fine, just quiet."
Logs are the last resort and the primary excuse for why this layer has been so hard to debug;
this is the layer where "healthy but idle" usually originates.

**Scope note that shapes naming:** this layer is **shared with execute** (`CommitReportsGTETimestamp`,
`ExecutedMessages`, `Nonces`, `Sync` are used only by execute; the rest feed commit). Measurements
here must therefore be **plugin-generic** — `ccip_reader_*` / `ccip_chainaccessor_*` with a
`query`/`method` label — *not* `ccip_commit_*`. This is the same principle as the shared
`consensus.go` fix in `commit-metrics.md`: don't hardcode a plugin prefix into code that another
plugin also runs.

## What exists today

Three instrumentation layers, and one of them has none at all. Every one of the three that does
exist is `promauto` (self-hosted Prometheus scrape only) — **none of them reach beholder**, which
is the only channel NOP-run nodes emit into production. That single fact is the biggest gap.

| Layer | File | Metrics today | Blind spot |
|---|---|---|---|
| Contract reader (below accessor) | `pkg/contractreader/observed.go` | `CrErrors`, `CrBatchRequestsDurations`, `CrDirectRequestsDurations`, `CrBatchSizes` (promauto) | Only fires on a **top-level** `err` return. Blind to the accessor's swallowed→empty/`nil` paths and to per-row/per-chain `continue` drops. |
| Chain accessor | `pkg/chainaccessor/*` | **none** | 100% logs-only. This is the layer with the worst silent-empty producers. |
| CCIP reader | `pkg/reader/observed_ccip.go` | `ccip_reader_query_duration{query}` histogram, `ccip_reader_data_set_size{query}` gauge, `ccip_reader_chain_fee_components` gauge (promauto) | Duration + *result-size* only — never the actual values, never errors, never staleness, never per-chain partials. Its own comment (`observed_ccip.go:67-68`) calls the accessor layer a future TODO: *"Every CCIPReader (and Chain Accessor Layer in the future) should be decorated with the observedCCIPReader."* |
| Config poller (background cache) | `pkg/reader/config_poller_v2.go` | **none** | Cache-age is computed and thrown away (`Debugw`); freshness timestamps are internal-only; the only health signal is a `Warnw`-and-service-health-buffer counter (`consecutiveFailedPolls`). Nothing is exported; a stale or wedged poller is invisible. |

`observedCCIPReader` wraps the live reader (`NewCCIPChainReader`, `pkg/reader/ccip_interface.go:83-113`)
and measures via `withObservedQueryAndResult` (`observed_ccip.go:332-359`): latency always counted,
data size only when `err == nil` and a size fn exists. Error is only `Debugw`. So even the one thin
promauto layer that exists tells you *how slow a read was*, never *what went wrong*.

## Cross-cutting themes (repeat across both layers)

1. **Error → empty/`nil` conversion (the false-idle factory).** ~25 distinct paths in the
   accessor turn a failure into an empty/partial result with a `nil` error (catalogued in the full
   findings). The commit plugin then reads "nothing" exactly as it reads "quiet chain."
2. **Batch fan-out swallows partials.** Cross-chain reads run one goroutine per chain
   (`pkg/reader/ccip.go:459,513`); each skipped chain is a log line at best.
3. **Two silent cache layers.** The `configPoller` (30s background poll; `reader/ccip.go:33`)
   caches router addresses, curse state, RMN remote config, source chain configs;
   the chainfee **async observer** (`commit/chainfee/observation_async.go:180-302`) freezes the
   *last successful* fee/price snapshot in memory on reader failure
   (`Debugw "returning cached value"`). Neither signals staleness, and the config poller's freshness
   timestamps — its entire reason for existing — are never exposed (see the
   [config-poller findings](#full-findings--config-poller-pkgreaderconfig_poller_v2go)).
   A stuck fee pump or a stale but "healthy" poller looks perfectly fine. This was already flagged
   in `commit-metrics.md`; the root cause is here, and it is un-instrumented.
4. **Empty-vs-broken is not distinguishable** for every read that returns `[]` / `0` / empty map
   with a `nil` error.
5. **Zero beholder anywhere**; the accessor has zero metrics of any kind.

## Proposed metrics at a glance — reader layer

All plugin-generic, keyed by `{queryName}`. ✅ = implemented; everything else is proposed.

All chain-identity labels below come from `pkg/reader/rcmetrics`'s private `chainAttrs()` helper
(mirrors `commit/metrics/prom.go`'s `destChainAttrs()`/`sourceChainAttrs()` pattern, but with no
dest/source prefix — every one of these metrics describes exactly one chain, never a lane): a
selector resolves to `chainID`, `chainFamily`, `chainName`, and `chainSelector` (the selector
value itself, as a string), replacing the old bare numeric-selector `chain` label. Every
`Record*` method on `ReaderMetrics`/`ObservedReaderMetrics`/`ConfigPollerMetrics`/`AccessorMetrics`
now takes a `chainSelector ccipocr3.ChainSelector` parameter directly instead of a pre-formatted
string, so callers never need `rcmetrics.ChainLabel(...)` (removed) at the call site.

| Metric | Status | Why |
|---|---|---|
| `ccip_reader_read_outcome{chainID,chainFamily,chainName,chainSelector,query,outcome="ok"\|"empty"\|"error"}` | ✅ Implemented | Per-query classification ok/empty/error — the false-idle primitive (empty-with-no-error). **Chain identity here is the reader instance's DESTINATION chain** (recorded once per query at the `observedCCIPReader` wrapper level), not whichever chain the query happens to be about — the one semantic exception to "chain = the chain this read is about" that every other row in this table follows. Since this metric now shares the exact same label *keys* as every other reader metric (no more `chain` vs `chainID` naming to flag it), don't assume `chainID` here means the same thing it means on `chain_gap`/`read_empty` two rows down — it doesn't. |
| `ccip_reader_read_empty{chainID,chainFamily,chainName,chainSelector,query}` | ✅ Implemented | "Returned nothing with no error" — the single-chain false-idle primitive, wired at the concrete reads (`MsgsBetweenSeqNums`, `LatestMsgSeqNum`, `GetChainFeePriceUpdate`) where the source chain is known. Chain identity is whichever chain that specific read was about (source or dest, depending on the query). |
| `ccip_reader_chain_gap{chainID,chainFamily,chainName,chainSelector,query,state}` | ✅ Implemented | Count of source chains in each state (`returned` vs `not_found`/`disabled`/`misconfigured`/`error`/`invalid`/`missing`/`stale`/`count_mismatch`/...). `count_mismatch` = a message read whose count doesn't match its requested range (partial/incomplete). This is the per-lane subset signal AND its reason. (Consolidated: the originally-proposed `ccip_reader_read_partial` counter was folded into `chain_gap{state!="returned"}` to avoid duplicate instrumentation.) |
| `ccip_reader_chain_fee_components{chainID,chainFamily,chainName,chainSelector,feeType}` | ✅ Implemented | Beholder port of the existing promauto gauge (chainFee `exec`/`da`) — now NOP-visible. The promauto side (`observed_ccip.go`'s `PromChainFeeGauge`) is a separate, unchanged registration still keyed on bare `chainFamily,chainID` (old scheme) — the two paths use different label sets, same as `commit/metrics/legacy_prom.go`'s relationship to `commit/metrics/prom.go`. |
| `ccip_reader_read_all_ok{query}` | Proposed | Folded into `read_outcome{outcome="ok"}`; not a separate instrument. |
| `ccip_reader_msg_dropped{query,reason}` | Instrument added; no call site yet | Reads that returned data but dropped rows on validation/cast. Registered but not yet recorded anywhere — wire when the accessor wrapper lands. |
| `ccip_reader_query_last_success_timestamp{query}` | Proposed | Per-member staleness for cache-feeding reads; the config-poller last-success timestamp below covers the poller side only. |
| `ccip_reader_txhash_blanked_total` | Proposed (Low) | A deliberate data wipe (`populateTxHashEnabled=false`) with no count today. |

## Proposed metrics at a glance — accessor / price layer

| Metric | Status | Why |
|---|---|---|
| `ccip_chainaccessor_batch_result_total{source,operation,outcome="ok\|empty\|partial\|error"}` | Proposed | One counter per batch call that explicitly tags the silent-empty outcomes (notably `GetChainFeePriceUpdate` → empty/`nil`). |
| `ccip_chainaccessor_row_drops_total{source,operation,reason}` | Proposed | Aggregates the ~20 per-row/per-chain/per-token `continue` drops (Nonces per-addr, fee-update per-chain, feed-price per-contract, event per-msg). |
| `ccip_chainaccessor_empty_read_total{source,operation}` | Proposed | The "returned empty, `nil` error" set (`LatestMessageTo`, `ExecutedMessages`, `MsgsBetweenSeqNums`, `CommitReportsGTETimestamp`) — the false-idle primitives. |
| `ccip_chainaccessor_finality_violated_total` | Proposed | `ErrFinalityViolated` is real but surfaced only as an unmetered error at this layer. |
| `ccip_chainaccessor_value_staleness_seconds{chain,source="feecomponents\|nativeprice\|chainfee\|feedprice\|quoter"}` | Proposed | Read age per chain since last successful value — the antidote to the configPoller/async-observer cache layers that freeze last-success silently. |
| `ccip_chainaccessor_no_bindings_total{source}` | Proposed (Low) | An intended-quiet case (address-discovery race) that is indistinguishable from "disabled" today; worth counting so discovery lag is visible. |

## Proposed metrics at a glance — config poller

| Metric | Status | Why |
|---|---|---|
| `ccip_reader_config_cache_age_seconds{chainID,chainFamily,chainName,chainSelector,kind="chain"\|"source"}` gauge | ✅ Implemented | The poller's whole job is freshness, and for that exact number it computed `time.Since(chainConfigRefresh)` and **discarded it into `Debugw`** (`config_poller_v2.go`). Now exported so every behind-the-cache read (curse, RMN config, config digest, router address) can be judged for staleness. |
| `ccip_reader_config_poll_success{chainID,chainFamily,chainName,chainSelector}` / `ccip_reader_config_poll_failure{...}` | ✅ Implemented | Per-chain refresh accounting, replacing the single chain-unlabeled `consecutiveFailedPolls` counter — now you can tell *which* chain is failing. |
| `ccip_reader_config_cache_overwritten_empty{chainID,chainFamily,chainName,chainSelector,kind}` | ✅ Implemented (+ write guard) | Flags when a refresh **time-stamped an empty snapshot as fresh**; the code now also refuses to write/mark-fresh an empty snapshot (the bug fix in this branch). The chain-identity labels were added alongside the `chainAttrs()` rollout — previously this metric had no chain label at all, even though the chain being refreshed (`chainSel`) was always in scope at the call site. |
| `ccip_reader_config_poller_last_success_timestamp{chainID,chainFamily,chainName,chainSelector}` gauge | ✅ Implemented | The background polling goroutine's liveness via last-success staleness — a wedged poller is now page-able instead of invisible. |
| `ccip_reader_config_poll_duration{chainID,...}` / `ccip_reader_config_miss_refresh_duration{chainID,...}` | Proposed | Per-poll and read-path miss-refresh latency; not implemented here. |
| `ccip_reader_config_cache_overwritten_empty{kind="source"}` | Proposed (source-cache variant) | The guard/metric cover the chain snapshot; the source-channel configs partial-overwrite is untouched. |

## How this extends the runbooks (all additive)

These signals sit **one causal level upstream** of the processor metrics the runbooks currently
read. Today the runbook is built entirely on *outcomes* — finished values the plugin already
computed (`onramp_max_seq_num`, `consensus_dropped`, `pending_messages`) — which can only tell you
*where* the pipeline is stuck, not *why the data feeding it went bad*. This section maps, step by
step and scenario by scenario, what the new metrics add.

### 1. A data-layer gate *before* the outcome tree (the highest-value addition)

step0's heartbeat only answers "is the process alive?" There is a failure mode it cannot see:
**process alive, rounds advancing normally, every outcome metric reading healthy — and the data it
is reading is empty or garbage.** When `onramp_max_seq_num` is flat, the runbook cannot currently
distinguish "no new messages (true idle)" from "every oracle's onramp read is returning empty
(false idle)." The reader's own counters can:

`read_empty`/`chain_gap`/`value_staleness` firing on the same query across oracles →
the onramp is stalled **because the read returns nothing**, and the outcome-level tree is moot.
This is effectively a **second heartbeat — data-source liveness** — orthogonal to process liveness,
and genuinely new: nothing in the current runbooks can answer "commit plugin healthy but fed
garbage" vs "commit plugin itself broken."

The runbook should therefore gain a *data-layer gate* step (run before the processor step tree, or
as the top of step2): if any reader `read_empty`/`chain_gap`/`value_staleness` series
for the lane (or dest chain) is non-zero, branch straight to `REPORT: chain-infra-oncall`, before
consensus or transmission logic is implicated.

### 2. Where it strengthens existing steps / scenarios

- **step2 — "onramp read is lagging."** Today this uses `ccip_commit_merkleroot_observation_errors`,
  which only fires for a handful of typed reasons. `chain_gap{NextSeqNum,state="misconfigured"}`,
  `read_empty` and `msg_dropped{MsgsBetweenSeqNums}` catch the *untyped* silent degradations
  (no-bindings, empty event index, cast/validation drops) that observation-errors never sees.
- **step2b — consensus dropped.** `read_empty` / non-"returned" `chain_gap` are the *per-oracle substrate* feeding
  `consensus_dropped{reason=insufficient_agreement|split}`. They let the runbook answer whether the
  disagreement was "because N oracles literally couldn't read the lane" vs "oracles read fine but
  observed different things" — before or in parallel with the consensus signal.
- **step3 / Scenario 3 (H1 vs H2).** Currently H1/H2 are inferred from silhouettes
  (`fchain_read_errors`, `config_digest_mismatch`). `read_empty` / non-"returned" `chain_gap` on the underlying
  source-chain queries give the **direct data-level observation** behind a consensus failure, for
  **all lanes at once**, rather than the single fChain the round needs. With the `node_id`/`csa_public_key`
  labels (see [known boundaries](#known-boundaries-and-deferred-work) below) this cleanly splits
  "DON-wide data-source failure" (all nodes empty) from "single bad node's RPC" (one series empty).
- **step4 — cursing.** `ccip_commit_rmn_curse_active` / `source_chain_cursed` are served from the
  **config poller cache** (`GetRmnCurseInfo`). If the poller froze or served an empty snapshot (the
  `config_cache_overwritten_empty` bug), step4 can **false-page on a curse already lifted** or miss a
  new one. `config_cache_age_seconds` / `config_poller_last_success_timestamp` let step4 say *"this
  curse reading is N minutes stale — re-read before acting"* (and the exculpatory mirror: don't trust
  a "halt" read from a stale cache).
- **step5 / Scenario 2 — `config_digest_mismatch`.** `GetOffRampConfigDigest` is also poller-cached;
  a stale or empty digest cache *manufactures* `config_digest_mismatch` pages across the DON. Poller
  freshness turns that branch from "mismatch = config sync issue" into "mismatch **and** the digest
  read is stale → reader issue, not config issue."
- **Scenario 2 — report-content rejections.** `read_empty`/`msg_dropped` on `MsgsBetweenSeqNums`
  explain *why* a built report is missing messages → spurious
  `empty_root`/`invalid_seqnum_range`/`stale` rejections. Splits "report-builder bug" (processor)
  from "reader-supplied partial data" (data layer).

### 3. Net-new branches that do not exist today

- **Chainfee / tokenprice health.** These processors have almost no runbook presence, and their
  silent-kill mode is exactly what this assessment surfaces: `GetChainFeePriceUpdate` /
  `GetChainsFeeComponents` / `GetWrappedNativeTokenPriceUSD` returning empty/partial → the async
  cache freezes → plugin looks healthy-idle. `batch_result{outcome="empty"}` +
  `value_staleness_seconds{chain,source="chainfee|tokenprice"}` create a first-class
  "fee/price pipeline healthy?" branch that does not exist today — including the
  `config_cache_overwritten_empty` class where a value is time-stamped fresh around empty data.
- **Temporal root-cause arrow.** Because data-layer failures are upstream of their consequences, a
  "did `read_empty`/non-"returned" `chain_gap` appear *before* `pending_messages`/`consensus_dropped` climbed?"
  ordering comparison gives a **causation direction**: "the read broke first, the pipeline
  consequence followed" vs "reads are fine, the breakdown is inside the pipeline." Purely from
  timestamp ordering — no extra instrumentation.
- **Report-content integrity** (as in §2/Scenario 2 above) doubles as a standalone health check on
  `MsgsBetweenSeqNums`: reads returning a subset are a warning even before any report is rejected.

### 4. The proactive payoff (`commit-plugin-health.md`)

The triage runbook is reactive — it needs a *missed* message to fire. Monitor-by-freshness metrics
(`chain_gap`, `value_staleness`, `read_empty`, `config_cache_age`) let the health runbook find a
silently-degraded lane **before a message is missed** — the incident in the making. Outcome-level
metrics cannot do this, because they legitimately don't move while nothing has been dropped yet.

---

## Known boundaries and deferred work

- **Complementary, not replacing.** The reader/accessor metrics say "the data entering the pipeline
  is bad"; the processor/consensus/transmission metrics say "how the pipeline reacted." Both are
  needed — the reader metrics alone still cannot distinguish a genuine message drought from a bad
  read. This is documented as goal-complementary: more useful signal is better.
- **`node_id` / `csa_public_key` labels resolve the per-oracle-vs-DON-wide question.** Every proposed
  beholder metric should carry a `node_id` and a `csa_public_key` label (`csa_public_key` is unique per node, the
  same value the heartbeat metrics already use). This turns §3's H1-vs-H2, and the data-layer-gate's
  "across oracles" check, into a real query instead of a hope: poll `group by (csa_public_key)` — a single
  bad node's empty read is one series, a DON-wide data failure is every series.
- **Event non-indexing is explicitly out of scope here.** These reader/accessor metrics make an
  empty result *visible and countable*, but they do not make it self-verifying: `read_empty` says
  "returned no data," not "should have returned data." Detecting a genuinely *un-indexed* event
  stream (the hardest false-idle) is the job of the components *below* this abstraction — e.g. an
  EVM log poller that pulls over HTTP vs websockets. Those lower-level components should export
  their own health metrics and are to be reviewed in a later session; this docs does not solve for
  them.
- Keep `read_empty`/`chain_gap` sovereign from the *correctness* of the batch — a
  partial read is not yet "wrong data," and a full read is not yet "correct data." Combined with the
  metric-reference note that "no such metric" must be treated as *unknown*, not `0`, the reference
  table below should be folded into the runbooks' reference tables as the metrics are implemented.

---

## Full findings — CCIP reader (`pkg/reader`)

Source: `pkg/reader/ccip.go` (`ccipChainReader`), interface `pkg/reader/ccip_interface.go`,
wrapper `observed_ccip.go`. Every read funnels through `getChainAccessor` into the accessor; the
reader itself holds **no error counter or aggregate metric of its own** — observability is wholly
a function of which log level each method uses, at what frequency.

### High

- **`NextSeqNum` is commit-merkleroot critical-path and hides dropped chains behind a 30-minute
  rate limit.** `ccip.go:346-412`: `fetchFreshSourceChainConfigsViaAccessor` (direct, cached-bypass
  read) returns a **subset** — chains that are not-found / disabled / misconfigured are silently
  dropped into buckets and only surfaced via
  `LogWhenExceedFrequency` (`readerLogFrequency = 30min`, `ccip.go:330,395-409`). Disabled/config-not-found
  are `Debugw`; only the misconfigured bucket gets `Errorw` with the sentinel
  `MsgMisconfiguredSourceChainsSkipped` (`ccip.go:334-341`, explicitly tied in the comment to "a
  plugin looking falsely idle"). A single dropped lane can be invisible for 30 minutes during an
  incident window, and the returned subset is indistinguishable from "all chains fine" to an
  outcome-only consumer. → `ccip_reader_chain_gap_total{query="NextSeqNum",chain,state=...}`.

- **`GetChainsFeeComponents` / `GetWrappedNativeTokenPriceUSD` return subsets and swallow causes.**
  `ccip.go:440-485` (chainfee critical path): per-chain failures are dropped from the result map;
  accessor-missing is rate-limited `Debugw` (451-453), timeout `Warnw` (462), other `Errorw` (464);
  non-positive fee validation `Errorw` (469-476). An empty return is "all chains failed"
  indistinguishable from "zero chains requested" (`ccip.go:445`). `ccip.go:495-573`
  (chainfee critical path): per-chain native-price reads in goroutines; two rate-limited `Debugw`
  miss/error necks (`519-521`, `530-532`), zero/empty native address `Debug` (539-541),
  `Timestamp == 0` stale → `Warnw MsgNoNativeTokenPriceAvailable` (555-557). The chainfee async
  observer then freezes whatever survived with `Debugw "returning cached value"` and never marks it
  stale (`observation_async.go:274-302`) — the classic "healthy but idle" chainfee state, rooted
  here. → `ccip_reader_chain_gap_total{query=...}` + value-staleness gauges.

- **Chain fee price updates can be wholly empty with `nil` error.** `GetChainFeePriceUpdate`
  `ccip.go:580-604`: accessor-miss → `Errorw`, `return nil` (583-586); query error → `Errorw`,
  `return nil` (593-596). Signature returns `map` with no `error`; the accessor failure mode farther
  down turns a full dest-chain failure into an empty map + `nil` (see accessor #34). The caller
  (`getChainFeePriceUpdates`, `observation_async.go:99-106`) treats `nil` and empty identically,
  so a complete dest-chain failure reads as a legitimately idle chain. → accessor
  `batch_result{outcome="empty"}`.

- **`printReports` destroys evidence for combined reports.** `ccip.go:187-206`: reports carrying a
  root have their price updates stripped before the `Debugw` (196-200) — deliberately, to save log
  space, but it means commit logs can never reconstruct a combined report's token/gas price data.
  → Low/Med: log-only concern worth flagging; the reader should rely on a value gauge, not the log.

### Med

- **`DiscoverContracts` silently drops source chains.** `ccip.go:754-833`: per-source-chain missing
  reader `Errorw` + `continue` (788-791), config fetch fail `Errorw` + `continue` (801-805); the
  result is best-effort with no discovered-vs-requested accounting. → `ccip_reader_chain_gap_total{query="DiscoverContracts",...}`.
- **`Sync` skips `ErrChainAccessorNotFound` with no log, and one bind failure fails the whole Sync
  error path.** `ccip.go:880-883` (no log), `886-897` (errgroup `Wait()` returns the first error
  even though discovery swallows it). → reader-global counters.
- **Config-poller-cached reads have no per-read freshness.** `GetRMNRemoteConfig` (618-648),
  `GetRmnCurseInfo` (652-661, comment: "Curse requires a dedicated cache, but for now fetching it
  in background..."), `GetOffRampConfigDigest` (1064-1078),
  `GetOffRampSourceChainsConfig` (937-977) all return background-polled data
  (`defaultRefreshPeriod = 30s`, `ccip.go:33`). Staleness governed solely by the poller's warn-only
  consecutive-failure counter. → `ccip_reader_query_last_success_timestamp`.
- **TxHash blanked when `populateTxHashEnabled=false`.** `ccip.go:278-282` — an analysis-relevant
  field wiped with no count. → `ccip_reader_txhash_blanked_total` (Low).

### Low

- **`GetExpectedNextSequenceNumber` is deprecated/dead** (`ccip_interface.go:179`, unused in
  commit/execute). Skip instrumentation.

- **`Close` / `GetLatestPriceSeqNr` error surfacing is correct** (bad for metrics-only purposes —
  caller-loggable).

## Full findings — default chain accessor & price reader (`pkg/chainaccessor`)

Sources: `default_accessor.go`, `config_processors.go`, `default_price_reader.go`. **None of the
three contain any metric — zero `prometheus`/`otel`/`beholder`/`meter` usage (confirmed; the only
"counter" hits are a local `int` variable).** The adjacent instrumentation is one layer below
(`contractreader/observed.go`, which is blind to swallowed paths) and one layer above
(`observed_ccip.go`, which measures only duration/size).

There is also **no quorum / BFT aggregation anywhere in this layer**: `DefaultAccessor` is a thin
single-reader facade (one per node), and "N-of-M" agreement or block-number alignment simply does
not exist here. Each oracle independently reads its own RPC; the only finality touchpoint is the
downstream health-check error `ErrFinalityViolated` (`contractreader/extended.go:126-129,169-172,205-207`).

### High — the silent-empty "false idle" producers (all log-only, no metric)

- **`GetChainFeePriceUpdate` batch failure → empty map, `nil`.** `default_accessor.go:834-837`
  (`Errorw`, `return make(...), nil`) and missing FeeQuoter → empty `nil` (850-853). This is the
  chainfee processor's data source; a dest-chain RPC/batch failure reads as "no price changes
  anywhere." **The single most valuable item in this whole assessment.**
- **Event reads return empty/`nil` on quiet.** `LatestMessageTo` → `0, nil` (444-446),
  `ExecutedMessages` → `nil, nil` when no seqs, empty map `nil` on empty (583-586, 620),
  `MsgsBetweenSeqNums` → `[], nil` (356-360), `CommitReportsGTETimestamp` → empty `[]`, `nil`
  (561, 1005-1009). A slow or empty event index is indistinguishable from a genuinely quiet chain,
  and the differencing log lines are Debug/Info.
- **`GetAllConfigsLegacy` → empty `ChainConfigSnapshot`, `nil` error on component failure.**
  `default_accessor.go:122-132` (Debug for no-bindings / Errorw otherwise, then
  `standardConfigs = ChainConfigSnapshot{}` and `nil`, `141`). This degrades the cached Router /
  curse / RMN config input silently.
- **Per-row drops across every read, aggregated nowhere.** Nonces per-addr
  (`731,739,745`), fee-update per-chain (`878-912`), commit-report validation/parse
  (`967-971,982-986,988-992`), feed-price per-contract (`default_price_reader.go:64-99`),
  fee-quoter token perdropped (`128-131` empty-`nil`, `157-171`, `237-264`),
  source-chain config rows (`config_processors.go:87-105`). Each is a `continue` with a log; the
  aggregate "rows dropped" does not exist.

### Med

- **Finality violation is real but unmetered at this layer.** `ErrFinalityViolated` bubbles as a
  top-level error (loud) but is only folded into the block-level `CrErrors` when it happens to be
  the outer `err`; a per-chain swallowed one is invisible. → `ccip_chainaccessor_finality_violated_total`.
- **no-bindings (address-discovery race) is intended-quiet but indistinguishable from disabled.**
  `config_processors.go:317-321,87-90`; `default_accessor.go:117-119` Debug. → low-count gauge.
- **Config poller staleness is warn-only.** `config_poller_v2.go:403-408` consecutive-failure
  counter with `Warnw` past `MaxFailedPolls`; no freshness/reorg-invalidation signal.
- **33 error paths return an error with no log and no metric** (loud-but-unmetered; e.g.
  `default_accessor.go:74-80,276-309,352-354,370-372,432-457,478-481,509-512,554-556,599-608,721-723,931-933`,
  `config_processors.go:122-278`, `default_price_reader.go:57-59,65,123-131,150-152,237-264`).
  These don't *cause* false-idle, but they are the errors that would aid diagnosis and are invisible
  today.

### Priority for instrumentation (in order)

1. `GetChainFeePriceUpdate` batch → empty/`nil` (`default_accessor.go:834-837,850-853`).
2. Silent-empty primitive reads: `LatestMessageTo`, `ExecutedMessages`, `MsgsBetweenSeqNums`,
   `CommitReportsGTETimestamp` (`default_accessor.go:444-446,356-360,583-586,561,1005-1009`).
3. Per-row/`continue` drops (`default_accessor.go:731-749,878-912,967-992`;
   `config_processors.go:87-105`; `default_price_reader.go:64-99,157-171`).
4. `GetAllConfigsLegacy` empty-snapshot-on-failure (`default_accessor.go:122-141`).
5. Reader-layer chain-gap + last-success freshness for the two cache layers.

---

## Full findings — config poller (`pkg/reader/config_poller_v2.go`)

Source background service (`configPollerV2`, created in `ccip.go:111-116` with
`defaultRefreshPeriod = 30s`, started at `ccip.go:127-133`). It is the cache behind
`GetRMNRemoteConfig`, `GetRmnCurseInfo`, `GetOffRampConfigDigest`, `GetOffRampSourceChainConfigs`,
the router address read in `GetWrappedNativeTokenPriceUSD`, and `DiscoverContracts` — i.e. almost
everything the commit plugin reads that isn't a seq-num/price. Its health *is* the plugin's view of
on-chain config, and none of it is observable.

### High

- **A refresh can time-stamp an *empty* snapshot as fresh and serve it to every consumer.**
  `batchRefreshChainAndSourceConfigs` (421-485) unconditionally writes
  `cache.chainConfigData = chainConfigSnapshot; cache.chainConfigRefresh = time.Now()` (457-460).
  The snapshot comes from `accessor.GetAllConfigsLegacy`, which — per the accessor findings — can
  return an **empty `ChainConfigSnapshot` with `nil` error** on component failure
  (`default_accessor.go:122-141`). End result: the cache records `chainConfigRefresh = now` around
  empty/garbage data. Every downstream read (`GetChainConfig` at 213-248 short-circuits on a non-zero
  refresh time, 228-233) then serves the empty snapshot **as fresh**, indefinitely — this silently
  reads like "config is fine, nothing written," the deepest false-idle in this layer. Nothing marks
  it; the only trace is the poller's `Debugw "Batch refreshed configs"` (478-483) that happens to log
  the (empty) snapshot. → `ccip_reader_config_cache_overwritten_empty_total{kind="chain"}`,
  plus guarding the write: only overwrite on a non-empty snapshot.

- **Cache freshness is internal-only.** Every `chainCache` tracks `chainConfigRefresh` /
  `sourceChainRefresh` (`config_poller_v2.go:79-85`), and `GetChainConfig` computes the one number
  that matters — `time.Since(chainConfigRefresh)` as `cacheAge` — then throws it into `Debugw`
  (232). Nothing exports it. A config that stopped refreshing looks identical to one refreshing every
  30s. → `ccip_reader_config_cache_age_seconds{chain,kind}`.

- **No per-chain poll accounting.** `refreshAllKnownChains` (388-409) refreshes every cached+known
  chain and folds all outcomes into a single `refreshFailed` bool and the
  `consecutiveFailedPolls` atomic (`Warnw` per-chain at 398, aggregate at 402-408). You can see
  *"10 consecutive failures somewhere"* but not *which* chain, or that only one lane's config broke
  while the rest are fine (exactly the lane-scoped failure the runbooks care about). →
  `ccip_reader_config_poll_success_total{chain}` / `ccip_reader_config_poll_failure_total{chain}`.

### Med

- **A wedged background goroutine is invisible.** `startBackgroundPolling` (172-185) is a bare
  `ticker.C` loop with no liveness/`lastSuccess` signal. If the goroutine dies or the ticker wedges,
  every cache ages out silently with no pageability. `HealthReport` (145-153) only appends to the
  service error buffer after `consecutiveFailedPolls >= MaxFailedPolls` *and* only counts failures
  that reached `refreshAllKnownChains`. → `ccip_reader_config_poller_last_success_timestamp{chain}`.
- **`sourceChainRefresh` is one timestamp for all source chains** (`85`), so a partial source-config
  refresh (some chains updated, others missed) is indistinguishable from a full one. Per-source-chain
  freshness is lost; `GetOffRampSourceChainConfigs` (304-346) returns only cached configs it happens
  to have. → per-`{chain}` staleness would need per-chain timestamps.
- **Read-path miss refreshes are synchronous and unmetered.** `GetChainConfig` (240) and
  `GetOfframpSourceChainConfigs` (333) perform the batch refresh inline on a cache miss (with a
  timeout), delaying the read; latency/outcome of these synchronous refreshes is not counted. →
  `ccip_reader_config_miss_refresh_duration{chain}`.

### Low

- **`consecutiveFailedPolls` is never growing a metric.** It has its intended `Warnw` + service-health
  purpose (403-408, `HealthReport`); graduating the counter into a gauge is a cosmetic, low-cardinality
  win (`ccip_reader_config_poll_consecutive_failures`).

---

## Metric reference (to be added to runbook reference tables as metrics land)

**Label convention:** every implemented beholder metric below carries `chainID`, `chainFamily`,
`chainName`, and `chainSelector` — resolved from a `ccipocr3.ChainSelector` by
`pkg/reader/rcmetrics`'s private `chainAttrs()` helper (mirrors `commit/metrics/prom.go`'s
`destChainAttrs()`/`sourceChainAttrs()`, minus the dest/source prefix, since each of these metrics
only ever describes one chain). This replaced the old bare numeric-selector `chain` label — every
`Record*` method on the `rcmetrics` interfaces now takes the selector directly
(`chainSelector ccipocr3.ChainSelector`) rather than a pre-formatted string, so there's no more
`rcmetrics.ChainLabel(...)` call at any site. **`ccip_reader_read_outcome` is the one exception**:
its chain-identity labels describe the reader instance's destination chain, not "the chain this
specific read is about" like every other row here — see its row below.

Proposed (not-yet-implemented) metrics below still carry `node_id` and `csa_public_key` labels
(unique per node) in addition to whatever chain-identity/per-metric labels are listed — this is
what makes the per-oracle-vs-DON-wide distinction queryable (see
[known boundaries](#known-boundaries-and-deferred-work)).

| metric | otel_type | key labels | registration |
|---|---|---|---|
| `ccip_reader_read_outcome` | Counter | chain(dest) + `query,outcome` | beholder |
| `ccip_reader_read_empty` | Counter | chain + `query` | beholder |
| `ccip_reader_chain_gap` | Counter | chain + `query,state` | beholder |
| `ccip_reader_query_last_success_timestamp` | Gauge | `query` | proposed, beholder |
| `ccip_reader_msg_dropped` | Counter | `query,reason` (no chain — instrument registered, no call site yet) | beholder |
| `ccip_reader_chain_fee_components` | Gauge | chain + `feeType` (beholder); promauto side unchanged, bare `chainFamily,chainID` | dual |
| `ccip_chainaccessor_batch_result` | Counter | `source,operation,outcome` | proposed, beholder |
| `ccip_chainaccessor_row_drops` | Counter | `source,operation,reason` | proposed, beholder |
| `ccip_chainaccessor_empty_read` | Counter | `source,operation` | proposed, beholder |
| `ccip_chainaccessor_finality_violated` | Counter | `source` | proposed, beholder |
| `ccip_chainaccessor_value_staleness_seconds` | Gauge | `chain,source` | proposed, beholder |
| `ccip_chainaccessor_no_bindings` | Counter | `source` | proposed, beholder (Low) |
| `ccip_reader_config_cache_age_seconds` | Gauge | chain + `kind` | beholder |
| `ccip_reader_config_poll_success_total` / `ccip_reader_config_poll_failure_total` | Counter | chain | beholder |
| `ccip_reader_config_cache_overwritten_empty` | Counter | chain + `kind` | beholder |
| `ccip_reader_config_poller_last_success_timestamp` | Gauge | chain | beholder |
| `ccip_reader_config_miss_refresh_duration` | Histogram | chain | proposed, beholder |

("chain" above = `chainID,chainFamily,chainName,chainSelector` from `chainAttrs()`.)
