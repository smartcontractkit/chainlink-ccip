// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";

contract TokenPool_setMinBlockConfirmation is AdvancedPoolHooksSetup {
  function test_setMinBlockConfirmation() public {
    uint16 newMinBlockConfirmations = 42;
    vm.expectEmit();
    emit TokenPool.MinBlockConfirmationSet(newMinBlockConfirmations);
    s_tokenPool.setMinBlockConfirmation(newMinBlockConfirmations);
    assertEq(s_tokenPool.getCustomMinBlockConfirmation(), newMinBlockConfirmations);
  }
}
