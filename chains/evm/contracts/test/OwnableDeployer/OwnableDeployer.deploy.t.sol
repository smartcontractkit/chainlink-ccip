// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IOwnable} from "@chainlink/contracts/src/v0.8/shared/interfaces/IOwnable.sol";

import {OwnableDeployerSetup} from "./OwnableDeployerSetup.t.sol";

contract OwnableDeployer_deploy is OwnableDeployerSetup {
  function test_deploy() public {
    address predictedAddress = s_ownableDeployer.computeAddress(OWNER, s_initCode, SALT);

    address deployedAddress = s_ownableDeployer.deployAndTransferOwnership(s_initCode, SALT);
    assertEq(deployedAddress, predictedAddress);

    // Verify that the deployer can accept ownership.
    IOwnable(deployedAddress).acceptOwnership();
    assertEq(IOwnable(deployedAddress).owner(), OWNER);
  }
}
