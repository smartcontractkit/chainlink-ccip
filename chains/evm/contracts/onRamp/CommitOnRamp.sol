// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFeeQuoterV2} from "../interfaces/IFeeQuoterV2.sol";

import {Client} from "../libraries/Client.sol";
import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {BaseOnRamp} from "./BaseOnRamp.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @notice The CommitOnRamp is a contract that handles lane-specific fee logic.
/// @dev The OnRamp and OffRamp form a cross chain upgradeable unit. Any change to one of them results in an onchain
/// upgrade of both contracts.
contract CommitOnRamp is Ownable2StepMsgSender, BaseOnRamp {
  error InvalidConfig();
  error OnlyCallableByOwnerOrAllowlistAdmin();

  event ConfigSet(DynamicConfig dynamicConfig);

  /// @dev Struct that contains the dynamic configuration.
  // solhint-disable-next-line gas-struct-packing
  struct DynamicConfig {
    address feeQuoter; // FeeQuoter address.
    address feeAggregator; // Fee aggregator address.
    address allowlistAdmin; // authorized admin to add or remove allowed senders.
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "CommitOnRamp 1.7.0-dev";
  /// @dev The address of the nonce manager.
  address private immutable i_nonceManager;

  // DYNAMIC CONFIG
  /// @dev The dynamic config for the onRamp.
  DynamicConfig private s_dynamicConfig;

  constructor(
    DynamicConfig memory dynamicConfig
  ) {
    _setDynamicConfig(dynamicConfig);
  }

  /// @notice Forwards a message from CCV proxy to this verifier for processing and returns verifier-specific data.
  /// @dev This function is called by the CCV proxy to delegate message verification to this specific verifier.
  /// It performs critical validation to ensure message integrity and proper sequencing.
  /// @param message Decoded MessageV1 payload for verification.
  /// @param 2nd parameter is messageId, unused.
  /// @param 3rd parameter is feeToken, unused.
  /// @param 4th parameter is feeTokenAmount, unused.
  /// @param 5th parameter is verifierArgs (opaque verifier-specific args), unused.
  /// @return verifierReturnData Verifier-specific encoded data
  function forwardToVerifier(
    MessageV1Codec.MessageV1 calldata message,
    bytes32,
    address,
    uint256,
    bytes calldata
  ) external returns (bytes memory verifierReturnData) {
    // For EVM, sender is expected to be 20 bytes.
    address senderAddress = address(bytes20(message.sender));
    _assertSenderIsAllowed(message.destChainSelector, senderAddress);

    // TODO: Call IFeeQuoterV2.processMessage
    // IFeeQuoterV2(s_dynamicConfig.feeQuoter).processMessageArgs(
    //   message.destChainSelector, feeToken, feeTokenAmount, verifierArgs, message.receiver
    // );
    return "";
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

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
    if (dynamicConfig.feeQuoter == address(0) || dynamicConfig.feeAggregator == address(0)) revert InvalidConfig();

    s_dynamicConfig = dynamicConfig;

    emit ConfigSet(dynamicConfig);
  }

  /// @notice Updates destination chains specific configs.
  /// @param destChainConfigArgs Array of destination chain specific configs.
  function applyDestChainConfigUpdates(
    DestChainConfigArgs[] calldata destChainConfigArgs
  ) external onlyOwner {
    _applyDestChainConfigUpdates(destChainConfigArgs);
  }

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

    _applyAllowlistUpdates(allowlistConfigArgsItems);
  }

  // ================================================================
  // │                             Fees                             │
  // ================================================================

  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage memory message,
    bytes memory // extraArgs
  ) external view returns (uint256) {
    return IFeeQuoterV2(s_dynamicConfig.feeQuoter).getValidatedFee(destChainSelector, message);
  }

  /// @notice Withdraws the outstanding fee token balances to the fee aggregator.
  /// @param feeTokens The fee tokens to withdraw.
  /// @dev This function can be permissionless as it only transfers tokens to the fee aggregator which is a trusted address.
  function withdrawFeeTokens(
    address[] calldata feeTokens
  ) external {
    _withdrawFeeTokens(feeTokens, s_dynamicConfig.feeAggregator);
  }
}
