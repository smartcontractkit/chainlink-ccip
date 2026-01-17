// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {ERC20LockBoxSetup} from "./ERC20LockBoxSetup.t.sol";

contract ERC20LockBox_withdraw is ERC20LockBoxSetup {
  function testFuzz_Withdraw_Success(
    uint256 amount
  ) public {
    vm.assume(amount != 0);
    amount = bound(amount, 1, type(uint256).max / 2);

    // Deposit tokens first
    _depositTokens(amount);

    uint256 recipientBalanceBefore = s_token.balanceOf(s_recipient);
    uint256 lockBoxBalanceBefore = s_token.balanceOf(address(s_erc20LockBox));

    vm.startPrank(s_allowedCaller);

    vm.expectEmit();
    emit ERC20LockBox.Withdrawal(address(s_token), s_recipient, amount);

    s_erc20LockBox.withdraw(address(s_token), amount, s_recipient);

    vm.stopPrank();

    // Verify balances
    assertEq(s_token.balanceOf(s_recipient), recipientBalanceBefore + amount);
    assertEq(s_token.balanceOf(address(s_erc20LockBox)), lockBoxBalanceBefore - amount);
  }

  function test_Withdraw_MultipleWithdrawals() public {
    uint256 amount1 = 1000e18;
    uint256 amount2 = 2000e18;

    // Deposit tokens
    _depositTokens(amount1);
    _depositTokens(amount2);

    uint256 recipientBalanceBefore = s_token.balanceOf(s_recipient);

    vm.startPrank(s_allowedCaller);

    // Withdraw from first chain selector
    vm.expectEmit();
    emit ERC20LockBox.Withdrawal(address(s_token), s_recipient, amount1);
    s_erc20LockBox.withdraw(address(s_token), amount1, s_recipient);

    // Withdraw from second chain selector
    vm.expectEmit();
    emit ERC20LockBox.Withdrawal(address(s_token), s_recipient, amount2);
    s_erc20LockBox.withdraw(address(s_token), amount2, s_recipient);

    vm.stopPrank();

    // Verify final balances
    assertEq(s_token.balanceOf(s_recipient), recipientBalanceBefore + amount1 + amount2);
    assertEq(s_token.balanceOf(address(s_erc20LockBox)), 0);
  }

  function test_Withdraw_PartialAmount() public {
    uint256 depositAmount = 1000e18;
    uint256 withdrawAmount = 300e18;

    // Deposit tokens
    _depositTokens(depositAmount);

    uint256 recipientBalanceBefore = s_token.balanceOf(s_recipient);
    uint256 expectedRemainingBalance = depositAmount - withdrawAmount;

    vm.startPrank(s_allowedCaller);

    vm.expectEmit();
    emit ERC20LockBox.Withdrawal(address(s_token), s_recipient, withdrawAmount);

    s_erc20LockBox.withdraw(address(s_token), withdrawAmount, s_recipient);

    vm.stopPrank();

    // Verify balances
    assertEq(s_token.balanceOf(s_recipient), recipientBalanceBefore + withdrawAmount);
    assertEq(s_token.balanceOf(address(s_erc20LockBox)), expectedRemainingBalance);
  }

  function test_Withdraw_EventEmission() public {
    uint256 amount = 1000e18;
    _depositTokens(amount);

    vm.startPrank(s_allowedCaller);

    vm.expectEmit();
    emit ERC20LockBox.Withdrawal(address(s_token), s_recipient, amount);

    s_erc20LockBox.withdraw(address(s_token), amount, s_recipient);

    vm.stopPrank();
  }

  function test_Withdraw_ToDifferentRecipients() public {
    uint256 amount = 1000e18;
    address recipient1 = makeAddr("recipient1");
    address recipient2 = makeAddr("recipient2");

    _depositTokens(amount * 2);

    vm.startPrank(s_allowedCaller);

    // Withdraw to first recipient
    vm.expectEmit();
    emit ERC20LockBox.Withdrawal(address(s_token), recipient1, amount);
    s_erc20LockBox.withdraw(address(s_token), amount, recipient1);

    // Withdraw to second recipient
    vm.expectEmit();
    emit ERC20LockBox.Withdrawal(address(s_token), recipient2, amount);
    s_erc20LockBox.withdraw(address(s_token), amount, recipient2);

    vm.stopPrank();

    // Verify both recipients received tokens
    assertEq(s_token.balanceOf(recipient1), amount);
    assertEq(s_token.balanceOf(recipient2), amount);
  }

  // ================================================================
  // │                        Revert Tests                          │
  // ================================================================

  function test_RevertWhen_Unauthorized() public {
    uint256 amount = 1000e18;
    _depositTokens(amount);

    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(ERC20LockBox.Unauthorized.selector, STRANGER));

    s_erc20LockBox.withdraw(address(s_token), amount, s_recipient);
  }

  function test_RevertWhen_RecipientIsZeroAddress() public {
    uint256 amount = 1000e18;
    _depositTokens(amount);

    vm.startPrank(s_allowedCaller);
    vm.expectRevert(ERC20LockBox.RecipientCannotBeZeroAddress.selector);

    s_erc20LockBox.withdraw(address(s_token), amount, address(0));
  }

  function test_RevertWhen_AmountIsZero() public {
    uint256 amount = 1000e18;
    _depositTokens(amount);

    vm.startPrank(s_allowedCaller);
    vm.expectRevert(ERC20LockBox.TokenAmountCannotBeZero.selector);

    s_erc20LockBox.withdraw(address(s_token), 0, s_recipient);
  }

  function test_RevertWhen_TokenIsZeroAddress() public {
    uint256 amount = 1000e18;
    _depositTokens(amount);

    vm.startPrank(s_allowedCaller);
    vm.expectRevert(ERC20LockBox.TokenAddressCannotBeZero.selector);

    s_erc20LockBox.withdraw(address(0), amount, s_recipient);
  }

  function test_RevertWhen_InsufficientBalance() public {
    uint256 depositAmount = 1000e18;
    uint256 withdrawAmount = 1500e18;

    _depositTokens(depositAmount);

    vm.startPrank(s_allowedCaller);
    vm.expectRevert(abi.encodeWithSelector(ERC20LockBox.InsufficientBalance.selector, withdrawAmount, depositAmount));

    s_erc20LockBox.withdraw(address(s_token), withdrawAmount, s_recipient);
  }
}
