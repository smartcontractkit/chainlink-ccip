// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCIPClientExampleWithCCVs} from "../../../applications/CCIPClientExampleWithCCVs.sol";
import {OnRampSetup} from "../../onRamp/OnRamp/OnRampSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract CCIPClientExampleWithCCVsSetup is OnRampSetup {
  CCIPClientExampleWithCCVs internal s_client;

  function setUp() public virtual override {
    super.setUp();

    s_client = new CCIPClientExampleWithCCVs(s_destRouter, IERC20(s_destFeeToken));
  }
}
