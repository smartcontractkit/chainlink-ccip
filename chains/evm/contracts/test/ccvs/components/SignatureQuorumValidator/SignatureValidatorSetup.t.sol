// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseTest} from "../../../BaseTest.t.sol";
import {SignatureQuorumValidatorHelper} from "../../../helpers/SignatureQuorumValidatorHelper.sol";

contract SignatureValidatorSetup is BaseTest {
  // 4 hardcoded private keys that are chosen to work with v=27 ecrecover
  // and have addresses in ascending order after sorting.
  uint256 internal constant PRIVATE_KEY_0 = 0x60b919c82f0b4791a5b7c6a7275970ace1748759ebdaa8a5c3a4b2f5a8b1e8d1;
  uint256 internal constant PRIVATE_KEY_1 = 0x70b919c82f0b4791a5b7c6a7275970ace1748759ebdaa8a5c3a4b2f5a8b1e8d2;
  uint256 internal constant PRIVATE_KEY_2 = 0x80b919c82f0b4791a5b7c6a7275970ace1748759ebdaa8a5c3a4b2f5a8b1e8d3;
  uint256 internal constant PRIVATE_KEY_3 = 0x90b919c82f0b4791a5b7c6a7275970ace1748759ebdaa8a5c3a4b2f5a8b1e8d4;

  address[] internal s_validSigners;
  uint256[] internal s_validSignerKeys;

  bytes32 internal constant DEFAULT_CONFIG_DIGEST = keccak256(abi.encode("defaultConfigDigest"));

  bytes internal constant REPORT = abi.encode("testReport");
  SignatureQuorumValidatorHelper internal s_sigQuorumVerifier;

  function setUp() public virtual override {
    BaseTest.setUp();

    s_validSignerKeys = new uint256[](4);
    s_validSigners = new address[](4);

    // Set up the keys and addresses.
    s_validSignerKeys[0] = PRIVATE_KEY_0;
    s_validSignerKeys[1] = PRIVATE_KEY_1;
    s_validSignerKeys[2] = PRIVATE_KEY_2;
    s_validSignerKeys[3] = PRIVATE_KEY_3;

    s_validSigners[0] = vm.addr(PRIVATE_KEY_0);
    s_validSigners[1] = vm.addr(PRIVATE_KEY_1);
    s_validSigners[2] = vm.addr(PRIVATE_KEY_2);
    s_validSigners[3] = vm.addr(PRIVATE_KEY_3);

    // Sort signers and keys by address to ensure proper ordering.
    _sortSignersByAddress();

    s_sigQuorumVerifier = new SignatureQuorumValidatorHelper();
    s_sigQuorumVerifier.setSignatureConfig(s_validSigners, 1);
  }

  function _sortSignersByAddress() internal {
    // Simple bubble sort by address.
    for (uint256 i = 0; i < s_validSigners.length; ++i) {
      for (uint256 j = i + 1; j < s_validSigners.length; ++j) {
        if (s_validSigners[i] > s_validSigners[j]) {
          // Swap addresses.
          address tempAddr = s_validSigners[i];
          s_validSigners[i] = s_validSigners[j];
          s_validSigners[j] = tempAddr;

          // Swap corresponding keys.
          uint256 tempKey = s_validSignerKeys[i];
          s_validSignerKeys[i] = s_validSignerKeys[j];
          s_validSignerKeys[j] = tempKey;
        }
      }
    }
  }

  /// @notice Helper to create a signature with v=27 (required by SignatureQuorumValidator)
  /// @param privateKey The private key to sign with
  /// @param hash The hash to sign
  /// @return r The r component of the signature
  /// @return s The s component of the signature (adjusted for v=27)
  function _signWithV27(uint256 privateKey, bytes32 hash) internal pure returns (bytes32 r, bytes32 s) {
    (uint8 v, bytes32 _r, bytes32 _s) = vm.sign(privateKey, hash);

    // SignatureQuorumValidator only supports sigs with v=27, so adjust if necessary.
    // Any valid ECDSA sig (r, s, v) can be "flipped" into (r, s*, v*) without knowing the private key.
    // https://github.com/kadenzipfel/smart-contract-vulnerabilities/blob/master/vulnerabilities/signature-malleability.md
    if (v == 28) {
      uint256 N = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141;
      _s = bytes32(N - uint256(_s));
    }

    return (_r, _s);
  }

  function _getSignatures(
    uint256[] memory signerKeys,
    bytes32 reportHash
  ) internal pure returns (bytes32[] memory rs, bytes32[] memory ss) {
    uint256 signatureCount = signerKeys.length;
    rs = new bytes32[](signatureCount);
    ss = new bytes32[](signatureCount);

    for (uint256 i; i < signatureCount; ++i) {
      (rs[i], ss[i]) = _signWithV27(signerKeys[i], reportHash);
    }

    return (rs, ss);
  }
}
