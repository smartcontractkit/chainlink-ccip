// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";

contract CommitOnRamp_forwardToVerifier is CommitOnRampSetup {
  function setUp() public override {
    super.setUp();
    vm.stopPrank();
  }

  function test_forwardToVerifier() public {
    bytes memory testData = "test data";

    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createMessageV1(DEST_CHAIN_SELECTOR, msg.sender, testData, msg.sender);

    vm.prank(s_ccvProxy);
    s_commitOnRamp.forwardToVerifier(
      DEST_CHAIN_SELECTOR, s_ccvProxy, message, messageId, s_sourceFeeTokens[0], 1000, ""
    );
  }

  function test_forwardToVerifier_RevertWhen_CallerIsNotARampOnRouter() public {
    bytes memory testData = "test data";

    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createMessageV1(DEST_CHAIN_SELECTOR, msg.sender, testData, msg.sender);

    vm.prank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(BaseOnRamp.CallerIsNotARampOnRouter.selector, STRANGER));
    s_commitOnRamp.forwardToVerifier(DEST_CHAIN_SELECTOR, STRANGER, message, messageId, s_sourceFeeTokens[0], 1000, "");
  }
}
