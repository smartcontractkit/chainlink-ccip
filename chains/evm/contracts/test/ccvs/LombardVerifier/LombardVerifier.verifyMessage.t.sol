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
  function _encodeCcvData(bytes memory rawPayload, bytes memory proof) internal pure returns (bytes memory) {
    return bytes.concat(
      VERSION_TAG_V1_7_0, bytes2(uint16(rawPayload.length)), rawPayload, bytes2(uint16(proof.length)), proof
    );
  }

  function test_verifyMessage() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), address(12));

    // Proofs are not used. Using raw bytes format.
    bytes memory ccvData = _encodeCcvData("", "");

    vm.startPrank(s_onRamp);

    // This sets the messageId in the mock mailbox to `messageId`.
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");

    vm.startPrank(s_offRamp);

    s_lombardVerifier.verifyMessage(message, messageId, ccvData);
  }

  function test_verifyMessage_RevertWhen_InvalidMessageId() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), address(12));

    vm.startPrank(s_onRamp);

    // This sets the messageId in the mock mailbox to `messageId`.
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");

    vm.startPrank(s_offRamp);

    // Wrong messageId.
    vm.expectRevert(
      abi.encodeWithSelector(LombardVerifier.InvalidMessageId.selector, keccak256("messageId"), messageId)
    );
    s_lombardVerifier.verifyMessage(
      message, keccak256("messageId"), _encodeCcvData(abi.encodePacked("", bytes32(uint256(0x01))), "")
    );
  }

  function test_verifyMessage_RevertWhen_InvalidMessageLength() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), address(12));

    bytes memory shortMessageId = new bytes(20);

    // This sets the messageId in the mock mailbox to `shortMessageId`.
    s_mockMailbox.setMessageId(shortMessageId);

    vm.startPrank(s_offRamp);

    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.InvalidMessageLength.selector, 36, shortMessageId.length));
    s_lombardVerifier.verifyMessage(message, messageId, _encodeCcvData("", ""));
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
    // Make the mailbox fail.
    s_mockMailbox.setShouldSucceed(false);

    vm.startPrank(s_offRamp);

    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.ExecutionError.selector));
    s_lombardVerifier.verifyMessage(_createBasicMessageV1(DEST_CHAIN_SELECTOR), bytes32(0), _encodeCcvData("", ""));
  }

  function test_verifyMessage_RevertWhen_CursedByRMN() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), address(12));

    // verifyMessage checks curse status using message.sourceChainSelector.
    _setMockRMNChainCurse(message.sourceChainSelector, true);

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CursedByRMN.selector, message.sourceChainSelector));
    s_lombardVerifier.verifyMessage(message, messageId, _encodeCcvData("", ""));
  }
}
