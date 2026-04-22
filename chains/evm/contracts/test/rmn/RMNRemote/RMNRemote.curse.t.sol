// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RMN} from "../../../rmn/RMN.sol";
import {RMNRemoteSetup} from "./RMNRemoteSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract RMNRemote_curse is RMNRemoteSetup {
  function test_curse() public {
    vm.expectEmit();
    emit RMN.Cursed(s_curseSubjects);

    s_rmn.curse(s_curseSubjects);

    assertEq(abi.encode(s_rmn.getCursedSubjects()), abi.encode(s_curseSubjects));
    assertTrue(s_rmn.isCursed(CURSE_SUBJ_1));
    assertTrue(s_rmn.isCursed(CURSE_SUBJ_2));
    // Should not have cursed a random subject
    assertFalse(s_rmn.isCursed(bytes16(keccak256("subject 3"))));
  }

  function test_curse_ByOwner_WhenCurseAdminsExist() public {
    address curseAdmin = makeAddr("curseAdmin");
    address[] memory adds = new address[](1);
    adds[0] = curseAdmin;
    s_rmn.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: adds, removedCallers: new address[](0)})
    );

    vm.expectEmit();
    emit RMN.Cursed(s_curseSubjects);
    s_rmn.curse(s_curseSubjects);

    assertTrue(s_rmn.isCursed(CURSE_SUBJ_1));
    assertTrue(s_rmn.isCursed(CURSE_SUBJ_2));
  }

  function test_curse_ByCurseAdmin() public {
    address curseAdmin = makeAddr("curseAdmin");
    address[] memory adds = new address[](1);
    adds[0] = curseAdmin;
    s_rmn.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: adds, removedCallers: new address[](0)})
    );

    vm.stopPrank();
    vm.prank(curseAdmin);
    vm.expectEmit();
    emit RMN.Cursed(s_curseSubjects);
    s_rmn.curse(s_curseSubjects);

    assertTrue(s_rmn.isCursed(CURSE_SUBJ_1));
    assertTrue(s_rmn.isCursed(CURSE_SUBJ_2));
  }

  function test_curse_ByNewlyAddedCurseAdmin() public {
    address newAdmin = makeAddr("newAdmin");
    address[] memory adds = new address[](1);
    adds[0] = newAdmin;
    s_rmn.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: adds, removedCallers: new address[](0)})
    );

    vm.stopPrank();
    vm.prank(newAdmin);
    vm.expectEmit();
    emit RMN.Cursed(s_curseSubjects);
    s_rmn.curse(s_curseSubjects);

    assertTrue(s_rmn.isCursed(CURSE_SUBJ_1));
    assertTrue(s_rmn.isCursed(CURSE_SUBJ_2));
  }

  function test_RevertWhen_curse_AlreadyCursed_duplicateSubject() public {
    s_curseSubjects.push(CURSE_SUBJ_1);

    vm.expectRevert(abi.encodeWithSelector(RMN.AlreadyCursed.selector, CURSE_SUBJ_1));
    s_rmn.curse(s_curseSubjects);
  }

  function test_RevertWhen_curse_calledByNonOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_rmn.curse(s_curseSubjects);
  }

  function test_RevertWhen_curse_calledByRemovedCurseAdmin() public {
    address curseAdmin = makeAddr("curseAdmin");
    address[] memory adds = new address[](1);
    adds[0] = curseAdmin;
    s_rmn.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: adds, removedCallers: new address[](0)})
    );
    address[] memory toRemove = new address[](1);
    toRemove[0] = curseAdmin;
    s_rmn.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: new address[](0), removedCallers: toRemove})
    );

    vm.stopPrank();
    vm.prank(curseAdmin);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, curseAdmin));
    s_rmn.curse(s_curseSubjects);
  }
}
