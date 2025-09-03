// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ExecutorOnRamp} from "../../../onRamp/ExecutorOnRamp.sol";
import {ExecutorOnRampSetup} from "./ExecutorOnRampSetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract ExecutorOnRamp_applyAllowedCCVUpdates is ExecutorOnRampSetup {
  function test_applyAllowedCCVUpdates_AddNewCCV() public {
    address[] memory newCCVs = new address[](1);
    address newCCV = makeAddr("newCCV");
    newCCVs[0] = newCCV;

    vm.expectEmit();
    emit ExecutorOnRamp.CCVAdded(newCCV);
    s_executorOnRamp.applyAllowedCCVUpdates(new address[](0), newCCVs, true);

    address[] memory currentCCVs = s_executorOnRamp.getAllowedCCVs();
    assertEq(currentCCVs.length, 2);
    bool found = false;
    for (uint256 i = 0; i < currentCCVs.length; ++i) {
      if (currentCCVs[i] == newCCV) {
        found = true;
        break;
      }
    }
    assertTrue(found, "New ccv should be supported");
    assertTrue(s_executorOnRamp.isCCVAllowlistEnabled(), "CCV allowlist should be enabled");
  }

  function test_applyAllowedCCVUpdates_AddExistingChain() public {
    address[] memory newCCVs = new address[](1);
    newCCVs[0] = INITIAL_CCV;

    vm.recordLogs();
    s_executorOnRamp.applyAllowedCCVUpdates(new address[](0), newCCVs, true);
    vm.assertEq(vm.getRecordedLogs().length, 0);

    address[] memory currentCCVs = s_executorOnRamp.getAllowedCCVs();
    assertEq(currentCCVs.length, 1);
  }

  function test_applyAllowedCCVUpdates_RemoveExistingCCV() public {
    address[] memory ccvsToRemove = new address[](1);
    ccvsToRemove[0] = INITIAL_CCV;

    vm.expectEmit();
    emit ExecutorOnRamp.CCVRemoved(INITIAL_CCV);
    s_executorOnRamp.applyAllowedCCVUpdates(ccvsToRemove, new address[](0), true);

    address[] memory currentCCVs = s_executorOnRamp.getAllowedCCVs();
    assertEq(currentCCVs.length, 0);
  }

  function test_applyAllowedCCVUpdates_RemoveNonexistentCCV() public {
    address[] memory ccvsToRemove = new address[](1);
    ccvsToRemove[0] = makeAddr("nonexistentCCV");

    vm.recordLogs();
    s_executorOnRamp.applyAllowedCCVUpdates(ccvsToRemove, new address[](0), true);
    vm.assertEq(vm.getRecordedLogs().length, 0);

    address[] memory currentCCVs = s_executorOnRamp.getAllowedCCVs();
    assertEq(currentCCVs.length, 1);
  }

  function test_applyAllowedCCVUpdates_DisableAllowlist() public {
    vm.expectEmit();
    emit ExecutorOnRamp.CCVAllowlistUpdated(false);
    s_executorOnRamp.applyAllowedCCVUpdates(new address[](0), new address[](0), false);

    assertFalse(s_executorOnRamp.isCCVAllowlistEnabled(), "CCV allowlist should be disabled");
  }

  function test_applyAllowedCCVUpdates_RevertWhen_NotOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_executorOnRamp.applyAllowedCCVUpdates(new address[](0), new address[](0), true);
  }

  function test_applyAllowedCCVUpdates_RevertWhen_CCVInvalid() public {
    address[] memory newCCVs = new address[](1);
    newCCVs[0] = address(0);

    vm.expectRevert(abi.encodeWithSelector(ExecutorOnRamp.InvalidCCV.selector, address(0)));
    s_executorOnRamp.applyAllowedCCVUpdates(new address[](0), newCCVs, true);
  }
}
