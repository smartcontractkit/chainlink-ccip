// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyTokenRemotePoolPolicyUpdates is PoolConfigManagerRolesSetup {
  /// @notice A batch assigns token A to OTHER_POLICY and removes token B's policy assignment. The connection store
  /// emits both changes.
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

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyTokenRemotePoolPolicyUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(PoolConfigManager.applyTokenRemotePoolPolicyUpdates, (new PolicyConfig.TokenPolicyArgs[](0))),
      _getTokenWriteRoles(s_poolConfigManager.REMOTE_POOL_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applyTokenRemotePoolPolicyUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(PoolConfigManager.applyTokenRemotePoolPolicyUpdates, (new PolicyConfig.TokenPolicyArgs[](0))),
      _getTokenWriteRoles(s_poolConfigManager.REMOTE_POOL_ROLE())
    );
  }
}
