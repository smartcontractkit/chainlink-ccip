// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../../../../interfaces/ICrossChainVerifierV1.sol";

import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

contract BaseVerifier_getStorageLocation is BaseVerifierSetup {
  function test_supportsInterface() public view {
    assertTrue(s_baseVerifier.supportsInterface(type(IERC165).interfaceId));
    assertTrue(s_baseVerifier.supportsInterface(type(ICrossChainVerifierV1).interfaceId));
  }
}
