# CCV (Cross-Chain Verifier) Invariants

## 1. Configuration Invariants

### 1.1 OnRamp DestChainConfig

- **INV-CFG-1**: At least one CCV must exist across `defaultCCVs` and `laneMandatedCCVs` combined. A lane cannot be configured with zero total CCVs.
- **INV-CFG-2**: No zero-value address is allowed in `defaultCCVs` or `laneMandatedCCVs`.
- **INV-CFG-3**: No duplicates are allowed within `defaultCCVs`, within `laneMandatedCCVs`, or across both sets.
- **INV-CFG-4**: `defaultCCVs` and `laneMandatedCCVs` are validated at config time. All downstream code assumes no duplicates or zero-value addresses in these sets.

### 1.2 OffRamp SourceChainConfig

- **INV-CFG-5**: `defaultCCVs` must be non-empty (at least one default CCV per source chain).
- **INV-CFG-6**: No zero-value address is allowed in `defaultCCVs` or `laneMandatedCCVs`.
- **INV-CFG-7**: No duplicates are allowed within or across `defaultCCVs` and `laneMandatedCCVs`, enforced by `CCVConfigValidation._validateDefaultAndMandatedCCVs`.

### 1.3 Pool CCV Config (AdvancedPoolHooks)

- **INV-CFG-8**: `outboundCCVs` and `thresholdOutboundCCVs` must not contain duplicates within themselves or between each other.
- **INV-CFG-9**: `inboundCCVs` and `thresholdInboundCCVs` must not contain duplicates within themselves or between each other.

---

## 2. Source Side (OnRamp) Invariants

### 2.1 User-Requested CCVs

- **INV-SRC-1**: User-provided CCVs (from `ExtraArgsV3`) must not contain duplicates.
- **INV-SRC-2**: `ccvs` and `ccvArgs` arrays must have the same length.
- **INV-SRC-3**: The zero-value address in the user CCV list is a placeholder meaning "include defaults". There can be at most one (since duplicates are rejected).
- **INV-SRC-4**: Legacy extraArgs (pre-V3) do not include CCVs. The user gets defaults applied automatically.
- **INV-SRC-5**: If the user provides no CCVs (empty list), defaults are applied for data/receiver messages. For token-only transfers, an empty list results in no user/default CCVs (only pool + lane-mandated apply).

### 2.2 Default CCVs

- **INV-SRC-6**: Defaults are used when the user provides no CCVs or uses the zero-value address as a placeholder.
- **INV-SRC-7**: Defaults are NOT applied for pure token-only transfers (no data, gasLimit=0, has tokens) unless the user explicitly requests them via the zero-value address placeholder.
- **INV-SRC-8**: When expanding the zero-value placeholder with defaults, defaults that already appear in the user list are skipped (deduplication).

### 2.3 Lane-Mandated CCVs

- **INV-SRC-9**: Lane-mandated CCVs are always added to every message on the lane, regardless of user or pool preferences.
- **INV-SRC-10**: Lane-mandated CCVs that already appear in the user/default list are skipped (deduplication).
- **INV-SRC-11**: If a user includes a lane-mandated CCV in their own CCV list with custom `ccvArgs`, those args are preserved (the user entry takes precedence due to deduplication). If the user does not include it, the lane-mandated CCV is added with empty `ccvArgs`.

### 2.4 Pool-Required CCVs

- **INV-SRC-12**: If a pool does not implement `IPoolV2`, it falls back to `defaultCCVs`.
- **INV-SRC-13**: If a pool implements `IPoolV2` but returns an empty array from `getRequiredCCVs`, it falls back to `defaultCCVs`.
- **INV-SRC-14**: If a pool returns the zero-value address in its CCV list, those entries are replaced with `defaultCCVs`.
- **INV-SRC-15**: Pool-required CCVs that already appear in the merged list (user + lane-mandated) are skipped (deduplication).
- **INV-SRC-16**: If a user includes a pool-required CCV in their own CCV list with custom `ccvArgs`, those args are preserved. If the user does not include it, the pool-required CCV is added with empty `ccvArgs`.

### 2.5 Merge Order

- **INV-SRC-17**: The final CCV list is built in order: user/default first, lane-mandated second, pool-required last. This order is deterministic but not a protocol-level requirement.
- **INV-SRC-18**: The final merged CCV list contains no duplicates.

### 2.6 CCV Fees

- **INV-SRC-19**: Each CCV's `getFee` is called to compute per-CCV fees (USD cents, gas, bytes overhead).

### 2.7 CCV+Executor Hash

- **INV-SRC-21**: `ccvAndExecutorHash` is computed as `keccak256([addressLength][ccv1..ccvN][executor])`. The length prefix encodes the address byte length for the chain family, preventing collisions between different array sizes and cross-chain-family ambiguity.
- **INV-SRC-22**: `ccvAndExecutorHash` is embedded in the message for offchain validation only. It is NOT verified on the destination chain.

### 2.8 Token-Only Transfer Behavior

- **INV-SRC-23**: A "token-only transfer" is defined as: `gasLimit == 0 && message.data.length == 0 && message.tokenAmounts.length > 0`.
- **INV-SRC-24**: For token-only transfers with no user-specified CCVs, only pool-required CCVs and lane-mandated CCVs are included. Default CCVs and receiver CCVs are excluded, since only the token issuer (pool) bears risk, not the receiver.

---

## 3. Destination Side (OffRamp) Invariants

### 3.1 Receiver-Required CCVs

- **INV-DST-1**: If the receiver implements `IAny2EVMMessageReceiverV2`, its `getCCVsAndMinBlockDepth` is called to get `requiredCCVs`, `optionalCCVs`, `optionalThreshold`, and `minBlockDepth`.
- **INV-DST-2**: If the receiver does NOT implement `IAny2EVMMessageReceiverV2`, it cannot specify custom CCVs. Default CCVs are used. FTF is also not supported.
- **INV-DST-3**: `requiredCCVs` returned by the receiver must not contain duplicates. Enforced by `CCVConfigValidation._assertNoDuplicates`.
- **INV-DST-4**: `optionalCCVs` returned by the receiver must not contain duplicates. Enforced by `CCVConfigValidation._assertNoDuplicates`.
- **INV-DST-5**: `optionalThreshold` must not exceed `optionalCCVs.length`. Otherwise reverts with `InvalidOptionalThreshold`.
- **INV-DST-6**: If the receiver returns empty `requiredCCVs` AND `optionalThreshold == 0`, the system falls back to default CCVs via a zero-value address placeholder.
- **INV-DST-7**: The zero-value address in required CCVs is a placeholder for "include defaults".

### 3.2 Pool-Required CCVs (Inbound)

- **INV-DST-8**: If a pool does not implement `IPoolV2`, it falls back to `defaultCCVs`.
- **INV-DST-9**: If a pool implements `IPoolV2` but returns an empty array, it falls back to `defaultCCVs`.
- **INV-DST-10**: If a pool returns the zero-value address, those entries trigger inclusion of `defaultCCVs`.
- **INV-DST-11**: Pool CCV requirements apply for both inbound and outbound directions. The pool's `getRequiredCCVs` is called with `MessageDirection.Inbound` on the OffRamp.

### 3.3 Lane-Mandated CCVs (Inbound)

- **INV-DST-12**: Lane-mandated CCVs from `SourceChainConfig` are always included in the required set.

### 3.4 Default CCVs (Inbound)

- **INV-DST-13**: Default CCVs are included whenever any zero-value address appears in the combined required list (receiver + pool + lane-mandated).
- **INV-DST-14**: Defaults are added at most once, even if multiple zero-value address entries exist.

### 3.5 Token-Only Transfer Behavior (Inbound)

- **INV-DST-15**: For token-only transfers with tokens, receiver CCVs and defaults are skipped. Only pool-required CCVs and lane-mandated CCVs are used.
- **INV-DST-16**: For token-only transfers without tokens (no-op), defaults are used.

### 3.6 Deduplication and Overlap

- **INV-DST-17**: The final required CCV list is deduplicated: duplicates and zero-value address entries are removed.
- **INV-DST-18**: If a CCV appears in both the required and optional lists, it is removed from the optional list and the `optionalThreshold` is decremented (but not below zero). This prevents double-counting.

### 3.7 Quorum Enforcement

- **INV-DST-19**: ALL required CCVs must be present in the `ccvs` array passed to `execute`. If any is missing, the execution reverts with `RequiredCCVMissing`.
- **INV-DST-20**: At least `optionalThreshold` of the optional CCVs must be present in the `ccvs` array. If not enough are found, the execution reverts with `OptionalCCVQuorumNotReached`.
- **INV-DST-21**: The executor may pass extra CCVs beyond what is required/optional. These are ignored (not queried for verification).

### 3.8 Verification

- **INV-DST-22**: Each CCV that passes quorum checks has `verifyMessage` called on it. Verification succeeds by not reverting. Any revert from `verifyMessage` causes the entire message execution to fail.
- **INV-DST-23**: `ccvs.length` must equal `verifierResults.length`. Each CCV has exactly one corresponding verifier result.

### 3.9 Execution State

- **INV-DST-25**: If CCV verification fails, the message state is set to `FAILURE` and can be retried later (possibly with different/updated verifier results).
- **INV-DST-26**: If CCV verification succeeds but `ccipReceive` fails, the message state is still `FAILURE` and can be retried.
- **INV-DST-27**: A message in `SUCCESS` state cannot be re-executed. Only `UNTOUCHED` or `FAILURE` states allow execution.
- **INV-DST-28**: A `FAILURE` re-execution that still fails reverts with `NoStateProgressMade` (no state change from FAILURE to FAILURE).

---

## 4. Finality / Faster-Than-Finality (FTF) Invariants

- **INV-FTF-1**: `message.finality` (blockConfirmations) of `0` means "wait for full finality".
- **INV-FTF-2**: `message.finality != 0` means the sender requested FTF with that many block confirmations.
- **INV-FTF-3**: If the receiver does NOT implement `IAny2EVMMessageReceiverV2`, `minBlockDepth` defaults to `0` (finality required). FTF messages are rejected.
- **INV-FTF-4**: If the receiver returns `minBlockDepth == 0`, only finalized messages are accepted. FTF messages (`finality != 0`) are rejected with `InvalidFinalityForReceiver`.
- **INV-FTF-5**: If the receiver returns `minBlockDepth > 0`, FTF is allowed only if `message.finality >= minBlockDepth`. Otherwise rejected with `InvalidFinalityForReceiver`.
- **INV-FTF-6**: For non-programmable accounts (e.g. EOAs on EVM), `minBlockDepth` is `0` and FTF is not supported.
- **INV-FTF-7**: `blockConfirmations` from `ExtraArgsV3` is passed to CCV `getFee` on the source side, allowing CCVs to price FTF differently.

---

## 5. Fee Invariants

- **INV-FEE-1**: Each CCV in the final merged list generates a receipt with `feeUSDCents`, `gasForVerification`, and `ccvPayloadSizeBytes`.
- **INV-FEE-2**: CCV fees are computed by calling `ICrossChainVerifierV1(impl).getFee(destChainSelector, message, ccvArgs, blockConfirmations)` on the resolved implementation.
- **INV-FEE-3**: The total message fee includes the sum of all CCV fees, token transfer fees, executor fee, and network fee.
- **INV-FEE-4**: CCV fees are denominated in USD cents and converted to fee token amounts using the FeeQuoter's token prices.
- **INV-FEE-5**: More CCVs = higher fees. Lane-mandated and pool-required CCVs add to the user's fee even if the user did not request them.

---

## 6. Encoding Invariants

- **INV-ENC-1**: `ExtraArgsV3` tag is `0xa69dd4aa`. Any extraArgs with this 4-byte prefix is decoded as V3.
- **INV-ENC-2**: CCV addresses in ExtraArgsV3 are encoded as `uint8 addressLength + bytes address`. `addressLength == 0` encodes the zero-value address (the default placeholder).
- **INV-ENC-3**: CCV args in ExtraArgsV3 are encoded as `uint16 argsLength + bytes args`.
- **INV-ENC-4**: The number of CCVs is encoded as a single `uint8`, limiting a message to 255 user-specified CCVs.

---

## 7. CCV Address Stability and Resolution

A CCV address is referenced by users, pools, lane configs, and receivers. Verifier implementations must be upgradeable without changing the address that all these parties reference. Each chain family must provide a mechanism to satisfy this requirement.

- **INV-ADDR-1**: A CCV address must remain stable across verifier upgrades. Users, pools, lane configs, and receivers all reference the CCV by this address — changing it would require coordinated reconfiguration across all parties.
- **INV-ADDR-2**: The CCV address must resolve to a concrete verifier implementation for each supported destination chain (outbound) and for each verifier result format (inbound). If resolution fails, the message cannot be sent or executed.
- **INV-ADDR-3**: Resolution must be transparent to the caller. The OnRamp, OffRamp, and fee quoting logic interact with the CCV address; the address is responsible for routing to the correct implementation.

### 7.1 EVM: Resolver Pattern

On EVM, address stability is achieved via resolvers that implement `ICrossChainVerifierResolver`:

- **INV-RES-1**: CCV addresses on EVM are resolver contract instances, not direct verifier implementations. Resolvers maintain a mapping from context (destination chain or version) to implementation.
- **INV-RES-2**: `VersionedVerifierResolver` maps outbound traffic by `destChainSelector` and inbound traffic by a 4-byte version prefix in `verifierResults`.
- **INV-RES-3**: If a resolver returns the zero-value address for outbound (`getOutboundImplementation`), the message cannot be sent.
- **INV-RES-4**: If a resolver returns the zero-value address for inbound (`getInboundImplementation`), the message cannot be executed.

---

## 8. Cross-Cutting Invariants

- **INV-CC-1**: Every non-token-only message is verified by at least one CCV (guaranteed by INV-CFG-1/5 and the default fallback mechanism).
- **INV-CC-2**: Token-only transfers may have zero receiver/default CCVs but will still have pool CCVs and/or lane-mandated CCVs if configured.
- **INV-CC-3**: The source-side CCV list (embedded in `ccvAndExecutorHash`) is informational for offchain systems. The destination side independently computes its own CCV requirements.
- **INV-CC-4**: Source and destination CCV lists are independently configured and may differ. The source-side list determines fees; the destination-side list determines verification requirements.
- **INV-CC-5**: A CCV that appears on both source and destination sides may resolve to different implementations, as resolution is handled independently per chain.

