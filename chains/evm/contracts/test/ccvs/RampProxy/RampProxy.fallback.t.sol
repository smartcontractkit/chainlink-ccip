// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RampProxySetup} from "./RampProxySetup.t.sol";

interface IMockRamp {
  function getValue() external returns (uint8);
}

contract RampProxy_fallback is RampProxySetup {
  function test_fallback() public {
    bytes memory revertData = "0x12345678";
    vm.mockCallRevert(s_rampProxy.s_ramp(), abi.encodeWithSelector(IMockRamp.getValue.selector), revertData);
    vm.expectRevert(revertData);
    IMockRamp(address(s_rampProxy)).getValue();

    uint8 expectedValue = 1;
    vm.mockCall(s_rampProxy.s_ramp(), abi.encodeWithSelector(IMockRamp.getValue.selector), abi.encode(1));
    uint8 value = IMockRamp(address(s_rampProxy)).getValue();
    assertEq(value, expectedValue);
  }
}
