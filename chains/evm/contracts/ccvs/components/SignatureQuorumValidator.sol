// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.4;

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

contract SignatureQuorumValidator is Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.UintSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;

  /// @param sourceChainSelector The selector of the source chain.
  /// @param signers List of valid signers of which only `threshold` are required to sign each report.
  /// @param threshold The number of signatures required for a report to be valid.
  event SignatureConfigSet(uint64 indexed sourceChainSelector, address[] signers, uint8 threshold);

  error InvalidSignatureConfig();
  error NoSignersSetForSource(uint64 sourceChainSelector);
  error ForkedChain(uint256 expected, uint256 actual);
  error WrongNumberOfSignatures();
  error UnauthorizedSigner();
  error NonOrderedOrNonUniqueSignatures();
  error SignerCannotBeZeroAddress();

  /// @dev Struct that contains the signer configuration
  struct SignerConfig {
    EnumerableSet.AddressSet signers; // List of valid signers of which only `threshold` are required to sign each report.
    uint8 threshold; // The number of signatures required for a report to be valid.
  }

  struct SignersUpdate {
    uint64 sourceChainSelector; // The source chain selector to apply the updates to.
    uint8 threshold; // The required threshold, must be non-zero and it can not be greater than signers.length.
    address[] signers; // The desired final set of signers for the source chain.
  }

  // STATIC CONFIG
  uint256 internal constant SIGNATURE_LENGTH = 64;
  uint256 internal constant SIGNATURE_COMPONENT_LENGTH = 32;

  uint256 internal immutable i_chainID;

  mapping(uint64 sourceChainSelector => SignerConfig cfg) private s_signerConfigs;

  EnumerableSet.UintSet private s_configuredChains;

  constructor() {
    i_chainID = block.chainid;
  }

  /// @notice Validates the signatures of a given report hash. IMPORTANT: the signatures must be provided in order of
  /// their signer addresses. This is required to efficiently check for duplicated signatures. If any signature is out
  /// of order, this function will revert with `NonOrderedOrNonUniqueSignatures`.
  /// @param sourceChainSelector The selector of the source chain.
  /// @param signedHash The hash that is signed.
  /// @param signatures The concatenated signatures to validate. Each signature is 64 bytes long, consisting of r
  /// (32 bytes) and s (32 bytes). The signatures must be provided in order of their signer addresses. For example, if
  /// the signers are [A, B, C] with addresses [0x1, 0x2, 0x3], the signatures must be provided ordered as [A, B, C].
  /// @dev The v values are assumed to be 27 for all signatures, this can be achieved by using ECDSA malleability.
  function _validateSignatures(uint64 sourceChainSelector, bytes32 signedHash, bytes calldata signatures) internal view {
    SignerConfig storage cfg = s_signerConfigs[sourceChainSelector];
    if (cfg.signers.length() == 0) {
      revert NoSignersSetForSource(sourceChainSelector);
    }

    // If the cached chainID at time of deployment doesn't match the current chainID, we reject all signed reports.
    // This avoids a (rare) scenario where chain A forks into chain A and A', and a report signed on A is replayed on A'.
    if (i_chainID != block.chainid) revert ForkedChain(i_chainID, block.chainid);

    uint256 numberOfSignatures = signatures.length / SIGNATURE_LENGTH;

    uint256 threshold = cfg.threshold;

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
      if (!cfg.signers.contains(signer)) revert UnauthorizedSigner();
      // This requires ordered signatures to check for duplicates. This also disallows the zero address.
      if (uint160(signer) <= lastSigner) revert NonOrderedOrNonUniqueSignatures();
      lastSigner = uint160(signer);
    }
  }

  /// @notice Returns the signer set and threshold.
  function getSignatureConfig(
    uint64 sourceChainSelector
  ) external view returns (address[] memory, uint8) {
    SignerConfig storage cfg = s_signerConfigs[sourceChainSelector];
    return (cfg.signers.values(), cfg.threshold);
  }

  /// @notice Returns all configured source chain selectors signer sets and thresholds.
  function getAllSignatureConfigs()
    external
    view
    returns (uint64[] memory sourceChainSelectors, address[][] memory signerSets, uint8[] memory thresholds)
  {
    uint256[] memory sourceChainSelectorSet = s_configuredChains.values();
    sourceChainSelectors = new uint64[](sourceChainSelectorSet.length);
    signerSets = new address[][](sourceChainSelectorSet.length);
    thresholds = new uint8[](sourceChainSelectorSet.length);

    for (uint256 i; i < sourceChainSelectorSet.length; ++i) {
      // Okay to cast down without checking as set is populated using uint64 sourceChainSelector input in `applySignersUpdates`.
      uint64 sourceChainSelector = uint64(sourceChainSelectorSet[i]);
      sourceChainSelectors[i] = sourceChainSelector;
      SignerConfig storage cfg = s_signerConfigs[sourceChainSelector];
      signerSets[i] = cfg.signers.values();
      thresholds[i] = cfg.threshold;
    }

    return (sourceChainSelectors, signerSets, thresholds);
  }

  /// @notice Removes source chain selectors, and applies multiple signers updates.
  /// @dev Last signers update wins. If a source chain selector is repeated in `signersUpdates` then the last one will be the state set.
  /// @param sourceChainSelectorsToRemove The selectors that should have their signer configuration cleared.
  /// @param signersUpdates The desired signer configuration updates to apply per source chain selector.
  function applySignersUpdates(
    uint64[] calldata sourceChainSelectorsToRemove,
    SignersUpdate[] calldata signersUpdates
  ) external onlyOwner {
    // Handle removals first.
    for (uint256 i = 0; i < sourceChainSelectorsToRemove.length; ++i) {
      if (s_configuredChains.contains(sourceChainSelectorsToRemove[i])) {
        SignerConfig storage cfg = s_signerConfigs[sourceChainSelectorsToRemove[i]];

        //Remove all signers.
        while (cfg.signers.length() > 0) {
          cfg.signers.remove(cfg.signers.at(0));
        }

        cfg.threshold = 0;
        s_configuredChains.remove(sourceChainSelectorsToRemove[i]);

        emit SignatureConfigSet(sourceChainSelectorsToRemove[i], new address[](0), 0);
      } // else noop
    }

    // Now handle signerUpdates.
    for (uint256 i = 0; i < signersUpdates.length; ++i) {
      SignersUpdate memory update = signersUpdates[i];

      if (update.threshold == 0 || update.threshold > update.signers.length) {
        revert InvalidSignatureConfig();
      }

      SignerConfig storage cfg = s_signerConfigs[update.sourceChainSelector];

      // We must remove all current signers first.
      while (cfg.signers.length() > 0) {
        cfg.signers.remove(cfg.signers.at(0));
      }

      // Add new signers.
      for (uint256 signerIndex = 0; signerIndex < update.signers.length; ++signerIndex) {
        if (update.signers[signerIndex] == address(0)) revert SignerCannotBeZeroAddress();

        // This checks for duplicates.
        if (!cfg.signers.add(update.signers[signerIndex])) {
          revert InvalidSignatureConfig();
        }
      }

      cfg.threshold = update.threshold;
      s_configuredChains.add(update.sourceChainSelector);

      emit SignatureConfigSet(update.sourceChainSelector, update.signers, update.threshold);
    }
  }
}
