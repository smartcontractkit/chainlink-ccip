// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {CommitteeVerifierSetup} from "./CommitteeVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CommitteeVerifier_updateStorageLocation is CommitteeVerifierSetup {
  function test_updateStorageLocation() public {
    string memory newStorageLocation = "new/location";

    vm.expectEmit();
    emit BaseVerifier.StorageLocationUpdated(STORAGE_LOCATION, newStorageLocation);

    s_committeeVerifier.updateStorageLocation(newStorageLocation);
  }

  // Reverts

  function test_updateStorageLocation_RevertWhen_NotOwner() public {
    vm.stopPrank();

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_committeeVerifier.updateStorageLocation("new/location");
  }
}
