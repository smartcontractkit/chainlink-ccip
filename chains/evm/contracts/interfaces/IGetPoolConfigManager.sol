// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Exposes the pool config manager associated with a pool.
interface IGetPoolConfigManager {
  /// @notice Returns the pool config manager used by the pool.
  /// @dev A zero address return indicates the pool does not use a pool config manager.
  /// @return poolConfigManager The pool config manager address, or zero if the pool does not use one.
  function getPoolConfigManager() external view returns (address poolConfigManager);
}
