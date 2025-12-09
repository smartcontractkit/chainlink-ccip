// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitteeVerifier} from "../../../ccvs/CommitteeVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {CommitteeVerifierSetup} from "./CommitteeVerifierSetup.t.sol";

contract CommitteeVerifier_updateStorageLocation is CommitteeVerifierSetup {
  function test_updateStorageLocation() public {
    string memory newStorageLocation = "new/location";

    vm.expectEmit();
    emit BaseVerifier.StorageLocationUpdated(STORAGE_LOCATION, newStorageLocation);

    s_committeeVerifier.updateStorageLocation(newStorageLocation);
  }

  // Reverts

  function test_updateStorageLocation_RevertWhen_NotStorageLocationAdmin() public {
    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(CommitteeVerifier.OnlyCallableByStorageLocationAdmin.selector);
    s_committeeVerifier.updateStorageLocation("new/location");
  }
}
