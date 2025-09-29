// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../../../interfaces/ICrossChainVerifierV1.sol";

import {VerifierProxy} from "../../../ccvs/VerifierProxy.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {CommitteeVerifierSetup} from "./CommitteeVerifierSetup.t.sol";

contract CommitteeVerifier_forwardToVerifier is CommitteeVerifierSetup {
  function setUp() public override {
    super.setUp();
    vm.stopPrank();
  }

  function test_forwardToVerifier() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _generateBasicMessageV1();

    vm.prank(s_ccvProxy);
    s_committeeVerifier.forwardToVerifier(s_ccvProxy, message, messageId, s_sourceFeeTokens[0], 1000, "");
  }

  function test_forwardToVerifier_ViaVerifierProxy() public {
    VerifierProxy verifierProxy = new VerifierProxy(address(s_committeeVerifier));

    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _generateBasicMessageV1();

    vm.prank(s_ccvProxy);
    ICrossChainVerifierV1(address(verifierProxy)).forwardToVerifier(
      s_ccvProxy, message, messageId, s_sourceFeeTokens[0], 1000, ""
    );
  }

  function test_forwardToVerifier_RevertWhen_CallerIsNotARampOnRouter() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _generateBasicMessageV1();

    vm.prank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CallerIsNotARampOnRouter.selector, STRANGER));
    s_committeeVerifier.forwardToVerifier(STRANGER, message, messageId, s_sourceFeeTokens[0], 1000, "");
  }
}
