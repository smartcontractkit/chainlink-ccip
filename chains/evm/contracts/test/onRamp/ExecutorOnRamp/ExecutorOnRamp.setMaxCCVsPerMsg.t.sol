// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ExecutorOnRamp} from "../../../onRamp/ExecutorOnRamp.sol";
import {ExecutorOnRampSetup} from "./ExecutorOnRampSetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract ExecutorOnRamp_setMaxCCVsPerMsg is ExecutorOnRampSetup {
  function test_setMaxCCVsPerMsg() public {
    uint8 maxCCVsPerMsg = 2;

    vm.expectEmit();
    emit ExecutorOnRamp.MaxCCVsPerMsgSet(maxCCVsPerMsg);
    s_executorOnRamp.setMaxCCVsPerMsg(maxCCVsPerMsg);

    assertEq(maxCCVsPerMsg, s_executorOnRamp.getMaxCCVsPerMsg());
  }

  function test_setMaxCCVsPerMsg_RevertWhen_NotOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_executorOnRamp.setMaxCCVsPerMsg(1);
  }

  function test_setMaxCCVsPerMsg_RevertWhen_InvalidMaxCCVsPerMsg() public {
    vm.expectRevert(abi.encodeWithSelector(ExecutorOnRamp.InvalidMaxPossibleCCVsPerMsg.selector, 0));
    s_executorOnRamp.setMaxCCVsPerMsg(0);
  }
}
