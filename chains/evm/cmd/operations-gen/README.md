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
    package_name: fee_quoter    # Optional: override default package name
    abi_file: "fee_quoter.json" # Optional: override default ABI filename
    no_deployment: false         # Optional: skip bytecode and Deploy operation (default: false)
    functions:
      - name: updatePrices
        access: owner  # Generates MCMS-compatible transaction
      - name: getPrice 
```

### Optional Fields

- `package_name`: Override the generated package name (default: snake_case of contract_name)
- `abi_file`: Override the ABI filename to use (default: {package_name}.json)
- `no_deployment`: Skip bytecode constant and Deploy operation (default: false, useful for contracts deployed elsewhere)

### Access Control

- `owner`: Generates write operation with MCMS support
- `public`: Generates read-only or public write operation

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
