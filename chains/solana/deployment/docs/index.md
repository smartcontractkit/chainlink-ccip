---
title: "Solana Implementation"
sidebar_label: "Overview"
sidebar_position: 1
---

# Solana Deployment Implementation

This documentation covers the Solana-specific implementation of the CCIP Deployment Tooling API. The Solana adapter handles the unique requirements of the Solana Virtual Machine (SVM), including account-based program deployment, PDA-derived addresses, and two-phase MCMS initialization.

For the shared interfaces this implementation fulfills, see [Interfaces Reference](../../../../deployment/docs/interfaces.md). For the shared types, see [Types Reference](../../../../deployment/docs/types.md).

---

## Package Layout

```
chains/solana/deployment/
├── utils/
│   ├── common.go                     # Contract types, PDA helpers, MCMS builders
│   ├── deploy.go                     # MaybeDeployContract, artifact download
│   ├── mcms.go                       # GetAllMCMS helper
│   ├── utils.go                      # Lookup tables, token utilities, program data
│   ├── upgrade_authority.go          # Upgrade authority management
│   ├── datastore.go                  # Address format conversions (base58)
│   ├── sequences.go                  # Shared sequence utilities
│   └── artifact_versions.go          # Program artifact version mappings
├── v1_6_0/
│   ├── adapters/
│   │   └── init.go                   # CurseAdapter registration
│   ├── operations/
│   │   ├── router/                   # Router program operations
│   │   ├── offramp/                  # OffRamp program operations
│   │   ├── fee_quoter/              # FeeQuoter program operations
│   │   ├── tokens/                   # Token deployment operations
│   │   ├── token_pools/             # Token pool operations
│   │   ├── mcms/                     # MCMS operations
│   │   ├── rmn_remote/              # RMN Remote operations
│   │   └── test_receiver/           # Test receiver operations
│   └── sequences/
│       ├── adapter.go                # SolanaAdapter struct + init() registration
│       ├── deploy_chain_contracts.go # Full chain deployment sequence
│       ├── connect_chains.go         # Lane configuration sequences
│       ├── mcms.go                   # MCMS deploy + finalize sequences
│       ├── ocr.go                    # OCR3 configuration
│       ├── fee_quoter.go            # Fee configuration sequences
│       ├── tokens.go                 # Token deployment + configuration
│       └── transfer_ownership.go     # Ownership transfer sequences
└── idl/                              # Anchor IDL definitions
```

## Key Differences from EVM

| Aspect | EVM | Solana |
|--------|-----|--------|
| **Address format** | Hex (20 bytes) | Base58 (32 bytes, `solana.PublicKey`) |
| **Adapter state** | Stateless empty struct | Stateful: `timelockAddr map[uint64]solana.PublicKey` |
| **Router = OnRamp** | Separate OnRamp contract | Router serves as OnRamp (`GetOnRampAddress` delegates to `GetRouterAddress`) |
| **Contract deployment** | Single-step `CREATE`/`CREATE2` | Program deployment with `DeployProgram` + separate initialization |
| **Address derivation** | Deployed address known at creation | PDAs derived from seeds (`FindConfigPDA`, `FindOfframpConfigPDA`, etc.) |
| **MCMS deployment** | Single-phase (EVM returns no-op for finalize) | Two-phase: `DeployMCMS` deploys + initializes, `FinalizeDeployMCMS` configures + sets up roles |
| **Token standards** | ERC-20 only | SPL Tokens + SPL 2022 Tokens |
| **Token pool deployment** | New contract deployed per pool | Token pools are program accounts (initialized, not deployed) |
| **Lookup tables** | Not needed | Address lookup tables required for OffRamp (extended during deployment) |
| **Ownership acceptance** | Atomic with transfer when caller is deployer | `ShouldAcceptOwnershipWithTransferOwnership` returns true when current owner is deployer key |
| **MCMS contract types** | Single MCM, Timelock | Access Controller + MCM + Timelock with PDA seeds (Proposer, Canceller, Bypasser, RBACTimelock) |
| **Token configuration** | Register via TokenAdminRegistry contract | Register via Router's TokenAdminRegistry + initialize pool with Router/RMN references |
| **Associated Token Accounts** | Not applicable | ATAs must be created for token pool signers before use |

## Documentation

- [SolanaAdapter Reference](adapter.md) -- adapter struct, interface implementations, registration
- [Operations and Sequences Reference](operations-and-sequences.md) -- all operations by program, all sequences, utilities
