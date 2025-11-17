// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {LockReleaseTokenPoolSetup} from "./LockReleaseTokenPoolSetup.t.sol";

contract LockReleaseTokenPool_provideLiquidity is LockReleaseTokenPoolSetup {
  function testFuzz_provideLiquidity(
    uint256 amount
  ) public {
    vm.assume(amount > 0);
    uint256 balancePre = s_token.balanceOf(OWNER);
    s_token.approve(address(s_lockReleaseTokenPool), amount);

    s_lockReleaseTokenPool.provideLiquidity(amount);

    assertEq(s_token.balanceOf(OWNER), balancePre - amount);
    assertEq(s_token.balanceOf(address(s_lockBox)), amount);
  }

  // Reverts.

  function test_RevertWhen_Unauthorized() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(TokenPool.Unauthorized.selector, STRANGER));

    s_lockReleaseTokenPool.provideLiquidity(1);
  }
}
