// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {MessageV1CodecSetup} from "./MessageV1CodecSetup.t.sol";

contract MessageV1Codec__decodeMessageV1 is MessageV1CodecSetup {
  function test__decodeMessageV1_WithData() public {
    MessageV1Codec.MessageV1 memory originalMessage = MessageV1Codec.MessageV1({
      sourceChainSelector: 5,
      destChainSelector: 10,
      sequenceNumber: 200,
      onRampAddress: abi.encodePacked(makeAddr("onRamp")),
      offRampAddress: abi.encodePacked(makeAddr("offRamp")),
      finality: 1000,
      sender: abi.encodePacked(makeAddr("sender")),
      receiver: abi.encodePacked(makeAddr("receiver")),
      destBlob: "destination blob data",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](0),
      data: "arbitrary message data"
    });

    MessageV1Codec.MessageV1 memory decodedMessage = s_helper.decodeMessageV1(s_helper.encodeMessageV1(originalMessage));

    _assertMessageEqual(originalMessage, decodedMessage);
  }

  function test__decodeMessageV1_WithTokenTransfer() public {
    MessageV1Codec.TokenTransferV1[] memory tokenTransfers = new MessageV1Codec.TokenTransferV1[](1);
    tokenTransfers[0] = MessageV1Codec.TokenTransferV1({
      amount: 1000000,
      sourcePoolAddress: abi.encodePacked(makeAddr("sourcePool")),
      sourceTokenAddress: abi.encodePacked(makeAddr("sourceToken")),
      destTokenAddress: abi.encodePacked(makeAddr("destToken")),
      extraData: "token extra data"
    });

    MessageV1Codec.MessageV1 memory originalMessage = MessageV1Codec.MessageV1({
      sourceChainSelector: 123,
      destChainSelector: 456,
      sequenceNumber: 789,
      onRampAddress: abi.encodePacked(makeAddr("onRamp")),
      offRampAddress: abi.encodePacked(makeAddr("offRamp")),
      finality: 2000,
      sender: abi.encodePacked(makeAddr("sender")),
      receiver: abi.encodePacked(makeAddr("receiver")),
      destBlob: "complex destination blob",
      tokenTransfer: tokenTransfers,
      data: "message with token data"
    });

    MessageV1Codec.MessageV1 memory decodedMessage = s_helper.decodeMessageV1(s_helper.encodeMessageV1(originalMessage));

    _assertMessageEqual(originalMessage, decodedMessage);
  }

  function test__decodeMessageV1_MaxLengthFields() public view {
    uint256 testDataLength = 1000; // Reasonable size for testing
    // Create maximum length fields to test boundary conditions
    bytes memory maxLengthBytes = new bytes(type(uint8).max); // 255 bytes
    bytes memory maxLengthData = new bytes(testDataLength);

    // Fill with some pattern for verification
    for (uint256 i = 0; i < maxLengthBytes.length; ++i) {
      maxLengthBytes[i] = bytes1(uint8(i % 256));
    }
    for (uint256 i = 0; i < testDataLength; ++i) {
      maxLengthData[i] = bytes1(uint8(i % 256));
    }

    MessageV1Codec.TokenTransferV1[] memory tokenTransfers = new MessageV1Codec.TokenTransferV1[](1);
    tokenTransfers[0] = MessageV1Codec.TokenTransferV1({
      amount: type(uint256).max,
      sourcePoolAddress: maxLengthBytes,
      sourceTokenAddress: maxLengthBytes,
      destTokenAddress: maxLengthBytes,
      extraData: maxLengthData
    });

    MessageV1Codec.MessageV1 memory originalMessage = MessageV1Codec.MessageV1({
      sourceChainSelector: type(uint64).max,
      destChainSelector: type(uint64).max,
      sequenceNumber: type(uint64).max,
      onRampAddress: maxLengthBytes,
      offRampAddress: maxLengthBytes,
      finality: type(uint16).max,
      sender: maxLengthBytes,
      receiver: maxLengthBytes,
      destBlob: maxLengthData,
      tokenTransfer: tokenTransfers,
      data: maxLengthData
    });

    bytes memory encoded = s_helper.encodeMessageV1(originalMessage);
    MessageV1Codec.MessageV1 memory decodedMessage = s_helper.decodeMessageV1(encoded);

    _assertMessageEqual(originalMessage, decodedMessage);
  }

  function test__decodeMessageV1_EmptyFields() public view {
    MessageV1Codec.MessageV1 memory originalMessage = MessageV1Codec.MessageV1({
      sourceChainSelector: 1,
      destChainSelector: 2,
      sequenceNumber: 3,
      onRampAddress: "",
      offRampAddress: "",
      finality: 0,
      sender: "",
      receiver: "",
      destBlob: "",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](0),
      data: ""
    });

    bytes memory encoded = s_helper.encodeMessageV1(originalMessage);
    MessageV1Codec.MessageV1 memory decodedMessage = s_helper.decodeMessageV1(encoded);

    _assertMessageEqual(originalMessage, decodedMessage);
  }

  // Reverts

  /// forge-config: default.allow_internal_expect_revert = true
  function test__decodeMessageV1_RevertWhen_InvalidVersion() public {
    // Create a valid encoded message first
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    bytes memory encoded = s_helper.encodeMessageV1(message);

    // Corrupt the version byte (first byte should be 1)
    bytes memory corruptedEncoded = encoded;
    corruptedEncoded[0] = 0x02; // Invalid version

    vm.expectRevert(abi.encodeWithSelector(MessageV1Codec.InvalidEncodingVersion.selector, 2));
    s_helper.decodeMessageV1(corruptedEncoded);
  }

  /// forge-config: default.allow_internal_expect_revert = true
  function test__decodeMessageV1_RevertWhen_InvalidTokenTransferVersion() public {
    // Create a message with token transfer
    MessageV1Codec.TokenTransferV1[] memory tokenTransfers = new MessageV1Codec.TokenTransferV1[](1);
    tokenTransfers[0] = MessageV1Codec.TokenTransferV1({
      amount: 1000,
      sourcePoolAddress: abi.encodePacked(makeAddr("pool")),
      sourceTokenAddress: abi.encodePacked(makeAddr("token")),
      destTokenAddress: abi.encodePacked(makeAddr("destToken")),
      extraData: ""
    });

    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    message.tokenTransfer = tokenTransfers;

    bytes memory encoded = s_helper.encodeMessageV1(message);

    // Find and corrupt the token transfer version byte
    // The token transfer version should be after the main message fields
    uint256 tokenVersionOffset = _findTokenTransferVersionOffset(encoded);

    bytes memory corruptedEncoded = encoded;
    corruptedEncoded[tokenVersionOffset] = 0x00; // Invalid token version

    vm.expectRevert(abi.encodeWithSelector(MessageV1Codec.InvalidEncodingVersion.selector, 0));
    s_helper.decodeMessageV1(corruptedEncoded);
  }

  /// forge-config: default.allow_internal_expect_revert = true
  function test__decodeMessageV1_RevertWhen_InsufficientDataLength() public {
    // Create a valid encoded message
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    bytes memory encoded = s_helper.encodeMessageV1(message);

    // Truncate the encoded data
    bytes memory truncatedEncoded = new bytes(10); // Too short
    for (uint256 i = 0; i < 10; ++i) {
      truncatedEncoded[i] = encoded[i];
    }

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.MESSAGE_MIN_SIZE
      )
    );
    s_helper.decodeMessageV1(truncatedEncoded);
  }

  function test__decodeMessageV1_RevertWhen_OnRampAddressContentTruncated() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    bytes memory encoded = s_helper.encodeMessageV1(message);

    // Truncate right after onRampAddress length byte to cause content error
    uint256 truncatePoint = 1 + 8 + 8 + 8 + 1; // version + selectors + seq + onRampLength
    bytes memory truncated = new bytes(truncatePoint + 12); // Not enough for the address but enough to be > 37
    for (uint256 i = 0; i < truncated.length; i++) {
      truncated[i] = encoded[i];
    }

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.MESSAGE_ONRAMP_ADDRESS_CONTENT
      )
    );
    s_helper.decodeMessageV1(truncated);
  }

  function test__decodeMessageV1_RevertWhen_OffRampAddressLengthTruncated() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    bytes memory encoded = s_helper.encodeMessageV1(message);

    // Truncate right at offRampAddress length byte position
    uint256 truncatePoint = 1 + 8 + 8 + 8 + 1 + 20; // version + selectors + seq + onRampLength + onRampAddress
    bytes memory truncated = new bytes(truncatePoint); // Missing offRampAddress length byte
    for (uint256 i = 0; i < truncated.length; i++) {
      truncated[i] = encoded[i];
    }

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.MESSAGE_OFFRAMP_ADDRESS_LENGTH
      )
    );
    s_helper.decodeMessageV1(truncated);
  }

  function test__decodeMessageV1_RevertWhen_FinalityTruncated() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    bytes memory encoded = s_helper.encodeMessageV1(message);

    // Truncate right before finality field (2 bytes)
    uint256 truncatePoint = 1 + 8 + 8 + 8 + 1 + 20 + 1 + 20; // Up to offRampAddress
    bytes memory truncated = new bytes(truncatePoint + 1); // Only 1 byte of finality instead of 2
    for (uint256 i = 0; i < truncated.length; i++) {
      truncated[i] = encoded[i];
    }

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.MESSAGE_FINALITY
      )
    );
    s_helper.decodeMessageV1(truncated);
  }

  function test__decodeMessageV1_RevertWhen_SenderLengthTruncated() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    bytes memory encoded = s_helper.encodeMessageV1(message);

    // Truncate right before sender length byte
    uint256 truncatePoint = 1 + 8 + 8 + 8 + 1 + 20 + 1 + 20 + 2; // Up to finality
    bytes memory truncated = new bytes(truncatePoint); // Missing sender length
    for (uint256 i = 0; i < truncated.length; i++) {
      truncated[i] = encoded[i];
    }

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.MESSAGE_SENDER_LENGTH
      )
    );
    s_helper.decodeMessageV1(truncated);
  }

  function test__decodeMessageV1_RevertWhen_DataContentTruncated() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    bytes memory encoded = s_helper.encodeMessageV1(message);

    // Truncate the data content at the very end
    bytes memory truncated = new bytes(encoded.length - 1); // Remove last byte of data
    for (uint256 i = 0; i < truncated.length; i++) {
      truncated[i] = encoded[i];
    }

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.MESSAGE_DATA_CONTENT
      )
    );
    s_helper.decodeMessageV1(truncated);
  }

  function test__decodeMessageV1_RevertWhen_FinalOffsetMismatch() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    bytes memory encoded = s_helper.encodeMessageV1(message);

    // Add extra bytes at the end to cause offset mismatch
    bytes memory extended = new bytes(encoded.length + 5);
    for (uint256 i = 0; i < encoded.length; i++) {
      extended[i] = encoded[i];
    }
    // Extra bytes will cause final offset mismatch

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.MESSAGE_FINAL_OFFSET
      )
    );
    s_helper.decodeMessageV1(extended);
  }

  function test__decodeMessageV1_RevertWhen_TokenTransferVersionTruncated() public {
    // Create message with token transfer but corrupt the token version
    MessageV1Codec.TokenTransferV1[] memory tokenTransfers = new MessageV1Codec.TokenTransferV1[](1);
    tokenTransfers[0] = MessageV1Codec.TokenTransferV1({
      amount: 1000,
      sourcePoolAddress: abi.encodePacked(makeAddr("pool")),
      sourceTokenAddress: abi.encodePacked(makeAddr("token")),
      destTokenAddress: abi.encodePacked(makeAddr("destToken")),
      extraData: "test"
    });

    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    message.tokenTransfer = tokenTransfers;
    bytes memory encoded = s_helper.encodeMessageV1(message);

    // Find token transfer section and truncate at version byte
    uint256 tokenVersionOffset = _findTokenTransferVersionOffset(encoded);
    bytes memory truncated = new bytes(tokenVersionOffset); // Cut off at version byte
    for (uint256 i = 0; i < truncated.length; i++) {
      truncated[i] = encoded[i];
    }

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.TOKEN_TRANSFER_VERSION
      )
    );
    s_helper.decodeMessageV1(truncated);
  }

  function test__decodeMessageV1_RevertWhen_TokenTransferAmountTruncated() public {
    // Create message with token transfer
    MessageV1Codec.TokenTransferV1[] memory tokenTransfers = new MessageV1Codec.TokenTransferV1[](1);
    tokenTransfers[0] = MessageV1Codec.TokenTransferV1({
      amount: 1000,
      sourcePoolAddress: abi.encodePacked(makeAddr("pool")),
      sourceTokenAddress: abi.encodePacked(makeAddr("token")),
      destTokenAddress: abi.encodePacked(makeAddr("destToken")),
      extraData: ""
    });

    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    message.tokenTransfer = tokenTransfers;
    bytes memory encoded = s_helper.encodeMessageV1(message);

    // Truncate in the middle of the amount field (32 bytes)
    uint256 tokenVersionOffset = _findTokenTransferVersionOffset(encoded);
    bytes memory truncated = new bytes(tokenVersionOffset + 1 + 16); // Version + partial amount
    for (uint256 i = 0; i < truncated.length; i++) {
      truncated[i] = encoded[i];
    }

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.TOKEN_TRANSFER_AMOUNT
      )
    );
    s_helper.decodeMessageV1(truncated);
  }

  function test__decodeMessageV1_RevertWhen_OffRampAddressContentTruncated() public {
    // Create a valid message and encode it
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    message.onRampAddress = abi.encodePacked(address(0x1234567890123456789012345678901234567890));
    message.offRampAddress = abi.encodePacked(address(0x9876543210987654321098765432109876543210));

    bytes memory encoded = s_helper.encodeMessageV1(message);

    // Find the off-ramp address length byte and truncate after it
    uint256 offset = 1 + 8 + 8 + 8; // version + source + dest + sequence
    offset += 1 + message.onRampAddress.length; // onRamp length + content
    uint8 offRampLength = uint8(encoded[offset]); // off-ramp length byte

    // Truncate the encoded data to include length but not full content
    bytes memory truncated = new bytes(offset + 1 + offRampLength - 5); // -5 to truncate content
    for (uint256 i = 0; i < truncated.length; i++) {
      truncated[i] = encoded[i];
    }

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.MESSAGE_OFFRAMP_ADDRESS_CONTENT
      )
    );
    s_helper.decodeMessageV1(truncated);
  }

  function test__decodeMessageV1_RevertWhen_ReceiverContentTruncated() public {
    // Create a valid message and encode it
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    message.onRampAddress = abi.encodePacked(address(0x1234567890123456789012345678901234567890));
    message.offRampAddress = abi.encodePacked(address(0x9876543210987654321098765432109876543210));
    message.sender = abi.encodePacked(address(0x1111111111111111111111111111111111111111));
    message.receiver = abi.encodePacked(address(0x2222222222222222222222222222222222222222));

    bytes memory encoded = s_helper.encodeMessageV1(message);

    // Find the receiver address length byte and truncate after it
    uint256 offset = 1 + 8 + 8 + 8; // version + source + dest + sequence
    offset += 1 + message.onRampAddress.length; // onRamp length + content
    offset += 1 + message.offRampAddress.length; // offRamp length + content
    offset += 8; // finality
    offset += 1 + message.sender.length; // sender length + content
    // Truncate the encoded data to include length but only partial content
    bytes memory truncated = new bytes(offset + 1 + 10); // Include length + only 10 bytes (less than full 20 bytes)
    for (uint256 i = 0; i < truncated.length; i++) {
      truncated[i] = encoded[i];
    }

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.MESSAGE_RECEIVER_CONTENT
      )
    );
    s_helper.decodeMessageV1(truncated);
  }

  function test__decodeMessageV1_RevertWhen_DestBlobContentTruncated() public {
    // Create a valid message and encode it
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    message.onRampAddress = abi.encodePacked(address(0x1234567890123456789012345678901234567890));
    message.offRampAddress = abi.encodePacked(address(0x9876543210987654321098765432109876543210));
    message.sender = abi.encodePacked(address(0x1111111111111111111111111111111111111111));
    message.receiver = abi.encodePacked(address(0x2222222222222222222222222222222222222222));
    message.destBlob = "test destination blob data";

    bytes memory encoded = s_helper.encodeMessageV1(message);

    // Find the dest blob length bytes and truncate after them
    uint256 offset = 1 + 8 + 8 + 8; // version + source + dest + sequence
    offset += 1 + message.onRampAddress.length; // onRamp length + content
    offset += 1 + message.offRampAddress.length; // offRamp length + content
    offset += 8; // finality
    offset += 1 + message.sender.length; // sender length + content
    offset += 1 + message.receiver.length; // receiver length + content
    // Truncate the encoded data to include length but only partial content
    bytes memory truncated = new bytes(offset + 2 + 10); // Include length + only 10 bytes (less than full blob)
    for (uint256 i = 0; i < truncated.length; i++) {
      truncated[i] = encoded[i];
    }

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.MESSAGE_DEST_BLOB_CONTENT
      )
    );
    s_helper.decodeMessageV1(truncated);
  }

  function _findTokenTransferVersionOffset(
    bytes memory encoded
  ) private pure returns (uint256) {
    // Parse through the encoded message to find the token transfer version byte
    uint256 offset = 1; // Skip message version

    // Skip fixed header fields
    offset += 8 + 8 + 8; // sourceChainSelector + destChainSelector + sequenceNumber

    // Skip onRampAddress
    uint8 onRampLength = uint8(encoded[offset++]);
    offset += onRampLength;

    // Skip offRampAddress
    uint8 offRampLength = uint8(encoded[offset++]);
    offset += offRampLength;

    // Skip finality
    offset += 2;

    // Skip sender
    uint8 senderLength = uint8(encoded[offset++]);
    offset += senderLength;

    // Skip receiver
    uint8 receiverLength = uint8(encoded[offset++]);
    offset += receiverLength;

    // Skip destBlob
    uint16 destBlobLength = uint16(uint8(encoded[offset]) << 8 | uint8(encoded[offset + 1]));
    offset += 2 + destBlobLength;

    // Skip tokenTransfer length
    uint16 tokenTransferLength = uint16(uint8(encoded[offset]) << 8 | uint8(encoded[offset + 1]));
    offset += 2;

    // If there are token transfers, the next byte should be the version
    if (tokenTransferLength > 0) {
      return offset; // This should be the token transfer version byte
    }

    revert("No token transfer found");
  }
}
