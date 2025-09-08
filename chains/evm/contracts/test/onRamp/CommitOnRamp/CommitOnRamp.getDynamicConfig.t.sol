// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitOnRamp} from "../../../onRamp/CommitOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";

contract CommitOnRamp_getDynamicConfig is CommitOnRampSetup {
  function test_GetDynamicConfig() public view {
    CommitOnRamp.DynamicConfig memory d = s_commitOnRamp.getDynamicConfig();
    assertEq(d.feeQuoter, address(s_feeQuoter));
    assertEq(d.feeAggregator, FEE_AGGREGATOR);
    assertEq(d.allowlistAdmin, ALLOWLIST_ADMIN);
  }
}
