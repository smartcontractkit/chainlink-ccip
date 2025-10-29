// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/// @notice Verifiers that expose an encoded version tag.
/// @dev Cross-chain verifiers that want compatibility with the VersionedVerifierResolver need to implement this interface.
interface IVersionedVerifier {
  /// @notice The verifier's version tag.
  /// @dev Example preimage for tag: bytes4(keccak256("ContractName Version"))
  function VERSION_TAG() external view returns (bytes4);
}
