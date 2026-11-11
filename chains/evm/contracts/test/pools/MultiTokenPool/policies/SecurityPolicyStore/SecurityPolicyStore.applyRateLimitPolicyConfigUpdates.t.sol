// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {RateLimiter} from "../../../../../libraries/RateLimiter.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract SecurityPolicyStore_applyRateLimitPolicyConfigUpdates is SecurityPolicyStoreSetup {
  /// @notice The first record set under a new name creates the policy and fixes its decimals. Each of the four configs
  /// has its own value, so a config in a wrong slot is caught on the read-back.
  function test_applyRateLimitPolicyConfigUpdates_CreatesPolicyAndSetsRecord() public {
    SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory policyConfigs =
      new SecurityPolicyStore.RateLimitPolicyConfigArgs[](1);
    policyConfigs[0] = _getPolicyConfigArgs(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, CAPACITY);
    // A rate equal to the capacity is valid.
    policyConfigs[0].rateLimitConfig.outbound.rate = CAPACITY;

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.RateLimitPolicyConfigSet(
      DEFAULT_POLICY, DEST_CHAIN_SELECTOR, DEFAULT_TOKEN_DECIMALS, policyConfigs[0].rateLimitConfig
    );
    _applyRateLimitPolicyConfigUpdates(policyConfigs);

    (uint32 tokenCount, uint8 localDecimals, SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory records) =
      s_store.getRateLimitPolicy(DEFAULT_POLICY);
    assertEq(0, tokenCount);
    assertEq(DEFAULT_TOKEN_DECIMALS, localDecimals);
    assertEq(abi.encode(policyConfigs[0]), abi.encode(records[0]));
  }

  /// @notice A bucket copies a new config of its policy record on its next use, and the copy refills the bucket. A set
  /// of the same config refills nothing.
  function test_applyRateLimitPolicyConfigUpdates_BucketSyncsOnChangeOnly() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 400e18);
    assertEq(CAPACITY - 400e18, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);

    // The same config changes no bucket.
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 100e18);
    assertEq(CAPACITY - 500e18, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);

    // Another rate alone refills the bucket on the next use.
    RateLimiter.Config memory newConfig = _getRateLimitConfig(CAPACITY);
    newConfig.rate = 2;
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, newConfig);
    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 100e18);
    assertEq(CAPACITY - 100e18, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);
  }

  /// @notice A record removal keeps the policy, and a policy removal needs a policy without records and without tokens.
  function test_applyRateLimitPolicyConfigUpdates_RemovalKeepsPolicy() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.RateLimitPolicyConfigRemoved(DEFAULT_POLICY, DEST_CHAIN_SELECTOR);
    _removeRateLimitPolicyRecords(DEFAULT_POLICY, DEST_CHAIN_SELECTOR);

    (,, SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory records) = s_store.getRateLimitPolicy(DEFAULT_POLICY);
    assertEq(0, records.length);
    // The policy still exists, so its removal needs no record.
    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.RateLimitPolicyRemoved(DEFAULT_POLICY);
    _removeRateLimitPolicy(DEFAULT_POLICY);
  }

  // Reverts

  /// @notice Every token of a rate limit policy shares its decimals, which the first record fixes.
  function test_applyRateLimitPolicyConfigUpdates_RevertWhen_InvalidRateLimitPolicyDecimals() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));

    SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory policyConfigs =
      new SecurityPolicyStore.RateLimitPolicyConfigArgs[](1);
    policyConfigs[0] = _getPolicyConfigArgs(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, CAPACITY);
    policyConfigs[0].localDecimals = DEFAULT_TOKEN_DECIMALS + 1;

    vm.expectRevert(
      abi.encodeWithSelector(
        SecurityPolicyStore.InvalidRateLimitPolicyDecimals.selector,
        DEFAULT_POLICY,
        DEFAULT_TOKEN_DECIMALS,
        DEFAULT_TOKEN_DECIMALS + 1
      )
    );
    _applyRateLimitPolicyConfigUpdates(policyConfigs);
  }

  /// @notice A record removal names a record that exists.
  function test_applyRateLimitPolicyConfigUpdates_RevertWhen_PolicyRecordNotFound() public {
    vm.expectRevert(
      abi.encodeWithSelector(PolicyConfig.PolicyRecordNotFound.selector, DEFAULT_POLICY, DEST_CHAIN_SELECTOR)
    );
    _removeRateLimitPolicyRecords(DEFAULT_POLICY, DEST_CHAIN_SELECTOR);
  }

  /// @notice Each config of a record follows the checks of RateLimiter, and the revert names the record.
  function test_applyRateLimitPolicyConfigUpdates_RevertWhen_InvalidRateLimitConfig() public {
    RateLimiter.Config memory invalidConfig = RateLimiter.Config({isEnabled: true, capacity: 10, rate: 11});

    SecurityPolicyStore.RateLimitPolicyConfigArgs memory args =
      _getPolicyConfigArgs(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, CAPACITY);
    args.rateLimitConfig.outbound = invalidConfig;
    _expectInvalidRateLimitPolicyConfigRevert(args);

    args = _getPolicyConfigArgs(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, CAPACITY);
    args.rateLimitConfig.inbound = invalidConfig;
    _expectInvalidRateLimitPolicyConfigRevert(args);

    args = _getPolicyConfigArgs(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, CAPACITY);
    args.rateLimitConfig.fastFinalityOutbound = invalidConfig;
    _expectInvalidRateLimitPolicyConfigRevert(args);

    // A disabled config needs a zero rate as well as a zero capacity.
    args = _getPolicyConfigArgs(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, CAPACITY);
    args.rateLimitConfig.fastFinalityInbound = RateLimiter.Config({isEnabled: false, capacity: 0, rate: 1});
    _expectInvalidRateLimitPolicyConfigRevert(args);
  }

  /// @notice Only an authorized pool config manager writes to the store.
  function test_applyRateLimitPolicyConfigUpdates_RevertWhen_UnauthorizedCaller() public {
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, address(this)));
    s_store.applyRateLimitPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), new SecurityPolicyStore.RateLimitPolicyConfigArgs[](0)
    );
  }

  // Helpers

  function _applyRateLimitPolicyConfigUpdates(
    SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory policyConfigs
  ) internal {
    vm.prank(s_poolConfigManager);
    s_store.applyRateLimitPolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs);
  }

  function _removeRateLimitPolicy(
    bytes32 policyName
  ) internal {
    vm.prank(s_poolConfigManager);
    s_store.applyRateLimitPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0),
      _arrayOfBytes32(policyName),
      new SecurityPolicyStore.RateLimitPolicyConfigArgs[](0)
    );
  }

  /// @dev Gives each of the four configs its own capacity.
  function _getPolicyConfigArgs(
    bytes32 policyName,
    uint64 remoteChainSelector,
    uint128 capacity
  ) internal pure returns (SecurityPolicyStore.RateLimitPolicyConfigArgs memory) {
    return SecurityPolicyStore.RateLimitPolicyConfigArgs({
      policyName: policyName,
      remoteChainSelector: remoteChainSelector,
      localDecimals: DEFAULT_TOKEN_DECIMALS,
      rateLimitConfig: SecurityPolicyStore.RateLimitConfig({
        outbound: _getRateLimitConfig(capacity),
        inbound: _getRateLimitConfig(capacity + 1),
        fastFinalityOutbound: _getRateLimitConfig(capacity + 2),
        fastFinalityInbound: _getRateLimitConfig(capacity + 3)
      })
    });
  }

  function _expectInvalidRateLimitPolicyConfigRevert(
    SecurityPolicyStore.RateLimitPolicyConfigArgs memory args
  ) internal {
    SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory policyConfigs =
      new SecurityPolicyStore.RateLimitPolicyConfigArgs[](1);
    policyConfigs[0] = args;
    vm.expectRevert(
      abi.encodeWithSelector(
        SecurityPolicyStore.InvalidRateLimitPolicyConfig.selector,
        args.policyName,
        args.remoteChainSelector,
        args.rateLimitConfig
      )
    );
    _applyRateLimitPolicyConfigUpdates(policyConfigs);
  }

  function _arrayOfBytes32(
    bytes32 value
  ) internal pure returns (bytes32[] memory arr) {
    arr = new bytes32[](1);
    arr[0] = value;
    return arr;
  }
}
