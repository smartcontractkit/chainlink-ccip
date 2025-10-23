// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Executor} from "../../../executor/Executor.sol";
import {ExecutorSetup} from "./ExecutorSetup.t.sol";

contract Executor_getMaxCCVsPerMessage is ExecutorSetup {
  function test_getMaxCCVsPerMessage() public {
    uint8 maxCCVs = s_executor.getMaxCCVsPerMessage();
    assertEq(maxCCVs, INITIAL_MAX_CCVS);

    uint8 newMaxCCVs = INITIAL_MAX_CCVS + 5;
    s_executor = new Executor(newMaxCCVs, 0);
    maxCCVs = s_executor.getMaxCCVsPerMessage();
    assertEq(maxCCVs, newMaxCCVs);
  }

  function test_constructor_RevertWhen_InvalidMaxPossibleCCVsPerMsg() public {
    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidMaxPossibleCCVsPerMsg.selector, 0));

    new Executor(0, 0);
  }
}
