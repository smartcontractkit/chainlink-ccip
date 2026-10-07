---
title: "Implementing the 2.0 Tooling API"
sidebar_label: "Implementing 2.0"
sidebar_position: 5
---

# Implementing the 2.0 Tooling API for a Chain Family

This guide adds a chain family to the **CCIP 2.0** changesets in [`v2_0_0/changesets`](../v2_0_0/changesets/): `DeployChainContracts`, `ConfigureChainsForLanesFromTopology`, the OnRamp upgrade phases, CCTP, Lombard, and the test verifier. Read [Architecture](architecture.md#v1-and-v2-dispatch-differ) first.

2.0 replaces OCR-based commit/exec with **committee verifiers (CCVs)** and **executors**. Most of their configuration comes from an **environment topology** (`offchain.EnvironmentTopology`), not from per-lane input.

Reference implementations:

| Family | Where | Registrations |
|---|---|---|
| EVM 2.0 | [`chains/evm/deployment/v2_0_0/adapters`](../../chains/evm/deployment/v2_0_0/adapters) | `init.go` |
| Solana 2.0 | `chainlink-ccip-solana` repo, `deployment/v2_0_0/adapters` | `init.go` (sequences are in `deployment/v2_0_0/sequences`, shared helpers in `deployment/solcommon`) |

## What 2.0 reuses from 1.6

A 2.0 chain family still registers these shared adapters. They are the same interfaces as in [Implementing 1.6](implementing-1.6.md#4-the-adapter-struct):

| Interface | 2.0 notes |
|---|---|
| `changesets.MCMSReader` | Required. Resolve the 2.0 MCMS refs (Solana 2.0 looks up MCM refs at version 2.0.0) |
| `deploy.TransferOwnershipAdapter` | Register at `2.0.0`. `DeployChainContracts` uses it to hand contracts to the timelock |
| `tokens.TokenAdapter` | Register under each **2.0 pool version** (EVM `2.0.0`; Solana `2.0.0` and `1.6.1`). Read CCVs from `RemoteChainConfig.OutboundCCVs`/`InboundCCVs`, and finality from `AllowedFinalityConfig`. 2.0 pools carry their own token transfer fees and allowed finality, so also implement `tokens.TokenFeeAdapter`. Its pool keys are counterpart addresses (see [TokenAdapter](interfaces.md#tokenadapter)) |
| `fastcurse` curse adapters | Register at the RMN version your chain uses (EVM registers RMN 2.1.0) |
| `fees.FeeAdapter`, `fees.FeeAggregatorAdapter` | Register at `2.0.0` |
| `authorizedcallers.AuthorizedCallersAdapter` | If your 2.0 contracts inherit `AuthorizedCallers` |
| `fees.FeeResolver`, `deploy.AddressNormalizer`, `tokens.TokenRefResolver`, `tokens.TokenAdminRegistryReader`/`Manager` | Keyed by family only, but **still required**: `SetTokenTransferFee`/`UpdateFeeQuoterDests` fail without a `FeeResolver`, and the token changesets need the resolver, normalizer and TAR reader. If your 1.6 package registers them, a pipeline that imports only your 2.0 package won't have them. Register them from the 2.0 package too, or make sure the version-agnostic package is imported |

You do **not** register a `lanes.LaneAdapter` or a `deploy.Deployer` for 2.0. `ChainFamily` and `DeployChainContractsAdapter` replace them. EVM still registers its 1.6 `Deployer` struct at `2.0.0`, so `DeployMCMS` and the other `deploy.*` MCMS changesets keep working when they are called with version 2.0.0.

## Checklist

Required:

- [ ] `adapters.DeployChainContractsAdapter`
- [ ] `adapters.ChainFamily`
- [ ] `adapters.CommitteeVerifierContractAdapter`
- [ ] The shared adapters listed above

Recommended or optional:

- [ ] `deploy.LaneVersionResolver` (`RegisterLaneVersionResolver`). Stops a stale 1.6 payload from downgrading a lane that is already on 2.0
- [ ] `deploy.ConfigImporter` (`RegisterConfigImporter`). Imports 1.6 config when migrating
- [ ] `adapters.GasPriceValidator` on your `ChainFamily`. Fails the whole changeset before any chain is written
- [ ] `adapters.OnRampUpgrader`, plus `OffRampSourceOnRampSetter`/`Reader` on your `ChainFamily`. Needed for OnRamp upgrades
- [ ] `adapters.CCTPChain`, `adapters.LombardChain`, `adapters.TestVerifierChainAdapter`
- [ ] If your nodes run CCV jobs: the off-chain adapters in `github.com/smartcontractkit/chainlink-ccv/deployment` (aggregator, executor, verifier, indexer, token verifier, lane config, protocol-contracts deploy, committee-verifier deploy). See [Off-chain adapters](#5-off-chain-ccv-adapters)

## 1. DeployChainContractsAdapter

The `DeployChainContracts` changeset calls the four methods in order, once per chain. Solana's full implementation is short (`chainlink-ccip-solana/deployment/v2_0_0/adapters/solana_deploy_chain_contracts_adapter.go`):

```go
type SolanaDeployChainContractsAdapter struct{}

var _ ccvadapters.DeployChainContractsAdapter = (*SolanaDeployChainContractsAdapter)(nil)

// 1. Defaults: contract versions and static config not derived from topology.
func (a *SolanaDeployChainContractsAdapter) GetDefaultDeployContractParams(_ uint64) ccvadapters.DeployContractParams {
    v200 := semver.MustParse("2.0.0")
    return ccvadapters.DeployContractParams{
        FeeQuoter: ccvadapters.FeeQuoterDeployParams{Version: v200, MaxFeeJuelsPerMsg: new(big.Int).Mul(big.NewInt(2e2), big.NewInt(1e9))},
        Executors: []ccvadapters.ExecutorDeployParams{{Version: v200, MaxCCVsPerMsg: 10}},
    }
}

// 2. Prerequisites. Solana needs none, so it returns a placeholder deployer.
func (a *SolanaDeployChainContractsAdapter) ResolveDeployAddresses(_ deployment.Environment, _ uint64) (ccvadapters.DeployChainResolvedAddresses, error) {
    return ccvadapters.DeployChainResolvedAddresses{DeployerContract: solana.SystemProgramID.String()}, nil
}

// 3. Merge defaults + topology committee verifiers + user overrides.
func (a *SolanaDeployChainContractsAdapter) BuildDeployContractParams(input ccvadapters.BuildDeployContractParamsInput) (ccvadapters.DeployContractParams, error) {
    params := input.Defaults
    params.CommitteeVerifiers = input.CommitteeVerifiers
    return ccvadapters.ApplyDeployContractParamsOverrides(params, input.Overrides), nil
}

// 4. The on-chain deploy sequence.
func (a *SolanaDeployChainContractsAdapter) DeployChainContracts() *operations.Sequence[ccvadapters.DeployChainContractsInput, ccvadapters.DeployChainContractsOutput, cldf_chain.BlockChains] {
    return sequences.DeployChainContracts
}
```

EVM differs in two places (`chains/evm/deployment/v2_0_0/adapters/deploy_chain_contracts_adapter.go`):

- `ResolveDeployAddresses` finds or deploys the **CREATE2 factory**, and returns it as `DeployerContract` plus a `NewAddressRefs` entry to persist.
- `DeployChainContracts()` wraps the EVM-typed sequence. It converts `DeployChainContractsInput` to EVM types and calls `cldf_ops.ExecuteSequence(b, sequences.DeployChainContracts, evmChain, evmInput)`.

What your deploy sequence must do:

- **Be idempotent.** Skip contracts already in `input.ExistingAddresses`.
- **Deploy one CommitteeVerifier per `input.ContractParams.CommitteeVerifiers` entry**, using each entry's `Qualifier` as the datastore qualifier. Also deploy the executors in `ContractParams.Executors`.
- **Return the ownership lists.** Return the refs the changeset must hand to the timelock in `RefsToTransferOwnership` (CLLCCIP timelock) and `RefsToTransferOwnershipRMN` (RMNMCMS timelock). This step is skipped when `input.DeployerKeyOwned` is true.

## 2. ChainFamily

`ChainFamily` is the 2.0 lane adapter. `ConfigureChainsForLanesFromTopology` runs three phases:

1. **Enrich.** Fetch signer keys from JD, resolve the verifier contracts (via your `CommitteeVerifierContractAdapter`), and compute signer quorums from the topology.
2. **Resolve.** Ask the **remote** family's `ChainFamily` for its OnRamp and OffRamp bytes, and your own for your local contracts.
3. **Dispatch.** Run your `ConfigureChainForLanes()` sequence once with all of the chain's remotes.

The EVM adapter shows the key encoding rule (`chains/evm/deployment/v2_0_0/adapters/chain_family.go`):

```go
type ChainFamilyAdapter struct{}

var _ ccvadapters.ChainFamily = (*ChainFamilyAdapter)(nil)

func (a *ChainFamilyAdapter) ConfigureChainForLanes() *operations.Sequence[ccvadapters.ConfigureChainForLanesInput, seq_core.OnChainOutput, cldf_chain.BlockChains] {
    return sequences.ConfigureChainForLanes
}

// OnRamp: the bytes this chain writes into a message's onRampAddress field.
// EVM abi-encodes it (20 bytes left-padded to 32). The remote OffRamp hashes exactly these bytes.
func (a *ChainFamilyAdapter) GetOnRampAddress(ds datastore.DataStore, sel uint64) ([]byte, error) {
    return datastore_utils.FindAndFormatCanonicalRef(ds, datastore.AddressRef{
        ChainSelector: sel, Type: datastore.ContractType(onramp.ContractType), Version: onramp.Version,
    }, sel, evm_datastore_utils.ToABIEncodedEVMAddress)
}

// OffRamp and other destination-side addresses: native, unpadded (20 bytes on EVM).
func (a *ChainFamilyAdapter) GetOffRampAddress(ds datastore.DataStore, sel uint64) ([]byte, error) {
    return datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
        ChainSelector: sel, Type: datastore.ContractType(offramp.ContractType), Version: offramp.Version,
    }, sel, evm_datastore_utils.ToEVMAddressBytes)
}
```

On Solana there is no separate OnRamp. `GetOnRampAddress` returns the Router, and so does `GetTestRouter`. `GetAddressBytesLength()` returns 32.

### What `ConfigureChainForLanes` must do

The input is [`ConfigureChainForLanesInput`](../v2_0_0/adapters/chain_family.go). It holds your local contracts as bytes, the `CommitteeVerifiers` with per-remote signature configs, and `RemoteChains map[uint64]RemoteChainConfig`. In each `RemoteChainConfig`:

- `OnRamps` and `OffRamp` are remote addresses, already in the remote family's encoding.
- CCV and executor values are local address strings.

Follow the order EVM uses (documented at the top of `chains/evm/deployment/v2_0_0/sequences/configure_chain_for_lanes.go`):

1. **Infrastructure.** OffRamp source configs (whitelist the remote `OnRamps`), OnRamp dest configs (default and lane-mandated CCVs, default executor), Executor dest configs, and FeeQuoter dest configs and gas prices.
2. **CommitteeVerifier.** Outbound config per remote, and inbound signer quorum per remote.
3. **Router, last, in its own BatchOperation.** In a proposal it executes after everything else, so traffic is never routed to a half-configured lane.

Rules:

- **Idempotent writes.** Read on-chain state first, and emit a write only when it differs, so proposals stay minimal.
- **Respect the safety flags.**
  - Don't replace a production Router OnRamp unless `AllowOnrampOverride` is set.
  - Don't lower `BaseExecutionGasCost` below the on-chain value unless `AllowLoweringBaseExecutionGasCost` is set.
- **`SkipExecutorConfig`.** When set on a remote, don't touch the executor for that lane. EVM writes its no-execution sentinel instead.
- **`FamilyExtras`.** A pass-through map for family-specific options. Document the keys you read.

### Defaults are per family

`GetDefaultFeeQuoterDestChainConfig`, `GetDefaultRemoteChainConfig`, `GetDefaultCommitteeVerifierRemoteChainConfig`, and `GetDefaultFinalityConfig` give the values users get when they don't override them. Return pointers in `FeeQuoterDestChainConfigOverrides`: `nil` means "not set", and an explicit zero is kept. Solana's values are a good template (`deployment/v2_0_0/sequences/adapter.go` in `chainlink-ccip-solana`):

| Setting | Solana default |
|---|---|
| `MaxDataBytes` | 30 000 |
| `MaxPerMsgGasLimit` | 3 000 000 |
| `BaseExecutionGasCost` | 150 000 |
| `GasForVerification` | 500 000 |
| `PayloadSizeBytes` | 390 |
| Finality | wait for finality, wait for safe, depth 1 |

`ValidateNOPsTopology` rejects topologies your family can't run safely. EVM requires at least 9 NOPs in production and Solana 15.

## 3. CommitteeVerifierContractAdapter

Map a committee qualifier to contract refs:

- `ResolveCommitteeVerifierContracts` returns the CommitteeVerifier contract(s).
- `GetCommitteeVerifierResolver` returns the address lanes use as the CCV. On EVM this is the `CommitteeVerifierResolver` proxy in front of the verifier. Solana has no resolver, so it returns the verifier itself:

```go
func (a *SolanaCommitteeVerifierContractAdapter) GetCommitteeVerifierResolver(ds datastore.DataStore, sel uint64, qualifier string) ([]datastore.AddressRef, error) {
    return a.ResolveCommitteeVerifierContracts(ds, sel, qualifier)
}
```

## 4. Register in `init()`

Trimmed from `chains/evm/deployment/v2_0_0/adapters/init.go`:

```go
func init() {
    v := semver.MustParse("2.0.0")

    // 2.0 chain API (keyed by family)
    ccvadapters.GetChainFamilyRegistry().RegisterChainFamily(chainsel.FamilyEVM, &ChainFamilyAdapter{})
    ccvadapters.GetCommitteeVerifierContractRegistry().Register(chainsel.FamilyEVM, &EVMCommitteeVerifierContractAdapter{})
    ccvadapters.GetDeployChainContractsRegistry().Register(chainsel.FamilyEVM, &EVMDeployChainContractsAdapter{})
    ccvadapters.GetDeployChainContractsRegistry().RegisterConfigImporter(chainsel.FamilyEVM, utils.Version_1_6_0, &adapters1_6.ConfigImportAdapter{})
    ccvadapters.GetDeployChainContractsRegistry().RegisterLaneVersionResolver(chainsel.FamilyEVM, &adapters1_2.LaneVersionResolver{})
    ccvadapters.GetOnRampUpgraderRegistry().Register(chainsel.FamilyEVM, &EVMOnRampUpgrader{})
    ccvadapters.GetTestVerifierChainRegistry().Register(chainsel.FamilyEVM, &EVMTestVerifierChainAdapter{})

    // shared APIs at 2.0.0
    tokens.GetTokenAdapterRegistry().RegisterTokenAdapter(chainsel.FamilyEVM, v, NewTokenAdapter())
    fees.GetRegistry().RegisterFeeAdapter(chainsel.FamilyEVM, v, NewFeesAdapter(&evmAdapter))
    fees.GetFeeAggregatorRegistry().RegisterFeeAggregatorAdapter(chainsel.FamilyEVM, v, NewFeeAggregatorAdapter())
}
```

Solana 2.0 registers the same three core 2.0 adapters. It also registers its `MCMSReader`, `TransferOwnershipAdapter` (2.0.0), `TokenAdapter` (2.0.0 and 1.6.1), curse, fee, and fee aggregator adapters, and its family selector via `utils.RegisterChainFamilySelector`, because it has no `LaneAdapter` to do that.

Things that commonly go wrong:

- **CCTP and Lombard adapters are registered by the consumer, not in `init()`.** Those registries have no singleton. The durable pipeline builds one (`adapters.NewCCTPChainRegistry()`), registers each family's adapter, and passes it to `DeployCCTPChains`. Examples: `chainlink-deployments/domains/ccip/<env>/durable_pipelines.go` and `domains/ccv/pkg/pipelines/shared.go`. Export your adapter type so it can be registered there.
- **`OffRampSourceOnRampSetter`/`Reader` need no registration.** Implement them on your `ChainFamily` struct and they are found by type assertion.
- **Version strings must not be pre-release.** If your contracts report `2.0.0-dev`, still register and store `2.0.0`. Pre-release versions sort lower and break version comparisons (see `solcommon/constants.go`).

## 5. Off-chain CCV adapters

Running CCV jobs on your chain (verifiers, executors, aggregator, indexer) needs a second set of family-keyed registries in `github.com/smartcontractkit/chainlink-ccv/deployment/adapters`. Solana 2.0 registers all of them in the same `init()`:

`GetAggregatorRegistry`, `GetExecutorRegistry`, `GetVerifierRegistry`, `GetIndexerRegistry`, `GetTokenVerifierRegistry`, `GetCommitteeVerifierOnchainRegistry`, `GetProtocolContractsDeployRegistry`, `GetLaneConfigRegistry`, `GetCommitteeVerifierDeployRegistry`, plus `shared.RegisterChainTypeFamily` and `RegisterSigningIdentityReader`.

Those interfaces belong to `chainlink-ccv`. Use the Solana adapters in `chainlink-ccip-solana/deployment/v2_0_0/adapters/*_ccv_adapter.go` and `*_config_adapter.go` as the reference.

## 6. Optional 2.0 features

| Feature | Implement | Changesets |
|---|---|---|
| OnRamp upgrade | `OnRampUpgrader`, and on the remote `ChainFamily`: `OffRampSourceOnRampSetter` + `Reader` | `UpgradeOnrampPhase1/2/3`, `UpgradeOnrampPhase3Rollback`, `UpgradeOnrampCleanup`, `OffRampSetSourceOnRamps` |
| USDC via CCTP | `CCTPChain` (one per USDC type). Optional: `tokens.RemotePoolRemover`, `tokens.TokenPoolRateLimitSetter`, `tokens.TokenPoolDynamicConfigAdapter` | `DeployCCTPChains`, `RemoveCCTPRemotePools`, `SetCCTPTokenPoolRateLimits`, `SetCCTPTokenPoolDynamicConfig` |
| Lombard (LBTC) | `LombardChain`. Optional: the same three, plus `LombardAuthoritiesUpdater` | `DeployLombardChains`, `RemoveLombardRemotePools`, `SetLombardTokenPoolRateLimits`, `SetLombardTokenPoolDynamicConfig` |
| Test verifier token (TESTVTR) | `TestVerifierChainAdapter` | `DeployTestVerifierChains` |

Each optional interface has a *remote* half (`Remote…Chain`) that only returns addresses. A sequence on chain A gets `RemoteChains map[uint64]Remote…Chain` in its deps, so it can ask chain B's family for its pool, token, or caller addresses without importing chain B's code.

## 7. Family-specific changesets

The shared 2.0 changesets cover the cross-family flows. Operations that only make sense on your family belong in your own `v2_0_0/changesets` package. Examples: EVM's CREATE2 factory, lockbox funding, and Glamsterdam gas updates; Solana's program upgrades, lookup tables, and IDL publishing.

Build simple ones with `changesets.NewFromOnChainSequence`. EVM's single-chain `DeployChainContracts` is an example (`chains/evm/deployment/v2_0_0/changesets/deploy_chain_contracts.go`):

```go
var DeployChainContracts = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
    sequences.DeployChainContractsInput, evm.Chain, DeployChainContractsCfg,
]{
    Sequence: wrappedDeployChainContracts,
    ResolveInput: func(e cldf_deployment.Environment, cfg DeployChainContractsCfg) (sequences.DeployChainContractsInput, error) {
        addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(cfg.ChainSel))
        return sequences.DeployChainContractsInput{ChainSelector: cfg.ChainSel, ExistingAddresses: addresses /* ... */}, nil
    },
    ResolveDep: evm_sequences.ResolveEVMChainDep[DeployChainContractsCfg],
})
```
