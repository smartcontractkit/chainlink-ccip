// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPool} from "../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolCCTPV2} from "../pools/USDC/USDCTokenPoolCCTPV2.sol";

/// @notice Library for encoding and decoding the source pool data for the USDC token pool based on the CCTP version.
library USDCSourcePoolDataCodec {
  error InvalidVersion(uint32 version);

  /// @notice Encodes the source token data payload into a bytes array.
  /// @param version The version of the source token data payload.
  /// @param sourceTokenDataPayload The source token data payload to encode.
  /// @return The encoded source token data payload.
  function _encodeSourceTokenDataPayloadV0(
    bytes4 version,
    USDCTokenPool.SourceTokenDataPayloadV0 memory sourceTokenDataPayload
  ) internal pure returns (bytes memory) {
    // By using encodePacked rather than abi.encode, significant amount of space on the source pool data is saved.
    // since abi.encode pads every field to the nearest 32 bytes.
    return abi.encodePacked(version, sourceTokenDataPayload.nonce, sourceTokenDataPayload.sourceDomain);
  }

  /// @notice Encodes the source token data payload into a bytes array.
  /// @param version The version of the source token data payload.
  /// @param sourceTokenDataPayload The source token data payload to encode.
  /// @return The encoded source token data payload.
  function _encodeSourceTokenDataPayloadV1(
    bytes4 version,
    USDCTokenPoolCCTPV2.SourceTokenDataPayloadV1 memory sourceTokenDataPayload
  ) internal pure returns (bytes memory) {
    // By using encodePacked rather than abi.encode, significant amount of space on the source pool data is saved.
    // since abi.encode pads every field to the nearest 32 bytes.
    return abi.encodePacked(version, sourceTokenDataPayload.sourceDomain, sourceTokenDataPayload.depositHash);
  }

  /// @notice Decodes the source pool data into its corresponding SourceTokenDataPayload struct.
  /// @param sourcePoolData The source pool data to decode.
  /// @return sourceTokenDataPayload The decoded source token data payload.
  function _decodeSourceTokenDataPayloadV1(
    bytes memory sourcePoolData
  ) internal pure returns (USDCTokenPoolCCTPV2.SourceTokenDataPayloadV1 memory sourceTokenDataPayload) {
    uint256 offset = 0;

    // Since memory arrays cannot be sliced in the same way as calldata arrays, we need to create new bytes arrays
    // to store the individual fields and then parse into their corresponding types.
    // Version (uint32)(4 bytes)
    bytes memory versionBytes = new bytes(4);
    for (uint256 i = 0; i < 4; ++i) {
      versionBytes[i] = sourcePoolData[offset + i];
    }
    uint32 version = uint32(bytes4(versionBytes));
    if (version != 1) revert InvalidVersion(version);
    offset += 4;

    // Source Domain (uint32)(4 bytes)
    bytes memory domainBytes = new bytes(4);
    for (uint256 i = 0; i < 4; ++i) {
      domainBytes[i] = sourcePoolData[offset + i];
    }
    sourceTokenDataPayload.sourceDomain = uint32(bytes4(domainBytes));
    offset += 4;

    // Deposit Hash (bytes32)(32 bytes)
    bytes memory hashBytes = new bytes(32);
    for (uint256 i = 0; i < 32; i++) {
      hashBytes[i] = sourcePoolData[offset + i];
    }
    sourceTokenDataPayload.depositHash = bytes32(hashBytes);

    return sourceTokenDataPayload;
  }

  /// @notice Decodes the source pool data into its corresponding SourceTokenDataPayload struct.
  /// @param sourcePoolData The source pool data to decode.
  /// @return sourceTokenDataPayload The decoded source token data payload.
  function _decodeSourceTokenDataPayloadV0(
    bytes memory sourcePoolData
  ) internal pure returns (USDCTokenPool.SourceTokenDataPayloadV0 memory sourceTokenDataPayload) {
    uint256 offset = 0;

    // Since memory arrays cannot be sliced in the same way as calldata arrays, we need to create new bytes arrays
    // to store the individual fields and then parse into their corresponding types.
    // Version (uint32)(4 bytes)
    bytes memory versionBytes = new bytes(4);
    for (uint256 i = 0; i < 4; ++i) {
      versionBytes[i] = sourcePoolData[offset + i];
    }
    uint32 version = uint32(bytes4(versionBytes));

    if (version != 0) revert InvalidVersion(version);
    offset += 4;

    // Nonce (uint64)(8 bytes)
    bytes memory nonceBytes = new bytes(8);
    for (uint256 i = 0; i < 8; ++i) {
      nonceBytes[i] = sourcePoolData[offset + i];
    }
    sourceTokenDataPayload.nonce = uint64(bytes8(nonceBytes));
    offset += 8;

    // Source Domain (uint32)(4 bytes)
    bytes memory domainBytes = new bytes(4);
    for (uint256 i = 0; i < 4; ++i) {
      domainBytes[i] = sourcePoolData[offset + i];
    }
    sourceTokenDataPayload.sourceDomain = uint32(bytes4(domainBytes));
    offset += 4;

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
  /// @return depositHash The deposit hash of the source pool data.
  function _calculateDepositHash(
    uint32 sourceDomain,
    uint256 amount,
    uint32 destinationDomain,
    bytes32 mintRecipient,
    address burnToken,
    bytes32 destinationCaller,
    uint256 maxFee,
    uint32 minFinalityThreshold
  ) internal pure returns (bytes32) {
    return keccak256(
      abi.encodePacked(
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
