// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @notice SP1Helios light client methods used by SuccinctZKVerifier.
/// @dev Anchored execution block hashes must be finalized and immutable.
interface ISP1Helios {
  /// @notice Returns the execution block hash anchored for a source chain block number, or zero if none.
  function executionBlockHashes(
    uint256 blockNumber
  ) external view returns (bytes32);

  /// @notice Returns the vkey for beacon chain update proofs.
  function lightClientVkey() external view returns (bytes32);

  /// @notice Returns the vkey for execution block proofs.
  function executionHeaderVkey() external view returns (bytes32);
}
