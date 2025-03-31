# Solana Onchain Utilities

A collection of utilities for Solana blockchain development.

## Contract Verification

https://solana.com/developers/guides/advanced/verified-builds

### Usage

1. cd in `chains/solana/scripts`
2. Enter into the nix shell: `nix-shell`
3. Edit the vars in verify.go based on your liking
4. Run `go run verify.go verify`
5. That will print the command to verify each program.
6. We print the command instead of running them, as the command spins up a docker container on every run.
7. That has been flaky for me. Running each command separately has been reliable (where you delete the docker container manually if the run get stuck)


## IDL Upload

### Usage

1. cd in `chains/solana/scripts`
2. Enter into the nix shell: `nix-shell`
3. Edit the vars in `verify.go` based on your liking
4. Run `go run verify.go idl`
5. That will prepare your idl folder in `scripts/idl` (this is different from the idl on github because it contains metada.address which the init command requires)
6. Overwrite `chains/solana/contracts/target/idl` with `scripts/idl`
7. In `chains/solana/contracts/Anchor.toml`, change anchor_version to whatever you get from `anchor -V` (or you will need avm in your shell which is not required according to me)
8. cd into `chains/solana/contracts` (because the `anchor idl init` command needs to be in an anchor workspace)
9. Run the commands printed in Step 2.