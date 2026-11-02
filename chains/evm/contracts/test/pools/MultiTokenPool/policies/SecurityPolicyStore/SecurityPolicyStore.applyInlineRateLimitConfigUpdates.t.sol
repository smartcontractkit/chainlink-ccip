// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {RateLimiter} from "../../../../../libraries/RateLimiter.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract SecurityPolicyStore_applyInlineRateLimitConfigUpdates is SecurityPolicyStoreSetup {
  /// @notice One call sets several records. Each of the four configs has its own value, so a config in a wrong slot is
  /// caught on the read-back.
  function test_applyInlineRateLimitConfigUpdates_SetsSeveralRecords() public {
    SecurityPolicyStore.InlineRateLimitConfigArgs[] memory inlineConfigs =
      new SecurityPolicyStore.InlineRateLimitConfigArgs[](2);
    inlineConfigs[0] = _getInlineConfigArgs(s_tokenA, DEST_CHAIN_SELECTOR, CAPACITY);
    inlineConfigs[1] = _getInlineConfigArgs(s_tokenA, OTHER_CHAIN_SELECTOR, 2 * CAPACITY);

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.InlineRateLimitConfigSet(s_tokenA, DEST_CHAIN_SELECTOR, inlineConfigs[0].rateLimitConfig);
    _setInlineRateLimitRecords(new IPoolConfigManager.LaneKey[](0), inlineConfigs);

    assertEq(abi.encode(inlineConfigs[0]), abi.encode(_getInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR)));
    assertEq(abi.encode(inlineConfigs[1]), abi.encode(_getInlineRateLimitConfig(s_tokenA, OTHER_CHAIN_SELECTOR)));
  }

  /// @notice A set replaces an existing record. A set of the same config keeps the state of the buckets, and a set of
  /// another config refills them.
  function test_applyInlineRateLimitConfigUpdates_ReplacesRecordAndRefillsOnChange() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 400e18);

    // The same config keeps the consumed amount.
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    assertEq(CAPACITY - 400e18, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);

    // Another capacity refills the bucket to the new capacity.
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(2 * CAPACITY));
    IPoolConfigManager.LaneRateLimiterState memory state = _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR);
    assertEq(2 * CAPACITY, state.outbound.tokens);
    assertEq(2 * CAPACITY, state.outbound.capacity);
  }

  /// @notice A removal and a set of the same config in one call recreates the buckets, so they refill, which a plain
  /// set of the same config does not do.
  function test_applyInlineRateLimitConfigUpdates_RemoveAndSetSameConfigRefillsBuckets() public {
    RateLimiter.Config memory config = _getRateLimitConfig(CAPACITY);
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, config);
    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 400e18);

    SecurityPolicyStore.InlineRateLimitConfigArgs[] memory inlineConfigs =
      new SecurityPolicyStore.InlineRateLimitConfigArgs[](1);
    inlineConfigs[0] = _getInlineConfigArgs(s_tokenA, DEST_CHAIN_SELECTOR, CAPACITY);
    _setInlineRateLimitRecords(_lanesOf(s_tokenA, DEST_CHAIN_SELECTOR), inlineConfigs);

    assertEq(CAPACITY, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);
  }

  /// @notice A removal deletes the buckets of the record.
  function test_applyInlineRateLimitConfigUpdates_RemovalDeletesBuckets() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _consume(s_tokenA, DEST_CHAIN_SELECTOR, 400e18);

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.InlineRateLimitConfigRemoved(s_tokenA, DEST_CHAIN_SELECTOR);
    _setInlineRateLimitRecords(
      _lanesOf(s_tokenA, DEST_CHAIN_SELECTOR), new SecurityPolicyStore.InlineRateLimitConfigArgs[](0)
    );

    assertEq(0, s_store.getInlineRateLimitConfigs(s_tokenA).length);
    assertEq(0, _getRateLimiterState(s_tokenA, DEST_CHAIN_SELECTOR).outbound.tokens);
  }

  // Reverts

  /// @notice Each config of a record follows the checks of RateLimiter, and the revert names the record.
  function test_applyInlineRateLimitConfigUpdates_RevertWhen_InvalidRateLimitConfig() public {
    SecurityPolicyStore.InlineRateLimitConfigArgs[] memory inlineConfigs =
      new SecurityPolicyStore.InlineRateLimitConfigArgs[](1);
    inlineConfigs[0] = _getInlineConfigArgs(s_tokenA, DEST_CHAIN_SELECTOR, CAPACITY);
    inlineConfigs[0].rateLimitConfig.fastFinalityInbound = RateLimiter.Config({isEnabled: true, capacity: 10, rate: 11});

    vm.expectRevert(
      abi.encodeWithSelector(
        SecurityPolicyStore.InvalidInlineRateLimitConfig.selector,
        s_tokenA,
        DEST_CHAIN_SELECTOR,
        inlineConfigs[0].rateLimitConfig
      )
    );
    _setInlineRateLimitRecords(new IPoolConfigManager.LaneKey[](0), inlineConfigs);

    RateLimiter.Config memory disabledConfig = RateLimiter.Config({isEnabled: false, capacity: 10, rate: 0});
    SecurityPolicyStore.RateLimitConfig memory disabledRecord = SecurityPolicyStore.RateLimitConfig({
      outbound: disabledConfig,
      inbound: disabledConfig,
      fastFinalityOutbound: disabledConfig,
      fastFinalityInbound: disabledConfig
    });
    vm.expectRevert(
      abi.encodeWithSelector(
        SecurityPolicyStore.InvalidInlineRateLimitConfig.selector, s_tokenA, DEST_CHAIN_SELECTOR, disabledRecord
      )
    );
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, disabledConfig);
  }

  /// @notice A token has no default record, so an inline record needs a chain.
  function test_applyInlineRateLimitConfigUpdates_RevertWhen_ZeroChainSelectorNotAllowed() public {
    vm.expectRevert(PolicyConfig.ZeroChainSelectorNotAllowed.selector);
    _setInlineRateLimitConfig(s_tokenA, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
  }

  /// @notice An inline record needs a registered token.
  function test_applyInlineRateLimitConfigUpdates_RevertWhen_LocalTokenNotRegistered() public {
    address unknownToken = makeAddr("unknownToken");
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, unknownToken));
    _setInlineRateLimitConfig(unknownToken, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
  }

  /// @notice A removal names an inline record that exists.
  function test_applyInlineRateLimitConfigUpdates_RevertWhen_InlineRecordNotFound() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.InlineRecordNotFound.selector, s_tokenA, DEST_CHAIN_SELECTOR));
    _setInlineRateLimitRecords(
      _lanesOf(s_tokenA, DEST_CHAIN_SELECTOR), new SecurityPolicyStore.InlineRateLimitConfigArgs[](0)
    );
  }

  /// @notice Only an authorized pool config manager writes to the store.
  function test_applyInlineRateLimitConfigUpdates_RevertWhen_UnauthorizedCaller() public {
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, address(this)));
    s_store.applyInlineRateLimitConfigUpdates(
      new IPoolConfigManager.LaneKey[](0), new SecurityPolicyStore.InlineRateLimitConfigArgs[](0)
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

  /// @dev Gives each of the four configs its own capacity.
  function _getInlineConfigArgs(
    address localToken,
    uint64 remoteChainSelector,
    uint128 capacity
  ) internal pure returns (SecurityPolicyStore.InlineRateLimitConfigArgs memory) {
    return SecurityPolicyStore.InlineRateLimitConfigArgs({
      localToken: localToken,
      remoteChainSelector: remoteChainSelector,
      rateLimitConfig: SecurityPolicyStore.RateLimitConfig({
        outbound: _getRateLimitConfig(capacity),
        inbound: _getRateLimitConfig(capacity + 1),
        fastFinalityOutbound: _getRateLimitConfig(capacity + 2),
        fastFinalityInbound: _getRateLimitConfig(capacity + 3)
      })
    });
  }
}
