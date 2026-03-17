// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {LombardVerifierSetup} from "./LombardVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract LombardVerifier_applyAllowlistUpdates is LombardVerifierSetup {
  address internal constant SENDER1 = address(0x1111);
  address internal constant SENDER2 = address(0x2222);
  address internal constant SENDER3 = address(0x3333);

  function setUp() public override {
    super.setUp();

    // Set up a remote chain config
    BaseVerifier.RemoteChainConfigArgs[] memory remoteChainConfigArgs = new BaseVerifier.RemoteChainConfigArgs[](1);
    remoteChainConfigArgs[0] = BaseVerifier.RemoteChainConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      allowlistEnabled: true,
      feeUSDCents: 100,
      gasForVerification: 50_000,
      payloadSizeBytes: 100
    });
    s_lombardVerifier.applyRemoteChainConfigUpdates(remoteChainConfigArgs);
  }

  function test_applyAllowlistUpdates_ByOwner() public {
    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigArgs = new BaseVerifier.AllowlistConfigArgs[](1);
    address[] memory sendersToAdd = new address[](2);
    sendersToAdd[0] = SENDER1;
    sendersToAdd[1] = SENDER2;

    allowlistConfigArgs[0] = BaseVerifier.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      addedAllowlistedSenders: sendersToAdd,
      removedAllowlistedSenders: new address[](0)
    });

    vm.expectEmit();
    emit BaseVerifier.AllowListSendersAdded(DEST_CHAIN_SELECTOR, SENDER1);
    vm.expectEmit();
    emit BaseVerifier.AllowListSendersAdded(DEST_CHAIN_SELECTOR, SENDER2);

    s_lombardVerifier.applyAllowlistUpdates(allowlistConfigArgs);

    (bool allowlistEnabled,, address[] memory allowedSenders) =
      s_lombardVerifier.getRemoteChainConfig(DEST_CHAIN_SELECTOR);

    assertTrue(allowlistEnabled);
    assertEq(allowedSenders.length, 2);
    assertEq(allowedSenders[0], SENDER1);
    assertEq(allowedSenders[1], SENDER2);
  }

  function test_applyAllowlistUpdates_RemoveSenders() public {
    // First add some senders
    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigArgs = new BaseVerifier.AllowlistConfigArgs[](1);
    address[] memory sendersToAdd = new address[](3);
    sendersToAdd[0] = SENDER1;
    sendersToAdd[1] = SENDER2;
    sendersToAdd[2] = SENDER3;

    allowlistConfigArgs[0] = BaseVerifier.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      addedAllowlistedSenders: sendersToAdd,
      removedAllowlistedSenders: new address[](0)
    });

    s_lombardVerifier.applyAllowlistUpdates(allowlistConfigArgs);

    // Now remove one sender
    address[] memory sendersToRemove = new address[](1);
    sendersToRemove[0] = SENDER2;

    allowlistConfigArgs[0] = BaseVerifier.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      addedAllowlistedSenders: new address[](0),
      removedAllowlistedSenders: sendersToRemove
    });

    vm.expectEmit();
    emit BaseVerifier.AllowListSendersRemoved(DEST_CHAIN_SELECTOR, SENDER2);

    s_lombardVerifier.applyAllowlistUpdates(allowlistConfigArgs);

    (,, address[] memory allowedSenders) = s_lombardVerifier.getRemoteChainConfig(DEST_CHAIN_SELECTOR);

    assertEq(allowedSenders.length, 2);
    // Verify SENDER2 is not in the list
    for (uint256 i = 0; i < allowedSenders.length; i++) {
      assertTrue(allowedSenders[i] != SENDER2);
    }
  }

  function test_applyAllowlistUpdates_DisableAllowlist() public {
    // First add some senders
    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigArgs = new BaseVerifier.AllowlistConfigArgs[](1);
    address[] memory sendersToAdd = new address[](1);
    sendersToAdd[0] = SENDER1;

    allowlistConfigArgs[0] = BaseVerifier.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      addedAllowlistedSenders: sendersToAdd,
      removedAllowlistedSenders: new address[](0)
    });

    s_lombardVerifier.applyAllowlistUpdates(allowlistConfigArgs);

    // Now disable the allowlist
    allowlistConfigArgs[0] = BaseVerifier.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false,
      addedAllowlistedSenders: new address[](0),
      removedAllowlistedSenders: new address[](0)
    });

    vm.expectEmit();
    emit BaseVerifier.AllowListStateChanged(DEST_CHAIN_SELECTOR, false);

    s_lombardVerifier.applyAllowlistUpdates(allowlistConfigArgs);

    (bool allowlistEnabled,,) = s_lombardVerifier.getRemoteChainConfig(DEST_CHAIN_SELECTOR);

    assertFalse(allowlistEnabled);
  }

  function test_applyAllowlistUpdates_RevertWhen_NotOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_lombardVerifier.applyAllowlistUpdates(new BaseVerifier.AllowlistConfigArgs[](1));
  }

  function test_applyAllowlistUpdates_RevertWhen_AddingSendersWithAllowlistDisabled() public {
    // First disable the allowlist
    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigArgs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigArgs[0] = BaseVerifier.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false,
      addedAllowlistedSenders: new address[](0),
      removedAllowlistedSenders: new address[](0)
    });

    s_lombardVerifier.applyAllowlistUpdates(allowlistConfigArgs);

    // Try to add senders with allowlist disabled
    address[] memory sendersToAdd = new address[](1);
    sendersToAdd[0] = SENDER1;

    allowlistConfigArgs[0] = BaseVerifier.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false,
      addedAllowlistedSenders: sendersToAdd,
      removedAllowlistedSenders: new address[](0)
    });

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.InvalidAllowListRequest.selector, DEST_CHAIN_SELECTOR));
    s_lombardVerifier.applyAllowlistUpdates(allowlistConfigArgs);
  }

  function test_applyAllowlistUpdates_RevertWhen_AddingZeroAddress() public {
    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigArgs = new BaseVerifier.AllowlistConfigArgs[](1);
    address[] memory sendersToAdd = new address[](1);
    sendersToAdd[0] = address(0);

    allowlistConfigArgs[0] = BaseVerifier.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      addedAllowlistedSenders: sendersToAdd,
      removedAllowlistedSenders: new address[](0)
    });

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.InvalidAllowListRequest.selector, DEST_CHAIN_SELECTOR));
    s_lombardVerifier.applyAllowlistUpdates(allowlistConfigArgs);
  }
}
