// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.4;

import {IVerifier} from "../interfaces/IVerifier.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice Onchain verification of reports from the offchain reporting protocol with multiple OCR plugin support.
contract OCRVerifier is IVerifier, ITypeAndVersion, Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;

  string public constant override typeAndVersion = "OCRVerifier 1.7.0-dev";

  // Maximum number of oracles the offchain reporting protocol is designed for
  uint256 internal constant MAX_NUM_ORACLES = 32;

  /// @notice Triggers a new run of the offchain reporting protocol.
  /// @param configDigest configDigest of this configuration.
  /// @param signers ith element is address ith oracle uses to sign a report.
  /// @param F maximum number of faulty/dishonest oracles the protocol can tolerate while still working correctly.
  event ConfigSet(bytes32 configDigest, address[] signers, uint8 F);

  /// @notice Optionally emitted to indicate the latest configDigest and sequence number
  /// for which a report was successfully transmitted. Alternatively, the contract may
  /// use latestConfigDigestAndEpoch with scanLogs set to false.
  event Transmitted(bytes32 configDigest, uint64 sequenceNumber);

  enum InvalidConfigErrorType {
    F_MUST_BE_POSITIVE,
    TOO_MANY_SIGNERS,
    F_TOO_HIGH
  }

  error InvalidConfig(InvalidConfigErrorType errorType);
  error ConfigDigestMismatch(bytes32 expected, bytes32 actual);
  error ForkedChain(uint256 expected, uint256 actual);
  error WrongNumberOfSignatures();
  error SignaturesOutOfRegistration();
  error UnauthorizedSigner();
  error NonUniqueSignatures();
  error OracleCannotBeZeroAddress();

  /// @notice OCR configuration for a single OCR plugin within a DON.
  struct OCRConfig {
    bytes32 configDigest;
    uint8 F; //  maximum number of faulty/dishonest oracles the system can tolerate.
    uint8 n; //  number of configured signers.
  }

  /// @notice Args to update an OCR Config.
  struct OCRConfigArgs {
    bytes32 configDigest; // The new config digest.
    uint8 F; //                              │ Maximum number of faulty/dishonest oracles.
    address[] signers; // signing address of each oracle.
  }

  struct OCRProof {
    bytes32 configDigest; // The config digest of the report.
    uint64 sequenceNumber; // The sequence number of the report.
    bytes32[] rs; // R components of the signatures.
    bytes32[] ss; // S components of the signatures.
  }

  OCRConfig internal s_ocrConfig;

  EnumerableSet.AddressSet internal s_signers;

  uint256 internal immutable i_chainID;

  constructor() {
    i_chainID = block.chainid;
  }

  function validateReport(bytes calldata rawReport, bytes calldata ocrProof, uint256) external {
    OCRProof memory report = abi.decode(ocrProof, (OCRProof));

    if (s_ocrConfig.configDigest != report.configDigest) {
      revert ConfigDigestMismatch(s_ocrConfig.configDigest, report.configDigest);
    }
    // If the cached chainID at time of deployment doesn't match the current chainID, we reject all signed reports.
    // This avoids a (rare) scenario where chain A forks into chain A and A', A' still has configDigest calculated
    // from chain A and so OCR reports will be valid on both forks.
    if (i_chainID != block.chainid) revert ForkedChain(i_chainID, block.chainid);

    if (report.rs.length != s_ocrConfig.F + 1) revert WrongNumberOfSignatures();
    if (report.rs.length != report.ss.length) revert SignaturesOutOfRegistration();

    _verifySignatures(
      keccak256(abi.encode(keccak256(rawReport), report.configDigest, report.sequenceNumber)), report.rs, report.ss
    );

    emit Transmitted(report.configDigest, uint64(uint256(report.sequenceNumber)));
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
  function setOCR3Config(
    OCRConfigArgs memory ocrConfigArgs
  ) external onlyOwner {
    if (ocrConfigArgs.F == 0) revert InvalidConfig(InvalidConfigErrorType.F_MUST_BE_POSITIVE);

    address[] memory newSigners = ocrConfigArgs.signers;

    if (newSigners.length > MAX_NUM_ORACLES) revert InvalidConfig(InvalidConfigErrorType.TOO_MANY_SIGNERS);
    if (newSigners.length <= 3 * ocrConfigArgs.F) revert InvalidConfig(InvalidConfigErrorType.F_TOO_HIGH);

    _clearAllSigners();

    // Add new signers to the set of oracles.
    for (uint256 i = 0; i < newSigners.length; ++i) {
      if (newSigners[i] == address(0)) revert OracleCannotBeZeroAddress();

      s_signers.add(newSigners[i]);
    }

    s_ocrConfig = OCRConfig({configDigest: ocrConfigArgs.configDigest, F: ocrConfigArgs.F, n: uint8(newSigners.length)});

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
  function latestConfigDetails() external view returns (OCRConfig memory ocrConfig) {
    return s_ocrConfig;
  }
}
