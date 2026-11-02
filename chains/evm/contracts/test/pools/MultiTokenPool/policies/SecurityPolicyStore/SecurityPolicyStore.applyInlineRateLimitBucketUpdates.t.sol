// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../../../interfaces/IPoolV2.sol";

import {RateLimiter} from "../../../../../libraries/RateLimiter.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract SecurityPolicyStore_applyInlineRateLimitBucketUpdates is SecurityPolicyStoreSetup {
  /// @notice Each direction and finality combination writes its own bucket only, and names it in the event.
  function test_applyInlineRateLimitBucketUpdates_OutboundDefault() public {
    _assertOneBucketUpdated(IPoolV2.MessageDirection.Outbound, false);
  }

  function test_applyInlineRateLimitBucketUpdates_OutboundFastFinality() public {
    _assertOneBucketUpdated(IPoolV2.MessageDirection.Outbound, true);
  }

  function test_applyInlineRateLimitBucketUpdates_InboundDefault() public {
    _assertOneBucketUpdated(IPoolV2.MessageDirection.Inbound, false);
  }

  function test_applyInlineRateLimitBucketUpdates_InboundFastFinality() public {
    _assertOneBucketUpdated(IPoolV2.MessageDirection.Inbound, true);
  }

  /// @notice A set of the same config keeps the state of the bucket.
  function test_applyInlineRateLimitBucketUpdates_SameConfigKeepsBucketState() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 400e18);

    _applyBucketUpdate(IPoolV2.MessageDirection.Outbound, false, _getRateLimitConfig(CAPACITY));

    assertEq(CAPACITY - 400e18, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);
  }

  // Reverts

  /// @notice The bucket update needs the record to exist.
  function test_applyInlineRateLimitBucketUpdates_RevertWhen_RateLimitRecordNotFound() public {
    SecurityPolicyStore.InlineRateLimitBucketArgs[] memory bucketConfigs =
      new SecurityPolicyStore.InlineRateLimitBucketArgs[](1);
    bucketConfigs[0] = SecurityPolicyStore.InlineRateLimitBucketArgs({
      localToken: s_tokenA,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      direction: IPoolV2.MessageDirection.Outbound,
      fastFinality: false,
      config: _getRateLimitConfig(CAPACITY)
    });

    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.RateLimitRecordNotFound.selector, s_tokenA, DEST_CHAIN_SELECTOR)
    );
    vm.prank(s_poolConfigManager);
    s_store.applyInlineRateLimitBucketUpdates(bucketConfigs);
  }

  /// @notice Each config follows the checks of RateLimiter, and the revert names the bucket.
  function test_applyInlineRateLimitBucketUpdates_RevertWhen_InvalidRateLimitConfig() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));

    RateLimiter.Config memory invalidConfig = RateLimiter.Config({isEnabled: true, capacity: 10, rate: 11});
    SecurityPolicyStore.InlineRateLimitBucketArgs[] memory bucketConfigs =
      new SecurityPolicyStore.InlineRateLimitBucketArgs[](1);
    bucketConfigs[0] = SecurityPolicyStore.InlineRateLimitBucketArgs({
      localToken: s_tokenA,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      direction: IPoolV2.MessageDirection.Outbound,
      fastFinality: false,
      config: invalidConfig
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        SecurityPolicyStore.InvalidInlineRateLimitBucketConfig.selector,
        s_tokenA,
        DEST_CHAIN_SELECTOR,
        IPoolV2.MessageDirection.Outbound,
        false,
        invalidConfig
      )
    );
    vm.prank(s_poolConfigManager);
    s_store.applyInlineRateLimitBucketUpdates(bucketConfigs);
  }

  /// @notice Only an authorized pool config manager writes to the store.
  function test_applyInlineRateLimitBucketUpdates_RevertWhen_UnauthorizedCaller() public {
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, address(this)));
    s_store.applyInlineRateLimitBucketUpdates(new SecurityPolicyStore.InlineRateLimitBucketArgs[](0));
  }

  // Helpers

  /// @dev Sets four distinct configs, updates one bucket, and checks that the other three kept their config.
  function _assertOneBucketUpdated(
    IPoolV2.MessageDirection direction,
    bool fastFinality
  ) internal {
    _setInlineRateLimitConfigPerSlot(
      s_tokenA,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );

    RateLimiter.Config memory newConfig = _getRateLimitConfig(7 * CAPACITY);
    _applyBucketUpdate(direction, fastFinality, newConfig);

    RateLimiter.Config memory outbound = _getRateLimitConfig(CAPACITY);
    RateLimiter.Config memory inbound = _getRateLimitConfig(CAPACITY + 1);
    RateLimiter.Config memory fastFinalityOutbound = _getRateLimitConfig(CAPACITY + 2);
    RateLimiter.Config memory fastFinalityInbound = _getRateLimitConfig(CAPACITY + 3);
    if (direction == IPoolV2.MessageDirection.Outbound) {
      if (fastFinality) {
        fastFinalityOutbound = newConfig;
      } else {
        outbound = newConfig;
      }
    } else if (fastFinality) {
      fastFinalityInbound = newConfig;
    } else {
      inbound = newConfig;
    }

    SecurityPolicyStore.RateLimitConfig memory expected = SecurityPolicyStore.RateLimitConfig({
      outbound: outbound,
      inbound: inbound,
      fastFinalityOutbound: fastFinalityOutbound,
      fastFinalityInbound: fastFinalityInbound
    });
    assertEq(abi.encode(expected), abi.encode(_getInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR).rateLimitConfig));
  }

  function _applyBucketUpdate(
    IPoolV2.MessageDirection direction,
    bool fastFinality,
    RateLimiter.Config memory config
  ) internal {
    SecurityPolicyStore.InlineRateLimitBucketArgs[] memory bucketConfigs =
      new SecurityPolicyStore.InlineRateLimitBucketArgs[](1);
    bucketConfigs[0] = SecurityPolicyStore.InlineRateLimitBucketArgs({
      localToken: s_tokenA,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      direction: direction,
      fastFinality: fastFinality,
      config: config
    });

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.InlineRateLimitBucketSet(s_tokenA, DEST_CHAIN_SELECTOR, direction, fastFinality, config);
    vm.prank(s_poolConfigManager);
    s_store.applyInlineRateLimitBucketUpdates(bucketConfigs);
  }

  function _consume(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount
  ) internal {
    _setInlineSecurityConfig(localToken, remoteChainSelector, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);
    vm.prank(s_poolConfigManager);
    s_store.validateLockOrBurn(localToken, remoteChainSelector, amount, bytes4(0));
  }
}
