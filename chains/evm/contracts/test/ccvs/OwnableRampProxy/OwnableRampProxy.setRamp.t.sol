// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OwnableRampProxySetup} from "./OwnableRampProxySetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract OwnableRampProxy_setRamps is OwnableRampProxySetup {
  function test_setRamps() public {
    s_rampProxy.setRamp(address(s_newMockCCVOnRamp));
    assertEq(s_rampProxy.s_ramp(), address(s_newMockCCVOnRamp));
  }

  function test_setRamps_RevertWhen_NotOwner() public {
    vm.stopPrank();
    address notOwner = makeAddr("NotOwner");
    vm.startPrank(notOwner);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_rampProxy.setRamp(address(s_newMockCCVOnRamp));
  }
}
