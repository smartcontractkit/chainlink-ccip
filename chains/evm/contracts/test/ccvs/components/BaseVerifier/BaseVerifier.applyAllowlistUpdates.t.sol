// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../../ccvs/components/BaseVerifier.sol";
import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

contract BaseVerifier_applyAllowlistUpdates is BaseVerifierSetup {
  function setUp() public override {
    super.setUp();

    // First enable allowlist for the destination chain.
    BaseVerifier.DestChainConfigArgs[] memory destChainConfigs = new BaseVerifier.DestChainConfigArgs[](1);
    destChainConfigs[0] = _getDestChainConfig(s_router, DEST_CHAIN_SELECTOR, true);
    s_baseVerifier.applyDestChainConfigUpdates(destChainConfigs);
  }

  function test_applyAllowlistUpdates() public {
    address[] memory senders = new address[](1);
    senders[0] = makeAddr("sender1");

    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, senders, new address[](0));

    vm.expectEmit();
    emit BaseVerifier.AllowListSendersAdded(DEST_CHAIN_SELECTOR, senders);

    s_baseVerifier.applyAllowlistUpdates(allowlistConfigs);

    // Verify sender was added.
    (,, address[] memory allowedSenders) = s_baseVerifier.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(allowedSenders.length, 1);
    assertEq(allowedSenders[0], senders[0]);
  }

  function test_applyAllowlistUpdates_RemoveSenders() public {
    address[] memory sendersToAdd = new address[](1);
    sendersToAdd[0] = makeAddr("sender1");

    // First add senders.
    BaseVerifier.AllowlistConfigArgs[] memory addConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    addConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, sendersToAdd, new address[](0));

    s_baseVerifier.applyAllowlistUpdates(addConfigs);

    // Now remove one and add another.
    address[] memory sendersToRemove = new address[](1);
    sendersToRemove[0] = sendersToAdd[0];

    address[] memory newSendersToAdd = new address[](1);
    newSendersToAdd[0] = makeAddr("sender2");

    BaseVerifier.AllowlistConfigArgs[] memory updateConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    updateConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, newSendersToAdd, sendersToRemove);

    vm.expectEmit();
    emit BaseVerifier.AllowListSendersAdded(DEST_CHAIN_SELECTOR, newSendersToAdd);
    vm.expectEmit();
    emit BaseVerifier.AllowListSendersRemoved(DEST_CHAIN_SELECTOR, sendersToRemove);

    s_baseVerifier.applyAllowlistUpdates(updateConfigs);

    (,, address[] memory allowedSenders) = s_baseVerifier.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(allowedSenders[0], newSendersToAdd[0]);
  }

  function test_applyAllowlistUpdates_DisableAllowlist() public {
    address[] memory senders = new address[](1);
    senders[0] = makeAddr("sender1");

    // First add a sender.
    BaseVerifier.AllowlistConfigArgs[] memory addConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    addConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, senders, new address[](0));
    s_baseVerifier.applyAllowlistUpdates(addConfigs);

    // Now disable allowlist.
    BaseVerifier.AllowlistConfigArgs[] memory disableConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    disableConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, false, new address[](0), new address[](0));

    s_baseVerifier.applyAllowlistUpdates(disableConfigs);

    (bool allowlistEnabled,,) = s_baseVerifier.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertFalse(allowlistEnabled);
  }

  // Reverts

  function test_applyAllowlistUpdates_RevertWhen_InvalidAllowListRequest_AddingToDisabledAllowlist() public {
    // First disable allowlist on the chain.
    BaseVerifier.DestChainConfigArgs[] memory destChainConfigs = new BaseVerifier.DestChainConfigArgs[](1);
    destChainConfigs[0] = _getDestChainConfig(s_router, DEST_CHAIN_SELECTOR, false);
    s_baseVerifier.applyDestChainConfigUpdates(destChainConfigs);

    address[] memory senders = new address[](1);
    senders[0] = makeAddr("sender");

    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, false, senders, new address[](0));

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.InvalidAllowListRequest.selector, DEST_CHAIN_SELECTOR));
    s_baseVerifier.applyAllowlistUpdates(allowlistConfigs);
  }

  function test_applyAllowlistUpdates_RevertWhen_InvalidAllowListRequest_ZeroAddressSender() public {
    address[] memory senders = new address[](1);
    senders[0] = address(0); // Zero address should be invalid.

    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, senders, new address[](0));

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.InvalidAllowListRequest.selector, DEST_CHAIN_SELECTOR));
    s_baseVerifier.applyAllowlistUpdates(allowlistConfigs);
  }
}
