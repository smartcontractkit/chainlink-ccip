// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @notice Minimal view of a light client that anchors finalized source chain execution blocks on this chain.
/// @dev SP1Helios exposes these getters natively. A different light client can satisfy the same interface.
interface IExecutionRootOracle {
  /// @notice Returns the finalized execution block hash for a source chain block number, or zero if not anchored.
  function executionBlockHashes(
    uint256 blockNumber
  ) external view returns (bytes32);

  /// @notice Returns the finalized receipts root for a source chain block number, or zero if not anchored.
  function executionReceiptsRoots(
    uint256 blockNumber
  ) external view returns (bytes32);
}
