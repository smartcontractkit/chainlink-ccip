// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_setCustomFinalityRateLimitConfig is TokenPoolV2Setup {
  function test_setCustomFinalityRateLimitConfig() public {
    RateLimiter.Config memory outboundConfig = RateLimiter.Config({isEnabled: true, capacity: 1e21, rate: 5e20});
    RateLimiter.Config memory inboundConfig = RateLimiter.Config({isEnabled: true, capacity: 2e21, rate: 1e21});
    TokenPool.CustomFinalityRateLimitConfigArgs[] memory args = new TokenPool.CustomFinalityRateLimitConfigArgs[](1);
    args[0] = TokenPool.CustomFinalityRateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundRateLimiterConfig: outboundConfig,
      inboundRateLimiterConfig: inboundConfig
    });

    s_tokenPool.setCustomFinalityRateLimitConfig(args);

    RateLimiter.TokenBucket memory outboundBucket =
      s_tokenPool.getCurrentOutboundCustomFinalityRateLimiterState(DEST_CHAIN_SELECTOR);
    assertTrue(outboundBucket.isEnabled);
    assertEq(outboundBucket.capacity, outboundConfig.capacity);
    assertEq(outboundBucket.rate, outboundConfig.rate);
    assertEq(outboundBucket.tokens, outboundConfig.capacity);

    RateLimiter.TokenBucket memory inboundBucket =
      s_tokenPool.getCurrentInboundCustomFinalityRateLimiterState(DEST_CHAIN_SELECTOR);
    assertTrue(inboundBucket.isEnabled);
    assertEq(inboundBucket.capacity, inboundConfig.capacity);
    assertEq(inboundBucket.rate, inboundConfig.rate);
    assertEq(inboundBucket.tokens, inboundConfig.capacity);
  }

  function test_setCustomFinalityRateLimitConfig_RevertWhen_NonExistentChain() public {
    TokenPool.CustomFinalityRateLimitConfigArgs[] memory args = new TokenPool.CustomFinalityRateLimitConfigArgs[](1);
    args[0] = TokenPool.CustomFinalityRateLimitConfigArgs({
      remoteChainSelector: 999,
      outboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1, rate: 1}),
      inboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1, rate: 1})
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPool.NonExistentChain.selector, args[0].remoteChainSelector));
    s_tokenPool.setCustomFinalityRateLimitConfig(args);
  }

  function test_setCustomFinalityRateLimitConfig_RevertWhen_InvalidRateLimitRate_Outbound() public {
    TokenPool.CustomFinalityRateLimitConfigArgs[] memory args = new TokenPool.CustomFinalityRateLimitConfigArgs[](1);
    args[0] = TokenPool.CustomFinalityRateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1, rate: 2}),
      inboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1, rate: 1})
    });

    vm.expectRevert(
      abi.encodeWithSelector(RateLimiter.InvalidRateLimitRate.selector, args[0].outboundRateLimiterConfig)
    );
    s_tokenPool.setCustomFinalityRateLimitConfig(args);
  }
}
