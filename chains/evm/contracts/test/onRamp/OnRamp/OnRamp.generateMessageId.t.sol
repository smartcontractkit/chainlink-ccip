// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

contract OnRamp_generateMessageId is OnRampSetup {
  function test_generateMessageId() public view {
    Internal.EVM2AnyRampMessage memory message =
      _messageToEvent(_generateSingleTokenMessage(s_sourceTokens[0], 1000e18), 1, 1, 100e18, OWNER);

    bytes32 expectedMessageId = message.header.messageId;
    message.header.messageId = "";

    bytes32 messageId = s_onRamp.generateMessageId(message);

    assertEq(expectedMessageId, messageId);
  }

  function test_generateMessageId_SameMessageProducesSameId() public {
    Internal.EVM2AnyRampMessage memory message = _messageToEvent(_generateEmptyMessage(), 1, 1, 100e18, OWNER);
    message.header.messageId = "";

    bytes32 messageId1 = s_onRamp.generateMessageId(message);

    // MessageId should remain the same when calculated in a different block
    vm.roll(block.number + 100);
    vm.warp(block.timestamp + 100);

    bytes32 messageId2 = s_onRamp.generateMessageId(message);

    assertEq(messageId1, messageId2);
  }

  // Reverts

  function test_generateMessageId_RevertWhen_MessageIdAlreadySet() public {
    Internal.EVM2AnyRampMessage memory message = _messageToEvent(_generateEmptyMessage(), 1, 1, 100e18, OWNER);

    vm.expectRevert(abi.encodeWithSelector(OnRamp.MessageIdUnexpectedlySet.selector, message.header.messageId));
    s_onRamp.generateMessageId(message);
  }
}
