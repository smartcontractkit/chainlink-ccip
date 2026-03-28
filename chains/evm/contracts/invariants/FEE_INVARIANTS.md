# Fee Invariants

## 1. Fee Structure

### 1.1 Components

A message fee is the sum of four components, computed in this order:

| Component | Source | Per-message count |
|-----------|--------|-------------------|
| CCV fees | Each CCV's `getFee` | One per CCV in the merged list |
| Pool fee | Pool's `getFee` (V2) or FeeQuoter fallback | 0 or 1 (only if token transfer) |
| Executor fee | Executor's `getFee` + execution cost | 1 |
| Network fee | Lane configuration | 1 |

- **INV-FEE-1**: The total message fee is the sum of all CCV fees, the pool fee (if applicable), the executor fee, and the network fee.
- **INV-FEE-2**: All fee components are initially denominated in USD cents and converted to fee token amounts using `feeTokenPrice`.

### 1.2 Fee Cap

- **INV-FEE-3**: The total fee in USD cents must not exceed a configured maximum (`maxUSDCentsPerMessage`). Exceeding the cap reverts the send.

### 1.3 Consistency

- **INV-FEE-4**: `getFee` (quoting) and the send path must compute the same fee using the same logic. CCV, pool, and executor fee functions must be deterministic between the two calls.

---

## 2. CCV Fees

See CCV_INVARIANTS.md (section 2.6) for CCV fee invariants.

---

## 3. Pool Fee

- **INV-FEE-5**: If the pool implements V2 and its fee config is enabled (`isEnabled == true`), the OnRamp uses the pool's `getFee` return values (`feeUSDCents`, `destGasOverhead`, `destBytesOverhead`, `tokenFeeBps`).
- **INV-FEE-6**: If the pool is V1 (legacy, relevant for v1 CCIP chains migrating to v2), or its fee config is not enabled (`isEnabled == false`), the OnRamp falls back to the FeeQuoter for token transfer fee parameters.
- **INV-FEE-7**: The FeeQuoter uses token-specific config if available, otherwise chain-level defaults.

---

## 4. Executor Fee

- **INV-FEE-8**: The executor fee has two parts: a flat fee (in USD cents, returned by the executor's fee function) and an execution cost (gas-based, computed separately).
- **INV-FEE-9**: When the no-execution address is used as the executor, the flat fee is zero and execution cost is not added. See MESSAGE_LIFECYCLE_INVARIANTS.md (section 9) for the no-execution address.
- **INV-FEE-10**: Execution cost is computed from gas estimates: base execution gas + CCV verification gas + pool gas overhead + user-specified gas limit, multiplied by `usdPerUnitGas`.

---

## 5. Network Fee

- **INV-FEE-11**: The network fee depends on whether the message includes a token transfer: `tokenNetworkFeeUSDCents` for messages with tokens, `messageNetworkFeeUSDCents` for messages without.
- **INV-FEE-12**: Exactly one of the two network fee values is used per message.

---

## 6. LINK Discount

- **INV-FEE-13**: When LINK is used as the fee token, a `linkFeeMultiplierPercent` discount applies to the USD-cent components of the fee (e.g. 90 = 10% discount). For non-LINK fee tokens, the multiplier is 100 (no discount).
- **INV-FEE-14**: The LINK discount does NOT apply to execution cost. Execution cost is converted from USD cents to fee token without the multiplier.

---

## 7. Price Requirements

- **INV-FEE-15**: Fee token price must be set and non-zero. A missing or zero price reverts the fee computation.
- **INV-FEE-16**: `usdPerUnitGas` for the destination chain must be set. A missing gas price reverts the fee computation.
- **INV-FEE-17**: Prices remain valid until overwritten. There is no staleness check.

---

## 8. Fee Distribution

- **INV-FEE-18**: CCV fees are transferred to their respective CCV addresses as defined in the receipts.
- **INV-FEE-19**: The executor fee is transferred to the executor address.
- **INV-FEE-20**: The pool fee is transferred to the pool only if it implements V2. For V1 pools (legacy), the fee stays on the OnRamp.
- **INV-FEE-21**: The network fee stays on the OnRamp.
