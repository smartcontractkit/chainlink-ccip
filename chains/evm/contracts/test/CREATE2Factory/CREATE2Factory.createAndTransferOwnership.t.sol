// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IOwnable} from "@chainlink/contracts/src/v0.8/shared/interfaces/IOwnable.sol";

import {CREATE2Factory} from "../../CREATE2Factory.sol";
import {CREATE2FactorySetup} from "./CREATE2FactorySetup.t.sol";

import {Create2} from "@openzeppelin/contracts@5.0.2/utils/Create2.sol";

contract CREATE2Factory_createAndTransferOwnership is CREATE2FactorySetup {
  function test_createAndTransferOwnership() public {
    address predictedAddress = s_create2Factory.computeAddress(getStorageCreationCode(1), SALT);

    vm.startPrank(s_allowedCaller);
    vm.expectEmit();
    emit CREATE2Factory.ContractDeployed(predictedAddress);
    vm.expectCall(predictedAddress, 0, abi.encodeWithSelector(IOwnable.transferOwnership.selector, OWNER));
    address deployedAddress = s_create2Factory.createAndTransferOwnership(getStorageCreationCode(1), SALT, OWNER);

    assertEq(deployedAddress, predictedAddress);

    vm.startPrank(OWNER);
    IOwnable(deployedAddress).acceptOwnership();
    assertEq(IOwnable(deployedAddress).owner(), OWNER);
  }

  function test_createAndTransferOwnership_RevertWhen_CallerNotAllowed() public {
    vm.startPrank(s_invalidCaller);
    vm.expectRevert(abi.encodeWithSelector(CREATE2Factory.CallerNotAllowed.selector, s_invalidCaller));
    s_create2Factory.createAndTransferOwnership(getStorageCreationCode(1), SALT, OWNER);
  }

  function test_createAndTransferOwnership_RevertWhen_Create2FailedDeployment() public {
    vm.startPrank(s_allowedCaller);
    vm.expectRevert(abi.encodeWithSelector(Create2.Create2FailedDeployment.selector));
    s_create2Factory.createAndTransferOwnership(getStorageCreationCode(0), SALT, OWNER);
  }
}
