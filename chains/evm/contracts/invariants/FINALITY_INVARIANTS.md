# Finality / Faster-Than-Finality (FTF) Invariants

## 1. Encoding

Finality is a single **`bytes4`** value carried on the wire as `MessageV1.finality` and selected by the sender as `requestedFinalityConfig` in ExtraArgsV3. Its layout is defined by the finality codec:

- The lower 16 bits encode an optional **block depth** (`1..65535`).
- The upper 16 bits encode **flags**, of which one is currently defined: `WAIT_FOR_SAFE_FLAG`. The remaining flag bits are reserved for future use and accepted on the wire.
- The all-zero value is `WAIT_FOR_FINALITY_FLAG` and means "wait for full finality".

### 1.1 Requested vs allowed finality

- **INV-FIN-ENC-1**: A *requested* finality (sender-side) must encode **exactly one mode**: either `WAIT_FOR_FINALITY_FLAG`, a pure block depth with no flag bits, or exactly one flag with no block depth. A flag combined with a non-zero depth is invalid.
- **INV-FIN-ENC-2**: An *allowed* finality (configured on participants such as pools, CCVs, executor and receiver) may set multiple bits, expressing the union of accepted requested modes.
- **INV-FIN-ENC-3**: A requested finality is permitted by an allowed finality iff one of the following holds:
  - it equals `WAIT_FOR_FINALITY_FLAG` (always allowed), or
  - it sets a flag bit that is also set in the allowed finality, or
  - it is depth-based and the allowed finality has a non-zero depth no greater than the requested depth.
- **INV-FIN-ENC-4**: Legacy ExtraArgs formats that do not carry a finality field default to `WAIT_FOR_FINALITY_FLAG`.

---

## 2. Source Side (OnRamp)

- **INV-FIN-SRC-1**: The OnRamp validates only the **wire shape** of the requested finality (one mode, see INV-FIN-ENC-1). Whether a given request is *permitted* is delegated to downstream participants (CCVs, pool, executor) and ultimately to the receiver on the destination side.
- **INV-FIN-SRC-2**: The requested finality is forwarded unchanged to every CCV, the pool, and the executor when computing fees and required CCVs, and is copied verbatim into `MessageV1.finality` so it is committed in the message ID.
- **INV-FIN-SRC-3**: A non-finality requested mode must be supported by every participant in the path; otherwise the send reverts. Participants without finality awareness (legacy pools, legacy receivers) implicitly require `WAIT_FOR_FINALITY_FLAG`.
- **INV-FIN-SRC-4**: Legacy pools cannot accept FTF requests or pool-specific token args. Sending such a message via a legacy pool reverts.

---

## 3. Destination Side (OffRamp)

- **INV-FIN-DST-1**: On execution the OffRamp checks `message.finality` against the receiver's allowed finality using the rule in INV-FIN-ENC-3.
- **INV-FIN-DST-2**: A receiver that does not declare an allowed finality is treated as allowing only `WAIT_FOR_FINALITY_FLAG`. This protects receivers that have not explicitly opted in to FTF.
- **INV-FIN-DST-3**: `message.finality` is forwarded to the pool on inbound for both CCV resolution and release/mint, so pool-side finality policy and fees apply on the destination as well.
- **INV-FIN-DST-4**: Legacy pools receive no finality parameter on inbound; they implicitly behave as if the request were `WAIT_FOR_FINALITY_FLAG`.

### 3.1 Token-Only Transfers

- **INV-FIN-DST-5**: For token-only transfers (no receiver callback), the receiver's allowed finality is not consulted; `message.finality` is still supplied to the pool for validation, fees, and rate limiting, with lane-mandated CCVs unchanged.

---

## 4. Pool Finality

### 4.1 Configuration

- **INV-FIN-POOL-1**: A pool exposes a configured **allowed finality** (`bytes4`) that bounds which requested modes it accepts. The default `WAIT_FOR_FINALITY_FLAG` means the pool does not support FTF.

### 4.2 Outbound and Inbound Validation

- **INV-FIN-POOL-2**: On both outbound and inbound, the pool validates the requested finality against its allowed finality using the rule in INV-FIN-ENC-3 and reverts otherwise.
- **INV-FIN-POOL-3**: Legacy `lockOrBurn` / `releaseOrMint` paths that take no finality parameter are treated as `WAIT_FOR_FINALITY_FLAG`.

### 4.3 Rate Limiting

- **INV-FIN-POOL-4**: Pools maintain separate inbound and outbound rate-limit buckets per remote chain for default-finality and fast-finality (FTF) transfers.
- **INV-FIN-POOL-5**: A request with `WAIT_FOR_FINALITY_FLAG` consumes from the default bucket; any other requested mode consumes from the FTF bucket when configured. If the FTF bucket is not configured the pool falls back to the default bucket.

### 4.4 Fees

- **INV-FIN-POOL-6**: Token transfer fee configuration carries separate fee fields for default-finality and FTF transfers (both flat USD-denominated fees and proportional bps fees deducted from the transferred amount).

### 4.5 Pool-required CCVs

- **INV-FIN-POOL-7**: The pool's required-CCV resolution receives the requested finality, so a pool may require additional or different CCVs for FTF transfers.

---

## 5. Executor Finality

- **INV-FIN-EXEC-1**: The executor exposes an **allowed finality** (`bytes4`) and validates the requested finality against it using the rule in INV-FIN-ENC-3.
- **INV-FIN-EXEC-2**: `WAIT_FOR_FINALITY_FLAG` is always accepted; other modes are accepted only when the executor's allowed finality permits them.

---

## 6. CCV Finality

- **INV-FIN-CCV-1**: A CCV exposes an **allowed finality** (`bytes4`) and validates each requested finality against it using the rule in INV-FIN-ENC-3.
- **INV-FIN-CCV-2**: CCVs may price requested modes differently (e.g. charge more for FTF) or reject a mode by reverting at fee time.
- **INV-FIN-CCV-3**: Waiting for the requested finality is an offchain concern of the CCV node(s). The onchain CCV contract validates encoding, configuration and proofs, not chain-head progression.

---

## 7. Re-org and Safety

- **INV-FIN-REORG-1**: `messageNumber` is unique per lane only for finalized messages. Under FTF a re-org may cause a send to be re-emitted with a different `messageNumber`.
- **INV-FIN-REORG-2**: FTF shifts re-org risk to the receiver, the pool and any downstream integrators. The protocol makes this explicit by requiring opt-in at every layer.

---

## 8. Opt-in Requirements (Defense in Depth)

FTF requires explicit opt-in at every layer of the stack. All layers default to allowing only `WAIT_FOR_FINALITY_FLAG`.

| Layer | Opt-in | Default |
|-------|--------|---------|
| Sender | Sets `requestedFinalityConfig` in ExtraArgsV3 to a non-finality mode | `WAIT_FOR_FINALITY_FLAG` |
| Receiver | Declares an allowed finality permitting the request | `WAIT_FOR_FINALITY_FLAG` |
| Pool | Configures an allowed finality permitting the request | `WAIT_FOR_FINALITY_FLAG` |
| Executor | Configures an allowed finality permitting the request | `WAIT_FOR_FINALITY_FLAG` |
| CCV | Configures an allowed finality permitting the request | `WAIT_FOR_FINALITY_FLAG` |

- **INV-FIN-OPTIN-1**: If any single layer rejects the requested mode, the message cannot be sent or executed in that mode.
- **INV-FIN-OPTIN-2**: Legacy pools and legacy receivers implicitly reject FTF without any code changes; backward compatibility is preserved.
