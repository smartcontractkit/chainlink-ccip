// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VersionedVerifierResolver} from "../../../ccvs/VersionedVerifierResolver.sol";
import {VersionedVerifierResolverSetup} from "./VersionedVerifierResolverSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract VersionedVerifierResolver_applyOutboundImplementationUpdates is VersionedVerifierResolverSetup {
  function test_applyOutboundImplementationUpdates() public {
    uint64 newDestChainSelector = 3;
    address addedVerifier = makeAddr("AddedVerifier");
    address updatedVerifier = makeAddr("UpdatedVerifier");

    VersionedVerifierResolver.OutboundImplementationArgs[] memory impls =
      new VersionedVerifierResolver.OutboundImplementationArgs[](3);
    impls[0] = VersionedVerifierResolver.OutboundImplementationArgs({ // Addition
      destChainSelector: newDestChainSelector,
      verifier: addedVerifier
    });
    impls[1] = VersionedVerifierResolver.OutboundImplementationArgs({ // Update
      destChainSelector: INITIAL_DEST_CHAIN_SELECTOR_1,
      verifier: updatedVerifier
    });
    impls[2] = VersionedVerifierResolver.OutboundImplementationArgs({ // Removal
      destChainSelector: INITIAL_DEST_CHAIN_SELECTOR_2,
      verifier: address(0)
    });

    vm.expectEmit();
    emit VersionedVerifierResolver.OutboundImplementationUpdated(newDestChainSelector, address(0), addedVerifier);
    vm.expectEmit();
    emit VersionedVerifierResolver.OutboundImplementationUpdated(
      INITIAL_DEST_CHAIN_SELECTOR_1, s_initialVerifier1, updatedVerifier
    );
    vm.expectEmit();
    emit VersionedVerifierResolver.OutboundImplementationRemoved(INITIAL_DEST_CHAIN_SELECTOR_2);

    s_versionedVerifierResolver.applyOutboundImplementationUpdates(impls);

    assertEq(s_versionedVerifierResolver.getOutboundImplementation(newDestChainSelector), addedVerifier);
    assertEq(s_versionedVerifierResolver.getOutboundImplementation(INITIAL_DEST_CHAIN_SELECTOR_1), updatedVerifier);
    assertEq(s_versionedVerifierResolver.getOutboundImplementation(INITIAL_DEST_CHAIN_SELECTOR_2), address(0));

    uint64[] memory supportedDestChains = s_versionedVerifierResolver.getSupportedDestChains();
    assertEq(supportedDestChains.length, 2);
    for (uint256 i = 0; i < supportedDestChains.length; ++i) {
      if (supportedDestChains[i] != newDestChainSelector && supportedDestChains[i] != INITIAL_DEST_CHAIN_SELECTOR_1) {
        revert("Unexpected supported dest chain");
      }
    }
  }

  function test_applyOutboundImplementationUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_versionedVerifierResolver.applyOutboundImplementationUpdates(
      new VersionedVerifierResolver.OutboundImplementationArgs[](0)
    );
  }

  function test_applyOutboundImplementationUpdates_RevertWhen_InvalidDestChainSelector() public {
    VersionedVerifierResolver.OutboundImplementationArgs[] memory impls =
      new VersionedVerifierResolver.OutboundImplementationArgs[](1);
    impls[0] =
      VersionedVerifierResolver.OutboundImplementationArgs({destChainSelector: 0, verifier: makeAddr("NewVerifier")});
    vm.expectRevert(abi.encodeWithSelector(VersionedVerifierResolver.InvalidDestChainSelector.selector, uint64(0)));
    s_versionedVerifierResolver.applyOutboundImplementationUpdates(impls);
  }
}
