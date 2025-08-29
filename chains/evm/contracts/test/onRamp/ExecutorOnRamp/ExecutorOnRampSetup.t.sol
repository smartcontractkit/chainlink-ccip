// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ExecutorOnRamp} from "../../../onRamp/ExecutorOnRamp.sol";
import {BaseTest} from "../../BaseTest.t.sol";

contract ExecutorOnRampSetup is BaseTest {
  ExecutorOnRamp internal s_executorOnRamp;
  address internal s_ccvOnRamp;
  address internal s_feeQuoter; // Not quoting fees for the executor yet, so we don't need FeeQuoterSetup
  address internal s_feeAggregator;

  function setUp() public virtual override {
    super.setUp();
    s_ccvOnRamp = makeAddr("CCVOnRamp");
    s_feeQuoter = makeAddr("FeeQuoter");
    s_feeAggregator = makeAddr("FeeAggregator");

    s_executorOnRamp = new ExecutorOnRamp(
      ExecutorOnRamp.DynamicConfig({
        feeQuoter: s_feeQuoter,
        feeAggregator: s_feeAggregator,
        maxPossibleCCVsPerMsg: 10,
        maxRequiredCCVsPerMsg: 5
      })
    );
  }
}
