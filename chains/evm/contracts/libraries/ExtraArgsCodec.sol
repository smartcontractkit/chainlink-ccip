// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/// @notice Gas-optimized assembly version of ExtraArgsCodec library.
/// @dev This is a carbon copy of ExtraArgsCodec but using inline assembly for better gas efficiency.
/// Key optimizations:
/// - Direct memory manipulation avoiding abi.encodePacked and bytes.concat
/// - Single-pass encoding with pre-calculated memory allocation
/// - Efficient pointer arithmetic
/// - Reduced memory copies
library ExtraArgsCodec {
  error InvalidDataLength(EncodingErrorLocation location);
  error InvalidEncodingVersion(uint8 version);
  error InvalidExecutorLength(uint256 length);
  error InvalidCCVAddressLength(uint256 length);
  error CCVArrayLengthMismatch(uint256 ccvsLength, uint256 ccvArgsLength);

  bytes4 public constant GENERIC_EXTRA_ARGS_V3_TAG = 0x302326cb;
  bytes4 public constant SVM_EXECUTOR_ARGS_V1_TAG = 0x1a2b3c4d;
  bytes4 public constant SUI_EXECUTOR_ARGS_V1_TAG = 0x5e6f7a8b;

  uint256 public constant MAX_CCVS = 16;
  uint256 public constant GENERIC_EXTRA_ARGS_V3_BASE_SIZE = 1 + 2 + 4 + 1 + 2 + 2 + 2;
  uint256 public constant SVM_EXECUTOR_ARGS_V1_BASE_SIZE = 1 + 8 + 2;
  uint256 public constant SUI_EXECUTOR_ARGS_V1_BASE_SIZE = 2;
  uint256 public constant CCV_BASE_SIZE = 1 + 2;

  enum EncodingErrorLocation {
    EXTRA_ARGS_CCVS_LENGTH,
    EXTRA_ARGS_CCV_ADDRESS_LENGTH,
    EXTRA_ARGS_CCV_ADDRESS_CONTENT,
    EXTRA_ARGS_CCV_ARGS_LENGTH,
    EXTRA_ARGS_CCV_ARGS_CONTENT,
    EXTRA_ARGS_FINALITY_CONFIG,
    EXTRA_ARGS_GAS_LIMIT,
    EXTRA_ARGS_EXECUTOR_LENGTH,
    EXTRA_ARGS_EXECUTOR_CONTENT,
    EXTRA_ARGS_EXECUTOR_ARGS_LENGTH,
    EXTRA_ARGS_EXECUTOR_ARGS_CONTENT,
    EXTRA_ARGS_TOKEN_RECEIVER_LENGTH,
    EXTRA_ARGS_TOKEN_RECEIVER_CONTENT,
    EXTRA_ARGS_TOKEN_ARGS_LENGTH,
    EXTRA_ARGS_TOKEN_ARGS_CONTENT,
    EXTRA_ARGS_FINAL_OFFSET,
    SVM_EXECUTOR_USE_ATA,
    SVM_EXECUTOR_ACCOUNT_BITMAP,
    SVM_EXECUTOR_ACCOUNTS_LENGTH,
    SVM_EXECUTOR_ACCOUNTS_CONTENT,
    SVM_EXECUTOR_FINAL_OFFSET,
    SUI_EXECUTOR_OBJECT_IDS_LENGTH,
    SUI_EXECUTOR_OBJECT_IDS_CONTENT,
    SUI_EXECUTOR_FINAL_OFFSET,
    ENCODE_CCVS_ARRAY_LENGTH,
    ENCODE_CCV_ARGS_ARRAY_LENGTH,
    ENCODE_CCV_ADDRESS_LENGTH,
    ENCODE_CCV_ARGS_LENGTH,
    ENCODE_EXECUTOR_LENGTH,
    ENCODE_EXECUTOR_ARGS_LENGTH,
    ENCODE_TOKEN_RECEIVER_LENGTH,
    ENCODE_TOKEN_ARGS_LENGTH,
    ENCODE_SVM_ACCOUNTS_LENGTH,
    ENCODE_SUI_OBJECT_IDS_LENGTH
  }

  struct GenericExtraArgsV3 {
    address[] ccvs;
    bytes[] ccvArgs;
    uint16 finalityConfig;
    uint32 gasLimit;
    address executor;
    bytes executorArgs;
    bytes tokenReceiver;
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

  /// @notice Encodes a GenericExtraArgsV3 struct into bytes using assembly for gas efficiency.
  /// @param extraArgs The GenericExtraArgsV3 struct to encode.
  /// @return encoded The encoded extra args as bytes.
  function _encodeGenericExtraArgsV3(
    GenericExtraArgsV3 memory extraArgs
  ) internal pure returns (bytes memory encoded) {
    // Validate ccvs and ccvArgs arrays have the same length
    uint256 ccvsLength = extraArgs.ccvs.length;
    uint256 ccvArgsLength = extraArgs.ccvArgs.length;
    if (ccvsLength != ccvArgsLength) {
      revert CCVArrayLengthMismatch(ccvsLength, ccvArgsLength);
    }

    // Validate field lengths
    if (ccvsLength > type(uint8).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_CCVS_ARRAY_LENGTH);
    }
    uint256 executorArgsLength = extraArgs.executorArgs.length;
    if (executorArgsLength > type(uint16).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_EXECUTOR_ARGS_LENGTH);
    }
    uint256 tokenReceiverLength = extraArgs.tokenReceiver.length;
    if (tokenReceiverLength > type(uint16).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_TOKEN_RECEIVER_LENGTH);
    }
    uint256 tokenArgsLength = extraArgs.tokenArgs.length;
    if (tokenArgsLength > type(uint16).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_TOKEN_ARGS_LENGTH);
    }

    // Calculate executor length
    address executor = extraArgs.executor;
    uint256 executorLength = executor == address(0) ? 0 : 20;

    // Calculate total CCV encoded size and validate
    uint256 ccvsEncodedSize = 0;
    for (uint256 i = 0; i < ccvsLength; ++i) {
      uint256 ccvAddrLength = extraArgs.ccvs[i] == address(0) ? 0 : 20;

      uint256 ccvArgLength = extraArgs.ccvArgs[i].length;
      if (ccvArgLength > type(uint16).max) {
        revert InvalidDataLength(EncodingErrorLocation.ENCODE_CCV_ARGS_LENGTH);
      }

      // 1 byte for address length + address bytes + 2 bytes for args length + args bytes
      ccvsEncodedSize += 1 + ccvAddrLength + 2 + ccvArgLength;
    }

    // Allocate memory
    // Calculate total size: tag(4) + ccvsLength(1) + ccvsData + finality(2) + gasLimit(4) +
    // executorLength(1) + executor + executorArgsLength(2) + executorArgs +
    // tokenReceiverLength(2) + tokenReceiver + tokenArgsLength(2) + tokenArgs
    encoded = new bytes(
      4 + 1 + ccvsEncodedSize + 2 + 4 + 1 + executorLength + 2 + executorArgsLength + 2 + tokenReceiverLength + 2
        + tokenArgsLength
    );

    uint256 ptr;
    assembly {
      ptr := add(encoded, 32) // Skip length prefix

      // Write tag (4 bytes)
      mstore(ptr, GENERIC_EXTRA_ARGS_V3_TAG)
      ptr := add(ptr, 4)

      // Write ccvs length (1 byte)
      mstore8(ptr, ccvsLength)
      ptr := add(ptr, 1)
    }

    // Write CCVs data
    for (uint256 i = 0; i < ccvsLength; ++i) {
      address ccvAddr = extraArgs.ccvs[i];
      uint256 ccvAddrLength = ccvAddr == address(0) ? 0 : 20;
      bytes memory ccvArg = extraArgs.ccvArgs[i];
      uint256 ccvArgLength = ccvArg.length;

      assembly {
        // Write CCV address length (1 byte)
        mstore8(ptr, ccvAddrLength)
        ptr := add(ptr, 1)

        // Write CCV address if non-zero
        if gt(ccvAddrLength, 0) {
          mstore(ptr, shl(96, ccvAddr)) // Shift address to align
          ptr := add(ptr, 20)
        }

        // Write CCV args length (2 bytes, big endian)
        mstore8(ptr, shr(8, ccvArgLength))
        mstore8(add(ptr, 1), and(ccvArgLength, 0xFF))
        ptr := add(ptr, 2)

        // Copy CCV args
        if gt(ccvArgLength, 0) {
          let ccvArgPtr := add(ccvArg, 32) // Skip bytes length prefix
          for { let end := add(ccvArgPtr, ccvArgLength) } lt(ccvArgPtr, end) { ccvArgPtr := add(ccvArgPtr, 32) } {
            mstore(ptr, mload(ccvArgPtr))
            ptr := add(ptr, 32)
          }
          // Adjust ptr if we overshot
          ptr := sub(ptr, sub(and(add(ccvArgLength, 31), not(31)), ccvArgLength))
        }
      }
    }

    uint16 finalityConfig = extraArgs.finalityConfig;
    uint32 gasLimit = extraArgs.gasLimit;

    assembly {
      // Write finality config (2 bytes, big endian)
      mstore8(ptr, shr(8, finalityConfig))
      mstore8(add(ptr, 1), and(finalityConfig, 0xFF))
      ptr := add(ptr, 2)

      // Write gas limit (4 bytes, big endian)
      mstore8(ptr, shr(24, gasLimit))
      mstore8(add(ptr, 1), and(shr(16, gasLimit), 0xFF))
      mstore8(add(ptr, 2), and(shr(8, gasLimit), 0xFF))
      mstore8(add(ptr, 3), and(gasLimit, 0xFF))
      ptr := add(ptr, 4)

      // Write executor length (1 byte)
      mstore8(ptr, executorLength)
      ptr := add(ptr, 1)

      // Write executor address if non-zero
      if gt(executorLength, 0) {
        mstore(ptr, shl(96, executor))
        ptr := add(ptr, 20)
      }
    }

    // Write executorArgs
    bytes memory executorArgs = extraArgs.executorArgs;
    assembly {
      // Write executorArgs length (2 bytes, big endian)
      mstore8(ptr, shr(8, executorArgsLength))
      mstore8(add(ptr, 1), and(executorArgsLength, 0xFF))
      ptr := add(ptr, 2)

      // Copy executorArgs
      if gt(executorArgsLength, 0) {
        let srcPtr := add(executorArgs, 32)
        for { let end := add(srcPtr, executorArgsLength) } lt(srcPtr, end) { srcPtr := add(srcPtr, 32) } {
          mstore(ptr, mload(srcPtr))
          ptr := add(ptr, 32)
        }
        ptr := sub(ptr, sub(and(add(executorArgsLength, 31), not(31)), executorArgsLength))
      }
    }

    // Write tokenReceiver
    bytes memory tokenReceiver = extraArgs.tokenReceiver;
    assembly {
      // Write tokenReceiver length (2 bytes, big endian)
      mstore8(ptr, shr(8, tokenReceiverLength))
      mstore8(add(ptr, 1), and(tokenReceiverLength, 0xFF))
      ptr := add(ptr, 2)

      // Copy tokenReceiver
      if gt(tokenReceiverLength, 0) {
        let srcPtr := add(tokenReceiver, 32)
        for { let end := add(srcPtr, tokenReceiverLength) } lt(srcPtr, end) { srcPtr := add(srcPtr, 32) } {
          mstore(ptr, mload(srcPtr))
          ptr := add(ptr, 32)
        }
        ptr := sub(ptr, sub(and(add(tokenReceiverLength, 31), not(31)), tokenReceiverLength))
      }
    }

    // Write tokenArgs
    bytes memory tokenArgs = extraArgs.tokenArgs;
    assembly {
      // Write tokenArgs length (2 bytes, big endian)
      mstore8(ptr, shr(8, tokenArgsLength))
      mstore8(add(ptr, 1), and(tokenArgsLength, 0xFF))
      ptr := add(ptr, 2)

      // Copy tokenArgs
      if gt(tokenArgsLength, 0) {
        let srcPtr := add(tokenArgs, 32)
        for { let end := add(srcPtr, tokenArgsLength) } lt(srcPtr, end) { srcPtr := add(srcPtr, 32) } {
          mstore(ptr, mload(srcPtr))
          ptr := add(ptr, 32)
        }
      }
    }

    return encoded;
  }

  /// @notice Decodes bytes into a GenericExtraArgsV3 struct using assembly for gas efficiency.
  /// @param encoded The encoded bytes to decode.
  /// @return extraArgs The decoded GenericExtraArgsV3 struct.
  function _decodeGenericExtraArgsV3(
    bytes calldata encoded
  ) internal pure returns (GenericExtraArgsV3 memory extraArgs) {
    uint256 encodedLength = encoded.length;
    uint256 offset = 4; // Skip tag

    // Read ccvs length (1 byte)
    if (offset + 1 > encodedLength) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_CCVS_LENGTH);
    uint256 ccvsLength;
    assembly {
      ccvsLength := byte(0, calldataload(add(encoded.offset, offset)))
      offset := add(offset, 1)
    }

    // Allocate arrays for CCVs
    extraArgs.ccvs = new address[](ccvsLength);
    extraArgs.ccvArgs = new bytes[](ccvsLength);

    // Decode CCVs
    for (uint256 i = 0; i < ccvsLength; ++i) {
      // Read CCV address length (1 byte)
      if (offset + 1 > encodedLength) {
        revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_CCV_ADDRESS_LENGTH);
      }

      uint256 ccvAddressLength;
      assembly {
        ccvAddressLength := byte(0, calldataload(add(encoded.offset, offset)))
        offset := add(offset, 1)
      }

      // Validate CCV address length (0 or 20 for EVM)
      if (ccvAddressLength != 0 && ccvAddressLength != 20) {
        revert InvalidCCVAddressLength(ccvAddressLength);
      }

      // Read CCV address
      if (offset + ccvAddressLength > encodedLength) {
        revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_CCV_ADDRESS_CONTENT);
      }

      if (ccvAddressLength == 20) {
        assembly {
          let addrData := calldataload(add(encoded.offset, offset))
          mstore(add(add(mload(extraArgs), 32), mul(i, 32)), shr(96, addrData))
        }
      }
      unchecked {
        offset += ccvAddressLength;
      }

      // Read CCV args length (2 bytes)
      if (offset + 2 > encodedLength) {
        revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_CCV_ARGS_LENGTH);
      }

      uint256 ccvArgsLength;
      assembly {
        let data := calldataload(add(encoded.offset, offset))
        ccvArgsLength := and(shr(240, data), 0xFFFF)
        offset := add(offset, 2)
      }

      // Read CCV args content
      if (offset + ccvArgsLength > encodedLength) {
        revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_CCV_ARGS_CONTENT);
      }

      extraArgs.ccvArgs[i] = encoded[offset:offset + ccvArgsLength];
      unchecked {
        offset += ccvArgsLength;
      }
    }

    // Read finality config (2 bytes)
    if (offset + 2 > encodedLength) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_FINALITY_CONFIG);
    assembly {
      let data := calldataload(add(encoded.offset, offset))
      mstore(add(extraArgs, 64), and(shr(240, data), 0xFFFF))
      offset := add(offset, 2)
    }

    // Read gas limit (4 bytes)
    if (offset + 4 > encodedLength) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_GAS_LIMIT);
    assembly {
      let data := calldataload(add(encoded.offset, offset))
      mstore(add(extraArgs, 96), and(shr(224, data), 0xFFFFFFFF))
      offset := add(offset, 4)
    }

    // Read executor length (1 byte)
    if (offset + 1 > encodedLength) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_EXECUTOR_LENGTH);
    uint256 executorLength;
    assembly {
      executorLength := byte(0, calldataload(add(encoded.offset, offset)))
      offset := add(offset, 1)
    }

    // Validate executor length (0 or 20 for EVM)
    if (executorLength != 0 && executorLength != 20) {
      revert InvalidExecutorLength(executorLength);
    }

    // Read executor
    if (offset + executorLength > encodedLength) {
      revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_EXECUTOR_CONTENT);
    }

    if (executorLength == 20) {
      assembly {
        let data := calldataload(add(encoded.offset, offset))
        mstore(add(extraArgs, 128), shr(96, data))
      }
    }
    unchecked {
      offset += executorLength;
    }

    // Read executorArgs length (2 bytes)
    if (offset + 2 > encodedLength) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_EXECUTOR_ARGS_LENGTH);
    uint256 executorArgsLength;
    assembly {
      let data := calldataload(add(encoded.offset, offset))
      executorArgsLength := and(shr(240, data), 0xFFFF)
      offset := add(offset, 2)
    }

    // Read executorArgs content
    if (offset + executorArgsLength > encodedLength) {
      revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_EXECUTOR_ARGS_CONTENT);
    }
    extraArgs.executorArgs = encoded[offset:offset + executorArgsLength];
    unchecked {
      offset += executorArgsLength;
    }

    // Read tokenReceiver length (2 bytes)
    if (offset + 2 > encodedLength) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_TOKEN_RECEIVER_LENGTH);
    uint256 tokenReceiverLength;
    assembly {
      let data := calldataload(add(encoded.offset, offset))
      tokenReceiverLength := and(shr(240, data), 0xFFFF)
      offset := add(offset, 2)
    }

    // Read tokenReceiver content
    if (offset + tokenReceiverLength > encodedLength) {
      revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_TOKEN_RECEIVER_CONTENT);
    }
    extraArgs.tokenReceiver = encoded[offset:offset + tokenReceiverLength];
    unchecked {
      offset += tokenReceiverLength;
    }

    // Read tokenArgs length (2 bytes)
    if (offset + 2 > encodedLength) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_TOKEN_ARGS_LENGTH);
    uint256 tokenArgsLength;
    assembly {
      let data := calldataload(add(encoded.offset, offset))
      tokenArgsLength := and(shr(240, data), 0xFFFF)
      offset := add(offset, 2)
    }

    // Read tokenArgs content
    if (offset + tokenArgsLength > encodedLength) {
      revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_TOKEN_ARGS_CONTENT);
    }
    extraArgs.tokenArgs = encoded[offset:offset + tokenArgsLength];
    unchecked {
      offset += tokenArgsLength;
    }

    // Ensure we've consumed all bytes
    if (offset != encodedLength) revert InvalidDataLength(EncodingErrorLocation.EXTRA_ARGS_FINAL_OFFSET);

    return extraArgs;
  }

  /// @notice Encodes a SVMExecutorArgsV1 struct into bytes using assembly.
  /// @param executorArgs The SVMExecutorArgsV1 struct to encode.
  /// @return encoded The encoded executor args as bytes.
  function _encodeSVMExecutorArgsV1(
    SVMExecutorArgsV1 memory executorArgs
  ) internal pure returns (bytes memory encoded) {
    uint256 accountsLength = executorArgs.accounts.length;
    if (accountsLength > type(uint8).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_SVM_ACCOUNTS_LENGTH);
    }

    // Total size: tag(4) + useATA(1) + bitmap(8) + accountsLength(1) + accounts(32 * n)
    uint256 totalSize = 4 + 1 + 8 + 1 + (accountsLength * 32);
    encoded = new bytes(totalSize);

    assembly {
      let ptr := add(encoded, 32)

      // Write tag
      mstore(ptr, SVM_EXECUTOR_ARGS_V1_TAG)
      ptr := add(ptr, 4)

      // Write useATA
      mstore8(ptr, iszero(iszero(mload(executorArgs))))
      ptr := add(ptr, 1)

      // Write accountIsWritableBitmap (8 bytes)
      let bitmap := mload(add(executorArgs, 32))
      mstore8(ptr, shr(56, bitmap))
      mstore8(add(ptr, 1), and(shr(48, bitmap), 0xFF))
      mstore8(add(ptr, 2), and(shr(40, bitmap), 0xFF))
      mstore8(add(ptr, 3), and(shr(32, bitmap), 0xFF))
      mstore8(add(ptr, 4), and(shr(24, bitmap), 0xFF))
      mstore8(add(ptr, 5), and(shr(16, bitmap), 0xFF))
      mstore8(add(ptr, 6), and(shr(8, bitmap), 0xFF))
      mstore8(add(ptr, 7), and(bitmap, 0xFF))
      ptr := add(ptr, 8)

      // Write accounts length
      mstore8(ptr, accountsLength)
      ptr := add(ptr, 1)

      // Write accounts
      let accountsPtr := add(mload(add(executorArgs, 64)), 32) // Skip array length

      for { let i := 0 } lt(i, accountsLength) { i := add(i, 1) } {
        mstore(ptr, mload(accountsPtr))
        ptr := add(ptr, 32)
        accountsPtr := add(accountsPtr, 32)
      }
    }

    return encoded;
  }

  /// @notice Decodes bytes into a SVMExecutorArgsV1 struct using assembly.
  /// @param encoded The encoded bytes to decode.
  /// @return executorArgs The decoded SVMExecutorArgsV1 struct.
  function _decodeSVMExecutorArgsV1(
    bytes calldata encoded
  ) internal pure returns (SVMExecutorArgsV1 memory executorArgs) {
    uint256 encodedLength = encoded.length;
    uint256 offset = 4; // Skip tag

    // Read useATA
    if (offset >= encodedLength) revert InvalidDataLength(EncodingErrorLocation.SVM_EXECUTOR_USE_ATA);
    assembly {
      let data := byte(0, calldataload(add(encoded.offset, offset)))
      mstore(executorArgs, iszero(iszero(data)))
    }
    offset += 1;

    // Read accountIsWritableBitmap (8 bytes)
    if (offset + 8 > encodedLength) revert InvalidDataLength(EncodingErrorLocation.SVM_EXECUTOR_ACCOUNT_BITMAP);
    assembly {
      let data := calldataload(add(encoded.offset, offset))
      mstore(add(executorArgs, 32), and(shr(192, data), 0xFFFFFFFFFFFFFFFF))
    }
    offset += 8;

    // Read accounts length
    if (offset + 1 > encodedLength) revert InvalidDataLength(EncodingErrorLocation.SVM_EXECUTOR_ACCOUNTS_LENGTH);
    uint256 accountsLength;
    assembly {
      accountsLength := byte(0, calldataload(add(encoded.offset, offset)))
    }
    offset += 1;

    // Read accounts
    if (offset + accountsLength * 32 > encodedLength) {
      revert InvalidDataLength(EncodingErrorLocation.SVM_EXECUTOR_ACCOUNTS_CONTENT);
    }

    executorArgs.accounts = new bytes32[](accountsLength);
    for (uint256 i = 0; i < accountsLength; ++i) {
      assembly {
        let data := calldataload(add(add(encoded.offset, offset), mul(i, 32)))
        let accountsArray := mload(add(executorArgs, 64))
        mstore(add(add(accountsArray, 32), mul(i, 32)), data)
      }
    }
    offset += accountsLength * 32;

    // Ensure we've consumed all bytes
    if (offset != encodedLength) revert InvalidDataLength(EncodingErrorLocation.SVM_EXECUTOR_FINAL_OFFSET);

    return executorArgs;
  }

  /// @notice Encodes a SuiExecutorArgsV1 struct into bytes using assembly.
  /// @param executorArgs The SuiExecutorArgsV1 struct to encode.
  /// @return encoded The encoded executor args as bytes.
  function _encodeSuiExecutorArgsV1(
    SuiExecutorArgsV1 memory executorArgs
  ) internal pure returns (bytes memory encoded) {
    uint256 objectIdsLength = executorArgs.receiverObjectIds.length;
    if (objectIdsLength > type(uint8).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_SUI_OBJECT_IDS_LENGTH);
    }

    // Total size: tag(4) + objectIdsLength(1) + objectIds(32 * n)
    uint256 totalSize = 4 + 1 + (objectIdsLength * 32);
    encoded = new bytes(totalSize);

    assembly {
      let ptr := add(encoded, 32)

      // Write tag
      mstore(ptr, SUI_EXECUTOR_ARGS_V1_TAG)
      ptr := add(ptr, 4)

      // Write objectIds length
      mstore8(ptr, objectIdsLength)
      ptr := add(ptr, 1)

      // Write objectIds
      let objectIdsPtr := mload(executorArgs)
      objectIdsPtr := add(objectIdsPtr, 32) // Skip array length

      for { let i := 0 } lt(i, objectIdsLength) { i := add(i, 1) } {
        mstore(ptr, mload(objectIdsPtr))
        ptr := add(ptr, 32)
        objectIdsPtr := add(objectIdsPtr, 32)
      }
    }

    return encoded;
  }

  /// @notice Decodes bytes into a SuiExecutorArgsV1 struct using assembly.
  /// @param encoded The encoded bytes to decode.
  /// @return executorArgs The decoded SuiExecutorArgsV1 struct.
  function _decodeSuiExecutorArgsV1(
    bytes calldata encoded
  ) internal pure returns (SuiExecutorArgsV1 memory executorArgs) {
    uint256 encodedLength = encoded.length;
    uint256 offset = 4; // Skip tag

    // Read objectIds length
    if (offset + 1 > encodedLength) revert InvalidDataLength(EncodingErrorLocation.SUI_EXECUTOR_OBJECT_IDS_LENGTH);
    uint256 objectIdsLength;
    assembly {
      objectIdsLength := byte(0, calldataload(add(encoded.offset, offset)))
    }
    offset += 1;

    // Read objectIds
    if (offset + objectIdsLength * 32 > encodedLength) {
      revert InvalidDataLength(EncodingErrorLocation.SUI_EXECUTOR_OBJECT_IDS_CONTENT);
    }

    executorArgs.receiverObjectIds = new bytes32[](objectIdsLength);
    for (uint256 i = 0; i < objectIdsLength; ++i) {
      assembly {
        let data := calldataload(add(add(encoded.offset, offset), mul(i, 32)))
        let objectIdsArray := mload(executorArgs)
        mstore(add(add(objectIdsArray, 32), mul(i, 32)), data)
      }
    }
    offset += objectIdsLength * 32;

    // Ensure we've consumed all bytes
    if (offset != encodedLength) revert InvalidDataLength(EncodingErrorLocation.SUI_EXECUTOR_FINAL_OFFSET);

    return executorArgs;
  }
}
