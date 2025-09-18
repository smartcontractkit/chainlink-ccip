// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {CommitRamp} from "../../../ccvs/CommitRamp.sol";
import {CommitRampSetup} from "./CommitRampSetup.t.sol";

contract CommitRamp_verifyMessage is CommitRampSetup {
  function test_verifyMessage_ExtractsSignatureLengthCorrectly() public view {
    (MessageV1Codec.MessageV1 memory message,) = _createMessageV1(DEST_CHAIN_SELECTOR, MOCK_SENDER, "", MOCK_RECEIVER);
    bytes32 messageHash = _generateMessageHash(message);

    (bytes32 r, bytes32 s) = _signWithV27(PRIVATE_KEY_0, messageHash);
    bytes memory ccvData = abi.encodePacked(uint16(64), r, s);

    s_commitRamp.verifyMessage(message, messageHash, ccvData);
  }

  function test_verifyMessage_ForwardCompatibilityWithExtraData() public view {
    (MessageV1Codec.MessageV1 memory message,) = _createMessageV1(DEST_CHAIN_SELECTOR, MOCK_SENDER, "", MOCK_RECEIVER);
    bytes32 messageHash = _generateMessageHash(message);

    // Extra data can be added to the ccvData without it affecting the signature validation.
    bytes memory extraFutureData = hex"0123456789012345678901234567890123456789012345678901234567890123";

    (bytes32 r, bytes32 s) = _signWithV27(PRIVATE_KEY_0, messageHash);
    bytes memory ccvData = abi.encodePacked(uint16(64), r, s, extraFutureData);

    s_commitRamp.verifyMessage(message, messageHash, ccvData);
  }

  // Reverts

  function test_verifyMessage_RevertWhen_InvalidCCVData_InvalidSignatureLength() public {
    (MessageV1Codec.MessageV1 memory message,) = _createMessageV1(DEST_CHAIN_SELECTOR, MOCK_SENDER, "", MOCK_RECEIVER);
    bytes32 messageHash = _generateMessageHash(message);

    // Create ccvData with only 1 byte (missing signature length - needs 2 bytes for uint16).
    bytes memory ccvData = hex"01";

    // Should revert with InvalidCCVData when ccvData is too short to contain signature length.
    vm.expectRevert(abi.encodeWithSelector(CommitRamp.InvalidCCVData.selector));
    s_commitRamp.verifyMessage(message, messageHash, ccvData);
  }

  function test_verifyMessage_RevertWhen_InvalidCCVData_SignatureLengthExceedsCCVData() public {
    (MessageV1Codec.MessageV1 memory message,) = _createMessageV1(DEST_CHAIN_SELECTOR, MOCK_SENDER, "", MOCK_RECEIVER);
    bytes32 messageHash = _generateMessageHash(message);

    // Set signature length to 100 but only provide 2 bytes of actual signature data.
    bytes memory ccvData = abi.encodePacked(uint16(100), hex"ab");

    // Should revert with InvalidCCVData when signature length exceeds available data.
    vm.expectRevert(abi.encodeWithSelector(CommitRamp.InvalidCCVData.selector));
    s_commitRamp.verifyMessage(message, messageHash, ccvData);
  }
}
