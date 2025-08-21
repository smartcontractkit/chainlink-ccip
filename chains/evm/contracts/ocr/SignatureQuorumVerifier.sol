// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.4;

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

contract SignatureQuorumVerifier is Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;

  /// @param configDigest configDigest of this configuration.
  /// @param signers ith element is address ith oracle uses to sign a report.
  /// @param F maximum number of faulty/dishonest singers the protocol can tolerate while still working correctly.
  event ConfigSet(bytes32 configDigest, address[] signers, uint8 F);
  event ConfigRevoked(bytes32 configDigest);

  error InvalidConfig();
  error InvalidConfigDigest(bytes32 configDigest);
  error ConfigDigestAlreadyExists(bytes32 configDigest);
  error ForkedChain(uint256 expected, uint256 actual);
  error WrongNumberOfSignatures();
  error SignaturesOutOfRegistration();
  error UnauthorizedSigner();
  error NonUniqueSignatures();
  error OracleCannotBeZeroAddress();

  /// @notice Args to update an OCR Config.
  struct SignatureConfigArgs {
    bytes32 configDigest; // The new config digest.
    uint8 F; // Maximum number of faulty/dishonest oracles.
    address[] signers; // signing address of each oracle.
  }

  struct SignatureConfigConfig {
    uint8 F; //  maximum number of faulty/dishonest oracles the system can tolerate.
    EnumerableSet.AddressSet signers;
  }

  struct SignatureProof {
    bytes32[] rs; // R components of the signatures.
    bytes32[] ss; // S components of the signatures.
  }

  mapping(bytes32 configDigest => SignatureConfigConfig config) internal s_signatureConfig;

  uint256 internal immutable i_chainID;

  constructor() {
    i_chainID = block.chainid;
  }

  function _validateSignatures(bytes32 reportHash, bytes32 configDigest, bytes memory signatureProof) internal view {
    // We allow proving of older messages that might have been signed by an older set of signers. This means we need to
    // get the set of signers for the given configDigest.
    SignatureConfigConfig storage config = s_signatureConfig[configDigest];
    if (config.signers.length() == 0) {
      revert InvalidConfigDigest(configDigest);
    }

    SignatureProof memory proofs = abi.decode(signatureProof, (SignatureProof));

    // If the cached chainID at time of deployment doesn't match the current chainID, we reject all signed reports.
    // This avoids a (rare) scenario where chain A forks into chain A and A', A' still has configDigest calculated
    // from chain A and so OCR reports will be valid on both forks.
    if (i_chainID != block.chainid) revert ForkedChain(i_chainID, block.chainid);

    uint256 numberOfSignatures = proofs.rs.length;

    if (numberOfSignatures != config.F + 1) revert WrongNumberOfSignatures();
    if (numberOfSignatures != proofs.ss.length) revert SignaturesOutOfRegistration();

    uint160 lastSigner = 0;

    for (uint256 i; i < numberOfSignatures; ++i) {
      // We use ECDSA malleability to only have signatures with a `v` value of 27.
      address signer = ecrecover(reportHash, 27, proofs.rs[i], proofs.ss[i]);
      // Check that the signer is registered.
      if (!config.signers.contains(signer)) revert UnauthorizedSigner();
      // This requires ordered signatures to check for duplicates. This also disallows the zero address.
      if (uint160(signer) <= lastSigner) revert NonUniqueSignatures();
      lastSigner = uint160(signer);
    }
  }

  /// @notice Sets offchain reporting protocol configuration incl. participating oracles for a single OCR plugin type.
  function setSignatureConfig(
    SignatureConfigArgs calldata signatureConfig
  ) external onlyOwner {
    if (signatureConfig.F == 0) revert InvalidConfig();

    SignatureConfigConfig storage configForDigest = s_signatureConfig[signatureConfig.configDigest];

    // If the configDigest already exists, we cannot modify it as there might be signed transactions that rely on this
    // exact signer set.
    if (configForDigest.signers.length() != 0) {
      revert ConfigDigestAlreadyExists(signatureConfig.configDigest);
    }

    // Add new signers.
    for (uint256 i = 0; i < signatureConfig.signers.length; ++i) {
      if (signatureConfig.signers[i] == address(0)) revert OracleCannotBeZeroAddress();

      configForDigest.signers.add(signatureConfig.signers[i]);
    }

    configForDigest.F = signatureConfig.F;

    emit ConfigSet(signatureConfig.configDigest, signatureConfig.signers, signatureConfig.F);
  }

  function revokeConfigDigest(
    bytes32 configDigest
  ) external onlyOwner {
    SignatureConfigConfig storage configForDigest = s_signatureConfig[configDigest];
    if (configForDigest.signers.length() == 0) {
      revert InvalidConfigDigest(configDigest);
    }

    delete s_signatureConfig[configDigest];

    emit ConfigRevoked(configDigest);
  }
}
