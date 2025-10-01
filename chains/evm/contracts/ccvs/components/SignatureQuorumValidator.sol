// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.4;

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.0.2/utils/structs/EnumerableSet.sol";

contract SignatureQuorumValidator is Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;

  /// @param signers List of valid signers of which only `threshold` are required to sign each report.
  /// @param threshold The number of signatures required for a report to be valid.
  event SignatureConfigSet(address[] signers, uint8 threshold);

  error InvalidSignatureConfig();
  error ForkedChain(uint256 expected, uint256 actual);
  error WrongNumberOfSignatures();
  error UnauthorizedSigner();
  error NonOrderedOrNonUniqueSignatures();
  error OracleCannotBeZeroAddress();

  uint256 internal constant SIGNATURE_LENGTH = 64;
  uint256 internal constant SIGNATURE_COMPONENT_LENGTH = 32;

  uint256 internal immutable i_chainID;

  /// @notice List of valid signers of which only `threshold` are required to sign each report.
  EnumerableSet.AddressSet internal s_signers;

  /// @notice The number of signatures required for a report to be valid.
  uint8 internal s_threshold;

  constructor() {
    i_chainID = block.chainid;
  }

  /// @notice Validates the signatures of a given report hash. IMPORTANT: the signatures must be provided in order of
  /// their signer addresses. This is required to efficiently check for duplicated signatures. If any signature is out
  /// of order, this function will revert with `NonOrderedOrNonUniqueSignatures`.
  /// @param signedHash The hash that is signed.
  /// @param signatures The concatenated signatures to validate. Each signature is 64 bytes long, consisting of r
  /// (32 bytes) and s (32 bytes). The signatures must be provided in order of their signer addresses. For example, if
  /// the signers are [A, B, C] with addresses [0x1, 0x2, 0x3], the signatures must be provided ordered as [A, B, C].
  /// @dev The v values are assumed to be 27 for all signatures, this can be achieved by using ECDSA malleability.
  function _validateSignatures(bytes32 signedHash, bytes calldata signatures) internal view {
    if (s_signers.length() == 0) {
      revert InvalidSignatureConfig();
    }

    // If the cached chainID at time of deployment doesn't match the current chainID, we reject all signed reports.
    // This avoids a (rare) scenario where chain A forks into chain A and A', and a report signed on A is replayed on A'.
    if (i_chainID != block.chainid) revert ForkedChain(i_chainID, block.chainid);

    uint256 numberOfSignatures = signatures.length / SIGNATURE_LENGTH;

    uint256 threshold = s_threshold;

    // We allow more signatures than the threshold, but we will only validate up to the threshold to save gas.
    // This still preserves the security properties while adding flexibility.
    if (numberOfSignatures < threshold) revert WrongNumberOfSignatures();

    uint160 lastSigner = 0;

    for (uint256 i; i < threshold; ++i) {
      uint256 offset = i * SIGNATURE_LENGTH;
      // We use ECDSA malleability to only have signatures with a `v` value of 27.
      address signer = ecrecover(
        signedHash,
        27,
        bytes32(signatures[offset:offset + SIGNATURE_COMPONENT_LENGTH]),
        bytes32(signatures[offset + SIGNATURE_COMPONENT_LENGTH:offset + SIGNATURE_LENGTH])
      );
      // Check that the signer is registered.
      if (!s_signers.contains(signer)) revert UnauthorizedSigner();
      // This requires ordered signatures to check for duplicates. This also disallows the zero address.
      if (uint160(signer) <= lastSigner) revert NonOrderedOrNonUniqueSignatures();
      lastSigner = uint160(signer);
    }
  }

  /// @notice Returns the signer sets, and F value.
  function getSignatureConfig() external view returns (address[] memory, uint8) {
    return (s_signers.values(), s_threshold);
  }

  /// @notice Sets a new signature configuration.
  function setSignatureConfig(address[] memory signers, uint8 threshold) external onlyOwner {
    if (threshold == 0 || threshold > signers.length) {
      revert InvalidSignatureConfig();
    }

    // We must remove all current signers first.
    while (s_signers.length() > 0) {
      s_signers.remove(s_signers.at(0));
    }

    // Add new signers.
    for (uint256 signerIndex = 0; signerIndex < signers.length; ++signerIndex) {
      if (signers[signerIndex] == address(0)) revert OracleCannotBeZeroAddress();

      // This checks for duplicates.
      if (!s_signers.add(signers[signerIndex])) {
        revert InvalidSignatureConfig();
      }
    }

    s_threshold = threshold;

    emit SignatureConfigSet(signers, threshold);
  }
}
