// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

/// @notice Library for encoding and decoding the source pool data for the USDC token pool based on the CCTP version.
/// @dev While every version of a CCTP-enabled pool has a different source pool data format, the encoding and decoding
/// schemes are similar in a few ways. Each source pool data format is prefixed with a version number, in bytes4 format.
/// It is then encodePacked() with each of the sourceTokenDataPayload fields. Decoding means parsing the bytes array
/// into the corresponding sourceTokenDataPayload struct by parsing each field individually and then assembling them
/// into the corresponding struct. This adds some additional gas overhead during decoding, but the benefits of saving
/// space on the source pool data outweigh the overhead.
/// @dev Any future versions of CCTP should include in this library a new function for encoding and decoding the source
/// pool data accordingly.
library USDCSourcePoolDataCodec {
  error InvalidVersion(bytes4 version);

  /// @dev The flag used to indicate that the source pool data is coming from a chain that does not have CCTP Support,
  /// and so the lock release pool should be used. The BurnMintWithLockReleaseTokenPool uses this flag as its source pool
  /// data to indicate that the tokens should be released from the lock release pool rather than attempting to be minted
  /// through CCTP.
  /// @dev The preimage is bytes4(keccak256("NO_CCTP_USE_LOCK_RELEASE")).
  bytes4 public constant LOCK_RELEASE_FLAG = 0xfa7c07de;

  /// @dev The preimage is bytes4(keccak256("CCTP_V1"))
  bytes4 public constant CCTP_VERSION_1_TAG = 0xf3567d18;

  /// @dev There are two different tags for CCTP V2 to allow for CCIP V1.7 Compatibility which will enable fast transfers.
  /// Both tags will route to the same CCTP V2 pool, but will allow for pools to identify the type of transfer (slow or fast).

  /// @dev The preimage is bytes4(keccak256("CCTP_V2"))
  bytes4 public constant CCTP_VERSION_2_TAG = 0xb148ea5f;

  /// @dev The preimage is bytes4(keccak256("CCTP_V2_CCV"))
  bytes4 public constant CCTP_VERSION_2_CCV_TAG = 0x3047587c;

  /// Note: Since this struct never exists in storage, only in memory after an ABI-decoding, proper struct-packing
  /// is not necessary and field ordering has been defined so as to best support off-chain code.
  /// @dev This struct has been titled for version 1 to indicate that it should be used for CCTP V1 messages.
  struct SourceTokenDataPayloadV1 {
    uint64 nonce; // Nonce of the message returned from the depositForBurnWithCaller() call to the CCTP contracts.
    uint32 sourceDomain; // Source domain of the message.
  }

  /// @dev This struct has been titled for version 2 to indicate that it should be used for CCTP V2 messages. Whether
  /// it is a slow or fast transfer is irrelevant as it will be routed to the same destination CCTP V2 pool regardless.
  struct SourceTokenDataPayloadV2 {
    uint32 sourceDomain;
    bytes32 depositHash;
  }

  /// @notice Encodes the source token data payload into a bytes array.
  /// @dev By using abi.encodePacked(), significant amount of space on the source pool data is saved.
  /// since abi.encode pads every field to the nearest 32 bytes. While it adds some overhead during decoding, the
  /// benefits of saving space on the source pool data outweigh the overhead.
  /// @param sourceTokenDataPayload The source token data payload to encode.
  /// @return The encoded source token data payload.
  function _encodeSourceTokenDataPayloadV1(
    SourceTokenDataPayloadV1 memory sourceTokenDataPayload
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(CCTP_VERSION_1_TAG, sourceTokenDataPayload.nonce, sourceTokenDataPayload.sourceDomain);
  }

  /// @notice Encodes the source token data payload into a bytes array using the CCTP V2 tag.
  /// @param sourceTokenDataPayload The source token data payload to encode.
  /// @return The encoded source token data payload.
  function _encodeSourceTokenDataPayloadV2(
    SourceTokenDataPayloadV2 memory sourceTokenDataPayload
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(CCTP_VERSION_2_TAG, sourceTokenDataPayload.sourceDomain, sourceTokenDataPayload.depositHash);
  }

  /// @notice Encodes the source token data payload into a bytes array using the CCTP V2 CCV tag.
  /// @param sourceTokenDataPayload The source token data payload to encode.
  /// @return The encoded source token data payload.
  function _encodeSourceTokenDataPayloadV2CCV(
    SourceTokenDataPayloadV2 memory sourceTokenDataPayload
  ) internal pure returns (bytes memory) {
    return
      abi.encodePacked(CCTP_VERSION_2_CCV_TAG, sourceTokenDataPayload.sourceDomain, sourceTokenDataPayload.depositHash);
  }

  /// @notice Decodes the abi.encodePacked() source pool data into its corresponding SourceTokenDataPayload struct.
  /// @param sourcePoolData The source pool data to decode in raw bytes.
  /// @return sourceTokenDataPayload The decoded source token data payload.
  function _decodeSourceTokenDataPayloadV2(
    bytes memory sourcePoolData
  ) internal pure returns (SourceTokenDataPayloadV2 memory sourceTokenDataPayload) {
    bytes4 version;
    uint32 sourceDomain;
    bytes32 depositHash;

    assembly {
      // Load version (first 4 bytes of data, offset 32 to skip the length slot)
      version := mload(add(sourcePoolData, 32))
      // Load sourceDomain (next 4 bytes, offset 36 (32 + 4)) - shift right by 224 bits to get left-most 4 bytes
      // offset 36 = 32 (length slot) + 4 (uint32)
      // shift right by 224 = 256 - 32 (uint32)
      sourceDomain := shr(224, mload(add(sourcePoolData, 36)))
      // Load depositHash (next 32 bytes) - Since depositHash is a bytes32, no shifting is needed.
      // offset 40 = 32 (length slot) + 4 (bytes4) + 4 (uint32)
      depositHash := mload(add(sourcePoolData, 40))
    }

    if (version != CCTP_VERSION_2_TAG && version != CCTP_VERSION_2_CCV_TAG) revert InvalidVersion(version);

    sourceTokenDataPayload.sourceDomain = sourceDomain;
    sourceTokenDataPayload.depositHash = depositHash;

    return sourceTokenDataPayload;
  }

  /// @notice Decodes the abi.encodePacked() source pool data into its corresponding SourceTokenDataPayload struct.
  /// @param sourcePoolData The source pool data to decode in raw bytes.
  /// @return sourceTokenDataPayload The decoded source token data payload.
  function _decodeSourceTokenDataPayloadV1(
    bytes memory sourcePoolData
  ) internal pure returns (SourceTokenDataPayloadV1 memory sourceTokenDataPayload) {
    bytes4 version;
    uint64 nonce;
    uint32 sourceDomain;

    assembly {
      // Load version (first 4 bytes of data, offset 32 to skip the length slot)
      version := mload(add(sourcePoolData, 32))
      // Load nonce (next 8 bytes, offset 36) - shift right by 192 bits to get left-most 8 bytes
      // offset 36 = 32 (length slot) + 4 (uint64)
      // shift right by 192 = 256 - 64 (uint64)
      nonce := shr(192, mload(add(sourcePoolData, 36)))
      // Load sourceDomain (next 4 bytes, offset 44) - shift right by 224 bits to get leftmost 4 bytes
      // offset 44 = 32 (length slot) + 4 (bytes4) 8 (uint64)
      // shift right by 224 = 256 - 32 (uint32)
      sourceDomain := shr(224, mload(add(sourcePoolData, 44)))
    }

    if (version != CCTP_VERSION_1_TAG) revert InvalidVersion(version);

    sourceTokenDataPayload.nonce = nonce;
    sourceTokenDataPayload.sourceDomain = sourceDomain;

    return sourceTokenDataPayload;
  }

  /// @notice Calculates the deposit hash for the source pool data.
  /// @param sourceDomain The source domain of the message.
  /// @param amount The amount of the message.
  /// @param destinationDomain The destination domain of the message.
  /// @param mintRecipient The mint recipient of the message.
  /// @param burnToken The burn token of the message.
  /// @param destinationCaller The destination caller of the message.
  /// @param maxFee The max fee of the message.
  /// @param minFinalityThreshold The min finality threshold of the message.
  /// @return depositHash The deposit hash of the source pool data which will be matched off-chain to its CCTP attestation.
  function _calculateDepositHash(
    uint32 sourceDomain,
    uint256 amount,
    uint32 destinationDomain,
    bytes32 mintRecipient,
    bytes32 burnToken,
    bytes32 destinationCaller,
    uint256 maxFee,
    uint32 minFinalityThreshold
  ) internal pure returns (bytes32) {
    return keccak256(
      abi.encode(
        sourceDomain,
        amount,
        destinationDomain,
        mintRecipient,
        burnToken,
        destinationCaller,
        maxFee,
        minFinalityThreshold
      )
    );
  }
}
