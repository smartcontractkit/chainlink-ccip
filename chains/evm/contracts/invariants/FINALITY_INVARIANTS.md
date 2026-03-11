# Finality / Faster-Than-Finality (FTF) Invariants

## 1. Encoding and Representation

### 1.1 ExtraArgs

- **INV-FIN-ENC-1**: `blockConfirmations` is a `uint16` field in `GenericExtraArgsV3`, encoded at byte offset 8 (after tag + gasLimit) in the wire format.
- **INV-FIN-ENC-2**: `blockConfirmations == 0` means "wait for default finality as determined by the CCV". This is the safe default.
- **INV-FIN-ENC-3**: `blockConfirmations != 0` means the sender requests Faster-Than-Finality (FTF) with that many block confirmations. This shifts re-org risk to the receiver and pool.
- **INV-FIN-ENC-4**: Legacy extraArgs (V1/V2) do not contain `blockConfirmations`. When legacy args are used, `blockConfirmations` defaults to `0` (full finality).
- **INV-FIN-ENC-5**: The helper `_getBasicEncodedExtraArgsV3(gasLimit, blockConfirmations)` encodes a minimal V3 with only these two fields set, all others zeroed.

### 1.2 Message

- **INV-FIN-ENC-6**: `MessageV1.finality` is a `uint16` populated from `resolvedExtraArgs.blockConfirmations` on the OnRamp. It is encoded at bytes 33–35 in the wire format.
- **INV-FIN-ENC-7**: `MessageV1.messageNumber` uniqueness is NOT guaranteed under FTF. After a re-org, a message could end up with a different message number. Only messages older than the source chain's finality delay have guaranteed unique message numbers.

---

## 2. Source Side (OnRamp) Invariants

### 2.1 General Flow

- **INV-FIN-SRC-1**: `blockConfirmations` flows from ExtraArgsV3 through the OnRamp to: the message's `finality` field, every CCV's `getFee`, the pool's `getFee`, the pool's `getRequiredCCVs`, the pool's `lockOrBurn`, and the executor's `getFee`.
- **INV-FIN-SRC-2**: The OnRamp does not validate or constrain `blockConfirmations` itself. Validation is delegated to downstream participants (CCVs, pools, executor).

### 2.2 CCV Interaction

- **INV-FIN-SRC-3**: Each CCV receives `blockConfirmations` in its `getFee(destChainSelector, message, ccvArgs, blockConfirmations)` call. CCVs may price FTF differently or reject the value by reverting.
- **INV-FIN-SRC-4**: CCVs receive `blockConfirmations` for fee quoting only. The actual "waiting" for block depth happens offchain. On-chain `verifyMessage` receives the full `MessageV1` (which includes `finality`).

### 2.3 Pool Interaction

- **INV-FIN-SRC-5**: V2 pools (`IPoolV2`) receive `blockConfirmationsRequested` in `lockOrBurn`, `getFee`, and `getRequiredCCVs`. They can reject FTF by reverting.
- **INV-FIN-SRC-6**: V1 pools (`IPoolV1`) do NOT receive `blockConfirmationsRequested`. If `blockConfirmationsRequested != 0`, the OnRamp reverts with `CustomBlockConfirmationsNotSupportedOnPoolV1`.
- **INV-FIN-SRC-7**: Similarly, if `tokenArgs` is non-empty with a V1 pool, the OnRamp reverts with `TokenArgsNotSupportedOnPoolV1`.

### 2.4 Executor Interaction

- **INV-FIN-SRC-8**: The executor receives `requestedBlockDepth` (= `blockConfirmations`) in `getFee`. The executor may reject the value if it is below its configured `minBlockConfirmations`.
- **INV-FIN-SRC-9**: `Executor.getFee` reverts with `Executor__RequestedBlockDepthTooLow` if `requestedBlockDepth != 0 && requestedBlockDepth < s_dynamicConfig.minBlockConfirmations`. This means `0` (full finality) is always accepted.
- **INV-FIN-SRC-10**: The executor's `getMinBlockConfirmations()` returns the minimum FTF depth it will service. Finality (`0`) is always accepted regardless of this value.

---

## 3. Destination Side (OffRamp) Invariants

### 3.1 Receiver Finality Requirements

- **INV-FIN-DST-1**: The OffRamp calls `IAny2EVMMessageReceiverV2.getCCVsAndMinBlockDepth(sourceChainSelector, sender)` on receivers that implement the V2 interface. This returns `minBlockDepth` alongside CCV requirements.
- **INV-FIN-DST-2**: `minBlockDepth == 0` means the receiver only accepts fully finalized messages. Any FTF message (`message.finality != 0`) is rejected with `InvalidFinalityForReceiver`.
- **INV-FIN-DST-3**: `minBlockDepth > 0` means the receiver allows FTF, but only if `message.finality >= minBlockDepth`. If the message's block depth is below the receiver's minimum, it is rejected with `InvalidFinalityForReceiver`.
- **INV-FIN-DST-4**: `minBlockDepth == 1` with a trusted sender pattern is the simplest way to allow any FTF level, as documented in the interface comments.

### 3.2 Non-V2 Receivers

- **INV-FIN-DST-5**: If the receiver does NOT implement `IAny2EVMMessageReceiverV2`, `minBlockDepth` defaults to `0`. FTF messages are rejected. This protects receivers that have not explicitly opted in.
- **INV-FIN-DST-6**: Non-programmable accounts (e.g. EOAs on EVM) cannot implement the receiver V2 interface. They cannot receive FTF messages.

### 3.3 Finality and Finalized Messages

- **INV-FIN-DST-7**: If `message.finality == 0` (full finality), the receiver's `minBlockDepth` check is skipped entirely. Finalized messages are always accepted regardless of the receiver's configuration.
- **INV-FIN-DST-8**: `minBlockDepth` is only enforced when `message.finality != 0`.

### 3.4 Token-Only Transfers

- **INV-FIN-DST-9**: For token-only transfers, receiver CCVs are skipped, but `message.finality` is still passed to the pool via `_releaseOrMintSingleToken`. The pool's own finality checks still apply.
- **INV-FIN-DST-10**: For token-only transfers, the receiver's `minBlockDepth` is NOT checked (since `_getCCVsFromReceiver` is skipped).

### 3.5 Pool Finality (Inbound)

- **INV-FIN-DST-11**: The OffRamp calls `IPoolV2.releaseOrMint(releaseOrMintIn, blockConfirmationsRequested)` with `message.finality` as `blockConfirmationsRequested`.
- **INV-FIN-DST-12**: V1 pools receive no finality parameter; the V1 `releaseOrMint` is called without `blockConfirmationsRequested` (internally treated as `WAIT_FOR_FINALITY`).
- **INV-FIN-DST-13**: The OffRamp also passes `message.finality` to `_getCCVsFromPool` so that pools can require different CCVs for FTF messages.

---

## 4. Pool Finality Invariants

### 4.1 Configuration

- **INV-FIN-POOL-1**: `WAIT_FOR_FINALITY = 0` is the constant representing default/full finality in `TokenPool`.
- **INV-FIN-POOL-2**: `s_minBlockConfirmations` is the pool's configured minimum block depth for FTF. A value of `0` means the pool does not support FTF.
- **INV-FIN-POOL-3**: `s_minBlockConfirmations` is set by the pool owner via `setMinBlockConfirmations`.

### 4.2 Outbound (lockOrBurn)

- **INV-FIN-POOL-4**: If `blockConfirmationsRequested != 0` and `s_minBlockConfirmations == 0`, the pool reverts with `CustomBlockConfirmationsNotEnabled`. Pools must explicitly opt in to FTF.
- **INV-FIN-POOL-5**: If `blockConfirmationsRequested != 0` and `blockConfirmationsRequested < s_minBlockConfirmations`, the pool reverts with `InvalidMinBlockConfirmations`. The requested depth must meet or exceed the pool's minimum.
- **INV-FIN-POOL-6**: V1 `lockOrBurn` (without finality parameter) internally passes `WAIT_FOR_FINALITY` to `_validateLockOrBurn`.

### 4.3 Inbound (releaseOrMint)

- **INV-FIN-POOL-7**: `_validateReleaseOrMint` applies the same finality-based rate limit routing as `_validateLockOrBurn` (see rate limiting below).
- **INV-FIN-POOL-8**: V1 `releaseOrMint` (without finality parameter) delegates to V2 `releaseOrMint` with `WAIT_FOR_FINALITY`.

### 4.4 Rate Limiting

- **INV-FIN-POOL-9**: Pools maintain separate rate limit buckets for default-finality and custom-block-confirmation transfers. This provides isolated risk profiles.
- **INV-FIN-POOL-10**: If `blockConfirmationsRequested != WAIT_FOR_FINALITY`, the custom-block-confirmations rate limiter is consumed (`_consumeCustomBlockConfirmationsOutboundRateLimit` / `_consumeCustomBlockConfirmationsInboundRateLimit`).
- **INV-FIN-POOL-11**: If the custom-block-confirmations rate limiter is not configured (not enabled), it falls back to the default rate limiter. This means FTF transfers consume from the same bucket as finalized transfers when no separate bucket is set.
- **INV-FIN-POOL-12**: If `blockConfirmationsRequested == WAIT_FOR_FINALITY`, the default rate limiter is consumed.
- **INV-FIN-POOL-13**: Rate limiter configuration is per-remote-chain. Each chain can have its own default and custom-block-confirmation buckets.
- **INV-FIN-POOL-14**: The `RateLimiterConfig` struct has a `customBlockConfirmations` boolean to indicate which bucket set is being configured.

### 4.5 Fees

- **INV-FIN-POOL-15**: `TokenTransferFeeConfig` has separate fee fields for default and custom block confirmations:
  - `defaultBlockConfirmationsFeeUSDCents` vs `customBlockConfirmationsFeeUSDCents` (flat fee)
  - `defaultBlockConfirmationsTransferFeeBps` vs `customBlockConfirmationsTransferFeeBps` (proportional fee deducted from transferred amount)
- **INV-FIN-POOL-16**: `_getFee` (internal) uses `customBlockConfirmationsTransferFeeBps` when `blockConfirmationsRequested != WAIT_FOR_FINALITY`, otherwise `defaultBlockConfirmationsTransferFeeBps`.
- **INV-FIN-POOL-17**: `getFee` (external, for quoting) returns the custom or default fee config based on `blockConfirmationsRequested`.
- **INV-FIN-POOL-18**: `customBlockConfirmationsTransferFeeBps` must be less than `BPS_DIVIDER` (10,000). Validated at config time.

### 4.6 Pool CCV Selection

- **INV-FIN-POOL-19**: `getRequiredCCVs` receives `blockConfirmationsRequested`. However, in the current `AdvancedPoolHooks` implementation, this parameter is unused — CCV selection is based on amount thresholds only, not finality.
- **INV-FIN-POOL-20**: Custom pool hook implementations may use `blockConfirmationsRequested` to require additional or different CCVs for FTF messages.
- **INV-FIN-POOL-21**: `preflightCheck` and `postflightCheck` in `AdvancedPoolHooks` receive `blockConfirmationsRequested` but do not use it in the current implementation.

---

## 5. Executor Finality Invariants

- **INV-FIN-EXEC-1**: The executor has a `minBlockConfirmations` in its `DynamicConfig`. This is the minimum FTF depth the executor will service.
- **INV-FIN-EXEC-2**: `Executor.getFee` reverts with `Executor__RequestedBlockDepthTooLow` if `requestedBlockDepth != 0 && requestedBlockDepth < minBlockConfirmations`. Full finality (`requestedBlockDepth == 0`) always passes this check because finality is the safest option.
- **INV-FIN-EXEC-3**: `minBlockConfirmations` on the executor is the floor for FTF block depth. For example, if `minBlockConfirmations == 10`, the executor will service finality (`0`) and any FTF depth `>= 10`, but reject depths `1–9`.
- **INV-FIN-EXEC-4**: The executor enforces a CCV allowlist (if enabled) and a max CCVs per message limit, both of which are independent of finality.

---

## 6. CCV Finality Invariants

- **INV-FIN-CCV-1**: CCVs receive `blockConfirmations` in `getFee` for pricing. They may charge more for FTF (lower block depth) due to higher risk.
- **INV-FIN-CCV-2**: CCVs can reject a `blockConfirmations` value by reverting in `getFee`, which prevents the message from being sent.
- **INV-FIN-CCV-3**: `verifyMessage` receives the full `MessageV1` which includes `finality`. The CCV can use this during verification (e.g., verify that enough blocks have passed).
- **INV-FIN-CCV-4**: The actual waiting for block confirmations is an offchain concern. The CCV contract verifies the proof, not the block depth. Block depth enforcement is the responsibility of the offchain system (executor + CCV nodes).

---

## 7. Re-org and Safety Invariants

- **INV-FIN-REORG-1**: `MessageV1.messageNumber` is unique per-lane only for finalized messages. Under FTF, a re-org may cause a message to receive a different `messageNumber`.
- **INV-FIN-REORG-2**: FTF shifts re-org risk to the receiver, pool, and any downstream integrators. The protocol makes this explicit by requiring opt-in at every layer.
- **INV-FIN-REORG-3**: The OffRamp's execution states are keyed by `messageId` (= `keccak256(encodedMessage)`). A re-orged message would have a different `messageId` and could theoretically be executed again with a different payload.
- **INV-FIN-REORG-4**: The separate rate limit buckets for FTF (INV-FIN-POOL-9) limit the financial exposure from re-org scenarios.

---

## 8. Opt-in Requirements (Defense in Depth)

FTF requires explicit opt-in at every layer of the stack:

| Layer | Opt-in Mechanism | Default |
|-------|-----------------|---------|
| **Sender** | Sets `blockConfirmations != 0` in ExtraArgsV3 | `0` (finality) |
| **Receiver** | Implements `IAny2EVMMessageReceiverV2` and returns `minBlockDepth > 0` | `0` (finality required) |
| **Pool** | Sets `s_minBlockConfirmations > 0` | `0` (FTF disabled) |
| **Executor** | Configures `minBlockConfirmations` in DynamicConfig | Finality always accepted; `minBlockConfirmations` is floor for FTF |
| **CCV** | Does not revert in `getFee` for the given `blockConfirmations` | Implementation-specific |

- **INV-FIN-OPTIN-1**: If any single layer rejects FTF, the message cannot be sent or executed with FTF.
- **INV-FIN-OPTIN-2**: All layers default to rejecting or not supporting FTF. The executor always accepts finality (`0`) and uses `minBlockConfirmations` as a floor for FTF depth, not as a toggle.
- **INV-FIN-OPTIN-3**: V1 pools and non-V2 receivers implicitly reject FTF with no code changes needed. Backward compatibility is preserved.

---

## 9. Cross-Cutting Invariants

- **INV-FIN-CC-1**: `finality` is a message-level field. All tokens in a message share the same finality requirement.
- **INV-FIN-CC-2**: `finality` affects fees (CCV, pool, executor may all charge differently for FTF), rate limiting (separate buckets), and CCV requirements (pools may require different CCVs).
- **INV-FIN-CC-3**: `finality` does NOT affect the `ccvAndExecutorHash` computation. The hash is based on CCV addresses and executor address only.
- **INV-FIN-CC-4**: The FeeQuoter does not use `blockConfirmations`. Finality-aware fees flow through the pool's own `getFee` and the CCV's `getFee`, not through the FeeQuoter.
