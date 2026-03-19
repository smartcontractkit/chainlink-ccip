---
title: "EVMAdapter Reference"
sidebar_label: "Adapter"
sidebar_position: 2
---

# EVMAdapter Reference

The `EVMAdapter` is a stateless struct that implements all required shared interfaces for the EVM chain family.

**Source:** [v1_6_0/sequences/adapter.go](../v1_6_0/sequences/adapter.go)

For the interfaces it implements, see [Interfaces Reference](../../../../deployment/docs/interfaces.md).

---

## Struct Definition

```go
type EVMAdapter struct{}
```

The EVM adapter is intentionally stateless -- all state is resolved from the DataStore and environment at execution time.

---

## Registration

### Main Adapter (`sequences/adapter.go`)

**Source:** [v1_6_0/sequences/adapter.go](../v1_6_0/sequences/adapter.go)

```go
func init() {
    v, _ := semver.NewVersion("1.6.0")

    laneapi.GetLaneAdapterRegistry().RegisterLaneAdapter(chain_selectors.FamilyEVM, v, &EVMAdapter{})
    deployapi.GetRegistry().RegisterDeployer(chain_selectors.FamilyEVM, v, &EVMAdapter{})
    deployapi.GetTransferOwnershipRegistry().RegisterAdapter(chain_selectors.FamilyEVM, v, &EVMAdapter{})
    mcmsreaderapi.GetRegistry().RegisterMCMSReader(chain_selectors.FamilyEVM, &EVMAdapter{})
    tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chain_selectors.FamilyEVM, v, &EVMAdapter{})
    lanes.GetPingPongAdapterRegistry().RegisterPingPongAdapter(chain_selectors.FamilyEVM, v, &EVMAdapter{})
}
```

### Separate Adapters (`adapters/init.go`)

**Source:** [v1_6_0/adapters/init.go](../v1_6_0/adapters/init.go)

```go
func init() {
    // Curse adapter
    fastcurse.GetCurseRegistry().RegisterNewCurse(fastcurse.CurseRegistryInput{
        CursingFamily:       chain_selectors.FamilyEVM,
        CursingVersion:      semver.MustParse("1.6.0"),
        CurseAdapter:        NewCurseAdapter(),
        CurseSubjectAdapter: NewCurseAdapter(),
    })

    // Lane migration adapters
    deploy.GetLaneMigratorRegistry().RegisterRampUpdater(chain_selectors.FamilyEVM, semver.MustParse("1.6.0"), &LaneMigrater{})
    deploy.GetLaneMigratorRegistry().RegisterRouterUpdater(chain_selectors.FamilyEVM, semver.MustParse("1.2.0"), &RouterUpdater{})
}
```

---

## Interface Implementations

### Deployer

| Method | Delegates To |
|--------|-------------|
| `DeployChainContracts()` | `DeployChainContracts` sequence |
| `DeployMCMS()` | v1.0.0 MCMS deployer |
| `FinalizeDeployMCMS()` | v1.0.0 MCMS finalize (no-op for EVM) |
| `SetOCR3Config()` | OCR3 configuration sequence |
| `GrantAdminRoleToTimelock()` | Timelock admin role sequence |

EVM delegates MCMS deployment to the v1.0.0 implementation since MCMS contracts are version-agnostic on EVM.

### LaneAdapter

| Method | Behavior |
|--------|----------|
| `GetOnRampAddress()` | Looks up OnRamp v1.6.0 from DataStore |
| `GetOffRampAddress()` | Looks up OffRamp v1.6.0 from DataStore |
| `GetRouterAddress()` | Looks up Router from DataStore |
| `GetFQAddress()` | Looks up latest FeeQuoter (v1.6.0 -- v2.0.0 range) |
| `ConfigureLaneLegAsSource()` | OnRamp dest config + FeeQuoter dest config + price updates |
| `ConfigureLaneLegAsDest()` | OffRamp source config + Router ramp updates |

All `Get*Address` methods return 20-byte EVM addresses.

### TokenAdapter

| Method | Behavior |
|--------|----------|
| `AddressRefToBytes()` | Hex-decodes the address string to bytes |
| `DeriveTokenAddress()` | Reads token address from the pool contract on-chain |
| `DeriveTokenDecimals()` | Reads decimals from the token contract on-chain |
| `DeriveTokenPoolCounterpart()` | Returns tokenPool unchanged (no PDA derivation needed on EVM) |
| `ConfigureTokenForTransfersSequence()` | Registers token + sets remote chain configs |
| `ManualRegistration()` | Registers customer token via proposeAdministrator |
| `SetTokenPoolRateLimits()` | Sets inbound/outbound rate limits |
| `DeployToken()` | Deploys ERC20 BurnMint token |
| `DeployTokenPoolForToken()` | Deploys token pool for existing token |
| `UpdateAuthorities()` | Transfers token/pool ownership to timelock |

### TokenPriceProvider (Optional)

```go
func (e *EVMAdapter) GetDefaultTokenPrices() map[datastore.ContractType]*big.Int
```

Returns a default price of $20 per token (18 decimals) for WETH and LINK.

### TransferOwnershipAdapter

| Method | Behavior |
|--------|----------|
| `InitializeTimelockAddress()` | No-op on EVM (timelock resolved at execution time) |
| `SequenceTransferOwnershipViaMCMS()` | Calls `transferOwnership()` on each contract |
| `SequenceAcceptOwnership()` | Calls `acceptOwnership()` on each contract |
| `ShouldAcceptOwnershipWithTransferOwnership()` | Returns `false` (two-step ownership on EVM) |

### MCMSReader

| Method | Behavior |
|--------|----------|
| `GetChainMetadata()` | Reads MCM starting op count from on-chain |
| `GetTimelockRef()` | Resolves RBACTimelock AddressRef from DataStore |
| `GetMCMSRef()` | Resolves ProposerManyChainMultiSig AddressRef |

---

## Specialized Adapters

### CurseAdapter (`adapters/fastcurse.go`)

Implements `CurseAdapter` and `CurseSubjectAdapter` for EVM RMN operations.

- `Initialize()` -- loads RMNRemote contract address
- `IsSubjectCursedOnChain()` -- calls `isCursed(subject)` on RMNRemote
- `Curse()` / `Uncurse()` -- returns sequences that call `curse()`/`uncurse()` on RMNRemote
- `SelectorToSubject()` -- converts chain selector to 16-byte curse subject
- `SubjectToSelector()` -- reverses the conversion

### FeeAdapter (`adapters/fees.go`)

Implements `FeeAdapter` for EVM fee configuration.

- `SetTokenTransferFee()` -- returns sequence that calls `applyTokenTransferFeeConfigUpdates` on FeeQuoter
- `GetOnchainTokenTransferFeeConfig()` -- reads current fee config from FeeQuoter
- `GetDefaultTokenTransferFeeConfig()` -- returns sensible defaults

### ConfigImporter (`adapters/configimport.go`)

Implements `ConfigImporter` for importing existing on-chain state into the DataStore.

### LaneMigrater (`adapters/lanemigrator.go`)

Implements `RampUpdateInRouter` and `RouterUpdateInRamp` for lane migration scenarios.

---

## EVM Utilities

### Address Conversion (`utils/datastore/datastore.go`)

```go
func ToByteArray(ref datastore.AddressRef) ([]byte, error)       // Hex string -> bytes
func ToEVMAddress(ref datastore.AddressRef) (common.Address, error)  // Hex string -> common.Address
func ToPaddedEVMAddress(ref datastore.AddressRef) ([]byte, error)   // 32-byte left-padded
```

### Contract Introspection (`utils/common.go`)

```go
func TypeAndVersion(addr common.Address, client bind.ContractBackend) (string, *semver.Version, error)
func ValidateEVMAddress(addr string, fieldName string) error
```
