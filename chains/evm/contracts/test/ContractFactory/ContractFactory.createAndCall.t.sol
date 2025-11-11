// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {IOwnable} from "@chainlink/contracts/src/v0.8/shared/interfaces/IOwnable.sol";

import {ContractFactory} from "../../ContractFactory.sol";
import {BaseTest} from "../BaseTest.t.sol";

contract ContractFactory_createAndCall is BaseTest {
  ContractFactory internal s_contractFactory;
  bytes internal s_creationCode;
  bytes32 internal constant SALT = keccak256("SALT");

  function setUp() public virtual override {
    super.setUp();

    s_creationCode = abi.encodePacked(type(Ownable2StepMsgSender).creationCode);
    s_contractFactory = new ContractFactory();
  }

  function test_createAndCall_noCalls() public {
    address predictedAddress = s_contractFactory.computeAddress(s_creationCode, SALT);
    address deployedAddress = s_contractFactory.createAndCall(s_creationCode, SALT, new bytes[](0));
    assertEq(deployedAddress, predictedAddress);
  }

  function test_createAndCall_transferOwnership() public {
    bytes[] memory calls = new bytes[](1);
    calls[0] = abi.encodeWithSelector(IOwnable.transferOwnership.selector, OWNER);

    address predictedAddress = s_contractFactory.computeAddress(s_creationCode, SALT);
    address deployedAddress = s_contractFactory.createAndCall(s_creationCode, SALT, calls);
    assertEq(deployedAddress, predictedAddress);

    IOwnable(deployedAddress).acceptOwnership();
    assertEq(IOwnable(deployedAddress).owner(), OWNER);
  }
}
