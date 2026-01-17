// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {CommitOffRamp} from "../../../offRamp/CommitOffRamp.sol";
import {BaseTest} from "../../BaseTest.t.sol";

contract CommitOffRampSetup is BaseTest {
  CommitOffRamp internal s_commitOffRamp;

  uint256 internal constant PRIVATE_KEY_0 = 0x60b919c82f0b4791a5b7c6a7275970ace1748759ebdaa8a5c3a4b2f5a8b1e8d1;

  function setUp() public virtual override {
    BaseTest.setUp();

    s_commitOffRamp = new CommitOffRamp();

    address[] memory validSigner = new address[](1);
    validSigner[0] = vm.addr(PRIVATE_KEY_0);

    s_commitOffRamp.setSignatureConfig(validSigner, 1);
  }

  /// @notice Helper to create a signature with v=27 (required by SignatureQuorumVerifier).
  /// @param privateKey The private key to sign with.
  /// @param hash The hash to sign.
  /// @return r The r component of the signature.
  /// @return s The s component of the signature (adjusted for v=27).
  function _signWithV27(uint256 privateKey, bytes32 hash) internal pure returns (bytes32 r, bytes32 s) {
    (uint8 v, bytes32 _r, bytes32 _s) = vm.sign(privateKey, hash);

    // SignatureQuorumVerifier only supports sigs with v=27, so adjust if necessary.
    // Any valid ECDSA sig (r, s, v) can be "flipped" into (r, s*, v*) without knowing the private key.
    // https://github.com/kadenzipfel/smart-contract-vulnerabilities/blob/master/vulnerabilities/signature-malleability.md
    if (v == 28) {
      uint256 N = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141;
      _s = bytes32(N - uint256(_s));
    }

    return (_r, _s);
  }

  function _generateBasicMessageV1() internal pure returns (MessageV1Codec.MessageV1 memory) {
    return MessageV1Codec.MessageV1({
      sourceChainSelector: 1,
      destChainSelector: 2,
      sequenceNumber: 1,
      onRampAddress: abi.encodePacked(address(0x1111111111111111111111111111111111111111)),
      offRampAddress: abi.encodePacked(address(0x2222222222222222222222222222222222222222)),
      finality: 100,
      sender: abi.encodePacked(address(0x3333333333333333333333333333333333333333)),
      receiver: abi.encodePacked(address(0x4444444444444444444444444444444444444444)),
      destBlob: "",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](0),
      data: ""
    });
  }

  function _generateMessageHash(
    MessageV1Codec.MessageV1 memory message
  ) internal pure returns (bytes32) {
    return keccak256(MessageV1Codec._encodeMessageV1(message));
  }
}
