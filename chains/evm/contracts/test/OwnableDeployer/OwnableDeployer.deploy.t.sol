// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IOwnable} from "@chainlink/contracts/src/v0.8/shared/interfaces/IOwnable.sol";

import {OwnableDeployer} from "../../OwnableDeployer.sol";
import {OwnableDeployerSetup} from "./OwnableDeployerSetup.t.sol";

import {Create2} from "@openzeppelin/contracts@5.0.2/utils/Create2.sol";

contract OwnableDeployer_deploy is OwnableDeployerSetup {
  function test_deploy() public {
    address predictedAddress = Create2.computeAddress(SALT, keccak256(s_initCode), address(s_ownableDeployer));

    address deployedAddress = s_ownableDeployer.deployAndTransferOwnership(s_initCode, SALT);
    assertEq(deployedAddress, predictedAddress);

    // Verify that the deployer can accept ownership.
    IOwnable(deployedAddress).acceptOwnership();
    assertEq(IOwnable(deployedAddress).owner(), OWNER);
  }
}
