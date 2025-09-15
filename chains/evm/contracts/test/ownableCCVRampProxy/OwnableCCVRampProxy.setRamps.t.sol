// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVRampProxy} from "../../CCVRampProxy.sol";
import {CCVRamp} from "../../libraries/CCVRamp.sol";
import {OwnableCCVRampProxySetup} from "./OwnableCCVRampProxySetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract OwnableCCVRampProxy_setRamps is OwnableCCVRampProxySetup {
  function test_setRamps_RevertWhen_NotOwner() public {
    CCVRampProxy.SetRampsArgs[] memory ramps = new CCVRampProxy.SetRampsArgs[](1);
    ramps[0] = CCVRampProxy.SetRampsArgs({
      remoteChainSelector: REMOTE_CHAIN_SELECTOR,
      version: CCVRamp.V1,
      rampAddress: makeAddr("NewRamp")
    });

    vm.startPrank(makeAddr("NotOwner"));
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_ccvRampProxy.setRamps(ramps);
  }
}
