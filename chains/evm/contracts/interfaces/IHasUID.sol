// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice A contract that exposes a UID. The implementer defines what the UID means.
interface IHasUID {
  /// @notice Returns the UID of the contract.
  /// @return The UID of the contract.
  function uid() external view returns (bytes32);
}
