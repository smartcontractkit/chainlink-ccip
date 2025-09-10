// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVRampProxy} from "../../CCVRampProxy.sol";
import {CCVRamp} from "../../libraries/CCVRamp.sol";
import {CCVRampProxySetup} from "./CCVRampProxySetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CCVRampProxy_setRamp is CCVRampProxySetup {
  function test_setRamp() public {
    address newRamp = makeAddr("NewRamp");

    vm.expectEmit();
    emit CCVRampProxy.RampSet(REMOTE_CHAIN_SELECTOR, CCVRamp.V1, newRamp);
    s_ccvRampProxy.setRamp(REMOTE_CHAIN_SELECTOR, CCVRamp.V1, newRamp);

    assertEq(newRamp, s_ccvRampProxy.getRamp(REMOTE_CHAIN_SELECTOR, CCVRamp.V1));
  }

  function test_setRamp_RevertWhen_NotOwner() public {
    vm.startPrank(makeAddr("NotOwner"));
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_ccvRampProxy.setRamp(REMOTE_CHAIN_SELECTOR, CCVRamp.V1, makeAddr("NewRamp"));
  }

  function test_setRamp_RevertWhen_InvalidRemoteChainSelector() public {
    vm.expectRevert(abi.encodeWithSelector(CCVRampProxy.InvalidRemoteChainSelector.selector, 0));
    s_ccvRampProxy.setRamp(0, CCVRamp.V1, makeAddr("NewRamp"));
  }

  function test_setRamp_RevertWhen_InvalidVersion() public {
    vm.expectRevert(abi.encodeWithSelector(CCVRampProxy.InvalidVersion.selector, bytes32(0)));
    s_ccvRampProxy.setRamp(REMOTE_CHAIN_SELECTOR, bytes32(0), makeAddr("NewRamp"));
  }

  function test_setRamp_RevertWhen_InvalidRampAddress() public {
    vm.expectRevert(abi.encodeWithSelector(CCVRampProxy.InvalidRampAddress.selector, address(0)));
    s_ccvRampProxy.setRamp(REMOTE_CHAIN_SELECTOR, CCVRamp.V1, address(0));
  }
}
