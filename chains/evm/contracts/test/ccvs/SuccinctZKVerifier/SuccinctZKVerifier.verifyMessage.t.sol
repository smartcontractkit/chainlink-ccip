// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SuccinctZKVerifier} from "../../../ccvs/SuccinctZKVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {MockSP1Helios} from "../../mocks/MockSP1Helios.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";

import {RLPWriter} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPWriter.sol";

contract SuccinctZKVerifier_verifyMessage is SuccinctZKVerifierSetup {
  uint256 internal constant HEADER_COUNT = 3;

  bytes32 internal s_messageId;
  bytes internal s_receipt;

  function setUp() public override {
    super.setUp();

    (, s_messageId) = _messageWithId();
    s_receipt = _encodeMessageSentReceipt(s_sourceOnRamp, s_messageId);
  }

  function _message() internal pure returns (MessageV1Codec.MessageV1 memory message) {
    (message,) = _messageWithId();
    return message;
  }

  function test_verifyMessage_HeaderChain() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);

    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_DirectAnchor() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, 0);

    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_HeaderChainAtMaxLength() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, MAX_HEADER_CHAIN_LENGTH);

    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_LegacyReceipt() public {
    bytes[] memory logs = new bytes[](1);
    logs[0] = _encodeLog(s_sourceOnRamp, _encodeMessageSentTopics(s_messageId));
    SuccinctZKVerifier.Witness memory witness = _buildWitness(_encodeReceipt(true, logs), HEADER_COUNT);

    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_SecondLog() public {
    bytes[] memory logs = new bytes[](2);
    logs[0] = _encodeLog(makeAddr("otherEmitter"), _encodeMessageSentTopics(keccak256("otherMessageId")));
    logs[1] = _encodeLog(s_sourceOnRamp, _encodeMessageSentTopics(s_messageId));
    SuccinctZKVerifier.Witness memory witness = _buildWitness(_encodeReceipt(true, logs), HEADER_COUNT);
    witness.logIndex = 1;

    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  // Reverts

  function test_verifyMessage_RevertWhen_CursedByRMN() public {
    bytes memory verifierResults = _encodeVerifierResults(_buildWitness(s_receipt, HEADER_COUNT));
    _setMockRMNChainCurse(SOURCE_CHAIN_SELECTOR, true);

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CursedByRMN.selector, SOURCE_CHAIN_SELECTOR));
    s_zkVerifier.verifyMessage(_message(), s_messageId, verifierResults);
  }

  function test_verifyMessage_RevertWhen_SourceChainNotSupported() public {
    bytes memory verifierResults = _encodeVerifierResults(_buildWitness(s_receipt, HEADER_COUNT));
    _setSourceChainConfig(MockSP1Helios(address(0)), s_sourceOnRamp, MAX_HEADER_CHAIN_LENGTH);

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.SourceChainNotSupported.selector, SOURCE_CHAIN_SELECTOR));
    s_zkVerifier.verifyMessage(_message(), s_messageId, verifierResults);
  }

  function test_verifyMessage_RevertWhen_InvalidVerifierResults() public {
    vm.expectRevert(SuccinctZKVerifier.InvalidVerifierResults.selector);
    s_zkVerifier.verifyMessage(_message(), s_messageId, hex"0102");
  }

  function test_verifyMessage_RevertWhen_InvalidCCVVersion() public {
    bytes memory verifierResults = _encodeVerifierResults(_buildWitness(s_receipt, HEADER_COUNT));
    verifierResults[0] ^= 0xff;

    vm.expectRevert(
      abi.encodeWithSelector(SuccinctZKVerifier.InvalidCCVVersion.selector, VERSION_TAG_V0_0_1, bytes4(verifierResults))
    );
    s_zkVerifier.verifyMessage(_message(), s_messageId, verifierResults);
  }

  function test_verifyMessage_RevertWhen_BlockNotAnchored_HeaderChain() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);
    s_mockHelios.setAnchor(ANCHOR_BLOCK_NUMBER, bytes32(0), bytes32(0));

    vm.expectRevert(
      abi.encodeWithSelector(SuccinctZKVerifier.BlockNotAnchored.selector, SOURCE_CHAIN_SELECTOR, ANCHOR_BLOCK_NUMBER)
    );
    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_BlockNotAnchored_DirectAnchor() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, 0);
    s_mockHelios.setAnchor(ANCHOR_BLOCK_NUMBER, bytes32(0), bytes32(0));

    vm.expectRevert(
      abi.encodeWithSelector(SuccinctZKVerifier.BlockNotAnchored.selector, SOURCE_CHAIN_SELECTOR, ANCHOR_BLOCK_NUMBER)
    );
    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_HeaderChainTooLong() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, MAX_HEADER_CHAIN_LENGTH + 1);

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.HeaderChainTooLong.selector, MAX_HEADER_CHAIN_LENGTH + 1, MAX_HEADER_CHAIN_LENGTH
      )
    );
    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidHeaderHash_FirstHeader() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);
    bytes32 anchorHash = keccak256(witness.headers[0]);
    witness.headers[0][40] ^= 0x01;

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.InvalidHeaderHash.selector, 0, anchorHash, keccak256(witness.headers[0])
      )
    );
    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidHeaderHash_MiddleHeader() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);
    bytes32 parentHash = keccak256(witness.headers[1]);
    witness.headers[1][40] ^= 0x01;

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.InvalidHeaderHash.selector, 1, parentHash, keccak256(witness.headers[1])
      )
    );
    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidProofNode() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);
    witness.proofNodes[0][40] ^= 0x01;

    vm.expectRevert("MerkleTrie: invalid root hash");
    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidTxIndex() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);
    witness.txIndex = 1;

    vm.expectRevert("MerkleTrie: path remainder must share all nibbles with key");
    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_ReceiptNotSuccessful() public {
    bytes[] memory logs = new bytes[](1);
    logs[0] = _encodeLog(s_sourceOnRamp, _encodeMessageSentTopics(s_messageId));
    SuccinctZKVerifier.Witness memory witness = _buildWitness(_encodeReceipt(false, logs), HEADER_COUNT);

    vm.expectRevert(SuccinctZKVerifier.ReceiptNotSuccessful.selector);
    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_LogIndexOutOfRange() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);
    witness.logIndex = 1;

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.LogIndexOutOfRange.selector, 1, 1));
    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidLogEmitter() public {
    address otherEmitter = makeAddr("otherEmitter");
    SuccinctZKVerifier.Witness memory witness =
      _buildWitness(_encodeMessageSentReceipt(otherEmitter, s_messageId), HEADER_COUNT);

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.InvalidLogEmitter.selector, s_sourceOnRamp, abi.encodePacked(otherEmitter)
      )
    );
    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidTopicCount() public {
    bytes[] memory topics = new bytes[](3);
    topics[0] = RLPWriter.writeBytes(abi.encodePacked(CCIP_MESSAGE_SENT_TOPIC));
    topics[1] = RLPWriter.writeBytes(abi.encodePacked(bytes32(uint256(DEST_CHAIN_SELECTOR))));
    topics[2] = RLPWriter.writeBytes(abi.encodePacked(s_messageId));
    SuccinctZKVerifier.Witness memory witness = _buildWitness(_receiptWithTopics(topics), HEADER_COUNT);

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidTopicCount.selector, 3));
    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidLogTopic() public {
    bytes32 otherTopic = keccak256("OtherEvent(uint64,bytes,bytes32,bytes)");
    bytes[] memory topics = _encodeMessageSentTopics(s_messageId);
    topics[0] = RLPWriter.writeBytes(abi.encodePacked(otherTopic));
    SuccinctZKVerifier.Witness memory witness = _buildWitness(_receiptWithTopics(topics), HEADER_COUNT);

    vm.expectRevert(
      abi.encodeWithSelector(SuccinctZKVerifier.InvalidLogTopic.selector, CCIP_MESSAGE_SENT_TOPIC, otherTopic)
    );
    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidMessageId() public {
    bytes32 otherMessageId = keccak256("otherMessageId");
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidMessageId.selector, otherMessageId, s_messageId));
    s_zkVerifier.verifyMessage(_message(), otherMessageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidFieldLength() public {
    bytes[] memory topics = _encodeMessageSentTopics(s_messageId);
    topics[0] = RLPWriter.writeBytes(hex"0102");
    SuccinctZKVerifier.Witness memory witness = _buildWitness(_receiptWithTopics(topics), HEADER_COUNT);

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidFieldLength.selector, 32, 2));
    s_zkVerifier.verifyMessage(_message(), s_messageId, _encodeVerifierResults(witness));
  }

  function _receiptWithTopics(
    bytes[] memory encodedTopics
  ) internal view returns (bytes memory) {
    bytes[] memory logs = new bytes[](1);
    logs[0] = _encodeLog(s_sourceOnRamp, encodedTopics);
    return bytes.concat(EIP1559_TRANSACTION_TYPE, _encodeReceipt(true, logs));
  }
}
