// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitRamp} from "../../../ccvs/CommitRamp.sol";
import {CommitRampSetup} from "./CommitRampSetup.t.sol";

contract CommitRamp_getDynamicConfig is CommitRampSetup {
  function test_GetDynamicConfig() public view {
    CommitRamp.DynamicConfig memory d = s_commitRamp.getDynamicConfig();
    assertEq(d.feeQuoter, address(s_feeQuoter));
    assertEq(d.feeAggregator, FEE_AGGREGATOR);
    assertEq(d.allowlistAdmin, ALLOWLIST_ADMIN);
  }
}
