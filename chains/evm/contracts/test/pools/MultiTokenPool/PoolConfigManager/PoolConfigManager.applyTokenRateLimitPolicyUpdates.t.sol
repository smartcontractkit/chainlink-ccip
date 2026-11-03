// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {RateLimiter} from "../../../../libraries/RateLimiter.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyTokenRateLimitPolicyUpdates is PoolConfigManagerRolesSetup {
  /// @notice CONFIG_ROLE sets the rate limit policy of several tokens in one call. Token A moves to the OTHER policy,
  /// and token B leaves its policy. The security store applies the call and emits the events.
  function test_applyTokenRateLimitPolicyUpdates() public {
    RateLimiter.Config memory disabledConfig;
    _setRateLimitPolicyConfig(OTHER_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, disabledConfig);
    PolicyConfig.TokenPolicyArgs[] memory tokenPolicies = new PolicyConfig.TokenPolicyArgs[](2);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: s_tokenA, policyName: OTHER_POLICY});
    tokenPolicies[1] = PolicyConfig.TokenPolicyArgs({localToken: s_tokenB, policyName: bytes32(0)});

    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.TokenRateLimitPolicySet(s_tokenA, DEFAULT_POLICY, OTHER_POLICY);
    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.TokenRateLimitPolicySet(s_tokenB, DEFAULT_POLICY, bytes32(0));
    s_poolConfigManager.applyTokenRateLimitPolicyUpdates(tokenPolicies);

    assertEq(OTHER_POLICY, _getTokenConfig(s_tokenA).rateLimitPolicyName);
    assertEq(bytes32(0), _getTokenConfig(s_tokenB).rateLimitPolicyName);
  }

  /// @notice CONFIG_ROLE and TOKEN_MANAGER_ROLE each pass alone.
  function test_applyTokenRateLimitPolicyUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(PoolConfigManager.applyTokenRateLimitPolicyUpdates, (new PolicyConfig.TokenPolicyArgs[](0))),
      _getTokenWriteRoles()
    );
  }

  // Reverts

  /// @notice Each other PCM role and a caller with no role revert. The error names the second role of the check.
  function test_applyTokenRateLimitPolicyUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(PoolConfigManager.applyTokenRateLimitPolicyUpdates, (new PolicyConfig.TokenPolicyArgs[](0))),
      _getTokenWriteRoles()
    );
  }
}
