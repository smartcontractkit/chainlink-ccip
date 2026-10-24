// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {MultiTokenPoolHooksSetup} from "./MultiTokenPoolHooksSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract MultiTokenPoolHooks_applyAuthorizedCallerUpdates is MultiTokenPoolHooksSetup {
  address internal s_newPool = makeAddr("newPool");

  function test_applyAuthorizedCallerUpdates() public {
    vm.expectEmit();
    emit AuthorizedCallers.AuthorizedCallerRemoved(s_pool);
    vm.expectEmit();
    emit AuthorizedCallers.AuthorizedCallerAdded(s_newPool);
    s_hooks.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: _arrayOf(s_newPool), removedCallers: _arrayOf(s_pool)})
    );

    assertEq(_arrayOf(s_newPool), s_hooks.getAllAuthorizedCallers());

    vm.startPrank(s_newPool);
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);

    vm.startPrank(s_pool);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, s_pool));
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  // Reverts

  function test_applyAuthorizedCallerUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_hooks.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: _arrayOf(s_newPool), removedCallers: new address[](0)})
    );
  }
}
