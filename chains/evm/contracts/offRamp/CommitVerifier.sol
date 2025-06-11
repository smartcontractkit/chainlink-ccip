// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRMNRemote} from "../interfaces/IRMNRemote.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Client} from "../libraries/Client.sol";
import {Internal} from "../libraries/Internal.sol";
import {MerkleMultiProof} from "../libraries/MerkleMultiProof.sol";
import {MultiOCR3Base} from "../ocr/MultiOCR3Base.sol";

import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice OffRamp enables OCR networks to execute multiple messages in an OffRamp in a single transaction.
/// @dev The OnRamp and OffRamp form a cross chain upgradeable unit. Any change to one of them results an
/// onchain upgrade of both contracts.
/// @dev MultiOCR3Base is used to store multiple OCR configs for the OffRamp. The execution plugin type has to be
/// configured without signature verification, and the commit plugin type with verification.
contract OffRamp is ITypeAndVersion, MultiOCR3Base {
  using EnumerableSet for EnumerableSet.UintSet;

  error ZeroChainSelectorNotAllowed();
  error SourceChainNotEnabled(uint64 sourceChainSelector);
  error RootAlreadyCommitted(uint64 sourceChainSelector, bytes32 merkleRoot);
  error InvalidRoot();
  error CursedByRMN(uint64 sourceChainSelector);
  error StaleCommitReport();
  error InvalidInterval(uint64 sourceChainSelector, uint64 min, uint64 max);
  error ZeroAddressNotAllowed();
  error SignatureVerificationRequiredInCommitPlugin();
  error SignatureVerificationNotAllowedInExecutionPlugin();
  error CommitOnRampMismatch(bytes reportOnRamp, bytes configOnRamp);
  error InvalidOnRampUpdate(uint64 sourceChainSelector);
  error RootBlessingMismatch(uint64 sourceChainSelector, bytes32 merkleRoot, bool isBlessed);

  /// @dev Atlas depends on various events, if changing, please notify Atlas.
  event StaticConfigSet(StaticConfig staticConfig);
  event DynamicConfigSet(DynamicConfig dynamicConfig);
  event SourceChainSelectorAdded(uint64 sourceChainSelector);
  event SourceChainConfigSet(uint64 indexed sourceChainSelector, SourceChainConfig sourceConfig);
  event CommitReportAccepted(
    Internal.MerkleRoot[] blessedMerkleRoots,
    Internal.MerkleRoot[] unblessedMerkleRoots,
    Internal.PriceUpdates priceUpdates
  );

  /// @dev Struct that contains the static configuration. The individual components are stored as immutable variables.
  // solhint-disable-next-line gas-struct-packing
  struct StaticConfig {
    uint64 chainSelector; // ─╮ Destination chainSelector
    IRMNRemote rmnRemote; // ─╯ RMN Verification Contract
  }

  /// @dev Per-chain source config (defining a lane from a Source Chain -> Dest OffRamp).
  struct SourceChainConfig {
    bool isEnabled; //                  │ Flag whether the source chain is enabled or not.
    uint64 minSeqNr; //                 │ The min sequence number expected for future messages.
    bool isRMNVerificationDisabled; // ─╯ Flag whether the RMN verification is disabled or not.
    bytes onRamp; // OnRamp address on the source chain.
  }

  /// @dev Same as SourceChainConfig but with source chain selector so that an array of these
  /// can be passed in the constructor and the applySourceChainConfigUpdates function.
  struct SourceChainConfigArgs {
    uint64 sourceChainSelector; //      │  Source chain selector of the config to update.
    bool isEnabled; //                  │  Flag whether the source chain is enabled or not.
    bool isRMNVerificationDisabled; // ─╯ Flag whether the RMN verification is disabled or not.
    bytes onRamp; // OnRamp address on the source chain.
  }

  /// @dev Dynamic offRamp config.
  /// @dev Since DynamicConfig is part of DynamicConfigSet event, if changing it, we should update the ABI on Atlas.
  struct DynamicConfig {
    address feeQuoter; // FeeQuoter address on the local chain.
  }

  /// @dev Report that is committed by the observing DON at the committing phase.
  struct CommitReport {
    Internal.PriceUpdates priceUpdates; // List of gas and price updates to commit.
    Internal.MerkleRoot[] merkleRoots; // List of merkle roots from source chains for which RMN is enabled.
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "OffRamp 1.6.0";
  /// @dev Hash of encoded address(0) used for empty address checks.
  bytes32 internal constant EMPTY_ENCODED_ADDRESS_HASH = keccak256(abi.encode(address(0)));
  /// @dev ChainSelector of this chain.
  uint64 internal immutable i_chainSelector;
  /// @dev The RMN verification contract.
  IRMNRemote internal immutable i_rmnRemote;

  // DYNAMIC CONFIG
  DynamicConfig internal s_dynamicConfig;

  /// @notice Set of source chain selectors.
  EnumerableSet.UintSet internal s_sourceChainSelectors;

  /// @notice SourceChainConfig per source chain selector.
  mapping(uint64 sourceChainSelector => SourceChainConfig sourceChainConfig) private s_sourceChainConfigs;

  // STATE
  /// @notice Commit timestamp of merkle roots per source chain.
  mapping(uint64 sourceChainSelector => mapping(bytes32 merkleRoot => uint256 timestamp)) internal s_roots;
  /// @dev The sequence number of the last price update.
  uint64 private s_latestPriceSequenceNumber;

  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    SourceChainConfigArgs[] memory sourceChainConfigs
  ) MultiOCR3Base() {
    if (address(staticConfig.rmnRemote) == address(0)) {
      revert ZeroAddressNotAllowed();
    }

    if (staticConfig.chainSelector == 0) {
      revert ZeroChainSelectorNotAllowed();
    }

    i_chainSelector = staticConfig.chainSelector;
    i_rmnRemote = staticConfig.rmnRemote;
    emit StaticConfigSet(staticConfig);

    _setDynamicConfig(dynamicConfig);
    _applySourceChainConfigUpdates(sourceChainConfigs);
  }

  // ================================================================
  // │                           Commit                             │
  // ================================================================

  /// @notice Commits a single merkle root. The blessing status has to match the source chain config.
  /// @dev An unblessed root means that RMN verification is disabled for the source chain. It does not mean there is
  /// some future point where the root will be blessed.
  /// @param root The merkle root to commit.
  /// @param isBlessed The blessing status of the root.
  function _commitRoot(Internal.MerkleRoot memory root, bool isBlessed) internal {
    uint64 sourceChainSelector = root.sourceChainSelector;

    if (i_rmnRemote.isCursed(bytes16(uint128(sourceChainSelector)))) {
      revert CursedByRMN(sourceChainSelector);
    }

    SourceChainConfig storage sourceChainConfig = _getEnabledSourceChainConfig(sourceChainSelector);

    // If the root is blessed but RMN blessing is disabled for the source chain, or if the root is not blessed but RMN
    // blessing is enabled, we revert.
    if (isBlessed == sourceChainConfig.isRMNVerificationDisabled) {
      revert RootBlessingMismatch(sourceChainSelector, root.merkleRoot, isBlessed);
    }

    if (keccak256(root.onRampAddress) != keccak256(sourceChainConfig.onRamp)) {
      revert CommitOnRampMismatch(root.onRampAddress, sourceChainConfig.onRamp);
    }

    if (sourceChainConfig.minSeqNr != root.minSeqNr || root.minSeqNr > root.maxSeqNr) {
      revert InvalidInterval(sourceChainSelector, root.minSeqNr, root.maxSeqNr);
    }

    bytes32 merkleRoot = root.merkleRoot;
    if (merkleRoot == bytes32(0)) revert InvalidRoot();
    // If we reached this section, the report should contain a valid root.
    // We disallow duplicate roots as that would reset the timestamp and delay potential manual execution.
    if (s_roots[sourceChainSelector][merkleRoot] != 0) {
      revert RootAlreadyCommitted(sourceChainSelector, merkleRoot);
    }

    sourceChainConfig.minSeqNr = root.maxSeqNr + 1;
    s_roots[sourceChainSelector][merkleRoot] = block.timestamp;
  }

  /// @notice Returns the timestamp of a potentially previously committed merkle root.
  /// If the root was never committed 0 will be returned.
  /// @param sourceChainSelector The source chain selector.
  /// @param root The merkle root to check the commit status for.
  /// @return timestamp The timestamp of the committed root or zero in the case that it was never committed.
  function getMerkleRoot(uint64 sourceChainSelector, bytes32 root) external view returns (uint256) {
    return s_roots[sourceChainSelector][root];
  }

  /// @notice Returns timestamp of when root was accepted or 0 if verification fails.
  /// @dev This method uses a merkle tree within a merkle tree, with the hashedLeaves,
  /// proofs and proofFlagBits being used to get the root of the inner tree.
  /// This root is then used as the singular leaf of the outer tree.
  /// @return timestamp The commit timestamp of the root.
  function _verify(
    uint64 sourceChainSelector,
    bytes32[] memory hashedLeaves,
    bytes32[] memory proofs,
    uint256 proofFlagBits
  ) internal view virtual returns (uint256 timestamp) {
    bytes32 root = MerkleMultiProof._merkleRoot(hashedLeaves, proofs, proofFlagBits);
    return s_roots[sourceChainSelector][root];
  }

  /// @inheritdoc MultiOCR3Base
  function _afterOCR3ConfigSet(
    uint8 ocrPluginType
  ) internal view override {
    bool isSignatureVerificationEnabled = s_ocrConfigs[ocrPluginType].configInfo.isSignatureVerificationEnabled;

    if (ocrPluginType == uint8(Internal.OCRPluginType.Commit)) {
      // Signature verification must be enabled for commit plugin.
      if (!isSignatureVerificationEnabled) {
        revert SignatureVerificationRequiredInCommitPlugin();
      }
    } else if (ocrPluginType == uint8(Internal.OCRPluginType.Execution)) {
      // Signature verification must be disabled for execution plugin.
      if (isSignatureVerificationEnabled) {
        revert SignatureVerificationNotAllowedInExecutionPlugin();
      }
    }
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Returns the static config.
  /// @dev This function will always return the same struct as the contents is static and can never change.
  /// @return staticConfig The static config.
  function getStaticConfig() external view returns (StaticConfig memory) {
    return StaticConfig({chainSelector: i_chainSelector, rmnRemote: i_rmnRemote});
  }

  /// @notice Returns the current dynamic config.
  /// @return dynamicConfig The current dynamic config.
  function getDynamicConfig() external view returns (DynamicConfig memory) {
    return s_dynamicConfig;
  }

  /// @notice Returns the source chain config for the provided source chain selector.
  /// @param sourceChainSelector chain to retrieve configuration for.
  /// @return sourceChainConfig The config for the source chain.
  function getSourceChainConfig(
    uint64 sourceChainSelector
  ) external view returns (SourceChainConfig memory) {
    return s_sourceChainConfigs[sourceChainSelector];
  }

  /// @notice Returns all source chain configs.
  /// @return sourceChainConfigs The source chain configs corresponding to all the supported chain selectors.
  function getAllSourceChainConfigs() external view returns (uint64[] memory, SourceChainConfig[] memory) {
    SourceChainConfig[] memory sourceChainConfigs = new SourceChainConfig[](s_sourceChainSelectors.length());
    uint64[] memory sourceChainSelectors = new uint64[](s_sourceChainSelectors.length());
    for (uint256 i = 0; i < s_sourceChainSelectors.length(); ++i) {
      sourceChainSelectors[i] = uint64(s_sourceChainSelectors.at(i));
      sourceChainConfigs[i] = s_sourceChainConfigs[sourceChainSelectors[i]];
    }
    return (sourceChainSelectors, sourceChainConfigs);
  }

  /// @notice Updates source configs.
  /// @param sourceChainConfigUpdates Source chain configs.
  function applySourceChainConfigUpdates(
    SourceChainConfigArgs[] memory sourceChainConfigUpdates
  ) external onlyOwner {
    _applySourceChainConfigUpdates(sourceChainConfigUpdates);
  }

  /// @notice Updates source configs.
  /// @param sourceChainConfigUpdates Source chain configs.
  function _applySourceChainConfigUpdates(
    SourceChainConfigArgs[] memory sourceChainConfigUpdates
  ) internal {
    for (uint256 i = 0; i < sourceChainConfigUpdates.length; ++i) {
      SourceChainConfigArgs memory sourceConfigUpdate = sourceChainConfigUpdates[i];
      uint64 sourceChainSelector = sourceConfigUpdate.sourceChainSelector;

      if (sourceChainSelector == 0) {
        revert ZeroChainSelectorNotAllowed();
      }

      SourceChainConfig storage currentConfig = s_sourceChainConfigs[sourceChainSelector];
      bytes memory newOnRamp = sourceConfigUpdate.onRamp;

      if (currentConfig.onRamp.length == 0) {
        currentConfig.minSeqNr = 1;
        emit SourceChainSelectorAdded(sourceChainSelector);
      } else {
        if (currentConfig.minSeqNr != 1 && keccak256(currentConfig.onRamp) != keccak256(newOnRamp)) {
          // OnRamp updates should only happens due to a misconfiguration.
          // If an OnRamp is misconfigured, no reports should have been committed and no messages should have been
          // executed. This is enforced by the onRamp address check in the commit function.
          revert InvalidOnRampUpdate(sourceChainSelector);
        }
      }

      // OnRamp can never be zero - if it is, then the source chain has been added for the first time.
      if (newOnRamp.length == 0 || keccak256(newOnRamp) == EMPTY_ENCODED_ADDRESS_HASH) {
        revert ZeroAddressNotAllowed();
      }

      currentConfig.onRamp = newOnRamp;
      currentConfig.isEnabled = sourceConfigUpdate.isEnabled;
      currentConfig.isRMNVerificationDisabled = sourceConfigUpdate.isRMNVerificationDisabled;

      // We don't need to check the return value, as inserting the item twice has no effect.
      s_sourceChainSelectors.add(sourceChainSelector);

      emit SourceChainConfigSet(sourceChainSelector, currentConfig);
    }
  }

  /// @notice Sets the dynamic config.
  /// @param dynamicConfig The new dynamic config.
  function setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) external onlyOwner {
    _setDynamicConfig(dynamicConfig);
  }

  /// @notice Sets the dynamic config.
  /// @param dynamicConfig The dynamic config.
  function _setDynamicConfig(
    DynamicConfig memory dynamicConfig
  ) internal {
    if (dynamicConfig.feeQuoter == address(0)) {
      revert ZeroAddressNotAllowed();
    }

    s_dynamicConfig = dynamicConfig;

    emit DynamicConfigSet(dynamicConfig);
  }

  /// @notice Returns a source chain config with a check that the config is enabled.
  /// @param sourceChainSelector Source chain selector to check for cursing.
  /// @return sourceChainConfig The source chain config storage pointer.
  function _getEnabledSourceChainConfig(
    uint64 sourceChainSelector
  ) internal view returns (SourceChainConfig storage) {
    SourceChainConfig storage sourceChainConfig = s_sourceChainConfigs[sourceChainSelector];
    if (!sourceChainConfig.isEnabled) {
      revert SourceChainNotEnabled(sourceChainSelector);
    }

    return sourceChainConfig;
  }

  // ================================================================
  // │                            Access                            │
  // ================================================================

  /// @notice Reverts as this contract should not be able to receive CCIP messages.
  function ccipReceive(
    Client.Any2EVMMessage calldata
  ) external pure {
    // solhint-disable-next-line
    revert();
  }
}
