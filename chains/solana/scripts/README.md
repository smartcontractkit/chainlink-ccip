# Solana Onchain Utilities

A collection of utilities for Solana blockchain development.

## Developer Notes

Can we automate both the processes completely? Maybe.

Will the automation be reliable? Maybe.
- There are some path assumptions we are making
- Anchor version might change
- Verify flow might change
- The verify flow spins up a docker container on every run which is flaky

Is it worth unreliably automating a process that will happen twice a year? Dont think so.

Hence, we are deciding to print the commands for the user to run them manually.
If we see the cadence of these proceseses changing, we can invest some more time in automating them.

## Contract Verification

https://solana.com/developers/guides/advanced/verified-builds

### Usage

1. cd into `chains/solana/scripts`
2. Enter into the nix shell: `nix-shell`
3. Edit the vars in verify.go based on your liking
4. Ensure your wallet is funded with SOL in the env that you have set
5. Run `go run verify.go verify`
6. That will print the command to verify each program.
7. We print the command instead of running them, as the command spins up a docker container on every run.
8. That has been flaky for me. Running each command separately has been reliable (where you delete the docker container manually if the run get stuck)


## IDL Upload

https://www.anchor-lang.com/docs/references/cli#idl-init

### Usage

1. cd into `chains/solana/scripts`
2. Enter into the nix shell: `nix-shell`
3. Edit the vars in `verify.go` based on your liking
4. Ensure your wallet is funded with SOL in the env that you have set
5. Run `go run verify.go idl`
6. That will prepare your idl folder in `scripts/idl` (this is different from the idl on github because it contains metada.address which the init command requires)
7. Overwrite `chains/solana/contracts/target/idl` with `scripts/idl`
8. In `chains/solana/contracts/Anchor.toml`, change anchor_version to whatever you get from `anchor -V` (or you will need avm in your shell which is not required according to me)
9. cd into `chains/solana/contracts` (because the `anchor idl init` command needs to be in an anchor workspace)
10. Run the commands printed in Step 2.