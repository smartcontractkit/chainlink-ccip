// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../libraries/Internal.sol";
import {SuperchainInterop} from "../libraries/SuperchainInterop.sol";
import {OffRamp} from "./OffRamp.sol";

import {ICrossL2Inbox} from "../vendor/optimism/interop-lib/v0/src/interfaces/ICrossL2Inbox.sol";
import {Identifier} from "../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";

/// @notice This OffRamp supports OP Superchain Interop. It leverages CrossL2Inbox to verify validity of a message,
/// as opposed to relying on roots signed by Commit DON.
/// @dev This OffRamp maintains the same Internal.ExecutionReport interface for execute, but it enforces
/// exactly 1 message per report. Batching is not supported, because this OffRamp only runs on OP L2s,
/// the benefit of batching is minimal, it is not worth the complexity.
contract OffRampOverSuperchainInterop is OffRamp {
  error CrossL2InboxCannotBeZero();
  error InvalidSourceChainSelector(uint64 sourceChainSelector, uint64 expected);
  error InvalidDestChainSelector(uint64 destChainSelector, uint64 expected);
  error InvalidSourceOnRamp(uint64 sourceChainSelector, address sourceOnRamp);
  error ZeroChainIdNotAllowed();
  error ChainIdNotConfiguredForSelector(uint64 chainSelector);
  error ChainIdMismatch(uint64 chainSelector, uint256 chainId, uint256 expectedChainId);
  error OperationNotSupportedByThisOffRampType();
  error ReportMustContainExactlyOneMessage();
  error InvalidProofsWordLength(uint256 length, uint256 expected);
  error ProofFlagBitsMustBeZero();
  error InvalidEncodingOfIdentifierInProofs(bytes32[] proofs);

  event ChainSelectorToChainIdConfigUpdated(uint64 indexed chainSelector, uint256 indexed chainId);
  event ChainSelectorToChainIdConfigRemoved(uint64 indexed chainSelector, uint256 indexed chainId);

  struct ChainSelectorToChainIdConfigArgs {
    uint64 chainSelector;
    uint256 chainId;
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "OffRampOverSuperchainInterop 1.6.2-dev";
  /// @dev CrossL2Inbox is a pre-deploy at a fixed address on OP L2s.
  ICrossL2Inbox internal immutable i_crossL2Inbox;

  /// @dev Resolve source chain selector to source chainId for message verification.
  mapping(uint64 chainSelector => uint256 chainId) private s_chainSelectorToChainId;

  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    SourceChainConfigArgs[] memory sourceChainConfigs,
    address crossL2Inbox,
    ChainSelectorToChainIdConfigArgs[] memory chainSelectorToChainIdConfigArgs
  ) OffRamp(staticConfig, dynamicConfig, sourceChainConfigs) {
    if (crossL2Inbox == address(0)) {
      revert CrossL2InboxCannotBeZero();
    }
    i_crossL2Inbox = ICrossL2Inbox(crossL2Inbox);

    _applyChainSelectorToChainIdConfigUpdates(new uint64[](0), chainSelectorToChainIdConfigArgs);
  }

  /// @notice Construct the Identifier from the proofData, then log hash from the message.
  /// @param message The message to construct proofData and log hash for.
  /// @param proofs Identifier split into an ordered array of 5 32-byte words.
  /// @return identifier The identifier constructed from the proofs.
  /// @return logHash The log hash constructed from the message.
  function _constructProofs(
    Internal.Any2EVMRampMessage memory message,
    bytes32[] memory proofs
  ) internal pure returns (Identifier memory identifier, bytes32 logHash) {
    if (proofs.length != 5) revert InvalidProofsWordLength(proofs.length, 5);

    // Sanity check that origin, blockNumber, timestamp, and chainId cannot be zero.
    // Log index can be zero.
    if (proofs[0] == 0 || proofs[1] == 0 || proofs[3] == 0 || proofs[4] == 0) {
      revert InvalidEncodingOfIdentifierInProofs(proofs);
    }

    identifier = Identifier({
      origin: address(uint160(uint256(proofs[0]))),
      blockNumber: uint256(proofs[1]),
      logIndex: uint256(proofs[2]),
      timestamp: uint256(proofs[3]),
      chainId: uint256(proofs[4])
    });

    // Concatenation as opposed to abi.encode all fields is necessary to construct correct log data
    // [0x00] event_selector
    // [0x20] dest_chain_selector
    // [0x40] sequence_number
    // [0x60] offset_to_message ← should be 0x20, but would be 0x80 if abi.encode all fields at once
    // [0x80] message...
    logHash = keccak256(
      bytes.concat(
        SuperchainInterop.CCIPSuperchainMessageSent.selector,
        bytes32(uint256(message.header.destChainSelector)),
        bytes32(uint256(message.header.sequenceNumber)),
        abi.encode(message)
      )
    );

    return (identifier, logHash);
  }

  /// @notice Verify the message was indeed sent on the source chain by checking against the CrossL2Inbox.
  /// @dev No trust assumption can be placed on the report, every field of the report can be forged.
  /// In addition to CrossL2Inbox validation, many checks are necessary,
  /// e.g. Identifier.chainId must match sourceChainSelector.
  /// @param sourceChainSelector The source chain selector of the message.
  /// @param report The execution report to verify.
  /// @return timestampCommitted The source timestamp of the message.
  /// @return hashedLeaves Array of 1 hashed message.
  function _verifyReport(
    uint64 sourceChainSelector,
    Internal.ExecutionReport memory report
  ) internal virtual override returns (uint256 timestampCommitted, bytes32[] memory hashedLeaves) {
    if (report.messages.length != 1) revert ReportMustContainExactlyOneMessage();
    Internal.Any2EVMRampMessage memory message = report.messages[0];

    // Sanity check that the report is constructed correctly.
    if (message.header.sourceChainSelector != sourceChainSelector) {
      revert InvalidSourceChainSelector(message.header.sourceChainSelector, sourceChainSelector);
    }
    if (report.proofFlagBits != 0) {
      revert ProofFlagBitsMustBeZero();
    }
    // Validate that the message is meant for this chain.
    if (message.header.destChainSelector != i_chainSelector) {
      revert InvalidDestChainSelector(message.header.destChainSelector, i_chainSelector);
    }

    (Identifier memory identifier, bytes32 logHash) = _constructProofs(message, report.proofs);
    address onRampAddress = abi.decode(_getEnabledSourceChainConfig(sourceChainSelector).onRamp, (address));

    // Validate that the message was emitted by the corresponding OnRamp.
    if (identifier.origin != onRampAddress) {
      revert InvalidSourceOnRamp(sourceChainSelector, identifier.origin);
    }
    // Validate that the chainId matches the given sourceChainSelector, such that it must have come from
    // the OnRamp on the correct source chain.
    uint256 expectedChainId = s_chainSelectorToChainId[sourceChainSelector];
    if (expectedChainId == 0) {
      revert ChainIdNotConfiguredForSelector(sourceChainSelector);
    }
    if (expectedChainId != identifier.chainId) {
      revert ChainIdMismatch(sourceChainSelector, identifier.chainId, expectedChainId);
    }

    // SECURITY CRITICAL CHECK.
    // Validate the exact log was emitted on the source chain.
    i_crossL2Inbox.validateMessage(identifier, logHash);

    hashedLeaves = new bytes32[](1);
    hashedLeaves[0] = SuperchainInterop._hashInteropMessage(message, onRampAddress);

    // Non-zero timestamp signals the message is verified.
    // The timestamp is additionally used to determine if manual exec window has opened.
    // Given there is no Commit, we use the Identifier timestmap, which is when the message was sent.
    // If this OffRamp only accepts low-latency messages from source chains within OP Mesh,
    // this will be very close to the commit timestamp.
    // If this OffRamp accepts messages from less-trusted sources, e.g a long-tail OP L2 in current
    // chain's dependency set, a higher satefy level may be required, leading to higher latency.
    // If so, `permissionLessExecutionThresholdSeconds` may need to be increased.
    return (identifier.timestamp, hashedLeaves);
  }

  /// @notice Commit is not supported by OffRamp over Superchain Interop, it is replaced by CrossL2Inbox.
  /// @dev This function is explicitly removed to allow the compiler's UnusedPruner step to remove most
  /// of the commit-related code from the contract bytecode size.
  function commit(
    bytes32[2] calldata,
    bytes calldata,
    bytes32[] calldata,
    bytes32[] calldata,
    bytes32
  ) external pure override {
    revert OperationNotSupportedByThisOffRampType();
  }

  // ================================================================
  // │                        Static Config                         │
  // ================================================================

  /// @notice Returns the CrossL2Inbox address.
  /// @return crossL2Inbox The address of the CrossL2Inbox.
  function getCrossL2Inbox() external view returns (address) {
    return address(i_crossL2Inbox);
  }

  // ================================================================
  // │                           ChainId                            │
  // ================================================================

  /// @notice Updates chainSelector to chainId mapping.
  /// @param chainSelectorsToRemove Array of selectors to remove from the mapping.
  /// @param chainSelectorsToAdd Array of selector to chainId mappings to add.
  function applyChainSelectorToChainIdConfigUpdates(
    uint64[] memory chainSelectorsToRemove,
    ChainSelectorToChainIdConfigArgs[] memory chainSelectorsToAdd
  ) external onlyOwner {
    _applyChainSelectorToChainIdConfigUpdates(chainSelectorsToRemove, chainSelectorsToAdd);
  }

  /// @notice Internal function to update the chainSelector to chainId mapping.
  /// @param chainSelectorsToRemove Array of selectors to remove from the mapping.
  /// @param chainSelectorsToAdd Array of selector to chainId mappings to add.
  function _applyChainSelectorToChainIdConfigUpdates(
    uint64[] memory chainSelectorsToRemove,
    ChainSelectorToChainIdConfigArgs[] memory chainSelectorsToAdd
  ) internal {
    for (uint256 i = 0; i < chainSelectorsToRemove.length; ++i) {
      uint64 chainSelector = chainSelectorsToRemove[i];
      uint256 chainId = s_chainSelectorToChainId[chainSelector];
      if (chainId != 0) {
        delete s_chainSelectorToChainId[chainSelector];
        emit ChainSelectorToChainIdConfigRemoved(chainSelector, chainId);
      }
    }

    for (uint256 i = 0; i < chainSelectorsToAdd.length; ++i) {
      ChainSelectorToChainIdConfigArgs memory config = chainSelectorsToAdd[i];
      if (config.chainId == 0) {
        revert ZeroChainIdNotAllowed();
      }
      if (config.chainSelector == 0) {
        revert ZeroChainSelectorNotAllowed();
      }
      s_chainSelectorToChainId[config.chainSelector] = config.chainId;
      emit ChainSelectorToChainIdConfigUpdated(config.chainSelector, config.chainId);
    }
  }

  /// @notice Returns the chainId for a given chainSelector.
  /// @param chainSelector The chain selector to get the chainId for.
  /// @return chainId The chainId for the given chainSelector.
  function getChainId(
    uint64 chainSelector
  ) external view returns (uint256) {
    return s_chainSelectorToChainId[chainSelector];
  }
}
