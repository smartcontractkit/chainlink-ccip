// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitOnRamp} from "../../../onRamp/CommitOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CommitOnRamp_setDynamicConfig is CommitOnRampSetup {
  function test_setDynamicConfig() public {
    CommitOnRamp.StaticConfig memory staticConfig = s_commitOnRamp.getStaticConfig();

    CommitOnRamp.DynamicConfig memory newConfig = CommitOnRamp.DynamicConfig({
      feeQuoter: makeAddr("feeQuoter2"),
      feeAggregator: makeAddr("feeAggregator2"),
      allowlistAdmin: makeAddr("allowlistAdmin2")
    });

    vm.expectEmit();
    emit CommitOnRamp.ConfigSet(staticConfig, newConfig);

    s_commitOnRamp.setDynamicConfig(newConfig);

    CommitOnRamp.DynamicConfig memory got = s_commitOnRamp.getDynamicConfig();
    assertEq(got.feeQuoter, newConfig.feeQuoter);
    assertEq(got.feeAggregator, newConfig.feeAggregator);
    assertEq(got.allowlistAdmin, newConfig.allowlistAdmin);
  }

  function test_setDynamicConfig_RevertWhen_InvalidConfig() public {
    // Zero feeQuoter should revert
    CommitOnRamp.DynamicConfig memory badConfig = CommitOnRamp.DynamicConfig({
      feeQuoter: address(0),
      feeAggregator: FEE_AGGREGATOR,
      allowlistAdmin: ALLOWLIST_ADMIN
    });
    vm.expectRevert(CommitOnRamp.InvalidConfig.selector);
    s_commitOnRamp.setDynamicConfig(badConfig);
  }

  function test_setDynamicConfig_RevertWhen_OnlyCallableByOwner() public {
    CommitOnRamp.DynamicConfig memory cfg;

    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_commitOnRamp.setDynamicConfig(cfg);
  }
}
