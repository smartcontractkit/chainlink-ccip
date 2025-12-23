// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeTokenHandler} from "../../../libraries/FeeTokenHandler.sol";
import {VersionedVerifierResolverSetup} from "./VersionedVerifierResolverSetup.t.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract VersionedVerifierResolver_withdrawFeeTokens is VersionedVerifierResolverSetup {
  address internal s_feeToken;
  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;

  function setUp() public override {
    super.setUp();
    s_feeToken = address(new BurnMintERC20("Chainlink Token", "LINK", 18, 0, 0));
    // Set the fee aggregator for the resolver
    s_versionedVerifierResolver.setFeeAggregator(FEE_AGGREGATOR);
    vm.stopPrank();
  }

  function test_withdrawFeeTokens() public {
    uint256 feeAmount = 1000e18;

    // Give the resolver some fee tokens.
    deal(s_feeToken, address(s_versionedVerifierResolver), feeAmount);

    uint256 initialAggregatorBalance = IERC20(s_feeToken).balanceOf(FEE_AGGREGATOR);
    uint256 initialResolverBalance = IERC20(s_feeToken).balanceOf(address(s_versionedVerifierResolver));

    assertEq(feeAmount, initialResolverBalance);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_feeToken;

    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(FEE_AGGREGATOR, s_feeToken, feeAmount);

    // Anyone can call withdrawFeeTokens since it's permissionless.
    vm.prank(STRANGER);
    s_versionedVerifierResolver.withdrawFeeTokens(feeTokens);

    uint256 finalAggregatorBalance = IERC20(s_feeToken).balanceOf(FEE_AGGREGATOR);
    uint256 finalResolverBalance = IERC20(s_feeToken).balanceOf(address(s_versionedVerifierResolver));

    assertEq(0, finalResolverBalance);
    assertEq(initialAggregatorBalance + feeAmount, finalAggregatorBalance);
  }

  function test_withdrawFeeTokens_MultipleTokens() public {
    uint256 feeAmount1 = 1000e18;
    uint256 feeAmount2 = 500e18;

    address token2 = address(new BurnMintERC20("Token2", "TK2", 18, 0, 0));

    // Give the resolver some fee tokens.
    deal(s_feeToken, address(s_versionedVerifierResolver), feeAmount1);
    deal(token2, address(s_versionedVerifierResolver), feeAmount2);

    uint256 initialAggregatorBalance1 = IERC20(s_feeToken).balanceOf(FEE_AGGREGATOR);
    uint256 initialAggregatorBalance2 = IERC20(token2).balanceOf(FEE_AGGREGATOR);

    address[] memory feeTokens = new address[](2);
    feeTokens[0] = s_feeToken;
    feeTokens[1] = token2;

    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(FEE_AGGREGATOR, s_feeToken, feeAmount1);
    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(FEE_AGGREGATOR, token2, feeAmount2);

    s_versionedVerifierResolver.withdrawFeeTokens(feeTokens);

    uint256 finalAggregatorBalance1 = IERC20(s_feeToken).balanceOf(FEE_AGGREGATOR);
    uint256 finalAggregatorBalance2 = IERC20(token2).balanceOf(FEE_AGGREGATOR);

    assertEq(0, IERC20(s_feeToken).balanceOf(address(s_versionedVerifierResolver)));
    assertEq(0, IERC20(token2).balanceOf(address(s_versionedVerifierResolver)));
    assertEq(initialAggregatorBalance1 + feeAmount1, finalAggregatorBalance1);
    assertEq(initialAggregatorBalance2 + feeAmount2, finalAggregatorBalance2);
  }
}

