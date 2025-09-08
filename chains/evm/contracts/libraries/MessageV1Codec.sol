// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/// @notice Library for CCIP MessageV1 encoding/decoding operations.
/// @dev This library handles the complete V1 message format protocol including:
/// - MessageV1 and TokenTransferV1 struct definitions
/// - Encoding/decoding functions with comprehensive error handling
/// - Detailed error location tracking for debugging
library MessageV1Codec {
  error InvalidDataLength(EncodingErrorLocation location);
  error InvalidEncodingVersion(uint8 version);

  uint256 public constant MAX_NUMBER_OF_TOKENS = 1;

  enum EncodingErrorLocation {
    // Message-level components.
    MESSAGE_MIN_SIZE,
    MESSAGE_ONRAMP_ADDRESS_CONTENT,
    MESSAGE_OFFRAMP_ADDRESS_LENGTH,
    MESSAGE_OFFRAMP_ADDRESS_CONTENT,
    MESSAGE_FINALITY,
    MESSAGE_SENDER_LENGTH,
    MESSAGE_SENDER_CONTENT,
    MESSAGE_RECEIVER_LENGTH,
    MESSAGE_RECEIVER_CONTENT,
    MESSAGE_DEST_BLOB_LENGTH,
    MESSAGE_DEST_BLOB_CONTENT,
    MESSAGE_TOKEN_TRANSFER_LENGTH,
    MESSAGE_TOKEN_TRANSFER_CONTENT,
    MESSAGE_DATA_LENGTH,
    MESSAGE_DATA_CONTENT,
    MESSAGE_FINAL_OFFSET,
    // Token transfer components.
    TOKEN_TRANSFER_VERSION,
    TOKEN_TRANSFER_AMOUNT,
    TOKEN_TRANSFER_SOURCE_POOL_LENGTH,
    TOKEN_TRANSFER_SOURCE_POOL_CONTENT,
    TOKEN_TRANSFER_SOURCE_TOKEN_LENGTH,
    TOKEN_TRANSFER_SOURCE_TOKEN_CONTENT,
    TOKEN_TRANSFER_DEST_TOKEN_LENGTH,
    TOKEN_TRANSFER_DEST_TOKEN_CONTENT,
    TOKEN_TRANSFER_EXTRA_DATA_LENGTH,
    TOKEN_TRANSFER_EXTRA_DATA_CONTENT,
    // Encoding validation components.
    ENCODE_ONRAMP_ADDRESS_LENGTH,
    ENCODE_OFFRAMP_ADDRESS_LENGTH,
    ENCODE_SENDER_LENGTH,
    ENCODE_RECEIVER_LENGTH,
    ENCODE_DEST_BLOB_LENGTH,
    ENCODE_TOKEN_TRANSFER_ARRAY_LENGTH,
    ENCODE_DATA_LENGTH,
    ENCODE_TOKEN_SOURCE_POOL_LENGTH,
    ENCODE_TOKEN_SOURCE_TOKEN_LENGTH,
    ENCODE_TOKEN_DEST_TOKEN_LENGTH,
    ENCODE_TOKEN_EXTRA_DATA_LENGTH
  }

  /// @notice Message format used in the v1 protocol.
  /// Protocol Header.
  ///   uint8 version;              Version, for future use and backwards compatibility.
  ///   uint64 sourceChainSelector; Source Chain Selector.
  ///   uint64 destChainSelector;   Destination Chain Selector.
  ///   uint64 sequenceNumber;      Auto-incrementing sequence number for the message.
  ///   uint8 onRampAddressLength;  Length of the onRamp Address in bytes.
  ///   bytes onRampAddress;        Source Chain OnRamp as unpadded bytes.
  ///   uint8 offRampAddressLength; Length of the offRamp Address in bytes.
  ///   bytes offRampAddress;       Destination Chain OffRamp as unpadded bytes.
  ///
  /// User controlled data.
  ///   uint16 finality;            Configurable per-message finality value.
  ///   uint8 senderLength;         Length of the Sender Address in bytes.
  ///   bytes sender;               Sender address as unpadded bytes.
  ///   uint8 receiverLength;       Length of the Receiver Address in bytes.
  ///   bytes receiver;             Receiver address on the destination chain as unpadded bytes.
  ///   uint16 destBlobLength;      Length of the Destination Blob in bytes.
  ///   bytes destBlob;             Destination chain-specific blob that contains data required for execution e.g.
  ///                               Solana accounts.
  ///   uint16 tokenTransferLength; Length of the Token Transfer structure in bytes.
  ///   bytes tokenTransfer;        Byte representation of the token transfer structure.
  ///   uint16 dataLength;          Length of the user specified data payload.
  ///   bytes data;                 Arbitrary data payload supplied by the message sender that is passed to the receiver.
  ///
  /// @dev None of the fields are abi encoded as this storage layout is used for non-EVMs as well. That means if the
  /// receiver is an EVM address, it is stored as 20 bytes without any padding.
  /// @dev Inefficient struct packing does not matter as this is not a storage struct, and it it would ever be written
  /// to storage it would be in its encoded form.
  // solhint-disable-next-line gas-struct-packing
  struct MessageV1 {
    // Protocol Header.
    uint64 sourceChainSelector; // ─╮ Source Chain Selector.
    uint64 destChainSelector; //    │ Destination Chain Selector.
    //                              │ Per-lane-unique sequence number for the message. When faster-than-finality is used
    //                              │ the guarantee that this value is unique no longer holds. After a re-org, a message
    //                              │ could end up with a different sequence number. Messages that are older than the
    uint64 sequenceNumber; //  ─────╯ chain finality delay should all have unique per-lane sequence numbers.
    // Source chain onRamp, NOT abi encoded but raw bytes. This means for EVM chains it is 20 bytes.
    bytes onRampAddress;
    // Destination chain offRamp, NOT abi encoded but raw bytes. This means for EVM chains it is 20 bytes.
    bytes offRampAddress;
    // Configurable per-message finality value.
    uint16 finality;
    // Source chain sender address, NOT abi encoded but raw bytes. This means for EVM chains it is 20 bytes.
    bytes sender;
    // Destination chain receiver address, NOT abi encoded but raw bytes. This means for EVM chains it is 20 bytes.
    bytes receiver;
    // Destination specific blob that contains chain-family specific data.
    bytes destBlob;
    // Contains either 0 or 1 token transfer structs. It is encoded as an array for gas efficiency.
    TokenTransferV1[] tokenTransfer;
    // Arbitrary data payload supplied by the message sender.
    bytes data;
  }

  struct TokenTransferV1 {
    uint256 amount; // Number of tokens.
    // This can be relied upon by the destination pool to validate the source pool. NOT abi encoded but raw bytes. This
    // means for EVM chains it is 20 bytes.
    bytes sourcePoolAddress;
    bytes sourceTokenAddress; // Address of source token, NOT abi encoded but raw bytes.
    bytes destTokenAddress; // Address of destination token, NOT abi encoded but raw bytes.
    // Optional pool data to be transferred to the destination chain. Be default this is capped at
    // CCIP_LOCK_OR_BURN_V1_RET_BYTES bytes. If more data is required, the TokenTransferFeeConfig.destBytesOverhead
    // has to be set for the specific token.
    bytes extraData;
  }

  /// @notice Encodes a TokenTransferV1 struct into bytes.
  /// @param tokenTransfer The TokenTransferV1 struct to encode.
  /// @return encoded The encoded token transfer as bytes.
  function _encodeTokenTransferV1(
    TokenTransferV1 memory tokenTransfer
  ) internal pure returns (bytes memory) {
    // Validate field lengths fit in their respective size limits.
    if (tokenTransfer.sourcePoolAddress.length > type(uint8).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_TOKEN_SOURCE_POOL_LENGTH);
    }
    if (tokenTransfer.sourceTokenAddress.length > type(uint8).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_TOKEN_SOURCE_TOKEN_LENGTH);
    }
    if (tokenTransfer.destTokenAddress.length > type(uint8).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_TOKEN_DEST_TOKEN_LENGTH);
    }
    if (tokenTransfer.extraData.length > type(uint16).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_TOKEN_EXTRA_DATA_LENGTH);
    }

    return abi.encodePacked(
      uint8(1), // version.
      tokenTransfer.amount,
      uint8(tokenTransfer.sourcePoolAddress.length),
      tokenTransfer.sourcePoolAddress,
      uint8(tokenTransfer.sourceTokenAddress.length),
      tokenTransfer.sourceTokenAddress,
      uint8(tokenTransfer.destTokenAddress.length),
      tokenTransfer.destTokenAddress,
      uint16(tokenTransfer.extraData.length),
      tokenTransfer.extraData
    );
  }

  /// @notice Decodes bytes into a TokenTransferV1 struct.
  /// @param encoded The encoded token transfer bytes to decode.
  /// @param offset The starting offset in the encoded bytes.
  /// @return tokenTransfer The decoded TokenTransferV1 struct.
  /// @return newOffset The new offset after decoding.
  function _decodeTokenTransferV1(
    bytes calldata encoded,
    uint256 offset
  ) internal pure returns (TokenTransferV1 memory tokenTransfer, uint256 newOffset) {
    // version (1 byte).
    if (offset >= encoded.length) revert InvalidDataLength(EncodingErrorLocation.TOKEN_TRANSFER_VERSION);
    uint8 version = uint8(encoded[offset++]);
    if (version != 1) revert InvalidEncodingVersion(version);

    // amount (32 bytes).
    if (offset + 32 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.TOKEN_TRANSFER_AMOUNT);
    tokenTransfer.amount = uint256(bytes32(encoded[offset:offset + 32]));
    offset += 32;

    // sourcePoolAddressLength and sourcePoolAddress.
    if (offset >= encoded.length) revert InvalidDataLength(EncodingErrorLocation.TOKEN_TRANSFER_SOURCE_POOL_LENGTH);
    uint8 sourcePoolAddressLength = uint8(encoded[offset++]);
    if (offset + sourcePoolAddressLength > encoded.length) {
      revert InvalidDataLength(EncodingErrorLocation.TOKEN_TRANSFER_SOURCE_POOL_CONTENT);
    }

    tokenTransfer.sourcePoolAddress = encoded[offset:offset + sourcePoolAddressLength];
    offset += sourcePoolAddressLength;

    // sourceTokenAddressLength and sourceTokenAddress.
    if (offset >= encoded.length) revert InvalidDataLength(EncodingErrorLocation.TOKEN_TRANSFER_SOURCE_TOKEN_LENGTH);
    uint8 sourceTokenAddressLength = uint8(encoded[offset++]);
    if (offset + sourceTokenAddressLength > encoded.length) {
      revert InvalidDataLength(EncodingErrorLocation.TOKEN_TRANSFER_SOURCE_TOKEN_CONTENT);
    }

    tokenTransfer.sourceTokenAddress = encoded[offset:offset + sourceTokenAddressLength];
    offset += sourceTokenAddressLength;

    // destTokenAddressLength and destTokenAddress.
    if (offset >= encoded.length) revert InvalidDataLength(EncodingErrorLocation.TOKEN_TRANSFER_DEST_TOKEN_LENGTH);
    uint8 destTokenAddressLength = uint8(encoded[offset++]);
    if (offset + destTokenAddressLength > encoded.length) {
      revert InvalidDataLength(EncodingErrorLocation.TOKEN_TRANSFER_DEST_TOKEN_CONTENT);
    }

    tokenTransfer.destTokenAddress = encoded[offset:offset + destTokenAddressLength];
    offset += destTokenAddressLength;

    // extraDataLength and extraData.
    if (offset + 2 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.TOKEN_TRANSFER_EXTRA_DATA_LENGTH);
    uint16 extraDataLength = uint16(bytes2(encoded[offset:offset + 2]));
    offset += 2;
    if (offset + extraDataLength > encoded.length) {
      revert InvalidDataLength(EncodingErrorLocation.TOKEN_TRANSFER_EXTRA_DATA_CONTENT);
    }

    tokenTransfer.extraData = encoded[offset:offset + extraDataLength];
    offset += extraDataLength;

    return (tokenTransfer, offset);
  }

  /// @notice Encodes a MessageV1 struct into bytes following the v1 protocol format.
  /// @param message The MessageV1 struct to encode.
  /// @return encoded The encoded message as bytes.
  function _encodeMessageV1(
    MessageV1 memory message
  ) internal pure returns (bytes memory) {
    // Validate field lengths fit in their respective size limits.
    if (message.onRampAddress.length > type(uint8).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_ONRAMP_ADDRESS_LENGTH);
    }
    if (message.offRampAddress.length > type(uint8).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_OFFRAMP_ADDRESS_LENGTH);
    }
    if (message.sender.length > type(uint8).max) revert InvalidDataLength(EncodingErrorLocation.ENCODE_SENDER_LENGTH);
    if (message.receiver.length > type(uint8).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_RECEIVER_LENGTH);
    }
    if (message.destBlob.length > type(uint16).max) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_DEST_BLOB_LENGTH);
    }
    if (message.tokenTransfer.length > MAX_NUMBER_OF_TOKENS) {
      revert InvalidDataLength(EncodingErrorLocation.ENCODE_TOKEN_TRANSFER_ARRAY_LENGTH);
    }
    if (message.data.length > type(uint16).max) revert InvalidDataLength(EncodingErrorLocation.ENCODE_DATA_LENGTH);

    // Encode token the transfer if present. We checked above that there is at most 1 token transfer.
    bytes memory encodedTokenTransfers;
    if (message.tokenTransfer.length > 0) {
      encodedTokenTransfers = _encodeTokenTransferV1(message.tokenTransfer[0]);
    }

    // Encoding has to be split into groups to avoid "Stack too deep" errors.
    return bytes.concat(
      abi.encodePacked(
        uint8(1), // version.
        message.sourceChainSelector,
        message.destChainSelector,
        message.sequenceNumber,
        uint8(message.onRampAddress.length),
        message.onRampAddress,
        uint8(message.offRampAddress.length),
        message.offRampAddress,
        message.finality
      ),
      abi.encodePacked(
        uint8(message.sender.length),
        message.sender,
        uint8(message.receiver.length),
        message.receiver,
        uint16(message.destBlob.length),
        message.destBlob,
        uint16(encodedTokenTransfers.length),
        encodedTokenTransfers,
        uint16(message.data.length),
        message.data
      )
    );
  }

  /// @notice Decodes bytes into a MessageV1 struct following the v1 protocol format.
  /// @param encoded The encoded message bytes to decode.
  /// @return message The decoded MessageV1 struct.
  function _decodeMessageV1(
    bytes calldata encoded
  ) internal pure returns (MessageV1 memory) {
    if (encoded.length < 37) revert InvalidDataLength(EncodingErrorLocation.MESSAGE_MIN_SIZE); // Minimum size check.

    MessageV1 memory message;
    uint256 offset = 0;

    uint8 version = uint8(encoded[offset++]);
    if (version != 1) revert InvalidEncodingVersion(version);

    // Protocol Header.
    // sourceChainSelector (8 bytes, big endian).
    message.sourceChainSelector = uint64(bytes8(encoded[offset:offset + 8]));
    offset += 8;

    // destChainSelector (8 bytes, big endian).
    message.destChainSelector = uint64(bytes8(encoded[offset:offset + 8]));
    offset += 8;

    // sequenceNumber (8 bytes, big endian).
    message.sequenceNumber = uint64(bytes8(encoded[offset:offset + 8]));
    offset += 8;

    // onRampAddressLength and onRampAddress.
    uint8 onRampAddressLength = uint8(encoded[offset++]);
    if (offset + onRampAddressLength > encoded.length) {
      revert InvalidDataLength(EncodingErrorLocation.MESSAGE_ONRAMP_ADDRESS_CONTENT);
    }

    message.onRampAddress = encoded[offset:offset + onRampAddressLength];
    offset += onRampAddressLength;

    // offRampAddressLength and offRampAddress.
    if (offset >= encoded.length) revert InvalidDataLength(EncodingErrorLocation.MESSAGE_OFFRAMP_ADDRESS_LENGTH);
    uint8 offRampAddressLength = uint8(encoded[offset++]);
    if (offset + offRampAddressLength > encoded.length) {
      revert InvalidDataLength(EncodingErrorLocation.MESSAGE_OFFRAMP_ADDRESS_CONTENT);
    }

    message.offRampAddress = encoded[offset:offset + offRampAddressLength];
    offset += offRampAddressLength;

    // User controlled data.
    if (offset + 2 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.MESSAGE_FINALITY);

    // finality (2 bytes, big endian).
    message.finality = uint16(bytes2(encoded[offset:offset + 2]));
    offset += 2;

    // senderLength and sender.
    if (offset >= encoded.length) revert InvalidDataLength(EncodingErrorLocation.MESSAGE_SENDER_LENGTH);
    uint8 senderLength = uint8(encoded[offset++]);
    if (offset + senderLength > encoded.length) revert InvalidDataLength(EncodingErrorLocation.MESSAGE_SENDER_CONTENT);

    message.sender = encoded[offset:offset + senderLength];
    offset += senderLength;

    // receiverLength and receiver.
    if (offset >= encoded.length) revert InvalidDataLength(EncodingErrorLocation.MESSAGE_RECEIVER_LENGTH);
    uint8 receiverLength = uint8(encoded[offset++]);
    if (offset + receiverLength > encoded.length) {
      revert InvalidDataLength(EncodingErrorLocation.MESSAGE_RECEIVER_CONTENT);
    }

    message.receiver = encoded[offset:offset + receiverLength];
    offset += receiverLength;

    // destBlobLength and destBlob.
    if (offset + 2 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.MESSAGE_DEST_BLOB_LENGTH);
    uint16 destBlobLength = uint16(bytes2(encoded[offset:offset + 2]));
    offset += 2;
    if (offset + destBlobLength > encoded.length) {
      revert InvalidDataLength(EncodingErrorLocation.MESSAGE_DEST_BLOB_CONTENT);
    }

    message.destBlob = encoded[offset:offset + destBlobLength];
    offset += destBlobLength;

    // tokenTransferLength and tokenTransfer.
    if (offset + 2 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.MESSAGE_TOKEN_TRANSFER_LENGTH);
    uint16 tokenTransferLength = uint16(bytes2(encoded[offset:offset + 2]));
    offset += 2;

    // Decode token transfer, which is either 0 or 1.
    if (tokenTransferLength == 0) {
      message.tokenTransfer = new TokenTransferV1[](0);
    } else {
      message.tokenTransfer = new TokenTransferV1[](1);
      uint256 expectedEnd = offset + tokenTransferLength;
      (message.tokenTransfer[0], offset) = _decodeTokenTransferV1(encoded, offset);
      if (offset != expectedEnd) revert InvalidDataLength(EncodingErrorLocation.MESSAGE_TOKEN_TRANSFER_CONTENT);
    }

    // dataLength and data.
    if (offset + 2 > encoded.length) revert InvalidDataLength(EncodingErrorLocation.MESSAGE_DATA_LENGTH);
    uint16 dataLength = uint16(bytes2(encoded[offset:offset + 2]));
    offset += 2;
    if (offset + dataLength > encoded.length) revert InvalidDataLength(EncodingErrorLocation.MESSAGE_DATA_CONTENT);

    message.data = encoded[offset:offset + dataLength];
    offset += dataLength;

    // Ensure we've consumed all bytes.
    if (offset != encoded.length) revert InvalidDataLength(EncodingErrorLocation.MESSAGE_FINAL_OFFSET);

    return message;
  }
}
