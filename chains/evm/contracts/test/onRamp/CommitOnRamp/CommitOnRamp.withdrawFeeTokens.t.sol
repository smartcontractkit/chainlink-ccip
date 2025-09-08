// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract CommitOnRamp_withdrawFeeTokens is CommitOnRampSetup {
  function setUp() public override {
    super.setUp();
    vm.stopPrank(); // Tests will manage their own pranks
  }

  function test_withdrawFeeTokens() public {
    uint256 feeAmount = 1000 ether;

    // Give the onRamp some fee tokens
    deal(s_sourceFeeToken, address(s_commitOnRamp), feeAmount);

    uint256 initialAggregatorBalance = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    uint256 initialOnRampBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_commitOnRamp));

    assertEq(initialOnRampBalance, feeAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceFeeToken;

    vm.expectEmit();
    emit BaseOnRamp.FeeTokenWithdrawn(FEE_AGGREGATOR, s_sourceFeeToken, feeAmount);

    // Anyone can call withdrawFeeTokens
    vm.prank(STRANGER);
    s_commitOnRamp.withdrawFeeTokens(feeTokens);

    uint256 finalAggregatorBalance = IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR);
    uint256 finalOnRampBalance = IERC20(s_sourceFeeToken).balanceOf(address(s_commitOnRamp));

    assertEq(finalOnRampBalance, 0);
    assertEq(finalAggregatorBalance, initialAggregatorBalance + feeAmount);
  }

  function test_withdrawFeeTokens_MultipleFeeTokens() public {
    // Use two different tokens
    address feeToken1 = s_sourceTokens[0];
    address feeToken2 = s_sourceTokens[1];
    uint256 feeAmount1 = 1000 ether;
    uint256 feeAmount2 = 500 ether;

    // Give the onRamp some fee tokens
    deal(feeToken1, address(s_commitOnRamp), feeAmount1);
    deal(feeToken2, address(s_commitOnRamp), feeAmount2);

    address[] memory feeTokens = new address[](2);
    feeTokens[0] = feeToken1;
    feeTokens[1] = feeToken2;

    vm.expectEmit();
    emit BaseOnRamp.FeeTokenWithdrawn(FEE_AGGREGATOR, feeToken1, feeAmount1);
    vm.expectEmit();
    emit BaseOnRamp.FeeTokenWithdrawn(FEE_AGGREGATOR, feeToken2, feeAmount2);

    s_commitOnRamp.withdrawFeeTokens(feeTokens);

    assertEq(IERC20(feeToken1).balanceOf(address(s_commitOnRamp)), 0);
    assertEq(IERC20(feeToken2).balanceOf(address(s_commitOnRamp)), 0);
    assertEq(IERC20(feeToken1).balanceOf(FEE_AGGREGATOR), feeAmount1);
    assertEq(IERC20(feeToken2).balanceOf(FEE_AGGREGATOR), feeAmount2);
  }

  function test_withdrawFeeTokens_PartialWithdrawal() public {
    uint256 feeAmount1 = 1000 ether;
    uint256 feeAmount2 = 500 ether;

    // Use two different tokens
    address token1 = s_sourceTokens[0];
    address token2 = s_sourceTokens[1];

    // Give the onRamp some fee tokens
    deal(token1, address(s_commitOnRamp), feeAmount1);
    deal(token2, address(s_commitOnRamp), feeAmount2);

    // Only withdraw one token
    address[] memory feeTokens = new address[](1);
    feeTokens[0] = token1;

    vm.expectEmit();
    emit BaseOnRamp.FeeTokenWithdrawn(FEE_AGGREGATOR, token1, feeAmount1);

    s_commitOnRamp.withdrawFeeTokens(feeTokens);

    // First token should be withdrawn
    assertEq(IERC20(token1).balanceOf(address(s_commitOnRamp)), 0);
    assertEq(IERC20(token1).balanceOf(FEE_AGGREGATOR), feeAmount1);

    // Second token should remain
    assertEq(IERC20(token2).balanceOf(address(s_commitOnRamp)), feeAmount2);
  }

  function test_withdrawFeeTokens_EmptyArray() public {
    uint256 feeAmount = 1000 ether;
    deal(s_sourceFeeToken, address(s_commitOnRamp), feeAmount);

    address[] memory feeTokens = new address[](0);

    // Should succeed but not withdraw anything
    s_commitOnRamp.withdrawFeeTokens(feeTokens);

    assertEq(IERC20(s_sourceFeeToken).balanceOf(address(s_commitOnRamp)), feeAmount);
  }

  function test_withdrawFeeTokens_PermissionlessExecution() public {
    uint256 feeAmount = 1000 ether;
    deal(s_sourceFeeToken, address(s_commitOnRamp), feeAmount);

    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceFeeToken;

    // Test that various actors can call withdrawFeeTokens
    address[] memory callers = new address[](3);
    callers[0] = OWNER;
    callers[1] = STRANGER;
    callers[2] = ALLOWLIST_ADMIN;

    for (uint256 i = 0; i < callers.length; ++i) {
      // Reset balance for each test
      deal(s_sourceFeeToken, address(s_commitOnRamp), feeAmount);

      vm.prank(callers[i]);
      s_commitOnRamp.withdrawFeeTokens(feeTokens);

      assertEq(IERC20(s_sourceFeeToken).balanceOf(address(s_commitOnRamp)), 0);
    }
  }
}
