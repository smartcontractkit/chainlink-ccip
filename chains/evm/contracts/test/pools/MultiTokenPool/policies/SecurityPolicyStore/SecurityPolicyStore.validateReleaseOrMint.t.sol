// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {RateLimiter} from "../../../../../libraries/RateLimiter.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

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

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.InboundRateLimitConsumed(DEST_CHAIN_SELECTOR, s_tokenA, 100e18);
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

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.FastFinalityInboundRateLimitConsumed(DEST_CHAIN_SELECTOR, s_tokenA, 100e18);
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

  // Policy records

  /// @notice An inbound transfer with no inline record uses the chain record of its policy, and consumes the inbound
  /// bucket of its token for the transfer chain.
  function test_validateReleaseOrMint_UsesPolicyChainRecord() public {
    _setRateLimitPolicyConfigPerSlot(
      DEFAULT_POLICY,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, new address[](0), _ccvsOf(s_ccv1, s_ccv2), s_hook);

    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 100e18);

    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(CAPACITY, state.outbound.tokens);
    assertEq(CAPACITY + 1 - 100e18, state.inbound.tokens);
    assertEq(CAPACITY + 2, state.fastFinalityOutbound.tokens);
    assertEq(CAPACITY + 3, state.fastFinalityInbound.tokens);
  }

  /// @notice An inbound transfer on a chain with no chain record of its policy uses the policy default record.
  function test_validateReleaseOrMint_UsesPolicyDefaultRecord() public {
    _setRateLimitPolicyConfigPerSlot(
      DEFAULT_POLICY,
      PolicyConfig.DEFAULT_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, new address[](0), _ccvsOf(s_ccv1, s_ccv2), s_hook);

    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 100e18);

    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(CAPACITY, state.outbound.tokens);
    assertEq(CAPACITY + 1 - 100e18, state.inbound.tokens);
  }

  /// @notice An FTF inbound transfer that uses a policy consumes the FTF bucket when the policy record enables it.
  function test_validateReleaseOrMint_FastFinalityPolicyEnabled() public {
    _setRateLimitPolicyConfigPerSlot(
      DEFAULT_POLICY,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, new address[](0), _ccvsOf(s_ccv1, s_ccv2), s_hook);

    vm.prank(s_poolConfigManager);
    s_store.validateReleaseOrMint(s_tokenA, DEST_CHAIN_SELECTOR, 100e18, _getFastFinalityConfig());

    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(CAPACITY + 1, state.inbound.tokens);
    assertEq(CAPACITY + 3 - 100e18, state.fastFinalityInbound.tokens);
  }

  /// @notice An FTF inbound transfer that uses a policy falls back to the default bucket when the policy disables FTF.
  function test_validateReleaseOrMint_FastFinalityPolicyDisabled() public {
    RateLimiter.Config memory disabled;
    _setRateLimitPolicyConfigPerSlot(
      DEFAULT_POLICY,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      disabled,
      disabled
    );
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, new address[](0), _ccvsOf(s_ccv1, s_ccv2), s_hook);

    vm.prank(s_poolConfigManager);
    s_store.validateReleaseOrMint(s_tokenA, DEST_CHAIN_SELECTOR, 100e18, _getFastFinalityConfig());

    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(CAPACITY, state.outbound.tokens);
    assertEq(CAPACITY + 1 - 100e18, state.inbound.tokens);
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

  /// @notice Only an authorized pool config manager consumes the rate limit and reads the hook.
  function test_validateReleaseOrMint_RevertWhen_UnauthorizedCaller() public {
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, address(this)));
    s_store.validateReleaseOrMint(s_tokenA, DEST_CHAIN_SELECTOR, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  // Helpers

  function _consume(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount
  ) internal {
    vm.prank(s_poolConfigManager);
    s_store.validateReleaseOrMint(localToken, remoteChainSelector, amount, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }
}
