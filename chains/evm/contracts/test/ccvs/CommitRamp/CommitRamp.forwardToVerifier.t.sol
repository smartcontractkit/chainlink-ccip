// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {BaseOnRamp} from "../../../ccvs/components/BaseOnRamp.sol";
import {CommitRampSetup} from "./CommitRampSetup.t.sol";

contract CommitRamp_forwardToVerifier is CommitRampSetup {
  function setUp() public override {
    super.setUp();
    vm.stopPrank();
  }

  function test_forwardToVerifier() public {
    bytes memory testData = "test data";

    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createMessageV1(DEST_CHAIN_SELECTOR, msg.sender, testData, msg.sender);

    vm.prank(s_ccvProxy);
    s_commitRamp.forwardToVerifier(
      message, messageId, s_sourceFeeTokens[0], 1000, ""
    );
  }

  function test_forwardToVerifier_RevertWhen_CallerIsNotARampOnRouter() public {
    bytes memory testData = "test data";

    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createMessageV1(DEST_CHAIN_SELECTOR, msg.sender, testData, msg.sender);

    vm.prank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(BaseOnRamp.CallerIsNotARampOnRouter.selector, STRANGER));
    s_commitRamp.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 1000, "");
  }
}
