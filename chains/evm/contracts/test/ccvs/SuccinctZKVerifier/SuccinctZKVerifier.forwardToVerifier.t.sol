// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";

contract SuccinctZKVerifier_forwardToVerifier is SuccinctZKVerifierSetup {
  function setUp() public override {
    super.setUp();

    vm.stopPrank();
  }

  function test_forwardToVerifier() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessageV1(SOURCE_CHAIN_SELECTOR);
    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));

    vm.prank(s_onRamp);
    bytes memory verifierData = s_zkVerifier.forwardToVerifier(message, messageId, address(0), 0, "");

    assertEq(abi.encodePacked(VERSION_TAG_V0_0_1), verifierData);
  }

  function test_forwardToVerifier_AllowlistedSender() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessageV1(SOURCE_CHAIN_SELECTOR);
    address sender = abi.decode(message.sender, (address));
    address[] memory addedSenders = new address[](1);
    addedSenders[0] = sender;
    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, addedSenders, new address[](0));
    vm.prank(OWNER);
    s_zkVerifier.applyAllowlistUpdates(allowlistConfigs);

    vm.prank(s_onRamp);
    bytes memory verifierData = s_zkVerifier.forwardToVerifier(message, bytes32(0), address(0), 0, "");

    assertEq(abi.encodePacked(VERSION_TAG_V0_0_1), verifierData);
  }

  // Reverts

  function test_forwardToVerifier_RevertWhen_CursedByRMN() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessageV1(SOURCE_CHAIN_SELECTOR);
    _setMockRMNChainCurse(message.destChainSelector, true);

    vm.prank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CursedByRMN.selector, message.destChainSelector));
    s_zkVerifier.forwardToVerifier(message, bytes32(0), address(0), 0, "");
  }

  function test_forwardToVerifier_RevertWhen_CallerIsNotARampOnRouter() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessageV1(SOURCE_CHAIN_SELECTOR);

    vm.prank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CallerIsNotARampOnRouter.selector, STRANGER));
    s_zkVerifier.forwardToVerifier(message, bytes32(0), address(0), 0, "");
  }

  function test_forwardToVerifier_RevertWhen_SenderNotAllowed() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessageV1(SOURCE_CHAIN_SELECTOR);
    address sender = abi.decode(message.sender, (address));
    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, new address[](0), new address[](0));
    vm.prank(OWNER);
    s_zkVerifier.applyAllowlistUpdates(allowlistConfigs);

    vm.prank(s_onRamp);
    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.SenderNotAllowed.selector, sender));
    s_zkVerifier.forwardToVerifier(message, bytes32(0), address(0), 0, "");
  }
}
