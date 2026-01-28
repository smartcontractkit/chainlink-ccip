// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract AdvancedPoolHooks_authorizedCallers is AdvancedPoolHooksSetup {
  address internal s_authorizedCaller = makeAddr("authorizedCaller");
  address internal s_unauthorizedCaller = makeAddr("unauthorizedCaller");

  AdvancedPoolHooks internal s_hooksWithRestrictedCallers;

  function setUp() public virtual override {
    super.setUp();

    // Create AdvancedPoolHooks with restricted callers (authorizedCallers.length > 0 enables restriction)
    address[] memory authorizedCallers = new address[](1);
    authorizedCallers[0] = s_authorizedCaller;
    s_hooksWithRestrictedCallers = new AdvancedPoolHooks(
      new address[](0), // no allowlist
      0, // no threshold
      address(0), // no policy engine
      authorizedCallers // only authorized callers can invoke
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

  function test_getAuthorizedCallersEnabled() public view {
    // Default setup allows anyone to invoke (authorizedCallersEnabled = false)
    assertFalse(s_advancedPoolHooks.getAuthorizedCallersEnabled());

    // Hooks with restricted callers (authorizedCallersEnabled = true)
    assertTrue(s_hooksWithRestrictedCallers.getAuthorizedCallersEnabled());
  }

  function test_preflightCheck_WhenAnyoneCanInvoke() public {
    // Default setup allows anyone to invoke (authorizedCallersEnabled = false)
    assertFalse(s_advancedPoolHooks.getAuthorizedCallersEnabled());

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn();

    // Any caller should be able to call when anyone can invoke
    vm.stopPrank();
    vm.prank(s_unauthorizedCaller);
    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, 5, "");
  }

  function test_postFlightCheck_WhenAnyoneCanInvoke() public {
    // Default setup allows anyone to invoke (authorizedCallersEnabled = false)
    assertFalse(s_advancedPoolHooks.getAuthorizedCallersEnabled());

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    // Any caller should be able to call when anyone can invoke
    vm.stopPrank();
    vm.prank(s_unauthorizedCaller);
    s_advancedPoolHooks.postFlightCheck(releaseOrMintIn, 100e18, 5);
  }

  function test_preflightCheck_WhenOnlyAuthorizedCallersCanInvoke() public {
    assertTrue(s_hooksWithRestrictedCallers.getAuthorizedCallersEnabled());

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn();

    // Authorized caller should succeed
    vm.stopPrank();
    vm.prank(s_authorizedCaller);
    s_hooksWithRestrictedCallers.preflightCheck(lockOrBurnIn, 5, "");
  }

  function test_postFlightCheck_WhenOnlyAuthorizedCallersCanInvoke() public {
    assertTrue(s_hooksWithRestrictedCallers.getAuthorizedCallersEnabled());

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    // Authorized caller should succeed
    vm.stopPrank();
    vm.prank(s_authorizedCaller);
    s_hooksWithRestrictedCallers.postFlightCheck(releaseOrMintIn, 100e18, 5);
  }

  // ================================================================
  // │                       Revert Tests                           │
  // ================================================================

  function test_preflightCheck_RevertWhen_UnauthorizedCaller() public {
    assertTrue(s_hooksWithRestrictedCallers.getAuthorizedCallersEnabled());

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn();

    vm.stopPrank();
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, s_unauthorizedCaller));
    vm.prank(s_unauthorizedCaller);
    s_hooksWithRestrictedCallers.preflightCheck(lockOrBurnIn, 5, "");
  }

  function test_postFlightCheck_RevertWhen_UnauthorizedCaller() public {
    assertTrue(s_hooksWithRestrictedCallers.getAuthorizedCallersEnabled());

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    vm.stopPrank();
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, s_unauthorizedCaller));
    vm.prank(s_unauthorizedCaller);
    s_hooksWithRestrictedCallers.postFlightCheck(releaseOrMintIn, 100e18, 5);
  }
}
