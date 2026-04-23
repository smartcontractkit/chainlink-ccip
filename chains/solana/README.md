# CCIP Solana Onchain

## Project Structure

- `contracts/` — Solana programs (Rust/Anchor) + integration tests (Go)
  - `programs/` — Anchor program source code (ccip-router, fee-quoter, token pools, mcm, timelock, etc.)
  - `tests/` — Go integration tests organized by area (`ccip/`, `mcms/`, `examples/`)
  - `target/deploy/` — compiled `.so` program binaries and IDL JSON files
  - `target/vendor/` — pre-built third-party program binaries (CCTP), committed to git
- `gobindings/` — auto-generated Go bindings for contracts using `anchor-go`
- `scripts/` — build and generation scripts (e.g. `anchor-go-gen.sh`)
- `utils/` — shared Go utility libraries

## Dependencies

- [Rust](https://www.rust-lang.org/tools/install) (version pinned via `contracts/rust-toolchain.toml`)
- [Go](https://go.dev/doc/install) (see `go.mod` for required version)
- [Solana CLI](https://docs.anza.xyz/cli/install) (provides `solana-test-validator`)
- [Anchor CLI](https://www.anchor-lang.com/docs/installation)
- [anchor-go](https://github.com/gagliardetto/anchor-go) v0.2.3 — `GOTOOLCHAIN=go1.20 go install github.com/gagliardetto/anchor-go@v0.2.3`
- [golangci-lint](https://golangci-lint.run/welcome/install/) v2.7.0 — `go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.7.0`
- [gotestloghelper](https://github.com/smartcontractkit/chainlink-testing-framework) — for `make go-tests` output formatting

### macOS (Apple Silicon) Note

The `solana-test-validator` requires GNU tar. Without it, you'll see errors like
`Archive error: extra entry found: "._genesis.bin"`. Install it and put it on your PATH:

```bash
brew install gnu-tar
export PATH="/opt/homebrew/opt/gnu-tar/libexec/gnubin:$PATH"
```

See [solana-labs/solana#35629](https://github.com/solana-labs/solana/issues/35629) for details.

## Environment Setup

Run the setup script to install and configure the required toolchain (Anchor `0.29.0`, Solana CLI `1.17.25`, and `anchor-go`):

```bash
./scripts/setup-contract-env.sh
```

The script will:
1. Verify that `rustc`, `cargo`, and `go` are installed
2. Install [AVM](https://github.com/coral-xyz/anchor) (Anchor Version Manager) and switch to the required Anchor version
3. Install the required Solana CLI version
4. Build and install `anchor-go` from source (with patched dependencies for Go 1.22+ compatibility)

> **Note:** You must have Rust, Cargo, and Go installed beforehand (see [Dependencies](#dependencies)).

## Development

### Build contracts

```bash
# Build on the host (requires Anchor CLI + Solana CLI + Rust toolchain)
make build-contracts

# Or build inside Docker for reproducibility
make docker-build-contracts
```

### Generate Go bindings

After any contract changes, regenerate the Go bindings:

```bash
make anchor-go-gen
```

This builds the contracts, then runs `anchor-go` against the IDL files to produce bindings in `gobindings/`.

### Run tests

```bash
# Go integration tests (spins up local solana-test-validator per test)
make go-tests

# Rust unit tests
make rust-tests

# Run a specific test suite
go test ./contracts/tests/ccip/ -run TestCCIPRouter -v -count=1
go test ./contracts/tests/mcms/ -run TestMcmSetConfig -v -count=1
go test ./contracts/tests/examples/ -run TestBaseTokenPoolHappyPath -v -count=1
```

Note: Go integration tests require that contracts have been built (`make build-contracts`)
and that the vendor `.so` files exist in `contracts/target/vendor/`.

### Format and lint

```bash
make format      # format Go + Rust
make lint-go     # run golangci-lint
```

### Full CI check

```bash
make solana-checks
```

This runs, in order: `clippy`, `anchor-go-gen`, `format`, `gomodtidy`, `lint-go`, `rust-tests`, `go-tests`, `build-contracts`.
