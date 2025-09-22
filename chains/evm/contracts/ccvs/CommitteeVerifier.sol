// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../interfaces/ICrossChainVerifierV1.sol";

import {Client} from "../libraries/Client.sol";
import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {BaseVerifier} from "./components/BaseVerifier.sol";
import {SignatureQuorumValidator} from "./components/SignatureQuorumValidator.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @notice The CommitteeVerifier is a contract that handles lane-specific fee logic and message verification.
/// @dev Source and destination responsibilities are combined to enable a single proxy address for a CCV on each chain.
contract CommitteeVerifier is Ownable2StepMsgSender, ICrossChainVerifierV1, SignatureQuorumValidator, BaseVerifier {
  error InvalidConfig();
  error InvalidCCVData();
  error OnlyCallableByOwnerOrAllowlistAdmin();

  event ConfigSet(DynamicConfig dynamicConfig);

  /// @dev Defines upgradeable configuration parameters.
  // solhint-disable-next-line gas-struct-packing
  struct DynamicConfig {
    address feeQuoter; // The contract used to quote fees on source.
    address feeAggregator; // Entity capable of withdrawing fees.
    address allowlistAdmin; // Entity capable adding or removing allowed senders.
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "CommitteeVerifier 1.7.0-dev";
  /// @dev The number of bytes allocated to encoding the signature length within the ccvData.
  uint256 internal constant SIGNATURE_LENGTH_BYTES = 2;

  // DYNAMIC CONFIG
  DynamicConfig private s_dynamicConfig;

  constructor(DynamicConfig memory dynamicConfig, string memory storageLocation) BaseVerifier(storageLocation) {
    _setDynamicConfig(dynamicConfig);
  }

  /// @inheritdoc ICrossChainVerifierV1
  function forwardToVerifier(
    address originalCaller,
    MessageV1Codec.MessageV1 calldata message,
    bytes32, // messageId
    address, // feeToken
    uint256, // feeTokenAmount
    bytes calldata // verifierArgs
  ) external view returns (bytes memory verifierReturnData) {
    // For EVM, sender is expected to be 20 bytes.
    address senderAddress = address(bytes20(message.sender));
    _assertSenderIsAllowed(message.destChainSelector, senderAddress, originalCaller);

    // TODO: Process msg & return verifier data
    return "";
  }

  /// @inheritdoc ICrossChainVerifierV1
  function verifyMessage(
    address, // originalCaller
    MessageV1Codec.MessageV1 calldata, // message
    bytes32 messageHash,
    bytes calldata ccvData
  ) external view {
    if (ccvData.length < SIGNATURE_LENGTH_BYTES) {
      revert InvalidCCVData();
    }

    uint256 signatureLength = uint16(bytes2(ccvData[:SIGNATURE_LENGTH_BYTES]));
    if (ccvData.length < SIGNATURE_LENGTH_BYTES + signatureLength) {
      revert InvalidCCVData();
    }

    // Even though the current version of this contract only expects signatures to be included in the ccvData, bounding
    // it to the given length allows potential forward compatibility with future formats that supply more data.
    _validateSignatures(messageHash, ccvData[SIGNATURE_LENGTH_BYTES:SIGNATURE_LENGTH_BYTES + signatureLength]);
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Returns the dynamic config.
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

  /// @notice Updates the storage location identifier.
  /// @param newLocation The new storage location identifier.
  function updateStorageLocation(
    string memory newLocation
  ) external onlyOwner {
    emit BaseVerifier.StorageLocationUpdated(s_storageLocation, newLocation);

    s_storageLocation = newLocation;
  }

  // ================================================================
  // │                             Fees                             │
  // ================================================================

  /// @inheritdoc ICrossChainVerifierV1
  function getFee(
    address, // originalCaller
    uint64, // destChainSelector
    Client.EVM2AnyMessage memory, // message
    bytes memory // extraArgs
  ) external pure returns (uint256) {
    // TODO: Process msg & return fee
    return 0;
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
