// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../interfaces/ICrossChainVerifierResolver.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.0.2/utils/structs/EnumerableSet.sol";

/// @notice Resolves and returns the appropriate verifier contract for the given outbound / inbound traffic.
/// @dev On source, the destChainSelector of a message is used to determine the verifier implementation to apply.
/// On destination, we must use the verifier version was applied on source, parsing this version from the ccvData.
contract VersionedVerifierResolver is ICrossChainVerifierResolver, ITypeAndVersion, Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.UintSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;

  error InvalidCCVDataLength();
  error InvalidDestChainSelector(uint64 destChainSelector);
  error InvalidVersion(bytes4 version);

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

  string public constant override typeAndVersion = "VersionedVerifierResolver 1.7.0-dev";

  /// @notice maps verifier versions to their implementation addresses, applied to inbound traffic.
  mapping(bytes4 version => address verifier) private s_versionToInboundImplementation;
  /// @notice all supported verifier versions.
  EnumerableSet.Bytes32Set private s_supportedVerifierVersions;
  /// @notice maps destination chain selectors to their implementation addresses, applied to outbound traffic.
  mapping(uint64 destChainSelector => address version) private s_destChainToOutboundImplementation;
  /// @notice all supported destination chains.
  EnumerableSet.UintSet private s_supportedDestChains;

  /// @inheritdoc ICrossChainVerifierResolver
  function getInboundImplementation(
    bytes calldata ccvData
  ) external view returns (address) {
    if (ccvData.length < 4) {
      revert InvalidCCVDataLength();
    }
    return s_versionToInboundImplementation[bytes4(ccvData[:4])];
  }

  /// @notice Returns the verifier contract for a given version.
  /// @param version The version of the verifier contract.
  /// @return verifierAddress The address of the verifier contract.
  function getInboundImplementationForVersion(
    bytes4 version
  ) external view returns (address) {
    return s_versionToInboundImplementation[version];
  }

  /// @notice Returns all supported verifier versions.
  function getSupportedVerifierVersions() external view returns (bytes4[] memory) {
    bytes4[] memory versions = new bytes4[](s_supportedVerifierVersions.length());
    for (uint256 i = 0; i < s_supportedVerifierVersions.length(); ++i) {
      versions[i] = bytes4(s_supportedVerifierVersions.at(i));
    }
    return versions;
  }

  /// @inheritdoc ICrossChainVerifierResolver
  function getOutboundImplementation(
    uint64 destChainSelector
  ) external view returns (address) {
    return s_destChainToOutboundImplementation[destChainSelector];
  }

  /// @notice Returns all supported destination chains.
  function getSupportedDestChains() external view returns (uint64[] memory) {
    uint64[] memory destChains = new uint64[](s_supportedDestChains.length());
    for (uint256 i = 0; i < s_supportedDestChains.length(); ++i) {
      destChains[i] = uint64(s_supportedDestChains.at(i));
    }
    return destChains;
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
        delete s_versionToInboundImplementation[implementation.version];
        s_supportedVerifierVersions.remove(bytes32(implementation.version));
        emit InboundImplementationRemoved(implementation.version);
        continue;
      }
      if (implementation.version == bytes4(0)) {
        revert InvalidVersion(implementation.version);
      }
      address previous = s_versionToInboundImplementation[implementation.version];
      s_versionToInboundImplementation[implementation.version] = implementation.verifier;
      s_supportedVerifierVersions.add(bytes32(implementation.version));
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
        delete s_destChainToOutboundImplementation[implementation.destChainSelector];
        s_supportedDestChains.remove(implementation.destChainSelector);
        emit OutboundImplementationRemoved(implementation.destChainSelector);
        continue;
      }
      if (implementation.destChainSelector == 0) {
        revert InvalidDestChainSelector(implementation.destChainSelector);
      }
      address previous = s_destChainToOutboundImplementation[implementation.destChainSelector];
      s_destChainToOutboundImplementation[implementation.destChainSelector] = implementation.verifier;
      s_supportedDestChains.add(implementation.destChainSelector);
      emit OutboundImplementationUpdated(implementation.destChainSelector, previous, implementation.verifier);
    }
  }
}
