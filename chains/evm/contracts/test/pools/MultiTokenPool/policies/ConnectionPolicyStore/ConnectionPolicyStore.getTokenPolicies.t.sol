// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_getTokenPolicies is ConnectionPolicyStoreSetup {
  function test_getTokenPolicies() public {
    (bool isRegistered, bytes32 connectionPolicyName, bytes32 remotePoolPolicyName) =
      s_connectionStore.getTokenPolicies(s_token);
    assertTrue(isRegistered);
    assertEq(CONNECTION_POLICY, connectionPolicyName);
    assertEq(REMOTE_POOL_POLICY, remotePoolPolicyName);

    (isRegistered, connectionPolicyName, remotePoolPolicyName) = s_connectionStore.getTokenPolicies(s_otherToken);
    assertTrue(isRegistered);
    assertEq(bytes32(0), connectionPolicyName);
    assertEq(bytes32(0), remotePoolPolicyName);

    (isRegistered, connectionPolicyName, remotePoolPolicyName) =
      s_connectionStore.getTokenPolicies(makeAddr("unknownToken"));
    assertFalse(isRegistered);
    assertEq(bytes32(0), connectionPolicyName);
    assertEq(bytes32(0), remotePoolPolicyName);
  }
}
