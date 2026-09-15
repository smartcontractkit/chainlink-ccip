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
  bytes32 internal constant LIGHT_CLIENT_VKEY = keccak256("lightClientVkey");
  bytes32 internal constant EXECUTION_HEADER_VKEY = keccak256("executionHeaderVkey");
  bytes1 internal constant EIP1559_TRANSACTION_TYPE = 0x02;
  uint16 internal constant MAX_HEADER_CHAIN_LENGTH = 8;
  uint256 internal constant ANCHOR_BLOCK_NUMBER = 1_000_000;

  SuccinctZKVerifier internal s_zkVerifier;
  MockSP1Helios internal s_mockHelios;
  address internal s_sourceOnRamp;

  function setUp() public virtual override {
    super.setUp();

    s_sourceOnRamp = abi.decode(_createBasicMessageV1(SOURCE_CHAIN_SELECTOR).onRampAddress, (address));
    s_mockHelios = new MockSP1Helios();
    s_mockHelios.setVkeys(LIGHT_CLIENT_VKEY, EXECUTION_HEADER_VKEY);
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
    address[] memory onRamps = new address[](1);
    onRamps[0] = onRamp;
    SuccinctZKVerifier.SourceChainConfigArgs[] memory sourceChainConfigs =
      new SuccinctZKVerifier.SourceChainConfigArgs[](1);
    sourceChainConfigs[0] = SuccinctZKVerifier.SourceChainConfigArgs({
      helios: helios,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      maxHeaderChainLength: maxHeaderChainLength,
      lightClientVkey: LIGHT_CLIENT_VKEY,
      executionHeaderVkey: EXECUTION_HEADER_VKEY,
      onRamps: onRamps
    });
    s_zkVerifier.applySourceChainConfigUpdates(sourceChainConfigs);
  }

  function _encodeVerifierResults(
    SuccinctZKVerifier.Witness memory witness
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(VERSION_TAG_V0_0_1, abi.encode(witness));
  }

  /// @notice Builds a witness for the receipt of transaction 0 in a block reached through headerCount headers from
  /// the anchored block. The last header is the message block, the first one is the anchored block itself.
  function _buildWitness(
    bytes memory receipt,
    uint256 headerCount
  ) internal returns (SuccinctZKVerifier.Witness memory) {
    (bytes[] memory proofNodes, bytes32 receiptsRoot) = _buildReceiptsTrie(receipt);

    bytes[] memory headers = new bytes[](headerCount);
    headers[headerCount - 1] = _encodeHeader(keccak256("parent"), receiptsRoot, ANCHOR_BLOCK_NUMBER - headerCount + 1);
    for (uint256 i = headerCount - 1; i > 0; --i) {
      headers[i - 1] = _encodeHeader(keccak256(headers[i]), keccak256("otherReceiptsRoot"), ANCHOR_BLOCK_NUMBER - i + 1);
    }
    s_mockHelios.setExecutionBlockHash(ANCHOR_BLOCK_NUMBER, keccak256(headers[0]));

    return SuccinctZKVerifier.Witness({
      anchorBlockNumber: ANCHOR_BLOCK_NUMBER, headers: headers, txIndex: 0, logIndex: 0, proofNodes: proofNodes
    });
  }

  /// @notice Builds a receipts trie holding the receipt at transaction 0 and an empty receipt at transaction 1, and
  /// returns the proof of transaction 0. The keys RLP(0) = 0x80 and RLP(1) = 0x01 differ in their first nibble, so
  /// the root is a branch node with a leaf under nibble 8 and a leaf under nibble 0.
  function _buildReceiptsTrie(
    bytes memory receipt
  ) internal pure returns (bytes[] memory proofNodes, bytes32 receiptsRoot) {
    // The leaf path holds the remaining odd nibble behind the leaf prefix 0x3.
    bytes memory leaf = _encodeLeaf(hex"30", receipt);
    bytes memory otherLeaf = _encodeLeaf(hex"31", _encodeReceipt(true, new bytes[](0)));

    bytes[] memory branch = new bytes[](17);
    for (uint256 i = 0; i < branch.length; ++i) {
      branch[i] = RLPWriter.writeBytes("");
    }
    branch[8] = RLPWriter.writeBytes(abi.encodePacked(keccak256(leaf)));
    branch[0] = RLPWriter.writeBytes(abi.encodePacked(keccak256(otherLeaf)));

    proofNodes = new bytes[](2);
    proofNodes[0] = RLPWriter.writeList(branch);
    proofNodes[1] = leaf;
    return (proofNodes, keccak256(proofNodes[0]));
  }

  function _encodeLeaf(
    bytes memory path,
    bytes memory value
  ) internal pure returns (bytes memory) {
    bytes[] memory fields = new bytes[](2);
    fields[0] = RLPWriter.writeBytes(path);
    fields[1] = RLPWriter.writeBytes(value);
    return RLPWriter.writeList(fields);
  }

  /// @notice Encodes a block header with the field layout of a Prague block. The verifier only reads parentHash and
  /// receiptsRoot, the other fields hold placeholder values of the right size.
  function _encodeHeader(
    bytes32 parentHash,
    bytes32 receiptsRoot,
    uint256 blockNumber
  ) internal pure returns (bytes memory) {
    bytes[] memory fields = new bytes[](21);
    fields[0] = RLPWriter.writeBytes(abi.encodePacked(parentHash));
    fields[1] = RLPWriter.writeBytes(abi.encodePacked(keccak256("ommersHash")));
    fields[2] = RLPWriter.writeBytes(new bytes(20)); // coinbase
    fields[3] = RLPWriter.writeBytes(abi.encodePacked(keccak256("stateRoot")));
    fields[4] = RLPWriter.writeBytes(abi.encodePacked(keccak256("transactionsRoot")));
    fields[5] = RLPWriter.writeBytes(abi.encodePacked(receiptsRoot));
    fields[6] = RLPWriter.writeBytes(new bytes(256)); // logsBloom
    fields[7] = RLPWriter.writeUint(0); // difficulty
    fields[8] = RLPWriter.writeUint(blockNumber);
    fields[9] = RLPWriter.writeUint(30_000_000); // gasLimit
    fields[10] = RLPWriter.writeUint(21_000); // gasUsed
    fields[11] = RLPWriter.writeUint(1_700_000_000); // timestamp
    fields[12] = RLPWriter.writeBytes(""); // extraData
    fields[13] = RLPWriter.writeBytes(abi.encodePacked(keccak256("mixHash")));
    fields[14] = RLPWriter.writeBytes(new bytes(8)); // nonce
    fields[15] = RLPWriter.writeUint(1 gwei); // baseFeePerGas
    fields[16] = RLPWriter.writeBytes(abi.encodePacked(keccak256("withdrawalsRoot")));
    fields[17] = RLPWriter.writeUint(0); // blobGasUsed
    fields[18] = RLPWriter.writeUint(0); // excessBlobGas
    fields[19] = RLPWriter.writeBytes(abi.encodePacked(keccak256("parentBeaconBlockRoot")));
    fields[20] = RLPWriter.writeBytes(abi.encodePacked(keccak256("requestsHash")));
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

  /// @notice Encodes the topics of CCIPMessageSent: the event signature, then the indexed destChainSelector, sender
  /// and messageId.
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
