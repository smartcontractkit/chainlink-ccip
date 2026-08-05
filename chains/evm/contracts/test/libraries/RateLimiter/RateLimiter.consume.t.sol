// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {RateLimiterSetup} from "./RateLimiterSetup.t.sol";

contract RateLimiter_consume is RateLimiterSetup {
  address internal s_token = address(100);

  function test_ConsumeAggregateValue() public {
    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(i_blockTime, rateLimiter.lastUpdated);

    uint256 requestTokens = 50;

    s_helper.consume(requestTokens, address(0));

    rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity - requestTokens, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(i_blockTime, rateLimiter.lastUpdated);
  }

  function test_Refill() public {
    uint256 requestTokens = 50;

    s_helper.consume(requestTokens, address(0));

    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity - requestTokens, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(i_blockTime, rateLimiter.lastUpdated);

    uint256 warpTime = 4;
    vm.warp(i_blockTime + warpTime);

    s_helper.consume(requestTokens, address(0));

    rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity - requestTokens * 2 + warpTime * s_config.rate, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(i_blockTime + warpTime, rateLimiter.lastUpdated);
  }

  function test_consume_ConsumeWhenDisabled() public {
    s_helper.consume(0, address(0));

    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.capacity, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);

    RateLimiter.Config memory disableConfig = RateLimiter.Config({isEnabled: false, rate: 0, capacity: 0});

    s_helper.setTokenBucketConfig(disableConfig);

    // Should not revert when consuming any amount of tokens when disabled.
    uint256 requestTokens = 50;
    s_helper.consume(requestTokens, address(0));

    rateLimiter = s_helper.getRateLimiter();
    assertEq(disableConfig.capacity, rateLimiter.tokens);
    assertEq(disableConfig.isEnabled, rateLimiter.isEnabled);
  }

  // Reverts

  function test_RevertWhen_TokenMaxCapacityExceeded() public {
    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    vm.expectRevert(
      abi.encodeWithSelector(
        RateLimiter.TokenMaxCapacityExceeded.selector, rateLimiter.capacity, rateLimiter.capacity + 1, s_token
      )
    );
    s_helper.consume(rateLimiter.capacity + 1, s_token);
  }

  function test_RevertWhen_ConsumingMoreThanUint128() public {
    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    uint256 request = uint256(type(uint128).max) + 1;

    vm.expectRevert(
      abi.encodeWithSelector(RateLimiter.TokenMaxCapacityExceeded.selector, rateLimiter.capacity, request, address(1))
    );
    s_helper.consume(request, address(1));
  }

  function test_RevertWhen_AggregateValueRateLimitReached() public {
    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    uint256 overLimit = 20;
    uint256 requestTokens1 = rateLimiter.capacity / 2;
    uint256 requestTokens2 = rateLimiter.capacity / 2 + overLimit;

    uint256 waitInSeconds = overLimit / rateLimiter.rate;

    s_helper.consume(requestTokens1, address(0));

    vm.expectRevert(
      abi.encodeWithSelector(
        RateLimiter.TokenRateLimitReached.selector, waitInSeconds, rateLimiter.capacity - requestTokens1, address(0)
      )
    );
    s_helper.consume(requestTokens2, address(0));
  }

  function test_RevertWhen_TokenRateLimitReached() public {
    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    uint256 overLimit = 20;
    uint256 requestTokens1 = rateLimiter.capacity / 2;
    uint256 requestTokens2 = rateLimiter.capacity / 2 + overLimit;

    uint256 waitInSeconds = overLimit / rateLimiter.rate;

    s_helper.consume(requestTokens1, s_token);

    vm.expectRevert(
      abi.encodeWithSelector(
        RateLimiter.TokenRateLimitReached.selector, waitInSeconds, rateLimiter.capacity - requestTokens1, s_token
      )
    );
    s_helper.consume(requestTokens2, s_token);
  }

  function test_RevertWhen_RateLimitReachedOverConsecutiveBlocks() public {
    uint256 initBlockTime = i_blockTime + 10000;
    vm.warp(initBlockTime);

    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    s_helper.consume(rateLimiter.capacity, address(0));

    vm.warp(initBlockTime + 1);

    // Over rate limit by 1, force 1 second wait
    uint256 overLimit = 1;

    vm.expectRevert(abi.encodeWithSelector(RateLimiter.TokenRateLimitReached.selector, 1, rateLimiter.rate, address(0)));
    s_helper.consume(rateLimiter.rate + overLimit, address(0));
  }
}
