# Encoding Invariants

## 1. MessageV1 Encoding

### 1.1 General

- **INV-MSG-1**: The message format is fully chain agnostic. All chains produce and consume the same wire format. `keccak256(encodedMessageV1)` produces the `messageId` on all chains.
- **INV-MSG-2**: The encoding version byte is `1`. Decoding rejects any other version.
- **INV-MSG-3**: Decoding is strict: the entire encoded byte array must be consumed exactly. Any trailing bytes cause a revert.
- **INV-MSG-4**: A message may contain 0 or 1 token transfers. `MAX_NUMBER_OF_TOKENS = 1`.

### 1.2 Static Fields

A fixed 69-byte big-endian header starts every encoded message:

| Offset | Size | Field |
|--------|------|-------|
| 0 | 1 | `version` (always `1`) |
| 1 | 8 | `sourceChainSelector` (uint64) |
| 9 | 8 | `destChainSelector` (uint64) |
| 17 | 8 | `messageNumber` (uint64) |
| 25 | 4 | `executionGasLimit` (uint32) |
| 29 | 4 | `ccipReceiveGasLimit` (uint32) |
| 33 | 4 | `finality` (bytes4, see FINALITY_INVARIANTS.md) |
| 37 | 32 | `ccvAndExecutorHash` (bytes32) |

The header is followed by length prefixes for each variable-length section (4 uint8 for the address fields, 3 uint16 for `destBlob`, `tokenTransfer` and `data`). Together these account for an additional 10 bytes when every payload is empty.

- **INV-MSG-5**: The minimum encoded message size is **79 bytes** (69-byte header + 10 bytes of length prefixes). Decoding rejects any shorter input.

### 1.3 Variable-Length Fields

After the static header, variable-length fields are encoded in order. Each uses a length prefix followed by content.

| Length Prefix | Field | Encoding |
|---------------|-------|----------|
| uint8 | `onRampAddress` | Source-side, padded (abi.encode for EVM = 32 bytes) |
| uint8 | `offRampAddress` | Dest-side, unpadded (native byte length for dest chain) |
| uint8 | `senderAddress` | Source-side, padded (abi.encode for EVM = 32 bytes) |
| uint8 | `receiverAddress` | Dest-side, unpadded (native byte length for dest chain) |
| uint16 | `destBlob` | Chain-family-specific execution data (e.g. Solana accounts) |
| uint16 | `tokenTransfer` | Encoded `TokenTransferV1` (0 bytes if no token transfer) |
| uint16 | `data` | Arbitrary user payload |

- **INV-MSG-6**: All variable-length fields use a length prefix followed by content: `uint8` for addresses, `uint16` for data blobs. This applies uniformly across MessageV1, TokenTransferV1, and ExtraArgsV3.
- **INV-MSG-7**: Addresses use the minimal native byte length for their chain family. Source-side addresses on must be padded for legacy reasons (e.g. 20-byte EVM addresses encoded as 32 bytes). Destination-side addresses are always unpadded.
- **INV-MSG-8**: Destination-side addresses must match the configured native byte length for the destination chain. Padded inputs are rejected unless the padding bytes are all zero, in which case they are trimmed to the native length.

### 1.4 TokenTransferV1 Encoding

Token transfers are encoded with their own version byte and field structure:

| Length Prefix | Field | Encoding |
|---------------|-------|----------|
| — | `version` (1 byte, always `1`) | |
| — | `amount` (32 bytes, uint256) | |
| uint8 | `sourcePoolAddress` | Source-side address |
| uint8 | `sourceTokenAddress` | Source-side address |
| uint8 | `destTokenAddress` | Destination-side address |
| uint8 | `tokenReceiver` | Destination-side address |
| uint16 | `extraData` | Pool-specific data |

- **INV-MSG-9**: Token transfer encoding follows the same address conventions as the message (INV-MSG-6, INV-MSG-7).
- **INV-MSG-10**: `extraData` in the token transfer carries pool-specific data (e.g. `destPoolData` from `lockOrBurn`).
- **INV-MSG-11**: The `tokenTransfer` field in the message is length-prefixed with `uint16`. The encoded token transfer bytes must exactly consume the declared length; a mismatch causes a revert.

---

## 2. ExtraArgs Encoding

### 2.1 GenericExtraArgsV3

- **INV-ENC-1**: `ExtraArgsV3` tag is `0xa69dd4aa`. Any extraArgs with this 4-byte prefix is decoded as V3.
- **INV-ENC-2**: The number of CCVs is encoded as a single `uint8`, limiting a message to 255 user-specified CCVs.
- **INV-ENC-3**: `ccvs` and `ccvArgs` must have the same length. Encoding rejects a mismatch.

A 13-byte leading prefix (`tag`, `gasLimit`, `requestedFinalityConfig`, `ccvsLength`) is followed by `ccvsLength` repetitions of:

|------------------------------|--|
| `uint8 ccvAddressLength` + `bytes ccvAddress` | CCV address |
| `uint16 ccvArgsLength` + `bytes ccvArgs` | CCV-specific args |

Then:

| Length Prefix | Field |
|---------------|-------|
| uint8 | `executor` address |
| uint16 | `executorArgs` |
| uint8 | `tokenReceiver` |
| uint16 | `tokenArgs` |

- **INV-ENC-4**: The minimum encoded ExtraArgsV3 size is **19 bytes** (the 13-byte prefix plus the five trailing length fields when no CCVs and no payloads are present). Decoding rejects any shorter input.
- **INV-ENC-5**: All addresses in ExtraArgsV3 use the same `uint8 addressLength + bytes` encoding. `addressLength == 0` is the zero-value address sentinel, which means: "use the default" for CCV and executor entries, and "use the message receiver" for `tokenReceiver`.
- **INV-ENC-6**: Decoding is strict: all bytes must be consumed exactly. Trailing bytes cause a revert.

### 2.2 Executor Args

- **INV-ENC-7**: SVM executor args use tag `0x1a2b3c4d` with base size 14 bytes: `tag(4) + useATA(1) + accountIsWritableBitmap(8) + accountsLength(1)`. Accounts are 32 bytes each.
- **INV-ENC-8**: Sui executor args use tag `0x5e6f7a8b` with base size 5 bytes: `tag(4) + objectIdsLength(1)`. Object IDs are 32 bytes each.
- **INV-ENC-9**: Both executor arg formats enforce strict decoding: all bytes must be consumed exactly.

---

## 3. Numeric Range Safety

- **INV-ENC-10**: All numeric fields in the wire format must survive a round-trip through the implementation's native type system without overflow, truncation, or precision loss. Specifically:
  - `amount` (uint256) must be representable in the implementation's token amount type at its full range. If the native type has a smaller range than uint256, the implementation must validate the value fits before processing.
  - `messageNumber` (uint64) must be representable without overflow.
  - `executionGasLimit` and `ccipReceiveGasLimit` (uint32) must be representable without overflow.
  - `finality` is a 32-bit packed bit pattern (see FINALITY_INVARIANTS.md) and must round-trip without altering bits.
- **INV-ENC-11**: Encoding functions must reject values that exceed the wire format's range. `encodeUint256` must reject values > 2^256 - 1. `encodeUint64` must reject values > 2^64 - 1. Encoding must never silently produce output longer or shorter than the specified field width.
- **INV-ENC-12**: Intermediate type conversions during encoding/decoding must not narrow the value range. If a conversion path routes a wide type through a narrower intermediate type (e.g., 256-bit → 64-bit → decimal), the narrowing must either be validated or eliminated.

---

## 4. Character and Byte-String Encoding

- **INV-ENC-13**: All text-to-bytes encoding functions must be injective (distinct inputs produce distinct outputs). Unsupported characters must cause an error, not a silent substitution. Hash inputs must be canonical — two semantically distinct values must never produce the same hash.
- **INV-ENC-14**: All byte-string inputs at system boundaries must be validated for well-formedness: even length, valid hex characters, and consistent case normalization. Malformed byte strings must be rejected at the point of entry, not silently accepted.

---

## 5. Token Pool Remote Configuration

- **INV-POOL-ENC-1**: `remotePoolAddresses` and `remoteTokenAddress` are stored as raw bytes. For remote EVM chains, these are abi-encoded addresses (32 bytes). For other chain families, they use the native address representation.
- **INV-POOL-ENC-2**: `remoteTokenAddress` must be non-empty (zero-length reverts).
- **INV-POOL-ENC-3**: Multiple remote pool addresses can be configured per remote chain. This supports pool upgrades on the remote chain — the old pool must remain configured until all inflight messages from it are processed.
- **INV-POOL-ENC-4**: Remote pool addresses must be non-empty. The encoding format of the address must match what the remote chain's OnRamp produces in `TokenTransferV1.sourcePoolAddress`.
