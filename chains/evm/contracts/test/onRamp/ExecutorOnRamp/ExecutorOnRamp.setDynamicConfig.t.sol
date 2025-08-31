// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ExecutorOnRamp} from "../../../onRamp/ExecutorOnRamp.sol";
import {ExecutorOnRampSetup} from "./ExecutorOnRampSetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract ExecutorOnRamp_setDynamicConfig is ExecutorOnRampSetup {
  function test_setDynamicConfig() public {
    ExecutorOnRamp.DynamicConfig memory newConfig;
    newConfig.maxCCVsPerMsg = 2;

    vm.expectEmit();
    emit ExecutorOnRamp.ConfigSet(newConfig);
    s_executorOnRamp.setDynamicConfig(newConfig);

    ExecutorOnRamp.DynamicConfig memory currentConfig = s_executorOnRamp.getDynamicConfig();
    assertEq(currentConfig.maxCCVsPerMsg, newConfig.maxCCVsPerMsg);
  }

  function test_setDynamicConfig_RevertWhen_NotOwner() public {
    ExecutorOnRamp.DynamicConfig memory newConfig;
    vm.prank(makeAddr("stranger"));

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_executorOnRamp.setDynamicConfig(newConfig);
  }

  function test_setDynamicConfig_RevertWhen_InvalidMaxCCVsPerMsg() public {
    ExecutorOnRamp.DynamicConfig memory newConfig;
    newConfig.maxCCVsPerMsg = 0; // Invalid, must be > 0

    vm.expectRevert(abi.encodeWithSelector(ExecutorOnRamp.InvalidMaxPossibleCCVsPerMsg.selector, 0));
    s_executorOnRamp.setDynamicConfig(newConfig);
  }
}
