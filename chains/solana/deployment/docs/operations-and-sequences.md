---
title: "Solana Operations and Sequences Reference"
sidebar_label: "Operations & Sequences"
sidebar_position: 3
---

# Solana Operations and Sequences Reference

This document covers all Solana-specific operations (atomic on-chain interactions) and sequences (composed workflows) for the CCIP deployment tooling.

For the shared interfaces these implement, see [Interfaces Reference](../../../../deployment/docs/interfaces.md). For the adapter that exposes these, see [SolanaAdapter Reference](adapter.md).

---

## Table of Contents

- [Operations by Program](#operations-by-program)
- [Sequences](#sequences)
- [Utilities](#utilities)

---

## Operations by Program

Solana operations are organized by the on-chain program they interact with. Each operation directory contains deploy, initialize, and state-modifying operations specific to that program.

**Source:** [v1_6_0/operations/](../v1_6_0/operations/)

### Router Operations

**Source:** [v1_6_0/operations/router/](../v1_6_0/operations/router/)

| Operation | Description |
|-----------|-------------|
| `Deploy` | Deploys the Router program using `MaybeDeployContract` |
| `Initialize` | Initializes the Router with FeeQuoter, LINK token, and RMN Remote references |
| `ConnectChains` | Adds a destination chain config to the Router (source-side lane setup) |
| `AddOffRamp` | Registers an OffRamp for a remote chain on the Router (dest-side lane setup) |
| `RegisterTokenAdminRegistry` | Registers a token in the admin registry via the Router |
| `AcceptTokenAdminRegistry` | Accepts token admin registration |
| `SetPool` | Associates a token pool with a token on the Router and FeeQuoter |
| `TransferOwnership` | Proposes ownership transfer to a new owner |
| `AcceptOwnership` | Accepts a proposed ownership transfer |

### OffRamp Operations

**Source:** [v1_6_0/operations/offramp/](../v1_6_0/operations/offramp/)

| Operation | Description |
|-----------|-------------|
| `Deploy` | Deploys the OffRamp program |
| `Initialize` | Initializes with FeeQuoter, Router, and RMN Remote references |
| `InitializeConfig` | Sets OffRamp config (e.g., `EnableExecutionAfter` threshold) |
| `ConnectChains` | Adds a source chain config to the OffRamp (dest-side lane setup) |
| `SetOcr3` | Sets OCR3 configuration on the OffRamp |
| `TransferOwnership` | Proposes ownership transfer |
| `AcceptOwnership` | Accepts ownership transfer |

### FeeQuoter Operations

**Source:** [v1_6_0/operations/fee_quoter/](../v1_6_0/operations/fee_quoter/)

| Operation | Description |
|-----------|-------------|
| `Deploy` | Deploys the FeeQuoter program |
| `Initialize` | Initializes with max fee, Router, OffRamp, and LINK token |
| `AddPriceUpdater` | Adds the OffRamp as a price updater |
| `ConnectChains` | Adds a destination chain config on the FeeQuoter (source-side lane setup) |
| `SetTokenTransferFeeConfig` | Sets per-token transfer fee configuration for a destination chain |
| `TransferOwnership` | Proposes ownership transfer |
| `AcceptOwnership` | Accepts ownership transfer |

### Token Operations

**Source:** [v1_6_0/operations/tokens/](../v1_6_0/operations/tokens/)

| Operation | Description |
|-----------|-------------|
| `DeployLINK` | Deploys a LINK token (SPL token with configurable decimals and private key) |
| `DeploySolanaToken` | Deploys an SPL or SPL2022 token with optional ATAs and pre-minting |
| `UpsertTokenMetadata` | Uploads or updates token metadata on-chain |
| `CreateTokenMultisig` | Creates a multisig authority for token operations (used for customer mint authorities) |

### Token Pool Operations

**Source:** [v1_6_0/operations/token_pools/](../v1_6_0/operations/token_pools/)

| Operation | Description |
|-----------|-------------|
| `DeployBurnMint` | Deploys BurnMint token pool program |
| `DeployLockRelease` | Deploys LockRelease token pool program |
| `InitializeBurnMint` | Initializes a BurnMint pool account with Router and RMN references |
| `InitializeLockRelease` | Initializes a LockRelease pool account with Router and RMN references |
| `UpsertRemoteChainConfigBurnMint` | Configures remote chain on a BurnMint pool |
| `UpsertRemoteChainConfigLockRelease` | Configures remote chain on a LockRelease pool |
| `UpsertRateLimitsBurnMint` | Sets rate limits on a BurnMint pool |
| `UpsertRateLimitsLockRelease` | Sets rate limits on a LockRelease pool |
| `TransferOwnershipBurnMint` | Proposes ownership transfer for BurnMint pool |
| `TransferOwnershipLockRelease` | Proposes ownership transfer for LockRelease pool |
| `AcceptOwnershipBurnMint` | Accepts ownership for BurnMint pool |
| `AcceptOwnershipLockRelease` | Accepts ownership for LockRelease pool |
| `UpdateRateLimitAdminBurnMint` | Updates the rate limit admin on a BurnMint pool |
| `UpdateRateLimitAdminLockRelease` | Updates the rate limit admin on a LockRelease pool |

### MCMS Operations

**Source:** [v1_6_0/operations/mcms/](../v1_6_0/operations/mcms/)

| Operation | Description |
|-----------|-------------|
| `AccessControllerDeploy` | Deploys Access Controller program |
| `McmDeploy` | Deploys MCM (ManyChainMultiSig) program |
| `TimelockDeploy` | Deploys Timelock program |
| `InitAccessControllerOp` | Initializes an access controller account for a specific role |
| `InitMCMOp` | Initializes MCM for a specific config type (Proposer, Canceller, Bypasser) |
| `ConfigureMCMOp` | Reconfigures MCM after deployment (used in `FinalizeDeployMCMS`) |
| `InitTimelockOp` | Initializes the Timelock with minimum delay |
| `AddAccessOp` | Grants access roles on Timelock (Proposer, Executor, Canceller, Bypasser) |
| `TransferOwnershipOp` | Proposes ownership transfer for an ownable MCMS contract |
| `AcceptOwnershipOp` | Accepts ownership for an ownable MCMS contract |

### RMN Remote Operations

**Source:** [v1_6_0/operations/rmn_remote/](../v1_6_0/operations/rmn_remote/)

| Operation | Description |
|-----------|-------------|
| `Deploy` | Deploys the RMN Remote program |
| `Initialize` | Initializes the RMN Remote |
| `TransferOwnership` | Proposes ownership transfer |
| `AcceptOwnership` | Accepts ownership transfer |

### Test Receiver Operations

**Source:** [v1_6_0/operations/test_receiver/](../v1_6_0/operations/test_receiver/)

| Operation | Description |
|-----------|-------------|
| `Deploy` | Deploys the Test Receiver program |
| `Initialize` | Initializes with Router reference |

---

## Sequences

Sequences compose multiple operations into complete workflows.

**Source:** [v1_6_0/sequences/](../v1_6_0/sequences/)

### Deploy Chain Contracts

**Source:** [v1_6_0/sequences/deploy_chain_contracts.go](../v1_6_0/sequences/deploy_chain_contracts.go)

**Variable:** `DeployChainContracts`
**Input:** `ContractDeploymentConfigPerChainWithAddress`
**Output:** `OnChainOutput`

Deploys and initializes all CCIP infrastructure on a Solana chain:

1. Deploy LINK token (SPL token with configurable private key/decimals)
2. Deploy Router program
3. Deploy FeeQuoter program
4. Deploy OffRamp program
5. Deploy RMN Remote program
6. Deploy BurnMint Token Pool program
7. Deploy LockRelease Token Pool program
8. Initialize FeeQuoter (with max fee, Router, OffRamp, LINK references; add OffRamp as price updater)
9. Initialize Router (with FeeQuoter, LINK, RMN Remote references)
10. Initialize OffRamp (with FeeQuoter, Router, RMN Remote references; set config with execution threshold)
11. Initialize RMN Remote
12. **Extend OffRamp lookup table** with PDAs for all deployed programs (OffRamp config, FeeQuoter config, Router config, RMN Remote config, token pool addresses)
13. Deploy and initialize Test Receiver

### Lane Configuration

**Source:** [v1_6_0/sequences/connect_chains.go](../v1_6_0/sequences/connect_chains.go)

#### ConfigureLaneLegAsSource

Configures this chain as the source end of a lane.

**Input:** `UpdateLanesInput`

**Steps:**
1. Add destination chain config to FeeQuoter (translates shared `FeeQuoterDestChainConfig` to Solana-specific `fee_quoter.DestChainConfig`)
2. Add destination chain to Router (with allowlist configuration)

#### ConfigureLaneLegAsDest

Configures this chain as the destination end of a lane.

**Input:** `UpdateLanesInput`

**Steps:**
1. Add OffRamp to Router for the remote chain
2. Add source chain config to OffRamp (with source OnRamp address, enabled state, RMN verification flag)

### MCMS Deployment

**Source:** [v1_6_0/sequences/mcms.go](../v1_6_0/sequences/mcms.go)

#### DeployMCMS (Phase 1)

Deploys and initializes MCMS infrastructure:

1. Deploy Access Controller program
2. Deploy MCM program
3. Deploy Timelock program
4. Initialize Access Controller accounts (Proposer, Executor, Canceller, Bypasser)
5. Initialize MCM for each config type (Proposer, Canceller, Bypasser)
6. Initialize Timelock with minimum delay

#### FinalizeDeployMCMS (Phase 2)

Completes MCMS setup after deployment:

1. Configure MCM (set_config for Proposer, Canceller, Bypasser -- updates signer/quorum config)
2. Set up Timelock roles:
   - Proposer role: MCM proposer signer PDA
   - Executor role: Deployer key
   - Canceller role: Canceller PDA + Proposer PDA + Bypasser PDA
   - Bypasser role: Bypasser PDA

#### UpdateMCMSConfig

Updates config of specified MCMS contracts. Applies the same MCM config to Canceller, Bypasser, and Proposer for each specified MCM contract.

#### GrantAdminRoleToTimelock

Not implemented for Solana (returns `nil`).

### OCR3 Configuration

**Source:** [v1_6_0/sequences/ocr.go](../v1_6_0/sequences/ocr.go)

**Variable:** `SetOCR3Config`

Sets OCR3 configuration on the OffRamp. Resolves the OffRamp address from the DataStore and delegates to `offrampops.SetOcr3`.

### Fee Configuration

**Source:** [v1_6_0/sequences/fee_quoter.go](../v1_6_0/sequences/fee_quoter.go)

#### SetTokenTransferFeeConfig

**Variable:** `SetTokenTransferFeeConfig`
**Input:** `FeeQuoterSetTokenTransferFeeConfigSequenceInput`

Sets per-token transfer fee configs on the FeeQuoter. Iterates over remote chain configs and applies fee config for each token/destination pair.

### Token Sequences

**Source:** [v1_6_0/sequences/tokens.go](../v1_6_0/sequences/tokens.go)

#### DeployToken

Deploys an SPL or SPL2022 token:
1. Checks if token already exists in DataStore (skips deployment if found)
2. Creates token with optional private key, ATAs for senders, pre-mint amount, and freeze authority disable
3. Optionally uploads token metadata

#### DeployTokenPoolForToken

Initializes a token pool account for an existing token:
1. Initializes BurnMint or LockRelease pool with Router and RMN Remote references
2. Creates associated token account (ATA) for the pool signer PDA
3. For BurnMint pools: sets mint authority to pool signer PDA (if deployer owns it)

#### ConfigureTokenForTransfersSequence

Full token configuration workflow:
1. Register token in admin registry via Router
2. Accept token admin registration
3. Set token pool on Router (associates pool with token on Router and FeeQuoter)
4. For each remote chain:
   - Upsert remote chain config on token pool (BurnMint or LockRelease variant)
   - Set inbound/outbound rate limits

#### ManualRegistration

For externally-owned tokens where mint authority is unavailable:
1. Register token admin registry (optionally with external admin)
2. Initialize token pool (BurnMint or LockRelease)
3. Transfer token pool ownership to proposed owner
4. Optionally create token multisig with customer mint authorities + pool signer PDA

#### SetTokenPoolRateLimits

Sets rate limits on a BurnMint or LockRelease token pool for a specific remote chain.

#### UpdateAuthorities

Transfers token pool ownership to the timelock signer PDA:
1. Update rate limit admin to timelock signer
2. Transfer ownership to timelock signer
3. Accept ownership (immediate acceptance since deployer is current owner)

### Ownership Transfer

**Source:** [v1_6_0/sequences/transfer_ownership.go](../v1_6_0/sequences/transfer_ownership.go)

#### SequenceTransferOwnershipViaMCMS

Dispatches ownership transfer based on contract type:
- **Router, OffRamp, FeeQuoter, RMN Remote**: Direct per-program transfer operation
- **AccessController type**: Transfers all MCMS contracts including access controller accounts
- **RBACTimelock type**: Transfers MCMS contracts (Timelock, Proposer, Canceller, Bypasser MCM accounts)

#### SequenceAcceptOwnership

Mirrors `SequenceTransferOwnershipViaMCMS` but calls accept-ownership operations for each contract type.

---

## Utilities

**Source:** [utils/](../utils/)

### Contract Deployment (`deploy.go`)

| Function | Description |
|----------|-------------|
| `MaybeDeployContract` | Checks DataStore for existing contract; deploys via `chain.DeployProgram` only if not found |
| `DownloadSolanaCCIPProgramArtifacts` | Downloads pre-built program artifacts from GitHub releases |

### MCMS Helpers (`common.go`)

| Function | Description |
|----------|-------------|
| `BuildMCMSBatchOperation` | Converts Solana instructions into MCMS `BatchOperation` transactions |
| `GetTimelockSignerPDA` | Resolves the timelock signer PDA from DataStore refs |
| `GetMCMSignerPDA` | Resolves an MCM signer PDA for a given signer type |
| `FundSolanaAccounts` | Airdrops SOL to accounts (testnet utility) |
| `FundFromDeployerKey` | Transfers SOL from deployer to accounts |
| `FundFromAddressIxs` | Creates SOL transfer instructions |

### MCMS Ref Resolution (`mcms.go`)

| Function | Description |
|----------|-------------|
| `GetAllMCMS` | Returns all MCMS-related address refs for a chain (access controller + accounts + timelock + MCM accounts) |

### Program Utilities (`utils.go`)

| Function | Description |
|----------|-------------|
| `ExtendLookupTable` | Extends an OffRamp address lookup table with new entries (deduplicates) |
| `GetTokenProgramID` | Maps `SPLTokens`/`SPL2022Tokens` contract types to Solana program IDs |
| `GetSolProgramSize` | Gets the byte size of a deployed program |
| `GetSolProgramData` | Reads program data account (data type + address) |
| `GetTokenDecimals` | Reads token decimals from on-chain mint account |
| `GetTokenMintAuthority` | Reads the current mint authority of a token |
| `MintTokens` | Mints tokens to specified addresses via their ATAs |
| `DisableFreezeAuthority` | Permanently disables freeze authority on token mints |

### DataStore Utilities (`datastore.go`)

Provides Solana-specific address format conversion functions, including `ToByteArray` for converting DataStore address refs to `[]byte` via base58 decoding.

### Upgrade Authority (`upgrade_authority.go`)

Manages Solana program upgrade authority for production deployments.

---

## Operation Pattern

Solana operations follow the same framework as EVM but interact with Solana chains:

```go
var MyOperation = operations.NewOperation(
    "my-operation",
    semver.MustParse("1.6.0"),
    "Description of what this operation does",
    func(b operations.Bundle, chain cldf_solana.Chain, input MyInput) (MyOutput, error) {
        // Build Solana instruction
        ixn, err := program.NewMyInstruction(input.Param, /* accounts */).ValidateAndBuild()
        if err != nil {
            return MyOutput{}, err
        }

        // Execute on-chain
        if err := chain.Confirm([]solana.Instruction{ixn}); err != nil {
            return MyOutput{}, err
        }

        return MyOutput{/* ... */}, nil
    },
)
```

Key differences from EVM operations:
- Uses `cldf_solana.Chain` instead of EVM chain type
- Builds Solana `Instruction` objects (not EVM transactions)
- Executes via `chain.Confirm(instructions)` (not contract bindings)
- No automatic MCMS fallback -- MCMS batch ops are built explicitly via `BuildMCMSBatchOperation`
- Deploy operations use `MaybeDeployContract` which calls `chain.DeployProgram`
