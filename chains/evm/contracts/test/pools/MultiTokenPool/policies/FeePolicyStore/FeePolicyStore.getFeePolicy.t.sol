// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";

contract FeePolicyStore_getFeePolicy is FeePolicyStoreSetup {
  function test_getFeePolicy_UnknownPolicy() public view {
    (uint32 tokenCount, FeePolicyStore.FeePolicyConfigArgs[] memory policyConfigs) =
      s_feeStore.getFeePolicy(OTHER_POLICY);

    assertEq(0, tokenCount);
    assertEq(0, policyConfigs.length);
  }
}
