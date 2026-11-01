// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_getPolicyNames is ConnectionPolicyStoreSetup {
  function test_getPolicyNames() public view {
    (bytes32[] memory connectionPolicyNames, bytes32[] memory remotePoolPolicyNames) =
      s_connectionStore.getPolicyNames();

    assertEq(1, connectionPolicyNames.length);
    assertEq(CONNECTION_POLICY, connectionPolicyNames[0]);
    assertEq(1, remotePoolPolicyNames.length);
    assertEq(REMOTE_POOL_POLICY, remotePoolPolicyNames[0]);
  }

  function test_getPolicyNames_NoPolicies() public {
    ConnectionPolicyStore connectionStore = new ConnectionPolicyStore(new address[](0));

    (bytes32[] memory connectionPolicyNames, bytes32[] memory remotePoolPolicyNames) = connectionStore.getPolicyNames();

    assertEq(0, connectionPolicyNames.length);
    assertEq(0, remotePoolPolicyNames.length);
  }
}
