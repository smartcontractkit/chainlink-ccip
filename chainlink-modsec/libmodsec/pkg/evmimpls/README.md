# EVM Implementations (`evmimpls`)

This module provides the EVM-specific implementations for the interfaces defined in `modsectypes`. It includes components for reading from source and destination chains, encoding and decoding messages, and transmitting transactions.

## Dependencies

The build and test processes for this module rely on the following tools:

- **Solidity Compiler (`solc`)**: Used to compile the test contracts. This is run via a Docker container, so you do not need to install `solc` locally. The version is specified in the `Makefile`.
- **Go Bindings Generator (`abigen`)**: Used to generate Go bindings for the compiled Solidity contracts. The `Makefile` will automatically install the correct version of `abigen` to your `$GOPATH/bin` directory.

## Building and Code Generation

A `Makefile` is provided to automate the compilation and code generation process.

### Compiling the Contracts

To compile the Solidity test contracts located in the `testcontracts/` directory, run:

```bash
make compile-contract
```

This command will:
1. Create `testcontracts/abis` and `testcontracts/bin` directories.
2. Run `solc` in a Docker container to generate the ABI (`.json`) and binary (`.bin`) files for the test contracts.

### Generating Go Bindings

To generate the Go bindings for the compiled contracts, run:

```bash
make generate-go
```

This command will:
1. Ensure `abigen` is installed.
2. Run `abigen` to generate the Go wrapper files in the `gethwrappers/` directory.

### All-in-One Command

To compile the contracts and generate the Go bindings in a single step, you can use the `all` target:

```bash
make all
```

This is the recommended command to use after making changes to the Solidity test contracts.

## Testing

To run the Go unit tests for this module, use the following command:

```bash
make test
```

This will execute all `_test.go` files in the module and display the results, including test coverage.

## Cleaning Up

To remove all generated artifacts, including the compiled contracts and Go bindings, run:

```bash
make clean
```

This will delete the `testcontracts/abis`, `testcontracts/bin`, and `gethwrappers` directories.
