# Runbooks

Operational docs for triaging and health-checking CCIP plugins, written to be followed
mechanically by a human **or** an AI agent — not read once and internalized. Each doc pairs a
structured, machine-checkable spec (a fenced YAML block: a decision graph or a checklist) with
prose explaining the *why*, sharing the same step/check IDs so the two can't silently drift
apart.

| doc | use when |
|---|---|
| [`uncommitted-message.md`](uncommitted-message.md) | reactive — you have a specific stuck message (source chain, dest chain, onramp seq num) and need to find where it's stuck |
| [`commit-plugin-health.md`](commit-plugin-health.md) | proactive — no specific incident; is this commit plugin instance healthy right now |

Both currently cover the **commit plugin / merkle root processor** only, built from the metric
coverage work in [`docs/metrics/commit-metrics.md`](../metrics/commit-metrics.md) and rendered
visually in [`devenv/dashboards/commit-plugin.json`](../../devenv/dashboards/commit-plugin.json)
(that dashboard's Health section mirrors `commit-plugin-health.md`; its Guided Debug section
mirrors `uncommitted-message.md`).

## As a human

1. Read the frontmatter first (`trigger`, `inputs`, `related`) — it tells you whether this is the
   right doc before you read anything else.
2. Skip the YAML block unless you want the exact query for a step; the prose under
   `## Steps`/`## Groups` explains the reasoning in the same order.
3. You need a Prometheus-compatible query endpoint with these metrics in it. Locally, that's the
   VictoriaMetrics stack started by `ccip obs up --victoria` (see `devenv/README.md`) — query it
   via Grafana Explore (`http://localhost:3000`, datasource `VictoriaMetrics`) or directly at
   `http://localhost:8428/api/v1/query`. Outside devenv, use whatever datasource your
   environment's Beholder pipeline actually feeds — see each doc's note on counter `_total`
   suffixing, which depends on that pipeline, not on the runbook.
4. `uncommitted-message.md`'s Guided Debug panels and `commit-plugin-health.md`'s Health panels
   in `devenv/dashboards/commit-plugin.json` run the same queries pre-wired with dashboard
   variables — often faster than copy-pasting PromQL by hand.
5. Every terminal outcome in these docs is `STOP` or `REPORT:<owner>` — never an instruction to
   page or act. If a doc leads you to a `REPORT`, that's a handoff for you to make, not something
   the doc did for you.

## Invoking an AI agent with a runbook

Point the agent at the file and the query endpoint, give it the specific inputs, and ask for the
doc's own output contract — don't ask it to "check if things look okay," the doc already defines
what "okay" means and how to report it.

**Template:**

```
Read <path to runbook> fully, then execute it against <query endpoint, e.g.
http://localhost:8428/api/v1/query for local devenv VictoriaMetrics>.

Inputs:
- destChain = "<value>"
- sourceChain = "<value>"       # uncommitted-message.md only
- seqNum = <value>              # uncommitted-message.md only
- msgID = "<value>"             # uncommitted-message.md only
- sourceChains = "<regex>"      # commit-plugin-health.md only, default ".*"

Follow the decision graph / checklist exactly as written, substituting the inputs above into
each query. Produce the doc's defined output as your final answer, fully filled in with real
query results -- not a description of what you'd do.

Do not page, escalate, or take any remediation action. Every terminal outcome in this doc is
`STOP` or `REPORT:<owner>` -- your job ends at reporting the finding and the suggested owner.
```

Two things worth being explicit about when you invoke one of these, because they're easy to get
wrong even for an agent reading the doc carefully:

- **The empty-result rule matters more than it looks like it should.** Both docs distinguish "no
  data because the pipeline is broken" from "no data because this counter has never fired" —
  getting this backwards silently turns a healthy system's report into a false alarm (or the
  reverse). If you're spot-checking an agent's output, this is the first thing to verify against
  the raw query results.
- **Label keys are not uniform across metrics** (`chainID` vs `chain_id`, see each doc's cheat
  sheet). An agent that "normalizes" these to one spelling will get silent empty results, not an
  error.

### Re-testing a runbook after editing it

If you change a decision graph, checklist, or severity rule, re-run it through an agent against
a real datasource before trusting the edit — a doc that reads correctly is not the same as a doc
that resolves correctly, and the fastest way either doc in this directory reached its current
state was by having an agent literally execute it and report back where it guessed. Ask the
agent to also report friction: places it had to guess, or where a real query result didn't match
what the doc told it to expect. That feedback is usually more valuable than the health/triage
result itself while a doc is still being iterated on.

## Adding a new runbook

Keep new docs consistent with the two here:

- Frontmatter: `name`, `description`, `trigger`, `severity`, `owner`, typed `inputs`, `related`,
  `status`.
- A `## For agents` section defining the outcome vocabulary you're using and any non-obvious
  interpretation rules (empty results, label inconsistencies, anything an agent would otherwise
  have to guess).
- A fenced YAML block that's the actual control structure (decision graph for
  reactive/triage docs, checklist for proactive/health docs) — prose explains it, doesn't replace
  it.
- A defined output contract, so "did the agent get the right answer" is checkable against a
  schema instead of read as prose.
- `STOP` / `CONTINUE:<id>` / `REPORT:<owner>` as the only outcomes. No outcome that implies the
  agent itself pages, escalates, or takes a remediation action.
