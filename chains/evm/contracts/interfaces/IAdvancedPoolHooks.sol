// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "./IPoolV2.sol";

/// @notice Interface for AdvancedPoolHooks contract
interface IAdvancedPoolHooks {
  /// @notice Checks if the sender is allowed to perform an operation
  /// @param sender The address to check
  function checkAllowList(
    address sender
  ) external view;

  /// @notice Gets whether the allowlist functionality is enabled.
  /// @return true is enabled, false if not.
  function getAllowListEnabled() external view returns (bool);

  /// @notice Gets the allowed addresses.
  /// @return The allowed addresses.
  function getAllowList() external view returns (address[] memory);

  /// @notice Returns the set of required CCVs for transfers in a specific direction.
  /// @param remoteChainSelector The remote chain selector for this transfer.
  /// @param amount The amount being transferred.
  /// @param direction The direction of the transfer (Inbound or Outbound).
  /// @return requiredCCVs Set of required CCV addresses.
  function getRequiredCCVs(
    uint64 remoteChainSelector,
    uint256 amount,
    IPoolV2.MessageDirection direction
  ) external view returns (address[] memory requiredCCVs);

  /// @notice Gets the threshold amount above which additional CCVs are required
  /// @return The threshold amount
  function getThresholdAmount() external view returns (uint256);
}
