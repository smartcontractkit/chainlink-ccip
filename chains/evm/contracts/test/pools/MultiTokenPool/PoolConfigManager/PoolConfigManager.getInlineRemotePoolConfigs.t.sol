// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT remote pool policy, which has a record on the destination chain. Token B has
/// a lane to the destination chain and no lane to the other chain.
contract PoolConfigManager_getInlineRemotePoolConfigs is PoolConfigManagerSetup {
  /// @notice Returns both inline remote pool records, including the record for a chain without a lane.
  function test_getInlineRemotePoolConfigs() public {
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory expectedConfigs =
      new ConnectionPolicyStore.InlineRemotePoolConfigArgs[](2);
    expectedConfigs[0] = _getInlineRemotePoolConfigArgs(s_tokenB, DEST_CHAIN_SELECTOR);
    expectedConfigs[1] = _getInlineRemotePoolConfigArgs(s_tokenB, OTHER_CHAIN_SELECTOR);
    expectedConfigs[1].remotePoolConfig.useLegacyPoolDataOnSend = true;
    expectedConfigs[1].remotePoolConfig.remotePools = new IPoolConfigManager.RemotePool[](0);
    _applyInlineRemotePoolConfigs(expectedConfigs);

    // Compare raw return data to avoid a stack-too-deep error when decoding the nested records.
    (bool success, bytes memory returnData) = address(s_poolConfigManager)
      .staticcall(abi.encodeCall(s_poolConfigManager.getInlineRemotePoolConfigs, (s_tokenB)));
    assertTrue(success);
    assertEq(abi.encode(expectedConfigs), returnData);
  }

  /// @dev Trusts the remote pool and enables allowLegacyPoolDataOnReceive.
  function _getInlineRemotePoolConfigArgs(
    address localToken,
    uint64 remoteChainSelector
  ) internal view returns (ConnectionPolicyStore.InlineRemotePoolConfigArgs memory) {
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePools[0].allowLegacyPoolDataOnReceive = true;
    return ConnectionPolicyStore.InlineRemotePoolConfigArgs({
      localToken: localToken,
      remoteChainSelector: remoteChainSelector,
      trustPolicyRemotePools: false,
      remotePoolConfig: remotePoolConfig
    });
  }

  function _applyInlineRemotePoolConfigs(
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs
  ) internal {
    s_poolConfigManager.applyInlineRemotePoolConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }
}
