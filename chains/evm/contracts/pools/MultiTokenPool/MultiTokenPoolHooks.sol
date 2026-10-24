// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IMultiTokenPoolHooks} from "../../interfaces/IMultiTokenPoolHooks.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {Pool} from "../../libraries/Pool.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Pool hooks for multi-token pools, with a sender allowlist and policy engine runs.
/// @dev This is AdvancedPoolHooks without the CCV config, as multi-token pools get their CCVs from the pool config
/// manager. One instance can serve many tokens and lanes, and they all share its allowlist and policy engine. Every
/// pool that calls these hooks must be an authorized caller.
contract MultiTokenPoolHooks is IMultiTokenPoolHooks, ITypeAndVersion, AuthorizedCallers {
  using EnumerableSet for EnumerableSet.AddressSet;

  error AllowListNotEnabled();
  error SenderNotAllowed(address sender);
  error PolicyEngineDetachReverted(IPolicyEngine oldPolicyEngine, bytes err);

  event AllowListAdd(address sender);
  event AllowListRemove(address sender);
  event PolicyEngineAttached(IPolicyEngine indexed policyEngine);
  event PolicyEngineDetachFailed(IPolicyEngine indexed policyEngine, bytes reason);

  /// @dev The immutable flag that indicates if the allowlist is access-controlled.
  bool private immutable i_allowlistEnabled;

  /// @dev A set of addresses allowed to trigger lockOrBurn as original senders.
  /// Only takes effect if i_allowlistEnabled is true.
  EnumerableSet.AddressSet private s_allowlist;

  /// @dev The policy engine to use. Value of 0 disables policy engine checks.
  IPolicyEngine private s_policyEngine;

  /// @param allowlist The addresses allowed to send tokens. An empty list disables the allowlist permanently.
  /// @param policyEngine The policy engine to attach. Zero disables policy engine checks.
  /// @param authorizedCallers The pools allowed to call the hooks.
  constructor(
    address[] memory allowlist,
    IPolicyEngine policyEngine,
    address[] memory authorizedCallers
  ) AuthorizedCallers(authorizedCallers) {
    // Allowlist can be set as enabled or disabled at deployment time only to save hot-path gas.
    i_allowlistEnabled = allowlist.length > 0;
    if (i_allowlistEnabled) {
      _applyAllowListUpdates(new address[](0), allowlist);
    }
    _setPolicyEngine(policyEngine, false);
  }

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "MultiTokenPoolHooks 2.1.0-dev";
  }

  /// @inheritdoc IMultiTokenPoolHooks
  /// @dev Performs allowlist check and policy engine validation for outbound transfers.
  function preflightCheck(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    bytes4, // requestedFinalityConfig
    bytes calldata tokenArgs,
    uint256 // amountPostFee
  ) public virtual {
    _validateCaller();
    checkAllowList(lockOrBurnIn.originalSender);

    IPolicyEngine policyEngine = s_policyEngine;
    if (address(policyEngine) == address(0)) return;

    policyEngine.run(
      IPolicyEngine.Payload({selector: msg.sig, sender: msg.sender, data: msg.data[4:], context: tokenArgs})
    );
  }

  /// @inheritdoc IMultiTokenPoolHooks
  /// @dev Performs policy engine validation for inbound transfers.
  function postflightCheck(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint256, // localAmount
    bytes4 // requestedFinalityConfig
  ) public virtual {
    _validateCaller();

    IPolicyEngine policyEngine = s_policyEngine;
    if (address(policyEngine) == address(0)) return;

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
  ) public view virtual {
    if (i_allowlistEnabled && !s_allowlist.contains(sender)) revert SenderNotAllowed(sender);
  }

  /// @notice Gets whether the allowlist functionality is enabled.
  /// @return allowlistEnabled True if enabled, false if not.
  function getAllowListEnabled() public view virtual returns (bool allowlistEnabled) {
    return i_allowlistEnabled;
  }

  /// @notice Gets the allowed addresses.
  /// @return allowlist The allowed addresses.
  function getAllowList() public view virtual returns (address[] memory allowlist) {
    return s_allowlist.values();
  }

  /// @notice Apply updates to the allow list.
  /// @param removes The addresses to be removed.
  /// @param adds The addresses to be added.
  function applyAllowListUpdates(
    address[] calldata removes,
    address[] calldata adds
  ) public virtual onlyOwner {
    _applyAllowListUpdates(removes, adds);
  }

  /// @notice Internal version of applyAllowListUpdates to allow for reuse in the constructor.
  /// @param removes The addresses to be removed.
  /// @param adds The addresses to be added.
  function _applyAllowListUpdates(
    address[] memory removes,
    address[] memory adds
  ) internal virtual {
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
  // │                       Policy Engine                          │
  // ================================================================

  /// @notice Sets a new policy engine.
  /// @param newPolicyEngine The address of the new policy engine.
  function setPolicyEngine(
    IPolicyEngine newPolicyEngine
  ) public virtual onlyOwner {
    _setPolicyEngine(newPolicyEngine, false);
  }

  /// @notice Sets a new policy engine while tolerating a pre-existing policy engine's detach reverting.
  /// @dev Use this to force update an old policy engine whose detach() reverts.
  /// @param newPolicyEngine The address of the new policy engine.
  function setPolicyEngineAllowFailedDetach(
    IPolicyEngine newPolicyEngine
  ) public virtual onlyOwner {
    _setPolicyEngine(newPolicyEngine, true);
  }

  /// @notice Internal function to set and attach to a policy engine.
  /// @param newPolicyEngine The address of the new policy engine, or address(0) to disable.
  /// @param allowFailedDetach Whether to continue if the old policy engine's detach reverts.
  function _setPolicyEngine(
    IPolicyEngine newPolicyEngine,
    bool allowFailedDetach
  ) internal virtual {
    IPolicyEngine oldPolicyEngine = s_policyEngine;
    if (address(newPolicyEngine) == address(oldPolicyEngine)) return;

    if (address(oldPolicyEngine) != address(0)) {
      // Guarding detach reverts to offer escape hatch from adversarial policy engine instances.
      try oldPolicyEngine.detach() {}
      catch (bytes memory err) {
        if (!allowFailedDetach) revert PolicyEngineDetachReverted(oldPolicyEngine, err);
        emit PolicyEngineDetachFailed(oldPolicyEngine, err);
      }
    }

    s_policyEngine = newPolicyEngine;
    if (address(newPolicyEngine) != address(0)) {
      newPolicyEngine.attach();
    }

    emit PolicyEngineAttached(newPolicyEngine);
  }

  /// @notice Gets the current policy engine address.
  /// @return policyEngine The address of the policy engine.
  function getPolicyEngine() public view virtual returns (IPolicyEngine policyEngine) {
    return s_policyEngine;
  }
}
