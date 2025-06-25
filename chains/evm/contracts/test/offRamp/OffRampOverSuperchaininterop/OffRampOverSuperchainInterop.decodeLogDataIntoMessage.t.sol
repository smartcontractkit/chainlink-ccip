// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {SuperchainInterop} from "../../../libraries/SuperchainInterop.sol";
import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";
import {OffRampOverSuperchainInteropSetup} from "./OffRampOverSuperchainInteropSetup.t.sol";
contract OffRampOverSuperchainInterop_decodeLogDataIntoMessage is OffRampOverSuperchainInteropSetup {

  function test_decodeLogDataIntoMessage() public {
    // Create a valid message
    uint64 sequenceNumber = 100;
    Internal.Any2EVMRampMessage memory originalMessage = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1, ON_RAMP_ADDRESS_1, sequenceNumber, 1, new Client.EVMTokenAmount[](0), true
    );

    // Encode log data
    bytes memory logData = _encodeLogData(DEST_CHAIN_SELECTOR, sequenceNumber, originalMessage);

    // Decode and verify
    Internal.Any2EVMRampMessage memory decodedMessage = s_offRampHelper.decodeLogDataIntoMessageExposed(logData);

    // Verify all fields match
    assertEq(decodedMessage.header.messageId, originalMessage.header.messageId);
    assertEq(decodedMessage.header.sourceChainSelector, originalMessage.header.sourceChainSelector);
    assertEq(decodedMessage.header.destChainSelector, originalMessage.header.destChainSelector);
    assertEq(decodedMessage.header.sequenceNumber, originalMessage.header.sequenceNumber);
    assertEq(decodedMessage.header.nonce, originalMessage.header.nonce);
    assertEq(decodedMessage.sender, originalMessage.sender);
    assertEq(decodedMessage.data, originalMessage.data);
    assertEq(decodedMessage.receiver, originalMessage.receiver);
    assertEq(decodedMessage.gasLimit, originalMessage.gasLimit);
    assertEq(decodedMessage.tokenAmounts.length, originalMessage.tokenAmounts.length);
  }

  function test_decodeLogDataIntoMessage_RevertWhen_InvalidSelector() public {
    // Create a valid message
    Internal.Any2EVMRampMessage memory message =
      _generateAny2EVMMessage(SOURCE_CHAIN_SELECTOR_1, ON_RAMP_ADDRESS_1, 100, 1, new Client.EVMTokenAmount[](0), true);

    // Encode log data with wrong selector
    bytes32 wrongSelector = keccak256("WrongSelector");
    bytes memory logData = abi.encodePacked(
      wrongSelector, abi.encode(DEST_CHAIN_SELECTOR, message.header.sequenceNumber), abi.encode(message)
    );

    // Should revert
    vm.expectRevert(
      abi.encodeWithSelector(OffRampOverSuperchainInterop.InvalidInteropLogSelector.selector, wrongSelector)
    );
    s_offRampHelper.decodeLogDataIntoMessageExposed(logData);
  }

  function test_decodeLogDataIntoMessage_RevertWhen_DestChainSelectorMismatch() public {
    // Create a message with wrong dest chain selector
    uint64 sequenceNumber = 100;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1, ON_RAMP_ADDRESS_1, sequenceNumber, 1, new Client.EVMTokenAmount[](0), true
    );

    uint64 wrongDestChainSelector = DEST_CHAIN_SELECTOR + 1;
    message.header.destChainSelector = wrongDestChainSelector;

    // Encode log data
    bytes memory logData = _encodeLogData(
      DEST_CHAIN_SELECTOR, // Correct selector in log data
      sequenceNumber,
      message // Message with wrong selector
    );

    // Should revert
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.MismatchedDestChainSelector.selector, DEST_CHAIN_SELECTOR, wrongDestChainSelector
      )
    );
    s_offRampHelper.decodeLogDataIntoMessageExposed(logData);
  }

  function test_decodeLogDataIntoMessage_RevertWhen_SequenceNumberMismatch() public {
    // Create a message
    uint64 sequenceNumber = 100;
    uint64 wrongSequenceNumber = 200;
    Internal.Any2EVMRampMessage memory message = _generateAny2EVMMessage(
      SOURCE_CHAIN_SELECTOR_1,
      ON_RAMP_ADDRESS_1,
      wrongSequenceNumber, // Wrong sequence number in message header
      1,
      new Client.EVMTokenAmount[](0),
      true
    );

    // Encode log data with different sequence number
    bytes memory logData = _encodeLogData(
      DEST_CHAIN_SELECTOR,
      sequenceNumber, // Different sequence number in log data
      message
    );

    // Should revert
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRampOverSuperchainInterop.MismatchedSequenceNumber.selector, sequenceNumber, wrongSequenceNumber
      )
    );
    s_offRampHelper.decodeLogDataIntoMessageExposed(logData);
  }

  function test_decodeLogDataIntoMessage_FuzzMessageSizes(
    uint256 dataLength,
    uint256 tokenCount,
    uint64 sequenceNumber,
    uint64 nonce
  ) public {
    // Bound inputs
    dataLength = bound(dataLength, 0, 1000);
    tokenCount = bound(tokenCount, 0, 10);
    sequenceNumber = uint64(bound(sequenceNumber, 1, type(uint64).max));
    nonce = uint64(bound(nonce, 0, type(uint64).max));

    // Create random data
    bytes memory data = new bytes(dataLength);
    for (uint256 i = 0; i < dataLength; i++) {
      data[i] = bytes1(uint8(uint256(keccak256(abi.encode(i))) % 256));
    }

    // Create token amounts
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](tokenCount);
    for (uint256 i = 0; i < tokenCount; i++) {
      tokenAmounts[i] = Client.EVMTokenAmount({
        token: makeAddr(string(abi.encodePacked("token", i))),
        amount: uint256(keccak256(abi.encode("amount", i))) % 1e18
      });
    }

    // Create message
    Internal.Any2EVMRampMessage memory originalMessage = Internal.Any2EVMRampMessage({
      header: Internal.RampMessageHeader({
        messageId: keccak256(abi.encode(sequenceNumber, nonce)),
        sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
        destChainSelector: DEST_CHAIN_SELECTOR,
        sequenceNumber: sequenceNumber,
        nonce: nonce
      }),
      sender: abi.encode(makeAddr("sender")),
      data: data,
      receiver: makeAddr("receiver"),
      gasLimit: 200_000,
      tokenAmounts: _convertToInternal(tokenAmounts)
    });

    // Encode log data
    bytes memory logData = _encodeLogData(DEST_CHAIN_SELECTOR, sequenceNumber, originalMessage);

    // Decode and verify
    Internal.Any2EVMRampMessage memory decodedMessage = s_offRampHelper.decodeLogDataIntoMessageExposed(logData);

    // Verify decoded correctly
    assertEq(decodedMessage.header.messageId, originalMessage.header.messageId);
    assertEq(decodedMessage.header.sequenceNumber, sequenceNumber);
    assertEq(decodedMessage.header.nonce, nonce);
    assertEq(decodedMessage.data, data);
    assertEq(decodedMessage.tokenAmounts.length, tokenCount);
  }

  function test_decodeLogDataIntoMessage_RevertWhen_InvalidLogDataLength() public {
    // Create log data that's too short
    bytes memory tooShortLogData = new bytes(32); // Only selector, missing required fields

    // Fill with the correct selector
    bytes32 selector = SuperchainInterop.SENT_MESSAGE_LOG_SELECTOR;
    assembly {
      mstore(add(tooShortLogData, 32), selector)
    }

    // Should revert with abi decoding error
    vm.expectRevert();
    s_offRampHelper.decodeLogDataIntoMessageExposed(tooShortLogData);
  }

}
