// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {RateLimiter} from "../../../../../libraries/RateLimiter.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

contract SecurityPolicyStore_validateLockOrBurn is SecurityPolicyStoreSetup {
  /// @notice A default finality transfer consumes the outbound default finality bucket only, and returns the hook of
  /// the security record. Each of the four configs has its own capacity, so a transfer that consumes the wrong
  /// direction or finality changes the wrong bucket.
  function test_validateLockOrBurn_ConsumesOutboundBucketOnly() public {
    _setInlineRateLimitConfigPerSlot(
      s_tokenA,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);

    vm.prank(s_poolConfigManager);
    address hook =
      s_store.validateLockOrBurn(s_tokenA, DEST_CHAIN_SELECTOR, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    assertEq(s_hook, hook);
    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(CAPACITY - 100e18, state.outbound.tokens);
    assertEq(CAPACITY + 1, state.inbound.tokens);
    assertEq(CAPACITY + 2, state.fastFinalityOutbound.tokens);
    assertEq(CAPACITY + 3, state.fastFinalityInbound.tokens);
  }

  /// @notice A fast finality transfer consumes the outbound FTF bucket only.
  function test_validateLockOrBurn_FastFinalityConsumesFastFinalityBucket() public {
    _setInlineRateLimitConfigPerSlot(
      s_tokenA,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);

    vm.prank(s_poolConfigManager);
    s_store.validateLockOrBurn(s_tokenA, DEST_CHAIN_SELECTOR, 100e18, _getFastFinalityConfig());

    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(CAPACITY, state.outbound.tokens);
    assertEq(CAPACITY + 1, state.inbound.tokens);
    assertEq(CAPACITY + 2 - 100e18, state.fastFinalityOutbound.tokens);
    assertEq(CAPACITY + 3, state.fastFinalityInbound.tokens);
  }

  /// @notice A record with no FTF config makes an FTF transfer use the default finality bucket, the same as TokenPool.
  function test_validateLockOrBurn_FastFinalityFallsBackToDefaultBucket() public {
    RateLimiter.Config memory disabled;
    _setInlineRateLimitConfigPerSlot(
      s_tokenA,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      disabled,
      disabled
    );
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);

    vm.prank(s_poolConfigManager);
    s_store.validateLockOrBurn(s_tokenA, DEST_CHAIN_SELECTOR, 100e18, _getFastFinalityConfig());

    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(CAPACITY - 100e18, state.outbound.tokens);
    assertEq(CAPACITY + 1, state.inbound.tokens);
  }

  /// @notice The event names the direction and finality that the transfer consumed.
  function test_validateLockOrBurn_EmitsConsumedEvent() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.OutboundRateLimitConsumed(DEST_CHAIN_SELECTOR, s_tokenA, 100e18);
    vm.prank(s_poolConfigManager);
    s_store.validateLockOrBurn(s_tokenA, DEST_CHAIN_SELECTOR, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  // Reverts

  /// @notice validateLockOrBurn reads the security record first, so a missing security record reverts before the rate
  /// limit is touched.
  function test_validateLockOrBurn_RevertWhen_SecurityRecordNotFound() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));

    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.SecurityRecordNotFound.selector, s_tokenA, DEST_CHAIN_SELECTOR)
    );
    vm.prank(s_poolConfigManager);
    s_store.validateLockOrBurn(s_tokenA, DEST_CHAIN_SELECTOR, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_validateLockOrBurn_RevertWhen_RateLimitRecordNotFound() public {
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);

    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.RateLimitRecordNotFound.selector, s_tokenA, DEST_CHAIN_SELECTOR)
    );
    vm.prank(s_poolConfigManager);
    s_store.validateLockOrBurn(s_tokenA, DEST_CHAIN_SELECTOR, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_validateLockOrBurn_RevertWhen_TokenMaxCapacityExceeded() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);

    vm.expectRevert(
      abi.encodeWithSelector(RateLimiter.TokenMaxCapacityExceeded.selector, CAPACITY, CAPACITY + 1, s_tokenA)
    );
    vm.prank(s_poolConfigManager);
    s_store.validateLockOrBurn(s_tokenA, DEST_CHAIN_SELECTOR, CAPACITY + 1, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }
}
