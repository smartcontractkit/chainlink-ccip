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

  /// @notice Identifies one policy record.
  struct PolicyRecordKey {
    bytes32 policyName; // The name of the policy.
    uint64 remoteChainSelector; // The remote chain selector, zero for the default record.
  }

  /// @notice The profile of one token. Setting it registers a new token without policies, or replaces the profile of a
  /// registered token.
  struct TokenProfileArgs {
    address localToken; // ──────────────────────────────────╮ The local token.
    uint8 localDecimals; //                                  │ Checked against decimals(). Fixed while registered.
    IPoolConfigManager.LockOrBurnFlavor lockOrBurnFlavor; // │ How the pool moves the token out.
    bytes4 allowedFinalityConfig; // ────────────────────────╯ The allowed FTF finality. Zero allows default only.
    address lockOrBurnAddress; // The burn address for BurnToAddress, the lock box for Lock, zero otherwise.
    address additionalPoolDataProvider; // The IAdditionalPoolDataProvider of the token, zero for none.
  }

  /// @notice The policy of one policy type for one token.
  struct TokenPolicyArgs {
    address localToken; // A registered token.
    bytes32 policyName; // The name of the policy. Zero removes the token from its current policy.
  }

  /// @notice The state of one policy. A policy exists from the time its first record is set until it is removed. It can
  /// exist without records.
  /// @dev A store must deregister only a token without a policy. Otherwise the token count of the policy stays above
  /// zero, and the policy can never be removed.
  struct Policy {
    EnumerableSet.UintSet remoteChainSelectors; // The chains of the policy records. Zero names the default record.
    uint32 tokenCount; // The number of tokens that use the policy.
  }

  /// @notice The policies of one policy type. Zero is never a policy name, so a token without a policy never finds a
  /// policy record.
  struct Policies {
    EnumerableSet.Bytes32Set policyNames; // The names of all policies, including policies that no token uses.
    mapping(bytes32 policyName => Policy policy) policies; // The state of each policy.
  }

  /// @notice Returns the level of the policy record that serves a chain. A record for the chain is used before the
  /// default record.
  /// @dev The caller checks the inline record of the token first, as an inline record overrides the policy. Chain
  /// selector zero never resolves to a record, as no lane uses it.
  /// @param s_policyChainSelectors The chains of the records of the token's policy. For a token without a policy, the
  /// empty set of policy name zero.
  /// @param remoteChainSelector The remote chain selector.
  /// @return The level of the policy record. It is never Inline.
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

  /// @notice Adds a record to a policy, or keeps the record if it exists. The first record that is set under a new name
  /// creates the policy.
  /// @dev The caller writes the values of the record. For a policy type without default records, the caller must reject
  /// chain selector zero.
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

  /// @notice Adds an inline record to a token, or keeps the record if it exists.
  /// @dev The caller writes the values of the record. There is no inline default record, so chain selector zero is
  /// rejected.
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
  /// @dev The caller deletes the values of the record.
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

  /// @notice Updates the token counts when a token changes from its old policy to its new policy. A zero policy name
  /// means that the token has no policy.
  /// @dev The caller checks that the token is registered. The caller then stores the new policy name for the token.
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
      --s_policies.policies[oldPolicyName].tokenCount;
    }
    if (newPolicyName != bytes32(0)) {
      ++s_policies.policies[newPolicyName].tokenCount;
    }
  }
}
