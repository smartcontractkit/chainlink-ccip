# CCIP Solana Onchain

- `contracts`: solana programs (rust) + tests (go) built on the anchor framework
- `gobindings`: auto-generated go bindings for contracts using `anchor-go`
- `scripts`: various scripts for generating artifacts

## Dependencies

- rust: https://www.rust-lang.org/tools/install
- go: https://go.dev/doc/install
- solana: https://docs.anza.xyz/cli/install
- anchor: https://www.anchor-lang.com/docs/installation

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

```bash
# install anchor-go if needed
go install github.com/gagliardetto/anchor-go@v0.2.3

# build contracts + IDL
anchor build

# go bindings need to be regenerated if contract changes were made
./scripts/anchor-go-gen.sh

# test contracts
go test ./... -v -count=1 -failfast
```
