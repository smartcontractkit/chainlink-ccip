---
title: "MCMS and Utilities Reference"
sidebar_label: "MCMS & Utilities"
sidebar_position: 7
---

# MCMS and Utilities Reference

This document covers the Multi-Chain Multi-Sig (MCMS) governance integration and shared utility functions used throughout the CCIP Deployment Tooling API.

For overall system design, see [Architecture Guide](architecture.md). For the full list of adapter interfaces including `MCMSReader`, see [Interfaces Reference](interfaces.md).

## MCMS Overview

MCMS (Multi-Chain Multi-Sig) is the governance layer for contract operations across chains. It allows write operations to be:

1. **Proposed** -- bundled into a `TimelockProposal` containing batch operations across one or more chains.
2. **Approved** -- signed by the required number of multi-sig signers.
3. **Executed** -- carried out via a timelock contract after an enforced delay.

When a changeset produces write operations that the deployer key cannot execute directly (because the contracts are owned by the timelock), those operations are collected as `BatchOperation` entries and assembled into a proposal. The `OutputBuilder` and `MCMSReader` interface handle this assembly automatically.

## mcms.Input

`mcms.Input` configures how an MCMS proposal is constructed. Every changeset that may produce governance operations accepts this as part of its config.

**Source:** [utils/mcms/mcms.go](../utils/mcms/mcms.go)

```go
type Input struct {
    OverridePreviousRoot bool
    ValidUntil           uint32
    TimelockDelay        mcms_types.Duration
    TimelockAction       mcms_types.TimelockAction
    Qualifier            string
    Description          string
}
```

| Field | Type | Description |
|-------|------|-------------|
| `OverridePreviousRoot` | `bool` | When `true`, overrides the existing root of the MCMS contract. Set this when a previous proposal was not executed and its root should be replaced. |
| `ValidUntil` | `uint32` | Unix timestamp after which the proposal can no longer be set or executed. Acts as an expiration deadline. |
| `TimelockDelay` | `mcms_types.Duration` | Minimum wait time between scheduling an operation and executing it. Enforced on-chain by the timelock contract. |
| `TimelockAction` | `mcms_types.TimelockAction` | One of `schedule`, `bypass`, or `cancel`. Controls what the timelock does with the operations: queue them for delayed execution, execute immediately (bypasser role), or cancel previously scheduled operations. |
| `Qualifier` | `string` | Qualifies which MCMS + Timelock contract addresses to use. Allows multiple MCMS deployments to coexist (e.g., `"CLLCCIP"` for CLL-managed CCIP contracts, `"RMNMCMS"` for RMN-specific governance). |
| `Description` | `string` | Human-readable description included in the proposal for review by signers. |

Validation enforces that `TimelockAction` is one of the three allowed values.

## MCMSReader Interface

The `MCMSReader` interface resolves chain-specific MCMS governance metadata. Each chain family registers one reader (keyed by `chainFamily` only, not versioned).

**Source:** [utils/changesets/output.go](../utils/changesets/output.go)
**Registry:** `MCMSReaderRegistry` via `changesets.GetRegistry()`
**Full definition:** See [Interfaces Reference -- MCMSReader](interfaces.md#mcmsreader)

```go
type MCMSReader interface {
    GetChainMetadata(e Environment, chainSelector uint64, input mcms.Input) (mcms_types.ChainMetadata, error)
    GetTimelockRef(e Environment, chainSelector uint64, input mcms.Input) (datastore.AddressRef, error)
    GetMCMSRef(e Environment, chainSelector uint64, input mcms.Input) (datastore.AddressRef, error)
}
```

Changesets use `MCMSReader` indirectly through `OutputBuilder`. When `Build()` is called, the builder iterates over the batch operations, determines each chain's family from its selector, looks up the registered reader, and calls:

- `GetTimelockRef` -- to resolve the timelock contract address that will execute the operations.
- `GetChainMetadata` -- to fetch the starting operation count and MCM address needed by the proposal builder.

The `MCMSReaderRegistry` is a thread-safe singleton. Readers register in `init()`:

```go
changesets.GetRegistry().RegisterMCMSReader(chain_selectors.FamilyEVM, &EVMMCMSReader{})
```

## OutputBuilder

`OutputBuilder` is a builder pattern that assembles a `ChangesetOutput` from sequence execution results, optionally including an MCMS `TimelockProposal` when there are batch operations to govern.

**Source:** [utils/changesets/output.go](../utils/changesets/output.go)

### Construction

```go
output, err := changesets.NewOutputBuilder(env, mcmsRegistry).
    WithReports(reports).
    WithBatchOps(batchOps).
    WithDataStore(ds).
    Build(mcmsInput)
```

### Methods

| Method | Signature | Description |
|--------|-----------|-------------|
| `NewOutputBuilder` | `NewOutputBuilder(e Environment, registry *MCMSReaderRegistry) *OutputBuilder` | Creates a new builder bound to an environment and MCMS reader registry. |
| `WithReports` | `WithReports(reports []operations.Report[any, any]) *OutputBuilder` | Attaches execution reports to the output. Reports contain operation-level traces for debugging and retry. |
| `WithBatchOps` | `WithBatchOps(ops []mcms_types.BatchOperation) *OutputBuilder` | Sets the MCMS batch operations. Automatically filters out any `BatchOperation` entries with zero transactions. |
| `WithDataStore` | `WithDataStore(ds datastore.MutableDataStore) *OutputBuilder` | Attaches a `MutableDataStore` containing newly deployed addresses and metadata. |
| `Build` | `Build(input mcms.Input) (ChangesetOutput, error)` | Finalizes the output. If batch operations exist, resolves timelock addresses and chain metadata via the `MCMSReaderRegistry`, then constructs and attaches a `TimelockProposal`. If no batch operations are present, returns the output without a proposal. |

### Build Internals

When `Build` is called with non-empty batch operations:

1. For each unique chain selector in the batch operations, the builder resolves the chain family and retrieves the registered `MCMSReader`.
2. `getTimelockAddresses` calls `GetTimelockRef` for each chain to build a `map[ChainSelector]string` of timelock addresses.
3. `getChainMetadata` calls `GetChainMetadata` for each chain to build a `map[ChainSelector]ChainMetadata`.
4. These maps are passed to `mcms.NewTimelockProposalBuilder` along with the `mcms.Input` fields to construct the final `TimelockProposal`.

## OnChainOutput

`OnChainOutput` is the standard return type for sequences that deploy contracts or perform on-chain write operations. It aggregates all artifacts produced during sequence execution.

**Source:** [utils/sequences/sequences.go](../utils/sequences/sequences.go)

```go
type OnChainOutput struct {
    Addresses []datastore.AddressRef
    Metadata  Metadata
    BatchOps  []mcms_types.BatchOperation
}

type Metadata struct {
    Contracts []datastore.ContractMetadata
    Chain     *datastore.ChainMetadata
    Env       *datastore.EnvMetadata
}
```

| Field | Description |
|-------|-------------|
| `Addresses` | Contract addresses deployed or managed by the sequence. Each `AddressRef` includes chain selector, address, contract type, version, and optional qualifier. |
| `Metadata.Contracts` | Per-contract metadata entries. Keyed by address + chain selector in the datastore (upsert semantics). |
| `Metadata.Chain` | Per-chain metadata. At most one per sequence; keyed by chain selector. |
| `Metadata.Env` | Per-environment metadata. At most one per environment. |
| `BatchOps` | Ordered list of `BatchOperation` entries for MCMS proposal construction. Each batch operation is executed atomically. Order is preserved during proposal assembly. |

## Sequence Utilities

### RunAndMergeSequence

Composes sub-sequences by executing them and merging their `OnChainOutput` into an aggregator. This is the primary mechanism for building complex workflows from smaller, focused sequences.

**Source:** [utils/sequences/sequences.go](../utils/sequences/sequences.go)

```go
func RunAndMergeSequence[IN any](
    b operations.Bundle,
    chains cldf_chain.BlockChains,
    seq *operations.Sequence[IN, OnChainOutput, cldf_chain.BlockChains],
    input IN,
    agg OnChainOutput,
) (OnChainOutput, error)
```

Merge behavior:

| Field | Strategy |
|-------|----------|
| `BatchOps` | Appended to the aggregator's list. |
| `Addresses` | Appended to the aggregator's list. |
| `Metadata.Contracts` | Appended to the aggregator's list. |
| `Metadata.Chain` | Set if aggregator has none; returns an error if both the aggregator and the sub-sequence provide conflicting chain metadata. |
| `Metadata.Env` | Set if aggregator has none; returns an error on conflict. |

### WriteMetadataToDatastore

Persists metadata from an `OnChainOutput` into a `MutableDataStore`.

**Source:** [utils/sequences/sequences.go](../utils/sequences/sequences.go)

```go
func WriteMetadataToDatastore(ds datastore.MutableDataStore, metadata Metadata) error
```

Upsert behavior:

- **Contract metadata**: Each entry is upserted individually. Key = address + chain selector.
- **Chain metadata**: Upserted if non-nil. Key = chain selector.
- **Env metadata**: Set if non-nil. One record per environment.

Because upsert replaces the entire record for a given key, callers must include all required fields when writing to a key that may already exist.

## Common Utility Functions

**Source:** [utils/common.go](../utils/common.go)

### NewRegistererID

Creates a composite registry key from a chain family and semver version:

```go
func NewRegistererID(chainFamily string, version *semver.Version) string
```

Returns `fmt.Sprintf("%s-%s", chainFamily, version.String())`, e.g., `"evm-1.6.0"` or `"solana-1.6.0"`.

### NewIDFromSelector

Resolves the chain family from a numeric chain selector and creates the registry key:

```go
func NewIDFromSelector(chainSelector uint64, version *semver.Version) string
```

Uses `chain_selectors.GetSelectorFamily` to determine the family, then delegates to `NewRegistererID`. Panics if the chain selector is invalid.

### GetSelectorHex

Returns the 4-byte on-chain family selector for a given chain selector:

```go
func GetSelectorHex(selector uint64) []byte
```

Maps chain families to their on-chain identifiers (derived from `keccak256` of the family name):

| Family | Hex Constant | Value |
|--------|-------------|-------|
| EVM | `EVMFamilySelector` | `0x2812d52c` |
| SVM (Solana) | `SVMFamilySelector` | `0x1e10bdc4` |
| Aptos | `AptosFamilySelector` | `0xac77ffec` |
| TVM (TON) | `TVMFamilySelector` | `0x647e2ba9` |
| Sui | `SuiFamilySelector` | `0xc4e05953` |

These selectors are defined in the CCIP Solidity library (`Internal.sol`) and must match across all chain families.

## Version Constants

**Source:** [utils/common.go](../utils/common.go)

Pre-parsed `semver.Version` values used for adapter registration and contract version identification:

| Constant | Value |
|----------|-------|
| `Version_1_0_0` | `1.0.0` |
| `Version_1_5_0` | `1.5.0` |
| `Version_1_5_1` | `1.5.1` |
| `Version_1_6_0` | `1.6.0` |
| `Version_1_6_1` | `1.6.1` |

These are created with `semver.MustParse` and are safe to use as map keys and in comparisons.

## Contract Type Constants

**Source:** [utils/common.go](../utils/common.go)

Shared `cldf.ContractType` strings used across all chain families for datastore lookups and adapter registration:

| Constant | String Value | Description |
|----------|-------------|-------------|
| `BypasserManyChainMultisig` | `"BypasserManyChainMultiSig"` | MCMS bypasser multi-sig contract |
| `CancellerManyChainMultisig` | `"CancellerManyChainMultiSig"` | MCMS canceller multi-sig contract |
| `ProposerManyChainMultisig` | `"ProposerManyChainMultiSig"` | MCMS proposer multi-sig contract |
| `RBACTimelock` | `"RBACTimelock"` | Role-based access control timelock |
| `CallProxy` | `"CallProxy"` | Call proxy contract |
| `CapabilitiesRegistry` | `"CapabilitiesRegistry"` | DON capabilities registry |
| `CCIPHome` | `"CCIPHome"` | CCIP home contract (chain config) |
| `RMNHome` | `"RMNHome"` | RMN home contract |
| `BurnMintTokenPool` | `"BurnMintTokenPool"` | Burn-and-mint token pool |
| `LockReleaseTokenPool` | `"LockReleaseTokenPool"` | Lock-and-release token pool |
| `TokenPoolLookupTable` | `"TokenPoolLookupTable"` | Token pool lookup table |
| `BurnWithFromMintTokenPool` | `"BurnWithFromMintTokenPool"` | Burn-with-from-mint token pool |
| `BurnFromMintTokenPool` | `"BurnFromMintTokenPool"` | Burn-from-mint token pool |
| `CCTPTokenPool` | `"CCTPTokenPool"` | CCTP (Circle) token pool |

Additional qualifier constants:

| Constant | Value | Purpose |
|----------|-------|---------|
| `CLLQualifier` | `"CLLCCIP"` | Qualifies CLL-managed CCIP contract addresses |
| `RMNTimelockQualifier` | `"RMNMCMS"` | Qualifies RMN-specific MCMS contract addresses |

These qualifiers are passed via `mcms.Input.Qualifier` to scope timelock and multi-sig resolution to the correct deployment set.

---

Cross-reference [Architecture Guide](architecture.md) for how MCMS proposal construction fits into the changeset execution flow, and [Interfaces Reference](interfaces.md) for the full `MCMSReader` interface definition and all other adapter interfaces.
