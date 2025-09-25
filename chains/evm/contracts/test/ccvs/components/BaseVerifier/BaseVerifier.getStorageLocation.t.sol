// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

contract BaseVerifier_getStorageLocation is BaseVerifierSetup {
  function test_getStorageLocation() public view {
    string memory storageLocation = s_baseVerifier.getStorageLocation();
    assertEq(storageLocation, STORAGE_LOCATION);
  }
}
