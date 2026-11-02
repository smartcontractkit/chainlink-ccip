// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {RateLimiter} from "../../../../../libraries/RateLimiter.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

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

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.OutboundRateLimitConsumed(DEST_CHAIN_SELECTOR, s_tokenA, 100e18);
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

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.FastFinalityOutboundRateLimitConsumed(DEST_CHAIN_SELECTOR, s_tokenA, 100e18);
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

  // Policy records

  /// @notice A transfer with no inline record uses the chain record of its policy, and consumes the bucket of its token
  /// for the transfer chain.
  function test_validateLockOrBurn_UsesPolicyChainRecord() public {
    _setRateLimitPolicyConfigPerSlot(
      DEFAULT_POLICY,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);

    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 100e18);

    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(CAPACITY - 100e18, state.outbound.tokens);
    assertEq(CAPACITY + 1, state.inbound.tokens);
    assertEq(CAPACITY + 2, state.fastFinalityOutbound.tokens);
    assertEq(CAPACITY + 3, state.fastFinalityInbound.tokens);
  }

  /// @notice A transfer on a chain with no chain record of its policy uses the policy default record.
  function test_validateLockOrBurn_UsesPolicyDefaultRecord() public {
    _setRateLimitPolicyConfigPerSlot(
      DEFAULT_POLICY,
      PolicyConfig.DEFAULT_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);

    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 100e18);

    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(CAPACITY - 100e18, state.outbound.tokens);
    assertEq(CAPACITY + 1, state.inbound.tokens);
  }

  /// @notice An FTF transfer that uses a policy consumes the FTF bucket when the policy record enables it.
  function test_validateLockOrBurn_FastFinalityPolicyEnabled() public {
    _setRateLimitPolicyConfigPerSlot(
      DEFAULT_POLICY,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);

    vm.prank(s_poolConfigManager);
    s_store.validateLockOrBurn(s_tokenA, DEST_CHAIN_SELECTOR, 100e18, _getFastFinalityConfig());

    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(CAPACITY, state.outbound.tokens);
    assertEq(CAPACITY + 2 - 100e18, state.fastFinalityOutbound.tokens);
  }

  /// @notice An FTF transfer that uses a policy falls back to the default bucket when the policy record disables FTF.
  function test_validateLockOrBurn_FastFinalityPolicyDisabled() public {
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
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);

    vm.prank(s_poolConfigManager);
    s_store.validateLockOrBurn(s_tokenA, DEST_CHAIN_SELECTOR, 100e18, _getFastFinalityConfig());

    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(CAPACITY - 100e18, state.outbound.tokens);
    assertEq(CAPACITY + 1, state.inbound.tokens);
  }

  /// @notice An inline record takes priority over the chain record and the default record of the policy. Removing the
  /// inline record moves the token to the chain record, and removing the chain record moves it to the default record.
  function test_validateLockOrBurn_InlineOverridesPolicyRecords() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(2 * CAPACITY));
    _setRateLimitPolicyConfig(DEFAULT_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _getRateLimitConfig(3 * CAPACITY));
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);

    // The inline record serves the transfer.
    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 100e18);
    assertEq(CAPACITY - 100e18, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);

    // After the inline removal, the chain record of the policy serves the transfer.
    _removeInlineRateLimitRecord(s_tokenA, DEST_CHAIN_SELECTOR);
    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 100e18);
    assertEq(2 * CAPACITY - 100e18, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);

    // After the chain record removal, the default record of the policy serves the transfer.
    _removeRateLimitPolicyRecord(DEFAULT_POLICY, DEST_CHAIN_SELECTOR);
    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 100e18);
    assertEq(3 * CAPACITY - 100e18, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);
  }

  /// @notice Two tokens that share a policy and two remote chains have separate buckets. A transfer changes only the
  /// bucket of its token and chain.
  function test_validateLockOrBurn_PolicySharedByTwoTokensTwoChains() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setRateLimitPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, _getRateLimitConfig(2 * CAPACITY));
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setTokenRateLimitPolicy(s_tokenB, DEFAULT_POLICY);
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);
    _setInlineSecurityConfig(s_tokenA, OTHER_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);
    _setInlineSecurityConfig(s_tokenB, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);

    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 100e18);

    assertEq(CAPACITY - 100e18, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);
    assertEq(2 * CAPACITY, _getRateLimiterState(s_tokenA, OTHER_CHAIN_SELECTOR).outbound.tokens);
    assertEq(CAPACITY, _getRateLimiterState(s_tokenB, DEST_CHAIN_SELECTOR).outbound.tokens);
  }

  /// @notice A policy config change refills the bucket of a token on its next use, while the same config keeps its
  /// state. The state getter follows the same rule and applies the rate refill of the elapsed time.
  function test_validateLockOrBurn_PolicyChangeAndTimeRefill() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);

    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 400e18);

    // The same config at the same timestamp keeps the consumed amount.
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 100e18);
    assertEq(CAPACITY - 500e18, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);

    // Another config refills the bucket to the new capacity, also in the state getter before the next use.
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(2 * CAPACITY));
    assertEq(2 * CAPACITY, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);

    // The rate refills the bucket with the elapsed time. The rate is one token unit per second.
    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 100e18);
    vm.warp(block.timestamp + 10);
    assertEq(2 * CAPACITY - 100e18 + 10, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);
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

  /// @notice Only an authorized pool config manager consumes the rate limit and reads the hook.
  function test_validateLockOrBurn_RevertWhen_UnauthorizedCaller() public {
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, address(this)));
    s_store.validateLockOrBurn(s_tokenA, DEST_CHAIN_SELECTOR, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  // Helpers

  function _consume(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount
  ) internal {
    vm.prank(s_poolConfigManager);
    s_store.validateLockOrBurn(localToken, remoteChainSelector, amount, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function _removeInlineRateLimitRecord(
    address localToken,
    uint64 remoteChainSelector
  ) internal {
    vm.prank(s_poolConfigManager);
    s_store.applyInlineRateLimitConfigUpdates(
      _lanesOf(localToken, remoteChainSelector), new SecurityPolicyStore.InlineRateLimitConfigArgs[](0)
    );
  }

  function _removeRateLimitPolicyRecord(
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal {
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove = new PolicyConfig.PolicyRecordKey[](1);
    recordsToRemove[0] =
      PolicyConfig.PolicyRecordKey({policyName: policyName, remoteChainSelector: remoteChainSelector});
    vm.prank(s_poolConfigManager);
    s_store.applyRateLimitPolicyConfigUpdates(
      recordsToRemove, new bytes32[](0), new SecurityPolicyStore.RateLimitPolicyConfigArgs[](0)
    );
  }
}
