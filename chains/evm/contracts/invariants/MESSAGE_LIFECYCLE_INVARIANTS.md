# Message Lifecycle Invariants

## 1. Message Sequencing

- **INV-SEQ-1**: `messageNumber` is per-lane: each `(OnRamp, destChainSelector)` pair maintains its own counter.
- **INV-SEQ-2**: `messageNumber` is strictly monotonic. It is pre-incremented before assignment — the first message on a lane has `messageNumber = 1`.
- **INV-SEQ-3**: `messageNumber = 0` is reserved and never assigned to a real message.
- **INV-SEQ-4**: `messageNumber` persists across lane configuration updates. Reconfiguring a lane does not reset the counter.

## 2. Message Identity

- **INV-ID-1**: `messageId = keccak256(encodedMessageV1)`. See ENCODING_INVARIANTS.md for the full wire format.
- **INV-ID-2**: Execution outcomes are keyed by `messageId`. Two messages with different content produce different IDs and are tracked independently.

---

## 3. Source Side (OnRamp) Flow

- **INV-SRC-1**: Fees are computed and distributed before `lockOrBurn`. Tokens are locked/burned only after fees are paid.
- **INV-SRC-2**: `messageId` is computed after `lockOrBurn`, because the token transfer encoding (including `destPoolData` from `lockOrBurn`) is part of the encoded message.
- **INV-SRC-3**: The computed fee must not exceed the fee token amount provided by the sender.
- **INV-SRC-4**: Token amount must be non-zero if a token transfer is included.

---

## 4. Destination Side (OffRamp) Flow

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

---

## 6. Receiver Callback (`ccipReceive`)

- **INV-RECV-1**: `ccipReceive` is called on `message.receiver`. Only configured OffRamps may trigger this call path.
- **INV-RECV-2**: `ccipReceive` is not called for token-only transfers.

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
- **INV-TO-4**: For token-only transfers, receiver CCVs and default CCVs are excluded on both source and destination. Only pool-required CCVs and lane-mandated CCVs apply. See CCV_INVARIANTS.md for details.
- **INV-TO-5**: Receiver finality requirements (`minBlockDepth`) are not checked for token-only transfers. See FINALITY_INVARIANTS.md for details.

---

## 9. No-Execution Address

- **INV-NOEXEC-1**: The no-execution address is `address(bytes20(0xeba517d2))` — the address-width representation of the `NO_EXECUTION_TAG` (`keccak256("NO_EXECUTION_TAG")[:4]`). This value is protocol-wide and must be the same on all chains.
- **INV-NOEXEC-2**: When the no-execution address is specified as the executor, no automated execution is expected on the destination. The executor fee is zero and execution cost is not added to the message fee.
- **INV-NOEXEC-3**: Tokens are still released/minted for no-execution messages. Only the automated execution is skipped — manual execution is still possible.
