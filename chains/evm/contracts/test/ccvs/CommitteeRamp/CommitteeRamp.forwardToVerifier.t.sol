// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICCVRampV1} from "../../../interfaces/ICCVRampV1.sol";

import {RampProxy} from "../../../ccvs/RampProxy.sol";
import {BaseOnRamp} from "../../../ccvs/components/BaseOnRamp.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {CommitteeRampSetup} from "./CommitteeRampSetup.t.sol";

contract CommitteeRamp_forwardToVerifier is CommitteeRampSetup {
  function setUp() public override {
    super.setUp();
    vm.stopPrank();
  }

  function test_forwardToVerifier() public {
    bytes memory testData = "test data";

    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _generateBasicMessageV1();

    vm.prank(s_ccvProxy);
    s_commitRamp.forwardToVerifier(s_ccvProxy, message, messageId, s_sourceFeeTokens[0], 1000, "");
  }

  function test_forwardToVerifier_ViaRampProxy() public {
    bytes memory testData = "test data";
    RampProxy rampProxy = new RampProxy(address(s_commitRamp));

    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _generateBasicMessageV1();

    vm.prank(s_ccvProxy);
    ICCVRampV1(address(rampProxy)).forwardToVerifier(s_ccvProxy, message, messageId, s_sourceFeeTokens[0], 1000, "");
  }

  function test_forwardToVerifier_RevertWhen_CallerIsNotARampOnRouter() public {
    bytes memory testData = "test data";

    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _generateBasicMessageV1();

    vm.prank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(BaseOnRamp.CallerIsNotARampOnRouter.selector, STRANGER));
    s_commitRamp.forwardToVerifier(STRANGER, message, messageId, s_sourceFeeTokens[0], 1000, "");
  }
}
