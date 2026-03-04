// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeTokenHandler} from "../../../../libraries/FeeTokenHandler.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract USDCTokenPoolProxy_withdrawFeeTokens is USDCTokenPoolProxySetup {
  address internal s_feeAggregator = makeAddr("feeAggregator");
  BurnMintERC20 internal s_feeToken1;
  BurnMintERC20 internal s_feeToken2;

  function setUp() public override {
    super.setUp();

    s_feeToken1 = new BurnMintERC20("FeeToken1", "FT1", 18, 0, 0);
    s_feeToken2 = new BurnMintERC20("FeeToken2", "FT2", 18, 0, 0);
  }

  function test_withdrawFeeTokens() public {
    s_usdcTokenPoolProxy.setFeeAggregator(s_feeAggregator);

    uint256 token1Amount = 1000e18;
    uint256 token2Amount = 500e18;

    deal(address(s_feeToken1), address(s_usdcTokenPoolProxy), token1Amount);
    deal(address(s_feeToken2), address(s_usdcTokenPoolProxy), token2Amount);

    address[] memory feeTokens = new address[](2);
    feeTokens[0] = address(s_feeToken1);
    feeTokens[1] = address(s_feeToken2);

    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(s_feeAggregator, address(s_feeToken1), token1Amount);
    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(s_feeAggregator, address(s_feeToken2), token2Amount);

    s_usdcTokenPoolProxy.withdrawFeeTokens(feeTokens);

    assertEq(s_feeToken1.balanceOf(s_feeAggregator), token1Amount);
    assertEq(s_feeToken2.balanceOf(s_feeAggregator), token2Amount);
    assertEq(s_feeToken1.balanceOf(address(s_usdcTokenPoolProxy)), 0);
    assertEq(s_feeToken2.balanceOf(address(s_usdcTokenPoolProxy)), 0);
  }

  function test_withdrawFeeTokens_CalledByNonOwner() public {
    s_usdcTokenPoolProxy.setFeeAggregator(s_feeAggregator);

    uint256 tokenAmount = 1000e18;
    deal(address(s_feeToken1), address(s_usdcTokenPoolProxy), tokenAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = address(s_feeToken1);

    address nonOwner = makeAddr("nonOwner");

    vm.expectEmit();
    emit FeeTokenHandler.FeeTokenWithdrawn(s_feeAggregator, address(s_feeToken1), tokenAmount);

    vm.startPrank(nonOwner);
    s_usdcTokenPoolProxy.withdrawFeeTokens(feeTokens);

    assertEq(s_feeToken1.balanceOf(s_feeAggregator), tokenAmount);
  }

  // Reverts

  function test_withdrawFeeTokens_RevertWhen_FeeAggregatorNotSet() public {
    uint256 tokenAmount = 1000e18;
    deal(address(s_feeToken1), address(s_usdcTokenPoolProxy), tokenAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = address(s_feeToken1);

    vm.expectRevert(FeeTokenHandler.ZeroAddressNotAllowed.selector);
    s_usdcTokenPoolProxy.withdrawFeeTokens(feeTokens);
  }
}

