// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolWithAllowListSetup} from "./TokenPoolWithAllowListSetup.t.sol";

contract TokenPoolWithAllowList_setDynamicConfig is TokenPoolWithAllowListSetup {
  function test_setDynamicConfig() public {
    address newRouter = makeAddr("newRouter");
    uint16 newMinBlockConfirmations = 5;
    uint256 newThresholdAmount = 1234;

    vm.expectEmit();
    emit TokenPool.DynamicConfigSet(newRouter, newMinBlockConfirmations, newThresholdAmount);

    s_tokenPool.setDynamicConfig(newRouter, newMinBlockConfirmations, newThresholdAmount);

    (address router, uint16 minBlockConfirmations, uint256 thresholdAmount) = s_tokenPool.getDynamicConfig();
    assertEq(newRouter, router);
    assertEq(newMinBlockConfirmations, minBlockConfirmations);
    assertEq(newThresholdAmount, thresholdAmount);
  }

  // Reverts

  function test_setDynamicConfig_RevertWhen_ZeroAddressInvalid() public {
    address newRouter = address(0);

    vm.expectRevert(TokenPool.ZeroAddressInvalid.selector);

    s_tokenPool.setDynamicConfig(newRouter, 0, 1234);
  }
}
