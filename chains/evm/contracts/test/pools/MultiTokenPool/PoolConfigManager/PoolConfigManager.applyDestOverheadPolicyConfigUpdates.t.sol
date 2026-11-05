// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyDestOverheadPolicyConfigUpdates is PoolConfigManagerRolesSetup {
  /// @notice Removes the last record and its policy, then sets a record for another policy. The fee store emits the
  /// events.
  function test_applyDestOverheadPolicyConfigUpdates() public {
    _setDestOverheadPolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR, 300_000, 320);
    FeePolicyStore.DestOverheadPolicyConfigArgs[] memory policyConfigs =
      new FeePolicyStore.DestOverheadPolicyConfigArgs[](1);
    policyConfigs[0] = FeePolicyStore.DestOverheadPolicyConfigArgs({
      policyName: DEFAULT_POLICY,
      remoteChainSelector: OTHER_CHAIN_SELECTOR,
      destOverheadConfig: FeePolicyStore.DestOverheadConfig({destGasOverhead: 400_000, destBytesOverhead: 352})
    });

    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.DestOverheadPolicyConfigRemoved(OTHER_POLICY, OTHER_CHAIN_SELECTOR);
    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.DestOverheadPolicyRemoved(OTHER_POLICY);
    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.DestOverheadPolicyConfigSet(
      DEFAULT_POLICY,
      OTHER_CHAIN_SELECTOR,
      FeePolicyStore.DestOverheadConfig({destGasOverhead: 400_000, destBytesOverhead: 352})
    );
    s_poolConfigManager.applyDestOverheadPolicyConfigUpdates(
      _policyRecordKeysOf(OTHER_POLICY, OTHER_CHAIN_SELECTOR), _policyNamesOf(OTHER_POLICY), policyConfigs
    );

    (,,,, bytes32[] memory destOverheadPolicyNames,) = s_poolConfigManager.getPolicyNames();
    assertEq(_policyNamesOf(DEFAULT_POLICY), destOverheadPolicyNames);
  }

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyDestOverheadPolicyConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyDestOverheadPolicyConfigUpdates,
        (new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), new FeePolicyStore.DestOverheadPolicyConfigArgs[](0))
      ),
      _getPolicyWriteRoles(s_poolConfigManager.DEST_OVERHEAD_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applyDestOverheadPolicyConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyDestOverheadPolicyConfigUpdates,
        (new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), new FeePolicyStore.DestOverheadPolicyConfigArgs[](0))
      ),
      _getPolicyWriteRoles(s_poolConfigManager.DEST_OVERHEAD_ROLE())
    );
  }
}
