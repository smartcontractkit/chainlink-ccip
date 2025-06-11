// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.4;

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice Onchain verification of reports from the offchain reporting protocol with multiple OCR plugin support.
contract OCRVerifier is ITypeAndVersion, Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;

  string public constant override typeAndVersion = "CommitVerifier 1.7.0-dev";

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

  OCRConfig internal s_ocrConfig;

  EnumerableSet.AddressSet internal s_signers;

  uint256 internal immutable i_chainID;

  constructor() {
    i_chainID = block.chainid;
  }

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

  /// @param report serialized report, which the signatures are signing.
  /// @param rs ith element is the R components of the ith signature on report. Must have at most MAX_NUM_ORACLES entries.
  /// @param ss ith element is the S components of the ith signature on report. Must have at most MAX_NUM_ORACLES entries.
  function validateReport(
    // NOTE: If these parameters are changed, expectedMsgDataLength and/or TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT
    // need to be changed accordingly.
    bytes32[2] calldata reportContext,
    bytes calldata report,
    bytes32[] memory rs,
    bytes32[] memory ss
  ) external {
    // reportContext consists of:
    // reportContext[0]: ConfigDigest.
    // reportContext[1]: 24 byte padding, 8 byte sequence number.
    bytes32 configDigest = reportContext[0];

    OCRConfig memory configInfo = s_ocrConfig;
    if (configInfo.configDigest != configDigest) {
      revert ConfigDigestMismatch(configInfo.configDigest, configDigest);
    }
    // If the cached chainID at time of deployment doesn't match the current chainID, we reject all signed reports.
    // This avoids a (rare) scenario where chain A forks into chain A and A', A' still has configDigest calculated
    // from chain A and so OCR reports will be valid on both forks.
    _whenChainNotForked();

    // Scoping to reduce stack pressure.
    {
      if (rs.length != configInfo.F + 1) revert WrongNumberOfSignatures();
      if (rs.length != ss.length) revert SignaturesOutOfRegistration();
    }

    bytes32 h = keccak256(abi.encodePacked(keccak256(report), reportContext));
    _verifySignatures(h, rs, ss);

    emit Transmitted(configDigest, uint64(uint256(reportContext[1])));
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
      // This requires ordered signatures, so we can check for duplicates. This also disallows the zero address.
      if (uint160(signer) <= lastSigner) revert NonUniqueSignatures();
    }
  }

  /// @notice Validates that the chain ID has not diverged after deployment. Reverts if the chain IDs do not match.
  function _whenChainNotForked() internal view {
    if (i_chainID != block.chainid) revert ForkedChain(i_chainID, block.chainid);
  }

  /// @notice Information about current offchain reporting protocol configuration.
  function latestConfigDetails() external view returns (OCRConfig memory ocrConfig) {
    return s_ocrConfig;
  }
}
