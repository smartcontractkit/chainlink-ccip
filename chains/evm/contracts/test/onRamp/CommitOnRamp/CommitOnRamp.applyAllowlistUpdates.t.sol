// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {CommitOnRamp} from "../../../onRamp/CommitOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";

contract CommitOnRamp_applyAllowlistUpdates is CommitOnRampSetup {
  function setUp() public override {
    super.setUp();

    // Enable allowlist for destination chain once
    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = _createDestChainConfigArgs(s_ccvProxy, DEST_CHAIN_SELECTOR, true);
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
}
