// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitRamp} from "../../../ccvs/CommitRamp.sol";
import {CommitRampSetup} from "./CommitRampSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CommitRamp_setDynamicConfig is CommitRampSetup {
  function test_setDynamicConfig() public {
    CommitRamp.DynamicConfig memory newConfig = CommitRamp.DynamicConfig({
      feeQuoter: makeAddr("feeQuoter2"),
      feeAggregator: makeAddr("feeAggregator2"),
      allowlistAdmin: makeAddr("allowlistAdmin2")
    });

    vm.expectEmit();
    emit CommitRamp.ConfigSet(newConfig);

    s_commitRamp.setDynamicConfig(newConfig);

    CommitRamp.DynamicConfig memory got = s_commitRamp.getDynamicConfig();
    assertEq(got.feeQuoter, newConfig.feeQuoter);
    assertEq(got.feeAggregator, newConfig.feeAggregator);
    assertEq(got.allowlistAdmin, newConfig.allowlistAdmin);
  }

  function test_setDynamicConfig_RevertWhen_InvalidConfig() public {
    // Zero feeQuoter should revert
    CommitRamp.DynamicConfig memory badConfig = CommitRamp.DynamicConfig({
      feeQuoter: address(0),
      feeAggregator: FEE_AGGREGATOR,
      allowlistAdmin: ALLOWLIST_ADMIN
    });
    vm.expectRevert(CommitRamp.InvalidConfig.selector);
    s_commitRamp.setDynamicConfig(badConfig);
  }

  function test_setDynamicConfig_RevertWhen_OnlyCallableByOwner() public {
    CommitRamp.DynamicConfig memory cfg;

    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_commitRamp.setDynamicConfig(cfg);
  }
}
