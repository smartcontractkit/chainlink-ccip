// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IVersionedVerifier} from "../../../interfaces/IVersionedVerifier.sol";

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

    vm.mockCall(addedVerifier, abi.encodeWithSelector(IVersionedVerifier.VERSION_TAG.selector), abi.encode(newVersion));
    vm.mockCall(
      updatedVerifier, abi.encodeWithSelector(IVersionedVerifier.VERSION_TAG.selector), abi.encode(INITIAL_VERSION_1)
    );
    s_versionedVerifierResolver.applyInboundImplementationUpdates(impls);

    assertEq(s_versionedVerifierResolver.getInboundImplementationForVersion(newVersion), addedVerifier);
    assertEq(s_versionedVerifierResolver.getInboundImplementation(abi.encodePacked(newVersion)), addedVerifier);
    assertEq(s_versionedVerifierResolver.getInboundImplementationForVersion(INITIAL_VERSION_1), updatedVerifier);
    assertEq(s_versionedVerifierResolver.getInboundImplementation(abi.encodePacked(INITIAL_VERSION_1)), updatedVerifier);
    assertEq(s_versionedVerifierResolver.getInboundImplementationForVersion(INITIAL_VERSION_2), address(0));
    assertEq(s_versionedVerifierResolver.getInboundImplementation(abi.encodePacked(INITIAL_VERSION_2)), address(0));
  }

  function test_applyInboundImplementationUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_versionedVerifierResolver.applyInboundImplementationUpdates(
      new VersionedVerifierResolver.InboundImplementationArgs[](0)
    );
  }

  function test_applyInboundImplementationUpdates_RevertWhen_VersionMismatch() public {
    bytes4 expected = 0x01010101;
    bytes4 inputted = 0x02020202;
    address newVerifier = makeAddr("NewVerifier");
    VersionedVerifierResolver.InboundImplementationArgs[] memory impls =
      new VersionedVerifierResolver.InboundImplementationArgs[](1);
    impls[0] = VersionedVerifierResolver.InboundImplementationArgs({version: inputted, verifier: newVerifier});
    vm.mockCall(newVerifier, abi.encodeWithSelector(IVersionedVerifier.VERSION_TAG.selector), abi.encode(expected));
    vm.expectRevert(
      abi.encodeWithSelector(VersionedVerifierResolver.VersionMismatch.selector, newVerifier, expected, inputted)
    );
    s_versionedVerifierResolver.applyInboundImplementationUpdates(impls);
  }
}
