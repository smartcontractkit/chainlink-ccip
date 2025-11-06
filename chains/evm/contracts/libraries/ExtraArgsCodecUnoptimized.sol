// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {ExtraArgsCodec} from "./ExtraArgsCodec.sol";

/// @notice Unoptimized version of CCIP Extra Arguments encoding/decoding operations for gas comparison.
/// @dev This library uses the same structs, constants, errors, and enums as ExtraArgsCodec but with
/// standard Solidity encoding/decoding instead of assembly optimizations. This allows for direct gas
/// comparison between optimized and unoptimized implementations using identical definitions.
/// @dev This library is used for differential fuzzing tests to ensure parity with the optimized version.
library ExtraArgsCodecUnoptimized {
  function _encodeGenericExtraArgsV3(
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs
  ) internal pure returns (bytes memory) {
    // Validate ccvs and ccvArgs arrays have the same length
    if (extraArgs.ccvs.length != extraArgs.ccvArgs.length) {
      revert ExtraArgsCodec.CCVArrayLengthMismatch(extraArgs.ccvs.length, extraArgs.ccvArgs.length);
    }

    // Validate field lengths fit in their respective size limits.
    if (extraArgs.ccvs.length > type(uint8).max) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.ENCODE_CCVS_ARRAY_LENGTH);
    }
    if (extraArgs.executorArgs.length > type(uint16).max) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.ENCODE_EXECUTOR_ARGS_LENGTH);
    }
    if (extraArgs.tokenReceiver.length > type(uint16).max) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.ENCODE_TOKEN_RECEIVER_LENGTH);
    }
    if (extraArgs.tokenArgs.length > type(uint16).max) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.ENCODE_TOKEN_ARGS_LENGTH);
    }

    // Encode executor as variable length (0 for address(0), 20 for non-zero addresses)
    bytes memory encodedExecutor = extraArgs.executor == address(0) ? bytes("") : abi.encodePacked(extraArgs.executor);
    if (encodedExecutor.length > type(uint8).max) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.ENCODE_EXECUTOR_LENGTH);
    }

    // Encode CCVs.
    bytes memory encodedCCVs;
    for (uint256 i = 0; i < extraArgs.ccvs.length; ++i) {
      // Encode CCV address as variable length (0 for address(0), 20 for non-zero addresses)
      bytes memory encodedCCVAddress = extraArgs.ccvs[i] == address(0) ? bytes("") : abi.encodePacked(extraArgs.ccvs[i]);
      if (encodedCCVAddress.length > type(uint8).max) {
        revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.ENCODE_CCV_ADDRESS_LENGTH);
      }
      if (extraArgs.ccvArgs[i].length > type(uint16).max) {
        revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.ENCODE_CCV_ARGS_LENGTH);
      }

      encodedCCVs = bytes.concat(
        encodedCCVs,
        abi.encodePacked(
          uint8(encodedCCVAddress.length), encodedCCVAddress, uint16(extraArgs.ccvArgs[i].length), extraArgs.ccvArgs[i]
        )
      );
    }

    // Split encoding to avoid stack too deep.
    bytes memory part1 = abi.encodePacked(
      ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_TAG,
      extraArgs.gasLimit,
      extraArgs.finalityConfig,
      uint8(extraArgs.ccvs.length),
      encodedCCVs,
      uint8(encodedExecutor.length),
      encodedExecutor
    );

    return bytes.concat(
      part1,
      abi.encodePacked(
        uint16(extraArgs.executorArgs.length),
        extraArgs.executorArgs,
        uint16(extraArgs.tokenReceiver.length),
        extraArgs.tokenReceiver,
        uint16(extraArgs.tokenArgs.length),
        extraArgs.tokenArgs
      )
    );
  }

  /// @notice Decodes bytes into a GenericExtraArgsV3 struct.
  /// @param encoded The encoded bytes to decode.
  /// @return extraArgs The decoded GenericExtraArgsV3 struct.
  function _decodeGenericExtraArgsV3(
    bytes calldata encoded
  ) internal pure returns (ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs) {
    if (encoded.length < ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_BASE_SIZE) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_STATIC_LENGTH_FIELDS);
    }

    // Tag (4 bytes) - already validated by caller typically, but we skip it here.
    uint256 offset = 4;

    // gasLimit (4 bytes).
    extraArgs.gasLimit = uint32(bytes4(encoded[offset:offset + 4]));
    offset += 4;

    // finalityConfig (2 bytes).
    extraArgs.finalityConfig = uint16(bytes2(encoded[offset:offset + 2]));
    offset += 2;

    // ccvs length (1 byte).
    uint256 ccvsLength = uint8(bytes1(encoded[offset:offset + 1]));
    offset += 1;

    // Decode CCVs.
    extraArgs.ccvs = new address[](ccvsLength);
    extraArgs.ccvArgs = new bytes[](ccvsLength);
    for (uint256 i = 0; i < ccvsLength; ++i) {
      // CCV address length (1 byte).
      if (offset + 1 > encoded.length) {
        revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_CCV_ADDRESS_LENGTH);
      }
      uint256 ccvAddressLength = uint8(bytes1(encoded[offset:offset + 1]));
      offset += 1;

      // CCV address content - must be 0 or 20 bytes for EVM
      if (ccvAddressLength != 0 && ccvAddressLength != 20) {
        revert ExtraArgsCodec.InvalidCCVAddressLength(ccvAddressLength);
      }
      if (offset + ccvAddressLength > encoded.length) {
        revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_CCV_ADDRESS_CONTENT);
      }
      extraArgs.ccvs[i] =
        ccvAddressLength == 0 ? address(0) : address(bytes20(encoded[offset:offset + ccvAddressLength]));
      offset += ccvAddressLength;

      // CCV argsLength (2 bytes).
      if (offset + 2 > encoded.length) {
        revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_CCV_ARGS_LENGTH);
      }
      uint256 ccvArgsLength = uint16(bytes2(encoded[offset:offset + 2]));
      offset += 2;

      // CCV args content.
      if (offset + ccvArgsLength > encoded.length) {
        revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_CCV_ARGS_CONTENT);
      }
      extraArgs.ccvArgs[i] = encoded[offset:offset + ccvArgsLength];
      offset += ccvArgsLength;
    }

    // executorLength (1 byte).
    if (offset + 1 > encoded.length) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_EXECUTOR_LENGTH);
    }
    uint256 executorLength = uint8(bytes1(encoded[offset:offset + 1]));
    offset += 1;

    // executor content - must be 0 or 20 bytes for EVM
    if (executorLength != 0 && executorLength != 20) {
      revert ExtraArgsCodec.InvalidExecutorLength(executorLength);
    }
    if (offset + executorLength > encoded.length) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_EXECUTOR_CONTENT);
    }
    extraArgs.executor = executorLength == 0 ? address(0) : address(bytes20(encoded[offset:offset + executorLength]));
    offset += executorLength;

    // executorArgsLength (2 bytes).
    if (offset + 2 > encoded.length) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_EXECUTOR_ARGS_LENGTH);
    }
    uint256 executorArgsLength = uint16(bytes2(encoded[offset:offset + 2]));
    offset += 2;

    // executorArgs content.
    if (offset + executorArgsLength > encoded.length) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_EXECUTOR_ARGS_CONTENT);
    }
    extraArgs.executorArgs = encoded[offset:offset + executorArgsLength];
    offset += executorArgsLength;

    // tokenReceiverLength (2 bytes).
    if (offset + 2 > encoded.length) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_TOKEN_RECEIVER_LENGTH);
    }
    uint256 tokenReceiverLength = uint16(bytes2(encoded[offset:offset + 2]));
    offset += 2;

    // tokenReceiver content.
    if (offset + tokenReceiverLength > encoded.length) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_TOKEN_RECEIVER_CONTENT);
    }
    extraArgs.tokenReceiver = encoded[offset:offset + tokenReceiverLength];
    offset += tokenReceiverLength;

    // tokenArgsLength (2 bytes).
    if (offset + 2 > encoded.length) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_TOKEN_ARGS_LENGTH);
    }
    uint256 tokenArgsLength = uint16(bytes2(encoded[offset:offset + 2]));
    offset += 2;

    // tokenArgs content.
    if (offset + tokenArgsLength > encoded.length) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_TOKEN_ARGS_CONTENT);
    }
    extraArgs.tokenArgs = encoded[offset:offset + tokenArgsLength];
    offset += tokenArgsLength;

    // Ensure we've consumed all bytes.
    if (offset != encoded.length) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_FINAL_OFFSET);
    }

    return extraArgs;
  }

  /// @notice Encodes a SVMExecutorArgsV1 struct into bytes.
  /// @param executorArgs The SVMExecutorArgsV1 struct to encode.
  /// @return encoded The encoded executor args as bytes.
  function _encodeSVMExecutorArgsV1(
    ExtraArgsCodec.SVMExecutorArgsV1 memory executorArgs
  ) internal pure returns (bytes memory) {
    if (executorArgs.accounts.length > type(uint8).max) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.ENCODE_SVM_ACCOUNTS_LENGTH);
    }

    return abi.encodePacked(
      ExtraArgsCodec.SVM_EXECUTOR_ARGS_V1_TAG,
      executorArgs.useATA,
      executorArgs.accountIsWritableBitmap,
      uint8(executorArgs.accounts.length),
      abi.encodePacked(executorArgs.accounts)
    );
  }

  /// @notice Decodes bytes into a SVMExecutorArgsV1 struct.
  /// @param encoded The encoded bytes to decode.
  /// @return executorArgs The decoded SVMExecutorArgsV1 struct.
  function _decodeSVMExecutorArgsV1(
    bytes calldata encoded
  ) internal pure returns (ExtraArgsCodec.SVMExecutorArgsV1 memory executorArgs) {
    if (encoded.length < ExtraArgsCodec.SVM_EXECUTOR_ARGS_V1_BASE_SIZE) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_STATIC_LENGTH_FIELDS);
    }
    // Tag (4 bytes) - skip.
    uint256 offset = 4;

    // useATA (1 byte).
    executorArgs.useATA = encoded[offset++] != 0;

    // accountIsWritableBitmap (8 bytes).
    executorArgs.accountIsWritableBitmap = uint64(bytes8(encoded[offset:offset + 8]));
    offset += 8;

    // accounts length (1 byte).
    uint256 accountsLength = uint8(bytes1(encoded[offset:offset + 1]));
    offset += 1;

    // accounts content (32 bytes each).
    if (offset + accountsLength * 32 > encoded.length) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.SVM_EXECUTOR_ACCOUNTS_CONTENT);
    }
    executorArgs.accounts = new bytes32[](accountsLength);
    for (uint256 i = 0; i < accountsLength; ++i) {
      executorArgs.accounts[i] = bytes32(encoded[offset:offset + 32]);
      offset += 32;
    }

    // Ensure we've consumed all bytes.
    if (offset != encoded.length) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.SVM_EXECUTOR_FINAL_OFFSET);
    }

    return executorArgs;
  }

  /// @notice Encodes a SuiExecutorArgsV1 struct into bytes.
  /// @param executorArgs The SuiExecutorArgsV1 struct to encode.
  /// @return encoded The encoded executor args as bytes.
  function _encodeSuiExecutorArgsV1(
    ExtraArgsCodec.SuiExecutorArgsV1 memory executorArgs
  ) internal pure returns (bytes memory) {
    if (executorArgs.receiverObjectIds.length > type(uint8).max) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.ENCODE_SUI_OBJECT_IDS_LENGTH);
    }

    return abi.encodePacked(
      ExtraArgsCodec.SUI_EXECUTOR_ARGS_V1_TAG,
      uint8(executorArgs.receiverObjectIds.length),
      abi.encodePacked(executorArgs.receiverObjectIds)
    );
  }

  /// @notice Decodes bytes into a SuiExecutorArgsV1 struct.
  /// @param encoded The encoded bytes to decode.
  /// @return executorArgs The decoded SuiExecutorArgsV1 struct.
  function _decodeSuiExecutorArgsV1(
    bytes calldata encoded
  ) internal pure returns (ExtraArgsCodec.SuiExecutorArgsV1 memory executorArgs) {
    // Tag (4 bytes) - skip.
    uint256 offset = 4;

    // receiverObjectIds length (1 byte).
    if (offset + 1 > encoded.length) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_STATIC_LENGTH_FIELDS);
    }
    uint256 objectIdsLength = uint16(bytes2(encoded[offset:offset + 1]));
    offset += 1;

    // receiverObjectIds content (32 bytes each).
    if (offset + objectIdsLength * 32 > encoded.length) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.SUI_EXECUTOR_OBJECT_IDS_CONTENT);
    }
    executorArgs.receiverObjectIds = new bytes32[](objectIdsLength);
    for (uint256 i = 0; i < objectIdsLength; ++i) {
      executorArgs.receiverObjectIds[i] = bytes32(encoded[offset:offset + 32]);
      offset += 32;
    }

    // Ensure we've consumed all bytes.
    if (offset != encoded.length) {
      revert ExtraArgsCodec.InvalidDataLength(ExtraArgsCodec.EncodingErrorLocation.SUI_EXECUTOR_FINAL_OFFSET);
    }

    return executorArgs;
  }
}
