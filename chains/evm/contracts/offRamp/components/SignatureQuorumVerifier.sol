// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.4;

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

contract SignatureQuorumVerifier is Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;

  /// @param signers List of valid signers of which only `threshold` are required to sign each report.
  /// @param threshold The number of signatures required for a report to be valid.
  event ConfigSet(address[] signers, uint8 threshold);

  error InvalidConfig();
  error ForkedChain(uint256 expected, uint256 actual);
  error WrongNumberOfSignatures();
  error SignaturesOutOfRegistration();
  error UnauthorizedSigner();
  error NonOrderedOrNonUniqueSignatures();
  error OracleCannotBeZeroAddress();

  uint256 internal immutable i_chainID;

  /// @notice List of valid signers of which only `threshold` are required to sign each report.
  EnumerableSet.AddressSet internal s_signers;

  /// @notice The number of signatures required for a report to be valid.
  uint8 internal s_threshold;

  constructor() {
    i_chainID = block.chainid;
  }

  /// @notice Validates the signatures of a given report hash.
  /// @param reportHash The hash of the report to validate signatures for.
  /// @param rs The r values of the signatures.
  /// @param ss The s values of the signatures.
  /// @dev The v values are assumed to be 27 for all signatures, this can be achieved by using ECDSA malleability.
  function _validateSignatures(bytes32 reportHash, bytes32[] memory rs, bytes32[] memory ss) internal view {
    if (s_signers.length() == 0) {
      revert InvalidConfig();
    }

    // If the cached chainID at time of deployment doesn't match the current chainID, we reject all signed reports.
    // This avoids a (rare) scenario where chain A forks into chain A and A', and a report signed on A is replayed on A'.
    if (i_chainID != block.chainid) revert ForkedChain(i_chainID, block.chainid);

    uint256 numberOfSignatures = rs.length;

    if (numberOfSignatures != s_threshold) revert WrongNumberOfSignatures();
    if (numberOfSignatures != ss.length) revert SignaturesOutOfRegistration();

    uint160 lastSigner = 0;

    for (uint256 i; i < numberOfSignatures; ++i) {
      // We use ECDSA malleability to only have signatures with a `v` value of 27.
      address signer = ecrecover(reportHash, 27, rs[i], ss[i]);
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
      revert InvalidConfig();
    }

    // We must remove all current signers first.
    for (uint256 i = 0; i < s_signers.length();) {
      s_signers.remove(s_signers.at(i));
    }

    // Add new signers.
    for (uint256 signerIndex = 0; signerIndex < signers.length; ++signerIndex) {
      if (signers[signerIndex] == address(0)) revert OracleCannotBeZeroAddress();

      // This checks for duplicates.
      if (!s_signers.add(signers[signerIndex])) {
        revert InvalidConfig();
      }
    }

    s_threshold = threshold;

    emit ConfigSet(signers, threshold);
  }
}
