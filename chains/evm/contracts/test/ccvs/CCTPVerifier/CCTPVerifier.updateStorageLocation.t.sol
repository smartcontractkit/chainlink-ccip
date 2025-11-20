// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {CCTPVerifierSetup} from "./CCTPVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CCTPVerifier_updateStorageLocation is CCTPVerifierSetup {
  function test_updateStorageLocation() public {
    string memory newStorageLocation = "new/location";

    vm.expectEmit();
    emit BaseVerifier.StorageLocationUpdated(STORAGE_LOCATION, newStorageLocation);

    s_cctpVerifier.updateStorageLocation(newStorageLocation);
  }

  // Reverts

  function test_updateStorageLocation_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_cctpVerifier.updateStorageLocation("new/location");
  }
}
