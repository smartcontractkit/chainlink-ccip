// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolRateLimitHelper} from "../../../pools/TokenPoolRateLimitHelper.sol";
import {TokenPoolRateLimitHelperSetup} from "./TokenPoolRateLimitHelperSetup.t.sol";

/// @notice Capacity-only ("quota") model: rate == 0 on both directions, so buckets never refill.
contract TokenPoolRateLimitHelper_zeroRateQuota is TokenPoolRateLimitHelperSetup {
  function setUp() public virtual override {
    super.setUp();

    TokenPool.RateLimitConfigArgs[] memory args = new TokenPool.RateLimitConfigArgs[](1);
    args[0] = TokenPool.RateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastFinality: false,
      outboundRateLimiterConfig: _cfg(true, 1000e18, 0),
      inboundRateLimiterConfig: _cfg(true, 2000e18, 0)
    });
    s_pool.setRateLimitConfig(args);
  }

  function test_zeroRateQuota_UntouchedRemainingQuotaPreserved() public {
    s_pool.consumeInbound(DEST_CHAIN_SELECTOR, 1500e18); // 500 left inbound.
    vm.warp(block.timestamp + 365 days); // Rate 0: no refill over time.
    (, RateLimiter.TokenBucket memory inbound) = _state(DEST_CHAIN_SELECTOR, false);
    assertEq(inbound.tokens, 500e18);

    s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 3000e18, 0));

    RateLimiter.TokenBucket memory outbound;
    (outbound, inbound) = _state(DEST_CHAIN_SELECTOR, false);
    assertEq(outbound.capacity, 3000e18);
    assertEq(outbound.tokens, 3000e18);
    assertEq(inbound.tokens, 500e18, "inbound remaining quota preserved");
    assertEq(inbound.capacity, 500e18, "stored inbound capacity now equals remaining");
    assertEq(inbound.rate, 0);

    // Exactly 500 more can come in, then it is exhausted.
    vm.expectRevert();
    s_pool.consumeInbound(DEST_CHAIN_SELECTOR, 500e18 + 1);
    s_pool.consumeInbound(DEST_CHAIN_SELECTOR, 500e18);
    vm.expectRevert();
    s_pool.consumeInbound(DEST_CHAIN_SELECTOR, 1);
  }

  function test_zeroRateQuota_ExhaustedUntouchedDirectionStaysExhausted() public {
    s_pool.consumeOutbound(DEST_CHAIN_SELECTOR, 1000e18); // Outbound fully used.

    s_helper.setInboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 50e18, 0));

    (RateLimiter.TokenBucket memory outbound, RateLimiter.TokenBucket memory inbound) =
      _state(DEST_CHAIN_SELECTOR, false);
    assertTrue(outbound.isEnabled);
    assertEq(outbound.capacity, 0);
    assertEq(outbound.tokens, 0);
    assertEq(inbound.capacity, 50e18);

    vm.expectRevert();
    s_pool.consumeOutbound(DEST_CHAIN_SELECTOR, 1);
  }

  /// @dev Documents pool semantics for the CHANGED direction: new capacity == new remaining quota (usage is forgotten).
  function test_zeroRateQuota_TargetDirectionCapacityBecomesRemainingQuota() public {
    s_pool.consumeOutbound(DEST_CHAIN_SELECTOR, 400e18);

    s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 1000e18, 0));

    (RateLimiter.TokenBucket memory outbound,) = _state(DEST_CHAIN_SELECTOR, false);
    assertEq(outbound.tokens, 1000e18);
  }

  function test_zeroRateQuota_EventCarriesOriginalUntouchedConfig() public {
    s_pool.consumeInbound(DEST_CHAIN_SELECTOR, 1500e18);

    vm.expectEmit(address(s_helper));
    emit TokenPoolRateLimitHelper.DirectionalRateLimitUpdated(
      DEST_CHAIN_SELECTOR,
      false,
      TokenPoolRateLimitHelper.Direction.Outbound,
      _helperCfg(true, 1000e18, 0),
      _helperCfg(true, 3000e18, 0),
      _helperCfg(true, 2000e18, 0),
      _helperCfg(true, 500e18, 0)
    );
    s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 3000e18, 0));
  }

  function test_zeroRateQuota_SwitchUntouchedToNonZeroRate() public {
    // Give inbound a refill rate, then change outbound; inbound config is kept as-is.
    s_helper.setInboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 2000e18, 1e18));
    s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 10e18, 0));

    (, RateLimiter.TokenBucket memory inbound) = _state(DEST_CHAIN_SELECTOR, false);
    assertEq(inbound.capacity, 2000e18);
    assertEq(inbound.rate, 1e18);
  }

  function testFuzz_zeroRateQuota_UntouchedRemainingNeverChanges(
    bool outbound,
    uint128 used,
    uint128 newCapacity
  ) public {
    used = uint128(bound(used, 0, outbound ? 2000e18 : 1000e18));
    if (outbound) s_pool.consumeInbound(DEST_CHAIN_SELECTOR, used);
    else s_pool.consumeOutbound(DEST_CHAIN_SELECTOR, used);

    (RateLimiter.TokenBucket memory outboundBefore, RateLimiter.TokenBucket memory inboundBefore) =
      _state(DEST_CHAIN_SELECTOR, false);
    uint256 remainingBefore = outbound ? inboundBefore.tokens : outboundBefore.tokens;

    if (outbound) s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, newCapacity, 0));
    else s_helper.setInboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, newCapacity, 0));

    (RateLimiter.TokenBucket memory outboundAfter, RateLimiter.TokenBucket memory inboundAfter) =
      _state(DEST_CHAIN_SELECTOR, false);
    assertEq(outbound ? inboundAfter.tokens : outboundAfter.tokens, remainingBefore);
    assertEq(outbound ? outboundAfter.tokens : inboundAfter.tokens, newCapacity);
  }
}
