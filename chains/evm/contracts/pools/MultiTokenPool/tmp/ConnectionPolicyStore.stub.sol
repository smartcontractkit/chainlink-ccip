// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";
import {PolicyConfig} from "../../../libraries/PolicyConfig.sol";

/// @dev Minimal stub of the WIP ConnectionPolicyStore, only so PoolConfigManager.sol can compile on this branch. The
/// real contract lives on the ConnectionPolicyStore branch.
abstract contract ConnectionPolicyStore {
  struct ConnectionPolicyConfigArgs {
    bytes32 policyName;
    uint64 remoteChainSelector;
    bool onlySubsetTokens;
  }

  struct RemotePoolConfig {
    bool useLegacyPoolDataOnSend;
    IPoolConfigManager.RemotePool[] remotePools;
  }

  struct RemotePoolPolicyConfigArgs {
    bytes32 policyName;
    uint64 remoteChainSelector;
    bytes[] remotePoolAddresses;
  }

  struct InlineConnectionConfigArgs {
    address localToken;
    uint64 remoteChainSelector;
    bytes remoteToken;
  }

  struct InlineRemotePoolConfigArgs {
    address localToken;
    uint64 remoteChainSelector;
    bool trustPolicyRemotePools;
    RemotePoolConfig remotePoolConfig;
  }

  function resolveOutboundLane(
    address localToken,
    uint64 remoteChainSelector,
    bytes4 requestedFinalityConfig,
    bytes calldata tokenArgs
  ) external view virtual returns (IPoolConfigManager.OutboundLane memory);

  function resolveInboundLane(
    address localToken,
    uint64 remoteChainSelector,
    bytes calldata sourcePoolAddress,
    bytes calldata sourcePoolData,
    uint256 sourceDenominatedAmount,
    bytes4 requestedFinalityConfig
  ) external view virtual returns (IPoolConfigManager.ResolvedInboundTransfer memory);

  function checkRequestedFinalityBatch(
    IPoolConfigManager.LaneKey[] calldata lanes,
    bytes4 requestedFinalityConfig
  ) external view virtual;

  function hasLaneBatch(
    IPoolConfigManager.LaneKey[] calldata lanes
  ) external view virtual returns (bool[] memory);

  function getLaneConnectionBatch(
    IPoolConfigManager.LaneKey[] calldata lanes
  ) external view virtual returns (IPoolConfigManager.LaneConnection[] memory);

  function getTokenConfigBatch(
    address[] calldata localTokens
  ) external view virtual returns (IPoolConfigManager.TokenConfig[] memory);

  function getAllConfiguredTokens(
    uint64 startIndex,
    uint64 maxCount
  ) external view virtual returns (address[] memory);

  function getLaneRecordLevelBatch(
    IPoolConfigManager.LaneKey[] calldata lanes
  )
    external
    view
    virtual
    returns (PolicyConfig.RecordLevel[] memory connectionLevels, PolicyConfig.RecordLevel[] memory remotePoolLevels);

  function getPolicyNames()
    external
    view
    virtual
    returns (bytes32[] memory connectionPolicyNames, bytes32[] memory remotePoolPolicyNames);
}
