# Message Lifecycle Invariants

## 1. Message Sequencing

- **INV-SEQ-1**: `messageNumber` is per-lane: each `(OnRamp, destChainSelector)` pair maintains its own counter.
- **INV-SEQ-2**: `messageNumber` is strictly monotonic. It is pre-incremented before assignment — the first message on a lane has `messageNumber = 1`.
- **INV-SEQ-3**: `messageNumber = 0` is reserved and never assigned to a real message.
- **INV-SEQ-4**: `messageNumber` persists across lane configuration updates. Reconfiguring a lane does not reset the counter.
- **INV-SEQ-5**: Under FTF, `messageNumber` uniqueness is not guaranteed. A re-org may cause a message to receive a different `messageNumber`. Only messages older than the source chain's finality delay have guaranteed unique message numbers.

## 2. Message Identity

- **INV-ID-1**: `messageId = keccak256(encodedMessageV1)`. The message ID is deterministic from the encoded message content.
- **INV-ID-2**: All chains compute `messageId` identically — the encoding format and hash function are protocol-level requirements.
- **INV-ID-3**: Execution outcomes are keyed by `messageId`. Two messages with different content produce different IDs and are tracked independently.

---

## 3. Source Side (OnRamp) Flow

The OnRamp processes a message in a fixed order:


| Step | Operation                                                    |
| ---- | ------------------------------------------------------------ |
| 1    | Parse extraArgs, apply defaults                              |
| 2    | Build message struct (assign `messageNumber`)                |
| 3    | Merge CCV lists (user + lane-mandated + pool)                |
| 4    | Compute `ccvAndExecutorHash`                                 |
| 5    | Compute fees (CCV fees, pool fee, executor fee, network fee) |
| 6    | Verify sufficient fee token provided                         |
| 7    | Distribute fees to CCVs, executor, pool                      |
| 8    | `lockOrBurn` (if tokens)                                     |
| 9    | Encode message, compute `messageId`                          |
| 10   | Forward to verifiers                                         |
| 11   | Emit event                                                   |


- **INV-SRC-1**: Fees are computed and distributed before `lockOrBurn`. Tokens are locked/burned only after fees are paid.
- **INV-SRC-2**: `messageId` is computed after `lockOrBurn`, because the token transfer encoding (including `destPoolData` from `lockOrBurn`) is part of the encoded message.
- **INV-SRC-3**: The computed fee must not exceed the fee token amount provided by the sender.
- **INV-SRC-4**: A message may contain at most one token transfer.
- **INV-SRC-5**: Token amount must be non-zero if a token transfer is included.

---

## 4. Destination Side (OffRamp) Flow

The OffRamp processes a message in a fixed order:


| Step | Operation                                                               |
| ---- | ----------------------------------------------------------------------- |
| 1    | Decode and validate message (source chain, OnRamp, OffRamp, dest chain) |
| 2    | Check execution state (must be untouched or failed)                     |
| 3    | Mark message as in-progress                                             |
| 4    | CCV verification (quorum check + `verifyMessage` calls)                 |
| 5    | `releaseOrMint` (if tokens)                                             |
| 6    | If token-only: return (no callback)                                     |
| 7    | Deliver message to receiver (`ccipReceive`)                             |
| 8    | Mark message as succeeded or failed                                     |


- **INV-DST-1**: CCV verification happens before token release. Tokens are not released until verification passes.
- **INV-DST-2**: `releaseOrMint` happens before `ccipReceive`. The receiver has access to the released tokens during the callback.
- **INV-DST-3**: For token-only transfers, `ccipReceive` is not called. Only CCV verification and `releaseOrMint` are performed.

---

## 5. Execution Semantics

### 5.1 Permissionless Execution

- **INV-EXEC-1**: Message execution is permissionless. Anyone can call `execute` with valid proofs.

### 5.2 Ordering

- **INV-EXEC-2**: There is no in-protocol ordering guarantee. Messages can be executed in any order regardless of `messageNumber`.
- **INV-EXEC-3**: `messageNumber` is a sequence number for identification and offchain coordination. It does not constrain on-chain execution order.

### 5.3 Execution Outcome

- **INV-EXEC-4**: A message that has never been executed may be executed.
- **INV-EXEC-5**: A message that failed may be retried.
- **INV-EXEC-6**: A successfully executed message cannot be re-executed.
- **INV-EXEC-7**: A retry that still fails must not produce a redundant state transition.

### 5.4 Gas Override

- **INV-EXEC-8**: The executor may provide a `gasLimitOverride` when executing a message.
- **INV-EXEC-9**: If `gasLimitOverride` is non-zero, it must be greater than or equal to the message's `ccipReceiveGasLimit`. This allows increasing gas for retries but not decreasing it.
- **INV-EXEC-10**: If `gasLimitOverride` is zero, the message's original `ccipReceiveGasLimit` is used.

---

## 6. Receiver Callback (`ccipReceive`)

- **INV-RECV-1**: `ccipReceive` is called on `message.receiver` via the router. Only configured OffRamps may trigger this call path.
- **INV-RECV-2**: The gas limit for `ccipReceive` is `gasLimitOverride` (if non-zero) or `message.ccipReceiveGasLimit`. The sender pays for the specified gas; unspent gas is not refunded.
- **INV-RECV-3**: `ccipReceive` is not called for token-only transfers.
- **INV-RECV-4**: If `ccipReceive` reverts, the message is marked as failed and may be retried.

---

## 7. Token Receiver vs Message Receiver

- **INV-TR-1**: `message.receiver` is the address that receives the `ccipReceive` callback.
- **INV-TR-2**: `tokenReceiver` (from ExtraArgsV3) is the address that receives transferred tokens. It may differ from `message.receiver`.
- **INV-TR-3**: If `tokenReceiver` is empty in ExtraArgsV3, `message.receiver` is used for token delivery.
- **INV-TR-4**: `tokenReceiverAllowed` is a per-lane configuration flag. If false, specifying a `tokenReceiver` that differs from the message receiver reverts.

---

## 8. Token-Only Transfer Behavior

- **INV-TO-1**: A "token-only transfer" is defined as: `gasLimit == 0`, `data` is empty, and at least one token transfer is present.
- **INV-TO-2**: For token-only transfers, `ccipReceive` is not called.
- **INV-TO-3**: CCV verification and `releaseOrMint` still apply to token-only transfers.
- **INV-TO-4**: For token-only transfers, receiver CCVs and default CCVs are excluded on both source and destination. Only pool-required CCVs and lane-mandated CCVs apply.
- **INV-TO-5**: Receiver finality requirements (`minBlockDepth`) are not checked for token-only transfers.

---

## 9. No-Execution Address

- **INV-NOEXEC-1**: A special "no-execution" address can be specified as the executor. This signals that no automated execution is expected on the destination.
- **INV-NOEXEC-2**: When the no-execution address is used, the executor fee is zero and execution cost is not added to the message fee.
- **INV-NOEXEC-3**: Tokens are still released/minted for no-execution messages. Only the automated execution is skipped — manual execution is still possible.

