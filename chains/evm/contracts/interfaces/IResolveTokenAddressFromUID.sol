// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Resolves a UID to a token address.
interface IResolveTokenAddressFromUID {
  /// @notice Returns the token address for a UID.
  /// @param uid The UID to resolve.
  /// @return The token address for the UID. Zero if no token has the UID.
  function tokenAddress(
    bytes32 uid
  ) external view returns (address);
}
