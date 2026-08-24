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
| [`chain-identifiers.md`](chain-identifiers.md) | reference, not a runbook — translating between chain ID / chain selector / chain name |
| [`evm.md`](evm.md) | chain-family deep dive — your commit runbook hit a `data_source` or `report_transmission` failure on an **EVM** destination chain and handed off to a txm/infra owner; this is what to look at in the chainlink-evm layer (RPC pool, log poller, gas, TXM) |
| [`solana.md`](solana.md) | chain-family deep dive — same, but for a **Solana** destination chain (log poller, block-history fee estimator, Txm) |

The last two are not standalone: they're entered from `commit-plugin-health.md` / `uncommitted-message.md`
when those hand off with `REPORT:...-txm-oncall` / `REPORT:...-infra-oncall` (chain-family-specific
versions of the commit plugin's own covered paths), and they use the **chain-family** repos'
metrics (chainlink-evm / chainlink-solana), not the `ccip_commit_*` / `ccip_reader_*` series the
other docs query.

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

## Identifying a chain

Every `destChain`/`sourceChain`/`sourceChains` input in these docs can be given to you as a chain
ID (`1`), a CCIP chain selector (`5009297550715157269`), or a chain name (`ethereum-mainnet`) —
and the PromQL label the query actually needs is not always the same one you were handed. See
[`chain-identifiers.md`](chain-identifiers.md) for what each of these means, where they're
defined, and how to translate between them before substituting into a query.

## Identifying a CCIP message

A message can be handed to you three ways: a message ID (32 bytes, `0x` + 64 hex chars), a
`(sourceChain, destChain, sequence number)` triple (sequence number is a strictly-positive
uint64), or a transaction hash plus the chain it was sent on. `uncommitted-message.md` wants the
triple plus the message ID (its `inputs`); if you were only given one of the other forms, resolve
the rest before starting the decision graph rather than guessing.

[`ccip-cli`](https://www.npmjs.com/package/@chainlink/ccip-cli) (from
[ccip-tools-ts](https://github.com/smartcontractkit/ccip-tools-ts); `npm install -g
@chainlink/ccip-cli` or run via `npx @chainlink/ccip-cli`) does this resolution for you. Its
`show` command takes either a transaction hash or a message ID and figures out the rest by
reading the chain(s) directly:

```bash
ccip-cli show <tx-hash-or-message-id> \
  --rpc <source-chain-rpc-url> \
  --rpc <dest-chain-rpc-url> \
  --format json
```

- You need at least the source chain's RPC (to find the send event); add the dest chain's RPC too
  if you want execution/commit status rather than just the send.
- Pass a transaction hash when that's what you have; pass the message ID directly when you don't
  have a tx hash (e.g. it came from a log or a report) — `show` accepts both as the same
  positional argument.
- `--format json` gives you a stable, parseable shape (source/dest chain, sequence number,
  message ID, status) instead of the pretty-printed table — use it when you're going to feed the
  result into the next step rather than read it yourself.
- If a transaction contains more than one CCIP message, `show` will ask you to disambiguate; pass
  `--log-index <n>` once you know which one you want.
- `show` reports chain name and chain selector, not necessarily the chain ID your PromQL queries
  need — see [Identifying a chain](#identifying-a-chain) above to translate.

## Providing credentials (e.g. a Grafana API key) to an agent

Don't paste a Grafana (or any other datasource) API key into chat — anything you type is part of
the conversation the agent sees and may retain. Instead, put the token in a local file the agent
can read but you never speak aloud:

```bash
mkdir -p ~/.config/ccip-runbooks
echo -n "<token>" > ~/.config/ccip-runbooks/grafana.token
chmod 600 ~/.config/ccip-runbooks/grafana.token
```

Then tell the agent the *path*, not the value, e.g. "query Grafana at `https://<org>.grafana.net`
using the token at `~/.config/ccip-runbooks/grafana.token`." Instruct it to:

- read the file only to build the request itself, e.g.
  `curl -H "Authorization: Bearer $(cat ~/.config/ccip-runbooks/grafana.token)" ...`;
- never `cat`/print/echo the file's contents on their own, and never include the token in any
  intermediate or final output (including the doc's own output contract);
- treat the file as write-only from the agent's perspective — if a query fails with an auth
  error, report that plainly rather than trying to dump the token to debug it.

Keep the file outside any git-tracked directory (`~/.config/...` is fine; a repo-local `.env` is
not, even if `.gitignore`d — it's too easy to `git add -f` by accident).

## Invoking an AI agent with a runbook

Point the agent at the file and the query endpoint, give it the specific inputs, and ask for the
doc's own output contract — don't ask it to "check if things look okay," the doc already defines
what "okay" means and how to report it.

**Template:**

```
Read <path to runbook> fully, then execute it against <query endpoint, e.g.
http://localhost:8428/api/v1/query for local devenv VictoriaMetrics>.

Inputs:
- destChain = "<chain ID, chain selector, or name -- translate per chain-identifiers.md as needed>"
- sourceChain = "<same>"        # uncommitted-message.md only
- seqNum = <value>              # uncommitted-message.md only
- msgID = "<value>"             # uncommitted-message.md only; if you don't have this, or only
                                 # have a tx hash, resolve it first with ccip-cli (see
                                 # README.md#identifying-a-ccip-message) before starting the doc
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
- **Label values aren't uniform either** — some labels want a chain name, others a chain ID, and
  neither is the chain selector you may have been handed. See
  [`chain-identifiers.md`](chain-identifiers.md); an agent that skips translating and passes
  whatever value it was given straight into every query will get silent empty results here too.

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
