// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IGetPoolConfigManager {
  /// @notice Returns the pool config manager used by the pool.
  /// @dev This method is named to never conflict with existing methods. A zero return indicates the pool does not
  /// use a pool config manager, letting offchain callers probe for pool config manager support.
  function getPoolConfigManager() external view returns (address poolConfigManager);
}
