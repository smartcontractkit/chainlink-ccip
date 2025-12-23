// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {ERC20LockBoxSetup} from "./ERC20LockBoxSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract ERC20LockBox_deposit is ERC20LockBoxSetup {
  function testFuzz_deposit_Success(
    uint256 amount
  ) public {
    amount = bound(amount, 1, type(uint256).max / 2);

    uint256 lockBoxBalanceBefore = s_token.balanceOf(address(s_erc20LockBox));
    uint256 callerBalanceBefore = s_token.balanceOf(s_allowedCaller);

    vm.startPrank(s_allowedCaller);
    s_token.approve(address(s_erc20LockBox), amount);

    vm.expectEmit();
    emit ERC20LockBox.Deposit(address(s_token), s_allowedCaller, amount);

    s_erc20LockBox.deposit(address(s_token), 0, amount);

    vm.stopPrank();

    // Verify balances
    assertEq(s_token.balanceOf(address(s_erc20LockBox)), lockBoxBalanceBefore + amount);
    assertEq(s_token.balanceOf(s_allowedCaller), callerBalanceBefore - amount);
  }

  function test_deposit_MultipleDeposits() public {
    uint256 amount1 = 1000e18;
    uint256 amount2 = 2000e18;

    vm.startPrank(s_allowedCaller);
    s_token.approve(address(s_erc20LockBox), amount1 + amount2);

    // First deposit
    vm.expectEmit();
    emit ERC20LockBox.Deposit(address(s_token), s_allowedCaller, amount1);
    s_erc20LockBox.deposit(address(s_token), 0, amount1);

    // Second deposit
    vm.expectEmit();
    emit ERC20LockBox.Deposit(address(s_token), s_allowedCaller, amount2);

    s_erc20LockBox.deposit(address(s_token), 0, amount2);

    vm.stopPrank();

    // Verify total balance
    assertEq(s_token.balanceOf(address(s_erc20LockBox)), amount1 + amount2);
  }

  function test_deposit_FromDifferentCallers() public {
    uint256 amount = 1000e18;
    address caller1 = makeAddr("caller1");
    address caller2 = makeAddr("caller2");

    // Give tokens to both callers
    deal(address(s_token), caller1, amount);
    deal(address(s_token), caller2, amount);

    // Configure both callers as allowed
    address[] memory callers = new address[](2);
    callers[0] = caller1;
    callers[1] = caller2;
    s_erc20LockBox.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: callers, removedCallers: new address[](0)})
    );

    // First caller deposits
    vm.startPrank(caller1);
    s_token.approve(address(s_erc20LockBox), amount);

    vm.expectEmit();
    emit ERC20LockBox.Deposit(address(s_token), caller1, amount);

    s_erc20LockBox.deposit(address(s_token), 0, amount);
    vm.stopPrank();

    // Second caller deposits
    vm.startPrank(caller2);
    s_token.approve(address(s_erc20LockBox), amount);

    vm.expectEmit();
    emit ERC20LockBox.Deposit(address(s_token), caller2, amount);

    s_erc20LockBox.deposit(address(s_token), 0, amount);
    vm.stopPrank();

    // Verify balances
    assertEq(s_token.balanceOf(address(s_erc20LockBox)), amount * 2);
  }

  // Reverts
  function test_deposit_RevertWhen_Unauthorized() public {
    uint256 amount = 1000e18;

    vm.startPrank(STRANGER);
    s_token.approve(address(s_erc20LockBox), amount);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));

    s_erc20LockBox.deposit(address(s_token), 0, amount);
  }

  function test_deposit_RevertWhen_TokenAmountCannotBeZero() public {
    vm.startPrank(s_allowedCaller);
    s_token.approve(address(s_erc20LockBox), 1);
    vm.expectRevert(ERC20LockBox.TokenAmountCannotBeZero.selector);

    s_erc20LockBox.deposit(address(s_token), 0, 0);
  }
}
