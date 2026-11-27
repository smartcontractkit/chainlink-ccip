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

# rebuild the programs for sBPFv3 (see below)
make build-contracts-sbfv3

# test contracts
go test ./... -v -count=1 -failfast
```

### sBPFv3 builds

Once [SIMD-0500](https://solana.com/upgrades/sbpfv3-programs) is active, the network only accepts deployments and upgrades of programs built for sBPFv3. Anchor `0.29.0` cannot target it, so the binaries are built in a second step:

- `anchor build` still generates the IDLs, and sBPFv0 binaries that are not deployed.
- `make build-contracts-sbfv3` rebuilds every program in `contracts/target/deploy` for sBPFv3. It needs `cargo-build-sbf` 4.2 or later (`cargo install cargo-build-sbf --locked`).
- `make docker-build-contracts` does both steps in docker, with no local toolchain. The deployment changesets use it to build the programs they deploy.
- Releases are built by `contracts/scripts/verify_build.sh`, with `solana-verify` and the same sBPFv3 target.

The tests need a `solana-test-validator` from Agave 4.2 or later. Older validators cannot load sBPFv3 programs and fail with `Program is not deployed`. The tests start it with `--alpenglow`, so they run on Alpenglow consensus.

The setup script leaves Solana CLI `1.17.25` active, because `anchor build` needs it. Keep it active and put a newer release first in `PATH` only when running the tests:

```bash
# installs 4.3.2 and makes it the active release
sh -c "$(curl -sSfL https://release.anza.xyz/v4.3.2/install)"
# makes 1.17.25 the active release again; 4.3.2 stays on disk
./scripts/setup-contract-env.sh

PATH="$HOME/.local/share/solana/install/releases/4.3.2/solana-release/bin:$PATH" go test ./... -v -count=1 -failfast
```

#### Compiler flag in `contracts/.cargo/config.toml`

The sBPFv3 build depends on one compiler flag, `-Zallow-features=...`, that Cargo picks up from `contracts/.cargo/config.toml` without any script mentioning it.

- **Why it exists:** `solana-frozen-abi` 1.17 pins `ahash` to `=0.8.5`. That version turns on `stdsimd`, a compiler feature that the toolchain building sBPFv3 no longer has, and the build fails with `unknown feature stdsimd`. The flag lists the features crates may turn on, and leaves `stdsimd` out.
- **Scope:** it only applies to sBPFv3 targets. `anchor build`, host builds and `cargo test` do not see it.
- **Risk, new dependencies:** the list is closed. If a dependency change brings in a crate that turns on another compiler feature, the sBPFv3 build fails with `the feature X is not in the list of allowed features`. Add the feature to the list in that file.
- **Risk, removing the file:** without it the sBPFv3 build fails, both locally and in the release workflow, while `anchor build` keeps passing.
