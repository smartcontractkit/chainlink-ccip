// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/// @notice Library for CCIP Extra Arguments encoding/decoding operations.
/// @dev This library handles the GenericExtraArgsV3 format and related structures including:
/// - GenericExtraArgsV3, SVMExecutorArgsV1, and SuiExecutorArgsV1 struct definitions
/// - Encoding/decoding functions with comprehensive error handling
/// - Detailed error location tracking for debugging
/// - Variable-length encoding for addresses and arguments for efficient payload size
library ExtraArgsCodecUnoptimized {
  error InvalidDataLength(EncodingErrorLocation location);
  error InvalidEncodingVersion(uint8 version);
  error InvalidExecutorLength(uint256 length);
  error InvalidCCVAddressLength(uint256 length);
  error CCVArrayLengthMismatch(uint256 ccvsLength, uint256 ccvArgsLength);

  bytes4 public constant GENERIC_EXTRA_ARGS_V3_TAG = 0x302326cb;
  bytes4 public constant SVM_EXECUTOR_ARGS_V1_TAG = 0x1a2b3c4d; // TODO: Define actual tag
  bytes4 public constant SUI_EXECUTOR_ARGS_V1_TAG = 0x5e6f7a8b; // TODO: Define actual tag

  uint256 public constant MAX_CCVS = 16;

  // Base size of GenericExtraArgsV3 without variable length fields.
  // 1 (ccvsLength) + 2 (finalityConfig) + 4 (gasLimit) + 1 (executorLength) +
  // 2 (executorArgsLength) + 2 (tokenReceiverLength) + 2 (tokenArgsLength)
  uint256 public constant GENERIC_EXTRA_ARGS_V3_BASE_SIZE = 1 + 2 + 4 + 1 + 2 + 2 + 2;

  // Base size of SVMExecutorArgsV1 without variable length accounts field.
  // 1 (useATA) + 8 (accountIsWritableBitmap) + 2 (accountsLength)
  uint256 public constant SVM_EXECUTOR_ARGS_V1_BASE_SIZE = 1 + 8 + 2;

  // Base size of SuiExecutorArgsV1 without variable length receiverObjectIds field.
  // 2 (receiverObjectIdsLength)
  uint256 public constant SUI_EXECUTOR_ARGS_V1_BASE_SIZE = 2;

  enum EncodingErrorLocation {
    // GenericExtraArgsV3 components.
    EXTRA_ARGS_CCVS_LENGTH, // 0
    EXTRA_ARGS_CCV_ADDRESS_LENGTH,
    EXTRA_ARGS_CCV_ADDRESS_CONTENT,
    EXTRA_ARGS_CCV_ARGS_LENGTH,
    EXTRA_ARGS_CCV_ARGS_CONTENT,
    EXTRA_ARGS_FINALITY_CONFIG, // 5
    EXTRA_ARGS_GAS_LIMIT,
    EXTRA_ARGS_EXECUTOR_LENGTH,
    EXTRA_ARGS_EXECUTOR_CONTENT,
    EXTRA_ARGS_EXECUTOR_ARGS_LENGTH,
    EXTRA_ARGS_EXECUTOR_ARGS_CONTENT, // 10
    EXTRA_ARGS_TOKEN_RECEIVER_LENGTH,
    EXTRA_ARGS_TOKEN_RECEIVER_CONTENT,
    EXTRA_ARGS_TOKEN_ARGS_LENGTH,
    EXTRA_ARGS_TOKEN_ARGS_CONTENT,
    EXTRA_ARGS_FINAL_OFFSET, // 15
    // SVMExecutorArgsV1 components.
    SVM_EXECUTOR_USE_ATA,
    SVM_EXECUTOR_ACCOUNT_BITMAP,
    SVM_EXECUTOR_ACCOUNTS_LENGTH,
    SVM_EXECUTOR_ACCOUNTS_CONTENT,
    SVM_EXECUTOR_FINAL_OFFSET, // 20
    // SuiExecutorArgsV1 components.
    SUI_EXECUTOR_OBJECT_IDS_LENGTH,
    SUI_EXECUTOR_OBJECT_IDS_CONTENT,
    SUI_EXECUTOR_FINAL_OFFSET,
    // Encoding validation components.
    ENCODE_CCVS_ARRAY_LENGTH,
    ENCODE_CCV_ARGS_ARRAY_LENGTH, // 25
    ENCODE_CCV_ADDRESS_LENGTH,
    ENCODE_CCV_ARGS_LENGTH,
    ENCODE_EXECUTOR_LENGTH,
    ENCODE_EXECUTOR_ARGS_LENGTH,
    ENCODE_TOKEN_RECEIVER_LENGTH, // 30
    ENCODE_TOKEN_ARGS_LENGTH,
    ENCODE_SVM_ACCOUNTS_LENGTH,
    ENCODE_SUI_OBJECT_IDS_LENGTH
  }

  /// @notice The GenericExtraArgsV3 struct is used to pass extra arguments for all destination chain families.
  /// @dev CCVs are split into separate address and args arrays for more efficient encoding.
  struct GenericExtraArgsV3 {
    /// @notice An array of CCV addresses to be used for the message.
    address[] ccvs;
    /// @notice An array of CCV-specific arguments, parallel to the ccvs array.
    bytes[] ccvArgs;
    /// @notice The finality config, 0 means the default finality that the CCV considers final.
    uint16 finalityConfig;
    /// @notice Gas limit for the callback on the destination chain.
    uint32 gasLimit;
    /// @notice Address of the executor contract on the source chain.
    address executor;
    /// @notice Destination chain family specific arguments for the executor.
    bytes executorArgs;
    /// @notice Address of the token receiver on the destination chain, in bytes format.
    bytes tokenReceiver;
    /// @notice Additional arguments for token transfers.
    bytes tokenArgs;
  }

  struct SVMExecutorArgsV1 {
    bool useATA;
    uint64 accountIsWritableBitmap;
    bytes32[] accounts;
  }

  struct SuiExecutorArgsV1 {
    bytes32[] receiverObjectIds;
  }

  /// @notice Encodes a GenericExtraArgsV3 struct into bytes.
  /// @param extraArgs The GenericExtraArgsV3 struct to encode.
  /// @return encoded The encoded extra args as bytes.
  function _encodeGenericExtraArgsV3(
    GenericExtraArgsV3 memory extraArgs
  ) internal pure returns (bytes memory) {
    // Validate ccvs and ccvArgs arrays have the same length
    if (extraArgs.ccvs.length != extraArgs.ccvArgs.length) {
      revert CCVArrayLengthMismatch(extraArgs.ccvs.length, extraArgs.ccvArgs.length);
    }

    // Validate field lengths fit in their respective size limits.
    if (extraArgs.ccvs.length > type(uint8).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_CCVS_ARRAY_LENGTH);
    }
    if (extraArgs.executorArgs.length > type(uint16).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_EXECUTOR_ARGS_LENGTH);
    }
    if (extraArgs.tokenReceiver.length > type(uint16).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_TOKEN_RECEIVER_LENGTH);
    }
    if (extraArgs.tokenArgs.length > type(uint16).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_TOKEN_ARGS_LENGTH);
    }

    // Encode executor as variable length (0 for address(0), 20 for non-zero addresses)
    bytes memory encodedExecutor = extraArgs.executor == address(0) ? bytes("") : abi.encodePacked(extraArgs.executor);
    if (encodedExecutor.length > type(uint8).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_EXECUTOR_LENGTH);
    }

    // Encode CCVs.
    bytes memory encodedCCVs;
    for (uint256 i = 0; i < extraArgs.ccvs.length; ++i) {
      // Encode CCV address as variable length (0 for address(0), 20 for non-zero addresses)
      bytes memory encodedCCVAddress = extraArgs.ccvs[i] == address(0) ? bytes("") : abi.encodePacked(extraArgs.ccvs[i]);
      if (encodedCCVAddress.length > type(uint8).max) {
        revert InvalidDataLength(EncodingErrorLocation.ENCODE_CCV_ADDRESS_LENGTH);
      }
      if (extraArgs.ccvArgs[i].length > type(uint16).max) {
        revert InvalidDataLength(EncodingErrorLocation.ENCODE_CCV_ARGS_LENGTH);
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
      GENERIC_EXTRA_ARGS_V3_TAG,
      uint8(extraArgs.ccvs.length),
      encodedCCVs,
      extraArgs.finalityConfig,
      extraArgs.gasLimit,
      uint8(encodedExecutor.length),
      encodedExecutor
    );

    return bytes.concat(part1, abi.encodePacked(
      uint16(extraArgs.executorArgs.length),
      extraArgs.executorArgs,
      uint16(extraArgs.tokenReceiver.length),
      extraArgs.tokenReceiver,
      uint16(extraArgs.tokenArgs.length),
      extraArgs.tokenArgs
    ));
  }

  /// @notice Decodes bytes into a GenericExtraArgsV3 struct.
  /// @param encoded The encoded bytes to decode.
  /// @return extraArgs The decoded GenericExtraArgsV3 struct.
  function _decodeGenericExtraArgsV3(
    bytes calldata encoded
  ) internal pure returns (GenericExtraArgsV3 memory extraArgs) {
    uint256 offset = 0;

    // Tag (4 bytes) - already validated by caller typically, but we skip it here.
    offset += 4;

    // ccvs length (1 byte).
    if (offset + 1 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_CCVS_LENGTH);
    uint256 ccvsLength = uint8(bytes1(encoded[offset:offset + 1]));
    offset += 1;

    // Decode CCVs.
    extraArgs.ccvs = new address[](ccvsLength);
    extraArgs.ccvArgs = new bytes[](ccvsLength);
    for (uint256 i = 0; i < ccvsLength; ++i) {
      // CCV address length (1 byte).
      if (offset + 1 > encoded.length) {
        revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_CCV_ADDRESS_LENGTH);
      }
      uint256 ccvAddressLength = uint8(bytes1(encoded[offset:offset + 1]));
      offset += 1;

      // CCV address content - must be 0 or 20 bytes for EVM
      if (ccvAddressLength != 0 && ccvAddressLength != 20) {
        revert InvalidCCVAddressLength(ccvAddressLength);
      }
      if (offset + ccvAddressLength > encoded.length) {
        revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_CCV_ADDRESS_CONTENT);
      }
      extraArgs.ccvs[i] =
        ccvAddressLength == 0 ? address(0) : address(bytes20(encoded[offset:offset + ccvAddressLength]));
      offset += ccvAddressLength;

      // CCV argsLength (2 bytes).
      if (offset + 2 > encoded.length) {
        revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_CCV_ARGS_LENGTH);
      }
      uint256 ccvArgsLength = uint16(bytes2(encoded[offset:offset + 2]));
      offset += 2;

      // CCV args content.
      if (offset + ccvArgsLength > encoded.length) {
        revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_CCV_ARGS_CONTENT);
      }
      extraArgs.ccvArgs[i] = encoded[offset:offset + ccvArgsLength];
      offset += ccvArgsLength;
    }

    // finalityConfig (2 bytes).
    if (offset + 2 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_FINALITY_CONFIG);
    extraArgs.finalityConfig = uint16(bytes2(encoded[offset:offset + 2]));
    offset += 2;

    // gasLimit (4 bytes).
    if (offset + 4 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_GAS_LIMIT);
    extraArgs.gasLimit = uint32(bytes4(encoded[offset:offset + 4]));
    offset += 4;

    // executorLength (1 byte).
    if (offset + 1 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_EXECUTOR_LENGTH);
    uint256 executorLength = uint8(bytes1(encoded[offset:offset + 1]));
    offset += 1;

    // executor content - must be 0 or 20 bytes for EVM
    if (executorLength != 0 && executorLength != 20) {
      revert InvalidExecutorLength(executorLength);
    }
    if (offset + executorLength > encoded.length) {
      revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_EXECUTOR_CONTENT);
    }
    extraArgs.executor = executorLength == 0 ? address(0) : address(bytes20(encoded[offset:offset + executorLength]));
    offset += executorLength;

    // executorArgsLength (2 bytes).
    if (offset + 2 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_EXECUTOR_ARGS_LENGTH);
    uint256 executorArgsLength = uint16(bytes2(encoded[offset:offset + 2]));
    offset += 2;

    // executorArgs content.
    if (offset + executorArgsLength > encoded.length) {
      revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_EXECUTOR_ARGS_CONTENT);
    }
    extraArgs.executorArgs = encoded[offset:offset + executorArgsLength];
    offset += executorArgsLength;

    // tokenReceiverLength (2 bytes).
    if (offset + 2 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_TOKEN_RECEIVER_LENGTH);
    uint256 tokenReceiverLength = uint16(bytes2(encoded[offset:offset + 2]));
    offset += 2;

    // tokenReceiver content.
    if (offset + tokenReceiverLength > encoded.length) {
      revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_TOKEN_RECEIVER_CONTENT);
    }
    extraArgs.tokenReceiver = encoded[offset:offset + tokenReceiverLength];
    offset += tokenReceiverLength;

    // tokenArgsLength (2 bytes).
    if (offset + 2 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_TOKEN_ARGS_LENGTH);
    uint256 tokenArgsLength = uint16(bytes2(encoded[offset:offset + 2]));
    offset += 2;

    // tokenArgs content.
    if (offset + tokenArgsLength > encoded.length) {
      revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_TOKEN_ARGS_CONTENT);
    }
    extraArgs.tokenArgs = encoded[offset:offset + tokenArgsLength];
    offset += tokenArgsLength;

    // Ensure we've consumed all bytes.
    if (offset != encoded.length) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_FINAL_OFFSET);

    return extraArgs;
  }

  /// @notice Encodes a SVMExecutorArgsV1 struct into bytes.
  /// @param executorArgs The SVMExecutorArgsV1 struct to encode.
  /// @return encoded The encoded executor args as bytes.
  function _encodeSVMExecutorArgsV1(
    SVMExecutorArgsV1 memory executorArgs
  ) internal pure returns (bytes memory) {
    if (executorArgs.accounts.length > type(uint8).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_SVM_ACCOUNTS_LENGTH);
    }

    return abi.encodePacked(
      SVM_EXECUTOR_ARGS_V1_TAG,
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
  ) internal pure returns (SVMExecutorArgsV1 memory executorArgs) {
    uint256 offset = 0;

    // Tag (4 bytes) - skip.
    offset += 4;

    // useATA (1 byte).
    if (offset >= encoded.length) revert InvalidDataLength(EncodingErrorLocation.SVM_EXECUTOR_USE_ATA);
    executorArgs.useATA = encoded[offset++] != 0;

    // accountIsWritableBitmap (8 bytes).
    if (offset + 8 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.SVM_EXECUTOR_ACCOUNT_BITMAP);
    executorArgs.accountIsWritableBitmap = uint64(bytes8(encoded[offset:offset + 8]));
    offset += 8;

    // accounts length (1 byte).
    if (offset + 1 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.SVM_EXECUTOR_ACCOUNTS_LENGTH);
    uint256 accountsLength = uint8(bytes1(encoded[offset:offset + 1]));
    offset += 1;

    // accounts content (32 bytes each).
    if (offset + accountsLength * 32 > encoded.length) {
      revert InvalidDataLength(EncodingErrorLocation.SVM_EXECUTOR_ACCOUNTS_CONTENT);
    }
    executorArgs.accounts = new bytes32[](accountsLength);
    for (uint256 i = 0; i < accountsLength; ++i) {
      executorArgs.accounts[i] = bytes32(encoded[offset:offset + 32]);
      offset += 32;
    }

    // Ensure we've consumed all bytes.
    if (offset != encoded.length) revert InvalidDataLength(EncodingErrorLocation.SVM_EXECUTOR_FINAL_OFFSET);

    return executorArgs;
  }

  /// @notice Encodes a SuiExecutorArgsV1 struct into bytes.
  /// @param executorArgs The SuiExecutorArgsV1 struct to encode.
  /// @return encoded The encoded executor args as bytes.
  function _encodeSuiExecutorArgsV1(
    SuiExecutorArgsV1 memory executorArgs
  ) internal pure returns (bytes memory) {
    if (executorArgs.receiverObjectIds.length > type(uint8).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_SUI_OBJECT_IDS_LENGTH);
    }

    return abi.encodePacked(
      SUI_EXECUTOR_ARGS_V1_TAG,
      uint8(executorArgs.receiverObjectIds.length),
      abi.encodePacked(executorArgs.receiverObjectIds)
    );
  }

  /// @notice Decodes bytes into a SuiExecutorArgsV1 struct.
  /// @param encoded The encoded bytes to decode.
  /// @return executorArgs The decoded SuiExecutorArgsV1 struct.
  function _decodeSuiExecutorArgsV1(
    bytes calldata encoded
  ) internal pure returns (SuiExecutorArgsV1 memory executorArgs) {
    uint256 offset = 0;

    // Tag (4 bytes) - skip.
    offset += 4;

    // receiverObjectIds length (1 byte).
    if (offset + 1 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.SUI_EXECUTOR_OBJECT_IDS_LENGTH);
    uint256 objectIdsLength = uint16(bytes2(encoded[offset:offset + 1]));
    offset += 1;

    // receiverObjectIds content (32 bytes each).
    if (offset + objectIdsLength * 32 > encoded.length) {
      revert InvalidDataLength(EncodingErrorLocation.SUI_EXECUTOR_OBJECT_IDS_CONTENT);
    }
    executorArgs.receiverObjectIds = new bytes32[](objectIdsLength);
    for (uint256 i = 0; i < objectIdsLength; ++i) {
      executorArgs.receiverObjectIds[i] = bytes32(encoded[offset:offset + 32]);
      offset += 32;
    }

    // Ensure we've consumed all bytes.
    if (offset != encoded.length) revert InvalidDataLength(EncodingErrorLocation.SUI_EXECUTOR_FINAL_OFFSET);

    return executorArgs;
  }
}
