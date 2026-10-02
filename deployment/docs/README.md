---
title: "CCIP Deployment Tooling API"
sidebar_label: "Overview"
sidebar_position: 1
---

# CCIP Deployment Tooling API

The tooling API is a Go library for deploying and operating CCIP on any chain family: EVM, Solana, Aptos, TON, Sui, and others.

- **`chainlink-ccip/deployment`** (this module) holds the chain-agnostic **changesets** you run, and the **interfaces** each chain family implements.
- **Chain families** implement those interfaces next to their contracts:
  - EVM: `chains/evm/deployment`
  - Solana v1: `chains/solana/deployment`
  - Solana v2: the `chainlink-ccip-solana` repo
  - Other families: their own repos
- Families register their implementations in global registries at `init()`. A changeset looks up the right implementation from the chain selector.

## v1 vs v2

There are two API generations. They share MCMS, ownership, tokens, fees, curse, and test plumbing, but deploy contracts and configure lanes differently.

| | **v1 (CCIP 1.6)** | **v2 (CCIP 2.0)** |
|---|---|---|
| Verification | OCR3 DONs | Committee verifiers (CCVs) + executors |
| Input | Per-chain / per-lane config | Environment **topology** (NOPs, committees, executor pools) + lane pairs |
| Lane setup | `lanes.ConnectChains` → `LaneAdapter` (source leg + dest leg) | `ConfigureChainsForLanesFromTopology` → `ChainFamily` (one call per chain, all remotes) |
| Deploy | `deploy.DeployContracts` → `Deployer` | `v2_0_0/changesets.DeployChainContracts` → `DeployChainContractsAdapter` |
| Code | `deploy/`, `lanes/`, `fees/`, … | `v2_0_0/` |

## Docs

| Page | Read it when you want to… |
|---|---|
| [Architecture](architecture.md) | Understand operations / sequences / changesets, registries, DataStore, MCMS proposals |
| [Consuming](consuming.md) | Run changesets: wiring, v1 and v2 examples, v1-vs-v2 per flow, full changeset catalog |
| [Implementing 1.6](implementing-1.6.md) | Add a chain family to the v1 API |
| [Implementing 2.0](implementing-2.0.md) | Add a chain family to the v2 API |
| [Interfaces](interfaces.md) | Look up an interface, its registry, and its key |
| [Changeset Style Guide](style-guide.md) | Write or review a changeset |

## Package map

| Package | Contents |
|---|---|
| `deploy/` | v1 deployer, MCMS deploy, OCR3, ownership, lane migration, FeeQuoter upgrade, address normalizer |
| `lanes/` | v1 `LaneAdapter`, `ConnectChains`, `DisableLane`, PingPong |
| `tokens/` | `TokenAdapter` + optional token interfaces, all token changesets (v1 and v2 pools) |
| `fees/` | Fee adapters, fee aggregator, token transfer fee, FeeQuoter dest changesets |
| `fastcurse/` | RMN curse adapters and changesets |
| `authorizedcallers/` | Authorized-callers adapter and changeset |
| `hooks/` | Pre/post/post-proposal hooks (verification, ownership, CCIP send, lane sanity) |
| `testadapters/` | Cross-chain message test adapters |
| `v2_0_0/adapters` | v2 interfaces: `ChainFamily`, `DeployChainContractsAdapter`, `CommitteeVerifierContractAdapter`, CCTP, Lombard, OnRamp upgrade, test verifier |
| `v2_0_0/changesets` | v2 changesets |
| `v2_0_0/offchain` | `EnvironmentTopology` and JD helpers |
| `utils/` | Versions, contract types, qualifiers, `OutputBuilder`, `OnChainOutput`, `mcms.Input`, datastore helpers |

## Agent skill

The [`tooling-api` skill](../.agents/skills/tooling-api/SKILL.md) walks an AI agent through implementing these interfaces for a chain family, adding a new interface, or adding a new changeset.
