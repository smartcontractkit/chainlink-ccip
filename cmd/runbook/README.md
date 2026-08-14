# runbook

Deterministic executor for the machine-checkable YAML specs behind
[`docs/runbooks`](../../docs/runbooks), against any Prometheus-compatible
datasource (defaults to the local devenv VictoriaMetrics at
`http://localhost:8428`).

The point of the tool is to **remove the mechanical walk from the AI-agent
loop**. Reading the doc, copy-pasting ~25 queries, applying the empty-result
rule, following the branches, and formatting the output contract is the
high-token, near-100% deterministic part. This binary does all of that —
an agent (or human) only has to spend tokens on the judgment the docs leave
open. Every executed query ships its raw results (`--raw`) so either layer can
still verify a verdict against ground truth.

## Build / test

```sh
make build
make checks
```

## Usage

```sh
runbook list
runbook run commit-plugin-health -D destChain=chain-b \
  -D sourceChains='.*' -D fRoleDON=2 \
  --endpoint http://localhost:8428 --raw
runbook run uncommitted-message \
  -D destChain=chain-b -D sourceChain=chain-a -D seqNum=42 \
  --endpoint http://localhost:8428
```

Runbook-specific inputs are passed with `-D name=value` (repeatable); which
ones a runbook needs is in its `inputs:` block (see `runbook list`).

## CLI flags (run)

| flag | default | purpose |
|---|---|---|
| `--endpoint` | `http://localhost:8428` | datasource base URL |
| `--bearer` / `--user` / `--pass` | | datasource auth |
| `--suffix auto\|keep\|strip` | `auto` | metric-name suffix policy (see below) |
| `--timeout` | `20s` | per-query timeout |
| `--raw` | off | also emit raw query results for verification |
| `-D input=value` | | runbook inputs |

## How the two runbook shapes map

- **checklist** (`commit-plugin-health.yaml`): run all independent checks, apply
  the empty-result rule per `always_emitted`, evaluate each check's severity
  rules, fold findings per the documented aggregation rules, and emit the
  `health_report` output contract.
- **graph** (`uncommitted-message.yaml`): walk the decision tree from `root`,
  executing each step's queries and following `STOP` / `CONTINUE:<id>` /
  `REPORT:<owner>` outcomes, emitting a trace. Steps with no query
  (`not_automatable`, e.g. chain-B TXM/explorer) or a deliberately fuzzy
  condition (`judgment`, e.g. "climbing vs flat") are surfaced as **AGENT**
  handoffs with their raw results — never guessed.

## The empty-result rule (the part that's easy to get wrong)

An empty PromQL result is not "all zero". The docs encode this as
`always_emitted: true` (a metric recorded every round — empty means the
pipeline is suspect → `UNKNOWN`) vs `always_emitted: false` (an event counter
that only gains a series after firing — empty means "never happened" → `OK`).
The engine implements this exactly; a bare `result` comparison over an
event-counter's empty result resolves to value `0`, and `any`/`all` quantifiers
are vacuous over an empty set.

## The `_total` / `_bucket` suffix adapter

Doc-form queries assume an OTel→`prometheusremotewrite` pipeline (counters gain
`_total`, histograms `_bucket`). A datasource scraped directly from
`promauto` base names has none. `--suffix auto` probes once (detecting which
spelling the datasource exposes), `keep` uses the queries as written, `strip`
removes the suffixes. This is a **one-time per-datasource** concern, not a
per-step one.

## Adding a runbook

Add a `.yaml` under `internal/runbooks/` matching the `type: checklist`
(checks + severity rules) or `type: graph` (root + steps) schema in
`internal/schema.go`. `runbook list` validates it on load, so a malformed
runbook fails fast rather than half-running. Keep it a faithful transcription
of the markdown doc — the shared step IDs are what stop the two from drifting.
