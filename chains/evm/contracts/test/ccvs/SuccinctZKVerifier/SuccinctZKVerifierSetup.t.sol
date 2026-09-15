// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";

import {SuccinctZKVerifier} from "../../../ccvs/SuccinctZKVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {MockSP1Helios} from "../../mocks/MockSP1Helios.sol";
import {BaseVerifierSetup} from "../components/BaseVerifier/BaseVerifierSetup.t.sol";

import {RLPWriter} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPWriter.sol";

contract SuccinctZKVerifierSetup is BaseVerifierSetup {
  bytes4 internal constant VERSION_TAG_V0_0_1 = bytes4(keccak256("SuccinctZKVerifier 0.0.1-dev"));
  bytes32 internal constant CCIP_MESSAGE_SENT_TOPIC =
    0x371bc2ff0a006f4ef863b1d27a065d4e9f938b6d883eb154572b4aea593b32cc;
  bytes1 internal constant EIP1559_TRANSACTION_TYPE = 0x02;
  uint16 internal constant MAX_HEADER_CHAIN_LENGTH = 8;
  uint256 internal constant ANCHOR_BLOCK_NUMBER = 1_000_000;

  SuccinctZKVerifier internal s_zkVerifier;
  MockSP1Helios internal s_mockHelios;
  address internal s_sourceOnRamp = makeAddr("sourceOnRamp");

  function setUp() public virtual override {
    super.setUp();

    s_mockHelios = new MockSP1Helios();
    s_zkVerifier = new SuccinctZKVerifier(
      SuccinctZKVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR}),
      s_storageLocations,
      address(s_mockRMNRemote),
      VERSION_TAG_V0_0_1
    );

    BaseVerifier.RemoteChainConfigArgs[] memory remoteChainConfigs = new BaseVerifier.RemoteChainConfigArgs[](1);
    remoteChainConfigs[0] = _getRemoteChainConfig(s_router, DEST_CHAIN_SELECTOR, false);
    s_zkVerifier.applyRemoteChainConfigUpdates(remoteChainConfigs);
    vm.mockCall(address(s_router), abi.encodeCall(IRouter.getOnRamp, (DEST_CHAIN_SELECTOR)), abi.encode(s_onRamp));

    _setSourceChainConfig(s_mockHelios, s_sourceOnRamp, MAX_HEADER_CHAIN_LENGTH);
  }

  function _setSourceChainConfig(
    MockSP1Helios helios,
    address onRamp,
    uint16 maxHeaderChainLength
  ) internal {
    SuccinctZKVerifier.SourceChainConfigArgs[] memory sourceChainConfigs =
      new SuccinctZKVerifier.SourceChainConfigArgs[](1);
    sourceChainConfigs[0] = SuccinctZKVerifier.SourceChainConfigArgs({
      helios: helios,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      maxHeaderChainLength: maxHeaderChainLength,
      onRamp: onRamp
    });
    s_zkVerifier.applySourceChainConfigUpdates(sourceChainConfigs);
  }

  function _encodeVerifierResults(
    SuccinctZKVerifier.Witness memory witness
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(VERSION_TAG_V0_0_1, abi.encode(witness));
  }

  /// @notice Builds a witness for a block with a single transaction, whose receipts trie is one leaf node. The block is
  /// anchored on the mock light client either directly or through headerCount headers.
  function _buildWitness(
    bytes memory receipt,
    uint256 headerCount
  ) internal returns (SuccinctZKVerifier.Witness memory) {
    bytes[] memory proofNodes = new bytes[](1);
    proofNodes[0] = _encodeLeaf(receipt);
    bytes32 receiptsRoot = keccak256(proofNodes[0]);

    bytes[] memory headers = new bytes[](headerCount);
    if (headerCount == 0) {
      s_mockHelios.setAnchor(ANCHOR_BLOCK_NUMBER, bytes32(0), receiptsRoot);
    } else {
      // The last header is the message block. Each header before it is the parent of the one after it.
      headers[headerCount - 1] = _encodeHeader(keccak256("parent"), receiptsRoot);
      for (uint256 i = headerCount - 1; i > 0; --i) {
        headers[i - 1] = _encodeHeader(keccak256(headers[i]), keccak256("otherReceiptsRoot"));
      }
      s_mockHelios.setAnchor(ANCHOR_BLOCK_NUMBER, keccak256(headers[0]), bytes32(0));
    }

    return SuccinctZKVerifier.Witness({
      anchorBlockNumber: ANCHOR_BLOCK_NUMBER, headers: headers, txIndex: 0, logIndex: 0, proofNodes: proofNodes
    });
  }

  /// @notice Encodes the leaf node of a receipts trie holding only transaction 0. The key is RLP(0), which is 0x80,
  /// so the leaf path is the two nibbles 8 and 0 with the even length leaf prefix 0x20.
  function _encodeLeaf(
    bytes memory receipt
  ) internal pure returns (bytes memory) {
    bytes[] memory fields = new bytes[](2);
    fields[0] = RLPWriter.writeBytes(hex"2080");
    fields[1] = RLPWriter.writeBytes(receipt);
    return RLPWriter.writeList(fields);
  }

  /// @notice Encodes a block header with only the fields the verifier reads set: parentHash and receiptsRoot.
  function _encodeHeader(
    bytes32 parentHash,
    bytes32 receiptsRoot
  ) internal pure returns (bytes memory) {
    bytes[] memory fields = new bytes[](6);
    fields[0] = RLPWriter.writeBytes(abi.encodePacked(parentHash));
    for (uint256 i = 1; i < 5; ++i) {
      fields[i] = RLPWriter.writeUint(0);
    }
    fields[5] = RLPWriter.writeBytes(abi.encodePacked(receiptsRoot));
    return RLPWriter.writeList(fields);
  }

  /// @notice Encodes an EIP-1559 receipt with a single CCIPMessageSent log.
  function _encodeMessageSentReceipt(
    address emitter,
    bytes32 messageId
  ) internal pure returns (bytes memory) {
    bytes[] memory logs = new bytes[](1);
    logs[0] = _encodeLog(emitter, _encodeMessageSentTopics(messageId));
    return bytes.concat(EIP1559_TRANSACTION_TYPE, _encodeReceipt(true, logs));
  }

  /// @notice Encodes a receipt as RLP([status, cumulativeGasUsed, logsBloom, logs]).
  function _encodeReceipt(
    bool success,
    bytes[] memory encodedLogs
  ) internal pure returns (bytes memory) {
    bytes[] memory fields = new bytes[](4);
    fields[0] = RLPWriter.writeUint(success ? 1 : 0);
    fields[1] = RLPWriter.writeUint(21_000);
    fields[2] = RLPWriter.writeBytes(new bytes(256));
    fields[3] = RLPWriter.writeList(encodedLogs);
    return RLPWriter.writeList(fields);
  }

  /// @notice Encodes a log as RLP([emitter, topics, data]).
  function _encodeLog(
    address emitter,
    bytes[] memory encodedTopics
  ) internal pure returns (bytes memory) {
    bytes[] memory fields = new bytes[](3);
    fields[0] = RLPWriter.writeBytes(abi.encodePacked(emitter));
    fields[1] = RLPWriter.writeList(encodedTopics);
    fields[2] = RLPWriter.writeBytes("");
    return RLPWriter.writeList(fields);
  }

  /// @notice Encodes the topics of CCIPMessageSent(uint64 indexed destChainSelector, bytes indexed sender, bytes32
  /// indexed messageId, bytes encodedMessage).
  function _encodeMessageSentTopics(
    bytes32 messageId
  ) internal pure returns (bytes[] memory) {
    bytes[] memory topics = new bytes[](4);
    topics[0] = RLPWriter.writeBytes(abi.encodePacked(CCIP_MESSAGE_SENT_TOPIC));
    topics[1] = RLPWriter.writeBytes(abi.encodePacked(bytes32(uint256(DEST_CHAIN_SELECTOR))));
    topics[2] = RLPWriter.writeBytes(abi.encodePacked(keccak256(abi.encode(OWNER))));
    topics[3] = RLPWriter.writeBytes(abi.encodePacked(messageId));
    return topics;
  }

  function _messageWithId() internal pure returns (MessageV1Codec.MessageV1 memory message, bytes32 messageId) {
    message = _createBasicMessageV1(SOURCE_CHAIN_SELECTOR);
    return (message, keccak256(MessageV1Codec._encodeMessageV1(message)));
  }
}
