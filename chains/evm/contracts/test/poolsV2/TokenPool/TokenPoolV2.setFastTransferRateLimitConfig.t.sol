// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";

import {TokenPool as TokenPoolV1} from "../../../pools/TokenPool.sol";
import {TokenPool} from "../../../poolsV2/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_setFastTransferRateLimitConfig is TokenPoolV2Setup {
  function test_setFastTransferRateLimitConfig() public {
    RateLimiter.Config memory outboundConfig = RateLimiter.Config({isEnabled: true, capacity: 1e21, rate: 5e20});
    RateLimiter.Config memory inboundConfig = RateLimiter.Config({isEnabled: true, capacity: 2e21, rate: 1e21});
    TokenPool.FastTransferRateLimitConfigArgs[] memory args = new TokenPool.FastTransferRateLimitConfigArgs[](1);
    args[0] = TokenPool.FastTransferRateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundRateLimiterConfig: outboundConfig,
      inboundRateLimiterConfig: inboundConfig
    });

    s_tokenPool.setFastTransferRateLimitConfig(args);

    RateLimiter.TokenBucket memory outboundBucket = s_tokenPool.getFastOutboundBucket(DEST_CHAIN_SELECTOR);
    assertTrue(outboundBucket.isEnabled);
    assertEq(outboundBucket.capacity, outboundConfig.capacity);
    assertEq(outboundBucket.rate, outboundConfig.rate);
    assertEq(outboundBucket.tokens, outboundConfig.capacity);

    RateLimiter.TokenBucket memory inboundBucket = s_tokenPool.getFastInboundBucket(DEST_CHAIN_SELECTOR);
    assertTrue(inboundBucket.isEnabled);
    assertEq(inboundBucket.capacity, inboundConfig.capacity);
    assertEq(inboundBucket.rate, inboundConfig.rate);
    assertEq(inboundBucket.tokens, inboundConfig.capacity);
  }

  function test_setFastTransferRateLimitConfig_RevertWhen_NonExistentChain() public {
    TokenPool.FastTransferRateLimitConfigArgs[] memory args = new TokenPool.FastTransferRateLimitConfigArgs[](1);
    args[0] = TokenPool.FastTransferRateLimitConfigArgs({
      remoteChainSelector: 999,
      outboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1, rate: 1}),
      inboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1, rate: 1})
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPoolV1.NonExistentChain.selector, uint64(999)));
    s_tokenPool.setFastTransferRateLimitConfig(args);
  }

  function test_setFastTransferRateLimitConfig_RevertWhen_InvalidRateLimitRate_Outbound() public {
    TokenPool.FastTransferRateLimitConfigArgs[] memory args = new TokenPool.FastTransferRateLimitConfigArgs[](1);
    args[0] = TokenPool.FastTransferRateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1, rate: 2}),
      inboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1, rate: 1})
    });

    vm.expectRevert(
      abi.encodeWithSelector(RateLimiter.InvalidRateLimitRate.selector, args[0].outboundRateLimiterConfig)
    );
    s_tokenPool.setFastTransferRateLimitConfig(args);
  }
}
