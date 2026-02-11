// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../interfaces/IAdvancedPoolHooks.sol";
import {IPoolV2} from "../interfaces/IPoolV2.sol";
import {IPolicyEngine} from "@chainlink/ace/policy-management/interfaces/IPolicyEngine.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {CCVConfigValidation} from "../libraries/CCVConfigValidation.sol";
import {Pool} from "../libraries/Pool.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Advanced pool hooks for additional security features like allowlists, CCV management, and policy engine runs.
/// @dev This is a standalone contract that can optionally be used by TokenPools.
contract AdvancedPoolHooks is IAdvancedPoolHooks, ITypeAndVersion, AuthorizedCallers {
  using EnumerableSet for EnumerableSet.AddressSet;

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "AdvancedPoolHooks 2.0.0-dev";
  }

  error AllowListNotEnabled();
  error AuthorizedCallersNotEnabled();
  error SenderNotAllowed(address sender);
  error MustSpecifyUnderThresholdCCVsForThresholdCCVs();
  error PolicyEngineDetachReverted(address oldPolicyEngine, bytes err);

  event AllowListAdd(address sender);
  event AllowListRemove(address sender);
  event CCVConfigUpdated(
    uint64 indexed remoteChainSelector,
    address[] outboundCCVs,
    address[] thresholdOutboundCCVs,
    address[] inboundCCVs,
    address[] thresholdInboundCCVs
  );
  event ThresholdAmountSet(uint256 thresholdAmount);
  event PolicyEngineAttached(address indexed policyEngine);
  event PolicyEngineDetachFailed(address indexed policyEngine, bytes reason);

  struct CCVConfig {
    address[] outboundCCVs; // CCVs required for outgoing messages to the remote chain.
    address[] thresholdOutboundCCVs; // Additional CCVs that are required for outgoing messages after reaching the threshold amount.
    address[] inboundCCVs; // CCVs required for incoming messages from the remote chain.
    address[] thresholdInboundCCVs; // Additional CCVs that are required for incoming messages after reaching the threshold amount.
  }

  struct CCVConfigArg {
    uint64 remoteChainSelector;
    address[] outboundCCVs;
    address[] thresholdOutboundCCVs;
    address[] inboundCCVs;
    address[] thresholdInboundCCVs;
  }

  /// @dev The immutable flag that indicates if the allowlist is access-controlled.
  bool internal immutable i_allowlistEnabled;

  /// @dev The immutable flag that indicates if preflightCheck/postflightCheck are access-controlled.
  bool internal immutable i_authorizedCallersEnabled;

  /// @dev A set of addresses allowed to trigger lockOrBurn as original senders.
  /// Only takes effect if i_allowlistEnabled is true.
  /// This can be used to ensure only token-issuer specified addresses can move tokens.
  EnumerableSet.AddressSet internal s_allowlist;

  /// @dev Threshold token transfer amount at which additional CCVs are required.
  /// Value of 0 means that there is no threshold and additional CCVs are not required for any transfer amount.
  uint256 internal s_thresholdAmountForAdditionalCCVs;

  /// @dev The policy engine to use. Value of 0 disables policy engine checks.
  IPolicyEngine internal s_policyEngine;

  /// @dev Stores verifier (CCV) requirements keyed by remote chain selector.
  mapping(uint64 remoteChainSelector => CCVConfig ccvConfig) internal s_verifierConfig;

  constructor(
    address[] memory allowlist,
    uint256 thresholdAmountForAdditionalCCVs,
    address policyEngine,
    address[] memory authorizedCallers
  ) AuthorizedCallers(authorizedCallers) {
    // Allowlist can be set as enabled or disabled at deployment time only to save hot-path gas.
    i_allowlistEnabled = allowlist.length > 0;
    if (i_allowlistEnabled) {
      _applyAllowListUpdates(new address[](0), allowlist);
    }
    s_thresholdAmountForAdditionalCCVs = thresholdAmountForAdditionalCCVs;

    i_authorizedCallersEnabled = authorizedCallers.length > 0;
    _setPolicyEngine(policyEngine, false);
  }

  /// @inheritdoc IAdvancedPoolHooks
  /// @dev Performs allowlist check and policy engine validation for outbound transfers.
  function preflightCheck(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16,
    bytes calldata tokenArgs,
    uint256
  ) external {
    validateCaller();
    checkAllowList(lockOrBurnIn.originalSender);

    IPolicyEngine policyEngine = s_policyEngine;
    if (address(policyEngine) == address(0)) {
      return;
    }

    policyEngine.run(
      IPolicyEngine.Payload({selector: msg.sig, sender: msg.sender, data: msg.data[4:], context: tokenArgs})
    );
  }

  /// @inheritdoc IAdvancedPoolHooks
  /// @dev Performs policy engine validation for inbound transfers.
  function postflightCheck(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint256,
    uint16
  ) external {
    validateCaller();

    IPolicyEngine policyEngine = s_policyEngine;
    if (address(policyEngine) == address(0)) {
      return;
    }

    policyEngine.run(
      IPolicyEngine.Payload({
        selector: msg.sig, sender: msg.sender, data: msg.data[4:], context: releaseOrMintIn.offchainTokenData
      })
    );
  }

  // ================================================================
  // │                          Allowlist                           │
  // ================================================================

  /// @notice Checks if the sender is allowed to perform an operation.
  /// @param sender The address to check.
  function checkAllowList(
    address sender
  ) public view {
    if (i_allowlistEnabled) {
      if (!s_allowlist.contains(sender)) {
        revert SenderNotAllowed(sender);
      }
    }
  }

  /// @notice Gets whether the allowlist functionality is enabled.
  /// @return true is enabled, false if not.
  function getAllowListEnabled() external view returns (bool) {
    return i_allowlistEnabled;
  }

  /// @notice Gets the allowed addresses.
  /// @return The allowed addresses.
  function getAllowList() external view returns (address[] memory) {
    return s_allowlist.values();
  }

  /// @notice Apply updates to the allow list.
  /// @param removes The addresses to be removed.
  /// @param adds The addresses to be added.
  function applyAllowListUpdates(
    address[] calldata removes,
    address[] calldata adds
  ) external onlyOwner {
    _applyAllowListUpdates(removes, adds);
  }

  /// @notice Internal version of applyAllowListUpdates to allow for reuse in the constructor.
  /// @param removes The addresses to be removed.
  /// @param adds The addresses to be added.
  function _applyAllowListUpdates(
    address[] memory removes,
    address[] memory adds
  ) internal {
    if (!i_allowlistEnabled) revert AllowListNotEnabled();

    for (uint256 i = 0; i < removes.length; ++i) {
      address toRemove = removes[i];
      if (s_allowlist.remove(toRemove)) {
        emit AllowListRemove(toRemove);
      }
    }
    for (uint256 i = 0; i < adds.length; ++i) {
      address toAdd = adds[i];
      if (toAdd == address(0)) {
        continue;
      }
      if (s_allowlist.add(toAdd)) {
        emit AllowListAdd(toAdd);
      }
    }
  }

  // ================================================================
  // │                          CCV                                 │
  // ================================================================

  /// @notice Updates the CCV configuration for specified remote chains.
  /// If the array includes address(0), it indicates that the default CCV should be used alongside any other specified CCVs.
  /// @dev Additional CCVs should only be configured for transfers at or above the threshold amount and should not duplicate base CCVs.
  /// Base CCVs are always required, while add-above-threshold CCVs are only required when the transfer amount exceeds the threshold.
  /// @param ccvConfigArgs The CCV configuration updates to apply.
  function applyCCVConfigUpdates(
    CCVConfigArg[] calldata ccvConfigArgs
  ) external onlyOwner {
    for (uint256 i = 0; i < ccvConfigArgs.length; ++i) {
      uint64 remoteChainSelector = ccvConfigArgs[i].remoteChainSelector;
      address[] calldata outboundCCVs = ccvConfigArgs[i].outboundCCVs;
      address[] calldata thresholdOutboundCCVs = ccvConfigArgs[i].thresholdOutboundCCVs;
      address[] calldata inboundCCVs = ccvConfigArgs[i].inboundCCVs;
      address[] calldata thresholdInboundCCVs = ccvConfigArgs[i].thresholdInboundCCVs;

      // Check for duplicates in outbound CCVs.
      CCVConfigValidation._assertNoDuplicates(outboundCCVs);

      // Check for duplicates in inbound CCVs.
      CCVConfigValidation._assertNoDuplicates(inboundCCVs);

      if (thresholdOutboundCCVs.length > 0) {
        // Must have base CCVs if specifying above-threshold CCVs. If the defaults are used below the threshold,
        // specify address(0) in the outboundCCVs array.
        if (outboundCCVs.length == 0) {
          revert MustSpecifyUnderThresholdCCVsForThresholdCCVs();
        }

        CCVConfigValidation._assertNoDuplicates(thresholdOutboundCCVs);
        CCVConfigValidation._assertNoDuplicatedBetweenLists(outboundCCVs, thresholdOutboundCCVs);
      }

      if (thresholdInboundCCVs.length > 0) {
        // Must have base CCVs if specifying above-threshold CCVs. If the defaults are used below the threshold,
        // specify address(0) in the inboundCCVs array.
        if (inboundCCVs.length == 0) {
          revert MustSpecifyUnderThresholdCCVsForThresholdCCVs();
        }

        CCVConfigValidation._assertNoDuplicates(thresholdInboundCCVs);
        CCVConfigValidation._assertNoDuplicatedBetweenLists(inboundCCVs, thresholdInboundCCVs);
      }

      s_verifierConfig[remoteChainSelector] = CCVConfig({
        outboundCCVs: outboundCCVs,
        thresholdOutboundCCVs: thresholdOutboundCCVs,
        inboundCCVs: inboundCCVs,
        thresholdInboundCCVs: thresholdInboundCCVs
      });
      emit CCVConfigUpdated({
        remoteChainSelector: remoteChainSelector,
        outboundCCVs: outboundCCVs,
        thresholdOutboundCCVs: thresholdOutboundCCVs,
        inboundCCVs: inboundCCVs,
        thresholdInboundCCVs: thresholdInboundCCVs
      });
    }
  }

  /// @notice Returns the set of required CCVs for transfers in a specific direction.
  /// @param remoteChainSelector The remote chain selector for this transfer.
  /// @param amount The amount being transferred.
  /// @param direction The direction of the transfer (Inbound or Outbound).
  /// This implementation returns base CCVs for all transfers, and includes additional CCVs when the transfer amount
  /// is above the configured threshold.
  /// @return requiredCCVs Set of required CCV addresses.
  function getRequiredCCVs(
    address,
    uint64 remoteChainSelector,
    uint256 amount,
    uint16,
    bytes calldata,
    IPoolV2.MessageDirection direction
  ) external view returns (address[] memory requiredCCVs) {
    CCVConfig storage config = s_verifierConfig[remoteChainSelector];
    if (direction == IPoolV2.MessageDirection.Inbound) {
      return _resolveRequiredCCVs(config.inboundCCVs, config.thresholdInboundCCVs, amount);
    }
    return _resolveRequiredCCVs(config.outboundCCVs, config.thresholdOutboundCCVs, amount);
  }

  /// @notice Gets the threshold amount above which additional CCVs are required.
  /// @return The threshold amount.
  function getThresholdAmount() external view returns (uint256) {
    return s_thresholdAmountForAdditionalCCVs;
  }

  /// @notice Sets the threshold amount above which additional CCVs are required.
  /// @param thresholdAmount The new threshold amount.
  function setThresholdAmount(
    uint256 thresholdAmount
  ) external onlyOwner {
    s_thresholdAmountForAdditionalCCVs = thresholdAmount;

    emit ThresholdAmountSet(thresholdAmount);
  }

  function _resolveRequiredCCVs(
    address[] memory baseCCVs,
    address[] storage requiredCCVsAboveThresholdStorage,
    uint256 amount
  ) internal view returns (address[] memory requiredCCVs) {
    // If amount is above threshold, combine base and additional CCVs.
    uint256 thresholdAmount = s_thresholdAmountForAdditionalCCVs;
    if (thresholdAmount != 0 && amount >= thresholdAmount) {
      address[] memory thresholdCCVs = requiredCCVsAboveThresholdStorage;
      if (thresholdCCVs.length > 0) {
        requiredCCVs = new address[](baseCCVs.length + thresholdCCVs.length);
        // Copy base CCVs.
        for (uint256 i = 0; i < baseCCVs.length; ++i) {
          requiredCCVs[i] = baseCCVs[i];
        }
        // Copy additional CCVs.
        for (uint256 i = 0; i < thresholdCCVs.length; ++i) {
          requiredCCVs[baseCCVs.length + i] = thresholdCCVs[i];
        }
        return requiredCCVs;
      }
    }
    return baseCCVs;
  }

  // ================================================================
  // │                       Policy Engine                          │
  // ================================================================

  /// @notice Sets a new policy engine.
  /// @param newPolicyEngine The address of the new policy engine.
  function setPolicyEngine(
    address newPolicyEngine
  ) external onlyOwner {
    _setPolicyEngine(newPolicyEngine, false);
  }

  /// @notice Sets a new policy engine while tolerating a pre-existing policy engine's detach reverting.
  /// @dev Use this to force update an old policy engine whose detach() reverts.
  /// @param newPolicyEngine The address of the new policy engine.
  function setPolicyEngineAllowFailedDetach(
    address newPolicyEngine
  ) external onlyOwner {
    _setPolicyEngine(newPolicyEngine, true);
  }

  /// @notice Internal function to set and attach to a policy engine.
  /// @param newPolicyEngine The address of the new policy engine, or address(0) to disable.
  /// @param allowFailedDetach Whether to revert if old policy engine's detach reverts.
  function _setPolicyEngine(
    address newPolicyEngine,
    bool allowFailedDetach
  ) internal {
    address oldPolicyEngine = address(s_policyEngine);

    if (newPolicyEngine == oldPolicyEngine) {
      return;
    }

    if (oldPolicyEngine != address(0)) {
      // Guarding detach reverts to offer escape hatch from adversarial policy engine instances.
      try IPolicyEngine(oldPolicyEngine).detach() {}
      catch (bytes memory err) {
        if (!allowFailedDetach) {
          revert PolicyEngineDetachReverted(oldPolicyEngine, err);
        }
        emit PolicyEngineDetachFailed(oldPolicyEngine, err);
      }
    }

    s_policyEngine = IPolicyEngine(newPolicyEngine);
    if (newPolicyEngine != address(0)) {
      IPolicyEngine(newPolicyEngine).attach();
    }

    emit PolicyEngineAttached(newPolicyEngine);
  }

  /// @notice Gets the current policy engine address.
  /// @return The address of the policy engine.
  function getPolicyEngine() external view returns (address) {
    return address(s_policyEngine);
  }

  // ================================================================
  // │                     Authorized Callers                       │
  // ================================================================

  /// @notice Checks the sender and reverts if it is anyone other than a listed authorized caller.
  function validateCaller() public view virtual {
    if (i_authorizedCallersEnabled) {
      _validateCaller();
    }
  }

  /// @notice Gets whether only authorized callers can invoke preflightCheck/postflightCheck.
  /// @return true if only authorized callers can call, false if anyone can call.
  function getAuthorizedCallersEnabled() external view returns (bool) {
    return i_authorizedCallersEnabled;
  }

  /// @notice Updates the list of authorized callers.
  /// @param authorizedCallerArgs Callers to add and remove. Removals are performed first.
  function applyAuthorizedCallerUpdates(
    AuthorizedCallerArgs memory authorizedCallerArgs
  ) external virtual override onlyOwner {
    if (!i_authorizedCallersEnabled) revert AuthorizedCallersNotEnabled();

    _applyAuthorizedCallerUpdates(authorizedCallerArgs);
  }
}
