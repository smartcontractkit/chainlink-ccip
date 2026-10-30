// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract FeePolicyStore_applyTokenDestOverheadPolicyUpdates is FeePolicyStoreSetup {
  uint32 internal constant OTHER_DEST_GAS_OVERHEAD = 150_000;

  function setUp() public override {
    super.setUp();
    FeePolicyStore.DestOverheadConfig memory otherDestOverheadConfig = _getDestOverheadConfig();
    otherDestOverheadConfig.destGasOverhead = OTHER_DEST_GAS_OVERHEAD;
    _setDestOverheadPolicyConfig(
      _getDestOverheadPolicyConfigArgs(OTHER_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, otherDestOverheadConfig)
    );
  }

  function test_applyTokenDestOverheadPolicyUpdates() public {
    PolicyConfig.TokenPolicyArgs[] memory tokenPolicies = new PolicyConfig.TokenPolicyArgs[](2);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: s_token, policyName: OTHER_POLICY});
    tokenPolicies[1] = PolicyConfig.TokenPolicyArgs({localToken: s_otherToken, policyName: DEST_OVERHEAD_POLICY});

    vm.expectEmit();
    emit FeePolicyStore.TokenDestOverheadPolicySet(s_token, OTHER_POLICY);
    vm.expectEmit();
    emit FeePolicyStore.TokenDestOverheadPolicySet(s_otherToken, DEST_OVERHEAD_POLICY);
    s_feeStore.applyTokenDestOverheadPolicyUpdates(tokenPolicies);

    (,, bytes32 destOverheadPolicyName) = s_feeStore.getTokenPolicies(s_token);
    assertEq(OTHER_POLICY, destOverheadPolicyName);
    (,, destOverheadPolicyName) = s_feeStore.getTokenPolicies(s_otherToken);
    assertEq(DEST_OVERHEAD_POLICY, destOverheadPolicyName);
    (uint32 tokenCount,) = s_feeStore.getDestOverheadPolicy(DEST_OVERHEAD_POLICY);
    assertEq(1, tokenCount);
    (tokenCount,) = s_feeStore.getDestOverheadPolicy(OTHER_POLICY);
    assertEq(1, tokenCount);
    // The token uses the records of its new policy at once.
    IPoolConfigManager.LaneFee[] memory fees = s_feeStore.getFeeBatch(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR), FinalityCodec.WAIT_FOR_FINALITY_FLAG, _getHasLanes(true)
    );
    assertEq(OTHER_DEST_GAS_OVERHEAD, fees[0].destGasOverhead);
  }

  function test_applyTokenDestOverheadPolicyUpdates_ZeroPolicyNameClearsTokenPolicy() public {
    vm.expectEmit();
    emit FeePolicyStore.TokenDestOverheadPolicySet(s_token, bytes32(0));
    s_feeStore.applyTokenDestOverheadPolicyUpdates(_getTokenPolicyArgs(s_token, bytes32(0)));

    (,, bytes32 destOverheadPolicyName) = s_feeStore.getTokenPolicies(s_token);
    assertEq(bytes32(0), destOverheadPolicyName);
    (uint32 tokenCount,) = s_feeStore.getDestOverheadPolicy(DEST_OVERHEAD_POLICY);
    assertEq(0, tokenCount);
  }

  function test_applyTokenDestOverheadPolicyUpdates_SamePolicy() public {
    vm.expectEmit();
    emit FeePolicyStore.TokenDestOverheadPolicySet(s_token, DEST_OVERHEAD_POLICY);
    s_feeStore.applyTokenDestOverheadPolicyUpdates(_getTokenPolicyArgs(s_token, DEST_OVERHEAD_POLICY));

    (uint32 tokenCount,) = s_feeStore.getDestOverheadPolicy(DEST_OVERHEAD_POLICY);
    assertEq(1, tokenCount);
  }

  function test_applyTokenDestOverheadPolicyUpdates_SameTokenTwice() public {
    // Each item reads the policy that the item before it set, so the token counts stay correct.
    PolicyConfig.TokenPolicyArgs[] memory tokenPolicies = new PolicyConfig.TokenPolicyArgs[](2);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: s_token, policyName: OTHER_POLICY});
    tokenPolicies[1] = PolicyConfig.TokenPolicyArgs({localToken: s_token, policyName: bytes32(0)});

    s_feeStore.applyTokenDestOverheadPolicyUpdates(tokenPolicies);

    (,, bytes32 destOverheadPolicyName) = s_feeStore.getTokenPolicies(s_token);
    assertEq(bytes32(0), destOverheadPolicyName);
    (uint32 tokenCount,) = s_feeStore.getDestOverheadPolicy(DEST_OVERHEAD_POLICY);
    assertEq(0, tokenCount);
    (tokenCount,) = s_feeStore.getDestOverheadPolicy(OTHER_POLICY);
    assertEq(0, tokenCount);
  }

  function test_applyTokenDestOverheadPolicyUpdates_InlineRecordStays() public {
    _setInlineDestOverheadConfig(
      _getInlineDestOverheadConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getDestOverheadConfig())
    );

    _setTokenDestOverheadPolicy(s_token, OTHER_POLICY);

    assertEq(1, s_feeStore.getInlineDestOverheadConfigs(s_token).length);
  }

  // Reverts

  function test_applyTokenDestOverheadPolicyUpdates_RevertWhen_LocalTokenNotRegistered() public {
    address unknownToken = makeAddr("unknownToken");

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, unknownToken));
    s_feeStore.applyTokenDestOverheadPolicyUpdates(_getTokenPolicyArgs(unknownToken, DEST_OVERHEAD_POLICY));
  }

  function test_applyTokenDestOverheadPolicyUpdates_RevertWhen_PolicyNotFound() public {
    // FEE_POLICY exists only as a fee policy.
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, FEE_POLICY));
    s_feeStore.applyTokenDestOverheadPolicyUpdates(_getTokenPolicyArgs(s_token, FEE_POLICY));
  }

  function test_applyTokenDestOverheadPolicyUpdates_RevertWhen_UnauthorizedCaller() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_feeStore.applyTokenDestOverheadPolicyUpdates(_getTokenPolicyArgs(s_token, OTHER_POLICY));
  }
}
