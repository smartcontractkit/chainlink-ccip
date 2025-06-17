// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

contract OnRamp_generateMessageId is OnRampSetup {
  function _getMetadataHash() internal view returns (bytes32) {
    return keccak256(
      abi.encode(Internal.EVM_2_ANY_MESSAGE_HASH, SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, address(s_onRamp))
    );
  }

  function test_BasicMessageId() public {
    Internal.EVM2AnyRampMessage memory message = _messageToEvent(_generateEmptyMessage(), 1, 1, 100e18, OWNER);

    bytes32 expectedMessageId = message.header.messageId;
    message.header.messageId = "";

    bytes32 messageId = s_onRamp.generateMessageId(message);

    assertEq(messageId, expectedMessageId);
  }

  function test_TokenMessageId() public {
    Internal.EVM2AnyRampMessage memory message =
      _messageToEvent(_generateSingleTokenMessage(s_sourceTokens[0], 1000e18), 1, 1, 100e18, OWNER);

    bytes32 expectedMessageId = message.header.messageId;
    message.header.messageId = "";

    bytes32 messageId = s_onRamp.generateMessageId(message);

    assertEq(messageId, expectedMessageId);
  }

  function test_SameMessageProducesSameId() public {
    Internal.EVM2AnyRampMessage memory message = _messageToEvent(_generateEmptyMessage(), 1, 1, 100e18, OWNER);
    message.header.messageId = "";

    bytes32 messageId1 = s_onRamp.generateMessageId(message);
    bytes32 messageId2 = s_onRamp.generateMessageId(message);

    assertEq(messageId1, messageId2);
  }

  function test_DifferentMessageProduceDifferentIds() public {
    Internal.EVM2AnyRampMessage memory message1 = _messageToEvent(_generateEmptyMessage(), 1, 1, 100e18, OWNER);
    message1.header.messageId = "";

    Internal.EVM2AnyRampMessage memory message2 = _messageToEvent(_generateEmptyMessage(), 2, 2, 100e18, OWNER);
    message2.header.messageId = "";

    bytes32 messageId1 = s_onRamp.generateMessageId(message1);
    bytes32 messageId2 = s_onRamp.generateMessageId(message2);

    assertTrue(messageId1 != messageId2);
  }

  // Reverts

  function test_RevertWhen_MessageIdAlreadySet() public {
    Internal.EVM2AnyRampMessage memory message = _messageToEvent(_generateEmptyMessage(), 1, 1, 100e18, OWNER);

    vm.expectRevert(abi.encodeWithSelector(OnRamp.MessageIdUnexpectedlySet.selector, message.header.messageId));
    s_onRamp.generateMessageId(message);
  }
}
