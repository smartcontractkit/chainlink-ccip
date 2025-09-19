// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseTest} from "../../BaseTest.t.sol";

import {MockRampImplementation} from "../../mocks/MockRampImplementation.sol";
import {MockRampProxy} from "../../mocks/MockRampProxy.sol";

contract RampProxySetup is BaseTest {
  MockRampProxy internal s_rampProxy;
  MockRampImplementation internal s_rampImpl;

  uint8 internal constant EXPECTED_VALUE = 1;

  function setUp() public override {
    super.setUp();

    s_rampImpl = new MockRampImplementation(EXPECTED_VALUE);
    s_rampProxy = new MockRampProxy(address(s_rampImpl));
  }
}
