---
name: chain-identifiers
description: Reference for translating between chain ID, chain selector, and chain name -- not a runbook itself, linked from every runbook input that takes a chain.
status: living
---

# Chain identifiers: chain ID, chain selector, chain name

Three different ways to name the same chain show up across this repo's runbooks, dashboards, and
tools. They are not interchangeable, none is optional to recognize, and no single one is "the"
identifier -- pick whichever the thing in front of you (a PromQL label, an on-chain event, a
human's incident report) happens to use.

| identifier | shape | example (Ethereum mainnet) | defined by |
|---|---|---|---|
| chain ID | uint, the chain's native/EVM chain ID (or family-specific equivalent) | `1` | the chain itself; chain-selectors mirrors it |
| chain selector | uint64, a CCIP-internal identifier, numerically unrelated to chain ID | `5009297550715157269` | [chain-selectors](https://github.com/smartcontractkit/chain-selectors) |
| chain name | human-readable slug | `ethereum-mainnet` | chain-selectors |

## Canonical source of truth

The YAML files in [chain-selectors](https://github.com/smartcontractkit/chain-selectors) (checked
out locally at `../chain-selectors` relative to this repo, if present -- otherwise clone it or use
its Go module). `selectors.yml` is keyed by chain ID, with each entry carrying a `selector:` and a
`name:`; family-specific chains that don't fit the EVM key space live in siblings
(`selectors_solana.yml`, `selectors_aptos.yml`, etc). `all_selectors.yml` is the merged view across
every family -- prefer it when you don't know or don't want to special-case the family.

## How to translate

Given any one of the three, find the other two by grepping the YAML -- these files are small,
flat, and stable, you don't need to write code or install anything:

- **Have chain ID** → it's the YAML key itself; look up that block directly.
- **Have chain selector** → grep for the value under `selector:` and read the enclosing block's
  key (chain ID) and `name:`.
- **Have chain name** → grep for `name: "<value>"` and read the enclosing block's key (chain ID)
  and `selector:`.

```bash
grep -B3 'name: "ethereum-mainnet"' path/to/chain-selectors/all_selectors.yml
```

If [`ccip-cli`](https://www.npmjs.com/package/@chainlink/ccip-cli) is available, its internal
`networkInfo()` resolver (used by `show`/`lane`) accepts any of the three and returns the other
two -- see [README.md's message-identification section](README.md#identifying-a-ccip-message).
For a bare "what does selector X map to" lookup with nothing else going on, reading the YAML
directly is more direct and doesn't require Node/npm.

## Why this matters for these runbooks specifically

PromQL label *values* in this repo are not uniformly one representation. The commit-plugin
metrics use chain ID for `chainID`/`chain_id` labels and chain name for
`source_network_name`/`dest_network_name` -- see each runbook's own label cheat sheet
(e.g. [`commit-plugin-health.md`](commit-plugin-health.md#label-key-cheat-sheet)). An incident
handed to you as a chain selector (common from an on-chain event or `ccip-cli show`'s output) or a
chain ID (common from a human's bug report) will usually need translating before it can go into a
query at all -- the query's label expects a specific one of the three, not whichever one you
happen to have.

Do the translation once, up front, and carry all three forms with you for the rest of the
investigation -- cheaper than re-deriving one from the others every time a different query needs a
different label.
