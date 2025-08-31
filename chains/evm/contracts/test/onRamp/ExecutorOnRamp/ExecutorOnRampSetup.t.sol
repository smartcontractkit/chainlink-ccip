// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ExecutorOnRamp} from "../../../onRamp/ExecutorOnRamp.sol";

import {Test} from "forge-std/Test.sol";

contract ExecutorOnRampSetup is Test {
  ExecutorOnRamp internal s_executorOnRamp;

  function setUp() public {
    s_executorOnRamp = new ExecutorOnRamp(ExecutorOnRamp.DynamicConfig({maxCCVsPerMsg: 1}));
  }
}
