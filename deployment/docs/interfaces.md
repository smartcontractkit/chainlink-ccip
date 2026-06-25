---
title: "Adapter Interfaces Reference"
sidebar_label: "Interfaces"
sidebar_position: 3
---

# Adapter Interfaces Reference

This document provides a complete API reference for all adapter interfaces that chain-specific implementations must or can provide. Each interface is accompanied by its registry, registration pattern, and method signatures.

For a step-by-step guide on implementing these interfaces for a new chain family, see [Implementing Adapters](implementing-adapters.md).

## Quick Reference

### Required Interfaces

Every chain family **must** implement these interfaces:

| Interface | Registry | Key Format | Source |
|-----------|----------|------------|--------|
| [Deployer](#deployer) | `DeployerRegistry` | `chainFamily-version` | [deploy/product.go](../deploy/product.go) |
| [LaneAdapter](#laneadapter) | `LaneAdapterRegistry` | `chainFamily-version` | [lanes/product.go](../lanes/product.go) |
| [TokenAdapter](#tokenadapter) | `TokenAdapterRegistry` | `chainFamily-version` | [tokens/product.go](../tokens/product.go) |
| [FeeAdapter](#feeadapter) | `FeeAdapterRegistry` | `chainFamily-version` | [fees/product.go](../fees/product.go) |
| [FeeAggregatorAdapter](#feeaggregatoradapter) | `FeeAggregatorAdapterRegistry` | `chainFamily-version` | [fees/fee_aggregator.go](../fees/fee_aggregator.go) |
| [MCMSReader](#mcmsreader) | `MCMSReaderRegistry` | `chainFamily` | [utils/changesets/output.go](../utils/changesets/output.go) |
| [TransferOwnershipAdapter](#transferownershipadapter) | `TransferOwnershipAdapterRegistry` | `chainFamily-version` | [deploy/product.go](../deploy/product.go) |
| [CurseAdapter](#curseadapter) | `CurseRegistry` | `chainFamily-version` | [fastcurse/product.go](../fastcurse/product.go) |
| [CurseSubjectAdapter](#cursesubjectadapter) | `CurseRegistry` | `chainFamily` | [fastcurse/product.go](../fastcurse/product.go) |

### Optional Interfaces

These interfaces are optional depending on what the chain family supports:

| Interface | Registry | Purpose | Source |
|-----------|----------|---------|--------|
| [TokenRefResolver](#tokenrefresolver) | `TokenAdapterRegistry` | Resolve token/pool refs from datastore or on-chain | [tokens/product.go](../tokens/product.go) |
| [TokenPriceProvider](#tokenpriceprovider) | None (embedded) | Provide default fee token prices | [lanes/product.go](../lanes/product.go) |
| [PingPongAdapter](#pingpongadapter) | `PingPongAdapterRegistry` | PingPong demo contract support | [lanes/pingpong.go](../lanes/pingpong.go) |
| [ConfigImporter](#configimporter) | None | Import config from existing deployments | [deploy/product.go](../deploy/product.go) |
| [RampUpdateInRouter](#rampupdateinrouter) | `LaneMigratorRegistry` | Lane migration: update router | [deploy/lanemigrator.go](../deploy/lanemigrator.go) |
| [RouterUpdateInRamp](#routerupdateinramp) | `LaneMigratorRegistry` | Lane migration: update ramps | [deploy/lanemigrator.go](../deploy/lanemigrator.go) |
| [TestAdapter](#testadapter) | `TestAdapterRegistry` | Cross-chain message testing | [testadapters/adapters.go](../testadapters/adapters.go) |

---

## Required Interfaces

### Deployer

Handles deployment of CCIP core contracts and MCMS governance contracts on a chain.

**Source:** [deploy/product.go](../deploy/product.go)
**Registry:** `DeployerRegistry` via `deploy.GetRegistry()`
**Key:** `chainFamily-version`

```go
type Deployer interface {
    // DeployChainContracts deploys all CCIP contracts required for a chain
    // (Router, OnRamp, OffRamp, FeeQuoter, RMNRemote, etc.)
    DeployChainContracts() *Sequence[ContractDeploymentConfigPerChainWithAddress, OnChainOutput, BlockChains]

    // DeployMCMS deploys Multi-Chain Multi-Sig governance contracts
    // (AccessController, MCM, Timelock)
    DeployMCMS() *Sequence[MCMSDeploymentConfigPerChainWithAddress, OnChainOutput, BlockChains]

    // FinalizeDeployMCMS finalizes MCMS deployment (e.g., timelock initialization on Solana).
    // Not all chains require this - can be a no-op sequence.
    FinalizeDeployMCMS() *Sequence[MCMSDeploymentConfigPerChainWithAddress, OnChainOutput, BlockChains]

    // SetOCR3Config sets OCR3 configuration on the chain's OffRamp
    SetOCR3Config() *Sequence[SetOCR3ConfigInput, OnChainOutput, BlockChains]

    // GrantAdminRoleToTimelock grants admin role from one timelock to another
    GrantAdminRoleToTimelock() *Sequence[GrantAdminRoleToTimelockConfigPerChainWithSelector, OnChainOutput, BlockChains]
}
```

**Registration:**
```go
deploy.GetRegistry().RegisterDeployer(chain_selectors.FamilyEVM, semver.MustParse("1.6.0"), &EVMAdapter{})
```

---

### LaneAdapter

Handles lane configuration between chains -- configuring a chain as a message source or destination for a given lane.

**Source:** [lanes/product.go](../lanes/product.go)
**Registry:** `LaneAdapterRegistry` via `lanes.GetLaneAdapterRegistry()`
**Key:** `chainFamily-version`

```go
type LaneAdapter interface {
    // ConfigureLaneLegAsSource configures this chain as the source for a lane.
    // Sets up OnRamp destination chain config, FeeQuoter destination config, and token prices.
    ConfigureLaneLegAsSource() *Sequence[UpdateLanesInput, OnChainOutput, BlockChains]

    // ConfigureLaneLegAsDest configures this chain as the destination for a lane.
    // Sets up OffRamp source chain config and Router integration.
    ConfigureLaneLegAsDest() *Sequence[UpdateLanesInput, OnChainOutput, BlockChains]

    // GetOnRampAddress returns the OnRamp contract address (as bytes) for the given chain.
    // On Solana, this returns the Router address since Solana has a unified contract.
    GetOnRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)

    // GetOffRampAddress returns the OffRamp contract address (as bytes) for the given chain.
    GetOffRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)

    // GetRouterAddress returns the Router contract address (as bytes) for the given chain.
    GetRouterAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)

    // GetFQAddress returns the FeeQuoter contract address (as bytes) for the given chain.
    GetFQAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
}
```

**Registration:**
```go
lanes.GetLaneAdapterRegistry().RegisterLaneAdapter(chain_selectors.FamilySolana, semver.MustParse("1.6.0"), &SolanaAdapter{})
```

**Note:** The `Get*Address` methods are used by the `ConnectChains` changeset to programmatically populate `ChainDefinition` fields. Address bytes must be chain-family encoded (e.g., 20-byte EVM address, 32-byte Solana public key).

---

### TokenAdapter

Handles token pool configuration, deployment, and cross-chain token transfer setup.

**Source:** [tokens/product.go](../tokens/product.go)
**Registry:** `TokenAdapterRegistry` via `tokens.GetTokenAdapterRegistry()`
**Key:** `chainFamily-version`

Each chain-family-version combination registers separately, because configuration differs by token pool version (e.g., 2.0.0 pools require CCV config, 1.5.0 pools require remote pool addresses).

```go
type TokenAdapter interface {
    // ConfigureTokenForTransfersSequence configures a token pool for cross-chain transfers.
    // Assumes the token and pool are already deployed and registered.
    ConfigureTokenForTransfersSequence() *Sequence[ConfigureTokenForTransfersInput, OnChainOutput, BlockChains]

    // AddressRefToBytes converts an AddressRef to a byte slice.
    // Each chain family serializes addresses differently (hex for EVM, base58 for Solana).
    AddressRefToBytes(ref datastore.AddressRef) ([]byte, error)

    // DeriveTokenAddress derives the token address from a token pool reference.
    // Used when the token address is stored on the pool contract.
    DeriveTokenAddress(e Environment, chainSelector uint64, poolRef datastore.AddressRef) (string, error)

    // DeriveTokenDecimals derives the token's decimal count from a pool reference.
    DeriveTokenDecimals(e Environment, chainSelector uint64, poolRef datastore.AddressRef, token []byte) (uint8, error)

    // DeriveTokenPoolCounterpart derives the effective pool address for chains where
    // the deployed address differs from the operational address (e.g., Solana PDAs).
    DeriveTokenPoolCounterpart(e Environment, chainSelector uint64, tokenPool []byte, token []byte) ([]byte, error)

    // ManualRegistration registers a customer token with the token admin registry.
    // Used when the customer has already deployed and no longer has mint authority.
    ManualRegistration() *Sequence[ManualRegistrationSequenceInput, OnChainOutput, BlockChains]

    // SetTokenPoolRateLimits sets rate limits on a token pool.
    SetTokenPoolRateLimits() *Sequence[TPRLRemotes, OnChainOutput, BlockChains]

    // DeployToken deploys a new token on the chain.
    DeployToken() *Sequence[DeployTokenInput, OnChainOutput, BlockChains]

    // DeployTokenVerify validates the DeployToken input before execution.
    DeployTokenVerify(e Environment, in DeployTokenInput) error

    // DeployTokenPoolForToken deploys a token pool for an existing token.
    DeployTokenPoolForToken() *Sequence[DeployTokenPoolInput, OnChainOutput, BlockChains]

    // UpdateAuthorities transfers token and pool ownership to the timelock signer.
    UpdateAuthorities() *Sequence[UpdateAuthoritiesInput, OnChainOutput, *Environment]
}
```

**Registration:**
```go
tokens.GetTokenAdapterRegistry().RegisterTokenAdapter(chain_selectors.FamilyEVM, semver.MustParse("1.6.0"), &EVMAdapter{})
```

---

### FeeAdapter

Handles token transfer fee configuration and retrieval, plus FeeQuoter destination chain config reads/writes.

**Source:** [fees/product.go](../fees/product.go), [fees/defaults.go](../fees/defaults.go)
**Registry:** `FeeAdapterRegistry` via `fees.GetRegistry()`
**Key:** `chainFamily-version` (adapters); `chainFamily` (resolvers)

Adapters take explicit dependencies instead of `Environment`: `Bundle` (logger + context), `BlockChains`, and `DataStore`. Changesets decompose `Environment` at the call site; private helpers may still accept `Environment` and decompose internally.

#### FeeResolver

Infers the on-ramp for a lane from on-chain state (via the router). Registered per chain family in `chains/<family>/deployment/v1_0_0/adapters/fees.go`.

```go
type FeeResolver interface {
    GetOnRampRef(b Bundle, chains BlockChains, ds DataStore, src uint64, dst uint64) (datastore.AddressRef, error)
}
```

**Registration:**
```go
fees.GetRegistry().RegisterFeeResolver(chain_selectors.FamilyEVM, &EVMFeeResolver{})
```

#### FeeAdapter

```go
type FeeAdapter interface {
    // GetFeeContractRef resolves the fee contract (OnRamp for v1.5, FeeQuoter for v1.6+/v2) from an on-ramp ref.
    GetFeeContractRef(b Bundle, chains BlockChains, ds DataStore, onRamp datastore.AddressRef, src uint64, dst uint64) (datastore.AddressRef, error)

    // SetTokenTransferFee returns a sequence that sets per-token transfer fees for each destination chain.
    SetTokenTransferFee(ds DataStore, fq datastore.AddressRef) *Sequence[SetTokenTransferFeeSequenceInput, OnChainOutput, BlockChains]

    // GetOnchainTokenTransferFeeConfig reads the current on-chain fee configuration for a token on a lane.
    GetOnchainTokenTransferFeeConfig(b Bundle, chains BlockChains, fq datastore.AddressRef, src uint64, dst uint64, token string) (TokenTransferFeeArgs, error)

    // GetDefaultTokenTransferFeeConfig returns default fee configuration for a token on a lane.
    GetDefaultTokenTransferFeeConfig(src uint64, dst uint64) TokenTransferFeeArgs

    // ApplyDestChainConfigUpdates returns a sequence that updates FeeQuoter destination chain configs.
    ApplyDestChainConfigUpdates(ds DataStore, fq datastore.AddressRef) *Sequence[ApplyDestChainConfigSequenceInput, OnChainOutput, BlockChains]

    // GetOnchainDestChainConfig reads the current FeeQuoter destination chain config for a lane.
    GetOnchainDestChainConfig(b Bundle, chains BlockChains, fq datastore.AddressRef, src uint64, dst uint64) (FeeQuoterDestChainConfig, error)

    // GetDefaultDestChainConfig returns default destination chain config for a lane.
    GetDefaultDestChainConfig(src, dst uint64) FeeQuoterDestChainConfig
}
```

#### ResolveFeeAdapter

Shared helper in [fees/defaults.go](../fees/defaults.go) that performs the full lane lookup: router → on-ramp → fee contract → adapter (using the **fee contract version**, not the on-ramp version).

```go
func ResolveFeeAdapter(b Bundle, chains BlockChains, ds DataStore, src, dst uint64) (FeeAdapter, datastore.AddressRef, error)
```

**Registration:**
```go
fees.GetRegistry().RegisterFeeAdapter(chain_selectors.FamilySolana, semver.MustParse("1.6.0"), &FeesAdapter{})
```

---

### FeeAggregatorAdapter

Handles setting and reading the fee aggregator address on a chain. The fee aggregator is the address that receives accumulated fees. The specific on-chain mechanism varies by chain family and version:

- **EVM 1.6:** Fee aggregator is stored in the OnRamp's DynamicConfig.
- **EVM 2.0:** Fee aggregator exists on multiple contracts (OnRamp, Proxy, Executor, USDCTokenPoolProxy). The adapter dispatches to the correct on-chain operation based on contract type.
- **Solana 1.6:** Fee aggregator is stored on the Router.

The `SetFeeAggregatorSequenceInput.Contracts` field allows callers to specify exactly which contracts to update using fully-qualified `datastore.AddressRef` values. This avoids ambiguity when the datastore contains multiple versions of the same contract type. When `Contracts` is empty, the adapter falls back to its default contract (Proxy for EVM 2.0, OnRamp for EVM 1.6, Router for Solana 1.6).

**Source:** [fees/fee_aggregator.go](../fees/fee_aggregator.go)
**Registry:** `FeeAggregatorAdapterRegistry` via `fees.GetFeeAggregatorRegistry()`
**Key:** `chainFamily-version`

```go
type FeeAggregatorAdapter interface {
    // SetFeeAggregator returns a sequence that sets the fee aggregator address on a chain.
    SetFeeAggregator(e Environment) *Sequence[SetFeeAggregatorSequenceInput, OnChainOutput, BlockChains]

    // GetFeeAggregator reads the current fee aggregator address from on-chain state.
    GetFeeAggregator(e Environment, chainSelector uint64) (string, error)
}
```

**Registration:**
```go
fees.GetFeeAggregatorRegistry().RegisterFeeAggregatorAdapter(chain_selectors.FamilyEVM, semver.MustParse("1.6.0"), &FeeAggregatorAdapter{})
```

---

### MCMSReader

Resolves MCMS governance metadata for a chain -- timelock addresses, MCMS contract references, and chain metadata needed to build proposals.

**Source:** [utils/changesets/output.go](../utils/changesets/output.go)
**Registry:** `MCMSReaderRegistry` via `changesets.GetRegistry()`
**Key:** `chainFamily` (version-agnostic -- one reader per chain family)

```go
type MCMSReader interface {
    // GetChainMetadata returns MCMS chain metadata (e.g., starting op count, MCM address).
    GetChainMetadata(e Environment, chainSelector uint64, input mcms.Input) (mcms_types.ChainMetadata, error)

    // GetTimelockRef returns the timelock contract AddressRef for a given MCMS input.
    GetTimelockRef(e Environment, chainSelector uint64, input mcms.Input) (datastore.AddressRef, error)

    // GetMCMSRef returns the MCMS contract AddressRef for a given MCMS input.
    GetMCMSRef(e Environment, chainSelector uint64, input mcms.Input) (datastore.AddressRef, error)
}
```

**Registration:**
```go
changesets.GetRegistry().RegisterMCMSReader(chain_selectors.FamilySolana, &SolanaAdapter{})
```

**Note:** Unlike other registries, this one is keyed by chain family only (no version), since MCMS metadata resolution is typically family-wide.

---

### TransferOwnershipAdapter

Handles transferring contract ownership via MCMS governance proposals.

**Source:** [deploy/product.go](../deploy/product.go)
**Registry:** `TransferOwnershipAdapterRegistry` via `deploy.GetTransferOwnershipRegistry()`
**Key:** `chainFamily-version`

```go
type TransferOwnershipAdapter interface {
    // InitializeTimelockAddress resolves and caches the timelock address for use in ownership sequences.
    InitializeTimelockAddress(e Environment, input mcms.Input) error

    // SequenceTransferOwnershipViaMCMS proposes ownership transfer of contracts through MCMS.
    SequenceTransferOwnershipViaMCMS() *Sequence[TransferOwnershipPerChainInput, OnChainOutput, BlockChains]

    // SequenceAcceptOwnership accepts previously proposed ownership transfers.
    SequenceAcceptOwnership() *Sequence[TransferOwnershipPerChainInput, OnChainOutput, BlockChains]

    // ShouldAcceptOwnershipWithTransferOwnership returns true if accept-ownership should be
    // called automatically as part of the transfer-ownership flow (chain-specific behavior).
    ShouldAcceptOwnershipWithTransferOwnership(e Environment, in TransferOwnershipPerChainInput) (bool, error)
}
```

**Registration:**
```go
deploy.GetTransferOwnershipRegistry().RegisterAdapter(chain_selectors.FamilyEVM, semver.MustParse("1.6.0"), &EVMAdapter{})
```

---

### CurseAdapter

Handles RMN (Risk Management Network) curse and uncurse operations on a chain.

**Source:** [fastcurse/product.go](../fastcurse/product.go)
**Registry:** `CurseRegistry` via `fastcurse.GetCurseRegistry()`
**Key:** `chainFamily-version`

```go
type CurseAdapter interface {
    // Initialize sets up the adapter state for a given chain (e.g., loads RMN contract addresses).
    Initialize(e Environment, selector uint64) error

    // IsSubjectCursedOnChain returns true if the given subject is cursed on the chain.
    // Does NOT follow EVM RMN behavior of returning true for global curse.
    // Use GlobalCurseSubject() to check global curse state.
    IsSubjectCursedOnChain(e Environment, selector uint64, subject Subject) (bool, error)

    // IsChainConnectedToTargetChain returns true if the chain is connected to the target chain.
    // E.g., on EVM, checks if router.isChainSupported(targetSel) returns true.
    IsChainConnectedToTargetChain(e Environment, selector uint64, targetSel uint64) (bool, error)

    // IsCurseEnabledForChain returns true if the chain supports cursing
    // (e.g., RMNRemote contract is deployed).
    IsCurseEnabledForChain(e Environment, selector uint64) (bool, error)

    // SubjectToSelector converts a Subject to a chain selector.
    SubjectToSelector(subject Subject) (uint64, error)

    // Curse returns a sequence that curses the given subjects on a chain.
    Curse() *Sequence[CurseInput, OnChainOutput, BlockChains]

    // Uncurse returns a sequence that lifts curses on the given subjects.
    Uncurse() *Sequence[CurseInput, OnChainOutput, BlockChains]

    // ListConnectedChains returns all chain selectors connected to this chain.
    // Used to determine which chains need to curse subjects derived from a given selector.
    ListConnectedChains(e Environment, selector uint64) ([]uint64, error)
}
```

---

### CurseSubjectAdapter

Maps between chain selectors and curse subjects, and derives the correct curse adapter version for a chain.

**Source:** [fastcurse/product.go](../fastcurse/product.go)
**Registry:** `CurseRegistry` via `fastcurse.GetCurseRegistry()`
**Key:** `chainFamily` (version-agnostic for subject mapping)

```go
type CurseSubjectAdapter interface {
    // SelectorToSubject converts a chain selector to a curse Subject.
    SelectorToSubject(selector uint64) Subject

    // DeriveCurseAdapterVersion derives which version of the curse adapter to use for a chain.
    // E.g., for EVM, this could check which RMN version is deployed on the chain.
    DeriveCurseAdapterVersion(e Environment, selector uint64) (*semver.Version, error)
}
```

**Registration (both adapters together):**
```go
fastcurse.GetCurseRegistry().RegisterNewCurse(fastcurse.CurseRegistryInput{
    CursingFamily:       chain_selectors.FamilyEVM,
    CursingVersion:      semver.MustParse("1.6.0"),
    CurseAdapter:        NewCurseAdapter(),
    CurseSubjectAdapter: NewCurseAdapter(),
})
```

---

## Optional Interfaces

### TokenRefResolver

Reconstructs `datastore.AddressRef` values for tokens and token pools when a changeset input only provides a partial ref (for example, an address string) or when the ref is not yet present in the environment datastore. Used by `ResolveTokenPoolRef`, `ResolveTokenRef`, and `ResolveAdapterAndRefs` in [tokens/token_expansion.go](../tokens/token_expansion.go).

**Source:** [tokens/product.go](../tokens/product.go)
**Registry:** `TokenAdapterRegistry` via `tokens.GetTokenAdapterRegistry()`
**Key:** `chainFamily` only (version-agnostic — one resolver per chain family)

```go
type TokenRefResolver interface {
    // ResolveTokenPoolRef reconstructs a token pool AddressRef from an on-chain address.
    // Implementations typically read type/version (and related metadata) from chain state.
    ResolveTokenPoolRef(b Bundle, chains BlockChains, ds DataStore, chainSelector uint64, address string) (AddressRef, error)

    // ResolveTokenRef reconstructs a token AddressRef from an on-chain address (for example, a mint or ERC20).
    ResolveTokenRef(b Bundle, chains BlockChains, ds DataStore, chainSelector uint64, address string) (AddressRef, error)
}
```

**Registration** (alongside `RegisterTokenAdapter` in the same `init()`):

```go
tokensapi.GetTokenAdapterRegistry().RegisterTokenRefResolver(chain_selectors.FamilyEVM, &EVMTokenBase{})
tokensapi.GetTokenAdapterRegistry().RegisterTokenRefResolver(chain_selectors.FamilySolana, &SolanaAdapter{})
```

First registration wins; a second `RegisterTokenRefResolver` for the same family is ignored.

**Resolution order** (package helpers, not part of the interface):

1. `TryNormalizeAddressRef` canonicalizes `ref.Address` when set (via `deploy.AddressNormalizer` for the chain family).
2. Datastore lookup with `AddressRefToFilters` — exactly one match returns that row; zero matches fall through to the resolver; more than one match is an error.
3. If no datastore row and `address` is set, `TokenRefResolver.ResolveTokenPoolRef` / `ResolveTokenRef` rebuild the ref from chain state.

**`ResolveAdapterAndRefs`** (used by token expansion, configure-for-transfers, rate limits):

1. `ResolveTokenPoolRef` → pick `TokenAdapter` from the resolved pool’s version.
2. `DeriveTokenAddress` on the resolved pool ref when the token address is stored on the pool (preferred).
3. If derivation fails, `ResolveTokenRef` on the input token ref (datastore, then resolver).

**Semantics by chain family:**

- **EVM:** Resolver reads pool `typeAndVersion` and token `symbol` via RPC. Reconstructed pool refs use the token address as `Qualifier` when `getToken()` succeeds.
- **Solana:** `ResolveTokenPoolRef` accepts either the **pool program ID** or the **pool config PDA**. For a PDA, the adapter loads the program ID from chain state, looks up the program in the datastore, and may attach an [`ArtificialAddressRefLabel`](../tokens/product.go) label so `DeriveTokenAddress` can read the mint from that PDA. `ResolveTokenRef` reads the mint’s token program and symbol (or a `{mint}-{programType}` placeholder qualifier).
- **`ArtificialAddressRefLabel`:** Optional label format `ArtificialAddressRef:<poolPDA>` on pool refs. Not required for all adapters; Solana uses it when resolving from a config PDA.

Implement `TokenRefResolver` on the same struct as `TokenAdapter` when the family should support address-only inputs (EVM and Solana in this repo register both on one adapter type).

---

### TokenPriceProvider

An optional interface that `LaneAdapter` implementations can also satisfy to provide default fee token prices. Primarily used by EVM chains.

**Source:** [lanes/product.go](../lanes/product.go)
**Registry:** None -- checked via Go type assertion on the `LaneAdapter` instance.

```go
type TokenPriceProvider interface {
    // GetDefaultTokenPrices returns default fee token prices for a chain.
    // Returns a map of contract type (e.g., "WETH", "LINK") to USD price (18 decimals).
    GetDefaultTokenPrices() map[datastore.ContractType]*big.Int
}
```

---

### PingPongAdapter

Supports the PingPong demo contract for testing lane connectivity. Chains that do not support PingPong (e.g., Solana) should not implement this.

**Source:** [lanes/pingpong.go](../lanes/pingpong.go)
**Registry:** `PingPongAdapterRegistry` via `lanes.GetPingPongAdapterRegistry()`
**Key:** `chainFamily-version`

```go
type PingPongAdapter interface {
    // GetPingPongDemoAddress returns the PingPongDemo contract address for the given chain.
    GetPingPongDemoAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)

    // ConfigurePingPong configures PingPong for a lane between source and dest.
    ConfigurePingPong() *Sequence[PingPongInput, PingPongOutput, BlockChains]
}
```

---

### ConfigImporter

Imports configuration from existing deployments to bootstrap the DataStore.

**Source:** [deploy/product.go](../deploy/product.go)
**Registry:** None currently.

```go
type ConfigImporter interface {
    // InitializeAdapter sets up the importer for the given chain selectors.
    InitializeAdapter(e Environment, selectors []uint64) error

    // ConnectedChains returns the chain selectors connected to a given chain.
    ConnectedChains(e Environment, chainsel uint64) ([]uint64, error)

    // SupportedTokensPerRemoteChain returns supported tokens per remote chain.
    SupportedTokensPerRemoteChain(e Environment, chainSelector uint64) (map[uint64][]common.Address, error)

    // SequenceImportConfig returns a sequence to import lane config from on-chain state.
    SequenceImportConfig() *Sequence[ImportConfigPerChainInput, OnChainOutput, BlockChains]
}
```

---

### RampUpdateInRouter

Updates router configuration for lane migration scenarios (pointing routers to new ramps).

**Source:** [deploy/lanemigrator.go](../deploy/lanemigrator.go)
**Registry:** `LaneMigratorRegistry` via `deploy.GetLaneMigratorRegistry()` (as RouterUpdater)
**Key:** `chainFamily-version`

```go
type RampUpdateInRouter interface {
    // UpdateRouter updates the router to point to new OnRamp/OffRamp contracts for remote chains.
    UpdateRouter() *Sequence[RouterUpdaterConfig, OnChainOutput, BlockChains]
}
```

---

### RouterUpdateInRamp

Updates ramp configuration with new router addresses for lane migration scenarios.

**Source:** [deploy/lanemigrator.go](../deploy/lanemigrator.go)
**Registry:** `LaneMigratorRegistry` via `deploy.GetLaneMigratorRegistry()` (as RampUpdater)
**Key:** `chainFamily-version`

```go
type RouterUpdateInRamp interface {
    // UpdateVersionWithRouter updates OnRamp/OffRamp contracts with a new router address.
    UpdateVersionWithRouter() *Sequence[RampUpdaterConfig, OnChainOutput, BlockChains]
}
```

---

### TestAdapter

Interface for integration testing of cross-chain message passing. Each adapter instance represents a concrete chain.

**Source:** [testadapters/adapters.go](../testadapters/adapters.go)
**Registry:** `TestAdapterRegistry` via `testadapters.GetTestAdapterRegistry()`
**Key:** `chainFamily-version`

```go
type TestAdapter interface {
    // ChainSelector returns the selector of the chain for this adapter.
    ChainSelector() uint64

    // Family returns the chain family string (e.g., "evm", "solana").
    Family() string

    // BuildMessage builds a chain-family-specific message from generic components.
    // E.g., EVM produces router.ClientEVM2AnyMessage, Solana produces ccip_router.SVM2AnyMessage.
    BuildMessage(components MessageComponents) (any, error)

    // SendMessage sends a CCIP message and returns the sequence number.
    SendMessage(ctx context.Context, destChainSelector uint64, msg any) (uint64, error)

    // CCIPReceiver returns the address of a CCIP receiver contract on this chain.
    CCIPReceiver() []byte

    // NativeFeeToken returns the native fee token identifier for this chain.
    NativeFeeToken() string

    // GetExtraArgs returns encoded extra args for sending to this chain from a given source family.
    // Extra args are source-family encoded (abi.encode for EVM, borsh for Solana).
    GetExtraArgs(receiver []byte, sourceFamily string, opts ...ExtraArgOpt) ([]byte, error)

    // GetInboundNonce returns the inbound nonce for a sender from a source chain.
    // Returns 0 for chains without nonce concepts.
    GetInboundNonce(ctx context.Context, sender []byte, srcSel uint64) (uint64, error)

    // ValidateCommit validates that a message was committed on this chain.
    ValidateCommit(t *testing.T, sourceSelector uint64, startBlock *uint64, seqNumRange ccipocr3.SeqNumRange)

    // ValidateExec validates that a message was executed on this chain and returns execution states.
    ValidateExec(t *testing.T, sourceSelector uint64, startBlock *uint64, seqNrs []uint64) (execStates map[uint64]int)

    // AllowRouterToWithdrawTokens approves the router to spend tokens from the deployer.
    AllowRouterToWithdrawTokens(ctx context.Context, tokenAddress string, amount *big.Int) error

    // GetTokenBalance returns the token balance for a given owner address.
    GetTokenBalance(ctx context.Context, tokenAddress string, ownerAddress []byte) (*big.Int, error)

    // GetTokenExpansionConfig returns default token expansion config for testing.
    GetTokenExpansionConfig() TokenExpansionInputPerChain

    // GetRegistryAddress returns the address of the token admin registry contract.
    GetRegistryAddress() (string, error)
}
```

**Note:** `TestAdapter` is instantiated via a factory function:
```go
type TestAdapterFactory = func(env *Environment, selector uint64) TestAdapter
```

---

## Registry Accessor Summary

| Registry | Accessor | Key Format |
|----------|----------|------------|
| `DeployerRegistry` | `deploy.GetRegistry()` | `chainFamily-version` |
| `TransferOwnershipAdapterRegistry` | `deploy.GetTransferOwnershipRegistry()` | `chainFamily-version` |
| `LaneAdapterRegistry` | `lanes.GetLaneAdapterRegistry()` | `chainFamily-version` |
| `PingPongAdapterRegistry` | `lanes.GetPingPongAdapterRegistry()` | `chainFamily-version` |
| `TokenAdapterRegistry` | `tokens.GetTokenAdapterRegistry()` | `chainFamily-version` (adapters); `chainFamily` (`TokenRefResolver`) |
| `FeeAdapterRegistry` | `fees.GetRegistry()` | `chainFamily-version` |
| `FeeAggregatorAdapterRegistry` | `fees.GetFeeAggregatorRegistry()` | `chainFamily-version` |
| `MCMSReaderRegistry` | `changesets.GetRegistry()` | `chainFamily` |
| `CurseRegistry` | `fastcurse.GetCurseRegistry()` | `chainFamily-version` / `chainFamily` |
| `LaneMigratorRegistry` | `deploy.GetLaneMigratorRegistry()` | `chainFamily-version` |
| `TestAdapterRegistry` | `testadapters.GetTestAdapterRegistry()` | `chainFamily-version` |
