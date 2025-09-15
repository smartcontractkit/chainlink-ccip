// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OwnableRampProxy} from "../../OwnableRampProxy.sol";
import {RampProxy} from "../../RampProxy.sol";
import {BaseTest} from "../BaseTest.t.sol";
import {MockCCVOnRamp} from "../mocks/MockCCVOnRamp.sol";

contract OwnableRampProxySetup is BaseTest {
  OwnableRampProxy internal s_rampProxy;
  MockCCVOnRamp internal s_mockCCVOnRamp;

  uint64 internal constant REMOTE_CHAIN_SELECTOR = 1111;

  function setUp() public override {
    super.setUp();

    s_rampProxy = new OwnableRampProxy();
    s_mockCCVOnRamp = new MockCCVOnRamp("");

    RampProxy.SetRampsArgs[] memory ramps = new RampProxy.SetRampsArgs[](1);
    ramps[0] =
      RampProxy.SetRampsArgs({remoteChainSelector: REMOTE_CHAIN_SELECTOR, rampAddress: address(s_mockCCVOnRamp)});
    s_rampProxy.setRamps(ramps);
  }
}
