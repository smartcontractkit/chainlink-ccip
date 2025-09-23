// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitteeVerifier} from "../../../ccvs/CommitteeVerifier.sol";
import {CommitteeVerifierSetup} from "./CommitteeVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CommitteeVerifier_setDynamicConfig is CommitteeVerifierSetup {
  function test_setDynamicConfig() public {
    CommitteeVerifier.DynamicConfig memory newConfig = CommitteeVerifier.DynamicConfig({
      feeQuoter: makeAddr("feeQuoter2"),
      feeAggregator: makeAddr("feeAggregator2"),
      allowlistAdmin: makeAddr("allowlistAdmin2")
    });

    vm.expectEmit();
    emit CommitteeVerifier.ConfigSet(newConfig);

    s_committeeVerifier.setDynamicConfig(newConfig);

    CommitteeVerifier.DynamicConfig memory got = s_committeeVerifier.getDynamicConfig();
    assertEq(got.feeQuoter, newConfig.feeQuoter);
    assertEq(got.feeAggregator, newConfig.feeAggregator);
    assertEq(got.allowlistAdmin, newConfig.allowlistAdmin);
  }

  function test_setDynamicConfig_RevertWhen_InvalidConfig() public {
    // Zero feeQuoter should revert
    CommitteeVerifier.DynamicConfig memory badConfig = CommitteeVerifier.DynamicConfig({
      feeQuoter: address(0),
      feeAggregator: FEE_AGGREGATOR,
      allowlistAdmin: ALLOWLIST_ADMIN
    });
    vm.expectRevert(CommitteeVerifier.InvalidConfig.selector);
    s_committeeVerifier.setDynamicConfig(badConfig);
  }

  function test_setDynamicConfig_RevertWhen_OnlyCallableByOwner() public {
    CommitteeVerifier.DynamicConfig memory cfg;

    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_committeeVerifier.setDynamicConfig(cfg);
  }
}
