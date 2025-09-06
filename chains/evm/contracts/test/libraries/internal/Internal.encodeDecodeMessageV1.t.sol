// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";

import {Test} from "forge-std/Test.sol";

// Helper contract to make the args calldata.
contract MessageV1CodecHelper {
  function decodeMessageV1(
    bytes calldata encoded
  ) external pure returns (MessageV1Codec.MessageV1 memory) {
    return MessageV1Codec._decodeMessageV1(encoded);
  }

  function encodeMessageV1(
    MessageV1Codec.MessageV1 memory message
  ) external pure returns (bytes memory) {
    return MessageV1Codec._encodeMessageV1(message);
  }
}

contract MessageV1Codec_encodeDecodeMessageV1 is Test {
  MessageV1CodecHelper private s_messageFormatHelper;

  function setUp() public {
    s_messageFormatHelper = new MessageV1CodecHelper();
  }

  function test_encodeDecodeMessageV1_WithData() public {
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

    MessageV1Codec.MessageV1 memory decodedMessage =
      s_messageFormatHelper.decodeMessageV1(s_messageFormatHelper.encodeMessageV1(originalMessage));

    _assertMessageEqual(originalMessage, decodedMessage);
  }

  function test_encodeDecodeMessageV1_WithTokenTransfer() public {
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

    MessageV1Codec.MessageV1 memory decodedMessage =
      s_messageFormatHelper.decodeMessageV1(s_messageFormatHelper.encodeMessageV1(originalMessage));

    _assertMessageEqual(originalMessage, decodedMessage);
  }

  function test_encodeDecodeMessageV1_MaxLengthFields() public view {
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

    bytes memory encoded = s_messageFormatHelper.encodeMessageV1(originalMessage);
    MessageV1Codec.MessageV1 memory decodedMessage = s_messageFormatHelper.decodeMessageV1(encoded);

    _assertMessageEqual(originalMessage, decodedMessage);
  }

  function test_decodeMessageV1_RevertWhen_InvalidVersion() public {
    // Create a valid encoded message first
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    bytes memory encoded = s_messageFormatHelper.encodeMessageV1(message);

    // Corrupt the version byte (first byte should be 1)
    bytes memory corruptedEncoded = encoded;
    corruptedEncoded[0] = 0x02; // Invalid version

    vm.expectRevert(abi.encodeWithSelector(MessageV1Codec.InvalidEncodingVersion.selector, 2));
    s_messageFormatHelper.decodeMessageV1(corruptedEncoded);
  }

  function test_decodeMessageV1_RevertWhen_InvalidTokenTransferVersion() public {
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

    bytes memory encoded = s_messageFormatHelper.encodeMessageV1(message);

    // Find and corrupt the token transfer version byte
    // The token transfer version should be after the main message fields
    // Look for the pattern where token transfer length is non-zero
    uint256 tokenVersionOffset = _findTokenTransferVersionOffset(encoded);

    bytes memory corruptedEncoded = encoded;
    corruptedEncoded[tokenVersionOffset] = 0x00; // Invalid token version

    vm.expectRevert(abi.encodeWithSelector(MessageV1Codec.InvalidEncodingVersion.selector, 0));
    s_messageFormatHelper.decodeMessageV1(corruptedEncoded);
  }

  function test_decodeMessageV1_RevertWhen_InsufficientDataLength() public {
    // Create a valid encoded message
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    bytes memory encoded = s_messageFormatHelper.encodeMessageV1(message);

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
    s_messageFormatHelper.decodeMessageV1(truncatedEncoded);
  }

  function _createBasicMessage() private pure returns (MessageV1Codec.MessageV1 memory) {
    return MessageV1Codec.MessageV1({
      sourceChainSelector: 1,
      destChainSelector: 2,
      sequenceNumber: 100,
      onRampAddress: abi.encodePacked(address(0x1234567890123456789012345678901234567890)),
      offRampAddress: abi.encodePacked(address(0x0987654321098765432109876543210987654321)),
      finality: 1000,
      sender: abi.encodePacked(address(0x1111111111111111111111111111111111111111)),
      receiver: abi.encodePacked(address(0x2222222222222222222222222222222222222222)),
      destBlob: "test blob",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](0),
      data: "test data"
    });
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

  function _assertMessageEqual(
    MessageV1Codec.MessageV1 memory expected,
    MessageV1Codec.MessageV1 memory actual
  ) private pure {
    assertEq(actual.sourceChainSelector, expected.sourceChainSelector, "sourceChainSelector mismatch");
    assertEq(actual.destChainSelector, expected.destChainSelector, "destChainSelector mismatch");
    assertEq(actual.sequenceNumber, expected.sequenceNumber, "sequenceNumber mismatch");
    assertEq(actual.onRampAddress, expected.onRampAddress, "onRampAddress mismatch");
    assertEq(actual.offRampAddress, expected.offRampAddress, "offRampAddress mismatch");
    assertEq(actual.finality, expected.finality, "finality mismatch");
    assertEq(actual.sender, expected.sender, "sender mismatch");
    assertEq(actual.receiver, expected.receiver, "receiver mismatch");
    assertEq(actual.destBlob, expected.destBlob, "destBlob mismatch");
    assertEq(actual.data, expected.data, "data mismatch");

    assertEq(actual.tokenTransfer.length, expected.tokenTransfer.length, "tokenTransfer length mismatch");
    for (uint256 i = 0; i < expected.tokenTransfer.length; ++i) {
      _assertTokenTransferEqual(expected.tokenTransfer[i], actual.tokenTransfer[i], i);
    }
  }

  function _assertTokenTransferEqual(
    MessageV1Codec.TokenTransferV1 memory expected,
    MessageV1Codec.TokenTransferV1 memory actual,
    uint256 index
  ) private pure {
    string memory indexStr = vm.toString(index);
    assertEq(actual.amount, expected.amount, string(abi.encodePacked("tokenTransfer[", indexStr, "].amount mismatch")));
    assertEq(
      actual.sourcePoolAddress,
      expected.sourcePoolAddress,
      string(abi.encodePacked("tokenTransfer[", indexStr, "].sourcePoolAddress mismatch"))
    );
    assertEq(
      actual.sourceTokenAddress,
      expected.sourceTokenAddress,
      string(abi.encodePacked("tokenTransfer[", indexStr, "].sourceTokenAddress mismatch"))
    );
    assertEq(
      actual.destTokenAddress,
      expected.destTokenAddress,
      string(abi.encodePacked("tokenTransfer[", indexStr, "].destTokenAddress mismatch"))
    );
    assertEq(
      actual.extraData, expected.extraData, string(abi.encodePacked("tokenTransfer[", indexStr, "].extraData mismatch"))
    );
  }
}
