// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitteeRamp} from "../../../ccvs/CommitteeRamp.sol";
import {CommitteeRampSetup} from "./CommitteeRampSetup.t.sol";

contract CommitteeRamp_getDynamicConfig is CommitteeRampSetup {
  function test_GetDynamicConfig() public view {
    CommitteeRamp.DynamicConfig memory d = s_commitRamp.getDynamicConfig();
    assertEq(d.feeQuoter, address(s_feeQuoter));
    assertEq(d.feeAggregator, FEE_AGGREGATOR);
    assertEq(d.allowlistAdmin, ALLOWLIST_ADMIN);
  }
}
