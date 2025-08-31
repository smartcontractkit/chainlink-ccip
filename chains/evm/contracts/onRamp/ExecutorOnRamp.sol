// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IExecutorOnRamp} from "../interfaces/IExecutorOnRamp.sol";

import {Client} from "../libraries/Client.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice The ExecutorOnRamp configures the supported destination chains and CCV limits for an executor.
contract ExecutorOnRamp is Ownable2StepMsgSender, IExecutorOnRamp {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.UintSet;

  error ExceedsMaxCCVs(uint256 provided, uint256 max);
  error InvalidCCV(address ccv);
  error InvalidDestChain(uint64 destChainSelector);
  error InvalidMaxPossibleCCVsPerMsg(uint256 maxPossibleCCVsPerMsg);

  event AllowlistUpdated(bool enabled);
  event CCVAdded(address indexed ccv);
  event CCVRemoved(address indexed ccv);
  event ConfigSet(DynamicConfig dynamicConfig);
  event DestChainAdded(uint64 indexed destChainSelector);
  event DestChainRemoved(uint64 indexed destChainSelector);

  /// @dev Struct that defines the dynamic configuration.
  // solhint-disable-next-line gas-struct-packing
  struct DynamicConfig {
    // Max(required ccvs + optional ccvs).
    // Limits the number of ccvs that the executor needs to search for results from.
    uint8 maxCCVsPerMsg;
  }

  /// @notice Whether or not the CCV allowlist is enabled.
  bool public s_allowlistEnabled;
  /// @notice The set of CCVs that the executor supports.
  /// @dev Addresses correspond to the CCVOnRamp for each CCV.
  EnumerableSet.AddressSet private s_allowedCCVs;
  /// @notice The set of destination chains that the executor supports.
  EnumerableSet.UintSet private s_allowedDestChains;
  /// @dev The dynamic config.
  DynamicConfig private s_dynamicConfig;

  string public constant typeAndVersion = "ExecutorOnRamp 1.7.0-dev";

  constructor(
    DynamicConfig memory dynamicConfig
  ) {
    _setDynamicConfig(dynamicConfig);
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Returns the dynamic config.
  /// @return dynamicConfig The config.
  function getDynamicConfig() external view returns (DynamicConfig memory dynamicConfig) {
    return s_dynamicConfig;
  }

  /// @notice Sets the dynamic config.
  /// @param dynamicConfig The config.
  function setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) external onlyOwner {
    _setDynamicConfig(dynamicConfig);
  }

  /// @notice Internal version of setDynamicConfig to allow for reuse in the constructor.
  function _setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) internal {
    if (dynamicConfig.maxCCVsPerMsg == 0) {
      revert InvalidMaxPossibleCCVsPerMsg(dynamicConfig.maxCCVsPerMsg);
    }

    s_dynamicConfig = dynamicConfig;

    emit ConfigSet(dynamicConfig);
  }

  /// @notice Returns the list of destination chains that the executor supports.
  /// @return destChains The list of destination chain selectors.
  function getDestChains() external view returns (uint64[] memory) {
    uint256 length = s_allowedDestChains.length();
    uint64[] memory destChains = new uint64[](length);
    for (uint256 i = 0; i < length; ++i) {
      destChains[i] = uint64(s_allowedDestChains.at(i));
    }
    return destChains;
  }

  /// @notice Updates the destination chains that the executor supports.
  /// @param destChainSelectorsToAdd The destination chain selectors to add.
  /// @param destChainSelectorsToRemove The destination chain selectors to remove.
  function applyDestChainUpdates(
    uint64[] calldata destChainSelectorsToAdd,
    uint64[] calldata destChainSelectorsToRemove
  ) external onlyOwner {
    for (uint256 i = 0; i < destChainSelectorsToAdd.length; ++i) {
      uint64 destChainSelector = destChainSelectorsToAdd[i];
      if (destChainSelector == 0) {
        revert InvalidDestChain(destChainSelector);
      }
      if (s_allowedDestChains.add(destChainSelector)) {
        emit DestChainAdded(destChainSelector);
      }
    }

    for (uint256 i = 0; i < destChainSelectorsToRemove.length; ++i) {
      uint64 destChainSelector = destChainSelectorsToRemove[i];
      if (s_allowedDestChains.remove(destChainSelector)) {
        emit DestChainRemoved(destChainSelector);
      }
    }
  }

  /// @notice Returns the list of CCVs that the executor supports.
  /// @return ccvs The list of ccv addresses.
  function getAllowedCCVs() external view returns (address[] memory) {
    return s_allowedCCVs.values();
  }

  /// @notice Updates CCV allowlist contents and enablement status.
  /// @param ccvsToAdd The CCV addresses to add.
  /// @param ccvsToRemove The CCV addresses to remove.
  /// @param allowlistEnabled Whether or not the allowlist is enabled.
  function applyAllowedCCVUpdates(
    address[] calldata ccvsToAdd,
    address[] calldata ccvsToRemove,
    bool allowlistEnabled
  ) external onlyOwner {
    for (uint256 i = 0; i < ccvsToAdd.length; ++i) {
      address ccv = ccvsToAdd[i];
      if (ccv == address(0)) {
        revert InvalidCCV(ccv);
      }
      if (s_allowedCCVs.add(ccv)) {
        emit CCVAdded(ccv);
      }
    }

    for (uint256 i = 0; i < ccvsToRemove.length; ++i) {
      if (s_allowedCCVs.remove(ccvsToRemove[i])) {
        emit CCVRemoved(ccvsToRemove[i]);
      }
    }

    s_allowlistEnabled = allowlistEnabled;
    emit AllowlistUpdated(allowlistEnabled);
  }

  // ================================================================
  // │                             Fees                             │
  // ================================================================

  // TODO: Needs fee token withdrawal function once fees are implemented

  /// @notice Validates whether or not the executor can process the message and returns the fee required to do so.
  /// @param destChainSelector The destination chain selector.
  /// @param requiredCCVs The CCVs that are required to execute the message.
  /// @param optionalCCVs The CCVs that can optionally be used to execute the message
  /// @return fee The fee required to execute the message.
  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata, // message
    Client.CCV[] calldata requiredCCVs,
    Client.CCV[] calldata optionalCCVs,
    bytes calldata // extraArgs
  ) external view returns (uint256) {
    if (!s_allowedDestChains.contains(destChainSelector)) {
      revert InvalidDestChain(destChainSelector);
    }

    if (s_allowlistEnabled) {
      for (uint256 i = 0; i < requiredCCVs.length; i++) {
        address ccvAddress = requiredCCVs[i].ccvAddress;
        if (!s_allowedCCVs.contains(ccvAddress)) {
          revert InvalidCCV(ccvAddress);
        }
      }

      for (uint256 i = 0; i < optionalCCVs.length; i++) {
        address ccvAddress = optionalCCVs[i].ccvAddress;
        if (!s_allowedCCVs.contains(ccvAddress)) {
          revert InvalidCCV(ccvAddress);
        }
      }
    }

    uint256 possibleCCVs = requiredCCVs.length + optionalCCVs.length;
    if (possibleCCVs > s_dynamicConfig.maxCCVsPerMsg) {
      revert ExceedsMaxCCVs(possibleCCVs, s_dynamicConfig.maxCCVsPerMsg);
    }

    // TODO: get execution fee, for now we just return 0
    return 0;
  }
}
