---
title: "Adapter Interfaces Reference"
sidebar_label: "Interfaces"
sidebar_position: 6
---

# Adapter Interfaces Reference

Every interface a chain family can implement, grouped by API generation. Each entry lists its registry, the key it is registered under, and what it does. The source file is the source of truth for exact signatures. Go to it when you need parameter types.

- For how to implement these, see [Implementing 1.6](implementing-1.6.md) and [Implementing 2.0](implementing-2.0.md).
- For how registries work, see [Architecture](architecture.md#registries).

**Key column:**

- `family-version`: one adapter per chain family *and* contract version (`"evm-1.6.0"`).
- `family`: one adapter per chain family.
- *type assertion*: there is no registry. The changeset checks whether the already-registered adapter also implements the interface (`adapter.(X)`).

All registries keep the **first** registration for a key and silently ignore later ones.

---

## Shared by v1 and v2

Both API generations use these. A v2 chain family still needs them.

| Interface | Registry accessor → register method | Key | Source |
|---|---|---|---|
| `MCMSReader` | `changesets.GetRegistry().RegisterMCMSReader` | family | [utils/changesets/output.go](../utils/changesets/output.go) |
| `TransferOwnershipAdapter` | `deploy.GetTransferOwnershipRegistry().RegisterAdapter` | family-version | [deploy/product.go](../deploy/product.go) |
| `AddressNormalizer` | `deploy.GetAddressNormalizerRegistry().RegisterAddressNormalizer` | family | [deploy/product.go](../deploy/product.go) |
| `TokenAdapter` (+ optional token interfaces below) | `tokens.GetTokenAdapterRegistry().RegisterTokenAdapter` | family-version (pool version) | [tokens/product.go](../tokens/product.go) |
| `CurseAdapter` + `CurseSubjectAdapter` | `fastcurse.GetCurseRegistry().RegisterNewCurse` | family-version / family | [fastcurse/product.go](../fastcurse/product.go) |
| `FeeAggregatorAdapter` | `fees.GetFeeAggregatorRegistry().RegisterFeeAggregatorAdapter` | family-version | [fees/fee_aggregator.go](../fees/fee_aggregator.go) |
| `AuthorizedCallersAdapter` | `authorizedcallers.GetAuthorizedCallersRegistry().RegisterAdapter` | family + contract type + version | [authorizedcallers/product.go](../authorizedcallers/product.go) |
| `TestAdapter` (factory) | `testadapters.GetTestAdapterRegistry().RegisterTestAdapter` | family-version | [testadapters/adapters.go](../testadapters/adapters.go) |
| Hook providers | `hooks.Get*Registry().Register` | family | [hooks/](../hooks/) |

### MCMSReader

Resolves the MCMS and timelock contracts and the chain metadata (MCM address, starting op count) that `OutputBuilder` needs to build a proposal. Every chain family needs one, or no proposal can be built for it.

| Method | Purpose |
|---|---|
| `GetChainMetadata(e, sel, mcms.Input)` | MCM address + starting op count for the proposal |
| `GetTimelockRef(e, sel, mcms.Input)` | Timelock `AddressRef` selected by `mcms.Input.Qualifier` |
| `GetMCMSRef(e, sel, mcms.Input)` | MCM `AddressRef` selected by `mcms.Input.Qualifier` |

### TransferOwnershipAdapter

| Method | Purpose |
|---|---|
| `InitializeTimelockAddress(e, mcms.Input)` | Resolve and cache the timelock that becomes the new owner |
| `SequenceTransferOwnershipViaMCMS()` | Propose an ownership transfer for the given refs |
| `SequenceAcceptOwnership()` | Accept a pending transfer |
| `ShouldAcceptOwnershipWithTransferOwnership(e, in)` | `true` if accept must run in the same changeset (e.g. Solana) |

### AddressNormalizer

Converts between the family's address strings and bytes. Token changesets use it before datastore lookups so that two spellings of the same address match.

`NormalizeAddress(string)`, `BytesToString([]byte)`, `StringToBytes(string)`.

### TokenAdapter

Registered once per **token pool version**, because pool versions configure differently. The big difference is v1 vs v2: v2 pools take CCVs, finality config, and pool-level fees. All pool versions must be able to connect to each other.

| Method | Purpose |
|---|---|
| `ConfigureTokenForTransfersSequence()` | Register the token in the TokenAdminRegistry and configure remote chains on the pool |
| `AddressRefToBytes(ref)` | Family address encoding (hex on EVM, base58 on Solana) |
| `DeriveTokenAddress(e, sel, poolRef)` | Read the token address stored on the pool |
| `DeriveTokenDecimals(e, sel, poolRef, token)` | Read the token's decimals |
| `DeriveTokenPoolCounterpart(e, sel, pool, token)` | Turn the deployed pool address into the per-token pool address: the one the remote side must store, and the key of `TokenFeeAdapter` inputs (Solana: pool config PDA). Return `pool` unchanged if not applicable |
| `ManualRegistration()` | Register a customer-deployed token whose mint authority the customer no longer holds |
| `SetTokenPoolRateLimits()` | Set outbound/inbound rate limits for one remote |
| `DeployToken()` / `DeployTokenVerify(e, in)` | Deploy a token, and validate the input first |
| `DeployTokenPoolForToken()` | Deploy or initialize a pool for an existing token |
| `UpdateAuthorities()` | Hand token/pool ownership to the timelock |
| `MigrateLockReleasePoolLiquiditySequence()` | Move liquidity from a legacy lock-release pool to a 2.0 lockbox pool. Return `nil` if unsupported |

**Optional token interfaces.** The changeset discovers these by type assertion on your `TokenAdapter`, or registers them separately where noted:

| Interface | Discovered by | Used by changeset |
|---|---|---|
| `TokenRefResolver` | `RegisterTokenRefResolver` (family) | Every token changeset that accepts address-only refs |
| `TokenAdminRegistryReader` | `RegisterTokenAdminRegistryReader` (family) | Upgrade-safety checks (`GetActivePool`) |
| `TokenAdminRegistryManager` (reader + `UnregisterToken`) | `RegisterTokenAdminRegistryManager` (family) | `RemoveRemotePools` |
| `TokenFeeAdapter` | type assertion | `ConfigureTokenPool`, `ConfigureTokensForTransfers` (pool-level fees, finality) |
| `TokenPoolDynamicConfigAdapter` | type assertion | `ConfigureTokenPool` (router, rate-limit admin, fee admin) |
| `TokenAdminRoleAdapter` | type assertion | `GrantTokenAdminRole`, `RevokeTokenAdminRole` |
| `RemotePoolRemover` | type assertion | `RemoveRemotePools` |
| `RateLimitReaderAdapter` | type assertion | `SetTokenPoolRateLimits` (outbound-only path), auto-migrate |
| `TokenPoolMigrator` | type assertion | `ConfigureTokensForTransfers` auto-migrate, `RemoveRemotePools` |

**Which pool address each interface receives.** On EVM it is always the pool contract. On families where one pool program serves many tokens, it depends on the interface:

- **`TokenFeeAdapter`**: the pool's counterpart address, from `tokens.TokenPoolCounterpartAddress`, which calls `DeriveTokenPoolCounterpart`. That's the pool address on EVM and the pool config PDA on Solana. The PDA identifies both the pool program (its owner) and the mint (in its state).
- **`TokenPoolMigrator`**: the pool as the datastore or TokenAdminRegistry holds it (the program ID on Solana), plus the token address.
- **Other optional interfaces**: the resolved `TokenPoolRef` and `TokenRef`.

On Solana, allowed finality and token transfer fees are stored per lane, in the pool's `ChainConfigV2`. EVM stores finality per pool. The Solana `SetAllowedFinalityConfig` therefore writes the value to every lane the pool is configured for.

`TokenRefResolver` and `DeriveTokenAddress` do different jobs. The resolver answers "what is the full `AddressRef` for this address?" `DeriveTokenAddress` answers "which token does this already resolved pool serve?" `ResolveAdapterAndRefs` in [tokens/token_expansion.go](../tokens/token_expansion.go) uses both, in this order:

1. It resolves the pool ref from the datastore, then falls back to `ResolveTokenPoolRef`.
2. It picks the `TokenAdapter` by the pool's version.
3. It calls `DeriveTokenAddress` on the resolved pool.
4. If that fails, it calls `ResolveTokenRef` on the user's token ref.

A resolver may tag reconstructed refs with the label `tokens.ArtificialAddressRefLabel` ("ArtificialAddressRef"). Solana does this when the user passes a pool config PDA instead of the program ID.

### CurseAdapter / CurseSubjectAdapter

Both are registered together with `RegisterNewCurse(fastcurse.CurseRegistryInput{CursingFamily, CursingVersion, CurseAdapter, CurseSubjectAdapter})`. The curse adapter is keyed by family-version, and the subject adapter by family.

- `CurseAdapter`: `Initialize`, `IsSubjectCursedOnChain`, `IsChainConnectedToTargetChain`, `IsCurseEnabledForChain`, `SubjectToSelector`, `Curse()`, `Uncurse()`, `ListConnectedChains`.
  - `IsSubjectCursedOnChain` returns `true` only when that exact subject is cursed. Unlike EVM RMN, it does not return `true` because of a global curse. To check for a global curse, query `GlobalCurseSubject()`.
- `CurseSubjectAdapter`: `SelectorToSubject`, and `DeriveCurseAdapterVersion` (picks which `CurseAdapter` version to use on a chain).

### FeeAggregatorAdapter

`SetFeeAggregator(e)`, `GetFeeAggregator(e, sel)`, `WithdrawFeeTokens(e)`. Where the aggregator lives differs by family and version:

| Family and version | Where the aggregator is stored |
|---|---|
| EVM 1.6 | OnRamp dynamic config |
| EVM 2.0 | Proxy (the default), OnRamp, Executor, and USDC proxy |
| Solana 1.6 | Router |

### AuthorizedCallersAdapter

Manages callers on any contract that inherits `AuthorizedCallers.sol`. It is keyed by `(family, ContractType, version)`.

`Initialize`, `GetAllAuthorizedCallers`, `ApplyAuthorizedCallerUpdates()`, `NormalizeCaller`.

### Test and hook providers

| Interface | Registry | Purpose |
|---|---|---|
| `TestAdapter` | `testadapters.GetTestAdapterRegistry().RegisterTestAdapter(family, version, factory)` | Send and validate CCIP messages in integration tests. Registered as `func(*cldf.Environment, uint64) TestAdapter` |
| `TestAdapterForFamily` / `ForkCCIPSendTestAdapter` | `RegisterTestAdapterForFamily` / `RegisterForkCCIPSendTestAdapter` | Narrow subsets of `TestAdapter` for fork smoke tests that don't need a live client |
| `MessageExecutor` | `testadapters.GetMessageExecutorRegistry().Register` | Execute a sent message manually on dest (optional post-send step) |
| `ContractVerification` | `hooks.GetContractVerificationRegistry().Register` | Block-explorer verification pre/post hooks |
| `ContractOwnership` | `hooks.GetContractOwnershipRegistry().Register` | "Contracts must be owned by the timelock" pre-hook |
| `PostProposalCCIPSend` | `hooks.GetPostProposalCCIPSendRegistry().Register` | Smoke-test a CCIP send after a proposal executes |
| `PostProposalLaneSanity` | `hooks.GetPostProposalLaneSanityRegistry().Register` | CLI lane sanity checks (extends `PostProposalCCIPSend`) |

---

## v1 (1.6) lane and chain API

These drive the v1 changesets (`DeployContracts`, `ConnectChains`, `SetOCR3Config`, `DisableLane`, `UpdateFeeQuoterDests`, …). They are keyed by `family-version`, and 1.6.0 is the version a new family should register.

| Interface | Registry accessor → register method | Key | Source |
|---|---|---|---|
| `Deployer` | `deploy.GetRegistry().RegisterDeployer` | family-version | [deploy/product.go](../deploy/product.go) |
| `LaneAdapter` (+ optional lane interfaces) | `lanes.GetLaneAdapterRegistry().RegisterLaneAdapter` | family-version | [lanes/product.go](../lanes/product.go) |
| `FeeAdapter` | `fees.GetRegistry().RegisterFeeAdapter` | family-version (fee contract version) | [fees/product.go](../fees/product.go) |
| `FeeResolver` | `fees.GetRegistry().RegisterFeeResolver` | family | [fees/product.go](../fees/product.go) |
| `PingPongAdapter` | `lanes.GetPingPongAdapterRegistry().RegisterPingPongAdapter` | family-version | [lanes/pingpong.go](../lanes/pingpong.go) |
| `RampUpdateInRouter` / `RouterUpdateInRamp` | `deploy.GetLaneMigratorRegistry().RegisterRouterUpdater` / `RegisterRampUpdater` | family-version | [deploy/lanemigrator.go](../deploy/lanemigrator.go) |
| `FeeQuoterUpdater`, `RampUpdater`, `ConfigImporter`, `LaneVersionResolver` | `deploy.GetFQAndRampUpdaterRegistry().Register*` | family-version (resolver: family) | [deploy/feequoterupdater.go](../deploy/feequoterupdater.go) |

### Deployer

| Method | Purpose |
|---|---|
| `DeployChainContracts()` | Deploy Router, OnRamp, OffRamp, FeeQuoter, RMNRemote, etc. on one chain |
| `DeployMCMS()` | Deploy MCM (proposer/canceller/bypasser), timelock, call proxy |
| `FinalizeDeployMCMS()` | Second MCMS phase (Solana timelock init). Return a no-op sequence if not needed |
| `SetOCR3Config()` | Write OCR3 config (read from CCIPHome) to the OffRamp |
| `GrantAdminRoleToTimelock()` | Make one timelock admin of another |
| `UpdateMCMSConfig()` | Change signers and quorums on an existing MCM |

### LaneAdapter

| Method | Purpose |
|---|---|
| `ConfigureLaneLegAsSource()` | This chain sends to the remote: OnRamp dest config, FeeQuoter dest config, prices |
| `ConfigureLaneLegAsDest()` | This chain receives from the remote: OffRamp source config, Router offramp |
| `DisableRemoteChain()` | Disable the remote on this chain (used by `DisableLane`) |
| `GetOnRampAddress` / `GetOffRampAddress` / `GetRouterAddress` / `GetFQAddress` | Addresses as **bytes in this family's encoding**. `ConnectChains` uses them to fill `ChainDefinition` |
| `GetFeeQuoterDestChainConfig()` | Default FeeQuoter config **for this chain as a destination**, applied on the remote source |
| `GetDefaultGasPrice()` | Default USD price (18 decimals) per gas unit for this chain as a destination |

**Optional lane interfaces.** The changeset discovers these by type assertion on your `LaneAdapter`:

| Interface | Effect |
|---|---|
| `ChainMetadataProvider` | `GetChainFamilySelector() [4]byte`. Registers your 4-byte family selector so `utils.GetSelectorHex` works without a hardcoded case |
| `TestRouterProvider` | `GetTestRouter`. Used when a lane sets `TestRouter: true` |
| `TokenPriceProvider` | `GetDefaultTokenPrices`. Default fee-token prices (EVM: LINK, WETH) |
| `DynamicFeeQuoter` | `GetFQAddressDynamic`. Resolve the FeeQuoter from chain state instead of the datastore |
| `FeeQuoterVersionProvider` | `GetFQVersion`. Report the FeeQuoter version (1.6 or 2.0) so `ConnectChains` picks the right ops |

### FeeAdapter and FeeResolver

`fees.ResolveFeeAdapter` ([fees/defaults.go](../fees/defaults.go)) looks up the adapter for a lane in this order:

1. `FeeResolver.GetOnRampRef` finds the lane's OnRamp from the router.
2. `FeeAdapter.GetFeeContractRef` finds the fee contract. This is the OnRamp for 1.5 and the FeeQuoter for 1.6 and later.
3. The adapter is picked by the **fee contract's version**, not the OnRamp's version.

FeeAdapter methods: `GetFeeContractRef`, `SetTokenTransferFee`, `GetOnchainTokenTransferFeeConfig`, `GetDefaultTokenTransferFeeConfig`, `ApplyDestChainConfigUpdates`, `GetOnchainDestChainConfig`, `GetDefaultDestChainConfig`. Adapters take a `Bundle`, `BlockChains`, and `DataStore` instead of a full `Environment`.

### Lane migration and FeeQuoter upgrade

- `RampUpdateInRouter.UpdateRouter()` points the Router at new ramps.
- `RouterUpdateInRamp.VerifyPreconditions` and `UpdateVersionWithRouter()` point the new ramps at the Router. These are used by `LaneMigrateToNewVersionChangeset`.
- `FeeQuoterUpdater`, `RampUpdater`, and `ConfigImporter` back `UpdateFeeQuoterChangeset`. That changeset rebuilds FeeQuoter config by importing existing 1.5 and 1.6 lane config from chain.

---

## v2 (2.0) chain API

These live in [`v2_0_0/adapters`](../v2_0_0/adapters/) and drive the 2.0 changesets in [`v2_0_0/changesets`](../v2_0_0/changesets/). Most of them are keyed by **family only**: 2.0 contract versions are chosen inside the adapter (via `GetDefaultDeployContractParams`), not by the registry key.

| Interface | Registry accessor → register method | Key | Required? | Source |
|---|---|---|---|---|
| `DeployChainContractsAdapter` | `adapters.GetDeployChainContractsRegistry().Register` | family | yes | [deploy_chain_contracts.go](../v2_0_0/adapters/deploy_chain_contracts.go) |
| `ChainFamily` | `adapters.GetChainFamilyRegistry().RegisterChainFamily` | family | yes | [chain_family.go](../v2_0_0/adapters/chain_family.go) |
| `CommitteeVerifierContractAdapter` | `adapters.GetCommitteeVerifierContractRegistry().Register` | family | yes | [committee_verifier_contract.go](../v2_0_0/adapters/committee_verifier_contract.go) |
| `LaneVersionResolver` | `adapters.GetDeployChainContractsRegistry().RegisterLaneVersionResolver` | family | recommended (blocks 2.0→1.6 lane downgrades) | [deploy/product.go](../deploy/product.go) |
| `ConfigImporter` | `adapters.GetDeployChainContractsRegistry().RegisterConfigImporter` | family-version | for migrations | [deploy/product.go](../deploy/product.go) |
| `OnRampUpgrader` | `adapters.GetOnRampUpgraderRegistry().Register` | family | to support OnRamp upgrades | [upgrade_onramp.go](../v2_0_0/adapters/upgrade_onramp.go) |
| `OffRampSourceOnRampSetter` / `Reader` | type assertion on `ChainFamily` | — | for OnRamp upgrades / `OffRampSetSourceOnRamps` | [offramp_source_onramps.go](../v2_0_0/adapters/offramp_source_onramps.go) |
| `GasPriceValidator` | type assertion on `ChainFamily` | — | optional preflight | [chain_family.go](../v2_0_0/adapters/chain_family.go) |
| `CCTPChain` | `adapters.NewCCTPChainRegistry().RegisterCCTPChain` | family + USDC type | for USDC/CCTP | [cctp.go](../v2_0_0/adapters/cctp.go) |
| `LombardChain` | `adapters.NewLombardChainRegistry().RegisterLombardChain` | family | for Lombard | [lombard.go](../v2_0_0/adapters/lombard.go) |
| `TestVerifierChainAdapter` | `adapters.GetTestVerifierChainRegistry().Register` | family | test envs | [test_verifier_chain.go](../v2_0_0/adapters/test_verifier_chain.go) |

The CCTP and Lombard registries have no global singleton, by design. The durable pipeline in `chainlink-deployments` builds one with `New*Registry()`, registers each family's adapter, and passes it to `DeployCCTPChains` or `DeployLombardChains`.

### DeployChainContractsAdapter

The `DeployChainContracts` (2.0) changeset calls these methods in this order, once per chain:

| # | Method | Purpose |
|---|---|---|
| 1 | `GetDefaultDeployContractParams(sel)` | Family defaults: contract versions, FeeQuoter static config, executors |
| 2 | `ResolveDeployAddresses(e, sel)` | Find or deploy prerequisites, at minimum the `DeployerContract` (EVM: CREATE2 factory). Returns `NewAddressRefs` to persist |
| 3 | `BuildDeployContractParams(input)` | Merge defaults + topology-derived `CommitteeVerifiers` + user `Overrides`. Call `adapters.ApplyDeployContractParamsOverrides` at the end |
| 4 | `DeployChainContracts()` | Sequence that deploys the contracts. Returns `DeployChainContractsOutput` with `RefsToTransferOwnership` (CLL timelock) and `RefsToTransferOwnershipRMN` |

### ChainFamily

The 2.0 replacement for `LaneAdapter`. Lanes are configured **per chain** (`ConfigureChainForLanes`) for all of that chain's remotes at once, not leg by leg.

| Method | Purpose |
|---|---|
| `ConfigureChainForLanes()` | Sequence. Configures OnRamp, OffRamp, FeeQuoter, CommitteeVerifiers, Executor, and Router (router **last**) for every remote in `ConfigureChainForLanesInput.RemoteChains`. Must be idempotent |
| `AddressRefToBytes(ref)` | Family address encoding |
| `GetOnRampAddress(ds, sel)` | OnRamp bytes **as this chain writes them into messages**. The remote OffRamp hashes these bytes. EVM: 20-byte address abi-encoded to 32 bytes |
| `GetOffRampAddress` / `GetFQAddress` / `GetRouterAddress` / `GetTestRouter` | Native-encoding bytes (destination-side addresses are never padded) |
| `ResolveExecutor(ds, sel, qualifier)` | Executor address for a qualifier |
| `GetAddressBytesLength()` | Address length on this family (EVM 20, Solana 32) |
| `GetChainFamilySelector()` | 4-byte family selector |
| `GetDefaultFeeQuoterDestChainConfig(sel, remote, familySel)` | FeeQuoter dest config defaults for `remote` as seen from `sel` |
| `GetDefaultRemoteChainConfig(src, remote)` | Executor fee, base execution gas, network fees, `SkipExecutorConfig`, … |
| `GetDefaultCommitteeVerifierRemoteChainConfig()` | Verifier fee, gas, and payload size defaults |
| `GetDefaultFinalityConfig()` | Default allowed finality |
| `ValidateNOPsTopology(sel, nopCount)` | Reject topologies the family can't support |

### CommitteeVerifierContractAdapter

`ResolveCommitteeVerifierContracts(ds, sel, qualifier)` returns the verifier contracts for a committee qualifier. `GetCommitteeVerifierResolver(ds, sel, qualifier)` returns the resolver contracts that are used as CCV addresses on lanes.

### OnRampUpgrader

Backs the five-phase OnRamp upgrade (`UpgradeOnrampPhase1/2/3`, `UpgradeOnrampPhase3Rollback`, `UpgradeOnrampCleanup`):

1. Deploy the new OnRamp.
2. Stage it behind the TestRouter.
3. Promote it to the production Router.
4. Roll back if needed.
5. Remove the legacy OnRamp from the remote OffRamp whitelists.

The remote `ChainFamily` must also implement `OffRampSourceOnRampSetter` and `OffRampSourceOnRampReader`.

### CCTPChain / LombardChain / TestVerifierChainAdapter

These follow the same pattern. Each is split into a *remote* interface and a *local* one:

- The **remote** interface (`RemoteCCTPChain`, `RemoteLombardChain`, `RemoteTestVerifierChain`) returns this chain's addresses for **other** chains to store.
- The **local** interface adds `Deploy…Chain()` and `Configure…ChainForLanes()` sequences.

Their dependency structs (`*Deps`) receive `RemoteChains map[uint64]Remote…Chain`, so a sequence on one family can ask another family for its addresses without importing that family's code.
