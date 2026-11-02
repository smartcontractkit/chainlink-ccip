// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";

/// @dev Minimal stub of the WIP ConnectionPolicyStore, only so PoolConfigManager.sol can compile on this branch. The
/// real contract lives on the ConnectionPolicyStore branch.
abstract contract ConnectionPolicyStore {
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
}
