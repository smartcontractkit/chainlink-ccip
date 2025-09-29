// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SignatureQuorumValidator} from "../../../../ccvs/components/SignatureQuorumValidator.sol";
import {SignatureQuorumValidatorHelper} from "../../../helpers/SignatureQuorumValidatorHelper.sol";
import {SignatureValidatorSetup} from "./SignatureValidatorSetup.t.sol";

contract SignatureQuorumValidator_validateSignatures is SignatureValidatorSetup {
  bytes32 internal constant TEST_HASH = keccak256("test message");

  function _createSignatures(uint256[] memory signerKeys, bytes32 hash) internal pure returns (bytes memory) {
    bytes memory signatures = "";

    for (uint256 i = 0; i < signerKeys.length; ++i) {
      (bytes32 r, bytes32 s) = _signWithV27(signerKeys[i], hash);
      signatures = bytes.concat(signatures, r, s);
    }

    return signatures;
  }

  function test_validateSignatures_MultipleSignatures() public {
    s_sigQuorumVerifier.setSignatureConfig(s_validSigners, 3);

    uint256[] memory signerKeys = new uint256[](3);
    signerKeys[0] = s_validSignerKeys[0];
    signerKeys[1] = s_validSignerKeys[1];
    signerKeys[2] = s_validSignerKeys[2];

    bytes memory signatures = _createSignatures(signerKeys, TEST_HASH);

    s_sigQuorumVerifier.validateSignatures(TEST_HASH, signatures);
  }

  function test_validateSignatures_ExtraSignatures() public {
    // Set threshold to 2 but provide 3 signatures.
    s_sigQuorumVerifier.setSignatureConfig(s_validSigners, 2);

    uint256[] memory signerKeys = new uint256[](3);
    signerKeys[0] = s_validSignerKeys[0];
    signerKeys[1] = s_validSignerKeys[1];
    signerKeys[2] = s_validSignerKeys[2];

    bytes memory signatures = _createSignatures(signerKeys, TEST_HASH);

    // Should not revert - extra signatures are allowed.
    s_sigQuorumVerifier.validateSignatures(TEST_HASH, signatures);
  }

  // Reverts

  function test_validateSignatures_RevertWhen_NoSignersConfigured() public {
    // Deploy new verifier with no signers.
    SignatureQuorumValidatorHelper newVerifier = new SignatureQuorumValidatorHelper();

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.InvalidSignatureConfig.selector));
    newVerifier.validateSignatures(TEST_HASH, "");
  }

  function test_validateSignatures_RevertWhen_ForkedChain() public {
    s_sigQuorumVerifier.setSignatureConfig(s_validSigners, 1);

    uint256[] memory signerKeys = new uint256[](1);
    signerKeys[0] = s_validSignerKeys[0];
    bytes memory signatures = _createSignatures(signerKeys, TEST_HASH);

    // Change chain ID to simulate fork.
    uint256 originalChainId = block.chainid;
    uint256 newChainId = originalChainId + 100000;
    vm.chainId(newChainId);

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.ForkedChain.selector, originalChainId, newChainId));
    s_sigQuorumVerifier.validateSignatures(TEST_HASH, signatures);
  }

  function test_validateSignatures_RevertWhen_WrongNumberOfSignatures() public {
    s_sigQuorumVerifier.setSignatureConfig(s_validSigners, 3);

    // Provide only 2 signatures when 3 are required.
    uint256[] memory signerKeys = new uint256[](2);
    signerKeys[0] = s_validSignerKeys[0];
    signerKeys[1] = s_validSignerKeys[1];

    bytes memory signatures = _createSignatures(signerKeys, TEST_HASH);

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.WrongNumberOfSignatures.selector));
    s_sigQuorumVerifier.validateSignatures(TEST_HASH, signatures);
  }

  function test_validateSignatures_RevertWhen_UnauthorizedSigner() public {
    s_sigQuorumVerifier.setSignatureConfig(s_validSigners, 1);

    // Use a signer key not in the valid set.
    uint256 unauthorizedKey = 0x1234567890abcdef;
    uint256[] memory signerKeys = new uint256[](1);
    signerKeys[0] = unauthorizedKey;

    bytes memory signatures = _createSignatures(signerKeys, TEST_HASH);

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.UnauthorizedSigner.selector));
    s_sigQuorumVerifier.validateSignatures(TEST_HASH, signatures);
  }

  function test_validateSignatures_RevertWhen_UnorderedSignatures() public {
    s_sigQuorumVerifier.setSignatureConfig(s_validSigners, 2);

    // Get two signers and ensure they're ordered, then reverse them.
    address addr0 = vm.addr(s_validSignerKeys[0]);
    address addr1 = vm.addr(s_validSignerKeys[1]);

    uint256 key0 = s_validSignerKeys[0];
    uint256 key1 = s_validSignerKeys[1];

    // Ensure addr0 < addr1 so we can reverse the order.
    if (addr0 > addr1) {
      (addr0, addr1) = (addr1, addr0);
      (key0, key1) = (key1, key0);
    }

    // Now create signatures in reverse order (addr1 first, addr0 second).
    (bytes32 r1, bytes32 s1) = _signWithV27(key1, TEST_HASH);
    (bytes32 r0, bytes32 s0) = _signWithV27(key0, TEST_HASH);
    bytes memory signatures = abi.encodePacked(r1, s1, r0, s0);

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.NonOrderedOrNonUniqueSignatures.selector));
    s_sigQuorumVerifier.validateSignatures(TEST_HASH, signatures);
  }

  function test_validateSignatures_RevertWhen_DuplicateSignatures() public {
    s_sigQuorumVerifier.setSignatureConfig(s_validSigners, 2);

    // Create duplicate signatures manually.
    (bytes32 r, bytes32 s) = _signWithV27(s_validSignerKeys[0], TEST_HASH);
    bytes memory signatures = abi.encodePacked(r, s, r, s); // Same signature twice.

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.NonOrderedOrNonUniqueSignatures.selector));
    s_sigQuorumVerifier.validateSignatures(TEST_HASH, signatures);
  }

  function test_validateSignatures_RevertWhen_InvalidSignatureLength() public {
    s_sigQuorumVerifier.setSignatureConfig(s_validSigners, 1);

    // Create signature with wrong length (63 bytes instead of 64).
    bytes memory invalidSignature = new bytes(63);

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.WrongNumberOfSignatures.selector));
    s_sigQuorumVerifier.validateSignatures(TEST_HASH, invalidSignature);
  }
}
