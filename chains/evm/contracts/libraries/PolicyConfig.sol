// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../interfaces/IPoolConfigManager.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Types and helpers that the policy stores of the pool config manager share. A token can attach several
/// policy types. For example, ConnectionPolicy, RateLimitPolicy, FeePolicy.
/// @dev A token uses its inline record for the chain. If it has no inline record, it uses the record of its policy for
/// the chain. If the per-chain record does not exist, it uses the default record of its policy. The connection policy
/// type has no default records.
library PolicyConfig {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;
  using EnumerableSet for EnumerableSet.UintSet;

  error ZeroPolicyNameNotAllowed();
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
    uint8 localDecimals; //                                  │ Must match decimals(). Fixed while registered.
    IPoolConfigManager.LockOrBurnFlavor lockOrBurnFlavor; // │ How the pool moves the token out.
    bytes4 allowedFinalityConfig; // ────────────────────────╯ The allowed finality. Zero allows only full finality.
    address lockOrBurnAddress; // The burn address for BurnToAddress, the lock box for Lock, zero otherwise.
    address additionalPoolDataProvider; // The IAdditionalPoolDataProvider of the token, zero for none.
  }

  /// @notice The policy of one policy type for one token.
  struct TokenPolicyArgs {
    address localToken; // A registered token.
    bytes32 policyName; // The name of the policy, zero removes the token from its current policy.
  }

  /// @notice The policies of one policy type. A policy exists from the time its first record is set until it is
  /// removed, and it can exist without records. Chain selector zero is the default record of a policy. Zero is never a
  /// policy name, so a token without a policy never finds a policy record.
  struct Policies {
    EnumerableSet.Bytes32Set policyNames; // The names of all policies, including policies that no token uses.
    mapping(bytes32 policyName => EnumerableSet.UintSet) remoteChainSelectors; // The chains of each policy's records.
    mapping(bytes32 policyName => uint32) tokenCounts; // The number of tokens that use each policy.
  }

  /// @notice Returns the level of the policy record that a token uses on a chain. A record for the chain is used before
  /// the default record.
  /// @dev The caller checks the inline record of the token first, as an inline record overrides the policy. This way a
  /// token with an inline record needs no storage read of its policy name. Chain selector zero never resolves to a
  /// record, as no lane uses it.
  /// @param s_policyChainSelectors The chains of the records of the token's policy.
  /// @param remoteChainSelector The remote chain selector.
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
  function _setPolicyRecord(
    Policies storage s_policies,
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal {
    if (policyName == bytes32(0)) {
      revert ZeroPolicyNameNotAllowed();
    }

    s_policies.policyNames.add(policyName);
    s_policies.remoteChainSelectors[policyName].add(remoteChainSelector);
  }

  /// @notice Removes a record from a policy. The policy stays, even when it has no records left.
  /// @dev The caller deletes the values of the record.
  function _removePolicyRecord(
    Policies storage s_policies,
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal {
    if (!s_policies.remoteChainSelectors[policyName].remove(remoteChainSelector)) {
      revert PolicyRecordNotFound(policyName, remoteChainSelector);
    }
  }

  /// @notice Removes a policy that has no tokens and no records.
  /// @dev Returns false if the policy does not exist, so the caller can revert with the error of its policy type. The
  /// caller deletes any other values it stores for the policy.
  function _removePolicy(
    Policies storage s_policies,
    bytes32 policyName
  ) internal returns (bool) {
    if (!s_policies.policyNames.contains(policyName)) {
      return false;
    }
    if (s_policies.tokenCounts[policyName] != 0) {
      revert PolicyInUse(policyName);
    }
    if (s_policies.remoteChainSelectors[policyName].length() != 0) {
      revert PolicyHasRecords(policyName);
    }

    s_policies.policyNames.remove(policyName);
    return true;
  }

  /// @notice Adds an inline record to a token, or keeps the record if it exists.
  /// @dev The caller writes the values of the record. There is no inline default record, so chain selector zero is
  /// rejected.
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
  function _removeInlineRecord(
    EnumerableSet.UintSet storage s_inlineChainSelectors,
    address localToken,
    uint64 remoteChainSelector
  ) internal {
    if (!s_inlineChainSelectors.remove(remoteChainSelector)) {
      revert InlineRecordNotFound(localToken, remoteChainSelector);
    }
  }

  /// @notice Moves a token from its old policy to its new policy, and updates the token count of both policies. A zero
  /// policy name means that the token has no policy.
  /// @dev Returns false without changes if the new policy does not exist, so the caller can revert with the error of
  /// its policy type. The caller checks that the token is registered, and stores the new policy name of the token. A
  /// store must only deregister a token that has no policy. Otherwise the old policy keeps counting the token, and the
  /// policy can never be removed.
  function _moveToken(
    Policies storage s_policies,
    bytes32 oldPolicyName,
    bytes32 newPolicyName
  ) internal returns (bool) {
    if (oldPolicyName == newPolicyName) {
      return true;
    }
    if (newPolicyName != bytes32(0) && !s_policies.policyNames.contains(newPolicyName)) {
      return false;
    }

    if (oldPolicyName != bytes32(0)) {
      --s_policies.tokenCounts[oldPolicyName];
    }
    if (newPolicyName != bytes32(0)) {
      ++s_policies.tokenCounts[newPolicyName];
    }
    return true;
  }
}
