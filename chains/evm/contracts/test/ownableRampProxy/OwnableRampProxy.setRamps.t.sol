// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RampProxy} from "../../RampProxy.sol";
import {OwnableRampProxySetup} from "./OwnableRampProxySetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract OwnableRampProxy_setRamps is OwnableRampProxySetup {
  function test_setRamps_RevertWhen_NotOwner() public {
    RampProxy.SetRampsArgs[] memory ramps = new RampProxy.SetRampsArgs[](1);
    ramps[0] = RampProxy.SetRampsArgs({remoteChainSelector: REMOTE_CHAIN_SELECTOR, rampAddress: makeAddr("NewRamp")});

    vm.startPrank(makeAddr("NotOwner"));
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_rampProxy.setRamps(ramps);
  }
}
