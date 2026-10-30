// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract FeePolicyStore_applyTokenFeePolicyUpdates is FeePolicyStoreSetup {
  uint16 internal constant OTHER_TRANSFER_FEE_BPS = 50;

  function setUp() public override {
    super.setUp();
    FeePolicyStore.FeeConfig memory otherFeeConfig = _getFeeConfig();
    otherFeeConfig.finalityTransferFeeBps = OTHER_TRANSFER_FEE_BPS;
    _setFeePolicyConfig(_getFeePolicyConfigArgs(OTHER_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, otherFeeConfig));
  }

  function test_applyTokenFeePolicyUpdates() public {
    PolicyConfig.TokenPolicyArgs[] memory tokenPolicies = new PolicyConfig.TokenPolicyArgs[](2);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: s_token, policyName: OTHER_POLICY});
    tokenPolicies[1] = PolicyConfig.TokenPolicyArgs({localToken: s_otherToken, policyName: FEE_POLICY});

    vm.expectEmit();
    emit FeePolicyStore.TokenFeePolicySet(s_token, OTHER_POLICY);
    vm.expectEmit();
    emit FeePolicyStore.TokenFeePolicySet(s_otherToken, FEE_POLICY);
    s_feeStore.applyTokenFeePolicyUpdates(tokenPolicies);

    (, bytes32 feePolicyName,) = s_feeStore.getTokenPolicies(s_token);
    assertEq(OTHER_POLICY, feePolicyName);
    (, feePolicyName,) = s_feeStore.getTokenPolicies(s_otherToken);
    assertEq(FEE_POLICY, feePolicyName);
    (uint32 tokenCount,) = s_feeStore.getFeePolicy(FEE_POLICY);
    assertEq(1, tokenCount);
    (tokenCount,) = s_feeStore.getFeePolicy(OTHER_POLICY);
    assertEq(1, tokenCount);
    // The token uses the records of its new policy at once.
    assertEq(
      (1e18 * uint256(OTHER_TRANSFER_FEE_BPS)) / BPS_DIVIDER,
      s_feeStore.resolveTokenFee(s_token, DEST_CHAIN_SELECTOR, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true)
    );
  }

  function test_applyTokenFeePolicyUpdates_ZeroPolicyNameClearsTokenPolicy() public {
    vm.expectEmit();
    emit FeePolicyStore.TokenFeePolicySet(s_token, bytes32(0));
    s_feeStore.applyTokenFeePolicyUpdates(_getTokenPolicyArgs(s_token, bytes32(0)));

    (, bytes32 feePolicyName,) = s_feeStore.getTokenPolicies(s_token);
    assertEq(bytes32(0), feePolicyName);
    (uint32 tokenCount,) = s_feeStore.getFeePolicy(FEE_POLICY);
    assertEq(0, tokenCount);
  }

  function test_applyTokenFeePolicyUpdates_SamePolicy() public {
    vm.expectEmit();
    emit FeePolicyStore.TokenFeePolicySet(s_token, FEE_POLICY);
    s_feeStore.applyTokenFeePolicyUpdates(_getTokenPolicyArgs(s_token, FEE_POLICY));

    (uint32 tokenCount,) = s_feeStore.getFeePolicy(FEE_POLICY);
    assertEq(1, tokenCount);
  }

  function test_applyTokenFeePolicyUpdates_SameTokenTwice() public {
    // Each item reads the policy that the item before it set, so the token counts stay correct.
    PolicyConfig.TokenPolicyArgs[] memory tokenPolicies = new PolicyConfig.TokenPolicyArgs[](2);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: s_token, policyName: OTHER_POLICY});
    tokenPolicies[1] = PolicyConfig.TokenPolicyArgs({localToken: s_token, policyName: bytes32(0)});

    s_feeStore.applyTokenFeePolicyUpdates(tokenPolicies);

    (, bytes32 feePolicyName,) = s_feeStore.getTokenPolicies(s_token);
    assertEq(bytes32(0), feePolicyName);
    (uint32 tokenCount,) = s_feeStore.getFeePolicy(FEE_POLICY);
    assertEq(0, tokenCount);
    (tokenCount,) = s_feeStore.getFeePolicy(OTHER_POLICY);
    assertEq(0, tokenCount);
  }

  function test_applyTokenFeePolicyUpdates_InlineRecordStays() public {
    _setInlineFeeConfig(_getInlineFeeConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getFeeConfig()));

    _setTokenFeePolicy(s_token, OTHER_POLICY);

    assertEq(1, s_feeStore.getInlineFeeConfigs(s_token).length);
  }

  // Reverts

  function test_applyTokenFeePolicyUpdates_RevertWhen_LocalTokenNotRegistered() public {
    address unknownToken = makeAddr("unknownToken");

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, unknownToken));
    s_feeStore.applyTokenFeePolicyUpdates(_getTokenPolicyArgs(unknownToken, FEE_POLICY));
  }

  function test_applyTokenFeePolicyUpdates_RevertWhen_PolicyNotFound() public {
    // DEST_OVERHEAD_POLICY exists only as a destination overhead policy.
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, DEST_OVERHEAD_POLICY));
    s_feeStore.applyTokenFeePolicyUpdates(_getTokenPolicyArgs(s_token, DEST_OVERHEAD_POLICY));
  }

  function test_applyTokenFeePolicyUpdates_RevertWhen_UnauthorizedCaller() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_feeStore.applyTokenFeePolicyUpdates(_getTokenPolicyArgs(s_token, OTHER_POLICY));
  }
}
