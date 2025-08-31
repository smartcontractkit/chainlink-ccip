// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ExecutorOnRamp} from "../../../onRamp/ExecutorOnRamp.sol";

import {Test} from "forge-std/Test.sol";

contract ExecutorOnRampSetup is Test {
  ExecutorOnRamp internal s_executorOnRamp;
  address internal constant INITIAL_CCV = address(121212);
  uint64 internal constant INITIAL_DEST = 1;

  function setUp() public {
    vm.expectEmit();
    emit ExecutorOnRamp.ConfigSet(ExecutorOnRamp.DynamicConfig({maxCCVsPerMsg: 1}));
    s_executorOnRamp = new ExecutorOnRamp(ExecutorOnRamp.DynamicConfig({maxCCVsPerMsg: 1}));

    address[] memory ccvs = new address[](1);
    ccvs[0] = INITIAL_CCV;
    vm.expectEmit();
    emit ExecutorOnRamp.CCVAdded(INITIAL_CCV);
    vm.expectEmit();
    emit ExecutorOnRamp.AllowlistUpdated(true);
    s_executorOnRamp.applyAllowedCCVUpdates(ccvs, new address[](0), true);

    uint64[] memory dests = new uint64[](1);
    dests[0] = INITIAL_DEST;
    vm.expectEmit();
    emit ExecutorOnRamp.DestChainAdded(INITIAL_DEST);
    s_executorOnRamp.applyDestChainUpdates(dests, new uint64[](0));
  }
}
