// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";

import {RateLimiter} from "../../../../libraries/RateLimiter.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyInlineRateLimitBucketUpdates is PoolConfigManagerRolesSetup {
  function setUp() public virtual override {
    super.setUp();
    _setInlineRateLimitConfig(s_tokenB, OTHER_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
  }

  /// @notice Updates one bucket config in an inline record.
  function test_applyInlineRateLimitBucketUpdates() public {
    RateLimiter.Config memory config = _getRateLimitConfig(2 * CAPACITY);

    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.InlineRateLimitBucketSet(
      s_tokenB, OTHER_CHAIN_SELECTOR, IPoolV2.MessageDirection.Inbound, false, config
    );
    _updateInlineRateLimitBucket(s_tokenB, OTHER_CHAIN_SELECTOR, false, IPoolV2.MessageDirection.Inbound, config);

    assertEq(
      abi.encode(config), abi.encode(s_poolConfigManager.getInlineRateLimitConfigs(s_tokenB)[0].rateLimitConfig.inbound)
    );
  }

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyInlineRateLimitBucketUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyInlineRateLimitBucketUpdates, (new SecurityPolicyStore.InlineRateLimitBucketArgs[](0))
      ),
      _getTokenWriteRoles(s_poolConfigManager.RATE_LIMIT_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applyInlineRateLimitBucketUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyInlineRateLimitBucketUpdates, (new SecurityPolicyStore.InlineRateLimitBucketArgs[](0))
      ),
      _getTokenWriteRoles(s_poolConfigManager.RATE_LIMIT_ROLE())
    );
  }

  function _getInlineRateLimitBucketArgs(
    address localToken,
    uint64 remoteChainSelector,
    bool fastFinality,
    IPoolV2.MessageDirection direction,
    RateLimiter.Config memory config
  ) internal pure returns (SecurityPolicyStore.InlineRateLimitBucketArgs memory) {
    return SecurityPolicyStore.InlineRateLimitBucketArgs({
      localToken: localToken,
      remoteChainSelector: remoteChainSelector,
      direction: direction,
      fastFinality: fastFinality,
      config: config
    });
  }

  /// @dev Updates one bucket config in an inline record.
  function _updateInlineRateLimitBucket(
    address localToken,
    uint64 remoteChainSelector,
    bool fastFinality,
    IPoolV2.MessageDirection direction,
    RateLimiter.Config memory config
  ) internal {
    SecurityPolicyStore.InlineRateLimitBucketArgs[] memory bucketConfigs =
      new SecurityPolicyStore.InlineRateLimitBucketArgs[](1);
    bucketConfigs[0] = _getInlineRateLimitBucketArgs(localToken, remoteChainSelector, fastFinality, direction, config);
    s_poolConfigManager.applyInlineRateLimitBucketUpdates(bucketConfigs);
  }
}
