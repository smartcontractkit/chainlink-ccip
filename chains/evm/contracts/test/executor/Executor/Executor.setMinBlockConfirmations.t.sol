// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Executor} from "../../../executor/Executor.sol";
import {ExecutorSetup} from "./ExecutorSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract Executor_setMinBlockConfirmations is ExecutorSetup {
  function test_setMinBlockConfirmations() public {
    uint16 newMinBlockConfirmations = 12;

    vm.expectEmit();
    emit Executor.MinBlockConfirmationsSet(newMinBlockConfirmations);

    s_executor.setMinBlockConfirmations(newMinBlockConfirmations);

    assertEq(s_executor.getMinBlockConfirmations(), newMinBlockConfirmations);
  }

  function test_constructor_RevertWhen_OnlyOwner() public {
    vm.stopPrank();

    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector, 0));

    new Executor(0, 0);
  }
}
