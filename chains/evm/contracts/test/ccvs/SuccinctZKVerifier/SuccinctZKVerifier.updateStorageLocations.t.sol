// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract SuccinctZKVerifier_updateStorageLocations is SuccinctZKVerifierSetup {
  function test_updateStorageLocations() public {
    string[] memory newStorageLocations = new string[](1);
    newStorageLocations[0] = "new/location";

    vm.expectEmit();
    emit BaseVerifier.StorageLocationsUpdated(s_storageLocations, newStorageLocations);

    s_zkVerifier.updateStorageLocations(newStorageLocations);

    assertEq(newStorageLocations, s_zkVerifier.getStorageLocations());
  }

  // Reverts

  function test_updateStorageLocations_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_zkVerifier.updateStorageLocations(new string[](0));
  }
}
