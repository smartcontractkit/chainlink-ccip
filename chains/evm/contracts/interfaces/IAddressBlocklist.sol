// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice An address blocklist that pool hooks can query.
interface IAddressBlocklist {
  /// @notice Returns whether an account is blocked.
  /// @param account The address to check.
  /// @return blocked True if the account is blocked.
  function isBlocked(
    address account
  ) external view returns (bool blocked);
}
