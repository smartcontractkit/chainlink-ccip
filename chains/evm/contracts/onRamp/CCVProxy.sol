// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICCVOnRamp} from "../interfaces/ICCVOnRamp.sol";
import {IEVM2AnyOnRampClient} from "../interfaces/IEVM2AnyOnRampClient.sol";
import {IFeeQuoterV2} from "../interfaces/IFeeQuoterV2.sol";
import {IPoolV1} from "../interfaces/IPool.sol";
import {IRMNRemote} from "../interfaces/IRMNRemote.sol";
import {IRouter} from "../interfaces/IRouter.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";

import {Client} from "../libraries/Client.sol";
import {Internal} from "../libraries/Internal.sol";
import {Pool} from "../libraries/Pool.sol";
import {USDPriceWith18Decimals} from "../libraries/USDPriceWith18Decimals.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

// TODO post process hooks?
contract CCVProxy is IEVM2AnyOnRampClient, ITypeAndVersion, Ownable2StepMsgSender {
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
  error DuplicateCCVInUserInput(address ccvAddress);
  error DuplicateCCVInConfig(address ccvAddress);

  event ConfigSet(StaticConfig staticConfig, DynamicConfig dynamicConfig);
  event DestChainConfigSet(uint64 indexed destChainSelector, uint64 sequenceNumber, IRouter router);
  event FeeTokenWithdrawn(address indexed feeAggregator, address indexed feeToken, uint256 amount);
  /// RMN depends on this event, if changing, please notify the RMN maintainers.
  event CCIPMessageSent(
    uint64 indexed destChainSelector,
    uint64 indexed sequenceNumber,
    Internal.EVM2AnyVerifierMessage message,
    bytes[] receiptBlobs
  );

  /// @dev Struct that contains the static configuration.
  /// RMN depends on this struct, if changing, please notify the RMN maintainers.
  // solhint-disable-next-line gas-struct-packing
  struct StaticConfig {
    uint64 chainSelector; // ────╮ Local chain selector.
    IRMNRemote rmnRemote; // ────╯ RMN remote address.
    address tokenAdminRegistry; // Token admin registry address.
  }

  /// @dev Struct that contains the dynamic configuration
  // solhint-disable-next-line gas-struct-packing
  struct DynamicConfig {
    address feeQuoter; // FeeQuoter address.
    bool reentrancyGuardEntered; // Reentrancy protection.
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
  }

  /// @dev Same as DestChainConfig but with the destChainSelector so that an array of these can be passed in the
  /// constructor and the applyDestChainConfigUpdates function.
  // solhint-disable gas-struct-packing
  struct DestChainConfigArgs {
    uint64 destChainSelector; // Destination chain selector.
    IRouter router; //           Source router address.
    address[] defaultCCVs; // Default CCVs to use for messages to this destination chain.
    address[] laneMandatedCCVs; // Required CCVs to use for all messages to this destination chain.
    address defaultExecutor;
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "CCVProxy 1.7.0-dev";
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
  ) external returns (bytes32) {
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

    // 1. parse extraArgs
    Client.EVMExtraArgsV3 memory resolvedExtraArgs = _parseExtraArgsWithDefaults(destChainConfig, message.extraArgs);
    // TODO where does the TokenReceiver go? Exec args feels strange but don't have a better place.
    bytes memory tokenReceiver =
      IFeeQuoterV2(s_dynamicConfig.feeQuoter).resolveTokenReceiver(resolvedExtraArgs.executorArgs);
    if (tokenReceiver.length == 0) {
      tokenReceiver = abi.encode(message.receiver);
    }

    // 2. get pool params, this potentially mutates CCV list

    // TODO pool call & fill receipt

    Internal.Receipt memory poolReceipt;

    (resolvedExtraArgs.requiredCCV, resolvedExtraArgs.optionalCCV, resolvedExtraArgs.optionalThreshold) =
    _mergeCCVsWithPoolAndLaneMandated(
      destChainConfig,
      new address[](0), // TODO pass in pool required CCVs
      resolvedExtraArgs.requiredCCV,
      resolvedExtraArgs.optionalCCV,
      resolvedExtraArgs.optionalThreshold
    );

    uint256 requiredCCVsCount = resolvedExtraArgs.requiredCCV.length;

    Internal.Receipt memory emptyReceipt;

    Internal.EVM2AnyVerifierMessage memory newMessage = Internal.EVM2AnyVerifierMessage({
      header: Internal.Header({
        // Should be generated after the message is complete.
        messageId: "",
        sourceChainSelector: i_localChainSelector,
        destChainSelector: destChainSelector,
        // We need the next available sequence number so we increment before we use the value.
        sequenceNumber: ++destChainConfig.sequenceNumber
      }),
      sender: originalSender,
      data: message.data,
      receiver: message.receiver,
      feeToken: message.feeToken,
      feeTokenAmount: feeTokenAmount,
      feeValueJuels: 0, // TODO
      tokenTransfer: new Internal.EVMTokenTransfer[](message.tokenAmounts.length),
      verifierReceipts: new Internal.Receipt[](requiredCCVsCount + resolvedExtraArgs.optionalCCV.length),
      executorReceipt: emptyReceipt
    });

    // 3. getFee on all verifiers & executor

    for (uint256 i = 0; i < requiredCCVsCount; ++i) {
      Client.CCV memory ccv = resolvedExtraArgs.requiredCCV[i];
      newMessage.verifierReceipts[i] = Internal.Receipt({
        issuer: ccv.ccvAddress,
        feeTokenAmount: 0, // TODO
        destGasLimit: 0, // TODO
        destBytesOverhead: 0, // TODO
        extraArgs: ccv.args
      });
    }

    for (uint256 i = 0; i < resolvedExtraArgs.optionalCCV.length; ++i) {
      Client.CCV memory verifier = resolvedExtraArgs.optionalCCV[i];
      newMessage.verifierReceipts[i + requiredCCVsCount] = Internal.Receipt({
        issuer: verifier.ccvAddress,
        feeTokenAmount: 0, // TODO
        destGasLimit: 0, // TODO
        destBytesOverhead: 0, // TODO
        extraArgs: verifier.args
      });
    }
    // TODO

    // 4. lockOrBurn

    if (message.tokenAmounts.length != 0) {
      if (message.tokenAmounts.length != 1) {
        revert CanOnlySendOneTokenPerMessage();
      }
      newMessage.tokenTransfer[0] = _lockOrBurnSingleToken(
        message.tokenAmounts[0], destChainSelector, tokenReceiver, originalSender, resolvedExtraArgs.tokenArgs
      );
      newMessage.tokenTransfer[0].receipt = poolReceipt;
    }

    // 5. calculate msg ID

    // Hash only after all fields have been set, but before it's sent to the verifiers.
    newMessage.header.messageId = Internal._hash(
      newMessage,
      // Metadata hash preimage to ensure global uniqueness, ensuring 2 identical messages sent to 2 different lanes
      // will have a distinct hash.
      keccak256(abi.encode(Internal.EVM_2_ANY_MESSAGE_HASH, i_localChainSelector, destChainSelector, address(this)))
    );

    // 6. call each verifier

    bytes memory encodedMessage = abi.encode(newMessage);
    bytes[] memory receiptBlobs = new bytes[](newMessage.verifierReceipts.length);

    for (uint256 i = 0; i < newMessage.verifierReceipts.length; ++i) {
      address verifier = newMessage.verifierReceipts[i].issuer;

      ICCVOnRamp(verifier).forwardToVerifier(encodedMessage, i);
    }

    // 7. emit event

    emit CCIPMessageSent(destChainSelector, newMessage.header.sequenceNumber, newMessage, receiptBlobs);

    s_dynamicConfig.reentrancyGuardEntered = false;

    return newMessage.header.messageId;
  }

  /// @notice Merges lane mandated and pool required CCVs with user-provided CCVs.
  /// This function ensures no duplicates are added and handles moving CCVs from optional to required.
  /// @param destChainConfig Destination chain configuration containing lane mandated CCVs.
  /// @param poolRequiredCCVs Pool-specific required CCVs.
  /// @param requiredCCV User-provided required CCVs.
  /// @param optionalCCV User-provided optional CCVs.
  /// @param optionalThreshold Threshold for optional CCVs.
  /// @return newRequiredCCVs Updated required CCVs list.
  /// @return newOptionalCCVs Updated optional CCVs list.
  /// @return newOptionalThreshold Updated optional threshold.
  function _mergeCCVsWithPoolAndLaneMandated(
    DestChainConfig storage destChainConfig,
    address[] memory poolRequiredCCVs,
    Client.CCV[] memory requiredCCV,
    Client.CCV[] memory optionalCCV,
    uint8 optionalThreshold
  )
    internal
    view
    returns (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold)
  {
    // Maximum possible CCVs to add
    uint256 totalMandatory = destChainConfig.laneMandatedCCVs.length + poolRequiredCCVs.length;
    Client.CCV[] memory toBeAdded = new Client.CCV[](totalMandatory);
    uint256 toBeAddedIndex = 0;

    // Process all mandatory CCVs in a single pass.
    // We iterate lane-mandated first, then pool-required for determinism only; there is no protocol-level
    // requirement on relative ordering. Duplicates across the two sources are removed below.
    for (uint256 i = 0; i < totalMandatory; ++i) {
      address mandatoryCCV = i < destChainConfig.laneMandatedCCVs.length
        ? destChainConfig.laneMandatedCCVs[i]
        : poolRequiredCCVs[i - destChainConfig.laneMandatedCCVs.length];

      // Skip CCVs we've already collected from a lane-mandated or pool-required
      // to avoid adding duplicates to requiredCCV.
      bool isDuplicateInToBeAdded = false;
      for (uint256 j = 0; j < toBeAddedIndex; ++j) {
        if (toBeAdded[j].ccvAddress == mandatoryCCV) {
          isDuplicateInToBeAdded = true;
          break;
        }
      }
      if (isDuplicateInToBeAdded) continue;

      // Check if already exists in user's required CCVs
      bool existsInUserRequired = false;
      for (uint256 reqCCVIndex = 0; reqCCVIndex < requiredCCV.length; ++reqCCVIndex) {
        if (mandatoryCCV == requiredCCV[reqCCVIndex].ccvAddress) {
          existsInUserRequired = true;
          break;
        }
      }

      // If not in user's required list, add it
      if (!existsInUserRequired) {
        toBeAdded[toBeAddedIndex++].ccvAddress = mandatoryCCV;

        // If the mandatory CCV is in the optional CCVs, remove it and adjust threshold
        for (uint256 optCCVIndex = 0; optCCVIndex < optionalCCV.length; ++optCCVIndex) {
          if (mandatoryCCV == optionalCCV[optCCVIndex].ccvAddress) {
            // Copy the args from optional CCV
            toBeAdded[toBeAddedIndex - 1].args = optionalCCV[optCCVIndex].args;

            // Remove from optional CCVs by swapping with last element
            optionalCCV[optCCVIndex] = optionalCCV[optionalCCV.length - 1];
            // Reduce array length
            assembly {
              mstore(optionalCCV, sub(mload(optionalCCV), 1))
            }

            // Decrement threshold to maintain security guarantee
            if (optionalThreshold > 0) {
              optionalThreshold--;
            }
            // Since each CCV address should be unique, we can break after finding the first match
            break;
          }
        }
      }
    }

    if (toBeAddedIndex > 0) {
      newRequiredCCVs = new Client.CCV[](requiredCCV.length + toBeAddedIndex);
      for (uint256 i = 0; i < toBeAddedIndex; ++i) {
        newRequiredCCVs[i] = toBeAdded[i];
      }
      for (uint256 i = 0; i < requiredCCV.length; ++i) {
        newRequiredCCVs[toBeAddedIndex + i] = requiredCCV[i];
      }

      return (newRequiredCCVs, optionalCCV, optionalThreshold);
    }

    return (requiredCCV, optionalCCV, optionalThreshold);
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

      if (resolvedArgs.optionalCCV.length != 0) {
        // Prevent impossible execution scenarios: if threshold >= array length, no combination of optional CCVs
        // could ever satisfy the requirement. Zero threshold defeats the purpose of having optional CCVs.
        if (resolvedArgs.optionalCCV.length <= resolvedArgs.optionalThreshold || resolvedArgs.optionalThreshold == 0) {
          revert InvalidOptionalCCVThreshold();
        }
      }

      // We need to ensure no duplicate CCVs are present across required and optional lists.
      uint256 requiredCCVLength = resolvedArgs.requiredCCV.length;
      uint256 optionalCCVLength = resolvedArgs.optionalCCV.length;
      uint256 totalInputCCV = requiredCCVLength + optionalCCVLength;
      for (uint256 i = 0; i < totalInputCCV; ++i) {
        address ccvAddressI = i < requiredCCVLength
          ? resolvedArgs.requiredCCV[i].ccvAddress
          : resolvedArgs.optionalCCV[i - requiredCCVLength].ccvAddress;

        for (uint256 j = i + 1; j < totalInputCCV; ++j) {
          address ccvAddressJ = j < requiredCCVLength
            ? resolvedArgs.requiredCCV[j].ccvAddress
            : resolvedArgs.optionalCCV[j - requiredCCVLength].ccvAddress;

          if (ccvAddressI == ccvAddressJ) {
            revert DuplicateCCVInUserInput(ccvAddressI);
          }
        }
      }

      // When users don't specify any required CCVs, default CCVs are chosen.
      if (resolvedArgs.requiredCCV.length == 0) {
        resolvedArgs.requiredCCV = new Client.CCV[](destChainConfig.defaultCCVs.length);
        for (uint256 i = 0; i < destChainConfig.defaultCCVs.length; ++i) {
          resolvedArgs.requiredCCV[i] = Client.CCV({ccvAddress: destChainConfig.defaultCCVs[i], args: ""});
        }
      }
    } else {
      // If old extraArgs are supplied, they are assumed to be for the default CCV and the default executor.
      // This means any default CCV/executor has to be able to process all prior extraArgs.
      resolvedArgs.executorArgs = extraArgs;
      resolvedArgs.requiredCCV = new Client.CCV[](destChainConfig.defaultCCVs.length);
      for (uint256 i = 0; i < destChainConfig.defaultCCVs.length; ++i) {
        resolvedArgs.requiredCCV[i] = Client.CCV({ccvAddress: destChainConfig.defaultCCVs[i], args: extraArgs});
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
      DestChainConfigArgs memory destChainConfigArg = destChainConfigArgs[i];
      uint64 destChainSelector = destChainConfigArg.destChainSelector;

      if (destChainSelector == 0) {
        revert InvalidDestChainConfig(destChainSelector);
      }

      // TODO : Validate no duplicate within and overlap in defaultCCVs
      // and laneMandatedCCVs, maybe use same dedup logic in _parseExtraArgsWithDefaults

      DestChainConfig storage destChainConfig = s_destChainConfigs[destChainSelector];
      // The router can be zero to pause the destination chain.
      destChainConfig.router = destChainConfigArg.router;
      destChainConfig.defaultCCVs = destChainConfigArg.defaultCCVs;
      destChainConfig.laneMandatedCCVs = destChainConfigArg.laneMandatedCCVs;
      destChainConfig.defaultExecutor = destChainConfigArg.defaultExecutor;

      emit DestChainConfigSet(destChainSelector, destChainConfig.sequenceNumber, destChainConfigArg.router);
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

  /// @notice Uses a pool to lock or burn a token.
  /// @param tokenAndAmount Token address and amount to lock or burn.
  /// @param destChainSelector Target destination chain selector of the message.
  /// @param receiver Message receiver.
  /// @param originalSender Message sender.
  /// @return EVM2AnyCommitVerifierTokenTransfer EVM2Any token and amount data.
  function _lockOrBurnSingleToken(
    Client.EVMTokenAmount memory tokenAndAmount,
    uint64 destChainSelector,
    bytes memory receiver,
    address originalSender,
    bytes memory // extraArgs
  ) internal returns (Internal.EVMTokenTransfer memory) {
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

    Internal.Receipt memory emptyReceipt;

    // NOTE: pool data validations are outsourced to the FeeQuoter to handle family-specific logic handling.
    return Internal.EVMTokenTransfer({
      sourceTokenAddress: tokenAndAmount.token,
      destTokenAddress: poolReturnData.destTokenAddress,
      extraData: poolReturnData.destPoolData,
      amount: tokenAndAmount.amount,
      receipt: emptyReceipt
    });
  }

  // ================================================================
  // │                             Fees                             │
  // ================================================================

  /// @inheritdoc IEVM2AnyOnRampClient
  /// @dev getFee MUST revert if the feeToken is not listed in the fee token config, as the router assumes it does.
  /// @param destChainSelector The destination chain selector.
  /// @param message The message to get quote for.
  /// @return feeTokenAmount The amount of fee token needed for the fee, in smallest denomination of the fee token.
  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message
  ) external view returns (uint256 feeTokenAmount) {
    if (i_rmnRemote.isCursed(bytes16(uint128(destChainSelector)))) revert CursedByRMN(destChainSelector);

    return IFeeQuoterV2(s_dynamicConfig.feeQuoter).getValidatedFee(destChainSelector, message);
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
