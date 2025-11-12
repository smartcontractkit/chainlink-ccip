// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_setCustomBlockConfirmationRateLimitConfig is TokenPoolV2Setup {
  function test_setCustomBlockConfirmationRateLimitConfig() public {
    RateLimiter.Config memory outboundConfig = RateLimiter.Config({isEnabled: true, capacity: 1e21, rate: 5e20});
    RateLimiter.Config memory inboundConfig = RateLimiter.Config({isEnabled: true, capacity: 2e21, rate: 1e21});
    TokenPool.CustomBlockConfirmationRateLimitConfigArgs[] memory args =
      new TokenPool.CustomBlockConfirmationRateLimitConfigArgs[](1);
    args[0] = TokenPool.CustomBlockConfirmationRateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundRateLimiterConfig: outboundConfig,
      inboundRateLimiterConfig: inboundConfig
    });

    vm.expectEmit();
    emit TokenPool.CustomBlockConfirmationRateLimitConfigured(DEST_CHAIN_SELECTOR, outboundConfig, inboundConfig);

    s_tokenPool.setCustomBlockConfirmationRateLimitConfig(args);

    (RateLimiter.TokenBucket memory outboundBucket, RateLimiter.TokenBucket memory inboundBucket) =
      s_tokenPool.getCurrentCustomBlockConfirmationRateLimiterState(DEST_CHAIN_SELECTOR);
    assertTrue(outboundBucket.isEnabled);
    assertEq(outboundBucket.capacity, outboundConfig.capacity);
    assertEq(outboundBucket.rate, outboundConfig.rate);
    assertEq(outboundBucket.tokens, outboundConfig.capacity);

    assertTrue(inboundBucket.isEnabled);
    assertEq(inboundBucket.capacity, inboundConfig.capacity);
    assertEq(inboundBucket.rate, inboundConfig.rate);
    assertEq(inboundBucket.tokens, inboundConfig.capacity);
  }

  function test_setCustomBlockConfirmationRateLimitConfig_RevertWhen_NonExistentChain() public {
    TokenPool.CustomBlockConfirmationRateLimitConfigArgs[] memory args =
      new TokenPool.CustomBlockConfirmationRateLimitConfigArgs[](1);
    args[0] = TokenPool.CustomBlockConfirmationRateLimitConfigArgs({
      remoteChainSelector: 999,
      outboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1, rate: 1}),
      inboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1, rate: 1})
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPool.NonExistentChain.selector, args[0].remoteChainSelector));
    s_tokenPool.setCustomBlockConfirmationRateLimitConfig(args);
  }

  function test_setCustomBlockConfirmationRateLimitConfig_RevertWhen_InvalidRateLimitRate_Outbound() public {
    TokenPool.CustomBlockConfirmationRateLimitConfigArgs[] memory args =
      new TokenPool.CustomBlockConfirmationRateLimitConfigArgs[](1);
    args[0] = TokenPool.CustomBlockConfirmationRateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1, rate: 2}),
      inboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1, rate: 1})
    });

    vm.expectRevert(
      abi.encodeWithSelector(RateLimiter.InvalidRateLimitRate.selector, args[0].outboundRateLimiterConfig)
    );
    s_tokenPool.setCustomBlockConfirmationRateLimitConfig(args);
  }

  function test_setCustomBlockConfirmationRateLimitConfig_RevertWhen_Unauthorized() public {
    vm.expectRevert(abi.encodeWithSelector(TokenPool.Unauthorized.selector, STRANGER));
    vm.startPrank(STRANGER);
    s_tokenPool.setCustomBlockConfirmationRateLimitConfig(new TokenPool.CustomBlockConfirmationRateLimitConfigArgs[](0));
  }
}
