// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract AdvancedPoolHooks_applyAuthorizedCallerUpdates is AdvancedPoolHooksSetup {
  address internal s_newCaller1 = makeAddr("newCaller1");

  function test_applyAuthorizedCallerUpdates_AddAndRemove() public {
    address[] memory addedCallers = new address[](1);
    addedCallers[0] = s_newCaller1;

    address[] memory removedCallers = new address[](1);
    removedCallers[0] = OWNER;

    vm.expectEmit();
    emit AuthorizedCallers.AuthorizedCallerRemoved(OWNER);
    vm.expectEmit();
    emit AuthorizedCallers.AuthorizedCallerAdded(s_newCaller1);

    s_advancedPoolHooks.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: addedCallers, removedCallers: removedCallers})
    );

    address[] memory allCallers = s_advancedPoolHooks.getAllAuthorizedCallers();
    assertEq(allCallers.length, 2); // s_tokenPool + s_newCaller1
  }

  // Reverts

  function test_applyAuthorizedCallerUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_advancedPoolHooks.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: new address[](0), removedCallers: new address[](0)})
    );
  }
}
