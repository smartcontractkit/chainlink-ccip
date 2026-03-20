---
title: "Architecture Guide"
sidebar_label: "Architecture"
sidebar_position: 2
---

# Architecture Guide

This document describes the design principles, patterns, and data flow of the CCIP Deployment Tooling API.

## Design Principles

1. **Chain Agnosticism**: Changesets operate without knowledge of chain-specific details. All chain-specific logic is encapsulated in adapters.
2. **Version-Aware Adapters**: Different contract versions have different adapters registering under the same interface. This allows multiple versions to coexist.
3. **Singleton Registries**: All adapter registries are global singletons, initialized once via Go `init()` functions and accessed thread-safely.
4. **Stateful Retry via Operations**: Operations produce reports that enable resuming long-running sequences from the point of failure.
5. **MCMS-First Governance**: All write operations can be routed through Multi-Chain Multi-Sig (MCMS) proposals when the deployer key is not the contract owner.

## Adapter-Registry Pattern

The core architectural pattern is a registry of chain-family adapters keyed by `chainFamily-version`:

```
Changeset (entry point)
    |
    v
Registry.GetAdapter(chainFamily, version)
    |
    v
Adapter (implements interface)
    |
    v
Sequence (ordered operations)
    |
    v
Operations (single side-effects)
```

### How Registries Work

Each registry is a singleton created via `sync.Once`:

```go
var (
    singletonRegistry *DeployerRegistry
    once              sync.Once
)

func GetRegistry() *DeployerRegistry {
    once.Do(func() {
        singletonRegistry = newDeployerRegistry()
    })
    return singletonRegistry
}
```

Adapters register themselves in Go `init()` functions, which run automatically when the package is imported:

```go
func init() {
    v := semver.MustParse("1.6.0")
    deploy.GetRegistry().RegisterDeployer(chain_selectors.FamilyEVM, v, &EVMAdapter{})
    lanes.GetLaneAdapterRegistry().RegisterLaneAdapter(chain_selectors.FamilyEVM, v, &EVMAdapter{})
    // ... more registrations
}
```

Registry keys are constructed using `utils.NewRegistererID(chainFamily, version)`, which produces strings like `"evm-1.6.0"` or `"solana-1.6.0"`. The `MCMSReaderRegistry` is the exception -- it uses only `chainFamily` as the key (one reader per family, not per version).

All registries use `sync.Mutex` for thread-safe concurrent access.

### Available Registries

| Registry | Key Format | Singleton Accessor | Source |
|----------|------------|-------------------|--------|
| `DeployerRegistry` | `chainFamily-version` | `deploy.GetRegistry()` | `deploy/product.go` |
| `LaneAdapterRegistry` | `chainFamily-version` | `lanes.GetLaneAdapterRegistry()` | `lanes/product.go` |
| `TokenAdapterRegistry` | `chainFamily-version` | `tokens.GetTokenAdapterRegistry()` | `tokens/product.go` |
| `FeeAdapterRegistry` | `chainFamily-version` | `fees.GetRegistry()` | `fees/product.go` |
| `MCMSReaderRegistry` | `chainFamily` | `changesets.GetRegistry()` | `utils/changesets/output.go` |
| `TransferOwnershipAdapterRegistry` | `chainFamily-version` | `deploy.GetTransferOwnershipRegistry()` | `deploy/product.go` |
| `CurseRegistry` | `chainFamily-version` | `fastcurse.GetCurseRegistry()` | `fastcurse/product.go` |

## The Three-Level Hierarchy

### Operations

An Operation is the most granular executable action -- a single contract deployment, transaction, or read call. Each operation has **at most one side effect**.

Operations produce reports that serialize their inputs and outputs, enabling stateful retries. If a sequence fails partway through, it can resume from the last successful operation.

Chain-specific implementations define operations differently:
- **EVM**: Uses `contract.NewRead`, `contract.NewWrite`, `contract.NewDeploy` helpers that wrap gethwrapper methods
- **Solana**: Uses `operations.NewOperation` with chain-specific transaction building

### Sequences

A Sequence is an ordered collection of operations that represents a complete workflow (e.g., "deploy all CCIP contracts on a chain" or "configure a lane between two chains").

Key properties:
- Accept a **serializable input** and return `OnChainOutput`
- Depend on minimal infrastructure (typically just `cldf_chain.BlockChains`)
- Target a **single chain** for simplicity
- Can be composed using `RunAndMergeSequence` to aggregate outputs

```go
type OnChainOutput struct {
    Addresses []datastore.AddressRef          // Deployed/managed contracts
    Metadata  Metadata                         // Contract, chain, and env metadata
    BatchOps  []mcms_types.BatchOperation      // Operations for MCMS proposals
}
```

### Changesets

A Changeset wraps a sequence with deployment environment context. It handles:
1. Reading addresses from a DataStore
2. Passing resolved addresses into sequences as input
3. Producing MCMS proposals from sequence output using `OutputBuilder`
4. Writing new addresses and metadata back to the DataStore

Changesets follow the `cldf.ChangeSetV2[Config]` pattern from chainlink-deployments-framework, with a `verify` function (validates config) and an `apply` function (executes the logic).

## Cross-Chain Dispatch Flow

When a changeset needs to operate across multiple chains, it follows this dispatch pattern:

1. The changeset receives a config with `map[uint64]PerChainConfig` (chain selector -> per-chain config)
2. For each chain selector, it resolves the chain family using `chain_selectors.GetSelectorFamily(chainSelector)`
3. It looks up the correct adapter from the registry using the family and version
4. It delegates to that adapter's sequence, passing chain-specific input
5. It aggregates outputs from all chains into a single `ChangesetOutput`

Example flow from `DeployContracts`:

```go
for chainSelector, chainConfig := range config.Chains {
    family, _ := chain_selectors.GetSelectorFamily(chainSelector)
    deployer, _ := registry.GetDeployer(family, chainConfig.Version)
    // Execute the deployer's sequence for this chain
    seq := deployer.DeployChainContracts()
    report, _ := operations.ExecuteSequence(bundle, seq, blockchains, input)
}
```

## DataStore Integration

The DataStore is the central persistence layer for deployment state. It stores contract addresses, metadata, and environment information.

### AddressRef

`datastore.AddressRef` is the universal contract pointer:

```go
type AddressRef struct {
    ChainSelector uint64
    Address       string
    Type          ContractType
    Version       *semver.Version
    Qualifier     string
}
```

### Flow

1. **Read**: Changesets read existing addresses from the DataStore to resolve dependencies (e.g., finding the Router address to pass to OffRamp deployment)
2. **Write**: Sequence outputs (`OnChainOutput.Addresses`) are written back to the DataStore after execution
3. **Metadata**: Contract, chain, and environment metadata is upserted via `WriteMetadataToDatastore`

## MCMS Proposal Construction

The `OutputBuilder` is a builder pattern for constructing changeset outputs that include MCMS (Multi-Chain Multi-Sig) proposals.

### Flow

1. Sequences return `BatchOperation` objects for write operations that require MCMS approval
2. The `OutputBuilder` collects all batch operations across sequences
3. On `Build(mcms.Input)`, it:
   - Resolves timelock addresses per chain via `MCMSReader.GetTimelockRef`
   - Fetches chain metadata via `MCMSReader.GetChainMetadata`
   - Constructs a `TimelockProposal` containing all batch operations
4. The proposal is included in the `ChangesetOutput.MCMSTimelockProposals`

```go
output, err := changesets.NewOutputBuilder(env, mcmsRegistry).
    WithReports(reports).
    WithBatchOps(batchOps).
    WithDataStore(ds).
    Build(mcmsInput)
```

The `MCMSReader` interface is what allows this to work across chain families -- each family implements its own way of resolving timelock addresses and chain metadata.

## Versioning Strategy

### Contract Versions

Multiple contract versions coexist in the system:
- **Version constants**: `Version_1_0_0`, `Version_1_5_0`, `Version_1_5_1`, `Version_1_6_0`, `Version_1_6_1` (defined in `utils/common.go`)
- **Adapter registration**: Each version registers its own adapter (e.g., EVM has adapters from v1_0_0 through v1_6_5)

### Token Adapter Versioning

Token adapters have a special versioning requirement: any token pool version must be connectable to any other. This means there is one `TokenAdapterRegistry` (not version-scoped), and each chain-family-version combination registers separately.

### Lane Adapter Versioning

Lane adapters are version-scoped -- a v1.6.0 lane and a v2.0.0 lane may have different configuration requirements. Separate registries can exist per major API version.

### Chain Family Selectors

Each chain family has a unique 4-byte selector used on-chain:

| Family | Selector | Constant |
|--------|----------|----------|
| EVM | `0x2812d52c` | `EVMFamilySelector` |
| SVM (Solana) | `0x1e10bdc4` | `SVMFamilySelector` |
| Aptos | `0xac77ffec` | `AptosFamilySelector` |
| TVM (TON) | `0x647e2ba9` | `TVMFamilySelector` |
| Sui | `0xc4e05953` | `SuiFamilySelector` |
