// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applySecurityPolicyConfigUpdates is PoolConfigManagerRolesSetup {
  address internal s_ccv = makeAddr("ccv");

  /// @notice Removes the last record and its policy, then sets a record for another policy. The security store
  /// emits the events.
  function test_applySecurityPolicyConfigUpdates() public {
    address[] memory ccvs = _arrayOf(s_ccv);
    _setSecurityPolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR, ccvs, ccvs, address(0));
    SecurityPolicyStore.SecurityPolicyConfigArgs[] memory policyConfigs =
      new SecurityPolicyStore.SecurityPolicyConfigArgs[](1);
    policyConfigs[0] = SecurityPolicyStore.SecurityPolicyConfigArgs({
      policyName: DEFAULT_POLICY,
      remoteChainSelector: OTHER_CHAIN_SELECTOR,
      securityConfig: SecurityPolicyStore.SecurityConfig({outboundCCVs: ccvs, inboundCCVs: ccvs, hook: address(0)})
    });

    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.SecurityPolicyConfigRemoved(OTHER_POLICY, OTHER_CHAIN_SELECTOR);
    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.SecurityPolicyRemoved(OTHER_POLICY);
    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.SecurityPolicyConfigSet(
      DEFAULT_POLICY,
      OTHER_CHAIN_SELECTOR,
      SecurityPolicyStore.SecurityConfig({outboundCCVs: ccvs, inboundCCVs: ccvs, hook: address(0)})
    );
    s_poolConfigManager.applySecurityPolicyConfigUpdates(
      _policyRecordKeysOf(OTHER_POLICY, OTHER_CHAIN_SELECTOR), _policyNamesOf(OTHER_POLICY), policyConfigs
    );

    (,,,,, bytes32[] memory securityPolicyNames) = s_poolConfigManager.getPolicyNames();
    assertEq(_policyNamesOf(DEFAULT_POLICY), securityPolicyNames);
  }

  /// @notice Each allowed role authorizes the call on its own.
  function test_applySecurityPolicyConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applySecurityPolicyConfigUpdates,
        (new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), new SecurityPolicyStore.SecurityPolicyConfigArgs[](0))
      ),
      _getPolicyWriteRoles(s_poolConfigManager.SECURITY_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applySecurityPolicyConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applySecurityPolicyConfigUpdates,
        (new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), new SecurityPolicyStore.SecurityPolicyConfigArgs[](0))
      ),
      _getPolicyWriteRoles(s_poolConfigManager.SECURITY_ROLE())
    );
  }
}
