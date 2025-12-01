// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract AdvancedPoolHooks_setThresholdAmount is AdvancedPoolHooksSetup {
  function test_setThresholdAmount() public {
    uint256 newThreshold = 5000e18;

    vm.expectEmit();
    emit AdvancedPoolHooks.ThresholdAmountSet(newThreshold);

    s_advancedPoolHooks.setThresholdAmount(newThreshold);

    assertEq(s_advancedPoolHooks.getThresholdAmount(), newThreshold);
  }

  function test_setThresholdAmount_ToZero() public {
    uint256 newThreshold = 0;

    vm.expectEmit();
    emit AdvancedPoolHooks.ThresholdAmountSet(newThreshold);

    s_advancedPoolHooks.setThresholdAmount(newThreshold);

    assertEq(s_advancedPoolHooks.getThresholdAmount(), newThreshold);
  }

  // Reverts

  function test_setThresholdAmount_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_advancedPoolHooks.setThresholdAmount(0);
  }
}
