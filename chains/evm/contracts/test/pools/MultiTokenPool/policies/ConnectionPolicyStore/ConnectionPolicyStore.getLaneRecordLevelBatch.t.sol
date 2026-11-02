// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_getLaneRecordLevelBatch is ConnectionPolicyStoreSetup {
  function test_getLaneRecordLevelBatch() public {
    // SOURCE gets an inline connection record and a remote pool policy chain record. OTHER gets a subset record that
    // does not serve s_token, and an inline remote pool record.
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
    // The remote pool level resolves also where the token has no lane.
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(connectionLevels[2]));
    assertEq(uint8(PolicyConfig.RecordLevel.Inline), uint8(remotePoolLevels[2]));
    // s_otherToken has no policy and no inline record.
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(connectionLevels[3]));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(remotePoolLevels[3]));
    // No policy of this store has a policy default record, so chain selector zero never resolves to a record.
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(connectionLevels[4]));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(remotePoolLevels[4]));
  }

  function test_getLaneRecordLevelBatch_TrustPolicyRemotePools() public {
    _setInlineRemotePoolConfigTrustingPolicy(s_token, s_otherRemotePool, true);

    (, PolicyConfig.RecordLevel[] memory remotePoolLevels) =
      s_connectionStore.getLaneRecordLevelBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));

    // The level is that of the inline record, also when it trusts the remote pools of the policy record.
    assertEq(uint8(PolicyConfig.RecordLevel.Inline), uint8(remotePoolLevels[0]));
  }
}
