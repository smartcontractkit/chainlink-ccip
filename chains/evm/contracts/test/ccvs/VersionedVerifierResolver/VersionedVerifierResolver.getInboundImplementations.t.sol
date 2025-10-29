// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VersionedVerifierResolver} from "../../../ccvs/VersionedVerifierResolver.sol";
import {VersionedVerifierResolverSetup} from "./VersionedVerifierResolverSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract VersionedVerifierResolver_getInboundImplementation is VersionedVerifierResolverSetup {
  function test_getInboundImplementation_InvalidCCVDataLength() public view {
    assertEq(s_versionedVerifierResolver.getInboundImplementation(""), address(0));
  }
}
