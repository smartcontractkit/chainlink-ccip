// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";

import {LombardVerifier} from "../../../ccvs/LombardVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {LombardVerifierSetup} from "./LombardVerifierSetup.t.sol";

contract LombardVerifier_verifyMessage is LombardVerifierSetup {
  /// @dev Encodes ccvData in the raw bytes format:
  /// [versionTag (4 bytes)][rawPayloadLength (2 bytes)][rawPayload][proofLength (2 bytes)][proof]
  function _encodeCcvData(
    bytes memory rawPayload,
    bytes memory proof
  ) internal pure returns (bytes memory) {
    return bytes.concat(
      VERSION_TAG_V1_7_0, bytes2(uint16(rawPayload.length)), rawPayload, bytes2(uint16(proof.length)), proof
    );
  }

  function test_verifyMessage() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), address(12));

    // Generate a valid rawPayload that matches the message token transfer data.
    bytes memory rawPayload = _generateValidRawPayload(
      message.tokenTransfer[0].destTokenAddress, message.tokenTransfer[0].tokenReceiver, message.tokenTransfer[0].amount
    );

    // Proofs are not used. Using raw bytes format.
    bytes memory ccvData = _encodeCcvData(rawPayload, "");

    vm.startPrank(s_onRamp);

    // This sets the messageId in the mock mailbox to `messageId`.
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");

    vm.startPrank(s_offRamp);

    s_lombardVerifier.verifyMessage(message, messageId, ccvData);
  }

  function test_verifyMessage_RevertWhen_InvalidMessageId() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), address(12));

    bytes32 wrongMessageId = keccak256("messageId");
    bytes memory rawPayload = _generateValidRawPayload(
      message.tokenTransfer[0].destTokenAddress, message.tokenTransfer[0].tokenReceiver, message.tokenTransfer[0].amount
    );

    vm.startPrank(s_onRamp);

    // This sets the messageId in the mock mailbox to `wrongMessageId`.
    s_lombardVerifier.forwardToVerifier(message, wrongMessageId, address(0), 0, "");

    vm.startPrank(s_offRamp);

    // Wrong messageId.
    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.InvalidMessageId.selector, messageId, wrongMessageId));
    s_lombardVerifier.verifyMessage(message, messageId, _encodeCcvData(rawPayload, ""));
  }

  function test_verifyMessage_RevertWhen_InvalidMessageLength() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), address(12));

    bytes memory rawPayload = _generateValidRawPayload(
      message.tokenTransfer[0].destTokenAddress, message.tokenTransfer[0].tokenReceiver, message.tokenTransfer[0].amount
    );

    bytes memory shortMessageId = new bytes(20);

    // This sets the messageId in the mock mailbox to `shortMessageId`.
    s_mockMailbox.setMessageId(shortMessageId);

    vm.startPrank(s_offRamp);

    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.InvalidMessageLength.selector, 36, shortMessageId.length));
    s_lombardVerifier.verifyMessage(message, messageId, _encodeCcvData(rawPayload, ""));
  }

  function test_verifyMessage_RevertWhen_CallerIsNotOffRamp() public {
    address invalidCaller = makeAddr("invalidCaller");

    // Mock the router to return false for isOffRamp with the invalid caller.
    vm.mockCall(
      address(s_router), abi.encodeCall(IRouter.isOffRamp, (DEST_CHAIN_SELECTOR, invalidCaller)), abi.encode(false)
    );

    MessageV1Codec.MessageV1 memory message = _createBasicMessageV1(DEST_CHAIN_SELECTOR);

    vm.startPrank(invalidCaller);

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CallerIsNotARampOnRouter.selector, invalidCaller));
    // Empty ccvData still needs valid format for parsing (though it will fail before parsing).
    s_lombardVerifier.verifyMessage(message, bytes32(0), _encodeCcvData("", ""));
  }

  function test_verifyMessage_RevertWhen_ExecutionError() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), address(12));

    bytes memory rawPayload = _generateValidRawPayload(
      message.tokenTransfer[0].destTokenAddress, message.tokenTransfer[0].tokenReceiver, message.tokenTransfer[0].amount
    );

    // Make the mailbox fail.
    s_mockMailbox.setShouldSucceed(false);

    vm.startPrank(s_offRamp);

    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.ExecutionError.selector));
    s_lombardVerifier.verifyMessage(message, messageId, _encodeCcvData(rawPayload, ""));
  }

  function test_verifyMessage_RevertWhen_CursedByRMN() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), address(12));

    // verifyMessage checks curse status using message.sourceChainSelector.
    _setMockRMNChainCurse(message.sourceChainSelector, true);

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CursedByRMN.selector, message.sourceChainSelector));
    s_lombardVerifier.verifyMessage(message, messageId, _encodeCcvData("", ""));
  }

  function test_verifyMessage_RevertWhen_InvalidVerifierResults_CcvDataTooShortForPayloadLengthField() public {
    vm.startPrank(s_offRamp);

    // ccvData with only 5 bytes (needs at least 6: 4 for version tag + 2 for rawPayloadLength).
    bytes memory tooShortCcvData = bytes.concat(VERSION_TAG_V1_7_0, bytes1(0x00));

    vm.expectRevert(LombardVerifier.InvalidVerifierResults.selector);
    s_lombardVerifier.verifyMessage(_createBasicMessageV1(DEST_CHAIN_SELECTOR), bytes32(0), tooShortCcvData);
  }

  function test_verifyMessage_RevertWhen_InvalidVerifierResults_CcvDataTooShortForProofLengthField() public {
    vm.startPrank(s_offRamp);

    // ccvData with version tag (4) + rawPayloadLength (2) claiming 10 bytes of raw payload,
    // but only providing 5 bytes of raw payload and no proof length field.
    // Total: 4 + 2 + 5 = 11 bytes, but needs at least 4 + 2 + 10 + 2 = 18 bytes.
    bytes memory tooShortCcvData = bytes.concat(
      VERSION_TAG_V1_7_0,
      bytes2(uint16(10)), // rawPayloadLength = 10
      bytes5(0) // only 5 bytes instead of 10 + 2 for proof length
    );

    vm.expectRevert(LombardVerifier.InvalidVerifierResults.selector);
    s_lombardVerifier.verifyMessage(_createBasicMessageV1(DEST_CHAIN_SELECTOR), bytes32(0), tooShortCcvData);
  }

  function test_verifyMessage_RevertWhen_InvalidVerifierResults_CcvDataTooShortForProof() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), address(12));

    bytes memory rawPayload = _generateValidRawPayload(
      message.tokenTransfer[0].destTokenAddress, message.tokenTransfer[0].tokenReceiver, message.tokenTransfer[0].amount
    );

    vm.startPrank(s_offRamp);

    // ccvData with version tag (4) + rawPayloadLength (2) + rawPayload (variable) + proofLength (2) claiming 10 bytes,
    // but only providing 5 bytes of proof.
    bytes memory tooShortCcvData = bytes.concat(
      VERSION_TAG_V1_7_0,
      bytes2(uint16(rawPayload.length)), // rawPayloadLength
      rawPayload,
      bytes2(uint16(10)), // proofLength = 10
      bytes5(0) // only 5 bytes instead of 10
    );

    vm.expectRevert(LombardVerifier.InvalidVerifierResults.selector);
    s_lombardVerifier.verifyMessage(message, messageId, tooShortCcvData);
  }

  function test_verifyMessage_RevertWhen_InvalidToken() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), address(12));

    // Generate a rawPayload with a different (invalid) token address.
    bytes memory invalidToken = abi.encodePacked(makeAddr("wrongToken"));
    bytes memory rawPayload =
      _generateValidRawPayload(invalidToken, message.tokenTransfer[0].tokenReceiver, message.tokenTransfer[0].amount);

    bytes memory ccvData = _encodeCcvData(rawPayload, "");

    vm.startPrank(s_onRamp);
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");

    vm.startPrank(s_offRamp);

    // The error will be InvalidMessageId because the contract reuses this error for token mismatches.
    vm.expectRevert(
      abi.encodeWithSelector(
        LombardVerifier.InvalidMessageId.selector,
        bytes32(message.tokenTransfer[0].destTokenAddress),
        bytes32(invalidToken)
      )
    );
    s_lombardVerifier.verifyMessage(message, messageId, ccvData);
  }

  function test_verifyMessage_RevertWhen_InvalidReceiver() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), address(12));

    // Generate a rawPayload with a different (invalid) receiver address.
    bytes memory invalidReceiver = abi.encodePacked(address(999));
    bytes memory rawPayload = _generateValidRawPayload(
      message.tokenTransfer[0].destTokenAddress, invalidReceiver, message.tokenTransfer[0].amount
    );

    bytes memory ccvData = _encodeCcvData(rawPayload, "");

    vm.startPrank(s_onRamp);
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");

    vm.startPrank(s_offRamp);

    vm.expectRevert(
      abi.encodeWithSelector(LombardVerifier.InvalidReceiver.selector, message.tokenTransfer[0].tokenReceiver)
    );
    s_lombardVerifier.verifyMessage(message, messageId, ccvData);
  }

  function test_verifyMessage_RevertWhen_InvalidAmount() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), address(12));

    // Generate a rawPayload with a different (invalid) amount.
    uint256 invalidAmount = message.tokenTransfer[0].amount + 100;
    bytes memory rawPayload = _generateValidRawPayload(
      message.tokenTransfer[0].destTokenAddress, message.tokenTransfer[0].tokenReceiver, invalidAmount
    );

    bytes memory ccvData = _encodeCcvData(rawPayload, "");

    vm.startPrank(s_onRamp);
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");

    vm.startPrank(s_offRamp);

    // The error will be InvalidMessageId because the contract reuses this error for amount mismatches.
    vm.expectRevert(
      abi.encodeWithSelector(
        LombardVerifier.InvalidMessageId.selector, bytes32(message.tokenTransfer[0].amount), bytes32(invalidAmount)
      )
    );
    s_lombardVerifier.verifyMessage(message, messageId, ccvData);
  }
}
