// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";

import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract TokenPool_setMinBlockConfirmation is AdvancedPoolHooksSetup {
  function test_setMinBlockConfirmation() public {
    uint16 newMinBlockConfirmations = 42;
    vm.expectEmit();
    emit TokenPool.MinBlockConfirmationSet(newMinBlockConfirmations);
    s_tokenPool.setMinBlockConfirmation(newMinBlockConfirmations);
    assertEq(s_tokenPool.getCustomMinBlockConfirmation(), newMinBlockConfirmations);
  }

  // Reverts
  function test_setMinBlockConfirmation_RevertWhen_OnlyCallableByOwner() public {
    uint16 newMinBlockConfirmations = 42;
    changePrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.setMinBlockConfirmation(newMinBlockConfirmations);
  }
}
