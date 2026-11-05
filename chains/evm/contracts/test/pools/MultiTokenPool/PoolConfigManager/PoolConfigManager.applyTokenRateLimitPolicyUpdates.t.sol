// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {RateLimiter} from "../../../../libraries/RateLimiter.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyTokenRateLimitPolicyUpdates is PoolConfigManagerRolesSetup {
  /// @notice A batch assigns token A to OTHER_POLICY and removes token B's policy assignment. The security store
  /// emits both changes.
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

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyTokenRateLimitPolicyUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(PoolConfigManager.applyTokenRateLimitPolicyUpdates, (new PolicyConfig.TokenPolicyArgs[](0))),
      _getTokenWriteRoles(s_poolConfigManager.RATE_LIMIT_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applyTokenRateLimitPolicyUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(PoolConfigManager.applyTokenRateLimitPolicyUpdates, (new PolicyConfig.TokenPolicyArgs[](0))),
      _getTokenWriteRoles(s_poolConfigManager.RATE_LIMIT_ROLE())
    );
  }
}
