# CCIP Router implementation in SVM

Collapsed Router + OnRamp + OffRamp Contracts Implementation for CCIP in SVM.

## Messages

### Initialization

`initialize`

1. Initialize the Config PDA with the SVM Chain Selector, the Default Extra Args and the Supported Destination Chain Selectors and the Sequence Number with 0.

#### Initialization Accounts

1. `PDA("config", program_id)`: **Init**, it is used to store the default configuration for Extra Args, the SVM Chain Selector and the supported list of Destination Chain Selectors.
1. `signer`: The sender of the message.

### Update Config

`add_chain_selector`, `remove_chain_selector`, `update_svm_chain_selector`

1. Update the Config PDA with the SVM Chain Selector.
1. Init/Close the Chain State PDA with the Supported Destination Chain Selectors (add/remove).

#### Update Config Accounts

1. `PDA("config", program_id)`: **Mut**, it is used to store the default configuration for Extra Args, the SVM Chain Selector and the supported list of Destination Chain Selectors.
1. `PDA("chain_state", chain_selector, program_id)`: **Init/Close**, it stores the latest sequence number per destination chain [only for `add_chain_selector` & `remove_chain_selector`].
1. `signer`: The sender of the message.

### Send Message

`ccipSend`

1. Check if the Destination Chain Selector is supported (enforced by the `chain_state` account).
1. Emit `CCIPMessageSent` event.
1. Update the Sequence Number.

#### Send Message Accounts

1. `PDA("config", program_id)`: **Read only**, it is used to read the default configuration for Extra Args, the SVM Chain Selector and the supported list of Destination Chain Selectors.
1. `PDA("chain_state", chain_selector, program_id)`: **Mut**, increases the latest sequence number per destination chain.
1. `signer`: The sender of the message.

### Commit Report

`commit`

1. Check if the Source Chain Selector is supported.
1. Check if the Interval is valid.
1. Check if the Merkle Root is not empty and is new.
1. Update the Source Chain Reports with the new report.
1. Emit `CommitReportAccepted` event.
1. Emit `Transmitted` event.

#### Commit Report Accounts

1. `PDA("chain_state", chain_selector, program_id)`: **Read only**, checks if the Source Chain Selector is supported.
1. `PDA("commit_report", chain_selector, merkle_root, program_id)`: **Init**, stores the Merkle Root for the new Report.
1. `signer`: The sender of the message.

### Execute Report

`executeReport`

1. Validates that the report was commited
1. Validates that the report was not executed (if not emit `SkippedAlreadyExecutedMessage` or `AlreadyAttempted` events)
1. Validates that the Merkle Proof is correct
1. Executes the message
1. Emit `ExecutionStateChanged` event

#### Execute Report Accounts

1. `PDA("config", program_id)`: **Read only**, checks if the Chain Selectors are supported.
1. `PDA("commit_report", chain_selector, merkle_root, program_id)`: **Mut**, verifies the Merkle Root for the Report and updates state.
1. `signer`: The sender of the message.

## Future Work

- _EMIT_: When Emitting `ExecutionStateChanged` events in the execute report, there are two values that are not being correctly populated: `message_id` & `new_state`.
- _EXTRA ARGS_:
  - What should we do when they are empty? Not serialize them in the hash? Now it's not allowed to be empty in the client.
  - Fix override for extra args: now it only works for `gasLimit`, it should work for `allowOutOfOrderExecution` too.
  - Decide if it makes sense to have a param named `gasLimit`
- _TYPES_
  - [blocked] Use `pub type Selector = u64;` instead of just `u64` in `ccip_send` args. It seems like there are some issues when parsing that in the Typescript tests, not only as a param in messages but also inside the Message struct.
    - Anchor v0.29.0 bug that requires parameter naming to match type - https://github.com/coral-xyz/anchor/issues/2820
    - Anchor-Go does not support code generation for aliased types
    - Attempted changes ([#92](https://github.com/smartcontractkit/chainlink-internal-integrations/pull/92/commits/2c700c430a78f3e63831d7cd0565bcc7206a1eeb)) for future reference
  - Use `[u128; 2]` to store the `Interval` in the `commit` message, this way we will be able to handle a maximum of 128 messages per report to be compatible with EVM.
- _ENABLE/DISABLE CHAINS_: Currently you can add/remove chain selectors, but maybe we need a way to enable/disable them (and only as destination or source).
- _UNIFY ERRORS_: Review the type of errors, and decide if we should have more granularity of them.
- _EXTERNAL EXECUTION_: Understand and documents the limitations over the external execution of the messages.

## Future Work (Production Readiness)

- _FEES_: Add fees for the OnRamp flow [WIP]
- _TOKEN POOLS_: Add the flow related to sending tokens and burning/minting in the Token Pools
- _RMN_: Add the flow related to RMN and cursed/blessed lanes or messages
- _RATE LIMIT_: Add rate limit for the ccipSend/executeReports messages
- _NONCES_: Use nonces for the ccipSend message per user, so they can execute ordered transactions on destination chain.
- _ORDERED EXECUTION_: Validate nonces when executing reports in the Off Ramp.

## Testing

- The Rust Tests execute logic specific to the program, to run them use the command

  ```bash
  cargo test
  ```

- The Anchor Tests are in the `tests` folder, written in Typescript and use the `@project-serum/solana-web3` library. The tests are run in a local network, so the tests are fast and don't require any real SVM network. To run them, use the command

  ```bash
  anchor test
  ```
