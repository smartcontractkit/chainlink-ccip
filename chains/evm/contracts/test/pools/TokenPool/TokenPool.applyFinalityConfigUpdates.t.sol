// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract TokenPoolV2_applyFinalityConfigUpdates is TokenPoolV2Setup {
  function test_applyFinalityConfigUpdates() public {
    uint16 finalityThreshold = 100;
    uint16 customFinalityTransferFeeBps = 500; // 5%
    RateLimiter.Config memory outboundFastConfig = RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24});
    RateLimiter.Config memory inboundFastConfig = RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24});
    TokenPool.CustomFinalityRateLimitConfigArgs[] memory rateLimitArgs =
      new TokenPool.CustomFinalityRateLimitConfigArgs[](1);
    rateLimitArgs[0] = TokenPool.CustomFinalityRateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundRateLimiterConfig: outboundFastConfig,
      inboundRateLimiterConfig: inboundFastConfig
    });

    vm.expectEmit();
    emit TokenPool.FinalityThresholdUpdated(finalityThreshold);
    s_tokenPool.applyFinalityConfigUpdates(finalityThreshold, rateLimitArgs);

    uint16 storedFinalityThreshold = s_tokenPool.getCustomFinalityConfig();
    assertEq(storedFinalityThreshold, finalityThreshold);

    RateLimiter.TokenBucket memory outboundBucket = s_tokenPool.getFastOutboundBucket(DEST_CHAIN_SELECTOR);
    assertTrue(outboundBucket.isEnabled);
    assertEq(outboundBucket.capacity, outboundFastConfig.capacity);
    assertEq(outboundBucket.rate, outboundFastConfig.rate);
    assertEq(outboundBucket.tokens, outboundFastConfig.capacity);
    assertEq(outboundBucket.lastUpdated, uint32(block.timestamp));

    RateLimiter.TokenBucket memory inboundBucket = s_tokenPool.getFastInboundBucket(DEST_CHAIN_SELECTOR);
    assertTrue(inboundBucket.isEnabled);
    assertEq(inboundBucket.capacity, inboundFastConfig.capacity);
    assertEq(inboundBucket.rate, inboundFastConfig.rate);
    assertEq(inboundBucket.tokens, inboundFastConfig.capacity);
    assertEq(inboundBucket.lastUpdated, uint32(block.timestamp));
  }

  // Reverts
  function test_applyFinalityConfigUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);

    uint16 finalityThreshold = 100;
    uint16 customFinalityTransferFeeBps = 500; // 5%
    TokenPool.CustomFinalityRateLimitConfigArgs[] memory emptyRateLimitArgs =
      new TokenPool.CustomFinalityRateLimitConfigArgs[](0);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.applyFinalityConfigUpdates(finalityThreshold, emptyRateLimitArgs);
  }
}
