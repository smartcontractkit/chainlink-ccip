// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IGetCCIPAdmin} from "../interfaces/IGetCCIPAdmin.sol";
import {IOwner} from "../interfaces/IOwner.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";
import {IExternalAccessControlRegistry} from "../interfaces/externalRegistry/IExternalAccessControlRegistry.sol";

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {AccessControl} from "@openzeppelin/contracts@5.3.0/access/AccessControl.sol";

/// @notice Registers token admins in batches using the token's getCCIPAdmin, owner, DEFAULT_ADMIN_ROLE, or
/// CCIP_TOKEN_REGISTRY_ADMIN_ROLE.
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
  /// @notice The holder of this role can register itself as the pending admin of the token.
  /// @dev In OZ AccessControl, only DEFAULT_ADMIN_ROLE can grant this role, unless the token sets another admin role
  /// for it. When many tokens share one AccessControl contract, one grant covers all of them. Revoking the role does
  /// not cancel a pending registration.
  bytes32 public constant CCIP_TOKEN_REGISTRY_ADMIN_ROLE = keccak256("CCIP_TOKEN_REGISTRY_ADMIN_ROLE");

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
      AccessControl accessControl = AccessControl(token);
      _registerAccessControlRole(token, accessControl, accessControl.DEFAULT_ADMIN_ROLE());
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
      _registerAccessControlRole(token, accessControl, accessControl.DEFAULT_ADMIN_ROLE());
    }
  }

  /// @notice Registers the admin of the tokens using the CCIP_TOKEN_REGISTRY_ADMIN_ROLE in each token's AccessControl.
  /// @param tokens The tokens to register the admin for.
  /// @dev The caller must have the CCIP_TOKEN_REGISTRY_ADMIN_ROLE in each token's AccessControl.
  function registerAccessControlCCIPTokenRegistryAdmin(
    address[] calldata tokens
  ) external {
    for (uint256 i = 0; i < tokens.length; ++i) {
      address token = tokens[i];
      _registerAccessControlRole(token, AccessControl(token), CCIP_TOKEN_REGISTRY_ADMIN_ROLE);
    }
  }

  /// @notice Registers the admin of the tokens using the CCIP_TOKEN_REGISTRY_ADMIN_ROLE of their external
  /// AccessControl contracts.
  /// @param tokens The tokens to register the admin for.
  /// @dev The caller must have the CCIP_TOKEN_REGISTRY_ADMIN_ROLE on the contract returned by
  /// `ACCESS_CONTROLLED_REGISTRY` for every token.
  function registerExternalAccessControlCCIPTokenRegistryAdmin(
    address[] calldata tokens
  ) external {
    for (uint256 i = 0; i < tokens.length; ++i) {
      address token = tokens[i];
      AccessControl accessControl = AccessControl(IExternalAccessControlRegistry(token).ACCESS_CONTROLLED_REGISTRY());
      _registerAccessControlRole(token, accessControl, CCIP_TOKEN_REGISTRY_ADMIN_ROLE);
    }
  }

  /// @notice Registers the admin of the token using a role of the given AccessControl contract.
  /// @param token The token to register the admin for.
  /// @param accessControl The contract that holds the roles for the token.
  /// @param role The role that the caller must have.
  function _registerAccessControlRole(
    address token,
    AccessControl accessControl,
    bytes32 role
  ) private {
    if (!accessControl.hasRole(role, msg.sender)) {
      revert RequiredRoleNotFound(msg.sender, role, token);
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
