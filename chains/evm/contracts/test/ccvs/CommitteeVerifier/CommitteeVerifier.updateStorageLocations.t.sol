// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitteeVerifier} from "../../../ccvs/CommitteeVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {CommitteeVerifierSetup} from "./CommitteeVerifierSetup.t.sol";

contract CommitteeVerifier_updateStorageLocations is CommitteeVerifierSetup {
  function test_updateStorageLocations() public {
    string[] memory newStorageLocations = new string[](1);
    newStorageLocations[0] = "new/location";

    vm.expectEmit();
    emit BaseVerifier.StorageLocationsUpdated(storageLocations, newStorageLocations);

    s_committeeVerifier.updateStorageLocations(newStorageLocations);
  }

  // Reverts

  function test_updateStorageLocation_RevertWhen_NotStorageLocationsAdmin() public {
    vm.stopPrank();
    string[] memory newStorageLocations = new string[](1);
    newStorageLocations[0] = "new/location";

    vm.prank(STRANGER);
    vm.expectRevert(CommitteeVerifier.OnlyCallableByStorageLocationsAdmin.selector);
    s_committeeVerifier.updateStorageLocations(newStorageLocations);
  }
}
