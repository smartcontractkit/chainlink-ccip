// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../interfaces/IAny2EVMMessageReceiver.sol";
import {IMessageInterceptor} from "../interfaces/IMessageInterceptor.sol";
import {INonceManager} from "../interfaces/INonceManager.sol";
import {IPoolV1} from "../interfaces/IPool.sol";

import {IRMNRemote} from "../interfaces/IRMNRemote.sol";
import {IRouter} from "../interfaces/IRouter.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";

import {Client} from "../libraries/Client.sol";
import {ERC165CheckerReverting} from "../libraries/ERC165CheckerReverting.sol";
import {Internal} from "../libraries/Internal.sol";
import {Pool} from "../libraries/Pool.sol";
import {SuperchainInterop} from "../libraries/SuperchainInterop.sol";

import {ICrossL2Inbox} from "../vendor/optimism/interop-lib/v0/src/interfaces/ICrossL2Inbox.sol";
import {Identifier} from "../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";

import {OffRamp} from "./OffRamp.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/token/ERC20/IERC20.sol";
import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice This OffRamp supports OP Superchain Interop. It leverages CrossL2Inbox to verify validity of a message,
/// as opposed to relying on roots signed by oracles.
/// @dev This OffRamp does not expect transmission to come from DONs, therefore it does not inherite MultiOCR3Base
contract OffRampOverSuperchainInterop is OffRamp {
  using ERC165CheckerReverting for address;
  using EnumerableSet for EnumerableSet.UintSet;
  using EnumerableSet for EnumerableSet.AddressSet;

  error MessageNotVerified(uint64 sourceChainSelector, bytes32 messageId);
  error InvalidDestChainSelector(uint64 destChainSelector);
  error InvalidSourceOnRamp(address sourceOnRamp);
  error InvalidInteropLogSelector(bytes32 selector);
  error MismatchedDestChainSelector(uint64 expectedSelector, uint64 gotSelector);
  error MismatchedSequenceNumber(uint64 expectedSequenceNumber, uint64 gotSequenceNumber);
  error ZeroChainIdNotAllowed();
  error ChainIdNotConfiguredForSelector(uint64 sourceChainSelector);

  /// @dev Atlas depends on various events, if changing, please notify Atlas.

  event AllowedTransmitterAdded(address indexed transmitter);
  event AllowedTransmitterRemoved(address indexed transmitter);
  event ChainSelectorToChainIdConfigAdded(uint64 indexed chainSelector, uint256 indexed chainId);
  event ChainSelectorToChainIdConfigRemoved(uint64 indexed chainSelector, uint256 indexed chainId);

  struct ChainSelectorToChainIdConfigArgs {
    uint64 chainSelector;
    uint256 chainId;
  }

  /// @dev The CrossL2Inbox interface.
  ICrossL2Inbox internal immutable i_crossL2Inbox;

  // Not using enumerable map to make contract size smaller.
  mapping(uint64 sourceChainSelector => uint256 chainId) private s_sourceChainSelectorToChainId;

  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    SourceChainConfigArgs[] memory sourceChainConfigs,
    address crossL2Inbox,
    ChainSelectorToChainIdConfigArgs[] memory chainSelectorToChainIdConfigArgs
  ) OffRamp(staticConfig, dynamicConfig, sourceChainConfigs) {
    i_crossL2Inbox = ICrossL2Inbox(crossL2Inbox);

    _applyChainSelectorToChainIdConfigUpdates(new uint64[](0), chainSelectorToChainIdConfigArgs);
  }

  // ================================================================================================
  // │                       Code below are new in Superchain Interop OffRamp                       │
  // ================================================================================================

  // ================================================================
  // │        Execution - Most logic, e.g. message status           │
  // │        transitions, stay the same, added interop-specific    │
  // │        message parsing and validation                        │
  // ================================================================

  function decodeLogDataIntoMessage(
    bytes calldata logData
  ) internal view returns (Internal.Any2EVMRampMessage memory message) {
    // Validate Selector (also reverts if LOG0 with no topics)
    bytes32 selector = abi.decode(logData[:32], (bytes32));
    if (selector != SuperchainInterop.SENT_MESSAGE_LOG_SELECTOR) revert InvalidInteropLogSelector(selector);

    uint64 destChainSelector;
    uint64 sequenceNumber;
    (destChainSelector, sequenceNumber) = abi.decode(logData[32:96], (uint64, uint64));

    // Data
    message = abi.decode(logData[96:], (Internal.Any2EVMRampMessage));
    if (message.header.destChainSelector != i_chainSelector) {
      revert MismatchedDestChainSelector(destChainSelector, message.header.destChainSelector);
    }
    if (message.header.sequenceNumber != sequenceNumber) {
      revert MismatchedSequenceNumber(sequenceNumber, message.header.sequenceNumber);
    }

    return message;
  }

  function manuallyExecute(
    SuperchainInterop.ExecutionReport calldata report,
    GasLimitOverride memory gasLimitOverride
  ) external {
    uint256 newLimit = gasLimitOverride.receiverExecutionGasLimit;
    Internal.Any2EVMRampMessage memory message = decodeLogDataIntoMessage(report.logData);
    if (newLimit != 0) {
      // Checks to ensure messages will not be executed with less gas than specified.
      if (newLimit < message.gasLimit) {
        revert InvalidManualExecutionGasLimit(message.header.sourceChainSelector, message.header.messageId, newLimit);
      }
    }
    if (message.tokenAmounts.length != gasLimitOverride.tokenGasOverrides.length) {
      revert ManualExecutionGasAmountCountMismatch(message.header.messageId, message.header.sequenceNumber);
    }

    // The gas limit can not be lowered as that could cause the message to fail. If manual execution is done
    // from an UNTOUCHED state and we would allow lower gas limit, anyone could grief by executing the message with
    // lower gas limit than the DON would have used. This results in the message being marked FAILURE and the DON
    // would not attempt it with the correct gas limit.
    for (uint256 tokenIndex = 0; tokenIndex < message.tokenAmounts.length; ++tokenIndex) {
      uint256 tokenGasOverride = gasLimitOverride.tokenGasOverrides[tokenIndex];
      if (tokenGasOverride != 0) {
        uint256 destGasAmount = message.tokenAmounts[tokenIndex].destGasAmount;
        if (tokenGasOverride < destGasAmount) {
          revert InvalidManualExecutionTokenGasOverride(
            message.header.messageId, tokenIndex, destGasAmount, tokenGasOverride
          );
        }
      }
    }

    _executeSingleReport(report, message, gasLimitOverride, true);
  }

  function execute(
    SuperchainInterop.ExecutionReport calldata report
  ) external {
    GasLimitOverride memory gasLimitOverride;
    _executeSingleReport(report, decodeLogDataIntoMessage(report.logData), gasLimitOverride, false);
  }

  // Easier to not batch, given CrossL2Inbox.validateMessage can fail.
  function _executeSingleReport(
    SuperchainInterop.ExecutionReport calldata report,
    Internal.Any2EVMRampMessage memory message,
    GasLimitOverride memory gasLimitOverride,
    bool manualExecution
  ) internal {
    uint64 sourceChainSelector = message.header.sourceChainSelector;
    address onRampAddress = abi.decode(_getEnabledSourceChainConfig(sourceChainSelector).onRamp, (address));

    // Validate that the message is meant for this chain.
    if (message.header.destChainSelector != i_chainSelector) {
      revert InvalidDestChainSelector(message.header.destChainSelector);
    }
    // Validate that the message was emitted by the corresponding OnRamp.
    if (report.identifier.origin != onRampAddress) {
      revert InvalidSourceOnRamp(report.identifier.origin);
    }
    // Validate that the chainId maps to the expected sourceChainSelector
    {
      // Scope to reduce stack depth.
      uint256 expectedChainId = s_sourceChainSelectorToChainId[sourceChainSelector];
      if (expectedChainId == 0) {
        revert ChainIdNotConfiguredForSelector(sourceChainSelector);
      }
      if (expectedChainId != report.identifier.chainId) {
        revert SourceChainSelectorMismatch(sourceChainSelector, sourceChainSelector);
      }
    }

    // SECURITY CRITICAL CHECK.
    // Validate the exact log was emitted on the source chain.
    i_crossL2Inbox.validateMessage(report.identifier, keccak256(report.logData));

    // Main body of the single message execution logic.
    uint256 gasStart = gasleft();
    message = _beforeExecuteSingleMessage(message);

    Internal.MessageExecutionState originalState = getExecutionState(sourceChainSelector, message.header.sequenceNumber);
    // Two valid cases here, we either have never touched this message before, or we tried to execute and failed. This
    // check protects against reentry and re-execution because the other state is IN_PROGRESS which should not be
    // allowed to execute.
    if (
      !(
        originalState == Internal.MessageExecutionState.UNTOUCHED
          || originalState == Internal.MessageExecutionState.FAILURE
      )
    ) {
      // If the message has already been executed, we skip it. We want to not revert on race conditions between
      // executing parties. This will allow us to open up manual exec while also attempting with the DON, without
      // reverting an entire DON batch when a user manually executes while the tx is inflight.
      emit SkippedAlreadyExecutedMessage(sourceChainSelector, message.header.sequenceNumber);
      return;
    }
    uint32[] memory tokenGasOverrides;
    if (manualExecution) {
      tokenGasOverrides = gasLimitOverride.tokenGasOverrides;
      // For Superchain Interop, we allow manual execution if we previously failed
      // or if enough time has passed since the message was created
      bool isOldMessage =
        (block.timestamp - message.header.sequenceNumber) > s_dynamicConfig.permissionLessExecutionThresholdSeconds;
      // Manually execution is fine if we previously failed or if the message is old enough.
      // Acceptable state transitions: UNTOUCHED->SUCCESS, UNTOUCHED->FAILURE, FAILURE->SUCCESS.
      if (!(isOldMessage || originalState == Internal.MessageExecutionState.FAILURE)) {
        revert ManualExecutionNotYetEnabled(sourceChainSelector);
      }

      // Manual execution gas limit can override gas limit specified in the message. Value of 0 indicates no override.
      if (gasLimitOverride.receiverExecutionGasLimit != 0) {
        message.gasLimit = gasLimitOverride.receiverExecutionGasLimit;
      }
    } else {
      // Relayer can only execute a message once.
      // Acceptable state transitions: UNTOUCHED->SUCCESS, UNTOUCHED->FAILURE.
      if (originalState != Internal.MessageExecutionState.UNTOUCHED) {
        emit AlreadyAttempted(sourceChainSelector, message.header.sequenceNumber);
        return;
      }
    }

    // Nonce changes per state transition (these only apply for ordered messages):
    // UNTOUCHED -> FAILURE  nonce bump.
    // UNTOUCHED -> SUCCESS  nonce bump.
    // FAILURE   -> SUCCESS  no nonce bump.
    // UNTOUCHED messages MUST be executed in order always.
    // If nonce == 0 then out of order execution is allowed.
    if (message.header.nonce != 0) {
      if (originalState == Internal.MessageExecutionState.UNTOUCHED) {
        // If a nonce is not incremented, that means it was skipped, and we can ignore the message.
        if (
          !INonceManager(i_nonceManager).incrementInboundNonce(sourceChainSelector, message.header.nonce, message.sender)
        ) return;
      }
    }

    // We check when executing as a defense in depth measure.
    if (message.tokenAmounts.length != report.offchainTokenData.length) {
      revert TokenDataMismatch(sourceChainSelector, message.header.sequenceNumber);
    }

    _setExecutionState(sourceChainSelector, message.header.sequenceNumber, Internal.MessageExecutionState.IN_PROGRESS);
    (Internal.MessageExecutionState newState, bytes memory returnData) =
      _trialExecute(message, report.offchainTokenData, tokenGasOverrides);
    _setExecutionState(sourceChainSelector, message.header.sequenceNumber, newState);

    // Since it's hard to estimate whether manual execution will succeed, we revert the entire transaction if it
    // fails. This will show the user if their manual exec will fail before they submit it.
    if (manualExecution) {
      if (newState == Internal.MessageExecutionState.FAILURE) {
        if (originalState != Internal.MessageExecutionState.UNTOUCHED) {
          // If manual execution fails, we revert the entire transaction, unless the originalState is UNTOUCHED as we
          // would still be making progress by changing the state from UNTOUCHED to FAILURE.
          revert ExecutionError(message.header.messageId, returnData);
        }
      }
    }

    // The only valid prior states are UNTOUCHED and FAILURE (checked above).
    // The only valid post states are FAILURE and SUCCESS (checked below).
    if (newState != Internal.MessageExecutionState.SUCCESS) {
      if (newState != Internal.MessageExecutionState.FAILURE) {
        revert InvalidNewState(sourceChainSelector, message.header.sequenceNumber, newState);
      }
    }

    emit ExecutionStateChanged(
      sourceChainSelector,
      message.header.sequenceNumber,
      message.header.messageId,
      SuperchainInterop._hashInteropMessage(message, onRampAddress),
      newState,
      returnData,
      // This emit covers not only the execution through the router, but also all of the overhead in executing the
      // message. This gives the most accurate representation of the gas used in the execution.
      gasStart - gasleft()
    );
  }

  // ================================================================
  // │                Commit - Completely Removed                   │
  // ================================================================

  // ================================================================
  // │         Access - Ported transmitter check from OCR3Base,     │
  // │         added chainSelector to chainId resolution            │
  // ================================================================

  /// @notice Updates sourceChainSelector to chainId mapping.
  /// @param chainSelectorsToUnset Array of chainIds to remove from the mapping.
  /// @param chainSelectorsToSet Array of chainId to sourceChainSelector mappings to add.
  function applyChainSelectorToChainIdConfigUpdates(
    uint64[] memory chainSelectorsToUnset,
    ChainSelectorToChainIdConfigArgs[] memory chainSelectorsToSet
  ) external onlyOwner {
    _applyChainSelectorToChainIdConfigUpdates(chainSelectorsToUnset, chainSelectorsToSet);
  }

  /// @notice Internal function to update the chainId to sourceChainSelector mapping.
  /// @param chainSelectorsToUnset Array of chainIds to remove from the mapping.
  /// @param chainSelectorsToSet Array of chainId to sourceChainSelector mappings to add.
  function _applyChainSelectorToChainIdConfigUpdates(
    uint64[] memory chainSelectorsToUnset,
    ChainSelectorToChainIdConfigArgs[] memory chainSelectorsToSet
  ) internal {
    for (uint256 i = 0; i < chainSelectorsToUnset.length; ++i) {
      uint64 chainSelector = chainSelectorsToUnset[i];
      uint256 chainId = s_sourceChainSelectorToChainId[chainSelector];
      delete s_sourceChainSelectorToChainId[chainSelector];
      if (chainId != 0) {
        emit ChainSelectorToChainIdConfigRemoved(chainSelector, chainId);
      }
    }

    for (uint256 i = 0; i < chainSelectorsToSet.length; ++i) {
      ChainSelectorToChainIdConfigArgs memory config = chainSelectorsToSet[i];
      if (config.chainId == 0) {
        revert ZeroChainIdNotAllowed();
      }
      if (config.chainSelector == 0) {
        revert ZeroChainSelectorNotAllowed();
      }
      // Not validating if chainSelector is within s_sourceChainSelectors for this config to
      // be able to evolve quickly.
      s_sourceChainSelectorToChainId[config.chainSelector] = config.chainId;
      emit ChainSelectorToChainIdConfigAdded(config.chainSelector, config.chainId);
    }
  }

  /// @notice Gets the chainId for a given sourceChainSelector.
  /// @param sourceChainSelector The sourceChainSelector to look up.
  /// @return chainId The corresponding chainId.
  function getChainIdBySourceChainSelector(
    uint64 sourceChainSelector
  ) external view returns (uint256) {
    return s_sourceChainSelectorToChainId[sourceChainSelector];
  }

  /// @notice Returns the CrossL2Inbox address.
  function getCrossL2Inbox() external view returns (address) {
    return address(i_crossL2Inbox);
  }
}
