// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IExecutorOnRamp} from "../interfaces/IExecutorOnRamp.sol";

import {Client} from "../libraries/Client.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.0.2/utils/structs/EnumerableSet.sol";

/// @notice The ExecutorOnRamp configures the supported destination chains and CCV limits for an executor.
contract ExecutorOnRamp is Ownable2StepMsgSender, IExecutorOnRamp {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.UintSet;

  error ExceedsMaxCCVs(uint256 provided, uint256 max);
  error InvalidCCV(address ccv);
  error InvalidDestChain(uint64 destChainSelector);
  error InvalidMaxPossibleCCVsPerMsg(uint256 maxPossibleCCVsPerMsg);

  event CCVAllowlistUpdated(bool enabled);
  event CCVAdded(address indexed ccv);
  event CCVRemoved(address indexed ccv);
  event DestChainAdded(uint64 indexed destChainSelector);
  event DestChainRemoved(uint64 indexed destChainSelector);
  event MaxCCVsPerMsgSet(uint8 maxCCVsPerMsg);

  /// @notice Limits the number of CCVs that the executor needs to search for results from.
  /// @dev Max(required CCVs + optional CCVs).
  uint8 private s_maxCCVsPerMsg;
  /// @notice Whether or not the CCV allowlist is enabled.
  bool private s_ccvAllowlistEnabled;
  /// @notice The set of CCVs that the executor supports.
  /// @dev Addresses correspond to the CCVOnRamp for each CCV.
  EnumerableSet.AddressSet private s_allowedCCVs;
  /// @notice The set of destination chains that the executor supports.
  EnumerableSet.UintSet private s_allowedDestChains;

  string public constant typeAndVersion = "ExecutorOnRamp 1.7.0-dev";

  constructor(
    uint8 maxCCVsPerMsg
  ) {
    _setMaxCCVsPerMsg(maxCCVsPerMsg);
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Gets the maximum number of CCVs that can be used in a single message.
  /// @return maxCCVsPerMsg The maximum number of CCVs.
  function getMaxCCVsPerMsg() external view returns (uint8) {
    return s_maxCCVsPerMsg;
  }

  /// @notice Sets the maximum number of CCVs that can be used in a single message.
  /// @param maxCCVsPerMsg The maximum number of CCVs.
  function setMaxCCVsPerMsg(
    uint8 maxCCVsPerMsg
  ) external onlyOwner {
    _setMaxCCVsPerMsg(maxCCVsPerMsg);
  }

  /// @notice Internal version of setMaxCCVsPerMsg to allow for reuse in the constructor.
  function _setMaxCCVsPerMsg(
    uint8 maxCCVsPerMsg
  ) internal {
    if (maxCCVsPerMsg == 0) {
      revert InvalidMaxPossibleCCVsPerMsg(maxCCVsPerMsg);
    }

    s_maxCCVsPerMsg = maxCCVsPerMsg;

    emit MaxCCVsPerMsgSet(maxCCVsPerMsg);
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
  /// @param destChainSelectorsToRemove The destination chain selectors to remove.
  /// @param destChainSelectorsToAdd The destination chain selectors to add.
  function applyDestChainUpdates(
    uint64[] calldata destChainSelectorsToRemove,
    uint64[] calldata destChainSelectorsToAdd
  ) external onlyOwner {
    for (uint256 i = 0; i < destChainSelectorsToRemove.length; ++i) {
      uint64 destChainSelector = destChainSelectorsToRemove[i];
      if (s_allowedDestChains.remove(destChainSelector)) {
        emit DestChainRemoved(destChainSelector);
      }
    }

    for (uint256 i = 0; i < destChainSelectorsToAdd.length; ++i) {
      uint64 destChainSelector = destChainSelectorsToAdd[i];
      if (destChainSelector == 0) {
        revert InvalidDestChain(destChainSelector);
      }
      if (s_allowedDestChains.add(destChainSelector)) {
        emit DestChainAdded(destChainSelector);
      }
    }
  }

  /// @notice Returns whether or not the CCV allowlist is enabled.
  /// @return enabled The enablement status.
  function isCCVAllowlistEnabled() external view returns (bool) {
    return s_ccvAllowlistEnabled;
  }

  /// @notice Returns the list of CCVs that the executor supports.
  /// @return ccvs The list of CCV addresses.
  function getAllowedCCVs() external view returns (address[] memory) {
    return s_allowedCCVs.values();
  }

  /// @notice Updates CCV allowlist contents and enablement status.
  /// @param ccvsToRemove The CCV addresses to remove.
  /// @param ccvsToAdd The CCV addresses to add.
  /// @param ccvAllowlistEnabled Whether or not the allowlist should be enabled.
  function applyAllowedCCVUpdates(
    address[] calldata ccvsToRemove,
    address[] calldata ccvsToAdd,
    bool ccvAllowlistEnabled
  ) external onlyOwner {
    for (uint256 i = 0; i < ccvsToRemove.length; ++i) {
      if (s_allowedCCVs.remove(ccvsToRemove[i])) {
        emit CCVRemoved(ccvsToRemove[i]);
      }
    }

    for (uint256 i = 0; i < ccvsToAdd.length; ++i) {
      address ccv = ccvsToAdd[i];
      if (ccv == address(0)) {
        revert InvalidCCV(ccv);
      }
      if (s_allowedCCVs.add(ccv)) {
        emit CCVAdded(ccv);
      }
    }

    if (s_ccvAllowlistEnabled != ccvAllowlistEnabled) {
      s_ccvAllowlistEnabled = ccvAllowlistEnabled;
      emit CCVAllowlistUpdated(ccvAllowlistEnabled);
    }
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

    if (s_ccvAllowlistEnabled) {
      for (uint256 i = 0; i < requiredCCVs.length; ++i) {
        address ccvAddress = requiredCCVs[i].ccvAddress;
        if (!s_allowedCCVs.contains(ccvAddress)) {
          revert InvalidCCV(ccvAddress);
        }
      }

      for (uint256 i = 0; i < optionalCCVs.length; ++i) {
        address ccvAddress = optionalCCVs[i].ccvAddress;
        if (!s_allowedCCVs.contains(ccvAddress)) {
          revert InvalidCCV(ccvAddress);
        }
      }
    }

    uint256 possibleCCVs = requiredCCVs.length + optionalCCVs.length;
    if (possibleCCVs > s_maxCCVsPerMsg) {
      revert ExceedsMaxCCVs(possibleCCVs, s_maxCCVsPerMsg);
    }

    // TODO: get execution fee, for now we just return 0
    return 0;
  }
}
