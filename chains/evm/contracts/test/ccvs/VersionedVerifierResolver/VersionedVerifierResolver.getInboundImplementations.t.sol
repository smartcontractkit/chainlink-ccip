// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VersionedVerifierResolver} from "../../../ccvs/VersionedVerifierResolver.sol";
import {VersionedVerifierResolverSetup} from "./VersionedVerifierResolverSetup.t.sol";

contract VersionedVerifierResolver_getInboundImplementation is VersionedVerifierResolverSetup {
  function test_getInboundImplementation_RevertWhen_InvalidVerifierResultsLength() public {
    vm.expectRevert(VersionedVerifierResolver.InvalidVerifierResultsLength.selector);
    s_versionedVerifierResolver.getInboundImplementation("");
  }
}
