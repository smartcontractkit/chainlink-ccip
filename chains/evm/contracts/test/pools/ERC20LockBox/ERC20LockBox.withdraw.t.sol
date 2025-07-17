// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {ERC20LockBoxSetup} from "./ERC20LockBoxSetup.t.sol";

contract ERC20LockBox_withdraw is ERC20LockBoxSetup {
  function testFuzz_Withdraw_Success(
    uint256 amount
  ) public {
    vm.assume(amount > 0 && amount <= type(uint256).max / 2);

    // Deposit tokens first
    _depositTokens(amount, DEST_CHAIN_SELECTOR);

    uint256 recipientBalanceBefore = s_token.balanceOf(s_recipient);
    uint256 lockBoxBalanceBefore = s_token.balanceOf(address(s_erc20LockBox));
    uint256 chainBalanceBefore = s_erc20LockBox.getBalance(address(s_token), DEST_CHAIN_SELECTOR);

    vm.startPrank(s_allowedCaller);

    vm.expectEmit();
    emit ERC20LockBox.Withdrawal(DEST_CHAIN_SELECTOR, s_recipient, amount);

    s_erc20LockBox.withdraw(address(s_token), amount, s_recipient, DEST_CHAIN_SELECTOR);

    vm.stopPrank();

    // Verify balances
    assertEq(s_token.balanceOf(s_recipient), recipientBalanceBefore + amount);
    assertEq(s_token.balanceOf(address(s_erc20LockBox)), lockBoxBalanceBefore - amount);
    assertEq(s_erc20LockBox.getBalance(address(s_token), DEST_CHAIN_SELECTOR), chainBalanceBefore - amount);
  }

  function test_Withdraw_WithMultipleChainSelectors() public {
    uint256 amount1 = 1000e18;
    uint256 amount2 = 2000e18;
    uint64 chainSelector1 = SOURCE_CHAIN_SELECTOR;
    uint64 chainSelector2 = DEST_CHAIN_SELECTOR;

    // Deposit tokens for two different chain selectors
    _depositTokens(amount1, chainSelector1);
    _depositTokens(amount2, chainSelector2);

    uint256 recipientBalanceBefore = s_token.balanceOf(s_recipient);

    vm.startPrank(s_allowedCaller);

    // Withdraw from first chain selector
    s_erc20LockBox.withdraw(address(s_token), amount1, s_recipient, chainSelector1);

    // Verify first chain selector balance is zero
    assertEq(s_erc20LockBox.getBalance(address(s_token), chainSelector1), 0);

    // Verify second chain selector balance is unchanged
    assertEq(s_erc20LockBox.getBalance(address(s_token), chainSelector2), amount2);

    // Withdraw from second chain selector
    s_erc20LockBox.withdraw(address(s_token), amount2, s_recipient, chainSelector2);

    vm.stopPrank();

    // Verify final balances
    assertEq(s_token.balanceOf(s_recipient), recipientBalanceBefore + amount1 + amount2);
    assertEq(s_erc20LockBox.getBalance(address(s_token), chainSelector1), 0);
    assertEq(s_erc20LockBox.getBalance(address(s_token), chainSelector2), 0);
  }

  function test_Withdraw_PartialAmount() public {
    uint256 depositAmount = 1000e18;
    uint256 withdrawAmount = 300e18;

    // Deposit tokens
    _depositTokens(depositAmount, DEST_CHAIN_SELECTOR);

    uint256 recipientBalanceBefore = s_token.balanceOf(s_recipient);
    uint256 expectedRemainingBalance = depositAmount - withdrawAmount;

    vm.startPrank(s_allowedCaller);

    s_erc20LockBox.withdraw(address(s_token), withdrawAmount, s_recipient, DEST_CHAIN_SELECTOR);

    vm.stopPrank();

    // Verify balances
    assertEq(s_token.balanceOf(s_recipient), recipientBalanceBefore + withdrawAmount);
    assertEq(s_erc20LockBox.getBalance(address(s_token), DEST_CHAIN_SELECTOR), expectedRemainingBalance);
  }

  // Reverts
  function test_RevertWhen_Unauthorized() public {
    uint256 amount = 1000e18;
    _depositTokens(amount, DEST_CHAIN_SELECTOR);

    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(ERC20LockBox.Unauthorized.selector, STRANGER));

    s_erc20LockBox.withdraw(address(s_token), amount, s_recipient, DEST_CHAIN_SELECTOR);
  }

  function test_RevertWhen_RecipientIsZeroAddress() public {
    uint256 amount = 1000e18;
    _depositTokens(amount, DEST_CHAIN_SELECTOR);

    vm.startPrank(s_allowedCaller);
    vm.expectRevert(ERC20LockBox.RecipientCannotBeZeroAddress.selector);

    s_erc20LockBox.withdraw(address(s_token), amount, address(0), DEST_CHAIN_SELECTOR);
  }

  function test_RevertWhen_AmountIsZero() public {
    uint256 amount = 1000e18;
    _depositTokens(amount, DEST_CHAIN_SELECTOR);

    vm.startPrank(s_allowedCaller);
    vm.expectRevert(ERC20LockBox.TokenAmountCannotBeZero.selector);

    s_erc20LockBox.withdraw(address(s_token), 0, s_recipient, DEST_CHAIN_SELECTOR);
  }

  function test_RevertWhen_TokenIsZeroAddress() public {
    uint256 amount = 1000e18;
    _depositTokens(amount, DEST_CHAIN_SELECTOR);

    vm.startPrank(s_allowedCaller);
    vm.expectRevert(ERC20LockBox.TokenAddressCannotBeZero.selector);

    s_erc20LockBox.withdraw(address(0), amount, s_recipient, DEST_CHAIN_SELECTOR);
  }

  function test_RevertWhen_InsufficientBalance() public {
    uint256 depositAmount = 1000e18;
    uint256 withdrawAmount = 1500e18;

    _depositTokens(depositAmount, DEST_CHAIN_SELECTOR);

    vm.startPrank(s_allowedCaller);
    vm.expectRevert(
      abi.encodeWithSelector(
        ERC20LockBox.InsufficientBalance.selector, DEST_CHAIN_SELECTOR, withdrawAmount, depositAmount
      )
    );

    s_erc20LockBox.withdraw(address(s_token), withdrawAmount, s_recipient, DEST_CHAIN_SELECTOR);
  }

  function test_RevertWhen_ChainSelectorHasNoBalance() public {
    uint64 emptyChainSelector = SOURCE_CHAIN_SELECTOR + DEST_CHAIN_SELECTOR + 1;

    vm.startPrank(s_allowedCaller);
    vm.expectRevert(abi.encodeWithSelector(ERC20LockBox.InsufficientBalance.selector, emptyChainSelector, 1, 0));

    s_erc20LockBox.withdraw(address(s_token), 1, s_recipient, emptyChainSelector);
  }

  function test_Withdraw_EventEmission() public {
    uint256 amount = 1000e18;
    _depositTokens(amount, DEST_CHAIN_SELECTOR);

    vm.startPrank(s_allowedCaller);

    vm.expectEmit(true, true, true, true);
    emit ERC20LockBox.Withdrawal(DEST_CHAIN_SELECTOR, s_recipient, amount);

    s_erc20LockBox.withdraw(address(s_token), amount, s_recipient, DEST_CHAIN_SELECTOR);

    vm.stopPrank();
  }

  function test_Withdraw_ToDifferentRecipients() public {
    uint256 amount = 1000e18;
    address recipient1 = makeAddr("recipient1");
    address recipient2 = makeAddr("recipient2");

    _depositTokens(amount * 2, DEST_CHAIN_SELECTOR);

    vm.startPrank(s_allowedCaller);

    // Withdraw to first recipient
    s_erc20LockBox.withdraw(address(s_token), amount, recipient1, DEST_CHAIN_SELECTOR);

    // Withdraw to second recipient
    s_erc20LockBox.withdraw(address(s_token), amount, recipient2, DEST_CHAIN_SELECTOR);

    vm.stopPrank();

    // Verify both recipients received tokens
    assertEq(s_token.balanceOf(recipient1), amount);
    assertEq(s_token.balanceOf(recipient2), amount);
    assertEq(s_erc20LockBox.getBalance(address(s_token), DEST_CHAIN_SELECTOR), 0);
  }
}
