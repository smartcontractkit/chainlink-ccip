---
title: "Implementing a New Chain Family Adapter"
sidebar_label: "Implementing Adapters"
sidebar_position: 6
---

# Implementing a New Chain Family Adapter

This guide provides a step-by-step walkthrough for adding support for a new chain family to the CCIP Deployment Tooling API.

For the complete interface definitions, see [Interfaces Reference](interfaces.md). For type definitions, see [Types Reference](types.md).

## Prerequisites

- Familiarity with the [Architecture](architecture.md) (adapter-registry pattern, operations-sequences-changesets hierarchy)
- The `chain-selectors` library must already define a family constant for your chain (e.g., `chain_selectors.FamilyAptos`)
- Smart contracts for your chain family are deployed or in development

## Step 1: Choose Your Module Location

By convention, chain-family adapters live alongside their contracts:

```
chainlink-ccip/chains/<family>/deployment/
```

Or in an external repository:

```
chainlink-<family>/deployment/
```

Create a Go module with its own `go.mod` that imports `chainlink-ccip/deployment` for the shared types and interfaces.

## Step 2: Set Up the Directory Structure

Follow this canonical layout:

```
chains/<family>/deployment/
├── go.mod
├── go.sum
├── utils/                          # Shared utilities for this chain family
│   ├── common.go                   # Chain-specific contract type constants
│   ├── deploy.go                   # Deployment helpers
│   ├── mcms.go                     # MCMS-related utilities
│   └── datastore.go                # DataStore helpers
├── v1_6_0/                         # Version-specific implementation
│   ├── adapters/                   # Adapters registered separately (curse, fees)
│   │   └── init.go                 # init() registrations for adapters
│   ├── operations/                 # Low-level contract operations
│   │   ├── router/
│   │   ├── offramp/
│   │   ├── fee_quoter/
│   │   ├── token_pools/
│   │   ├── tokens/
│   │   ├── rmn_remote/
│   │   └── mcms/
│   ├── sequences/                  # High-level operation orchestration
│   │   ├── adapter.go              # Main adapter struct + init() registration
│   │   ├── deploy_chain_contracts.go
│   │   ├── connect_chains.go
│   │   ├── ocr.go
│   │   ├── mcms.go
│   │   ├── tokens.go
│   │   ├── fee_quoter.go
│   │   └── transfer_ownership.go
│   └── testadapter/                # Test adapter implementation
└── docs/                           # Chain-specific documentation
```

## Step 3: Create the Adapter Struct

Define a single struct that will implement all required interfaces:

```go
// sequences/adapter.go
package sequences

type MyChainAdapter struct {
    // Add any stateful fields your chain needs.
    // For example, Solana caches timelock addresses:
    //   timelockAddr map[uint64]solana.PublicKey
    //
    // EVM uses a stateless empty struct:
    //   (no fields)
}
```

## Step 4: Implement Required Interfaces

### Interface Checklist

Each method returns a `*operations.Sequence` that the framework executes. Your job is to implement the sequence logic using chain-specific operations.

#### 4.1 Deployer

Deploy all CCIP infrastructure contracts on your chain:

```go
func (a *MyChainAdapter) DeployChainContracts() *Sequence[ContractDeploymentConfigPerChainWithAddress, OnChainOutput, BlockChains] {
    return DeployChainContractsSequence // your sequence variable
}

func (a *MyChainAdapter) DeployMCMS() *Sequence[MCMSDeploymentConfigPerChainWithAddress, OnChainOutput, BlockChains] {
    return DeployMCMSSequence
}

func (a *MyChainAdapter) FinalizeDeployMCMS() *Sequence[MCMSDeploymentConfigPerChainWithAddress, OnChainOutput, BlockChains] {
    return FinalizeDeployMCMSSequence // can be no-op if not needed
}

func (a *MyChainAdapter) SetOCR3Config() *Sequence[SetOCR3ConfigInput, OnChainOutput, BlockChains] {
    return SetOCR3ConfigSequence
}

func (a *MyChainAdapter) GrantAdminRoleToTimelock() *Sequence[GrantAdminRoleToTimelockConfigPerChainWithSelector, OnChainOutput, BlockChains] {
    return GrantAdminRoleToTimelockSequence
}
```

#### 4.2 LaneAdapter

Address retrieval methods must encode addresses in your chain's native format:

```go
func (a *MyChainAdapter) GetOnRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
    // Look up the OnRamp address from the DataStore and return as bytes
}

func (a *MyChainAdapter) GetOffRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) { ... }
func (a *MyChainAdapter) GetRouterAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) { ... }
func (a *MyChainAdapter) GetFQAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) { ... }

func (a *MyChainAdapter) ConfigureLaneLegAsSource() *Sequence[UpdateLanesInput, OnChainOutput, BlockChains] {
    return ConfigureLaneLegAsSourceSequence
}

func (a *MyChainAdapter) ConfigureLaneLegAsDest() *Sequence[UpdateLanesInput, OnChainOutput, BlockChains] {
    return ConfigureLaneLegAsDestSequence
}
```

#### 4.3 TokenAdapter

The token adapter requires several helper methods for address derivation:

```go
func (a *MyChainAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
    // Convert the AddressRef.Address string to bytes using your chain's encoding
    // EVM: hex decoding, Solana: base58 decoding
}

func (a *MyChainAdapter) DeriveTokenAddress(e Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
    // Read the token address from the pool contract on-chain
}

func (a *MyChainAdapter) DeriveTokenDecimals(e Environment, chainSelector uint64, poolRef datastore.AddressRef, token []byte) (uint8, error) {
    // Read the token decimals on-chain
}

func (a *MyChainAdapter) DeriveTokenPoolCounterpart(e Environment, chainSelector uint64, tokenPool []byte, token []byte) ([]byte, error) {
    // For chains where the operational pool address differs from the deployed address
    // (e.g., Solana PDAs). Return tokenPool unchanged if not applicable.
}
```

#### 4.4 FeeAdapter, MCMSReader, TransferOwnershipAdapter, CurseAdapter, CurseSubjectAdapter

Follow the same pattern -- see [Interfaces Reference](interfaces.md) for the full method signatures.

## Step 5: Create Operations

Operations are the atomic building blocks. Define them in the `operations/` directory, organized by contract.

Each operation follows this pattern:

```go
var Deploy = operations.NewOperation(
    "my-contract:deploy",                   // Operation ID: contract:method
    semver.MustParse("1.6.0"),              // Contract version
    "Deploys the MyContract program",       // Description
    func(b operations.Bundle, chain MyChain, input DeployInput) (DeployOutput, error) {
        // Chain-specific deployment logic
    },
)
```

## Step 6: Compose Sequences

Sequences orchestrate multiple operations into complete workflows. Define them in the `sequences/` directory.

```go
var DeployChainContractsSequence = operations.NewSequence(
    "deploy-chain-contracts",
    semver.MustParse("1.6.0"),
    "Deploys all CCIP contracts on a chain",
    func(b operations.Bundle, chains BlockChains, input ContractDeploymentConfigPerChainWithAddress) (OnChainOutput, error) {
        var agg sequences.OnChainOutput

        // Deploy Router
        agg, err := sequences.RunAndMergeSequence(b, chains, deployRouterSeq, routerInput, agg)
        if err != nil {
            return agg, err
        }

        // Deploy FeeQuoter
        agg, err = sequences.RunAndMergeSequence(b, chains, deployFeeQuoterSeq, fqInput, agg)
        if err != nil {
            return agg, err
        }

        // ... more contract deployments

        return agg, nil
    },
)
```

Key patterns:
- Use `sequences.RunAndMergeSequence` to compose sub-sequences and aggregate their outputs
- Each sequence should target a single chain for simplicity
- Return `OnChainOutput` with deployed addresses, metadata, and MCMS batch operations

## Step 7: Register via init()

Create `init()` functions that run automatically when your package is imported.

**Main adapter registration** (`sequences/adapter.go`):

```go
func init() {
    v, err := semver.NewVersion("1.6.0")
    if err != nil {
        panic(err)
    }

    // Required registrations
    laneapi.GetLaneAdapterRegistry().RegisterLaneAdapter(chain_selectors.FamilyMyChain, v, &MyChainAdapter{})
    deployapi.GetRegistry().RegisterDeployer(chain_selectors.FamilyMyChain, v, &MyChainAdapter{})
    deployapi.GetTransferOwnershipRegistry().RegisterAdapter(chain_selectors.FamilyMyChain, v, &MyChainAdapter{})
    mcmsreaderapi.GetRegistry().RegisterMCMSReader(chain_selectors.FamilyMyChain, &MyChainAdapter{})
    tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chain_selectors.FamilyMyChain, v, &MyChainAdapter{})
}
```

**Separate adapter registration** (`adapters/init.go`):

```go
func init() {
    // Curse/RMN adapter
    fastcurse.GetCurseRegistry().RegisterNewCurse(fastcurse.CurseRegistryInput{
        CursingFamily:       chain_selectors.FamilyMyChain,
        CursingVersion:      semver.MustParse("1.6.0"),
        CurseAdapter:        NewCurseAdapter(),
        CurseSubjectAdapter: NewCurseAdapter(),
    })

    // Fee adapter
    fees.GetRegistry().RegisterFeeAdapter(chain_selectors.FamilyMyChain, semver.MustParse("1.6.0"), &FeesAdapter{})
}
```

## Step 8: Wire Into chainlink-deployments

In `chainlink-deployments/domains/<CCIP_DOMAIN>/<ENVIRONMENT>/pipelines.go`, import your adapter package to trigger the `init()` functions:

```go
import (
    _ "github.com/smartcontractkit/chainlink-ccip/chains/mychain/deployment/v1_6_0/sequences"
    _ "github.com/smartcontractkit/chainlink-ccip/chains/mychain/deployment/v1_6_0/adapters"
)
```

The blank import is sufficient -- it triggers `init()` and registers all adapters into the global registries.

## Reference Implementations

Study these existing implementations:

| Chain | Adapter Source | Key Patterns |
|-------|---------------|--------------|
| **EVM** | [chains/evm/deployment/v1_6_0/sequences/adapter.go](../../chains/evm/deployment/v1_6_0/sequences/adapter.go) | Stateless adapter, auto-generated operations from gethwrappers, multi-version support (v1_0_0 through v1_6_5) |
| **Solana** | [chains/solana/deployment/v1_6_0/sequences/adapter.go](../../chains/solana/deployment/v1_6_0/sequences/adapter.go) | Stateful adapter (caches timelock addresses), two-phase MCMS deployment, PDA-based address derivation |

## Common Considerations

### Address Encoding

Each chain family must handle address serialization consistently:
- `AddressRefToBytes`: Convert string addresses from the DataStore to bytes
- `Get*Address` methods on `LaneAdapter`: Return addresses as bytes in your chain's native format
- Cross-chain references: When configuring a lane, the source chain receives the destination's addresses as `[]byte`

### Two-Phase MCMS

Some chains (like Solana) require a two-phase MCMS deployment:
1. `DeployMCMS()` -- deploy the contracts
2. `FinalizeDeployMCMS()` -- initialize the timelock and configure roles

If your chain can do everything in one step, `FinalizeDeployMCMS()` can return a no-op sequence.

### MCMS Batch Operations

When a write operation cannot be executed directly by the deployer key (because the contract is owned by a timelock), the operation should produce `mcms_types.BatchOperation` entries instead. These get collected by the `OutputBuilder` and assembled into an MCMS proposal.

### DataStore Integration

Sequences should return all deployed contract addresses in `OnChainOutput.Addresses` so the changeset can persist them to the DataStore. Use `datastore.AddressRef` with your chain's contract types, versions, and qualifiers.

### Contract Type Constants

Define your chain-specific contract types in `utils/common.go`. Use the shared constants from `deployment/utils/common.go` where applicable (e.g., `RBACTimelock`, `BurnMintTokenPool`).
