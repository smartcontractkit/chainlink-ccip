// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPVerifier} from "../../../ccvs/CCTPVerifier.sol";
import {CCTPVerifierSetup} from "./CCTPVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CCTPVerifier_setDynamicConfig is CCTPVerifierSetup {
  function test_setDynamicConfig() public {
    CCTPVerifier.DynamicConfig memory newConfig = CCTPVerifier.DynamicConfig({
      feeAggregator: makeAddr("feeAggregator2"),
      allowlistAdmin: makeAddr("allowlistAdmin2")
    });

    vm.expectEmit();
    emit CCTPVerifier.DynamicConfigSet(newConfig);

    s_cctpVerifier.setDynamicConfig(newConfig);

    CCTPVerifier.DynamicConfig memory got = s_cctpVerifier.getDynamicConfig();
    assertEq(got.feeAggregator, newConfig.feeAggregator);
    assertEq(got.allowlistAdmin, newConfig.allowlistAdmin);
  }

  function test_setDynamicConfig_RevertWhen_OnlyCallableByOwner() public {
    CCTPVerifier.DynamicConfig memory cfg;

    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_cctpVerifier.setDynamicConfig(cfg);
  }
}
