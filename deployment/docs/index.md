---
title: "CCIP Deployment Tooling API"
sidebar_label: "Overview"
sidebar_position: 1
---

# CCIP Deployment Tooling API

The CCIP Deployment Tooling API provides a unified, strongly-typed Go-based operational layer for deploying, configuring, and managing CCIP across all chain families (EVM, Solana, Aptos, TON, Sui, etc.).

This shared library (`chainlink-ccip/deployment`) contains all generic, chain-agnostic types, registries, and utilities. Each chain family implements its own adapter alongside its contracts, either in `chainlink-ccip/chains/<chain>/deployment` or in a dedicated repository.

## Architecture Overview

```mermaid
flowchart LR
    subgraph chainlink-deployments
        dp[Durable Pipelines]
        ds[DataStore]
    end

    subgraph Chain Family Repo
        cfmr[MCMSReader]
        subgraph Sequences
            subgraph v1.6.0
                clls1_6[ConfigureLaneLegAsSource]
                clld1_6[ConfigureLaneLegAsDest]
                cllb1_6[ConfigureLaneLegBidirectionally]
                ctft1_6[ConfigureTokenForTransfers]
            end
            subgraph v2.0.0
                clls1_7[ConfigureChainForLanes]
                ctft1_7[ConfigureTokenForTransfers]
            end
            subgraph v1.5.0
                ctft1_5[ConfigureTokenForTransfers]
            end
        end
        subgraph Helpers
            artb[AddressRefToBytes]
            dtfp[DeriveTokenFromPool]
        end
    end

    subgraph chainlink-ccip/deployment
        subgraph changesets
            ctft[ConfigureTokensForTransfers]
            cc1_6[v1_6.ConnectChains]
            cc1_7[v1_7.ConnectChains]
        end
        subgraph interfaces
            mr[MCMSReader]
            ta[TokensAdapter]
            ca1_6[v1_6.ChainAdapter]
            ca1_7[v1_7.ChainAdapter]
        end
        subgraph registries
            tar[TokenAdapterRegistry]
            car1_6[v1_6.ChainAdapterRegistry]
            car1_7[v1_7.ChainAdapterRegistry]
            mrr[MCMSReaderRegistry]
        end
    end

    ctft --> tar
    ctft --> mrr

    tar --> ta
    mrr --> mr

    ta --> artb
    ta --> dtfp
    ta --> ctft1_5
    ta --> ctft1_6
    ta --> ctft1_7
    mr ------> cfmr

    dp ---init---> registries
    dp ---run---> changesets

    cc1_6 --> mrr & car1_6
    car1_6 --> ca1_6
    ca1_6 ------> cllb1_6
    ca1_6 ------> clls1_6
    ca1_6 ------> clld1_6
```

## Three-Level Hierarchy

The API is structured in three levels of granularity:

| Level | Description | Use When |
|-------|-------------|----------|
| **Changesets** | Environment-aware entry points that read from DataStore, invoke sequences, and produce MCMS proposals | Executing operations via Durable Pipelines or full deployment environments |
| **Sequences** | Ordered collections of operations. Accept serializable input and minimal dependencies | Completing an operational workflow without a full deployment environment |
| **Operations** | Single side-effect actions (deploy, read, write). Produce reports for stateful retries | Making a single contract call or deployment |

## Package Layout

| Package | Purpose |
|---------|---------|
| `deploy/` | Contract deployment, MCMS deployment, OCR3 config, ownership transfer changesets and interfaces |
| `lanes/` | Lane configuration and inter-chain connection changesets and interfaces |
| `tokens/` | Token pool configuration, expansion, manual registration, rate limits |
| `fees/` | Fee configuration and token transfer fee management |
| `fastcurse/` | RMN curse/uncurse operations |
| `utils/changesets/` | `MCMSReader` interface, `OutputBuilder`, changeset utilities |
| `utils/sequences/` | `OnChainOutput` type, sequence execution utilities |
| `utils/mcms/` | MCMS input types |
| `utils/` | Common types, version constants, contract type constants |
| `testadapters/` | Test adapter framework for cross-chain message testing |

## Documentation

| Document | Description |
|----------|-------------|
| [Architecture](architecture.md) | Design principles, adapter-registry pattern, dispatch flow, DataStore and MCMS integration |
| [Interfaces](interfaces.md) | Complete API reference for all adapter interfaces and their registries |
| [Types](types.md) | All input/output types, config structs, and constants |
| [Changesets](changesets.md) | Reference for all changesets (entry points) with config types and usage |
| [Implementing Adapters](implementing-adapters.md) | Step-by-step guide for adding a new chain family |
| [MCMS and Utilities](mcms-and-utilities.md) | MCMS integration, `OutputBuilder`, DataStore and sequence utilities |

## Chain-Specific Documentation

| Chain Family | Documentation |
|-------------|---------------|
| EVM | [EVM Deployment Docs](../../chains/evm/deployment/docs/index.md) |
| Solana | [Solana Deployment Docs](../../chains/solana/deployment/docs/index.md) |
