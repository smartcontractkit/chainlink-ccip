// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IGetCCIPAdmin} from "../interfaces/IGetCCIPAdmin.sol";
import {IOwner} from "../interfaces/IOwner.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";
import {IExternalAccessControlRegistry} from "../interfaces/externalRegistry/IExternalAccessControlRegistry.sol";

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {AccessControl} from "@openzeppelin/contracts@5.3.0/access/AccessControl.sol";

/// @notice Proposes token administrators in batches using token ownership, CCIP administration, or default admin roles
/// held on the tokens themselves or in external access control registries.
/// @dev The module must be enabled in TokenAdminRegistry. Each token is checked independently and any failure reverts
/// the entire batch. Empty batches do nothing; duplicate tokens are processed in order. The caller is only proposed as
/// pending administrator and must separately call TokenAdminRegistry.acceptAdminRole for each token. An accepted
/// administrator cannot be replaced through this module; a pending administrator can be replaced by a current authority.
contract RegistryModuleExternalAccessControl is ITypeAndVersion {
  error CanOnlySelfRegister(address admin, address token);
  error RequiredRoleNotFound(address msgSender, bytes32 role, address token);
  error AddressZero();

  event AdministratorRegistered(address indexed token, address indexed administrator);

  string public constant override typeAndVersion = "RegistryModuleExternalAccessControl 2.1.0-dev";

  ITokenAdminRegistry private immutable i_tokenAdminRegistry;

  /// @notice Creates a batch registration module.
  /// @param tokenAdminRegistry The registry in which to propose token administrators.
  constructor(
    address tokenAdminRegistry
  ) {
    if (tokenAdminRegistry == address(0)) {
      revert AddressZero();
    }
    i_tokenAdminRegistry = ITokenAdminRegistry(tokenAdminRegistry);
  }

  /// @notice Proposes the caller as administrator of each token using its getCCIPAdmin method.
  /// @param tokens The tokens to register.
  /// @dev The caller must equal getCCIPAdmin for every token. Any failure reverts the entire batch.
  function registerAdminViaGetCCIPAdmin(
    address[] calldata tokens
  ) external {
    for (uint256 i = 0; i < tokens.length; ++i) {
      address token = tokens[i];
      _registerAdmin(token, IGetCCIPAdmin(token).getCCIPAdmin());
    }
  }

  /// @notice Proposes the caller as administrator of each token using its owner method.
  /// @param tokens The tokens to register.
  /// @dev The caller must equal owner for every token. Any failure reverts the entire batch.
  function registerAdminViaOwner(
    address[] calldata tokens
  ) external {
    for (uint256 i = 0; i < tokens.length; ++i) {
      address token = tokens[i];
      _registerAdmin(token, IOwner(token).owner());
    }
  }

  /// @notice Proposes the caller as administrator of each token using its own DEFAULT_ADMIN_ROLE.
  /// @param tokens The tokens to register.
  /// @dev The caller must hold DEFAULT_ADMIN_ROLE on every token. Any failure reverts the entire batch.
  function registerAccessControlDefaultAdmin(
    address[] calldata tokens
  ) external {
    for (uint256 i = 0; i < tokens.length; ++i) {
      address token = tokens[i];
      _registerAccessControlDefaultAdmin(token, AccessControl(token));
    }
  }

  /// @notice Proposes the caller as administrator of each token using its external registry's DEFAULT_ADMIN_ROLE.
  /// @param tokens The tokens to register.
  /// @dev The caller must have the DEFAULT_ADMIN_ROLE on the contract returned by `ACCESS_CONTROLLED_REGISTRY` of
  /// every token. If the caller does not have the role for one of the tokens, the call reverts and no token is
  /// registered.
  function registerExternalAccessControlDefaultAdmin(
    address[] calldata tokens
  ) external {
    for (uint256 i = 0; i < tokens.length; ++i) {
      address token = tokens[i];
      AccessControl accessControl = AccessControl(IExternalAccessControlRegistry(token).ACCESS_CONTROLLED_REGISTRY());
      _registerAccessControlDefaultAdmin(token, accessControl);
    }
  }

  function _registerAccessControlDefaultAdmin(
    address token,
    AccessControl accessControl
  ) internal {
    bytes32 defaultAdminRole = accessControl.DEFAULT_ADMIN_ROLE();
    if (!accessControl.hasRole(defaultAdminRole, msg.sender)) {
      revert RequiredRoleNotFound(msg.sender, defaultAdminRole, token);
    }

    _registerAdmin(token, msg.sender);
  }

  function _registerAdmin(
    address token,
    address admin
  ) internal {
    if (admin != msg.sender) {
      revert CanOnlySelfRegister(admin, token);
    }

    i_tokenAdminRegistry.proposeAdministrator(token, admin);

    emit AdministratorRegistered(token, admin);
  }
}
