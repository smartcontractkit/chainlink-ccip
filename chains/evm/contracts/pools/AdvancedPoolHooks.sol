// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../interfaces/IAdvancedPoolHooks.sol";
import {IPoolV2} from "../interfaces/IPoolV2.sol";

import {CCVConfigValidation} from "../libraries/CCVConfigValidation.sol";
import {Pool} from "../libraries/Pool.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Advanced pool hooks for additional security features like allowlists and CCV management.
/// @dev This is a standalone contract that can optionally be used by TokenPools.
contract AdvancedPoolHooks is IAdvancedPoolHooks, Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;

  error AllowListNotEnabled();
  error SenderNotAllowed(address sender);

  event AllowListAdd(address sender);
  event AllowListRemove(address sender);
  event CCVConfigUpdated(
    uint64 indexed remoteChainSelector,
    address[] outboundCCVs,
    address[] outboundCCVsToAddAboveThreshold,
    address[] inboundCCVs,
    address[] inboundCCVsToAddAboveThreshold
  );
  event ThresholdAmountSet(uint256 thresholdAmount);

  struct CCVConfig {
    address[] outboundCCVs; // CCVs required for outgoing messages to the remote chain.
    address[] outboundCCVsToAddAboveThreshold; // Additional CCVs that are required for outgoing messages above threshold to the remote chain.
    address[] inboundCCVs; // CCVs required for incoming messages from the remote chain.
    address[] inboundCCVsToAddAboveThreshold; // Additional CCVs that are required for incoming messages above threshold from the remote chain.
  }

  struct CCVConfigArg {
    uint64 remoteChainSelector;
    address[] outboundCCVs;
    address[] outboundCCVsToAddAboveThreshold;
    address[] inboundCCVs;
    address[] inboundCCVsToAddAboveThreshold;
  }

  /// @dev The immutable flag that indicates if the allowlist is access-controlled.
  bool internal immutable i_allowlistEnabled;

  /// @dev A set of addresses allowed to trigger lockOrBurn as original senders.
  /// Only takes effect if i_allowlistEnabled is true.
  /// This can be used to ensure only token-issuer specified addresses can move tokens.
  EnumerableSet.AddressSet internal s_allowlist;

  /// @dev Threshold token transfer amount above which additional CCVs are required.
  /// Value of 0 means that there is no threshold and additional CCVs are not required for any transfer amount.
  uint256 internal s_thresholdAmountForAdditionalCCVs;

  /// @dev Stores verifier (CCV) requirements keyed by remote chain selector.
  mapping(uint64 remoteChainSelector => CCVConfig ccvConfig) internal s_verifierConfig;

  constructor(address[] memory allowlist, uint256 thresholdAmountForAdditionalCCVs) {
    // Allowlist can be set as enabled or disabled at deployment time only to save hot-path gas.
    i_allowlistEnabled = allowlist.length > 0;
    if (i_allowlistEnabled) {
      _applyAllowListUpdates(new address[](0), allowlist);
    }
    s_thresholdAmountForAdditionalCCVs = thresholdAmountForAdditionalCCVs;
  }

  /// @inheritdoc IAdvancedPoolHooks
  function preflightCheck(Pool.LockOrBurnInV1 calldata lockOrBurnIn, uint16, bytes calldata) external view {
    checkAllowList(lockOrBurnIn.originalSender);
  }

  function postFlightCheck(Pool.ReleaseOrMintInV1 calldata, uint256, uint16) external pure {}

  // ================================================================
  // │                          Allowlist                           │
  // ================================================================

  /// @notice Checks if the sender is allowed to perform an operation.
  /// @param sender The address to check.
  function checkAllowList(
    address sender
  ) public view {
    if (i_allowlistEnabled) {
      if (!s_allowlist.contains(sender)) {
        revert SenderNotAllowed(sender);
      }
    }
  }

  /// @notice Gets whether the allowlist functionality is enabled.
  /// @return true is enabled, false if not.
  function getAllowListEnabled() external view returns (bool) {
    return i_allowlistEnabled;
  }

  /// @notice Gets the allowed addresses.
  /// @return The allowed addresses.
  function getAllowList() external view returns (address[] memory) {
    return s_allowlist.values();
  }

  /// @notice Apply updates to the allow list.
  /// @param removes The addresses to be removed.
  /// @param adds The addresses to be added.
  function applyAllowListUpdates(address[] calldata removes, address[] calldata adds) external onlyOwner {
    _applyAllowListUpdates(removes, adds);
  }

  /// @notice Internal version of applyAllowListUpdates to allow for reuse in the constructor.
  function _applyAllowListUpdates(address[] memory removes, address[] memory adds) internal {
    if (!i_allowlistEnabled) revert AllowListNotEnabled();

    for (uint256 i = 0; i < removes.length; ++i) {
      address toRemove = removes[i];
      if (s_allowlist.remove(toRemove)) {
        emit AllowListRemove(toRemove);
      }
    }
    for (uint256 i = 0; i < adds.length; ++i) {
      address toAdd = adds[i];
      if (toAdd == address(0)) {
        continue;
      }
      if (s_allowlist.add(toAdd)) {
        emit AllowListAdd(toAdd);
      }
    }
  }

  // ================================================================
  // │                          CCV                                 │
  // ================================================================

  /// @notice Updates the CCV configuration for specified remote chains.
  /// If the array includes address(0), it indicates that the default CCV should be used alongside any other specified CCVs.
  /// @dev Additional CCVs should only be configured for transfers above the threshold amount and should not duplicate base CCVs.
  /// Base CCVs are always required, while add-above-threshold CCVs are only required when the transfer amount exceeds the threshold.
  function applyCCVConfigUpdates(
    CCVConfigArg[] calldata ccvConfigArgs
  ) external onlyOwner {
    for (uint256 i = 0; i < ccvConfigArgs.length; ++i) {
      uint64 remoteChainSelector = ccvConfigArgs[i].remoteChainSelector;
      address[] calldata outboundCCVs = ccvConfigArgs[i].outboundCCVs;
      address[] calldata outboundCCVsToAddAboveThreshold = ccvConfigArgs[i].outboundCCVsToAddAboveThreshold;
      address[] calldata inboundCCVs = ccvConfigArgs[i].inboundCCVs;
      address[] calldata inboundCCVsToAddAboveThreshold = ccvConfigArgs[i].inboundCCVsToAddAboveThreshold;

      // check for duplicates in outbound CCVs.
      CCVConfigValidation._assertNoDuplicates(outboundCCVs);
      CCVConfigValidation._assertNoDuplicates(outboundCCVsToAddAboveThreshold);

      // check for duplicates in inbound CCVs.
      CCVConfigValidation._assertNoDuplicates(inboundCCVs);
      CCVConfigValidation._assertNoDuplicates(inboundCCVsToAddAboveThreshold);

      s_verifierConfig[remoteChainSelector] = CCVConfig({
        outboundCCVs: outboundCCVs,
        outboundCCVsToAddAboveThreshold: outboundCCVsToAddAboveThreshold,
        inboundCCVs: inboundCCVs,
        inboundCCVsToAddAboveThreshold: inboundCCVsToAddAboveThreshold
      });
      emit CCVConfigUpdated({
        remoteChainSelector: remoteChainSelector,
        outboundCCVs: outboundCCVs,
        outboundCCVsToAddAboveThreshold: outboundCCVsToAddAboveThreshold,
        inboundCCVs: inboundCCVs,
        inboundCCVsToAddAboveThreshold: inboundCCVsToAddAboveThreshold
      });
    }
  }

  /// @notice Returns the set of required CCVs for transfers in a specific direction.
  /// @param remoteChainSelector The remote chain selector for this transfer.
  /// @param amount The amount being transferred.
  /// @param direction The direction of the transfer (Inbound or Outbound).
  /// This implementation returns base CCVs for all transfers, and includes additional CCVs when the transfer amount.
  /// is above the configured threshold.
  /// @return requiredCCVs Set of required CCV addresses.
  function getRequiredCCVs(
    address,
    uint64 remoteChainSelector,
    uint256 amount,
    uint16,
    bytes calldata,
    IPoolV2.MessageDirection direction
  ) external view returns (address[] memory requiredCCVs) {
    CCVConfig storage config = s_verifierConfig[remoteChainSelector];
    if (direction == IPoolV2.MessageDirection.Inbound) {
      return _resolveRequiredCCVs(config.inboundCCVs, config.inboundCCVsToAddAboveThreshold, amount);
    }
    return _resolveRequiredCCVs(config.outboundCCVs, config.outboundCCVsToAddAboveThreshold, amount);
  }

  /// @notice Gets the threshold amount above which additional CCVs are required.
  /// @return The threshold amount.
  function getThresholdAmount() external view returns (uint256) {
    return s_thresholdAmountForAdditionalCCVs;
  }

  /// @notice Sets the threshold amount above which additional CCVs are required.
  /// @param thresholdAmount The new threshold amount.
  function setThresholdAmount(
    uint256 thresholdAmount
  ) external onlyOwner {
    s_thresholdAmountForAdditionalCCVs = thresholdAmount;

    emit ThresholdAmountSet(thresholdAmount);
  }

  function _resolveRequiredCCVs(
    address[] memory baseCCVs,
    address[] storage requiredCCVsAboveThresholdStorage,
    uint256 amount
  ) internal view returns (address[] memory requiredCCVs) {
    // If amount is above threshold, combine base and additional CCVs.
    uint256 thresholdAmount = s_thresholdAmountForAdditionalCCVs;
    if (thresholdAmount != 0 && amount >= thresholdAmount) {
      address[] memory thresholdCCVs = requiredCCVsAboveThresholdStorage;
      if (thresholdCCVs.length > 0) {
        requiredCCVs = new address[](baseCCVs.length + thresholdCCVs.length);
        // Copy base CCVs.
        for (uint256 i = 0; i < baseCCVs.length; ++i) {
          requiredCCVs[i] = baseCCVs[i];
        }
        // Copy additional CCVs.
        for (uint256 i = 0; i < thresholdCCVs.length; ++i) {
          requiredCCVs[baseCCVs.length + i] = thresholdCCVs[i];
        }
        return requiredCCVs;
      }
    }
    return baseCCVs;
  }
}
