# operations-gen

Generates Go operation wrappers for EVM smart contracts from their ABIs.

## Usage

From `chains/evm/`:

```bash
make operations
```

Or run directly:

```bash
go run ./cmd/operations-gen/main.go -config deployment/operations_gen_config.yaml
```

## Configuration

Edit `chains/evm/deployment/operations_gen_config.yaml`:

```yaml
version: "1.0"

input:
  base_path: ".."  # Directory containing abi/ and bytecode/ folders

output:
  base_path: "."   # Directory for generated operations/

contracts:
  - contract_name: FeeQuoter
    version: "1.6.0"
    functions:
      - name: updatePrices
        access: owner  # Generates MCMS-compatible transaction
      - name: getPrice 
```

### Access Control

- `owner`: Generates write operation with MCMS support
- Omit `access`: For view functions or functions without special access control

## Output

Generates files in `deployment/{version}/operations/{contract}/`:

- Type-safe operation structs
- Contract deployment helpers
- Read and write operation functions
- MCMS transaction builders (for owner functions)

## Requirements

- ABIs in `chains/evm/abi/{version}/`
- Bytecode in `chains/evm/bytecode/{version}/`

Paths in config are relative to the config file location.
