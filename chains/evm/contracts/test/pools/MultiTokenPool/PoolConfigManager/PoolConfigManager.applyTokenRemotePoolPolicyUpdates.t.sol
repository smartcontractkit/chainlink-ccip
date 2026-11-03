// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyTokenRemotePoolPolicyUpdates is PoolConfigManagerRolesSetup {
  /// @notice CONFIG_ROLE sets the remote pool policy of several tokens in one call. Token A moves to the OTHER policy,
  /// and token B leaves its policy. The connection store applies the call and emits the events.
  function test_applyTokenRemotePoolPolicyUpdates() public {
    _setRemotePoolPolicyConfig(OTHER_POLICY, DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(abi.encode(s_remotePool)));
    PolicyConfig.TokenPolicyArgs[] memory tokenPolicies = new PolicyConfig.TokenPolicyArgs[](2);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: s_tokenA, policyName: OTHER_POLICY});
    tokenPolicies[1] = PolicyConfig.TokenPolicyArgs({localToken: s_tokenB, policyName: bytes32(0)});

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.TokenRemotePoolPolicySet(s_tokenA, DEFAULT_POLICY, OTHER_POLICY);
    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.TokenRemotePoolPolicySet(s_tokenB, DEFAULT_POLICY, bytes32(0));
    s_poolConfigManager.applyTokenRemotePoolPolicyUpdates(tokenPolicies);

    assertEq(OTHER_POLICY, _getTokenConfig(s_tokenA).remotePoolPolicyName);
    assertEq(bytes32(0), _getTokenConfig(s_tokenB).remotePoolPolicyName);
  }

  /// @notice CONFIG_ROLE and TOKEN_MANAGER_ROLE each pass alone.
  function test_applyTokenRemotePoolPolicyUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(PoolConfigManager.applyTokenRemotePoolPolicyUpdates, (new PolicyConfig.TokenPolicyArgs[](0))),
      _getTokenWriteRoles()
    );
  }

  // Reverts

  /// @notice Each other PCM role and a caller with no role revert. The error names the second role of the check.
  function test_applyTokenRemotePoolPolicyUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(PoolConfigManager.applyTokenRemotePoolPolicyUpdates, (new PolicyConfig.TokenPolicyArgs[](0))),
      _getTokenWriteRoles()
    );
  }
}
