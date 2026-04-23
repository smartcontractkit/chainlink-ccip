---
title: "SolanaAdapter Reference"
sidebar_label: "Adapter"
sidebar_position: 2
---

# SolanaAdapter Reference

The `SolanaAdapter` implements the shared deployment interfaces for the Solana Virtual Machine (SVM). Unlike the stateless EVM adapter, the Solana adapter carries state for timelock addresses.

**Source:** [v1_6_0/sequences/adapter.go](../v1_6_0/sequences/adapter.go)

---

## Struct Definition

```go
type SolanaAdapter struct {
    timelockAddr map[uint64]solana.PublicKey
}
```

The `timelockAddr` map caches resolved timelock public keys by chain selector, avoiding repeated datastore lookups.

---

## Registration

The Solana adapter registers itself with **5 shared registries** via two `init()` functions:

### Main Registration (`v1_6_0/sequences/adapter.go`)

**Source:** [v1_6_0/sequences/adapter.go](../v1_6_0/sequences/adapter.go)

```go
func init() {
    v := semver.MustParse("1.6.0")
    laneapi.GetLaneAdapterRegistry().RegisterLaneAdapter(chain_selectors.FamilySolana, v, &SolanaAdapter{})
    deployapi.GetRegistry().RegisterDeployer(chain_selectors.FamilySolana, v, &SolanaAdapter{})
    deployapi.GetTransferOwnershipRegistry().RegisterAdapter(chain_selectors.FamilySolana, v, &SolanaAdapter{})
    mcmsreaderapi.GetRegistry().RegisterMCMSReader(chain_selectors.FamilySolana, &SolanaAdapter{})
    tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chain_selectors.FamilySolana, v, &SolanaAdapter{})
}
```

### Curse Adapter Registration (`v1_6_0/adapters/init.go`)

**Source:** [v1_6_0/adapters/init.go](../v1_6_0/adapters/init.go)

```go
func init() {
    curseRegistry := fastcurse.GetCurseRegistry()
    curseRegistry.RegisterNewCurse(fastcurse.CurseRegistryInput{
        CursingFamily:       chainsel.FamilySolana,
        CursingVersion:      semver.MustParse("1.6.0"),
        CurseAdapter:        NewCurseAdapter(),
        CurseSubjectAdapter: NewCurseAdapter(),
    })
}
```

### Registration Summary

| Registry | Interface | Key |
|----------|-----------|-----|
| `LaneAdapterRegistry` | `LaneAdapter` | `solana-1.6.0` |
| `DeployerRegistry` | `Deployer` | `solana-1.6.0` |
| `TransferOwnershipAdapterRegistry` | `TransferOwnershipAdapter` | `solana-1.6.0` |
| `MCMSReaderRegistry` | `MCMSReader` | `solana` (no version) |
| `TokenAdapterRegistry` | `TokenAdapter` | `solana-1.6.0` |
| `CurseRegistry` | `CurseAdapter` + `CurseSubjectAdapter` | `solana-1.6.0` |

---

## Interface Implementations

### Deployer Interface

| Method | Implementation |
|--------|---------------|
| `DeployChainContracts()` | Returns `DeployChainContracts` sequence (deploys LINK, Router, FeeQuoter, OffRamp, RMNRemote, BurnMint/LockRelease token pools, TestReceiver; initializes all; extends lookup table) |
| `DeployMCMS()` | Deploys Access Controller, MCM, and Timelock programs; initializes access controller roles, MCM configs, and timelock |
| `FinalizeDeployMCMS()` | Configures MCM (set_config), sets up timelock roles (Proposer, Executor, Canceller, Bypasser) |
| `SetOCR3Config()` | Returns `SetOCR3Config` sequence (sets OCR3 config on OffRamp) |
| `GrantAdminRoleToTimelock()` | Returns `nil` (not implemented for Solana) |
| `UpdateMCMSConfig()` | Updates config of specified MCMS contracts |

### LaneAdapter Interface

| Method | Implementation |
|--------|---------------|
| `ConfigureLaneLegAsSource()` | Adds destination chain config to FeeQuoter, adds destination chain to Router |
| `ConfigureLaneLegAsDest()` | Adds OffRamp to Router for remote chain, adds source chain config to OffRamp |
| `GetOnRampAddress()` | **Delegates to `GetRouterAddress()`** -- on Solana, the Router serves as the OnRamp |
| `GetOffRampAddress()` | Looks up OffRamp from DataStore |
| `GetFQAddress()` | Looks up FeeQuoter from DataStore |
| `GetRouterAddress()` | Looks up Router from DataStore |
| `GetRMNRemoteAddress()` | Looks up RMNRemote from DataStore |

### TokenAdapter Interface

| Method | Implementation |
|--------|---------------|
| `ConfigureTokenForTransfersSequence()` | Registers token in admin registry, accepts admin, sets pool on Router, configures remote chains with rate limits |
| `AddressRefToBytes()` | Converts base58 address to `[]byte` via `solana.MustPublicKeyFromBase58` |
| `DeriveTokenAddress()` | Not implemented (returns error) |
| `DeriveTokenDecimals()` | Reads decimals from on-chain mint account data |
| `DeriveTokenPoolCounterpart()` | Derives PDA using `TokenPoolConfigAddress(tokenMint, tokenPool)` |
| `ManualRegistration()` | Registers token admin registry, initializes token pool, optionally creates token multisig with customer mint authorities |
| `SetTokenPoolRateLimits()` | Sets inbound/outbound rate limits for BurnMint or LockRelease pools |
| `DeployToken()` | Deploys SPL/SPL2022 token, creates ATAs for senders, optionally uploads metadata |
| `DeployTokenVerify()` | No-op (returns nil) |
| `DeployTokenPoolForToken()` | Initializes pool account (BurnMint or LockRelease), creates pool signer ATA, sets mint authority for BurnMint pools |
| `UpdateAuthorities()` | Transfers token pool ownership to timelock signer PDA, updates rate limit admin |

### TransferOwnershipAdapter Interface

| Method | Implementation |
|--------|---------------|
| `InitializeTimelockAddress()` | No-op (returns nil) |
| `SequenceTransferOwnershipViaMCMS()` | Dispatches per contract type: Router, OffRamp, FeeQuoter, RMNRemote, AccessController, RBACTimelock |
| `ShouldAcceptOwnershipWithTransferOwnership()` | Returns `true` when current owner equals the chain's deployer key |
| `SequenceAcceptOwnership()` | Dispatches accept-ownership per contract type (same types as transfer) |

### MCMSReader Interface

| Method | Implementation |
|--------|---------------|
| `GetChainMetadata()` | Resolves MCM contract address based on timelock action (Schedule/Cancel/Bypass), gets op count from inspector, builds Solana MCMS chain metadata with access controller accounts |
| `GetTimelockRef()` | Looks up RBACTimelock from DataStore |
| `GetMCMSRef()` | Looks up MCM program from DataStore |

---

## Solana-Specific Patterns

### Router = OnRamp

On Solana, the Router program functions as both the Router and OnRamp. `GetOnRampAddress()` delegates directly to `GetRouterAddress()`:

```go
func (a *SolanaAdapter) GetOnRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
    return a.GetRouterAddress(ds, chainSelector)
}
```

### Two-Phase MCMS Deployment

Unlike EVM where `FinalizeDeployMCMS()` is a no-op, Solana requires:

1. **Phase 1 (`DeployMCMS`)**: Deploy programs + initialize (Access Controller accounts, MCM init, Timelock init)
2. **Phase 2 (`FinalizeDeployMCMS`)**: Configure MCM (set_config for Proposer/Canceller/Bypasser) + set up timelock roles

### PDA-Based Address Derivation

Solana uses Program Derived Addresses (PDAs) extensively:

- `FindConfigPDA(routerProgram)` -- Router config
- `FindOfframpConfigPDA(offRampProgram)` -- OffRamp config
- `FindFqConfigPDA(feeQuoterProgram)` -- FeeQuoter config
- `FindRMNRemoteConfigPDA(rmnRemoteProgram)` -- RMN Remote config
- `state.GetTimelockSignerPDA(timelockID, seed)` -- Timelock signer
- `state.GetMCMSignerPDA(mcmID, seed)` -- MCM signer
- `tokens.TokenPoolConfigAddress(tokenMint, tokenPool)` -- Token pool config
- `tokens.TokenPoolSignerAddress(tokenMint, tokenPool)` -- Token pool signer

### Base58 Address Handling

All Solana addresses use base58 encoding. The adapter converts between `datastore.AddressRef` (string addresses) and `solana.PublicKey`:

```go
func (a *SolanaAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
    return solana.MustPublicKeyFromBase58(ref.Address).Bytes(), nil
}
```

### Solana-Specific Contract Types

Defined in `utils/common.go`:

| Constant | Type |
|----------|------|
| `TimelockProgramType` | `"RBACTimelockProgram"` |
| `McmProgramType` | `"ManyChainMultiSigProgram"` |
| `AccessControllerProgramType` | `"AccessControllerProgram"` |
| `ProposerSeed` | `"ProposerSeed"` |
| `CancellerSeed` | `"CancellerSeed"` |
| `BypasserSeed` | `"BypasserSeed"` |
| `RBACTimelockSeed` | `"RBACTimelockSeed"` |
| `ProposerAccessControllerAccount` | `"ProposerAccessControllerAccount"` |
| `ExecutorAccessControllerAccount` | `"ExecutorAccessControllerAccount"` |
| `CancellerAccessControllerAccount` | `"CancellerAccessControllerAccount"` |
| `BypasserAccessControllerAccount` | `"BypasserAccessControllerAccount"` |
| `SPLTokens` | `"SPLTokens"` |
| `SPL2022Tokens` | `"SPL2022Tokens"` |
