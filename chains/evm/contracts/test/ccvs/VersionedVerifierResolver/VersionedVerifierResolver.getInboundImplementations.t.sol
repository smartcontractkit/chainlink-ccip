// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VersionedVerifierResolver} from "../../../ccvs/VersionedVerifierResolver.sol";
import {VersionedVerifierResolverSetup} from "./VersionedVerifierResolverSetup.t.sol";

contract VersionedVerifierResolver_getInboundImplementation is VersionedVerifierResolverSetup {
  function test_getInboundImplementation_RevertWhen_InvalidCCVDataLength() public {
    vm.expectRevert(VersionedVerifierResolver.InvalidCCVDataLength.selector);
    s_versionedVerifierResolver.getInboundImplementation("");
  }
}
