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
- **INV-SRC-2**: The zero-value address in the user CCV list is a placeholder meaning "include defaults". There can be at most one (since duplicates are rejected).
- **INV-SRC-3**: Legacy extraArgs (pre-V3) do not include CCVs. The user gets defaults applied automatically.
- **INV-SRC-4**: If the user provides no CCVs (empty list), defaults are applied for data/receiver messages. For token-only transfers, an empty list results in no user/default CCVs (only pool related + lane-mandated apply).

### 2.2 Default CCVs

- **INV-SRC-5**: Defaults are used when the user provides no CCVs or uses the zero-value address as a placeholder.
- **INV-SRC-6**: Defaults are NOT applied for pure token-only transfers (no data, gasLimit=0, has tokens) unless the user explicitly requests them via the zero-value address placeholder.
- **INV-SRC-7**: When expanding the zero-value placeholder with defaults, defaults that already appear in the user list are skipped (deduplication).

### 2.3 Lane-Mandated CCVs

- **INV-SRC-8**: Lane-mandated CCVs are always added to every message on the lane, regardless of user or pool preferences.
- **INV-SRC-9**: Lane-mandated CCVs that already appear in the user/default list are skipped (deduplication).
- **INV-SRC-10**: If a user includes a lane-mandated CCV in their own CCV list with custom `ccvArgs`, those args are preserved (the user entry takes precedence due to deduplication). If the user does not include it, the lane-mandated CCV is added with empty `ccvArgs`.

### 2.4 Pool-Required CCVs

- **INV-SRC-11**: If a pool does not implement `IPoolV2`, it falls back to `defaultCCVs`.
- **INV-SRC-12**: If a pool implements `IPoolV2` but returns an empty array from `getRequiredCCVs`, it falls back to `defaultCCVs`.
- **INV-SRC-13**: If a pool returns the zero-value address in its CCV list, those entries are replaced with `defaultCCVs`.
- **INV-SRC-14**: Pool-required CCVs that already appear in the merged list (user + lane-mandated) are skipped (deduplication).
- **INV-SRC-15**: If a user includes a pool-required CCV in their own CCV list with custom `ccvArgs`, those args are preserved. If the user does not include it, the pool-required CCV is added with empty `ccvArgs`.

### 2.5 Merge Order

- **INV-SRC-16**: The final CCV list is built in order: user/default first, lane-mandated second, pool-required last. This order is deterministic but not a protocol-level requirement.
- **INV-SRC-17**: The final merged CCV list contains no duplicates.

### 2.6 CCV Fees

- **INV-SRC-18**: Each CCV's `getFee` is called to compute per-CCV fees (USD cents, gas, bytes overhead).
- **INV-SRC-19**: Each CCV in the final merged list generates a receipt with `feeUSDCents`, `gasForVerification`, and `ccvPayloadSizeBytes`.
See FEE_INVARIANTS.md for overall fee structure, distribution, and conversion to fee token amounts.

### 2.7 CCV+Executor Hash

- **INV-SRC-20**: `ccvAndExecutorHash` is computed as `keccak256([addressLength][ccv1..ccvN][executor])`. The length prefix encodes the address byte length for the chain family, preventing collisions between different array sizes and cross-chain-family ambiguity.
- **INV-SRC-21**: `ccvAndExecutorHash` is embedded in the message for offchain validation only. It is NOT verified on the destination chain.

### 2.8 Token-Only Transfer Behavior

- **INV-SRC-22**: For token-only transfers (see MESSAGE_LIFECYCLE_INVARIANTS.md INV-TO-1) with no user-specified CCVs, only pool-related CCVs and lane-mandated CCVs are included. Default CCVs and receiver CCVs are excluded, since only the token issuer (pool) bears risk, not the receiver.

---

## 3. Destination Side (OffRamp) Invariants

### 3.1 Receiver-Required CCVs

- **INV-DST-1**: If the receiver implements `IAny2EVMMessageReceiverV2`, `getCCVsAndFinalityConfig` is called to get `requiredCCVs`, `optionalCCVs`, `optionalThreshold`, and `allowedFinalityConfig` (`bytes4`, see `FinalityCodec`).
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

### 3.3 Lane-Mandated CCVs (Inbound)

- **INV-DST-11**: Lane-mandated CCVs from `SourceChainConfig` are always included in the required set.

### 3.4 Default CCVs (Inbound)

- **INV-DST-12**: Default CCVs are included whenever any zero-value address appears in the combined required list (receiver + pool + lane-mandated).
- **INV-DST-13**: Defaults are added at most once, even if multiple zero-value address entries exist.

### 3.5 Token-Only Transfer Behavior (Inbound)

- **INV-DST-14**: For token-only transfers with tokens, receiver CCVs and defaults are skipped. Only pool-required CCVs and lane-mandated CCVs are used.
- **INV-DST-15**: For no-op tx (no token, no execution), defaults are used.

### 3.6 Deduplication and Overlap

- **INV-DST-16**: The final required CCV list is deduplicated: duplicates and zero-value address entries are removed.
- **INV-DST-17**: If a CCV appears in both the required and optional lists, it is removed from the optional list and the `optionalThreshold` is decremented (but not below zero). This prevents double-counting.

### 3.7 Quorum Enforcement

- **INV-DST-18**: ALL required CCVs must be present in the `ccvs` array passed to `execute`. If any is missing, the execution reverts with `RequiredCCVMissing`.
- **INV-DST-19**: At least `optionalThreshold` of the optional CCVs must be present in the `ccvs` array. If not enough are found, the execution reverts with `OptionalCCVQuorumNotReached`.
- **INV-DST-20**: The executor may pass extra CCVs beyond what is required/optional. These are ignored (not queried for verification).

### 3.8 Verification

- **INV-DST-21**: Each CCV that passes quorum checks has `verifyMessage` called on it. Verification succeeds by not reverting. Any revert from `verifyMessage` causes the entire message execution to fail.
- **INV-DST-22**: `ccvs.length` must equal `verifierResults.length`. Each CCV has exactly one corresponding verifier result.

### 3.9 Execution Outcome

See MESSAGE_LIFECYCLE_INVARIANTS.md (section 5.3) for execution outcome invariants (retry, terminal success, redundant state transitions).

---

## 4. CCV Address Stability

A CCV address is referenced by users, pools, lane configs, and receivers. Verifier implementations must be upgradeable without changing the address that all these parties reference. Each chain family must provide a mechanism to satisfy this requirement.

- **INV-ADDR-1**: A CCV address must remain stable across verifier upgrades. Users, pools, lane configs, and receivers all reference the CCV by this address — changing it would require coordinated reconfiguration across all parties.
- **INV-ADDR-2**: The CCV address must resolve to a concrete verifier implementation for each supported destination chain (outbound) and for each verifier result format (inbound). If resolution fails, the message cannot be sent or executed.
- **INV-ADDR-3**: Resolution must be transparent to the caller. The OnRamp, OffRamp, and fee quoting logic interact with the CCV address; the address is responsible for routing to the correct implementation.

### 4.1 EVM: Resolver Pattern

On EVM, address stability is achieved via resolvers that implement `ICrossChainVerifierResolver`:

- **INV-RES-1**: CCV addresses on EVM are resolver contract instances, not direct verifier implementations. Resolvers maintain a mapping from context (destination chain or version) to implementation.
- **INV-RES-2**: `VersionedVerifierResolver` maps outbound traffic by `destChainSelector` and inbound traffic by a 4-byte version prefix in `verifierResults`.
- **INV-RES-3**: If a resolver returns the zero-value address for outbound (`getOutboundImplementation`), the message cannot be sent.
- **INV-RES-4**: If a resolver returns the zero-value address for inbound (`getInboundImplementation`), the message cannot be executed.

---

## 5. Cross-Cutting Invariants

- **INV-CC-1**: Every message is verified by at least one CCV. This is guaranteed by INV-CFG-1/5 (at least one default or mandated CCV must exist) and the fallback mechanisms that ensure pools and receivers always resolve to at least the default CCVs. For token-only transfers, even when receiver CCVs and defaults are excluded (see INV-DST-14), pool-required CCVs or lane-mandated CCVs must provide at least one CCV. If no CCVs resolve for a token-only transfer, the message must revert — not proceed unverified.
- **INV-CC-2**: The source-side CCV list (embedded in `ccvAndExecutorHash`) is informational for offchain systems. The destination side independently computes its own CCV requirements.
- **INV-CC-3**: Source and destination CCV lists are independently configured and may differ. The source-side list determines fees; the destination-side list determines verification requirements.
- **INV-CC-4**: The same verifier may have different addresses on different chains. CCV addresses are chain-specific identifiers.
