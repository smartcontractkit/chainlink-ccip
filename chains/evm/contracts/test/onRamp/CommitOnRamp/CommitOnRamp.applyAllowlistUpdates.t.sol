// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {CommitOnRamp} from "../../../onRamp/CommitOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";

contract CommitOnRamp_applyAllowlistUpdates is CommitOnRampSetup {
  function setUp() public override {
    super.setUp();

    // First enable allowlist for the destination chain
    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = _getDestChainConfig(s_ccvProxy, DEST_CHAIN_SELECTOR, true);
    s_commitOnRamp.applyDestChainConfigUpdates(destChainConfigs);
  }

  function test_applyAllowlistUpdates_AsOwner() public {
    address[] memory senders = new address[](1);
    senders[0] = makeAddr("sender1");

    BaseOnRamp.AllowlistConfigArgs[] memory allowlistConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = BaseOnRamp.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      addedAllowlistedSenders: senders,
      removedAllowlistedSenders: new address[](0)
    });

    vm.expectEmit();
    emit BaseOnRamp.AllowListSendersAdded(DEST_CHAIN_SELECTOR, senders);

    s_commitOnRamp.applyAllowlistUpdates(allowlistConfigs);
  }

  function test_applyAllowlistUpdates_AsAllowlistAdmin() public {
    vm.stopPrank();
    vm.startPrank(ALLOWLIST_ADMIN);

    address[] memory senders = new address[](1);
    senders[0] = makeAddr("sender1");

    BaseOnRamp.AllowlistConfigArgs[] memory allowlistConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = BaseOnRamp.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      addedAllowlistedSenders: senders,
      removedAllowlistedSenders: new address[](0)
    });

    vm.expectEmit();
    emit BaseOnRamp.AllowListSendersAdded(DEST_CHAIN_SELECTOR, senders);

    s_commitOnRamp.applyAllowlistUpdates(allowlistConfigs);
  }

  function test_applyAllowlistUpdates_AddAndRemoveSenders() public {
    address[] memory sendersToAdd = new address[](1);
    sendersToAdd[0] = makeAddr("sender1");

    BaseOnRamp.AllowlistConfigArgs[] memory addConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    addConfigs[0] = BaseOnRamp.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      addedAllowlistedSenders: sendersToAdd,
      removedAllowlistedSenders: new address[](0)
    });

    s_commitOnRamp.applyAllowlistUpdates(addConfigs);

    // Now remove one and add another
    address[] memory sendersToRemove = new address[](1);
    sendersToRemove[0] = sendersToAdd[0];

    address[] memory newSendersToAdd = new address[](1);
    newSendersToAdd[0] = makeAddr("sender2");

    BaseOnRamp.AllowlistConfigArgs[] memory updateConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    updateConfigs[0] = BaseOnRamp.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      addedAllowlistedSenders: newSendersToAdd,
      removedAllowlistedSenders: sendersToRemove
    });

    vm.expectEmit();
    emit BaseOnRamp.AllowListSendersAdded(DEST_CHAIN_SELECTOR, newSendersToAdd);
    vm.expectEmit();
    emit BaseOnRamp.AllowListSendersRemoved(DEST_CHAIN_SELECTOR, sendersToRemove);

    s_commitOnRamp.applyAllowlistUpdates(updateConfigs);
  }

  function test_applyAllowlistUpdates_DisableAllowlist() public {
    address[] memory senders = new address[](1);
    senders[0] = makeAddr("sender1");

    BaseOnRamp.AllowlistConfigArgs[] memory addConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    addConfigs[0] = BaseOnRamp.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      addedAllowlistedSenders: senders,
      removedAllowlistedSenders: new address[](0)
    });

    s_commitOnRamp.applyAllowlistUpdates(addConfigs);

    // Now disable allowlist
    BaseOnRamp.AllowlistConfigArgs[] memory disableConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    disableConfigs[0] = BaseOnRamp.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false,
      addedAllowlistedSenders: new address[](0),
      removedAllowlistedSenders: new address[](0)
    });

    s_commitOnRamp.applyAllowlistUpdates(disableConfigs);

    (bool allowlistEnabled,,) = s_commitOnRamp.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertFalse(allowlistEnabled);
  }

  // Reverts

  function test_applyAllowlistUpdates_RevertWhen_OnlyCallableByOwnerOrAllowlistAdmin() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);

    BaseOnRamp.AllowlistConfigArgs[] memory allowlistConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = BaseOnRamp.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      addedAllowlistedSenders: new address[](0),
      removedAllowlistedSenders: new address[](0)
    });

    vm.expectRevert(CommitOnRamp.OnlyCallableByOwnerOrAllowlistAdmin.selector);
    s_commitOnRamp.applyAllowlistUpdates(allowlistConfigs);
  }

  function test_applyAllowlistUpdates_RevertWhen_InvalidAllowListRequest_AddingToDisabledAllowlist() public {
    // First disable allowlist on the chain
    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = _getDestChainConfig(s_ccvProxy, DEST_CHAIN_SELECTOR, false);
    s_commitOnRamp.applyDestChainConfigUpdates(destChainConfigs);

    address[] memory senders = new address[](1);
    senders[0] = makeAddr("sender");

    BaseOnRamp.AllowlistConfigArgs[] memory allowlistConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = BaseOnRamp.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false,
      addedAllowlistedSenders: senders,
      removedAllowlistedSenders: new address[](0)
    });

    vm.expectRevert(abi.encodeWithSelector(BaseOnRamp.InvalidAllowListRequest.selector, DEST_CHAIN_SELECTOR));
    s_commitOnRamp.applyAllowlistUpdates(allowlistConfigs);
  }

  function test_applyAllowlistUpdates_RevertWhen_InvalidAllowListRequest_ZeroAddressSender() public {
    address[] memory senders = new address[](1);
    senders[0] = address(0); // Zero address should be invalid

    BaseOnRamp.AllowlistConfigArgs[] memory allowlistConfigs = new BaseOnRamp.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = BaseOnRamp.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      addedAllowlistedSenders: senders,
      removedAllowlistedSenders: new address[](0)
    });

    vm.expectRevert(abi.encodeWithSelector(BaseOnRamp.InvalidAllowListRequest.selector, DEST_CHAIN_SELECTOR));
    s_commitOnRamp.applyAllowlistUpdates(allowlistConfigs);
  }
}
