// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

contract SecurityPolicyStore_getPolicyNames is SecurityPolicyStoreSetup {
  /// @notice getPolicyNames reports both policy types.
  function test_getPolicyNames() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setSecurityPolicyConfig(OTHER_POLICY, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);

    (bytes32[] memory rateLimitPolicyNames, bytes32[] memory securityPolicyNames) = s_store.getPolicyNames();
    assertEq(1, rateLimitPolicyNames.length);
    assertEq(DEFAULT_POLICY, rateLimitPolicyNames[0]);
    assertEq(1, securityPolicyNames.length);
    assertEq(OTHER_POLICY, securityPolicyNames[0]);
  }
}
