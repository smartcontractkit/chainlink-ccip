---
title: "Changesets Reference"
sidebar_label: "Changesets"
sidebar_position: 5
---

# Changesets Reference

Changesets are the top-level entry points for the CCIP Deployment Tooling API. Each changeset accepts a configuration input, resolves the appropriate chain-family adapters, executes sequences, and returns a `ChangesetOutput` that may include DataStore updates and MCMS proposals.

For the types used by these changesets, see [Types Reference](types.md). For the adapter interfaces they delegate to, see [Interfaces Reference](interfaces.md).

---

## Table of Contents

- [Contract Deployment](#contract-deployment)
- [MCMS Deployment](#mcms-deployment)
- [Lane Configuration](#lane-configuration)
- [OCR3 Configuration](#ocr3-configuration)
- [Token Configuration](#token-configuration)
- [Fee Configuration](#fee-configuration)
- [Ownership Management](#ownership-management)
- [RMN Curse Operations](#rmn-curse-operations)
- [Lane Migration](#lane-migration)

---

## Contract Deployment

### DeployContracts

Deploys all CCIP infrastructure contracts (Router, OnRamp, OffRamp, FeeQuoter, RMNRemote, etc.) on one or more chains.

**Source:** [deploy/contracts.go](../deploy/contracts.go)

**Constructor:**
```go
func DeployContracts(deployerReg *DeployerRegistry) cldf.ChangeSetV2[ContractDeploymentConfig]
```

**Input:** [`ContractDeploymentConfig`](types.md#contractdeploymentconfig)

**Behavior:**
1. For each chain in `Chains`, resolves the chain family from the selector
2. Looks up the `Deployer` adapter from the `DeployerRegistry` using family + version
3. Executes `deployer.DeployChainContracts()` sequence
4. Collects deployed addresses into a DataStore
5. Returns output via `OutputBuilder`

**Example:**
```go
changeset := deploy.DeployContracts(deploy.GetRegistry())
output, err := changeset.Apply(env, deploy.ContractDeploymentConfig{
    Chains: map[uint64]deploy.ContractDeploymentConfigPerChain{
        chainSelector: {
            Version:                                 semver.MustParse("1.6.0"),
            MaxFeeJuelsPerMsg:                       big.NewInt(1e18),
            TokenPriceStalenessThreshold:            86400,
            LinkPremiumMultiplier:                    9e17,
            NativeTokenPremiumMultiplier:             18e17,
            PermissionLessExecutionThresholdSeconds:  86400,
        },
    },
})
```

---

## MCMS Deployment

### DeployMCMS

Deploys Multi-Chain Multi-Sig governance contracts (AccessController, MCM, Timelock) on one or more chains.

**Source:** [deploy/mcms.go](../deploy/mcms.go)

**Constructor:**
```go
func DeployMCMS(deployerReg *DeployerRegistry, mcmsRegistry *MCMSReaderRegistry) cldf.ChangeSetV2[MCMSDeploymentConfig]
```

**Input:** [`MCMSDeploymentConfig`](types.md#mcmsdeploymentconfig)

**Behavior:**
1. For each chain, resolves family and looks up `Deployer` adapter
2. Executes `deployer.DeployMCMS()` sequence
3. Collects deployed MCMS contract addresses into a DataStore

### FinalizeDeployMCMS

Finalizes MCMS deployment. Required for chains that need a two-phase deployment (e.g., Solana timelock initialization).

**Constructor:**
```go
func FinalizeDeployMCMS(deployerReg *DeployerRegistry, mcmsRegistry *MCMSReaderRegistry) cldf.ChangeSetV2[MCMSDeploymentConfig]
```

Same input as `DeployMCMS`. Executes `deployer.FinalizeDeployMCMS()` instead.

### GrantAdminRoleToTimelock

Grants admin role from one timelock to another. Used when setting up multi-tiered MCMS governance.

**Constructor:**
```go
func GrantAdminRoleToTimelock(deployerReg *DeployerRegistry, mcmsRegistry *MCMSReaderRegistry) cldf.ChangeSetV2[GrantAdminRoleToTimelockConfig]
```

**Input:** [`GrantAdminRoleToTimelockConfig`](types.md#grantadminroletotimelockconfig)

**Validation:**
- Both timelock refs must have type `RBACTimelock`
- Both must have non-empty qualifier and version

---

## Lane Configuration

### ConnectChains

Configures bidirectional lanes between chains. For each lane, configures both the source-side and destination-side on each chain.

**Source:** [lanes/connect_chains.go](../lanes/connect_chains.go)

**Constructor:**
```go
func ConnectChains(laneRegistry *LaneAdapterRegistry, mcmsRegistry *MCMSReaderRegistry) cldf.ChangeSetV2[ConnectChainsConfig]
```

**Input:** [`ConnectChainsConfig`](types.md#connectchainsconfig)

**Behavior:**
For each `LaneConfig` in the input:
1. Resolves `LaneAdapter` for both ChainA and ChainB families
2. Populates contract addresses (OnRamp, OffRamp, Router, FeeQuoter) from the DataStore via the adapter's `Get*Address` methods
3. Executes `ConfigureLaneLegAsSource()` and `ConfigureLaneLegAsDest()` for both directions:
   - A as source, B as dest
   - B as source, A as dest

**Example:**
```go
changeset := lanes.ConnectChains(lanes.GetLaneAdapterRegistry(), changesets.GetRegistry())
output, err := changeset.Apply(env, lanes.ConnectChainsConfig{
    Lanes: []lanes.LaneConfig{
        {
            ChainA:  lanes.ChainDefinition{Selector: evmSel, GasPrice: big.NewInt(2e12)},
            ChainB:  lanes.ChainDefinition{Selector: solSel, GasPrice: big.NewInt(2e12)},
            Version: semver.MustParse("1.6.0"),
        },
    },
    MCMS: mcms.Input{TimelockAction: mcms_types.TimelockActionSchedule},
})
```

---

## OCR3 Configuration

### SetOCR3Config

Sets OCR3 configuration on remote chain OffRamps. Reads the OCR3 config from CCIPHome on the home chain and applies it to the specified remote chains.

**Source:** [deploy/set_ocr3_config.go](../deploy/set_ocr3_config.go)

**Constructor:**
```go
func SetOCR3Config(deployerReg *DeployerRegistry, mcmsReg *MCMSReaderRegistry) cldf.ChangeSetV2[SetOCR3ConfigArgs]
```

**Input:** [`SetOCR3ConfigArgs`](types.md#setocr3configargs)

**Behavior:**
1. For each remote chain selector, resolves the `Deployer` adapter (always version 1.6.0)
2. Reads OCR3 config from CCIPHome on the home chain
3. Builds `OCR3ConfigArgs` with signers, transmitters, and config digests
4. Executes `deployer.SetOCR3Config()` on each remote chain

---

## Token Configuration

### TokenExpansion

Comprehensive changeset that deploys tokens, token pools, configures them for cross-chain transfers, and transfers ownership to the timelock.

**Source:** [tokens/token_expansion.go](../tokens/token_expansion.go)

**Constructor:**
```go
func TokenExpansion() cldf.ChangeSetV2[TokenExpansionInput]
```

**Input:** [`TokenExpansionInput`](types.md#tokenexpansioninput)

**Behavior:**
1. **Deploy tokens** (if `DeployTokenInput` is non-nil): Executes `adapter.DeployToken()`
2. **Deploy token pools** (if `DeployTokenPoolInput` is non-nil): Executes `adapter.DeployTokenPoolForToken()`
3. **Configure for transfers** (if `TokenTransferConfig` is non-nil): Calls `processTokenConfigForChain()` which executes `adapter.ConfigureTokenForTransfersSequence()`
4. **Update authorities** (unless `SkipOwnershipTransfer` is true): Executes `adapter.UpdateAuthorities()` to transfer ownership to the timelock

### ConfigureTokensForTransfers

Configures already-deployed tokens and token pools for cross-chain transfers. Does not deploy new contracts.

**Source:** [tokens/configure_tokens_for_transfers.go](../tokens/configure_tokens_for_transfers.go)

**Constructor:**
```go
func ConfigureTokensForTransfers(tokenRegistry *TokenAdapterRegistry, mcmsRegistry *MCMSReaderRegistry) cldf.ChangeSetV2[ConfigureTokensForTransfersConfig]
```

**Input:** [`ConfigureTokensForTransfersConfig`](types.md#configuretokensfortransfersconfig)

**Behavior:**
For each token configuration:
1. Resolves the `TokenAdapter` for the chain family
2. Resolves remote chain addresses and decimals
3. Derives inbound rate limiter config from counterpart's outbound config
4. Executes `adapter.ConfigureTokenForTransfersSequence()`

### ManualRegistration

Registers customer tokens with the token admin registry. Used when the customer has already deployed the token and no longer has mint authority.

**Source:** [tokens/manual_registration.go](../tokens/manual_registration.go)

**Constructor:**
```go
func ManualRegistration() cldf.ChangeSetV2[ManualRegistrationInput]
```

**Input:** [`ManualRegistrationInput`](types.md#manualregistrationinput)

### SetTokenPoolRateLimits

Sets rate limits on token pools for specific remote chains.

**Source:** [tokens/rate_limits.go](../tokens/rate_limits.go)

**Constructor:**
```go
func SetTokenPoolRateLimits() cldf.ChangeSetV2[TPRLInput]
```

**Input:** [`TPRLInput`](types.md#tprlinput)

**Behavior:**
- Resolves token decimals on both local and remote chains
- Scales float-based user inputs by token decimals
- Derives inbound rate limiter from counterpart's outbound config
- Inbound capacity is scaled by 1.1x to avoid accidental rate limit hits

**Validation:**
- If rate limiter is enabled, capacity and rate must be positive

---

## Fee Configuration

### SetTokenTransferFee

Sets per-token transfer fee configurations on the FeeQuoter for each source-destination pair.

**Source:** [fees/set_token_transfer_fee.go](../fees/set_token_transfer_fee.go)

**Constructor:**
```go
func SetTokenTransferFee(feeRegistry *FeeAdapterRegistry, mcmsRegistry *MCMSReaderRegistry) cldf.ChangeSetV2[SetTokenTransferFeeInput]
```

**Input:**
```go
type SetTokenTransferFeeInput struct {
    Version *semver.Version
    Args    []TokenTransferFeeForSrc
    MCMS    mcms.Input
}
```

**Fee Resolution:**
For each token, the changeset uses a three-tier fallback:
1. **User-specified values** (if `Valid` is true in `UnresolvedTokenTransferFeeArgs`)
2. **On-chain values** (if the token already has fee config set)
3. **Adapter defaults** (from `adapter.GetDefaultTokenTransferFeeConfig()`)

**Validation:**
- No duplicate source chain selectors
- No duplicate destination selectors within the same source
- Source and destination selectors cannot be the same
- Token addresses cannot be empty
- Same address cannot appear in both updates and resets

---

## Ownership Management

### TransferOwnershipChangeset

Proposes ownership transfer of contracts through MCMS governance. On chains where `ShouldAcceptOwnershipWithTransferOwnership()` returns true (e.g., Solana), also accepts ownership atomically.

**Source:** [deploy/transfer_ownership.go](../deploy/transfer_ownership.go)

**Constructor:**
```go
func TransferOwnershipChangeset(cr *TransferOwnershipAdapterRegistry, mcmsRegistry *MCMSReaderRegistry) cldf.ChangeSetV2[TransferOwnershipInput]
```

**Input:** [`TransferOwnershipInput`](types.md#transferownershipinput)

**Behavior:**
1. For each chain, resolves the `TransferOwnershipAdapter`
2. Initializes timelock address via `adapter.InitializeTimelockAddress()`
3. Resolves partial contract refs to full refs
4. Executes `adapter.SequenceTransferOwnershipViaMCMS()`
5. If `adapter.ShouldAcceptOwnershipWithTransferOwnership()` returns true, also executes `adapter.SequenceAcceptOwnership()`

### AcceptOwnershipChangeset

Accepts previously proposed ownership transfers.

**Constructor:**
```go
func AcceptOwnershipChangeset(cr *TransferOwnershipAdapterRegistry, mcmsRegistry *MCMSReaderRegistry) cldf.ChangeSetV2[TransferOwnershipInput]
```

Same input as `TransferOwnershipChangeset`. Only executes `adapter.SequenceAcceptOwnership()`.

---

## RMN Curse Operations

**Source:** [fastcurse/fastcurse.go](../fastcurse/fastcurse.go)

### CurseChangeset

Curses specified subjects on specified chains via the RMNRemote contract.

**Constructor:**
```go
func CurseChangeset(cr *CurseRegistry, mcmsRegistry *MCMSReaderRegistry) cldf.ChangeSetV2[RMNCurseConfig]
```

**Input:**
```go
type RMNCurseConfig struct {
    CurseActions []CurseActionInput
    Force        bool        // Include already-cursed subjects
    MCMS         mcms.Input
}

type CurseActionInput struct {
    IsGlobalCurse        bool
    ChainSelector        uint64
    SubjectChainSelector uint64
    Version              *semver.Version
}
```

**Behavior:**
- Groups curse actions by chain selector
- Filters out already-cursed subjects (unless `Force` is true)
- Executes `adapter.Curse()` sequence for each chain

### UncurseChangeset

Lifts curses on specified subjects.

**Constructor:**
```go
func UncurseChangeset(cr *CurseRegistry, mcmsRegistry *MCMSReaderRegistry) cldf.ChangeSetV2[RMNCurseConfig]
```

Same input as `CurseChangeset`. Executes `adapter.Uncurse()` instead.

### GloballyCurseChainChangeset

Globally curses chains across the entire network. Automatically discovers connected chains and creates curse actions for all of them.

**Constructor:**
```go
func GloballyCurseChainChangeset(cr *CurseRegistry, mcmsRegistry *MCMSReaderRegistry) cldf.ChangeSetV2[GlobalCurseOnNetworkInput]
```

**Input:**
```go
type GlobalCurseOnNetworkInput struct {
    ChainSelectors map[uint64]*semver.Version  // Chains to curse -> RMN version
    Force          bool
    MCMS           mcms.Input
}
```

### GloballyUncurseChainChangeset

Reverses a global curse across the network.

**Constructor:**
```go
func GloballyUncurseChainChangeset(cr *CurseRegistry, mcmsRegistry *MCMSReaderRegistry) cldf.ChangeSetV2[GlobalCurseOnNetworkInput]
```

---

## Lane Migration

### LaneMigrateToNewVersionChangeset

Migrates lanes to a new contract version by updating router configurations and ramp contracts.

**Source:** [deploy/lanemigrator.go](../deploy/lanemigrator.go)

**Constructor:**
```go
func LaneMigrateToNewVersionChangeset(migratorReg *LaneMigratorRegistry, mcmsRegistry *MCMSReaderRegistry) cldf.ChangeSetV2[LaneMigratorConfig]
```

**Input:**
```go
type LaneMigratorConfig struct {
    Input map[uint64]LaneMigratorConfigPerChain
    MCMS  mcms.Input
}

type LaneMigratorConfigPerChain struct {
    RemoteChains  []uint64           // Remote chains to migrate
    RouterVersion *semver.Version    // Router version to use
    RampVersion   *semver.Version    // Ramp version to upgrade to (must be >= 1.6.0)
}
```

**Behavior:**
1. Looks up OnRamp, OffRamp, and Router refs from the DataStore
2. Executes `routerUpdater.UpdateRouter()` to point the router to new ramps
3. Executes `rampUpdater.UpdateVersionWithRouter()` to configure ramps with the new router

**Validation:**
- Ramp version must be >= 1.6.0
- Both router updater and ramp updater must be registered
- All chain selectors must exist in the environment
- Existing addresses must be present in the DataStore
