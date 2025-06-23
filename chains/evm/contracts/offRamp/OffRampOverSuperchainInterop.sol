// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../interfaces/IAny2EVMMessageReceiver.sol";
import {IMessageInterceptor} from "../interfaces/IMessageInterceptor.sol";
import {INonceManager} from "../interfaces/INonceManager.sol";
import {IPoolV1} from "../interfaces/IPool.sol";
import {IRouter} from "../interfaces/IRouter.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";

import {Client} from "../libraries/Client.sol";
import {ERC165CheckerReverting} from "../libraries/ERC165CheckerReverting.sol";
import {Internal} from "../libraries/Internal.sol";
import {Pool} from "../libraries/Pool.sol";
import {SuperchainInterop} from "../libraries/SuperchainInterop.sol";

import {ICrossL2Inbox} from "../vendor/optimism/interop-lib/v0/src/interfaces/ICrossL2Inbox.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {CallWithExactGas} from "@chainlink/contracts/src/v0.8/shared/call/CallWithExactGas.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/token/ERC20/IERC20.sol";
import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice This OffRamp supports OP Superchain Interop. It leverages CrossL2Inbox to verify validity of a message,
/// as opposed to relying on roots signed by oracles.
/// @dev This OffRamp does not expect transmission to come from DONs, therefore it does not inherite MultiOCR3Base
contract OffRampOverSuperchainInterop is ITypeAndVersion, Ownable2StepMsgSender {
  using ERC165CheckerReverting for address;
  using EnumerableSet for EnumerableSet.UintSet;
  using EnumerableSet for EnumerableSet.AddressSet;

  error ZeroChainSelectorNotAllowed();
  error ExecutionError(bytes32 messageId, bytes err);
  error SourceChainNotEnabled(uint64 sourceChainSelector);
  error TokenDataMismatch(uint64 sourceChainSelector, uint64 sequenceNumber);
  error ManualExecutionNotYetEnabled(uint64 sourceChainSelector);
  error InvalidManualExecutionGasLimit(uint64 sourceChainSelector, bytes32 messageId, uint256 newLimit);
  error InvalidManualExecutionTokenGasOverride(
    bytes32 messageId, uint256 tokenIndex, uint256 oldLimit, uint256 tokenGasOverride
  );
  error ManualExecutionGasAmountCountMismatch(bytes32 messageId, uint64 sequenceNumber);
  error CanOnlySelfCall();
  error ReceiverError(bytes err);
  error TokenHandlingError(address target, bytes err);
  error ReleaseOrMintBalanceMismatch(uint256 amountReleased, uint256 balancePre, uint256 balancePost);
  error NotACompatiblePool(address notPool);
  error InvalidDataLength(uint256 expected, uint256 got);
  error InvalidNewState(uint64 sourceChainSelector, uint64 sequenceNumber, Internal.MessageExecutionState newState);
  error InvalidInterval(uint64 sourceChainSelector, uint64 min, uint64 max);
  error ZeroAddressNotAllowed();
  error InvalidMessageDestChainSelector(uint64 messageDestChainSelector);
  error SourceChainSelectorMismatch(uint256 expectedChainId, uint256 gotChainId);
  error InvalidOnRampUpdate(uint64 sourceChainSelector);
  error MessageNotVerified(uint64 sourceChainSelector, bytes32 messageId);
  error UnauthorizedTransmitter();
  error InsufficientGasToCompleteTx(bytes4 err);
  error ForkedChain(uint256 expected, uint256 actual);
  error InvalidDestChainSelector(uint64 destChainSelector);
  error InvalidSourceOnRamp(address sourceOnRamp);
  error InvalidInteropLogSelector(bytes32 selector);
  error MismatchedDestChainSelector(uint64 expectedSelector, uint64 gotSelector);
  error MismatchedSequenceNumber(uint64 expectedSequenceNumber, uint64 gotSequenceNumber);
  error ZeroChainIdNotAllowed();
  error ChainIdNotConfiguredForSelector(uint64 sourceChainSelector);

  /// @dev Atlas depends on various events, if changing, please notify Atlas.
  event StaticConfigSet(StaticConfig staticConfig);
  event DynamicConfigSet(DynamicConfig dynamicConfig);
  event ExecutionStateChanged(
    uint64 indexed sourceChainSelector,
    uint64 indexed sequenceNumber,
    bytes32 indexed messageId,
    bytes32 messageHash,
    Internal.MessageExecutionState state,
    bytes returnData,
    uint256 gasUsed
  );
  event SourceChainSelectorAdded(uint64 sourceChainSelector);
  event SourceChainConfigSet(uint64 indexed sourceChainSelector, SourceChainConfig sourceConfig);
  event SkippedAlreadyExecutedMessage(uint64 sourceChainSelector, uint64 sequenceNumber);
  event AlreadyAttempted(uint64 sourceChainSelector, uint64 sequenceNumber);
  event SkippedReportExecution(uint64 sourceChainSelector);
  event AllowedTransmitterAdded(address indexed transmitter);
  event AllowedTransmitterRemoved(address indexed transmitter);
  event ChainSelectorToChainIdConfigAdded(uint64 indexed chainSelector, uint256 indexed chainId);
  event ChainSelectorToChainIdConfigRemoved(uint64 indexed chainSelector, uint256 indexed chainId);

  /// @dev Struct that contains the static configuration. The individual components are stored as immutable variables.
  // solhint-disable-next-line gas-struct-packing
  struct StaticConfig {
    uint64 chainSelector; // ───────╮ Destination chainSelector
    uint16 gasForCallExactCheck; // | Gas for call exact check
    ICrossL2Inbox crossL2Inbox; // ─╯ CrossL2Inbox for message verification
    address tokenAdminRegistry; // Token admin registry address
    address nonceManager; // Nonce manager address
  }

  /// @dev Per-chain source config (defining a lane from a Source Chain -> Dest OffRamp).
  struct SourceChainConfig {
    IRouter router; // ─────────────────╮ Local router to use for messages coming from this source chain.
    bool isEnabled; //                  │ Flag whether the source chain is enabled or not.
    uint64 minSeqNr; // ────────────────╯ The min sequence number expected for future messages.
    bytes onRamp; // OnRamp address on the source chain.
  }

  /// @dev Same as SourceChainConfig but with source chain selector so that an array of these
  /// can be passed in the constructor and the applySourceChainConfigUpdates function.
  struct SourceChainConfigArgs {
    IRouter router; // ─────────────────╮  Local router to use for messages coming from this source chain.
    uint64 sourceChainSelector; //      │  Source chain selector of the config to update.
    bool isEnabled; // ─────────────────╯  Flag whether the source chain is enabled or not.
    bytes onRamp; // OnRamp address on the source chain.
  }

  struct ChainSelectorToChainIdConfigArgs {
    uint64 chainSelector;
    uint256 chainId;
  }

  /// @dev Dynamic offRamp config.
  /// @dev Since DynamicConfig is part of DynamicConfigSet event, if changing it, we should update the ABI on Atlas.
  struct DynamicConfig {
    address feeQuoter; // ──────────────────────────────╮ FeeQuoter address on the local chain.
    uint32 permissionLessExecutionThresholdSeconds; // ─╯ Waiting time before manual execution is enabled.
    address messageInterceptor; // Optional, validates incoming messages (zero address = no interceptor).
  }

  /// @dev Both receiverExecutionGasLimit and tokenGasOverrides are optional. To indicate no override, set the value
  /// to 0. The length of tokenGasOverrides must match the length of tokenAmounts, even if it only contains zeros.
  struct GasLimitOverride {
    uint256 receiverExecutionGasLimit; // Overrides EVM2EVMMessage.gasLimit.
    uint32[] tokenGasOverrides; // Overrides EVM2EVMMessage.sourceTokenData.destGasAmount, length must be same as tokenAmounts.
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "OffRampOverSuperchainInterop 1.6.1-dev";
  /// @dev Hash of encoded address(0) used for empty address checks.
  bytes32 internal constant EMPTY_ENCODED_ADDRESS_HASH = keccak256(abi.encode(address(0)));
  /// @dev ChainSelector of this chain.
  uint64 internal immutable i_chainSelector;
  /// @dev The CrossL2Inbox interface.
  ICrossL2Inbox internal immutable i_crossL2Inbox;
  /// @dev The address of the token admin registry.
  address internal immutable i_tokenAdminRegistry;
  /// @dev The address of the nonce manager.
  address internal immutable i_nonceManager;
  /// @dev The minimum amount of gas to perform the call with exact gas.
  /// We include this in the offRamp so that we can redeploy to adjust it should a hardfork change the gas costs of
  /// relevant opcodes in callWithExactGas.
  uint16 internal immutable i_gasForCallExactCheck;

  // DYNAMIC CONFIG
  DynamicConfig internal s_dynamicConfig;

  /// @notice Set of source chain selectors.
  EnumerableSet.UintSet internal s_sourceChainSelectors;

  // Not using enumerable map to make contract size smaller.
  mapping(uint64 sourceChainSelector => uint256 chainId) private s_sourceChainSelectorToChainId;

  /// @notice SourceChainConfig per source chain selector.
  mapping(uint64 sourceChainSelector => SourceChainConfig sourceChainConfig) private s_sourceChainConfigs;

  // STATE
  /// @dev A mapping of sequence numbers (per source chain) to execution state using a bitmap with each execution
  /// state only taking up 2 bits of the uint256, packing 128 states into a single slot.
  /// Message state is tracked to ensure message can only be executed successfully once.
  mapping(uint64 sourceChainSelector => mapping(uint64 seqNum => uint256 executionStateBitmap)) internal
    s_executionStates;

  /// @dev Set of addresses allowed to call execute.
  EnumerableSet.AddressSet private s_allowedTransmitters;

  /// @dev The chain ID of the chain at deployment.
  uint256 internal immutable i_chainID;

  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    SourceChainConfigArgs[] memory sourceChainConfigs,
    address[] memory allowedTransmitters,
    ChainSelectorToChainIdConfigArgs[] memory chainSelectorToChainIdConfigArgs
  ) {
    if (
      address(staticConfig.crossL2Inbox) == address(0) || staticConfig.tokenAdminRegistry == address(0)
        || staticConfig.nonceManager == address(0)
    ) {
      revert ZeroAddressNotAllowed();
    }

    if (staticConfig.chainSelector == 0) {
      revert ZeroChainSelectorNotAllowed();
    }

    i_chainSelector = staticConfig.chainSelector;
    i_crossL2Inbox = staticConfig.crossL2Inbox;
    i_tokenAdminRegistry = staticConfig.tokenAdminRegistry;
    i_nonceManager = staticConfig.nonceManager;
    i_gasForCallExactCheck = staticConfig.gasForCallExactCheck;
    i_chainID = block.chainid;
    emit StaticConfigSet(staticConfig);

    _setDynamicConfig(dynamicConfig);
    _applySourceChainConfigUpdates(sourceChainConfigs);
    _applyAllowedTransmitterUpdates(new address[](0), allowedTransmitters);
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

  /// @notice Validates that the chain ID has not diverged after deployment. Reverts if the chain IDs do not match.
  function _whenChainNotForked() internal view {
    if (i_chainID != block.chainid) revert ForkedChain(i_chainID, block.chainid);
  }

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
  ) external onlyAllowedTransmitter {
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
        revert SourceChainSelectorMismatch(expectedChainId, report.identifier.chainId);
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
  // │             Static Config - Added crossL2Inbox               │
  // ================================================================

  /// @notice Returns the static config.
  /// @dev This function will always return the same struct as the contents is static and can never change.
  /// @return staticConfig The static config.
  function getStaticConfig() external view returns (StaticConfig memory) {
    return StaticConfig({
      chainSelector: i_chainSelector,
      gasForCallExactCheck: i_gasForCallExactCheck,
      crossL2Inbox: i_crossL2Inbox,
      tokenAdminRegistry: i_tokenAdminRegistry,
      nonceManager: i_nonceManager
    });
  }

  // ================================================================
  // │         Access - Ported transmitter check from OCR3Base,     │
  // │         added chainSelector to chainId resolution            │
  // ================================================================

  /// @notice Modifier to restrict access to allowed transmitters only.
  modifier onlyAllowedTransmitter() {
    if (!s_allowedTransmitters.contains(msg.sender)) {
      if (msg.sender != Internal.GAS_ESTIMATION_SENDER) {
        revert UnauthorizedTransmitter();
      }
    }
    _;
  }

  /// @notice Updates the allowed transmitter list.
  /// @param transmittersToRemove Array of transmitter addresses to remove.
  /// @param transmittersToAdd Array of transmitter addresses to add.
  function applyAllowedTransmitterUpdates(
    address[] memory transmittersToRemove,
    address[] memory transmittersToAdd
  ) external onlyOwner {
    _applyAllowedTransmitterUpdates(transmittersToRemove, transmittersToAdd);
  }

  /// @notice Updates the allowed transmitter list.
  /// @param transmittersToRemove Array of transmitter addresses to remove.
  /// @param transmittersToAdd Array of transmitter addresses to add.
  function _applyAllowedTransmitterUpdates(
    address[] memory transmittersToRemove,
    address[] memory transmittersToAdd
  ) internal {
    for (uint256 i = 0; i < transmittersToRemove.length; ++i) {
      address transmitter = transmittersToRemove[i];
      if (s_allowedTransmitters.remove(transmitter)) {
        emit AllowedTransmitterRemoved(transmitter);
      }
    }

    for (uint256 i = 0; i < transmittersToAdd.length; ++i) {
      address transmitter = transmittersToAdd[i];
      if (transmitter == address(0)) {
        revert ZeroAddressNotAllowed();
      }
      if (s_allowedTransmitters.add(transmitter)) {
        emit AllowedTransmitterAdded(transmitter);
      }
    }
  }

  /// @notice get getAllAllowedTransmittes List configured transmitters.
  /// @return configuredAddresses This is always populated with the list of allowed transmitters.
  function getAllAllowedTransmittes() external view returns (address[] memory) {
    return s_allowedTransmitters.values();
  }

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

  // ================================================================================================
  // │               Code below are same as regular Offramp, with updated comments                  │
  // ================================================================================================

  // ================================================================
  // │                           Execution                          │
  // ================================================================

  // The size of the execution state in bits.
  uint256 private constant MESSAGE_EXECUTION_STATE_BIT_WIDTH = 2;
  // The mask for the execution state bits.
  uint256 private constant MESSAGE_EXECUTION_STATE_MASK = (1 << MESSAGE_EXECUTION_STATE_BIT_WIDTH) - 1;

  /// @notice Returns the current execution state of a message based on its sequenceNumber.
  /// @param sourceChainSelector The source chain to get the execution state for.
  /// @param sequenceNumber The sequence number of the message to get the execution state for.
  /// @return executionState The current execution state of the message.
  /// @dev We use the literal number 128 because using a constant increased gas usage.
  function getExecutionState(
    uint64 sourceChainSelector,
    uint64 sequenceNumber
  ) public view returns (Internal.MessageExecutionState) {
    return Internal.MessageExecutionState(
      (
        _getSequenceNumberBitmap(sourceChainSelector, sequenceNumber)
          >> ((sequenceNumber % 128) * MESSAGE_EXECUTION_STATE_BIT_WIDTH)
      ) & MESSAGE_EXECUTION_STATE_MASK
    );
  }

  /// @notice Sets a new execution state for a given sequence number. It will overwrite any existing state.
  /// @param sourceChainSelector The source chain to set the execution state for.
  /// @param sequenceNumber The sequence number for which the state will be saved.
  /// @param newState The new value the state will be in after this function is called.
  /// @dev We use the literal number 128 because using a constant increased gas usage.
  function _setExecutionState(
    uint64 sourceChainSelector,
    uint64 sequenceNumber,
    Internal.MessageExecutionState newState
  ) internal {
    uint256 offset = (sequenceNumber % 128) * MESSAGE_EXECUTION_STATE_BIT_WIDTH;
    uint256 bitmap = _getSequenceNumberBitmap(sourceChainSelector, sequenceNumber);
    // To unset any potential existing state we zero the bits of the section the state occupies,
    // then we do an AND operation to blank out any existing state for the section.
    bitmap &= ~(MESSAGE_EXECUTION_STATE_MASK << offset);
    // Set the new state.
    bitmap |= uint256(newState) << offset;

    s_executionStates[sourceChainSelector][sequenceNumber / 128] = bitmap;
  }

  /// @param sourceChainSelector remote source chain selector to get sequence number bitmap for.
  /// @param sequenceNumber sequence number to get bitmap for.
  /// @return bitmap Bitmap of the given sequence number for the provided source chain selector. One bitmap represents
  /// 128 sequence numbers.
  function _getSequenceNumberBitmap(
    uint64 sourceChainSelector,
    uint64 sequenceNumber
  ) internal view returns (uint256 bitmap) {
    return s_executionStates[sourceChainSelector][sequenceNumber / 128];
  }

  /// @notice Try executing a message.
  /// @param message Internal.Any2EVMRampMessage memory message.
  /// @param offchainTokenData Data provided by the DON for token transfers.
  /// @return executionState The new state of the message, being either SUCCESS or FAILURE.
  /// @return errData Revert data in bytes if CCIP receiver reverted during execution.
  function _trialExecute(
    Internal.Any2EVMRampMessage memory message,
    bytes[] memory offchainTokenData,
    uint32[] memory tokenGasOverrides
  ) internal returns (Internal.MessageExecutionState executionState, bytes memory) {
    try this.executeSingleMessage(message, offchainTokenData, tokenGasOverrides) {}
    catch (bytes memory err) {
      if (msg.sender == Internal.GAS_ESTIMATION_SENDER) {
        if (
          CallWithExactGas.NOT_ENOUGH_GAS_FOR_CALL_SIG == bytes4(err)
            || CallWithExactGas.NO_GAS_FOR_CALL_EXACT_CHECK_SIG == bytes4(err)
            || ERC165CheckerReverting.InsufficientGasForStaticCall.selector == bytes4(err)
        ) {
          revert InsufficientGasToCompleteTx(bytes4(err));
        }
      }
      // return the message execution state as FAILURE and the revert data.
      // Max length of revert data is Router.MAX_RET_BYTES, max length of err is 4 + Router.MAX_RET_BYTES.
      return (Internal.MessageExecutionState.FAILURE, err);
    }
    // If message execution succeeded, no CCIP receiver return data is expected, return with empty bytes.
    return (Internal.MessageExecutionState.SUCCESS, "");
  }

  /// @notice hook for applying custom logic to the input message before executeSingleMessage()
  /// @param message initial message
  /// @return transformedMessage modified message
  function _beforeExecuteSingleMessage(
    Internal.Any2EVMRampMessage memory message
  ) internal virtual returns (Internal.Any2EVMRampMessage memory transformedMessage) {
    return message;
  }

  /// @notice Executes a single message.
  /// @param message The message that will be executed.
  /// @param offchainTokenData Token transfer data to be passed to TokenPool.
  /// @dev We make this external and callable by the contract itself, in order to try/catch
  /// its execution and enforce atomicity among successful message processing and token transfer.
  /// @dev We use ERC-165 to check for the ccipReceive interface to permit sending tokens to contracts, for example
  /// smart contract wallets, without an associated message.
  function executeSingleMessage(
    Internal.Any2EVMRampMessage memory message,
    bytes[] calldata offchainTokenData,
    uint32[] calldata tokenGasOverrides
  ) external {
    if (msg.sender != address(this)) revert CanOnlySelfCall();

    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](0);
    if (message.tokenAmounts.length > 0) {
      destTokenAmounts = _releaseOrMintTokens(
        message.tokenAmounts,
        message.sender,
        message.receiver,
        message.header.sourceChainSelector,
        offchainTokenData,
        tokenGasOverrides
      );
    }

    Client.Any2EVMMessage memory any2EvmMessage = Client.Any2EVMMessage({
      messageId: message.header.messageId,
      sourceChainSelector: message.header.sourceChainSelector,
      sender: message.sender,
      data: message.data,
      destTokenAmounts: destTokenAmounts
    });

    // The main message interceptor is the aggregate rate limiter, but we also allow for a custom interceptor. This is
    // why we always have to call into the contract when it's enabled, even when there are no tokens in the message.
    address messageInterceptor = s_dynamicConfig.messageInterceptor;
    if (messageInterceptor != address(0)) {
      try IMessageInterceptor(messageInterceptor).onInboundMessage(any2EvmMessage) {}
      catch (bytes memory err) {
        revert IMessageInterceptor.MessageValidationError(err);
      }
    }

    // There are three cases in which we skip calling the receiver:
    // 1. If the message data is empty AND the gas limit is 0.
    //          This indicates a message that only transfers tokens. It is valid to only send tokens to a contract
    //          that supports the IAny2EVMMessageReceiver interface, but without this first check we would call the
    //          receiver without any gas, which would revert the transaction.
    // 2. If the receiver is not a contract.
    // 3. If the receiver is a contract but it does not support the IAny2EVMMessageReceiver interface.
    //
    // The ordering of these checks is important, as the first check is the cheapest to execute.
    //
    // To prevent message delivery bypass issues, a modified version of the ERC165Checker is used
    // which checks for sufficient gas before making the external call.
    if (
      (message.data.length == 0 && message.gasLimit == 0) || message.receiver.code.length == 0
        || !message.receiver._supportsInterfaceReverting(type(IAny2EVMMessageReceiver).interfaceId)
    ) return;

    (bool success, bytes memory returnData,) = s_sourceChainConfigs[message.header.sourceChainSelector]
      .router
      .routeMessage(any2EvmMessage, i_gasForCallExactCheck, message.gasLimit, message.receiver);
    // If CCIP receiver execution is not successful, revert the call including token transfers.
    if (!success) revert ReceiverError(returnData);
  }

  // ================================================================
  // │                      Tokens and pools                        │
  // ================================================================

  /// @notice Uses a pool to release or mint a token to a receiver address, with balance checks before and after the
  /// transfer. This is done to ensure the exact number of tokens the pool claims to release are actually transferred.
  /// @dev The local token address is validated through the TokenAdminRegistry. If, due to some misconfiguration, the
  /// token is unknown to the registry, the offRamp will revert. The tx, and the tokens, can be retrieved by registering
  /// the token on this chain, and re-trying the msg.
  /// @param sourceTokenAmount Amount and source data of the token to be released/minted.
  /// @param originalSender The message sender on the source chain.
  /// @param receiver The address that will receive the tokens.
  /// @param sourceChainSelector The remote source chain selector
  /// @param offchainTokenData Data fetched offchain by the DON.
  /// @return destTokenAmount local token address with amount.
  function _releaseOrMintSingleToken(
    Internal.Any2EVMTokenTransfer memory sourceTokenAmount,
    bytes memory originalSender,
    address receiver,
    uint64 sourceChainSelector,
    bytes memory offchainTokenData
  ) internal returns (Client.EVMTokenAmount memory destTokenAmount) {
    // We need to safely decode the token address from the sourceTokenData as it could be wrong, in which case it
    // doesn't have to be a valid EVM address.
    address localToken = sourceTokenAmount.destTokenAddress;
    // We check with the token admin registry if the token has a pool on this chain.
    address localPoolAddress = ITokenAdminRegistry(i_tokenAdminRegistry).getPool(localToken);
    // This will call the supportsInterface through the ERC165Checker, and not directly on the pool address.
    // This is done to prevent a pool from reverting the entire transaction if it doesn't support the interface.
    // The call gets a max or 30k gas per instance, of which there are three. This means offchain gas estimations should
    // account for 90k gas overhead due to the interface check.
    if (localPoolAddress == address(0) || !localPoolAddress._supportsInterfaceReverting(Pool.CCIP_POOL_V1)) {
      revert NotACompatiblePool(localPoolAddress);
    }

    // We retrieve the local token balance of the receiver before the pool call.
    (uint256 balancePre, uint256 gasLeft) = _getBalanceOfReceiver(receiver, localToken, sourceTokenAmount.destGasAmount);

    // We determined that the pool address is a valid EVM address, but that does not mean the code at this address is a
    // (compatible) pool contract. _callWithExactGasSafeReturnData will check if the location contains a contract. If it
    // doesn't it reverts with a known error. We call the pool with exact gas  to increase resistance against malicious
    // tokens or token pools. We protect against return data bombs by capping the return data size at MAX_RET_BYTES.
    (bool success, bytes memory returnData, uint256 gasUsedReleaseOrMint) = CallWithExactGas
      ._callWithExactGasSafeReturnData(
      abi.encodeCall(
        IPoolV1.releaseOrMint,
        Pool.ReleaseOrMintInV1({
          originalSender: originalSender,
          receiver: receiver,
          sourceDenominatedAmount: sourceTokenAmount.amount,
          localToken: localToken,
          remoteChainSelector: sourceChainSelector,
          sourcePoolAddress: sourceTokenAmount.sourcePoolAddress,
          sourcePoolData: sourceTokenAmount.extraData,
          offchainTokenData: offchainTokenData
        })
      ),
      localPoolAddress,
      gasLeft,
      i_gasForCallExactCheck,
      Internal.MAX_RET_BYTES
    );

    // Wrap and rethrow the error so we can catch it lower in the stack.
    if (!success) revert TokenHandlingError(localPoolAddress, returnData);

    // If the call was successful, the returnData should be the amount released or minted denominated in the local
    // token's decimals.
    if (returnData.length != Pool.CCIP_POOL_V1_RET_BYTES) {
      revert InvalidDataLength(Pool.CCIP_POOL_V1_RET_BYTES, returnData.length);
    }
    uint256 localAmount = abi.decode(returnData, (uint256));

    // We don't need to do balance checks if the pool is the receiver, as they would always fail in the case
    // of a lockRelease pool.
    if (receiver != localPoolAddress) {
      (uint256 balancePost,) = _getBalanceOfReceiver(receiver, localToken, gasLeft - gasUsedReleaseOrMint);

      // First we check if the subtraction would result in an underflow to ensure we revert with a clear error.
      if (balancePost < balancePre || balancePost - balancePre != localAmount) {
        revert ReleaseOrMintBalanceMismatch(localAmount, balancePre, balancePost);
      }
    }

    return Client.EVMTokenAmount({token: localToken, amount: localAmount});
  }

  /// @notice Retrieves the balance of a receiver address for a given token.
  /// @param receiver The address to check the balance of.
  /// @param token The token address.
  /// @param gasLimit The gas limit to use for the call.
  /// @return balance The balance of the receiver.
  /// @return gasLeft The gas left after the call.
  function _getBalanceOfReceiver(
    address receiver,
    address token,
    uint256 gasLimit
  ) internal returns (uint256 balance, uint256 gasLeft) {
    (bool success, bytes memory returnData, uint256 gasUsed) = CallWithExactGas._callWithExactGasSafeReturnData(
      abi.encodeCall(IERC20.balanceOf, (receiver)), token, gasLimit, i_gasForCallExactCheck, Internal.MAX_RET_BYTES
    );
    if (!success) revert TokenHandlingError(token, returnData);

    // If the call was successful, the returnData should contain only the balance.
    if (returnData.length != Internal.MAX_BALANCE_OF_RET_BYTES) {
      revert InvalidDataLength(Internal.MAX_BALANCE_OF_RET_BYTES, returnData.length);
    }

    // Return the decoded balance, which cannot fail as we checked the length, and the gas that is left
    // after this call.
    return (abi.decode(returnData, (uint256)), gasLimit - gasUsed);
  }

  /// @notice Uses pools to release or mint a number of different tokens to a receiver address.
  /// @param sourceTokenAmounts List of token amounts with source data of the tokens to be released/minted.
  /// @param originalSender The message sender on the source chain.
  /// @param receiver The address that will receive the tokens.
  /// @param sourceChainSelector The remote source chain selector.
  /// @param offchainTokenData Array of token data fetched offchain by the DON.
  /// @param tokenGasOverrides Array of override gas limits to use for token transfers. If empty, the normal gas limit
  /// as defined on the source chain is used.
  /// @return destTokenAmounts local token addresses with amounts.
  function _releaseOrMintTokens(
    Internal.Any2EVMTokenTransfer[] memory sourceTokenAmounts,
    bytes memory originalSender,
    address receiver,
    uint64 sourceChainSelector,
    bytes[] calldata offchainTokenData,
    uint32[] calldata tokenGasOverrides
  ) internal returns (Client.EVMTokenAmount[] memory destTokenAmounts) {
    destTokenAmounts = new Client.EVMTokenAmount[](sourceTokenAmounts.length);
    bool isTokenGasOverridesEmpty = tokenGasOverrides.length == 0;
    for (uint256 i = 0; i < sourceTokenAmounts.length; ++i) {
      if (!isTokenGasOverridesEmpty) {
        if (tokenGasOverrides[i] != 0) {
          sourceTokenAmounts[i].destGasAmount = tokenGasOverrides[i];
        }
      }
      destTokenAmounts[i] = _releaseOrMintSingleToken(
        sourceTokenAmounts[i], originalSender, receiver, sourceChainSelector, offchainTokenData[i]
      );
    }

    return destTokenAmounts;
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Returns the current dynamic config.
  /// @return dynamicConfig The current dynamic config.
  function getDynamicConfig() external view returns (DynamicConfig memory) {
    return s_dynamicConfig;
  }

  /// @notice Returns the source chain config for the provided source chain selector.
  /// @param sourceChainSelector chain to retrieve configuration for.
  /// @return sourceChainConfig The config for the source chain.
  function getSourceChainConfig(
    uint64 sourceChainSelector
  ) external view returns (SourceChainConfig memory) {
    return s_sourceChainConfigs[sourceChainSelector];
  }

  /// @notice Returns all source chain configs.
  /// @return sourceChainConfigs The source chain configs corresponding to all the supported chain selectors.
  function getAllSourceChainConfigs() external view returns (uint64[] memory, SourceChainConfig[] memory) {
    SourceChainConfig[] memory sourceChainConfigs = new SourceChainConfig[](s_sourceChainSelectors.length());
    uint64[] memory sourceChainSelectors = new uint64[](s_sourceChainSelectors.length());
    for (uint256 i = 0; i < s_sourceChainSelectors.length(); ++i) {
      sourceChainSelectors[i] = uint64(s_sourceChainSelectors.at(i));
      sourceChainConfigs[i] = s_sourceChainConfigs[sourceChainSelectors[i]];
    }
    return (sourceChainSelectors, sourceChainConfigs);
  }

  /// @notice Updates source configs.
  /// @param sourceChainConfigUpdates Source chain configs.
  function applySourceChainConfigUpdates(
    SourceChainConfigArgs[] memory sourceChainConfigUpdates
  ) external onlyOwner {
    _applySourceChainConfigUpdates(sourceChainConfigUpdates);
  }

  /// @notice Updates source configs.
  /// @param sourceChainConfigUpdates Source chain configs.
  function _applySourceChainConfigUpdates(
    SourceChainConfigArgs[] memory sourceChainConfigUpdates
  ) internal {
    for (uint256 i = 0; i < sourceChainConfigUpdates.length; ++i) {
      SourceChainConfigArgs memory sourceConfigUpdate = sourceChainConfigUpdates[i];
      uint64 sourceChainSelector = sourceConfigUpdate.sourceChainSelector;

      if (sourceChainSelector == 0) {
        revert ZeroChainSelectorNotAllowed();
      }

      if (address(sourceConfigUpdate.router) == address(0)) {
        revert ZeroAddressNotAllowed();
      }

      SourceChainConfig storage currentConfig = s_sourceChainConfigs[sourceChainSelector];
      bytes memory newOnRamp = sourceConfigUpdate.onRamp;

      if (currentConfig.onRamp.length == 0) {
        currentConfig.minSeqNr = 1;
        emit SourceChainSelectorAdded(sourceChainSelector);
      } else {
        if (currentConfig.minSeqNr != 1 && keccak256(currentConfig.onRamp) != keccak256(newOnRamp)) {
          // OnRamp updates should only happen due to a misconfiguration.
          // If an OnRamp is misconfigured, no messages should have been executed.
          revert InvalidOnRampUpdate(sourceChainSelector);
        }
      }

      // OnRamp can never be zero - if it is, then the source chain has been added for the first time.
      if (newOnRamp.length == 0 || keccak256(newOnRamp) == EMPTY_ENCODED_ADDRESS_HASH) {
        revert ZeroAddressNotAllowed();
      }

      currentConfig.onRamp = newOnRamp;
      currentConfig.isEnabled = sourceConfigUpdate.isEnabled;
      currentConfig.router = sourceConfigUpdate.router;

      // We don't need to check the return value, as inserting the item twice has no effect.
      s_sourceChainSelectors.add(sourceChainSelector);

      emit SourceChainConfigSet(sourceChainSelector, currentConfig);
    }
  }

  /// @notice Sets the dynamic config.
  /// @param dynamicConfig The new dynamic config.
  function setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) external onlyOwner {
    _setDynamicConfig(dynamicConfig);
  }

  /// @notice Sets the dynamic config.
  /// @param dynamicConfig The dynamic config.
  function _setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) internal {
    if (dynamicConfig.feeQuoter == address(0)) {
      revert ZeroAddressNotAllowed();
    }

    s_dynamicConfig = dynamicConfig;

    emit DynamicConfigSet(dynamicConfig);
  }

  /// @notice Returns a source chain config with a check that the config is enabled.
  /// @param sourceChainSelector Source chain selector to check.
  /// @return sourceChainConfig The source chain config storage pointer.
  function _getEnabledSourceChainConfig(
    uint64 sourceChainSelector
  ) internal view returns (SourceChainConfig storage) {
    SourceChainConfig storage sourceChainConfig = s_sourceChainConfigs[sourceChainSelector];
    if (!sourceChainConfig.isEnabled) {
      revert SourceChainNotEnabled(sourceChainSelector);
    }

    return sourceChainConfig;
  }

  // ================================================================
  // │                            Access                            │
  // ================================================================

  /// @notice Reverts as this contract should not be able to receive CCIP messages.
  function ccipReceive(
    Client.Any2EVMMessage calldata
  ) external pure {
    // solhint-disable-next-line
    revert();
  }
}
