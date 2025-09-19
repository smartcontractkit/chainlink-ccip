// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MockRampImplementation} from "../../mocks/MockRampImplementation.sol";
import {RampProxySetup} from "./RampProxySetup.t.sol";

contract RampProxy_fallback is RampProxySetup {
  function test_fallback() public {
    uint8 value = MockRampImplementation(address(s_rampProxy)).getValue();
    assertEq(value, EXPECTED_VALUE);

    vm.expectRevert(abi.encodeWithSelector(MockRampImplementation.Failed.selector));
    MockRampImplementation(address(s_rampProxy)).revertWithError();
  }
}
