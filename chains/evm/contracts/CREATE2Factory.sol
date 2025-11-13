// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IOwnable} from "@chainlink/contracts/src/v0.8/shared/interfaces/IOwnable.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {Create2} from "@openzeppelin/contracts@5.0.2/utils/Create2.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.0.2/utils/structs/EnumerableSet.sol";

/// @notice A contract for deploying and configuring contracts via CREATE2.
/// @dev This contract is used to deploy static, user-facing contracts (e.g. proxies)
/// such that the resulting addresses are the same across multiple chains. Optionally, you can provide
/// one or more initialization calls (i.e. transfer ownership, configure roles) to perform post-deployment.
/// To achieve deterministic addresses across chains, this contract must be deployed with a reserved key.
/// This is because the factory address is used in the CREATE2 address computation.
contract CREATE2Factory is ITypeAndVersion, Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;

  error CallFailed(uint256 index, bytes result);
  error CallerNotAllowed(address caller);

  event CallerAdded(address indexed caller);
  event CallerRemoved(address indexed caller);
  event ContractDeployed(address indexed contractAddress);

  string public constant override typeAndVersion = "CREATE2Factory 1.7.0";

  /// @notice Addresses that are allowed to call createAndCall.
  EnumerableSet.AddressSet private s_allowList;

  constructor(
    address[] memory allowList
  ) {
    _applyAllowListUpdates(new address[](0), allowList);
  }

  /// @notice Deploys a contract with the given creation code and salt and optionally calls it.
  /// @dev The deployed address is deterministic based on address(this), salt, and creation code.
  /// Creation code includes constructor arguments, which must be taken into account if address parity is desired.
  /// Example: For token pools to have the same address across chains, their tokens must also have the same address across chains.
  /// This method does not support deploying contracts with payable constructors (sets amount to 0).
  /// This function is allowlisted to prevent unexpected accounts from claiming important addresses on new chains.
  /// Concatenating msg.sender with the salt is an alternative way to approach this problem, but prevents the ability
  /// to rotate keys. Taking that approach, you would need to use the same key for createAndCall on every chain in perpetuity.
  /// @param creationCode The creation code of the contract to deploy.
  /// @param salt The salt used to ensure a unique deployment.
  /// @param calls Any calls to perform post-deployment.
  /// @return contractAddress The address of the contract deployed.
  function createAndCall(bytes calldata creationCode, bytes32 salt, bytes[] memory calls) external returns (address) {
    return _createAndCall(creationCode, salt, calls);
  }

  /// @notice Internal helper for createAndCall.
  /// @param creationCode The creation code of the contract to deploy.
  /// @param salt The salt used to ensure a unique deployment.
  /// @param calls Any calls to perform post-deployment.
  /// @return contractAddress The address of the contract deployed.
  function _createAndCall(bytes calldata creationCode, bytes32 salt, bytes[] memory calls) internal returns (address) {
    if (!s_allowList.contains(msg.sender)) {
      revert CallerNotAllowed(msg.sender);
    }

    address contractAddress = Create2.deploy(0, salt, creationCode);
    emit ContractDeployed(contractAddress);

    for (uint256 i = 0; i < calls.length; ++i) {
      // solhint-disable-next-line avoid-low-level-calls
      (bool success, bytes memory result) = contractAddress.call(calls[i]);
      if (!success) {
        revert CallFailed(i, result);
      }
    }

    return contractAddress;
  }

  /// @notice Deploys a contract with the given creation code and salt and transfers ownership to the given address.
  /// @param creationCode The creation code of the contract to deploy.
  /// @param salt The salt used to ensure a unique deployment.
  /// @param to The address to transfer ownership to.
  /// @return contractAddress The address of the contract deployed.
  function createAndTransferOwnership(bytes calldata creationCode, bytes32 salt, address to) external returns (address) {
    bytes[] memory calls = new bytes[](1);
    calls[0] = abi.encodeWithSelector(IOwnable.transferOwnership.selector, to);
    return _createAndCall(creationCode, salt, calls);
  }

  /// @notice Computes the address of a contract if deployed with the given creation code and salt.
  /// @param creationCode The creation code of the contract.
  /// @param salt The salt used to ensure a unique deployment.
  /// @return contractAddress The address that would result from the deployment.
  function computeAddress(bytes memory creationCode, bytes32 salt) external view returns (address) {
    return Create2.computeAddress(salt, keccak256(creationCode), address(this));
  }

  /// @notice Updates the addresses that are allowed to call createAndCall.
  /// @param adds Array of addresses to add.
  /// @param removes Array of addresses to remove.
  function applyAllowListUpdates(address[] calldata removes, address[] calldata adds) external onlyOwner {
    _applyAllowListUpdates(removes, adds);
  }

  /// @notice Updates the addresses that are allowed to call createAndCall.
  /// @dev Internal helper for applyAllowListUpdates and constructor.
  /// @param adds Array of addresses to add.
  /// @param removes Array of addresses to remove.
  function _applyAllowListUpdates(address[] memory removes, address[] memory adds) internal {
    for (uint256 i = 0; i < removes.length; ++i) {
      if (s_allowList.remove(removes[i])) {
        emit CallerRemoved(removes[i]);
      }
    }

    for (uint256 i = 0; i < adds.length; ++i) {
      if (s_allowList.add(adds[i])) {
        emit CallerAdded(adds[i]);
      }
    }
  }

  /// @notice Returns all addresses that are allowed to call createAndCall.
  function getAllowList() external view returns (address[] memory) {
    return s_allowList.values();
  }
}
