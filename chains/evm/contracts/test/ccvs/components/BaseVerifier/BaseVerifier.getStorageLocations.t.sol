// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

contract BaseVerifier_getStorageLocations is BaseVerifierSetup {
  function test_getStorageLocations() public view {
    string[] memory storageLocations = s_baseVerifier.getStorageLocations();
    assertEq(storageLocations.length, 1);
    assertEq(storageLocations[0], storageLocations[0]);
  }
}
