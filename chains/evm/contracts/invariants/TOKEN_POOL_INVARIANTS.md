# Token Pool Invariants

## 1. Pool Types

- **INV-POOL-1**: All pool types implement the same protocol-level interface. The pool mechanism (lock/release, burn/mint, or variants) is an implementation detail transparent to the OnRamp and OffRamp.
- **INV-POOL-2**: Lock-release pools move tokens into a lockbox on the source chain and withdraw from a lockbox on the destination chain.
- **INV-POOL-3**: Burn-mint pools destroy tokens on the source chain and create tokens on the destination chain.

---

## 2. Access Control

- **INV-POOL-4**: Only the OnRamp (resolved via the router for the remote chain) may call `lockOrBurn`. The pool verifies that the caller is the OnRamp for the given remote chain.
- **INV-POOL-5**: Only the OffRamp may call `releaseOrMint`. The pool verifies that the caller is a registered OffRamp for the given remote chain.
- **INV-POOL-6**: The remote chain must be configured on the pool before any `lockOrBurn` or `releaseOrMint` call. Unconfigured chains are rejected.

---

## 3. V1 vs V2 Interface

- **INV-POOL-7**: V1 pools support `lockOrBurn` and `releaseOrMint` without finality or fee parameters. The full token amount is locked/burned and released/minted. V1 is legacy, relevant for v1 CCIP chains migrating to v2.
- **INV-POOL-8**: V2 pools extend V1 with: `getFee`, `getTokenTransferFeeConfig`, `getRequiredCCVs`, and finality-aware `lockOrBurn`/`releaseOrMint` (receiving `blockConfirmationsRequested` and `tokenArgs`).
- **INV-POOL-9**: V1 pools cannot support FTF or custom token args. See FINALITY_INVARIANTS.md (section 2.3) for details on how V1 pools interact with finality.

---

## 4. Fee Deduction (V2)

- **INV-POOL-10**: V2 pools deduct a proportional fee from the token amount before lock/burn: `destTokenAmount = amount - (amount * feeBps / 10000)`. The fee stays in the pool.
- **INV-POOL-11**: The fee rate (`feeBps`) depends on finality: `defaultBlockConfirmationsTransferFeeBps` for finalized transfers, `customBlockConfirmationsTransferFeeBps` for FTF. See FINALITY_INVARIANTS.md for details.
- **INV-POOL-12**: All `feeBps` values (default and custom block confirmations) must be less than 10,000 (100%). Validated at config time.
- **INV-POOL-13**: When fee config is not enabled for a remote chain, `feeBps` defaults to `0` and no fee is deducted.

See FEE_INVARIANTS.md (section 3) for how pool fee quoting integrates with the OnRamp and FeeQuoter fallback.

---

## 5. Decimal Handling

- **INV-POOL-14**: Pools encode their local token decimals into `sourcePoolData` (as part of the `lockOrBurn` return). On the destination, this is used to convert the source-denominated amount to the local token's decimals.
- **INV-POOL-15**: If `sourcePoolData` is empty, the destination pool assumes the source token has the same decimals as the local token.
- **INV-POOL-16**: If `sourcePoolData` is present, it must be exactly 32 bytes encoding a `uint8` decimal value.
- **INV-POOL-17**: When converting from higher to lower decimals, the amount is divided and rounds down. Dust is lost.
- **INV-POOL-18**: When converting from lower to higher decimals, the amount is multiplied. If the multiplication overflows or the decimal difference exceeds 77, the transfer reverts.

---

## 6. destBytesOverhead

- **INV-POOL-19**: `destBytesOverhead` is the data availability byte budget attributed to the token transfer for fee calculation.
- **INV-POOL-20**: The minimum `destBytesOverhead` is 32 bytes (the size of a standard `sourcePoolData` encoding).
- **INV-POOL-21**: If `destPoolData` returned by `lockOrBurn` exceeds `destBytesOverhead`, the OnRamp reverts. Pools returning larger data must have a correspondingly larger `destBytesOverhead` configured.

---

## 7. Rate Limiting

### 7.1 Token Bucket Mechanism

- **INV-RL-1**: Rate limiting uses a token bucket: each bucket has `capacity` (maximum tokens), `rate` (tokens refilled per second), and `tokens` (current available tokens).
- **INV-RL-2**: On each consumption, the bucket first refills based on elapsed time: `tokens = min(capacity, tokens + elapsedSeconds * rate)`.
- **INV-RL-3**: If the requested amount exceeds `capacity`, the transfer always reverts regardless of available tokens.
- **INV-RL-4**: If the requested amount exceeds available `tokens`, the transfer reverts. If `rate > 0`, a minimum wait time is calculable. If `rate == 0`, the wait is effectively infinite.
- **INV-RL-5**: If the bucket is disabled or the requested amount is zero, no consumption occurs.

### 7.2 Inbound vs Outbound

- **INV-RL-6**: Pools maintain separate rate limit buckets for outbound (`lockOrBurn`) and inbound (`releaseOrMint`) per remote chain.
- **INV-RL-7**: FTF transfers may use additional separate rate limit buckets. See FINALITY_INVARIANTS.md (section 4.4) for details.

### 7.3 Configuration

- **INV-RL-8**: Rate limiter configuration is per-remote-chain. Each remote chain can have its own outbound and inbound bucket parameters.
- **INV-RL-9**: When a bucket is enabled, `rate` must not exceed `capacity`.
- **INV-RL-10**: When a bucket is disabled, both `rate` and `capacity` must be zero.
- **INV-RL-11**: Rate limit configuration can be updated by the contract owner or a designated rate limit admin.

---

## 8. Pool Configuration

- **INV-PCFG-1**: Remote chains are added and removed via chain update operations. Adding a chain requires a non-empty `remoteTokenAddress`. Removing a chain clears all associated configuration including rate limiters.
- **INV-PCFG-2**: Remote pool addresses are added and removed individually. See ENCODING_INVARIANTS.md (section 3) for encoding format and validation requirements.
- **INV-PCFG-3**: Pool configuration updates (chain adds/removes, remote pool adds/removes) are restricted to the contract owner.

---

## 9. Source Pool Validation (Inbound)

- **INV-PVAL-1**: On `releaseOrMint`, the pool validates that `sourcePoolAddress` (from the inbound message) matches one of the configured remote pool addresses for the source chain. Unrecognized source pools are rejected.
