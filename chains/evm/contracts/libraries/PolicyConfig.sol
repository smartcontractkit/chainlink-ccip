// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../interfaces/IPoolConfigManager.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Types and helpers that the policy stores of the pool config manager share. A token can use one policy of
/// each policy type, such as connection, rate limit and fee.
/// @dev A token uses its inline record for the chain. If it has no inline record, it uses the record of its policy for
/// the chain. If the per-chain record does not exist, it uses the default record of its policy. The connection policy
/// type has no default records.
library PolicyConfig {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;
  using EnumerableSet for EnumerableSet.UintSet;

  error ZeroPolicyNameNotAllowed();
  error PolicyNotFound(bytes32 policyName);
  error PolicyRecordNotFound(bytes32 policyName, uint64 remoteChainSelector);
  error PolicyInUse(bytes32 policyName);
  error PolicyHasRecords(bytes32 policyName);
  error LocalTokenNotRegistered(address localToken);
  error ZeroChainSelectorNotAllowed();
  error InlineRecordNotFound(address localToken, uint64 remoteChainSelector);

  /// @notice The chain selector of the policy default record. No chain has selector zero, so this record can serve
  /// every chain where the policy has no record of its own.
  uint64 internal constant DEFAULT_CHAIN_SELECTOR = 0;

  /// @notice The record that a token uses for one policy type on one chain.
  enum RecordLevel {
    None, // The token has no record for the chain.
    Inline, // The inline record of the token for the chain.
    PolicyChain, // The record of the token's policy for the chain.
    PolicyDefault // The default record of the token's policy.
  }

  /// @notice Specifies one policy record.
  struct PolicyRecordKey {
    bytes32 policyName; // The name of the policy.
    uint64 remoteChainSelector; // The remote chain selector, zero for the default record.
  }

  /// @notice The profile of one token, independent of policies.
  struct TokenProfileArgs {
    address localToken; // ──────────────────────────────────╮ The local token.
    uint8 localDecimals; //                                  │ Checked against decimals(). Fixed while registered.
    IPoolConfigManager.LockOrBurnFlavor lockOrBurnFlavor; // │ Specifies token's burn flavor, or to use lock/release.
    bytes4 allowedFinalityConfig; // ────────────────────────╯ The allowed FTF finality. Zero allows default only.
    address lockOrBurnAddress; // The burn address for BurnToAddress, the lock box for Lock, zero otherwise.
    address additionalPoolDataProvider; // The IAdditionalPoolDataProvider of the token, zero for none.
  }

  /// @notice Specifies a token's policy for one policy type.
  struct TokenPolicyArgs {
    address localToken; // A registered token.
    bytes32 policyName; // The name of the policy. Zero removes the token from its current policy.
  }

  /// @notice The state of one policy instance.
  /// @dev A policy cannot be removed if there are any tokens using it.
  struct Policy {
    EnumerableSet.UintSet remoteChainSelectors; // The chains of the policy records. Zero names the default record.
    uint32 tokenCount; // The number of tokens that use the policy.
  }

  /// @notice The policy registry of one policy type. Zero is never a policy name.
  struct Policies {
    EnumerableSet.Bytes32Set policyNames; // The names of all policies, including policies that no token uses.
    mapping(bytes32 policyName => Policy policy) policies; // The state of each policy.
  }

  /// @notice Helper function that returns the level of the policy record that serves a chain.
  /// A chain-specific record takes higher precedence over the default record.
  /// @dev The caller checks the inline record of the token first, as an inline record overrides the policy.
  /// @param s_policyChainSelectors The chains of the records of the token's policy.
  /// @param remoteChainSelector The remote chain selector.
  /// @return The level of the policy record.
  function _resolvePolicyRecordLevel(
    EnumerableSet.UintSet storage s_policyChainSelectors,
    uint64 remoteChainSelector
  ) internal view returns (RecordLevel) {
    if (remoteChainSelector == DEFAULT_CHAIN_SELECTOR) {
      return RecordLevel.None;
    }
    if (s_policyChainSelectors.contains(remoteChainSelector)) {
      return RecordLevel.PolicyChain;
    }
    if (s_policyChainSelectors.contains(DEFAULT_CHAIN_SELECTOR)) {
      return RecordLevel.PolicyDefault;
    }
    return RecordLevel.None;
  }

  /// @notice Adds a record to a policy, or updates the record if it exists.
  /// @param s_policies The policies of the policy type.
  /// @param policyName The name of the policy.
  /// @param remoteChainSelector The chain of the record. Zero names the default record.
  function _setPolicyRecord(
    Policies storage s_policies,
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal {
    if (policyName == bytes32(0)) {
      revert ZeroPolicyNameNotAllowed();
    }

    s_policies.policyNames.add(policyName);
    s_policies.policies[policyName].remoteChainSelectors.add(remoteChainSelector);
  }

  /// @notice Removes a record from a policy. The policy stays, even when it has no records left.
  /// @dev The caller deletes the values of the record.
  /// @param s_policies The policies of the policy type.
  /// @param policyName The name of the policy.
  /// @param remoteChainSelector The chain of the record. Zero names the default record.
  function _removePolicyRecord(
    Policies storage s_policies,
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal {
    if (!s_policies.policies[policyName].remoteChainSelectors.remove(remoteChainSelector)) {
      revert PolicyRecordNotFound(policyName, remoteChainSelector);
    }
  }

  /// @notice Removes a policy that has no tokens and no records.
  /// @dev The caller deletes any other values it stores for the policy.
  /// @param s_policies The policies of the policy type.
  /// @param policyName The name of the policy.
  function _removePolicy(
    Policies storage s_policies,
    bytes32 policyName
  ) internal {
    if (!s_policies.policyNames.contains(policyName)) {
      revert PolicyNotFound(policyName);
    }
    Policy storage policy = s_policies.policies[policyName];
    if (policy.tokenCount != 0) {
      revert PolicyInUse(policyName);
    }
    if (policy.remoteChainSelectors.length() != 0) {
      revert PolicyHasRecords(policyName);
    }

    s_policies.policyNames.remove(policyName);
  }

  /// @notice Adds an inline record to a token, or updates the record if it exists.
  /// @dev There are no defaults for inline records, chain selector zero is rejected.
  /// @param s_tokens The registered tokens.
  /// @param s_inlineChainSelectors The chains of the inline records of the token.
  /// @param localToken The token.
  /// @param remoteChainSelector The chain of the record.
  function _setInlineRecord(
    EnumerableSet.AddressSet storage s_tokens,
    EnumerableSet.UintSet storage s_inlineChainSelectors,
    address localToken,
    uint64 remoteChainSelector
  ) internal {
    if (!s_tokens.contains(localToken)) {
      revert LocalTokenNotRegistered(localToken);
    }
    if (remoteChainSelector == DEFAULT_CHAIN_SELECTOR) {
      revert ZeroChainSelectorNotAllowed();
    }

    s_inlineChainSelectors.add(remoteChainSelector);
  }

  /// @notice Removes an inline record from a token.
  /// @param s_inlineChainSelectors The chains of the inline records of the token.
  /// @param localToken The token.
  /// @param remoteChainSelector The chain of the record.
  function _removeInlineRecord(
    EnumerableSet.UintSet storage s_inlineChainSelectors,
    address localToken,
    uint64 remoteChainSelector
  ) internal {
    if (!s_inlineChainSelectors.remove(remoteChainSelector)) {
      revert InlineRecordNotFound(localToken, remoteChainSelector);
    }
  }

  /// @notice Updates the token counts when a token changes from its old policy to its new policy.
  /// A zero policy name means that the token has no policy.
  /// @param s_policies The policies of the policy type.
  /// @param oldPolicyName The current policy of the token, as the caller stores it.
  /// @param newPolicyName The new policy of the token.
  function _moveToken(
    Policies storage s_policies,
    bytes32 oldPolicyName,
    bytes32 newPolicyName
  ) internal {
    if (oldPolicyName == newPolicyName) {
      return;
    }
    if (newPolicyName != bytes32(0) && !s_policies.policyNames.contains(newPolicyName)) {
      revert PolicyNotFound(newPolicyName);
    }

    if (oldPolicyName != bytes32(0)) {
      s_policies.policies[oldPolicyName].tokenCount -= 1;
    }
    if (newPolicyName != bytes32(0)) {
      s_policies.policies[newPolicyName].tokenCount += 1;
    }
  }
}
