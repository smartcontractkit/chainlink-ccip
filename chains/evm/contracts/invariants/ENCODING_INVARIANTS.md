# Encoding Invariants

## 1. MessageV1 Encoding

### 1.1 General

- **INV-MSG-1**: The message format is fully chain agnostic. All chains produce and consume the same wire format. `keccak256(encodedMessageV1)` produces the `messageId` on all chains.
- **INV-MSG-2**: The encoding version byte is `1`. Decoding rejects any other version.
- **INV-MSG-3**: Decoding is strict: the entire encoded byte array must be consumed exactly. Any trailing bytes cause a revert.
- **INV-MSG-4**: A message may contain 0 or 1 token transfers. `MAX_NUMBER_OF_TOKENS = 1`.

### 1.2 Static Fields

The static header is 77 bytes, encoded in big-endian order:

| Offset | Size | Field |
|--------|------|-------|
| 0 | 1 | `version` (always `1`) |
| 1 | 8 | `sourceChainSelector` (uint64) |
| 9 | 8 | `destChainSelector` (uint64) |
| 17 | 8 | `messageNumber` (uint64) |
| 25 | 4 | `executionGasLimit` (uint32) |
| 29 | 4 | `ccipReceiveGasLimit` (uint32) |
| 33 | 2 | `finality` (uint16) |
| 35 | 32 | `ccvAndExecutorHash` (bytes32) |

- **INV-MSG-5**: `MESSAGE_V1_BASE_SIZE = 77`. Decoding rejects any input shorter than 77 bytes.

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
- **INV-MSG-7**: All addresses use the minimal native byte length for their chain family (e.g. 20 bytes for EVM, 32 bytes for Solana). The only exception is source-side EVM addresses (`onRampAddress`, `sender`, `sourcePoolAddress`, `sourceTokenAddress`), which are abi-encoded (32 bytes, left-padded) for legacy reasons.
- **INV-MSG-8**: Destination-side addresses on the OnRamp are validated against the configured `addressBytesLength` for the destination chain. If a 32-byte abi-encoded address is provided, it is trimmed to `addressBytesLength` after verifying no non-zero padding exists outside the expected range.

### 1.4 TokenTransferV1 Encoding

Token transfers are encoded with their own version byte and field structure:

| Length Prefix | Field | Encoding |
|---------------|-------|----------|
| — | `version` (1 byte, always `1`) | |
| — | `amount` (32 bytes, uint256) | |
| uint8 | `sourcePoolAddress` | Source-side, padded |
| uint8 | `sourceTokenAddress` | Source-side, padded |
| uint8 | `destTokenAddress` | Dest-side, unpadded |
| uint8 | `tokenReceiver` | Dest-side, unpadded |
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

Static fields (17 bytes total):

| Offset | Size | Field |
|--------|------|-------|
| 0 | 4 | `tag` (`0xa69dd4aa`) |
| 4 | 4 | `gasLimit` (uint32) |
| 8 | 2 | `blockConfirmations` (uint16) |
| 10 | 1 | `ccvsLength` (uint8) |

Followed by variable-length fields:

| Repeated `ccvsLength` times | |
|------------------------------|--|
| `uint8 ccvAddressLength` + `bytes ccvAddress` | CCV address |
| `uint16 ccvArgsLength` + `bytes ccvArgs` | CCV-specific args |

Then:

| Length Prefix | Field |
|---------------|-------|
| uint8 | `executor` address |
| uint16 | `executorArgs` |
| uint8 | `tokenReceiver` address |
| uint16 | `tokenArgs` |

- **INV-ENC-4**: `GENERIC_EXTRA_ARGS_V3_BASE_SIZE = 17`. Decoding rejects any input shorter than 17 bytes.
- **INV-ENC-5**: All addresses in ExtraArgsV3 (CCVs, executor, tokenReceiver) use the same `uint8 addressLength + bytes` encoding. `addressLength == 0` encodes the zero-value address (default placeholder for CCVs, default executor, or "use message receiver" for tokenReceiver).
- **INV-ENC-6**: Decoding is strict: all bytes must be consumed exactly. Trailing bytes cause a revert.

### 2.2 Executor Args

- **INV-ENC-7**: SVM executor args use tag `0x1a2b3c4d` with base size 14 bytes: `tag(4) + useATA(1) + accountIsWritableBitmap(8) + accountsLength(1)`. Accounts are 32 bytes each.
- **INV-ENC-8**: Sui executor args use tag `0x5e6f7a8b` with base size 5 bytes: `tag(4) + objectIdsLength(1)`. Object IDs are 32 bytes each.
- **INV-ENC-9**: Both executor arg formats enforce strict decoding: all bytes must be consumed exactly.

---

## 3. Token Pool Remote Configuration

- **INV-POOL-ENC-1**: `remotePoolAddresses` and `remoteTokenAddress` are stored as raw bytes. For remote EVM chains, these are abi-encoded addresses (32 bytes). For other chain families, they use the native address representation.
- **INV-POOL-ENC-2**: `remoteTokenAddress` must be non-empty (zero-length reverts).
- **INV-POOL-ENC-3**: Multiple remote pool addresses can be configured per remote chain. This supports pool upgrades on the remote chain — the old pool must remain configured until all inflight messages from it are processed.
- **INV-POOL-ENC-4**: Remote pool addresses must be non-empty. The encoding format of the address must match what the remote chain's OnRamp produces in `TokenTransferV1.sourcePoolAddress`.
