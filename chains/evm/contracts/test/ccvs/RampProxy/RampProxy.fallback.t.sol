// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RampProxySetup, MockImplementation} from "./RampProxySetup.t.sol";

contract RampProxy_fallback is RampProxySetup {
  function test_fallback() public {
    uint8 value = MockImplementation(address(s_rampProxy)).getValue();
    assertEq(value, CURR_EXPECTED_VALUE);

    vm.expectRevert(abi.encodeWithSelector(MockImplementation.Failed.selector));
    MockImplementation(address(s_rampProxy)).revertWithError();
  }
}
