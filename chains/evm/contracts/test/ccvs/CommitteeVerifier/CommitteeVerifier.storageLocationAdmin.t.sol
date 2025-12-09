// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitteeVerifier} from "../../../ccvs/CommitteeVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {CommitteeVerifierSetup} from "./CommitteeVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CommitteeVerifier_storageLocationAdmin is CommitteeVerifierSetup {
  function test_transferStorageLocationAdmin_SetsPendingAndEmits() public {
    address newAdmin = makeAddr("NewStorageLocationAdmin");

    vm.expectEmit();
    emit CommitteeVerifier.StorageLocationAdminTransferRequested(OWNER, newAdmin);

    s_committeeVerifier.transferStorageLocationAdmin(newAdmin);

    assertEq(s_committeeVerifier.getPendingStorageLocationAdmin(), newAdmin);
    assertEq(s_committeeVerifier.getStorageLocationAdmin(), OWNER);
  }

  function test_transferStorageLocationAdmin_RevertWhen_NotStorageLocationAdmin() public {
    address newAdmin = makeAddr("UnauthorizedStorageLocationAdmin");

    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(CommitteeVerifier.OnlyCallableByStorageLocationAdmin.selector);
    s_committeeVerifier.transferStorageLocationAdmin(newAdmin);
  }

  function test_transferStorageLocationAdmin_RevertWhen_TransferToSelf() public {
    vm.expectRevert(Ownable2Step.CannotTransferToSelf.selector);
    s_committeeVerifier.transferStorageLocationAdmin(OWNER);
  }

  function test_acceptStorageLocationAdmin_SetsNewAdminAndEmits() public {
    address newAdmin = makeAddr("AcceptedStorageLocationAdmin");
    s_committeeVerifier.transferStorageLocationAdmin(newAdmin);

    vm.stopPrank();
    vm.prank(newAdmin);
    vm.expectEmit();
    emit CommitteeVerifier.StorageLocationAdminTransferred(OWNER, newAdmin);

    s_committeeVerifier.acceptStorageLocationAdmin();

    assertEq(s_committeeVerifier.getStorageLocationAdmin(), newAdmin);
    assertEq(s_committeeVerifier.getPendingStorageLocationAdmin(), address(0));
  }

  function test_acceptStorageLocationAdmin_RevertWhen_NotPendingAdmin() public {
    vm.stopPrank();
    vm.expectRevert(CommitteeVerifier.MustBeProposedStorageLocationAdmin.selector);
    s_committeeVerifier.acceptStorageLocationAdmin();
  }

  function test_updateStorageLocation_AfterAdminTransfer() public {
    address newAdmin = makeAddr("PostTransferStorageLocationAdmin");
    string memory newStorageLocation = "storage/location/admin";

    s_committeeVerifier.transferStorageLocationAdmin(newAdmin);

    vm.stopPrank();
    vm.startPrank(newAdmin);
    s_committeeVerifier.acceptStorageLocationAdmin();
    vm.expectEmit();
    emit BaseVerifier.StorageLocationUpdated(STORAGE_LOCATION, newStorageLocation);

    s_committeeVerifier.updateStorageLocation(newStorageLocation);
    vm.stopPrank();
  }
}
