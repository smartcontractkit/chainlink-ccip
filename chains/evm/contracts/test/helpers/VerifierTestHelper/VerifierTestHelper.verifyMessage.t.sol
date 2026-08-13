// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {VerifierTestHelper} from "../VerifierTestHelper.sol";
import {VerifierTestHelperSetup} from "./VerifierTestHelperSetup.t.sol";

contract VerifierTestHelper_verifyMessage is VerifierTestHelperSetup {
  function test_verifyMessage() public {
    _configureRemoteChain(SOURCE_CHAIN_SELECTOR);
    _allowSender(SOURCE_CHAIN_SELECTOR, s_receiver);
    _mockOffRampSourceChainConfig(SOURCE_CHAIN_SELECTOR, s_testRouter);
    MessageV1Codec.MessageV1 memory message =
      _createMessage(SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, address(1), s_testToken, s_sender, s_receiver);

    vm.prank(s_offRamp);
    s_verifier.verifyMessage(message, bytes32(0), "");
  }

  function test_verifyMessage_RevertWhen_DestTokenIsNotTestToken() public {
    _configureRemoteChain(SOURCE_CHAIN_SELECTOR);
    _mockOffRampSourceChainConfig(SOURCE_CHAIN_SELECTOR, s_testRouter);
    MessageV1Codec.MessageV1 memory message =
      _createMessage(SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, address(1), address(1), s_sender, s_receiver);

    vm.prank(s_offRamp);
    vm.expectRevert(abi.encodeWithSelector(VerifierTestHelper.MustUseTestToken.selector, s_testToken));
    s_verifier.verifyMessage(message, bytes32(0), "");
  }

  function test_verifyMessage_RevertWhen_OffRampRouterIsNotTestRouter() public {
    _configureRemoteChain(SOURCE_CHAIN_SELECTOR);
    _mockOffRampSourceChainConfig(SOURCE_CHAIN_SELECTOR, makeAddr("wrongRouter"));
    MessageV1Codec.MessageV1 memory message =
      _createMessage(SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, address(1), s_testToken, s_sender, s_receiver);

    vm.prank(s_offRamp);
    vm.expectRevert(VerifierTestHelper.MustUseTestRouter.selector);
    s_verifier.verifyMessage(message, bytes32(0), "");
  }

  function test_verifyMessage_RevertWhen_ReceiverIsNotAllowed() public {
    _configureRemoteChain(SOURCE_CHAIN_SELECTOR);
    _mockOffRampSourceChainConfig(SOURCE_CHAIN_SELECTOR, s_testRouter);
    address allowedReceiver = makeAddr("allowedReceiver");
    _allowSender(SOURCE_CHAIN_SELECTOR, allowedReceiver);
    MessageV1Codec.MessageV1 memory message =
      _createMessage(SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, address(1), s_testToken, s_sender, s_receiver);

    vm.prank(s_offRamp);
    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.SenderNotAllowed.selector, s_receiver));
    s_verifier.verifyMessage(message, bytes32(0), "");
  }
}
