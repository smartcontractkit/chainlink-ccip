// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

contract PoolConfigManager_getPolicyNames is PoolConfigManagerSetup {
  /// @notice The first record set under a new name adds the name to the list of its type only. A name used in two types
  /// is in both lists. No token uses these new policies, and they are listed.
  function test_getPolicyNames() public {
    _setFeePolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR, _getFeeConfig(100, 10));

    bytes32[] memory defaultOnly = _namesOf(DEFAULT_POLICY);
    bytes32[] memory defaultAndOther = _namesOf(DEFAULT_POLICY, OTHER_POLICY);
    _assertPolicyNames(defaultOnly, defaultOnly, defaultOnly, defaultAndOther, defaultOnly, defaultOnly);

    _setConnectionPolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR, false);

    _assertPolicyNames(defaultAndOther, defaultOnly, defaultOnly, defaultAndOther, defaultOnly, defaultOnly);
  }

  function _namesOf(
    bytes32 policyName
  ) internal pure returns (bytes32[] memory policyNames) {
    policyNames = new bytes32[](1);
    policyNames[0] = policyName;
    return policyNames;
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

  /// @dev Checks the names of each type, in the order of the view.
  function _assertPolicyNames(
    bytes32[] memory expectedConnectionPolicyNames,
    bytes32[] memory expectedRemotePoolPolicyNames,
    bytes32[] memory expectedRateLimitPolicyNames,
    bytes32[] memory expectedFeePolicyNames,
    bytes32[] memory expectedDestOverheadPolicyNames,
    bytes32[] memory expectedSecurityPolicyNames
  ) internal view {
    (
      bytes32[] memory connectionPolicyNames,
      bytes32[] memory remotePoolPolicyNames,
      bytes32[] memory rateLimitPolicyNames,
      bytes32[] memory feePolicyNames,
      bytes32[] memory destOverheadPolicyNames,
      bytes32[] memory securityPolicyNames
    ) = s_poolConfigManager.getPolicyNames();
    assertEq(expectedConnectionPolicyNames, connectionPolicyNames);
    assertEq(expectedRemotePoolPolicyNames, remotePoolPolicyNames);
    assertEq(expectedRateLimitPolicyNames, rateLimitPolicyNames);
    assertEq(expectedFeePolicyNames, feePolicyNames);
    assertEq(expectedDestOverheadPolicyNames, destOverheadPolicyNames);
    assertEq(expectedSecurityPolicyNames, securityPolicyNames);
  }
}
