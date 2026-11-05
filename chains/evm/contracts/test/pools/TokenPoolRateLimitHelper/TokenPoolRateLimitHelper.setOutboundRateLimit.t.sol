// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolRateLimitHelper} from "../../../pools/TokenPoolRateLimitHelper.sol";
import {TokenPoolRateLimitHelperSetup} from "./TokenPoolRateLimitHelperSetup.t.sol";

contract TokenPoolRateLimitHelper_setOutboundRateLimit is TokenPoolRateLimitHelperSetup {
  function test_setOutboundRateLimit_LeavesInboundConfig() public {
    vm.startPrank(s_manager);
    s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 3000e18, 30e18));

    (RateLimiter.TokenBucket memory outbound, RateLimiter.TokenBucket memory inbound) =
      _state(DEST_CHAIN_SELECTOR, false);
    assertTrue(outbound.isEnabled);
    assertEq(outbound.capacity, 3000e18);
    assertEq(outbound.rate, 30e18);
    assertTrue(inbound.isEnabled);
    assertEq(inbound.capacity, 2000e18);
    assertEq(inbound.rate, 20e18);
  }

  function test_setOutboundRateLimit_UntouchedInboundDisabledStaysDisabled() public {
    s_helper.setInboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(false, 0, 0));

    s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 7e18, 1e18));

    (RateLimiter.TokenBucket memory outbound, RateLimiter.TokenBucket memory inbound) =
      _state(DEST_CHAIN_SELECTOR, false);
    assertFalse(inbound.isEnabled);
    assertEq(inbound.capacity, 0);
    assertEq(outbound.capacity, 7e18);
  }

  function test_setOutboundRateLimit_OtherChainUntouched() public {
    vm.startPrank(s_manager);
    s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 1e18, 1e18));

    (RateLimiter.TokenBucket memory outbound, RateLimiter.TokenBucket memory inbound) =
      _state(OTHER_CHAIN_SELECTOR, false);
    assertEq(outbound.capacity, 500e18);
    assertEq(inbound.capacity, 600e18);
  }

  function test_setOutboundRateLimit_FastFinalityBucketsIndependent() public {
    // Fast-finality buckets start disabled; set only the fast-finality outbound bucket.
    vm.startPrank(s_manager);
    s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, true, _helperCfg(true, 100e18, 1e18));

    (RateLimiter.TokenBucket memory ffOutbound, RateLimiter.TokenBucket memory ffInbound) =
      _state(DEST_CHAIN_SELECTOR, true);
    assertTrue(ffOutbound.isEnabled);
    assertEq(ffOutbound.capacity, 100e18);
    assertFalse(ffInbound.isEnabled);

    (RateLimiter.TokenBucket memory outbound, RateLimiter.TokenBucket memory inbound) =
      _state(DEST_CHAIN_SELECTOR, false);
    assertEq(outbound.capacity, 1000e18);
    assertEq(inbound.capacity, 2000e18);
  }

  /// @dev Documents the unavoidable side effect: a non-zero-rate untouched bucket is refilled to capacity.
  function test_setOutboundRateLimit_UntouchedNonZeroRateBucketIsRefilled() public {
    s_pool.consumeInbound(DEST_CHAIN_SELECTOR, 1500e18);
    (, RateLimiter.TokenBucket memory inbound) = _state(DEST_CHAIN_SELECTOR, false);
    assertEq(inbound.tokens, 500e18);

    vm.startPrank(s_manager);
    s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 3000e18, 30e18));

    (, inbound) = _state(DEST_CHAIN_SELECTOR, false);
    assertEq(inbound.capacity, 2000e18); // Config unchanged...
    assertEq(inbound.tokens, 2000e18); // ...but the bucket is refilled.
  }

  function test_setOutboundRateLimit_EmitsEvent() public {
    vm.expectEmit(address(s_helper));
    emit TokenPoolRateLimitHelper.DirectionalRateLimitUpdated(
      DEST_CHAIN_SELECTOR,
      false,
      TokenPoolRateLimitHelper.Direction.Outbound,
      _helperCfg(true, 1000e18, 10e18),
      _helperCfg(true, 3000e18, 30e18),
      _helperCfg(true, 2000e18, 20e18),
      _helperCfg(true, 2000e18, 20e18)
    );
    s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 3000e18, 30e18));
  }

  function testFuzz_setOutboundRateLimit_OnlyTargetDirectionChanges(
    bool fastFinality,
    uint128 capacity,
    uint128 rate
  ) public {
    rate = uint128(bound(rate, 0, capacity));
    (, RateLimiter.TokenBucket memory inboundBefore) = _state(DEST_CHAIN_SELECTOR, fastFinality);

    vm.startPrank(s_manager);
    s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, fastFinality, _helperCfg(true, capacity, rate));

    (RateLimiter.TokenBucket memory outboundAfter, RateLimiter.TokenBucket memory inboundAfter) =
      _state(DEST_CHAIN_SELECTOR, fastFinality);
    assertEq(inboundAfter.isEnabled, inboundBefore.isEnabled);
    assertEq(inboundAfter.capacity, inboundBefore.capacity);
    assertEq(inboundAfter.rate, inboundBefore.rate);
    assertTrue(outboundAfter.isEnabled);
    assertEq(outboundAfter.capacity, capacity);
    assertEq(outboundAfter.rate, rate);
  }

  // Reverts

  function test_RevertWhen_CallerIsNotOwnerOrManager() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(TokenPoolRateLimitHelper.Unauthorized.selector, STRANGER));
    s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 1e18, 1e18));
  }

  function test_RevertWhen_HelperIsNotRateLimitAdmin() public {
    s_pool.setDynamicConfig(address(s_sourceRouter), STRANGER, address(0));
    assertFalse(s_helper.isActiveRateLimitAdmin());

    vm.expectRevert(abi.encodeWithSelector(TokenPool.Unauthorized.selector, address(s_helper)));
    s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 1e18, 1e18));
  }

  function test_RevertWhen_ChainNotSupported() public {
    vm.expectRevert(abi.encodeWithSelector(TokenPool.NonExistentChain.selector, uint64(999)));
    s_helper.setOutboundRateLimit(999, false, _helperCfg(true, 1e18, 1e18));
  }

  function test_RevertWhen_InvalidRateLimitRate() public {
    vm.expectRevert(abi.encodeWithSelector(RateLimiter.InvalidRateLimitRate.selector, _cfg(true, 1e18, 2e18)));
    s_helper.setOutboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 1e18, 2e18));
  }
}
