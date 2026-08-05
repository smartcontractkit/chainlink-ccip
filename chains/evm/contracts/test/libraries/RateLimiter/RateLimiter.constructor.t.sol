// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {RateLimiterSetup} from "./RateLimiterSetup.t.sol";

contract RateLimiter_constructor is RateLimiterSetup {
  function test_Constructor() public view {
    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(i_blockTime, rateLimiter.lastUpdated);
  }
}
