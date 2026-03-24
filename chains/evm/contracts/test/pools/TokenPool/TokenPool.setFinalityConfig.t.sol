// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";

import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract TokenPool_setFinalityConfig is AdvancedPoolHooksSetup {
  function test_setFinalityConfig() public {
    bytes2 newMinFinality = bytes2(uint16(42));
    vm.expectEmit();
    emit TokenPool.FinalityConfigSet(newMinFinality);
    s_tokenPool.setFinalityConfig(newMinFinality);
    assertEq(s_tokenPool.getFinalityConfig(), newMinFinality);
  }

  // Reverts
  function test_setFinalityConfig_RevertWhen_OnlyCallableByOwner() public {
    bytes2 newMinFinality = bytes2(uint16(42));
    changePrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.setFinalityConfig(newMinFinality);
  }
}
