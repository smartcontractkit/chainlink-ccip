# Chainlink Solana Contracts (Programs)

## Prerequisites

- [Rust](https://www.rust-lang.org/tools/install)
- [Solana CLI](https://docs.anza.xyz/cli/install) (provides `solana-test-validator`)
- [Anchor CLI](https://www.anchor-lang.com/docs/installation)
- [Go](https://go.dev/doc/install) (see `go.mod` for required version)
- [anchor-go](https://github.com/gagliardetto/anchor-go) v1.0.0 — `go install github.com/gagliardetto/anchor-go@v1.0.0`

## Build

To build on the host:

```bash
anchor build
# or from the repo root:
make build-contracts
```

To build inside a Docker environment (reproducible builds):

```bash
make docker-build-contracts
```

## Test

### Rust unit tests

```bash
cargo test
# or from the repo root:
make rust-tests
```

### Go integration tests

The Go tests spin up a local `solana-test-validator`, deploy the compiled programs, and run
integration tests against them.

Prerequisites:

- `solana-test-validator` installed and on your PATH
- Contracts built (`anchor build` or `make build-contracts`)
- Vendor programs present in `target/vendor/` (pre-built CCTP `.so` files, committed to git)

```bash
# from the repo root:
make go-tests
```

To run specific test suites:

```bash
go test ./tests/ccip/ -run TestCCIPRouter -v -count=1
go test ./tests/mcms/ -run TestMcmSetConfig -v -count=1
go test ./tests/examples/ -run TestBaseTokenPoolHappyPath -v -count=1
```

Note: subtests within a single top-level `Test*` function share sequential state, so running
an individual subtest in isolation (e.g. `-run TestCCIPRouter/Config`) will likely fail because
earlier setup subtests won't have executed.

## Go bindings generation

Install `anchor-go` and regenerate bindings after contract changes:

```bash
go install github.com/gagliardetto/anchor-go@v1.0.0
make anchor-go-gen
```

This builds the contracts, generates Go bindings from the IDL files in `target/idl/` and
`target/vendor/`, and outputs them to `gobindings/`.
