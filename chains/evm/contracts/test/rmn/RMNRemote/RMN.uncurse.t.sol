// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RMN} from "../../../rmn/RMN.sol";
import {RMNRemoteSetup} from "./RMNSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract RMNRemote_uncurse is RMNRemoteSetup {
  function setUp() public override {
    super.setUp();
    s_rmn.curse(s_curseSubjects);
  }

  function test_uncurse() public {
    vm.expectEmit();
    emit RMN.Uncursed(s_curseSubjects);

    s_rmn.uncurse(s_curseSubjects);

    assertEq(s_rmn.getCursedSubjects().length, 0);
    assertFalse(s_rmn.isCursed(CURSE_SUBJ_1));
    assertFalse(s_rmn.isCursed(CURSE_SUBJ_2));
  }

  function test_RevertWhen_uncurse_NotCursed_duplicatedUncurseSubject() public {
    s_curseSubjects.push(CURSE_SUBJ_1);

    vm.expectRevert(abi.encodeWithSelector(RMN.NotCursed.selector, CURSE_SUBJ_1));
    s_rmn.uncurse(s_curseSubjects);
  }

  function test_RevertWhen_uncurse_calledByNonOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_rmn.uncurse(s_curseSubjects);
  }

  function test_RevertWhen_uncurse_calledByCurseAdmin() public {
    address curseAdmin = makeAddr("curseAdmin");
    address[] memory adds = new address[](1);
    adds[0] = curseAdmin;
    s_rmn.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: adds, removedCallers: new address[](0)})
    );

    vm.stopPrank();
    vm.prank(curseAdmin);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_rmn.uncurse(s_curseSubjects);
  }
}
