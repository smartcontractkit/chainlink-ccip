// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract FeePolicyStore_applyTokenProfileUpdates is FeePolicyStoreSetup {
  address internal s_newToken = makeAddr("newToken");

  function test_applyTokenProfileUpdates() public {
    vm.expectEmit();
    emit FeePolicyStore.TokenProfileSet(s_newToken);
    s_feeStore.applyTokenProfileUpdates(new address[](0), _getTokenProfileArgs(s_newToken));

    (bool isRegistered, bytes32 feePolicyName, bytes32 destOverheadPolicyName) = s_feeStore.getTokenPolicies(s_newToken);
    assertTrue(isRegistered);
    assertEq(bytes32(0), feePolicyName);
    assertEq(bytes32(0), destOverheadPolicyName);
  }

  function test_applyTokenProfileUpdates_RegisteredTokenKeepsPolicies() public {
    vm.expectEmit();
    emit FeePolicyStore.TokenProfileSet(s_token);
    s_feeStore.applyTokenProfileUpdates(new address[](0), _getTokenProfileArgs(s_token));

    (bool isRegistered, bytes32 feePolicyName, bytes32 destOverheadPolicyName) = s_feeStore.getTokenPolicies(s_token);
    assertTrue(isRegistered);
    assertEq(FEE_POLICY, feePolicyName);
    assertEq(DEST_OVERHEAD_POLICY, destOverheadPolicyName);
  }

  function test_applyTokenProfileUpdates_RemoveToken() public {
    vm.expectEmit();
    emit FeePolicyStore.TokenRemoved(s_otherToken);
    s_feeStore.applyTokenProfileUpdates(_arrayOf(s_otherToken), new PolicyConfig.TokenProfileArgs[](0));

    (bool isRegistered,,) = s_feeStore.getTokenPolicies(s_otherToken);
    assertFalse(isRegistered);
  }

  function test_applyTokenProfileUpdates_RemovalsRunFirst() public {
    vm.expectEmit();
    emit FeePolicyStore.TokenRemoved(s_otherToken);
    vm.expectEmit();
    emit FeePolicyStore.TokenProfileSet(s_otherToken);
    s_feeStore.applyTokenProfileUpdates(_arrayOf(s_otherToken), _getTokenProfileArgs(s_otherToken));

    (bool isRegistered,,) = s_feeStore.getTokenPolicies(s_otherToken);
    assertTrue(isRegistered);
  }

  // Reverts

  function test_applyTokenProfileUpdates_RevertWhen_ZeroAddressInvalid() public {
    vm.expectRevert(FeePolicyStore.ZeroAddressInvalid.selector);
    s_feeStore.applyTokenProfileUpdates(new address[](0), _getTokenProfileArgs(address(0)));
  }

  function test_applyTokenProfileUpdates_RevertWhen_LocalTokenNotRegistered() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, s_newToken));
    s_feeStore.applyTokenProfileUpdates(_arrayOf(s_newToken), new PolicyConfig.TokenProfileArgs[](0));
  }

  function test_applyTokenProfileUpdates_RevertWhen_TokenHasConfig() public {
    // Each of the four kinds of config blocks the deregistration on its own.
    _setTokenFeePolicy(s_otherToken, FEE_POLICY);
    _expectTokenHasConfigRevert();
    _setTokenFeePolicy(s_otherToken, bytes32(0));

    _setTokenDestOverheadPolicy(s_otherToken, DEST_OVERHEAD_POLICY);
    _expectTokenHasConfigRevert();
    _setTokenDestOverheadPolicy(s_otherToken, bytes32(0));

    _setInlineFeeConfig(_getInlineFeeConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, _getFeeConfig()));
    _expectTokenHasConfigRevert();
    s_feeStore.applyInlineFeeConfigUpdates(
      _getLaneKeys(s_otherToken, DEST_CHAIN_SELECTOR), new FeePolicyStore.InlineFeeConfigArgs[](0)
    );

    _setInlineDestOverheadConfig(
      _getInlineDestOverheadConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, _getDestOverheadConfig())
    );
    _expectTokenHasConfigRevert();
  }

  function test_applyTokenProfileUpdates_RevertWhen_UnauthorizedCaller() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_feeStore.applyTokenProfileUpdates(new address[](0), _getTokenProfileArgs(s_newToken));
  }

  function _expectTokenHasConfigRevert() private {
    vm.expectRevert(abi.encodeWithSelector(FeePolicyStore.TokenHasConfig.selector, s_otherToken));
    s_feeStore.applyTokenProfileUpdates(_arrayOf(s_otherToken), new PolicyConfig.TokenProfileArgs[](0));
  }
}
