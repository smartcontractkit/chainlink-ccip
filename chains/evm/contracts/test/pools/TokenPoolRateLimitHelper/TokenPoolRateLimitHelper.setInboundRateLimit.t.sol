// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {TokenPoolRateLimitHelper} from "../../../pools/TokenPoolRateLimitHelper.sol";
import {TokenPoolRateLimitHelperSetup} from "./TokenPoolRateLimitHelperSetup.t.sol";

contract TokenPoolRateLimitHelper_setInboundRateLimit is TokenPoolRateLimitHelperSetup {
  function test_setInboundRateLimit_LeavesOutboundConfig() public {
    vm.startPrank(s_manager);
    s_helper.setInboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 50e18, 1e18));

    (RateLimiter.TokenBucket memory outbound, RateLimiter.TokenBucket memory inbound) =
      _state(DEST_CHAIN_SELECTOR, false);
    assertEq(outbound.capacity, 1000e18);
    assertEq(outbound.rate, 10e18);
    assertEq(inbound.capacity, 50e18);
    assertEq(inbound.rate, 1e18);
  }

  function test_setInboundRateLimit_Disable() public {
    s_helper.setInboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(false, 0, 0));

    (RateLimiter.TokenBucket memory outbound, RateLimiter.TokenBucket memory inbound) =
      _state(DEST_CHAIN_SELECTOR, false);
    assertFalse(inbound.isEnabled);
    assertTrue(outbound.isEnabled);
    assertEq(outbound.capacity, 1000e18);
  }

  function testFuzz_setInboundRateLimit_OnlyTargetDirectionChanges(
    bool fastFinality,
    uint128 capacity,
    uint128 rate
  ) public {
    rate = uint128(bound(rate, 0, capacity));
    (RateLimiter.TokenBucket memory outboundBefore,) = _state(DEST_CHAIN_SELECTOR, fastFinality);

    vm.startPrank(s_manager);
    s_helper.setInboundRateLimit(DEST_CHAIN_SELECTOR, fastFinality, _helperCfg(true, capacity, rate));

    (RateLimiter.TokenBucket memory outboundAfter, RateLimiter.TokenBucket memory inboundAfter) =
      _state(DEST_CHAIN_SELECTOR, fastFinality);
    assertEq(outboundAfter.isEnabled, outboundBefore.isEnabled);
    assertEq(outboundAfter.capacity, outboundBefore.capacity);
    assertEq(outboundAfter.rate, outboundBefore.rate);
    assertTrue(inboundAfter.isEnabled);
    assertEq(inboundAfter.capacity, capacity);
    assertEq(inboundAfter.rate, rate);
  }

  function test_RevertWhen_ManagerRevoked() public {
    s_helper.setManager(s_manager, false);

    vm.startPrank(s_manager);
    vm.expectRevert(abi.encodeWithSelector(TokenPoolRateLimitHelper.Unauthorized.selector, s_manager));
    s_helper.setInboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 1e18, 1e18));
  }
}
