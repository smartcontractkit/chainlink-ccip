// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract FeePolicyStore_applyFeePolicyConfigUpdates is FeePolicyStoreSetup {
  function test_applyFeePolicyConfigUpdates() public {
    FeePolicyStore.FeeConfig memory chainFeeConfig = _getFeeConfig();
    chainFeeConfig.finalityFeeUSDCents = 500;
    FeePolicyStore.FeePolicyConfigArgs[] memory policyConfigs = new FeePolicyStore.FeePolicyConfigArgs[](2);
    policyConfigs[0] = _getFeePolicyConfigArgs(OTHER_POLICY, DEST_CHAIN_SELECTOR, chainFeeConfig)[0];
    policyConfigs[1] = _getFeePolicyConfigArgs(OTHER_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _getFeeConfig())[0];

    vm.expectEmit();
    emit FeePolicyStore.FeePolicyConfigSet(OTHER_POLICY, DEST_CHAIN_SELECTOR, chainFeeConfig);
    vm.expectEmit();
    emit FeePolicyStore.FeePolicyConfigSet(OTHER_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _getFeeConfig());
    s_feeStore.applyFeePolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs);

    // The first record set under a new name creates the policy.
    (uint32 tokenCount, FeePolicyStore.FeePolicyConfigArgs[] memory storedConfigs) =
      s_feeStore.getFeePolicy(OTHER_POLICY);
    assertEq(0, tokenCount);
    assertEq(abi.encode(policyConfigs), abi.encode(storedConfigs));
    (bytes32[] memory feePolicyNames,) = s_feeStore.getPolicyNames();
    assertEq(2, feePolicyNames.length);
    assertEq(FEE_POLICY, feePolicyNames[0]);
    assertEq(OTHER_POLICY, feePolicyNames[1]);
  }

  function test_applyFeePolicyConfigUpdates_ReplaceRecord() public {
    FeePolicyStore.FeeConfig memory newFeeConfig = _getFeeConfig();
    newFeeConfig.fastFinalityTransferFeeBps = 70;
    FeePolicyStore.FeePolicyConfigArgs[] memory policyConfigs =
      _getFeePolicyConfigArgs(FEE_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, newFeeConfig);

    vm.expectEmit();
    emit FeePolicyStore.FeePolicyConfigSet(FEE_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, newFeeConfig);
    s_feeStore.applyFeePolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs);

    (uint32 tokenCount, FeePolicyStore.FeePolicyConfigArgs[] memory storedConfigs) = s_feeStore.getFeePolicy(FEE_POLICY);
    assertEq(1, tokenCount);
    assertEq(abi.encode(policyConfigs), abi.encode(storedConfigs));
  }

  function test_applyFeePolicyConfigUpdates_BoundaryValues() public {
    // A disabled record with zero in every field, and an enabled record with the highest fees.
    FeePolicyStore.FeeConfig memory disabledFeeConfig;
    FeePolicyStore.FeeConfig memory maxFeeConfig = FeePolicyStore.FeeConfig({
      finalityFeeUSDCents: type(uint32).max,
      fastFinalityFeeUSDCents: type(uint32).max,
      finalityTransferFeeBps: uint16(BPS_DIVIDER - 1),
      fastFinalityTransferFeeBps: uint16(BPS_DIVIDER - 1),
      isEnabled: true
    });
    FeePolicyStore.FeePolicyConfigArgs[] memory policyConfigs = new FeePolicyStore.FeePolicyConfigArgs[](2);
    policyConfigs[0] = _getFeePolicyConfigArgs(OTHER_POLICY, DEST_CHAIN_SELECTOR, disabledFeeConfig)[0];
    policyConfigs[1] = _getFeePolicyConfigArgs(OTHER_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, maxFeeConfig)[0];

    s_feeStore.applyFeePolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs);

    (, FeePolicyStore.FeePolicyConfigArgs[] memory storedConfigs) = s_feeStore.getFeePolicy(OTHER_POLICY);
    assertEq(abi.encode(policyConfigs), abi.encode(storedConfigs));
  }

  function test_applyFeePolicyConfigUpdates_RemoveRecord() public {
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, DEST_CHAIN_SELECTOR, _getFeeConfig()));
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove = new PolicyConfig.PolicyRecordKey[](1);
    recordsToRemove[0] =
      PolicyConfig.PolicyRecordKey({policyName: FEE_POLICY, remoteChainSelector: DEST_CHAIN_SELECTOR});

    vm.expectEmit();
    emit FeePolicyStore.FeePolicyConfigRemoved(FEE_POLICY, DEST_CHAIN_SELECTOR);
    s_feeStore.applyFeePolicyConfigUpdates(
      recordsToRemove, new bytes32[](0), new FeePolicyStore.FeePolicyConfigArgs[](0)
    );

    (, FeePolicyStore.FeePolicyConfigArgs[] memory storedConfigs) = s_feeStore.getFeePolicy(FEE_POLICY);
    assertEq(1, storedConfigs.length);
    assertEq(PolicyConfig.DEFAULT_CHAIN_SELECTOR, storedConfigs[0].remoteChainSelector);
    // The token falls back to the policy default record at once.
    (PolicyConfig.RecordLevel[] memory feeLevels,) =
      s_feeStore.getLaneRecordLevelBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));
    assertEq(uint8(PolicyConfig.RecordLevel.PolicyDefault), uint8(feeLevels[0]));
  }

  function test_applyFeePolicyConfigUpdates_RemovePolicy() public {
    _setTokenFeePolicy(s_token, bytes32(0));
    // The record removals run before the policy removals, so one call can remove the records and the policy.
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove = new PolicyConfig.PolicyRecordKey[](1);
    recordsToRemove[0] =
      PolicyConfig.PolicyRecordKey({policyName: FEE_POLICY, remoteChainSelector: PolicyConfig.DEFAULT_CHAIN_SELECTOR});
    bytes32[] memory policiesToRemove = new bytes32[](1);
    policiesToRemove[0] = FEE_POLICY;

    vm.expectEmit();
    emit FeePolicyStore.FeePolicyConfigRemoved(FEE_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR);
    vm.expectEmit();
    emit FeePolicyStore.FeePolicyRemoved(FEE_POLICY);
    s_feeStore.applyFeePolicyConfigUpdates(
      recordsToRemove, policiesToRemove, new FeePolicyStore.FeePolicyConfigArgs[](0)
    );

    (bytes32[] memory feePolicyNames, bytes32[] memory destOverheadPolicyNames) = s_feeStore.getPolicyNames();
    assertEq(0, feePolicyNames.length);
    assertEq(1, destOverheadPolicyNames.length);
    (uint32 tokenCount, FeePolicyStore.FeePolicyConfigArgs[] memory storedConfigs) = s_feeStore.getFeePolicy(FEE_POLICY);
    assertEq(0, tokenCount);
    assertEq(0, storedConfigs.length);
  }

  function test_applyFeePolicyConfigUpdates_RemovalsRunFirst() public {
    _setTokenFeePolicy(s_token, bytes32(0));
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove = new PolicyConfig.PolicyRecordKey[](1);
    recordsToRemove[0] =
      PolicyConfig.PolicyRecordKey({policyName: FEE_POLICY, remoteChainSelector: PolicyConfig.DEFAULT_CHAIN_SELECTOR});
    bytes32[] memory policiesToRemove = new bytes32[](1);
    policiesToRemove[0] = FEE_POLICY;
    FeePolicyStore.FeeConfig memory newFeeConfig = _getFeeConfig();
    newFeeConfig.finalityFeeUSDCents = 500;
    FeePolicyStore.FeePolicyConfigArgs[] memory policyConfigs =
      _getFeePolicyConfigArgs(FEE_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, newFeeConfig);

    s_feeStore.applyFeePolicyConfigUpdates(recordsToRemove, policiesToRemove, policyConfigs);

    // The set runs last, so it creates the policy again with the new record.
    (, FeePolicyStore.FeePolicyConfigArgs[] memory storedConfigs) = s_feeStore.getFeePolicy(FEE_POLICY);
    assertEq(abi.encode(policyConfigs), abi.encode(storedConfigs));
  }

  // Reverts

  function test_applyFeePolicyConfigUpdates_RevertWhen_DisabledNonZeroFeeConfig() public {
    // Each fee field of a disabled config must be zero.
    for (uint256 i = 0; i < 4; ++i) {
      FeePolicyStore.FeeConfig memory feeConfig;
      if (i == 0) feeConfig.finalityFeeUSDCents = 1;
      if (i == 1) feeConfig.fastFinalityFeeUSDCents = 1;
      if (i == 2) feeConfig.finalityTransferFeeBps = 1;
      if (i == 3) feeConfig.fastFinalityTransferFeeBps = 1;
      FeePolicyStore.FeePolicyConfigArgs[] memory policyConfigs =
        _getFeePolicyConfigArgs(FEE_POLICY, DEST_CHAIN_SELECTOR, feeConfig);

      vm.expectRevert(abi.encodeWithSelector(FeePolicyStore.DisabledNonZeroFeeConfig.selector, feeConfig));
      s_feeStore.applyFeePolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs);
    }
  }

  function test_applyFeePolicyConfigUpdates_RevertWhen_InvalidTransferFeeBps() public {
    FeePolicyStore.FeeConfig memory feeConfig = _getFeeConfig();
    feeConfig.finalityTransferFeeBps = uint16(BPS_DIVIDER);
    FeePolicyStore.FeePolicyConfigArgs[] memory policyConfigs =
      _getFeePolicyConfigArgs(FEE_POLICY, DEST_CHAIN_SELECTOR, feeConfig);

    vm.expectRevert(abi.encodeWithSelector(FeePolicyStore.InvalidTransferFeeBps.selector, BPS_DIVIDER));
    s_feeStore.applyFeePolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs);

    feeConfig = _getFeeConfig();
    feeConfig.fastFinalityTransferFeeBps = uint16(BPS_DIVIDER);
    policyConfigs = _getFeePolicyConfigArgs(FEE_POLICY, DEST_CHAIN_SELECTOR, feeConfig);

    vm.expectRevert(abi.encodeWithSelector(FeePolicyStore.InvalidTransferFeeBps.selector, BPS_DIVIDER));
    s_feeStore.applyFeePolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs);
  }

  function test_applyFeePolicyConfigUpdates_RevertWhen_ZeroPolicyNameNotAllowed() public {
    FeePolicyStore.FeePolicyConfigArgs[] memory policyConfigs =
      _getFeePolicyConfigArgs(bytes32(0), DEST_CHAIN_SELECTOR, _getFeeConfig());

    vm.expectRevert(PolicyConfig.ZeroPolicyNameNotAllowed.selector);
    s_feeStore.applyFeePolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs);
  }

  function test_applyFeePolicyConfigUpdates_RevertWhen_PolicyRecordNotFound() public {
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove = new PolicyConfig.PolicyRecordKey[](1);
    recordsToRemove[0] =
      PolicyConfig.PolicyRecordKey({policyName: FEE_POLICY, remoteChainSelector: DEST_CHAIN_SELECTOR});

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyRecordNotFound.selector, FEE_POLICY, DEST_CHAIN_SELECTOR));
    s_feeStore.applyFeePolicyConfigUpdates(
      recordsToRemove, new bytes32[](0), new FeePolicyStore.FeePolicyConfigArgs[](0)
    );
  }

  function test_applyFeePolicyConfigUpdates_RevertWhen_PolicyNotFound() public {
    // DEST_OVERHEAD_POLICY exists only as a destination overhead policy.
    bytes32[] memory policiesToRemove = new bytes32[](1);
    policiesToRemove[0] = DEST_OVERHEAD_POLICY;

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, DEST_OVERHEAD_POLICY));
    s_feeStore.applyFeePolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), policiesToRemove, new FeePolicyStore.FeePolicyConfigArgs[](0)
    );
  }

  function test_applyFeePolicyConfigUpdates_RevertWhen_PolicyInUse() public {
    bytes32[] memory policiesToRemove = new bytes32[](1);
    policiesToRemove[0] = FEE_POLICY;

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyInUse.selector, FEE_POLICY));
    s_feeStore.applyFeePolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), policiesToRemove, new FeePolicyStore.FeePolicyConfigArgs[](0)
    );
  }

  function test_applyFeePolicyConfigUpdates_RevertWhen_PolicyHasRecords() public {
    _setTokenFeePolicy(s_token, bytes32(0));
    bytes32[] memory policiesToRemove = new bytes32[](1);
    policiesToRemove[0] = FEE_POLICY;

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyHasRecords.selector, FEE_POLICY));
    s_feeStore.applyFeePolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), policiesToRemove, new FeePolicyStore.FeePolicyConfigArgs[](0)
    );
  }

  function test_applyFeePolicyConfigUpdates_RevertWhen_UnauthorizedCaller() public {
    FeePolicyStore.FeePolicyConfigArgs[] memory policyConfigs =
      _getFeePolicyConfigArgs(OTHER_POLICY, DEST_CHAIN_SELECTOR, _getFeeConfig());
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_feeStore.applyFeePolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs);
  }
}
