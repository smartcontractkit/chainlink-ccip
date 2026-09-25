---
name: evm-chain-family
description: Chain-family deep dive for an EVM chain feeding or receiving a CCIP commit plugin. Drills into the EVM chain-family library (chainlink-evm) layers -- RPC client pool, log poller, gas/fee estimators, and the transaction manager (TXM) -- that underpin the commit plugin's data-source (reader) and report-transmission paths. Enter this doc when the commit runbooks hand off with REPORT:chain-b-txm-oncall or REPORT:chain-infra-oncall, or when a data_source / report_transmission check in commit-plugin-health.md fires.
trigger: "entered from docs/runbooks/commit-plugin-health.md (data_source or report_transmission group CRIT/WARN) or docs/runbooks/uncommitted-message.md (step2c, or Scenario 2-B 'transmitted, reverted/stuck on-chain'). Not triggered standalone."
severity: page
owner: chain-evm-oncall
inputs:
  chainID: {type: string, description: "the EVM chain ID of the chain under investigation. Crash-important: the commit runbooks filter on the commit plugin's destChainID label; the chainlink-evm metrics in THIS doc carry a bare chainID label (the EVM chain ID), so the value you substitute here is the same number you used for destChainID there, but it must be applied to a different label key -- see the label cheat sheet before running any query"}
  entry: {type: string, description: "which commit-plugin symptom brought you here", default: "data_source"}
  txHash: {type: string, required: false, description: "optional. If you have the specific stuck/reverted commit-report tx hash (from logs or an explorer), you can cross-reference it in stage B; otherwise run the checks blind and let them tell you what to look at"}
related:
  - docs/runbooks/commit-plugin-health.md # the health checklist this doc deep-dives for EVM dest chains
  - docs/runbooks/uncommitted-message.md   # the triage doc this doc deep-dives for EVM dest chains
  - docs/runbooks/solana.md                # the same deep-dive for the Solana chain family
  - docs/runbooks/chain-identifiers.md     # translating chain ID / chain selector / chain name
status: living
---

- [Runbook: EVM chain family (chainlink-evm)](#runbook-evm-chain-family-chainlink-evm)
   * [For agents](#for-agents)
      + [Label key cheat sheet -- the crash-important part](#label-key-cheat-sheet--the-crash-important-part)
      + [The success label inversion on rpc_call_latency](#the-success-label-inversion-on-rpc_call_latency)
      + [Which TXM metrics actually carry a chainID label (not all do)](#which-txm-metrics-actually-carry-a-chainid-label-not-all-do)
      + [Empty result sets](#empty-result-sets)
      + [Counter _total suffixing](#counter-total-suffixing)
   * [Decision graph](#decision-graph)
   * [Steps](#steps)
      + [Stage A0 -- RPC client pool: are the nodes answering?](#stage-a0--rpc-client-pool-are-the-nodes-answering)
      + [Stage A1 -- log poller: is it keeping up with the chain?](#stage-a1--log-poller-is-it-keeping-up-with-the-chain)
      + [Stage B0 -- gas/fee layer: can it price and afford a send?](#stage-b0--gasfee-layer-can-it-price-and-afford-a-send)
      + [Stage B1 -- TXM broadcast: are reports being sent at all?](#stage-b1--txm-broadcast-are-reports-being-sent-at-all)
      + [Stage B2 -- TXM stuck: killed attempts, nonce gaps, lifecycle failures](#stage-b2--txm-stuck-killed-attempts-nonce-gaps-lifecycle-failures)
   * [Metric reference](#metric-reference)
   * [Output contract](#output-contract)

# Runbook: EVM chain family (chainlink-evm)

`commit-plugin-health.md` and `uncommitted-message.md` decide *whether* the commit plugin's
destination chain is fine *as a plugin*. That decision bottoms out in `REPORT:chain-b-txm-oncall`
(report transmission) and `REPORT:chain-infra-oncall` / the `data_source` group (reader feeding)
-- but those handoffs deliberately don't tell the on-call **what to look at in the EVM layer**.
This doc fills that gap: it's the runbook those `REPORT` outcomes and those two groups point to,
using the metrics chainlink-evm itself emits (checked out at `../chainlink-evm`, one level above
this repo's root) for its own critical components.

Two layers matter to CCIP commit on an EVM destination chain, and they map 1:1 onto how the
plugin got its data and how it gets its reports out:

- **Read path (data source).** CCIP's reader reads on/offramp state and messages over RPC and
  ingests on-chain events through the **log poller**. If either degrades, the commit plugin is
  fed empty/partial/stale data and looks "fine but idle" -- exactly the 
  [false idle](../runbooks/README.md#identifying-a-ccip-message) trap the reader-metrics gate in
  `uncommitted-message.md` step2c is designed to catch. Stages **A0/A1**.
- **Write path (report transmission).** Commit reports are submitted as normal transactions
  through the **transaction manager (TXM)**, which must price them (gas estimators), broadcast
  them, and shepherd them to inclusion. When `report_transmission_gave_up` fires in the commit
  plugin, the root cause usually lives in one of: can't afford a send (fee layer), can't broadcast
  (RPC), or the tx is stuck/backing up the send queue. Stages **B0/B1/B2**.

Read A0 and A1 if your `entry` is `data_source`; read B0-B2 if it's `report_transmission`. The
two halves are independent and you can enter directly at the relevant quarter.

## For agents

This is a reactive triage doc: walk the [decision graph](#decision-graph) mechanically, fill in
`$chainID` (and `$entry`) from the inputs, evaluate each query's condition, and follow the output.
Outcome vocabulary is the same three-way contract as `uncommitted-message.md`:
`STOP` / `CONTINUE:<step>` / `REPORT:<owner>`. `REPORT` is a handover, never an action.

### Label key cheat sheet -- the crash-important part

The commit runbooks filter on the **commit plugin's** labels: `destChainID`, `sourceChainName`,
etc. (from `destChainAttrs()`/`sourceChainAttrs()`). **The chainlink-evm metrics in this doc do
not carry those.** They carry a bare `chainID` (the EVM chain ID), and in the older promauto
downstream of the pool they carry `evmChainID`. There is no `destChainID` label anywhere in this
doc. So the number you substitute for `$chainID` is the **same numeric EVM chain ID** you used for
`destChainID` in the commit doc -- the key is just spelled differently. Do not try to swap in a
chain selector, a chain name, or `destChainID` as a label key; every one of those silently returns
empty here.

| label key | appears on | meaning |
|---|---|---|
| `chainID` | all EVM beholder metrics | the EVM chain ID (== commit's `destChainID` value for the dest chain). **Live gas/TXM metrics only exist on the destination/tx-sending chain** -- a query scoped to a chain the node doesn't send txs on returns empty, legitimately |
| `chainFamily` | all EVM beholder metrics | `"EVM"` on the pool counters, `"evm"` on the log-poller and gas-updater gauges -- **case is not consistent**; don't regex against a single spelling |
| `evmChainID` | the promauto pool counters (`evm_pool_rpc_node_calls_*`), gas-updater gauges (promauto side only) | same value, old promauto spelling. Only on the direct-scrape path, not the beholder path |
| `nodeName` | `evm_pool_rpc_node_calls_*` | which RPC node -- use this to split a single bad RPC from the whole pool failing |
| `rpcDomain` | `evm_pool_rpc_node_calls_*` (beholder), `rpc_call_latency` | the RPC endpoint/domain (sanitized) the call went to |
| `callName` / `rpcCallName` | RPC metrics | which JSON-RPC method / logical call |
| `percentile` | the gas-updater gauges | which price percentile, formatted as a **string like `"60%"`** (not `p25/p50`) |
| `mode` | `block_history_estimator_connectivity_failure_count` | `legacy` or `eip1559` |

The pool metrics are effectively **dual-registered** (same name on promauto and beholder), which
means the label set you see depends on which path your datasource consumes -- see
[the note below](#counter-total-suffixing). On a beholder-fed Prometheus datasource expect
`chainID,nodeName,rpcDomain,callName`; on a direct promauto scrape expect `evmChainID,nodeName`.
Confirm against the metric browser before trusting an empty result.

One metric in this doc is **not** part of the EVM pool's beholder trio and has to be reached via
the chainlink-framework latency instrument -- **`rpc_call_latency`** (a histogram). It is the
thinnest, most information-dense read-path signal (it exposes per-`callName` latency and success
for *every* RPC call), so it earns the inversion warning below.

### The success label inversion on rpc_call_latency

`rpc_call_latency`'s `success` label is set from `strconv.FormatBool(err != nil)` in
chainlink-framework's `RecordRequest`. That means **`success="true"` is the label attached to a
failed call** (an error was present) and `success="false"` to a successful one. This is the
opposite of intuition and of the metric's own `Help` text. If you write `{success="true"}` because
you want "the successful calls," you get the failures. For latency-by-failure, author your queries
against `success="true"` (err != nil) and `success="false"` (clean) -- and call out this inversion
in any finding so the human doesn't re-derive it from the metric name. (The older
`evm_pool_rpc_node_rpc_call_time` histogram shares the same `success` semantics on the promauto
side.)

### Which TXM metrics actually carry a chainID label (not all do)

**What is actually live in a CCIP EVM deployment:** the counters `tx_manager_num_broadcasted_total`,
`tx_manager_num_confirmed_transactions_total`, `tx_manager_num_successful_transactions_total`,
`tx_manager_num_finalized_transactions_total`, the gauges `txm_pending_tx_queue_utilization`,
`tx_manager_tx_oldest_non_terminal_age_seconds`, `tx_manager_tx_attempt_count`, and (on a revert)
`tx_manager_num_tx_reverted_total`. **Every one of these carries a `chainID` label**, 
so all write-path queries in this doc are chain-scoped reliably. They appear
only on the **destination (tx-sending) chain** -- the same chain the commit plugin transmits on --
so filter `$chainID` against that dest chain specifically.

Concretely, the live family maps onto the EVM V1 vocabulary:
- no per-address nonce-gap counter → the "one stuck thing" read is `txm_pending_tx_queue_utilization`
  climbing + `tx_manager_tx_oldest_non_terminal_age_seconds` climbing + `tx_manager_tx_attempt_count`
  rising, *not* the V2 `txm_num_nonce_gaps`;
- no lifecycle-failure stage breakdown → owe the *kind* of stuck (`max_in_flight` vs `broadcast`)
  to logs/TXM config, not to a metric;
- broadcast-vs-confirm gap is `tx_manager_num_broadcasted_total` vs `tx_manager_num_confirmed_transactions_total`.

### Empty result sets

No metric in this doc is an unconditional per-round heartbeat like the commit plugin's -- most are
event counters that only get a time series after they fire, or gauges reported on their own
ticker. The same two rules from `commit-plugin-health.md#empty-result-sets` apply:

1. The **continuous gauges** -- `evm_log_poller_last_processed_block`, `evm_pool_rpc_node_calls_*`
   (counters, but accumulate rapidly), `gas_updater_current_base_fee`, and the txmgr gauges
   `txm_pending_tx_queue_utilization`, `tx_manager_tx_oldest_non_terminal_age_seconds`,
   `tx_manager_tx_attempt_count` -- are all emitted continuously while the process runs (the gauges)
   or as fast-moving counters, so an **empty** result means the metric isn't reaching your
   datasource (or you're looking at a chain the node isn't sending txs on -- the gas/TXM family only
   exists on the destination/tx-sending chain), not "value zero." Report `UNKNOWN`, don't grade it.
2. Everything else is an **event counter** fired only on occurrence
   (`evm_pool_rpc_node_calls_failed`, `tx_manager_num_*_total`, `tx_manager_num_tx_reverted_total`,
   `block_history_estimator_connectivity_failure_count`): **empty == the event never happened ==
   value 0.** Grade normally. The one exception is the same as the commit docs: if a `UNKNOWN`-class
   gauge in the same run is itself empty, don't trust counter-emptiness -- say so in `concerns`.

`rpc_call_latency_milliseconds` is a histogram; if you only have `_count`/`_bucket` forms on your
datasource, derive `rate()`s and quantiles from those, treating an absent `_count` for a window as
"no calls (or none clearing the first bucket) in that window."

### Counter _total suffixing

OTel counters that reach Prometheus via the otel-collector's
`prometheusremotewrite` exporter get a `_total` appended **unless the name already ends in `_total`**.
`evm_pool_rpc_node_calls_total` (already suffixed) stayed as-is, while `evm_pool_rpc_node_calls_success`
became `evm_pool_rpc_node_calls_success_total` and `tx_manager_num_broadcasted` became
`tx_manager_num_broadcasted_total`. Two unit-suffix gotchas also seen live: the framework `rpc_call_latency`
histogram reaches Prometheus as **`rpc_call_latency_milliseconds_*`** (OTel appends the `ms` unit),
and the `gas_updater_*` gauges carry `percentile="60%"`-style values, not `p25/p50`. **Check one
query against your datasource's metric browser before trusting the rest.** The doc lists the metric
as registered; your pipeline may suffix/rename it.

## Decision graph

```yaml
# entry defaults to "data_source"; set to "report_transmission" when dropped from that path.
steps:
      # ---- READ PATH (data source) ----
  - id: A0
    check: rpc_node_pool_health
    queries:
      - 'sum(rate(evm_pool_rpc_node_calls_failed{chainID="$chainID"}[5m])) by (nodeName)'
      - 'sum(rate(evm_pool_rpc_node_calls_total{chainID="$chainID"}[5m])) by (nodeName)'
    condition: "any nodeName shows a failed/total ratio clearly above its peers, OR all nodeNames show failures"
    if_true:
      one_bad_node: {action: "REPORT:chain-evm-oncall", reason: "a single RPC node is failing calls that its peers aren't (nodeName differs) -- per-node RPC/infra issue, not the chain or the pool. The commit plugin auto-selects away from a failed node, so this usually self-heals, but a persistently failing primary stalls reads until failover. Hand off with the failing nodeName."}
      all_nodes:     {action: "REPORT:chain-evm-oncall", reason: "every node in the pool is failing RPC calls for $chainID -- this is a chain-wide or pool-config outage, and it will surface upstream as reader read_outcome_error / ccip_commit_fchain_read_errors / onramp lag (uncommitted-message.md step2/step2c). Do not attribute it to the commit plugin."}
    if_false: {action: "CONTINUE:A1", reason: "no node's failure ratio warrants a finding; move to log-poller catch-up"}
    automatable: true

  - id: A1
    check: log_poller_catchup
    query: 'evm_log_poller_last_processed_block{chainID="$chainID"}'
    condition: "value stale (no sample in the last ~2 poll intervals) OR lagging the chain head by more than your environment's tolerance"
    if_true: {action: "REPORT:chain-evm-oncall", reason: "the log poller for $chainID is stuck (no fresh samples) or far behind -- the reader's per-filter queries read logs through it, so a wedge here produces exactly the read_empty / MsgsBetweenSeqNums-empty / config_poller batch failures the commit runbooks attribute to 'data source'. Cross-check the reader metrics (ccip_reader_read_empty_total, ccip_reader_config_poll_failure_total reason=\"batch_fetch_failed\") to confirm they coincide with this log-poller stall before reporting."}
    if_false: {action: "CONTINUE:$entry == \"report_transmission\" ? B0 : STOP", reason: "read path is healthy for $chainID. If you came in via data_source, stop here; if via report_transmission, continue to the write path."}
    automatable: true

  # ---- WRITE PATH (report transmission) ----
  - id: B0
    check: fee_layer_can_send
    queries:
      - 'sum(rate(block_history_estimator_connectivity_failure_count{chainID="$chainID"}[15m])) by (mode)'
      - 'max(gas_updater_current_base_fee{chainID="$chainID"})'
    condition: "connectivity_failure_count any mode > 0, OR base_fee gauge missing/zero when the chain is EIP-1559"
    if_true: {action: "CONTINUE:B0_detail", reason: "gas estimation flagged a connectivity failure (bumps being suppressed) or has no base fee to price against -- TXM may refuse to bump or underprice"}
    if_false: {action: "CONTINUE:B1", reason: "fee layer looks able to price sends; move to broadcast"}
    automatable: true

  - id: B0_detail
    check: connectivity_failure_mode
    query: 'sum(rate(block_history_estimator_connectivity_failure_count{chainID="$chainID"}[15m])) by (mode)'
    condition: "any series > 0"
    if_true: {action: "REPORT:chain-evm-oncall", reason: "gas bumps are being suppressed for $chainID due to a detected network propagation/connectivity issue (mode=legacy or mode=eip1559). This frequently pairs with report_transmission_gave_up: TXM keeps trying to bump but is blocked from re-pricing, so reports underprice and never land. Treat fee-layer connectivity as the root cause for the transmission give-up, not a secondary observation."}
    if_false: {action: "REPORT:chain-evm-oncall", reason: "no connectivity-failure counter, but base fee is missing/zero on an EIP-1559 chain -- the estimator isn't obtaining a usable base fee, so TXM cannot price EIP-1559 sends correctly."}
    automatable: true

  - id: B1
    check: txm_broadcast_flow
    queries:
      - 'sum(rate(tx_manager_num_confirmed_transactions{chainID="$chainID"}[15m]))'
      - 'sum(rate(tx_manager_num_broadcasted{chainID="$chainID"}[15m]))'
    condition: "confirmed == 0 AND broadcasted == 0 over the window"
    note: "these are the framework/txmgr counters actually emitted by a CCIP EVM node (verified live -- the V2 txm_* variants are NOT emitted, see the cheat sheet). All carry chainID, so {chainID=...} is reliable here. Both are counters with a _total suffix on an OTel pipeline."
    if_true: {action: "CONTINUE:B2", reason: "nothing being broadcast at all -- the give-up is a 'no reports are even being sent' situation, dig into the queue/stuck gauges"}
    if_false:
      confirmed_zero_broadcasted_positive: {action: "CONTINUE:B2", reason: "reports ARE being broadcast but NONE confirm in window -- a land/drop/revert problem, not a send problem"}
      both_positive: {action: "STOP", reason: "TXN is broadcasting and confirming reports normally for $chainID. If report_transmission_gave_up still fires, it's almost certainly a per-report revert or a backlog-size / round-cadence issue, not the EVM node stack. Escalate with this evidence rather than assuming the EVM layer broken. (If you were handed a specific $txHash, sanity-check it on the explorer before stopping.)"}
    automatable: true

  - id: B2
    check: txm_stuck_signals
    queries:
      - 'max(txm_pending_tx_queue_utilization{chainID="$chainID"})'
      - 'max(tx_manager_tx_oldest_non_terminal_age_seconds{chainID="$chainID"})'
      - 'max(tx_manager_tx_attempt_count{chainID="$chainID"})'
    condition: "queue_utilization rising toward 1, OR oldest_non_terminal_age climbing (well above normal block-time), OR attempt_count climbing"
    note: "these three framework gauges (live on the dest chain) are the 'one stuck thing' read this deployment actually exposes -- there is no nonce-gap or lifecycle-stage metric in the live family (V2 txm_* is not emitted). Treat 'queue near full' + 'oldest non-terminal age climbing' + 'attempts rising' together as the stuck signature."
    if_true: {action: "REPORT:chain-evm-oncall", reason: "the TXM queue is backing up / the oldest tx is aging / attempts are rising -- a stuck transaction is blocking the send path, which surfaces upstream as report_transmission_gave_up. Say which of the three gauges is the driver and its value. Pair with B0 if the fee layer is also suppressing bumps."}
    if_false: {action: "STOP", reason: "no nonce/queue stuck signature and nothing aging. Broadcasts were happening but nothing confirms and no stuck marker is firing -- most likely an RPC/indexing lag on the confirmation side (tx landed but isn't seen as finalized). Verify against the chain explorer / $txHash before escalating."}
    automatable: true
```

## Steps

### Stage A0 -- RPC client pool: are the nodes answering?

The commit reader's reads and the TXM's sends both go through the EVM pool's RPC clients. Group
the pool counters by node to tell "one bad RPC" from "whole pool down":

```promql
sum(rate(evm_pool_rpc_node_calls_failed{chainID="$chainID"}[5m])) by (nodeName)
sum(rate(evm_pool_rpc_node_calls_total{chainID="$chainID"}[5m])) by (nodeName)
```

Corroborate at the wire level with per-node latency (the framework latency instrument; on an OTel
pipeline it surfaces as `rpc_call_latency_milliseconds_*`, and `success="true"` means *failed* --
see [the inversion](#the-success-label-inversion-on-rpc_call_latency)):

```promql
histogram_quantile(0.95, sum by (le) (rate(rpc_call_latency_milliseconds_bucket{chainID="$chainID", success="true"}[5m])))
```

A single `nodeName` failing while peers pass → hand off that node (`chain-evm-oncall`). Every node
failing → a pool/chain-wide outage that will masquerade as reader `read_outcome_error` /
`fchain_read_errors` / onramp lag in the commit docs -- own it as a chain-family finding, not a
plugin one.

### Stage A1 -- log poller: is it keeping up with the chain?

`ccip`'s reader ingests on-chain events (message sends, commits, config) through the EVM **log
poller**. A log poller wedge is the single most common *silent* false-idle cause on an EVM dest
chain: every per-filter log read returns empty, so the reader reports clean `read_empty` /
`chain_gap` and the commit plugin sits idle looking healthy.

The one gauge is `evm_log_poller_last_processed_block`:

```promql
evm_log_poller_last_processed_block{chainID="$chainID"}
```

Two failure shapes:
- **No fresh samples** (stale for a couple of poll intervals): the poller goroutine itself is
  wedged. This is the more dangerous one because `last_processed_block` parks at its last value
  and looks "normal, just old."
- **Lagging block head**: the poller is running but can't keep up (slow RPC, backlog of
  reorg/dedupe work). Eventually bounded if the RPC catches up; unbounded if it doesn't.

Either way report `chain-evm-oncall`, and corroborate with the reader's own signals
(`ccip_reader_read_empty_total`, `ccip_reader_config_poll_failure_total{reason="batch_fetch_failed"}`)
so the finding lands on the EVM layer rather than being re-attributed upstairs.

### Stage B0 -- gas/fee layer: can it price and afford a send?

Commit-report transactions must be priced. Two estimator families feed the TXM:
- **block-history estimator** (EIP-1559 and legacy price percentiles via `gas_updater_*`
  gauges, plus the connectivity-failure counter)
- **fee-history estimator** (`base_fee_updater`, `gas_price_updater`,
  `max_fee_per_gas_updater`, `max_priority_fee_per_gas_updater` gauges)

The actionable signal is the **connectivity-failure counter** -- the estimator *detected* it
couldn't trust the fee data and is refusing to bump, which starves transmission at the pricing
layer:

```promql
sum(rate(block_history_estimator_connectivity_failure_count{chainID="$chainID"}[15m])) by (mode)
```

`mode` is `legacy` or `eip1559`. If it fires alongside `report_transmission_gave_up`, the bump
was suppressed by the fee estimator's connectivity guard -- root cause at the fee layer, not a
TXM logic bug. Also sanity-check `gas_updater_current_base_fee` exists and is nonzero
on an EIP-1559 chain; a missing base fee means the estimator has nothing to price against.

### Stage B1 -- TXM broadcast: are reports being sent at all?

The two live counters that divide "nothing is being sent" from "sent but not landing":

```promql
sum(rate(tx_manager_num_confirmed_transactions{chainID="$chainID"}[15m]))
sum(rate(tx_manager_num_broadcasted{chainID="$chainID"}[15m]))
```

- Both `> 0` → the node stack is broadcasting **and confirming** reports normally. A persisted
  `report_transmission_gave_up` together with a healthy TXM is a per-report revert or a
  backlog/round-cadence issue, not the EVM stack. Stop here (verify the actual tx on-chain if you
  have a `$txHash`).
- Broadcasted `> 0`, confirmed `== 0` → reports go out but never land → **B2** (stuck/revert).
- Both `== 0` → nothing is being sent at all → **B2** (stuck/never-created).

`tx_manager_num_finalized_transactions` / `tx_manager_num_successful_transactions` are the
supporting confirmation counters (finalized = past finality horizon; successful = included and
executed); a widening broadcast→confirmed↔finalized gap with confirms stuck at zero is the "sent
but not landing" signature. (There is no `_time_until_confirmed` histogram in the live family --
the V2 `txm_time_until_tx_confirmed` is not emitted -- so latency has to be inferred from how long
the broadcast counter advances before the confirmed counter follows.)

### Stage B2 -- TXM stuck: queue back-up, oldest-tx age, attempt count

When broadcasts aren't confirming, name *why* from the three framework gauges the deployment
actually exposes -- there is **no** V2 lifecycle-stage or nonce-gap metric in the live family:
`txm_pending_tx_queue_utilization` (queue fullness), `tx_manager_tx_oldest_non_terminal_age_seconds`
(age of the oldest not-yet-terminal tx), and `tx_manager_tx_attempt_count` (how many attempts are
stacked up, ~retries/bumps):

```promql
max(txm_pending_tx_queue_utilization{chainID="$chainID"})
max(tx_manager_tx_oldest_non_terminal_age_seconds{chainID="$chainID"})
max(tx_manager_tx_attempt_count{chainID="$chainID"})
```

The **stuck signature** is the combination: queue near `1.0` (or climbing), oldest-non-terminal-age
far above a healthy block time (climbing → the stuck thing is old, not just momentarily delayed),
and attempt-count climbing (the TXM keeps re-attempting/bumping the same txs). This is the EVM
analogue of a `reached_max_attempts` / "parked nonce" outage -- one blocked transaction holds the
send path open and the backlog up, so `report_transmission_gave_up` fires even while the reader is
healthy. The fix lives in the EVM operation (identify the stuck tx, recover the sequence), not in
the commit plugin. On a revert, additionally check `tx_manager_num_tx_reverted_total` rising.

Cross-check the *kind* of backing-up from the log/tx-sender rather than a metric: in this family
there is no `max_in_flight`-vs-`broadcast` stage breakdown like the unemitted V2
`txm_transaction_lifecycle_failure_total` would give -- if you need to distinguish "throttling" from
"RPC rejecting sends," pivot to the RPC pool (`evm_pool_rpc_node_calls_failed`) and the fee layer
(B0) for the mechanism.

If nothing in B2 fires at all but confirms are still zero, the remaining honest conclusion is an
RPC confirmation/indexing lag rather than a TXM failure -- say so and verify against the explorer
rather than inventing a stuck marker.

## Metric reference

`otel_type` governs `_total`/`_bucket` suffixing (see [the note](#counter-total-suffixing)).
`dual` means promauto + beholder side-by-side; `beholder-only` means only the OTel/beholder path
(chainlink-evm patterns mirror the commit-plugin ones in `uncommitted-message.md#metric-reference`).
Many TXM beholder series carry **no `chainID` attribute** -- see
[which TXM metrics carry chainID](#which-txm-metrics-actually-carry-a-chainid-label-not-all-do).

`otel_type` governs `_total`/`_bucket` suffixing; **status** reflects what was confirmed live (or
absent) on a `ccip obs up --victoria` devel env while reports were transmitting:

| metric | otel_type | key labels | status in a CCIP EVM deploy |
|---|---|---|---|
| `evm_pool_rpc_node_calls_total` | Counter | `chainID,nodeName,rpcDomain,callName` (beholder) | **live** |
| `evm_pool_rpc_node_calls_success` | Counter | same | **live** |
| `evm_pool_rpc_node_calls_failed` | Counter | same | **live** (fires only on failure) |
| `rpc_call_latency` | Histogram | `chainFamily,chainID,rpcDomain,isSendOnly,success,callName` (success inverted) | **live** -- surfaces as `rpc_call_latency_milliseconds_*` |
| `evm_log_poller_last_processed_block` | Gauge | `chainFamily,chainID` | **live** |
| `gas_updater_set_gas_price` | Gauge | `chainID,percentile` (e.g. `"60%"`) | **live** (dest/tx-sending chain) |
| `gas_updater_all_gas_price_percentiles` | Gauge | `chainID,percentile` | **live** (dest chain) |
| `gas_updater_current_base_fee` | Gauge | `chainID` | **live** (dest chain) |
| `gas_updater_set_tip_cap` / `gas_updater_all_tip_cap_percentiles` | Gauge | `chainID,percentile` | conditional -- only EIP-1559; absent on a legacy chain |
| `block_history_estimator_connectivity_failure_count` | Counter | `chainID,mode` | **live** (fires only on a detected connectivity failure) |
| `gas_price_updater` / `base_fee_updater` / `max_priority_fee_per_gas_updater` / `max_fee_per_gas_updater` | Gauge | `chainID` | **not emitted** here -- these are the *fee-history* estimator, which the local block-history config does not run |
| `txm_pending_tx_queue_utilization` | Gauge | `chainID` | **live** |
| `tx_manager_tx_oldest_non_terminal_age_seconds` | Gauge | `chainID` | **live** |
| `tx_manager_tx_attempt_count` | Gauge | `chainID` | **live** |
| `tx_manager_num_broadcasted` | Counter | `chainID` | **live** (`_total` suffix) |
| `tx_manager_num_confirmed_transactions` | Counter | `chainID` | **live** (`_total` suffix) |
| `tx_manager_num_successful_transactions` | Counter | `chainID` | **live** (`_total` suffix) |
| `tx_manager_num_finalized_transactions` | Counter | `chainID` | **live** (`_total` suffix) |
| `tx_manager_num_tx_reverted` | Counter | `chainID` | **live** (fires only on a revert) |
| `tx_manager_fwd_tx_count` | Counter | `chainID` | conditional (forwarding enabled) |
| `tx_manager_num_gas_bumps` / `tx_manager_gas_bump_exceeds_limit` | Counter | `chainID` | conditional (fires only on a bump) |
| `txm_num_broadcasted_transactions`, `txm_num_confirmed_transactions`, `txm_num_nonce_gaps`, `txm_reached_max_attempts`, `txm_time_until_tx_confirmed`, `txm_rpc_nonce`, `txm_transaction_lifecycle_failure_total` | V2 instruments | (V2 code would omit `chainID` on beholder) | **NOT emitted** -- V2 engine not the active transmit path; do not build on these |
| `txm_multicall_duration_ms` | Histogram | `chainID,method,blockTag,success,timedOut` | **NOT emitted** here (V2 wrapper path) |
| `ofa_send_tx_status` / `ofa_send_tx_latency` | Counter / Histogram | `chainID,backend,status` | MEV/dual-broadcast only |
| `meta_endpoint_status_codes` / `meta_endpoint_latency` | Counter / Histogram | `chainID,feedAddress` | MEV/dual-broadcast only |
| `meta_bids_per_transaction` / `meta_errors` | Histogram / Counter | `chainID,feedAddress,errorType` | MEV/dual-broadcast only |

Read/interp notes for the table:

- The **live** markers are empirical (local devel env, node 2.61.0, chains 1337/2337). A "heavy"
  row absent on *your* datasource means it's the **not emitted** class, not a healthy zero -- check
  the class before reporting.
- **`gas_updater_*` and the whole TXM family exist only on the destination/tx-sending chain** --
  a query scoped to a chain the node doesn't send txs on legitimately returns empty.
- The `ofa_*`/`meta_*` families are only live on deployments that route sends through a
  dual-broadcast/MEV backend (Flashbots/Nova/FastLane Atlas); on a plain NOP commit deployment they
  will legitimately be absent.
- The commit-plugin's *own* metrics (the `ccip_commit_*`/`ccip_reader_*` series
  `commit-plugin-health.md` and `uncommitted-message.md` query) are **not** part of this doc --
  this doc is the layer underneath, reached by their `REPORT:...-oncall` handoffs.

## Output contract

```yaml
evm_finding:
  chainID: string
  entry: data_source | report_transmission
  timestamp: string           # ISO8601
  verdict: OK | ISSUE
  steps_run:
    - id: string              # A0 | A0_detail | A1 | B0 | B0_detail | B1 | B2 | B2_detail | B2_nonce
      outcome: STOP | CONTINUE:<step> | REPORT:<owner>
      value: string           # the raw number(s) that drove the outcome -- carry the gauges/counters as-is
  findings:                   # one per REPORT; empty list if none
    - owner: string           # chain-evm-oncall, or whoever stage named
      summary: string
      stage: string           # name the lifecycle stage / RPC node / queue that told you
      evidence: string        # the exact query + result that produced it
      commit_handoff: string  # which commit-runbook owner/check this replaces or justifies
```

Do not page, escalate, or take remediation -- your job ends by handing this report to the named
owner. The commit runbook that dropped you here already decided *whether* there's a problem; this
doc only locates *which EVM-layer component* owns it.
