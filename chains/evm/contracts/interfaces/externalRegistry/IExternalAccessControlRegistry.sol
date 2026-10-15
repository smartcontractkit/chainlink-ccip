// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Exposes a token's external access control registry to RegistryModuleExternalAccessControl.
interface IExternalAccessControlRegistry {
  /// @notice Returns the OZ AccessControl contract that holds the roles of this contract.
  /// @return The address of the external access control registry.
  /// @dev The getter name preserves compatibility with tokens exposing this ABI.
  // solhint-disable-next-line func-name-mixedcase
  function ACCESS_CONTROLLED_REGISTRY() external view returns (address);
}
