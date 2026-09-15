// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeTokenHandler} from "../../../libraries/FeeTokenHandler.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract SuccinctZKVerifier_withdrawFeeTokens is SuccinctZKVerifierSetup {
  function test_withdrawFeeTokens() public {
    uint256 feeAmount = 1000 ether;
    deal(s_sourceFeeToken, address(s_zkVerifier), feeAmount);
    uint256 initialAggregatorBalance = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceFeeToken;

    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(FEE_AGGREGATOR, s_sourceFeeToken, feeAmount);

    // Anyone can call withdrawFeeTokens since it's permissionless.
    vm.stopPrank();
    vm.prank(STRANGER);
    s_zkVerifier.withdrawFeeTokens(feeTokens);

    assertEq(0, IERC20(s_sourceFeeToken).balanceOf(address(s_zkVerifier)));
    assertEq(initialAggregatorBalance + feeAmount, IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR));
  }
}
