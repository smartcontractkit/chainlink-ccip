// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract FeePolicyStore_applyDestOverheadPolicyConfigUpdates is FeePolicyStoreSetup {
  function test_applyDestOverheadPolicyConfigUpdates() public {
    // The smallest valid overheads, and the largest.
    FeePolicyStore.DestOverheadConfig memory minDestOverheadConfig =
      FeePolicyStore.DestOverheadConfig({destGasOverhead: 1, destBytesOverhead: MIN_DEST_BYTES_OVERHEAD});
    FeePolicyStore.DestOverheadConfig memory maxDestOverheadConfig =
      FeePolicyStore.DestOverheadConfig({destGasOverhead: type(uint32).max, destBytesOverhead: type(uint32).max});
    FeePolicyStore.DestOverheadPolicyConfigArgs[] memory policyConfigs =
      new FeePolicyStore.DestOverheadPolicyConfigArgs[](2);
    policyConfigs[0] = _getDestOverheadPolicyConfigArgs(OTHER_POLICY, DEST_CHAIN_SELECTOR, minDestOverheadConfig)[0];
    policyConfigs[1] =
      _getDestOverheadPolicyConfigArgs(OTHER_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, maxDestOverheadConfig)[0];

    vm.expectEmit();
    emit FeePolicyStore.DestOverheadPolicyConfigSet(OTHER_POLICY, DEST_CHAIN_SELECTOR, minDestOverheadConfig);
    vm.expectEmit();
    emit FeePolicyStore.DestOverheadPolicyConfigSet(
      OTHER_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, maxDestOverheadConfig
    );
    s_feeStore.applyDestOverheadPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );

    // The first record set under a new name creates the policy.
    (uint32 tokenCount, FeePolicyStore.DestOverheadPolicyConfigArgs[] memory storedConfigs) =
      s_feeStore.getDestOverheadPolicy(OTHER_POLICY);
    assertEq(0, tokenCount);
    assertEq(abi.encode(policyConfigs), abi.encode(storedConfigs));
    (, bytes32[] memory destOverheadPolicyNames) = s_feeStore.getPolicyNames();
    assertEq(2, destOverheadPolicyNames.length);
    assertEq(DEST_OVERHEAD_POLICY, destOverheadPolicyNames[0]);
    assertEq(OTHER_POLICY, destOverheadPolicyNames[1]);
  }

  function test_applyDestOverheadPolicyConfigUpdates_ReplaceRecord() public {
    FeePolicyStore.DestOverheadConfig memory newDestOverheadConfig = _getDestOverheadConfig();
    newDestOverheadConfig.destBytesOverhead = MIN_DEST_BYTES_OVERHEAD * 2;
    FeePolicyStore.DestOverheadPolicyConfigArgs[] memory policyConfigs = _getDestOverheadPolicyConfigArgs(
      DEST_OVERHEAD_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, newDestOverheadConfig
    );

    vm.expectEmit();
    emit FeePolicyStore.DestOverheadPolicyConfigSet(
      DEST_OVERHEAD_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, newDestOverheadConfig
    );
    s_feeStore.applyDestOverheadPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );

    (uint32 tokenCount, FeePolicyStore.DestOverheadPolicyConfigArgs[] memory storedConfigs) =
      s_feeStore.getDestOverheadPolicy(DEST_OVERHEAD_POLICY);
    assertEq(1, tokenCount);
    assertEq(abi.encode(policyConfigs), abi.encode(storedConfigs));
  }

  function test_applyDestOverheadPolicyConfigUpdates_RemoveRecord() public {
    _setDestOverheadPolicyConfig(
      _getDestOverheadPolicyConfigArgs(DEST_OVERHEAD_POLICY, DEST_CHAIN_SELECTOR, _getDestOverheadConfig())
    );
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove = new PolicyConfig.PolicyRecordKey[](1);
    recordsToRemove[0] =
      PolicyConfig.PolicyRecordKey({policyName: DEST_OVERHEAD_POLICY, remoteChainSelector: DEST_CHAIN_SELECTOR});

    vm.expectEmit();
    emit FeePolicyStore.DestOverheadPolicyConfigRemoved(DEST_OVERHEAD_POLICY, DEST_CHAIN_SELECTOR);
    s_feeStore.applyDestOverheadPolicyConfigUpdates(
      recordsToRemove, new bytes32[](0), new FeePolicyStore.DestOverheadPolicyConfigArgs[](0)
    );

    (, FeePolicyStore.DestOverheadPolicyConfigArgs[] memory storedConfigs) =
      s_feeStore.getDestOverheadPolicy(DEST_OVERHEAD_POLICY);
    assertEq(1, storedConfigs.length);
    assertEq(PolicyConfig.DEFAULT_CHAIN_SELECTOR, storedConfigs[0].remoteChainSelector);
    // The token falls back to the policy default record at once.
    (, PolicyConfig.RecordLevel[] memory destOverheadLevels) =
      s_feeStore.getLaneRecordLevelBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));
    assertEq(uint8(PolicyConfig.RecordLevel.PolicyDefault), uint8(destOverheadLevels[0]));
  }

  function test_applyDestOverheadPolicyConfigUpdates_RemovePolicy() public {
    _setTokenDestOverheadPolicy(s_token, bytes32(0));
    // The record removals run before the policy removals, so one call can remove the records and the policy.
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove = new PolicyConfig.PolicyRecordKey[](1);
    recordsToRemove[0] = PolicyConfig.PolicyRecordKey({
      policyName: DEST_OVERHEAD_POLICY, remoteChainSelector: PolicyConfig.DEFAULT_CHAIN_SELECTOR
    });
    bytes32[] memory policiesToRemove = new bytes32[](1);
    policiesToRemove[0] = DEST_OVERHEAD_POLICY;

    vm.expectEmit();
    emit FeePolicyStore.DestOverheadPolicyConfigRemoved(DEST_OVERHEAD_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR);
    vm.expectEmit();
    emit FeePolicyStore.DestOverheadPolicyRemoved(DEST_OVERHEAD_POLICY);
    s_feeStore.applyDestOverheadPolicyConfigUpdates(
      recordsToRemove, policiesToRemove, new FeePolicyStore.DestOverheadPolicyConfigArgs[](0)
    );

    (bytes32[] memory feePolicyNames, bytes32[] memory destOverheadPolicyNames) = s_feeStore.getPolicyNames();
    assertEq(1, feePolicyNames.length);
    assertEq(0, destOverheadPolicyNames.length);
    (uint32 tokenCount, FeePolicyStore.DestOverheadPolicyConfigArgs[] memory storedConfigs) =
      s_feeStore.getDestOverheadPolicy(DEST_OVERHEAD_POLICY);
    assertEq(0, tokenCount);
    assertEq(0, storedConfigs.length);
  }

  function test_applyDestOverheadPolicyConfigUpdates_RemovalsRunFirst() public {
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove = new PolicyConfig.PolicyRecordKey[](1);
    recordsToRemove[0] = PolicyConfig.PolicyRecordKey({
      policyName: DEST_OVERHEAD_POLICY, remoteChainSelector: PolicyConfig.DEFAULT_CHAIN_SELECTOR
    });
    FeePolicyStore.DestOverheadConfig memory newDestOverheadConfig = _getDestOverheadConfig();
    newDestOverheadConfig.destGasOverhead = 150_000;
    FeePolicyStore.DestOverheadPolicyConfigArgs[] memory policyConfigs = _getDestOverheadPolicyConfigArgs(
      DEST_OVERHEAD_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, newDestOverheadConfig
    );

    s_feeStore.applyDestOverheadPolicyConfigUpdates(recordsToRemove, new bytes32[](0), policyConfigs);

    (, FeePolicyStore.DestOverheadPolicyConfigArgs[] memory storedConfigs) =
      s_feeStore.getDestOverheadPolicy(DEST_OVERHEAD_POLICY);
    assertEq(abi.encode(policyConfigs), abi.encode(storedConfigs));
  }

  // Reverts

  function test_applyDestOverheadPolicyConfigUpdates_RevertWhen_InvalidDestOverheadPolicyConfig() public {
    // A zero gas overhead, and a bytes overhead below the minimum.
    FeePolicyStore.DestOverheadConfig[] memory invalidConfigs = new FeePolicyStore.DestOverheadConfig[](2);
    invalidConfigs[0] = _getDestOverheadConfig();
    invalidConfigs[0].destGasOverhead = 0;
    invalidConfigs[1] = _getDestOverheadConfig();
    invalidConfigs[1].destBytesOverhead = MIN_DEST_BYTES_OVERHEAD - 1;

    for (uint256 i = 0; i < invalidConfigs.length; ++i) {
      FeePolicyStore.DestOverheadPolicyConfigArgs[] memory policyConfigs =
        _getDestOverheadPolicyConfigArgs(DEST_OVERHEAD_POLICY, DEST_CHAIN_SELECTOR, invalidConfigs[i]);

      vm.expectRevert(
        abi.encodeWithSelector(
          FeePolicyStore.InvalidDestOverheadPolicyConfig.selector,
          DEST_OVERHEAD_POLICY,
          DEST_CHAIN_SELECTOR,
          invalidConfigs[i]
        )
      );
      s_feeStore.applyDestOverheadPolicyConfigUpdates(
        new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
      );
    }
  }

  function test_applyDestOverheadPolicyConfigUpdates_RevertWhen_ZeroPolicyNameNotAllowed() public {
    FeePolicyStore.DestOverheadPolicyConfigArgs[] memory policyConfigs =
      _getDestOverheadPolicyConfigArgs(bytes32(0), DEST_CHAIN_SELECTOR, _getDestOverheadConfig());

    vm.expectRevert(PolicyConfig.ZeroPolicyNameNotAllowed.selector);
    s_feeStore.applyDestOverheadPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  function test_applyDestOverheadPolicyConfigUpdates_RevertWhen_PolicyRecordNotFound() public {
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove = new PolicyConfig.PolicyRecordKey[](1);
    recordsToRemove[0] =
      PolicyConfig.PolicyRecordKey({policyName: DEST_OVERHEAD_POLICY, remoteChainSelector: DEST_CHAIN_SELECTOR});

    vm.expectRevert(
      abi.encodeWithSelector(PolicyConfig.PolicyRecordNotFound.selector, DEST_OVERHEAD_POLICY, DEST_CHAIN_SELECTOR)
    );
    s_feeStore.applyDestOverheadPolicyConfigUpdates(
      recordsToRemove, new bytes32[](0), new FeePolicyStore.DestOverheadPolicyConfigArgs[](0)
    );
  }

  function test_applyDestOverheadPolicyConfigUpdates_RevertWhen_PolicyNotFound() public {
    // FEE_POLICY exists only as a fee policy.
    bytes32[] memory policiesToRemove = new bytes32[](1);
    policiesToRemove[0] = FEE_POLICY;

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, FEE_POLICY));
    s_feeStore.applyDestOverheadPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), policiesToRemove, new FeePolicyStore.DestOverheadPolicyConfigArgs[](0)
    );
  }

  function test_applyDestOverheadPolicyConfigUpdates_RevertWhen_PolicyInUse() public {
    bytes32[] memory policiesToRemove = new bytes32[](1);
    policiesToRemove[0] = DEST_OVERHEAD_POLICY;

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyInUse.selector, DEST_OVERHEAD_POLICY));
    s_feeStore.applyDestOverheadPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), policiesToRemove, new FeePolicyStore.DestOverheadPolicyConfigArgs[](0)
    );
  }

  function test_applyDestOverheadPolicyConfigUpdates_RevertWhen_PolicyHasRecords() public {
    _setTokenDestOverheadPolicy(s_token, bytes32(0));
    bytes32[] memory policiesToRemove = new bytes32[](1);
    policiesToRemove[0] = DEST_OVERHEAD_POLICY;

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyHasRecords.selector, DEST_OVERHEAD_POLICY));
    s_feeStore.applyDestOverheadPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), policiesToRemove, new FeePolicyStore.DestOverheadPolicyConfigArgs[](0)
    );
  }

  function test_applyDestOverheadPolicyConfigUpdates_RevertWhen_UnauthorizedCaller() public {
    FeePolicyStore.DestOverheadPolicyConfigArgs[] memory policyConfigs =
      _getDestOverheadPolicyConfigArgs(OTHER_POLICY, DEST_CHAIN_SELECTOR, _getDestOverheadConfig());
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_feeStore.applyDestOverheadPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }
}
