// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IOwnable} from "@chainlink/contracts/src/v0.8/shared/interfaces/IOwnable.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Create2} from "@openzeppelin/contracts@5.0.2/utils/Create2.sol";

/// @notice A contract for deploying Ownable contracts with deterministic addresses via CREATE2.
/// @dev This contract is used for deploying static, user-facing contracts (i.e. a proxy for a CCV)
/// such that the addresses are the same across multiple chains. This contract transfers ownership of
/// deployed contracts to the caller, ensuring that the caller can accept ownership. It does not accept
/// a desired owner parameter to reduce opportunity for misconfiguration.
contract OwnableDeployer is ITypeAndVersion {
  string public constant typeAndVersion = "DeterministicDeployer 1.7.0";

  /// @notice Deploys and transfers ownership of a contract with the given init code and salt.
  /// @dev The deployed address is deterministic based on the deployer address, salt, and init code.
  /// This method does not supmport deploying contracts with payable constructors (sets amount to 0).
  /// Thi method assumes that the deployer of the contract will be set as its owner upon construction.
  /// @return contractAddress The address of the contract deployed.
  function deployAndTransferOwnership(bytes memory initCode, bytes32 salt) external returns (address) {
    address contractAddress = Create2.deploy(0, salt, initCode);
    IOwnable(contractAddress).transferOwnership(msg.sender);

    return contractAddress;
  }
}
