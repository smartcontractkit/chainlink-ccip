// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

contract PoolConfigManager_getLaneConnectionBatch is PoolConfigManagerSetup {
  /// @notice Returns connection config, remote pool config, and effective pause state for each lane. Missing lanes
  /// still return their pause state. Token A uses its inline remote token address. Token B uses its policy's
  /// provider.
  function test_getLaneConnectionBatch() public {
    s_poolConfigManager.pauseTokens(_arrayOf(s_tokenB), ConnectionPolicyStore.PauseDirection.Outbound);

    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](3);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_tokenB, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_tokenB, remoteChainSelector: OTHER_CHAIN_SELECTOR});
    lanes[2] = IPoolConfigManager.LaneKey({localToken: s_tokenA, remoteChainSelector: DEST_CHAIN_SELECTOR});
    IPoolConfigManager.LaneConnection[] memory connections = s_poolConfigManager.getLaneConnectionBatch(lanes);

    assertTrue(connections[0].hasConnection);
    assertTrue(connections[0].isOutboundPaused);
    assertFalse(connections[0].isInboundPaused);
    assertFalse(connections[0].useLegacyPoolDataOnSend);
    assertEq(abi.encode(s_tokenB), connections[0].remoteToken);
    assertEq(1, connections[0].remotePools.length);
    assertFalse(connections[0].remotePools[0].allowLegacyPoolDataOnReceive);
    assertEq(abi.encode(s_remotePool), connections[0].remotePools[0].remotePoolAddress);

    assertFalse(connections[1].hasConnection);
    assertTrue(connections[1].isOutboundPaused);
    assertEq(0, connections[1].remoteToken.length);
    assertEq(0, connections[1].remotePools.length);

    assertTrue(connections[2].hasConnection);
    assertFalse(connections[2].isOutboundPaused);
    assertEq(abi.encode(s_remoteTokenA), connections[2].remoteToken);
  }
}
