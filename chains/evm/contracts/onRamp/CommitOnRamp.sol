// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICCVOnRamp} from "../interfaces/ICCVOnRamp.sol";
import {IFeeQuoterV2} from "../interfaces/IFeeQuoterV2.sol";
import {INonceManager} from "../interfaces/INonceManager.sol";
import {IRMNRemote} from "../interfaces/IRMNRemote.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Client} from "../libraries/Client.sol";
import {Internal} from "../libraries/Internal.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice The CommitOnRamp is a contract that handles lane-specific fee logic.
/// @dev The OnRamp and OffRamp form a cross chain upgradeable unit. Any change to one of them results in an onchain
/// upgrade of both contracts.
contract CommitOnRamp is ICCVOnRamp, ITypeAndVersion, Ownable2StepMsgSender {
  using SafeERC20 for IERC20;
  using EnumerableSet for EnumerableSet.AddressSet;

  error CannotSendZeroTokens();
  error UnsupportedToken(address token);
  error MustBeCalledByVerifierAggregator();
  error RouterMustSetOriginalSender();
  error InvalidConfig();
  error CursedByRMN(uint64 destChainSelector);
  error InvalidDestChainConfig(uint64 destChainSelector);
  error OnlyCallableByOwnerOrAllowlistAdmin();
  error SenderNotAllowed(address sender);
  error InvalidAllowListRequest(uint64 destChainSelector);

  event ConfigSet(StaticConfig staticConfig, DynamicConfig dynamicConfig);
  event DestChainConfigSet(uint64 indexed destChainSelector, address router, bool allowlistEnabled);
  event FeeTokenWithdrawn(address indexed feeAggregator, address indexed feeToken, uint256 amount);
  event AllowListSendersAdded(uint64 indexed destChainSelector, address[] senders);
  event AllowListSendersRemoved(uint64 indexed destChainSelector, address[] senders);

  struct StaticConfig {
    IRMNRemote rmnRemote; // RMN remote address.
    address nonceManager; // Nonce manager address.
  }

  /// @dev Struct that contains the dynamic configuration
  // solhint-disable-next-line gas-struct-packing
  struct DynamicConfig {
    address feeQuoter; // FeeQuoter address.
    bool reentrancyGuardEntered; // Reentrancy protection.
    address feeAggregator; // Fee aggregator address.
    address allowlistAdmin; // authorized admin to add or remove allowed senders.
  }

  /// @dev Struct to hold the configs for a single destination chain.
  struct DestChainConfig {
    bool allowlistEnabled; // ──────╮ True if the allowlist is enabled.
    address verifierAggregator; // ─╯ Local router address  that is allowed to send messages to the destination chain.
    EnumerableSet.AddressSet allowedSendersList; // The list of addresses allowed to send messages.
  }

  /// @dev Same as DestChainConfig but with the destChainSelector so that an array of these can be passed in the
  /// constructor and the applyDestChainConfigUpdates function.
  // solhint-disable gas-struct-packing
  struct DestChainConfigArgs {
    uint64 destChainSelector; // ──╮ Destination chain selector.
    address verifierAggregator; // │ Source router address.
    bool allowlistEnabled; // ─────╯ True if the allowlist is enabled.
  }

  /// @dev Struct to hold the allowlist configuration args per dest chain.
  struct AllowlistConfigArgs {
    uint64 destChainSelector; // ─╮ Destination chain selector.
    bool allowlistEnabled; // ────╯ True if the allowlist is enabled.
    address[] addedAllowlistedSenders; // list of senders to be added to the allowedSendersList.
    address[] removedAllowlistedSenders; // list of senders to be removed from the allowedSendersList.
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "CommitVerifier 1.7.0-dev";
  /// @dev The rmn contract.
  IRMNRemote private immutable i_rmnRemote;
  /// @dev The address of the nonce manager.
  address private immutable i_nonceManager;

  // DYNAMIC CONFIG
  /// @dev The dynamic config for the onRamp.
  DynamicConfig private s_dynamicConfig;

  /// @dev The destination chain specific configs.
  mapping(uint64 destChainSelector => DestChainConfig destChainConfig) private s_destChainConfigs;

  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    DestChainConfigArgs[] memory destChainConfigArgs
  ) {
    if (address(staticConfig.rmnRemote) == address(0) || staticConfig.nonceManager == address(0)) {
      revert InvalidConfig();
    }

    i_rmnRemote = staticConfig.rmnRemote;
    i_nonceManager = staticConfig.nonceManager;

    _setDynamicConfig(dynamicConfig);
    _applyDestChainConfigUpdates(destChainConfigArgs);
  }

  // ================================================================
  // │                          Messaging                           │
  // ================================================================

  function forwardToVerifier(bytes calldata rawMessage, uint256 verifierIndex) external returns (bytes memory) {
    Internal.EVM2AnyVerifierMessage memory message = abi.decode(rawMessage, (Internal.EVM2AnyVerifierMessage));

    _assertNotCursed(message.header.destChainSelector);

    // If the allowlist is enabled, check if the original sender is allowed.
    DestChainConfig storage destChainConfig = s_destChainConfigs[message.header.destChainSelector];
    // VerifierAggregator address may be zero intentionally to pause, which should stop all messages.
    if (msg.sender != destChainConfig.verifierAggregator) revert MustBeCalledByVerifierAggregator();

    if (destChainConfig.allowlistEnabled) {
      if (!destChainConfig.allowedSendersList.contains(message.sender)) {
        revert SenderNotAllowed(message.sender);
      }
    }

    (, bool isOutOfOrderExecution,,) = IFeeQuoterV2(s_dynamicConfig.feeQuoter).processMessageArgs(
      message.header.destChainSelector,
      message.feeToken,
      message.feeTokenAmount,
      message.verifierReceipts[verifierIndex].extraArgs,
      message.receiver
    );

    uint64 nonce = 0;
    if (!isOutOfOrderExecution) {
      // If the message is not out of order execution, we need to increment the nonce.
      nonce =
        INonceManager(i_nonceManager).getIncrementedOutboundNonce(message.header.destChainSelector, message.sender);
    }

    return abi.encode(nonce);
  }

  function _assertNotCursed(
    uint64 destChainSelector
  ) internal view {
    if (i_rmnRemote.isCursed(bytes16(uint128(destChainSelector)))) revert CursedByRMN(destChainSelector);
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Returns the static onRamp config.
  /// @dev RMN depends on this function, if modified, please notify the RMN maintainers.
  /// @return staticConfig the static configuration.
  function getStaticConfig() external view returns (StaticConfig memory) {
    return StaticConfig({rmnRemote: i_rmnRemote, nonceManager: i_nonceManager});
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

    emit ConfigSet(StaticConfig({rmnRemote: i_rmnRemote, nonceManager: i_nonceManager}), dynamicConfig);
  }

  /// @notice Updates destination chains specific configs.
  /// @param destChainConfigArgs Array of destination chain specific configs.
  function applyDestChainConfigUpdates(
    DestChainConfigArgs[] memory destChainConfigArgs
  ) external onlyOwner {
    _applyDestChainConfigUpdates(destChainConfigArgs);
  }

  /// @notice Internal version of applyDestChainConfigUpdates.
  function _applyDestChainConfigUpdates(
    DestChainConfigArgs[] memory destChainConfigArgs
  ) internal {
    for (uint256 i = 0; i < destChainConfigArgs.length; ++i) {
      DestChainConfigArgs memory destChainConfigArg = destChainConfigArgs[i];
      uint64 destChainSelector = destChainConfigArgs[i].destChainSelector;

      if (destChainSelector == 0) {
        revert InvalidDestChainConfig(destChainSelector);
      }

      DestChainConfig storage destChainConfig = s_destChainConfigs[destChainSelector];
      // The router can be zero to pause the destination chain
      destChainConfig.verifierAggregator = destChainConfigArg.verifierAggregator;
      destChainConfig.allowlistEnabled = destChainConfigArg.allowlistEnabled;

      emit DestChainConfigSet(
        destChainSelector, destChainConfigArg.verifierAggregator, destChainConfig.allowlistEnabled
      );
    }
  }

  /// @notice get ChainConfig configured for the DestinationChainSelector.
  /// @param destChainSelector The destination chain selector.
  /// @return allowlistEnabled boolean indicator to specify if allowlist check is enabled.
  /// @return verifierAggregator address of the verifierAggregator.
  function getDestChainConfig(
    uint64 destChainSelector
  ) external view returns (bool allowlistEnabled, address verifierAggregator) {
    DestChainConfig storage config = s_destChainConfigs[destChainSelector];
    allowlistEnabled = config.allowlistEnabled;
    verifierAggregator = config.verifierAggregator;
    return (allowlistEnabled, verifierAggregator);
  }

  /// @notice get allowedSenders List configured for the DestinationChainSelector.
  /// @param destChainSelector The destination chain selector.
  /// @return isEnabled True if allowlist is enabled.
  /// @return configuredAddresses This is always populated with the list of allowed senders, even if the allowlist
  /// is turned off. This is because the only way to know what addresses are configured is through this function. If
  /// it would return an empty list when the allowlist is disabled, it would be impossible to know what addresses are
  /// configured.
  function getAllowedSendersList(
    uint64 destChainSelector
  ) external view returns (bool isEnabled, address[] memory configuredAddresses) {
    return (
      s_destChainConfigs[destChainSelector].allowlistEnabled,
      s_destChainConfigs[destChainSelector].allowedSendersList.values()
    );
  }

  // ================================================================
  // │                          Allowlist                           │
  // ================================================================

  /// @notice Updates allowlistConfig for Senders.
  /// @dev configuration used to set the list of senders who are authorized to send messages.
  /// @param allowlistConfigArgsItems Array of AllowlistConfigArguments where each item is for a destChainSelector.
  function applyAllowlistUpdates(
    AllowlistConfigArgs[] calldata allowlistConfigArgsItems
  ) external {
    if (msg.sender != owner()) {
      if (msg.sender != s_dynamicConfig.allowlistAdmin) {
        revert OnlyCallableByOwnerOrAllowlistAdmin();
      }
    }

    for (uint256 i = 0; i < allowlistConfigArgsItems.length; ++i) {
      AllowlistConfigArgs memory allowlistConfigArgs = allowlistConfigArgsItems[i];

      DestChainConfig storage destChainConfig = s_destChainConfigs[allowlistConfigArgs.destChainSelector];
      destChainConfig.allowlistEnabled = allowlistConfigArgs.allowlistEnabled;

      if (allowlistConfigArgs.addedAllowlistedSenders.length > 0) {
        if (allowlistConfigArgs.allowlistEnabled) {
          for (uint256 j = 0; j < allowlistConfigArgs.addedAllowlistedSenders.length; ++j) {
            address toAdd = allowlistConfigArgs.addedAllowlistedSenders[j];
            if (toAdd == address(0)) {
              revert InvalidAllowListRequest(allowlistConfigArgs.destChainSelector);
            }
            destChainConfig.allowedSendersList.add(toAdd);
          }

          emit AllowListSendersAdded(allowlistConfigArgs.destChainSelector, allowlistConfigArgs.addedAllowlistedSenders);
        } else {
          revert InvalidAllowListRequest(allowlistConfigArgs.destChainSelector);
        }
      }

      for (uint256 j = 0; j < allowlistConfigArgs.removedAllowlistedSenders.length; ++j) {
        destChainConfig.allowedSendersList.remove(allowlistConfigArgs.removedAllowlistedSenders[j]);
      }

      if (allowlistConfigArgs.removedAllowlistedSenders.length > 0) {
        emit AllowListSendersRemoved(
          allowlistConfigArgs.destChainSelector, allowlistConfigArgs.removedAllowlistedSenders
        );
      }
    }
  }

  // ================================================================
  // │                             Fees                             │
  // ================================================================

  function getFee(uint64 destChainSelector, Client.EVM2AnyMessage memory, bytes memory) external view returns (uint256) {
    _assertNotCursed(destChainSelector);

    return 0;
  }

  /// @dev getFee MUST revert if the feeToken is not listed in the fee token config, as the router assumes it does.
  /// @param destChainSelector The destination chain selector.
  /// @param message The message to get quote for.
  /// @return feeTokenAmount The amount of fee token needed for the fee, in smallest denomination of the fee token.
  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message
  ) external view returns (uint256 feeTokenAmount) {
    _assertNotCursed(destChainSelector);

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
