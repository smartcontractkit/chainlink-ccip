// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolSetup} from "./TokenPoolSetup.t.sol";

contract TokenPool_setRateLimitAdmin is TokenPoolSetup {
  function test_GrantRateLimitAdminRole() public {
    assertFalse(s_tokenPool.hasRateLimitAdminRole(OWNER));
    vm.expectEmit();
    emit TokenPool.RateLimitAdminRoleGranted(OWNER);
    s_tokenPool.grantRateLimitAdminRole(OWNER);
    assertTrue(s_tokenPool.hasRateLimitAdminRole(OWNER));
  }

  function test_RevokeRateLimitAdminRole() public {
    // First grant the role
    s_tokenPool.grantRateLimitAdminRole(OWNER);
    assertTrue(s_tokenPool.hasRateLimitAdminRole(OWNER));
    
    // Then revoke it
    vm.expectEmit();
    emit TokenPool.RateLimitAdminRoleRevoked(OWNER);
    s_tokenPool.revokeRateLimitAdminRole(OWNER);
    assertFalse(s_tokenPool.hasRateLimitAdminRole(OWNER));
  }

  // Reverts

  function test_RevertWhen_GrantRateLimitAdminRole_NotOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.grantRateLimitAdminRole(STRANGER);
  }

  function test_RevertWhen_RevokeRateLimitAdminRole_NotOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.revokeRateLimitAdminRole(OWNER);
  }
}
