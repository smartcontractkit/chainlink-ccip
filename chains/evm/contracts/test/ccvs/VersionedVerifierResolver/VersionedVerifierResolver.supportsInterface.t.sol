// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VersionedVerifierResolver} from "../../../ccvs/VersionedVerifierResolver.sol";
import {ICrossChainVerifierResolver} from "../../../interfaces/ICrossChainVerifierResolver.sol";
import {VersionedVerifierResolverSetup} from "./VersionedVerifierResolverSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";

contract VersionedVerifierResolver_supportsInterface is VersionedVerifierResolverSetup {
  function test_supportsInterface_ICrossChainVerifierResolver() public view {
    assertTrue(s_versionedVerifierResolver.supportsInterface(type(ICrossChainVerifierResolver).interfaceId));
  }

  function test_supportsInterface_IERC165() public view {
    assertTrue(s_versionedVerifierResolver.supportsInterface(type(IERC165).interfaceId));
  }

  function test_supportsInterface_Unsupported() public view {
    assertFalse(s_versionedVerifierResolver.supportsInterface(bytes4(0x12345678)));
  }
}
