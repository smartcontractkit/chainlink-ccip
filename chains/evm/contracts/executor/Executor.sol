// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IExecutor} from "../interfaces/IExecutor.sol";

import {Client} from "../libraries/Client.sol";
import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.0.2/utils/structs/EnumerableSet.sol";

/// @notice The Executor configures the supported destination chains and CCV limits for an executor.
contract Executor is IExecutor, Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.UintSet;
  using SafeERC20 for IERC20;

  error ExceedsMaxCCVs(uint256 provided, uint256 max);
  error InvalidCCV(address ccv);
  error InvalidDestChain(uint64 destChainSelector);
  error Executor__RequestedBlockDepthTooLow(uint16 requestedBlockDepth, uint16 minBlockConfirmations);
  error InvalidMaxPossibleCCVsPerMsg(uint256 maxPossibleCCVsPerMsg);
  error InvalidConfig();

  event CCVAllowlistUpdated(bool enabled);
  event CCVAdded(address indexed ccv);
  event CCVRemoved(address indexed ccv);
  event DestChainAdded(uint64 indexed destChainSelector, RemoteChainConfig config);
  event DestChainRemoved(uint64 indexed destChainSelector);
  event FeeTokenWithdrawn(address indexed receiver, address indexed feeToken, uint256 amount);
  event ConfigSet(DynamicConfig dynamicConfig);

  struct RemoteChainConfig {
    uint16 usdCentsFee; // ──────────╮ The fee charged by the executor for processing messages to this chain, USD cents.
    uint32 baseExecGas; //           │ The base gas cost to execute messages, excluding pool/CCV/receiver gas.
    uint8 destAddressLengthBytes; // │ The length of addresses on the destination chain, in bytes.
    bool enabled; // ────────────────╯ Whether or not this destination chain is enabled.
  }

  struct RemoteChainConfigArgs {
    uint64 destChainSelector;
    RemoteChainConfig config;
  }

  struct DynamicConfig {
    address feeAggregator; // ───────╮ Address to send withdrawn fees to.
    uint16 minBlockConfirmations; // │ Minimum number of block confirmations allowed (0 = finality).
    bool ccvAllowlistEnabled; // ────╯ Whether the CCV allowlist is enabled.
  }

  string public constant typeAndVersion = "Executor 1.7.0-dev";

  /// @notice Limits the number of CCVs that the executor needs to search for results from.
  /// @dev Max(required CCVs + optional CCVs).
  uint8 internal immutable i_maxCCVsPerMsg;
  /// @notice Dynamic configuration.
  DynamicConfig internal s_dynamicConfig;
  /// @notice The set of CCVs that the executor supports.
  EnumerableSet.AddressSet internal s_allowedCCVs;
  /// @notice The set of destination chains that the executor supports.
  EnumerableSet.UintSet internal s_allowedDestChains;
  /// @notice The remote chain configurations for supported destination chains.
  mapping(uint64 => RemoteChainConfig) internal s_remoteChainConfigs;

  constructor(uint8 maxCCVsPerMsg, DynamicConfig memory dynamicConfig) {
    if (maxCCVsPerMsg == 0) {
      revert InvalidMaxPossibleCCVsPerMsg(maxCCVsPerMsg);
    }
    i_maxCCVsPerMsg = maxCCVsPerMsg;
    _setDynamicConfig(dynamicConfig);
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Gets the maximum number of CCVs that can be used in a single message.
  /// @return maxCCVsPerMsg The maximum number of CCVs.
  function getMaxCCVsPerMessage() external view virtual returns (uint8 maxCCVsPerMsg) {
    return i_maxCCVsPerMsg;
  }

  /// @notice Returns the list of destination chains that the executor supports.
  /// @return destChains The list of destination chain selectors.
  function getDestChains() external view virtual returns (RemoteChainConfigArgs[] memory) {
    uint256 numberOfChains = s_allowedDestChains.length();
    RemoteChainConfigArgs[] memory destChains = new RemoteChainConfigArgs[](numberOfChains);
    for (uint256 i = 0; i < numberOfChains; ++i) {
      destChains[i].destChainSelector = uint64(s_allowedDestChains.at(i));
      destChains[i].config = s_remoteChainConfigs[destChains[i].destChainSelector];
    }
    return destChains;
  }

  /// @notice Updates the destination chains that the executor supports.
  /// @param destChainSelectorsToRemove The destination chain selectors to remove.
  /// @param destChainSelectorsToAdd The destination chain selectors to add.
  function applyDestChainUpdates(
    uint64[] calldata destChainSelectorsToRemove,
    RemoteChainConfigArgs[] calldata destChainSelectorsToAdd
  ) external virtual onlyOwner {
    for (uint256 i = 0; i < destChainSelectorsToRemove.length; ++i) {
      uint64 destChainSelector = destChainSelectorsToRemove[i];
      if (s_allowedDestChains.remove(destChainSelector)) {
        delete s_remoteChainConfigs[destChainSelector];
        emit DestChainRemoved(destChainSelector);
      }
    }

    for (uint256 i = 0; i < destChainSelectorsToAdd.length; ++i) {
      RemoteChainConfigArgs calldata args = destChainSelectorsToAdd[i];
      if (args.destChainSelector == 0) {
        revert InvalidDestChain(args.destChainSelector);
      }

      s_allowedDestChains.add(args.destChainSelector);
      s_remoteChainConfigs[args.destChainSelector] = args.config;
      emit DestChainAdded(args.destChainSelector, args.config);
    }
  }

  /// @notice Returns the dynamic configuration.
  /// @return dynamicConfig The dynamic configuration.
  function getDynamicConfig() external view virtual returns (DynamicConfig memory) {
    return s_dynamicConfig;
  }

  /// @notice Sets the dynamic configuration.
  /// @param dynamicConfig The dynamic configuration.
  function setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) external virtual onlyOwner {
    _setDynamicConfig(dynamicConfig);
  }

  /// @notice Internal function to set the dynamic configuration.
  /// @param dynamicConfig The dynamic configuration.
  function _setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) internal {
    if (dynamicConfig.feeAggregator == address(0)) {
      revert InvalidConfig();
    }
    // Zero is a valid value for minBlockConfirmations, indicating that finality is requested.
    s_dynamicConfig = dynamicConfig;

    emit ConfigSet(dynamicConfig);
  }

  /// @inheritdoc IExecutor
  function getMinBlockConfirmations() external view virtual returns (uint16) {
    return s_dynamicConfig.minBlockConfirmations;
  }

  /// @notice Returns the list of CCVs that the executor supports.
  /// @return ccvs The list of CCV addresses.
  function getAllowedCCVs() external view virtual returns (address[] memory) {
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
  ) external virtual onlyOwner {
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

    if (s_dynamicConfig.ccvAllowlistEnabled != ccvAllowlistEnabled) {
      s_dynamicConfig.ccvAllowlistEnabled = ccvAllowlistEnabled;
      emit CCVAllowlistUpdated(ccvAllowlistEnabled);
    }
  }

  // ================================================================
  // │                             Fees                             │
  // ================================================================

  /// @notice Validates whether or not the executor can process the message and returns the fee required to do so.
  /// @param destChainSelector The destination chain selector.
  /// @param requestedBlockDepth The requested block depth for the message. `0` indicates waiting for finality.
  /// @param dataLength The length of the message data in bytes.
  /// @param numberOfTokens The number of tokens being transferred in the message.
  /// @param ccvs The CCVs that are requested on source.
  /// @return usdCentsFee The USD denominated fee for the executor.
  /// @return execGasCost The gas required to execute the message on destination, excluding pool/CCV/receiver gas.
  /// @return execBytes The byte overhead required to execute the message on destination, excluding pool/CCV bytes.
  function getFee(
    uint64 destChainSelector,
    uint16 requestedBlockDepth,
    uint32 dataLength,
    uint8 numberOfTokens,
    Client.CCV[] calldata ccvs,
    bytes calldata extraArgs
  ) external view virtual returns (uint16 usdCentsFee, uint32 execGasCost, uint32 execBytes) {
    RemoteChainConfig memory remoteChainConfig = s_remoteChainConfigs[destChainSelector];
    if (!remoteChainConfig.enabled) {
      revert InvalidDestChain(destChainSelector);
    }
    if (requestedBlockDepth != 0 && requestedBlockDepth < s_dynamicConfig.minBlockConfirmations) {
      revert Executor__RequestedBlockDepthTooLow(requestedBlockDepth, s_dynamicConfig.minBlockConfirmations);
    }

    if (s_dynamicConfig.ccvAllowlistEnabled) {
      for (uint256 i = 0; i < ccvs.length; ++i) {
        address ccvAddress = ccvs[i].ccvAddress;
        if (!s_allowedCCVs.contains(ccvAddress)) {
          revert InvalidCCV(ccvAddress);
        }
      }
    }

    if (ccvs.length > i_maxCCVsPerMsg) {
      revert ExceedsMaxCCVs(ccvs.length, i_maxCCVsPerMsg);
    }

    // Since the message payload is the same on source and destination chains with the V1 codec, we can use the
    // same calculation for execBytes on destination.
    execBytes = uint32(
      MessageV1Codec.MESSAGE_V1_EVM_SOURCE_BASE_SIZE + dataLength + extraArgs.length
        + (MessageV1Codec.MESSAGE_V1_REMOTE_CHAIN_ADDRESSES * remoteChainConfig.destAddressLengthBytes)
        + (
          numberOfTokens
            * (MessageV1Codec.TOKEN_TRANSFER_V1_EVM_SOURCE_BASE_SIZE + remoteChainConfig.destAddressLengthBytes)
        )
    );

    execGasCost += remoteChainConfig.baseExecGas;

    return (remoteChainConfig.usdCentsFee, execGasCost, execBytes);
  }

  /// @notice Withdraws the outstanding fee token balances to the fee aggregator.
  /// @param feeTokens The fee tokens to withdraw.
  function withdrawFeeTokens(
    address[] calldata feeTokens
  ) external virtual onlyOwner {
    address feeAggregator = s_dynamicConfig.feeAggregator;
    for (uint256 i = 0; i < feeTokens.length; ++i) {
      IERC20 feeToken = IERC20(feeTokens[i]);
      uint256 feeTokenBalance = feeToken.balanceOf(address(this));

      if (feeTokenBalance > 0) {
        feeToken.safeTransfer(feeAggregator, feeTokenBalance);

        emit FeeTokenWithdrawn(feeAggregator, address(feeToken), feeTokenBalance);
      }
    }
  }
}
