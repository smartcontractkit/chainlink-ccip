// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {OffRampSetup} from "./OffRampSetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract OffRamp_setMaxGasBufferToUpdateState is OffRampSetup {
  function test_setMaxGasBufferToUpdateState() public {
    uint32 oldMaxGasBufferToUpdateState = s_offRamp.getmaxGasBufferToUpdateState();
    uint32 newMaxGasBufferToUpdateState = 15000;

    vm.expectEmit();
    emit OffRamp.MaxGasBufferToUpdateStateUpdated(oldMaxGasBufferToUpdateState, newMaxGasBufferToUpdateState);

    s_offRamp.setMaxGasBufferToUpdateState(newMaxGasBufferToUpdateState);

    assertEq(s_offRamp.getmaxGasBufferToUpdateState(), newMaxGasBufferToUpdateState);
  }

  function test_setMaxGasBufferToUpdateState_RevertWhen_GasCannotBeZero() public {
    vm.expectRevert(OffRamp.GasCannotBeZero.selector);
    s_offRamp.setMaxGasBufferToUpdateState(0);
  }

  function test_setMaxGasBufferToUpdateState_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_offRamp.setMaxGasBufferToUpdateState(15000);
  }
}

