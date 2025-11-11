// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ContractFactory} from "../../ContractFactory.sol";
import {ContractFactorySetup} from "./ContractFactorySetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract ContractFactory_applyAllowListUpdates is ContractFactorySetup {
  function test_applyAllowListUpdates() public {
    address[] memory removes = new address[](1);
    removes[0] = s_allowedCaller;

    address newAllowedCaller = makeAddr("NEW_ALLOWED_CALLER");
    address[] memory adds = new address[](1);
    adds[0] = newAllowedCaller;

    vm.expectEmit();
    emit ContractFactory.CallerRemoved(s_allowedCaller);
    vm.expectEmit();
    emit ContractFactory.CallerAdded(newAllowedCaller);
    s_contractFactory.applyAllowListUpdates(removes, adds);

    address[] memory allowList = s_contractFactory.getAllowList();
    assertEq(allowList.length, 1);
    assertEq(allowList[0], newAllowedCaller);
  }

  function test_applyAllowListUpdates_SkipEmitWhen_NoOp() public {
    address[] memory adds = new address[](1);
    adds[0] = s_allowedCaller;

    address unknownCaller = makeAddr("UNKNOWN_CALLER");
    address[] memory removes = new address[](1);
    removes[0] = unknownCaller;

    vm.recordLogs();
    s_contractFactory.applyAllowListUpdates(removes, adds);
    vm.assertEq(vm.getRecordedLogs().length, 0);

    address[] memory allowList = s_contractFactory.getAllowList();
    assertEq(allowList.length, 1);
    assertEq(allowList[0], s_allowedCaller);
  }

  function test_applyAllowListUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_contractFactory.applyAllowListUpdates(new address[](0), new address[](0));
  }
}
