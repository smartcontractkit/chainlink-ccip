// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {RateLimiter} from "../../../../libraries/RateLimiter.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyRateLimitPolicyBucketUpdates is PoolConfigManagerRolesSetup {
  /// @notice Updates one bucket config in a policy record.
  function test_applyRateLimitPolicyBucketUpdates() public {
    RateLimiter.Config memory config = _getRateLimitConfig(CAPACITY);

    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.RateLimitPolicyBucketSet(
      DEFAULT_POLICY,
      PolicyConfig.DEFAULT_CHAIN_SELECTOR,
      DEFAULT_TOKEN_DECIMALS,
      IPoolV2.MessageDirection.Outbound,
      false,
      config
    );
    _updateRateLimitPolicyBucket(
      DEFAULT_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, false, IPoolV2.MessageDirection.Outbound, config
    );

    (,, SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory policyConfigs) =
      s_poolConfigManager.getRateLimitPolicy(DEFAULT_POLICY);
    assertEq(abi.encode(config), abi.encode(policyConfigs[0].rateLimitConfig.outbound));
  }

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyRateLimitPolicyBucketUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyRateLimitPolicyBucketUpdates, (new SecurityPolicyStore.RateLimitPolicyBucketArgs[](0))
      ),
      _getPolicyWriteRoles(s_poolConfigManager.RATE_LIMIT_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applyRateLimitPolicyBucketUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyRateLimitPolicyBucketUpdates, (new SecurityPolicyStore.RateLimitPolicyBucketArgs[](0))
      ),
      _getPolicyWriteRoles(s_poolConfigManager.RATE_LIMIT_ROLE())
    );
  }

  function _getRateLimitPolicyBucketArgs(
    bytes32 policyName,
    uint64 remoteChainSelector,
    bool fastFinality,
    IPoolV2.MessageDirection direction,
    RateLimiter.Config memory config
  ) internal pure returns (SecurityPolicyStore.RateLimitPolicyBucketArgs memory) {
    return SecurityPolicyStore.RateLimitPolicyBucketArgs({
      policyName: policyName,
      remoteChainSelector: remoteChainSelector,
      direction: direction,
      fastFinality: fastFinality,
      config: config
    });
  }

  /// @dev Updates one bucket config in a policy record.
  function _updateRateLimitPolicyBucket(
    bytes32 policyName,
    uint64 remoteChainSelector,
    bool fastFinality,
    IPoolV2.MessageDirection direction,
    RateLimiter.Config memory config
  ) internal {
    SecurityPolicyStore.RateLimitPolicyBucketArgs[] memory bucketConfigs =
      new SecurityPolicyStore.RateLimitPolicyBucketArgs[](1);
    bucketConfigs[0] = _getRateLimitPolicyBucketArgs(policyName, remoteChainSelector, fastFinality, direction, config);
    s_poolConfigManager.applyRateLimitPolicyBucketUpdates(bucketConfigs);
  }
}
