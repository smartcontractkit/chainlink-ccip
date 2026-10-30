// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVConfigValidation} from "../../../../../libraries/CCVConfigValidation.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

contract SecurityPolicyStore_applySecurityPolicyConfigUpdates is SecurityPolicyStoreSetup {
  /// @notice The first record set under a new name creates the policy. The two CCV lists differ, so a swap of the two
  /// is caught on the read-back.
  function test_applySecurityPolicyConfigUpdates_CreatesPolicyAndSetsRecord() public {
    SecurityPolicyStore.SecurityPolicyConfigArgs[] memory configs =
      new SecurityPolicyStore.SecurityPolicyConfigArgs[](1);
    configs[0] = SecurityPolicyStore.SecurityPolicyConfigArgs({
      policyName: DEFAULT_POLICY,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      securityConfig: SecurityPolicyStore.SecurityConfig({
        outboundCCVs: _ccvsOf(s_ccv1, s_ccv2), inboundCCVs: _arrayOf(s_ccv2), hook: s_hook
      })
    });

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.SecurityPolicyConfigSet(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, configs[0].securityConfig);
    _setSecurityPolicyConfig(
      DEFAULT_POLICY,
      DEST_CHAIN_SELECTOR,
      configs[0].securityConfig.outboundCCVs,
      configs[0].securityConfig.inboundCCVs,
      s_hook
    );

    (uint32 tokenCount, SecurityPolicyStore.SecurityPolicyConfigArgs[] memory records) =
      s_store.getSecurityPolicy(DEFAULT_POLICY);
    assertEq(0, tokenCount);
    assertEq(1, records.length);
    assertEq(abi.encode(configs[0]), abi.encode(records[0]));
  }

  /// @notice A record removal keeps the policy. A policy removal then needs a policy without records and without tokens.
  function test_applySecurityPolicyConfigUpdates_RemovalKeepsPolicy() public {
    _setSecurityPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.SecurityPolicyConfigRemoved(DEFAULT_POLICY, DEST_CHAIN_SELECTOR);
    _removeSecurityPolicyRecords(DEFAULT_POLICY, DEST_CHAIN_SELECTOR);

    (, SecurityPolicyStore.SecurityPolicyConfigArgs[] memory records) = s_store.getSecurityPolicy(DEFAULT_POLICY);
    assertEq(0, records.length);

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.SecurityPolicyRemoved(DEFAULT_POLICY);
    _removeSecurityPolicy(DEFAULT_POLICY);
  }

  // Reverts

  /// @notice A CCV list has no duplicate.
  function test_applySecurityPolicyConfigUpdates_RevertWhen_DuplicateCCVNotAllowed() public {
    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, s_ccv1));
    _setSecurityPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv1), _arrayOf(s_ccv2), s_hook);
  }

  /// @notice A record removal names a record that exists.
  function test_applySecurityPolicyConfigUpdates_RevertWhen_PolicyRecordNotFound() public {
    vm.expectRevert(
      abi.encodeWithSelector(PolicyConfig.PolicyRecordNotFound.selector, DEFAULT_POLICY, DEST_CHAIN_SELECTOR)
    );
    _removeSecurityPolicyRecords(DEFAULT_POLICY, DEST_CHAIN_SELECTOR);
  }

  /// @notice A policy removal needs a policy that exists and has no records.
  function test_applySecurityPolicyConfigUpdates_RevertWhen_PolicyNotFound() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, DEFAULT_POLICY));
    _removeSecurityPolicy(DEFAULT_POLICY);
  }

  // Helpers

  function _removeSecurityPolicyRecords(
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal {
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove = new PolicyConfig.PolicyRecordKey[](1);
    recordsToRemove[0] =
      PolicyConfig.PolicyRecordKey({policyName: policyName, remoteChainSelector: remoteChainSelector});
    vm.prank(s_poolConfigManager);
    s_store.applySecurityPolicyConfigUpdates(
      recordsToRemove, new bytes32[](0), new SecurityPolicyStore.SecurityPolicyConfigArgs[](0)
    );
  }

  function _removeSecurityPolicy(
    bytes32 policyName
  ) internal {
    bytes32[] memory policiesToRemove = new bytes32[](1);
    policiesToRemove[0] = policyName;
    vm.prank(s_poolConfigManager);
    s_store.applySecurityPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), policiesToRemove, new SecurityPolicyStore.SecurityPolicyConfigArgs[](0)
    );
  }
}
