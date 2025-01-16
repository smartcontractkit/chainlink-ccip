# Chainlink Solana contracts (programs)

## Prerequisites

Install Rust, Solana & Anchor. See https://solana.com/docs/intro/installation

## Build

To build on the host:

```
anchor build
```

To build inside a docker environment:

```bash
anchor build --verifiable
```

To build for a specific network, specify via a cargo feature:

```bash
anchor build -- --features mainnet
```

Available networks with declared IDs:

- mainnet
- testnet
- devnet
- localnet (default)

## Test

Make sure to run `pnpm i` to fetch mocha and other test dependencies.

Start a dockerized shell that contains Solana and Anchor:

```bash
./scripts/anchor-shell.sh
```

Next, generate a keypair for anchor:

```bash
solana-keygen new -o id.json
```

### Run anchor TypeScript tests (automatically tests against a local node)

```bash
anchor test
```

### Run GoLang tests (automatically tests against a local node)

Pre-requisites:

- Have the `solana-test-validator` command installed
- Run `anchor build` if there have been any changes to the program under test.

```bash
make contracts-go-tests
```

#### `anchor-go` bindings generation

Install `https://github.com/gagliardetto/anchor-go`

Current version: [v0.2.3](https://github.com/gagliardetto/anchor-go/tree/v0.2.3)

To install `anchor-go` locally so that you can use the `anchor-go` command globally, follow these steps:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/gagliardetto/anchor-go.git
   cd anchor-go
   git checkout v0.2.3
   ```

2. **Install the command globally:**

   Run the following command to install the `anchor-go` command globally:

   ```bash
   go install
   ```

   This will install the `anchor-go` binary to your `$GOPATH/bin` directory. Make sure that this directory is included in your system's `PATH` environment variable.

3. **Then run the following command to generate the Go bindings:**

   ```bash
   make anchor-go-gen
   ```
