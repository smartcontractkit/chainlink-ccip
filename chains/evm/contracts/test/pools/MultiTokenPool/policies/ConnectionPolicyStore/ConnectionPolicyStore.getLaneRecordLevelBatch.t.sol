// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_getLaneRecordLevelBatch is ConnectionPolicyStoreSetup {
  function test_getLaneRecordLevelBatch() public {
    // Configure connection and remote pool records independently. The subset restriction excludes s_token on OTHER.
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, SOURCE_CHAIN_SELECTOR, abi.encode(s_token)));
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(REMOTE_POOL_POLICY, SOURCE_CHAIN_SELECTOR, _getRemotePoolAddresses(s_remotePool))
    );
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, OTHER_CHAIN_SELECTOR, true));
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_token, OTHER_CHAIN_SELECTOR, _getRemotePoolConfig()));

    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](5);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: SOURCE_CHAIN_SELECTOR});
    lanes[2] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: OTHER_CHAIN_SELECTOR});
    lanes[3] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[4] =
      IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: PolicyConfig.DEFAULT_CHAIN_SELECTOR});

    (PolicyConfig.RecordLevel[] memory connectionLevels, PolicyConfig.RecordLevel[] memory remotePoolLevels) =
      s_connectionStore.getLaneRecordLevelBatch(lanes);

    assertEq(uint8(PolicyConfig.RecordLevel.PolicyChain), uint8(connectionLevels[0]));
    assertEq(uint8(PolicyConfig.RecordLevel.PolicyChain), uint8(remotePoolLevels[0]));
    assertEq(uint8(PolicyConfig.RecordLevel.Inline), uint8(connectionLevels[1]));
    assertEq(uint8(PolicyConfig.RecordLevel.PolicyChain), uint8(remotePoolLevels[1]));
    // Remote pool records resolve even without a connection.
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(connectionLevels[2]));
    assertEq(uint8(PolicyConfig.RecordLevel.Inline), uint8(remotePoolLevels[2]));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(connectionLevels[3]));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(remotePoolLevels[3]));
    // Neither policy type supports a default record at chain selector zero.
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(connectionLevels[4]));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(remotePoolLevels[4]));
  }

  function test_getLaneRecordLevelBatch_TrustPolicyRemotePools() public {
    _setInlineRemotePoolConfigTrustingPolicy(s_token, s_otherRemotePool, true);

    (, PolicyConfig.RecordLevel[] memory remotePoolLevels) =
      s_connectionStore.getLaneRecordLevelBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));

    // Trusting policy pools does not change the resolved record level from Inline.
    assertEq(uint8(PolicyConfig.RecordLevel.Inline), uint8(remotePoolLevels[0]));
  }
}
