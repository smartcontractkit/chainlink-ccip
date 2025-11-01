// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VersionedVerifierResolver} from "../../../ccvs/VersionedVerifierResolver.sol";
import {VersionedVerifierResolverSetup} from "./VersionedVerifierResolverSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract VersionedVerifierResolver_applyInboundImplementationUpdates is VersionedVerifierResolverSetup {
  function test_applyInboundImplementationUpdates() public {
    bytes4 newVersion = 0x01010101;
    address addedVerifier = makeAddr("AddedVerifier");
    address updatedVerifier = makeAddr("UpdatedVerifier");

    VersionedVerifierResolver.InboundImplementationArgs[] memory impls =
      new VersionedVerifierResolver.InboundImplementationArgs[](3);
    impls[0] = VersionedVerifierResolver.InboundImplementationArgs({version: newVersion, verifier: addedVerifier}); // Addition
    impls[1] =
      VersionedVerifierResolver.InboundImplementationArgs({version: INITIAL_VERSION_1, verifier: updatedVerifier}); // Update
    impls[2] = VersionedVerifierResolver.InboundImplementationArgs({version: INITIAL_VERSION_2, verifier: address(0)}); // Removal

    vm.expectEmit();
    emit VersionedVerifierResolver.InboundImplementationUpdated(newVersion, address(0), addedVerifier);
    vm.expectEmit();
    emit VersionedVerifierResolver.InboundImplementationUpdated(INITIAL_VERSION_1, s_initialVerifier1, updatedVerifier);
    vm.expectEmit();
    emit VersionedVerifierResolver.InboundImplementationRemoved(INITIAL_VERSION_2);

    s_versionedVerifierResolver.applyInboundImplementationUpdates(impls);

    assertEq(s_versionedVerifierResolver.getInboundImplementation(abi.encodePacked(newVersion)), addedVerifier);
    assertEq(s_versionedVerifierResolver.getInboundImplementation(abi.encodePacked(INITIAL_VERSION_1)), updatedVerifier);
    assertEq(s_versionedVerifierResolver.getInboundImplementation(abi.encodePacked(INITIAL_VERSION_2)), address(0));

    VersionedVerifierResolver.InboundImplementationArgs[] memory inboundImpls =
      s_versionedVerifierResolver.getAllInboundImplementations();
    assertEq(inboundImpls.length, 2);
    for (uint256 i = 0; i < inboundImpls.length; ++i) {
      if (inboundImpls[i].version != newVersion && inboundImpls[i].version != INITIAL_VERSION_1) {
        revert("Unexpected supported version");
      }
    }
  }

  function test_applyInboundImplementationUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_versionedVerifierResolver.applyInboundImplementationUpdates(
      new VersionedVerifierResolver.InboundImplementationArgs[](0)
    );
  }

  function test_applyInboundImplementationUpdates_RevertWhen_InvalidVersion() public {
    address newVerifier = makeAddr("NewVerifier");
    VersionedVerifierResolver.InboundImplementationArgs[] memory impls =
      new VersionedVerifierResolver.InboundImplementationArgs[](1);
    impls[0] = VersionedVerifierResolver.InboundImplementationArgs({version: bytes4(0), verifier: newVerifier});
    vm.expectRevert(abi.encodeWithSelector(VersionedVerifierResolver.InvalidVersion.selector, bytes4(0)));
    s_versionedVerifierResolver.applyInboundImplementationUpdates(impls);
  }
}
