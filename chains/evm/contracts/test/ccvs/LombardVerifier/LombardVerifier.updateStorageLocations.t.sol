// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {LombardVerifierSetup} from "./LombardVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract LombardVerifier_updateStorageLocations is LombardVerifierSetup {
  function test_updateStorageLocations() public {
    string[] memory newStorageLocations = new string[](1);
    newStorageLocations[0] = "new/location";

    vm.expectEmit();
    emit BaseVerifier.StorageLocationsUpdated(s_storageLocations, newStorageLocations);

    s_lombardVerifier.updateStorageLocations(newStorageLocations);
  }

  function test_updateStorageLocation_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_lombardVerifier.updateStorageLocations(new string[](0));
  }
}
