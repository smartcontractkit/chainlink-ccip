// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Executor} from "../../../Executor/Executor.sol";
import {ExecutorSetup} from "./ExecutorSetup.t.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract Executor_withdrawFeeTokens is ExecutorSetup {
  function test_withdrawFeeTokens() public {
    uint256 feeAmount = 1000 ether;

    // Give the Executor some fee tokens.
    deal(s_sourceFeeToken, address(s_executor), feeAmount);

    uint256 initialAggregatorBalance = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    uint256 initialExecutorBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_executor));

    assertEq(feeAmount, initialExecutorBalance);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceFeeToken;

    vm.expectEmit();
    emit Executor.FeeTokenWithdrawn(FEE_AGGREGATOR, s_sourceFeeToken, feeAmount);
    s_executor.withdrawFeeTokens(feeTokens);

    uint256 finalAggregatorBalance = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    uint256 finalExecutorBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_executor));

    assertEq(0, finalExecutorBalance);
    assertEq(initialAggregatorBalance + feeAmount, finalAggregatorBalance);
  }

  function test_withdrawFeeTokens_MultipleTokens() public {
    uint256 feeAmount1 = 1000 ether;
    uint256 feeAmount2 = 500 ether;

    address token2 = address(new BurnMintERC20("Token2", "TK2", 18, 0, 0));

    // Give the Executor some fee tokens.
    deal(s_sourceFeeToken, address(s_executor), feeAmount1);
    deal(token2, address(s_executor), feeAmount2);

    uint256 initialAggregatorBalance1 = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    uint256 initialAggregatorBalance2 = IERC20(token2).balanceOf(FEE_AGGREGATOR);

    address[] memory feeTokens = new address[](2);
    feeTokens[0] = s_sourceFeeToken;
    feeTokens[1] = token2;

    vm.expectEmit();
    emit Executor.FeeTokenWithdrawn(FEE_AGGREGATOR, s_sourceFeeToken, feeAmount1);
    vm.expectEmit();
    emit Executor.FeeTokenWithdrawn(FEE_AGGREGATOR, token2, feeAmount2);
    s_executor.withdrawFeeTokens(feeTokens);

    uint256 finalAggregatorBalance1 = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    uint256 finalAggregatorBalance2 = IERC20(token2).balanceOf(FEE_AGGREGATOR);

    assertEq(0, IERC20(s_sourceFeeToken).balanceOf(address(s_executor)));
    assertEq(0, IERC20(token2).balanceOf(address(s_executor)));
    assertEq(initialAggregatorBalance1 + feeAmount1, finalAggregatorBalance1);
    assertEq(initialAggregatorBalance2 + feeAmount2, finalAggregatorBalance2);
  }
}
