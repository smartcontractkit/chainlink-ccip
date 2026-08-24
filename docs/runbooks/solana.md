---
name: solana-chain-family
description: Chain-family deep dive for a Solana chain feeding or receiving a CCIP commit plugin. Drills into the Solana chain-family library (chainlink-solana) layers -- log poller, block-history fee estimator, and the Solana transaction manager (Txm) -- that underpin the commit plugin's data-source (reader) and report-transmission paths. Enter this doc when the commit runbooks hand off with REPORT:...-txm-oncall / REPORT:...-infra-oncall, or when a data_source / report_transmission check in commit-plugin-health.md fires for a Solana chain.
trigger: "entered from docs/runbooks/commit-plugin-health.md (data_source or report_transmission group CRIT/WARN) or docs/runbooks/uncommitted-message.md (step2c, or Scenario 2 'report transmission stuck') when the destination chain is Solana. Not triggered standalone."
severity: page
owner: chain-solana-oncall
inputs:
  chainID: {type: string, description: "the Solana chain ID (chain family / cluster identifier) of the chain under investigation. Like evm.md, the chainlink-solana metrics carry a bare chainID label -- same value you used for the commit plugin's destChainID where applicable, different label-key spelling"}
  entry: {type: string, description: "which commit-plugin symptom brought you here", default: "data_source"}
  txSignature: {type: string, required: false, description: "optional. A Solana transaction signature for the specific stuck/reverted commit-report tx, to cross-check on the explorer in stage B"}
related:
  - docs/runbooks/commit-plugin-health.md # the health checklist this doc deep-dives for Solana dest chains
  - docs/runbooks/uncommitted-message.md   # the triage doc this doc deep-dives for Solana dest chains
  - docs/runbooks/evm.md                   # the same deep-dive for the EVM chain family
  - docs/runbooks/chain-identifiers.md     # translating chain ID / chain selector / chain name
status: living
---

- [Runbook: Solana chain family (chainlink-solana)](#runbook-solana-chain-family-chainlinksolana)
   * [For agents](#for-agents)
      + [Label key cheat sheet](#label-key-cheat-sheet)
      + [What Solana's TXM calls a failure -- and why the taxonomy matters](#what-solanas-txm-calls-a-failure--and-why-the-taxonomy-matters)
      + [No on-chain per-address nonce, so no nonce-gap equivalent](#no-on-chain-per-address-nonce-so-no-nonce-gap-equivalent)
      + [Empty result sets](#empty-result-sets)
      + [Counter _total suffixing and the promauto-only read-path gap](#counter-total-suffixing-and-the-promauto-only-read-path-gap)
   * [Decision graph](#decision-graph)
   * [Steps](#steps)
      + [Stage A0 -- log poller: keeping up with slots?](#stage-a0--log-poller-keeping-up-with-slots)
      + [Stage B0 -- fee layer (block-history estimator): pricing sends?](#stage-b0--fee-layer-block-history-estimator-pricing-sends)
      + [Stage B1 -- Txm broadcast: are reports landing at all?](#stage-b1--txm-broadcast-are-reports-landing-at-all)
      + [Stage B2 -- Txm failure taxonomy: which kind of non-landing?](#stage-b2--txm-failure-taxonomy-which-kind-of-non-landing)
   * [Metric reference](#metric-reference)
   * [Output contract](#output-contract)

# Runbook: Solana chain family (chainlink-solana)

The CCIP commit plugin reaches Solana through chainlink-solana (checked out at
`../chainlink-solana`, one level above this repo's root). Where EVM's
transaction manager tracks discrete wallet nonces and gas price bumps, Solana's Txm operates in
*priority-fee* (compute-unit-price) and *simulation* terms: a transaction either simulates cleanly
and gets included, or it fails one of several distinct ways. Because of that, the metric story on
Solana is a **taxonomy of why transactions don't land** plus a slot-based log poller -- and there
is deliberately **no EVM-style nonce-gap concept** (see the note below).

Two layers, matching the two commit-plugin paths exactly as `evm.md` does for EVM:

- **Read path (data source).** The commit reader ingests CCIP events through the Solana **log
  poller**, which tracks processed **slots** (Solana's equivalent of EVM block numbers). A slot
  wedge or skip storm produces clean `read_empty` / `chain_gap` results upstream -- the classic
  false-idle shape. Stage **A0**.
- **Write path (report transmission).** Commit reports are submitted through the Solana **Txm**,
  priced by the **block-history estimator** (compute-unit-price). That's the whole `B` half:
  **B0** fee pricing, **B1** broadcast flow, **B2** the failure taxonomy.

Read A0 for `data_source`, B0-B2 for `report_transmission`. The halves are independent.

## For agents

Reactive triage doc: walk the [decision graph](#decision-graph), substitute `$chainID` (and
`$entry`), evaluate conditions, follow the output. Outcome vocabulary is the standard
`STOP` / `CONTINUE:<step>` / `REPORT:<owner>`; `REPORT` is a handover, never a remediation.

### Label key cheat sheet

Every chainlink-solana metric in this doc carries a bare `chainID` label (from
`metrics.Labeler().With("chainID", chainID)` in each package, exposed to beholder via
`GetOtelAttributes()`). It is **not** the commit plugin's `destChainID`/`sourceChainName` family
and there is no selector here. Same-number/different-key rule as `evm.md`: use the numeric chain
ID you already had.

| label key | appears on | meaning |
|---|---|---|
| `chainID` | all metrics here | the Solana chain ID / cluster. Every beholder series also inherits `node_id` + `csa_public_key` from the beholder client for per-node splitting |
| `# (outcome suffix)` | `solana_log_poller_txs_truncated_*`, `solana_log_poller_txs_log_parsing_error_*` | `_succeeded` vs `_reverted` -- whether the truncated/parse-failed tx ultimately succeeded or reverted on-chain |

### What Solana's TXM calls a failure -- and why the taxonomy matters

EVM gives you a boolean stuck signal (`txm_reached_max_attempts`) and one lifecycle-failure
counter. Solana instead gives you six purpose-built error counters that *subclass* the problem
before you even look at a log header. Get these right or every drought looks like "the chain is
down":

- `solana_txm_tx_error_reject` — the **RPC rejected the transaction outright** (before broadcast).
  This is almost always RPC/node-side, or a transaction that violates a validation rule (e.g. too
  cheap a fee is a common real cause). Check RPC health / fee first, not the report.
- `solana_txm_tx_error_drop` — timed out during confirmation; **most likely dropped** from the
  chain (may still be included). Under-funding/compute-unit-price that's too low to get pulled in
  is the archetypal cause → pair with `solana_bhe_compute_unit_price` and `solana_txm_fee_bumps`.
- `solana_txm_tx_error_sim_revert` — **reverted during simulation** (never attempted). On a commit
  report this usually points at a *state/account* condition the report hit at send-time, not the
  transaction manager. Verify on-chain state before blaming either.
- `solana_txm_tx_error_sim_other` — simulation failed with an **unrecognized error**. Same
  treatment as sim_revert but explicitly "we don't know why." Say that you can't name it rather
  than guessing; check logs.
- `solana_txm_tx_error_revert` — **included but failed on-chain** (sometimes post-simulate
  state change). The report landed in a block and reverted.
- `solana_txm_tx_error_dependency` — failed because a **dependency transaction failed**.
- `solana_txm_tx_error` — the umbrella sink for transactions that error-tested but didn't fall
  into a named bucket.

### No on-chain per-address nonce, so no nonce-gap equivalent

`evm.md` leans on `txm_num_nonce_gaps` / `txm_rpc_nonce` because EVM serializes an account's
transactions by nonce. Solana does not: a transaction identifies its instruction accounts and
*blockhash/priority-fee* season, and there is no ordering nonce to park. Consequently there is **no
`txm_num_nonce_gaps` or `txm_reached_max_attempts` here** -- do not go looking for them, and do not
translate an EVM nonce diagnosis into a Solana one. The Solana "one stuck thing blocking the rest"
equivalent is a **recurring `revert`/`sim_revert`/`dependency` on the same program/account**, which
repeats per report until the underlying state or the fee budget changes. That is exactly what
`report_transmission_gave_up` sees as a sustained drought.

### Empty result sets

Same discipline as the other runbooks. The gauges that should be continuously present while the
process runs are `solana_log_poller_last_processed_slot` (A0) and `solana_bhe_compute_unit_price`
(B0): an **empty** result on those means the metric isn't reaching your datasource, not "zero" --
report `UNKNOWN`. Every counter here (`solana_log_poller_blocks_skipped`,
`solana_log_poller_txs_*`, all `solana_txm_*` error/success counters, `solana_txm_fee_bumps`) is
an event counter: **empty == it never happened == 0**, grade normally. As ever, if a continuously-
emitted gauge is itself empty in the same run, don't trust counter-emptiness to mean "confirmed
clean."

### Counter _total suffixing and the promauto-only read-path gap

Same pipeline rule as `evm.md`: OTel counters gain `_total` when converted to Prometheus unless the
name already ends in it; the doc lists names as registered -- check your datasource's metric
browser. (Among these, `solana_log_poller_blocks_skipped`, `solana_txm_tx_*`,
`solana_txm_fee_bumps` will typically get a `_total` appended by an OTel→Prometheus exporter;
`solana_log_poller_txs_truncated_*` etc. likewise.)

**About the read-path RPC signal:** unlike EVM, chainlink-solana has **no beholder record for its
RPC client calls** -- the per-node/client metrics (`solana_client_latency_ms`, deprecated; and the
chainlink-framework `rpc_call_latency` promauto gauge) are **promauto-only**. They are still
available on a direct-scrape datasource, so this doc can use them, but know that they will not
appear on a strictly beholder-fed pipeline, and there is no `solana_pool_rpc_node_calls_*`
beholder equivalent to EVM's. On a beholder-only datasource, leave the read-path RPC question to
the commit-plugin reader metrics (`ccip_reader_read_outcome_error`) rather than this doc.

## Decision graph

```yaml
steps:
  # ---- READ PATH (data source) ----
  - id: A0
    check: log_poller_slot_catchup
    query: 'solana_log_poller_last_processed_slot{chainID="$chainID"}'
    condition: "value stale (no fresh sample in the last couple of poll intervals) OR clearly lagging the current slot (cross-check on the explorer) OR blocks_skipped climbing"
    if_true:
      stale_or_lag: {action: "REPORT:chain-solana-oncall", reason: "the Solana log poller for $chainID is stuck or behind -- every commit/reader log read routes through it, so a wedge here surfaces as clean reader read_empty / chain_gap (false idle) in the commit docs, or batch_fetch_failed in the config poller. Report as a Solana chain-family / infra finding, not a reader one. Corroborate blocks_skipped: solana_log_poller_blocks_skipped{chainID=\"$chainID\"}."}
    if_false: {action: "CONTINUE:$entry == \"report_transmission\" ? B0 : STOP", reason: "log poller keeping up. Stop if you came in via data_source."}
    automatable: true

  # ---- WRITE PATH (report transmission) ----
  - id: B0
    check: fee_estimator_pricing
    query: 'solana_bhe_compute_unit_price{chainID="$chainID"}'
    condition: "value absent/stale (UNKNOWN) OR materially below historical floor for $chainID (compare against recent samples)"
    if_true: {action: "CONTINUE:B0_detail", reason: "the block-history estimator is not producing a usable compute-unit-price -- underpriced txs get dropped/rejected and never land"}
    if_false: {action: "CONTINUE:B1", reason: "BHE pricing normally; move to broadcast flow"}
    automatable: true

  - id: B0_detail
    check: underpricing_signature
    queries:
      - 'sum(rate(solana_txm_tx_error_drop{chainID="$chainID"}[15m]))'
      - 'sum(rate(solana_txm_tx_error_reject{chainID="$chainID"}[15m]))'
      - 'sum(rate(solana_txm_fee_bumps{chainID="$chainID"}[15m]))'
    condition: "drop OR reject climbing, combined with low/absent compute_unit_price"
    if_true: {action: "REPORT:chain-solana-oncall", reason: "the fee/drop/reject signature points at costs: compute_unit_price is low/absent AND drop/reject are climbing. Underpriced Solana transactions are dropped by the cluster or rejected by RPCs, so every commit report struggles to land (report_transmission_gave_up) regardless of the reader being healthy. The fix is the block-history estimator (wider block window, higher default floor), an infra/ops decision, not a commit-plugin one."}
    if_false: {action: "CONTINUE:B1", reason: "no concurrent drop/reject/fee-bump storm; treat the estimator note as a caveat, continue"}
    automatable: true

  - id: B1
    check: txm_broadcast_flow
    queries:
      - 'sum(rate(solana_txm_tx_success{chainID="$chainID"}[15m]))'
      - 'sum(rate(solana_txm_tx_finalized{chainID="$chainID"}[15m]))'
      - 'max(solana_txm_tx_pending{chainID="$chainID"})'
    condition: "success==0 AND finalized==0 over the window"
    if_true: {action: "CONTINUE:B2", reason: "no reports confirming at all -- the give-up is a 'nothing lands' situation, classify why in B2"}
    if_false:
      pending_climbing_or_flat_high: {action: "CONTINUE:B2", reason: "txs are inflight but success/finalized are zero -- more are queued than landing; classify why"}
      success_positive: {action: "STOP", reason: "Solana Txm is landing reports ($chainID). If report_transmission_gave_up still fires it's per-report (revert on a specific report) or a backlog/round-cadence issue, not the Solana node/Txm stack. If you have a $txSignature, sanity-check it on the explorer before stopping."}
    automatable: true

  - id: B2
    check: txm_failure_taxonomy
    queries:
      - 'sum(rate(solana_txm_tx_error_reject{chainID="$chainID"}[15m]))'
      - 'sum(rate(solana_txm_tx_error_revert{chainID="$chainID"}[15m]))'
      - 'sum(rate(solana_txm_tx_error_sim_revert{chainID="$chainID"}[15m]))'
      - 'sum(rate(solana_txm_tx_error_sim_other{chainID="$chainID"}[15m]))'
      - 'sum(rate(solana_txm_tx_error_drop{chainID="$chainID"}[15m]))'
      - 'sum(rate(solana_txm_tx_error_dependency{chainID="$chainID"}[15m]))'
    condition: "classify the dominant nonzero class (report the counts)"
    if_true:
      reject_dominant:       {action: "REPORT:chain-solana-oncall", reason: "RPCs are rejecting the commit-report txs outright. Cross back to B0 (underpriced fees are a common real trigger) and check RPC/cluster health. Not a report-content bug."}
      drop_dominant:         {action: "REPORT:chain-solana-oncall", reason: "commit txs are timing out during confirmation and being dropped. Pair with B0 fee pricing and solana_txm_fee_bumps. Likely cluster congestion or underpricing, not content."}
      sim_revert_dominant:   {action: "REPORT:ccip-commit-oncall", reason: "commit reports are reverting at simulation (never attempted). This is usually a state/account condition the report hit at send-time -- e.g. the report references accounts/state that has already moved. Confirm on-chain state for $chainID before blaming the Txm; this can genuinely be a commit-report state issue."}
      sim_other_dominant:    {action: "REPORT:ccip-commit-oncall", reason: "simulation failed with an unrecognized error. Do NOT translate this into a specific cause -- report the count and that the reason is unknown; pull node logs for the sim error text."}
      revert_dominant:       {action: "REPORT:ccip-commit-oncall", reason: "commit reports are included in a block but revert on-chain. Confirm against $txSignature on the explorer; often a state/priority-fee interaction at execution."}
      dependency_dominant:   {action: "REPORT:chain-solana-oncall", reason: "commit txs are failing because a dependency transaction failed. Naming the dependency requires the account list in the failing tx -- look to Txm logs / transaction detail."}
      multiple:              {action: "REPORT:chain-solana-oncall", reason: "more than one class is firing; report each with its rate, order by impact, and note which likely share one root (e.g. drop+reject share the fee cause)."}
    if_false: {action: "STOP", reason: "no tx in the window hit any named failure, yet nothing confirms and nothing is pending-ish -- the most honest reading is a confirmation/indexing lag on the RPC side (txs landed but aren't being observed as finalized). Verify on-chain via $txSignature / explorer before escalating further."}
    automatable: true
```

## Steps

### Stage A0 -- log poller: keeping up with slots?

The commit reader ingests CCIP sends/commits through the Solana log poller, which tracks
**processed slots**, not a chain head block. Two gauges:

```promql
solana_log_poller_last_processed_slot{chainID="$chainID"}
```

Continuous absence → the poller goroutine is wedged (parks at its last value). Clear lag behind
the current slot, or `solana_log_poller_blocks_skipped` climbing, → it's running but can't keep
up and is dropping ranges on retry exhaustion:

```promql
sum(rate(solana_log_poller_blocks_skipped{chainID="$chainID"}[15m]))
```

A wedged or skipping Solana log poller is the highest-probability *silent false idle* on a Solana
dest chain: every per-filter log read comes back empty and clean, so the commit docs' reader gates
fire `read_empty`/`chain_gap` and the plugin looks securely idle. Also worth a glance:
`solana_log_poller_txs_truncated_*` and `solana_log_poller_txs_log_parsing_error_*` -- these
count txs whose logs were truncated or failed to parse (split by whether the tx ultimately
`_succeeded` or `_reverted`), and a parsing-error storm can drop *individually* *invisible*
messages even while the poller as a whole advances. If either climbs and you're investigating a
specific message that "should have been seen," that is the first place to look.

### Stage B0 -- fee layer (block-history estimator): pricing sends?

Solana transactions are priced by a **compute-unit-price** the block-history estimator derives
from recent block fee data. If the estimator can't produce one (or produces one far below the
cluster's actual floor), the Txm's broadcasts get dropped or rejected no matter how healthy the
reader is:

```promql
solana_bhe_compute_unit_price{chainID="$chainID"}
```

Absent/stale = estimator not running (UNKNOWN, not zero). Abnormally low or gliding toward the
default floor = it has no good recent data. Corroborate the "underpriced" diagnosis with:

```promql
sum(rate(solana_txm_tx_error_drop{chainID="$chainID"}[15m]))
sum(rate(solana_txm_tx_error_reject{chainID="$chainID"}[15m]))
sum(rate(solana_txm_fee_bumps{chainID="$chainID"}[15m]))
```

`fee_bumps` is the Txm re-raising the price to chase inclusion; a persistent bump storm with drops
on top is a costs problem (estimator tuning / cluster congestion), a `chain-solana-oncall`
decision, and it sits *underneath* the commit plugin's `report_transmission_gave_up`.

### Stage B1 -- Txm broadcast: are reports landing at all?

```promql
sum(rate(solana_txm_tx_success{chainID="$chainID"}[15m]))
sum(rate(solana_txm_tx_finalized{chainID="$chainID"}[15m]))
max(solana_txm_tx_pending{chainID="$chainID"})
```

Solana separates **success** (included and executed) from **finalized** (committed past the
finality horizon). A healthy node fronts them by roughly a block; both should be `> 0` whenever
the plugin is transmitting. Both `== 0` with `pending` parked flat → nothing is landing → **B2**.
`success`/`finalized` positive → the Txm is landing reports; a sustained commit-plugin
`report_transmission_gave_up` on top of that is per-report or cadence, not the Solana stack.

### Stage B2 -- Txm failure taxonomy: which kind of non-landing?

Classify with the six error counters (queries in the decision graph). Read the taxonomy in
[For agents](#what-solanas-txm-calls-a-failure--and-why-the-taxonomy-matters) -- the short version:
`reject`/`drop` → RPC/costs (own as `chain-solana-oncall`, cross to B0); `sim_revert`/
`sim_other`/`revert` → report/state-level, verify on-chain and own as `ccip-commit-oncall` when a
real report-state interaction (be honest that sim_other is an unexplained reason); `dependency` →
an account/program dependency failed, name it from the tx detail. If nothing fires but nothing
lands either, conclude a confirmation/indexing lag and verify against the explorer rather than
inventing a cause.

## Metric reference

`otel_type` governs `_total` suffixing (see the pipeline note). All carry `chainID`; beholder
series also inherit `node_id` + `csa_public_key`. Registration is effectively **dual** (promauto +
beholder) for the log-poller and Txm families; the read-path RPC metric is **promauto-only** (see
[the gap note](#counter-total-suffixing-and-the-promauto-only-read-path-gap)).

| metric | otel_type | key labels | registration |
|---|---|---|---|
| `solana_log_poller_last_processed_slot` | Gauge | `chainID` | dual |
| `solana_log_poller_blocks_skipped` | Counter | `chainID` | dual |
| `solana_log_poller_txs_truncated_succeeded` / `_reverted` | Counter | `chainID` | dual |
| `solana_log_poller_txs_log_parsing_error_succeeded` / `_reverted` | Counter | `chainID` | dual |
| `solana_bhe_compute_unit_price` | Gauge | `chainID` | dual |
| `solana_txm_tx_success` | Counter | `chainID` | dual |
| `solana_txm_tx_finalized` | Counter | `chainID` | dual |
| `solana_txm_tx_pending` | Gauge | `chainID` | dual |
| `solana_txm_tx_error` | Counter | `chainID` | dual |
| `solana_txm_tx_error_reject` | Counter | `chainID` | dual |
| `solana_txm_tx_error_revert` | Counter | `chainID` | dual |
| `solana_txm_tx_error_drop` | Counter | `chainID` | dual |
| `solana_txm_tx_error_sim_revert` | Counter | `chainID` | dual |
| `solana_txm_tx_error_sim_other` | Counter | `chainID` | dual |
| `solana_txm_tx_error_dependency` | Counter | `chainID` | dual |
| `solana_txm_fee_bumps` | Counter | `chainID` | dual |
| `solana_client_latency_ms` | Gauge | `request,url` | promauto-only (deprecated) |
| `rpc_call_latency` (Solana) | Histogram | `chainFamily,chainID,rpcUrl,isSendOnly,success,callName` | promauto-only via chainlink-framework |

(`--succeeded`/`--reverted` here are suffixes on the two log-poller tx-outcome counters, produced
by chainlink-solana's `outcomeDependantMetric` helper.)

The `ccip_commit_*` / `ccip_reader_*` series the commit runbooks actually query are **not** in this
doc; the Solana layer is what sits underneath them and what this doc's `REPORT:...-oncall`
outcomes hand you to investigate.

## Output contract

```yaml
solana_finding:
  chainID: string
  entry: data_source | report_transmission
  timestamp: string           # ISO8601
  verdict: OK | ISSUE
  steps_run:
    - id: string              # A0 | B0 | B0_detail | B1 | B2
      outcome: STOP | CONTINUE:<step> | REPORT:<owner>
      value: string           # the gauges/counter rates that drove the outcome, verbatim
  findings:                   # one per REPORT; empty list if none
    - owner: string           # chain-solana-oncall | ccip-commit-oncall
      summary: string
      failure_class: string   # reject | drop | sim_revert | sim_other | revert | dependency | fee | log_poller
      evidence: string        # exact query + result that produced it
      commit_handoff: string  # which commit-runbook owner/check this replaces or justifies
```

Do not page, escalate, or remediate. Hand this report to the named owner; the commit runbook you
dropped in from already established *whether* there's a problem, this doc only locates *which
Solana-layer component* owns it.
