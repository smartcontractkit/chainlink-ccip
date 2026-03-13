---
title: "EVM Sequences Reference"
sidebar_label: "Sequences"
sidebar_position: 4
---

# EVM Sequences Reference

Sequences compose multiple operations into complete workflows. Each sequence is defined as a package-level variable using `operations.NewSequence`.

**Source:** [v1_6_0/sequences/](../v1_6_0/sequences/)

---

## Sequence Catalog

### Deploy Chain Contracts

**Source:** [v1_6_0/sequences/deploy_chain_contracts.go](../v1_6_0/sequences/deploy_chain_contracts.go)

Deploys all CCIP infrastructure contracts on a single EVM chain. This is the primary deployment sequence.

**Variable:** `DeployChainContracts`
**Input:** `ContractDeploymentConfigPerChainWithAddress`
**Output:** `OnChainOutput`

**Deployment order:**
1. WETH9 (wrapped native token)
2. LINK token (or use existing)
3. RMNRemote + RMNProxy (set ARM on proxy)
4. Router + TestRouter
5. TokenAdminRegistry + RegistryModule
6. NonceManager (authorize OnRamp/OffRamp as callers)
7. FeeQuoter (configure authorized callers, LINK/native token pricing)
8. OffRamp (set source chain configs)
9. OnRamp (set destination chain configs)
10. PingPongDemo (optional, if `DeployPingPongDapp` is true)

### MCMS Sequences

**Source:** [v1_6_0/sequences/mcms.go](../v1_6_0/sequences/mcms.go)

EVM delegates MCMS deployment to the v1.0.0 deployer since MCMS contracts are not version-specific on EVM.

| Method | Behavior |
|--------|----------|
| `DeployMCMS()` | Returns v1.0.0 deployer's `DeployMCMS()` sequence |
| `FinalizeDeployMCMS()` | Returns v1.0.0 deployer's `FinalizeDeployMCMS()` (no-op) |
| `UpdateMCMSConfig()` | Returns v1.0.0 update config sequence |
| `GrantAdminRoleToTimelock()` | Returns admin role granting sequence |

### Lane Configuration

**Source:** [v1_6_0/sequences/update_lanes.go](../v1_6_0/sequences/update_lanes.go)

#### ConfigureLaneLegAsSource

Configures this chain as the source end of a lane.

**Input:** `UpdateLanesInput`

**Steps:**
1. Apply OnRamp destination chain config updates
2. Apply FeeQuoter destination chain config updates
3. Update gas and token prices on FeeQuoter

#### ConfigureLaneLegAsDest

Configures this chain as the destination end of a lane.

**Input:** `UpdateLanesInput`

**Steps:**
1. Apply OffRamp source chain config updates
2. Apply Router ramp updates (register OnRamp/OffRamp for remote chain)

### FeeQuoter Sequences

**Source:** [v1_6_0/sequences/fee_quoter.go](../v1_6_0/sequences/fee_quoter.go)

| Sequence | Description |
|----------|-------------|
| `FeeQuoterApplyDestChainConfigUpdatesSequence` | Updates destination chain fee configs |
| `FeeQuoterUpdatePricesSequence` | Updates gas and token prices |
| `FeeQuoterApplyTokenTransferFeeConfigUpdatesSequence` | Sets per-token transfer fees |
| `FeeQuoterImportConfigSequence` | Imports existing on-chain fee config |

### OnRamp Sequences

**Source:** [v1_6_0/sequences/onramp.go](../v1_6_0/sequences/onramp.go)

| Sequence | Description |
|----------|-------------|
| `OnRampApplyDestChainConfigUpdatesSequence` | Applies destination chain config |
| `OnRampImportConfigSequence` | Imports existing on-chain config |

### OffRamp Sequences

**Source:** [v1_6_0/sequences/offramp.go](../v1_6_0/sequences/offramp.go)

| Sequence | Description |
|----------|-------------|
| `OffRampApplySourceChainConfigUpdatesSequence` | Applies source chain config |
| `OffRampImportConfigSequence` | Imports existing on-chain config |

### Router Sequences

**Source:** [v1_6_0/sequences/router.go](../v1_6_0/sequences/router.go)

| Sequence | Description |
|----------|-------------|
| `RouterApplyRampUpdatesSequence` | Updates OnRamp/OffRamp references on Router |

### RMN Remote Sequences

**Source:** [v1_6_0/sequences/rmn_remote.go](../v1_6_0/sequences/rmn_remote.go)

RMN configuration sequences for blessed/cursed state management.

### OCR3 Sequences

**Source:** [v1_6_0/sequences/ocr.go](../v1_6_0/sequences/ocr.go)

Sets OCR3 configuration on the OffRamp contract.

### Token Sequences

**Source:** [v1_6_0/sequences/token.go](../v1_6_0/sequences/token.go), [v1_6_0/sequences/token_and_pools.go](../v1_6_0/sequences/token_and_pools.go)

| Sequence | Description |
|----------|-------------|
| `DeployToken` | Deploys ERC20 BurnMint/BurnMintWithDrip token |
| `ConfigureTokenForTransfersSequence` | Registers token + configures remote chains |
| `ManualRegistration` | Registers customer token via proposeAdministrator |
| `SetTokenPoolRateLimits` | Sets inbound/outbound rate limits |
| `DeployTokenPoolForToken` | Deploys token pool for existing token |
| `UpdateAuthorities` | Transfers token/pool ownership to timelock |

### Token Pool Sequences

**Source:** [v1_6_0/sequences/deploy_token_pool_contracts.go](../v1_6_0/sequences/deploy_token_pool_contracts.go)

Handles deployment of various token pool types:
- BurnMintTokenPool
- BurnWithFromMintTokenPool
- BurnFromMintTokenPool
- LockReleaseTokenPool
- BurnMintWithExternalMinterTokenPool
- CCTPTokenPool

Supports both v1.5.1 and v1.6.1 pool versions.

### Token Governor Sequences

**Source:** [v1_6_0/sequences/token_governor.go](../v1_6_0/sequences/token_governor.go)

Sequences for deploying and configuring the Token Governor contract.

### PingPong Sequences

| Sequence | Description |
|----------|-------------|
| `ConfigurePingPongSequence` | Configures PingPong demo for a lane |

---

## Sequence Composition Pattern

EVM sequences follow a consistent pattern:

```go
var MySequence = operations.NewSequence(
    "my-sequence",
    semver.MustParse("1.6.0"),
    "Description of what this sequence does",
    func(b operations.Bundle, chains cldf_chain.BlockChains, input MyInput) (sequences.OnChainOutput, error) {
        var output sequences.OnChainOutput

        // Execute operations
        writeResult, err := operations.ExecuteOperation(b, MyWriteOp, chain, contract.FunctionInput[Args]{
            Address:       contractAddr,
            ChainSelector: input.ChainSelector,
            Args:          myArgs,
        })
        if err != nil {
            return output, err
        }

        // Collect MCMS batch operations from writes
        batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{writeResult.Output})
        if err != nil {
            return output, err
        }
        output.BatchOps = append(output.BatchOps, batchOp)

        return output, nil
    },
)
```

Key patterns:
- Operations return `WriteOutput` which may or may not have been executed directly
- `NewBatchOperationFromWrites` collects un-executed writes into MCMS proposals
- Deploy operations return `datastore.AddressRef` which gets added to `output.Addresses`
- Sub-sequences can be composed via `sequences.RunAndMergeSequence`

---

## EVM-Specific Changesets

The EVM implementation also provides version-specific changesets in `v1_5_0/changesets/` and `v1_6_0/changesets/` for MCMS configuration updates and other EVM-specific operations not covered by the shared changeset layer.
