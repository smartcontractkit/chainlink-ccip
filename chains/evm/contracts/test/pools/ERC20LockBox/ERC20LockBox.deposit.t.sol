// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {ERC20LockBoxSetup} from "./ERC20LockBoxSetup.t.sol";

contract ERC20LockBox_deposit is ERC20LockBoxSetup {
  function testFuzz_Deposit_Success(
    uint256 amount
  ) public {
    amount = bound(amount, 1, type(uint256).max / 2);

    uint256 lockBoxBalanceBefore = s_token.balanceOf(address(s_erc20LockBox));
    uint256 callerBalanceBefore = s_token.balanceOf(s_allowedCaller);

    vm.startPrank(s_allowedCaller);
    s_token.approve(address(s_erc20LockBox), amount);

    vm.expectEmit();
    emit ERC20LockBox.Deposit(address(s_token), s_allowedCaller, amount);

    s_erc20LockBox.deposit(address(s_token), amount);

    vm.stopPrank();

    // Verify balances
    assertEq(s_token.balanceOf(address(s_erc20LockBox)), lockBoxBalanceBefore + amount);
    assertEq(s_token.balanceOf(s_allowedCaller), callerBalanceBefore - amount);
  }

  function test_Deposit_WithMultipleChainSelectors() public {
    uint256 amount1 = 1000e18;
    uint256 amount2 = 2000e18;

    vm.startPrank(s_allowedCaller);
    s_token.approve(address(s_erc20LockBox), amount1 + amount2);

    // Deposit tokens
    s_erc20LockBox.deposit(address(s_token), amount1);
    s_erc20LockBox.deposit(address(s_token), amount2);

    vm.stopPrank();

    // Verify balances
    assertEq(s_token.balanceOf(address(s_erc20LockBox)), amount1 + amount2);
  }

  function test_Deposit_MultipleDepositsToSameChain() public {
    uint256 amount1 = 1000e18;
    uint256 amount2 = 2000e18;

    vm.startPrank(s_allowedCaller);
    s_token.approve(address(s_erc20LockBox), amount1 + amount2);

    // First deposit
    s_erc20LockBox.deposit(address(s_token), amount1);

    // Second deposit to same chain
    s_erc20LockBox.deposit(address(s_token), amount2);

    vm.stopPrank();

    // Verify total balance
    assertEq(s_token.balanceOf(address(s_erc20LockBox)), amount1 + amount2);
  }

  function test_Deposit_EventEmission() public {
    uint256 amount = 1000e18;

    vm.startPrank(s_allowedCaller);
    s_token.approve(address(s_erc20LockBox), amount);

    vm.expectEmit(true, true, true, true);
    emit ERC20LockBox.Deposit(address(s_token), s_allowedCaller, amount);

    s_erc20LockBox.deposit(address(s_token), amount);

    vm.stopPrank();
  }

  function test_Deposit_FromDifferentCallers() public {
    uint256 amount = 1000e18;
    address caller1 = makeAddr("caller1");
    address caller2 = makeAddr("caller2");

    // Give tokens to both callers
    deal(address(s_token), caller1, amount);
    deal(address(s_token), caller2, amount);

    // Configure both callers as allowed
    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](2);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({token: address(s_token), caller: caller1, allowed: true});
    configArgs[1] = ERC20LockBox.AllowedCallerConfigArgs({token: address(s_token), caller: caller2, allowed: true});
    s_erc20LockBox.configureAllowedCallers(configArgs);

    // First caller deposits
    vm.startPrank(caller1);
    s_token.approve(address(s_erc20LockBox), amount);
    s_erc20LockBox.deposit(address(s_token), amount);
    vm.stopPrank();

    // Second caller deposits
    vm.startPrank(caller2);
    s_token.approve(address(s_erc20LockBox), amount);
    s_erc20LockBox.deposit(address(s_token), amount);
    vm.stopPrank();

    // Verify balances
    assertEq(s_token.balanceOf(address(s_erc20LockBox)), amount * 2);
  }

  function test_Deposit_ChainSelectorZero() public {
    uint256 amount = 1000e18;

    vm.startPrank(s_allowedCaller);
    s_token.approve(address(s_erc20LockBox), amount);

    s_erc20LockBox.deposit(address(s_token), amount);

    vm.stopPrank();
  }

  function test_Deposit_MaxAmount() public {
    uint256 maxAmount = type(uint256).max;

    // Give max tokens to caller
    deal(address(s_token), s_allowedCaller, maxAmount);

    vm.startPrank(s_allowedCaller);
    s_token.approve(address(s_erc20LockBox), maxAmount);

    s_erc20LockBox.deposit(address(s_token), maxAmount);

    vm.stopPrank();

    assertEq(s_token.balanceOf(address(s_erc20LockBox)), maxAmount);
  }

  // ================================================================
  // │                        Revert Tests                          │
  // ================================================================

  function test_RevertWhen_Unauthorized() public {
    uint256 amount = 1000e18;

    vm.startPrank(STRANGER);
    s_token.approve(address(s_erc20LockBox), amount);
    vm.expectRevert(abi.encodeWithSelector(ERC20LockBox.Unauthorized.selector, STRANGER));

    s_erc20LockBox.deposit(address(s_token), amount);
  }

  function test_RevertWhen_AmountIsZero() public {
    vm.startPrank(s_allowedCaller);
    s_token.approve(address(s_erc20LockBox), 1);
    vm.expectRevert(ERC20LockBox.TokenAmountCannotBeZero.selector);

    s_erc20LockBox.deposit(address(s_token), 0);
  }

  function test_RevertWhen_TokenIsZeroAddress() public {
    uint256 amount = 1000e18;

    vm.startPrank(s_allowedCaller);
    s_token.approve(address(s_erc20LockBox), amount);
    vm.expectRevert(ERC20LockBox.TokenAddressCannotBeZero.selector);

    s_erc20LockBox.deposit(address(0), amount);
  }

  function test_RevertWhen_InsufficientAllowance() public {
    uint256 amount = 1000e18;

    vm.startPrank(s_allowedCaller);
    s_token.approve(address(s_erc20LockBox), amount - 1); // Approve less than amount
    vm.expectRevert("ERC20: insufficient allowance");

    s_erc20LockBox.deposit(address(s_token), amount);
  }

  function test_RevertWhen_NoAllowance() public {
    uint256 amount = 1000e18;

    vm.startPrank(s_allowedCaller);
    vm.expectRevert("ERC20: insufficient allowance");

    s_erc20LockBox.deposit(address(s_token), amount);
  }
}
