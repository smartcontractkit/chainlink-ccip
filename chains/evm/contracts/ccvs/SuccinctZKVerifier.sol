// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../interfaces/ICrossChainVerifierV1.sol";
import {ISP1Helios} from "../interfaces/succinct/ISP1Helios.sol";

import {FeeTokenHandler} from "../libraries/FeeTokenHandler.sol";
import {Internal} from "../libraries/Internal.sol";
import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {BaseVerifier} from "./components/BaseVerifier.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {Bytes} from "@eth-optimism/contracts-bedrock/src/libraries/Bytes.sol";
import {RLPReader} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPReader.sol";
import {RLPWriter} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPWriter.sol";
import {MerkleTrie} from "@eth-optimism/contracts-bedrock/src/libraries/trie/MerkleTrie.sol";

/// @notice The SuccinctZKVerifier verifies a message by proving that its CCIPMessageSent receipt is included in a
/// source chain block anchored by the SP1Helios light client on this chain. It does not rely on any signer set.
/// @dev Only EVM source chains are supported, as the proof is built from EVM block headers and receipts.
/// @dev Source and destination responsibilities are combined to enable a single proxy address for a CCV on each chain.
contract SuccinctZKVerifier is Ownable2StepMsgSender, ICrossChainVerifierV1, BaseVerifier {
  using RLPReader for RLPReader.RLPItem;
  using RLPReader for bytes;

  error InvalidVerifierResults();
  error InvalidCCVVersion(bytes4 expected, bytes4 got);
  error InvalidSourceChainConfig(uint64 sourceChainSelector);
  error SourceChainNotSupported(uint64 sourceChainSelector);
  error BlockNotAnchored(uint64 sourceChainSelector, uint256 blockNumber);
  error HeaderChainTooLong(uint256 headerCount, uint256 maxHeaderChainLength);
  error InvalidHeaderHash(uint256 headerIndex, bytes32 expected, bytes32 got);
  error InvalidFieldLength(uint256 expected, uint256 got);
  error ReceiptNotSuccessful();
  error LogIndexOutOfRange(uint256 logIndex, uint256 logCount);
  error InvalidLogEmitter(address expected, bytes got);
  error InvalidTopicCount(uint256 topicCount);
  error InvalidLogTopic(bytes32 expected, bytes32 got);
  error InvalidMessageId(bytes32 expected, bytes32 got);

  event DynamicConfigSet(DynamicConfig dynamicConfig);
  event SourceChainConfigSet(
    uint64 indexed sourceChainSelector, address helios, address onRamp, uint16 maxHeaderChainLength
  );

  /// @dev Defines upgradeable configuration parameters.
  struct DynamicConfig {
    address feeAggregator; // The entity receiving the withdrawn fees.
  }

  struct SourceChainConfigArgs {
    ISP1Helios helios; // ───────────╮ The light client anchoring this source chain. Can be zero to pause the chain.
    uint64 sourceChainSelector; // ──│ Source chain selector.
    uint16 maxHeaderChainLength; // ─╯ Max number of headers in a proof, bounds gas and calldata.
    address onRamp; // The OnRamp on the source chain that emits CCIPMessageSent.
  }

  struct SourceChainConfig {
    ISP1Helios helios; // ───────────╮ The light client anchoring this source chain. Zero means paused.
    uint16 maxHeaderChainLength; // ─╯ Max number of headers in a proof, bounds gas and calldata.
    address onRamp; // The OnRamp on the source chain that emits CCIPMessageSent.
  }

  /// @dev The proof carried in the verifierResults after the version tag, ABI encoded.
  struct Witness {
    uint256 anchorBlockNumber; // Source block anchored by SP1Helios.
    bytes[] headers; // RLP headers from the anchored block down to the message block. Empty when they are the same.
    uint256 txIndex; // Index of the transaction in the message block.
    uint256 logIndex; // Index of the CCIPMessageSent log in the receipt.
    bytes[] proofNodes; // Receipts trie proof, root node first.
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "SuccinctZKVerifier 0.0.1-dev";

  /// @dev The number of bytes allocated to encoding the verifier version.
  uint256 internal constant VERIFIER_VERSION_BYTES = 4;
  /// @dev keccak256("CCIPMessageSent(uint64,bytes,bytes32,bytes)").
  bytes32 internal constant CCIP_MESSAGE_SENT_TOPIC =
    0x371bc2ff0a006f4ef863b1d27a065d4e9f938b6d883eb154572b4aea593b32cc;
  /// @dev The event signature and three indexed fields: destChainSelector, sender and messageId.
  uint256 internal constant CCIP_MESSAGE_SENT_TOPIC_COUNT = 4;
  uint256 internal constant MESSAGE_ID_TOPIC_INDEX = 3;
  /// @dev Field positions in an RLP encoded block header.
  uint256 internal constant HEADER_PARENT_HASH_INDEX = 0;
  uint256 internal constant HEADER_RECEIPTS_ROOT_INDEX = 5;
  /// @dev Field positions in an RLP encoded receipt and log.
  uint256 internal constant RECEIPT_STATUS_INDEX = 0;
  uint256 internal constant RECEIPT_LOGS_INDEX = 3;
  uint256 internal constant LOG_EMITTER_INDEX = 0;
  uint256 internal constant LOG_TOPICS_INDEX = 1;
  /// @dev EIP-2718 receipts are prefixed with the transaction type, which is at most 0x7f. Legacy receipts have no
  /// prefix and start with an RLP list byte, which is at least 0xc0.
  uint8 internal constant MAX_TRANSACTION_TYPE = 0x7f;

  // DYNAMIC CONFIG
  DynamicConfig private s_dynamicConfig;

  mapping(uint64 sourceChainSelector => SourceChainConfig sourceChainConfig) private s_sourceChainConfigs;

  constructor(
    DynamicConfig memory dynamicConfig,
    string[] memory storageLocations,
    address rmn,
    bytes4 versionTag
  ) BaseVerifier(storageLocations, rmn, versionTag) {
    _setDynamicConfig(dynamicConfig);
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

    return abi.encodePacked(versionTag());
  }

  /// @inheritdoc ICrossChainVerifierV1
  function verifyMessage(
    MessageV1Codec.MessageV1 calldata message,
    bytes32 messageId,
    bytes calldata verifierResults
  ) external view {
    _assertNotCursedByRMN(message.sourceChainSelector);

    SourceChainConfig memory sourceChainConfig = s_sourceChainConfigs[message.sourceChainSelector];
    if (address(sourceChainConfig.helios) == address(0)) revert SourceChainNotSupported(message.sourceChainSelector);

    if (verifierResults.length < VERIFIER_VERSION_BYTES) revert InvalidVerifierResults();

    // Any verifierResults submitted to this verifier should have the expected version.
    bytes4 verifierVersion = bytes4(verifierResults[:VERIFIER_VERSION_BYTES]);
    if (verifierVersion != versionTag()) revert InvalidCCVVersion(versionTag(), verifierVersion);

    Witness memory witness = abi.decode(verifierResults[VERIFIER_VERSION_BYTES:], (Witness));

    bytes32 receiptsRoot = _getReceiptsRoot(message.sourceChainSelector, sourceChainConfig, witness);

    // Receipts are keyed by the RLP encoded transaction index. This reverts unless the proof nodes hash up to the
    // receipts root along the path of that key, so the returned receipt is the one included in the message block.
    bytes memory receipt = MerkleTrie.get(RLPWriter.writeUint(witness.txIndex), witness.proofNodes, receiptsRoot);

    // The OffRamp computes messageId from the message it executes. Matching it to the messageId in the proven log
    // binds the full message to the proof.
    _validateLog(receipt, witness.logIndex, sourceChainConfig.onRamp, messageId);
  }

  /// @notice Resolves the receipts root of the message block from an anchored block hash.
  /// @dev SP1Helios anchors one block per update, so the message block is usually not anchored itself. The witness
  /// then carries the headers from the anchored block down to the message block. Each header must hash to the parent
  /// hash of the header before it, starting from the anchored block hash. Only the real headers hash this way.
  function _getReceiptsRoot(
    uint64 sourceChainSelector,
    SourceChainConfig memory sourceChainConfig,
    Witness memory witness
  ) internal view returns (bytes32 receiptsRoot) {
    uint256 headerCount = witness.headers.length;
    if (headerCount == 0) {
      receiptsRoot = sourceChainConfig.helios.executionReceiptsRoots(witness.anchorBlockNumber);
      if (receiptsRoot == bytes32(0)) revert BlockNotAnchored(sourceChainSelector, witness.anchorBlockNumber);
      return receiptsRoot;
    }
    if (headerCount > sourceChainConfig.maxHeaderChainLength) {
      revert HeaderChainTooLong(headerCount, sourceChainConfig.maxHeaderChainLength);
    }

    bytes32 expectedHash = sourceChainConfig.helios.executionBlockHashes(witness.anchorBlockNumber);
    if (expectedHash == bytes32(0)) revert BlockNotAnchored(sourceChainSelector, witness.anchorBlockNumber);

    RLPReader.RLPItem[] memory headerFields;
    for (uint256 i = 0; i < headerCount; ++i) {
      bytes32 headerHash = keccak256(witness.headers[i]);
      if (headerHash != expectedHash) revert InvalidHeaderHash(i, expectedHash, headerHash);

      headerFields = witness.headers[i].readList();
      expectedHash = _readBytes32(headerFields[HEADER_PARENT_HASH_INDEX]);
    }

    // The last header is the message block header.
    return _readBytes32(headerFields[HEADER_RECEIPTS_ROOT_INDEX]);
  }

  /// @notice Validates that the receipt is successful and that the selected log is CCIPMessageSent for the given
  /// messageId, emitted by the configured OnRamp.
  function _validateLog(
    bytes memory receipt,
    uint256 logIndex,
    address onRamp,
    bytes32 messageId
  ) internal pure {
    if (uint8(receipt[0]) <= MAX_TRANSACTION_TYPE) receipt = Bytes.slice(receipt, 1);

    RLPReader.RLPItem[] memory receiptFields = receipt.readList();

    // A reverted transaction cannot have emitted the event. Pre-Byzantium receipts carry a state root here instead
    // of a status and are rejected as well.
    bytes memory status = receiptFields[RECEIPT_STATUS_INDEX].readBytes();
    if (status.length != 1 || status[0] != 0x01) revert ReceiptNotSuccessful();

    RLPReader.RLPItem[] memory logs = receiptFields[RECEIPT_LOGS_INDEX].readList();
    if (logIndex >= logs.length) revert LogIndexOutOfRange(logIndex, logs.length);

    RLPReader.RLPItem[] memory log = logs[logIndex].readList();

    // Only the OnRamp can emit CCIPMessageSent, so any other emitter means the log is not the one we are looking for.
    bytes memory emitter = log[LOG_EMITTER_INDEX].readBytes();
    if (Internal._leftPadBytesToBytes32(emitter) != bytes32(uint256(uint160(onRamp)))) {
      revert InvalidLogEmitter(onRamp, emitter);
    }

    RLPReader.RLPItem[] memory topics = log[LOG_TOPICS_INDEX].readList();
    if (topics.length != CCIP_MESSAGE_SENT_TOPIC_COUNT) revert InvalidTopicCount(topics.length);

    bytes32 eventSignature = _readBytes32(topics[0]);
    if (eventSignature != CCIP_MESSAGE_SENT_TOPIC) revert InvalidLogTopic(CCIP_MESSAGE_SENT_TOPIC, eventSignature);

    bytes32 loggedMessageId = _readBytes32(topics[MESSAGE_ID_TOPIC_INDEX]);
    if (loggedMessageId != messageId) revert InvalidMessageId(messageId, loggedMessageId);
  }

  /// @notice Reads an RLP item that must be exactly 32 bytes, like a hash or a topic.
  function _readBytes32(
    RLPReader.RLPItem memory item
  ) internal pure returns (bytes32) {
    bytes memory data = item.readBytes();
    if (data.length != 32) revert InvalidFieldLength(32, data.length);
    return bytes32(data);
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
  /// @dev FeeTokenHandler will revert if feeAggregator is zero when withdrawing fees.
  /// @dev A zero address fee aggregator is valid, and intentionally reverts calls to withdraw fee tokens.
  function setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) external onlyOwner {
    _setDynamicConfig(dynamicConfig);
  }

  /// @notice Internal version of setDynamicConfig to allow for reuse in the constructor.
  function _setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) internal {
    s_dynamicConfig = dynamicConfig;

    emit DynamicConfigSet(dynamicConfig);
  }

  /// @notice Returns the config of a source chain.
  /// @param sourceChainSelector The source chain selector.
  /// @return sourceChainConfig The source chain config.
  function getSourceChainConfig(
    uint64 sourceChainSelector
  ) external view returns (SourceChainConfig memory sourceChainConfig) {
    return s_sourceChainConfigs[sourceChainSelector];
  }

  /// @notice Updates source chain specific configs.
  /// @param sourceChainConfigArgs Array of source chain specific configs.
  function applySourceChainConfigUpdates(
    SourceChainConfigArgs[] calldata sourceChainConfigArgs
  ) external onlyOwner {
    for (uint256 i = 0; i < sourceChainConfigArgs.length; ++i) {
      SourceChainConfigArgs calldata args = sourceChainConfigArgs[i];
      // SP1Helios only anchors one block per update, so a proof almost always needs at least one header.
      if (args.sourceChainSelector == 0 || args.onRamp == address(0) || args.maxHeaderChainLength == 0) {
        revert InvalidSourceChainConfig(args.sourceChainSelector);
      }

      // The light client can be zero to pause the source chain.
      s_sourceChainConfigs[args.sourceChainSelector] =
        SourceChainConfig({helios: args.helios, maxHeaderChainLength: args.maxHeaderChainLength, onRamp: args.onRamp});

      emit SourceChainConfigSet(args.sourceChainSelector, address(args.helios), args.onRamp, args.maxHeaderChainLength);
    }
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
  ) external onlyOwner {
    _applyAllowlistUpdates(allowlistConfigArgsItems);
  }

  /// @notice Sets the finality config according to the FinalityCodec library encoding.
  /// @param allowedFinality The finality settings allowed by this verifier.
  function setAllowedFinalityConfig(
    bytes4 allowedFinality
  ) external onlyOwner {
    _setAllowedFinalityConfig(allowedFinality);
  }

  /// @notice Updates the storage locations identifiers.
  /// @param newLocations The new storage locations identifiers.
  function updateStorageLocations(
    string[] memory newLocations
  ) external onlyOwner {
    _setStorageLocations(newLocations);
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
    FeeTokenHandler._withdrawFeeTokens(feeTokens, s_dynamicConfig.feeAggregator);
  }
}
