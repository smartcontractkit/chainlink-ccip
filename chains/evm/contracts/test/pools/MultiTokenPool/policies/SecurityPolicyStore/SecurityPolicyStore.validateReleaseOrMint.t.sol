// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {RateLimiter} from "../../../../../libraries/RateLimiter.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

contract SecurityPolicyStore_validateReleaseOrMint is SecurityPolicyStoreSetup {
  /// @notice A default finality transfer consumes the inbound default finality bucket only, and returns the hook of the
  /// security record. Each of the four configs has its own capacity, so a transfer that consumes the wrong direction
  /// or finality changes the wrong bucket.
  function test_validateReleaseOrMint_ConsumesInboundBucketOnly() public {
    _setInlineRateLimitConfigPerSlot(
      s_tokenA,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, new address[](0), _ccvsOf(s_ccv1, s_ccv2), s_hook);

    vm.prank(s_poolConfigManager);
    address hook =
      s_store.validateReleaseOrMint(s_tokenA, DEST_CHAIN_SELECTOR, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    assertEq(s_hook, hook);
    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(CAPACITY, state.outbound.tokens);
    assertEq(CAPACITY + 1 - 100e18, state.inbound.tokens);
    assertEq(CAPACITY + 2, state.fastFinalityOutbound.tokens);
    assertEq(CAPACITY + 3, state.fastFinalityInbound.tokens);
  }

  /// @notice A fast finality transfer consumes the inbound FTF bucket only.
  function test_validateReleaseOrMint_FastFinalityConsumesFastFinalityBucket() public {
    _setInlineRateLimitConfigPerSlot(
      s_tokenA,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, new address[](0), _ccvsOf(s_ccv1, s_ccv2), s_hook);

    vm.prank(s_poolConfigManager);
    s_store.validateReleaseOrMint(s_tokenA, DEST_CHAIN_SELECTOR, 100e18, _getFastFinalityConfig());

    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(CAPACITY, state.outbound.tokens);
    assertEq(CAPACITY + 1, state.inbound.tokens);
    assertEq(CAPACITY + 2, state.fastFinalityOutbound.tokens);
    assertEq(CAPACITY + 3 - 100e18, state.fastFinalityInbound.tokens);
  }

  /// @notice A record with no FTF config makes an FTF transfer use the default finality bucket, the same as TokenPool.
  function test_validateReleaseOrMint_FastFinalityFallsBackToDefaultBucket() public {
    RateLimiter.Config memory disabled;
    _setInlineRateLimitConfigPerSlot(
      s_tokenA,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      disabled,
      disabled
    );
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, new address[](0), _ccvsOf(s_ccv1, s_ccv2), s_hook);

    vm.prank(s_poolConfigManager);
    s_store.validateReleaseOrMint(s_tokenA, DEST_CHAIN_SELECTOR, 100e18, _getFastFinalityConfig());

    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(CAPACITY, state.outbound.tokens);
    assertEq(CAPACITY + 1 - 100e18, state.inbound.tokens);
  }

  /// @notice The event names the direction and finality that the transfer consumed.
  function test_validateReleaseOrMint_EmitsConsumedEvent() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, new address[](0), _ccvsOf(s_ccv1, s_ccv2), s_hook);

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.InboundRateLimitConsumed(DEST_CHAIN_SELECTOR, s_tokenA, 100e18);
    vm.prank(s_poolConfigManager);
    s_store.validateReleaseOrMint(s_tokenA, DEST_CHAIN_SELECTOR, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  // Reverts

  /// @notice validateReleaseOrMint reads the rate limit record first, so a missing rate limit record reverts before the
  /// security record is read.
  function test_validateReleaseOrMint_RevertWhen_RateLimitRecordNotFound() public {
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, new address[](0), _ccvsOf(s_ccv1, s_ccv2), s_hook);

    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.RateLimitRecordNotFound.selector, s_tokenA, DEST_CHAIN_SELECTOR)
    );
    vm.prank(s_poolConfigManager);
    s_store.validateReleaseOrMint(s_tokenA, DEST_CHAIN_SELECTOR, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_validateReleaseOrMint_RevertWhen_SecurityRecordNotFound() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));

    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.SecurityRecordNotFound.selector, s_tokenA, DEST_CHAIN_SELECTOR)
    );
    vm.prank(s_poolConfigManager);
    s_store.validateReleaseOrMint(s_tokenA, DEST_CHAIN_SELECTOR, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_validateReleaseOrMint_RevertWhen_TokenMaxCapacityExceeded() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, new address[](0), _ccvsOf(s_ccv1, s_ccv2), s_hook);

    vm.expectRevert(
      abi.encodeWithSelector(RateLimiter.TokenMaxCapacityExceeded.selector, CAPACITY, CAPACITY + 1, s_tokenA)
    );
    vm.prank(s_poolConfigManager);
    s_store.validateReleaseOrMint(s_tokenA, DEST_CHAIN_SELECTOR, CAPACITY + 1, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }
}
