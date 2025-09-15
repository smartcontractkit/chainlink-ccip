// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RampProxy} from "../../RampProxy.sol";
import {BaseTest} from "../BaseTest.t.sol";
import {MockCCVOnRamp} from "../mocks/MockCCVOnRamp.sol";
import {MockRampProxy} from "../mocks/MockRampProxy.sol";

contract RampProxySetup is BaseTest {
  MockRampProxy internal s_rampProxy;
  MockCCVOnRamp internal s_mockCCVOnRamp;
  MockCCVOnRamp internal s_otherMockCCVOnRamp;

  uint64 internal constant REMOTE_CHAIN_SELECTOR = 1111;
  uint64 internal constant UNSUPPORTED_REMOTE_CHAIN_SELECTOR = 2222;
  bytes internal constant EXPECTED_VERIFIER_RESULT = "Hello, World!";

  function setUp() public override {
    super.setUp();

    s_rampProxy = new MockRampProxy();
    s_mockCCVOnRamp = new MockCCVOnRamp(EXPECTED_VERIFIER_RESULT);
    s_otherMockCCVOnRamp = new MockCCVOnRamp(EXPECTED_VERIFIER_RESULT);

    RampProxy.SetRampsArgs[] memory ramps = new RampProxy.SetRampsArgs[](1);
    ramps[0] =
      RampProxy.SetRampsArgs({remoteChainSelector: REMOTE_CHAIN_SELECTOR, rampAddress: address(s_mockCCVOnRamp)});
    s_rampProxy.setRamps(ramps);
  }
}
