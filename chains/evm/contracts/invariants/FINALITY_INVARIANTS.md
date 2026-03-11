# Finality / Faster-Than-Finality (FTF) Invariants

## 1. Encoding and Representation

### 1.1 ExtraArgs

- **INV-FIN-ENC-1**: `blockConfirmations == 0` means "wait for default finality as determined by the CCV". This is the safe default.
- **INV-FIN-ENC-2**: `blockConfirmations != 0` means the sender requests Faster-Than-Finality (FTF) with that many block confirmations. This shifts re-org risk to the receiver and pool.
- **INV-FIN-ENC-3**: Legacy extraArgs (V1/V2) do not contain `blockConfirmations`. When legacy args are used, `blockConfirmations` defaults to `0` (full finality).

---

## 2. Source Side (OnRamp) Invariants

### 2.1 General Flow

- **INV-FIN-SRC-1**: `blockConfirmations` flows from ExtraArgsV3 through the OnRamp to: the message's `finality` field, every CCV's `getFee`, the pool's `getFee`, the pool's `getRequiredCCVs`, the pool's `lockOrBurn`, and the executor's `getFee`.
- **INV-FIN-SRC-2**: The OnRamp does not validate or constrain `blockConfirmations` itself. Validation is delegated to downstream participants (CCVs, pools, executor).

### 2.2 CCV Interaction

- **INV-FIN-SRC-3**: Each CCV receives `blockConfirmations` in its `getFee(destChainSelector, message, ccvArgs, blockConfirmations)` call. CCVs may price FTF differently or reject the value by reverting.

### 2.3 Pool Interaction

- **INV-FIN-SRC-4**: V2 pools receive `blockConfirmationsRequested` in `lockOrBurn`, `getFee`, and `getRequiredCCVs`. They can reject FTF by reverting.
- **INV-FIN-SRC-5**: V1 pools do NOT receive `blockConfirmationsRequested`. If `blockConfirmationsRequested != 0`, the OnRamp reverts. Pools must be upgraded to V2 to support FTF.
- **INV-FIN-SRC-6**: Similarly, if `tokenArgs` is non-empty with a V1 pool, the OnRamp reverts.

### 2.4 Executor Interaction

- **INV-FIN-SRC-7**: The executor receives `requestedBlockDepth` (= `blockConfirmations`) in `getFee`. The executor may reject the value if it is below its configured `minBlockConfirmations`.
- **INV-FIN-SRC-8**: The executor rejects FTF requests where `requestedBlockDepth != 0 && requestedBlockDepth < minBlockConfirmations`. Full finality (`requestedBlockDepth == 0`) is always accepted.
- **INV-FIN-SRC-9**: The executor's `getMinBlockConfirmations()` returns the minimum FTF depth it will service. Finality (`0`) is always accepted regardless of this value.

---

## 3. Destination Side (OffRamp) Invariants

### 3.1 Receiver Finality Requirements

- **INV-FIN-DST-1**: The OffRamp queries the receiver for its CCV and finality requirements. This returns `minBlockDepth` alongside CCV requirements.
- **INV-FIN-DST-2**: `minBlockDepth == 0` means the receiver only accepts fully finalized messages. Any FTF message (`message.finality != 0`) is rejected.
- **INV-FIN-DST-3**: `minBlockDepth > 0` means the receiver allows FTF, but only if `message.finality >= minBlockDepth`. If the message's block depth is below the receiver's minimum, it is rejected.

### 3.2 Non-V2 Receivers

- **INV-FIN-DST-4**: If the receiver does NOT implement the V2 receiver interface, `minBlockDepth` defaults to `0`. FTF messages are rejected. This protects receivers that have not explicitly opted in.
- **INV-FIN-DST-5**: Non-programmable accounts (e.g. EOAs on EVM) cannot implement the V2 receiver interface. They cannot receive FTF messages unless the transfer is token-only (in which case receiver finality checks are skipped).

### 3.3 Token-Only Transfers

- **INV-FIN-DST-6**: For token-only transfers, receiver CCVs are skipped, but `message.finality` is still passed to the pool via `releaseOrMint`. The pool's own finality checks still apply.
- **INV-FIN-DST-7**: For token-only transfers, the receiver's `minBlockDepth` is NOT checked (since receiver CCV resolution is skipped).

### 3.4 Pool Finality (Inbound)

- **INV-FIN-DST-8**: The OffRamp passes `message.finality` as `blockConfirmationsRequested` to the V2 pool's `releaseOrMint`.
- **INV-FIN-DST-9**: V1 pools receive no finality parameter; they are called without `blockConfirmationsRequested` (internally treated as `WAIT_FOR_FINALITY`).
- **INV-FIN-DST-10**: The OffRamp also passes `message.finality` to the pool's CCV resolution so that pools can require different CCVs for FTF messages.

---

## 4. Pool Finality Invariants

### 4.1 Configuration

- **INV-FIN-POOL-1**: `WAIT_FOR_FINALITY = 0` is the constant representing default/full finality.
- **INV-FIN-POOL-2**: A pool's configured minimum block depth for FTF determines whether it supports FTF. A value of `0` means the pool does not support FTF.
- **INV-FIN-POOL-3**: The pool owner configures the minimum block depth.

### 4.2 Outbound (lockOrBurn)

- **INV-FIN-POOL-4**: If `blockConfirmationsRequested != 0` and the pool has not enabled FTF (minimum block depth == 0), the pool reverts. Pools must explicitly opt in to FTF.
- **INV-FIN-POOL-5**: If `blockConfirmationsRequested != 0` and `blockConfirmationsRequested` is below the pool's configured minimum block depth, the pool reverts. The requested depth must meet or exceed the pool's minimum.
- **INV-FIN-POOL-6**: V1 `lockOrBurn` (without finality parameter) internally passes `WAIT_FOR_FINALITY` to validation.

### 4.3 Inbound (releaseOrMint)

- **INV-FIN-POOL-7**: `releaseOrMint` validation applies the same finality-based rate limit routing as `lockOrBurn` (see rate limiting below).

### 4.4 Rate Limiting

- **INV-FIN-POOL-8**: Pools maintain separate rate limit buckets for default-finality and custom-block-confirmation transfers. This provides isolated risk profiles.
- **INV-FIN-POOL-9**: If `blockConfirmationsRequested != WAIT_FOR_FINALITY`, the custom-block-confirmations rate limiter is consumed.
- **INV-FIN-POOL-10**: If the custom-block-confirmations rate limiter is not configured (not enabled), it falls back to the default rate limiter. This means FTF transfers consume from the same bucket as finalized transfers when no separate bucket is set.
- **INV-FIN-POOL-11**: If `blockConfirmationsRequested == WAIT_FOR_FINALITY`, the default rate limiter is consumed.
- **INV-FIN-POOL-12**: Rate limiter configuration is per-remote-chain. Each chain can have its own default and custom-block-confirmation buckets.

### 4.5 Fees

- **INV-FIN-POOL-13**: Token transfer fee configuration has separate fee fields for default and custom block confirmations: separate flat fees and separate proportional fees (deducted from transferred amount).
- **INV-FIN-POOL-14**: `customBlockConfirmationsTransferFeeBps` must be less than `BPS_DIVIDER` (10,000). Validated at config time.

### 4.6 Pool CCV Selection

- **INV-FIN-POOL-15**: `getRequiredCCVs` receives `blockConfirmationsRequested`. Pool implementations may use this to require additional or different CCVs for FTF messages.

---

## 5. Executor Finality Invariants

- **INV-FIN-EXEC-1**: The executor has a configurable `minBlockConfirmations`. This is the minimum FTF depth the executor will service.
- **INV-FIN-EXEC-2**: The executor rejects FTF requests where `requestedBlockDepth != 0 && requestedBlockDepth < minBlockConfirmations`. Full finality (`requestedBlockDepth == 0`) always passes this check because finality is the safest option.
- **INV-FIN-EXEC-3**: `minBlockConfirmations` on the executor is the floor for FTF block depth. For example, if `minBlockConfirmations == 10`, the executor will service finality (`0`) and any FTF depth `>= 10`, but reject depths `1–9`.
- **INV-FIN-EXEC-4**: The executor enforces a CCV allowlist (if enabled) and a max CCVs per message limit, both of which are independent of finality.

---

## 6. CCV Finality Invariants

- **INV-FIN-CCV-1**: CCVs receive `blockConfirmations` in `getFee` for pricing. They may charge more for FTF (lower block depth) due to higher risk.
- **INV-FIN-CCV-2**: CCVs can reject a `blockConfirmations` value by reverting in `getFee`, which prevents the message from being sent.
- **INV-FIN-CCV-3**: The actual waiting for block confirmations is an offchain concern. The CCV contract verifies the proof, not the block depth. Block depth enforcement is the responsibility of the offchain CCV.

---

## 7. Re-org and Safety Invariants

- **INV-FIN-REORG-1**: `messageNumber` is unique per-lane only for finalized messages. Under FTF, a re-org may cause a message to receive a different `messageNumber`.
- **INV-FIN-REORG-2**: FTF shifts re-org risk to the receiver, pool, and any downstream integrators. The protocol makes this explicit by requiring opt-in at every layer.
- **INV-FIN-REORG-3**: Execution outcomes are keyed by `messageId` (= `keccak256(encodedMessage)`).

---

## 8. Opt-in Requirements (Defense in Depth)

FTF requires explicit opt-in at every layer of the stack:

| Layer | Opt-in Mechanism | Default |
|-------|------------------|---------|
| **Sender** | Sets `blockConfirmations != 0` in ExtraArgsV3 | `0` (finality) |
| **Receiver** | Implements V2 receiver interface and returns `minBlockDepth > 0` | `0` (finality required) |
| **Pool** | Configures minimum block depth > 0 | `0` (FTF disabled) |
| **Executor** | Configures `minBlockConfirmations` | Finality always accepted; `minBlockConfirmations` is floor for FTF |
| **CCV** | Does not revert in `getFee` for the given `blockConfirmations` | Implementation-specific |

- **INV-FIN-OPTIN-1**: If any single layer rejects FTF, the message cannot be sent or executed with FTF.
- **INV-FIN-OPTIN-2**: All layers default to rejecting or not supporting FTF.
- **INV-FIN-OPTIN-3**: V1 pools and non-V2 receivers implicitly reject FTF with no code changes needed. Backward compatibility is preserved.

---

## 9. Cross-Cutting Invariants

- **INV-FIN-CC-1**: `finality` is a message-level field. All tokens in a message share the same finality requirement.
- **INV-FIN-CC-2**: `finality` may affect fees (CCV, pool, executor may all charge differently for FTF), rate limiting (separate buckets), and CCV requirements (pools may require different CCVs).
