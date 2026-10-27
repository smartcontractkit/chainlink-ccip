// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice A token that a token factory deployed for a UID. The issuer gives the same asset the same UID on every
/// chain.
interface ITokenWithUID {
  /// @notice Returns the UID of the token.
  /// @return The UID that the token factory deployed this token for.
  function uid() external view returns (bytes32);
}
