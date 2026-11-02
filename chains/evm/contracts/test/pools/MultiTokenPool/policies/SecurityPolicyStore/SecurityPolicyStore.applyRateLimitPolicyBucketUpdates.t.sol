// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../../../interfaces/IPoolV2.sol";

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {RateLimiter} from "../../../../../libraries/RateLimiter.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract SecurityPolicyStore_applyRateLimitPolicyBucketUpdates is SecurityPolicyStoreSetup {
  /// @notice Each direction and finality combination writes its own bucket only, and names it in the event.
  function test_applyRateLimitPolicyBucketUpdates_OutboundDefault() public {
    _assertOneBucketUpdated(DEST_CHAIN_SELECTOR, IPoolV2.MessageDirection.Outbound, false);
  }

  function test_applyRateLimitPolicyBucketUpdates_OutboundFastFinality() public {
    _assertOneBucketUpdated(DEST_CHAIN_SELECTOR, IPoolV2.MessageDirection.Outbound, true);
  }

  function test_applyRateLimitPolicyBucketUpdates_InboundDefault() public {
    _assertOneBucketUpdated(DEST_CHAIN_SELECTOR, IPoolV2.MessageDirection.Inbound, false);
  }

  function test_applyRateLimitPolicyBucketUpdates_InboundFastFinality() public {
    _assertOneBucketUpdated(DEST_CHAIN_SELECTOR, IPoolV2.MessageDirection.Inbound, true);
  }

  /// @notice Chain selector zero names the policy default record, which the bucket update also accepts.
  function test_applyRateLimitPolicyBucketUpdates_DefaultRecord() public {
    _assertOneBucketUpdated(PolicyConfig.DEFAULT_CHAIN_SELECTOR, IPoolV2.MessageDirection.Outbound, false);
  }

  // Reverts

  /// @notice The bucket update needs the record to exist.
  function test_applyRateLimitPolicyBucketUpdates_RevertWhen_PolicyRecordNotFound() public {
    SecurityPolicyStore.RateLimitPolicyBucketArgs[] memory bucketConfigs =
      new SecurityPolicyStore.RateLimitPolicyBucketArgs[](1);
    bucketConfigs[0] = SecurityPolicyStore.RateLimitPolicyBucketArgs({
      policyName: DEFAULT_POLICY,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      direction: IPoolV2.MessageDirection.Outbound,
      fastFinality: false,
      config: _getRateLimitConfig(CAPACITY)
    });

    vm.expectRevert(
      abi.encodeWithSelector(PolicyConfig.PolicyRecordNotFound.selector, DEFAULT_POLICY, DEST_CHAIN_SELECTOR)
    );
    vm.prank(s_poolConfigManager);
    s_store.applyRateLimitPolicyBucketUpdates(bucketConfigs);
  }

  /// @notice Each config follows the checks of RateLimiter, and the revert names the bucket.
  function test_applyRateLimitPolicyBucketUpdates_RevertWhen_InvalidRateLimitConfig() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));

    RateLimiter.Config memory invalidConfig = RateLimiter.Config({isEnabled: false, capacity: 10, rate: 0});
    SecurityPolicyStore.RateLimitPolicyBucketArgs[] memory bucketConfigs =
      new SecurityPolicyStore.RateLimitPolicyBucketArgs[](1);
    bucketConfigs[0] = SecurityPolicyStore.RateLimitPolicyBucketArgs({
      policyName: DEFAULT_POLICY,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      direction: IPoolV2.MessageDirection.Inbound,
      fastFinality: true,
      config: invalidConfig
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        SecurityPolicyStore.InvalidRateLimitPolicyBucketConfig.selector,
        DEFAULT_POLICY,
        DEST_CHAIN_SELECTOR,
        IPoolV2.MessageDirection.Inbound,
        true,
        invalidConfig
      )
    );
    vm.prank(s_poolConfigManager);
    s_store.applyRateLimitPolicyBucketUpdates(bucketConfigs);
  }

  /// @notice Only an authorized pool config manager writes to the store.
  function test_applyRateLimitPolicyBucketUpdates_RevertWhen_UnauthorizedCaller() public {
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, address(this)));
    s_store.applyRateLimitPolicyBucketUpdates(new SecurityPolicyStore.RateLimitPolicyBucketArgs[](0));
  }

  // Helpers

  /// @dev Sets four distinct configs on a record, updates one bucket, and checks that the other three kept their config.
  function _assertOneBucketUpdated(
    uint64 remoteChainSelector,
    IPoolV2.MessageDirection direction,
    bool fastFinality
  ) internal {
    _setRateLimitPolicyConfigPerSlot(
      DEFAULT_POLICY,
      remoteChainSelector,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );

    RateLimiter.Config memory newConfig = _getRateLimitConfig(7 * CAPACITY);
    SecurityPolicyStore.RateLimitPolicyBucketArgs[] memory bucketConfigs =
      new SecurityPolicyStore.RateLimitPolicyBucketArgs[](1);
    bucketConfigs[0] = SecurityPolicyStore.RateLimitPolicyBucketArgs({
      policyName: DEFAULT_POLICY,
      remoteChainSelector: remoteChainSelector,
      direction: direction,
      fastFinality: fastFinality,
      config: newConfig
    });

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.RateLimitPolicyBucketSet(
      DEFAULT_POLICY, remoteChainSelector, DEFAULT_TOKEN_DECIMALS, direction, fastFinality, newConfig
    );
    vm.prank(s_poolConfigManager);
    s_store.applyRateLimitPolicyBucketUpdates(bucketConfigs);

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
    assertEq(
      abi.encode(expected), abi.encode(_getRateLimitPolicyConfig(DEFAULT_POLICY, remoteChainSelector).rateLimitConfig)
    );
  }

  function _getRateLimitPolicyConfig(
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal view returns (SecurityPolicyStore.RateLimitPolicyConfigArgs memory) {
    (,, SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory records) = s_store.getRateLimitPolicy(policyName);
    for (uint256 i = 0; i < records.length; ++i) {
      if (records[i].remoteChainSelector == remoteChainSelector) return records[i];
    }
    revert("policy record not found");
  }
}
