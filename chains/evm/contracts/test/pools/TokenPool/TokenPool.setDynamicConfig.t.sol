// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";

contract TokenPoolWithAllowList_setDynamicConfig is AdvancedPoolHooksSetup {
  function test_setDynamicConfig() public {
    address newRouter = makeAddr("newRouter");
    address newRateLimitAdmin = makeAddr("newRateLimitAdmin");
    address newFeeAggregator = makeAddr("newFeeAggregator");

    vm.expectEmit();
    emit TokenPool.DynamicConfigSet(newRouter, newRateLimitAdmin, newFeeAggregator);

    s_tokenPool.setDynamicConfig(newRouter, newRateLimitAdmin, newFeeAggregator);

    (address router, address rateLimitAdmin, address feeAggregator) = s_tokenPool.getDynamicConfig();
    assertEq(newRouter, router);
    assertEq(newRateLimitAdmin, rateLimitAdmin);
    assertEq(newFeeAggregator, feeAggregator);
  }

  // Reverts

  function test_setDynamicConfig_RevertWhen_ZeroAddressInvalid() public {
    address newRouter = address(0);

    vm.expectRevert(TokenPool.ZeroAddressInvalid.selector);

    s_tokenPool.setDynamicConfig(newRouter, address(0), address(0));
  }
}
