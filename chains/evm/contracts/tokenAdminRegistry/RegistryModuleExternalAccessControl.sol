// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IGetCCIPAdmin} from "../interfaces/IGetCCIPAdmin.sol";
import {IOwner} from "../interfaces/IOwner.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";
import {IExternalAccessControlRegistry} from "../interfaces/externalRegistry/IExternalAccessControlRegistry.sol";

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {AccessControl} from "@openzeppelin/contracts@5.3.0/access/AccessControl.sol";

/// @notice Registers token admins in batches using the token's getCCIPAdmin, owner, or DEFAULT_ADMIN_ROLE.
/// @dev AccessControl roles can be stored on the token or in the contract returned by ACCESS_CONTROLLED_REGISTRY.
/// This module must be added to TokenAdminRegistry before it can register tokens. Registration sets the caller as the
/// pending admin. The caller must call acceptAdminRole on TokenAdminRegistry for each token to accept the role.
/// If any token cannot be registered, the call reverts and no tokens are registered. Empty arrays do nothing, and
/// duplicate tokens are processed in order. Pending admins can be replaced, but an admin that has accepted the role
/// cannot be replaced through this module.
contract RegistryModuleExternalAccessControl is ITypeAndVersion {
  error CanOnlySelfRegister(address admin, address token);
  error RequiredRoleNotFound(address msgSender, bytes32 role, address token);
  error AddressZero();

  event AdministratorRegistered(address indexed token, address indexed administrator);

  string public constant override typeAndVersion = "RegistryModuleExternalAccessControl 2.1.0-dev";

  ITokenAdminRegistry private immutable i_tokenAdminRegistry;

  /// @param tokenAdminRegistry The TokenAdminRegistry contract to register token admins in.
  constructor(
    address tokenAdminRegistry
  ) {
    if (tokenAdminRegistry == address(0)) {
      revert AddressZero();
    }
    i_tokenAdminRegistry = ITokenAdminRegistry(tokenAdminRegistry);
  }

  /// @notice Registers the admin of the tokens using the `getCCIPAdmin` method.
  /// @param tokens The tokens to register the admin for.
  /// @dev The caller must be the admin returned by the `getCCIPAdmin` method for every token.
  function registerAdminViaGetCCIPAdmin(
    address[] calldata tokens
  ) external {
    for (uint256 i = 0; i < tokens.length; ++i) {
      address token = tokens[i];
      _registerAdmin(token, IGetCCIPAdmin(token).getCCIPAdmin());
    }
  }

  /// @notice Registers the admin of the tokens using the `owner` method.
  /// @param tokens The tokens to register the admin for.
  /// @dev The caller must be the admin returned by the `owner` method for every token.
  function registerAdminViaOwner(
    address[] calldata tokens
  ) external {
    for (uint256 i = 0; i < tokens.length; ++i) {
      address token = tokens[i];
      _registerAdmin(token, IOwner(token).owner());
    }
  }

  /// @notice Registers the admin of the tokens using OZ's AccessControl DEFAULT_ADMIN_ROLE.
  /// @param tokens The tokens to register the admin for.
  /// @dev The caller must have the DEFAULT_ADMIN_ROLE as defined by each token itself.
  function registerAccessControlDefaultAdmin(
    address[] calldata tokens
  ) external {
    for (uint256 i = 0; i < tokens.length; ++i) {
      address token = tokens[i];
      _registerAccessControlDefaultAdmin(token, AccessControl(token));
    }
  }

  /// @notice Registers the admin of the tokens using the DEFAULT_ADMIN_ROLE of their external AccessControl contracts.
  /// @param tokens The tokens to register the admin for.
  /// @dev The caller must have the DEFAULT_ADMIN_ROLE on the contract returned by `ACCESS_CONTROLLED_REGISTRY`
  /// for every token.
  function registerExternalAccessControlDefaultAdmin(
    address[] calldata tokens
  ) external {
    for (uint256 i = 0; i < tokens.length; ++i) {
      address token = tokens[i];
      AccessControl accessControl = AccessControl(IExternalAccessControlRegistry(token).ACCESS_CONTROLLED_REGISTRY());
      _registerAccessControlDefaultAdmin(token, accessControl);
    }
  }

  /// @notice Registers the admin of the token using the DEFAULT_ADMIN_ROLE of the given AccessControl contract.
  /// @param token The token to register the admin for.
  /// @param accessControl The contract that holds the DEFAULT_ADMIN_ROLE for the token.
  function _registerAccessControlDefaultAdmin(
    address token,
    AccessControl accessControl
  ) private {
    bytes32 defaultAdminRole = accessControl.DEFAULT_ADMIN_ROLE();
    if (!accessControl.hasRole(defaultAdminRole, msg.sender)) {
      revert RequiredRoleNotFound(msg.sender, defaultAdminRole, token);
    }

    _registerAdmin(token, msg.sender);
  }

  /// @notice Registers the caller as the pending admin of the token.
  /// @param token The token to register the admin for.
  /// @param admin The admin to register, which must be the caller.
  function _registerAdmin(
    address token,
    address admin
  ) private {
    if (admin != msg.sender) {
      revert CanOnlySelfRegister(admin, token);
    }

    i_tokenAdminRegistry.proposeAdministrator(token, admin);

    emit AdministratorRegistered(token, admin);
  }
}
