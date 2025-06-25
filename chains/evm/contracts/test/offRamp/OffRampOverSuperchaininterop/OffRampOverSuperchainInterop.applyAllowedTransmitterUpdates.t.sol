// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";
import {OffRampOverSuperchainInteropSetup} from "./OffRampOverSuperchainInteropSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract OffRampOverSuperchainInterop_applyAllowedTransmitterUpdates is OffRampOverSuperchainInteropSetup {
  function test_applyAllowedTransmitterUpdates_AddTransmitters() public {
    address newTransmitter1 = makeAddr("newTransmitter1");
    address newTransmitter2 = makeAddr("newTransmitter2");
    
    address[] memory transmittersToAdd = new address[](2);
    transmittersToAdd[0] = newTransmitter1;
    transmittersToAdd[1] = newTransmitter2;
    
    vm.expectEmit();
    emit OffRampOverSuperchainInterop.AllowedTransmitterAdded(newTransmitter1);
    vm.expectEmit();
    emit OffRampOverSuperchainInterop.AllowedTransmitterAdded(newTransmitter2);
    
    s_offRamp.applyAllowedTransmitterUpdates(new address[](0), transmittersToAdd);
    
    // Verify transmitters were added
    address[] memory transmitters = s_offRamp.getAllAllowedTransmittes();
    assertEq(transmitters.length, 4); // 2 original + 2 new
  }

  function test_applyAllowedTransmitterUpdates_RemoveTransmitters() public {
    address[] memory transmittersToRemove = new address[](1);
    transmittersToRemove[0] = s_transmitter1;
    
    vm.expectEmit();
    emit OffRampOverSuperchainInterop.AllowedTransmitterRemoved(s_transmitter1);
    
    s_offRamp.applyAllowedTransmitterUpdates(transmittersToRemove, new address[](0));
    
    // Verify transmitter was removed
    address[] memory transmitters = s_offRamp.getAllAllowedTransmittes();
    assertEq(transmitters.length, 1);
    assertEq(transmitters[0], s_transmitter2);
  }

  function test_applyAllowedTransmitterUpdates_AddAndRemove() public {
    address newTransmitter = makeAddr("newTransmitter");
    
    address[] memory transmittersToRemove = new address[](1);
    transmittersToRemove[0] = s_transmitter1;
    
    address[] memory transmittersToAdd = new address[](1);
    transmittersToAdd[0] = newTransmitter;
    
    vm.expectEmit();
    emit OffRampOverSuperchainInterop.AllowedTransmitterRemoved(s_transmitter1);
    vm.expectEmit();
    emit OffRampOverSuperchainInterop.AllowedTransmitterAdded(newTransmitter);
    
    s_offRamp.applyAllowedTransmitterUpdates(transmittersToRemove, transmittersToAdd);
    
    // Verify changes
    address[] memory transmitters = s_offRamp.getAllAllowedTransmittes();
    assertEq(transmitters.length, 2);
  }

  function test_applyAllowedTransmitterUpdates_RevertWhen_ZeroAddress() public {
    address[] memory transmittersToAdd = new address[](1);
    transmittersToAdd[0] = address(0);
    
    vm.expectRevert(OffRampOverSuperchainInterop.ZeroAddressNotAllowed.selector);
    s_offRamp.applyAllowedTransmitterUpdates(new address[](0), transmittersToAdd);
  }

  function test_applyAllowedTransmitterUpdates_RevertWhen_NotOwner() public {
    address[] memory transmittersToAdd = new address[](1);
    transmittersToAdd[0] = makeAddr("newTransmitter");
    
    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_offRamp.applyAllowedTransmitterUpdates(new address[](0), transmittersToAdd);
  }
}