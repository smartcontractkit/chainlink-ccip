// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyTokenFeePolicyUpdates is PoolConfigManagerRolesSetup {
  /// @notice CONFIG_ROLE sets the fee policy of several tokens in one call. Token A moves to the OTHER policy, and
  /// token B leaves its policy. The fee store applies the call and emits the events.
  function test_applyTokenFeePolicyUpdates() public {
    _setFeePolicyConfig(OTHER_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _getFeeConfig(100, 10));
    PolicyConfig.TokenPolicyArgs[] memory tokenPolicies = new PolicyConfig.TokenPolicyArgs[](2);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: s_tokenA, policyName: OTHER_POLICY});
    tokenPolicies[1] = PolicyConfig.TokenPolicyArgs({localToken: s_tokenB, policyName: bytes32(0)});

    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.TokenFeePolicySet(s_tokenA, DEFAULT_POLICY, OTHER_POLICY);
    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.TokenFeePolicySet(s_tokenB, DEFAULT_POLICY, bytes32(0));
    s_poolConfigManager.applyTokenFeePolicyUpdates(tokenPolicies);

    assertEq(OTHER_POLICY, _getTokenConfig(s_tokenA).feePolicyName);
    assertEq(bytes32(0), _getTokenConfig(s_tokenB).feePolicyName);
  }

  /// @notice CONFIG_ROLE and TOKEN_MANAGER_ROLE each pass alone.
  function test_applyTokenFeePolicyUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(PoolConfigManager.applyTokenFeePolicyUpdates, (new PolicyConfig.TokenPolicyArgs[](0))),
      _getTokenWriteRoles()
    );
  }

  // Reverts

  /// @notice Each other PCM role and a caller with no role revert. The error names the second role of the check.
  function test_applyTokenFeePolicyUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(PoolConfigManager.applyTokenFeePolicyUpdates, (new PolicyConfig.TokenPolicyArgs[](0))),
      _getTokenWriteRoles()
    );
  }
}
