// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolWithAllowListSetup} from "./TokenPoolWithAllowListSetup.t.sol";

contract TokenPoolWithAllowList_setDynamicConfig is TokenPoolWithAllowListSetup {
  function test_setDynamicConfig() public {
    address newRouter = makeAddr("newRouter");
    uint96 newThresholdAmount = 1234;

    vm.expectEmit();
    emit TokenPool.DynamicConfigSet(newRouter, newThresholdAmount);

    s_tokenPool.setDynamicConfig(newRouter, newThresholdAmount);

    (address router, uint96 thresholdAmount) = s_tokenPool.getDynamicConfig();
    assertEq(newRouter, router);
    assertEq(newThresholdAmount, thresholdAmount);
  }

  // Reverts

  function test_setDynamicConfig_RevertWhen_ZeroAddressInvalid() public {
    address newRouter = address(0);

    vm.expectRevert(TokenPool.ZeroAddressInvalid.selector);

    s_tokenPool.setDynamicConfig(newRouter, 1234);
  }
}
