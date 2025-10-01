// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../../ccvs/components/BaseVerifier.sol";
import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract BaseVerifier_withdrawFeeTokens is BaseVerifierSetup {
  function setUp() public override {
    super.setUp();
    vm.stopPrank();
  }

  function test_withdrawFeeTokens() public {
    uint256 feeAmount = 1000 ether;

    // Give the onRamp some fee tokens.
    deal(s_sourceFeeToken, address(s_baseVerifier), feeAmount);

    uint256 initialAggregatorBalance = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    uint256 initialOnRampBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_baseVerifier));

    assertEq(initialOnRampBalance, feeAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceFeeToken;

    vm.expectEmit();
    emit BaseVerifier.FeeTokenWithdrawn(FEE_AGGREGATOR, s_sourceFeeToken, feeAmount);

    // Anyone can call withdrawFeeTokens.
    vm.prank(STRANGER);
    s_baseVerifier.withdrawFeeTokens(feeTokens, FEE_AGGREGATOR);

    uint256 finalAggregatorBalance = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    uint256 finalOnRampBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_baseVerifier));

    assertEq(finalOnRampBalance, 0);
    assertEq(finalAggregatorBalance, initialAggregatorBalance + feeAmount);
  }
}
