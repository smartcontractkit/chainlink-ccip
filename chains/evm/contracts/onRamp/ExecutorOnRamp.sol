// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IExecutorOnRamp} from "../interfaces/IExecutorOnRamp.sol";

import {Client} from "../libraries/Client.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice The ExecutorOnRamp configures executor fees, destination support, and CCV limits.
contract ExecutorOnRamp is Ownable2StepMsgSender, IExecutorOnRamp {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.UintSet;

  error ExceedsMaxPossibleCCVs(uint256 provided, uint256 max);
  error ExceedsMaxRequiredCCVs(uint256 provided, uint256 max);
  error InvalidCCV(address ccv);
  error InvalidConfig();
  error InvalidDestChain(uint64 destChainSelector);
  error InvalidExtraArgsVersion(bytes4 provided);

  event ConfigSet(DynamicConfig dynamicConfig);
  event DestChainAdded(uint64 indexed destChainSelector);
  event DestChainRemoved(uint64 indexed destChainSelector);
  event CCVAdded(address indexed ccv);
  event CCVRemoved(address indexed ccv);

  /// @dev Struct that defines the dynamic configuration.
  // solhint-disable-next-line gas-struct-packing
  struct DynamicConfig {
    // Address that quotes a fee for each message
    address feeQuoter;
    // Address that receives fees, regardless of who calls withdraw
    address feeAggregator;
    // Max(required ccvs + optional ccvs)
    // Limits the number of ccvs that the executor needs to search for results from
    uint8 maxPossibleCCVsPerMsg;
    // Max(required ccvs + optional ccv threshold)
    // Limits the number of ccvs that the executor needs to submit on-chain
    uint8 maxRequiredCCVsPerMsg;
  }

  /// @notice Whether or not the CCV allow list is enabled.
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
    if (
      dynamicConfig.feeQuoter == address(0) || dynamicConfig.feeAggregator == address(0)
        || dynamicConfig.maxRequiredCCVsPerMsg > dynamicConfig.maxPossibleCCVsPerMsg
    ) revert InvalidConfig();

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
  /// @return ccvs The list of verifier addresses.
  function getAllowedCCVs() external view returns (address[] memory) {
    uint256 length = s_allowedCCVs.length();
    address[] memory verifiers = new address[](length);
    for (uint256 i = 0; i < length; ++i) {
      verifiers[i] = s_allowedCCVs.at(i);
    }
    return verifiers;
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
  }

  // ================================================================
  // │                             Fees                             │
  // ================================================================

  // TODO: Needs fee token withdrawal function once fees are implemented

  /// @notice Validates whether or not the executor can process the message and returns the fee required to do so.
  /// @param destChainSelector The destination chain selector.
  /// @param extraArgs The extra arguments for the message execution.
  /// @return fee The fee required to execute the message.
  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage memory, // message
    bytes calldata extraArgs
  ) external view returns (uint256) {
    if (!s_allowedDestChains.contains(destChainSelector)) {
      revert InvalidDestChain(destChainSelector);
    }

    if (bytes4(extraArgs[0:4]) != Client.GENERIC_EXTRA_ARGS_V3_TAG) {
      revert InvalidExtraArgsVersion(bytes4(extraArgs[0:4]));
    }

    Client.EVMExtraArgsV3 memory resolvedArgs = abi.decode(extraArgs[4:], (Client.EVMExtraArgsV3));

    if (s_allowlistEnabled) {
      for (uint256 i = 0; i < resolvedArgs.requiredCCV.length; i++) {
        address ccvAddress = resolvedArgs.requiredCCV[i].ccvAddress;
        if (!s_allowedCCVs.contains(ccvAddress)) {
          revert InvalidCCV(ccvAddress);
        }
      }

      // TODO: Should we instead just check until we reach the optionalThreshold?
      for (uint256 i = 0; i < resolvedArgs.optionalCCV.length; i++) {
        address ccvAddress = resolvedArgs.optionalCCV[i].ccvAddress;
        if (!s_allowedCCVs.contains(ccvAddress)) {
          revert InvalidCCV(ccvAddress);
        }
      }
    }

    uint256 possibleVerifiers = resolvedArgs.requiredCCV.length + resolvedArgs.optionalCCV.length;
    if (possibleVerifiers > s_dynamicConfig.maxPossibleCCVsPerMsg) {
      revert ExceedsMaxPossibleCCVs(possibleVerifiers, s_dynamicConfig.maxPossibleCCVsPerMsg);
    }

    uint256 requiredVerifiers = resolvedArgs.requiredCCV.length + resolvedArgs.optionalThreshold;
    if (requiredVerifiers > s_dynamicConfig.maxRequiredCCVsPerMsg) {
      revert ExceedsMaxRequiredCCVs(requiredVerifiers, s_dynamicConfig.maxRequiredCCVsPerMsg);
    }

    // TODO: get execution fee using extraArgs, for now we just return 0
    return 0;
  }
}
