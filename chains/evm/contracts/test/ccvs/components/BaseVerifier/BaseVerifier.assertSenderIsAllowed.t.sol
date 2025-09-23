// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../../ccvs/components/BaseVerifier.sol";
import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

contract BaseVerifier_assertSenderIsAllowed is BaseVerifierSetup {
  function setUp() public override {
    super.setUp();
    vm.stopPrank();
  }

  function test_assertSenderIsAllowed() public {
    // Should allow any sender when allowlist is disabled.
    address anySender = makeAddr("anySender");

    vm.prank(s_ccvProxy);
    s_baseVerifier.assertSenderIsAllowed(DEST_CHAIN_SELECTOR, anySender, s_ccvProxy);
  }

  function test_assertSenderIsAllowed_AllowlistEnabledWithAllowedSender() public {
    // Enable allowlist and add a sender.
    BaseVerifier.DestChainConfigArgs[] memory destChainConfigs = new BaseVerifier.DestChainConfigArgs[](1);
    destChainConfigs[0] = _getDestChainConfig(s_router, DEST_CHAIN_SELECTOR, true);

    vm.prank(OWNER);
    s_baseVerifier.applyDestChainConfigUpdates(destChainConfigs);

    address allowedSender = makeAddr("allowedSender");
    address[] memory sendersToAdd = new address[](1);
    sendersToAdd[0] = allowedSender;

    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, sendersToAdd, new address[](0));

    vm.prank(OWNER);
    s_baseVerifier.applyAllowlistUpdates(allowlistConfigs);

    vm.prank(s_ccvProxy);
    s_baseVerifier.assertSenderIsAllowed(DEST_CHAIN_SELECTOR, allowedSender, s_ccvProxy);
  }

  // Reverts

  function test_assertSenderIsAllowed_RevertWhen_CallerIsNotARampOnRouter() public {
    address sender = makeAddr("sender");

    // Try calling from non-ccvProxy address
    vm.prank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CallerIsNotARampOnRouter.selector, STRANGER));
    s_baseVerifier.assertSenderIsAllowed(DEST_CHAIN_SELECTOR, sender, STRANGER);
  }

  function test_assertSenderIsAllowed_RevertWhen_SenderNotAllowed() public {
    // Enable allowlist and add one sender.
    BaseVerifier.DestChainConfigArgs[] memory destChainConfigs = new BaseVerifier.DestChainConfigArgs[](1);
    destChainConfigs[0] = _getDestChainConfig(s_router, DEST_CHAIN_SELECTOR, true);

    vm.prank(OWNER);
    s_baseVerifier.applyDestChainConfigUpdates(destChainConfigs);

    address allowedSender = makeAddr("allowedSender");
    address notAllowedSender = makeAddr("notAllowedSender");

    address[] memory sendersToAdd = new address[](1);
    sendersToAdd[0] = allowedSender;

    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, sendersToAdd, new address[](0));

    vm.prank(OWNER);
    s_baseVerifier.applyAllowlistUpdates(allowlistConfigs);

    // Should revert for non-allowed sender.
    vm.prank(s_ccvProxy);
    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.SenderNotAllowed.selector, notAllowedSender));
    s_baseVerifier.assertSenderIsAllowed(DEST_CHAIN_SELECTOR, notAllowedSender, s_ccvProxy);
  }
}
