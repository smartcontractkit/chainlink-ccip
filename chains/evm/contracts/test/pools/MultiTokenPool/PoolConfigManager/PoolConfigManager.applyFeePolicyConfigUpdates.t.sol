// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyFeePolicyConfigUpdates is PoolConfigManagerRolesSetup {
  /// @notice Removes the last record and its policy, then sets a record for another policy. The fee store emits the
  /// events.
  function test_applyFeePolicyConfigUpdates() public {
    FeePolicyStore.FeeConfig memory feeConfig = _getFeeConfig(10, 5);
    _setFeePolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR, feeConfig);
    FeePolicyStore.FeePolicyConfigArgs[] memory policyConfigs = new FeePolicyStore.FeePolicyConfigArgs[](1);
    policyConfigs[0] = FeePolicyStore.FeePolicyConfigArgs({
      policyName: DEFAULT_POLICY, remoteChainSelector: OTHER_CHAIN_SELECTOR, feeConfig: feeConfig
    });

    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.FeePolicyConfigRemoved(OTHER_POLICY, OTHER_CHAIN_SELECTOR);
    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.FeePolicyRemoved(OTHER_POLICY);
    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.FeePolicyConfigSet(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, feeConfig);
    s_poolConfigManager.applyFeePolicyConfigUpdates(
      _policyRecordKeysOf(OTHER_POLICY, OTHER_CHAIN_SELECTOR), _policyNamesOf(OTHER_POLICY), policyConfigs
    );

    (,,, bytes32[] memory feePolicyNames,,) = s_poolConfigManager.getPolicyNames();
    assertEq(_policyNamesOf(DEFAULT_POLICY), feePolicyNames);
  }

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyFeePolicyConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyFeePolicyConfigUpdates,
        (new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), new FeePolicyStore.FeePolicyConfigArgs[](0))
      ),
      _getPolicyWriteRoles(s_poolConfigManager.FEE_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applyFeePolicyConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyFeePolicyConfigUpdates,
        (new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), new FeePolicyStore.FeePolicyConfigArgs[](0))
      ),
      _getPolicyWriteRoles(s_poolConfigManager.FEE_ROLE())
    );
  }
}
