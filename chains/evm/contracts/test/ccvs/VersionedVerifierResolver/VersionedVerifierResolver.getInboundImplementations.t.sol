// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VersionedVerifierResolverSetup} from "./VersionedVerifierResolverSetup.t.sol";

contract VersionedVerifierResolver_getInboundImplementation is VersionedVerifierResolverSetup {
  function test_getInboundImplementation_InvalidCCVDataLength() public view {
    assertEq(s_versionedVerifierResolver.getInboundImplementation(""), address(0));
  }
}
