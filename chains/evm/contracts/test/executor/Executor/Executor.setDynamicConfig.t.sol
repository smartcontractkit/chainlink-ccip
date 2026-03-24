// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Executor} from "../../../executor/Executor.sol";
import {ExecutorSetup} from "./ExecutorSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract Executor_setDynamicConfig is ExecutorSetup {
  function test_setDynamicConfig() public {
    Executor.DynamicConfig memory newConfig = Executor.DynamicConfig({
      feeAggregator: makeAddr("newFeeAggregator"),
      allowedFinalityConfig: bytes2(uint16(123)),
      ccvAllowlistEnabled: false
    });

    vm.expectEmit();
    emit Executor.ConfigSet(newConfig);
    s_executor.setDynamicConfig(newConfig);

    Executor.DynamicConfig memory config = s_executor.getDynamicConfig();
    assertEq(newConfig.feeAggregator, config.feeAggregator);
    assertEq(newConfig.allowedFinalityConfig, config.allowedFinalityConfig);
    assertEq(newConfig.ccvAllowlistEnabled, config.ccvAllowlistEnabled);

    // Verify getAllowedFinalityConfig getter works
    assertEq(s_executor.getAllowedFinalityConfig(), newConfig.allowedFinalityConfig);
  }

  function test_setDynamicConfig_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();

    Executor.DynamicConfig memory newConfig;

    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector));
    s_executor.setDynamicConfig(newConfig);
  }
}
