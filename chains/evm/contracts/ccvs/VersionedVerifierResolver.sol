// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../interfaces/ICrossChainVerifierResolver.sol";
import {IVersionedVerifier} from "../interfaces/IVersionedVerifier.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @notice Resolves and returns the appropriate verifier contract for the given outbound / inbound traffic.
/// @dev On source, the destChainSelector of a message is used to determine the verifier implementation to apply.
/// On destination, we must use the verifier version was applied on source, parsing this version from the ccvData.
contract VersionedVerifierResolver is ICrossChainVerifierResolver, Ownable2StepMsgSender {
  error InvalidCCVDataLength();
  error InvalidDestChainSelector(uint64 destChainSelector);
  error VersionMismatch(address verifier, bytes4 expected, bytes4 got);

  event InboundImplementationRemoved(bytes4 version);
  event OutboundImplementationRemoved(uint64 destChainSelector);
  event InboundImplementationUpdated(bytes4 version, address prevImpl, address newImpl);
  event OutboundImplementationUpdated(uint64 destChainSelector, address prevImpl, address newImpl);

  struct InboundImplementationArgs {
    bytes4 version; // ────╮ Verifier version.
    address verifier; // ──╯ Address of the verifier contract.
  }

  struct OutboundImplementationArgs {
    uint64 destChainSelector; // ──╮ Destination chain selector.
    address verifier; // ──────────╯ Address of the verifier contract.
  }

  /// @notice maps verifier versions to their implementation addresses, applied to inbound traffic.
  mapping(bytes4 version => address verifier) private s_inboundImplementations;
  /// @notice maps destination chain selectors to their implementation addresses, applied to outbound traffic.
  mapping(uint64 destChainSelector => address version) private s_outboundImplementations;

  /// @inheritdoc ICrossChainVerifierResolver
  function getInboundImplementation(
    bytes calldata ccvData
  ) external view returns (address) {
    if (ccvData.length < 4) {
      revert InvalidCCVDataLength();
    }
    return s_inboundImplementations[bytes4(ccvData[:4])];
  }

  /// @notice Returns the verifier contract for a given version.
  /// @param version The version of the verifier contract.
  /// @return verifierAddress The address of the verifier contract.
  function getInboundImplementationForVersion(
    bytes4 version
  ) external view returns (address) {
    return s_inboundImplementations[version];
  }

  /// @inheritdoc ICrossChainVerifierResolver
  function getOutboundImplementation(
    uint64 destChainSelector
  ) external view returns (address) {
    return s_outboundImplementations[destChainSelector];
  }

  /// @notice Updates inbound implementations.
  /// @param implementations Verifier versions and their corresponding contracts.
  function applyInboundImplementationUpdates(
    InboundImplementationArgs[] calldata implementations
  ) external onlyOwner {
    for (uint256 i = 0; i < implementations.length; ++i) {
      InboundImplementationArgs memory implementation = implementations[i];
      if (implementation.verifier == address(0)) {
        // If the verifier address is zero, we clear the implementation for this version.
        delete s_inboundImplementations[implementation.version];
        emit InboundImplementationRemoved(implementation.version);
        continue;
      }
      bytes4 expectedVersion = IVersionedVerifier(implementation.verifier).VERSION_TAG();
      if (expectedVersion != implementation.version) {
        revert VersionMismatch(implementation.verifier, expectedVersion, implementation.version);
      }
      address previous = s_inboundImplementations[implementation.version];
      s_inboundImplementations[implementation.version] = implementation.verifier;
      emit InboundImplementationUpdated(implementation.version, previous, implementation.verifier);
    }
  }

  /// @notice Updates outbound implementations.
  /// @param implementations Destination chain selectors and their corresponding verifier contracts.
  function applyOutboundImplementationUpdates(
    OutboundImplementationArgs[] calldata implementations
  ) external onlyOwner {
    for (uint256 i = 0; i < implementations.length; ++i) {
      OutboundImplementationArgs memory implementation = implementations[i];
      if (implementation.verifier == address(0)) {
        // If the verifier address is zero, we clear the implementation for this destination chain.
        delete s_outboundImplementations[implementation.destChainSelector];
        emit OutboundImplementationRemoved(implementation.destChainSelector);
        continue;
      }
      if (implementation.destChainSelector == 0) {
        revert InvalidDestChainSelector(implementation.destChainSelector);
      }
      address previous = s_outboundImplementations[implementation.destChainSelector];
      s_outboundImplementations[implementation.destChainSelector] = implementation.verifier;
      emit OutboundImplementationUpdated(implementation.destChainSelector, previous, implementation.verifier);
    }
  }
}
