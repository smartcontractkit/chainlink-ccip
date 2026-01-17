// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ExecutorOnRamp} from "../../../onRamp/ExecutorOnRamp.sol";
import {BaseTest} from "../../BaseTest.t.sol";

contract ExecutorOnRampSetup is BaseTest {
  ExecutorOnRamp internal s_executorOnRamp;
  address internal constant INITIAL_CCV = address(121212);
  uint64 internal constant INITIAL_DEST = 1;
  uint8 internal constant INITIAL_MAX_CCVS = 1;

  function setUp() public override {
    super.setUp();

    s_executorOnRamp = new ExecutorOnRamp(INITIAL_MAX_CCVS);

    address[] memory ccvs = new address[](1);
    ccvs[0] = INITIAL_CCV;
    s_executorOnRamp.applyAllowedCCVUpdates(new address[](0), ccvs, true);

    uint64[] memory dests = new uint64[](1);
    dests[0] = INITIAL_DEST;
    s_executorOnRamp.applyDestChainUpdates(new uint64[](0), dests);
  }
}
