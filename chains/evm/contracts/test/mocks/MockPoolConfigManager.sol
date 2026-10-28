// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IMultiTokenPool} from "../../interfaces/IMultiTokenPool.sol";
import {IPoolConfigManager} from "../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../interfaces/IPoolV2.sol";

import {Pool} from "../../libraries/Pool.sol";

/// @dev A pool config manager with settable return values. Records the finality, token args and fee flag of the last
/// validate call. Batch views repeat the set value for each queried lane or token.
contract MockPoolConfigManager is IPoolConfigManager {
  IPoolConfigManager.OutboundLane internal s_outboundLane;
  IPoolConfigManager.ResolvedInboundTransfer internal s_inboundTransfer;
  IPoolConfigManager.TokenConfig internal s_tokenConfig;
  IPoolConfigManager.LaneConnection internal s_laneConnection;
  IPoolConfigManager.LaneFee internal s_laneFee;
  IPoolV2.TokenTransferFeeConfig internal s_feeConfig;
  address[] internal s_outboundCCVs;
  address[] internal s_inboundCCVs;

  bool internal s_lastChargeTokenFee;
  bytes4 internal s_lastRequestedFinalityConfig;
  bytes internal s_lastTokenArgs;

  function setOutboundLane(
    IPoolConfigManager.OutboundLane memory lane
  ) external {
    s_outboundLane = lane;
  }

  function setInboundTransfer(
    IPoolConfigManager.ResolvedInboundTransfer memory transfer
  ) external {
    s_inboundTransfer = transfer;
  }

  function setTokenConfig(
    IPoolConfigManager.TokenConfig memory config
  ) external {
    s_tokenConfig = config;
  }

  function setLaneConnection(
    IPoolConfigManager.LaneConnection memory connection
  ) external {
    s_laneConnection = connection;
  }

  function setLaneFee(
    IPoolConfigManager.LaneFee memory fee
  ) external {
    s_laneFee = fee;
  }

  function setTokenTransferFeeConfig(
    IPoolV2.TokenTransferFeeConfig memory feeConfig
  ) external {
    s_feeConfig = feeConfig;
  }

  function setRequiredCCVs(
    address[] memory outboundCCVs,
    address[] memory inboundCCVs
  ) external {
    s_outboundCCVs = outboundCCVs;
    s_inboundCCVs = inboundCCVs;
  }

  function lastChargeTokenFee() external view returns (bool) {
    return s_lastChargeTokenFee;
  }

  function lastRequestedFinalityConfig() external view returns (bytes4) {
    return s_lastRequestedFinalityConfig;
  }

  function lastTokenArgs() external view returns (bytes memory) {
    return s_lastTokenArgs;
  }

  function validateLockOrBurn(
    address,
    uint64,
    uint256,
    bytes4 requestedFinalityConfig,
    bytes calldata tokenArgs,
    bool chargeTokenFee
  ) external returns (OutboundLane memory) {
    s_lastChargeTokenFee = chargeTokenFee;
    s_lastRequestedFinalityConfig = requestedFinalityConfig;
    s_lastTokenArgs = tokenArgs;
    return s_outboundLane;
  }

  function validateReleaseOrMint(
    Pool.ReleaseOrMintInV1 calldata,
    bytes4 requestedFinalityConfig
  ) external returns (ResolvedInboundTransfer memory) {
    s_lastRequestedFinalityConfig = requestedFinalityConfig;
    return s_inboundTransfer;
  }

  function getFeeBatch(
    LaneKey[] calldata lanes,
    bytes4
  ) external view returns (LaneFee[] memory fees) {
    fees = new LaneFee[](lanes.length);
    for (uint256 i = 0; i < lanes.length; ++i) {
      fees[i] = s_laneFee;
    }
    return fees;
  }

  function getTokenTransferFeeConfigBatch(
    LaneKey[] calldata lanes
  ) external view returns (IPoolV2.TokenTransferFeeConfig[] memory feeConfigs) {
    feeConfigs = new IPoolV2.TokenTransferFeeConfig[](lanes.length);
    for (uint256 i = 0; i < lanes.length; ++i) {
      feeConfigs[i] = s_feeConfig;
    }
    return feeConfigs;
  }

  function getCurrentRateLimiterStateBatch(
    LaneKey[] calldata lanes
  ) external pure returns (LaneRateLimiterState[] memory states) {
    return new LaneRateLimiterState[](lanes.length);
  }

  function getLaneConnectionBatch(
    LaneKey[] calldata lanes
  ) external view returns (LaneConnection[] memory connections) {
    connections = new LaneConnection[](lanes.length);
    for (uint256 i = 0; i < lanes.length; ++i) {
      connections[i] = s_laneConnection;
    }
    return connections;
  }

  function getLaneVerificationBatch(
    LaneKey[] calldata lanes,
    bytes4
  ) external view returns (address[] memory hooks, address[][] memory outboundCCVs, address[][] memory inboundCCVs) {
    hooks = new address[](lanes.length);
    outboundCCVs = new address[][](lanes.length);
    inboundCCVs = new address[][](lanes.length);
    for (uint256 i = 0; i < lanes.length; ++i) {
      outboundCCVs[i] = s_outboundCCVs;
      inboundCCVs[i] = s_inboundCCVs;
    }
    return (hooks, outboundCCVs, inboundCCVs);
  }

  function getTokenConfigBatch(
    address[] calldata localTokens
  ) external view returns (TokenConfig[] memory configs) {
    configs = new TokenConfig[](localTokens.length);
    for (uint256 i = 0; i < localTokens.length; ++i) {
      configs[i] = s_tokenConfig;
    }
    return configs;
  }

  function getAllConfiguredTokens(
    uint64,
    uint64
  ) external pure returns (address[] memory tokens) {
    return tokens;
  }

  function getTokenAdminRegistry() external pure returns (address) {
    return address(0);
  }

  /// @dev The pool accepts fee withdrawals only from its pool config manager.
  function withdrawFeeTokens(
    IMultiTokenPool pool,
    address[] calldata feeTokens,
    address recipient
  ) external {
    pool.withdrawFeeTokens(feeTokens, recipient);
  }
}
