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
  (RPC), or the tx is stuck/reverting (nonce gap, max attempts, lifecycle failure). Stages
  **B0/B1/B2**.

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
| `chainID` | all EVM beholder metrics | the EVM chain ID (== commit's `destChainID` value for the dest chain) |
| `evmChainID` | the promauto pool counters (`evm_pool_rpc_node_calls_*`), gas-updater gauges (promauto side only) | same value, old promauto spelling. Only on the direct-scrape path, not the beholder path |
| `nodeName` | `evm_pool_rpc_node_calls_*` | which RPC node -- use this to split a single bad RPC from the whole pool failing |
| `rpcDomain` | `evm_pool_rpc_node_calls_*` (beholder), `rpc_call_latency` | the RPC endpoint/domain (sanitized) the call went to |
| `callName` / `rpcCallName` | RPC metrics | which JSON-RPC method / logical call |
| `address` | `txm_rpc_nonce` | the wallet address whose on-chain nonce is reported |
| `stage` | `txm_transaction_lifecycle_failure_total` | where in the lifecycle the failure happened (`create`, `in_flight_subset`, `max_in_flight`, `broadcast`, `nonce_at`) |
| `percentile` | the gas-updater gauges | which price percentile (`p25`, `p50`, ...) the gauge carries |
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

`pkg/txm/metrics.go` constructs a `metrics.Labeler` with `chainID` and carries it on the struct --
but the beholder `.Add(...)`/`.Record(...)` calls for **most** of the `txm_*` instruments pass
**no attributes** (`m.numBroadcastedTxs.Add(ctx, 1)`, etc.), so those beholder series get only the
beholder client's inherited identity labels (`node_id`, `csa_public_key`), **no `chainID`**. Only
**`txm_transaction_lifecycle_failure_total`** explicitly attaches `chainID` + `stage`. The
consequences, and the escape hatches:

- On a **beholder-fed** datasource, the counters `txm_num_broadcasted_transactions`,
  `txm_num_confirmed_transactions`, `txm_num_nonce_gaps`, the gauge `txm_reached_max_attempts`,
  the histogram `txm_time_until_tx_confirmed`, and the gauge `txm_rpc_nonce` **are not filterable
  by `$chainID`**. Queries against them on a node hosting multiple chains will mix chains. Scope
  them per `node_id`/`csa_public_key` instead, or use the **`txm_transaction_lifecycle_failure_total`**
  (`chainID` + `stage`) as the chain-scoped backbone of stage B2.
- The **promauto** dual-writes for all of these **do** carry `chainID` (and `txm_rpc_nonce` adds
  `address`). If your datasource scrapes the node's prometheus endpoint directly, prefer the
  pramaauto names for chain-scoped TXM questions.

The queries in this doc are written to be correct on either path, but the label caveat is why stage
B2 leans on `txm_transaction_lifecycle_failure_total` rather than assuming every TXM counter is
chain-filterable.

### Empty result sets

No metric in this doc is an unconditional per-round heartbeat like the commit plugin's -- most are
event counters that only get a time series after they fire, or gauges reported on their own
ticker. The same two rules from `commit-plugin-health.md#empty-result-sets` apply:

1. The **staleness/liveness gauges** -- `evm_log_poller_last_processed_block` (A1),
   `txm_reached_max_attempts`, `txm_rpc_nonce` (B2) -- are reported continuously while the process
   runs, so an **empty** result (no series at all) means the metric isn't reaching your datasource,
   not "value zero." Report `UNKNOWN` there, don't grade it against thresholds.
2. Everything else is an **event counter** (`evm_pool_rpc_node_calls_*`, `rpc_call_latency`
   buckets, `block_history_estimator_connectivity_failure_count`, `txm_num_*`,
   `txm_transaction_lifecycle_failure_total`): **empty == the event never happened == value 0.**
   Grade normally. The one exception is the same as the commit docs: if a `UNKNOWN`-class gauge in
   the same run is itself empty, don't trust counter-emptiness -- say so in `concerns`.

`rpc_call_latency` and `txm_time_until_tx_confirmed` are histograms; if you only have their
`_count`/`_bucket` forms on your datasource, derive `rate()`s and quantiles from those, treating an
absent `_count` for a window as "no calls (or none clearing the first bucket) in that window."

### Counter _total suffixing

Same pipeline caveat as the commit runbooks: OTel counters that reach Prometheus via a
prometheusremotewrite/OTel-collector exporter get a `_total` appended **unless the name already
ends in `_total`**. Several EVM metrics are registered with `_total` already baked in
(`evm_pool_rpc_node_calls_total`, `txm_transaction_lifecycle_failure_total`,
`block_history_estimator_connectivity_failure_count` does *not*), so on such a datasource these
names are *not* doubled -- but the fee/other counters also registered without `_total`
(`gas_updater_set_gas_price` is a gauge; `txm_num_nonce_gaps` is a counter and would gain a
`_total`). **Check one query against your datasource's metric browser before trusting the rest.**
The doc lists the metric as registered; your pipeline may suffix it.

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
      - 'sum(rate(txm_num_confirmed_transactions{chainID="$chainID"}[15m]))'
      - 'sum(rate(txm_num_broadcasted_transactions{chainID="$chainID"}[15m]))'
    condition: "confirmed == 0 AND broadcasted == 0 over the window"
    note: "these two counters do NOT carry chainID on a beholder-only datasource (see the label cheat sheet) -- the {chainID=...} filter here works only if your datasource also has the promauto forms. If {chainID=...} returns empty on a datasource you believe is beholder-only, re-run WITHOUT chainID and split by node_id/csa_public_key instead; an empty-result treated as 'zero' here is exactly the trap the cheat sheet warns about."
    if_true: {action: "CONTINUE:B2", reason: "nothing being broadcast at all -- the give-up is a 'no reports are even being sent' situation, dig into stuck/lifecycle markers"}
    if_false:
      confirmed_zero_broadcasted_positive: {action: "CONTINUE:B2", reason: "reports ARE being broadcast but NONE confirm in window -- a land/drop/revert problem, not a send problem"}
      both_positive: {action: "STOP", reason: "TXN is broadcasting and confirming reports normally for $chainID. If report_transmission_gave_up still fires, it's almost certainly a per-report revert or a backlog-size / round-cadence issue, not the EVM node stack. Escalate with this evidence rather than assuming the EVM layer broken. (If you were handed a specific $txHash, sanity-check it on the explorer before stopping.)"}
    automatable: true

  - id: B2
    check: txm_stuck_signals
    query: 'sum by (stage) (rate(txm_transaction_lifecycle_failure_total{chainID="$chainID"}[15m]))'
    condition: "any series > 0"
    if_true: {action: "CONTINUE:B2_detail", reason: "transaction lifecycle failures present; stage names the phase"}
    if_false: {action: "CONTINUE:B2_nonce", reason: "no lifecycle-failure counter; the stuck signal (give-up, zero confirms) must come from nonce or max-attempts instead"}
    automatable: true

  - id: B2_detail
    check: lifecycle_failure_stage
    query: 'sum by (stage) (rate(txm_transaction_lifecycle_failure_total{chainID="$chainID"}[15m]))'
    condition: "report the dominant stage(s)"
    if_true:
      nonce_at:       {action: "REPORT:chain-evm-oncall", reason: "lifecycle failing at the nonce stage -- nonce tracking issue. Cross-check txm_rpc_nonce (per address) vs expected and the nonce gap counter (nonce_gap) below."}
      max_in_flight:  {action: "REPORT:chain-evm-oncall", reason: "lifecycle failing at max_in_flight -- the TXM's max inflight budget is exhausted (slow finality or too many concurrent reports). Throttling, not an RPC failure."}
      broadcast:      {action: "REPORT:chain-evm-oncall", reason: "lifecycle failing at broadcast -- the RPC rejected/dropped the send; corroborate with evm_pool_rpc_node_calls_failed at the same time."}
      create:         {action: "REPORT:chain-evm-oncall", reason: "lifecycle failing at create -- the tx could not be constructed (e.g. missing wallet/crypto, keystore). Check infrastructure, not the chain."}
      default:        {action: "REPORT:chain-evm-oncall", reason: "dominant/most relevant of the above, or a stage not named here -- report the exact stage string with the rate, don't guess."}
    if_false: {action: "CONTINUE:B2_nonce", reason: "no strongly dominant stage; continue to the nonce/max-attempt check"}
    automatable: true

  - id: B2_nonce
    check: nonce_gaps_and_max_attempts
    queries:
      - 'sum(rate(txm_num_nonce_gaps{chainID="$chainID"}[15m]))'
      - 'max(txm_reached_max_attempts{chainID="$chainID"})'
    condition: "nonce gaps rate > 0 OR reached_max_attempts == 1"
    note: "same chainID caveat as B1 -- txm_num_nonce_gaps / txm_reached_max_attempts / txm_rpc_nonce lack chainID on beholder-only datasources. Prefer the chainID-bearing txm_transaction_lifecycle_failure_total from B2 when the datasource gives you no promauto forms; treat an empty {chainID=...} result here as 'cannot scope', not 'zero'."
    if_true:
      nonce_gaps:
        action: "REPORT:chain-evm-oncall"
        reason: "the TXM is filling nonce gaps for $chainID -- a stalled out-of-order tx is parking the sender's nonce and blocking subsequent reports (classic report_transmission_gave_up root cause). Cross-check txm_rpc_nonce per address to see the parked nonce."
        followup_query: 'txm_rpc_nonce{chainID="$chainID"}'
      max_attempts:
        action: "REPORT:chain-evm-oncall"
        reason: "reached_max_attempts == 1 -- the TXM gave up on a tx after exhausting its bump budget (combined with B0 if the bump was suppressed by a connectivity failure, or with B2's broadcast stage if the RPC dropped it). Any report behind that tx will report_transmission_gave_up until the wallet nonce recovers."
    if_false: {action: "STOP", reason: "no nonce gaps and no max-attempt kill. Broadcasts were happening but nothing confirms and no stuck marker is firing -- this is most likely an RPC/indexing lag on the confirmation side (tx landed but isn't seen as finalized). Verify against the chain explorer / $txHash before escalating further."}
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

Optionally corroborate at the wire level with the chain-client multicall latency histogram
(`txm_multicall_duration_ms{success="false"}`, beholder-only) if your datasource has it, and with
per-node latency (beholder histogram from chainlink-framework; `success="true"` means *failed* --
see [the inversion](#the-success-label-inversion-on-rpc_call_latency)):

```promql
histogram_quantile(0.95, sum by (le) (rate(rpc_call_latency_bucket{chainID="$chainID", callName=~".*", success="true"}[5m])))
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

The two counters that divide "nothing is being sent" from "sent but not landing":

```promql
sum(rate(txm_num_confirmed_transactions{chainID="$chainID"}[15m]))
sum(rate(txm_num_broadcasted_transactions{chainID="$chainID"}[15m]))
```

- Both `> 0` → the node stack is broadcasting **and confirming** reports normally. A persisted
  `report_transmission_gave_up` together with a healthy TXM is a per-report revert or a
  backlog/round-cadence issue, not the EVM stack. Stop here (verify the actual tx on-chain if you
  have a `$txHash`).
- Broadcasted `> 0`, confirmed `== 0` → reports go out but never land → **B2** (land/drop/revert /
  nonce).
- Both `== 0` → nothing is being sent at all → **B2** (stuck/never-created).

`txm_time_until_tx_confirmed` (histogram) is the supporting latency signal; a healthy chain spends
most of it near the chain's block time, and a value pinned at the histogram ceiling means confirms
are being delayed, not that everything is fine.

### Stage B2 -- TXM stuck: killed attempts, nonce gaps, lifecycle failures

When broadcasts aren't confirming, name *why*. The chain-scoped backbone is the lifecycle-failure
counter (the one TXM metric that reliably carries `chainID`):

```promql
sum by (stage) (rate(txm_transaction_lifecycle_failure_total{chainID="$chainID"}[15m]))
```

`stage` values: `create` (couldn't construct the tx -- keystore/wallet/infra),
`in_flight_subset` / `max_in_flight` (concurrency budget exhausted -- throttling), `broadcast`
(RPC rejected/dropped the send), `nonce_at` (nonce tracking broken). Pair each with the matching
corroboration from the note on the decision-graph step.

Then the classic stuck-nonce pair:

```promql
sum(rate(txm_num_nonce_gaps{chainID="$chainID"}[15m]))
max(txm_reached_max_attempts{chainID="$chainID"})
txm_rpc_nonce{chainID="$chainID"}   # per wallet address
```

A **nonce gap** (or `reached_max_attempts==1`) means one parked out-of-order tx is holding the
sender's nonce, so every subsequent report (a) can't get its required nonce and (b) keeps bouncing
off the RPC with a nonce-too-low/known-tx rejection. That is the archetypal
`report_transmission_gave_up` + zero-confirms combination on EVM, and the fix lives in the EVM
operation (find/nothing the stuck tx, recover the nonce), not in the commit plugin.

If nothing in B2 fires at all but confirms are still zero, the remaining honest conclusion is an
RPC confirmation/indexing lag rather than a TXM failure -- say so and verify against the explorer
rather than inventing a stuck marker.

## Metric reference

`otel_type` governs `_total`/`_bucket` suffixing (see [the note](#counter-total-suffixing)).
`dual` means promauto + beholder side-by-side; `beholder-only` means only the OTel/beholder path
(chainlink-evm patterns mirror the commit-plugin ones in `uncommitted-message.md#metric-reference`).
Many TXM beholder series carry **no `chainID` attribute** -- see
[which TXM metrics carry chainID](#which-txm-metrics-actually-carry-a-chainid-label-not-all-do).

| metric | otel_type | key labels | registration |
|---|---|---|---|
| `evm_pool_rpc_node_calls_total` | Counter | beholder: `chainID,nodeName,rpcDomain,callName`; promauto: `evmChainID,nodeName` | dual |
| `evm_pool_rpc_node_calls_success` | Counter | same as above | dual |
| `evm_pool_rpc_node_calls_failed` | Counter | same as above | dual |
| `rpc_call_latency` | Histogram | `chainFamily,chainID,rpcDomain,isSendOnly,success,callName` (success inverted) | dual (chainlink-framework) |
| `evm_log_poller_last_processed_block` | Gauge | `chainFamily,chainID` | dual |
| `gas_updater_set_gas_price` | Gauge | `chainID,percentile` | dual |
| `gas_updater_set_tip_cap` | Gauge | `chainID,percentile` | dual |
| `gas_updater_all_gas_price_percentiles` | Gauge | `chainID,percentile` | dual |
| `gas_updater_all_tip_cap_percentiles` | Gauge | `chainID,percentile` | dual |
| `gas_updater_current_base_fee` | Gauge | `chainID` | dual |
| `block_history_estimator_connectivity_failure_count` | Counter | `chainID,mode` (`legacy`\|`eip1559`) | dual |
| `gas_price_updater` | Gauge | `chainID` | dual |
| `base_fee_updater` | Gauge | `chainID` | dual |
| `max_priority_fee_per_gas_updater` | Gauge | `chainID` | dual |
| `max_fee_per_gas_updater` | Gauge | `chainID` | dual |
| `txm_transaction_lifecycle_failure_total` | Counter | `chainID,stage` | both (chainID attached) |
| `txm_num_broadcasted_transactions` | Counter | **no chainID on beholder**; promauto `chainID` | both |
| `txm_num_confirmed_transactions` | Counter | **no chainID on beholder**; promauto `chainID` | both |
| `txm_num_nonce_gaps` | Counter | **no chainID on beholder**; promauto `chainID` | both |
| `txm_reached_max_attempts` | Gauge | **no chainID on beholder**; promauto `chainID` | both |
| `txm_time_until_tx_confirmed` | Histogram | **no chainID on beholder**; promauto `chainID` | both |
| `txm_rpc_nonce` | Gauge | **no chainID on beholder**; promauto `chainID,address` | both |
| `txm_multicall_duration_ms` | Histogram | `chainID,method,blockTag,success,timedOut` | beholder-only |
| `ofa_send_tx_status` / `ofa_send_tx_latency` | Counter / Histogram | `chainID,backend,status` | beholder-only (dual-broadcast/MEV only) |
| `meta_endpoint_status_codes` / `meta_endpoint_latency` | Counter / Histogram | `chainID,feedAddress` | beholder-only (FastLane/Atlas only) |
| `meta_bids_per_transaction` / `meta_errors` | Histogram / Counter | `chainID,feedAddress,errorType` | beholder-only (FastLane/Atlas only) |

The `ofa_*` and `meta_*` families are only live on deployments that route sends through a
dual-broadcast/MEV backend (Flashbots/Nova/FastLane Atlas); on a plain NOP commit deployment they
will legitimately be absent. The commit-plugin's *own* `metrics` (the `ccip_commit_*` /
`ccip_reader_*` series `commit-plugin-health.md` and `uncommitted-message.md` query) are **not**
part of this doc -- this doc is the layer underneath, reached by their `REPORT:...-oncall`
handoffs.

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
