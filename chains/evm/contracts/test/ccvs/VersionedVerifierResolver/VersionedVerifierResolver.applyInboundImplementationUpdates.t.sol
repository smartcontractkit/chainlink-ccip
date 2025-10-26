// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VersionedVerifierResolver} from "../../../ccvs/VersionedVerifierResolver.sol";
import {VersionedVerifierResolverSetup} from "./VersionedVerifierResolverSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract VersionedVerifierResolver_applyInboundImplementationUpdates is VersionedVerifierResolverSetup {
  function test_applyInboundImplementationUpdates() public {
    bytes4 newVersion = 0x01010101;
    address newVerifier = makeAddr("NewVerifier");

    VersionedVerifierResolver.InboundImplementationArgs[] memory implementationsToAdd =
      new VersionedVerifierResolver.InboundImplementationArgs[](1);
    implementationsToAdd[0] =
      VersionedVerifierResolver.InboundImplementationArgs({version: newVersion, verifier: newVerifier});
    bytes4[] memory versionsToRemove = new bytes4[](1);
    versionsToRemove[0] = INITIAL_VERSION;

    vm.expectEmit();
    emit VersionedVerifierResolver.InboundImplementationRemoved(INITIAL_VERSION);
    vm.expectEmit();
    emit VersionedVerifierResolver.InboundImplementationAdded(newVersion, newVerifier);
    s_versionedVerifierResolver.applyInboundImplementationUpdates(versionsToRemove, implementationsToAdd);

    assertEq(s_versionedVerifierResolver.getInboundImplementationForVersion(newVersion), newVerifier);
    assertEq(s_versionedVerifierResolver.getInboundImplementation(abi.encodePacked(newVersion)), newVerifier);

    // Attempts to get the implementation for the initial version should revert
    vm.expectRevert(
      abi.encodeWithSelector(VersionedVerifierResolver.InboundImplementationNotFound.selector, INITIAL_VERSION)
    );
    s_versionedVerifierResolver.getInboundImplementationForVersion(INITIAL_VERSION);
    vm.expectRevert(
      abi.encodeWithSelector(VersionedVerifierResolver.InboundImplementationNotFound.selector, INITIAL_VERSION)
    );
    s_versionedVerifierResolver.getInboundImplementation(abi.encodePacked(INITIAL_VERSION));
  }

  function test_applyInboundImplementationUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_versionedVerifierResolver.applyInboundImplementationUpdates(
      new bytes4[](0), new VersionedVerifierResolver.InboundImplementationArgs[](0)
    );
  }

  function test_applyInboundImplementationUpdates_RevertWhen_InboundImplementationNotFound() public {
    bytes4[] memory versionsToRemove = new bytes4[](1);
    versionsToRemove[0] = UNKNOWN_VERSION;
    vm.expectRevert(
      abi.encodeWithSelector(VersionedVerifierResolver.InboundImplementationNotFound.selector, UNKNOWN_VERSION)
    );
    s_versionedVerifierResolver.applyInboundImplementationUpdates(
      versionsToRemove, new VersionedVerifierResolver.InboundImplementationArgs[](0)
    );
  }

  function test_applyInboundImplementationUpdates_RevertWhen_ZeroAddressNotAllowed() public {
    VersionedVerifierResolver.InboundImplementationArgs[] memory implementationsToAdd =
      new VersionedVerifierResolver.InboundImplementationArgs[](1);
    implementationsToAdd[0] =
      VersionedVerifierResolver.InboundImplementationArgs({version: INITIAL_VERSION, verifier: address(0)});
    vm.expectRevert(abi.encodeWithSelector(VersionedVerifierResolver.ZeroAddressNotAllowed.selector));
    s_versionedVerifierResolver.applyInboundImplementationUpdates(new bytes4[](0), implementationsToAdd);
  }

  function test_applyInboundImplementationUpdates_RevertWhen_InvalidVersion() public {
    VersionedVerifierResolver.InboundImplementationArgs[] memory implementationsToAdd =
      new VersionedVerifierResolver.InboundImplementationArgs[](1);
    implementationsToAdd[0] =
      VersionedVerifierResolver.InboundImplementationArgs({version: bytes4(0), verifier: makeAddr("NewVerifier")});
    vm.expectRevert(abi.encodeWithSelector(VersionedVerifierResolver.InvalidVersion.selector, bytes4(0)));
    s_versionedVerifierResolver.applyInboundImplementationUpdates(new bytes4[](0), implementationsToAdd);
  }

  function test_applyInboundImplementationUpdates_RevertWhen_InboundImplementationAlreadyExists() public {
    VersionedVerifierResolver.InboundImplementationArgs[] memory implementationsToAdd =
      new VersionedVerifierResolver.InboundImplementationArgs[](1);
    implementationsToAdd[0] =
      VersionedVerifierResolver.InboundImplementationArgs({version: INITIAL_VERSION, verifier: s_initialVerifier});
    vm.expectRevert(
      abi.encodeWithSelector(VersionedVerifierResolver.InboundImplementationAlreadyExists.selector, INITIAL_VERSION)
    );
    s_versionedVerifierResolver.applyInboundImplementationUpdates(new bytes4[](0), implementationsToAdd);
  }
}
