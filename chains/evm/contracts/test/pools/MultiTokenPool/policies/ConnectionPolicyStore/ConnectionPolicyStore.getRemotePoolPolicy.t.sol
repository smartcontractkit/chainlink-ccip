// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_getRemotePoolPolicy is ConnectionPolicyStoreSetup {
  function test_getRemotePoolPolicy_UnknownPolicy() public view {
    _assertRemotePoolPolicy(OTHER_POLICY, 0, new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](0));
  }
}
