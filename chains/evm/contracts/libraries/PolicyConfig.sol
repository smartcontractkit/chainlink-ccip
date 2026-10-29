// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../interfaces/IPoolConfigManager.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Types and helpers that the policy stores of multi-token pools share. A token has five types of config on
/// each remote chain: connection, chain, rate limit, fee and security. For each type, a token can use one policy and
/// can have inline records. A policy is a named set of records that many tokens share. An inline record belongs to one
/// token and one chain.
/// @dev A token uses its inline record for the chain. If it has no inline record, it uses the record of its policy for
/// the chain. If that record does not exist, it uses the default record of its policy. The connection type has no
/// default records. Setting a record writes every value of the record. The one exception is a connection policy
/// record, which keeps its subset tokens.
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

  /// @notice The record that a token uses for one type on one chain.
  enum RecordLevel {
    None, //          The token has no record of the type for the chain.
    Inline, //        The inline record of the token for the chain.
    PolicyChain, //   The record of the token's policy for the chain.
    PolicyDefault //  The default record of the token's policy.
  }

  /// @notice Identifies one policy record. Chain selector zero is the default record of the policy.
  struct PolicyRecordKey {
    bytes32 policyName; //         The policy.
    uint64 remoteChainSelector; // The remote chain.
  }

  /// @notice The profile of one token. Setting it registers a new token with no policy, or replaces the profile of a
  /// registered token.
  struct TokenProfileArgs {
    address localToken; // ──────────────────────────────────╮ The local token.
    uint8 localDecimals; //                                  │ Checked against decimals(). Fixed while registered.
    IPoolConfigManager.LockOrBurnFlavor lockOrBurnFlavor; // │ How the pool moves the token out.
    bytes4 allowedFinalityConfig; // ────────────────────────╯ The allowed finality. Zero allows only full finality.
    address lockOrBurnAddress; //                              The burn address or the lock box. Zero for other flavors.
    address additionalPoolDataProvider; //                     The IAdditionalPoolDataProvider. Zero means none.
  }

  /// @notice The policy of one type for one token.
  struct TokenPolicyArgs {
    address localToken; // A registered token.
    bytes32 policyName; // The policy. Zero removes the token from its policy of the type.
  }

  /// @dev A policy exists from the time its first record is set until it is removed. It can exist without records.
  /// Zero is never a policy name, so a token without a policy finds no policy record.
  struct Policy {
    EnumerableSet.UintSet remoteChainSelectors; // The chains of the policy's records. Zero is the default record.
    uint32 tokenCount; //                          The number of tokens that use the policy.
  }

  /// @dev The policies of one type. A name is in names only while its policy exists. So a view can list every policy,
  /// even a policy that no token uses.
  struct Policies {
    mapping(bytes32 policyName => Policy policy) byName; // The policies, by name.
    EnumerableSet.Bytes32Set names; //                      The name of each policy that exists.
  }

  /// @notice Returns the level of the policy record that a token uses on a chain. The record for the chain wins over
  /// the default record.
  /// @dev The caller checks the inline record of the token first, because an inline record wins over the policy. So a
  /// token with an inline record needs no read of its policy name. Chain selector zero returns None, because no lane
  /// uses it.
  /// @param policyChains The chains of the records of the token's policy.
  /// @param remoteChainSelector The remote chain.
  /// @return level PolicyChain, PolicyDefault or None.
  function _resolvePolicyRecordLevel(
    EnumerableSet.UintSet storage policyChains,
    uint64 remoteChainSelector
  ) internal view returns (RecordLevel level) {
    if (remoteChainSelector == DEFAULT_CHAIN_SELECTOR) return RecordLevel.None;
    if (policyChains.contains(remoteChainSelector)) return RecordLevel.PolicyChain;
    if (policyChains.contains(DEFAULT_CHAIN_SELECTOR)) return RecordLevel.PolicyDefault;
    return RecordLevel.None;
  }

  /// @notice Adds a policy record, or keeps it if it exists. The first record that is set under a new name creates
  /// the policy.
  /// @dev The caller writes the values of the record. For a type without default records, the caller rejects chain
  /// selector zero.
  function _setPolicyRecord(
    Policies storage policies,
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal {
    if (policyName == bytes32(0)) revert ZeroPolicyNameNotAllowed();
    policies.byName[policyName].remoteChainSelectors.add(remoteChainSelector);
    policies.names.add(policyName);
  }

  /// @notice Removes a policy record. The policy stays, even when it has no records left.
  /// @dev The caller deletes the values of the record.
  function _removePolicyRecord(
    Policies storage policies,
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal {
    if (!policies.byName[policyName].remoteChainSelectors.remove(remoteChainSelector)) {
      revert PolicyRecordNotFound(policyName, remoteChainSelector);
    }
  }

  /// @notice Removes a policy that has no tokens and no records.
  /// @dev Returns false for a policy that does not exist, so that the caller reverts with the error of its type. The
  /// caller deletes the other values that it holds for the policy.
  function _removePolicy(
    Policies storage policies,
    bytes32 policyName
  ) internal returns (bool isPolicyFound) {
    if (!policies.names.remove(policyName)) return false;
    Policy storage policy = policies.byName[policyName];
    if (policy.tokenCount != 0) revert PolicyInUse(policyName);
    if (policy.remoteChainSelectors.length() != 0) revert PolicyHasRecords(policyName);
    return true;
  }

  /// @notice Adds an inline record, or keeps it if it exists.
  /// @dev The caller writes the values of the record. There is no inline default record, so the chain selector must
  /// not be zero.
  function _setInlineRecord(
    EnumerableSet.AddressSet storage tokens,
    EnumerableSet.UintSet storage inlineChains,
    address localToken,
    uint64 remoteChainSelector
  ) internal {
    if (!tokens.contains(localToken)) revert LocalTokenNotRegistered(localToken);
    if (remoteChainSelector == DEFAULT_CHAIN_SELECTOR) revert ZeroChainSelectorNotAllowed();
    inlineChains.add(remoteChainSelector);
  }

  /// @notice Removes an inline record.
  /// @dev The caller deletes the values of the record.
  function _removeInlineRecord(
    EnumerableSet.UintSet storage inlineChains,
    address localToken,
    uint64 remoteChainSelector
  ) internal {
    if (!inlineChains.remove(remoteChainSelector)) revert InlineRecordNotFound(localToken, remoteChainSelector);
  }

  /// @notice Moves a token from its old policy to its new policy, and updates the token counts.
  /// @dev Returns false and changes nothing if the new policy does not exist, so that the caller reverts with the error
  /// of its type. The caller checks that the token is registered. The caller also stores the new policy name of the
  /// token. A token never uses a policy that does not exist, because a policy in use cannot be removed. A store must
  /// deregister a token only when the token has no policy. Otherwise the policy keeps the token in its count, and the
  /// policy can never be removed.
  function _moveToken(
    Policies storage policies,
    bytes32 oldPolicyName,
    bytes32 newPolicyName
  ) internal returns (bool isNewPolicyFound) {
    if (oldPolicyName == newPolicyName) return true;
    if (newPolicyName != bytes32(0) && !policies.names.contains(newPolicyName)) return false;
    if (oldPolicyName != bytes32(0)) --policies.byName[oldPolicyName].tokenCount;
    if (newPolicyName != bytes32(0)) ++policies.byName[newPolicyName].tokenCount;
    return true;
  }
}
