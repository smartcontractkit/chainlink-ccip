// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {Test} from "forge-std/Test.sol";

contract MessageV1CodecSetup is Test {
  // Helper contract to make the args calldata for decode functions.
  MessageV1CodecHelper internal s_helper;

  function setUp() public virtual {
    s_helper = new MessageV1CodecHelper();
  }

  // Helper functions used across multiple test files.
  function _createBasicMessage() internal pure returns (MessageV1Codec.MessageV1 memory) {
    return MessageV1Codec.MessageV1({
      sourceChainSelector: 1,
      destChainSelector: 2,
      sequenceNumber: 100,
      onRampAddress: abi.encodePacked(address(0x1234567890123456789012345678901234567890)),
      offRampAddress: abi.encodePacked(address(0x0987654321098765432109876543210987654321)),
      finality: 1000,
      sender: abi.encodePacked(address(0x1111111111111111111111111111111111111111)),
      receiver: abi.encodePacked(address(0x2222222222222222222222222222222222222222)),
      destBlob: "test blob",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](0),
      data: "test data"
    });
  }

  function _createBasicTokenTransfer() internal pure returns (MessageV1Codec.TokenTransferV1 memory) {
    return MessageV1Codec.TokenTransferV1({
      amount: 100,
      sourcePoolAddress: hex"1234567890abcdef",
      sourceTokenAddress: hex"abcdef1234567890",
      destTokenAddress: hex"fedcba0987654321",
      extraData: hex"deadbeef"
    });
  }

  // Assertion helpers for complex types used in multiple tests.
  function _assertMessageEqual(
    MessageV1Codec.MessageV1 memory expected,
    MessageV1Codec.MessageV1 memory actual
  ) internal pure {
    assertEq(expected.sourceChainSelector, actual.sourceChainSelector, "sourceChainSelector mismatch");
    assertEq(expected.destChainSelector, actual.destChainSelector, "destChainSelector mismatch");
    assertEq(expected.sequenceNumber, actual.sequenceNumber, "sequenceNumber mismatch");
    assertEq(expected.onRampAddress, actual.onRampAddress, "onRampAddress mismatch");
    assertEq(expected.offRampAddress, actual.offRampAddress, "offRampAddress mismatch");
    assertEq(expected.finality, actual.finality, "finality mismatch");
    assertEq(expected.sender, actual.sender, "sender mismatch");
    assertEq(expected.receiver, actual.receiver, "receiver mismatch");
    assertEq(expected.destBlob, actual.destBlob, "destBlob mismatch");
    assertEq(expected.data, actual.data, "data mismatch");

    assertEq(expected.tokenTransfer.length, actual.tokenTransfer.length, "tokenTransfer length mismatch");
    for (uint256 i = 0; i < expected.tokenTransfer.length; ++i) {
      _assertTokenTransferEqual(expected.tokenTransfer[i], actual.tokenTransfer[i], i);
    }
  }

  function _assertTokenTransferEqual(
    MessageV1Codec.TokenTransferV1 memory expected,
    MessageV1Codec.TokenTransferV1 memory actual,
    uint256 index
  ) internal pure {
    string memory indexStr = vm.toString(index);
    assertEq(expected.amount, actual.amount, string(abi.encodePacked("tokenTransfer[", indexStr, "].amount mismatch")));
    assertEq(
      expected.sourcePoolAddress,
      actual.sourcePoolAddress,
      string(abi.encodePacked("tokenTransfer[", indexStr, "].sourcePoolAddress mismatch"))
    );
    assertEq(
      expected.sourceTokenAddress,
      actual.sourceTokenAddress,
      string(abi.encodePacked("tokenTransfer[", indexStr, "].sourceTokenAddress mismatch"))
    );
    assertEq(
      expected.destTokenAddress,
      actual.destTokenAddress,
      string(abi.encodePacked("tokenTransfer[", indexStr, "].destTokenAddress mismatch"))
    );
    assertEq(
      expected.extraData, actual.extraData, string(abi.encodePacked("tokenTransfer[", indexStr, "].extraData mismatch"))
    );
  }
}

// Helper contract to make the args calldata for decode functions.
contract MessageV1CodecHelper {
  function decodeMessageV1(
    bytes calldata encoded
  ) external pure returns (MessageV1Codec.MessageV1 memory) {
    return MessageV1Codec._decodeMessageV1(encoded);
  }

  function decodeTokenTransferV1(
    bytes calldata encoded
  ) external pure returns (MessageV1Codec.TokenTransferV1 memory) {
    (MessageV1Codec.TokenTransferV1 memory tokenTransfer,) = MessageV1Codec._decodeTokenTransferV1(encoded, 0);
    return tokenTransfer;
  }

  function encodeMessageV1(
    MessageV1Codec.MessageV1 memory message
  ) external pure returns (bytes memory) {
    return MessageV1Codec._encodeMessageV1(message);
  }

  function encodeTokenTransferV1(
    MessageV1Codec.TokenTransferV1 memory tokenTransfer
  ) external pure returns (bytes memory) {
    return MessageV1Codec._encodeTokenTransferV1(tokenTransfer);
  }
}
