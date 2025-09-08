// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {MessageV1CodecSetup} from "./MessageV1CodecSetup.t.sol";

contract MessageV1Codec__encodeMessageV1 is MessageV1CodecSetup {
  function test__encodeMessageV1_WithData() public {
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

    bytes memory encoded = s_helper.encodeMessageV1(originalMessage);
    MessageV1Codec.MessageV1 memory decodedMessage = s_helper.decodeMessageV1(encoded);

    _assertMessageEqual(originalMessage, decodedMessage);
  }

  function test__encodeMessageV1_WithTokenTransfer() public {
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

    bytes memory encoded = s_helper.encodeMessageV1(originalMessage);
    MessageV1Codec.MessageV1 memory decodedMessage = s_helper.decodeMessageV1(encoded);

    _assertMessageEqual(originalMessage, decodedMessage);
  }

  function test__encodeMessageV1_MaxLengthFields() public view {
    uint256 testDataLength = 1000; // Reasonable size for testing
    // Create maximum length fields to test boundary conditions.
    bytes memory maxLengthBytes = new bytes(type(uint8).max); // 255 bytes
    bytes memory maxLengthData = new bytes(testDataLength);

    // Fill with some pattern for verification.
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

  function test__encodeMessageV1_EmptyFields() public view {
    MessageV1Codec.MessageV1 memory originalMessage = MessageV1Codec.MessageV1({
      sourceChainSelector: 0,
      destChainSelector: 0,
      sequenceNumber: 0,
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

  function test__encodeMessageV1_RevertWhen_OnRampAddressTooLong() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    message.onRampAddress = new bytes(uint256(type(uint8).max) + 1); // uint8 max + 1

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.ENCODE_ONRAMP_ADDRESS_LENGTH
      )
    );
    s_helper.encodeMessageV1(message);
  }

  function test__encodeMessageV1_RevertWhen_OffRampAddressTooLong() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    message.offRampAddress = new bytes(uint256(type(uint8).max) + 1); // uint8 max + 1

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.ENCODE_OFFRAMP_ADDRESS_LENGTH
      )
    );
    s_helper.encodeMessageV1(message);
  }

  function test__encodeMessageV1_RevertWhen_SenderAddressTooLong() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    message.sender = new bytes(uint256(type(uint8).max) + 1);

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.ENCODE_SENDER_LENGTH
      )
    );
    s_helper.encodeMessageV1(message);
  }

  function test__encodeMessageV1_RevertWhen_ReceiverAddressTooLong() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    message.receiver = new bytes(uint256(type(uint8).max) + 1);

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.ENCODE_RECEIVER_LENGTH
      )
    );
    s_helper.encodeMessageV1(message);
  }

  function test__encodeMessageV1_RevertWhen_DestBlobTooLong() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    message.destBlob = new bytes(uint256(type(uint16).max) + 1);

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.ENCODE_DEST_BLOB_LENGTH
      )
    );
    s_helper.encodeMessageV1(message);
  }

  function test__encodeMessageV1_RevertWhen_DataTooLong() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();
    message.data = new bytes(uint256(type(uint16).max) + 1);

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.ENCODE_DATA_LENGTH
      )
    );
    s_helper.encodeMessageV1(message);
  }

  function test__encodeMessageV1_RevertWhen_TokenTransferArrayTooLong() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessage();

    // Create array with MAX_NUMBER_OF_TOKENS token transfers.
    message.tokenTransfer = new MessageV1Codec.TokenTransferV1[](MessageV1Codec.MAX_NUMBER_OF_TOKENS + 1);

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector,
        MessageV1Codec.EncodingErrorLocation.ENCODE_TOKEN_TRANSFER_ARRAY_LENGTH
      )
    );
    s_helper.encodeMessageV1(message);
  }
}
