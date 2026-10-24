// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IMultiTokenPoolHooks} from "../../interfaces/IMultiTokenPoolHooks.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {Pool} from "../../libraries/Pool.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice The hooks of the lanes of multi-token pools: a sender allowlist and policy engine runs. Both work the same as
/// in AdvancedPoolHooks.
/// @dev The pool config manager gives the CCVs of each lane, so these hooks have no CCV config. Each pool that serves a
/// token on a lane that uses these hooks must be an authorized caller. All tokens and lanes that use these hooks share
/// the allowlist and the policy engine.
contract MultiTokenPoolHooks is IMultiTokenPoolHooks, ITypeAndVersion, AuthorizedCallers {
  using EnumerableSet for EnumerableSet.AddressSet;

  error AllowListNotEnabled();
  error SenderNotAllowed(address sender);
  error PolicyEngineDetachReverted(IPolicyEngine oldPolicyEngine, bytes err);

  event AllowListAdd(address sender);
  event AllowListRemove(address sender);
  event PolicyEngineAttached(IPolicyEngine indexed policyEngine);
  event PolicyEngineDetachFailed(IPolicyEngine indexed policyEngine, bytes reason);

  /// @dev The allowlist can only be enabled at deployment, which keeps the check cheap on the hot path.
  bool private immutable i_allowlistEnabled;

  /// @dev The original senders that can send tokens. Only used when i_allowlistEnabled is true.
  EnumerableSet.AddressSet private s_allowlist;

  /// @dev Zero disables the policy engine runs.
  IPolicyEngine private s_policyEngine;

  /// @param allowlist The original senders that can send tokens. An empty list disables the allowlist for good.
  /// @param policyEngine The policy engine to attach. Zero disables the policy engine runs.
  /// @param authorizedCallers The pools that can call the hooks.
  constructor(
    address[] memory allowlist,
    IPolicyEngine policyEngine,
    address[] memory authorizedCallers
  ) AuthorizedCallers(authorizedCallers) {
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
  /// @dev Checks the original sender against the allowlist, then runs the policy engine with tokenArgs as context.
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
  /// @dev Runs the policy engine with offchainTokenData as context. The allowlist only applies on send.
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

  /// @notice Reverts if the allowlist is enabled and does not hold the sender.
  /// @param sender The original sender to check.
  function checkAllowList(
    address sender
  ) public view virtual {
    if (i_allowlistEnabled && !s_allowlist.contains(sender)) revert SenderNotAllowed(sender);
  }

  /// @notice Returns whether the allowlist is enabled.
  /// @return allowlistEnabled True if the allowlist is enabled.
  function getAllowListEnabled() public view virtual returns (bool allowlistEnabled) {
    return i_allowlistEnabled;
  }

  /// @notice Returns the original senders that can send tokens.
  /// @return allowlist The allowed senders.
  function getAllowList() public view virtual returns (address[] memory allowlist) {
    return s_allowlist.values();
  }

  /// @notice Removes and adds original senders on the allowlist. Removals run first.
  /// @dev Zero addresses in adds are skipped. An address that is already in the target state emits no event.
  /// @param removes The senders to remove.
  /// @param adds The senders to add.
  function applyAllowListUpdates(
    address[] calldata removes,
    address[] calldata adds
  ) public virtual onlyOwner {
    _applyAllowListUpdates(removes, adds);
  }

  /// @dev Shared by the constructor and applyAllowListUpdates.
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

  /// @notice Detaches the old policy engine and attaches the new one.
  /// @dev Reverts if the detach of the old policy engine reverts.
  /// @param newPolicyEngine The new policy engine. Zero disables the policy engine runs.
  function setPolicyEngine(
    IPolicyEngine newPolicyEngine
  ) public virtual onlyOwner {
    _setPolicyEngine(newPolicyEngine, false);
  }

  /// @notice Detaches the old policy engine and attaches the new one, also when the detach reverts.
  /// @dev Use this to replace an old policy engine whose detach reverts.
  /// @param newPolicyEngine The new policy engine. Zero disables the policy engine runs.
  function setPolicyEngineAllowFailedDetach(
    IPolicyEngine newPolicyEngine
  ) public virtual onlyOwner {
    _setPolicyEngine(newPolicyEngine, true);
  }

  /// @param newPolicyEngine The new policy engine. Zero disables the policy engine runs.
  /// @param allowFailedDetach True to continue when the detach of the old policy engine reverts. False to revert.
  function _setPolicyEngine(
    IPolicyEngine newPolicyEngine,
    bool allowFailedDetach
  ) internal virtual {
    IPolicyEngine oldPolicyEngine = s_policyEngine;
    if (address(newPolicyEngine) == address(oldPolicyEngine)) return;

    if (address(oldPolicyEngine) != address(0)) {
      // A detach that reverts must not stop the owner from replacing the policy engine.
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

  /// @notice Returns the policy engine.
  /// @return policyEngine The policy engine. Zero means the policy engine runs are disabled.
  function getPolicyEngine() public view virtual returns (IPolicyEngine policyEngine) {
    return s_policyEngine;
  }
}
