// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "../interfaces/IPolicyEngine.sol";
import {IPoolV2} from "../interfaces/IPoolV2.sol";

import {CCIPPolicyEnginePayloads} from "../libraries/CCIPPolicyEnginePayloads.sol";
import {CCVConfigValidation} from "../libraries/CCVConfigValidation.sol";
import {Pool} from "../libraries/Pool.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Advanced pool hooks for additional security features like allowlists, CCV management, and policy engine integration.
/// @dev This is a standalone contract that can optionally be used by TokenPools.
contract AdvancedPoolHooks is IAdvancedPoolHooks, AuthorizedCallers {
  using EnumerableSet for EnumerableSet.AddressSet;

  error AllowListNotEnabled();
  error SenderNotAllowed(address sender);
  error MustSpecifyUnderThresholdCCVsForThresholdCCVs();
  error PolicyEngineDetachFailed(address oldPolicyEngine, bytes err);

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
  event PolicyEngineSet(address indexed oldPolicyEngine, address indexed newPolicyEngine);
  event AllowAnyoneToInvokeThisHookSet(bool allowed);

  // bytes4(keccak256("PoolHookOutboundPolicyDataV1"))
  bytes4 internal constant POOL_HOOK_OUTBOUND_POLICY_DATA_V1_TAG = 0x12bebcb8;

  // bytes4(keccak256("PoolHookInboundPolicyDataV1"))
  bytes4 internal constant POOL_HOOK_INBOUND_POLICY_DATA_V1_TAG = 0x44d1de78;

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

  /// @dev A set of addresses allowed to trigger lockOrBurn as original senders.
  /// Only takes effect if i_allowlistEnabled is true.
  /// This can be used to ensure only token-issuer specified addresses can move tokens.
  EnumerableSet.AddressSet internal s_allowlist;

  /// @dev Threshold token transfer amount at which additional CCVs are required.
  /// Value of 0 means that there is no threshold and additional CCVs are not required for any transfer amount.
  uint256 internal s_thresholdAmountForAdditionalCCVs;

  /// @dev The policy engine to use for additional validation. If set to address(0), no policy engine will be used.
  IPolicyEngine internal s_policyEngine;

  /// @dev When true, anyone can call preflightCheck/postFlightCheck. When false, only authorized callers can call.
  bool internal s_allowAnyoneToInvokeThisHook;

  /// @dev Stores verifier (CCV) requirements keyed by remote chain selector.
  mapping(uint64 remoteChainSelector => CCVConfig ccvConfig) internal s_verifierConfig;

  constructor(
    address[] memory allowlist,
    uint256 thresholdAmountForAdditionalCCVs,
    address policyEngine,
    address[] memory authorizedCallers,
    bool allowAnyoneToInvokeThisHook
  ) AuthorizedCallers(authorizedCallers) {
    // Allowlist can be set as enabled or disabled at deployment time only to save hot-path gas.
    i_allowlistEnabled = allowlist.length > 0;
    if (i_allowlistEnabled) {
      _applyAllowListUpdates(new address[](0), allowlist);
    }
    s_thresholdAmountForAdditionalCCVs = thresholdAmountForAdditionalCCVs;
    _setPolicyEngine(policyEngine, false);
    s_allowAnyoneToInvokeThisHook = allowAnyoneToInvokeThisHook;
  }

  /// @inheritdoc IAdvancedPoolHooks
  /// @dev Performs allowlist check and policy engine validation for outbound transfers.
  function preflightCheck(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  ) external {
    if (!s_allowAnyoneToInvokeThisHook) {
      _validateCaller();
    }
    checkAllowList(lockOrBurnIn.originalSender);

    IPolicyEngine policyEngine = s_policyEngine;
    if (address(policyEngine) == address(0)) {
      return;
    }

    CCIPPolicyEnginePayloads.PoolHookOutboundPolicyDataV1 memory outboundData = CCIPPolicyEnginePayloads.PoolHookOutboundPolicyDataV1({
      originalSender: lockOrBurnIn.originalSender,
      blockConfirmationRequested: blockConfirmationRequested,
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      receiver: lockOrBurnIn.receiver,
      amount: lockOrBurnIn.amount,
      localToken: lockOrBurnIn.localToken,
      tokenArgs: tokenArgs
    });
    bytes memory policyData = abi.encodeWithSelector(POOL_HOOK_OUTBOUND_POLICY_DATA_V1_TAG, outboundData);
    policyEngine.run(
      IPolicyEngine.Payload({
        selector: IAdvancedPoolHooks.preflightCheck.selector,
        sender: msg.sender,
        data: policyData,
        context: ""
      })
    );
  }

  /// @inheritdoc IAdvancedPoolHooks
  /// @dev Performs policy engine validation for inbound transfers.
  function postFlightCheck(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint256 localAmount,
    uint16 blockConfirmationRequested
  ) external {
    if (!s_allowAnyoneToInvokeThisHook) {
      _validateCaller();
    }

    IPolicyEngine policyEngine = s_policyEngine;
    if (address(policyEngine) == address(0)) {
      return;
    }

    CCIPPolicyEnginePayloads.PoolHookInboundPolicyDataV1 memory inboundData = CCIPPolicyEnginePayloads.PoolHookInboundPolicyDataV1({
      originalSender: releaseOrMintIn.originalSender,
      blockConfirmationRequested: blockConfirmationRequested,
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      receiver: releaseOrMintIn.receiver,
      amount: releaseOrMintIn.sourceDenominatedAmount,
      localToken: releaseOrMintIn.localToken,
      sourcePoolAddress: releaseOrMintIn.sourcePoolAddress,
      sourcePoolData: releaseOrMintIn.sourcePoolData,
      offchainTokenData: releaseOrMintIn.offchainTokenData,
      localAmount: localAmount
    });
    bytes memory policyData = abi.encodeWithSelector(POOL_HOOK_INBOUND_POLICY_DATA_V1_TAG, inboundData);
    policyEngine.run(
      IPolicyEngine.Payload({
        selector: IAdvancedPoolHooks.postFlightCheck.selector,
        sender: msg.sender,
        data: policyData,
        context: ""
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
  function applyAllowListUpdates(address[] calldata removes, address[] calldata adds) external onlyOwner {
    _applyAllowListUpdates(removes, adds);
  }

  /// @notice Internal version of applyAllowListUpdates to allow for reuse in the constructor.
  /// @param removes The addresses to be removed.
  /// @param adds The addresses to be added.
  function _applyAllowListUpdates(address[] memory removes, address[] memory adds) internal {
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
  /// @dev Use this when the old policy engine is unresponsive or has a bug in its detach() implementation.
  /// @param newPolicyEngine The address of the new policy engine.
  function setPolicyEngineAllowFailedDetach(
    address newPolicyEngine
  ) external onlyOwner {
    _setPolicyEngine(newPolicyEngine, true);
  }

  /// @notice Internal function to set and attach to a policy engine.
  /// @param newPolicyEngine The address of the new policy engine, or address(0) to disable.
  /// @param allowFailedDetach Whether to revert if old policy engine's detach reverts.
  function _setPolicyEngine(address newPolicyEngine, bool allowFailedDetach) internal {
    address oldPolicyEngine = address(s_policyEngine);

    if (newPolicyEngine == oldPolicyEngine) {
      return;
    }

    if (oldPolicyEngine != address(0)) {
      try IPolicyEngine(oldPolicyEngine).detach() {}
      catch (bytes memory err) {
        if (!allowFailedDetach) {
          revert PolicyEngineDetachFailed(oldPolicyEngine, err);
        }
      }
    }

    s_policyEngine = IPolicyEngine(newPolicyEngine);
    if (newPolicyEngine != address(0)) {
      IPolicyEngine(newPolicyEngine).attach();
    }

    emit PolicyEngineSet(oldPolicyEngine, newPolicyEngine);
  }

  /// @notice Gets the current policy engine address.
  /// @return The address of the policy engine, address(0) if none is set.
  function getPolicyEngine() external view returns (address) {
    return address(s_policyEngine);
  }

  // ================================================================
  // │                     Authorized Callers                       │
  // ================================================================

  /// @notice Gets whether anyone can invoke preflightCheck/postFlightCheck.
  /// @return true if anyone can call, false if only authorized callers can call.
  function getAllowAnyoneToInvokeThisHook() external view returns (bool) {
    return s_allowAnyoneToInvokeThisHook;
  }

  /// @notice Sets whether anyone can invoke preflightCheck/postFlightCheck.
  /// @param allowed When true, anyone can call preflightCheck/postFlightCheck. When false, only authorized callers.
  function setAllowAnyoneToInvokeThisHook(
    bool allowed
  ) external onlyOwner {
    s_allowAnyoneToInvokeThisHook = allowed;
    emit AllowAnyoneToInvokeThisHookSet(allowed);
  }
}
