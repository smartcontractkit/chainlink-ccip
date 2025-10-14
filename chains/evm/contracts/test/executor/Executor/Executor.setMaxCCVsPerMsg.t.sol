// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Executor} from "../../../executor/Executor.sol";
import {ExecutorSetup} from "./ExecutorSetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract Executor_setMaxCCVsPerMsg is ExecutorSetup {
  function test_setMaxCCVsPerMsg() public {
    uint8 maxCCVsPerMsg = 2;

    vm.expectEmit();
    emit Executor.MaxCCVsPerMsgSet(maxCCVsPerMsg);
    s_Executor.setMaxCCVsPerMsg(maxCCVsPerMsg);

    assertEq(maxCCVsPerMsg, s_Executor.getMaxCCVsPerMsg());
  }

  function test_setMaxCCVsPerMsg_RevertWhen_NotOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_Executor.setMaxCCVsPerMsg(1);
  }

  function test_setMaxCCVsPerMsg_RevertWhen_InvalidMaxCCVsPerMsg() public {
    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidMaxPossibleCCVsPerMsg.selector, 0));
    s_Executor.setMaxCCVsPerMsg(0);
  }
}
