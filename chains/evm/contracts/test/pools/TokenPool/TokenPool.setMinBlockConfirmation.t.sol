// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";

import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract TokenPool_setMinBlockConfirmations is AdvancedPoolHooksSetup {
  function test_setMinBlockConfirmations() public {
    uint16 newMinBlockConfirmations = 42;
    vm.expectEmit();
    emit TokenPool.MinBlockConfirmationsSet(newMinBlockConfirmations);
    s_tokenPool.setMinBlockConfirmations(newMinBlockConfirmations);
    assertEq(s_tokenPool.getCustomMinBlockConfirmations(), newMinBlockConfirmations);
  }

  // Reverts
  function test_setMinBlockConfirmations_RevertWhen_OnlyCallableByOwner() public {
    uint16 newMinBlockConfirmations = 42;
    changePrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.setMinBlockConfirmations(newMinBlockConfirmations);
  }
}
