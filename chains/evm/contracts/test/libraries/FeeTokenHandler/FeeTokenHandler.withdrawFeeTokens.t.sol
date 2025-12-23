// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeTokenHandler} from "../../../libraries/FeeTokenHandler.sol";
import {BaseTest} from "../../BaseTest.t.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract FeeTokenHandlerTestHarness {
  function withdrawFeeTokens(
    address[] calldata feeTokens,
    address feeAggregator
  ) external {
    FeeTokenHandler._withdrawFeeTokens(feeTokens, feeAggregator);
  }
}

contract FeeTokenHandler_withdrawFeeTokens is BaseTest {
  FeeTokenHandlerTestHarness internal s_harness;
  address internal s_feeToken;

  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;

  function setUp() public override {
    super.setUp();
    s_harness = new FeeTokenHandlerTestHarness();
    s_feeToken = address(new BurnMintERC20("FeeToken", "FEE", 18, 0, 0));
    // BaseTest leaves an OWNER prank active; stop it so tests can freely prank other senders.
    vm.stopPrank();
  }

  function test_withdrawFeeTokens() public {
    uint256 feeAmount = 1000 ether;

    // Give the harness some fee tokens.
    deal(s_feeToken, address(s_harness), feeAmount);

    uint256 initialAggregatorBalance = IERC20(s_feeToken).balanceOf(FEE_AGGREGATOR);
    uint256 initialHarnessBalance = IERC20(s_feeToken).balanceOf(address(s_harness));

    assertEq(initialHarnessBalance, feeAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_feeToken;

    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(FEE_AGGREGATOR, s_feeToken, feeAmount);

    // Anyone can call withdrawFeeTokens.
    vm.prank(STRANGER);
    s_harness.withdrawFeeTokens(feeTokens, FEE_AGGREGATOR);

    uint256 finalAggregatorBalance = IERC20(s_feeToken).balanceOf(FEE_AGGREGATOR);
    uint256 finalHarnessBalance = IERC20(s_feeToken).balanceOf(address(s_harness));

    assertEq(finalHarnessBalance, 0);
    assertEq(finalAggregatorBalance, initialAggregatorBalance + feeAmount);
  }
}
