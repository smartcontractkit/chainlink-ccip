// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Create2} from "@openzeppelin/contracts@5.0.2/utils/Create2.sol";

/// @notice A contract for deploying and configuring contracts via CREATE2.
/// @dev This contract is used to deploy static, user-facing contracts (e.g. proxies)
/// such that the resulting addresses are the same across multiple chains. Optionally, you can provide
/// one or more initialization calls (i.e. transfer ownership, configure roles) to perform post-deployment.
/// To achieve deterministic addresses across chains, this contract must be deployed with a reserved key.
/// This is because the factory address is used in the CREATE2 address computation.
contract ContractFactory is ITypeAndVersion {
  string public constant typeAndVersion = "ContractFactory 1.7.0";

  error CallFailed(uint256 index, bytes result);

  /// @notice Deploys a contract with the given creation code and salt and optionally calls it.
  /// @dev The deployed address is deterministic based on address(this), salt, and creation code.
  /// This method does not support deploying contracts with payable constructors (sets amount to 0).
  /// @param creationCode The creation code of the contract to deploy.
  /// @param salt The salt used to ensure a unique deployment.
  /// @param calls Any calls to perform post-deployment.
  /// @return contractAddress The address of the contract deployed.
  function createAndCall(bytes calldata creationCode, bytes32 salt, bytes[] calldata calls) external returns (address) {
    address contractAddress = Create2.deploy(0, salt, creationCode);

    for (uint256 i = 0; i < calls.length; ++i) {
      // solhint-disable-next-line avoid-low-level-calls
      (bool success, bytes memory result) = contractAddress.call(calls[i]);
      if (!success) {
        revert CallFailed(i, result);
      }
    }

    return contractAddress;
  }

  /// @notice Computes the address of a contract if deployed with the given creation code and salt.
  /// @param creationCode The creation code of the contract.
  /// @param salt The salt used to ensure a unique deployment.
  /// @return contractAddress The address that would result from the deployment.
  function computeAddress(bytes memory creationCode, bytes32 salt) external view returns (address) {
    return Create2.computeAddress(salt, keccak256(creationCode), address(this));
  }
}
