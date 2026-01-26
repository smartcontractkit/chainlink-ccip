# Automatic Operations Generation from Solidity Contracts

## Overview

This system automatically generates CCIP operations files from Solidity contracts with **minimal configuration**.  
You only need to specify:
1. The Solidity contract path
2. Which functions to generate operations for

Everything else is **automatically extracted** from the contract and gobindings.

## Quick Example

### Input Configuration (`operations_gen_config_simple.yaml`)

```yaml
version: "1.0"

output:
  base_path: chains/evm/deployment
  version_prefix: v1_6_0

contracts:
  - solidity_path: chains/evm/contracts/rmn/RMNRemote.sol
    functions:
      - curse
      - uncurse
      - isCursed
```

### Generate Operations

```bash
go run chains/evm/cmd/operations-gen-simple/main.go
```

### Result

✓ Generates `chains/evm/deployment/v1_6_0/operations/rmn_remote/rmn_remote.go`  
✓ **Byte-for-byte identical** to handwritten version  
✓ Complete with Deploy, Write, and Read operations

## What Gets Auto-Generated

From the Solidity contract and gobindings, the generator automatically extracts:

### ✅ Contract Metadata
- Contract name (from filename)
- Version (from `typeAndVersion` constant)
- Package naming (snake_case conversion)

### ✅ Constructor Information
- Parameter names and types
- Go type mappings

### ✅ Function Details
- Function signatures
- Parameter names and Solidity types
- Return types
- State mutability (view/pure vs. non-payable)
- Overload detection (automatically uses `Curse0`, `Uncurse0`, etc.)

### ✅ Access Control
- Detects `onlyOwner` modifiers from Solidity source
- Auto-generates `contract.OnlyOwner` vs. `contract.AllCallersAllowed`

### ✅ Type Mappings
- Solidity → Go type conversion
- Custom types (e.g., `bytes16` → `fastcurse.Subject`)
- Array types
- Struct types

### ✅ Import Management
- Required package imports
- Conditional imports (e.g., `fastcurse` only when needed)

## Configuration Reference

### Full Configuration Format

```yaml
version: "1.0"

# Output configuration (applies to all contracts)
output:
  base_path: chains/evm/deployment       # Base deployment path
  version_prefix: v1_6_0                 # Version folder name

contracts:
  - solidity_path: chains/evm/contracts/rmn/RMNRemote.sol  # REQUIRED
    
    # Optional overrides (usually auto-detected):
    contract_name: RMNRemote   # Default: extracted from filename
    version: "1.6.0"           # Default: extracted from typeAndVersion()
    
    # Functions to generate (REQUIRED)
    functions:
      - curse      # Will generate Write operation
      - uncurse    # Will generate Write operation  
      - isCursed   # Will generate Read operation
```

### Minimal Configuration

The absolute minimum you need:

```yaml
contracts:
  - solidity_path: chains/evm/contracts/rmn/RMNRemote.sol
    functions: [curse, uncurse, isCursed]
```

## How It Works

### 1. Extract Contract Name
```
chains/evm/contracts/rmn/RMNRemote.sol → RMNRemote
```

### 2. Find Gobinding
```
RMNRemote → rmn_remote (snake_case)
→ chains/evm/gobindings/generated/v1_6_0/rmn_remote/rmn_remote.go
```

### 3. Extract ABI
Parses the ABI JSON embedded in the gobinding's `MetaData` struct:
```go
var RMNRemoteMetaData = &bind.MetaData{
    ABI: "[{\"type\":\"function\",\"name\":\"curse\",...}]",
    ...
}
```

### 4. Parse Functions
From ABI entries:
```json
{
  "type": "function",
  "name": "curse",
  "inputs": [{"name": "subjects", "type": "bytes16[]"}],
  "stateMutability": "nonpayable"
}
```

Generates:
- Operation type: Write (non-payable, not view/pure)
- Args struct: `CurseArgs`
- Call method: `Curse0` (detects array overload)
- Parameters: `Subject []fastcurse.Subject`

### 5. Detect Access Control
Scans Solidity source for:
```solidity
function curse(bytes16[] memory subjects) public onlyOwner {
                                                   ^^^^^^^^
```
→ Generates `contract.OnlyOwner`

### 6. Generate Code
Uses templates to create the complete operations file.

## Type Mapping

| Solidity Type | Go Type | Notes |
|---------------|---------|-------|
| `uint8` - `uint64` | `uint8` - `uint64` | Direct mapping |
| `uint256`, `int256` | `*big.Int` | Large integers |
| `address` | `common.Address` | Ethereum addresses |
| `bool` | `bool` | Direct mapping |
| `string` | `string` | Direct mapping |
| `bytes` | `[]byte` | Dynamic bytes |
| `bytes16`, `bytes32` | `[16]byte`, `[32]byte` | Fixed bytes |
| `bytes16` (Subject) | `fastcurse.Subject` | **Custom type detection** |
| `type[]` | `[]type` | Arrays |
| `contract IFoo` | `common.Address` | Contract references |

## Operation Type Detection

### Write Operations
Functions with `stateMutability`:
- `nonpayable` → Write operation
- `payable` → Write operation

### Read Operations  
Functions with `stateMutability`:
- `view` → Read operation
- `pure` → Read operation

### Constructor
Special `type: "constructor"` entry → Deploy operation

## Function Overload Handling

When multiple functions exist with the same name (e.g., `curse(bytes16)` and `curse(bytes16[])`):

1. **Detects overloads** by checking for array parameters
2. **Prefers array version** for operations (more flexible)
3. **Uses correct call method**: `Curse0` instead of `Curse`

Example from gobinding:
```go
func (c *RMNRemote) Curse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error)
func (c *RMNRemote) Curse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error)
                    ^^^^^                                                                 
                    Overload suffix
```

## Adding New Contracts

### Step 1: Ensure Gobinding Exists

Check that the contract has a gobinding:
```bash
ls chains/evm/gobindings/generated/v1_6_0/your_contract/
```

If not, add to `go_generate.go`:
```go
//go:generate go run ./wrap ccip YourContract your_contract
```

And run:
```bash
go generate ./...
```

### Step 2: Add to Config

```yaml
contracts:
  - solidity_path: chains/evm/contracts/path/YourContract.sol
    functions:
      - functionOne
      - functionTwo
      - functionThree
```

### Step 3: Generate

```bash
go run chains/evm/cmd/operations-gen-simple/main.go
```

Done! ✓

## Validation

The generator has been validated to produce **byte-for-byte identical output** to handwritten operations files:

```bash
$ go run chains/evm/cmd/operations-gen-simple/main.go
✓ Generated operations for RMNRemote at chains/evm/deployment/v1_6_0/operations/rmn_remote/rmn_remote.go

$ diff rmn_remote_original.go rmn_remote.go
# No output = Perfect match!
```

## Example: Generated vs. Handwritten

For the RMNRemote contract with functions `[curse, uncurse, isCursed]`:

### Input (3 lines)
```yaml
- solidity_path: chains/evm/contracts/rmn/RMNRemote.sol
  functions: [curse, uncurse, isCursed]
```

### Output (78 lines)
Complete operations file with:
- Package and imports (12 lines)
- Type definitions (4 lines)
- Constructor args struct (4 lines)
- Deploy operation (11 lines)
- CurseArgs struct (3 lines)
- Curse write operation (11 lines)
- Uncurse write operation (11 lines)
- IsCursed read operation (9 lines)

**Reduction: 26x fewer lines of configuration needed!**

## Files

| File | Purpose | Lines |
|------|---------|-------|
| `operations_gen_config_simple.yaml` | Minimal config | ~10 |
| `cmd/operations-gen-simple/main.go` | Generator | ~600 |
| `README_OPERATIONS_GEN.md` | Documentation | This file |

## Comparison: Old vs. New System

### Old System (operations-gen)
❌ Required full YAML config with all details  
❌ Manual type mappings  
❌ Manual import management  
❌ Manual call method names  
❌ ~70 lines of config per contract  

### New System (operations-gen-simple)
✅ Only contract path + function names  
✅ Auto type detection  
✅ Auto import management  
✅ Auto overload detection  
✅ ~3 lines of config per contract  

## Future Enhancements

Potential improvements:
- **Function filtering** - Regex patterns to select functions
- **Custom validators** - Generate validation logic from Solidity requires
- **Documentation extraction** - Pull NatSpec comments into operations
- **Multi-version** - Generate for multiple contract versions at once
- **Watch mode** - Auto-regenerate on Solidity changes

## Troubleshooting

### Error: "ABI not found in gobinding"

**Cause**: Gobinding file doesn't exist or is malformed.

**Solution**:
```bash
# Regenerate gobindings
cd chains/evm
go generate ./...
```

### Error: "function X not found in ABI"

**Cause**: Function name typo or function doesn't exist.

**Solution**: Check function name in Solidity contract.

### Generated file doesn't match expected

**Cause**: Gobinding out of sync with Solidity.

**Solution**:
```bash
# Rebuild contracts and regenerate gobindings
FOUNDRY_PROFILE=ccip forge build
go generate ./...
```

## Dependencies

- `gopkg.in/yaml.v3` - YAML parsing
- Existing gobindings - Source of ABI and type info
- Solidity source - For access control detection

## License

Same as parent project.
