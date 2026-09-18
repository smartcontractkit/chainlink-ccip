// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SuccinctZKVerifier} from "../../../ccvs/SuccinctZKVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {MockSP1Helios} from "../../mocks/MockSP1Helios.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";

import {RLPWriter} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPWriter.sol";
import {Vm} from "forge-std/Vm.sol";

contract SuccinctZKVerifier_verifyMessage is SuccinctZKVerifierSetup {
  uint256 internal constant HEADER_COUNT = 3;

  bytes32 internal s_messageId;
  bytes internal s_receipt;

  function setUp() public override {
    super.setUp();

    (, s_messageId) = _createMessageWithId();
    s_receipt = _encodeMessageSentReceipt(s_sourceOnRamp, s_messageId);
  }

  function _createMessage() internal pure returns (MessageV1Codec.MessageV1 memory message) {
    (message,) = _createMessageWithId();
    return message;
  }

  function test_verifyMessage_HeaderChain() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);

    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_MessageBlockIsProven() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, 1);

    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_ProvenBlock() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);
    // Cache the message block hash so verification only needs that block's header.
    bytes[] memory provingHeaders = new bytes[](HEADER_COUNT - 1);
    for (uint256 i = 0; i < provingHeaders.length; ++i) {
      provingHeaders[i] = witness.headers[i];
    }
    s_zkVerifier.proveBlockHash(SOURCE_CHAIN_SELECTOR, PROVEN_BLOCK_NUMBER, provingHeaders);

    bytes[] memory headers = new bytes[](1);
    headers[0] = witness.headers[HEADER_COUNT - 1];
    witness.provenBlockNumber = PROVEN_BLOCK_NUMBER - provingHeaders.length;
    witness.headers = headers;

    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_LegacyReceipt() public {
    bytes[] memory logs = new bytes[](1);
    logs[0] = _encodeLog(s_sourceOnRamp, _encodeMessageSentTopics(s_messageId));
    SuccinctZKVerifier.Witness memory witness = _buildWitness(_encodeReceipt(true, logs), HEADER_COUNT);

    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_LastOfManyLogs() public {
    bytes[] memory logs = new bytes[](100);
    for (uint256 i = 0; i < logs.length - 1; ++i) {
      logs[i] = _encodeLog(makeAddr("otherEmitter"), _encodeMessageSentTopics(keccak256("otherMessageId")));
    }
    logs[logs.length - 1] = _encodeLog(s_sourceOnRamp, _encodeMessageSentTopics(s_messageId));
    SuccinctZKVerifier.Witness memory witness = _buildWitness(_encodeReceipt(true, logs), HEADER_COUNT);
    witness.logIndex = logs.length - 1;

    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_SkipsItemsOfEveryRLPSizeClass() public {
    // Use synthetic list items to cover RLP prefixes that do not occur in receipt logs.
    bytes[] memory logs = new bytes[](5);
    logs[0] = hex"01"; // single byte
    logs[1] = RLPWriter.writeBytes(hex"0102"); // short string
    logs[2] = RLPWriter.writeBytes(new bytes(64)); // long string
    logs[3] = RLPWriter.writeList(new bytes[](0)); // short list
    logs[4] = _encodeLog(s_sourceOnRamp, _encodeMessageSentTopics(s_messageId)); // long list
    SuccinctZKVerifier.Witness memory witness = _buildWitness(_encodeReceipt(true, logs), HEADER_COUNT);
    witness.logIndex = 4;

    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_OnRampEvent() public {
    // The event is emitted by this contract, so the message must name it as the OnRamp.
    MessageV1Codec.MessageV1 memory message = _createMessage();
    message.onRampAddress = abi.encode(address(this));
    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));

    vm.recordLogs();
    emit OnRamp.CCIPMessageSent({
      destChainSelector: message.destChainSelector,
      sender: abi.decode(message.sender, (address)),
      messageId: messageId,
      feeToken: address(0),
      tokenAmountBeforeTokenPoolFees: 0,
      encodedMessage: MessageV1Codec._encodeMessageV1(message),
      receipts: new OnRamp.Receipt[](0),
      verifierBlobs: new bytes[](0)
    });
    Vm.Log memory recordedLog = vm.getRecordedLogs()[0];

    bytes[] memory topics = new bytes[](recordedLog.topics.length);
    for (uint256 i = 0; i < topics.length; ++i) {
      topics[i] = RLPWriter.writeBytes(abi.encodePacked(recordedLog.topics[i]));
    }
    bytes[] memory fields = new bytes[](3);
    fields[0] = RLPWriter.writeBytes(abi.encodePacked(recordedLog.emitter));
    fields[1] = RLPWriter.writeList(topics);
    fields[2] = RLPWriter.writeBytes(recordedLog.data);
    bytes[] memory logs = new bytes[](1);
    logs[0] = RLPWriter.writeList(fields);
    SuccinctZKVerifier.Witness memory witness = _buildWitness(_encodeReceipt(true, logs), HEADER_COUNT);

    s_zkVerifier.verifyMessage(message, messageId, _encodeVerifierResults(witness));
  }

  // Reverts

  function test_verifyMessage_RevertWhen_CursedByRMN() public {
    bytes memory verifierResults = _encodeVerifierResults(_buildWitness(s_receipt, HEADER_COUNT));
    _setMockRMNChainCurse(SOURCE_CHAIN_SELECTOR, true);

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CursedByRMN.selector, SOURCE_CHAIN_SELECTOR));
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, verifierResults);
  }

  function test_verifyMessage_RevertWhen_SourceChainNotSupported() public {
    bytes memory verifierResults = _encodeVerifierResults(_buildWitness(s_receipt, HEADER_COUNT));
    _setSourceChainConfig(MockSP1Helios(address(0)));

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.SourceChainNotSupported.selector, SOURCE_CHAIN_SELECTOR));
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, verifierResults);
  }

  function test_verifyMessage_RevertWhen_InvalidVerifierResults() public {
    vm.expectRevert(SuccinctZKVerifier.InvalidVerifierResults.selector);
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, hex"0102");
  }

  function test_verifyMessage_RevertWhen_InvalidCCVVersion() public {
    bytes memory verifierResults = _encodeVerifierResults(_buildWitness(s_receipt, HEADER_COUNT));
    verifierResults[0] ^= 0xff;

    vm.expectRevert(
      abi.encodeWithSelector(SuccinctZKVerifier.InvalidCCVVersion.selector, VERSION_TAG_V0_0_1, bytes4(verifierResults))
    );
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, verifierResults);
  }

  function test_verifyMessage_RevertWhen_InvalidVkey_LightClient() public {
    bytes memory verifierResults = _encodeVerifierResults(_buildWitness(s_receipt, HEADER_COUNT));
    bytes32 otherVkey = keccak256("otherVkey");
    s_mockHelios.setVkeys(otherVkey, EXECUTION_HEADER_VKEY);

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidVkey.selector, LIGHT_CLIENT_VKEY, otherVkey));
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, verifierResults);
  }

  function test_verifyMessage_RevertWhen_InvalidVkey_ExecutionHeader() public {
    bytes memory verifierResults = _encodeVerifierResults(_buildWitness(s_receipt, HEADER_COUNT));
    bytes32 otherVkey = keccak256("otherVkey");
    s_mockHelios.setVkeys(LIGHT_CLIENT_VKEY, otherVkey);

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidVkey.selector, EXECUTION_HEADER_VKEY, otherVkey));
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, verifierResults);
  }

  function test_verifyMessage_RevertWhen_BlockNotProven() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);
    witness.provenBlockNumber = PROVEN_BLOCK_NUMBER + 1;

    vm.expectRevert(
      abi.encodeWithSelector(SuccinctZKVerifier.BlockNotProven.selector, SOURCE_CHAIN_SELECTOR, PROVEN_BLOCK_NUMBER + 1)
    );
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_EmptyHeaderChain() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);
    witness.headers = new bytes[](0);

    vm.expectRevert(SuccinctZKVerifier.EmptyHeaderChain.selector);
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidHeaderHash_FirstHeader() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);
    bytes32 provenBlockHash = keccak256(witness.headers[0]);
    witness.headers[0][40] ^= 0x01;

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.InvalidHeaderHash.selector, 0, provenBlockHash, keccak256(witness.headers[0])
      )
    );
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
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
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidRootNode() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);
    witness.proofNodes[0][40] ^= 0x01;

    vm.expectRevert("MerkleTrie: invalid root hash");
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidLeafNode() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);
    witness.proofNodes[1][40] ^= 0x01;

    vm.expectRevert("MerkleTrie: invalid large internal hash");
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidTxIndex() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);
    witness.txIndex = 1;

    vm.expectRevert("MerkleTrie: invalid large internal hash");
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_ReceiptNotSuccessful() public {
    bytes[] memory logs = new bytes[](1);
    logs[0] = _encodeLog(s_sourceOnRamp, _encodeMessageSentTopics(s_messageId));
    SuccinctZKVerifier.Witness memory witness = _buildWitness(_encodeReceipt(false, logs), HEADER_COUNT);

    vm.expectRevert(SuccinctZKVerifier.ReceiptNotSuccessful.selector);
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_LogIndexOutOfRange() public {
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);
    witness.logIndex = 1;

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.LogIndexOutOfRange.selector, 1, 1));
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
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
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidLogEmitter_MessageNamesOtherOnRamp() public {
    bytes memory verifierResults = _encodeVerifierResults(_buildWitness(s_receipt, HEADER_COUNT));
    address otherOnRamp = makeAddr("otherOnRamp");
    MessageV1Codec.MessageV1 memory message = _createMessage();
    message.onRampAddress = abi.encode(otherOnRamp);

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.InvalidLogEmitter.selector, otherOnRamp, abi.encodePacked(s_sourceOnRamp)
      )
    );
    s_zkVerifier.verifyMessage(message, s_messageId, verifierResults);
  }

  function test_verifyMessage_RevertWhen_InvalidTopicCount() public {
    bytes[] memory topics = new bytes[](3);
    topics[0] = RLPWriter.writeBytes(abi.encodePacked(CCIP_MESSAGE_SENT_TOPIC));
    topics[1] = RLPWriter.writeBytes(abi.encodePacked(bytes32(uint256(DEST_CHAIN_SELECTOR))));
    topics[2] = RLPWriter.writeBytes(abi.encodePacked(s_messageId));
    SuccinctZKVerifier.Witness memory witness = _buildWitness(_encodeReceiptWithTopics(topics), HEADER_COUNT);

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidTopicCount.selector, 3));
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidLogTopic() public {
    bytes32 otherTopic = keccak256("OtherEvent()");
    bytes[] memory topics = _encodeMessageSentTopics(s_messageId);
    topics[0] = RLPWriter.writeBytes(abi.encodePacked(otherTopic));
    SuccinctZKVerifier.Witness memory witness = _buildWitness(_encodeReceiptWithTopics(topics), HEADER_COUNT);

    vm.expectRevert(
      abi.encodeWithSelector(SuccinctZKVerifier.InvalidLogTopic.selector, CCIP_MESSAGE_SENT_TOPIC, otherTopic)
    );
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidMessageId() public {
    bytes32 otherMessageId = keccak256("otherMessageId");
    SuccinctZKVerifier.Witness memory witness = _buildWitness(s_receipt, HEADER_COUNT);

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidMessageId.selector, otherMessageId, s_messageId));
    s_zkVerifier.verifyMessage(_createMessage(), otherMessageId, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidFieldLength() public {
    bytes[] memory topics = _encodeMessageSentTopics(s_messageId);
    topics[0] = RLPWriter.writeBytes(hex"0102");
    SuccinctZKVerifier.Witness memory witness = _buildWitness(_encodeReceiptWithTopics(topics), HEADER_COUNT);

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidFieldLength.selector, 32, 2));
    s_zkVerifier.verifyMessage(_createMessage(), s_messageId, _encodeVerifierResults(witness));
  }

  function _encodeReceiptWithTopics(
    bytes[] memory encodedTopics
  ) internal view returns (bytes memory) {
    bytes[] memory logs = new bytes[](1);
    logs[0] = _encodeLog(s_sourceOnRamp, encodedTopics);
    return bytes.concat(EIP1559_TRANSACTION_TYPE, _encodeReceipt(true, logs));
  }
}
