// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVRampProxy} from "../../CCVRampProxy.sol";
import {CCVRamp} from "../../libraries/CCVRamp.sol";
import {CCVRampProxySetup} from "./CCVRampProxySetup.t.sol";

contract CCVRampProxy_setRamps is CCVRampProxySetup {
  function test_setRamps() public {
    address newRamp = makeAddr("NewRamp");
    CCVRampProxy.SetRampsArgs[] memory ramps = new CCVRampProxy.SetRampsArgs[](1);
    ramps[0] =
      CCVRampProxy.SetRampsArgs({remoteChainSelector: REMOTE_CHAIN_SELECTOR, version: CCVRamp.V1, rampAddress: newRamp});

    vm.expectEmit();
    emit CCVRampProxy.RampSet(REMOTE_CHAIN_SELECTOR, CCVRamp.V1, newRamp);
    s_ccvRampProxy.setRamps(ramps);

    assertEq(newRamp, s_ccvRampProxy.getRamp(REMOTE_CHAIN_SELECTOR, CCVRamp.V1));
  }

  function test_setRamps_RevertWhen_InvalidRemoteChainSelector() public {
    CCVRampProxy.SetRampsArgs[] memory ramps = new CCVRampProxy.SetRampsArgs[](1);
    ramps[0] =
      CCVRampProxy.SetRampsArgs({remoteChainSelector: 0, version: CCVRamp.V1, rampAddress: makeAddr("NewRamp")});

    vm.expectRevert(abi.encodeWithSelector(CCVRampProxy.InvalidRemoteChainSelector.selector, 0));
    s_ccvRampProxy.setRamps(ramps);
  }

  function test_setRamps_RevertWhen_InvalidVersion() public {
    CCVRampProxy.SetRampsArgs[] memory ramps = new CCVRampProxy.SetRampsArgs[](1);
    ramps[0] = CCVRampProxy.SetRampsArgs({
      remoteChainSelector: REMOTE_CHAIN_SELECTOR,
      version: INVALID_RAMP_VERSION,
      rampAddress: makeAddr("NewRamp")
    });

    vm.expectRevert(abi.encodeWithSelector(CCVRampProxy.InvalidVersion.selector, INVALID_RAMP_VERSION));
    s_ccvRampProxy.setRamps(ramps);
  }
}
