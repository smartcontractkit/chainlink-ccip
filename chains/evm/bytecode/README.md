# Bytecode Directory

This directory contains extracted bytecode from versioned gobindings, organized by version.

## Overview

Bytecode is automatically extracted from the gobindings using the extraction script located at:
`gobindings/cmd/extract_bytecode/main.go`

The script:
- Scans all versioned directories in `gobindings/generated/` (e.g., `v1_0_0`, `v1_2_0`, etc.)
- Excludes the `latest/` directory (unaudited code)
- Extracts the `Bin` field from each `*MetaData` variable in `.go` files
- Writes the bytecode to corresponding `.bin` files in this directory
- Also extracts ABIs (as JSON) to the `../abi` directory

## Usage

### Running Manually

From the `chains/evm` directory:

```
make extract-bytecode
```

### Running in CI

The script can be run in CI pipelines. It will:
- Create the bytecode directory structure if it doesn't exist
- Extract bytecode from all versioned gobindings
- Report any errors to stderr with a non-zero exit code

Example CI integration:

```yaml
- name: Extract bytecode and ABIs
  run: |
    cd chains/evm
    make extract-bytecode
```

Note: The extraction script extracts both bytecode and ABIs in a single pass.

## Directory Structure

The bytecode is organized by version with all binaries directly in the version directory:

```
bytecode/
├── v1_0_0/
│   ├── lock_release_token_pool.bin
│   ├── rmn_proxy_contract.bin
│   └── ...
├── v1_2_0/
│   ├── burn_mint_token_pool.bin
│   ├── commit_store.bin
│   └── ...
└── ...
```

Each `.bin` file contains the raw bytecode hex string (with `0x` prefix) extracted from the corresponding gobinding.

## How It Works

The extraction script uses Go's AST parser to:
1. Parse each `.go` file in versioned directories
2. Find variable declarations ending with `MetaData`
3. Extract the `Bin` field from the `&bind.MetaData{...}` composite literal
4. Write the bytecode to a corresponding `.bin` file

This approach is robust and doesn't rely on regex or string manipulation, ensuring accurate extraction even as the gobinding format evolves.