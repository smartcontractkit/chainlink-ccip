// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ExecutorOnRamp} from "../../../onRamp/ExecutorOnRamp.sol";
import {ExecutorOnRampSetup} from "./ExecutorOnRampSetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract ExecutorOnRamp_applyDestChainUpdates is ExecutorOnRampSetup {
  function test_applyDestChainUpdates_AddNewChain() public {
    uint64[] memory newDests = new uint64[](1);
    newDests[0] = INITIAL_DEST + 1;

    vm.expectEmit();
    emit ExecutorOnRamp.DestChainAdded(INITIAL_DEST + 1);
    s_executorOnRamp.applyDestChainUpdates(new uint64[](0), newDests);

    uint64[] memory currentDestChains = s_executorOnRamp.getDestChains();
    assertEq(currentDestChains.length, 2);
    bool found = false;
    for (uint256 i = 0; i < currentDestChains.length; ++i) {
      if (currentDestChains[i] == 2) {
        found = true;
        break;
      }
    }
    assertTrue(found, "New dest chain should be supported");
  }

  function test_applyDestChainUpdates_AddExistingChain() public {
    uint64[] memory newDests = new uint64[](1);
    newDests[0] = INITIAL_DEST;

    vm.recordLogs();
    s_executorOnRamp.applyDestChainUpdates(new uint64[](0), newDests);
    vm.assertEq(vm.getRecordedLogs().length, 0);

    uint64[] memory currentDestChains = s_executorOnRamp.getDestChains();
    assertEq(currentDestChains.length, 1);
  }

  function test_applyDestChainUpdates_RemoveExistingChain() public {
    uint64[] memory destsToRemove = new uint64[](1);
    destsToRemove[0] = INITIAL_DEST;

    vm.expectEmit();
    emit ExecutorOnRamp.DestChainRemoved(INITIAL_DEST);
    s_executorOnRamp.applyDestChainUpdates(destsToRemove, new uint64[](0));

    uint64[] memory currentDestChains = s_executorOnRamp.getDestChains();
    assertEq(currentDestChains.length, 0);
  }

  function test_applyDestChainUpdates_RemoveNonexistentChain() public {
    uint64[] memory destsToRemove = new uint64[](1);
    destsToRemove[0] = 999;

    vm.recordLogs();
    s_executorOnRamp.applyDestChainUpdates(destsToRemove, new uint64[](0));
    vm.assertEq(vm.getRecordedLogs().length, 0);

    uint64[] memory currentDestChains = s_executorOnRamp.getDestChains();
    assertEq(currentDestChains.length, 1);
  }

  function test_applyDestChainUpdates_RevertWhen_NotOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_executorOnRamp.applyDestChainUpdates(new uint64[](0), new uint64[](0));
  }

  function test_applyDestChainUpdates_RevertWhen_DestChainInvalid() public {
    uint64[] memory newDests = new uint64[](1);
    newDests[0] = 0;

    vm.expectRevert(abi.encodeWithSelector(ExecutorOnRamp.InvalidDestChain.selector, 0));
    s_executorOnRamp.applyDestChainUpdates(new uint64[](0), newDests);
  }
}
