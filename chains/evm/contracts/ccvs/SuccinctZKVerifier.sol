// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../interfaces/ICrossChainVerifierV1.sol";
import {IExecutionRootOracle} from "../interfaces/IExecutionRootOracle.sol";

import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {Bytes} from "../libraries/mpt/Bytes.sol";
import {MerkleTrie} from "../libraries/mpt/MerkleTrie.sol";
import {RLPReader} from "../libraries/mpt/RLPReader.sol";
import {BaseVerifier} from "./components/BaseVerifier.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @notice Verifies a CCIPMessageSent event by proving receipt inclusion in a source block anchored by a ZK light
/// client on this chain.
/// @dev The light client anchors one block per update. The witness carries a header chain from the anchored block down
/// to the message block, then a Merkle Patricia Trie proof of the receipt. An empty header chain means the message
/// block itself is anchored.
contract SuccinctZKVerifier is Ownable2StepMsgSender, ICrossChainVerifierV1, BaseVerifier {
  using RLPReader for RLPReader.RLPItem;
  using RLPReader for bytes;

  error InvalidVerifierResults();
  error InvalidCCVVersion(bytes4 verifierVersion);
  error SourceChainNotConfigured(uint64 sourceChainSelector);
  error AnchorNotAvailable(uint64 sourceChainSelector, uint64 anchorBlockNumber);
  error HeaderChainTooLong(uint256 headerCount, uint256 maxHeaderChainLength);
  error HeaderHashMismatch(uint256 headerIndex, bytes32 expected, bytes32 actual);
  error ReceiptNotSuccessful();
  error LogIndexOutOfRange(uint256 logIndex, uint256 logCount);
  error UnexpectedLogEmitter(address emitter, address onRamp);
  error UnexpectedLogTopic(bytes32 topic);
  error InvalidMessageId(bytes32 expected, bytes32 actual);
  error InvalidFieldLength(uint256 expected, uint256 actual);

  event SourceChainConfigSet(
    uint64 indexed sourceChainSelector, address rootOracle, address onRamp, uint16 maxHeaderChainLength
  );

  struct SourceChainConfigArgs {
    uint64 sourceChainSelector;
    IExecutionRootOracle rootOracle; // Light client that anchors the source chain on this chain.
    address onRamp; // Expected emitter of CCIPMessageSent on the source chain.
    uint16 maxHeaderChainLength; // Upper bound on headers in one witness.
  }

  struct SourceChainConfig {
    IExecutionRootOracle rootOracle;
    address onRamp;
    uint16 maxHeaderChainLength;
  }

  /// @dev Decoded witness. See `_decodeWitness` for the byte layout.
  struct Witness {
    uint64 anchorBlockNumber;
    bytes[] headers; // RLP headers, anchor block first, message block last.
    uint16 txIndex;
    uint8 logIndex; // Index within the receipt's logs.
    bytes[] proofNodes; // Receipts trie proof, root first.
  }

  string public constant override typeAndVersion = "SuccinctZKVerifier 2.0.0";

  /// @dev keccak256 of the OnRamp CCIPMessageSent event signature.
  bytes32 internal constant CCIP_MESSAGE_SENT_TOPIC =
    0x371bc2ff0a006f4ef863b1d27a065d4e9f938b6d883eb154572b4aea593b32cc;
  /// @dev Number of topics of the CCIPMessageSent event: signature, destChainSelector, sender, messageId.
  uint256 internal constant CCIP_MESSAGE_SENT_TOPIC_COUNT = 4;
  uint256 internal constant MESSAGE_ID_TOPIC_INDEX = 3;

  uint256 internal constant HEADER_PARENT_HASH_INDEX = 0;
  uint256 internal constant HEADER_RECEIPTS_ROOT_INDEX = 5;
  uint256 internal constant RECEIPT_STATUS_INDEX = 0;
  uint256 internal constant RECEIPT_LOGS_INDEX = 3;
  uint256 internal constant LOG_EMITTER_INDEX = 0;
  uint256 internal constant LOG_TOPICS_INDEX = 1;

  uint256 internal constant VERSION_TAG_BYTES = 4;
  uint256 internal constant ANCHOR_BLOCK_NUMBER_BYTES = 8;
  uint256 internal constant COUNT_BYTES = 2;
  uint256 internal constant LENGTH_PREFIX_BYTES = 2;
  uint256 internal constant TX_INDEX_BYTES = 2;
  uint256 internal constant LOG_INDEX_BYTES = 1;
  uint256 internal constant NODE_COUNT_BYTES = 1;

  mapping(uint64 sourceChainSelector => SourceChainConfig sourceChainConfig) private s_sourceChainConfigs;

  constructor(
    string[] memory storageLocations,
    address rmn,
    bytes4 versionTag
  ) BaseVerifier(storageLocations, rmn, versionTag) {}

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

    SourceChainConfig memory config = s_sourceChainConfigs[message.sourceChainSelector];
    if (address(config.rootOracle) == address(0)) {
      revert SourceChainNotConfigured(message.sourceChainSelector);
    }

    Witness memory witness = _decodeWitness(verifierResults);

    bytes32 receiptsRoot = _resolveReceiptsRoot(message.sourceChainSelector, config, witness);
    bytes memory receipt = MerkleTrie.get(_receiptTrieKey(witness.txIndex), witness.proofNodes, receiptsRoot);

    _assertLogMatches(receipt, witness.logIndex, config.onRamp, messageId);
  }

  /// @notice Decodes the witness envelope.
  /// @dev Byte layout:
  /// [versionTag 4][anchorBlockNumber 8][headerCount 2]
  /// headerCount times: [headerLen 2][header RLP]
  /// [txIndex 2][logIndex 1][proofNodeCount 1]
  /// proofNodeCount times: [nodeLen 2][node RLP]
  function _decodeWitness(
    bytes calldata data
  ) internal view returns (Witness memory witness) {
    if (data.length < VERSION_TAG_BYTES + ANCHOR_BLOCK_NUMBER_BYTES + COUNT_BYTES) {
      revert InvalidVerifierResults();
    }

    bytes4 verifierVersion = bytes4(data[:VERSION_TAG_BYTES]);
    if (verifierVersion != versionTag()) {
      revert InvalidCCVVersion(verifierVersion);
    }
    uint256 offset = VERSION_TAG_BYTES;

    witness.anchorBlockNumber = uint64(bytes8(data[offset:offset + ANCHOR_BLOCK_NUMBER_BYTES]));
    offset += ANCHOR_BLOCK_NUMBER_BYTES;

    uint256 headerCount = uint16(bytes2(data[offset:offset + COUNT_BYTES]));
    offset += COUNT_BYTES;
    witness.headers = new bytes[](headerCount);
    for (uint256 i = 0; i < headerCount; ++i) {
      (witness.headers[i], offset) = _readLengthPrefixed(data, offset);
    }

    if (data.length < offset + TX_INDEX_BYTES + LOG_INDEX_BYTES + NODE_COUNT_BYTES) {
      revert InvalidVerifierResults();
    }
    witness.txIndex = uint16(bytes2(data[offset:offset + TX_INDEX_BYTES]));
    offset += TX_INDEX_BYTES;
    witness.logIndex = uint8(data[offset]);
    offset += LOG_INDEX_BYTES;
    uint256 nodeCount = uint8(data[offset]);
    offset += NODE_COUNT_BYTES;

    witness.proofNodes = new bytes[](nodeCount);
    for (uint256 i = 0; i < nodeCount; ++i) {
      (witness.proofNodes[i], offset) = _readLengthPrefixed(data, offset);
    }

    if (offset != data.length) {
      revert InvalidVerifierResults();
    }
    return witness;
  }

  function _readLengthPrefixed(
    bytes calldata data,
    uint256 offset
  ) internal pure returns (bytes memory item, uint256 newOffset) {
    if (data.length < offset + LENGTH_PREFIX_BYTES) {
      revert InvalidVerifierResults();
    }
    uint256 length = uint16(bytes2(data[offset:offset + LENGTH_PREFIX_BYTES]));
    offset += LENGTH_PREFIX_BYTES;
    if (data.length < offset + length) {
      revert InvalidVerifierResults();
    }
    return (data[offset:offset + length], offset + length);
  }

  /// @notice Returns the receipts root of the message block. With headers present, the first header must hash to the
  /// anchored block hash and each next header must hash to the parent hash of the previous one.
  function _resolveReceiptsRoot(
    uint64 sourceChainSelector,
    SourceChainConfig memory config,
    Witness memory witness
  ) internal view returns (bytes32 receiptsRoot) {
    uint256 headerCount = witness.headers.length;
    if (headerCount == 0) {
      receiptsRoot = config.rootOracle.executionReceiptsRoots(witness.anchorBlockNumber);
      if (receiptsRoot == bytes32(0)) {
        revert AnchorNotAvailable(sourceChainSelector, witness.anchorBlockNumber);
      }
      return receiptsRoot;
    }
    if (headerCount > config.maxHeaderChainLength) {
      revert HeaderChainTooLong(headerCount, config.maxHeaderChainLength);
    }

    bytes32 expectedHash = config.rootOracle.executionBlockHashes(witness.anchorBlockNumber);
    if (expectedHash == bytes32(0)) {
      revert AnchorNotAvailable(sourceChainSelector, witness.anchorBlockNumber);
    }

    RLPReader.RLPItem[] memory headerFields;
    for (uint256 i = 0; i < headerCount; ++i) {
      bytes32 actualHash = keccak256(witness.headers[i]);
      if (actualHash != expectedHash) {
        revert HeaderHashMismatch(i, expectedHash, actualHash);
      }
      headerFields = witness.headers[i].readList();
      expectedHash = _toBytes32(headerFields[HEADER_PARENT_HASH_INDEX].readBytes());
    }

    return _toBytes32(headerFields[HEADER_RECEIPTS_ROOT_INDEX].readBytes());
  }

  /// @notice Checks that the receipt succeeded and that the selected log is CCIPMessageSent for messageId from onRamp.
  function _assertLogMatches(
    bytes memory receipt,
    uint8 logIndex,
    address onRamp,
    bytes32 messageId
  ) internal pure {
    // Typed receipts carry a leading transaction type byte before the RLP list.
    if (uint8(receipt[0]) < 0xc0) {
      receipt = Bytes.slice(receipt, 1);
    }
    RLPReader.RLPItem[] memory receiptFields = receipt.readList();

    bytes memory status = receiptFields[RECEIPT_STATUS_INDEX].readBytes();
    if (status.length != 1 || status[0] != 0x01) {
      revert ReceiptNotSuccessful();
    }

    RLPReader.RLPItem[] memory logs = receiptFields[RECEIPT_LOGS_INDEX].readList();
    if (logIndex >= logs.length) {
      revert LogIndexOutOfRange(logIndex, logs.length);
    }
    RLPReader.RLPItem[] memory log = logs[logIndex].readList();

    address emitter = _toAddress(log[LOG_EMITTER_INDEX].readBytes());
    if (emitter != onRamp) {
      revert UnexpectedLogEmitter(emitter, onRamp);
    }

    RLPReader.RLPItem[] memory topics = log[LOG_TOPICS_INDEX].readList();
    if (topics.length != CCIP_MESSAGE_SENT_TOPIC_COUNT) {
      revert InvalidVerifierResults();
    }
    bytes32 eventTopic = _toBytes32(topics[0].readBytes());
    if (eventTopic != CCIP_MESSAGE_SENT_TOPIC) {
      revert UnexpectedLogTopic(eventTopic);
    }
    bytes32 loggedMessageId = _toBytes32(topics[MESSAGE_ID_TOPIC_INDEX].readBytes());
    if (loggedMessageId != messageId) {
      revert InvalidMessageId(messageId, loggedMessageId);
    }
  }

  /// @notice Returns the receipts trie key, which is the RLP encoding of the transaction index.
  function _receiptTrieKey(
    uint16 txIndex
  ) internal pure returns (bytes memory) {
    if (txIndex == 0) {
      return hex"80";
    }
    if (txIndex < 0x80) {
      return abi.encodePacked(uint8(txIndex));
    }
    if (txIndex < 0x100) {
      return abi.encodePacked(uint8(0x81), uint8(txIndex));
    }
    return abi.encodePacked(uint8(0x82), txIndex);
  }

  function _toBytes32(
    bytes memory data
  ) internal pure returns (bytes32 out) {
    if (data.length != 32) {
      revert InvalidFieldLength(32, data.length);
    }
    // solhint-disable-next-line no-inline-assembly
    assembly {
      out := mload(add(data, 32))
    }
  }

  function _toAddress(
    bytes memory data
  ) internal pure returns (address out) {
    if (data.length != 20) {
      revert InvalidFieldLength(20, data.length);
    }
    // solhint-disable-next-line no-inline-assembly
    assembly {
      out := shr(96, mload(add(data, 32)))
    }
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Returns the source chain config for the given selector.
  function getSourceChainConfig(
    uint64 sourceChainSelector
  ) external view returns (SourceChainConfig memory) {
    return s_sourceChainConfigs[sourceChainSelector];
  }

  /// @notice Sets the root oracle, expected OnRamp and header chain bound per source chain.
  function applySourceChainConfigUpdates(
    SourceChainConfigArgs[] calldata sourceChainConfigArgs
  ) external onlyOwner {
    for (uint256 i = 0; i < sourceChainConfigArgs.length; ++i) {
      SourceChainConfigArgs calldata args = sourceChainConfigArgs[i];
      if (args.sourceChainSelector == 0 || address(args.rootOracle) == address(0) || args.onRamp == address(0)) {
        revert ZeroAddressNotAllowed();
      }
      s_sourceChainConfigs[args.sourceChainSelector] = SourceChainConfig({
        rootOracle: args.rootOracle, onRamp: args.onRamp, maxHeaderChainLength: args.maxHeaderChainLength
      });

      emit SourceChainConfigSet(
        args.sourceChainSelector, address(args.rootOracle), args.onRamp, args.maxHeaderChainLength
      );
    }
  }

  /// @notice Updates the remote chain configs.
  function applyRemoteChainConfigUpdates(
    RemoteChainConfigArgs[] calldata remoteChainConfigArgs
  ) external onlyOwner {
    _applyRemoteChainConfigUpdates(remoteChainConfigArgs);
  }

  /// @notice Updates the allowlist per destination chain.
  function applyAllowlistUpdates(
    AllowlistConfigArgs[] calldata allowlistConfigArgsItems
  ) external onlyOwner {
    _applyAllowlistUpdates(allowlistConfigArgsItems);
  }

  /// @notice Sets the allowed finality config, see FinalityCodec.
  function setAllowedFinalityConfig(
    bytes4 allowedFinality
  ) external onlyOwner {
    _setAllowedFinalityConfig(allowedFinality);
  }

  /// @notice Updates the storage locations.
  function updateStorageLocations(
    string[] calldata storageLocations
  ) external onlyOwner {
    _setStorageLocations(storageLocations);
  }
}
