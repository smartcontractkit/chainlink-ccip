// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {CCTPVerifierSetup} from "./CCTPVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CCTPVerifier_updateStorageLocations is CCTPVerifierSetup {
  function test_updateStorageLocations() public {
    string[] memory newStorageLocations = new string[](1);
    newStorageLocations[0] = "new/location";

    vm.expectEmit();
    emit BaseVerifier.StorageLocationsUpdated(storageLocations, newStorageLocations);

    s_cctpVerifier.updateStorageLocations(newStorageLocations);
  }

  function test_updateStorageLocation_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_cctpVerifier.updateStorageLocations(new string[](0));
  }
}
