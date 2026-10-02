---
title: "Architecture"
sidebar_label: "Architecture"
sidebar_position: 2
---

# Architecture

This page explains how the tooling API is put together: the layers, how a changeset finds the right chain-specific code, and how results end up in the DataStore and in MCMS proposals. It applies to both 1.6 and 2.0. Where they differ, the difference is called out.

## The idea in one paragraph

Changesets in `chainlink-ccip/deployment` are **chain-agnostic**. They never import EVM or Solana code. Instead, each chain family implements a set of Go **interfaces** ("adapters") and registers them in global **registries** from an `init()` function. At run time a changeset:

1. Takes a chain selector.
2. Looks up that chain's family (`evm`, `solana`, …).
3. Fetches the registered adapter.
4. Runs the **sequence** the adapter returns.

```
Changeset  (chainlink-ccip/deployment, chain-agnostic)
   │  family := chain_selectors.GetSelectorFamily(selector)
   ▼
Registry.Get(family[, version])
   ▼
Adapter    (chains/<family>/deployment, chain-specific)
   ▼
Sequence   ordered operations on one chain
   ▼
Operation  one side effect (deploy, write, read)
```

## Three levels: operations, sequences, changesets

| Level | What it is | Inputs | Output |
|---|---|---|---|
| **Operation** | A single side effect: deploy one contract, send one transaction, or do one read. Produces a report, so a retried run skips completed work | Serializable input + one chain | Typed output |
| **Sequence** | An ordered list of operations for **one chain** that completes a workflow ("deploy all contracts", "configure this chain for these lanes") | Serializable input + `cldf_chain.BlockChains` | `sequences.OnChainOutput` |
| **Changeset** | The entry point users run. Reads the DataStore, resolves adapters, runs sequences on every requested chain, and builds an MCMS proposal | `cldf.Environment` + config | `cldf.ChangesetOutput` |

Operations and sequences come from `chainlink-deployments-framework/operations`. Changesets implement `cldf.ChangeSetV2[Config]`, which is a `VerifyPreconditions` function plus an `Apply` function.

### OnChainOutput

Every sequence returns [`sequences.OnChainOutput`](../utils/sequences/sequences.go):

```go
type OnChainOutput struct {
    Addresses []datastore.AddressRef      // contracts deployed or managed
    Metadata  Metadata                    // contract / chain / env metadata to upsert
    BatchOps  []mcms_types.BatchOperation // writes that must go through MCMS
}
```

- `sequences.RunAndMergeSequence(b, chains, seq, input, agg)` runs a sub-sequence and merges its output into `agg`. Addresses, batch ops, and contract metadata are appended. Chain and env metadata may be set only once; a conflicting second value is an error.
- `sequences.WriteMetadataToDatastore(ds, metadata)` upserts the metadata. Upserts replace the whole record, so always write complete values.

## Registries

Each registry is a process-wide singleton (`sync.Once` + mutex), and each one has a `Get…Registry()` accessor. Adapters register from `init()`. The changeset caller only needs a blank import of the adapter package to trigger it:

```go
import _ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
```

Rules every registry follows:

- **Key.** The key is either `family` (`"evm"`) or `family-version` (`"evm-1.6.0"`, built by `utils.NewRegistererID`). For the key each registry uses, see [Interfaces](interfaces.md).
- **First registration wins.** A second registration for the same key is silently ignored. This makes registration safe to repeat, but if you get the key wrong, your adapter is silently never used.
- **Optional capabilities are type assertions, not registrations.** For example, `ConnectChains` checks `adapter.(lanes.TestRouterProvider)`. Implement the method on the registered struct and the changeset picks it up.

Most changesets accept their registries as arguments (`ConnectChains(lanes.GetLaneAdapterRegistry(), changesets.GetRegistry(), …)`), so tests can pass in isolated registries.

## v1 and v2 dispatch differ

| | v1 (1.6) | v2 (2.0) |
|---|---|---|
| Lane model | **Leg by leg**: `ConfigureLaneLegAsSource` on A and `ConfigureLaneLegAsDest` on B, then the reverse | **Chain by chain**: `ConfigureChainForLanes` on each chain with all its remotes at once |
| Lane interface | `lanes.LaneAdapter` | `v2_0_0/adapters.ChainFamily` |
| Deploy interface | `deploy.Deployer` | `v2_0_0/adapters.DeployChainContractsAdapter` |
| Registry key | `family-version` (contract version) | `family` (versions are picked inside the adapter) |
| Config input | Per-lane structs (`ConnectChainsConfig`) | **Topology** (`offchain.EnvironmentTopology`: NOPs, committees, executor pools) + lane pairs |
| Verification | OCR3 commit/exec DONs (`SetOCR3Config`) | Committee verifiers (CCVs) + executors, derived from topology |
| Changesets live in | `deploy/`, `lanes/`, `fees/`, … | `v2_0_0/changesets/` |

Some parts are **shared by both**: MCMS readers, ownership transfer, token adapters (keyed by pool version, with 2.0 pools as just another version), curse, fee aggregator, authorized callers, test adapters, and hooks. See [Interfaces → Shared](interfaces.md#shared-by-v1-and-v2).

## DataStore

The DataStore holds deployment state. Its main record is `datastore.AddressRef`:

```go
type AddressRef struct {
    ChainSelector uint64
    Address       string            // family-native string (hex on EVM, base58 on Solana)
    Type          ContractType      // e.g. "OnRamp", "RBACTimelock"
    Version       *semver.Version
    Qualifier     string            // disambiguates multiple instances
    Labels        ...
}
```

- **Read.** Changesets look up existing addresses and pass them into sequences (`ExistingAddresses`, resolved refs, or raw bytes).
- **Write.** The changeset adds every `OnChainOutput.Addresses` entry to a fresh `MemoryDataStore` and returns it in `ChangesetOutput.DataStore`. Sequences never write the DataStore directly.
- **Find helpers** live in [`utils/datastore`](../utils/datastore/datastore.go): `FindAndFormatRef`, `FindAndFormatCanonicalRef`, and others. Pass a family formatter (for example EVM `ToEVMAddressBytes`) to get bytes back.

Shared qualifiers ([utils/common.go](../utils/common.go)):

| Constant | Value | Selects |
|---|---|---|
| `CLLQualifier` | `CLLCCIP` | The CLL-managed MCMS + timelock (default owner of CCIP contracts) |
| `RMNTimelockQualifier` | `RMNMCMS` | The RMN MCMS + timelock |
| `UltraFastCurseMCMSQualifier` | `UltraFastCurse` | The dedicated fast-curse MCMS |

## MCMS proposals

Contracts are usually owned by an RBAC timelock, not by the deployer key. An operation that writes to such a contract must not send the transaction. Instead it returns a `mcms_types.BatchOperation`. The rule every family follows:

> If the contract's owner/authority is the deployer key, execute the transaction directly. Otherwise, return it as a `BatchOperation`.

The changeset collects the batch ops from every chain and builds a single timelock proposal:

```go
return changesets.NewOutputBuilder(e, mcmsRegistry).
    WithReports(reports).
    WithBatchOps(batchOps).   // empty batch ops are dropped
    WithDataStore(ds).
    Build(cfg.MCMS)           // no batch ops → no proposal
```

During `Build`, each chain in the batch ops goes through these steps:

1. Its family's `MCMSReader` is looked up.
2. `GetTimelockRef` and `GetChainMetadata` are called (MCM address and starting op count).
3. Everything is assembled into a `TimelockProposal`.

This is why **every chain family must register an `MCMSReader`**. Without one, no proposal can include that chain.

### mcms.Input

[`mcms.Input`](../utils/mcms/mcms.go) is embedded in every changeset config that can produce a proposal:

| Field | Meaning |
|---|---|
| `TimelockAction` | `schedule` (default), `bypass`, or `cancel` |
| `ValidUntil` | Unix expiry. If unset, it defaults to a randomized time in the next ~24h, so identical payloads get distinct operation IDs |
| `Qualifier` | Which MCMS/timelock pair to use (see qualifiers above) |
| `OverridePreviousRoot` | Replace an unexecuted root on the MCM |
| `Description` | Shown to signers |
| `TimelockDelay` | **Deprecated, ignored.** The delay is read from chain |

### Building a changeset from one sequence

For a changeset that runs one sequence on one chain, use [`changesets.NewFromOnChainSequence`](../utils/changesets/changesets.go). You supply the sequence, a function that turns the user config into sequence input (`ResolveInput`), and a function that resolves the chain dependency (`ResolveDep`). The helper does the rest: it executes the sequence, writes the DataStore, and builds the proposal. The resulting changeset takes `changesets.WithMCMS[Cfg]{MCMS, Cfg}`. EVM 2.0's own changesets (for example `chains/evm/deployment/v2_0_0/changesets/deploy_chain_contracts.go`) use this pattern.

## Versioning

The split that matters is **v1 vs v2**:

- **v1 (1.x, 1.6 being the current one)** uses the `deploy`/`lanes`/`fees` interfaces, keyed by `family-version`. A family registers under whichever 1.x versions it supports. Minor-version differences stay inside the family's adapter.
- **v2 (2.0)** uses the `v2_0_0/adapters` interfaces, keyed by `family`. The adapter picks contract versions itself in `GetDefaultDeployContractParams`.

Two exceptions:

- **Token adapters** are keyed by *pool* version (for example, a v1 pool and a v2 pool are separate adapters). Every pool version must be able to connect to every other, so a 1.6 pool can be connected to a 2.0 pool.
- **Version-agnostic helpers** (MCMS reader, address normalizer, fee resolver, TAR manager) live in a `v1_0_0` package by convention, and are registered by family only.

Version constants are in [utils/common.go](../utils/common.go) (`utils.Version_1_6_0`, `utils.Version_2_0_0`, …).

## Chain family selectors

Every family has a 4-byte on-chain selector, which the FeeQuoter uses. A family registers its selector by implementing `GetChainFamilySelector() [4]byte` on its `LaneAdapter` (1.6) or `ChainFamily` (2.0), or by calling `utils.RegisterChainFamilySelector` directly. `utils.GetSelectorHex(chainSelector)` returns it.

| Family | Selector |
|---|---|
| EVM | `0x2812d52c` |
| SVM (Solana) | `0x1e10bdc4` |
| Aptos | `0xac77ffec` |
| TVM (TON) | `0x647e2ba9` |
| Sui | `0xc4e05953` |

Aptos, TON, and Sui still have hardcoded fallbacks in `GetSelectorHex`. A new family must register its selector itself.
