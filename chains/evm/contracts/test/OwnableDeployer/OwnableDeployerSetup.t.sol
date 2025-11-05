// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OwnableDeployer} from "../../OwnableDeployer.sol";
import {BaseTest} from "../BaseTest.t.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

contract OwnableDeployerSetup is BaseTest {
  OwnableDeployer internal s_ownableDeployer;
  bytes internal s_initCode;

  bytes32 internal constant SALT = keccak256("SALT");

  function setUp() public virtual override {
    super.setUp();

    s_initCode = abi.encodePacked(type(Ownable2StepMsgSender).creationCode);
    s_ownableDeployer = new OwnableDeployer();
  }
}
