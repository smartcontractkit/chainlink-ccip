// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SignatureQuorumVerifier} from "../../../offRamp/components/SignatureQuorumVerifier.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {SignatureQuorumVerifierHelper} from "../../helpers/SignatureQuorumVerifierHelper.sol";

contract SignatureVerifierSetup is BaseTest {
  // Signer private keys used for these test
  uint256 internal constant PRIVATE0 = 0x7b2e97fe057e6de99d6872a2ef2abf52c9b4469bc848c2465ac3fcd8d336e81d;

  address[] internal s_validSigners;
  uint256[] internal s_validSignerKeys;

  bytes32 internal constant DEFAULT_CONFIG_DIGEST = keccak256(abi.encode("defaultConfigDigest"));

  bytes internal constant REPORT = abi.encode("testReport");
  SignatureQuorumVerifierHelper internal s_sigQuorumVerifier;

  function setUp() public virtual override {
    BaseTest.setUp();

    uint160 numSigners = 4;
    s_validSignerKeys = new uint256[](numSigners);
    s_validSigners = new address[](numSigners);

    for (uint160 i; i < numSigners; ++i) {
      s_validSignerKeys[i] = PRIVATE0 + i;
      s_validSigners[i] = vm.addr(s_validSignerKeys[i]);
    }

    s_sigQuorumVerifier = new SignatureQuorumVerifierHelper();

    s_sigQuorumVerifier.setSignatureConfig(
      SignatureQuorumVerifier.SignatureConfigArgs({signers: s_validSigners, threshold: 1})
    );
  }

  function _getSignatures(
    uint256[] memory signerKeys,
    bytes32 reportHash
  ) internal pure returns (bytes32[] memory rs, bytes32[] memory ss) {
    uint256 signatureCount = signerKeys.length;
    rs = new bytes32[](signatureCount);
    ss = new bytes32[](signatureCount);

    // Calculate signatures
    for (uint256 i; i < signatureCount; ++i) {
      (, rs[i], ss[i]) = vm.sign(signerKeys[i], reportHash);
    }

    return (rs, ss);
  }
}
