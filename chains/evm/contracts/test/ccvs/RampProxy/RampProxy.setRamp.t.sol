// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RampProxy} from "../../../ccvs/RampProxy.sol";
import {RampProxySetup} from "./RampProxySetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract RampProxy_setRamp is RampProxySetup {
  function test_setRamp() public {
    address newRamp = makeAddr("NewRamp");
    s_rampProxy.setRamp(newRamp);
    assertEq(s_rampProxy.s_ramp(), newRamp);
  }

  function test_setRamp_RevertWhen_ZeroAddressNotAllowed() public {
    vm.expectRevert(RampProxy.ZeroAddressNotAllowed.selector);
    s_rampProxy.setRamp(address(0));
  }

  function test_setRamp_RevertWhen_NotOwner() public {
    vm.stopPrank();
    address notOwner = makeAddr("NotOwner");
    vm.startPrank(notOwner);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_rampProxy.setRamp(makeAddr("NewRamp"));
  }
}
