// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../libraries/PolicyConfig.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Exposes `PolicyConfig` internal functions for unit tests. It holds the policies of one type, the registered
/// tokens and the inline chains of each token, the same as a policy store.
contract PolicyConfigHelper {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;
  using EnumerableSet for EnumerableSet.UintSet;

  PolicyConfig.Policies private s_policies;
  EnumerableSet.AddressSet private s_tokens;
  mapping(address localToken => EnumerableSet.UintSet remoteChainSelectors) private s_inlineChains;

  function registerToken(
    address localToken
  ) external {
    s_tokens.add(localToken);
  }

  function resolvePolicyRecordLevel(
    bytes32 policyName,
    uint64 remoteChainSelector
  ) external view returns (PolicyConfig.RecordLevel) {
    return
      PolicyConfig._resolvePolicyRecordLevel(s_policies.byName[policyName].remoteChainSelectors, remoteChainSelector);
  }

  function setPolicyRecord(
    bytes32 policyName,
    uint64 remoteChainSelector
  ) external {
    PolicyConfig._setPolicyRecord(s_policies, policyName, remoteChainSelector);
  }

  function removePolicyRecord(
    bytes32 policyName,
    uint64 remoteChainSelector
  ) external {
    PolicyConfig._removePolicyRecord(s_policies, policyName, remoteChainSelector);
  }

  function removePolicy(
    bytes32 policyName
  ) external returns (bool isPolicyFound) {
    return PolicyConfig._removePolicy(s_policies, policyName);
  }

  function setInlineRecord(
    address localToken,
    uint64 remoteChainSelector
  ) external {
    PolicyConfig._setInlineRecord(s_tokens, s_inlineChains[localToken], localToken, remoteChainSelector);
  }

  function removeInlineRecord(
    address localToken,
    uint64 remoteChainSelector
  ) external {
    PolicyConfig._removeInlineRecord(s_inlineChains[localToken], localToken, remoteChainSelector);
  }

  function moveToken(
    bytes32 oldPolicyName,
    bytes32 newPolicyName
  ) external returns (bool isNewPolicyFound) {
    return PolicyConfig._moveToken(s_policies, oldPolicyName, newPolicyName);
  }

  function getPolicyNames() external view returns (bytes32[] memory policyNames) {
    return s_policies.names.values();
  }

  function getPolicyChains(
    bytes32 policyName
  ) external view returns (uint256[] memory remoteChainSelectors) {
    return s_policies.byName[policyName].remoteChainSelectors.values();
  }

  function getTokenCount(
    bytes32 policyName
  ) external view returns (uint32 tokenCount) {
    return s_policies.byName[policyName].tokenCount;
  }

  function getInlineChains(
    address localToken
  ) external view returns (uint256[] memory remoteChainSelectors) {
    return s_inlineChains[localToken].values();
  }
}
