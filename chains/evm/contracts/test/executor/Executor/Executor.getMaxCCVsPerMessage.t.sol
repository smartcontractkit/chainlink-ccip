// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Executor} from "../../../executor/Executor.sol";
import {ExecutorSetup} from "./ExecutorSetup.t.sol";

contract Executor_getMaxCCVsPerMessage is ExecutorSetup {
  function test_getMaxCCVsPerMessage() public {
    uint8 maxCCVs = s_executor.getMaxCCVsPerMessage();
    assertEq(INITIAL_MAX_CCVS, maxCCVs);

    uint8 newMaxCCVs = INITIAL_MAX_CCVS + 5;
    Executor.DynamicConfig memory dynamicConfig = Executor.DynamicConfig({
      feeAggregator: FEE_AGGREGATOR,
      minBlockConfirmations: MIN_BLOCK_CONFIRMATIONS,
      ccvAllowlistEnabled: false
    });
    s_executor = new Executor(newMaxCCVs, dynamicConfig);
    maxCCVs = s_executor.getMaxCCVsPerMessage();
    assertEq(newMaxCCVs, maxCCVs);
  }

  function test_constructor_RevertWhen_InvalidMaxPossibleCCVsPerMsg() public {
    Executor.DynamicConfig memory dynamicConfig =
      Executor.DynamicConfig({feeAggregator: FEE_AGGREGATOR, minBlockConfirmations: 0, ccvAllowlistEnabled: false});

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidMaxPossibleCCVsPerMsg.selector, 0));
    new Executor(0, dynamicConfig);
  }
}
