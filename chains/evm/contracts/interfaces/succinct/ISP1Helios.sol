// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @notice The subset of the SP1Helios light client used by the SuccinctZKVerifier. SP1Helios anchors finalized
/// source chain execution blocks on this chain after verifying a ZK proof of the beacon chain sync protocol.
/// @dev Both mappings are append-only: once a block number is anchored its values never change.
interface ISP1Helios {
  /// @notice Returns the execution block hash anchored for a source chain block number, or zero if none.
  function executionBlockHashes(
    uint256 blockNumber
  ) external view returns (bytes32);

  /// @notice Returns the receipts root anchored for a source chain block number, or zero if none.
  function executionReceiptsRoots(
    uint256 blockNumber
  ) external view returns (bytes32);
}
