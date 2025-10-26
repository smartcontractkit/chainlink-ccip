// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VersionedVerifierResolver} from "../../../ccvs/VersionedVerifierResolver.sol";
import {VersionedVerifierResolverSetup} from "./VersionedVerifierResolverSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract VersionedVerifierResolver_applyOutboundImplementationUpdates is VersionedVerifierResolverSetup {
  function test_applyOutboundImplementationUpdates() public {
    uint64 newDestChainSelector = 3;
    address newVerifier = makeAddr("NewVerifier");

    VersionedVerifierResolver.OutboundImplementationArgs[] memory implementationsToAdd =
      new VersionedVerifierResolver.OutboundImplementationArgs[](1);
    implementationsToAdd[0] = VersionedVerifierResolver.OutboundImplementationArgs({
      destChainSelector: newDestChainSelector,
      verifier: newVerifier
    });
    uint64[] memory chainsToRemove = new uint64[](1);
    chainsToRemove[0] = INITIAL_DEST_CHAIN_SELECTOR;

    vm.expectEmit();
    emit VersionedVerifierResolver.OutboundImplementationRemoved(INITIAL_DEST_CHAIN_SELECTOR);
    vm.expectEmit();
    emit VersionedVerifierResolver.OutboundImplementationAdded(newDestChainSelector, newVerifier);
    s_versionedVerifierResolver.applyOutboundImplementationUpdates(chainsToRemove, implementationsToAdd);

    assertEq(s_versionedVerifierResolver.getOutboundImplementation(newDestChainSelector), newVerifier);

    // Attempts to get the implementation for the initial destination chain selector should revert
    vm.expectRevert(
      abi.encodeWithSelector(
        VersionedVerifierResolver.OutboundImplementationNotFound.selector, INITIAL_DEST_CHAIN_SELECTOR
      )
    );
    s_versionedVerifierResolver.getOutboundImplementation(INITIAL_DEST_CHAIN_SELECTOR);
  }

  function test_applyOutboundImplementationUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_versionedVerifierResolver.applyOutboundImplementationUpdates(
      new uint64[](0), new VersionedVerifierResolver.OutboundImplementationArgs[](0)
    );
  }

  function test_applyOutboundImplementationUpdates_RevertWhen_OutboundImplementationNotFound() public {
    uint64[] memory chainsToRemove = new uint64[](1);
    chainsToRemove[0] = UNKNOWN_DEST_CHAIN_SELECTOR;
    vm.expectRevert(
      abi.encodeWithSelector(
        VersionedVerifierResolver.OutboundImplementationNotFound.selector, UNKNOWN_DEST_CHAIN_SELECTOR
      )
    );
    s_versionedVerifierResolver.applyOutboundImplementationUpdates(
      chainsToRemove, new VersionedVerifierResolver.OutboundImplementationArgs[](0)
    );
  }

  function test_applyOutboundImplementationUpdates_RevertWhen_ZeroAddressNotAllowed() public {
    VersionedVerifierResolver.OutboundImplementationArgs[] memory implementationsToAdd =
      new VersionedVerifierResolver.OutboundImplementationArgs[](1);
    implementationsToAdd[0] = VersionedVerifierResolver.OutboundImplementationArgs({
      destChainSelector: INITIAL_DEST_CHAIN_SELECTOR,
      verifier: address(0)
    });
    vm.expectRevert(abi.encodeWithSelector(VersionedVerifierResolver.ZeroAddressNotAllowed.selector));
    s_versionedVerifierResolver.applyOutboundImplementationUpdates(new uint64[](0), implementationsToAdd);
  }

  function test_applyOutboundImplementationUpdates_RevertWhen_InvalidDestChainSelector() public {
    VersionedVerifierResolver.OutboundImplementationArgs[] memory implementationsToAdd =
      new VersionedVerifierResolver.OutboundImplementationArgs[](1);
    implementationsToAdd[0] =
      VersionedVerifierResolver.OutboundImplementationArgs({destChainSelector: 0, verifier: makeAddr("NewVerifier")});
    vm.expectRevert(abi.encodeWithSelector(VersionedVerifierResolver.InvalidDestChainSelector.selector, uint64(0)));
    s_versionedVerifierResolver.applyOutboundImplementationUpdates(new uint64[](0), implementationsToAdd);
  }

  function test_applyOutboundImplementationUpdates_RevertWhen_OutboundImplementationAlreadyExists() public {
    VersionedVerifierResolver.OutboundImplementationArgs[] memory implementationsToAdd =
      new VersionedVerifierResolver.OutboundImplementationArgs[](1);
    implementationsToAdd[0] = VersionedVerifierResolver.OutboundImplementationArgs({
      destChainSelector: INITIAL_DEST_CHAIN_SELECTOR,
      verifier: s_initialVerifier
    });
    vm.expectRevert(
      abi.encodeWithSelector(
        VersionedVerifierResolver.OutboundImplementationAlreadyExists.selector, INITIAL_DEST_CHAIN_SELECTOR
      )
    );
    s_versionedVerifierResolver.applyOutboundImplementationUpdates(new uint64[](0), implementationsToAdd);
  }
}
