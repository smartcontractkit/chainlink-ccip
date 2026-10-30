// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../../../../interfaces/IPoolV2.sol";

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
    emit SecurityPolicyStore.InlineRateLimitConfigSet(
      s_tokenA,
      DEST_CHAIN_SELECTOR,
      inlineConfigs[0].outboundRateLimiterConfig,
      inlineConfigs[0].inboundRateLimiterConfig,
      inlineConfigs[0].fastFinalityOutboundRateLimiterConfig,
      inlineConfigs[0].fastFinalityInboundRateLimiterConfig
    );
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

  /// @notice The direction update writes the two buckets of one direction only, and keeps the other direction.
  function test_applyInlineRateLimitDirectionUpdates_UpdatesOneDirectionOnly() public {
    _setInlineRateLimitConfigPerSlot(
      s_tokenA,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );

    SecurityPolicyStore.InlineRateLimitDirectionArgs[] memory directionConfigs =
      new SecurityPolicyStore.InlineRateLimitDirectionArgs[](1);
    directionConfigs[0] = SecurityPolicyStore.InlineRateLimitDirectionArgs({
      localToken: s_tokenA,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      direction: IPoolV2.MessageDirection.Inbound,
      defaultFinalityConfig: _getRateLimitConfig(7 * CAPACITY),
      fastFinalityConfig: _getRateLimitConfig(8 * CAPACITY)
    });
    vm.prank(s_poolConfigManager);
    s_store.applyInlineRateLimitDirectionUpdates(directionConfigs);

    SecurityPolicyStore.InlineRateLimitConfigArgs memory config =
      _getInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR);
    // The inbound direction changed.
    assertEq(7 * CAPACITY, config.inboundRateLimiterConfig.capacity);
    assertEq(8 * CAPACITY, config.fastFinalityInboundRateLimiterConfig.capacity);
    // The outbound direction kept its config.
    assertEq(CAPACITY, config.outboundRateLimiterConfig.capacity);
    assertEq(CAPACITY + 2, config.fastFinalityOutboundRateLimiterConfig.capacity);
  }

  // Reverts

  /// @notice Each config of a record follows the checks of RateLimiter.
  function test_applyInlineRateLimitConfigUpdates_RevertWhen_InvalidRateLimitConfig() public {
    RateLimiter.Config memory invalidRateConfig = RateLimiter.Config({isEnabled: true, capacity: 10, rate: 11});
    SecurityPolicyStore.InlineRateLimitConfigArgs[] memory inlineConfigs =
      new SecurityPolicyStore.InlineRateLimitConfigArgs[](1);
    inlineConfigs[0] = _getInlineConfigArgs(s_tokenA, DEST_CHAIN_SELECTOR, CAPACITY);
    inlineConfigs[0].fastFinalityInboundRateLimiterConfig = invalidRateConfig;

    vm.expectRevert(abi.encodeWithSelector(RateLimiter.InvalidRateLimitRate.selector, invalidRateConfig));
    _setInlineRateLimitRecords(new IPoolConfigManager.LaneKey[](0), inlineConfigs);

    RateLimiter.Config memory disabledConfig = RateLimiter.Config({isEnabled: false, capacity: 10, rate: 0});
    vm.expectRevert(abi.encodeWithSelector(RateLimiter.DisabledNonZeroRateLimit.selector, disabledConfig));
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

  /// @notice The direction update needs the record to exist.
  function test_applyInlineRateLimitDirectionUpdates_RevertWhen_RateLimitRecordNotFound() public {
    SecurityPolicyStore.InlineRateLimitDirectionArgs[] memory directionConfigs =
      new SecurityPolicyStore.InlineRateLimitDirectionArgs[](1);
    directionConfigs[0] = SecurityPolicyStore.InlineRateLimitDirectionArgs({
      localToken: s_tokenA,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      direction: IPoolV2.MessageDirection.Outbound,
      defaultFinalityConfig: _getRateLimitConfig(CAPACITY),
      fastFinalityConfig: _getRateLimitConfig(CAPACITY)
    });

    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.RateLimitRecordNotFound.selector, s_tokenA, DEST_CHAIN_SELECTOR)
    );
    vm.prank(s_poolConfigManager);
    s_store.applyInlineRateLimitDirectionUpdates(directionConfigs);
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
      outboundRateLimiterConfig: _getRateLimitConfig(capacity),
      inboundRateLimiterConfig: _getRateLimitConfig(capacity + 1),
      fastFinalityOutboundRateLimiterConfig: _getRateLimitConfig(capacity + 2),
      fastFinalityInboundRateLimiterConfig: _getRateLimitConfig(capacity + 3)
    });
  }
}
