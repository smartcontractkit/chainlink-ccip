// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {RateLimiter} from "../../../../../libraries/RateLimiter.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

contract SecurityPolicyStore_applyTokenRateLimitPolicyUpdates is SecurityPolicyStoreSetup {
  /// @notice A set binds a token to a policy and moves its token count. Zero removes the token from its policy.
  function test_applyTokenRateLimitPolicyUpdates_SetsAndRemovesPolicy() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.TokenRateLimitPolicySet(s_tokenA, DEFAULT_POLICY);
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);

    (,, bytes32 rateLimitPolicyName,) = s_store.getTokenPolicies(s_tokenA);
    assertEq(DEFAULT_POLICY, rateLimitPolicyName);
    (uint32 tokenCount,,) = s_store.getRateLimitPolicy(DEFAULT_POLICY);
    assertEq(1, tokenCount);

    _setTokenRateLimitPolicy(s_tokenA, bytes32(0));
    (,, rateLimitPolicyName,) = s_store.getTokenPolicies(s_tokenA);
    assertEq(bytes32(0), rateLimitPolicyName);
    (tokenCount,,) = s_store.getRateLimitPolicy(DEFAULT_POLICY);
    assertEq(0, tokenCount);
  }

  // Reverts

  function test_applyTokenRateLimitPolicyUpdates_RevertWhen_PolicyNotFound() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, DEFAULT_POLICY));
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
  }

  /// @notice Every token of a rate limit policy shares its decimals.
  function test_applyTokenRateLimitPolicyUpdates_RevertWhen_InvalidRateLimitPolicyDecimals() public {
    _createRateLimitPolicyWithDecimals(DEFAULT_POLICY, DEFAULT_TOKEN_DECIMALS + 1);

    vm.expectRevert(
      abi.encodeWithSelector(
        SecurityPolicyStore.InvalidRateLimitPolicyDecimals.selector,
        DEFAULT_POLICY,
        DEFAULT_TOKEN_DECIMALS + 1,
        DEFAULT_TOKEN_DECIMALS
      )
    );
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
  }

  function test_applyTokenRateLimitPolicyUpdates_RevertWhen_LocalTokenNotRegistered() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));

    address unknownToken = makeAddr("unknownToken");
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, unknownToken));
    _setTokenRateLimitPolicy(unknownToken, DEFAULT_POLICY);
  }

  // Helpers

  function _createRateLimitPolicyWithDecimals(
    bytes32 policyName,
    uint8 localDecimals
  ) internal {
    SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory policyConfigs =
      new SecurityPolicyStore.RateLimitPolicyConfigArgs[](1);
    policyConfigs[0].policyName = policyName;
    policyConfigs[0].remoteChainSelector = PolicyConfig.DEFAULT_CHAIN_SELECTOR;
    policyConfigs[0].localDecimals = localDecimals;
    vm.prank(s_poolConfigManager);
    s_store.applyRateLimitPolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs);
  }
}
