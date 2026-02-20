// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeTokenHandler} from "../../../libraries/FeeTokenHandler.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract TokenPool_withdrawFee is AdvancedPoolHooksSetup {
  address internal s_feeAdmin;

  function setUp() public override {
    super.setUp();
    vm.stopPrank();
    s_feeAdmin = makeAddr("fee_admin");

    vm.prank(OWNER);
    s_tokenPool.setDynamicConfig(makeAddr("router"), makeAddr("rateLimitAdmin"), s_feeAdmin);
  }

  function test_withdrawFeeTokens() public {
    uint256 feeAmount = 20 ether;
    address recipient = makeAddr("fee_recipient");

    deal(address(s_token), address(s_tokenPool), feeAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = address(s_token);

    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(recipient, address(s_token), feeAmount);

    vm.prank(s_feeAdmin);
    s_tokenPool.withdrawFeeTokens(feeTokens, recipient);

    assertEq(s_token.balanceOf(recipient), feeAmount);
    assertEq(s_token.balanceOf(address(s_tokenPool)), 0);
  }

  function test_withdrawFeeTokens_MultipleTokens() public {
    uint256 feeAmount1 = 20 ether;
    uint256 feeAmount2 = 10 ether;
    address recipient = makeAddr("fee_recipient");

    address token2 = address(new BurnMintERC20("Token2", "TK2", 18, 0, 0));

    deal(address(s_token), address(s_tokenPool), feeAmount1);
    deal(token2, address(s_tokenPool), feeAmount2);

    address[] memory feeTokens = new address[](2);
    feeTokens[0] = address(s_token);
    feeTokens[1] = token2;

    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(recipient, address(s_token), feeAmount1);
    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(recipient, token2, feeAmount2);

    vm.prank(s_feeAdmin);
    s_tokenPool.withdrawFeeTokens(feeTokens, recipient);

    assertEq(s_token.balanceOf(recipient), feeAmount1);
    assertEq(IERC20(token2).balanceOf(recipient), feeAmount2);
    assertEq(s_token.balanceOf(address(s_tokenPool)), 0);
    assertEq(IERC20(token2).balanceOf(address(s_tokenPool)), 0);
  }

  function test_withdrawFeeTokens_UpdateFeeAdmin() public {
    uint256 feeAmount1 = 20 ether;
    uint256 feeAmount2 = 10 ether;
    address newFeeAdmin = makeAddr("new_fee_admin");
    address recipient = makeAddr("fee_recipient");

    deal(address(s_token), address(s_tokenPool), feeAmount1);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = address(s_token);

    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(recipient, address(s_token), feeAmount1);
    vm.prank(s_feeAdmin);
    s_tokenPool.withdrawFeeTokens(feeTokens, recipient);

    assertEq(s_token.balanceOf(recipient), feeAmount1);

    // Add more fees and update fee admin
    deal(address(s_token), address(s_tokenPool), feeAmount2);
    vm.prank(OWNER);
    s_tokenPool.setDynamicConfig(makeAddr("router"), makeAddr("rateLimitAdmin"), newFeeAdmin);

    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(recipient, address(s_token), feeAmount2);
    vm.prank(newFeeAdmin);
    s_tokenPool.withdrawFeeTokens(feeTokens, recipient);

    assertEq(s_token.balanceOf(recipient) - feeAmount1, feeAmount2);
  }

  function test_withdrawFeeTokens_CallableByOwnerAndFeeAdmin() public {
    uint256 feeAmount = 20 ether;
    address recipient = makeAddr("fee_recipient");

    deal(address(s_token), address(s_tokenPool), feeAmount / 2);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = address(s_token);

    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(recipient, address(s_token), feeAmount / 2);
    vm.prank(s_feeAdmin);
    s_tokenPool.withdrawFeeTokens(feeTokens, recipient);

    deal(address(s_token), address(s_tokenPool), feeAmount / 2);

    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(recipient, address(s_token), feeAmount / 2);
    vm.prank(OWNER);
    s_tokenPool.withdrawFeeTokens(feeTokens, recipient);

    assertEq(s_token.balanceOf(recipient), feeAmount);
  }

  function test_withdrawFeeTokens_RevertsWhen_CalledByNonOwnerOrFeeAdmin() public {
    uint256 feeAmount = 20 ether;
    address recipient = makeAddr("fee_recipient");

    deal(address(s_token), address(s_tokenPool), feeAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = address(s_token);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.CallerIsNotOwnerOrFeeAdmin.selector, STRANGER));
    vm.prank(STRANGER);
    s_tokenPool.withdrawFeeTokens(feeTokens, recipient);
  }
}
