// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ExecutorOnRamp} from "../../../onRamp/ExecutorOnRamp.sol";
import {ExecutorOnRampSetup} from "./ExecutorOnRampSetup.t.sol";

contract ExecutorOnRamp_setDynamicConfig is ExecutorOnRampSetup {
  function setUp() public virtual override {
    super.setUp();
  }

  function test_setDynamicConfig_RevertWhen_NotOwner() public {
    vm.stopPrank();
    vm.startPrank(makeAddr("notOwner"));
    ExecutorOnRamp.DynamicConfig memory newConfig = ExecutorOnRamp.DynamicConfig({
      feeQuoter: makeAddr("newFeeQuoter"),
      feeAggregator: makeAddr("newFeeAggregator"),
      maxPossibleCCVsPerMsg: 5,
      maxRequiredCCVsPerMsg: 3
    });
    vm.stopPrank();
    vm.expectRevert("OnlyCallableByOwner()");
    s_executorOnRamp.setDynamicConfig(newConfig);
  }

  function test_SetDynamicConfig_Success() public {
    ExecutorOnRamp.DynamicConfig memory newConfig = ExecutorOnRamp.DynamicConfig({
      feeQuoter: makeAddr("newFeeQuoter"),
      feeAggregator: makeAddr("newFeeAggregator"),
      maxPossibleCCVsPerMsg: 5,
      maxRequiredCCVsPerMsg: 3
    });
    s_executorOnRamp.setDynamicConfig(newConfig);
    ExecutorOnRamp.DynamicConfig memory currentConfig = s_executorOnRamp.getDynamicConfig();
    assertEq(currentConfig.feeQuoter, newConfig.feeQuoter);
    assertEq(currentConfig.feeAggregator, newConfig.feeAggregator);
    assertEq(currentConfig.maxPossibleCCVsPerMsg, newConfig.maxPossibleCCVsPerMsg);
    assertEq(currentConfig.maxRequiredCCVsPerMsg, newConfig.maxRequiredCCVsPerMsg);
  }
}
