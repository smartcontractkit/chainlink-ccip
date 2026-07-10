---
title: "EVM Operations Reference"
sidebar_label: "Operations"
sidebar_position: 3
---

# EVM Operations Reference

Operations are atomic building blocks that perform a single contract interaction. The EVM implementation provides a typed framework for creating read, write, and deploy operations with built-in MCMS support.

**Source:** [chainlink-deployments-framework `chain/evm/operations2/contract`](https://github.com/smartcontractkit/chainlink-deployments-framework/tree/main/chain/evm/operations2/contract)

**Helpers:** [utils/operations/](../utils/operations/) (`ExecuteRead`, `ExecuteWrite`, `ExecuteDeploy`, `MaybeDeployContract`, idempotency keys)

---

## Operation Framework

### NewRead

Creates a read-only contract call operation.

**Source:** `chainlink-deployments-framework/chain/evm/operations2/contract`

Generated operations export factory functions that take a pre-bound contract:

```go
func NewReadGetDestChainConfig(c gobindings.OnRampInterface) *Operation[contract.FunctionInput[uint64], gobindings.GetDestChainConfig, evm.Chain]
```

Call sites bind the contract and pass the factory result to `ExecuteOperation` (or use `evmops.ExecuteRead`):

```go
report, err := evmops.ExecuteRead(b, chain, onRampAddr, onrampbindings.NewOnRamp, onrampops.NewReadGetDestChainConfig, remoteSelector)
```

### NewWrite

Creates a state-modifying contract call operation with automatic MCMS fallback.

**Source:** `chainlink-deployments-framework/chain/evm/operations2/contract`

Write factories take a pre-bound contract (`Contract` field). `IsAllowedCaller` receives the bound contract instance.

**Behavior:**
1. Validates input via `Validate` (if provided)
2. Checks if the deployer key is an allowed caller via `IsAllowedCaller`
3. **If allowed:** Executes the transaction directly, waits for confirmation, returns `WriteOutput` with `ExecInfo` populated
4. **If not allowed:** Encodes the transaction as an MCMS batch operation, returns `WriteOutput` without `ExecInfo`

**Output:**
```go
type WriteOutput struct {
    ChainSelector uint64
    Tx            mcms_types.Transaction
    ExecInfo      *ExecInfo  // nil if transaction was deferred to MCMS
}

func (o WriteOutput) Executed() bool  // Returns true if executed directly
```

**Common Caller Checks:**
- `OnlyOwner()` -- checks if caller is the contract owner (with retry for testnet flakiness)
- `AllCallersAllowed()` -- always returns true (permissive)

### NewDeploy

Creates a contract deployment operation.

**Source:** `chainlink-deployments-framework/chain/evm/operations2/contract`

`DeployInput` no longer includes `ChainSelector` (chain comes from the `evm.Chain` dependency). Use `evmops.MaybeDeployContract` for idempotent deploys.

### Common Input Type

```go
type FunctionInput[ARGS any] struct {
    Args     ARGS
    GasLimit uint64 `json:"gasLimit,omitempty"`  // optional write/deploy override
    GasPrice uint64 `json:"gasPrice,omitempty"`
}
```

Use `operations.WithIdempotencyKey` (or `evmops.ContractIdempotencyKey` / `ChainIdempotencyKey` via helpers) to scope report reuse per contract address or chain.

### Batch Operation Helper

```go
func NewBatchOperationFromWrites(writes []WriteOutput) (mcms_types.BatchOperation, error)
```

Converts a slice of `WriteOutput` into an MCMS `BatchOperation`. Filters out already-executed transactions.

---

## Operations by Contract

All operations are in [v1_6_0/operations/](../v1_6_0/operations/).

### OnRamp (`operations/onramp/`)

| Operation | Type | Description |
|-----------|------|-------------|
| Deploy | Deploy | Deploys OnRamp contract |
| GetDestChainConfig | Read | Gets destination chain configuration |
| GetDynamicConfig | Read | Gets dynamic configuration |
| GetStaticConfig | Read | Gets static configuration |
| ApplyDestChainConfigUpdates | Write | Updates destination chain configs for remote chains |

### OffRamp (`operations/offramp/`)

| Operation | Type | Description |
|-----------|------|-------------|
| Deploy | Deploy | Deploys OffRamp contract |
| GetSourceChainConfig | Read | Gets source chain configuration |
| GetStaticConfig | Read | Gets static configuration |
| GetDynamicConfig | Read | Gets dynamic configuration |
| ApplySourceChainConfigUpdates | Write | Updates source chain configs for remote chains |

### FeeQuoter (`operations/fee_quoter/`)

| Operation | Type | Description |
|-----------|------|-------------|
| Deploy | Deploy | Deploys FeeQuoter contract |
| ApplyDestChainConfigUpdates | Write | Updates destination chain fee configs |
| UpdatePrices | Write | Updates gas and token prices |
| ApplyTokenTransferFeeConfigUpdates | Write | Sets per-token transfer fee configurations |
| GetDestChainConfig | Read | Gets destination chain fee config |
| GetTokenTransferFeeConfig | Read | Gets token transfer fee configuration |

### NonceManager (`operations/nonce_manager/`)

| Operation | Type | Description |
|-----------|------|-------------|
| Deploy | Deploy | Deploys NonceManager contract |
| ApplyAuthorizedCallerUpdates | Write | Authorizes OnRamp/OffRamp as callers |

### RMN Remote (`operations/rmn_remote/`)

| Operation | Type | Description |
|-----------|------|-------------|
| Deploy | Deploy | Deploys RMNRemote contract |

### CCIP Home (`operations/ccip_home/`)

Operations for the CCIPHome contract on the home chain.

### Token Pool Operations

| Module | Description |
|--------|-------------|
| `burn_mint_with_external_minter_token_pool/` | BurnMint pool with external minter |
| `hybrid_with_external_minter_token_pool/` | Hybrid pool with external minter |

### Token Governor (`operations/token_governor/`)

Operations for the Token Governor contract.

---

## Creating New Operations

### Read Operation Example (generated factory)

```go
func NewReadGetDestChainConfig(c gobindings.OnRampInterface) *cld_ops.Operation[contract.FunctionInput[uint64], gobindings.GetDestChainConfig, cldf_evm.Chain] {
    return contract.NewRead(contract.ReadParams[uint64, gobindings.GetDestChainConfig, gobindings.OnRampInterface]{
        Name:         "onramp:get-dest-chain-config",
        Version:      Version,
        ContractType: ContractType,
        Contract:     c,
        CallContract: func(c gobindings.OnRampInterface, opts *bind.CallOpts, destChainSelector uint64) (gobindings.GetDestChainConfig, error) {
            return c.GetDestChainConfig(opts, destChainSelector)
        },
    })
}
```

### Write Operation Example (generated factory)

```go
func NewWriteApplyDestChainConfigUpdates(c gobindings.OnRampInterface) *cld_ops.Operation[contract.FunctionInput[[]gobindings.OnRampDestChainConfigArgs], contract.WriteOutput, cldf_evm.Chain] {
    return contract.NewWrite(contract.WriteParams[[]gobindings.OnRampDestChainConfigArgs, gobindings.OnRampInterface]{
        Name:         "onramp:apply-dest-chain-config-updates",
        Version:      Version,
        ContractType: ContractType,
        ContractABI:  gobindings.OnRampMetaData.ABI,
        Contract:     c,
        IsAllowedCaller: func(c gobindings.OnRampInterface, opts *bind.CallOpts, caller common.Address, args []gobindings.OnRampDestChainConfigArgs) (bool, error) {
            return contract.OnlyOwner(c, opts, caller, args)
        },
        CallContract: func(c gobindings.OnRampInterface, opts *bind.TransactOpts, args []gobindings.OnRampDestChainConfigArgs) (*types.Transaction, error) {
            return c.ApplyDestChainConfigUpdates(opts, args)
        },
    })
}
```

### Deploy Operation Example

```go
var Deploy = contract.NewDeploy(contract.DeployParams[DeployInput]{
    Name:             "onramp:deploy",
    Version:          semver.MustParse("1.6.0"),
    Description:      "Deploys the OnRamp contract",
    ContractMetadata: onramp.OnRampMetaData,
})
```

---

## Code Generation

Many EVM operations are auto-generated from gobindings via `operations-gen@v0.1.0`. Configuration is in `operations_gen_config.yaml`. Regenerate with:

```bash
cd chains/evm && make operations-fast
```
