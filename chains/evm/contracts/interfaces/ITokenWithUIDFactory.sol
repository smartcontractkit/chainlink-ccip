// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice A token factory that deploys at most one ITokenWithUID for each UID.
interface ITokenWithUIDFactory {
  /// @notice Returns the token that the factory deployed for a UID.
  /// @param uid The UID of the token.
  /// @return The token for the UID. Zero if the factory has not deployed a token for the UID.
  function tokenAddress(
    bytes32 uid
  ) external view returns (address);
}
