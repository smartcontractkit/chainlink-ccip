// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

contract PoolConfigManager_getAllConfiguredTokens is PoolConfigManagerSetup {
  /// @notice Passes the page arguments to the connection store.
  function test_getAllConfiguredTokens() public view {
    assertEq(_arrayOf(s_tokenA), s_poolConfigManager.getAllConfiguredTokens(0, 1));
    assertEq(_arrayOf(s_tokenB), s_poolConfigManager.getAllConfiguredTokens(1, 1));
  }
}
