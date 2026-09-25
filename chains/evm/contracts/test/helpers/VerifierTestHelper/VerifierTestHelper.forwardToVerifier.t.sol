// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {VerifierTestHelper} from "../VerifierTestHelper.sol";
import {VerifierTestHelperSetup} from "./VerifierTestHelperSetup.t.sol";

contract VerifierTestHelper_forwardToVerifier is VerifierTestHelperSetup {
  function test_forwardToVerifier() public {
    _configureRemoteChain(DEST_CHAIN_SELECTOR);
    _allowSender(DEST_CHAIN_SELECTOR, s_sender);
    MessageV1Codec.MessageV1 memory message = _createMessage(
      SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, s_verifier.getTestToken(), address(1), s_sender, s_receiver
    );

    vm.prank(s_onRamp);
    bytes memory returnData = s_verifier.forwardToVerifier(message, bytes32(0), address(0), 0, "");

    assertEq(returnData, abi.encodePacked(VERSION_TAG));
  }

  function test_forwardToVerifier_RevertWhen_MessageCannotHaveSideEffects() public {
    _configureRemoteChain(DEST_CHAIN_SELECTOR);
    _allowSender(DEST_CHAIN_SELECTOR, s_sender);
    MessageV1Codec.MessageV1 memory message = _createMessage(
      SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, s_verifier.getTestToken(), address(1), s_sender, s_receiver
    );
    message.data = "side effect";

    vm.prank(s_onRamp);
    vm.expectRevert(VerifierTestHelper.MessageCannotHaveSideEffects.selector);
    s_verifier.forwardToVerifier(message, bytes32(0), address(0), 0, "");
  }

  function test_forwardToVerifier_RevertWhen_MustUseTestToken() public {
    _configureRemoteChain(DEST_CHAIN_SELECTOR);
    _allowSender(DEST_CHAIN_SELECTOR, s_sender);
    MessageV1Codec.MessageV1 memory message =
      _createMessage(SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, address(1), address(1), s_sender, s_receiver);
    address testToken = s_verifier.getTestToken();

    vm.prank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(VerifierTestHelper.MustUseTestToken.selector, testToken));
    s_verifier.forwardToVerifier(message, bytes32(0), address(0), 0, "");
  }

  function test_forwardToVerifier_RevertWhen_SenderNotAllowed() public {
    _configureRemoteChain(DEST_CHAIN_SELECTOR);
    _allowSender(DEST_CHAIN_SELECTOR, s_sender);
    address notAllowedSender = makeAddr("notAllowedSender");
    MessageV1Codec.MessageV1 memory message = _createMessage(
      SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, s_verifier.getTestToken(), address(1), notAllowedSender, s_receiver
    );

    vm.prank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.SenderNotAllowed.selector, notAllowedSender));
    s_verifier.forwardToVerifier(message, bytes32(0), address(0), 0, "");
  }
}
