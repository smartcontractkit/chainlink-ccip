// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {BaseTest} from "../../../../BaseTest.t.sol";

contract FeePolicyStore_getPolicyNames is BaseTest {
  function test_getPolicyNames_NoPolicies() public {
    FeePolicyStore feeStore = new FeePolicyStore(new address[](0));

    (bytes32[] memory feePolicyNames, bytes32[] memory destOverheadPolicyNames) = feeStore.getPolicyNames();

    assertEq(0, feePolicyNames.length);
    assertEq(0, destOverheadPolicyNames.length);
  }
}
