> **DRAFT — pre-implementation.** Every `ccip_exec_*` metric referenced below is a *proposal*
> from [`docs/metrics/execute-metrics.md`](../metrics/execute-metrics.md) unless explicitly marked
> `(existing)`. The queries below are the dry run that stress-tested those proposals; see
> [Gaps the dry run exposed](#gaps-the-dry-run-exposed) for what the exercise changed.

- [Runbook sketch: unexecuted message](#runbook-sketch-unexecuted-message)
   * [Trigger and inputs](#trigger-and-inputs)
   * [Reading conventions](#reading-conventions)
   * [Decision tree](#decision-tree)
      + [step0 — is the exec plugin alive at all?](#step0--is-the-exec-plugin-alive-at-all)
      + [step1 — does exec know this message exists?](#step1--does-exec-know-this-message-exists)
      + [step2 — is the message being read from the source chain?](#step2--is-the-message-being-read-from-the-source-chain)
      + [step3 — can the DON agree on this lane's data?](#step3--can-the-don-agree-on-this-lanes-data)
      + [step4 — is token data (attestation) the blocker?](#step4--is-token-data-attestation-the-blocker)
      + [step5 — is a nonce the blocker?](#step5--is-a-nonce-the-blocker)
      + [step6 — is a report being built but discarded?](#step6--is-a-report-being-built-but-discarded)
      + [step7 — is the report accepted but not landing onchain?](#step7--is-the-report-accepted-but-not-landing-onchain)
   * [Worked scenarios](#worked-scenarios)
   * [Gaps the dry run exposed](#gaps-the-dry-run-exposed)

# Runbook sketch: unexecuted message

The exec-side counterpart of the [uncommitted-message runbook](uncommitted-message.md): a root
covering the message has landed on the destination offramp (commit's job is done — that runbook
STOPs here and hands off), but the message has not been executed within the expected window.

The dry-run premise: assume the High-priority metrics from
[`execute-metrics.md`](../metrics/execute-metrics.md) exist where proposed. This document is the
human-readable layer only; the agent-followable YAML decision graph is written after the metrics
are implemented and the tree has survived a forced-failure test (same bar as the commit runbooks).

## Trigger and inputs

**Trigger**: "message in flight > N minutes: commit root covering seq X is finalized on the dest
offramp, but no execution."

**Inputs**: `sourceChain`, `destChain`, `seqNum` (the message's onramp sequence number), `msgID`
(hex, for log/explorer cross-referencing only — see [reading conventions](#reading-conventions)).

## Reading conventions

1. **Every metric query is lane-scoped, never message-scoped.** All proposed exec metrics key on
   `source_network_name` (plus fixed enums like `reason`) — deliberately *not* on `msgID` or
   `seqNum` (cardinality). The runbook isolates the message to its lane, then reasons about the
   lane; the specific message is re-attached at the end via logs, joined on
   `(plugin, round/ocrSeqNr, sourceChain, seqNum)` in the classic way. A step that fires on the
   lane tells you the *class* of blocker; the logs tell you *which* messages it holds.
2. **Outcome vocabulary** is the same as the commit runbook: `STOP` (resolved), `CONTINUE:<step>`
   (next check), `REPORT:<owner>` (root cause named, but the fix lives outside exec metrics —
   surface it, don't act).
3. **Metric names** are written as registered in code (no `_total`); whether your datasource shows
   a `_total` suffix depends on the exposition path — same caveat as the
   [commit runbook's naming note](uncommitted-message.md#a-note-on-counter-naming). Check one
   query against your datasource's browser before trusting the rest.
4. **Per-node vs DON-wide.** These are per-oracle-node metrics. Steps that diagnose
   *lane-wide* failure want `max by (...)`/`sum by (...)` across nodes; steps that diagnose
   *node-local* failure (a bad RPC on one oracle) want the per-instance breakdown. Each step says
   which it wants and why.

## Decision tree

The tree mirrors the exec state machine (`Initialized → GetCommitReports → GetMessages → Filter →
report → transmit`), so "which step did I stop at" is also "which phase is wedged". Steps are
ordered so that the cheapest and most-discriminating checks come first.

### step0 — is the exec plugin alive at all?

`(existing)` `ccip_exec_plugin_heartbeat{phase="observation"|"outcome"}` — per node, incremented
unconditionally every round.

```
sum by (instance) (rate(ccip_exec_plugin_heartbeat{phase="outcome"}[1m]))
```

- **0 for one node** → that node is wedged/crashed; the DON may still be healthy. `REPORT` with
  per-node detail (node-local incident).
- **0 DON-wide** → the whole exec DON is not cycling rounds. Every metric below may be *stale*
  rather than bad — same "gauge went flat vs gauge shows a bad value" trap as the commit runbook's
  step0. Follow up with the live-oracle-headcount query (count instances with a heartbeat in the
  last minute) and `REPORT:ccip-exec-oncall`.
- **Alive** → `CONTINUE:step1`.

Sanity companion: `ccip_exec_current_state{state}` — a DON alive-but-stuck shows heartbeats
advancing while the state gauge never leaves `get_messages` (or `filter`). If the state
distribution across nodes is frozen at one non-idle state, skip ahead to that state's step
(`get_commit_reports` → step1, `get_messages` → steps 2–6, `filter` → steps 5–6) — you've been
handed the answer early.

### step1 — does exec know this message exists?

Two `(existing)` gauges bracket the answer:

- `ccip_exec_max_sequence_number{source_network_name="$sourceChain"}` — the highest seq this node
  has observed in a commit report for the lane.
- `ccip_exec_pending_reports{source_network_name="$sourceChain"}` — commit reports currently
  carried with unexecuted work, and
  `ccip_exec_oldest_pending_commit_age_seconds{source_network_name="$sourceChain"}` — how long the
  oldest has been waiting.

| Reading | Meaning | Outcome |
|---|---|---|
| `max_seq_num ≥ $seqNum` **and** `pending_reports > 0` | The lane has work pending and exec sees it. The message is plausibly inside a pending report. | `CONTINUE:step2` |
| `max_seq_num < $seqNum` | Exec has never observed a commit report covering this seq. Either the root hasn't landed (commit's problem) or it landed and exec's commit-report cache hasn't picked it up. Check the commit side first — if the root *is* landed on the offramp (the trigger condition says it is), this is the exec commit-report cache: **`ccip_exec_commit_report_cache_last_refresh_age_seconds`** (proposed, Medium — see [gaps](#gaps-the-dry-run-exposed)) and its refresh-error counter. High age + no refresh errors = cache looking in the wrong window; high age + refresh errors = reader outage. | `REPORT:ccip-exec-oncall` (cache) or hand back to the [uncommitted-message runbook](uncommitted-message.md) (root never landed) |
| `pending_reports == 0` **and** no `oldest_pending_commit_age` series for the lane | Exec believes there is nothing to execute on this lane. Given the trigger, the most likely truth is the message (or the range it lives in) is already executed onchain. Confirm on the explorer against the offramp's executed set — this runbook has no query for that. | `REPORT:verify-onchain` — if executed, the alert was a false positive; if not, exec and commit disagree about pending work (grab logs) |

The age gauge's *interpretation rules* (from the metric design) apply from here on: growing age +
advancing heartbeat + state cycling = the plugin is alive and seeing work it can't finish; flat or
absent age = idle; **frozen** age with a frozen heartbeat = wedged (you should have stopped at
step0). One benign growth pattern to keep in mind: reports executed-but-unfinalized (snoozed
roots) also hold the age up — cross-check the root-cache snooze signals before paging on age
alone.

### step2 — is the message being read from the source chain?

`ccip_exec_message_read_errors_total{source_network_name, reason="rpc"|"range_mismatch"}` —
per-node counter.

```
sum by (instance, reason) (rate(ccip_exec_message_read_errors_total{source_network_name="$sourceChain"}[15m]))
```

- **Firing on all nodes, `reason="range_mismatch"`** — the log poller returns fewer messages than
  the commit root claims for the range: the classic unindexed-event / false-idle signature.
  Lane-wide, so *every* oracle starves the lane identically and it looks like quiet from the
  outside. `REPORT:source-chain-infra` (log poller / RPC).
- **Firing on all nodes, `reason="rpc"`** — source-chain RPC degradation DON-wide.
  `REPORT:source-chain-infra`.
- **Firing on one node only** — node-local RPC problem. This oracle silently drops the lane's
  reports from *its* observations; if the DON is otherwise healthy the lane survives, but this
  node is a liveness liability. `REPORT:node-ops` with the instance.
- **Clean** → `CONTINUE:step3`.

Also check the silent exclusion paths here, since they're observation-stage too:
`ccip_exec_messages_skipped_total{source_network_name, reason="already_inflight"}` — a sustained
rate with no execution progress is the orphaned-inflight-entry signature (the inflight cache never
deletes in production; a crash between mark and transmit silently benches messages for a full
TTL). `reason="already_executed"` here means the *commit data* claims execution — if the trigger
says it isn't, exec's and the chain's views disagree. Either way these are
`STOP`-quality findings once confirmed in logs for `$msgID`.

### step3 — can the DON agree on this lane's data?

- `ccip_exec_consensus_dropped_total{objectName="fChain", key="$sourceChain", reason}` — if this
  fires, the lane fell out of consensus fChain and **everything** downstream (messages, roots,
  hashes, token data) silently drops for it.
  - `reason="insufficient_agreement"` → not enough oracles are observing the lane (peer with
    step2: usually an upstream read failure on a subset of nodes).
  - `reason="split"` → oracles observe *different* fChain values — integrity-adjacent, treat
    seriously.
  - `reason="threshold_not_defined"` → config bug, not an oracle health issue.
  → `STOP` with the reason; this is the execute twin of commit's fChain incident.
- `ccip_exec_message_consensus_conflicts_total{source_network_name, kind}` — two different
  messages/hashes reached consensus for one seqNum; the message is dropped from execution
  entirely, and with ordered execution **everything behind it is blocked too**. Persistent rate on
  a lane where `$seqNum` sits explains "this message never executes" in one query.
  → `STOP`, `REPORT:ccip-exec-oncall` (needs log inspection of the conflicting observations).
- **Clean** → `CONTINUE:step4`.

### step4 — is token data (attestation) the blocker?

Gate check: `ccip_exec_messages_skipped_total{source_network_name, reason="token_data_not_ready"}`.

- **Flat / zero** → not a token-data problem. `CONTINUE:step5`.
- **Rising** → token data is the blocker for this lane; now identify *which layer*, using the
  tokendata client metrics (all High proposals):

  | Layer check | Reading | Meaning |
  |---|---|---|
  | `ccip_exec_tokendata_observer_error_total{observer, stage}` rate > 0 | any | The composite observer swallowed a client failure — messages degrade to not-ready and no plugin-level error ever fires. `observer` tells you which adapter (usdc_v1/lbtc/cctpv2). `REPORT:attestation-infra` |
  | `ccip_exec_tokendata_http_cooldown_active{api} == 1` | true | In a rate-limit cooldown — every fetch fails until it lifts. Expected to self-resolve; page only if it never lifts (check `http_request_total{outcome="rate_limited"}` rate). |
  | `ccip_exec_tokendata_http_request_total{outcome="not_ready_404"}` dominant | — | **Healthy pending**: the attestation simply hasn't been produced upstream yet. This is "waiting", not "broken". Check backlog shape (`ccip_exec_tokendata_attestations_waiting{token}`, background `queue_depth`) and let it ride unless the backlog grows unbounded. |
  | `ccip_exec_tokendata_http_request_total{outcome="http_error"\|"timeout"}` dominant | — | Attestation API outage. `REPORT:attestation-infra`. |
  | `outcome="data_missing"` (usdc/lbtc fetch counters) | — | The deposit/never-attested signature: the upstream system never saw the deposit hash. False-idle-shaped upstream failure. `REPORT:token-vendor`. |

  → In every row, `STOP` once the layer is named — the runbook's job (locate the stage) is done;
  the fix belongs to the attestation pipeline.

### step5 — is a nonce the blocker?

- **`ccip_exec_nonce_read_errors_total` rate > 0** → the root cause: one destination-chain nonce
  read failure empties `Observation.Nonces`, and *every* ordered message on the lane then skips
  with `nonces_missing_for_chain` next state — a total-execution-stop that masquerades as N
  per-message errors. `STOP`, `REPORT:dest-chain-rpc`.
- **`ccip_exec_messages_skipped_total{reason="invalid_nonce"\|"missing_nonce"}` rising** → a nonce
  gap: the message's seqNum is executable but an *earlier* seq on the lane hasn't executed, and
  ordered execution refuses to skip ahead. The real patient is the first unexecuted seq on the
  lane, not `$seqNum` — loop back to the earliest unexecuted seq (the runbook's
  `$seqNum` becomes that value) and re-enter at step2. If the gap-message is *itself* stuck on
  token data, this branch and step4 are the same incident wearing two hats; step4's evidence wins.
- **Clean** → `CONTINUE:step6`.

### step6 — is a report being built but discarded?

The report-builder funnel. Any fire here means messages passed all per-message checks and died at
report granularity:

- `ccip_exec_report_build_errors_total{reason}` — token-data length mismatch, merkle
  reconstruction failures, malformed commit data. Persistent `reason="empty_report"` is the
  strongest "builder is wedged" signal (checks passed, nothing built, and today that path is
  fully silent). `STOP` with the reason.
- `ccip_exec_report_validation_rejections_total{reason="size_limit"\|"gas_limit"}` + the
  `ccip_exec_report_size_bytes`/`ccip_exec_report_estimated_gas` gauges — the report is *too
  big for the chain*, i.e. a throughput ceiling, not a fault. Suspect when a lane's backlog is
  large and age grows while every earlier step is clean. `REPORT:config-tuning` (limits) — and
  note sustained firing means the lane executes at less than commit throughput forever.
- `ccip_exec_messages_skipped_total{reason="deferred"}` — ready messages evicted from the
  in-progress report by the same limits; the per-message view of the above.

### step7 — is the report accepted but not landing onchain?

The report made it through OCR — now the acceptance/transmission boundary.

`ccip_exec_report_validation_rejected_total{phase="should_accept"|"should_transmit", reason}`:

| reason | meaning | outcome |
|---|---|---|
| `config_digest_mismatch` | OffRamp config drifted mid-flight | `REPORT:config-drift` — expect `ccip_exec_config_digest_mismatch` `(existing)` to corroborate |
| `dest_not_supported` | transmission-schedule misconfiguration (this oracle isn't a dest reader) | `REPORT:config-drift` |
| `already_executed` | onchain state says executed — the "all-or-nothing" rejection when a range is partially executed | `STOP` if the trigger was a false positive; otherwise `REPORT:ccip-exec-oncall` (out-of-order-transmission residue) |
| `cursed` | RMN curse | cross-check `ccip_exec_cursed{...}` — hand to the curse playbook |
| `validation_error` | transient DB/RPC in the acceptance path | watch; page if sustained |

- **Rejections at `should_accept` firing DON-wide** → no report ever gets transmitted; zero
  onchain trace. Distinguish carefully from the next row — different remediation entirely.
- **`phase="should_transmit"` clean but the message still isn't executed** → the report was
  accepted and handed to transmission; from here it's the txm and the chain, which exec metrics
  cannot see (the same epistemic boundary the commit runbook hits at its step5). Confirm on the
  explorer whether the execute tx was even *submitted*, then either `REPORT:txm` (never
  submitted) or `REPORT:verify-onchain` (submitted, reverted/stuck).
  `REPORT:ccip-exec-oncall` in both cases — this is the boundary of the runbook.

## Worked scenarios

The dry run: four realistic incidents walked through the tree, checking each terminates at a
*distinct* outcome (a tree where everything ends at the same step is a tree that can't triage).

| # | Incident (forced failure mode) | Path | Terminates at | Discriminated from the others by |
|---|---|---|---|---|
| 1 | Source-chain RPC outage: log poller can't read messages in range | step0 ✅ → step1 pending>0, age growing → step2 `read_errors{reason="range_mismatch"}` on all nodes | `REPORT:source-chain-infra` | step2's per-node split + reason; commit-side metrics stay clean |
| 2 | Circle rate-limits the DON: 5-min cooldown, attestations stall | step0 ✅ → step1 ✅ → step2 clean → step3 clean → step4 `token_data_not_ready` rising + `cooldown_active=1` | `REPORT:attestation-infra` (self-resolving) | step4's layer table — cooldown vs 404-pending vs observer-error are three different outcomes on the *same* skip reason |
| 3 | One oracle's dest-chain nonce read fails persistently | step0 ✅ → …→ step5 `nonce_read_errors > 0` on one node | `STOP`+`REPORT:dest-chain-rpc` | The upstream counter fires even though the *symptom* is `nonces_missing_for_chain` on every message — root cause one query instead of log-chain across two packages |
| 4 | Config digest drift mid-rollout: reports rejected at acceptance | step0 ✅ → …→ step7 `rejected{phase="should_accept", reason="config_digest_mismatch"}` | `REPORT:config-drift` | Phase label (`should_accept` vs `should_transmit`) separates "never left the DON" from "submitted but stuck" — zero onchain trace vs some |

All four walk *through* steps 1–3 without firing, which is the point: those steps exist to
bracket (alive? seen? agreed?) and their silence is load-bearing. The dry run's honest result: the
tree discriminates the four canonical failure classes cleanly, but two of the branches lean on
metrics that are currently only **proposed Medium** — see below.

## Gaps the dry run exposed

Things the walkthrough hit that the High list alone couldn't answer — feedback for
[`execute-metrics.md`](../metrics/execute-metrics.md):

1. **The commit-report cache refresh age is load-bearing on the main path, not a Medium nice-to-have.**
   Step1's "root landed but exec never saw it" branch needs
   `ccip_exec_commit_report_cache_last_refresh_age_seconds` + its refresh-error counter to avoid
   dead-ending at `REPORT` with "suspect the cache, check logs". ~~Promote both to High.~~
   **Done — promoted to High in [`execute-metrics.md`](../metrics/execute-metrics.md).**
2. **No "last executed seq num" gauge exists.** Step1's `pending_reports == 0` branch (and
   step7's `already_executed` row) both want "what does exec believe is executed on this lane?" as
   a metric. Proposal: `ccip_exec_last_executed_seq_num{source_network_name}` — the exact analogue
   of commit's `offramp_next_seq_num` gauges, derivable from `ExecutedMessages` already flowing
   through commit data. This is the single cheapest way to close the "already executed vs never
   committed" ambiguity that currently forces an explorer check. **Done — `ccip_exec_last_executed_seq_num` added to the High table in [`execute-metrics.md`](../metrics/execute-metrics.md).**
3. **`token_data_not_ready` doesn't say which adapter.** The lane-level skip counter and the
   tokendata client counters reconcile only by inference (which adapters are configured, which of
   their counters move). Acceptable — but if a fourth adapter lands, consider a `token`/`observer`
   label on the skip counter. Kept out of the proposal for now to avoid premature label axes.
4. **The "first unexecuted seq" hunt in step5 is log-only by design.** A per-lane
   "lowest unexecuted seqNum" gauge would collapse the loop-back, but it's derivable from
   `ccip_exec_last_executed_seq_num` (gap #2) without new cardinality risk — covered once that
   exists.
