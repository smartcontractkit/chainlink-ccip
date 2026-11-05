// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {ITokenPoolV2RateLimits, TokenPoolRateLimitHelper} from "../../../pools/TokenPoolRateLimitHelper.sol";
import {TokenPoolRateLimitHelperSetup} from "./TokenPoolRateLimitHelperSetup.t.sol";

contract TokenPoolRateLimitHelper_constructor is TokenPoolRateLimitHelperSetup {
  function test_constructor() public view {
    assertEq(address(s_helper.i_pool()), address(s_pool));
    assertEq(s_helper.owner(), OWNER);
  }

  /// @dev The helper declares a local interface to avoid a BUSL dependency; its selectors must match the real pool.
  function test_constructor_InterfaceSelectorsMatchTokenPool() public pure {
    assertEq(ITokenPoolV2RateLimits.setRateLimitConfig.selector, TokenPool.setRateLimitConfig.selector);
    assertEq(ITokenPoolV2RateLimits.getCurrentRateLimiterState.selector, TokenPool.getCurrentRateLimiterState.selector);
    assertEq(ITokenPoolV2RateLimits.getDynamicConfig.selector, TokenPool.getDynamicConfig.selector);
  }

  function test_RevertWhen_ZeroPoolAddress() public {
    vm.expectRevert(TokenPoolRateLimitHelper.ZeroAddressNotAllowed.selector);
    new TokenPoolRateLimitHelper(address(0), OWNER);
  }
}
