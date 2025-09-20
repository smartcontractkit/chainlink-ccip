// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseOnRamp} from "../../../../ccvs/components/BaseOnRamp.sol";
import {BaseOnRampSetup} from "./BaseOnRampSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract BaseOnRamp_withdrawFeeTokens is BaseOnRampSetup {
  function setUp() public override {
    super.setUp();
    vm.stopPrank();
  }

  function test_withdrawFeeTokens() public {
    uint256 feeAmount = 1000 ether;

    // Give the onRamp some fee tokens.
    deal(s_sourceFeeToken, address(s_baseOnRamp), feeAmount);

    uint256 initialAggregatorBalance = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    uint256 initialOnRampBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_baseOnRamp));

    assertEq(initialOnRampBalance, feeAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceFeeToken;

    vm.expectEmit();
    emit BaseOnRamp.FeeTokenWithdrawn(FEE_AGGREGATOR, s_sourceFeeToken, feeAmount);

    // Anyone can call withdrawFeeTokens.
    vm.prank(STRANGER);
    s_baseOnRamp.withdrawFeeTokens(feeTokens, FEE_AGGREGATOR);

    uint256 finalAggregatorBalance = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    uint256 finalOnRampBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_baseOnRamp));

    assertEq(finalOnRampBalance, 0);
    assertEq(finalAggregatorBalance, initialAggregatorBalance + feeAmount);
  }
}
