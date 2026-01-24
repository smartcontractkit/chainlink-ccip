// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract AdvancedPoolHooks_authorizedCallers is AdvancedPoolHooksSetup {
  address internal s_authorizedCaller = makeAddr("authorizedCaller");
  address internal s_unauthorizedCaller = makeAddr("unauthorizedCaller");

  AdvancedPoolHooks internal s_hooksWithRestrictedCallers;

  function setUp() public virtual override {
    super.setUp();

    // Create AdvancedPoolHooks with restricted callers (allowAnyoneToInvokeThisHook = false)
    address[] memory authorizedCallers = new address[](1);
    authorizedCallers[0] = s_authorizedCaller;
    s_hooksWithRestrictedCallers = new AdvancedPoolHooks(
      new address[](0), // no allowlist
      0, // no threshold
      address(0), // no policy engine
      authorizedCallers,
      false // only authorized callers can invoke
    );
  }

  function _createLockOrBurnIn() internal view returns (Pool.LockOrBurnInV1 memory) {
    return Pool.LockOrBurnInV1({
      receiver: s_receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: OWNER,
      amount: 100e18,
      localToken: address(s_token)
    });
  }

  function _createReleaseOrMintIn() internal view returns (Pool.ReleaseOrMintInV1 memory) {
    return Pool.ReleaseOrMintInV1({
      originalSender: abi.encode(s_sender),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      receiver: OWNER,
      sourceDenominatedAmount: 100e18,
      localToken: address(s_token),
      sourcePoolAddress: abi.encode(address(s_tokenPool)),
      sourcePoolData: "",
      offchainTokenData: ""
    });
  }

  // ================================================================
  // │                       Success Tests                          │
  // ================================================================

  function test_setAllowAnyoneToInvokeThisHook() public {
    assertTrue(s_advancedPoolHooks.getAllowAnyoneToInvokeThisHook());

    vm.expectEmit();
    emit AdvancedPoolHooks.AllowAnyoneToInvokeThisHookSet(false);

    s_advancedPoolHooks.setAllowAnyoneToInvokeThisHook(false);

    assertFalse(s_advancedPoolHooks.getAllowAnyoneToInvokeThisHook());

    // Toggle back to true
    vm.expectEmit();
    emit AdvancedPoolHooks.AllowAnyoneToInvokeThisHookSet(true);

    s_advancedPoolHooks.setAllowAnyoneToInvokeThisHook(true);

    assertTrue(s_advancedPoolHooks.getAllowAnyoneToInvokeThisHook());
  }

  function test_getAllowAnyoneToInvokeThisHook() public view {
    // Default setup allows anyone to invoke
    assertTrue(s_advancedPoolHooks.getAllowAnyoneToInvokeThisHook());

    // Hooks with restricted callers (only authorized callers)
    assertFalse(s_hooksWithRestrictedCallers.getAllowAnyoneToInvokeThisHook());
  }

  function test_preflightCheck_WhenAnyoneCanInvoke() public {
    // Default setup allows anyone to invoke
    assertTrue(s_advancedPoolHooks.getAllowAnyoneToInvokeThisHook());

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn();

    // Any caller should be able to call when anyone can invoke
    vm.stopPrank();
    vm.prank(s_unauthorizedCaller);
    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, 5, "");
  }

  function test_postFlightCheck_WhenAnyoneCanInvoke() public {
    // Default setup allows anyone to invoke
    assertTrue(s_advancedPoolHooks.getAllowAnyoneToInvokeThisHook());

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    // Any caller should be able to call when anyone can invoke
    vm.stopPrank();
    vm.prank(s_unauthorizedCaller);
    s_advancedPoolHooks.postFlightCheck(releaseOrMintIn, 100e18, 5);
  }

  function test_preflightCheck_WhenOnlyAuthorizedCallersCanInvoke() public {
    assertFalse(s_hooksWithRestrictedCallers.getAllowAnyoneToInvokeThisHook());

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn();

    // Authorized caller should succeed
    vm.stopPrank();
    vm.prank(s_authorizedCaller);
    s_hooksWithRestrictedCallers.preflightCheck(lockOrBurnIn, 5, "");
  }

  function test_postFlightCheck_WhenOnlyAuthorizedCallersCanInvoke() public {
    assertFalse(s_hooksWithRestrictedCallers.getAllowAnyoneToInvokeThisHook());

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    // Authorized caller should succeed
    vm.stopPrank();
    vm.prank(s_authorizedCaller);
    s_hooksWithRestrictedCallers.postFlightCheck(releaseOrMintIn, 100e18, 5);
  }

  // ================================================================
  // │                       Revert Tests                           │
  // ================================================================

  function test_setAllowAnyoneToInvokeThisHook_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_advancedPoolHooks.setAllowAnyoneToInvokeThisHook(false);
  }

  function test_preflightCheck_RevertWhen_UnauthorizedCaller() public {
    assertFalse(s_hooksWithRestrictedCallers.getAllowAnyoneToInvokeThisHook());

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn();

    vm.stopPrank();
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, s_unauthorizedCaller));
    vm.prank(s_unauthorizedCaller);
    s_hooksWithRestrictedCallers.preflightCheck(lockOrBurnIn, 5, "");
  }

  function test_postFlightCheck_RevertWhen_UnauthorizedCaller() public {
    assertFalse(s_hooksWithRestrictedCallers.getAllowAnyoneToInvokeThisHook());

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    vm.stopPrank();
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, s_unauthorizedCaller));
    vm.prank(s_unauthorizedCaller);
    s_hooksWithRestrictedCallers.postFlightCheck(releaseOrMintIn, 100e18, 5);
  }
}
