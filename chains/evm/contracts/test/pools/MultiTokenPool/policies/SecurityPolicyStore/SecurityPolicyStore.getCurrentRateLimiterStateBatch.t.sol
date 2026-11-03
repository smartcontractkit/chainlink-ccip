// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

contract SecurityPolicyStore_getCurrentRateLimiterStateBatch is SecurityPolicyStoreSetup {
  /// @notice getCurrentRateLimiterStateBatch shows a bucket that waits for a policy copy as a full bucket with the
  /// config of its record.
  function test_getCurrentRateLimiterStateBatch() public {
    _setRateLimitPolicyConfigPerSlot(
      DEFAULT_POLICY,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);

    IPoolConfigManager.LaneRateLimiterState[] memory states =
      s_store.getCurrentRateLimiterStateBatch(_lanesOf(s_tokenA, DEST_CHAIN_SELECTOR));
    _assertFullBucket(_getRateLimitConfig(CAPACITY), states[0].outbound);
    _assertFullBucket(_getRateLimitConfig(CAPACITY + 1), states[0].inbound);
    _assertFullBucket(_getRateLimitConfig(CAPACITY + 2), states[0].fastFinalityOutbound);
    _assertFullBucket(_getRateLimitConfig(CAPACITY + 3), states[0].fastFinalityInbound);

    // A token with no rate limit record gives zero values, because the store does not resolve the lane.
    states = s_store.getCurrentRateLimiterStateBatch(_lanesOf(s_tokenB, DEST_CHAIN_SELECTOR));
    assertEq(0, states[0].outbound.tokens);
    assertEq(0, states[0].outbound.capacity);
  }
}
