// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RMN} from "../../../rmn/RMN.sol";
import {RMNRemoteSetup} from "./RMNRemoteSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract RMNRemote_curse is RMNRemoteSetup {
  function test_curse() public {
    vm.expectEmit();
    emit RMN.Cursed(s_curseSubjects);

    s_rmnRemote.curse(s_curseSubjects);

    assertEq(abi.encode(s_rmnRemote.getCursedSubjects()), abi.encode(s_curseSubjects));
    assertTrue(s_rmnRemote.isCursed(CURSE_SUBJ_1));
    assertTrue(s_rmnRemote.isCursed(CURSE_SUBJ_2));
    // Should not have cursed a random subject
    assertFalse(s_rmnRemote.isCursed(bytes16(keccak256("subject 3"))));
  }

  function test_RevertWhen_curse_AlreadyCursed_duplicateSubject() public {
    s_curseSubjects.push(CURSE_SUBJ_1);

    vm.expectRevert(abi.encodeWithSelector(RMN.AlreadyCursed.selector, CURSE_SUBJ_1));
    s_rmnRemote.curse(s_curseSubjects);
  }

  function test_RevertWhen_curse_calledByNonOwner() public {
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    vm.stopPrank();
    vm.prank(STRANGER);
    s_rmnRemote.curse(s_curseSubjects);
  }
}

contract RMNRemote_curseAdmin is RMNRemoteSetup {
  address internal s_curseAdmin = makeAddr("curseAdmin");

  function setUp() public override {
    super.setUp();
    address[] memory adds = new address[](1);
    adds[0] = s_curseAdmin;
    s_rmnRemote.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: adds, removedCallers: new address[](0)})
    );
  }

  function test_curse_byCurseAdmin_Success() public {
    vm.expectEmit();
    emit RMN.Cursed(s_curseSubjects);

    vm.stopPrank();
    vm.prank(s_curseAdmin);
    s_rmnRemote.curse(s_curseSubjects);

    assertTrue(s_rmnRemote.isCursed(CURSE_SUBJ_1));
    assertTrue(s_rmnRemote.isCursed(CURSE_SUBJ_2));
  }

  function test_applyAuthorizedCallerUpdates_addsAndRemoves() public {
    address newAdmin = makeAddr("newAdmin");
    address[] memory adds = new address[](1);
    adds[0] = newAdmin;

    vm.expectEmit();
    emit AuthorizedCallers.AuthorizedCallerAdded(newAdmin);
    s_rmnRemote.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: adds, removedCallers: new address[](0)})
    );

    address[] memory adminList = s_rmnRemote.getAllAuthorizedCallers();
    assertEq(adminList.length, 2); // s_curseAdmin (from setUp) + newAdmin

    // Remove
    address[] memory toRemove = new address[](1);
    toRemove[0] = newAdmin;
    vm.expectEmit();
    emit AuthorizedCallers.AuthorizedCallerRemoved(newAdmin);
    s_rmnRemote.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: new address[](0), removedCallers: toRemove})
    );

    adminList = s_rmnRemote.getAllAuthorizedCallers();
    assertEq(adminList.length, 1); // only s_curseAdmin remains

    // Remove of non-member is idempotent (no event emitted)
    vm.recordLogs();
    s_rmnRemote.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: new address[](0), removedCallers: toRemove})
    );
    assertEq(vm.getRecordedLogs().length, 0);
  }

  function test_RevertWhen_applyAuthorizedCallerUpdates_calledByNonOwner() public {
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    vm.stopPrank();
    vm.prank(STRANGER);
    s_rmnRemote.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: new address[](0), removedCallers: new address[](0)})
    );
  }

  function test_curse_byOwner_SuccessWhenCurseAdminsExist() public {
    vm.expectEmit();
    emit RMN.Cursed(s_curseSubjects);
    s_rmnRemote.curse(s_curseSubjects);

    assertTrue(s_rmnRemote.isCursed(CURSE_SUBJ_1));
    assertTrue(s_rmnRemote.isCursed(CURSE_SUBJ_2));
  }

  function test_curse_byNewlyAddedCurseAdmin_Success() public {
    address newAdmin = makeAddr("newAdmin");
    address[] memory adds = new address[](1);
    adds[0] = newAdmin;
    s_rmnRemote.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: adds, removedCallers: new address[](0)})
    );

    vm.expectEmit();
    emit RMN.Cursed(s_curseSubjects);

    vm.stopPrank();
    vm.prank(newAdmin);
    s_rmnRemote.curse(s_curseSubjects);

    assertTrue(s_rmnRemote.isCursed(CURSE_SUBJ_1));
    assertTrue(s_rmnRemote.isCursed(CURSE_SUBJ_2));
  }

  function test_RevertWhen_curse_calledByRemovedCurseAdmin() public {
    address[] memory toRemove = new address[](1);
    toRemove[0] = s_curseAdmin;
    s_rmnRemote.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: new address[](0), removedCallers: toRemove})
    );

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, s_curseAdmin));
    vm.stopPrank();
    vm.prank(s_curseAdmin);
    s_rmnRemote.curse(s_curseSubjects);
  }

  function test_RevertWhen_uncurse_calledByCurseAdmin() public {
    s_rmnRemote.curse(s_curseSubjects);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    vm.stopPrank();
    vm.prank(s_curseAdmin);
    s_rmnRemote.uncurse(s_curseSubjects);
  }
}
