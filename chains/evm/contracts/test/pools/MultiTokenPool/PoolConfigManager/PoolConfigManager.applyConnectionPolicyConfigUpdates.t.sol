// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

import {Errors} from "@openzeppelin/contracts@5.3.0/utils/Errors.sol";

contract PoolConfigManager_applyConnectionPolicyConfigUpdates is PoolConfigManagerRolesSetup {
  /// @notice CONFIG_ROLE removes the last record of a policy, removes the policy and sets a record of another policy in
  /// one call. The connection store applies the call and emits the events.
  function test_applyConnectionPolicyConfigUpdates() public {
    _setConnectionPolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR, false);
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs =
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](1);
    policyConfigs[0] = ConnectionPolicyStore.ConnectionPolicyConfigArgs({
      policyName: DEFAULT_POLICY, remoteChainSelector: OTHER_CHAIN_SELECTOR, onlySubsetTokens: true
    });

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.ConnectionPolicyConfigRemoved(OTHER_POLICY, OTHER_CHAIN_SELECTOR);
    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.ConnectionPolicyRemoved(OTHER_POLICY);
    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.ConnectionPolicyConfigSet(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, true);
    s_poolConfigManager.applyConnectionPolicyConfigUpdates(
      _policyRecordKeysOf(OTHER_POLICY, OTHER_CHAIN_SELECTOR), _policyNamesOf(OTHER_POLICY), policyConfigs
    );

    (bytes32[] memory connectionPolicyNames,,,,,) = s_poolConfigManager.getPolicyNames();
    assertEq(_policyNamesOf(DEFAULT_POLICY), connectionPolicyNames);
  }

  /// @notice The PCM runs no lane check for a policy record. A set opens its chain for every token of the policy, and
  /// no store lists those tokens. A removal opens no lane. Token B has no fee record on any chain.
  function test_applyConnectionPolicyConfigUpdates_PolicyRecordsAreNotChecked() public {
    _setTokenFeePolicy(s_tokenB, bytes32(0));
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs =
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](1);
    policyConfigs[0] = ConnectionPolicyStore.ConnectionPolicyConfigArgs({
      policyName: DEFAULT_POLICY, remoteChainSelector: OTHER_CHAIN_SELECTOR, onlySubsetTokens: false
    });

    s_poolConfigManager.applyConnectionPolicyConfigUpdates(
      _policyRecordKeysOf(DEFAULT_POLICY, DEST_CHAIN_SELECTOR), new bytes32[](0), policyConfigs
    );

    assertTrue(_getLaneConnection(s_tokenB, OTHER_CHAIN_SELECTOR).hasConnection);
    assertFalse(_hasFeeRecord(s_tokenB, OTHER_CHAIN_SELECTOR));
    assertFalse(_getLaneConnection(s_tokenB, DEST_CHAIN_SELECTOR).hasConnection);
  }

  /// @notice CONFIG_ROLE, POLICY_ROLE and the policy role of the type each pass alone.
  function test_applyConnectionPolicyConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyConnectionPolicyConfigUpdates,
        (
          new PolicyConfig.PolicyRecordKey[](0),
          new bytes32[](0),
          new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](0)
        )
      ),
      _getPolicyWriteRoles(s_poolConfigManager.CONNECTION_POLICY_ROLE())
    );
  }

  // Reverts

  /// @notice A store revert reaches the caller unchanged, with the error and the arguments of the store.
  function test_applyConnectionPolicyConfigUpdates_RevertWhen_PolicyNotFound() public {
    vm.expectRevert(
      abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, OTHER_POLICY), address(s_connectionStore)
    );
    s_poolConfigManager.applyConnectionPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0),
      _policyNamesOf(OTHER_POLICY),
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](0)
    );
  }

  /// @notice Address.functionCall turns a store revert without data into FailedCall.
  function test_applyConnectionPolicyConfigUpdates_RevertWhen_StoreRevertsWithoutData() public {
    PolicyConfig.PolicyRecordKey[] memory policyConfigsToRemove = new PolicyConfig.PolicyRecordKey[](0);
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigsToSet =
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](0);
    vm.mockCallRevert(
      address(s_connectionStore),
      abi.encodeCall(
        ConnectionPolicyStore.applyConnectionPolicyConfigUpdates,
        (policyConfigsToRemove, _policyNamesOf(OTHER_POLICY), policyConfigsToSet)
      ),
      ""
    );

    vm.expectRevert(Errors.FailedCall.selector);
    s_poolConfigManager.applyConnectionPolicyConfigUpdates(
      policyConfigsToRemove, _policyNamesOf(OTHER_POLICY), policyConfigsToSet
    );
  }

  /// @notice Each other PCM role and a caller with no role revert. The error names the second role of the check.
  function test_applyConnectionPolicyConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyConnectionPolicyConfigUpdates,
        (
          new PolicyConfig.PolicyRecordKey[](0),
          new bytes32[](0),
          new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](0)
        )
      ),
      _getPolicyWriteRoles(s_poolConfigManager.CONNECTION_POLICY_ROLE())
    );
  }

  function _hasFeeRecord(
    address localToken,
    uint64 remoteChainSelector
  ) internal view returns (bool) {
    (PolicyConfig.RecordLevel[] memory levels,) =
      s_feeStore.getLaneRecordLevelBatch(_lanesOf(localToken, remoteChainSelector));
    return levels[0] != PolicyConfig.RecordLevel.None;
  }
}
