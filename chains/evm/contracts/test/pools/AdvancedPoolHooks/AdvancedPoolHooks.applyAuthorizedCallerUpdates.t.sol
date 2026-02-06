// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract AdvancedPoolHooks_applyAuthorizedCallerUpdates is AdvancedPoolHooksSetup {
  address internal s_authorizedCaller = makeAddr("authorizedCaller");
  address internal s_newCaller1 = makeAddr("newCaller1");

  AdvancedPoolHooks internal s_hooksWithAuthorizedCallers;

  function setUp() public virtual override {
    super.setUp();

    address[] memory authorizedCallers = new address[](1);
    authorizedCallers[0] = s_authorizedCaller;
    s_hooksWithAuthorizedCallers = new AdvancedPoolHooks(new address[](0), 0, address(0), authorizedCallers);
  }

  function test_applyAuthorizedCallerUpdates_AddAndRemove() public {
    assertTrue(s_hooksWithAuthorizedCallers.getAuthorizedCallersEnabled());

    address[] memory addedCallers = new address[](1);
    addedCallers[0] = s_newCaller1;

    address[] memory removedCallers = new address[](1);
    removedCallers[0] = s_authorizedCaller;

    vm.expectEmit();
    emit AuthorizedCallers.AuthorizedCallerRemoved(s_authorizedCaller);
    vm.expectEmit();
    emit AuthorizedCallers.AuthorizedCallerAdded(s_newCaller1);

    s_hooksWithAuthorizedCallers.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: addedCallers, removedCallers: removedCallers})
    );

    address[] memory allCallers = s_hooksWithAuthorizedCallers.getAllAuthorizedCallers();
    assertEq(allCallers.length, 1);
  }

  // Reverts

  function test_applyAuthorizedCallerUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_hooksWithAuthorizedCallers.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: new address[](0), removedCallers: new address[](0)})
    );
  }

  function test_applyAuthorizedCallerUpdates_RevertWhen_AuthorizedCallersNotEnabled() public {
    assertFalse(s_advancedPoolHooks.getAuthorizedCallersEnabled());

    vm.expectRevert(AdvancedPoolHooks.AuthorizedCallersNotEnabled.selector);
    s_advancedPoolHooks.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: new address[](0), removedCallers: new address[](0)})
    );
  }
}
