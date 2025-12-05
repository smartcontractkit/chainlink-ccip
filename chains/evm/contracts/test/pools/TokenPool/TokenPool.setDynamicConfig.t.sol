// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";

contract TokenPoolWithAllowList_setDynamicConfig is AdvancedPoolHooksSetup {
  function test_setDynamicConfig() public {
    address newRouter = makeAddr("newRouter");
    address newRateLimitAdmin = makeAddr("newRateLimitAdmin");

    vm.expectEmit();
    emit TokenPool.DynamicConfigSet(newRouter, newRateLimitAdmin);

    s_tokenPool.setDynamicConfig(newRouter, newRateLimitAdmin);

    (address router, address rateLimitAdmin) = s_tokenPool.getDynamicConfig();
    assertEq(newRouter, router);
    assertEq(newRateLimitAdmin, rateLimitAdmin);
  }

  // Reverts

  function test_setDynamicConfig_RevertWhen_ZeroAddressInvalid() public {
    address newRouter = address(0);

    vm.expectRevert(TokenPool.ZeroAddressInvalid.selector);

    s_tokenPool.setDynamicConfig(newRouter, address(0));
  }
}
