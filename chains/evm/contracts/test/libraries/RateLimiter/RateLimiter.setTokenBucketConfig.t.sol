// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {RateLimiterSetup} from "./RateLimiterSetup.t.sol";

contract RateLimiter_setTokenBucketConfig is RateLimiterSetup {
  function test_setTokenBucketConfig_SettingConfigRefills() public {
    RateLimiter.TokenBucket memory bucket = s_helper.currentTokenBucketState();
    assertEq(s_config.rate, bucket.rate);
    assertEq(s_config.capacity, bucket.capacity);
    assertEq(s_config.capacity, bucket.tokens);
    assertEq(s_config.isEnabled, bucket.isEnabled);
    assertEq(i_blockTime, bucket.lastUpdated);

    s_config = RateLimiter.Config({isEnabled: true, rate: uint128(bucket.rate * 2), capacity: bucket.capacity * 8});

    s_helper.setTokenBucketConfig(s_config);

    bucket = s_helper.currentTokenBucketState();
    assertEq(s_config.rate, bucket.rate);
    assertEq(s_config.capacity, bucket.capacity);
    assertEq(s_config.capacity, bucket.tokens);
    assertEq(s_config.isEnabled, bucket.isEnabled);
    assertEq(i_blockTime, bucket.lastUpdated);
  }
}
