// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../interfaces/IAny2EVMMessageReceiver.sol";
import {IMessageInterceptor} from "../interfaces/IMessageInterceptor.sol";
import {INonceManager} from "../interfaces/INonceManager.sol";
import {IPoolV1} from "../interfaces/IPool.sol";

import {IRMNRemote} from "../interfaces/IRMNRemote.sol";
import {IRouter} from "../interfaces/IRouter.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";

import {Client} from "../libraries/Client.sol";
import {Internal} from "../libraries/Internal.sol";
import {Pool} from "../libraries/Pool.sol";
import {SuperchainInterop} from "../libraries/SuperchainInterop.sol";

import {ICrossL2Inbox} from "../vendor/optimism/interop-lib/v0/src/interfaces/ICrossL2Inbox.sol";
import {Identifier} from "../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";

import {OffRamp} from "./OffRamp.sol";

/// @notice This OffRamp supports OP Superchain Interop. It leverages CrossL2Inbox to verify validity of a message,
/// as opposed to relying on roots signed by oracles.
contract OffRampOverSuperchainInterop is OffRamp {
  error InvalidSourceChainSelector(uint64 sourceChainSelector, uint64 expected);
  error InvalidDestChainSelector(uint64 destChainSelector, uint64 expected);
  error InvalidSourceOnRamp(address sourceOnRamp);
  error ZeroChainIdNotAllowed();
  error ChainIdNotConfiguredForSelector(uint64 sourceChainSelector);
  error OperationNotSupportedbyThisOffRampType();
  error InvalidMessageCountInReport(uint256 numMessages, uint256 expected);
  error InvalidProofsWordLength(uint256 length, uint256 expected);

  event ChainSelectorToChainIdConfigUpdated(uint64 indexed chainSelector, uint256 indexed chainId);
  event ChainSelectorToChainIdConfigRemoved(uint64 indexed chainSelector, uint256 indexed chainId);

  struct ChainSelectorToChainIdConfigArgs {
    uint64 chainSelector;
    uint256 chainId;
  }

  /// @dev The CrossL2Inbox interface.
  ICrossL2Inbox internal immutable i_crossL2Inbox;

  mapping(uint64 sourceChainSelector => uint256 chainId) private s_sourceChainSelectorToChainId;

  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    SourceChainConfigArgs[] memory sourceChainConfigs,
    address crossL2Inbox,
    ChainSelectorToChainIdConfigArgs[] memory chainSelectorToChainIdConfigArgs
  ) OffRamp(staticConfig, dynamicConfig, sourceChainConfigs) {
    i_crossL2Inbox = ICrossL2Inbox(crossL2Inbox);

    _applyChainSelectorToChainIdConfigUpdates(new uint64[](0), chainSelectorToChainIdConfigArgs);
  }

  /// @notice Constructs the Identifier from the proofData, then log hash from the message.
  /// @param message The message to construct proofData and log hash.
  /// @param proofs Identifier split into 5 words, each representing a field in order.
  /// @return identifier The identifier constructed from the proofs.
  /// @return logHash The log hash constructed from the message.
  function _constructProofs(
    Internal.Any2EVMRampMessage memory message,
    bytes32[] memory proofs
  ) internal pure returns (Identifier memory identifier, bytes32 logHash) {
    if (proofs.length != 5) revert InvalidProofsWordLength(proofs.length, 5);

    identifier = Identifier({
      origin: address(uint160(uint256(proofs[0]))),
      blockNumber: uint256(proofs[1]),
      logIndex: uint256(proofs[2]),
      timestamp: uint256(proofs[3]),
      chainId: uint256(proofs[4])
    });

    logHash = keccak256(
      abi.encode(
        SuperchainInterop.SENT_MESSAGE_LOG_SELECTOR,
        message.header.destChainSelector,
        message.header.sequenceNumber,
        message
      )
    );

    return (identifier, logHash);
  }

  function _verifyMessage(
    uint64 sourceChainSelector,
    Internal.ExecutionReport memory report
  ) internal virtual override returns (uint256 timestampCommitted, bytes32[] memory hashedLeaves) {
    if (report.messages.length != 1) revert InvalidMessageCountInReport(report.messages.length, 1);
    Internal.Any2EVMRampMessage memory message = report.messages[0];

    // Validate that the message is meant for this chain.
    if (message.header.sourceChainSelector != sourceChainSelector) {
      revert InvalidSourceChainSelector(message.header.sourceChainSelector, sourceChainSelector);
    }
    if (message.header.destChainSelector != i_chainSelector) {
      revert InvalidDestChainSelector(message.header.destChainSelector, i_chainSelector);
    }

    (Identifier memory identifier, bytes32 logHash) = _constructProofs(message, report.proofs);
    address onRampAddress = abi.decode(_getEnabledSourceChainConfig(sourceChainSelector).onRamp, (address));

    // Validate that the message was emitted by the corresponding OnRamp.
    if (identifier.origin != onRampAddress) {
      revert InvalidSourceOnRamp(identifier.origin);
    }
    // Validate that the chainId maps to the expected sourceChainSelector
    // Scope to reduce stack depth.
    uint256 expectedChainId = s_sourceChainSelectorToChainId[sourceChainSelector];
    if (expectedChainId == 0) {
      revert ChainIdNotConfiguredForSelector(sourceChainSelector);
    }
    if (expectedChainId != identifier.chainId) {
      revert SourceChainSelectorMismatch(sourceChainSelector, sourceChainSelector);
    }

    // SECURITY CRITICAL CHECK.
    // Validate the exact log was emitted on the source chain.
    i_crossL2Inbox.validateMessage(identifier, logHash);

    hashedLeaves = new bytes32[](1);
    hashedLeaves[0] = SuperchainInterop._hashInteropMessage(message, onRampAddress);

    // Non-zero timestamp signals the message is verified.
    // Because there is no Commit timestamp, the timestmap of the message is taken from the time it is sent.
    // If this OffRamp only accepts low-latency messages from source chains within OP Mesh,
    // this will be very close to the commit timestamp.
    // If this OffRamp can accept messages from high-latency sources,
    // `permissionLessExecutionThresholdSeconds` needs to be adjusted..
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
  ) external override {
    revert OperationNotSupportedbyThisOffRampType();
  }

  // ================================================================
  // │                        Custom Configs                        │
  // ================================================================

  /// @notice Updates sourceChainSelector to chainId mapping.
  /// @param chainSelectorsToUnset Array of chainIds to remove from the mapping.
  /// @param chainSelectorsToSet Array of chainId to sourceChainSelector mappings to add.
  function applyChainSelectorToChainIdConfigUpdates(
    uint64[] memory chainSelectorsToUnset,
    ChainSelectorToChainIdConfigArgs[] memory chainSelectorsToSet
  ) external onlyOwner {
    _applyChainSelectorToChainIdConfigUpdates(chainSelectorsToUnset, chainSelectorsToSet);
  }

  /// @notice Internal function to update the sourceChainSelector tp chainId mapping.
  /// @param chainSelectorsToUnset Array of chain selector to remove.
  /// @param chainSelectorsToSet Array of chain selectors to add.
  function _applyChainSelectorToChainIdConfigUpdates(
    uint64[] memory chainSelectorsToUnset,
    ChainSelectorToChainIdConfigArgs[] memory chainSelectorsToSet
  ) internal {
    for (uint256 i = 0; i < chainSelectorsToUnset.length; ++i) {
      uint64 chainSelector = chainSelectorsToUnset[i];
      uint256 chainId = s_sourceChainSelectorToChainId[chainSelector];
      delete s_sourceChainSelectorToChainId[chainSelector];
      if (chainId != 0) {
        emit ChainSelectorToChainIdConfigRemoved(chainSelector, chainId);
      }
    }

    for (uint256 i = 0; i < chainSelectorsToSet.length; ++i) {
      ChainSelectorToChainIdConfigArgs memory config = chainSelectorsToSet[i];
      if (config.chainId == 0) {
        revert ZeroChainIdNotAllowed();
      }
      if (config.chainSelector == 0) {
        revert ZeroChainSelectorNotAllowed();
      }
      s_sourceChainSelectorToChainId[config.chainSelector] = config.chainId;
      emit ChainSelectorToChainIdConfigUpdated(config.chainSelector, config.chainId);
    }
  }

  /// @notice Returns the chainId for a given sourceChainSelector.
  function getChainId(
    uint64 sourceChainSelector
  ) external view returns (uint256) {
    return s_sourceChainSelectorToChainId[sourceChainSelector];
  }

  /// @notice Returns the CrossL2Inbox address.
  function getCrossL2Inbox() external view returns (address) {
    return address(i_crossL2Inbox);
  }
}
