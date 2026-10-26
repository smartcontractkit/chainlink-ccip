// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";

import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_isSupportedToken is MultiTokenPoolSetup {
  function test_isSupportedToken() public view {
    assertTrue(s_pool.isSupportedToken(address(s_token)));
  }

  function test_isSupportedToken_FalseWhenNotRegistered() public {
    IPoolConfigManager.TokenConfig memory config = _defaultTokenConfig();
    config.isRegistered = false;
    s_poolConfigManager.setTokenConfig(config);

    assertFalse(s_pool.isSupportedToken(address(s_token)));
  }

  function test_isSupportedToken_FalseWhenFlavorNotServed() public {
    IPoolConfigManager.TokenConfig memory config = _defaultTokenConfig();
    config.lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.Lock;
    s_poolConfigManager.setTokenConfig(config);

    assertFalse(s_pool.isSupportedToken(address(s_token)));
  }
}
