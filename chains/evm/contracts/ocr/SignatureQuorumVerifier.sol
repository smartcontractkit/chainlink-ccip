// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.4;

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice Onchain verification of reports from the offchain reporting protocol with multiple OCR plugin support.
contract SignatureQuorumVerifier is ITypeAndVersion, Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;

  string public constant override typeAndVersion = "SignatureQuorumVerifier 1.7.0-dev";

  /// @notice Triggers a new run of the offchain reporting protocol.
  /// @param configDigest configDigest of this configuration.
  /// @param signers ith element is address ith oracle uses to sign a report.
  /// @param F maximum number of faulty/dishonest oracles the protocol can tolerate while still working correctly.
  event ConfigSet(bytes32 configDigest, address[] signers, uint8 F);

  error InvalidConfig();
  error ConfigDigestMismatch(bytes32 expected, bytes32 actual);
  error ForkedChain(uint256 expected, uint256 actual);
  error WrongNumberOfSignatures();
  error SignaturesOutOfRegistration();
  error UnauthorizedSigner();
  error NonUniqueSignatures();
  error OracleCannotBeZeroAddress();

  struct SignatureConfigConfig {
    bytes32 configDigest;
    uint8 F; //  maximum number of faulty/dishonest oracles the system can tolerate.
    uint8 n; //  number of configured signers.
  }

  /// @notice Args to update an OCR Config.
  struct SignatureConfigArgs {
    bytes32 configDigest; // The new config digest.
    uint8 F; // Maximum number of faulty/dishonest oracles.
    address[] signers; // signing address of each oracle.
  }

  struct SignatureProof {
    bytes32 configDigest; // The config digest of the report.
    bytes32[] rs; // R components of the signatures.
    bytes32[] ss; // S components of the signatures.
  }

  SignatureConfigConfig internal s_ocrConfig;

  EnumerableSet.AddressSet internal s_signers;

  uint256 internal immutable i_chainID;

  constructor() {
    i_chainID = block.chainid;
  }

  // TODO allow older digests.
  function _validateConfigDigest(
    bytes32 configDigest
  ) internal view {
    if (s_ocrConfig.configDigest != configDigest) {
      revert ConfigDigestMismatch(s_ocrConfig.configDigest, configDigest);
    }
  }

  function _validateOCRSignatures(bytes32 reportHash, bytes32 blobHash, bytes calldata ocrProof) internal view {
    SignatureProof memory report = abi.decode(ocrProof, (SignatureProof));

    // If the cached chainID at time of deployment doesn't match the current chainID, we reject all signed reports.
    // This avoids a (rare) scenario where chain A forks into chain A and A', A' still has configDigest calculated
    // from chain A and so OCR reports will be valid on both forks.
    if (i_chainID != block.chainid) revert ForkedChain(i_chainID, block.chainid);

    if (report.rs.length != s_ocrConfig.F + 1) revert WrongNumberOfSignatures();
    if (report.rs.length != report.ss.length) revert SignaturesOutOfRegistration();

    _verifySignatures(keccak256(abi.encode(reportHash, blobHash, report.configDigest)), report.rs, report.ss);
  }

  /// @notice Verifies the signatures of a hashed report value for one OCR plugin type.
  /// @param hashedReport hashed encoded packing of report + reportContext.
  /// @param rs ith element is the R components of the ith signature on report. Must have at most MAX_NUM_ORACLES entries.
  /// @param ss ith element is the S components of the ith signature on report. Must have at most MAX_NUM_ORACLES entries.
  /// @dev we assume all signatures use a `v` value of 27.
  function _verifySignatures(bytes32 hashedReport, bytes32[] memory rs, bytes32[] memory ss) internal view {
    // Verify signatures attached to report. Using a uint256 means we can only verify up to 256 oracles.
    uint160 lastSigner = 0;

    uint256 numberOfSignatures = rs.length;
    for (uint256 i; i < numberOfSignatures; ++i) {
      // Safe from ECDSA malleability here since we check for duplicate signers.
      address signer = ecrecover(hashedReport, 27, rs[i], ss[i]);

      // Check that the signer is registered as an oracle.
      if (!s_signers.contains(signer)) revert UnauthorizedSigner();
      // This requires ordered signatures to check for duplicates. This also disallows the zero address.
      if (uint160(signer) <= lastSigner) revert NonUniqueSignatures();
    }
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Sets offchain reporting protocol configuration incl. participating oracles for a single OCR plugin type.
  /// @param ocrConfigArgs OCR config update args.
  function setSignatureConfig(
    SignatureConfigArgs memory ocrConfigArgs
  ) external onlyOwner {
    if (ocrConfigArgs.F == 0) revert InvalidConfig();

    address[] memory newSigners = ocrConfigArgs.signers;

    _clearAllSigners();

    // Add new signers to the set of oracles.
    for (uint256 i = 0; i < newSigners.length; ++i) {
      if (newSigners[i] == address(0)) revert OracleCannotBeZeroAddress();

      s_signers.add(newSigners[i]);
    }

    s_ocrConfig =
      SignatureConfigConfig({configDigest: ocrConfigArgs.configDigest, F: ocrConfigArgs.F, n: uint8(newSigners.length)});

    emit ConfigSet(ocrConfigArgs.configDigest, newSigners, ocrConfigArgs.F);
  }

  /// @notice Clears all the signers.
  function _clearAllSigners() internal {
    address[] memory signers = s_signers.values();

    for (uint256 i = 0; i < signers.length; ++i) {
      s_signers.remove(signers[i]);
    }
  }

  /// @notice Information about current offchain reporting protocol configuration.
  function latestConfigDetails() external view returns (SignatureConfigConfig memory ocrConfig) {
    return s_ocrConfig;
  }
}
