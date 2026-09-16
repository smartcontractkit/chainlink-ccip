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

/// @notice Verifies CCIPMessageSent receipt inclusion in EVM source blocks proven via the SP1Helios light client.
/// @dev The proveBlockHash function caches block hashes, allowing callers to bridge large gaps between consecutive proven blocks.
/// @dev The log emitter must match message.onRampAddress. Callers must validate the OnRamp and compute messageId from the message.
/// The OffRamp handles both before calling this verifier.
contract SuccinctZKVerifier is Ownable2StepMsgSender, ICrossChainVerifierV1, BaseVerifier {
  using RLPReader for RLPReader.RLPItem;
  using RLPReader for bytes;

  error InvalidVerifierResults();
  error InvalidCCVVersion(bytes4 expected, bytes4 got);
  error InvalidSourceChainConfig(uint64 sourceChainSelector);
  error SourceChainNotSupported(uint64 sourceChainSelector);
  error InvalidVkey(bytes32 expected, bytes32 got);
  error BlockNotAnchored(uint64 sourceChainSelector, uint256 blockNumber);
  error EmptyHeaderChain();
  error InvalidHeaderHash(uint256 headerIndex, bytes32 expected, bytes32 got);
  error InvalidFieldLength(uint256 expected, uint256 got);
  error ReceiptNotSuccessful();
  error LogIndexOutOfRange(uint256 logIndex, uint256 logCount);
  error InvalidLogEmitter(address expected, bytes got);
  error InvalidTopicCount(uint256 topicCount);
  error InvalidLogTopic(bytes32 expected, bytes32 got);
  error InvalidMessageId(bytes32 expected, bytes32 got);

  event DynamicConfigSet(DynamicConfig dynamicConfig);
  event SourceChainConfigSet(uint64 indexed sourceChainSelector, SourceChainConfigArgs sourceChainConfigArgs);
  event BlockHashProven(uint64 indexed sourceChainSelector, uint256 indexed blockNumber, bytes32 blockHash);

  struct DynamicConfig {
    address feeAggregator; // The entity receiving the withdrawn fees.
  }

  struct SourceChainConfigArgs {
    ISP1Helios helios; // ──────────╮ The light client proving this source chain. Can be zero to pause the chain.
    uint64 sourceChainSelector; // ─╯ Source chain selector.
    bytes32 lightClientVkey; // Expected vkey for beacon chain update proofs.
    bytes32 executionHeaderVkey; // Expected vkey for execution block proofs.
  }

  struct SourceChainConfig {
    ISP1Helios helios; // The light client proving this source chain. Zero means paused.
    bytes32 lightClientVkey; // Expected vkey for beacon chain update proofs.
    bytes32 executionHeaderVkey; // Expected vkey for execution block proofs.
  }

  /// @dev Verifier results format.
  ///     * Field                      Bytes      Type       Index
  ///     * verifierVersion            4          bytes4     0
  ///     * witness                    dynamic    Witness    4
  /// The witness is ABI encoded.
  struct Witness {
    uint256 anchorBlockNumber; // Source block anchored by SP1Helios or proven through proveBlockHash.
    bytes[] headers; // RLP headers from the anchor block down to the message block, both included.
    uint256 txIndex; // Index of the transaction in the message block.
    uint256 logIndex; // Index of the CCIPMessageSent log in the receipt.
    bytes[] proofNodes; // Receipts trie proof, root node first.
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "SuccinctZKVerifier 0.0.1-dev";

  uint256 internal constant VERIFIER_VERSION_BYTES = 4;
  /// @dev keccak256("CCIPMessageSent(uint64,address,bytes32,address,uint256,bytes,(address,uint32,uint32,uint256,bytes)[],bytes[])").
  bytes32 internal constant CCIP_MESSAGE_SENT_TOPIC =
    0x371bc2ff0a006f4ef863b1d27a065d4e9f938b6d883eb154572b4aea593b32cc;
  /// @dev CCIPMessageSent log topics. The event has three indexed fields, which follow the signature.
  ///     * Field                      Bytes      Type       Index
  ///     * signature                  32         bytes32    0
  ///     * destChainSelector          32         uint64     1
  ///     * sender                     32         address    2
  ///     * messageId                  32         bytes32    3
  uint256 internal constant CCIP_MESSAGE_SENT_TOPIC_COUNT = 4;
  uint256 internal constant MESSAGE_ID_TOPIC_INDEX = 3;
  /// @dev Block header format. The header is an RLP list and the index is the position in that list. Fields after
  /// receiptsRoot depend on the fork and are not read.
  ///     * Field                      Bytes      Type       Index
  ///     * parentHash                 32         bytes32    0
  ///     * ommersHash                 32         bytes32    1
  ///     * beneficiary                20         address    2
  ///     * stateRoot                  32         bytes32    3
  ///     * transactionsRoot           32         bytes32    4
  ///     * receiptsRoot               32         bytes32    5
  ///     * ...
  uint256 internal constant HEADER_PARENT_HASH_INDEX = 0;
  uint256 internal constant HEADER_RECEIPTS_ROOT_INDEX = 5;
  /// @dev Receipt format. A typed receipt is the transaction type byte followed by the RLP list. A legacy receipt is
  /// the RLP list alone. The index is the position in that list.
  ///     * Field                      Bytes      Type       Index
  ///     * status                     1          uint8      0
  ///     * cumulativeGasUsed          dynamic    uint256    1
  ///     * logsBloom                  256        bytes      2
  ///     * logs                       dynamic    Log[]      3
  /// @dev Log format. Each log is an RLP list and the index is the position in that list.
  ///     * Field                      Bytes      Type       Index
  ///     * emitter                    20         address    0
  ///     * topics                     dynamic    bytes32[]  1
  ///     * data                       dynamic    bytes      2
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

  /// @dev Keyed by light client so a replacement cannot use hashes proven under the previous client.
  mapping(ISP1Helios helios => mapping(uint256 blockNumber => bytes32 blockHash)) private s_provenBlockHashes;

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

    // For EVM, onRampAddress is abi encoded.
    address onRamp = abi.decode(message.onRampAddress, (address));

    if (verifierResults.length < VERIFIER_VERSION_BYTES) revert InvalidVerifierResults();

    // Any verifierResults submitted to this verifier should have the expected version.
    bytes4 verifierVersion = bytes4(verifierResults[:VERIFIER_VERSION_BYTES]);
    if (verifierVersion != versionTag()) revert InvalidCCVVersion(versionTag(), verifierVersion);

    Witness memory witness = abi.decode(verifierResults[VERIFIER_VERSION_BYTES:], (Witness));

    bytes32 anchorHash =
      _getAnchoredBlockHash(message.sourceChainSelector, sourceChainConfig, witness.anchorBlockNumber);
    RLPReader.RLPItem[] memory messageBlockHeader = _readHeaderChain(anchorHash, witness.headers);
    bytes32 receiptsRoot = _readBytes32(messageBlockHeader[HEADER_RECEIPTS_ROOT_INDEX]);

    // The receipts trie is keyed by the RLP encoded transaction index.
    bytes memory receipt = MerkleTrie.get(RLPWriter.writeUint(witness.txIndex), witness.proofNodes, receiptsRoot);

    // Matching the messageId binds the full message to the proven log.
    _validateLog(receipt, witness.logIndex, onRamp, messageId);
  }

  /// @notice Proves and stores a block hash from a chain of headers starting at a proven block.
  /// @dev Calls can start from a previously proven block to traverse gaps larger than the gas limit allows in one call.
  /// @param sourceChainSelector The source chain selector.
  /// @param anchorBlockNumber The number of a proven block.
  /// @param headers RLP headers from the anchor block down to the block above the one being proven, both included.
  function proveBlockHash(
    uint64 sourceChainSelector,
    uint256 anchorBlockNumber,
    bytes[] memory headers
  ) external {
    SourceChainConfig memory sourceChainConfig = s_sourceChainConfigs[sourceChainSelector];
    if (address(sourceChainConfig.helios) == address(0)) revert SourceChainNotSupported(sourceChainSelector);

    bytes32 anchorHash = _getAnchoredBlockHash(sourceChainSelector, sourceChainConfig, anchorBlockNumber);
    RLPReader.RLPItem[] memory lastHeader = _readHeaderChain(anchorHash, headers);

    // The parent of the last header is headers.length blocks before the anchor.
    uint256 blockNumber = anchorBlockNumber - headers.length;
    bytes32 blockHash = _readBytes32(lastHeader[HEADER_PARENT_HASH_INDEX]);
    s_provenBlockHashes[sourceChainConfig.helios][blockNumber] = blockHash;

    emit BlockHashProven(sourceChainSelector, blockNumber, blockHash);
  }

  /// @notice Returns an anchored or proven block hash after checking the light client's vkeys.
  function _getAnchoredBlockHash(
    uint64 sourceChainSelector,
    SourceChainConfig memory sourceChainConfig,
    uint256 blockNumber
  ) internal view returns (bytes32 blockHash) {
    bytes32 vkey = sourceChainConfig.helios.lightClientVkey();
    if (vkey != sourceChainConfig.lightClientVkey) revert InvalidVkey(sourceChainConfig.lightClientVkey, vkey);
    vkey = sourceChainConfig.helios.executionHeaderVkey();
    if (vkey != sourceChainConfig.executionHeaderVkey) revert InvalidVkey(sourceChainConfig.executionHeaderVkey, vkey);

    blockHash = sourceChainConfig.helios.executionBlockHashes(blockNumber);
    if (blockHash == bytes32(0)) blockHash = s_provenBlockHashes[sourceChainConfig.helios][blockNumber];
    if (blockHash == bytes32(0)) revert BlockNotAnchored(sourceChainSelector, blockNumber);
    return blockHash;
  }

  /// @notice Validates a chain of RLP block headers against an anchor hash.
  /// @dev Each subsequent header must hash to the previous header's parent hash.
  /// @return lastHeader The RLP fields of the last header in the chain.
  function _readHeaderChain(
    bytes32 anchorHash,
    bytes[] memory headers
  ) internal pure returns (RLPReader.RLPItem[] memory lastHeader) {
    if (headers.length == 0) revert EmptyHeaderChain();

    bytes32 expectedHash = anchorHash;
    for (uint256 i = 0; i < headers.length; ++i) {
      bytes32 headerHash = keccak256(headers[i]);
      if (headerHash != expectedHash) revert InvalidHeaderHash(i, expectedHash, headerHash);

      lastHeader = headers[i].readList();
      expectedHash = _readBytes32(lastHeader[HEADER_PARENT_HASH_INDEX]);
    }
    return lastHeader;
  }

  /// @notice Validates that the receipt is successful and that the selected log is CCIPMessageSent for the given
  /// messageId, emitted by the given OnRamp.
  function _validateLog(
    bytes memory receipt,
    uint256 logIndex,
    address onRamp,
    bytes32 messageId
  ) internal pure {
    if (uint8(receipt[0]) <= MAX_TRANSACTION_TYPE) receipt = Bytes.slice(receipt, 1);

    RLPReader.RLPItem[] memory receiptFields = receipt.readList();

    // Require a successful transaction. This also rejects pre-Byzantium receipts, which contain a state root.
    bytes memory status = receiptFields[RECEIPT_STATUS_INDEX].readBytes();
    if (status.length != 1 || status[0] != 0x01) revert ReceiptNotSuccessful();

    RLPReader.RLPItem[] memory log =
      _readListItem(receiptFields[RECEIPT_LOGS_INDEX].readRawBytes(), logIndex).readList();

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

  /// @notice Returns one item of an RLP encoded list without decoding the other items.
  /// @dev RLPReader.readList caps lists at 32 items, but receipts can contain more logs. Walk by item length to avoid
  /// that limit. The receipt proof authenticates the encoding, so no additional RLP validation is needed here.
  function _readListItem(
    bytes memory encodedList,
    uint256 index
  ) internal pure returns (bytes memory) {
    (uint256 offset, uint256 length) = _readItemBounds(encodedList, 0);
    uint256 end = offset + length;

    uint256 itemCount = 0;
    while (offset < end) {
      (uint256 itemOffset, uint256 itemLength) = _readItemBounds(encodedList, offset);
      if (itemCount == index) return Bytes.slice(encodedList, offset, itemOffset + itemLength - offset);
      offset = itemOffset + itemLength;
      ++itemCount;
    }
    revert LogIndexOutOfRange(index, itemCount);
  }

  /// @notice Reads the payload offset and length of the RLP item that starts at the given offset.
  function _readItemBounds(
    bytes memory encoded,
    uint256 offset
  ) internal pure returns (uint256 payloadOffset, uint256 payloadLength) {
    uint8 prefix = uint8(encoded[offset]);
    // A single byte below 0x80 is its own encoding.
    if (prefix < 0x80) return (offset, 1);

    // Short strings and lists carry the payload length in the prefix. Long ones carry the size of the length in the
    // prefix and the length itself in the bytes that follow.
    uint256 lengthSize = 0;
    if (prefix <= 0xb7) {
      payloadLength = prefix - 0x80;
    } else if (prefix <= 0xbf) {
      lengthSize = prefix - 0xb7;
    } else if (prefix <= 0xf7) {
      payloadLength = prefix - 0xc0;
    } else {
      lengthSize = prefix - 0xf7;
    }

    for (uint256 i = 1; i <= lengthSize; ++i) {
      payloadLength = (payloadLength << 8) | uint8(encoded[offset + i]);
    }
    return (offset + 1 + lengthSize, payloadLength);
  }

  /// @notice Reads an RLP value as bytes32.
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
  /// @dev A zero feeAggregator disables fee withdrawals.
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

  /// @notice Returns the hash of a source chain block proven through proveBlockHash, or zero if none.
  /// @param sourceChainSelector The source chain selector.
  /// @param blockNumber The source chain block number.
  function getProvenBlockHash(
    uint64 sourceChainSelector,
    uint256 blockNumber
  ) external view returns (bytes32 blockHash) {
    return s_provenBlockHashes[s_sourceChainConfigs[sourceChainSelector].helios][blockNumber];
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
      if (args.sourceChainSelector == 0) revert InvalidSourceChainConfig(args.sourceChainSelector);
      // The light client can be zero to pause the source chain. Otherwise both of its vkeys must be pinned.
      if (
        address(args.helios) != address(0)
          && (args.lightClientVkey == bytes32(0) || args.executionHeaderVkey == bytes32(0))
      ) {
        revert InvalidSourceChainConfig(args.sourceChainSelector);
      }

      s_sourceChainConfigs[args.sourceChainSelector] = SourceChainConfig({
        helios: args.helios, lightClientVkey: args.lightClientVkey, executionHeaderVkey: args.executionHeaderVkey
      });

      emit SourceChainConfigSet(args.sourceChainSelector, args);
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
