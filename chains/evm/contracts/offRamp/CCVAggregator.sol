// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../interfaces/IAny2EVMMessageReceiver.sol";
import {IAny2EVMMessageReceiverV2} from "../interfaces/IAny2EVMMessageReceiverV2.sol";
import {ICCVOffRamp} from "../interfaces/ICCVOffRamp.sol";
import {IPoolV1} from "../interfaces/IPool.sol";
import {IRMNRemote} from "../interfaces/IRMNRemote.sol";
import {IRouter} from "../interfaces/IRouter.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Client} from "../libraries/Client.sol";
import {ERC165CheckerReverting} from "../libraries/ERC165CheckerReverting.sol";
import {Internal} from "../libraries/Internal.sol";
import {Pool} from "../libraries/Pool.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {CallWithExactGas} from "@chainlink/contracts/src/v0.8/shared/call/CallWithExactGas.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/token/ERC20/IERC20.sol";
import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

contract CCVAggregator is ITypeAndVersion, Ownable2StepMsgSender {
  using ERC165CheckerReverting for address;
  using EnumerableSet for EnumerableSet.UintSet;

  error ZeroChainSelectorNotAllowed();
  error ExecutionError(bytes32 messageId, bytes err);
  error NoCCVQuorumReached();
  error SourceChainNotEnabled(uint64 sourceChainSelector);
  error CanOnlySelfCall();
  error ReceiverError(bytes err);
  error TokenHandlingError(address target, bytes err);
  error ReleaseOrMintBalanceMismatch(uint256 amountReleased, uint256 balancePre, uint256 balancePost);
  error CursedByRMN(uint64 sourceChainSelector);
  error NotACompatiblePool(address notPool);
  error InvalidProofLength(uint256 expected, uint256 got);
  error InvalidNewState(uint64 sourceChainSelector, uint64 sequenceNumber, Internal.MessageExecutionState newState);
  error ZeroAddressNotAllowed();
  error InvalidMessageDestChainSelector(uint64 messageDestChainSelector);
  error InsufficientGasToCompleteTx(bytes4 err);
  error SkippedAlreadyExecutedMessage(uint64 sourceChainSelector, uint64 sequenceNumber);
  error InvalidVerifierSelector(bytes4 selector);
  error ReentrancyGuardReentrantCall();

  /// @dev Atlas depends on various events, if changing, please notify Atlas.
  event StaticConfigSet(StaticConfig staticConfig);
  event ExecutionStateChanged(
    uint64 indexed sourceChainSelector,
    uint64 indexed sequenceNumber,
    bytes32 indexed messageId,
    Internal.MessageExecutionState state,
    bytes returnData
  );
  event SourceChainConfigSet(uint64 indexed sourceChainSelector, SourceChainConfig sourceConfig);

  /// @dev Struct that contains the static configuration. The individual components are stored as immutable variables.
  // solhint-disable-next-line gas-struct-packing
  struct StaticConfig {
    uint64 localChainSelector; // ──╮ Local chainSelector
    uint16 gasForCallExactCheck; // │ Gas for call exact check
    IRMNRemote rmnRemote; // ───────╯ RMN Verification Contract
    address tokenAdminRegistry; // Token admin registry address
  }

  /// @dev Per-chain source config (defining a lane from a Source Chain -> Dest OffRamp).
  struct SourceChainConfig {
    IRouter router; // ─╮ Local router to use for messages coming from this source chain.
    bool isEnabled; // ─╯ Flag whether the source chain is enabled or not.
    address defaultCCV; // Default CCV to use for messages from this source chain.
    address requiredCCV; // Required CCV to use for all messages from this source chain.
  }

  /// @dev Same as SourceChainConfig but with source chain selector so that an array of these
  /// can be passed in the constructor and the applySourceChainConfigUpdates function.
  struct SourceChainConfigArgs {
    IRouter router; // ────────────╮  Local router to use for messages coming from this source chain.
    uint64 sourceChainSelector; // │  Source chain selector of the config to update.
    bool isEnabled; // ────────────╯  Flag whether the source chain is enabled or not.
    bytes onRamp; // OnRamp address on the source chain.
  }

  struct AggregatedReport {
    /// @notice The message that is being executed.
    Internal.Any2EVMMessage message; // The message is attested to by each CCV in the report.
    /// @notice CCVs that attested to the message. They must match the CCVs specified by the receiver of the message,
    /// and the pool of the token being transferred. They can be a superset, but the ones not specified by the receiver
    /// will be ignored. If there is no token transfer, no additional token CCVs are required. If the receiver is an EOA
    /// or a contract that does not support the IAny2EVMMessageReceiver2 interface, the default and required CCVs are
    /// used.
    /// @dev Must be the same length and proofs.
    address[] ccvs;
    /// @notice This data is specific to the CCV implementation and is used to verify the message.
    bytes[] ccvData;
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "CCVAggregator 1.7.0-dev";
  /// @dev Hash of encoded address(0) used for empty address checks.
  bytes32 internal constant EMPTY_ENCODED_ADDRESS_HASH = keccak256(abi.encode(address(0)));
  /// @dev ChainSelector of this chain.
  uint64 internal immutable i_chainSelector;
  /// @dev The RMN verification contract.
  IRMNRemote internal immutable i_rmnRemote;
  /// @dev The address of the token admin registry.
  address internal immutable i_tokenAdminRegistry;
  /// @dev The minimum amount of gas to perform the call with exact gas.
  /// We include this in the offRamp so that we can redeploy to adjust it should a hardfork change the gas costs of
  /// relevant opcodes in callWithExactGas.
  uint16 internal immutable i_gasForCallExactCheck;

  // DYNAMIC CONFIG
  bool private s_reentrancyGuardEntered;

  /// @notice Set of source chain selectors.
  EnumerableSet.UintSet internal s_sourceChainSelectors;

  /// @notice SourceChainConfig per source chain selector.
  mapping(uint64 sourceChainSelector => SourceChainConfig sourceChainConfig) private s_sourceChainConfigs;

  // STATE
  /// @dev A mapping of sequence numbers (per source chain) to execution state using a bitmap with each execution
  /// state only taking up 2 bits of the uint256, packing 128 states into a single slot.
  /// Message state is tracked to ensure message can only be executed successfully once.
  mapping(uint64 sourceChainSelector => mapping(uint64 seqNum => uint256 executionStateBitmap)) internal
    s_executionStates;

  constructor(
    StaticConfig memory staticConfig
  ) {
    if (address(staticConfig.rmnRemote) == address(0) || staticConfig.tokenAdminRegistry == address(0)) {
      revert ZeroAddressNotAllowed();
    }

    if (staticConfig.localChainSelector == 0) {
      revert ZeroChainSelectorNotAllowed();
    }

    i_chainSelector = staticConfig.localChainSelector;
    i_rmnRemote = staticConfig.rmnRemote;
    i_tokenAdminRegistry = staticConfig.tokenAdminRegistry;
    i_gasForCallExactCheck = staticConfig.gasForCallExactCheck;
    emit StaticConfigSet(staticConfig);
  }

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

  /// @notice Executes a report, executing each message in order.
  /// @param report The execution report containing the messages and proofs.
  /// @dev If called from the DON, this array is always empty.
  /// @dev If called from manual execution, this array is always same length as messages.
  /// @dev This function can fully revert in some cases, reverting potentially valid other reports with it. The reasons
  /// for these reverts are so severe that we prefer to revert the entire batch instead of silently failing.
  function execute(
    AggregatedReport calldata report
  ) external {
    if (s_reentrancyGuardEntered) revert ReentrancyGuardReentrantCall();
    s_reentrancyGuardEntered = true;

    Internal.Any2EVMMessage memory message = _beforeExecuteSingleMessage(report.message);

    uint64 sourceChainSelector = report.message.header.sourceChainSelector;

    if (i_rmnRemote.isCursed(bytes16(uint128(sourceChainSelector)))) {
      revert CursedByRMN(sourceChainSelector);
    }
    if (!s_sourceChainConfigs[sourceChainSelector].isEnabled) {
      revert SourceChainNotEnabled(sourceChainSelector);
    }
    if (report.message.header.destChainSelector != i_chainSelector) {
      revert InvalidMessageDestChainSelector(report.message.header.destChainSelector);
    }

    /////// Original state checks ///////

    Internal.MessageExecutionState originalState =
      getExecutionState(sourceChainSelector, report.message.header.sequenceNumber);

    // Two valid cases here, we either have never touched this message before, or we tried to execute and failed. This
    // check protects against reentry and re-execution because the other state is IN_PROGRESS which should not be
    // allowed to execute.
    if (
      !(
        originalState == Internal.MessageExecutionState.UNTOUCHED
          || originalState == Internal.MessageExecutionState.FAILURE
      )
    ) {
      revert SkippedAlreadyExecutedMessage(sourceChainSelector, message.header.sequenceNumber);
    }

    /////// SECURITY CRITICAL CHECKS ///////

    _ensureCCVQuorumIsReached(
      report.message.header.sourceChainSelector,
      report.message.receiver,
      report.message.tokenAmounts.destTokenAddress, // TODO fix
      report.ccvs
    );

    {
      bytes memory encodedMessage = abi.encode(report.message);
      // TODO real hash
      bytes32 messageHash = keccak256(encodedMessage);
      for (uint256 i = 0; i < report.ccvs.length; ++i) {
        ICCVOffRamp(report.ccvs[i]).validateReport(encodedMessage, messageHash, report.ccvData[i], originalState);
      }
    }

    /////// Execution ///////

    _setExecutionState(sourceChainSelector, message.header.sequenceNumber, Internal.MessageExecutionState.IN_PROGRESS);
    (Internal.MessageExecutionState newState, bytes memory returnData) = _trialExecute(message);
    _setExecutionState(sourceChainSelector, message.header.sequenceNumber, newState);

    // The only valid prior states are UNTOUCHED and FAILURE (checked above).
    // The only valid post states are FAILURE and SUCCESS (checked below).
    if (newState != Internal.MessageExecutionState.SUCCESS) {
      if (newState != Internal.MessageExecutionState.FAILURE) {
        revert InvalidNewState(sourceChainSelector, message.header.sequenceNumber, newState);
      }
    }

    emit ExecutionStateChanged(
      sourceChainSelector, message.header.sequenceNumber, message.header.messageId, newState, returnData
    );
    s_reentrancyGuardEntered = false;
  }

  function _ensureCCVQuorumIsReached(
    uint64 sourceChainSelector,
    address receiver,
    address tokenPool,
    address[] calldata CCVs
  ) internal view {
    (address[] memory requiredCCV, address[] memory optionalCCVs, uint8 optionalThreshold) =
      _getCCVsFromReceiverAndPool(sourceChainSelector, receiver, tokenPool);

    for (uint256 i = 0; i < requiredCCV.length; ++i) {
      bool found = false;
      for (uint256 j = 0; j < CCVs.length; ++j) {
        if (CCVs[j] == requiredCCV[i]) {
          found = true;
          break;
        }
      }
      if (!found) {
        revert NoCCVQuorumReached(); // TODO better error message
      }
    }

    uint256 optionalCCVsToFind = optionalThreshold;
    for (uint256 i = 0; i < optionalCCVs.length; ++i) {
      for (uint256 j = 0; j < CCVs.length && optionalCCVsToFind > 0; ++j) {
        if (CCVs[j] == optionalCCVs[i]) {
          optionalCCVsToFind--;
          break;
        }
      }
    }

    if (optionalCCVsToFind > 0) {
      revert NoCCVQuorumReached(); // TODO better error message
    }
  }

  // TODO check pool as well.
  function _getCCVsFromReceiverAndPool(
    uint64 sourceChainSelector,
    address receiver,
    address // poolAddress
  ) internal view returns (address[] memory requiredCCV, address[] memory optionalCCVs, uint8 optionalThreshold) {
    // If the receiver is not a contract, or it doesn't support the required interface, we return the default.
    if (receiver.code.length == 0 || !receiver._supportsInterfaceReverting(type(IAny2EVMMessageReceiverV2).interfaceId))
    {
      SourceChainConfig memory sourceConfig = s_sourceChainConfigs[sourceChainSelector];
      if (sourceConfig.requiredCCV == address(0)) {
        requiredCCV = new address[](1);
        requiredCCV[0] = sourceConfig.defaultCCV;

        return (requiredCCV, new address[](0), 0);
      }

      requiredCCV = new address[](2);
      requiredCCV[0] = sourceConfig.defaultCCV;
      requiredCCV[1] = sourceConfig.requiredCCV;

      return (requiredCCV, new address[](0), 0);
    }

    // If it does, we return the required CCVs from the receiver.
    return IAny2EVMMessageReceiverV2(receiver).getCCVs(sourceChainSelector);
  }

  /// @notice Try executing a message.
  /// @param message Internal.Any2EVMMultiProofMessage memory message.
  /// @return executionState The new state of the message, being either SUCCESS or FAILURE.
  /// @return errData Revert data in bytes if CCIP receiver reverted during execution.
  function _trialExecute(
    Internal.Any2EVMMessage memory message
  ) internal returns (Internal.MessageExecutionState executionState, bytes memory) {
    (bool success, bytes memory returnData,) = CallWithExactGas._callWithExactGasSafeReturnData(
      abi.encodeCall(this.executeSingleMessage, (message)),
      address(this),
      gasleft(),
      i_gasForCallExactCheck,
      Internal.MAX_RET_BYTES
    );

    if (!success) {
      if (msg.sender == Internal.GAS_ESTIMATION_SENDER) {
        if (
          CallWithExactGas.NOT_ENOUGH_GAS_FOR_CALL_SIG == bytes4(returnData)
            || CallWithExactGas.NO_GAS_FOR_CALL_EXACT_CHECK_SIG == bytes4(returnData)
            || ERC165CheckerReverting.InsufficientGasForStaticCall.selector == bytes4(returnData)
        ) {
          revert InsufficientGasToCompleteTx(bytes4(returnData));
        }
      }
      // return the message execution state as FAILURE and the revert data.
      // Max length of revert data is Router.MAX_RET_BYTES, max length of err is 4 + Router.MAX_RET_BYTES.
      return (Internal.MessageExecutionState.FAILURE, returnData);
    }

    // If message execution succeeded, no CCIP receiver return data is expected, return with empty bytes.
    return (Internal.MessageExecutionState.SUCCESS, "");
  }

  /// @notice hook for applying custom logic to the input message before executeSingleMessage()
  /// @param message initial message
  /// @return transformedMessage modified message
  function _beforeExecuteSingleMessage(
    Internal.Any2EVMMessage memory message
  ) internal virtual returns (Internal.Any2EVMMessage memory transformedMessage) {
    return message;
  }

  /// @notice Executes a single message.
  /// @param message The message that will be executed.
  /// @dev We make this external and callable by the contract itself, in order to try/catch
  /// its execution and enforce atomicity among successful message processing and token transfer.
  /// @dev We use ERC-165 to check for the ccipReceive interface to permit sending tokens to contracts, for example
  /// smart contract wallets, without an associated message.
  function executeSingleMessage(
    Internal.Any2EVMMessage memory message
  ) external {
    if (msg.sender != address(this)) revert CanOnlySelfCall();

    bool hasToken = message.tokenAmounts.destTokenAddress != address(0);

    Client.EVMTokenAmount[] memory destTokenAmounts = new Client.EVMTokenAmount[](hasToken ? 1 : 0);
    if (hasToken) {
      destTokenAmounts[0] = _releaseOrMintSingleToken(
        message.tokenAmounts, message.sender, message.receiver, message.header.sourceChainSelector
      );
    }

    Client.Any2EVMMessage memory any2EvmMessage = Client.Any2EVMMessage({
      messageId: message.header.messageId,
      sourceChainSelector: message.header.sourceChainSelector,
      sender: message.sender,
      data: message.data,
      destTokenAmounts: destTokenAmounts
    });

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
  /// @return destTokenAmount local token address with amount.
  function _releaseOrMintSingleToken(
    Internal.TokenTransfer memory sourceTokenAmount,
    bytes memory originalSender,
    address receiver,
    uint64 sourceChainSelector
  ) internal returns (Client.EVMTokenAmount memory destTokenAmount) {
    address localToken = sourceTokenAmount.destTokenAddress;
    // We check with the token admin registry if the token has a pool on this chain.
    address localPoolAddress = ITokenAdminRegistry(i_tokenAdminRegistry).getPool(localToken);
    // This will call the supportsInterface through the ERC165Checker, and not directly on the pool address.
    // This is done to prevent a pool from reverting the entire transaction if it doesn't support the interface.
    // The call gets a max or 30k gas per instance, of which there are three. This means offchain gas estimations should
    // account for 90k gas overhead due to the interface check.
    if (localPoolAddress == address(0)) {
      revert NotACompatiblePool(localPoolAddress);
    }

    // Check V2 first, as it is the most recent version of the pool interface.
    if (localPoolAddress._supportsInterfaceReverting(Pool.CCIP_POOL_V2)) {
      // Revert for now
      // TODO write IPoolV2
      revert NotACompatiblePool(localPoolAddress);
    }

    if (!localPoolAddress._supportsInterfaceReverting(Pool.CCIP_POOL_V1)) {
      // If the pool does not support the v1 interface, we revert.
      revert NotACompatiblePool(localPoolAddress);
    }
    // We retrieve the local token balance of the receiver before the pool call.
    uint256 balancePre = _getBalanceOfReceiver(receiver, localToken);

    Pool.ReleaseOrMintOutV1 memory returnData;
    try IPoolV1(localPoolAddress).releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: originalSender,
        receiver: receiver,
        sourceDenominatedAmount: sourceTokenAmount.amount,
        localToken: localToken,
        remoteChainSelector: sourceChainSelector,
        sourcePoolAddress: sourceTokenAmount.sourcePoolAddress,
        sourcePoolData: sourceTokenAmount.extraData,
        // All use cases that use offchain token data in IPoolV1 have to upgrade to the modular security interface.
        offchainTokenData: ""
      })
    ) returns (Pool.ReleaseOrMintOutV1 memory result) {
      returnData = result;
    } catch (bytes memory err) {
      revert TokenHandlingError(localToken, err);
    }

    // We don't need to do balance checks if the pool is the receiver, as they would always fail in the case
    // of a lockRelease pool.
    if (receiver != localPoolAddress) {
      uint256 balancePost = _getBalanceOfReceiver(receiver, localToken);

      // First we check if the subtraction would result in an underflow to ensure we revert with a clear error.
      if (balancePost < balancePre || balancePost - balancePre != returnData.destinationAmount) {
        revert ReleaseOrMintBalanceMismatch(returnData.destinationAmount, balancePre, balancePost);
      }
    }

    return Client.EVMTokenAmount({token: localToken, amount: returnData.destinationAmount});
  }

  /// @notice Retrieves the balance of a receiver address for a given token.
  /// @param receiver The address to check the balance of.
  /// @param token The token address.
  /// @return balance The balance of the receiver.
  function _getBalanceOfReceiver(address receiver, address token) internal view returns (uint256) {
    try IERC20(token).balanceOf(receiver) returns (uint256 balance_) {
      // If the call succeeds, we return the balance and the gas left.
      return (balance_);
    } catch (bytes memory err) {
      // If the call fails, we revert with a known error.
      revert TokenHandlingError(token, err);
    }
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Returns the static config.
  /// @dev This function will always return the same struct as the contents is static and can never change.
  /// @return staticConfig The static config.
  function getStaticConfig() external view returns (StaticConfig memory) {
    return StaticConfig({
      localChainSelector: i_chainSelector,
      gasForCallExactCheck: i_gasForCallExactCheck,
      rmnRemote: i_rmnRemote,
      tokenAdminRegistry: i_tokenAdminRegistry
    });
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
    SourceChainConfigArgs[] calldata sourceChainConfigUpdates
  ) external onlyOwner {
    for (uint256 i = 0; i < sourceChainConfigUpdates.length; ++i) {
      SourceChainConfigArgs memory sourceConfigUpdate = sourceChainConfigUpdates[i];
      uint64 sourceChainSelector = sourceConfigUpdate.sourceChainSelector;

      if (sourceChainSelector == 0) {
        revert ZeroChainSelectorNotAllowed();
      }

      if (address(sourceConfigUpdate.router) == address(0)) {
        revert ZeroAddressNotAllowed();
      }

      // TODO check replay protection if onRamp changes
      SourceChainConfig storage currentConfig = s_sourceChainConfigs[sourceChainSelector];
      bytes memory newOnRamp = sourceConfigUpdate.onRamp;

      // OnRamp can never be zero - if it is, then the source chain has been added for the first time.
      if (newOnRamp.length == 0 || keccak256(newOnRamp) == EMPTY_ENCODED_ADDRESS_HASH) {
        revert ZeroAddressNotAllowed();
      }

      currentConfig.isEnabled = sourceConfigUpdate.isEnabled;
      currentConfig.router = sourceConfigUpdate.router;

      // We don't need to check the return value, as inserting the item twice has no effect.
      s_sourceChainSelectors.add(sourceChainSelector);

      emit SourceChainConfigSet(sourceChainSelector, currentConfig);
    }
  }
}
