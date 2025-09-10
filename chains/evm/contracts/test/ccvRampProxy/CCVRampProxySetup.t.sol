// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVRampProxy} from "../../CCVRampProxy.sol";

import {CCVRamp} from "../../libraries/CCVRamp.sol";
import {BaseTest} from "../BaseTest.t.sol";
import {MockCCVOnRamp} from "../mocks/MockCCVOnRamp.sol";

contract CCVRampProxySetup is BaseTest {
  CCVRampProxy internal s_ccvRampProxy;
  MockCCVOnRamp internal s_mockCCVOnRamp;

  uint64 internal constant REMOTE_CHAIN_SELECTOR = 1111;

  function setUp() public override {
    super.setUp();

    s_ccvRampProxy = new CCVRampProxy();
    s_mockCCVOnRamp = new MockCCVOnRamp();

    s_ccvRampProxy.setRamp(REMOTE_CHAIN_SELECTOR, CCVRamp.V1, address(s_mockCCVOnRamp));
  }
}
