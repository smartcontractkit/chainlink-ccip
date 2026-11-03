// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {PoolConfigManagerTransferSetup} from "./PoolConfigManagerTransferSetup.t.sol";

contract PoolConfigManager_getCurrentRateLimiterStateBatch is PoolConfigManagerTransferSetup {
  /// @notice Each lane gets its four buckets. A send draws only from the outbound bucket of its finality class. A token
  /// without a lane gets zero buckets, also if a rate limit record would apply to it.
  function test_getCurrentRateLimiterStateBatch() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setRateLimitPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setCaller(address(s_pool));
    s_poolConfigManager.validateLockOrBurn(
      s_tokenB, DEST_CHAIN_SELECTOR, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", true
    );

    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](2);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_tokenB, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_tokenB, remoteChainSelector: OTHER_CHAIN_SELECTOR});
    IPoolConfigManager.LaneRateLimiterState[] memory states = s_poolConfigManager.getCurrentRateLimiterStateBatch(lanes);

    assertEq(CAPACITY - 100e18, states[0].outbound.tokens);
    assertEq(CAPACITY, states[0].outbound.capacity);
    assertTrue(states[0].outbound.isEnabled);
    assertEq(CAPACITY, states[0].inbound.tokens);
    assertEq(CAPACITY, states[0].fastFinalityOutbound.tokens);
    assertEq(CAPACITY, states[0].fastFinalityInbound.tokens);
    IPoolConfigManager.LaneRateLimiterState memory zeroState;
    assertEq(abi.encode(zeroState), abi.encode(states[1]));
  }
}
