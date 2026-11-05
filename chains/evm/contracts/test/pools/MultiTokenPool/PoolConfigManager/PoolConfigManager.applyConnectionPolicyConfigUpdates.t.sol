// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

import {Errors} from "@openzeppelin/contracts@5.3.0/utils/Errors.sol";

contract PoolConfigManager_applyConnectionPolicyConfigUpdates is PoolConfigManagerRolesSetup {
  /// @notice Removes the last record and its policy, then sets a record for another policy. The connection store
  /// emits the events.
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

  /// @notice Policy updates do not check lane completeness because the stores do not track the tokens assigned to
  /// each policy. Token B gains a lane despite having no fee record.
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

  /// @notice Each allowed role authorizes the call on its own.
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
      _getPolicyWriteRoles(s_poolConfigManager.CONNECTION_ROLE())
    );
  }

  // Reverts

  /// @notice Store errors reach the caller with their arguments unchanged.
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

  /// @notice Calls without an allowed role revert.
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
      _getPolicyWriteRoles(s_poolConfigManager.CONNECTION_ROLE())
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
