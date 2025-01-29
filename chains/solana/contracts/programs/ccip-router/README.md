# CCIP Router implementation in SVM

Collapsed Router + OnRamp + OffRamp Contracts Implementation for CCIP in SVM.

## Testing

- The Rust Tests execute logic specific to the program, to run them use the command

  ```bash
  cargo test
  ```

- The Golang Tests are in the `../contracts/tests` folder, using generated bindings from `gagliardetto/anchor-go` library. The tests are run in a local network, so the tests are fast and don't require any real SVM network. To run them, use the command

  ```bash
  anchor build # build contract artifacts
  go test ./... -count=1 -failfast
  ```

  Note: the [solana-test-validator](https://docs.anza.xyz/cli/examples/test-validator) must be installed
