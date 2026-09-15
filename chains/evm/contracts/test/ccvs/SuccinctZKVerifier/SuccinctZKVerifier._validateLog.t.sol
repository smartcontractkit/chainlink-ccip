// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SuccinctZKVerifier} from "../../../ccvs/SuccinctZKVerifier.sol";
import {SuccinctZKVerifierHelper} from "../../helpers/SuccinctZKVerifierHelper.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";

import {RLPWriter} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPWriter.sol";

contract SuccinctZKVerifier_validateLog is SuccinctZKVerifierSetup {
  bytes32 internal constant CCIP_MESSAGE_SENT_TOPIC =
    0x371bc2ff0a006f4ef863b1d27a065d4e9f938b6d883eb154572b4aea593b32cc;
  bytes1 internal constant EIP1559_TRANSACTION_TYPE = 0x02;

  SuccinctZKVerifierHelper internal s_helper;
  address internal s_onRampEmitter = makeAddr("onRampEmitter");
  bytes32 internal s_messageId = keccak256("messageId");

  function setUp() public override {
    super.setUp();

    s_helper = new SuccinctZKVerifierHelper(
      SuccinctZKVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR}),
      s_storageLocations,
      address(s_mockRMNRemote),
      VERSION_TAG_V0_0_1
    );
  }

  function test_validateLog_LegacyReceipt() public view {
    bytes memory receipt = _encodeReceipt(true, _encodeMessageSentLog(s_onRampEmitter, s_messageId));

    s_helper.validateLog(receipt, 0, s_onRampEmitter, s_messageId);
  }

  function test_validateLog_TypedReceipt() public view {
    bytes memory receipt =
      bytes.concat(EIP1559_TRANSACTION_TYPE, _encodeReceipt(true, _encodeMessageSentLog(s_onRampEmitter, s_messageId)));

    s_helper.validateLog(receipt, 0, s_onRampEmitter, s_messageId);
  }

  function test_validateLog_SecondLog() public {
    bytes[] memory logs = new bytes[](2);
    logs[0] = _encodeMessageSentLog(makeAddr("other"), keccak256("otherMessageId"));
    logs[1] = _encodeMessageSentLog(s_onRampEmitter, s_messageId);
    bytes memory receipt = _encodeReceipt(true, logs);

    s_helper.validateLog(receipt, 1, s_onRampEmitter, s_messageId);
  }

  function test_readBytes32() public view {
    bytes32 value = keccak256("value");

    assertEq(value, s_helper.readBytes32(RLPWriter.writeBytes(abi.encodePacked(value))));
  }

  // Reverts

  function test_validateLog_RevertWhen_ReceiptNotSuccessful() public {
    bytes memory receipt = _encodeReceipt(false, _encodeMessageSentLog(s_onRampEmitter, s_messageId));

    vm.expectRevert(SuccinctZKVerifier.ReceiptNotSuccessful.selector);
    s_helper.validateLog(receipt, 0, s_onRampEmitter, s_messageId);
  }

  function test_validateLog_RevertWhen_LogIndexOutOfRange() public {
    bytes memory receipt = _encodeReceipt(true, _encodeMessageSentLog(s_onRampEmitter, s_messageId));

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.LogIndexOutOfRange.selector, 1, 1));
    s_helper.validateLog(receipt, 1, s_onRampEmitter, s_messageId);
  }

  function test_validateLog_RevertWhen_InvalidLogEmitter() public {
    address otherEmitter = makeAddr("otherEmitter");
    bytes memory receipt = _encodeReceipt(true, _encodeMessageSentLog(otherEmitter, s_messageId));

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.InvalidLogEmitter.selector, s_onRampEmitter, abi.encodePacked(otherEmitter)
      )
    );
    s_helper.validateLog(receipt, 0, s_onRampEmitter, s_messageId);
  }

  function test_validateLog_RevertWhen_InvalidTopicCount() public {
    bytes[] memory topics = new bytes[](3);
    topics[0] = RLPWriter.writeBytes(abi.encodePacked(CCIP_MESSAGE_SENT_TOPIC));
    topics[1] = RLPWriter.writeBytes(abi.encodePacked(bytes32(uint256(DEST_CHAIN_SELECTOR))));
    topics[2] = RLPWriter.writeBytes(abi.encodePacked(s_messageId));
    bytes memory receipt = _encodeReceipt(true, _encodeLog(s_onRampEmitter, topics));

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidTopicCount.selector, 3));
    s_helper.validateLog(receipt, 0, s_onRampEmitter, s_messageId);
  }

  function test_validateLog_RevertWhen_InvalidLogTopic() public {
    bytes32 otherTopic = keccak256("OtherEvent(uint64,bytes,bytes32,bytes)");
    bytes[] memory topics = _messageSentTopics(s_messageId);
    topics[0] = RLPWriter.writeBytes(abi.encodePacked(otherTopic));
    bytes memory receipt = _encodeReceipt(true, _encodeLog(s_onRampEmitter, topics));

    vm.expectRevert(
      abi.encodeWithSelector(SuccinctZKVerifier.InvalidLogTopic.selector, CCIP_MESSAGE_SENT_TOPIC, otherTopic)
    );
    s_helper.validateLog(receipt, 0, s_onRampEmitter, s_messageId);
  }

  function test_validateLog_RevertWhen_InvalidMessageId() public {
    bytes32 otherMessageId = keccak256("otherMessageId");
    bytes memory receipt = _encodeReceipt(true, _encodeMessageSentLog(s_onRampEmitter, otherMessageId));

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidMessageId.selector, s_messageId, otherMessageId));
    s_helper.validateLog(receipt, 0, s_onRampEmitter, s_messageId);
  }

  function test_validateLog_RevertWhen_InvalidFieldLength() public {
    bytes[] memory topics = _messageSentTopics(s_messageId);
    topics[0] = RLPWriter.writeBytes(hex"0102");
    bytes memory receipt = _encodeReceipt(true, _encodeLog(s_onRampEmitter, topics));

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidFieldLength.selector, 32, 2));
    s_helper.validateLog(receipt, 0, s_onRampEmitter, s_messageId);
  }

  function test_readBytes32_RevertWhen_InvalidFieldLength() public {
    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidFieldLength.selector, 32, 20));
    s_helper.readBytes32(RLPWriter.writeBytes(abi.encodePacked(s_onRampEmitter)));
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

  function _encodeReceipt(
    bool success,
    bytes memory encodedLog
  ) internal pure returns (bytes memory) {
    bytes[] memory encodedLogs = new bytes[](1);
    encodedLogs[0] = encodedLog;
    return _encodeReceipt(success, encodedLogs);
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

  function _encodeMessageSentLog(
    address emitter,
    bytes32 messageId
  ) internal pure returns (bytes memory) {
    return _encodeLog(emitter, _messageSentTopics(messageId));
  }

  /// @notice Topics of CCIPMessageSent(uint64 indexed destChainSelector, bytes indexed sender, bytes32 indexed messageId).
  function _messageSentTopics(
    bytes32 messageId
  ) internal pure returns (bytes[] memory) {
    bytes[] memory topics = new bytes[](4);
    topics[0] = RLPWriter.writeBytes(abi.encodePacked(CCIP_MESSAGE_SENT_TOPIC));
    topics[1] = RLPWriter.writeBytes(abi.encodePacked(bytes32(uint256(DEST_CHAIN_SELECTOR))));
    topics[2] = RLPWriter.writeBytes(abi.encodePacked(keccak256(abi.encode(OWNER))));
    topics[3] = RLPWriter.writeBytes(abi.encodePacked(messageId));
    return topics;
  }
}
