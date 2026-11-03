// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {RateLimiter} from "../../../../libraries/RateLimiter.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT rate limit policy, which has a policy default record. Token B has a lane to
/// the destination chain and no lane to the other chain.
contract PoolConfigManager_getInlineRateLimitConfigs is PoolConfigManagerSetup {
  /// @notice The view returns each inline rate limit record of the token in the shape of
  /// applyInlineRateLimitConfigUpdates, in the order of the sets. Each of the four configs keeps its place. A record on
  /// a chain with a lane and a record on a chain without a lane both show.
  function test_getInlineRateLimitConfigs() public {
    SecurityPolicyStore.InlineRateLimitConfigArgs[] memory expectedConfigs =
      new SecurityPolicyStore.InlineRateLimitConfigArgs[](2);
    expectedConfigs[0] = _getInlineRateLimitConfigArgs(s_tokenB, DEST_CHAIN_SELECTOR, CAPACITY);
    expectedConfigs[1] = _getInlineRateLimitConfigArgs(s_tokenB, OTHER_CHAIN_SELECTOR, CAPACITY * 2);
    _applyInlineRateLimitConfigs(expectedConfigs);

    assertEq(abi.encode(expectedConfigs), abi.encode(s_poolConfigManager.getInlineRateLimitConfigs(s_tokenB)));
  }

  /// @dev Four different configs, so each one shows in its own place. The FTF inbound config is disabled.
  function _getInlineRateLimitConfigArgs(
    address localToken,
    uint64 remoteChainSelector,
    uint128 capacity
  ) internal pure returns (SecurityPolicyStore.InlineRateLimitConfigArgs memory) {
    return SecurityPolicyStore.InlineRateLimitConfigArgs({
      localToken: localToken,
      remoteChainSelector: remoteChainSelector,
      rateLimitConfig: SecurityPolicyStore.RateLimitConfig({
        outbound: RateLimiter.Config({isEnabled: true, capacity: capacity, rate: 1}),
        inbound: RateLimiter.Config({isEnabled: true, capacity: capacity + 1, rate: 2}),
        fastFinalityOutbound: RateLimiter.Config({isEnabled: true, capacity: capacity + 2, rate: 3}),
        fastFinalityInbound: RateLimiter.Config({isEnabled: false, capacity: 0, rate: 0})
      })
    });
  }

  function _applyInlineRateLimitConfigs(
    SecurityPolicyStore.InlineRateLimitConfigArgs[] memory inlineConfigs
  ) internal {
    s_poolConfigManager.applyInlineRateLimitConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }
}
