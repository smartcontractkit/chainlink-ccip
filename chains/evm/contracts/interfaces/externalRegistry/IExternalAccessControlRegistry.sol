// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Token methods used by RegistryModuleExternalAccessControl to read the external AccessControl contract.
interface IExternalAccessControlRegistry {
  /// @notice Returns the AccessControl contract that holds the roles for this token.
  /// @return The address of the AccessControl contract.
  /// @dev This method is named to match the token's public getter.
  // solhint-disable-next-line func-name-mixedcase
  function ACCESS_CONTROLLED_REGISTRY() external view returns (address);
}
