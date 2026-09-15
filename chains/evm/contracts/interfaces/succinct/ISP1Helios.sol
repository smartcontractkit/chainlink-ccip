// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @notice The subset of the SP1Helios light client used by the SuccinctZKVerifier. SP1Helios anchors finalized
/// source chain execution blocks on this chain after verifying a ZK proof of the beacon chain sync protocol.
/// @dev The verifier assumes that an anchored block hash never changes once set.
interface ISP1Helios {
  /// @notice Returns the execution block hash anchored for a source chain block number, or zero if none.
  function executionBlockHashes(
    uint256 blockNumber
  ) external view returns (bytes32);

  /// @notice Returns the vkey that beacon chain update proofs are verified against.
  function lightClientVkey() external view returns (bytes32);

  /// @notice Returns the vkey that execution block anchor proofs are verified against.
  function executionHeaderVkey() external view returns (bytes32);
}
