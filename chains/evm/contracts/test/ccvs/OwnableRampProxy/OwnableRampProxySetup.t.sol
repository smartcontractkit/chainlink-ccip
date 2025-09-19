// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OwnableRampProxy} from "../../../ccvs/OwnableRampProxy.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {MockCCVOnRamp} from "../../mocks/MockCCVOnRamp.sol";

contract OwnableRampProxySetup is BaseTest {
  OwnableRampProxy internal s_rampProxy;
  MockCCVOnRamp internal s_mockCCVOnRamp;
  MockCCVOnRamp internal s_newMockCCVOnRamp;

  uint64 internal constant REMOTE_CHAIN_SELECTOR = 1111;

  function setUp() public override {
    super.setUp();

    s_mockCCVOnRamp = new MockCCVOnRamp("");
    s_rampProxy = new OwnableRampProxy(address(s_mockCCVOnRamp));

    s_newMockCCVOnRamp = new MockCCVOnRamp("");
  }
}
