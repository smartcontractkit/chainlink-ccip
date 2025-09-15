// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RampProxy} from "../../RampProxy.sol";
import {RampProxySetup} from "./RampProxySetup.t.sol";

contract RampProxy_setRamps is RampProxySetup {
  function test_setRamps_UpdateRamp() public {
    RampProxy.SetRampsArgs[] memory ramps = new RampProxy.SetRampsArgs[](1);
    ramps[0] =
      RampProxy.SetRampsArgs({remoteChainSelector: REMOTE_CHAIN_SELECTOR, rampAddress: address(s_otherMockCCVOnRamp)});

    vm.expectEmit();
    emit RampProxy.RampUpdated(REMOTE_CHAIN_SELECTOR, address(s_mockCCVOnRamp), address(s_otherMockCCVOnRamp));
    s_rampProxy.setRamps(ramps);

    assertEq(address(s_otherMockCCVOnRamp), s_rampProxy.getRamp(REMOTE_CHAIN_SELECTOR));
  }

  function test_setRamps_RemoveRamp() public {
    RampProxy.SetRampsArgs[] memory ramps = new RampProxy.SetRampsArgs[](1);
    ramps[0] = RampProxy.SetRampsArgs({remoteChainSelector: REMOTE_CHAIN_SELECTOR, rampAddress: address(0)});

    vm.expectEmit();
    emit RampProxy.RampUpdated(REMOTE_CHAIN_SELECTOR, address(s_mockCCVOnRamp), address(0));
    s_rampProxy.setRamps(ramps);

    assertEq(address(0), s_rampProxy.getRamp(REMOTE_CHAIN_SELECTOR));
  }

  function test_setRamps_RevertWhen_InvalidRemoteChainSelector() public {
    RampProxy.SetRampsArgs[] memory ramps = new RampProxy.SetRampsArgs[](1);
    ramps[0] = RampProxy.SetRampsArgs({remoteChainSelector: 0, rampAddress: address(s_otherMockCCVOnRamp)});

    vm.expectRevert(abi.encodeWithSelector(RampProxy.InvalidRemoteChainSelector.selector, 0));
    s_rampProxy.setRamps(ramps);
  }

  function test_setRamps_RevertWhen_InvalidRampAddress() public {
    address invalidRamp = makeAddr("InvalidRamp"); // no code at this address
    RampProxy.SetRampsArgs[] memory ramps = new RampProxy.SetRampsArgs[](1);
    ramps[0] = RampProxy.SetRampsArgs({remoteChainSelector: REMOTE_CHAIN_SELECTOR, rampAddress: invalidRamp});

    vm.expectRevert(abi.encodeWithSelector(RampProxy.InvalidRampAddress.selector, invalidRamp));
    s_rampProxy.setRamps(ramps);
  }
}
