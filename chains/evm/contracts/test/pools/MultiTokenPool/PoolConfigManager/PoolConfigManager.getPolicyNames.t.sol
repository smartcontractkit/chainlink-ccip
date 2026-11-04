// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

contract PoolConfigManager_getPolicyNames is PoolConfigManagerSetup {
  /// @notice Creating a policy adds its name only to the list for its type. Policies need no assigned tokens to
  /// appear. The same name can appear in multiple lists.
  function test_getPolicyNames() public {
    _addPolicyOfEachType();

    (
      bytes32[] memory connectionPolicyNames,
      bytes32[] memory remotePoolPolicyNames,
      bytes32[] memory rateLimitPolicyNames,
      bytes32[] memory feePolicyNames,
      bytes32[] memory destOverheadPolicyNames,
      bytes32[] memory securityPolicyNames
    ) = s_poolConfigManager.getPolicyNames();
    assertEq(_namesOf(DEFAULT_POLICY, CONNECTION_POLICY), connectionPolicyNames);
    assertEq(_namesOf(DEFAULT_POLICY, REMOTE_POOL_POLICY), remotePoolPolicyNames);
    assertEq(_namesOf(DEFAULT_POLICY, RATE_LIMIT_POLICY), rateLimitPolicyNames);
    assertEq(_namesOf(DEFAULT_POLICY, FEE_POLICY), feePolicyNames);
    assertEq(_namesOf(DEFAULT_POLICY, DEST_OVERHEAD_POLICY), destOverheadPolicyNames);
    assertEq(_namesOf(DEFAULT_POLICY, SECURITY_POLICY), securityPolicyNames);
  }

  function _namesOf(
    bytes32 firstPolicyName,
    bytes32 secondPolicyName
  ) internal pure returns (bytes32[] memory policyNames) {
    policyNames = new bytes32[](2);
    policyNames[0] = firstPolicyName;
    policyNames[1] = secondPolicyName;
    return policyNames;
  }
}
