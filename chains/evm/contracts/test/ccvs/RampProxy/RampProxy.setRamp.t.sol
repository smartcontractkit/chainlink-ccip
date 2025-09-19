// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RampProxy} from "../../../ccvs/RampProxy.sol";
import {RampProxySetup} from "./RampProxySetup.t.sol";

contract RampProxy_setRamp is RampProxySetup {
  function test_setRamp() public {
    address newRamp = makeAddr("NewRamp");
    s_rampProxy.setRamp(newRamp);
    assertEq(s_rampProxy.s_ramp(), newRamp);
  }

  function test_setRamps_RevertWhen_ZeroAddressNotAllowed() public {
    vm.expectRevert(RampProxy.ZeroAddressNotAllowed.selector);
    s_rampProxy.setRamp(address(0));
  }
}
