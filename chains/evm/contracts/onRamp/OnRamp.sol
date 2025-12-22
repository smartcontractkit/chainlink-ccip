// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../interfaces/ICrossChainVerifierResolver.sol";
import {ICrossChainVerifierV1} from "../interfaces/ICrossChainVerifierV1.sol";
import {IEVM2AnyOnRampClient} from "../interfaces/IEVM2AnyOnRampClient.sol";
import {IExecutor} from "../interfaces/IExecutor.sol";
import {IFeeQuoter} from "../interfaces/IFeeQuoter.sol";
import {IPoolV1} from "../interfaces/IPool.sol";
import {IPoolV2} from "../interfaces/IPoolV2.sol";
import {IRMNRemote} from "../interfaces/IRMNRemote.sol";
import {IRouter} from "../interfaces/IRouter.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {CCVConfigValidation} from "../libraries/CCVConfigValidation.sol";
import {Client} from "../libraries/Client.sol";
import {ExtraArgsCodec} from "../libraries/ExtraArgsCodec.sol";
import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {Pool} from "../libraries/Pool.sol";
import {USDPriceWith18Decimals} from "../libraries/USDPriceWith18Decimals.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

contract OnRamp is IEVM2AnyOnRampClient, ITypeAndVersion, Ownable2StepMsgSender {
  using SafeERC20 for IERC20;
  using EnumerableSet for EnumerableSet.AddressSet;
  using USDPriceWith18Decimals for uint224;

  error CannotSendZeroTokens();
  error DestinationChainNotSupportedByCCV(address ccvAddress, uint64 destChainSelector);
  error UnsupportedToken(address token);
  error CanOnlySendOneTokenPerMessage();
  error MustBeCalledByRouter();
  error RouterMustSetOriginalSender();
  error InvalidConfig();
  error CursedByRMN(uint64 destChainSelector);
  error GetSupportedTokensFunctionalityRemovedCheckAdminRegistry();
  error InvalidDestChainConfig(uint64 destChainSelector);
  error ReentrancyGuardReentrantCall();
  error DestinationChainNotSupported(uint64 destChainSelector);
  error InvalidDestChainAddress(bytes destChainAddress);
  error CustomBlockConfirmationNotSupportedOnPoolV1();
  error TokenArgsNotSupportedOnPoolV1();
  error InsufficientFeeTokenAmount();
  error TokenReceiverNotAllowed(uint64 destChainSelector);
  error SourceTokenDataTooLarge(address token, uint256 actualLength, uint32 maxLength);

  event ConfigSet(StaticConfig staticConfig, DynamicConfig dynamicConfig);
  event DestChainConfigSet(uint64 indexed destChainSelector, uint64 messageNumber, DestChainConfigArgs config);
  event FeeTokenWithdrawn(address indexed feeAggregator, address indexed feeToken, uint256 amount);
  event CCIPMessageSent(
    uint64 indexed destChainSelector,
    uint64 indexed messageNumber,
    bytes32 indexed messageId,
    address feeToken,
    bytes encodedMessage,
    Receipt[] receipts,
    bytes[] verifierBlobs
  );

  /// @dev A helper struct that holds data for the CCIPMessageSent event to avoid stack too deep errors.
  struct CCIPMessageSentEventData {
    bytes encodedMessage;
    Receipt[] receipts;
    bytes[] verifierBlobs;
  }

  /// @dev Struct that contains the static configuration.
  // solhint-disable-next-line gas-struct-packing
  struct StaticConfig {
    uint64 chainSelector; // ────╮ Local chain selector.
    IRMNRemote rmnRemote; // ────╯ RMN remote address.
    address tokenAdminRegistry; // Token admin registry address.
  }

  /// @dev Struct that contains the dynamic configuration
  // solhint-disable-next-line gas-struct-packing
  struct DynamicConfig {
    address feeQuoter; // ───────────╮ FeeQuoter address.
    bool reentrancyGuardEntered; // ─╯ Reentrancy protection.
    address feeAggregator; // Fee aggregator address.
  }

  /// @dev Struct to hold the configs for a single destination chain.
  struct DestChainConfig {
    IRouter router; // ────────────╮ Local router address  that is allowed to send messages to the destination chain.
    // The last used message number. This is zero in the case where no messages have yet been sent.
    // 0 is not a valid message number for any real transaction as this value will be incremented before use.
    uint64 messageNumber; //       │
    uint8 addressBytesLength; //   │ The length of an address on this chain in bytes, e.g. 20 for EVM, 32 for SVM.
    uint16 networkFeeUSDCents; //  │ Network fee in USD cents for messages to this destination chain.
    bool tokenReceiverAllowed; // ─╯ Whether specifying `tokenReceiver` in extraArgs is allowed at all.
    uint32 baseExecutionGasCost; // Base gas cost for executing a message on the destination chain.
    address defaultExecutor; // Default executor to use for messages to this destination chain.
    address[] laneMandatedCCVs; // Required CCVs to use for all messages to this destination chain.
    address[] defaultCCVs; // Default CCVs to use for messages to this destination chain.
    bytes offRamp; // Destination OffRamp address, NOT abi encoded but raw bytes.
  }

  /// @dev Same as DestChainConfig but with the destChainSelector so that an array of these can be passed in the
  /// constructor and the applyDestChainConfigUpdates function.
  // solhint-disable gas-struct-packing
  struct DestChainConfigArgs {
    uint64 destChainSelector; // Destination chain selector.
    IRouter router; //  Source router address  that is allowed to send messages to the destination chain.
    uint8 addressBytesLength; // The length of an address on this chain in bytes, e.g. 20 for EVM, 32 for SVM.
    uint16 networkFeeUSDCents; // Network fee in USD cents for messages to this destination chain.
    bool tokenReceiverAllowed; // Whether specifying `tokenReceiver` in extraArgs is allowed at all.
    uint32 baseExecutionGasCost; // Base gas cost for executing a message on the destination chain.
    address[] defaultCCVs; // Default CCVs to use for messages to this destination chain.
    address[] laneMandatedCCVs; // Required CCVs to use for all messages to this destination chain.
    address defaultExecutor; // If no executor is specified in the extraArgs, this executor will be used.
    bytes offRamp; // Destination OffRamp address, NOT abi encoded but raw bytes.
  }

  /// @notice Receipt structure used to record gas limits and fees for message processing.
  /// @dev The ordering of receipts in a message is as follows:
  /// - Verifier receipts in the order of the CCV list.
  /// - Token transfer receipt (if any tokens are being transferred).
  /// - Executor receipt.
  /// - Network fee receipt.
  struct Receipt {
    // The address of the entity that issued the receipt. For token receipts this is the token address, not the pool.
    // for verifiers and executors, this is the user specified value, even if the call is ultimately handled by some
    // underlying contract.
    address issuer; // ───────────╮
    uint32 destGasLimit; //       │ The gas limit for the actions taken on the destination chain for this entity.
    uint32 destBytesOverhead; // ─╯ The byte overhead for the actions taken on the destination chain for this entity.
    // The fee amount for this entity, in smallest denomination of the fee token.
    // NOTE: While building receipts in `_getReceipts`, this field is temporarily populated with a USD-cent value
    // (and converted to fee token amount later in the same function).
    uint256 feeTokenAmount;
    bytes extraArgs; // Extra args that have been passed in on the source chain. May be empty.
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "OnRamp 1.7.0-dev";
  /// @dev The chain ID of the source chain that this contract is deployed to.
  uint64 private immutable i_localChainSelector;
  /// @dev The rmn contract.
  IRMNRemote private immutable i_rmnRemote;
  /// @dev The address of the token admin registry.
  address private immutable i_tokenAdminRegistry;

  // DYNAMIC CONFIG
  /// @dev The dynamic config for the onRamp.
  DynamicConfig private s_dynamicConfig;

  /// @dev The destination chain specific configs.
  mapping(uint64 destChainSelector => DestChainConfig destChainConfig) internal s_destChainConfigs;

  constructor(StaticConfig memory staticConfig, DynamicConfig memory dynamicConfig) {
    if (
      staticConfig.chainSelector == 0 || address(staticConfig.rmnRemote) == address(0)
        || staticConfig.tokenAdminRegistry == address(0)
    ) {
      revert InvalidConfig();
    }

    i_localChainSelector = staticConfig.chainSelector;
    i_rmnRemote = staticConfig.rmnRemote;
    i_tokenAdminRegistry = staticConfig.tokenAdminRegistry;

    _setDynamicConfig(dynamicConfig);
  }

  // ================================================================
  // │                          Messaging                           │
  // ================================================================

  /// @notice Gets the next message number to be used in the onRamp.
  /// @param destChainSelector The destination chain selector.
  /// @return nextMessageNumber The next message number to be used.
  function getExpectedNextMessageNumber(
    uint64 destChainSelector
  ) external view returns (uint64) {
    return s_destChainConfigs[destChainSelector].messageNumber + 1;
  }

  /// @inheritdoc IEVM2AnyOnRampClient
  /// @param destChainSelector The destination chain selector.
  /// @param message The message being sent.
  /// @param feeTokenAmount The amount of fee token provided by the router for this message.
  /// @param originalSender The original sender of the message on the source chain.
  function forwardFromRouter(
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message,
    uint256 feeTokenAmount,
    address originalSender
  ) external returns (bytes32 messageId) {
    if (i_rmnRemote.isCursed(bytes16(uint128(destChainSelector)))) {
      revert CursedByRMN(destChainSelector);
    }
    // We rely on a reentrancy guard here due to the untrusted calls performed to the pools. This enables some
    // optimizations by not following the CEI pattern.
    if (s_dynamicConfig.reentrancyGuardEntered) revert ReentrancyGuardReentrantCall();
    s_dynamicConfig.reentrancyGuardEntered = true;

    DestChainConfig storage destChainConfig = s_destChainConfigs[destChainSelector];

    // NOTE: The router is expected to call `getFee` prior to `forwardFromRouter`.
    // This function still performs its own checks (e.g. router/originalSender, dest address validation, CCV list
    // finalization) but relies on the router's pre-checks to avoid duplicating validation work.
    // Validate originalSender is set and allowed. Not validated in `getFee` since it is not user-driven.
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    // Router address may be zero intentionally to pause, which should stop all messages.
    if (msg.sender != address(destChainConfig.router)) revert MustBeCalledByRouter();

    // 1. parse extraArgs.

    ExtraArgsCodec.GenericExtraArgsV3 memory resolvedExtraArgs = _parseExtraArgsWithDefaults(
      destChainSelector,
      destChainConfig,
      message.extraArgs,
      (message.data.length == 0 && message.tokenAmounts.length > 0)
    );

    MessageV1Codec.MessageV1 memory newMessage = MessageV1Codec.MessageV1({
      sourceChainSelector: i_localChainSelector,
      destChainSelector: destChainSelector,
      messageNumber: ++destChainConfig.messageNumber,
      executionGasLimit: 0, // Populated after getting receipts.
      ccipReceiveGasLimit: resolvedExtraArgs.gasLimit,
      finality: resolvedExtraArgs.blockConfirmations,
      ccvAndExecutorHash: bytes32(0), // Will be set after CCV list is finalized.
      onRampAddress: abi.encode(address(this)), // Source address, so abi encoded.
      offRampAddress: destChainConfig.offRamp, // Dest address, so unpadded bytes.
      sender: abi.encode(originalSender), // Source address, so abi encoded.
      receiver: _validateDestChainAddress(message.receiver, destChainConfig.addressBytesLength), // Dest address, so unpadded bytes.
      // Executor args hold security critical execution args, like Solana accounts or Sui object IDs. Because of this,
      // they have to be part of the message that is signed off on by the verifiers.
      destBlob: resolvedExtraArgs.executorArgs,
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](message.tokenAmounts.length), //  values are populated with _lockOrBurnSingleToken.
      data: message.data
    });

    // 2. get pool params, this potentially mutates the CCV list.

    {
      address[] memory poolRequiredCCVs = new address[](0);
      if (message.tokenAmounts.length != 0) {
        poolRequiredCCVs = _getCCVsForPool(
          destChainSelector,
          message.tokenAmounts[0].token,
          message.tokenAmounts[0].amount,
          resolvedExtraArgs.blockConfirmations,
          resolvedExtraArgs.tokenArgs
        );
      }
      (resolvedExtraArgs.ccvs, resolvedExtraArgs.ccvArgs) = _mergeCCVLists(
        resolvedExtraArgs.ccvs, resolvedExtraArgs.ccvArgs, destChainConfig.laneMandatedCCVs, poolRequiredCCVs
      );
    }

    // Set the ccvAndExecutorHash now that the CCV list is finalized.
    newMessage.ccvAndExecutorHash =
      MessageV1Codec._computeCCVAndExecutorHash(resolvedExtraArgs.ccvs, resolvedExtraArgs.executor);

    // 3. getFee on all verifiers, pool and executor.

    CCIPMessageSentEventData memory eventData;
    {
      uint256 computedFeeTokenAmount;
      // Populate receipts for verifiers, pool (if applicable), executor and network fee in that order.
      (eventData.receipts, newMessage.executionGasLimit, computedFeeTokenAmount) =
        _getReceipts(destChainSelector, destChainConfig.networkFeeUSDCents, message, resolvedExtraArgs);

      // Any third party (ccv, pool, or executor) could theoretically return different values for getFee on every call.
      // Without this guard, the onRamp would pay until it ran out of funds, potentially using accrued protocol fees to
      // pay for further calls.
      if (computedFeeTokenAmount > feeTokenAmount) {
        revert InsufficientFeeTokenAmount();
      }
      _distributeFees(message, eventData.receipts);
    }

    // 4. lockOrBurn.

    if (message.tokenAmounts.length != 0) {
      if (message.tokenAmounts.length != 1) revert CanOnlySendOneTokenPerMessage();
      newMessage.tokenTransfer[0] = _lockOrBurnSingleToken(
        message.tokenAmounts[0],
        destChainSelector,
        // At this point `resolvedExtraArgs.tokenReceiver` and `message.receiver` are the raw inputs provided by caller.
        // The receiver is passed as-is to the TokenPool (which expects abi-encoded format for EVM source chains),
        // then validated and trimmed to minimal bytes for destination chain encoding in the message.
        resolvedExtraArgs.tokenReceiver.length > 0 ? resolvedExtraArgs.tokenReceiver : message.receiver,
        originalSender,
        resolvedExtraArgs.blockConfirmations,
        resolvedExtraArgs.tokenArgs
      );

      // Enforce that the token pool payload (`destPoolData` -> TokenTransferV1.extraData) is not larger than the
      // bytes overhead that was quoted and paid for in the token transfer receipt.
      uint32 maxExtraDataLength = eventData.receipts[resolvedExtraArgs.ccvs.length].destBytesOverhead;
      uint256 actualExtraDataLength = newMessage.tokenTransfer[0].extraData.length;
      if (actualExtraDataLength > maxExtraDataLength) {
        revert SourceTokenDataTooLarge(message.tokenAmounts[0].token, actualExtraDataLength, maxExtraDataLength);
      }
    }

    // 5. encode message and calculate messageId.

    eventData.encodedMessage = MessageV1Codec._encodeMessageV1(newMessage);
    messageId = keccak256(eventData.encodedMessage);

    eventData.verifierBlobs = new bytes[](resolvedExtraArgs.ccvs.length);

    // 6. call each verifier.
    for (uint256 i = 0; i < resolvedExtraArgs.ccvs.length; ++i) {
      address implAddress = ICrossChainVerifierResolver(resolvedExtraArgs.ccvs[i]).getOutboundImplementation(
        destChainSelector, resolvedExtraArgs.ccvArgs[i]
      );
      if (implAddress == address(0)) {
        revert DestinationChainNotSupportedByCCV(resolvedExtraArgs.ccvs[i], destChainSelector);
      }
      // NOTE: this verifier blob is *not* the same as the verifier data that will be delivered to the destination chain.
      // This field is meant for the offchain verifier. The verifier *may* submit this data as part of the verifier data
      // on the destination chain, but it may also choose to submit different data or no data at all. This means there
      // should be no check on the length of this and the CCV bytes overhead, as there is no relationship between the two.
      eventData.verifierBlobs[i] = ICrossChainVerifierV1(implAddress).forwardToVerifier(
        newMessage, messageId, message.feeToken, feeTokenAmount, resolvedExtraArgs.ccvArgs[i]
      );
    }

    // 7. emit event.
    emit CCIPMessageSent({
      destChainSelector: destChainSelector,
      messageNumber: newMessage.messageNumber,
      messageId: messageId,
      feeToken: message.feeToken,
      encodedMessage: eventData.encodedMessage,
      receipts: eventData.receipts,
      verifierBlobs: eventData.verifierBlobs
    });

    s_dynamicConfig.reentrancyGuardEntered = false;

    return messageId;
  }

  /// @notice Distributes the fee token to each receipt issuer.
  /// @dev Token pool receipt payments are routed to the pool only if it supports IPoolV2 interface.
  /// @param message The message containing the fee token and token transfer info.
  /// @param receipts The receipts to pay out, in protocol-defined order.
  function _distributeFees(Client.EVM2AnyMessage calldata message, Receipt[] memory receipts) internal {
    IERC20 feeToken = IERC20(message.feeToken);
    uint256 tokenReceiptIndex = type(uint256).max;
    if (message.tokenAmounts.length > 0) {
      // Layout with tokens: verifiers..., token, executor, network fee.
      tokenReceiptIndex = receipts.length - 3;
      address tokenPool = receipts[tokenReceiptIndex].issuer;
      // In case the token pool supports the IPoolV2 interface, the pool receive the fee share as fee handling logic built in.
      // V1 pools intentionally leave the balance sitting on the OnRamp so it can be withdrawn later.
      if (IERC165(tokenPool).supportsInterface(type(IPoolV2).interfaceId)) {
        feeToken.safeTransfer(address(tokenPool), receipts[tokenReceiptIndex].feeTokenAmount);
      }
    }
    // We iterate up to receipts.length - 1 to skip the network fee receipt which must remain in the onRamp.
    uint256 networkFeeReceiptIndex = receipts.length - 1;
    for (uint256 i = 0; i < networkFeeReceiptIndex; ++i) {
      // We skip fee distribution if:
      // - The fee is 0.
      // - The receipt is the token receipt as that's handled above.
      // - The network fee receipt, as explained above.
      if (i == tokenReceiptIndex) continue;
      uint256 receiptFee = receipts[i].feeTokenAmount;
      if (receiptFee == 0) continue;
      feeToken.safeTransfer(receipts[i].issuer, receiptFee);
    }
  }

  /// @notice Merges lane mandated and pool required CCVs with user-provided CCVs.
  /// @dev This function assumes there are no duplicates in the userRequestedOrDefaultCCVs list.
  /// @dev There is no protocol-level requirement on the ordering of CCVs in the final list, but for determinism we
  /// process user requested first, then lane-mandated second, pool-required last.
  /// @param userRequestedOrDefaultCCVs User-provided required CCV addresses. Can not be empty, as defaults are applied earlier if needed.
  /// @param userRequestedOrDefaultCCVArgs User-provided CCV arguments, parallel to userRequestedOrDefaultCCVs.
  /// @param laneMandatedCCVs Lane mandated CCVs are always added, regardless of what a user/pool chooses. Can be empty.
  /// @param poolRequiredCCVs Pool-specific required CCVs.
  /// @return ccvs Updated list of CCV addresses.
  /// @return ccvArgs Updated list of CCV arguments, parallel to ccvs.
  function _mergeCCVLists(
    address[] memory userRequestedOrDefaultCCVs,
    bytes[] memory userRequestedOrDefaultCCVArgs,
    address[] memory laneMandatedCCVs,
    address[] memory poolRequiredCCVs
  ) internal pure returns (address[] memory ccvs, bytes[] memory ccvArgs) {
    // Maximum possible CCVs: user + lane + pool.
    uint256 totalCCVs = userRequestedOrDefaultCCVs.length + laneMandatedCCVs.length + poolRequiredCCVs.length;
    ccvs = new address[](totalCCVs);
    ccvArgs = new bytes[](totalCCVs);
    uint256 toBeAddedIndex = 0;

    // First add all user requested CCVs.
    for (uint256 i = 0; i < userRequestedOrDefaultCCVs.length; ++i) {
      ccvs[toBeAddedIndex] = userRequestedOrDefaultCCVs[i];
      ccvArgs[toBeAddedIndex++] = userRequestedOrDefaultCCVArgs[i];
    }
    // Add lane mandated CCVs, skipping duplicates.
    for (uint256 i = 0; i < laneMandatedCCVs.length; ++i) {
      address laneMandatedCCV = laneMandatedCCVs[i];
      bool found = false;
      for (uint256 j = 0; j < toBeAddedIndex; ++j) {
        if (ccvs[j] == laneMandatedCCV) {
          found = true;
          break;
        }
      }
      if (!found) {
        ccvs[toBeAddedIndex++] = laneMandatedCCV;
      }
    }
    // Add pool required CCVs, skipping duplicates.
    for (uint256 i = 0; i < poolRequiredCCVs.length; ++i) {
      address poolRequiredCCV = poolRequiredCCVs[i];
      bool found = false;
      for (uint256 j = 0; j < toBeAddedIndex; ++j) {
        if (ccvs[j] == poolRequiredCCV) {
          found = true;
          break;
        }
      }
      if (!found) {
        ccvs[toBeAddedIndex++] = poolRequiredCCV;
      }
    }

    // Resize both arrays to the actual number of CCVs added.
    assembly {
      mstore(ccvs, toBeAddedIndex)
      mstore(ccvArgs, toBeAddedIndex)
    }

    return (ccvs, ccvArgs);
  }

  /// @notice Validates a destination-chain address and strips leading ABI padding for sub-32-byte addresses.
  /// @dev Assumes `addressBytesLength < 32` only in the padded branch; otherwise length must match exactly.
  /// @param rawAddress The address bytes (may be 32-byte ABI-encoded or exact-length raw bytes).
  /// @param addressBytesLength The expected address length on the destination chain.
  /// @return validatedAddress The validated address with any leading padding removed.
  function _validateDestChainAddress(
    bytes memory rawAddress,
    uint256 addressBytesLength
  ) internal pure returns (bytes memory validatedAddress) {
    uint256 len = rawAddress.length;
    if (addressBytesLength < 32 && len == 32) {
      uint256 word;
      // assembly equivalent: word = uint256(bytes32(rawAddress));
      assembly {
        word := mload(add(rawAddress, 32))
      }
      // Shift out the actual address bytes; any residue means non-zero padding.
      if (word >> (addressBytesLength * 8) != 0) {
        revert InvalidDestChainAddress(rawAddress);
      }
      validatedAddress = new bytes(addressBytesLength);
      // assembly equivalent:
      //  uint256 offset = 32 - addressBytesLength;
      //  for (uint256 i = 0; i < addressBytesLength; ++i) {
      //    validatedAddress[i] = rawAddress[offset + i];
      //  }
      assembly {
        let shift := mul(sub(32, addressBytesLength), 8)
        mstore(add(validatedAddress, 32), shl(shift, word))
      }
      return validatedAddress;
    }

    if (len != addressBytesLength) {
      revert InvalidDestChainAddress(rawAddress);
    }
    return rawAddress;
  }

  /// @notice Parses and validates extra arguments, applying defaults from destination chain configuration.
  /// The function ensures all messages have the required CCVs and executor needed for processing,
  /// even when users don't explicitly specify them.
  /// @dev `tokenReceiver` is NOT validated here, and is validated in `_lockOrBurnSingleToken`.
  /// @param destChainSelector The destination chain selector.
  /// @param destChainConfig Configuration for the destination chain including default values.
  /// @param extraArgs User-provided extra arguments in either V3 or legacy format.
  /// @param isTokenTransferWithoutData Indicates if the message has no data but includes a token transfer. This is meant to
  /// signal token-only transfers to avoid adding default CCVs when not needed.
  /// @return resolvedArgs Complete EVMExtraArgsV3 struct with all defaults applied.
  function _parseExtraArgsWithDefaults(
    uint64 destChainSelector,
    DestChainConfig memory destChainConfig,
    bytes calldata extraArgs,
    bool isTokenTransferWithoutData
  ) internal view returns (ExtraArgsCodec.GenericExtraArgsV3 memory resolvedArgs) {
    // If ExtraArgsV3 are provided, decode them.
    if (extraArgs.length >= 4 && bytes4(extraArgs[:4]) == ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_TAG) {
      resolvedArgs = ExtraArgsCodec._decodeGenericExtraArgsV3(extraArgs);

      // We need to ensure no duplicate CCVs are present in the ccv list.
      CCVConfigValidation._assertNoDuplicates(resolvedArgs.ccvs);
    } else {
      // Populate the fields that could be present in legacy extraArgs.
      (resolvedArgs.tokenReceiver, resolvedArgs.gasLimit, resolvedArgs.executorArgs) =
        IFeeQuoter(s_dynamicConfig.feeQuoter).resolveLegacyArgs(destChainSelector, extraArgs);
    }

    // We remove the need for sender/receiver CCVs if the transfer is a pure token transfer.
    //
    // A pure token transfer is defined as:
    // - receiver callback gas limit is 0
    // - the message has no data
    // - the message sends a token
    //
    // This has always existed in CCIP: the earliest versions skipped calling the receiver when this was true, and on
    // the destination chain also checked (via ERC165) whether the receiver is a contract and supports the required
    // interfaces. Those destination-side checks are not available on the source chain, so we use the three
    // conditions above as the source-chain definition of a pure token transfer.
    //
    // When a transfer is pure, we do not add sender/receiver (default) CCVs. This is safe because the only entity at
    // risk is the token issuer, who already defines their required CCVs (or falls back to defaults), so their risk is
    // not increased by omitting sender/receiver CCVs in this case. This enables token-only transfers to use
    // token-specific CCVs (e.g. CCTP) without the user having to know about CCVs or use the new extraArgs format.
    //
    // For example, token-only USDC transfers can use only CCTP (without committee verification), since CCTP is fully
    // trusted for that token flow.
    if (resolvedArgs.ccvs.length == 0 && !(isTokenTransferWithoutData && resolvedArgs.gasLimit == 0)) {
      resolvedArgs.ccvs = destChainConfig.defaultCCVs;
      resolvedArgs.ccvArgs = new bytes[](resolvedArgs.ccvs.length);
    }

    // Normalize and validate tokenReceiver if specified.
    if (resolvedArgs.tokenReceiver.length != 0) {
      // Some lanes disallow specifying tokenReceiver entirely.
      if (!destChainConfig.tokenReceiverAllowed) {
        revert TokenReceiverNotAllowed(destChainSelector);
      }
    }

    // When users don't specify an executor, default executor is chosen.
    if (resolvedArgs.executor == address(0)) {
      resolvedArgs.executor = destChainConfig.defaultExecutor;
    }

    return resolvedArgs;
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Returns the static onRamp config.
  /// @return staticConfig the static configuration.
  function getStaticConfig() public view returns (StaticConfig memory) {
    return StaticConfig({
      chainSelector: i_localChainSelector,
      rmnRemote: i_rmnRemote,
      tokenAdminRegistry: i_tokenAdminRegistry
    });
  }

  /// @notice Returns the dynamic onRamp config.
  /// @return dynamicConfig the dynamic configuration.
  function getDynamicConfig() external view returns (DynamicConfig memory dynamicConfig) {
    return s_dynamicConfig;
  }

  /// @notice Sets the dynamic configuration.
  /// @param dynamicConfig The configuration.
  function setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) external onlyOwner {
    _setDynamicConfig(dynamicConfig);
  }

  /// @notice Internal version of setDynamicConfig to allow for reuse in the constructor.
  /// @param dynamicConfig The configuration.
  function _setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) internal {
    if (
      dynamicConfig.feeQuoter == address(0) || dynamicConfig.feeAggregator == address(0)
        || dynamicConfig.reentrancyGuardEntered
    ) revert InvalidConfig();

    s_dynamicConfig = dynamicConfig;

    emit ConfigSet(getStaticConfig(), dynamicConfig);
  }

  /// @notice Updates destination chains specific configs.
  /// @param destChainConfigArgs Array of destination chain specific configs.
  function applyDestChainConfigUpdates(
    DestChainConfigArgs[] calldata destChainConfigArgs
  ) external onlyOwner {
    for (uint256 i = 0; i < destChainConfigArgs.length; ++i) {
      DestChainConfigArgs calldata destChainConfigArg = destChainConfigArgs[i];
      uint64 destChainSelector = destChainConfigArg.destChainSelector;

      if (
        destChainSelector == 0 || destChainSelector == i_localChainSelector
          || destChainConfigArg.addressBytesLength == 0 || destChainConfigArg.baseExecutionGasCost == 0
      ) {
        revert InvalidDestChainConfig(destChainSelector);
      }

      // Ensure at least one default or mandated CCV exists, and check for duplicates or zero addresses in both sets.
      CCVConfigValidation._validateDefaultAndMandatedCCVs(
        destChainConfigArg.defaultCCVs, destChainConfigArg.laneMandatedCCVs
      );

      DestChainConfig storage destChainConfig = s_destChainConfigs[destChainSelector];
      // The router can be zero to pause the destination chain.
      destChainConfig.router = destChainConfigArg.router;
      destChainConfig.addressBytesLength = destChainConfigArg.addressBytesLength;
      destChainConfig.networkFeeUSDCents = destChainConfigArg.networkFeeUSDCents;
      destChainConfig.tokenReceiverAllowed = destChainConfigArg.tokenReceiverAllowed;
      destChainConfig.baseExecutionGasCost = destChainConfigArg.baseExecutionGasCost;
      destChainConfig.defaultCCVs = destChainConfigArg.defaultCCVs;
      destChainConfig.laneMandatedCCVs = destChainConfigArg.laneMandatedCCVs;
      // Require a default executor so messages that rely on older/defaulted args still resolve to a concrete
      // executor. A zero executor would break backward compatibility and cause otherwise-valid traffic to revert.
      if (destChainConfigArg.defaultExecutor == address(0)) revert InvalidConfig();
      destChainConfig.defaultExecutor = destChainConfigArg.defaultExecutor;

      // Make sure that offRamp length matches addressBytesLength.
      if (destChainConfigArg.offRamp.length != destChainConfigArg.addressBytesLength) {
        revert InvalidDestChainAddress(destChainConfigArg.offRamp);
      }
      destChainConfig.offRamp = destChainConfigArg.offRamp;

      emit DestChainConfigSet(destChainSelector, destChainConfig.messageNumber, destChainConfigArg);
    }
  }

  /// @notice get ChainConfig configured for the DestinationChainSelector.
  /// @param destChainSelector The destination chain selector.
  /// @return destChainConfig The destination chain configuration.
  function getDestChainConfig(
    uint64 destChainSelector
  ) external view returns (DestChainConfig memory destChainConfig) {
    return s_destChainConfigs[destChainSelector];
  }

  // ================================================================
  // │                      Tokens and pools                        │
  // ================================================================

  /// @inheritdoc IEVM2AnyOnRampClient
  /// @param sourceToken The source token.
  function getPoolBySourceToken(
    uint64,
    /*destChainSelector*/
    IERC20 sourceToken
  ) public view returns (IPoolV1) {
    return IPoolV1(ITokenAdminRegistry(i_tokenAdminRegistry).getPool(address(sourceToken)));
  }

  /// @inheritdoc IEVM2AnyOnRampClient
  function getSupportedTokens(
    uint64 // destChainSelector
  ) external pure returns (address[] memory) {
    revert GetSupportedTokensFunctionalityRemovedCheckAdminRegistry();
  }

  /// @notice Uses a pool to lock or burn a token and returns MessageV1 token transfer data.
  /// @param tokenAndAmount Token address and amount to lock or burn.
  /// @param destChainSelector Target destination chain selector of the message.
  /// @param receiver Message receiver in abi-encoded format (as expected by the pool on EVM source chains).
  /// @param originalSender Message sender.
  /// @param blockConfirmationRequested Requested block confirmation.
  /// @param tokenArgs Additional token arguments from the message.
  /// @return TokenTransferV1 token transfer encoding for MessageV1.
  function _lockOrBurnSingleToken(
    Client.EVMTokenAmount memory tokenAndAmount,
    uint64 destChainSelector,
    bytes memory receiver,
    address originalSender,
    uint16 blockConfirmationRequested,
    bytes memory tokenArgs
  ) internal returns (MessageV1Codec.TokenTransferV1 memory) {
    if (tokenAndAmount.amount == 0) revert CannotSendZeroTokens();

    IPoolV1 sourcePool = getPoolBySourceToken(destChainSelector, IERC20(tokenAndAmount.token));
    // We don't have to check if it supports the pool version in a non-reverting way here because
    // if we revert here, there is no effect on CCIP. Therefore we directly call the supportsInterface
    // function and not through the ERC165Checker.
    if (address(sourcePool) == address(0) || !sourcePool.supportsInterface(Pool.CCIP_POOL_V1)) {
      revert UnsupportedToken(tokenAndAmount.token);
    }

    // For v1 pools, the destination amount is set equal to the source amount.
    // For v2 pools, the destination amount may be modified in the following logic.
    uint256 destTokenAmount = tokenAndAmount.amount;
    Pool.LockOrBurnOutV1 memory poolReturnData;

    {
      Pool.LockOrBurnInV1 memory lockOrBurnInput = Pool.LockOrBurnInV1({
        receiver: receiver,
        remoteChainSelector: destChainSelector,
        originalSender: originalSender,
        amount: tokenAndAmount.amount,
        localToken: tokenAndAmount.token
      });

      // If the pool declares support for IPoolV2, it can handle `finality` and `tokenArgs`.
      // Use the V2 overload which returns a potentially adjusted destination amount.
      if (IERC165(address(sourcePool)).supportsInterface(type(IPoolV2).interfaceId)) {
        (poolReturnData, destTokenAmount) =
          IPoolV2(address(sourcePool)).lockOrBurn(lockOrBurnInput, blockConfirmationRequested, tokenArgs);
      } else {
        // V1 pools don't understand `blockConfirmationRequested`/`tokenArgs`.
        // We enforce default for `blockConfirmationRequested` and no `tokenArgs` to avoid silent mis-interpretation.
        if (blockConfirmationRequested != 0) {
          revert CustomBlockConfirmationNotSupportedOnPoolV1();
        }
        if (tokenArgs.length != 0) {
          revert TokenArgsNotSupportedOnPoolV1();
        }
        poolReturnData = sourcePool.lockOrBurn(lockOrBurnInput);
      }
    }

    return MessageV1Codec.TokenTransferV1({
      amount: destTokenAmount,
      sourcePoolAddress: abi.encode(sourcePool), // Source address, so abi encoded.
      sourceTokenAddress: abi.encode(tokenAndAmount.token), // Source address, so abi encoded.
      destTokenAddress: _validateDestChainAddress(
        poolReturnData.destTokenAddress, s_destChainConfigs[destChainSelector].addressBytesLength
      ), // Dest address so unpadded bytes.
      tokenReceiver: _validateDestChainAddress(receiver, s_destChainConfigs[destChainSelector].addressBytesLength), // Dest address so unpadded bytes.
      extraData: poolReturnData.destPoolData
    });
  }

  /// @notice Gets the required CCVs from the pool for token transfers.
  /// @dev Resolves address(0) returned by the pool into the destination defaults.
  /// If the pool does not specify any CCVs, we fall back to the default CCVs.
  /// @param destChainSelector The destination chain selector.
  /// @param token The token address being transferred.
  /// @param amount The amount of tokens being transferred.
  /// @param finality The finality configuration from the message.
  /// @param tokenArgs Additional token arguments from the message.
  /// @return requiredCCVs The list of CCV addresses the pool requires with defaults expanded if requested.
  function _getCCVsForPool(
    uint64 destChainSelector,
    address token,
    uint256 amount,
    uint16 finality,
    bytes memory tokenArgs
  ) internal view returns (address[] memory requiredCCVs) {
    address[] memory defaultCCVs = s_destChainConfigs[destChainSelector].defaultCCVs;
    IPoolV1 pool = getPoolBySourceToken(destChainSelector, IERC20(token));

    // Pool not specifying CCVs or lacking V2 support falls back to destination defaults so the lane still enforces a
    // minimum verifier set.
    if (!IERC165(pool).supportsInterface(type(IPoolV2).interfaceId)) {
      return defaultCCVs;
    }

    requiredCCVs = IPoolV2(address(pool)).getRequiredCCVs(
      token, destChainSelector, amount, finality, tokenArgs, IPoolV2.MessageDirection.Outbound
    );

    if (requiredCCVs.length == 0) {
      return defaultCCVs;
    }

    address[] memory resolvedCCVs = new address[](requiredCCVs.length + defaultCCVs.length);
    uint256 writeIndex = 0;
    bool includeDefaults = false;

    for (uint256 i = 0; i < requiredCCVs.length; ++i) {
      address poolCCV = requiredCCVs[i];
      if (poolCCV == address(0)) {
        includeDefaults = true;
        continue;
      }
      resolvedCCVs[writeIndex++] = poolCCV;
    }

    if (includeDefaults) {
      for (uint256 i = 0; i < defaultCCVs.length; ++i) {
        resolvedCCVs[writeIndex++] = defaultCCVs[i];
      }
    }

    assembly {
      mstore(resolvedCCVs, writeIndex)
    }

    return resolvedCCVs;
  }

  // ================================================================
  // │                             Fees                             │
  // ================================================================

  /// @inheritdoc IEVM2AnyOnRampClient
  /// @dev getFee MUST revert if the feeToken is not listed in the fee token config, as the router assumes it does.
  /// @param destChainSelector The destination chain selector.
  /// @param message The message to price.
  /// @return feeTokenAmount The amount of fee token needed for the fee, in smallest denomination of the fee token.
  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message
  ) external view returns (uint256 feeTokenAmount) {
    DestChainConfig storage destChainConfig = s_destChainConfigs[destChainSelector];
    if (address(destChainConfig.router) == address(0)) {
      revert DestinationChainNotSupported(destChainSelector);
    }

    ExtraArgsCodec.GenericExtraArgsV3 memory resolvedExtraArgs = _parseExtraArgsWithDefaults(
      destChainSelector,
      destChainConfig,
      message.extraArgs,
      (message.data.length == 0 && message.tokenAmounts.length > 0)
    );
    // Update the CCVs list to include lane mandated and pool required CCVs.
    address[] memory poolRequiredCCVs = new address[](0);
    if (message.tokenAmounts.length != 0) {
      poolRequiredCCVs = _getCCVsForPool(
        destChainSelector,
        message.tokenAmounts[0].token,
        message.tokenAmounts[0].amount,
        resolvedExtraArgs.blockConfirmations,
        resolvedExtraArgs.tokenArgs
      );
    }
    (resolvedExtraArgs.ccvs, resolvedExtraArgs.ccvArgs) = _mergeCCVLists(
      resolvedExtraArgs.ccvs, resolvedExtraArgs.ccvArgs, destChainConfig.laneMandatedCCVs, poolRequiredCCVs
    );

    // We sum the fees for the verifier, executor and the pool (if any).
    (,, feeTokenAmount) =
      _getReceipts(destChainSelector, destChainConfig.networkFeeUSDCents, message, resolvedExtraArgs);

    return feeTokenAmount;
  }

  /// @notice Gets the receipts for a message. The ordering of receipts is as follows:
  /// - Verifier receipts in the order of the CCV list.
  /// - Token transfer receipt if any tokens are being transferred.
  /// - Executor receipt.
  /// - Network fee receipt.
  /// @param destChainSelector The destination chain selector.
  /// @param networkFeeUSDCents The flat network fee charged in USD cents.
  /// @param message The message being sent.
  /// @param extraArgs The extra arguments for the message.
  /// @return receipts The list of receipts for verifiers, token transfer and executor.
  /// @return gasLimitSum The total gas limit required to execute the transaction on the destination chain.
  /// @return feeTokenAmount The total fee amount in fee token smallest denomination.
  function _getReceipts(
    uint64 destChainSelector,
    uint16 networkFeeUSDCents,
    Client.EVM2AnyMessage calldata message,
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs
  ) internal view returns (Receipt[] memory receipts, uint32 gasLimitSum, uint256 feeTokenAmount) {
    // Preallocate receipts with room for: verifiers, token transfer (0 or 1), executor, and network fee.
    receipts = new Receipt[](extraArgs.ccvs.length + message.tokenAmounts.length + 2);
    uint32 bytesOverheadSum = 0;

    for (uint256 i = 0; i < extraArgs.ccvs.length; ++i) {
      address implAddress = ICrossChainVerifierResolver(extraArgs.ccvs[i]).getOutboundImplementation(
        destChainSelector, extraArgs.ccvArgs[i]
      );
      if (implAddress == address(0)) {
        revert DestinationChainNotSupportedByCCV(extraArgs.ccvs[i], destChainSelector);
      }

      (uint256 feeUSDCents, uint32 gasForVerification, uint32 ccvPayloadSizeBytes) = ICrossChainVerifierV1(implAddress)
        .getFee(destChainSelector, message, extraArgs.ccvArgs[i], extraArgs.blockConfirmations);

      receipts[i] = Receipt({
        issuer: extraArgs.ccvs[i],
        destGasLimit: gasForVerification,
        destBytesOverhead: ccvPayloadSizeBytes,
        feeTokenAmount: feeUSDCents,
        extraArgs: extraArgs.ccvArgs[i]
      });

      gasLimitSum += gasForVerification;
      bytesOverheadSum += ccvPayloadSizeBytes;
    }

    if (message.tokenAmounts.length > 0) {
      IPoolV1 pool = getPoolBySourceToken(destChainSelector, IERC20(message.tokenAmounts[0].token));
      bool hasCustomFeeConfig = false;

      // Since the ordering is known, we can directly calculate the index for the pool receipt.
      uint256 poolReceiptIndex = extraArgs.ccvs.length;

      // issuer is set to the token pool address.
      receipts[poolReceiptIndex] = Receipt({
        issuer: address(pool),
        destGasLimit: 0,
        destBytesOverhead: 0,
        feeTokenAmount: 0,
        extraArgs: extraArgs.tokenArgs
      });

      // Try to call `IPoolV2.getFee` to fetch fee components if the pool supports IPoolV2. If the specified pool is not
      // a contract, we want it to revert. Using `ERC165Checker` here would mean it doesn't revert, and it costs more
      // gas.
      if (pool.supportsInterface(type(IPoolV2).interfaceId)) {
        (
          receipts[poolReceiptIndex].feeTokenAmount,
          receipts[poolReceiptIndex].destGasLimit,
          receipts[poolReceiptIndex].destBytesOverhead,
          ,
          hasCustomFeeConfig
        ) = IPoolV2(address(pool)).getFee(
          message.tokenAmounts[0].token,
          destChainSelector,
          message.tokenAmounts[0].amount,
          message.feeToken,
          extraArgs.blockConfirmations,
          extraArgs.tokenArgs
        );
      }

      // If the pool doesn't support IPoolV2 or config is disabled, fall back to FeeQuoter.
      if (!hasCustomFeeConfig) {
        (
          receipts[poolReceiptIndex].feeTokenAmount,
          receipts[poolReceiptIndex].destGasLimit,
          receipts[poolReceiptIndex].destBytesOverhead
        ) = IFeeQuoter(s_dynamicConfig.feeQuoter).getTokenTransferFee(destChainSelector, message.tokenAmounts[0].token);
      }

      gasLimitSum += receipts[poolReceiptIndex].destGasLimit;
      bytesOverheadSum += receipts[poolReceiptIndex].destBytesOverhead;
    }

    uint256 executorIndex = receipts.length - 2;
    // This includes the user callback gas limit.
    receipts[executorIndex] =
      _getExecutionFee(destChainSelector, message.data.length, message.tokenAmounts.length, extraArgs);

    gasLimitSum += receipts[executorIndex].destGasLimit;
    bytesOverheadSum += receipts[executorIndex].destBytesOverhead;

    // Tag the calling router as the network fee issuer so CCIPMessageSent surfaces which router forwarded the
    // message. Third-party verifiers can pin on the router address even though the flat network fee stays on
    // the onRamp for later aggregation.
    receipts[receipts.length - 1] = Receipt({
      issuer: msg.sender,
      destGasLimit: 0,
      destBytesOverhead: 0,
      feeTokenAmount: networkFeeUSDCents,
      extraArgs: ""
    });

    (uint32 updatedGasLimitSum, uint256 execCostInUSDCents, uint256 feeTokenPrice, uint256 percentMultiplier) =
    IFeeQuoter(s_dynamicConfig.feeQuoter).quoteGasForExec(
      destChainSelector, gasLimitSum, bytesOverheadSum, message.feeToken
    );

    // Transform the USD based fees into fee token amounts & sum them. For the executor, if the executor isn't
    // NO_EXECUTION_ADDRESS we also add the execution cost.
    for (uint256 i = 0; i < receipts.length; ++i) {
      // Example:
      // - feeTokenPrice = $15 = 15e18
      // - usdFeeCents = $1.50 = 150
      // - feeTokenAmount = 150 * 1e34 / 15e18 = 1e17 (0.1 tokens of the fee token)
      // Normally we'd multiple by 1e36, but since usdFeeCents has 2 decimals and bpsMultiplier has 2 decimals, we use
      // 1e32 here.
      receipts[i].feeTokenAmount *= percentMultiplier * 1e32 / feeTokenPrice;

      if (i == executorIndex) {
        // Update the fee of the executor to include execution costs.
        if (extraArgs.executor != Client.NO_EXECUTION_ADDRESS) {
          // Add execution cost to the executor's fee. Execution cost should not be multiplied by bpsMultiplier.
          receipts[i].feeTokenAmount += execCostInUSDCents * 1e34 / feeTokenPrice;
        }
      }

      feeTokenAmount += receipts[i].feeTokenAmount;
    }

    return (receipts, updatedGasLimitSum, feeTokenAmount);
  }

  /// @notice Gets the execution fee receipt. Takes into account specifying the NO_EXECUTION_ADDRESS.
  /// @param destChainSelector The destination chain selector.
  /// @param dataLength The length of the message data.
  /// @param numberOfTokens The number of tokens being transferred.
  /// @param extraArgs The extra arguments for the message.
  /// @return receipt The execution fee receipt.
  function _getExecutionFee(
    uint64 destChainSelector,
    uint256 dataLength,
    uint256 numberOfTokens,
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs
  ) internal view returns (Receipt memory) {
    DestChainConfig storage destChainConfig = s_destChainConfigs[destChainSelector];
    uint8 remoteChainAddressLengthBytes = destChainConfig.addressBytesLength;

    // Even if no automated execution is requested, we still need to fill out the receipt for proper accounting.
    // The gas limit and byte overhead are still relevant for estimating total message cost.
    return Receipt({
      issuer: extraArgs.executor,
      destGasLimit: destChainConfig.baseExecutionGasCost + extraArgs.gasLimit,
      // Since the message payload is the same on source and destination chains with the V1 codec, we can use the
      // same calculation for execBytes on destination.
      destBytesOverhead: uint32(
        MessageV1Codec.MESSAGE_V1_EVM_SOURCE_BASE_SIZE + dataLength + extraArgs.executorArgs.length
          + (MessageV1Codec.MESSAGE_V1_REMOTE_CHAIN_ADDRESSES * remoteChainAddressLengthBytes)
          + (numberOfTokens * (MessageV1Codec.TOKEN_TRANSFER_V1_EVM_SOURCE_BASE_SIZE + remoteChainAddressLengthBytes))
      ),
      // Only bill a flat fee when automated execution is enabled.
      feeTokenAmount: extraArgs.executor == Client.NO_EXECUTION_ADDRESS
        ? 0
        : IExecutor(extraArgs.executor).getFee(
          destChainSelector, extraArgs.blockConfirmations, extraArgs.ccvs, extraArgs.executorArgs
        ),
      extraArgs: extraArgs.executorArgs
    });
  }

  /// @notice Withdraws the outstanding fee token balances to the fee aggregator.
  /// @param feeTokens The fee tokens to withdraw.
  /// @dev This function can be permissionless as it only transfers tokens to the fee aggregator which is a trusted address.
  function withdrawFeeTokens(
    address[] calldata feeTokens
  ) external {
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
