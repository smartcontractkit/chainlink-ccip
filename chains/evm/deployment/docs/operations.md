---
title: "EVM Operations Reference"
sidebar_label: "Operations"
sidebar_position: 3
---

# EVM Operations Reference

Operations are atomic building blocks that perform a single contract interaction. The EVM implementation provides a typed framework for creating read, write, and deploy operations with built-in MCMS support.

**Source:** [utils/operations/contract/](../utils/operations/contract/)

---

## Operation Framework

### NewRead

Creates a read-only contract call operation.

**Source:** [utils/operations/contract/read.go](../utils/operations/contract/read.go)

```go
func NewRead[ARGS any, RET any, C any](params ReadParams[ARGS, RET, C]) *Operation[FunctionInput[ARGS], RET, evm.Chain]
```

**Parameters:**
```go
type ReadParams[ARGS any, RET any, C any] struct {
    Name         string                    // Operation ID
    Version      *semver.Version           // Contract version
    Description  string
    ContractType deployment.ContractType
    NewContract  func(address common.Address, backend bind.ContractBackend) (C, error)
    CallContract func(contract C, opts *bind.CallOpts, input ARGS) (RET, error)
}
```

**Behavior:** Instantiates the contract using `NewContract`, then calls `CallContract` with the provided args. Returns the result directly.

### NewWrite

Creates a state-modifying contract call operation with automatic MCMS fallback.

**Source:** [utils/operations/contract/write.go](../utils/operations/contract/write.go)

```go
func NewWrite[ARGS any, C any](params WriteParams[ARGS, C]) *Operation[FunctionInput[ARGS], WriteOutput, evm.Chain]
```

**Parameters:**
```go
type WriteParams[ARGS any, C any] struct {
    Name            string
    Version         *semver.Version
    Description     string
    ContractType    deployment.ContractType
    ContractABI     string
    NewContract     func(address common.Address, backend bind.ContractBackend) (C, error)
    IsAllowedCaller func(contract C, opts *bind.CallOpts, caller common.Address, input ARGS) (bool, error)
    Validate        func(input ARGS) error
    CallContract    func(contract C, opts *bind.TransactOpts, input ARGS) (*eth_types.Transaction, error)
}
```

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

**Source:** [utils/operations/contract/deploy.go](../utils/operations/contract/deploy.go)

```go
func NewDeploy[ARGS any](params DeployParams[ARGS]) *Operation[DeployInput[ARGS], datastore.AddressRef, evm.Chain]
```

**Parameters:**
```go
type DeployParams[ARGS any] struct {
    Name                     string
    Version                  *semver.Version
    Description              string
    ContractMetadata         *bind.MetaData
    BytecodeByTypeAndVersion map[string]Bytecode
    Validate                 func(input ARGS) error
}
```

**Behavior:** Deploys the contract using the chain's deployer key. Supports both standard EVM and ZkSync VM bytecodes. Returns a `datastore.AddressRef` with the deployed address, type, and version.

### Common Input Type

```go
type FunctionInput[ARGS any] struct {
    Address       common.Address
    ChainSelector uint64
    Args          ARGS
}
```

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

### Read Operation Example

```go
var GetDestChainConfig = contract.NewRead(contract.ReadParams[uint64, DestChainConfig, *onramp.OnRamp]{
    Name:         "onramp:getDestChainConfig",
    Version:      semver.MustParse("1.6.0"),
    Description:  "Gets the destination chain config for a remote chain",
    ContractType: "OnRamp",
    NewContract:  onramp.NewOnRamp,
    CallContract: func(c *onramp.OnRamp, opts *bind.CallOpts, destChainSelector uint64) (DestChainConfig, error) {
        return c.GetDestChainConfig(opts, destChainSelector)
    },
})
```

### Write Operation Example

```go
var ApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]DestChainConfigArgs, *onramp.OnRamp]{
    Name:            "onramp:applyDestChainConfigUpdates",
    Version:         semver.MustParse("1.6.0"),
    Description:     "Updates destination chain configurations",
    ContractType:    "OnRamp",
    ContractABI:     onramp.OnRampABI,
    NewContract:     onramp.NewOnRamp,
    IsAllowedCaller: contract.OnlyOwner[[]DestChainConfigArgs](),
    CallContract: func(c *onramp.OnRamp, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
        return c.ApplyDestChainConfigUpdates(opts, args)
    },
})
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

Many EVM operations are auto-generated from Go bindings (gethwrappers). The generation configuration is in `operations_gen_config.yaml`. Generated operations follow the same `NewRead`/`NewWrite`/`NewDeploy` patterns but are produced automatically from the contract ABI.
