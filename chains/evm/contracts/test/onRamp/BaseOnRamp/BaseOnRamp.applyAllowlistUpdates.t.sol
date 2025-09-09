// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {BaseOnRampSetup} from "./BaseOnRampSetup.t.sol";

contract BaseOnRamp_applyAllowlistUpdates is BaseOnRampSetup {
  function setUp() public override {
    super.setUp();

    // First enable allowlist for the destination chain.
    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = _getDestChainConfig(s_ccvProxy, DEST_CHAIN_SELECTOR, true);
    s_baseOnRamp.applyDestChainConfigUpdates(destChainConfigs);
  }

  function test_applyAllowlistUpdates() public {
    address[] memory senders = new address[](1);
    senders[0] = makeAddr("sender1");

    BaseOnRamp.AllowlistConfigArgs[] memory allowlistConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, senders, new address[](0));

    vm.expectEmit();
    emit BaseOnRamp.AllowListSendersAdded(DEST_CHAIN_SELECTOR, senders);

    s_baseOnRamp.applyAllowlistUpdates(allowlistConfigs);

    // Verify sender was added.
    (,, address[] memory allowedSenders) = s_baseOnRamp.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(allowedSenders.length, 1);
    assertEq(allowedSenders[0], senders[0]);
  }

  function test_applyAllowlistUpdates_RemoveSenders() public {
    address[] memory sendersToAdd = new address[](1);
    sendersToAdd[0] = makeAddr("sender1");

    // First add senders.
    BaseOnRamp.AllowlistConfigArgs[] memory addConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    addConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, sendersToAdd, new address[](0));

    s_baseOnRamp.applyAllowlistUpdates(addConfigs);

    // Now remove one and add another.
    address[] memory sendersToRemove = new address[](1);
    sendersToRemove[0] = sendersToAdd[0];

    address[] memory newSendersToAdd = new address[](1);
    newSendersToAdd[0] = makeAddr("sender2");

    BaseOnRamp.AllowlistConfigArgs[] memory updateConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    updateConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, newSendersToAdd, sendersToRemove);

    vm.expectEmit();
    emit BaseOnRamp.AllowListSendersAdded(DEST_CHAIN_SELECTOR, newSendersToAdd);
    vm.expectEmit();
    emit BaseOnRamp.AllowListSendersRemoved(DEST_CHAIN_SELECTOR, sendersToRemove);

    s_baseOnRamp.applyAllowlistUpdates(updateConfigs);

    (,, address[] memory allowedSenders) = s_baseOnRamp.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(allowedSenders[0], newSendersToAdd[0]);
  }

  function test_applyAllowlistUpdates_DisableAllowlist() public {
    address[] memory senders = new address[](1);
    senders[0] = makeAddr("sender1");

    // First add a sender.
    BaseOnRamp.AllowlistConfigArgs[] memory addConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    addConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, senders, new address[](0));
    s_baseOnRamp.applyAllowlistUpdates(addConfigs);

    // Now disable allowlist.
    BaseOnRamp.AllowlistConfigArgs[] memory disableConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    disableConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, false, new address[](0), new address[](0));

    s_baseOnRamp.applyAllowlistUpdates(disableConfigs);

    (bool allowlistEnabled,,) = s_baseOnRamp.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertFalse(allowlistEnabled);
  }

  // Reverts

  function test_applyAllowlistUpdates_RevertWhen_InvalidAllowListRequest_AddingToDisabledAllowlist() public {
    // First disable allowlist on the chain.
    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = _getDestChainConfig(s_ccvProxy, DEST_CHAIN_SELECTOR, false);
    s_baseOnRamp.applyDestChainConfigUpdates(destChainConfigs);

    address[] memory senders = new address[](1);
    senders[0] = makeAddr("sender");

    BaseOnRamp.AllowlistConfigArgs[] memory allowlistConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, false, senders, new address[](0));

    vm.expectRevert(abi.encodeWithSelector(BaseOnRamp.InvalidAllowListRequest.selector, DEST_CHAIN_SELECTOR));
    s_baseOnRamp.applyAllowlistUpdates(allowlistConfigs);
  }

  function test_applyAllowlistUpdates_RevertWhen_InvalidAllowListRequest_ZeroAddressSender() public {
    address[] memory senders = new address[](1);
    senders[0] = address(0); // Zero address should be invalid.

    BaseOnRamp.AllowlistConfigArgs[] memory allowlistConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, senders, new address[](0));

    vm.expectRevert(abi.encodeWithSelector(BaseOnRamp.InvalidAllowListRequest.selector, DEST_CHAIN_SELECTOR));
    s_baseOnRamp.applyAllowlistUpdates(allowlistConfigs);
  }
}
