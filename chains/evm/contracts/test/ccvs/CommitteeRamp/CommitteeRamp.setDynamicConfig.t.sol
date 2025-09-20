// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitteeRamp} from "../../../ccvs/CommitteeRamp.sol";
import {CommitteeRampSetup} from "./CommitteeRampSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CommitteeRamp_setDynamicConfig is CommitteeRampSetup {
  function test_setDynamicConfig() public {
    CommitteeRamp.DynamicConfig memory newConfig = CommitteeRamp.DynamicConfig({
      feeQuoter: makeAddr("feeQuoter2"),
      feeAggregator: makeAddr("feeAggregator2"),
      allowlistAdmin: makeAddr("allowlistAdmin2")
    });

    vm.expectEmit();
    emit CommitteeRamp.ConfigSet(newConfig);

    s_commitRamp.setDynamicConfig(newConfig);

    CommitteeRamp.DynamicConfig memory got = s_commitRamp.getDynamicConfig();
    assertEq(got.feeQuoter, newConfig.feeQuoter);
    assertEq(got.feeAggregator, newConfig.feeAggregator);
    assertEq(got.allowlistAdmin, newConfig.allowlistAdmin);
  }

  function test_setDynamicConfig_RevertWhen_InvalidConfig() public {
    // Zero feeQuoter should revert
    CommitteeRamp.DynamicConfig memory badConfig = CommitteeRamp.DynamicConfig({
      feeQuoter: address(0),
      feeAggregator: FEE_AGGREGATOR,
      allowlistAdmin: ALLOWLIST_ADMIN
    });
    vm.expectRevert(CommitteeRamp.InvalidConfig.selector);
    s_commitRamp.setDynamicConfig(badConfig);
  }

  function test_setDynamicConfig_RevertWhen_OnlyCallableByOwner() public {
    CommitteeRamp.DynamicConfig memory cfg;

    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_commitRamp.setDynamicConfig(cfg);
  }
}
