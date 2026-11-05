// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyRemotePoolPolicyConfigUpdates is PoolConfigManagerRolesSetup {
  /// @notice Removes the last record and its policy, then sets a record for another policy. The connection store
  /// emits the events.
  function test_applyRemotePoolPolicyConfigUpdates() public {
    bytes[] memory remotePoolAddresses = _getRemotePoolAddresses(abi.encode(s_remotePool));
    _setRemotePoolPolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR, remotePoolAddresses);
    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory policyConfigs =
      new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](1);
    policyConfigs[0] = ConnectionPolicyStore.RemotePoolPolicyConfigArgs({
      policyName: DEFAULT_POLICY, remoteChainSelector: OTHER_CHAIN_SELECTOR, remotePoolAddresses: remotePoolAddresses
    });

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.RemotePoolPolicyConfigRemoved(OTHER_POLICY, OTHER_CHAIN_SELECTOR);
    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.RemotePoolPolicyRemoved(OTHER_POLICY);
    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.RemotePoolPolicyConfigSet(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, remotePoolAddresses);
    s_poolConfigManager.applyRemotePoolPolicyConfigUpdates(
      _policyRecordKeysOf(OTHER_POLICY, OTHER_CHAIN_SELECTOR), _policyNamesOf(OTHER_POLICY), policyConfigs
    );

    (, bytes32[] memory remotePoolPolicyNames,,,,) = s_poolConfigManager.getPolicyNames();
    assertEq(_policyNamesOf(DEFAULT_POLICY), remotePoolPolicyNames);
  }

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyRemotePoolPolicyConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyRemotePoolPolicyConfigUpdates,
        (
          new PolicyConfig.PolicyRecordKey[](0),
          new bytes32[](0),
          new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](0)
        )
      ),
      _getPolicyWriteRoles(s_poolConfigManager.REMOTE_POOL_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applyRemotePoolPolicyConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyRemotePoolPolicyConfigUpdates,
        (
          new PolicyConfig.PolicyRecordKey[](0),
          new bytes32[](0),
          new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](0)
        )
      ),
      _getPolicyWriteRoles(s_poolConfigManager.REMOTE_POOL_ROLE())
    );
  }
}
