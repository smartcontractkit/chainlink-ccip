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
import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {Pool} from "../libraries/Pool.sol";
import {USDPriceWith18Decimals} from "../libraries/USDPriceWith18Decimals.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.0.2/utils/structs/EnumerableSet.sol";

// TODO post process hooks?
contract OnRamp is IEVM2AnyOnRampClient, ITypeAndVersion, Ownable2StepMsgSender {
  using SafeERC20 for IERC20;
  using EnumerableSet for EnumerableSet.AddressSet;
  using USDPriceWith18Decimals for uint224;

  error CannotSendZeroTokens();
  error UnsupportedToken(address token);
  error CanOnlySendOneTokenPerMessage();
  error MustBeCalledByRouter();
  error RouterMustSetOriginalSender();
  error InvalidConfig();
  error CursedByRMN(uint64 destChainSelector);
  error GetSupportedTokensFunctionalityRemovedCheckAdminRegistry();
  error InvalidDestChainConfig(uint64 destChainSelector);
  error ReentrancyGuardReentrantCall();
  error InvalidOptionalCCVThreshold();
  error DestinationChainNotSupported(uint64 destChainSelector);

  event ConfigSet(StaticConfig staticConfig, DynamicConfig dynamicConfig);
  event DestChainConfigSet(
    uint64 indexed destChainSelector,
    uint64 sequenceNumber,
    IRouter router,
    address[] defaultCCVs,
    address[] laneMandatedCCVs,
    address defaultExecutor,
    bytes offRamp
  );
  event FeeTokenWithdrawn(address indexed feeAggregator, address indexed feeToken, uint256 amount);
  event CCIPMessageSent(
    uint64 indexed destChainSelector,
    uint64 indexed sequenceNumber,
    bytes32 indexed messageId,
    bytes encodedMessage,
    Receipt[] receipts,
    bytes[] verifierBlobs
  );

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
    IRouter router; // ─────────╮ Local router address  that is allowed to send messages to the destination chain.
    // The last used sequence number. This is zero in the case where no messages have yet been sent.
    // 0 is not a valid sequence number for any real transaction as this value will be incremented before use.
    uint64 sequenceNumber; // ──╯
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
    IRouter router; // Source router address  that is allowed to send messages to the destination chain.
    address[] defaultCCVs; // Default CCVs to use for messages to this destination chain.
    address[] laneMandatedCCVs; // Required CCVs to use for all messages to this destination chain.
    address defaultExecutor;
    bytes offRamp; // Destination OffRamp address, NOT abi encoded but raw bytes.
  }

  /// @notice Receipt structure used to record gas limits and fees for verifiers, executors and token transfers.
  struct Receipt {
    address issuer; // ───────────╮ The address of the entity that issued the receipt.
    uint64 destGasLimit; //       │ The gas limit for the actions taken on the destination chain for this entity.
    uint32 destBytesOverhead; // ─╯ The byte overhead for the actions taken on the destination chain for this entity.
    uint256 feeTokenAmount; // The fee amount in the fee token for this entity.
    bytes extraArgs; // Extra args that have been passed in on the source chain.
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

  /// @notice Gets the next sequence number to be used in the onRamp.
  /// @param destChainSelector The destination chain selector.
  /// @return nextSequenceNumber The next sequence number to be used.
  function getExpectedNextSequenceNumber(
    uint64 destChainSelector
  ) external view returns (uint64) {
    return s_destChainConfigs[destChainSelector].sequenceNumber + 1;
  }

  /// @inheritdoc IEVM2AnyOnRampClient
  function forwardFromRouter(
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message,
    uint256 feeTokenAmount,
    address originalSender
  ) external returns (bytes32 messageId) {
    // We rely on a reentrancy guard here due to the untrusted calls performed to the pools. This enables some
    // optimizations by not following the CEI pattern.
    if (s_dynamicConfig.reentrancyGuardEntered) revert ReentrancyGuardReentrantCall();
    s_dynamicConfig.reentrancyGuardEntered = true;

    DestChainConfig storage destChainConfig = s_destChainConfigs[destChainSelector];

    // NOTE: assumes the message has already been validated through the getFee call.
    // Validate originalSender is set and allowed. Not validated in `getFee` since it is not user-driven.
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    // Router address may be zero intentionally to pause, which should stop all messages.
    if (msg.sender != address(destChainConfig.router)) revert MustBeCalledByRouter();

    // 1. parse extraArgs.

    Client.EVMExtraArgsV3 memory resolvedExtraArgs = _parseExtraArgsWithDefaults(destChainConfig, message.extraArgs);
    MessageV1Codec.MessageV1 memory newMessage = MessageV1Codec.MessageV1({
      sourceChainSelector: i_localChainSelector,
      destChainSelector: destChainSelector,
      sequenceNumber: ++destChainConfig.sequenceNumber,
      onRampAddress: abi.encodePacked(address(this)),
      offRampAddress: destChainConfig.offRamp,
      finality: resolvedExtraArgs.finalityConfig,
      sender: abi.encodePacked(originalSender),
      // The user encodes the receiver with abi.encode when creating EVM2AnyMessage
      // whereas MessageV1 expects just the raw bytes, so we strip the first 12 bytes.
      // TODO handle non-EVM chain families, maybe through fee quoter
      receiver: message.receiver[12:],
      // Executor args hold security critical execution args, like Solana accounts or Sui object IDs. Because of this,
      // they have to part of the message that is signed off on by the verifiers.
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
          resolvedExtraArgs.finalityConfig,
          resolvedExtraArgs.tokenArgs
        );
      }
      resolvedExtraArgs.ccvs =
        _mergeCCVLists(resolvedExtraArgs.ccvs, destChainConfig.laneMandatedCCVs, poolRequiredCCVs);
    }

    // 3. getFee on all verifiers, pool and executor.

    CCIPMessageSentEventData memory eventData;
    // Populate receipts for verifiers, pool and executor in that order.
    eventData.receipts = _getReceipts(destChainSelector, message, resolvedExtraArgs);

    // 4. lockOrBurn

    if (message.tokenAmounts.length != 0) {
      if (message.tokenAmounts.length != 1) revert CanOnlySendOneTokenPerMessage();
      // TODO where does the TokenReceiver go? Exec args feels strange but don't have a better place.
      bytes memory tokenReceiver =
        IFeeQuoter(s_dynamicConfig.feeQuoter).resolveTokenReceiver(resolvedExtraArgs.executorArgs);
      if (tokenReceiver.length == 0) {
        tokenReceiver = abi.encode(message.receiver);
      }
      newMessage.tokenTransfer[0] = _lockOrBurnSingleToken(
        message.tokenAmounts[0], destChainSelector, tokenReceiver, originalSender, resolvedExtraArgs.tokenArgs
      );
    }

    // 5. encode message and calculate messageId.

    eventData.encodedMessage = MessageV1Codec._encodeMessageV1(newMessage);
    messageId = keccak256(eventData.encodedMessage);

    eventData.verifierBlobs = new bytes[](resolvedExtraArgs.ccvs.length);

    // 6. resolve and call each verifier.
    {
      for (uint256 i = 0; i < resolvedExtraArgs.ccvs.length; ++i) {
        address implAddress =
          ICrossChainVerifierResolver(resolvedExtraArgs.ccvs[i].ccvAddress).getOutboundImplementation(destChainSelector);
        eventData.verifierBlobs[i] = ICrossChainVerifierV1(implAddress).forwardToVerifier(
          newMessage, messageId, message.feeToken, feeTokenAmount, resolvedExtraArgs.ccvs[i].args
        );
      }
    }

    // 7. emit event
    emit CCIPMessageSent(
      destChainSelector,
      newMessage.sequenceNumber,
      messageId,
      eventData.encodedMessage,
      eventData.receipts,
      eventData.verifierBlobs
    );

    s_dynamicConfig.reentrancyGuardEntered = false;

    return messageId;
  }

  /// @notice Merges lane mandated and pool required CCVs with user-provided CCVs.
  /// @dev This function assumes there are no duplicates in the userRequestedOrDefaultCCVs list.
  /// @dev There is no protocol-level requirement on the ordering of CCVs in the final list, but for determinism we
  /// process user requested first, then lane-mandated second, pool-required last.
  /// @param userRequestedOrDefaultCCVs User-provided required CCVs. Can not be empty, as defaults are applied earlier
  /// if needed. This list does not only contain addresses, but also arguments
  /// @param laneMandatedCCVs Lane mandated CCVs are always added, regardless of what a user/pool chooses. Can be empty.
  /// @param poolRequiredCCVs Pool-specific required CCVs.
  /// @return ccvs Updated list of CCVs.
  function _mergeCCVLists(
    Client.CCV[] memory userRequestedOrDefaultCCVs,
    address[] memory laneMandatedCCVs,
    address[] memory poolRequiredCCVs
  ) internal pure returns (Client.CCV[] memory ccvs) {
    // Maximum possible CCVs: user + lane + pool.
    uint256 totalCCVs = userRequestedOrDefaultCCVs.length + laneMandatedCCVs.length + poolRequiredCCVs.length;
    ccvs = new Client.CCV[](totalCCVs);
    uint256 toBeAddedIndex = 0;

    // First add all user requested CCVs.
    for (uint256 i = 0; i < userRequestedOrDefaultCCVs.length; ++i) {
      ccvs[toBeAddedIndex++] = userRequestedOrDefaultCCVs[i];
    }
    // Add lane mandated CCVs, skipping duplicates.
    for (uint256 i = 0; i < laneMandatedCCVs.length; ++i) {
      address laneMandatedCCV = laneMandatedCCVs[i];
      bool found = false;
      for (uint256 j = 0; j < toBeAddedIndex; ++j) {
        if (ccvs[j].ccvAddress == laneMandatedCCV) {
          found = true;
          break;
        }
      }
      if (!found) {
        ccvs[toBeAddedIndex++] = Client.CCV({ccvAddress: laneMandatedCCV, args: ""});
      }
    }
    // Add pool required CCVs, skipping duplicates.
    for (uint256 i = 0; i < poolRequiredCCVs.length; ++i) {
      address poolRequiredCCV = poolRequiredCCVs[i];
      bool found = false;
      for (uint256 j = 0; j < toBeAddedIndex; ++j) {
        if (ccvs[j].ccvAddress == poolRequiredCCV) {
          found = true;
          break;
        }
      }
      if (!found) {
        ccvs[toBeAddedIndex++] = Client.CCV({ccvAddress: poolRequiredCCV, args: ""});
      }
    }

    // Resize the array to the actual number of CCVs added.
    assembly {
      mstore(ccvs, toBeAddedIndex)
    }

    return ccvs;
  }

  /// @notice Parses and validates extra arguments, applying defaults from destination chain configuration.
  /// The function ensures all messages have the required CCVs and executor needed for processing,
  /// even when users don't explicitly specify them.
  /// @param destChainConfig Configuration for the destination chain including default values.
  /// @param extraArgs User-provided extra arguments in either V3 or legacy format.
  /// @return resolvedArgs Complete EVMExtraArgsV3 struct with all defaults applied.
  function _parseExtraArgsWithDefaults(
    DestChainConfig memory destChainConfig,
    bytes calldata extraArgs
  ) internal pure returns (Client.EVMExtraArgsV3 memory resolvedArgs) {
    if (extraArgs.length >= 4 && bytes4(extraArgs[0:4]) == Client.GENERIC_EXTRA_ARGS_V3_TAG) {
      resolvedArgs = abi.decode(extraArgs[4:], (Client.EVMExtraArgsV3));

      // We need to ensure no duplicate CCVs are present in the ccv list.
      uint256 length = resolvedArgs.ccvs.length;
      for (uint256 i = 0; i < length; ++i) {
        for (uint256 j = i + 1; j < length; ++j) {
          if (resolvedArgs.ccvs[i].ccvAddress == resolvedArgs.ccvs[j].ccvAddress) {
            revert CCVConfigValidation.DuplicateCCVNotAllowed(resolvedArgs.ccvs[i].ccvAddress);
          }
        }
      }

      // When users don't specify any CCVs, default CCVs are chosen.
      if (resolvedArgs.ccvs.length == 0) {
        resolvedArgs.ccvs = new Client.CCV[](destChainConfig.defaultCCVs.length);
        for (uint256 i = 0; i < destChainConfig.defaultCCVs.length; ++i) {
          resolvedArgs.ccvs[i] = Client.CCV({ccvAddress: destChainConfig.defaultCCVs[i], args: ""});
        }
      }
    } else {
      // If old extraArgs are supplied, they are assumed to be for the default CCV and the default executor.
      // This means any default CCV/executor has to be able to process all prior extraArgs.
      resolvedArgs.executorArgs = extraArgs;
      resolvedArgs.ccvs = new Client.CCV[](destChainConfig.defaultCCVs.length);
      for (uint256 i = 0; i < destChainConfig.defaultCCVs.length; ++i) {
        resolvedArgs.ccvs[i] = Client.CCV({ccvAddress: destChainConfig.defaultCCVs[i], args: ""});
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
  /// @dev RMN depends on this function, if modified, please notify the RMN maintainers.
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

      if (destChainSelector == 0 || destChainSelector == i_localChainSelector) {
        revert InvalidDestChainConfig(destChainSelector);
      }

      // Ensure at least one default or mandated CCV exists, and check for duplicates or zero addresses in both sets.
      CCVConfigValidation._validateDefaultAndMandatedCCVs(
        destChainConfigArg.defaultCCVs, destChainConfigArg.laneMandatedCCVs
      );

      DestChainConfig storage destChainConfig = s_destChainConfigs[destChainSelector];
      // The router can be zero to pause the destination chain.
      destChainConfig.router = destChainConfigArg.router;
      destChainConfig.defaultCCVs = destChainConfigArg.defaultCCVs;
      destChainConfig.laneMandatedCCVs = destChainConfigArg.laneMandatedCCVs;
      // Require a default executor so messages that rely on older/defaulted args still resolve to a concrete
      // executor. A zero executor would break backward compatibility and cause otherwise-valid traffic to revert.
      if (destChainConfigArg.defaultExecutor == address(0)) revert InvalidConfig();
      destChainConfig.defaultExecutor = destChainConfigArg.defaultExecutor;
      destChainConfig.offRamp = destChainConfigArg.offRamp;

      emit DestChainConfigSet(
        destChainSelector,
        destChainConfig.sequenceNumber,
        destChainConfigArg.router,
        destChainConfigArg.defaultCCVs,
        destChainConfigArg.laneMandatedCCVs,
        destChainConfigArg.defaultExecutor,
        destChainConfigArg.offRamp
      );
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
  function getPoolBySourceToken(uint64, /*destChainSelector*/ IERC20 sourceToken) public view returns (IPoolV1) {
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
  /// @param receiver Message receiver.
  /// @param originalSender Message sender.
  /// @return TokenTransferV1 token transfer encoding for MessageV1.
  function _lockOrBurnSingleToken(
    Client.EVMTokenAmount memory tokenAndAmount,
    uint64 destChainSelector,
    bytes memory receiver,
    address originalSender,
    bytes memory // extraArgs
  ) internal returns (MessageV1Codec.TokenTransferV1 memory) {
    if (tokenAndAmount.amount == 0) revert CannotSendZeroTokens();

    IPoolV1 sourcePool = getPoolBySourceToken(destChainSelector, IERC20(tokenAndAmount.token));
    // We don't have to check if it supports the pool version in a non-reverting way here because
    // if we revert here, there is no effect on CCIP. Therefore we directly call the supportsInterface
    // function and not through the ERC165Checker.
    if (address(sourcePool) == address(0) || !sourcePool.supportsInterface(Pool.CCIP_POOL_V1)) {
      revert UnsupportedToken(tokenAndAmount.token);
    }

    // TODO support CCIP_POOL_V2

    Pool.LockOrBurnOutV1 memory poolReturnData = sourcePool.lockOrBurn(
      Pool.LockOrBurnInV1({
        receiver: receiver,
        remoteChainSelector: destChainSelector,
        originalSender: originalSender,
        amount: tokenAndAmount.amount,
        localToken: tokenAndAmount.token
      })
    );

    // NOTE: pool data validations are outsourced to the FeeQuoter to handle family-specific logic handling.
    return MessageV1Codec.TokenTransferV1({
      amount: tokenAndAmount.amount,
      sourcePoolAddress: abi.encodePacked(address(sourcePool)),
      sourceTokenAddress: abi.encodePacked(tokenAndAmount.token),
      destTokenAddress: IFeeQuoter(s_dynamicConfig.feeQuoter).validateEncodedAddressAndEncodePacked(
        destChainSelector, poolReturnData.destTokenAddress
      ),
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
      token, destChainSelector, amount, finality, tokenArgs, IPoolV2.CCVDirection.Outbound
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
  /// @return feeTokenAmount The amount of fee token needed for the fee, in smallest denomination of the fee token.
  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message
  ) external view returns (uint256 feeTokenAmount) {
    DestChainConfig storage destChainConfig = s_destChainConfigs[destChainSelector];
    if (address(destChainConfig.router) == address(0)) {
      revert DestinationChainNotSupported(destChainSelector);
    }

    Client.EVMExtraArgsV3 memory resolvedExtraArgs = _parseExtraArgsWithDefaults(destChainConfig, message.extraArgs);
    // Update the CCVs list to include lane mandated and pool required CCVs.
    address[] memory poolRequiredCCVs = new address[](0);
    if (message.tokenAmounts.length != 0) {
      poolRequiredCCVs = _getCCVsForPool(
        destChainSelector,
        message.tokenAmounts[0].token,
        message.tokenAmounts[0].amount,
        resolvedExtraArgs.finalityConfig,
        resolvedExtraArgs.tokenArgs
      );
    }
    resolvedExtraArgs.ccvs = _mergeCCVLists(resolvedExtraArgs.ccvs, destChainConfig.laneMandatedCCVs, poolRequiredCCVs);

    // We sum the fees for the verifier, executor and the pool (if any).
    Receipt[] memory receipts = _getReceipts(destChainSelector, message, resolvedExtraArgs);

    uint256 destGasLimit = 0;
    uint256 destBytesOverhead = 0;

    // Sum all receipts.
    for (uint256 i = 0; i < receipts.length; ++i) {
      feeTokenAmount += receipts[i].feeTokenAmount;
      destGasLimit += receipts[i].destGasLimit;
      destBytesOverhead += receipts[i].destBytesOverhead;
    }

    // TODO: handle destBytes & gas

    // TODO: it currently returns dollars, not fee token amount.
    return feeTokenAmount;
  }

  function _getReceipts(
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message,
    Client.EVMExtraArgsV3 memory extraArgs
  ) internal view returns (Receipt[] memory verifierReceipts) {
    // Already ensure there's room for the token transfer and executor receipts.
    verifierReceipts = new Receipt[](extraArgs.ccvs.length + message.tokenAmounts.length + 1);

    for (uint256 i = 0; i < extraArgs.ccvs.length; ++i) {
      Client.CCV memory verifier = extraArgs.ccvs[i];
      address implAddress =
        ICrossChainVerifierResolver(verifier.ccvAddress).getOutboundImplementation(destChainSelector);

      (uint256 feeUSDCents, uint32 gasForVerification, uint32 payloadSizeBytes) =
        ICrossChainVerifierV1(implAddress).getFee(destChainSelector, message, verifier.args, extraArgs.finalityConfig);

      verifierReceipts[i] = Receipt({
        issuer: verifier.ccvAddress,
        destGasLimit: gasForVerification,
        destBytesOverhead: payloadSizeBytes,
        feeTokenAmount: feeUSDCents,
        extraArgs: verifier.args
      });
    }

    (uint16 usdCentsFee, uint64 execGasCost, uint32 execBytes) = IExecutor(extraArgs.executor).getFee(
      destChainSelector,
      extraArgs.finalityConfig,
      uint32(message.data.length),
      uint8(message.tokenAmounts.length),
      extraArgs.ccvs,
      extraArgs.executorArgs
    );
    verifierReceipts[verifierReceipts.length - 1] = Receipt({
      issuer: extraArgs.executor,
      destGasLimit: execGasCost, // TODO add user gas limit
      destBytesOverhead: execBytes,
      feeTokenAmount: usdCentsFee,
      extraArgs: extraArgs.executorArgs
    });

    if (message.tokenAmounts.length > 0) {
      // TODO pool fees
      verifierReceipts[verifierReceipts.length - 2] = Receipt({
        issuer: message.tokenAmounts[0].token,
        destGasLimit: 0,
        destBytesOverhead: 0,
        feeTokenAmount: 0,
        extraArgs: extraArgs.tokenArgs
      });
    }

    return verifierReceipts;
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
