// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPVerifier} from "../../../ccvs/CCTPVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {CCTPVerifierSetup} from "./CCTPVerifierSetup.t.sol";

contract CCTPVerifier_applyAllowlistUpdates is CCTPVerifierSetup {
  function setUp() public override {
    super.setUp();

    // Enable allowlist for destination chain once.
    BaseVerifier.RemoteChainConfigArgs[] memory remoteChainConfigs = new BaseVerifier.RemoteChainConfigArgs[](1);
    remoteChainConfigs[0] = _getRemoteChainConfig(s_router, DEST_CHAIN_SELECTOR, true);
    s_cctpVerifier.applyRemoteChainConfigUpdates(remoteChainConfigs);
  }

  function test_applyAllowlistUpdates_AsOwner() public {
    address[] memory senders = new address[](1);
    senders[0] = makeAddr("sender1");

    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = BaseVerifier.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      addedAllowlistedSenders: senders,
      removedAllowlistedSenders: new address[](0)
    });

    vm.expectEmit();
    emit BaseVerifier.AllowListSendersAdded(DEST_CHAIN_SELECTOR, senders);

    s_cctpVerifier.applyAllowlistUpdates(allowlistConfigs);

    (bool allowlistEnabled,, address[] memory allowlistSender) =
      s_cctpVerifier.getRemoteChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(allowlistEnabled, allowlistConfigs[0].allowlistEnabled);
    assertEq(allowlistSender.length, allowlistConfigs[0].addedAllowlistedSenders.length);
    assertEq(allowlistSender[0], allowlistConfigs[0].addedAllowlistedSenders[0]);
  }

  function test_applyAllowlistUpdates_AsAllowlistAdmin() public {
    vm.stopPrank();
    vm.startPrank(ALLOWLIST_ADMIN);

    address[] memory senders = new address[](1);
    senders[0] = makeAddr("sender1");

    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = BaseVerifier.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      addedAllowlistedSenders: senders,
      removedAllowlistedSenders: new address[](0)
    });

    vm.expectEmit();
    emit BaseVerifier.AllowListSendersAdded(DEST_CHAIN_SELECTOR, senders);

    s_cctpVerifier.applyAllowlistUpdates(allowlistConfigs);
  }

  function test_applyAllowlistUpdates_RevertWhen_OnlyCallableByOwnerOrAllowlistAdmin() public {
    vm.stopPrank();

    vm.expectRevert(CCTPVerifier.OnlyCallableByOwnerOrAllowlistAdmin.selector);
    s_cctpVerifier.applyAllowlistUpdates(new BaseVerifier.AllowlistConfigArgs[](1));
  }
}
