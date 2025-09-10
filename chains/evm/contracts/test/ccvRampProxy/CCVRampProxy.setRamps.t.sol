// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVRampProxy} from "../../CCVRampProxy.sol";
import {CCVRamp} from "../../libraries/CCVRamp.sol";
import {CCVRampProxySetup} from "./CCVRampProxySetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CCVRampProxy_setRamps is CCVRampProxySetup {
  function test_setRamps() public {
    address newRamp = makeAddr("NewRamp");
    CCVRampProxy.SetRampsArgs[] memory ramps = new CCVRampProxy.SetRampsArgs[](1);
    ramps[0] =
      CCVRampProxy.SetRampsArgs({remoteChainSelector: REMOTE_CHAIN_SELECTOR, version: CCVRamp.V1, addr: newRamp});

    vm.expectEmit();
    emit CCVRampProxy.RampSet(REMOTE_CHAIN_SELECTOR, CCVRamp.V1, newRamp);
    s_ccvRampProxy.setRamps(ramps);

    assertEq(newRamp, s_ccvRampProxy.getRamp(REMOTE_CHAIN_SELECTOR, CCVRamp.V1));
  }

  function test_setRamps_RevertWhen_NotOwner() public {
    CCVRampProxy.SetRampsArgs[] memory ramps = new CCVRampProxy.SetRampsArgs[](1);
    ramps[0] = CCVRampProxy.SetRampsArgs({
      remoteChainSelector: REMOTE_CHAIN_SELECTOR,
      version: CCVRamp.V1,
      addr: makeAddr("NewRamp")
    });

    vm.startPrank(makeAddr("NotOwner"));
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_ccvRampProxy.setRamps(ramps);
  }

  function test_setRamps_RevertWhen_InvalidRemoteChainSelector() public {
    CCVRampProxy.SetRampsArgs[] memory ramps = new CCVRampProxy.SetRampsArgs[](1);
    ramps[0] = CCVRampProxy.SetRampsArgs({remoteChainSelector: 0, version: CCVRamp.V1, addr: makeAddr("NewRamp")});

    vm.expectRevert(abi.encodeWithSelector(CCVRampProxy.InvalidRemoteChainSelector.selector, 0));
    s_ccvRampProxy.setRamps(ramps);
  }

  function test_setRamps_RevertWhen_InvalidVersion() public {
    CCVRampProxy.SetRampsArgs[] memory ramps = new CCVRampProxy.SetRampsArgs[](1);
    ramps[0] = CCVRampProxy.SetRampsArgs({
      remoteChainSelector: REMOTE_CHAIN_SELECTOR,
      version: bytes32(0),
      addr: makeAddr("NewRamp")
    });

    vm.expectRevert(abi.encodeWithSelector(CCVRampProxy.InvalidVersion.selector, bytes32(0)));
    s_ccvRampProxy.setRamps(ramps);
  }

  function test_setRamps_RevertWhen_InvalidRampAddress() public {
    CCVRampProxy.SetRampsArgs[] memory ramps = new CCVRampProxy.SetRampsArgs[](1);
    ramps[0] =
      CCVRampProxy.SetRampsArgs({remoteChainSelector: REMOTE_CHAIN_SELECTOR, version: CCVRamp.V1, addr: address(0)});

    vm.expectRevert(abi.encodeWithSelector(CCVRampProxy.InvalidRampAddress.selector, address(0)));
    s_ccvRampProxy.setRamps(ramps);
  }
}
