---
title: "Types Reference"
sidebar_label: "Types"
sidebar_position: 4
---

# Types Reference

This document provides a complete reference for all input/output types used across the CCIP Deployment Tooling API.

For the interfaces that use these types, see [Interfaces Reference](interfaces.md). For the changesets that accept these types as input, see [Changesets Reference](changesets.md).

---

## Table of Contents

- [Deployment Types](#deployment-types)
- [MCMS Deployment Types](#mcms-deployment-types)
- [Lane Types](#lane-types)
- [OCR3 Types](#ocr3-types)
- [Token Types](#token-types)
- [Fee Types](#fee-types)
- [Ownership Types](#ownership-types)
- [MCMS Input Types](#mcms-input-types)
- [Output Types](#output-types)
- [Curse Types](#curse-types)
- [Test Adapter Types](#test-adapter-types)
- [Constants](#constants)

---

## Deployment Types

**Source:** [deploy/contracts.go](../deploy/contracts.go)

### ContractDeploymentConfig

Top-level input for the `DeployContracts` changeset.

```go
type ContractDeploymentConfig struct {
    Chains map[uint64]ContractDeploymentConfigPerChain
    MCMS   mcms.Input
}
```

### ContractDeploymentConfigPerChain

Per-chain configuration for deploying CCIP contracts.

```go
type ContractDeploymentConfigPerChain struct {
    Version                                 *semver.Version
    TokenPrivKey                            string    // Private key for LINK token deployment (Solana: base58)
    TokenDecimals                           uint8     // LINK token decimals
    MaxFeeJuelsPerMsg                       *big.Int  // FeeQuoter: max fee per message
    TokenPriceStalenessThreshold            uint32    // FeeQuoter: staleness threshold
    LinkPremiumMultiplier                   uint64    // FeeQuoter: LINK premium (Wei per ETH on EVM)
    NativeTokenPremiumMultiplier            uint64    // FeeQuoter: native token premium (Wei per ETH on EVM)
    PermissionLessExecutionThresholdSeconds uint32    // OffRamp: threshold for manual execution
    GasForCallExactCheck                    uint16    // OffRamp: EVM only
    MessageInterceptor                      string    // OffRamp: EVM only, validates incoming messages
    LegacyRMN                               string    // RMN Remote config
    ContractVersion                         string    // Contract version string
    DeployPingPongDapp                      bool      // Deploy PingPongDemo contract
}
```

### ContractDeploymentConfigPerChainWithAddress

Extends `ContractDeploymentConfigPerChain` with chain selector and existing addresses (populated by the framework).

```go
type ContractDeploymentConfigPerChainWithAddress struct {
    ContractDeploymentConfigPerChain
    ChainSelector     uint64
    ExistingAddresses []datastore.AddressRef
}
```

---

## MCMS Deployment Types

**Source:** [deploy/mcms.go](../deploy/mcms.go)

### MCMSDeploymentConfig

Top-level input for `DeployMCMS` and `FinalizeDeployMCMS` changesets.

```go
type MCMSDeploymentConfig struct {
    Chains         map[uint64]MCMSDeploymentConfigPerChain
    AdapterVersion *semver.Version
    MCMS           mcms.Input
}
```

### MCMSDeploymentConfigPerChain

Per-chain MCMS configuration specifying multi-sig roles and timelock settings.

```go
type MCMSDeploymentConfigPerChain struct {
    Canceller        mcmstypes.Config   // MCM canceller role configuration
    Bypasser         mcmstypes.Config   // MCM bypasser role configuration
    Proposer         mcmstypes.Config   // MCM proposer role configuration
    TimelockMinDelay *big.Int           // Minimum delay for timelock operations
    Label            *string            // Optional label for the MCMS instance
    Qualifier        *string            // Optional qualifier for the MCMS instance
    TimelockAdmin    common.Address     // Admin address for the timelock
    ContractVersion  string             // Contract version string
}
```

### MCMSDeploymentConfigPerChainWithAddress

Extends `MCMSDeploymentConfigPerChain` with chain selector and existing addresses.

```go
type MCMSDeploymentConfigPerChainWithAddress struct {
    MCMSDeploymentConfigPerChain
    ChainSelector     uint64
    ExistingAddresses []datastore.AddressRef
}
```

### GrantAdminRoleToTimelockConfig

Input for the `GrantAdminRoleToTimelock` changeset.

```go
type GrantAdminRoleToTimelockConfig struct {
    Chains         map[uint64]GrantAdminRoleToTimelockConfigPerChain
    AdapterVersion *semver.Version
}
```

### GrantAdminRoleToTimelockConfigPerChain

Per-chain config specifying which timelock transfers admin rights and which receives them.

```go
type GrantAdminRoleToTimelockConfigPerChain struct {
    TimelockToTransferRef datastore.AddressRef  // Timelock that transfers its admin rights
    NewAdminTimelockRef   datastore.AddressRef  // Timelock that will be granted admin
}
```

---

## Lane Types

**Source:** [lanes/lane_update.go](../lanes/lane_update.go)

### ConnectChainsConfig

Top-level input for the `ConnectChains` changeset. Configures bidirectional lanes between chains.

```go
type ConnectChainsConfig struct {
    Lanes []LaneConfig
    MCMS  mcms.Input
}
```

### LaneConfig

Defines a single bidirectional lane between two chains.

```go
type LaneConfig struct {
    ChainA       ChainDefinition
    ChainB       ChainDefinition
    Version      *semver.Version
    IsDisabled   bool
    TestRouter   bool          // Use test router instead of production
    ExtraConfigs ExtraConfigs
}
```

### ChainDefinition

Complete definition of a chain's role in a lane. Some fields are user-provided, others are populated programmatically.

```go
type ChainDefinition struct {
    // User-provided fields
    Selector                   uint64                       // Chain selector
    GasPrice                   *big.Int                     // USD price (18 decimals) per unit gas
    TokenPrices                map[string]*big.Int          // Token USD prices (18 decimals)
    FeeQuoterDestChainConfig   FeeQuoterDestChainConfig     // Fee config when this chain is a destination
    RMNVerificationEnabled     bool                         // RMN blessing for messages FROM this chain
    AllowListEnabled           bool                         // Allowlist for messages TO this chain
    AllowList                  []string                     // Allowed sender addresses

    // Populated programmatically (do not set)
    OnRamp    []byte  // OnRamp contract address
    OffRamp   []byte  // OffRamp contract address
    Router    []byte  // Router contract address
    FeeQuoter []byte  // FeeQuoter contract address
}
```

### FeeQuoterDestChainConfig

Fee configuration applied on a source chain when the target chain is a destination.

```go
type FeeQuoterDestChainConfig struct {
    IsEnabled                         bool
    MaxNumberOfTokensPerMsg           uint16
    MaxDataBytes                      uint32
    MaxPerMsgGasLimit                 uint32
    DestGasOverhead                   uint32
    DestGasPerPayloadByteBase         uint8
    DestGasPerPayloadByteHigh         uint8
    DestGasPerPayloadByteThreshold    uint16
    DestDataAvailabilityOverheadGas   uint32
    DestGasPerDataAvailabilityByte    uint16
    DestDataAvailabilityMultiplierBps uint16
    ChainFamilySelector               uint32
    EnforceOutOfOrder                 bool
    DefaultTokenFeeUSDCents           uint16
    DefaultTokenDestGasOverhead       uint32
    DefaultTxGasLimit                 uint32
    GasMultiplierWeiPerEth            uint64
    GasPriceStalenessThreshold        uint32
    NetworkFeeUSDCents                uint32
}
```

A default configuration is available via `DefaultFeeQuoterDestChainConfig(configEnabled bool, selector uint64)`.

### UpdateLanesInput

Input passed to `LaneAdapter.ConfigureLaneLegAsSource()` and `ConfigureLaneLegAsDest()` sequences.

```go
type UpdateLanesInput struct {
    Source       *ChainDefinition
    Dest         *ChainDefinition
    IsDisabled   bool
    TestRouter   bool
    ExtraConfigs ExtraConfigs
}
```

### ExtraConfigs

Additional lane configuration options.

```go
type ExtraConfigs struct {
    OnRampVersion []byte
}
```

---

## OCR3 Types

**Source:** [deploy/set_ocr3_config.go](../deploy/set_ocr3_config.go)

### SetOCR3ConfigArgs

Top-level input for the `SetOCR3Config` changeset.

```go
type SetOCR3ConfigArgs struct {
    HomeChainSel    uint64           // Home chain selector (where CCIPHome lives)
    RemoteChainSels []uint64         // Remote chains to configure
    ConfigType      utils.ConfigType // "active" or "candidate"
    MCMS            mcms.Input
}
```

### SetOCR3ConfigInput

Per-chain input passed to the `Deployer.SetOCR3Config()` sequence.

```go
type SetOCR3ConfigInput struct {
    ChainSelector uint64
    Datastore     datastore.DataStore
    Configs       map[ccipocr3.PluginType]OCR3ConfigArgs
}
```

### OCR3ConfigArgs

OCR3 configuration parameters for a single plugin.

```go
type OCR3ConfigArgs struct {
    ConfigDigest                   [32]byte
    PluginType                     ccipocr3.PluginType
    F                              uint8      // Faulty nodes tolerance
    IsSignatureVerificationEnabled bool
    Signers                        [][]byte
    Transmitters                   [][]byte
}
```

---

## Token Types

**Source:** [tokens/product.go](../tokens/product.go), [tokens/token_expansion.go](../tokens/token_expansion.go), [tokens/configure_tokens_for_transfers.go](../tokens/configure_tokens_for_transfers.go), [tokens/manual_registration.go](../tokens/manual_registration.go), [tokens/rate_limits.go](../tokens/rate_limits.go)

### TokenExpansionInput

Top-level input for the `TokenExpansion` changeset. Deploys tokens, pools, and configures them for cross-chain transfers.

```go
type TokenExpansionInput struct {
    TokenExpansionInputPerChain map[uint64]TokenExpansionInputPerChain
    ChainAdapterVersion         *semver.Version
    MCMS                        mcms.Input
}
```

### TokenExpansionInputPerChain

Per-chain token expansion configuration.

```go
type TokenExpansionInputPerChain struct {
    TokenPoolVersion      *semver.Version
    DeployTokenInput      *DeployTokenInput      // nil = token already deployed
    DeployTokenPoolInput  *DeployTokenPoolInput   // nil = pool already deployed
    TokenTransferConfig   *TokenTransferConfig    // nil = skip transfer configuration
    SkipOwnershipTransfer bool                    // Skip timelock ownership transfer
}
```

### DeployTokenInput

Input for deploying a new token.

```go
type DeployTokenInput struct {
    Name                  string              // Token name
    Symbol                string              // Token symbol
    Decimals              uint8               // Token decimals
    Supply                *big.Int            // Total supply
    PreMint               *big.Int            // Amount to pre-mint
    ExternalAdmin         string              // External admin address (chain-agnostic string)
    CCIPAdmin             string              // CCIP admin address (defaults to timelock)
    Senders               []string            // Addresses needing special processing (e.g., Solana ATAs)
    Type                  cldf.ContractType   // SPLToken, ERC20, etc.
    TokenPrivKey          string              // Solana: base58 private key for vanity addresses
    DisableFreezeAuthority bool               // Solana: revoke freeze authority permanently
    TokenMetadata         *TokenMetadata      // Solana: token metadata to upload
    // Populated programmatically
    ChainSelector         uint64
    ExistingDataStore     datastore.DataStore
}
```

### TokenMetadata

Token metadata for Solana tokens (extensible to other VMs).

```go
type TokenMetadata struct {
    TokenPubkey     string  // Populated programmatically
    MetadataJSONPath string // Path to metadata JSON (initial upload)
    UpdateAuthority string  // Update authority for metadata PDA
    UpdateName      string  // Update token name after upload
    UpdateSymbol    string  // Update token symbol after upload
    UpdateURI       string  // Update token URI after upload
}
```

### DeployTokenPoolInput

Input for deploying a token pool.

```go
type DeployTokenPoolInput struct {
    TokenRef           *datastore.AddressRef  // Reference to the token
    TokenPoolQualifier string                  // Pool qualifier in DataStore
    PoolType           string                  // BurnMintTokenPool, LockReleaseTokenPool, etc.
    TokenPoolVersion   *semver.Version
    Allowlist          []string
    AcceptLiquidity    *bool                   // LockReleaseTokenPool v1.5.1 only
    BurnAddress        string                  // BurnToAddressMintTokenPool only
    TokenGovernor      string                  // BurnMintWithExternalMinterTokenPool only
    // Populated programmatically
    ChainSelector     uint64
    ExistingDataStore datastore.DataStore
}
```

### TokenTransferConfig

Configuration for enabling a token for cross-chain transfers.

```go
type TokenTransferConfig struct {
    ChainSelector uint64                                                    // Target chain
    TokenPoolRef  datastore.AddressRef                                      // Token pool reference
    TokenRef      datastore.AddressRef                                      // Token reference
    ExternalAdmin string                                                    // External admin (leave empty for internal)
    RegistryRef   datastore.AddressRef                                      // Token admin registry reference
    RemoteChains  map[uint64]RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]
}
```

### ConfigureTokensForTransfersConfig

Top-level input for the `ConfigureTokensForTransfers` changeset.

```go
type ConfigureTokensForTransfersConfig struct {
    ChainAdapterVersion *semver.Version
    Tokens              []TokenTransferConfig
    MCMS                mcms.Input
}
```

### ConfigureTokenForTransfersInput

Per-chain input passed to the `TokenAdapter.ConfigureTokenForTransfersSequence()`.

```go
type ConfigureTokenForTransfersInput struct {
    ChainSelector    uint64
    TokenPoolAddress string
    RemoteChains     map[uint64]RemoteChainConfig[[]byte, string]
    ExternalAdmin    string
    RegistryAddress  string
    // Populated programmatically
    ExistingDataStore datastore.DataStore
    PoolType          string
    TokenRef          datastore.AddressRef
}
```

### RemoteChainConfig

Generic configuration for a remote chain on a token pool. Parameterized by address type (`R`) and cross-chain verifier type (`CCV`).

```go
type RemoteChainConfig[R any, CCV any] struct {
    RemoteToken                 R                           // Token on remote chain
    RemotePool                  R                           // Token pool on remote chain
    RemoteDecimals              uint8                       // Token decimals on remote chain
    InboundRateLimiterConfig    RateLimiterConfigFloatInput // Derived from counterpart outbound
    OutboundRateLimiterConfig   RateLimiterConfigFloatInput // User-specified
    OutboundCCVs                []CCV                       // Outbound cross-chain verifiers
    InboundCCVs                 []CCV                       // Inbound cross-chain verifiers
}
```

### RateLimiterConfig

On-chain rate limiter configuration (uses `*big.Int` for precision).

```go
type RateLimiterConfig struct {
    IsEnabled bool
    Capacity  *big.Int  // Maximum tokens in bucket
    Rate      *big.Int  // Refill rate (tokens per second)
}
```

### RateLimiterConfigFloatInput

User-friendly rate limiter input (float values scaled by token decimals internally).

```go
type RateLimiterConfigFloatInput struct {
    IsEnabled bool
    Capacity  float64
    Rate      float64
}
```

### ManualRegistrationInput

Top-level input for the `ManualRegistration` changeset.

```go
type ManualRegistrationInput struct {
    ChainAdapterVersion *semver.Version
    Registrations       []RegisterTokenConfig
    MCMS                mcms.Input
}
```

### RegisterTokenConfig

Per-registration configuration for manual token registration.

```go
type RegisterTokenConfig struct {
    TokenPoolRef  datastore.AddressRef  // Token pool reference (always required on SVM)
    TokenRef      datastore.AddressRef  // Token reference (always required on SVM)
    ChainSelector uint64                // Chain selector (required)
    ProposedOwner string                // Proposed owner address (required)
    SVMExtraArgs  *SVMExtraArgs         // SVM-specific extra args (optional)
}
```

### SVMExtraArgs

Solana-specific arguments for manual registration.

```go
type SVMExtraArgs struct {
    CustomerMintAuthorities []solana.PublicKey
    SkipTokenPoolInit       bool
}
```

### TPRLInput

Top-level input for the `SetTokenPoolRateLimits` changeset.

```go
type TPRLInput struct {
    Configs map[uint64]TPRLConfig
    MCMS    mcms.Input
}
```

### TPRLConfig

Per-chain rate limit configuration.

```go
type TPRLConfig struct {
    ChainAdapterVersion *semver.Version
    TokenRef            datastore.AddressRef
    TokenPoolRef        datastore.AddressRef
    RemoteOutbounds     map[uint64]RateLimiterConfigFloatInput  // Remote chain -> outbound limits
}
```

### TPRLRemotes

Per-remote-chain input passed to the `TokenAdapter.SetTokenPoolRateLimits()` sequence.

```go
type TPRLRemotes struct {
    OutboundRateLimiterConfig RateLimiterConfig
    InboundRateLimiterConfig  RateLimiterConfig
    ChainSelector             uint64
    RemoteChainSelector       uint64
    TokenRef                  datastore.AddressRef
    TokenPoolRef              datastore.AddressRef
    ExistingDataStore         datastore.DataStore
}
```

### UpdateAuthoritiesInput

Input for transferring token and pool ownership to the timelock.

```go
type UpdateAuthoritiesInput struct {
    ChainSelector uint64
    TokenRef      datastore.AddressRef
    TokenPoolRef  datastore.AddressRef
}
```

---

## Fee Types

**Source:** [fees/models.go](../fees/models.go), [fees/product.go](../fees/product.go)

### SetTokenTransferFeeSequenceInput

Input passed to the `FeeAdapter.SetTokenTransferFee()` sequence.

```go
type SetTokenTransferFeeSequenceInput struct {
    // Settings maps destination chain selector -> token address -> fee args
    Settings map[uint64]map[string]*TokenTransferFeeArgs
    Selector uint64
}
```

### TokenTransferFeeArgs

Standardized token transfer fee configuration for all chain families.

```go
type TokenTransferFeeArgs struct {
    DestBytesOverhead uint32  // Additional bytes overhead on destination
    DestGasOverhead   uint32  // Additional gas overhead on destination
    MinFeeUSDCents    uint32  // Minimum fee in USD cents
    MaxFeeUSDCents    uint32  // Maximum fee in USD cents
    DeciBps           uint16  // Fee in deci-basis points (1/10th of a basis point)
    IsEnabled         bool    // Whether fee is enabled
}
```

### UnresolvedTokenTransferFeeArgs

Allows partial specification of fee configuration. Unset fields are auto-filled from on-chain data or defaults.

```go
type UnresolvedTokenTransferFeeArgs struct {
    DestBytesOverhead TokenTransferFeeValue[uint32]
    DestGasOverhead   TokenTransferFeeValue[uint32]
    MinFeeUSDCents    TokenTransferFeeValue[uint32]
    MaxFeeUSDCents    TokenTransferFeeValue[uint32]
    DeciBps           TokenTransferFeeValue[uint16]
    IsEnabled         TokenTransferFeeValue[bool]
}
```

### TokenTransferFeeValue

A wrapper that indicates whether a value was explicitly set or should use a fallback.

```go
type TokenTransferFeeValue[T any] struct {
    Valid bool  // If false, Value is auto-filled from on-chain data or defaults
    Value T     // Only used when Valid is true
}
```

---

## Ownership Types

**Source:** [deploy/transfer_ownership.go](../deploy/transfer_ownership.go)

### TransferOwnershipInput

Top-level input for `TransferOwnership` and `AcceptOwnership` changesets.

```go
type TransferOwnershipInput struct {
    ChainInputs    []TransferOwnershipPerChainInput
    AdapterVersion *semver.Version
    MCMS           mcms.Input
}
```

### TransferOwnershipPerChainInput

Per-chain ownership transfer configuration.

```go
type TransferOwnershipPerChainInput struct {
    ChainSelector uint64
    ContractRef   []datastore.AddressRef  // Contracts to transfer ownership of
    CurrentOwner  string
    ProposedOwner string
}
```

---

## MCMS Input Types

**Source:** [utils/mcms/mcms.go](../utils/mcms/mcms.go)

### mcms.Input

Configuration for MCMS proposal construction. Included in most changeset inputs.

```go
type Input struct {
    OverridePreviousRoot bool                     // Override existing MCMS root
    ValidUntil           uint32                   // Unix timestamp for proposal expiry
    TimelockDelay        mcms_types.Duration      // Delay before operations can execute
    TimelockAction       mcms_types.TimelockAction // schedule, bypass, or cancel
    Qualifier            string                   // Qualifier for MCMS + Timelock addresses
    Description          string                   // Human-readable proposal description
}
```

**TimelockAction values:**
- `mcms_types.TimelockActionSchedule` -- schedule operations through the timelock
- `mcms_types.TimelockActionBypass` -- bypass the timelock (immediate execution)
- `mcms_types.TimelockActionCancel` -- cancel pending operations

---

## Output Types

**Source:** [utils/sequences/sequences.go](../utils/sequences/sequences.go)

### OnChainOutput

Standard output type returned by all sequences.

```go
type OnChainOutput struct {
    Addresses []datastore.AddressRef       // Deployed contract addresses
    Metadata  Metadata                     // Execution metadata
    BatchOps  []mcms_types.BatchOperation  // MCMS batch operations
}
```

### Metadata

Metadata about sequence execution, persisted to the DataStore.

```go
type Metadata struct {
    Contracts []datastore.ContractMetadata  // Contract-level metadata
    Chain     *datastore.ChainMetadata      // Chain-level metadata
    Env       *datastore.EnvMetadata        // Environment-level metadata
}
```

---

## Curse Types

**Source:** [fastcurse/product.go](../fastcurse/product.go)

See [Interfaces Reference](interfaces.md#curseadapter) for `CurseInput`, `Subject`, and related types used by `CurseAdapter` and `CurseSubjectAdapter`.

---

## Test Adapter Types

**Source:** [testadapters/adapters.go](../testadapters/adapters.go)

### MessageComponents

Generic CCIP message components, independent of chain family.

```go
type MessageComponents struct {
    DestChainSelector uint64    // Destination chain selector
    Receiver          []byte    // Dest-chain-encoded receiver address
    Data              []byte    // Message payload
    FeeToken          string    // Fee token identifier (source-chain encoded)
    ExtraArgs         []byte    // Message extra args (chain-family encoded)
    TokenAmounts      []TokenAmount
}
```

### TokenAmount

Token and amount for cross-chain transfers.

```go
type TokenAmount struct {
    Token  string    // Source-chain-encoded token address
    Amount *big.Int
}
```

### ExtraArgOpt

Chain-agnostic representation of a message extra arg.

```go
type ExtraArgOpt struct {
    Name  string
    Value any
}
```

Constructors: `NewOutOfOrderExtraArg(bool)`, `NewGasLimitExtraArg(*big.Int)`.

---

## Constants

**Source:** [utils/common.go](../utils/common.go)

### Contract Type Constants

```go
const (
    BypasserManyChainMultisig  = "BypasserManyChainMultiSig"
    CancellerManyChainMultisig = "CancellerManyChainMultiSig"
    ProposerManyChainMultisig  = "ProposerManyChainMultiSig"
    RBACTimelock               = "RBACTimelock"
    CallProxy                  = "CallProxy"
    CapabilitiesRegistry       = "CapabilitiesRegistry"
    CCIPHome                   = "CCIPHome"
    RMNHome                    = "RMNHome"
    BurnMintTokenPool          = "BurnMintTokenPool"
    LockReleaseTokenPool       = "LockReleaseTokenPool"
    TokenPoolLookupTable       = "TokenPoolLookupTable"
    BurnWithFromMintTokenPool  = "BurnWithFromMintTokenPool"
    BurnFromMintTokenPool      = "BurnFromMintTokenPool"
    CCTPTokenPool              = "CCTPTokenPool"
)
```

### Version Constants

```go
var (
    Version_1_0_0 = semver.MustParse("1.0.0")
    Version_1_5_0 = semver.MustParse("1.5.0")
    Version_1_5_1 = semver.MustParse("1.5.1")
    Version_1_6_0 = semver.MustParse("1.6.0")
    Version_1_6_1 = semver.MustParse("1.6.1")
)
```

### Chain Family Selectors

On-chain identifiers for each chain family, derived from `keccak256`:

| Family | Selector | Constant |
|--------|----------|----------|
| EVM | `0x2812d52c` | `EVMFamilySelector` |
| SVM (Solana) | `0x1e10bdc4` | `SVMFamilySelector` |
| Aptos | `0xac77ffec` | `AptosFamilySelector` |
| TVM (TON) | `0x647e2ba9` | `TVMFamilySelector` |
| Sui | `0xc4e05953` | `SuiFamilySelector` |

### Qualifier Constants

```go
const (
    CLLQualifier         = "CLLCCIP"     // Standard CLL qualifier
    RMNTimelockQualifier = "RMNMCMS"     // RMN timelock qualifier
)
```

### Execution State Constants

```go
const (
    EXECUTION_STATE_UNTOUCHED  = 0
    EXECUTION_STATE_INPROGRESS = 1
    EXECUTION_STATE_SUCCESS    = 2
    EXECUTION_STATE_FAILURE    = 3
)
```
