---
title: "Implementing the 1.6 Tooling API"
sidebar_label: "Implementing 1.6"
sidebar_position: 4
---

# Implementing the 1.6 Tooling API for a Chain Family

This guide adds a chain family to the **1.6** changesets: `DeployContracts`, `DeployMCMS`, `ConnectChains`, `SetOCR3Config`, token, fee, ownership, and curse changesets. Read [Architecture](architecture.md) first. For 2.0, see [Implementing 2.0](implementing-2.0.md); much of this page (MCMS, tokens, ownership) applies there too.

Reference implementations:

| Family | Where | Notes |
|---|---|---|
| EVM | [`chains/evm/deployment`](../../chains/evm/deployment) | One stateless `EVMAdapter`, plus a version-agnostic `v1_0_0` layer. Operations are generated from gethwrappers |
| Solana | [`chains/solana/deployment`](../../chains/solana/deployment) | One `SolanaAdapter` that implements almost everything. Two-phase MCMS, PDAs, Router doubles as OnRamp. See [Solana notes](#8-solana-notes) |

## Checklist

Required for the core 1.6 flows:

- [ ] `changesets.MCMSReader`, registered once per family
- [ ] `deploy.Deployer`
- [ ] `lanes.LaneAdapter` (plus `ChainMetadataProvider` so your family selector is registered)
- [ ] `deploy.TransferOwnershipAdapter`
- [ ] `tokens.TokenAdapter` for each pool version you support, plus `TokenRefResolver`
- [ ] `deploy.AddressNormalizer`
- [ ] `fees.FeeResolver` + `fees.FeeAdapter`
- [ ] `fastcurse.CurseAdapter` + `CurseSubjectAdapter`
- [ ] `fees.FeeAggregatorAdapter`

Optional, add as needed: `TokenAdminRegistryManager`, the [optional token interfaces](interfaces.md#tokenadapter), `PingPongAdapter`, lane-migration adapters, `TestAdapter`, and the hook providers.

## 1. Module layout

Put the adapter next to your contracts (`chainlink-ccip/chains/<family>/deployment`) or in your own repo (`chainlink-<family>/deployment`). Give it its own `go.mod`, which imports `github.com/smartcontractkit/chainlink-ccip/deployment`. Mirror this layout:

```
<family>/deployment/
├── utils/                 # contract types, address helpers, MCMS batch-op builder
├── v1_0_0/adapters/       # version-agnostic: MCMSReader, AddressNormalizer, FeeResolver, TAR manager  + init.go
├── v1_6_0/
│   ├── operations/<contract>/   # one package per contract/program
│   ├── sequences/
│   │   ├── adapter.go           # the adapter struct + init() registrations
│   │   ├── deploy_chain_contracts.go, connect_chains.go, ocr.go, mcms.go, tokens.go, ...
│   ├── adapters/                # separately registered adapters (curse, fees, fee aggregator) + init.go
│   └── testadapter/             # TestAdapter (optional)
└── docs/
```

Put version-agnostic code in `v1_0_0` and register it by family only. Put contract-version-specific code in `v1_6_x` and register it under that version.

## 2. Write operations

An operation does exactly one thing. Keep the "deployer or MCMS" decision inside the write operation, so sequences don't need to care who owns the contract.

**EVM.** Use the CLDF contract helpers, `contract.NewDeploy`, `NewWrite`, and `NewRead`, from `chainlink-deployments-framework/chain/evm/operations/contract`. `NewWrite` checks `IsAllowedCaller`. If the deployer key isn't allowed, it returns a batch op instead of sending (`chains/evm/deployment/v1_6_0/operations/onramp/onramp.go`):

```go
var ApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]DestChainConfigArgs, *OnRampContract]{
    Name:            "onramp:apply-dest-chain-config-updates",
    Version:         Version,
    Description:     "Calls applyDestChainConfigUpdates on the contract",
    ContractType:    ContractType,
    ContractABI:     OnRampABI,
    NewContract:     NewOnRampContract,
    IsAllowedCaller: contract.OnlyOwner[*OnRampContract, []DestChainConfigArgs],
    Validate:        func([]DestChainConfigArgs) error { return nil },
    CallContract: func(c *OnRampContract, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
        return c.ApplyDestChainConfigUpdates(opts, args)
    },
})
```

**Non-EVM.** Use `operations.NewOperation` and do the branching yourself (`chains/solana/deployment/v1_6_0/operations/fee_quoter/fee_quoter.go`):

```go
var Deploy = operations.NewOperation(
    "fee-quoter:deploy",
    Version,
    "Deploys the FeeQuoter program",
    func(b operations.Bundle, chain cldf_solana.Chain, input []datastore.AddressRef) (datastore.AddressRef, error) {
        return utils.MaybeDeployContract(b, chain, input, ContractType, Version, "", ProgramName)
    },
)
```

Solana write ops compare the program authority with `chain.DeployerKey`. If they match, the op calls `chain.Confirm(...)`. Otherwise it returns `utils.BuildMCMSBatchOperation(...)` in `OnChainOutput.BatchOps`. For an example, see `operations/router/router.go`.

Operation rules:

- IDs are `"<contract>:<action>"` and must be unique per version. Reports are keyed on the ID plus the input.
- Deploy ops must be idempotent. Check `ExistingAddresses` and return the existing ref instead of redeploying (`MaybeDeployContract`).
- Reads used for decisions must not come from cached reports. Re-read on-chain state. See the [style guide](style-guide.md#avoid-stale-reads-from-cached-operations).

## 3. Compose sequences

A sequence targets one chain, takes serializable input, and returns `OnChainOutput`. Use `sequences.RunAndMergeSequence` to compose (`chains/evm/deployment/v1_6_0/sequences/update_lanes.go`, abridged):

```go
var ConfigureLaneLegAsSource = operations.NewSequence(
    "ConfigureLaneLegAsSource",
    semver.MustParse("1.6.0"),
    "Configures lane leg as source on CCIP 1.6.0",
    func(b operations.Bundle, chains cldf_chain.BlockChains, input lanes.UpdateLanesInput) (sequences.OnChainOutput, error) {
        var result sequences.OnChainOutput
        // 1. FeeQuoter dest chain config + prices for input.Dest
        result, err := sequences.RunAndMergeSequence(b, chains, FeeQuoterApplyDestChainConfigUpdatesSequence, fqInput, result)
        if err != nil { return result, err }
        // 2. OnRamp dest chain config (+ allowlist)
        result, err = sequences.RunAndMergeSequence(b, chains, OnRampApplyDestChainConfigUpdatesSequence, onRampInput, result)
        if err != nil { return result, err }
        // 3. Router: point dest at the OnRamp last (zero address when input.IsDisabled)
        return sequences.RunAndMergeSequence(b, chains, RouterApplyRampUpdatesSequence, routerInput, result)
    },
)
```

Update the **Router last**, so the lane only opens once everything behind it is configured.

## 4. The adapter struct

A single struct can implement most interfaces. EVM uses an empty struct. Solana's `SolanaAdapter` spreads its methods over `adapter.go`, `connect_chains.go`, `mcms.go`, `tokens.go`, `transfer_ownership.go`, and other files. Methods that return sequences just return a package-level sequence variable:

```go
type MyChainAdapter struct{}

func (a *MyChainAdapter) ConfigureLaneLegAsSource() *operations.Sequence[lanes.UpdateLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
    return ConfigureLaneLegAsSource
}
```

Add compile-time checks so a changed interface breaks your build instead of silently dropping a registration:

```go
var (
    _ lanes.LaneAdapter               = (*MyChainAdapter)(nil)
    _ lanes.ChainMetadataProvider     = (*MyChainAdapter)(nil)
    _ deploy.Deployer                 = (*MyChainAdapter)(nil)
    _ deploy.TransferOwnershipAdapter = (*MyChainAdapter)(nil)
    _ tokens.TokenAdapter             = (*MyChainAdapter)(nil)
    _ tokens.TokenRefResolver         = (*MyChainAdapter)(nil)
    _ changesets.MCMSReader           = (*MyChainAdapter)(nil)
)
```

### What each interface must get right

**`Deployer`**

- `DeployChainContracts` must be idempotent. The input has `ExistingAddresses`, so skip contracts that already exist.
- If your family deploys and initializes MCMS in one step, `FinalizeDeployMCMS` returns a no-op sequence. Solana uses this phase to initialize the timelock and roles.
- `SetOCR3Config` receives configs that were already read from CCIPHome. It only writes them to your OffRamp.

**`LaneAdapter`**

- `Get*Address` return bytes in **your** encoding: 20 bytes on EVM, a 32-byte pubkey on Solana. The remote chain stores these bytes as-is.
  - If your family has no separate OnRamp, return the contract that plays that role. On Solana, `GetOnRampAddress` returns the Router.
- `GetFeeQuoterDestChainConfig` and `GetDefaultGasPrice` describe your chain **as a destination**. They are applied on the remote source.
- `DisableRemoteChain` must close the lane on your side. `DisableLane` calls it on both chains.
- Implement `ChainMetadataProvider.GetChainFamilySelector()`, so `RegisterLaneAdapter` records your 4-byte selector.

**`MCMSReader`**

- Resolve the MCM and timelock refs using `input.Qualifier` (default `CLLCCIP`).
- `GetChainMetadata` must return the MCM address and the **current op count** read from chain.

**`TransferOwnershipAdapter`**

- If your contracts need the new owner to accept, and the deployer can do it in the same run, return `true` from `ShouldAcceptOwnershipWithTransferOwnership`. Solana does this when the current owner is the deployer key.

**`TokenAdapter`**

- Register one per pool version.
- `DeriveTokenPoolCounterpart` exists for families where the deployed pool address differs from the address the remote side must store. On Solana that is the config PDA, derived from (pool program, mint).
- Return `nil` from `MigrateLockReleasePoolLiquiditySequence` if you have no lockbox migration.

**`TokenRefResolver`**

- Lets users pass address-only refs. Rebuild a full `AddressRef` (type, version, qualifier) from on-chain state.

**`FeeResolver` / `FeeAdapter`**

- The resolver finds the OnRamp for a lane, and the adapter is selected by the **fee contract's** version. See [Interfaces](interfaces.md#feeadapter-and-feeresolver).

**`CurseAdapter`**

- `Initialize` loads the RMN addresses from the DataStore. It is called before every other method.
- `IsSubjectCursedOnChain` must **not** treat a global curse as "subject cursed".

## 5. Register in `init()`

Version-agnostic adapters (`v1_0_0/adapters/init.go`; EVM version shown):

```go
func init() {
    v := utils.Version_1_0_0
    deploy.GetTransferOwnershipRegistry().RegisterAdapter(chain_selectors.FamilyMyChain, v, &MyTransferOwnershipAdapter{})
    changesets.GetRegistry().RegisterMCMSReader(chain_selectors.FamilyMyChain, &MyMCMSReader{})
    tokens.GetTokenAdapterRegistry().RegisterTokenRefResolver(chain_selectors.FamilyMyChain, &MyTokenBase{})
    tokens.GetTokenAdapterRegistry().RegisterTokenAdminRegistryManager(chain_selectors.FamilyMyChain, &MyTokenBase{})
    fees.GetRegistry().RegisterFeeResolver(chain_selectors.FamilyMyChain, &MyFeeResolver{})
    deploy.GetAddressNormalizerRegistry().RegisterAddressNormalizer(chain_selectors.FamilyMyChain, &MyAddressNormalizer{})
}
```

1.6 adapters (`v1_6_0/sequences/adapter.go`; this is Solana's actual `init`):

```go
func init() {
    v := semver.MustParse("1.6.0")
    lanes.GetLaneAdapterRegistry().RegisterLaneAdapter(chain_selectors.FamilySolana, v, &SolanaAdapter{})
    deploy.GetRegistry().RegisterDeployer(chain_selectors.FamilySolana, v, &SolanaAdapter{})
    deploy.GetTransferOwnershipRegistry().RegisterAdapter(chain_selectors.FamilySolana, v, &SolanaAdapter{})
    changesets.GetRegistry().RegisterMCMSReader(chain_selectors.FamilySolana, &SolanaAdapter{})
    tokens.GetTokenAdapterRegistry().RegisterTokenAdapter(chain_selectors.FamilySolana, v, &SolanaAdapter{})
    tokens.GetTokenAdapterRegistry().RegisterTokenRefResolver(chain_selectors.FamilySolana, &SolanaAdapter{})
}
```

Curse, fee, and fee aggregator adapters (`v1_6_0/adapters/init.go`):

```go
func init() {
    fastcurse.GetCurseRegistry().RegisterNewCurse(fastcurse.CurseRegistryInput{
        CursingFamily:       chain_selectors.FamilyMyChain,
        CursingVersion:      utils.Version_1_6_0,
        CurseAdapter:        NewCurseAdapter(),
        CurseSubjectAdapter: NewCurseAdapter(),
    })
    fees.GetRegistry().RegisterFeeAdapter(chain_selectors.FamilyMyChain, utils.Version_1_6_0, NewFeesAdapter())
    fees.GetFeeAggregatorRegistry().RegisterFeeAggregatorAdapter(chain_selectors.FamilyMyChain, utils.Version_1_6_0, NewFeeAggregatorAdapter())
}
```

Things that commonly go wrong:

- **Each `&Adapter{}` literal is a separate instance.** Don't rely on struct fields to share state between registrations.
- **A wrong key fails silently.** For example, registering 1.6.0 while users pass 1.6.1. The changeset then fails with "no adapter registered".
- **The `init()` only runs if something imports the package.** The registry stays empty otherwise. See [Consuming](consuming.md#1-register-adapters-with-blank-imports).

## 6. Test

- Unit-test sequences against a simulated chain (EVM: `environment.New(ctx, environment.WithEVMSimulated(t, sels))`).
- Integration-test the shared changesets end to end in [`integration-tests/deployment`](../../integration-tests/deployment). These tests blank-import the adapter packages and run, for example, `fastcurse.CurseChangeset(...).Apply(*env, cfg)`.
- Implement `testadapters.TestAdapter` and register a factory with `RegisterTestAdapter(family, version, NewMyTestAdapter)`. This lets the shared message tests send and validate CCIP messages to and from your chain.

## 7. Address encoding cheat sheet

| Where | Encoding |
|---|---|
| `AddressRef.Address` | Your family's native string (hex, base58, …) |
| `TokenAdapter.AddressRefToBytes`, `LaneAdapter.Get*Address` | Native bytes (EVM 20, Solana 32) |
| `AddressNormalizer` | Canonical string ⇄ bytes. Used before datastore lookups |
| Remote addresses in your inputs (`[]byte`) | The **remote** family's encoding. Never re-encode them |

## 8. Solana notes

Solana is the reference for a non-EVM family. Where it differs from EVM:

| Aspect | EVM | Solana |
|---|---|---|
| Address | Hex, 20 bytes | Base58 `PublicKey`, 32 bytes |
| OnRamp | Separate contract | The Router plays the OnRamp role (`GetOnRampAddress` returns the Router) |
| Deployment | Deploy contract (CREATE/CREATE2) | Deploy a program, then initialize its config PDA. Each op calls the binding's `SetProgramID` first |
| Derived addresses | Known at deploy | PDAs from seeds (config, signer, and token pool PDAs) |
| MCMS | One phase; `FinalizeDeployMCMS` is a no-op | `DeployMCMS` deploys and initializes; `FinalizeDeployMCMS` sets MCM configs and timelock roles. MCMS refs are `programID.seed` strings |
| MCMS vs direct | `contract.NewWrite` checks the owner | Each op compares the program authority with the deployer key and returns `BuildMCMSBatchOperation(...)` when they differ |
| Token pools | One contract per pool | One shared pool program; per-mint state lives in a config PDA. Datastore pool refs are the **program ID**; `DeriveTokenPoolCounterpart` returns the PDA |
| Tokens | ERC-20 | SPL and SPL-2022; pool signer ATAs must exist |
| Lookup tables | — | The OffRamp lookup table is extended as lanes and pools are added |
| Accept ownership | Separate step | `ShouldAcceptOwnershipWithTransferOwnership` is true when the deployer is the current owner, so accept runs in the same changeset |
