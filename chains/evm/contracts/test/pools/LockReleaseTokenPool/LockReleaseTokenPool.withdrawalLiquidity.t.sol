// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {LockReleaseTokenPool} from "../../../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {LockReleaseTokenPoolSetup} from "./LockReleaseTokenPoolSetup.t.sol";

contract LockReleaseTokenPool_withdrawalLiquidity is LockReleaseTokenPoolSetup {
  function testFuzz_withdrawLiquidity(
    uint256 amount
  ) public {
    amount = bound(amount, 1, type(uint256).max);
    uint256 balancePre = s_token.balanceOf(OWNER);
    s_token.approve(address(s_lockReleaseTokenPool), amount);
    s_lockReleaseTokenPool.provideLiquidity(amount);

    s_lockReleaseTokenPool.withdrawLiquidity(amount);

    assertEq(s_token.balanceOf(OWNER), balancePre);
  }

  // Reverts.
  function test_withdrawLiquidity_RevertWhen_Unauthorized() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(TokenPool.Unauthorized.selector, STRANGER));

    s_lockReleaseTokenPool.withdrawLiquidity(1);
  }

  function test_withdrawLiquidity_RevertWhen_InsufficientBalance() public {
    uint256 amount = 1000;
    s_token.approve(address(s_lockReleaseTokenPool), amount);
    s_lockReleaseTokenPool.provideLiquidity(amount);

    vm.expectRevert(abi.encodeWithSelector(ERC20LockBox.InsufficientBalance.selector, amount + 1, amount));
    s_lockReleaseTokenPool.withdrawLiquidity(amount + 1);
  }
}
