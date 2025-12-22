// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../interfaces/ICrossChainVerifierV1.sol";

import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {BaseVerifier} from "./components/BaseVerifier.sol";
import {SignatureQuorumValidator} from "./components/SignatureQuorumValidator.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @notice The CommitteeVerifier is a contract that handles lane-specific fee logic and message verification.
/// @dev Source and destination responsibilities are combined to enable a single proxy address for a CCV on each chain.
contract CommitteeVerifier is Ownable2StepMsgSender, ICrossChainVerifierV1, SignatureQuorumValidator, BaseVerifier {
  error InvalidConfig();
  error InvalidVerifierResults();
  error InvalidCCVVersion(bytes4 verifierVersion);
  error OnlyCallableByOwnerOrAllowlistAdmin();
  error MustBeProposedStorageLocationsAdmin();
  error OnlyCallableByStorageLocationsAdmin();

  event ConfigSet(DynamicConfig dynamicConfig);
  event StorageLocationsAdminTransferRequested(address indexed from, address indexed to);
  event StorageLocationsAdminTransferred(address indexed from, address indexed to);

  /// @dev Defines upgradeable configuration parameters.
  struct DynamicConfig {
    address feeAggregator; // The entity receiving the withdrawn fees.
    address allowlistAdmin; // Entity capable adding or removing allowed senders.
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "CommitteeVerifier 1.7.0-dev";
  /// @dev The preimage is bytes4(keccak256("CommitteeVerifier 1.7.0")).
  bytes4 internal constant VERSION_TAG_V1_7_0 = 0x49ff34ed;
  /// @dev The number of bytes allocated to encoding the verifier version.
  uint256 internal constant VERIFIER_VERSION_BYTES = 4;
  /// @dev The number of bytes allocated to encoding the signature length within the verifierResults.
  uint256 internal constant SIGNATURE_LENGTH_BYTES = 2;

  // DYNAMIC CONFIG
  DynamicConfig private s_dynamicConfig;

  /// @dev Account capable of changing the storage location, and transferring admin.
  address private s_storageLocationsAdmin;
  /// @dev Pending storage locations admin.
  address private s_pendingStorageLocationsAdmin;

  constructor(
    DynamicConfig memory dynamicConfig,
    string[] memory storageLocations,
    address rmn
  ) BaseVerifier(storageLocations, rmn) {
    _setDynamicConfig(dynamicConfig);
    s_storageLocationsAdmin = msg.sender;
  }

  /// @inheritdoc ICrossChainVerifierV1
  function forwardToVerifier(
    MessageV1Codec.MessageV1 calldata message,
    bytes32, // messageId
    address, // feeToken
    uint256, // feeTokenAmount
    bytes calldata // verifierArgs
  ) external view returns (bytes memory verifierReturnData) {
    _assertNotCursedByRMN(message.destChainSelector);

    // For EVM, sender is abi encoded.
    address senderAddress = abi.decode(message.sender, (address));
    _assertSenderIsAllowed(message.destChainSelector, senderAddress);

    return abi.encodePacked(VERSION_TAG_V1_7_0);
  }

  /// @inheritdoc ICrossChainVerifierV1
  function verifyMessage(
    MessageV1Codec.MessageV1 calldata message,
    bytes32 messageHash,
    bytes calldata verifierResults
  ) external view {
    _assertNotCursedByRMN(message.sourceChainSelector);
    if (verifierResults.length < VERIFIER_VERSION_BYTES + SIGNATURE_LENGTH_BYTES) {
      revert InvalidVerifierResults();
    }

    // Any verifierResults submitted to this verifier should have the expected version.
    bytes4 verifierVersion = bytes4(verifierResults[:VERIFIER_VERSION_BYTES]);
    if (verifierVersion != VERSION_TAG_V1_7_0) {
      revert InvalidCCVVersion(verifierVersion);
    }

    uint256 signatureLength =
      uint16(bytes2(verifierResults[VERIFIER_VERSION_BYTES:VERIFIER_VERSION_BYTES + SIGNATURE_LENGTH_BYTES]));
    if (verifierResults.length < VERIFIER_VERSION_BYTES + SIGNATURE_LENGTH_BYTES + signatureLength) {
      revert InvalidVerifierResults();
    }

    // Even though the current version of this contract only expects verifier version and signatures to be included in the verifierResults,
    // bounding it to the given length allows potential forward compatibility with future formats that supply more data.
    _validateSignatures(
      message.sourceChainSelector,
      // Verifiers sign a concatenation of the verifier version and the message hash.
      // The version is included so that a resolver can return the correct verifier implementation on destination.
      // The version must be signed, otherwise any version could be inserted post-signatures.
      keccak256(bytes.concat(verifierVersion, messageHash)),
      verifierResults[VERIFIER_VERSION_BYTES
          + SIGNATURE_LENGTH_BYTES:VERIFIER_VERSION_BYTES + SIGNATURE_LENGTH_BYTES + signatureLength
      ]
    );
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
    if (dynamicConfig.feeAggregator == address(0)) revert InvalidConfig();

    s_dynamicConfig = dynamicConfig;

    emit ConfigSet(dynamicConfig);
  }

  /// @notice Updates remote chains specific configs.
  /// @param remoteChainConfigArgs Array of remote chain specific configs.
  function applyRemoteChainConfigUpdates(
    RemoteChainConfigArgs[] calldata remoteChainConfigArgs
  ) external onlyOwner {
    _applyRemoteChainConfigUpdates(remoteChainConfigArgs);
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

  /// @notice Returns the account currently authorized to manage the storage location.
  /// @return storageLocationsAdmin The active storage locations admin.
  function getStorageLocationsAdmin() external view returns (address) {
    return s_storageLocationsAdmin;
  }

  /// @notice Returns the account that has been nominated as the next storage locations admin.
  /// @return pendingStorageLocationsAdmin The address that must call acceptStorageLocationsAdmin.
  function getPendingStorageLocationsAdmin() external view returns (address) {
    return s_pendingStorageLocationsAdmin;
  }

  /// @notice Initiates the two-step transfer for the storage locations admin role.
  /// @param to The address proposed to become the next admin.
  function transferStorageLocationsAdmin(
    address to
  ) external {
    address currentAdmin = s_storageLocationsAdmin;
    if (msg.sender != currentAdmin) revert OnlyCallableByStorageLocationsAdmin();
    if (to == msg.sender) revert CannotTransferToSelf();

    s_pendingStorageLocationsAdmin = to;

    emit StorageLocationsAdminTransferRequested(currentAdmin, to);
  }

  /// @notice Completes the transfer of the storage locations admin role.
  /// @dev Only the nominated pending admin can call this function.
  function acceptStorageLocationsAdmin() external {
    if (msg.sender != s_pendingStorageLocationsAdmin) revert MustBeProposedStorageLocationsAdmin();

    address oldAdmin = s_storageLocationsAdmin;
    s_storageLocationsAdmin = msg.sender;
    s_pendingStorageLocationsAdmin = address(0);

    emit StorageLocationsAdminTransferred(oldAdmin, msg.sender);
  }

  /// @notice Updates the storage locations identifiers.
  /// @dev Caller must be the storage locations admin.
  /// @param newLocations The new storage locations identifiers.
  function updateStorageLocations(
    string[] memory newLocations
  ) external {
    if (msg.sender != s_storageLocationsAdmin) revert OnlyCallableByStorageLocationsAdmin();

    _setStorageLocations(newLocations);
  }

  /// @notice Exposes the version tag.
  function versionTag() public pure returns (bytes4) {
    return VERSION_TAG_V1_7_0;
  }

  // ================================================================
  // │                             Fees                             │
  // ================================================================

  /// @notice Withdraws the outstanding fee token balances to the fee aggregator.
  /// @param feeTokens The fee tokens to withdraw.
  /// @dev This function can be permissionless as it only transfers tokens to the fee aggregator which is a trusted address.
  function withdrawFeeTokens(
    address[] calldata feeTokens
  ) external {
    _withdrawFeeTokens(feeTokens, s_dynamicConfig.feeAggregator);
  }
}
