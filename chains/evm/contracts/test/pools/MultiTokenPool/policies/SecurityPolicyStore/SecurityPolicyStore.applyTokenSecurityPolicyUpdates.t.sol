// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract SecurityPolicyStore_applyTokenSecurityPolicyUpdates is SecurityPolicyStoreSetup {
  /// @notice A set binds a token to a policy and moves its token count. Zero removes the token from its policy.
  function test_applyTokenSecurityPolicyUpdates_SetsAndRemovesPolicy() public {
    _setSecurityPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.TokenSecurityPolicySet(s_tokenA, bytes32(0), DEFAULT_POLICY);
    _setTokenSecurityPolicy(s_tokenA, DEFAULT_POLICY);

    (,,, bytes32 securityPolicyName) = s_store.getTokenPolicies(s_tokenA);
    assertEq(DEFAULT_POLICY, securityPolicyName);
    (uint32 tokenCount,) = s_store.getSecurityPolicy(DEFAULT_POLICY);
    assertEq(1, tokenCount);

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.TokenSecurityPolicySet(s_tokenA, DEFAULT_POLICY, bytes32(0));
    _setTokenSecurityPolicy(s_tokenA, bytes32(0));
    (,,, securityPolicyName) = s_store.getTokenPolicies(s_tokenA);
    assertEq(bytes32(0), securityPolicyName);
    (tokenCount,) = s_store.getSecurityPolicy(DEFAULT_POLICY);
    assertEq(0, tokenCount);
  }

  // Reverts

  function test_applyTokenSecurityPolicyUpdates_RevertWhen_PolicyNotFound() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, DEFAULT_POLICY));
    _setTokenSecurityPolicy(s_tokenA, DEFAULT_POLICY);
  }

  function test_applyTokenSecurityPolicyUpdates_RevertWhen_LocalTokenNotRegistered() public {
    _setSecurityPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);

    address unknownToken = makeAddr("unknownToken");
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, unknownToken));
    _setTokenSecurityPolicy(unknownToken, DEFAULT_POLICY);
  }

  /// @notice Only an authorized pool config manager writes to the store.
  function test_applyTokenSecurityPolicyUpdates_RevertWhen_UnauthorizedCaller() public {
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, address(this)));
    s_store.applyTokenSecurityPolicyUpdates(new PolicyConfig.TokenPolicyArgs[](0));
  }
}
