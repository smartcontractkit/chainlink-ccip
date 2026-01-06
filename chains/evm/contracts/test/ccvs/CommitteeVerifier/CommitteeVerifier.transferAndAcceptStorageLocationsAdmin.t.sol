// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitteeVerifier} from "../../../ccvs/CommitteeVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {CommitteeVerifierSetup} from "./CommitteeVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CommitteeVerifier_storageLocationsAdmin is CommitteeVerifierSetup {
  function test_transferStorageLocationsAdmin_SetsPendingAndEmits() public {
    address newAdmin = makeAddr("NewStorageLocationsAdmin");

    vm.expectEmit();
    emit CommitteeVerifier.StorageLocationsAdminTransferRequested(OWNER, newAdmin);

    s_committeeVerifier.transferStorageLocationsAdmin(newAdmin);

    assertEq(s_committeeVerifier.getPendingStorageLocationsAdmin(), newAdmin);
    assertEq(s_committeeVerifier.getStorageLocationsAdmin(), OWNER);
  }

  function test_transferStorageLocationsAdmin_RevertWhen_NotStorageLocationsAdmin() public {
    address newAdmin = makeAddr("UnauthorizedStorageLocationsAdmin");

    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(CommitteeVerifier.OnlyCallableByStorageLocationsAdmin.selector);
    s_committeeVerifier.transferStorageLocationsAdmin(newAdmin);
  }

  function test_transferStorageLocationsAdmin_RevertWhen_TransferToSelf() public {
    vm.expectRevert(Ownable2Step.CannotTransferToSelf.selector);
    s_committeeVerifier.transferStorageLocationsAdmin(OWNER);
  }

  function test_acceptStorageLocationsAdmin_SetsNewAdminAndEmits() public {
    address newAdmin = makeAddr("AcceptedStorageLocationsAdmin");
    s_committeeVerifier.transferStorageLocationsAdmin(newAdmin);

    vm.stopPrank();
    vm.prank(newAdmin);
    vm.expectEmit();
    emit CommitteeVerifier.StorageLocationsAdminTransferred(OWNER, newAdmin);

    s_committeeVerifier.acceptStorageLocationsAdmin();

    assertEq(s_committeeVerifier.getStorageLocationsAdmin(), newAdmin);
    assertEq(s_committeeVerifier.getPendingStorageLocationsAdmin(), address(0));
  }

  function test_acceptStorageLocationsAdmin_RevertWhen_NotPendingAdmin() public {
    vm.stopPrank();
    vm.expectRevert(CommitteeVerifier.MustBeProposedStorageLocationsAdmin.selector);
    s_committeeVerifier.acceptStorageLocationsAdmin();
  }

  function test_updateStorageLocation_AfterAdminTransfer() public {
    address newAdmin = makeAddr("PostTransferStorageLocationsAdmin");
    string[] memory newStorageLocations = new string[](1);
    newStorageLocations[0] = "storage/location/admin";

    s_committeeVerifier.transferStorageLocationsAdmin(newAdmin);

    vm.stopPrank();
    vm.startPrank(newAdmin);
    s_committeeVerifier.acceptStorageLocationsAdmin();
    vm.expectEmit();
    emit BaseVerifier.StorageLocationsUpdated(s_storageLocations, newStorageLocations);

    s_committeeVerifier.updateStorageLocations(newStorageLocations);
  }
}
