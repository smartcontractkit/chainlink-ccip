// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RampProxy} from "../../../ccvs/RampProxy.sol";
import {BaseTest} from "../../BaseTest.t.sol";

contract RampProxySetup is BaseTest {
  RampProxy internal s_rampProxy;

  function setUp() public override {
    super.setUp();

    s_rampProxy = new RampProxy(makeAddr("MockImpl"));
  }
}
