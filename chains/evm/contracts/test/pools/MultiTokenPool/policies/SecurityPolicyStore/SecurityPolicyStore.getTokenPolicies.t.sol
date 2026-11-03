// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

contract SecurityPolicyStore_getTokenPolicies is SecurityPolicyStoreSetup {
  /// @notice getTokenPolicies reports the registration, the decimals and the policies of a token.
  function test_getTokenPolicies() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setSecurityPolicyConfig(OTHER_POLICY, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setTokenSecurityPolicy(s_tokenA, OTHER_POLICY);

    (bool isRegistered, uint8 localDecimals, bytes32 rateLimitPolicyName, bytes32 securityPolicyName) =
      s_store.getTokenPolicies(s_tokenA);
    assertTrue(isRegistered);
    assertEq(DEFAULT_TOKEN_DECIMALS, localDecimals);
    assertEq(DEFAULT_POLICY, rateLimitPolicyName);
    assertEq(OTHER_POLICY, securityPolicyName);

    // A token with no policy reports zero for both policies.
    (isRegistered, localDecimals, rateLimitPolicyName, securityPolicyName) = s_store.getTokenPolicies(s_tokenB);
    assertTrue(isRegistered);
    assertEq(DEFAULT_TOKEN_DECIMALS, localDecimals);
    assertEq(bytes32(0), rateLimitPolicyName);
    assertEq(bytes32(0), securityPolicyName);
  }
}
