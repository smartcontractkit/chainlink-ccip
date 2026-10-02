---
name: tooling-api
description: Use when working on the CCIP deployment tooling API in chainlink-ccip/deployment — implementing its adapters (Deployer, LaneAdapter, TokenAdapter, MCMSReader, ChainFamily, DeployChainContractsAdapter, CommitteeVerifierContractAdapter, …) for a new or existing chain family, adding a new adapter interface or registry, or adding a new chain-agnostic changeset. Covers both the v1 (1.6) and v2 (2.0) APIs.
---

# CCIP Tooling API: extend it or implement it for a chain family

The tooling API is the chain-agnostic changesets plus the adapter interfaces in `deployment/`. Chain families implement the interfaces and register them from `init()`.

## 0. Orient first (always)

1. Read `deployment/docs/README.md` and `deployment/docs/architecture.md`. They are short.
2. Decide which API generation the task targets. It is either **v1 (1.6)** or **v2 (2.0)**; ignore 1.x minor versions unless the task is specifically about them.

   | Task touches… | Generation | Guide |
   |---|---|---|
   | `deploy/`, `lanes/`, `fees/` FeeAdapter, `ConnectChains`, `DeployContracts`, OCR3 | v1 | `deployment/docs/implementing-1.6.md` |
   | `v2_0_0/adapters`, `v2_0_0/changesets`, topology, CCVs, executors | v2 | `deployment/docs/implementing-2.0.md` |
   | `tokens/`, MCMS, ownership, curse, fee aggregator, authorized callers, hooks, test adapters | shared | `deployment/docs/interfaces.md#shared-by-v1-and-v2` |

3. Open `deployment/docs/interfaces.md` and find each interface involved: its registry, its key (`family` or `family-version`), and whether it is discovered by type assertion.
4. Before writing anything, read the reference implementation for that interface:

   | | v1 | v2 |
   |---|---|---|
   | EVM | `chains/evm/deployment/v1_0_0/adapters`, `v1_6_0/{sequences,adapters,operations}` | `chains/evm/deployment/v2_0_0/{adapters,sequences,operations}` |
   | Solana | `chains/solana/deployment/v1_0_0/adapters`, `v1_6_0/{sequences,adapters,operations}` | `../chainlink-ccip-solana/deployment/v2_0_0/{adapters,sequences,operations}`, `solcommon/` |

   Find a family's registrations with `grep -rn 'func init()' <module>`, then read the `init.go` files.

Then follow the workflow below that matches the task.

---

## Workflow A: Implement the API for a chain family

Use this for a new family, or to fill gaps in an existing one.

1. **Inventory.** List the interfaces needed from the checklist in the matching guide: `implementing-1.6.md` "Checklist", or `implementing-2.0.md` "Checklist" plus "What 2.0 reuses from 1.6". Check which ones the family already registers (`grep -rn 'Register' <module> | grep -v _test`).
2. **Layout.** Mirror the reference layout:
   - `utils/`
   - `v1_0_0/adapters` (version-agnostic, keyed by family)
   - `v1_6_0/{operations,sequences,adapters}` or `v2_0_0/{operations,sequences,adapters,changesets}`
3. **Operations, bottom up.** Write one side effect per operation, with ID `"<contract>:<action>"`.
   - Writes must branch on ownership: if the deployer key owns the contract, execute; otherwise return an MCMS `BatchOperation` in `OnChainOutput.BatchOps`.
   - Deploys must be idempotent: check `ExistingAddresses` first.
   - EVM uses the CLDF `contract.NewDeploy`, `NewWrite`, and `NewRead`. Non-EVM uses `operations.NewOperation`.
4. **Sequences.** Each sequence works on one chain, takes serializable input, and returns `sequences.OnChainOutput`. Compose sub-sequences with `sequences.RunAndMergeSequence`. Configure the Router **last** (v2: in its own BatchOperation). Read on-chain state, and only emit writes that change something.
5. **Adapter struct(s).** Methods return package-level sequence vars. Add `var _ <pkg>.<Interface> = (*MyAdapter)(nil)` for **every** interface, including optional ones discovered by type assertion. A mismatch then fails the build instead of being silently skipped.
6. **Register in `init()`.** Use the exact accessor and key from `interfaces.md`. Templates are in [references/templates.md](references/templates.md).
   - v1: `family-version` with `utils.Version_1_6_0`, plus version-agnostic helpers keyed by family in `v1_0_0`.
   - v2: register by family; shared ones (tokens, fees, ownership) use `utils.Version_2_0_0`.
   - Register the 4-byte family selector: implement `GetChainFamilySelector()` on the `LaneAdapter`, or call `utils.RegisterChainFamilySelector` if there is no `LaneAdapter`.
7. **Address encoding.**
   - Datastore: the native string.
   - `AddressRefToBytes` / `Get*Address`: native bytes.
   - v2 `GetOnRampAddress`: the bytes **as written into messages**. EVM abi-encodes to 32 bytes.
   - Remote addresses in inputs are already in the remote family's encoding. Never re-encode them.
8. **Tests.** Cover each sequence on a simulated chain. Add a changeset-level test that blank-imports the adapter package and calls the shared changeset's `.Apply(...)` (see "Consuming" below). Implement a `testadapters.TestAdapter` if message tests should cover the family.
9. **Docs.** Update the family's docs (if any), and `deployment/docs/interfaces.md` or the guides if you found a gap or a stale statement.

## Workflow B: Add a new adapter interface or capability to the shared API

1. **Pick the narrowest mechanism** (see the style guide, "prefer the narrowest clear abstraction"):
   - *New optional capability for an existing adapter* → define a small interface next to the existing one, and use `adapter.(NewIface)` in the changeset. No registry is needed. Examples: `tokens.TokenFeeAdapter`, `lanes.TestRouterProvider`, `adapters.GasPriceValidator`.
   - *New concept with its own lifecycle* → new interface + new registry (singleton accessor, first registration wins, keyed by `family` or `family-version`). Example: `authorizedcallers/product.go`.
2. **Where it goes.** v1/shared goes in the relevant package's `product.go`. v2-only goes in `deployment/v2_0_0/adapters/<feature>.go`.
3. **Shape.** Methods that change chain state return `*cldf_ops.Sequence[In, sequences.OnChainOutput, cldf_chain.BlockChains]`. Read helpers take explicit deps (`Bundle`, `BlockChains`, `DataStore`) rather than a whole `Environment` where possible. Put the input structs next to the interface, with doc comments, and `json`/`yaml` camelCase tags on user-facing fields.
4. **Doc-comment every method,** including encoding expectations (raw bytes vs. string; which family's encoding) and idempotency requirements. Implementers rely on these.
5. **Implement it for EVM** (and Solana v1 if it applies) in the same PR, or explicitly state which families support it. Register in the existing `init.go`.
6. **Regenerate mocks** if the interface is listed in `deployment/.mockery.yaml` (`v2_0_0/mocks` covers v2 adapters): run `mockery` from `deployment/`.
7. **Update `deployment/docs/interfaces.md`.** Add a row to the right section table (registry, key, source), plus a method table or a one-line summary. If consumers need to know, also update `consuming.md`.

## Workflow C: Add a new chain-agnostic changeset

1. Read `deployment/docs/style-guide.md`. It covers idempotency, `AddressRef` resolution, defaults, stale reads, and qualifiers.
2. **Constructor.** Take the registries as arguments (so tests can inject them) and return `cldf.ChangeSetV2[Cfg]` from `cldf.CreateChangeSet(apply, verify)`.
   - For a single sequence on a single chain, use `changesets.NewFromOnChainSequence` instead; its config is wrapped in `changesets.WithMCMS[Cfg]`.
3. **`verify`** must reject what `apply` cannot do without side effects:
   - unknown selectors
   - families with no registered adapter
   - duplicate selectors
   - invalid `mcms.Input` when needed
4. **`apply`:**
   1. For each chain: `chain_selectors.GetSelectorFamily(sel)` → registry lookup → `cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.X(), e.BlockChains, input)`.
   2. Collect `report.Output.Addresses` into a `datastore.NewMemoryDataStore()`.
   3. Collect `BatchOps` and `ExecutionReports`.
   4. Return `changesets.NewOutputBuilder(e, mcmsRegistry).WithReports(...).WithBatchOps(...).WithDataStore(ds).Build(cfg.MCMS)`.
5. **Config.** Infer well-known addresses from the DataStore via the adapter (users shouldn't pass a Router address). Use camelCase `json`/`yaml` tags. Use pointer fields for "optional, nil means default".
6. **Tests.** Cover verify errors and the apply path with a mocked or real adapter. Then add an end-to-end test in `integration-tests/deployment` or the family module.
7. **Wire into durable pipelines.** Changesets run through durable pipelines in `../chainlink-deployments` (`domains/<domain>/<env>/durable_pipelines.go`, `registry.Add(...)`). Add the pipeline there, plus input YAML under `durable_pipelines/inputs/`. For the end-to-end flow and MCMS signer groups, see `../chainlink-deployments/domains/ccip/.agents/skills/solana-durable-pipelines-mcms/SKILL.md`.
8. **Docs.** Add the changeset to the catalog in `deployment/docs/consuming.md`, and to the "v1 vs v2 per flow" table if it changes a flow.

---

## Consuming (for tests and examples)

```go
import _ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters" // triggers init()

cs := v2changesets.ConfigureChainsForLanesFromTopology(
    adapters.GetCommitteeVerifierContractRegistry(), adapters.GetChainFamilyRegistry(), changesets.GetRegistry())
out, err := cs.Apply(env, cfg) // out.DataStore, out.MCMSTimelockProposals, out.Reports
```

For v1 and v2 examples and the full catalog, see `deployment/docs/consuming.md`.

## Verify before finishing

Run these from each touched module (`deployment/`, `chains/<family>/deployment/`, or the external repo):

```bash
go build ./...
go vet ./...
go test ./<touched packages>/...
```

Non-obvious points:

- `chains/evm/deployment` and `chains/solana/deployment` have `replace github.com/smartcontractkit/chainlink-ccip/deployment => ../../../deployment`, so they compile against your local changes.
- External repos (for example `chainlink-ccip-solana`) pin a version. After merging an interface change, bump them with `go get github.com/smartcontractkit/chainlink-ccip/deployment@<sha>`.
- **Changing an interface breaks every implementer.** Grep the whole workspace, including `../chainlink-ccip-solana`, for implementations, and fix or flag them.

## Pitfalls checklist

- [ ] Registry key matches what callers pass. A family-version mismatch fails silently at registration and loudly at run time ("no adapter registered").
- [ ] No state is shared through struct fields across registrations. Each `&Adapter{}` literal in `init()` is a separate instance.
- [ ] Decision reads are live on-chain reads, not cached operation reports.
- [ ] v2 versions are not pre-release (`2.0.0`, not `2.0.0-dev`).
- [ ] CCTP and Lombard adapters are registered in the `chainlink-deployments` durable pipelines (`New*Registry()`), not in `init()`. A new family adapter must be added there too.
- [ ] `MCMSReader` is registered for the family. Without it, `OutputBuilder.Build` cannot create proposals for that family's chains.
- [ ] If one pool program serves many tokens (as on Solana), `DeriveTokenPoolCounterpart` returns the per-token address, and `TokenFeeAdapter` accepts that address as its pool key. A changeset that calls `TokenFeeAdapter` gets the key from `tokens.TokenPoolCounterpartAddress`, not from `poolRef.Address`.
- [ ] Docs updated: `interfaces.md` for new interfaces, `consuming.md` for new changesets.
