// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";

import {TokenPool} from "../../../pools/TokenPool.sol";
import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract TokenPool_updateAdvancedPoolHooks is AdvancedPoolHooksSetup {
  IAdvancedPoolHooks internal constant NEW_HOOK = IAdvancedPoolHooks(address(bytes20(keccak256("MESSAGE_RECEIVER"))));

  function test_updateAdvancedPoolHooks() public {
    IAdvancedPoolHooks oldHook = s_tokenPool.getAdvancedPoolHooks();
    vm.expectEmit();
    emit TokenPool.AdvancedPoolHooksUpdated(oldHook, NEW_HOOK);
    s_tokenPool.updateAdvancedPoolHooks(NEW_HOOK);
    assertEq(address(s_tokenPool.getAdvancedPoolHooks()), address(NEW_HOOK));
  }

  // Reverts
  function test_updateAdvancedPoolHooks_OnlyCallableByOwner() public {
    changePrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.updateAdvancedPoolHooks(NEW_HOOK);
  }
}
