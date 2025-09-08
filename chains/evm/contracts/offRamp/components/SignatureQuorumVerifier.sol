// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.4;

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

contract SignatureQuorumVerifier is Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;

  /// @param signers ith element is address ith oracle uses to sign a report.
  /// @param F maximum number of faulty/dishonest singers the protocol can tolerate while still working correctly.
  event ConfigSet(address[] signers, uint8 F);

  error InvalidConfig();
  error ForkedChain(uint256 expected, uint256 actual);
  error WrongNumberOfSignatures();
  error SignaturesOutOfRegistration();
  error UnauthorizedSigner();
  error NonUniqueSignatures();
  error OracleCannotBeZeroAddress();

  struct SignatureConfigArgs {
    uint8 threshold; // The expected number of signatures for a report.
    address[] signers; // The signing address of each signer.
  }

  struct SignatureConfigConfig {
    uint8 threshold; // The expected number of signatures for a report.
    EnumerableSet.AddressSet signers; // The signing address of each signer.
  }

  SignatureConfigConfig private s_config;

  uint256 internal immutable i_chainID;

  constructor() {
    i_chainID = block.chainid;
  }

  /// @notice Validates the signatures of a given report hash.
  /// @param reportHash The hash of the report to validate signatures for.
  /// @param rs The r values of the signatures.
  /// @param ss The s values of the signatures.
  /// @dev The v values are assumed to be 27 for all signatures, this can be achieved by using ECDSA malleability.
  function _validateSignatures(bytes32 reportHash, bytes32[] memory rs, bytes32[] memory ss) internal view {
    SignatureConfigConfig storage config = s_config;
    if (config.signers.length() == 0) {
      revert InvalidConfig();
    }

    // If the cached chainID at time of deployment doesn't match the current chainID, we reject all signed reports.
    // This avoids a (rare) scenario where chain A forks into chain A and A', and a report signed on A is replayed on A'.
    if (i_chainID != block.chainid) revert ForkedChain(i_chainID, block.chainid);

    uint256 numberOfSignatures = rs.length;

    if (numberOfSignatures != config.threshold) revert WrongNumberOfSignatures();
    if (numberOfSignatures != ss.length) revert SignaturesOutOfRegistration();

    uint160 lastSigner = 0;

    for (uint256 i; i < numberOfSignatures; ++i) {
      // We use ECDSA malleability to only have signatures with a `v` value of 27.
      address signer = ecrecover(reportHash, 27, rs[i], ss[i]);
      // Check that the signer is registered.
      if (!config.signers.contains(signer)) revert UnauthorizedSigner();
      // This requires ordered signatures to check for duplicates. This also disallows the zero address.
      if (uint160(signer) <= lastSigner) revert NonUniqueSignatures();
      lastSigner = uint160(signer);
    }
  }

  /// @notice Returns the signer sets, and F value.
  function getSignatureConfig() external view returns (SignatureConfigArgs memory) {
    return SignatureConfigArgs({threshold: s_config.threshold, signers: s_config.signers.values()});
  }

  /// @notice Sets a new signature configuration.
  function setSignatureConfig(
    SignatureConfigArgs calldata signatureConfig
  ) external onlyOwner {
    if (signatureConfig.threshold == 0 || signatureConfig.threshold > signatureConfig.signers.length) {
      revert InvalidConfig();
    }

    SignatureConfigConfig storage config = s_config;

    // We must remove all current signers first.
    for (uint256 i = 0; i < config.signers.length();) {
      config.signers.remove(config.signers.at(i));
    }

    // Add new signers.
    for (uint256 signerIndex = 0; signerIndex < signatureConfig.signers.length; ++signerIndex) {
      if (signatureConfig.signers[signerIndex] == address(0)) revert OracleCannotBeZeroAddress();

      // This checks for duplicates.
      if (!config.signers.add(signatureConfig.signers[signerIndex])) {
        revert InvalidConfig();
      }
    }

    config.threshold = signatureConfig.threshold;

    emit ConfigSet(signatureConfig.signers, signatureConfig.threshold);
  }
}
