// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

contract SecurityPolicyStore_getTokenPolicyNamesBatch is SecurityPolicyStoreSetup {
  /// @notice getTokenPolicyNamesBatch reports both policies of each token, aligned by index.
  function test_getTokenPolicyNamesBatch() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setSecurityPolicyConfig(OTHER_POLICY, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setTokenSecurityPolicy(s_tokenA, OTHER_POLICY);

    (bytes32[] memory rateLimitPolicyNames, bytes32[] memory securityPolicyNames) =
      s_store.getTokenPolicyNamesBatch(_arrayOf(s_tokenA));
    assertEq(DEFAULT_POLICY, rateLimitPolicyNames[0]);
    assertEq(OTHER_POLICY, securityPolicyNames[0]);

    // A token with no policy reports zero for both.
    (rateLimitPolicyNames, securityPolicyNames) = s_store.getTokenPolicyNamesBatch(_arrayOf(s_tokenB));
    assertEq(bytes32(0), rateLimitPolicyNames[0]);
    assertEq(bytes32(0), securityPolicyNames[0]);
  }
}
