// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolSetup} from "./TokenPoolSetup.t.sol";

contract TokenPool_grantRateLimitAdminRole is TokenPoolSetup {
  function test_grantRateLimitAdminRole() public {
    assertFalse(s_tokenPool.hasRole(s_tokenPool.RATE_LIMITER_ADMIN_ROLE(), OWNER));
    vm.expectEmit();
    emit TokenPool.RateLimitAdminRoleGranted(OWNER);
    s_tokenPool.grantRateLimitAdminRole(OWNER);
    assertTrue(s_tokenPool.hasRole(s_tokenPool.RATE_LIMITER_ADMIN_ROLE(), OWNER));
  }

  // Reverts
  function test_RevertWhen_grantRateLimitAdminRole_NotOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.grantRateLimitAdminRole(STRANGER);
  }
}
