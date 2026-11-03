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

    // Another capacity refills the bucket to the new capacity on the next use.
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(2 * CAPACITY));
    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 100e18);
    assertEq(2 * CAPACITY - 100e18, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);
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
    SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory policyConfigs =
      new SecurityPolicyStore.RateLimitPolicyConfigArgs[](1);
    policyConfigs[0] = _getPolicyConfigArgs(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, CAPACITY);
    policyConfigs[0].rateLimitConfig.fastFinalityOutbound =
      RateLimiter.Config({isEnabled: true, capacity: 10, rate: 11});

    vm.expectRevert(
      abi.encodeWithSelector(
        SecurityPolicyStore.InvalidRateLimitPolicyConfig.selector,
        DEFAULT_POLICY,
        DEST_CHAIN_SELECTOR,
        policyConfigs[0].rateLimitConfig
      )
    );
    _applyRateLimitPolicyConfigUpdates(policyConfigs);
  }

  /// @notice Only an authorized pool config manager writes to the store.
  function test_applyRateLimitPolicyConfigUpdates_RevertWhen_UnauthorizedCaller() public {
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, address(this)));
    s_store.applyRateLimitPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), new SecurityPolicyStore.RateLimitPolicyConfigArgs[](0)
    );
  }

  // Helpers

  function _consume(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount
  ) internal {
    _setInlineSecurityConfig(localToken, remoteChainSelector, _ccvsOf(s_ccv1, s_ccv2), new address[](0), s_hook);
    vm.prank(s_poolConfigManager);
    s_store.validateLockOrBurn(localToken, remoteChainSelector, amount, bytes4(0));
  }

  function _applyRateLimitPolicyConfigUpdates(
    SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory policyConfigs
  ) internal {
    vm.prank(s_poolConfigManager);
    s_store.applyRateLimitPolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs);
  }

  function _removeRateLimitPolicyRecords(
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal {
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove = new PolicyConfig.PolicyRecordKey[](1);
    recordsToRemove[0] =
      PolicyConfig.PolicyRecordKey({policyName: policyName, remoteChainSelector: remoteChainSelector});
    vm.prank(s_poolConfigManager);
    s_store.applyRateLimitPolicyConfigUpdates(
      recordsToRemove, new bytes32[](0), new SecurityPolicyStore.RateLimitPolicyConfigArgs[](0)
    );
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

  function _arrayOfBytes32(
    bytes32 value
  ) internal pure returns (bytes32[] memory arr) {
    arr = new bytes32[](1);
    arr[0] = value;
    return arr;
  }
}
