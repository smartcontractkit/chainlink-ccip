// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../interfaces/IAny2EVMMessageReceiver.sol";
import {IAny2EVMMessageReceiverV2} from "../interfaces/IAny2EVMMessageReceiverV2.sol";
import {ICrossChainVerifierResolver} from "../interfaces/ICrossChainVerifierResolver.sol";
import {ICrossChainVerifierV1} from "../interfaces/ICrossChainVerifierV1.sol";
import {IPoolV1} from "../interfaces/IPool.sol";
import {IPoolV2} from "../interfaces/IPoolV2.sol";
import {IRMNRemote} from "../interfaces/IRMNRemote.sol";
import {IRouter} from "../interfaces/IRouter.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {CCVConfigValidation} from "../libraries/CCVConfigValidation.sol";
import {Client} from "../libraries/Client.sol";
import {ERC165CheckerReverting} from "../libraries/ERC165CheckerReverting.sol";
import {Internal} from "../libraries/Internal.sol";
import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {Pool} from "../libraries/Pool.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

contract OffRamp is ITypeAndVersion, Ownable2StepMsgSender {
  using ERC165CheckerReverting for address;
  using EnumerableSet for EnumerableSet.UintSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;

  error ZeroChainSelectorNotAllowed();
  error ExecutionError(bytes32 messageId, bytes err);
  error OptionalCCVQuorumNotReached(uint256 wanted, uint256 got);
  error SourceChainNotEnabled(uint64 sourceChainSelector);
  error CanOnlySelfCall();
  error ReceiverError(bytes err);
  error TokenHandlingError(address target, bytes err);
  error CursedByRMN(uint64 sourceChainSelector);
  error NotACompatiblePool(address notPool);
  error InvalidVerifierResultsLength(uint256 expected, uint256 got);
  error ZeroAddressNotAllowed();
  error InvalidMessageDestChainSelector(uint64 messageDestChainSelector);
  error InsufficientGasToCompleteTx(bytes4 err);
  error SkippedAlreadyExecutedMessage(bytes32 messageId, uint64 sourceChainSelector, uint64 messageNumber);
  error ReentrancyGuardReentrantCall();
  error RequiredCCVMissing(address requiredCCV);
  error InvalidNumberOfTokens(uint256 numTokens);
  error InvalidOnRamp(bytes got);
  error InvalidOffRamp(address expected, bytes got);
  error InboundImplementationNotFound(address ccvAddress, bytes verifierResults);
  error InvalidGasLimitOverride(uint32 messageGasLimit, uint32 gasLimitOverride);
  error GasCannotBeZero();

  /// @dev Atlas depends on various events, if changing, please notify Atlas.
  event StaticConfigSet(StaticConfig staticConfig);
  event ExecutionStateChanged(
    uint64 indexed sourceChainSelector,
    uint64 indexed messageNumber,
    bytes32 indexed messageId,
    Internal.MessageExecutionState state,
    bytes returnData
  );
  event SourceChainConfigSet(uint64 indexed sourceChainSelector, SourceChainConfigArgs sourceConfig);

  // 5k for updating the state + 5k for the event and misc costs.
  uint256 internal constant MAX_GAS_BUFFER_TO_UPDATE_STATE = 5000 + 5000 + 2000;

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
    bytes[] onRamps; // OnRamp address on the source chain.
    address[] defaultCCVs; // Default CCVs to use for messages from this source chain.
    address[] laneMandatedCCVs; // Required CCVs to use for all messages from this source chain.
  }

  /// @dev Same as SourceChainConfig but with source chain selector so that an array of these
  /// can be passed in the constructor and the applySourceChainConfigUpdates function.
  // solhint-disable-next-line gas-struct-packing
  struct SourceChainConfigArgs {
    IRouter router; // ────────────╮ Local router to use for messages coming from this source chain.
    uint64 sourceChainSelector; // │ Source chain selector of the config to update.
    bool isEnabled; // ────────────╯ Flag whether the source chain is enabled or not.
    bytes[] onRamps; // OnRamp address on the source chain. For EVM source chains, these should be abi-encoded (32 bytes).
    address[] defaultCCVs; // Default CCVs to use for messages from this source chain.
    address[] laneMandatedCCVs; // Required CCVs to use for all messages from this source chain.
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "OffRamp 1.7.0-dev";
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

  // At the top to pack it with the `owner` variable from Ownable2StepMsgSender.
  bool private s_reentrancyGuardEntered;

  // DYNAMIC CONFIG

  /// @notice Set of source chain selectors.
  EnumerableSet.UintSet internal s_sourceChainSelectors;

  /// @notice SourceChainConfig per source chain selector.
  mapping(uint64 sourceChainSelector => SourceChainConfig sourceChainConfig) private s_sourceChainConfigs;

  /// @notice Set of allowed onRamp address hashes per source chain selector. We hash the onRamp addresses to save on
  /// gas during retrieval. These sets are duplicated in the source chain config in their raw form to enable lookups.
  mapping(uint64 sourceChainSelector => EnumerableSet.Bytes32Set allowedOnRampHashes) private s_allowedOnRampHashes;

  // STATE

  /// Message state is tracked to ensure message can only be executed successfully once.
  mapping(bytes32 execStateKey => Internal.MessageExecutionState state) internal s_executionStates;

  constructor(
    StaticConfig memory staticConfig
  ) {
    if (address(staticConfig.rmnRemote) == address(0) || staticConfig.tokenAdminRegistry == address(0)) {
      revert ZeroAddressNotAllowed();
    }

    if (staticConfig.localChainSelector == 0) {
      revert ZeroChainSelectorNotAllowed();
    }
    if (staticConfig.gasForCallExactCheck == 0) {
      revert GasCannotBeZero();
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

  /// @notice Returns the current execution state of a message.
  /// @return executionState The current execution state of the message.
  function getExecutionState(
    bytes32 messageId
  ) public view returns (Internal.MessageExecutionState) {
    return s_executionStates[messageId];
  }

  /// @notice Executes a message from a source chain.
  /// @param encodedMessage The message that is being executed, encoded as bytes.
  /// @param ccvs CCVs that attested to the message. Must match the CCVs specified by the receiver and token pool.
  /// @param verifierResults CCV-specific data used to verify the message. Must be same length as ccvs array.
  function execute(
    bytes calldata encodedMessage,
    address[] calldata ccvs,
    bytes[] calldata verifierResults,
    uint32 gasLimitOverride
  ) external {
    if (s_reentrancyGuardEntered) revert ReentrancyGuardReentrantCall();
    s_reentrancyGuardEntered = true;

    MessageV1Codec.MessageV1 memory message =
      _beforeExecuteSingleMessage(MessageV1Codec._decodeMessageV1(encodedMessage));

    if (i_rmnRemote.isCursed(bytes16(uint128(message.sourceChainSelector)))) {
      revert CursedByRMN(message.sourceChainSelector);
    }
    if (!s_sourceChainConfigs[message.sourceChainSelector].isEnabled) {
      revert SourceChainNotEnabled(message.sourceChainSelector);
    }
    if (!s_allowedOnRampHashes[message.sourceChainSelector].contains(keccak256(message.onRampAddress))) {
      revert InvalidOnRamp(message.onRampAddress);
    }
    if (message.offRampAddress.length != 20 || address(bytes20(message.offRampAddress)) != address(this)) {
      revert InvalidOffRamp(address(this), message.offRampAddress);
    }
    if (message.destChainSelector != i_chainSelector) {
      revert InvalidMessageDestChainSelector(message.destChainSelector);
    }
    if (ccvs.length != verifierResults.length) {
      revert InvalidVerifierResultsLength(ccvs.length, verifierResults.length);
    }
    if (message.receiver.length != 20) {
      revert Internal.InvalidEVMAddress(message.receiver);
    }
    // gasLimitOverride == 0 means "no override" (use message.ccipReceiveGasLimit).
    // A non-zero override must not be lower than the message-provided gas limit.
    if (gasLimitOverride != 0 && gasLimitOverride < message.ccipReceiveGasLimit) {
      revert InvalidGasLimitOverride(message.ccipReceiveGasLimit, gasLimitOverride);
    }

    /////// Original state checks ///////

    bytes32 messageId = keccak256(encodedMessage);

    Internal.MessageExecutionState originalState = s_executionStates[messageId];

    // Two valid cases here, we either have never touched this message before, or we tried to execute and failed. This
    // check protects against reentry and re-execution because the other state is IN_PROGRESS which should not be
    // allowed to execute.
    if (!(originalState == Internal.MessageExecutionState.UNTOUCHED
          || originalState == Internal.MessageExecutionState.FAILURE)) {
      revert SkippedAlreadyExecutedMessage(messageId, message.sourceChainSelector, message.messageNumber);
    }

    /////// Execution ///////

    s_executionStates[messageId] = Internal.MessageExecutionState.IN_PROGRESS;

    (bool success, bytes memory err) = _callWithGasBuffer(
      abi.encodeCall(this.executeSingleMessage, (message, messageId, ccvs, verifierResults, gasLimitOverride))
    );
    Internal.MessageExecutionState newState =
      success ? Internal.MessageExecutionState.SUCCESS : Internal.MessageExecutionState.FAILURE;

    s_executionStates[messageId] = newState;

    emit ExecutionStateChanged(message.sourceChainSelector, message.messageNumber, messageId, newState, err);
    s_reentrancyGuardEntered = false;
  }

  function _callWithGasBuffer(
    bytes memory payload
  ) internal returns (bool success, bytes memory retData) {
    // allocate retData memory ahead of time.
    retData = new bytes(Internal.MAX_RET_BYTES);
    uint16 maxReturnBytes = Internal.MAX_RET_BYTES;

    uint256 gasLeft = gasleft();
    if (gasLeft <= MAX_GAS_BUFFER_TO_UPDATE_STATE) {
      revert InsufficientGasToCompleteTx(bytes4(uint32(gasleft())));
    }

    uint256 gasLimit = gasLeft - MAX_GAS_BUFFER_TO_UPDATE_STATE;

    assembly {
      // Call and return whether we succeeded.
      // call(gas, addr, value, argsOffset, argsLength, retOffset, retLength).
      success := call(gasLimit, address(), 0, add(payload, 0x20), mload(payload), 0x0, 0x0)

      // Limit our copy to maxReturnBytes bytes.
      let toCopy := returndatasize()
      if gt(toCopy, maxReturnBytes) { toCopy := maxReturnBytes }
      // Store the length of the copied bytes.
      mstore(retData, toCopy)
      // copy the bytes from retData[0:_toCopy].
      returndatacopy(add(retData, 0x20), 0x0, toCopy)
    }
    return (success, retData);
  }

  /// @notice Executes a single message.
  /// @param message The message that will be executed.
  /// @dev We make this external and callable by the contract itself, in order to try/catch
  /// its execution and enforce atomicity among successful message processing and token transfer.
  /// @dev We use ERC-165 to check for the ccipReceive interface to permit sending tokens to contracts, for example
  /// smart contract wallets, without an associated message.
  function executeSingleMessage(
    MessageV1Codec.MessageV1 calldata message,
    bytes32 messageId,
    address[] calldata ccvs,
    bytes[] calldata verifierResults,
    uint32 gasLimitOverride
  ) external {
    if (msg.sender != address(this)) revert CanOnlySelfCall();

    // We track the balance of the receiver prior to verification because a verifier may be responsible for releasing or
    // minting the token.
    uint256 balancePre = 0;
    if (message.tokenTransfer.length > 0) {
      if (message.tokenTransfer[0].destTokenAddress.length != 20) {
        revert Internal.InvalidEVMAddress(message.tokenTransfer[0].destTokenAddress);
      }
      if (message.tokenTransfer[0].tokenReceiver.length != 20) {
        revert Internal.InvalidEVMAddress(message.tokenTransfer[0].tokenReceiver);
      }
      balancePre = _getBalanceOfReceiver(
        address(bytes20(message.tokenTransfer[0].tokenReceiver)),
        address(bytes20(message.tokenTransfer[0].destTokenAddress))
      );
    }

    address receiver = address(bytes20(message.receiver));

    /////// SECURITY CRITICAL CHECKS ///////
    bool isTokenOnlyTransfer = _isTokenOnlyTransfer(message.data.length, message.ccipReceiveGasLimit, receiver);

    {
      (address[] memory ccvsToQuery, uint256[] memory verifierResultsIndex) = _ensureCCVQuorumIsReached(
        message.sourceChainSelector, receiver, message.tokenTransfer, message.finality, ccvs, isTokenOnlyTransfer
      );

      for (uint256 i = 0; i < ccvsToQuery.length; ++i) {
        address implAddress = ICrossChainVerifierResolver(ccvsToQuery[i])
          .getInboundImplementation(verifierResults[verifierResultsIndex[i]]);
        if (implAddress == address(0)) {
          revert InboundImplementationNotFound(ccvsToQuery[i], verifierResults[verifierResultsIndex[i]]);
        }
        ICrossChainVerifierV1(implAddress)
          .verifyMessage({
            message: message, messageId: messageId, verifierResults: verifierResults[verifierResultsIndex[i]]
          });
      }
    }

    Client.EVMTokenAmount[] memory tokenTransfer = new Client.EVMTokenAmount[](message.tokenTransfer.length);

    if (message.tokenTransfer.length > 0) {
      address localPoolAddress;
      (tokenTransfer[0], localPoolAddress) = _releaseOrMintSingleToken(
        message.tokenTransfer[0], message.sender, message.sourceChainSelector, message.finality
      );

      address tokenReceiver = address(bytes20(message.tokenTransfer[0].tokenReceiver));
      // If a lock-release pool is the receiver, balancePost - balancePre would not reflect the amount transferred.
      // Therefore, if the receiver is the token pool, we trust the value returned by the pool.
      // Otherwise, we trust balancePost - balancePre as the amount given to the receiver.
      if (tokenReceiver != localPoolAddress) {
        uint256 balancePost = _getBalanceOfReceiver(tokenReceiver, tokenTransfer[0].token);
        tokenTransfer[0].amount = balancePost - balancePre;
      }
    }

    // Short circuit if we don't need to call the receiver.
    if (isTokenOnlyTransfer) return;

    _callReceiver(
      Client.Any2EVMMessage({
        messageId: messageId,
        sourceChainSelector: message.sourceChainSelector,
        sender: message.sender,
        data: message.data,
        destTokenAmounts: tokenTransfer
      }),
      receiver,
      gasLimitOverride != 0 ? gasLimitOverride : message.ccipReceiveGasLimit
    );
  }

  /// @notice Calls the receiver contract via the Router.
  /// @param message The message to call the receiver for.
  /// @param receiver The receiver address.
  /// @param gasLimit The gas limit to use for the call.
  function _callReceiver(
    Client.Any2EVMMessage memory message,
    address receiver,
    uint32 gasLimit
  ) internal {
    (bool success, bytes memory returnData,) = s_sourceChainConfigs[message.sourceChainSelector].router
      .routeMessage(message, i_gasForCallExactCheck, gasLimit, receiver);

    // If CCIP receiver execution is not successful, revert the call including token transfers.
    if (!success) revert ReceiverError(returnData);
  }

  /// @notice There are three cases in which we skip calling the receiver:
  /// 1. If the message data is empty AND the gas limit is 0.
  ///          This indicates a message that only transfers tokens. It is valid to only send tokens to a contract
  ///          that supports the IAny2EVMMessageReceiver interface, but without this first check we would call the
  ///          receiver without any gas, which would revert the transaction.
  /// 2. If the receiver is not a contract.
  /// 3. If the receiver is a contract but it does not support the IAny2EVMMessageReceiver interface.
  ///
  /// The ordering of these checks is important, as the first check is the cheapest to execute.
  ///
  /// To prevent message delivery bypass issues, a modified version of the ERC165Checker is used
  /// which checks for sufficient gas before making the external call.
  /// @param dataLength The length of the message data.
  /// @param ccipReceiveGasLimit The gas limit specified for the CCIP receive call.
  /// @param receiver The receiver address.
  /// @return Whether the message is a token-only transfer.
  function _isTokenOnlyTransfer(
    uint256 dataLength,
    uint256 ccipReceiveGasLimit,
    address receiver
  ) internal view returns (bool) {
    return (dataLength == 0 && ccipReceiveGasLimit == 0) || receiver.code.length == 0
      || !receiver._supportsInterfaceReverting(type(IAny2EVMMessageReceiver).interfaceId);
  }

  // ================================================================
  // │                            CCVs                              │
  // ================================================================

  /// @notice Returns the CCVs required to execute a message. There can be duplicates between the required and optional
  // CCVs, but all duplicated within the required CCVs are removed.
  /// @dev This function exists for offchain purposes, calling it onchain might not be gas efficient.
  function getCCVsForMessage(
    bytes calldata encodedMessage
  ) external view returns (address[] memory requiredCCVs, address[] memory optionalCCVs, uint8 threshold) {
    MessageV1Codec.MessageV1 memory message = MessageV1Codec._decodeMessageV1(encodedMessage);
    address receiver = address(bytes20(message.receiver));

    return _getCCVsForMessage(
      message.sourceChainSelector,
      receiver,
      message.tokenTransfer,
      message.finality,
      _isTokenOnlyTransfer(message.data.length, message.ccipReceiveGasLimit, receiver)
    );
  }

  /// @notice Returns the CCVs required by the receiver, pool and lane for a message. Duplicates are removed and
  /// defaults are added if necessary. This function handles all the logic of combining the various sources of CCVs.
  /// @param sourceChainSelector The source chain selector of the message.
  /// @param receiver The receiver of the message.
  /// @param tokenTransfer The tokens transferred in the message.
  /// @param finality The finality requirement of the message.
  /// @param isTokenOnlyTransfer Whether the message is a token-only transfer (no exec).
  /// @return requiredCCVs The deduplicated list of required CCVs for the message.
  /// @return optionalCCVs The list of optional CCVs for the message, with duplicates removed against required CCVs.
  /// @return optionalThreshold The threshold of optional CCVs, adjusted for any duplicates with required CCVs.
  /// @dev This function is quite complex as it needs to handle multiple sources of CCVs, deduplication and adding of
  /// defaults. The function looks quite gas intensive, but the expected lengths of the various CCV arrays are small, so
  /// the gas usage should be acceptable.
  /// @dev The offchain system relies on this functions logic as well, meaning both onchain and offchain have the same
  /// source of truth for which CCVs are needed for a message.
  function _getCCVsForMessage(
    uint64 sourceChainSelector,
    address receiver,
    MessageV1Codec.TokenTransferV1[] memory tokenTransfer,
    uint16 finality,
    bool isTokenOnlyTransfer
  ) internal view returns (address[] memory requiredCCVs, address[] memory optionalCCVs, uint8 optionalThreshold) {
    address[] memory requiredPoolCCVs = new address[](0);
    if (tokenTransfer.length > 0) {
      if (tokenTransfer.length != 1) {
        revert InvalidNumberOfTokens(tokenTransfer.length);
      }
      // Not validated because it gets validated in the executeSingleMessage function.
      // Only other path that calls _getCCVsForMessage is getCCVsForMessage, which is a view function used by executors.
      address localTokenAddress = address(bytes20(tokenTransfer[0].destTokenAddress));

      // If the pool returns does not specify any CCVs, we fall back to the default CCVs. These will be deduplicated
      // in the ensureCCVQuorumIsReached function. This is to maintain the same pre-1.7.0 security level for pools
      // that do not support the V2 interface.
      requiredPoolCCVs = _getCCVsFromPool(
        localTokenAddress, sourceChainSelector, tokenTransfer[0].amount, finality, tokenTransfer[0].extraData
      );
    }

    // Get the CCVs for the receiver, if any.
    address[] memory requiredReceiverCCVs;
    if (isTokenOnlyTransfer) {
      if (tokenTransfer.length > 0) {
        // For token-only transfers, we skip querying the receiver for CCVs, and don't add the defaults. This enables
        // pure token transfers to only require the pool CCVs, as the token issuer is the only party that takes any risk.
        requiredReceiverCCVs = new address[](0);
        optionalCCVs = new address[](0);
        optionalThreshold = 0;
      } else {
        // The transfer is token-only but doesn't contain any tokens. This is a no-op transfer, we fall back to
        // requiring the default CCV.
        requiredReceiverCCVs = new address[](1);
      }
    } else {
      // The transfer is not token-only, we query the receiver for its CCV requirements.
      (requiredReceiverCCVs, optionalCCVs, optionalThreshold) = _getCCVsFromReceiver(sourceChainSelector, receiver);
    }

    address[] memory laneMandatedCCVs = s_sourceChainConfigs[sourceChainSelector].laneMandatedCCVs;
    address[] storage defaultCCVs = s_sourceChainConfigs[sourceChainSelector].defaultCCVs;

    // We allocate the memory for all possible CCVs upfront to avoid multiple allocations.
    address[] memory allRequiredCCVs = new address[](
      requiredReceiverCCVs.length + requiredPoolCCVs.length + laneMandatedCCVs.length + defaultCCVs.length
    );

    uint256 index = 0;
    for (uint256 i = 0; i < requiredReceiverCCVs.length; ++i) {
      allRequiredCCVs[index++] = requiredReceiverCCVs[i];
    }

    for (uint256 i = 0; i < requiredPoolCCVs.length; ++i) {
      allRequiredCCVs[index++] = requiredPoolCCVs[i];
    }

    for (uint256 i = 0; i < laneMandatedCCVs.length; ++i) {
      allRequiredCCVs[index++] = laneMandatedCCVs[i];
    }

    // Add defaults if any address(0) was found.
    for (uint256 i = 0; i < index; ++i) {
      if (allRequiredCCVs[i] == address(0)) {
        uint256 numberOfDefaults = defaultCCVs.length;
        for (uint256 j = 0; j < numberOfDefaults; ++j) {
          allRequiredCCVs[index++] = defaultCCVs[j];
        }
        // Break to ensure they're only added once.
        break;
      }
    }

    // Remove duplicates and address(0) entries.
    uint256 writeIndex = 0;
    for (uint256 readIndex = 0; readIndex < index; ++readIndex) {
      address currentCCV = allRequiredCCVs[readIndex];

      // Skip address(0) entries, effectively removing them.
      if (currentCCV == address(0)) {
        continue;
      }

      // Check if this address already exists in the deduplicated portion.
      bool isDuplicate = false;
      for (uint256 j = 0; j < writeIndex; ++j) {
        if (allRequiredCCVs[j] == currentCCV) {
          isDuplicate = true;
          break;
        }
      }

      // If not a duplicate, add it to the deduplicated portion.
      if (!isDuplicate) {
        allRequiredCCVs[writeIndex++] = currentCCV;
      }
    }

    assembly {
      // set the length of the array to the new index which we used to track the number of unique CCVs.
      mstore(allRequiredCCVs, writeIndex)
    }

    // Remove duplicates between required and optional CCVs.
    uint256 newOptionalLength = optionalCCVs.length;
    for (uint256 i = 0; i < allRequiredCCVs.length; ++i) {
      for (uint256 j = 0; j < newOptionalLength;) {
        if (optionalCCVs[j] == allRequiredCCVs[i]) {
          // Remove the duplicate by replacing it with the last element and reducing the length of the array.
          optionalCCVs[j] = optionalCCVs[--newOptionalLength];

          // Since we moved one CCV from optional to required, we can reduce the threshold by one, but not below zero.
          if (optionalThreshold > 0) {
            --optionalThreshold;
          }

          // Break is safe because we asserted no duplicates in _getCCVsFromReceiver.
          break;
        }
        ++j;
      }
    }

    assembly {
      // set the length of the array to the new index which we used to track the number of unique CCVs.
      mstore(optionalCCVs, newOptionalLength)
    }

    // Return the deduplicated required CCVs, the unchanged optional CCVs and the optional threshold.
    return (allRequiredCCVs, optionalCCVs, optionalThreshold);
  }

  /// @notice Ensures that the provided CCVs meet the quorum required by the receiver, pool and lane.
  /// @param sourceChainSelector The source chain selector of the message.
  /// @param receiver The receiver of the message.
  /// @param tokenTransfer The tokens transferred in the message.
  /// @param ccvs The CCVs that provided data for the message.
  /// @param finality The finality requirement of the message.
  /// @return ccvsToQuery The CCVs that need to be queried to verify the message.
  /// @return dataIndexes The indexes of the CCVs in the provided ccvs array that correspond to ccvsToQuery.
  function _ensureCCVQuorumIsReached(
    uint64 sourceChainSelector,
    address receiver,
    MessageV1Codec.TokenTransferV1[] memory tokenTransfer,
    uint16 finality,
    address[] calldata ccvs,
    bool isTokenOnlyTransfer
  ) internal view returns (address[] memory ccvsToQuery, uint256[] memory dataIndexes) {
    (address[] memory requiredCCV, address[] memory optionalCCVs, uint8 optionalThreshold) =
      _getCCVsForMessage(sourceChainSelector, receiver, tokenTransfer, finality, isTokenOnlyTransfer);

    ccvsToQuery = new address[](ccvs.length);
    dataIndexes = new uint256[](ccvs.length);
    uint256 numCCVsToQuery = 0;

    for (uint256 i = 0; i < requiredCCV.length; ++i) {
      bool found = false;
      for (uint256 j = 0; j < ccvs.length; ++j) {
        if (ccvs[j] == requiredCCV[i]) {
          found = true;
          ccvsToQuery[numCCVsToQuery] = ccvs[j];
          dataIndexes[numCCVsToQuery++] = j;
          break;
        }
      }
      if (!found) {
        revert RequiredCCVMissing(requiredCCV[i]);
      }
    }

    uint256 optionalCCVsToFind = optionalThreshold;
    for (uint256 i = 0; i < optionalCCVs.length; ++i) {
      for (uint256 j = 0; j < ccvs.length && optionalCCVsToFind > 0; ++j) {
        if (ccvs[j] == optionalCCVs[i]) {
          optionalCCVsToFind--;

          ccvsToQuery[numCCVsToQuery] = ccvs[j];
          dataIndexes[numCCVsToQuery++] = j;
          break;
        }
      }
    }

    if (optionalCCVsToFind > 0) {
      revert OptionalCCVQuorumNotReached(optionalThreshold, optionalThreshold - optionalCCVsToFind);
    }

    if (numCCVsToQuery != ccvsToQuery.length) {
      // Resize the array to the actual number of CCVs found.
      assembly {
        mstore(ccvsToQuery, numCCVsToQuery)
        mstore(dataIndexes, numCCVsToQuery)
      }
    }
    return (ccvsToQuery, dataIndexes);
  }

  /// @notice Retrieves the required and optional CCVs from a receiver contract. If the receiver does not specify any
  /// CCVs, we fall back to the default CCVs.
  /// @dev This function reverts if the receiver returns duplicates in either the required or optional CCVs.
  /// @param sourceChainSelector The source chain selector.
  /// @param receiver The receiver address.
  /// @return requiredCCV The required CCVs.
  /// @return optionalCCVs The optional CCVs.
  /// @return optionalThreshold The threshold of optional CCVs.
  function _getCCVsFromReceiver(
    uint64 sourceChainSelector,
    address receiver
  ) internal view returns (address[] memory requiredCCV, address[] memory optionalCCVs, uint8 optionalThreshold) {
    // Only query for custom CCVs if the receiver supports the interface..
    if (receiver._supportsInterfaceReverting(type(IAny2EVMMessageReceiverV2).interfaceId)) {
      (requiredCCV, optionalCCVs, optionalThreshold) = IAny2EVMMessageReceiverV2(receiver).getCCVs(sourceChainSelector);

      CCVConfigValidation._assertNoDuplicates(requiredCCV);
      CCVConfigValidation._assertNoDuplicates(optionalCCVs);

      // If the receiver specified empty required and optional CCVs, we fall back to the default CCVs.
      // If they did specify something, we use what they specified.
      if (requiredCCV.length != 0 || optionalThreshold != 0) {
        return (requiredCCV, optionalCCVs, optionalThreshold);
      }
    }

    // Returning new address[](1) means we add the default, as address(0) is the marker for that.
    return (new address[](1), new address[](0), 0);
  }

  /// @notice Retrieves the required CCVs from a pool. If the pool does not specify any CCVs, we fall back to the
  /// default CCVs.
  /// @dev The params passed into getRequiredCCVs could influence the CCVs returned.
  /// @param localToken The local token address.
  /// @param sourceChainSelector The source chain selector.
  /// @param amount The amount of the token to be released/minted.
  /// @param extraData The extra data for the pool.
  /// @return requiredCCV The required CCVs.
  function _getCCVsFromPool(
    address localToken,
    uint64 sourceChainSelector,
    uint256 amount,
    uint16 finality,
    bytes memory extraData
  ) internal view returns (address[] memory requiredCCV) {
    address pool = ITokenAdminRegistry(i_tokenAdminRegistry).getPool(localToken);

    if (pool._supportsInterfaceReverting(type(IPoolV2).interfaceId)) {
      requiredCCV = IPoolV2(pool)
        .getRequiredCCVs(localToken, sourceChainSelector, amount, finality, extraData, IPoolV2.MessageDirection.Inbound);
      CCVConfigValidation._assertNoDuplicates(requiredCCV);
    }

    // If the pool does not specify any CCVs, or the pool does not support the V2 interface, we fall back to the
    // default CCVs. If this wasn't done, any pool not specifying CCVs would allow any arbitrary CCV to mint infinite
    // tokens by fabricating messages. Since CCVs are permissionless, this would mean anyone would be able to mint.
    if (requiredCCV.length == 0) {
      // A list with address(0) indicates that we should use the default CCVs.
      return new address[](1);
    }
    return requiredCCV;
  }

  // ================================================================
  // │                      Tokens and pools                        │
  // ================================================================

  /// @notice Uses a pool to release or mint a token to a receiver address.
  /// @dev The local token address is validated through the TokenAdminRegistry. If, due to some misconfiguration, the
  /// token is unknown to the registry, the offRamp will revert. The tx, and the tokens, can be retrieved by registering
  /// the token on this chain, and re-trying the msg.
  /// @dev Returns the local pool address so that the registry doesn't have to be queried again by executeSingleMessage.
  /// @param tokenTransfer Amount and source data of the token to be released/minted.
  /// @param originalSender The message sender on the source chain.
  /// @param sourceChainSelector The remote source chain selector
  /// @param blockConfirmationRequested Requested block confirmation.
  function _releaseOrMintSingleToken(
    MessageV1Codec.TokenTransferV1 memory tokenTransfer,
    bytes memory originalSender,
    uint64 sourceChainSelector,
    uint16 blockConfirmationRequested
  ) internal returns (Client.EVMTokenAmount memory destTokenAmount, address localPoolAddress) {
    address receiver = address(bytes20(tokenTransfer.tokenReceiver));

    address localToken = address(bytes20(tokenTransfer.destTokenAddress));
    // We check with the token admin registry if the token has a pool on this chain.
    localPoolAddress = ITokenAdminRegistry(i_tokenAdminRegistry).getPool(localToken);
    // This will call the supportsInterface through the ERC165Checker, and not directly on the pool address.
    // This is done to prevent a pool from reverting the entire transaction if it doesn't support the interface.
    // The call gets a max or 30k gas per instance, of which there are three. This means offchain gas estimations should
    // account for 90k gas overhead due to the interface check.
    if (localPoolAddress == address(0)) {
      revert NotACompatiblePool(localPoolAddress);
    }

    Pool.ReleaseOrMintOutV1 memory returnData;

    Pool.ReleaseOrMintInV1 memory releaseOrMintInput = Pool.ReleaseOrMintInV1({
      originalSender: originalSender,
      receiver: receiver,
      sourceDenominatedAmount: tokenTransfer.amount,
      localToken: localToken,
      remoteChainSelector: sourceChainSelector,
      // The source chain has encoded this in the expected format.
      sourcePoolAddress: tokenTransfer.sourcePoolAddress,
      sourcePoolData: tokenTransfer.extraData,
      // All use cases that use offchain token data in IPoolV1 have to upgrade to the modular security interface.
      offchainTokenData: ""
    });

    // This will call the supportsInterface through the ERC165Checker, and not directly on the pool address.
    // This is done to prevent a pool from reverting the entire transaction if it doesn't support the interface.
    // The call gets a max or 30k gas per instance, of which there are three. This means offchain gas estimations should
    // account for 90k gas overhead due to the interface check.
    if (localPoolAddress._supportsInterfaceReverting(type(IPoolV2).interfaceId)) {
      try IPoolV2(localPoolAddress).releaseOrMint(releaseOrMintInput, blockConfirmationRequested) returns (
        Pool.ReleaseOrMintOutV1 memory result
      ) {
        returnData = result;
      } catch (bytes memory err) {
        revert TokenHandlingError(localToken, err);
      }
    } else if (localPoolAddress._supportsInterfaceReverting(Pool.CCIP_POOL_V1)) {
      try IPoolV1(localPoolAddress).releaseOrMint(releaseOrMintInput) returns (Pool.ReleaseOrMintOutV1 memory result) {
        returnData = result;
      } catch (bytes memory err) {
        revert TokenHandlingError(localToken, err);
      }
    } else {
      // If the pool does not support the v1 interface, we revert.
      revert NotACompatiblePool(localPoolAddress);
    }

    return (Client.EVMTokenAmount({token: localToken, amount: returnData.destinationAmount}), localPoolAddress);
  }

  /// @notice Retrieves the balance of a receiver address for a given token.
  /// @param receiver The address to check the balance of.
  /// @param token The token address.
  /// @return balance The balance of the receiver.
  function _getBalanceOfReceiver(
    address receiver,
    address token
  ) internal view returns (uint256) {
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
  /// @return sourceChainSelectors The supported source chain selectors.
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
      SourceChainConfigArgs memory configUpdate = sourceChainConfigUpdates[i];

      if (configUpdate.sourceChainSelector == 0) {
        revert ZeroChainSelectorNotAllowed();
      }
      if (address(configUpdate.router) == address(0) || configUpdate.defaultCCVs.length == 0) {
        revert ZeroAddressNotAllowed();
      }

      for (uint256 j = 0; j < configUpdate.defaultCCVs.length; ++j) {
        if (configUpdate.defaultCCVs[j] == address(0)) {
          revert ZeroAddressNotAllowed();
        }
      }
      for (uint256 j = 0; j < configUpdate.laneMandatedCCVs.length; ++j) {
        if (configUpdate.laneMandatedCCVs[j] == address(0)) {
          revert ZeroAddressNotAllowed();
        }
      }
      CCVConfigValidation._validateDefaultAndMandatedCCVs(configUpdate.defaultCCVs, configUpdate.laneMandatedCCVs);

      SourceChainConfig storage currentConfig = s_sourceChainConfigs[configUpdate.sourceChainSelector];
      EnumerableSet.Bytes32Set storage allowedOnRampHashes = s_allowedOnRampHashes[configUpdate.sourceChainSelector];

      // Remove all current onRamps.
      allowedOnRampHashes.clear();

      // Populate allowed onRamps. This list could be empty, which would mean no onRamps are allowed and the lane is
      // disabled, even for existing messages.
      for (uint256 j = 0; j < configUpdate.onRamps.length; ++j) {
        bytes memory onRamp = configUpdate.onRamps[j];
        bytes32 onRampHash = keccak256(onRamp);
        if (onRamp.length == 0 || onRampHash == EMPTY_ENCODED_ADDRESS_HASH) {
          revert ZeroAddressNotAllowed();
        }
        allowedOnRampHashes.add(onRampHash);
      }

      currentConfig.onRamps = configUpdate.onRamps;
      currentConfig.isEnabled = configUpdate.isEnabled;
      currentConfig.router = configUpdate.router;
      currentConfig.defaultCCVs = configUpdate.defaultCCVs;
      currentConfig.laneMandatedCCVs = configUpdate.laneMandatedCCVs;

      // We don't need to check the return value, as inserting the item twice has no effect.
      s_sourceChainSelectors.add(configUpdate.sourceChainSelector);

      emit SourceChainConfigSet(configUpdate.sourceChainSelector, configUpdate);
    }
  }

  /// @notice hook for applying custom logic to the input message before executeSingleMessage()
  /// @param message initial message
  /// @return transformedMessage modified message
  function _beforeExecuteSingleMessage(
    MessageV1Codec.MessageV1 memory message
  ) internal view virtual returns (MessageV1Codec.MessageV1 memory transformedMessage) {
    return message;
  }
}
