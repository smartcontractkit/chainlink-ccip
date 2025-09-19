// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitteeRamp} from "../../../ccvs/CommitteeRamp.sol";
import {BaseOnRamp} from "../../../ccvs/components/BaseOnRamp.sol";
import {CommitteeRampSetup} from "./CommitteeRampSetup.t.sol";

contract CommitteeRamp_applyAllowlistUpdates is CommitteeRampSetup {
  function setUp() public override {
    super.setUp();

    // Enable allowlist for destination chain once.
    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = _getDestChainConfig(s_router, DEST_CHAIN_SELECTOR, true);
    s_commitRamp.applyDestChainConfigUpdates(destChainConfigs);
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

    s_commitRamp.applyAllowlistUpdates(allowlistConfigs);

    (bool allowlistEnabled,, address[] memory allowlistSender) = s_commitRamp.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(allowlistEnabled, allowlistConfigs[0].allowlistEnabled);
    assertEq(allowlistSender.length, allowlistConfigs[0].addedAllowlistedSenders.length);
    assertEq(allowlistSender[0], allowlistConfigs[0].addedAllowlistedSenders[0]);
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

    s_commitRamp.applyAllowlistUpdates(allowlistConfigs);
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

    vm.expectRevert(CommitteeRamp.OnlyCallableByOwnerOrAllowlistAdmin.selector);
    s_commitRamp.applyAllowlistUpdates(allowlistConfigs);
  }
}
