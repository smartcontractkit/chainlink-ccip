// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../interfaces/ICrossChainVerifierResolver.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

/// @notice Resolves and returns the appropriate verifier contract for the given outbound / inbound traffic.
contract VersionedVerifierResolver is ICrossChainVerifierResolver, Ownable2StepMsgSender {
  error InvalidCCVDataLength();
  error InvalidDestChainSelector(uint64 destChainSelector);
  error InvalidVersion(bytes4 version);
  error InboundImplementationAlreadyExists(bytes4 version);
  error InboundImplementationNotFound(bytes4 version);
  error OutboundImplementationAlreadyExists(uint64 destChainSelector);
  error OutboundImplementationNotFound(uint64 destChainSelector);
  error ZeroAddressNotAllowed();

  event InboundImplementationAdded(bytes4 version, address verifier);
  event InboundImplementationRemoved(bytes4 version);
  event OutboundImplementationAdded(uint64 destChainSelector, address verifier);
  event OutboundImplementationRemoved(uint64 destChainSelector);

  struct InboundImplementationArgs {
    bytes4 version; // ────╮ Verifier version
    address verifier; // ──╯ Address of the verifier contract
  }

  struct OutboundImplementationArgs {
    uint64 destChainSelector; // ──╮ Destination chain selector
    address verifier; // ──────────╯ Address of the verifier contract
  }

  /// @notice maps verifier versions to their implementation addresses, applied to inbound traffic
  mapping(bytes4 => address) private s_inboundImplementations;
  /// @notice maps destination chain selectors to their implementation addresses, applied to outbound traffic
  mapping(uint64 => address) private s_outboundImplementations;

  constructor() Ownable2StepMsgSender() {}

  /// @inheritdoc ICrossChainVerifierResolver
  function getInboundImplementation(
    bytes calldata ccvData
  ) external view returns (address) {
    if (ccvData.length < 4) {
      revert InvalidCCVDataLength();
    }
    return _getInboundImplementationForVersion(bytes4(ccvData[:4]));
  }

  /// @notice Returns the verifier contract for a given version.
  /// @param version The version of the verifier contract.
  /// @return verifierAddress The address of the verifier contract.
  function getInboundImplementationForVersion(
    bytes4 version
  ) external view returns (address) {
    return _getInboundImplementationForVersion(version);
  }

  /// @dev Internal function that enables reuse between functions.
  /// Validates that the inbound implementation exists before returning it.
  /// @param version The version of the verifier contract.
  /// @return verifierAddress The address of the verifier contract.
  function _getInboundImplementationForVersion(
    bytes4 version
  ) internal view returns (address) {
    if (s_inboundImplementations[version] == address(0)) {
      revert InboundImplementationNotFound(version);
    }
    return s_inboundImplementations[version];
  }

  /// @inheritdoc ICrossChainVerifierResolver
  function getOutboundImplementation(
    uint64 destChainSelector
  ) external view returns (address) {
    if (s_outboundImplementations[destChainSelector] == address(0)) {
      revert OutboundImplementationNotFound(destChainSelector);
    }
    return s_outboundImplementations[destChainSelector];
  }

  /// @notice Updates inbound implementations.
  /// @param versionsToRemove Versions that must no longer be supported by the resolver.
  /// @param implementationsToAdd New verifier contracts and their versions.
  function applyInboundImplementationUpdates(
    bytes4[] calldata versionsToRemove,
    InboundImplementationArgs[] calldata implementationsToAdd
  ) external onlyOwner {
    for (uint256 i = 0; i < versionsToRemove.length; i++) {
      bytes4 version = versionsToRemove[i];
      if (s_inboundImplementations[version] == address(0)) {
        revert InboundImplementationNotFound(version);
      }
      delete s_inboundImplementations[version];
      emit InboundImplementationRemoved(version);
    }
    for (uint256 i = 0; i < implementationsToAdd.length; i++) {
      InboundImplementationArgs memory implementation = implementationsToAdd[i];
      if (implementation.verifier == address(0)) {
        revert ZeroAddressNotAllowed();
      }
      if (implementation.version == bytes4(0)) {
        revert InvalidVersion(implementation.version);
      }
      if (s_inboundImplementations[implementation.version] != address(0)) {
        revert InboundImplementationAlreadyExists(implementation.version);
      }
      s_inboundImplementations[implementation.version] = implementation.verifier;
      emit InboundImplementationAdded(implementation.version, implementation.verifier);
    }
  }

  /// @notice Updates outbound implementations.
  /// @param chainsToRemove Destinations that must no longer be supported by the resolver.
  /// @param implementationsToAdd New verifier contracts and their destination chain selectors.
  function applyOutboundImplementationUpdates(
    uint64[] calldata chainsToRemove,
    OutboundImplementationArgs[] calldata implementationsToAdd
  ) external onlyOwner {
    for (uint256 i = 0; i < chainsToRemove.length; i++) {
      uint64 destChainSelector = chainsToRemove[i];
      if (s_outboundImplementations[destChainSelector] == address(0)) {
        revert OutboundImplementationNotFound(destChainSelector);
      }
      delete s_outboundImplementations[destChainSelector];
      emit OutboundImplementationRemoved(destChainSelector);
    }
    for (uint256 i = 0; i < implementationsToAdd.length; i++) {
      OutboundImplementationArgs memory implementation = implementationsToAdd[i];
      if (implementation.verifier == address(0)) {
        revert ZeroAddressNotAllowed();
      }
      if (implementation.destChainSelector == 0) {
        revert InvalidDestChainSelector(implementation.destChainSelector);
      }
      if (s_outboundImplementations[implementation.destChainSelector] != address(0)) {
        revert OutboundImplementationAlreadyExists(implementation.destChainSelector);
      }
      s_outboundImplementations[implementation.destChainSelector] = implementation.verifier;
      emit OutboundImplementationAdded(implementation.destChainSelector, implementation.verifier);
    }
  }

  /// @inheritdoc IERC165
  function supportsInterface(
    bytes4 interfaceId
  ) external pure returns (bool) {
    return interfaceId == type(ICrossChainVerifierResolver).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}
