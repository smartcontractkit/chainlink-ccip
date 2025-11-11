// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IOwnable} from "@chainlink/contracts/src/v0.8/shared/interfaces/IOwnable.sol";

import {ContractFactory} from "../../ContractFactory.sol";
import {ContractFactorySetup, Storage} from "./ContractFactorySetup.t.sol";

import {Create2} from "@openzeppelin/contracts@5.0.2/utils/Create2.sol";

contract ContractFactory_createAndCall is ContractFactorySetup {
  function test_createAndCall_NoCalls() public {
    address predictedAddress = s_contractFactory.computeAddress(getStorageCreationCode(1), SALT);
    vm.startPrank(s_allowedCaller);
    vm.expectEmit();
    emit ContractFactory.ContractDeployed(predictedAddress);
    address deployedAddress = s_contractFactory.createAndCall(getStorageCreationCode(1), SALT, new bytes[](0));
    assertEq(deployedAddress, predictedAddress);
  }

  function test_createAndCall_SingleCall() public {
    bytes[] memory calls = new bytes[](1);
    calls[0] = abi.encodeWithSelector(IOwnable.transferOwnership.selector, OWNER);

    address predictedAddress = s_contractFactory.computeAddress(getStorageCreationCode(1), SALT);
    vm.startPrank(s_allowedCaller);
    vm.expectEmit();
    emit ContractFactory.ContractDeployed(predictedAddress);
    address deployedAddress = s_contractFactory.createAndCall(getStorageCreationCode(1), SALT, calls);
    assertEq(deployedAddress, predictedAddress);

    vm.startPrank(OWNER);
    IOwnable(deployedAddress).acceptOwnership();
    assertEq(IOwnable(deployedAddress).owner(), OWNER);
  }

  function test_createAndCall_MultipleCalls() public {
    bytes[] memory calls = new bytes[](2);
    calls[0] = abi.encodeWithSelector(IOwnable.transferOwnership.selector, OWNER);
    calls[1] = abi.encodeWithSelector(Storage.setValue.selector, 2);

    address predictedAddress = s_contractFactory.computeAddress(getStorageCreationCode(1), SALT);
    vm.startPrank(s_allowedCaller);
    vm.expectEmit();
    emit ContractFactory.ContractDeployed(predictedAddress);
    address deployedAddress = s_contractFactory.createAndCall(getStorageCreationCode(1), SALT, calls);
    assertEq(deployedAddress, predictedAddress);

    vm.startPrank(OWNER);
    IOwnable(deployedAddress).acceptOwnership();
    assertEq(IOwnable(deployedAddress).owner(), OWNER);
    assertEq(Storage(deployedAddress).getValue(), 2);
  }

  function test_createAndCall_RevertWhen_CallerNotAllowed() public {
    vm.startPrank(s_invalidCaller);
    vm.expectRevert(abi.encodeWithSelector(ContractFactory.CallerNotAllowed.selector, s_invalidCaller));
    s_contractFactory.createAndCall(getStorageCreationCode(1), SALT, new bytes[](0));
  }

  function test_createAndCall_RevertWhen_Create2FailedDeployment() public {
    vm.startPrank(s_allowedCaller);
    vm.expectRevert(abi.encodeWithSelector(Create2.Create2FailedDeployment.selector));
    s_contractFactory.createAndCall(getStorageCreationCode(0), SALT, new bytes[](0));
  }

  function test_createAndCall_RevertWhen_CallFailed() public {
    bytes[] memory calls = new bytes[](1);
    calls[0] = abi.encodeWithSelector(Storage.setValue.selector, 0);

    vm.startPrank(s_allowedCaller);
    vm.expectRevert(
      abi.encodeWithSelector(
        ContractFactory.CallFailed.selector, 0, abi.encodeWithSelector(Storage.InvalidValue.selector)
      )
    );
    s_contractFactory.createAndCall(getStorageCreationCode(1), SALT, calls);
  }
}
