// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitteeVerifier} from "../../../ccvs/CommitteeVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {CommitteeVerifierSetup} from "./CommitteeVerifierSetup.t.sol";

contract CommitteeVerifier_verifyMessage is CommitteeVerifierSetup {
  function test_verifyMessage_ExtractsSignatureLengthCorrectly() public view {
    (MessageV1Codec.MessageV1 memory message,) = _generateBasicMessageV1();
    bytes32 messageHash = _generateMessageHash(message);

    (bytes32 r, bytes32 s) = _signWithV27(PRIVATE_KEY_0, keccak256(bytes.concat(s_versionTag, messageHash)));
    bytes memory ccvData = abi.encodePacked(s_versionTag, uint16(64), r, s);

    s_committeeVerifier.verifyMessage(message, messageHash, ccvData);
  }

  function test_verifyMessage_ForwardCompatibilityWithExtraData() public view {
    (MessageV1Codec.MessageV1 memory message,) = _generateBasicMessageV1();
    bytes32 messageHash = _generateMessageHash(message);

    // Extra data can be added to the ccvData without it affecting the signature validation.
    bytes memory extraFutureData = hex"0123456789012345678901234567890123456789012345678901234567890123";

    (bytes32 r, bytes32 s) = _signWithV27(PRIVATE_KEY_0, keccak256(bytes.concat(s_versionTag, messageHash)));
    bytes memory ccvData = abi.encodePacked(s_versionTag, uint16(64), r, s, extraFutureData);

    s_committeeVerifier.verifyMessage(message, messageHash, ccvData);
  }

  // Reverts

  function test_verifyMessage_RevertWhen_InvalidCCVData_InvalidPrefix() public {
    (MessageV1Codec.MessageV1 memory message,) = _generateBasicMessageV1();
    bytes32 messageHash = _generateMessageHash(message);

    // Create ccvData with only 1 byte (missing signature length & verifier version).
    // Reverts because we need 2 bytes for uint16 & 4 bytes for verifier version.
    bytes memory ccvData = hex"0102030405";

    // Should revert with InvalidCCVData when ccvData is too short to contain signature length & verifier version.
    vm.expectRevert(abi.encodeWithSelector(CommitteeVerifier.InvalidCCVData.selector));
    s_committeeVerifier.verifyMessage(message, messageHash, ccvData);
  }

  function test_verifyMessage_RevertWhen_InvalidCCVData_SignatureLengthExceedsCCVData() public {
    (MessageV1Codec.MessageV1 memory message,) = _generateBasicMessageV1();
    bytes32 messageHash = _generateMessageHash(message);

    // Set signature length to 100 but only provide 2 bytes of actual signature data.
    bytes memory ccvData = abi.encodePacked(s_versionTag, uint16(100), hex"ab");

    // Should revert with InvalidCCVData when signature length exceeds available data.
    vm.expectRevert(abi.encodeWithSelector(CommitteeVerifier.InvalidCCVData.selector));
    s_committeeVerifier.verifyMessage(message, messageHash, ccvData);
  }

  function test_verifyMessage_RevertWhen_InvalidVersion() public {
    (MessageV1Codec.MessageV1 memory message,) = _generateBasicMessageV1();
    bytes32 messageHash = _generateMessageHash(message);

    // Has expected length, but the version (0x01020304) is incorrect.
    bytes memory ccvData = hex"010203040506";

    // Should revert with InvalidCCVVersion when the version is incorrect.
    vm.expectRevert(abi.encodeWithSelector(CommitteeVerifier.InvalidCCVVersion.selector, bytes4(0x01020304)));
    s_committeeVerifier.verifyMessage(message, messageHash, ccvData);
  }
}
