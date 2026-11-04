// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_hasLaneBatch is ConnectionPolicyStoreSetup {
  function test_hasLaneBatch() public {
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, OTHER_CHAIN_SELECTOR, true));
    _setInlineConnectionConfig(
      _getInlineConnectionConfigArgs(s_otherToken, SOURCE_CHAIN_SELECTOR, abi.encode(s_otherToken))
    );

    bool[] memory hasLanes = s_connectionStore.hasLaneBatch(_getLanes());

    assertEq(6, hasLanes.length);
    assertTrue(hasLanes[0]);
    assertFalse(hasLanes[1]);
    assertFalse(hasLanes[2]);
    assertTrue(hasLanes[3]);
    assertFalse(hasLanes[4]);
    assertFalse(hasLanes[5]);
  }

  function test_hasLaneBatch_IgnoresPause() public {
    s_connectionStore.setGlobalPause(true);
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Both, true);

    bool[] memory hasLanes = s_connectionStore.hasLaneBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));

    assertTrue(hasLanes[0]);
  }

  function _getLanes() private view returns (IPoolConfigManager.LaneKey[] memory lanes) {
    lanes = new IPoolConfigManager.LaneKey[](6);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: SOURCE_CHAIN_SELECTOR});
    lanes[2] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[3] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: SOURCE_CHAIN_SELECTOR});
    lanes[4] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: OTHER_CHAIN_SELECTOR});
    lanes[5] =
      IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: PolicyConfig.DEFAULT_CHAIN_SELECTOR});
    return lanes;
  }
}
