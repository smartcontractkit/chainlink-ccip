// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {CommitteeVerifierSetup} from "./CommitteeVerifierSetup.t.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract CommitteeVerifier_withdrawFeeTokens is CommitteeVerifierSetup {
  function setUp() public override {
    super.setUp();
    vm.stopPrank();
  }

  function test_withdrawFeeTokens() public {
    uint256 feeAmount = 1000 ether;

    // Give the verifier some fee tokens.
    deal(s_sourceFeeToken, address(s_committeeVerifier), feeAmount);

    uint256 initialAggregatorBalance = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    uint256 initialVerifierBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_committeeVerifier));

    assertEq(feeAmount, initialVerifierBalance);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceFeeToken;

    vm.expectEmit();
    emit BaseVerifier.FeeTokenWithdrawn(FEE_AGGREGATOR, s_sourceFeeToken, feeAmount);

    // Anyone can call withdrawFeeTokens since it's permissionless.
    vm.prank(STRANGER);
    s_committeeVerifier.withdrawFeeTokens(feeTokens);

    uint256 finalAggregatorBalance = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    uint256 finalVerifierBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_committeeVerifier));

    assertEq(0, finalVerifierBalance);
    assertEq(initialAggregatorBalance + feeAmount, finalAggregatorBalance);
  }

  function test_withdrawFeeTokens_MultipleTokens() public {
    uint256 feeAmount1 = 1000 ether;
    uint256 feeAmount2 = 500 ether;

    address token2 = address(new BurnMintERC20("Token2", "TK2", 18, 0, 0));

    // Give the verifier some fee tokens.
    deal(s_sourceFeeToken, address(s_committeeVerifier), feeAmount1);
    deal(token2, address(s_committeeVerifier), feeAmount2);

    uint256 initialAggregatorBalance1 = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    uint256 initialAggregatorBalance2 = IERC20(token2).balanceOf(FEE_AGGREGATOR);

    address[] memory feeTokens = new address[](2);
    feeTokens[0] = s_sourceFeeToken;
    feeTokens[1] = token2;

    vm.expectEmit();
    emit BaseVerifier.FeeTokenWithdrawn(FEE_AGGREGATOR, s_sourceFeeToken, feeAmount1);
    vm.expectEmit();
    emit BaseVerifier.FeeTokenWithdrawn(FEE_AGGREGATOR, token2, feeAmount2);

    s_committeeVerifier.withdrawFeeTokens(feeTokens);

    uint256 finalAggregatorBalance1 = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    uint256 finalAggregatorBalance2 = IERC20(token2).balanceOf(FEE_AGGREGATOR);

    assertEq(0, IERC20(s_sourceFeeToken).balanceOf(address(s_committeeVerifier)));
    assertEq(0, IERC20(token2).balanceOf(address(s_committeeVerifier)));
    assertEq(initialAggregatorBalance1 + feeAmount1, finalAggregatorBalance1);
    assertEq(initialAggregatorBalance2 + feeAmount2, finalAggregatorBalance2);
  }
}
