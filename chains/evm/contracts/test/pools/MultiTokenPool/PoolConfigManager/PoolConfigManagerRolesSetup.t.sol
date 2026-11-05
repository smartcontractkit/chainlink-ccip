// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@5.3.0/access/IAccessControl.sol";

/// @notice Checks each allowed and disallowed role independently. Tests supply call data and the allowed roles. The
/// last allowed role is the one reported by _checkRoleOr on failure.
abstract contract PoolConfigManagerRolesSetup is PoolConfigManagerSetup {
  /// @dev Grants one role to a caller whose address is derived from that role. Must be called by the default admin.
  function _grantNewCaller(
    bytes32 role
  ) internal returns (address caller) {
    caller = makeAddr(vm.toString(role));
    s_poolConfigManager.grantRole(role, caller);
    return caller;
  }

  function _getMissingRoleError(
    address caller,
    bytes32 role
  ) internal pure returns (bytes memory) {
    return abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, caller, role);
  }

  /// @dev CONFIG_ROLE, POLICY_ROLE, and the role of the type.
  function _getPolicyWriteRoles(
    bytes32 typeRole
  ) internal view returns (bytes32[] memory roles) {
    roles = new bytes32[](3);
    roles[0] = s_poolConfigManager.CONFIG_ROLE();
    roles[1] = s_poolConfigManager.POLICY_ROLE();
    roles[2] = typeRole;
    return roles;
  }

  /// @dev CONFIG_ROLE, TOKEN_MANAGER_ROLE, and the role of the type.
  function _getTokenWriteRoles(
    bytes32 typeRole
  ) internal view returns (bytes32[] memory roles) {
    roles = new bytes32[](3);
    roles[0] = s_poolConfigManager.CONFIG_ROLE();
    roles[1] = s_poolConfigManager.TOKEN_MANAGER_ROLE();
    roles[2] = typeRole;
    return roles;
  }

  function _getProfileWriteRoles() internal view returns (bytes32[] memory) {
    return _rolesOf(s_poolConfigManager.CONFIG_ROLE(), s_poolConfigManager.TOKEN_MANAGER_ROLE());
  }

  function _getPauseRoles() internal view returns (bytes32[] memory) {
    return _rolesOf(s_poolConfigManager.DEFAULT_ADMIN_ROLE(), s_poolConfigManager.PAUSER_ROLE());
  }

  function _getUnpauseRoles() internal view returns (bytes32[] memory) {
    return _rolesOf(s_poolConfigManager.DEFAULT_ADMIN_ROLE(), s_poolConfigManager.UNPAUSER_ROLE());
  }

  /// @dev Checks that each allowed role authorizes the call on its own. Revokes CONFIG_ROLE from the owner to
  /// test DEFAULT_ADMIN_ROLE alone.
  function _assertRolesPass(
    bytes memory callData,
    bytes32[] memory allowedRoles
  ) internal {
    s_poolConfigManager.revokeRole(s_poolConfigManager.CONFIG_ROLE(), OWNER);
    for (uint256 i = 0; i < allowedRoles.length; ++i) {
      address caller = _getCallerOf(allowedRoles[i]);
      _setCaller(caller);
      (bool success, bytes memory returnData) = address(s_poolConfigManager).call(callData);
      assertTrue(success, vm.toString(returnData));
      _setCaller(OWNER);
    }
  }

  /// @dev Calls with a disallowed role or no role must revert. The error must identify the last allowed role.
  /// Revokes CONFIG_ROLE from the owner to test DEFAULT_ADMIN_ROLE alone.
  function _assertRolesRevert(
    bytes memory callData,
    bytes32[] memory allowedRoles
  ) internal {
    bytes32 expectedRole = allowedRoles[allowedRoles.length - 1];
    s_poolConfigManager.revokeRole(s_poolConfigManager.CONFIG_ROLE(), OWNER);
    bytes32[] memory roles = _getAllRoles();
    for (uint256 i = 0; i < roles.length; ++i) {
      if (_contains(allowedRoles, roles[i])) continue;
      _assertMissingRole(_getCallerOf(roles[i]), callData, expectedRole);
    }
    _assertMissingRole(STRANGER, callData, expectedRole);
  }

  function _assertMissingRole(
    address caller,
    bytes memory callData,
    bytes32 expectedRole
  ) private {
    _setCaller(caller);
    (bool success, bytes memory returnData) = address(s_poolConfigManager).call(callData);
    assertFalse(success);
    assertEq(_getMissingRoleError(caller, expectedRole), returnData);
    _setCaller(OWNER);
  }

  /// @dev Uses the owner for DEFAULT_ADMIN_ROLE because grantRole cannot assign it. Creates a separate caller for
  /// each other role.
  function _getCallerOf(
    bytes32 role
  ) private returns (address) {
    if (role == s_poolConfigManager.DEFAULT_ADMIN_ROLE()) return OWNER;
    return _grantNewCaller(role);
  }

  function _getAllRoles() private view returns (bytes32[] memory roles) {
    roles = new bytes32[](13);
    roles[0] = s_poolConfigManager.DEFAULT_ADMIN_ROLE();
    roles[1] = s_poolConfigManager.CONFIG_ROLE();
    roles[2] = s_poolConfigManager.POLICY_ROLE();
    roles[3] = s_poolConfigManager.CONNECTION_ROLE();
    roles[4] = s_poolConfigManager.REMOTE_POOL_ROLE();
    roles[5] = s_poolConfigManager.RATE_LIMIT_ROLE();
    roles[6] = s_poolConfigManager.FEE_ROLE();
    roles[7] = s_poolConfigManager.DEST_OVERHEAD_ROLE();
    roles[8] = s_poolConfigManager.SECURITY_ROLE();
    roles[9] = s_poolConfigManager.TOKEN_MANAGER_ROLE();
    roles[10] = s_poolConfigManager.PAUSER_ROLE();
    roles[11] = s_poolConfigManager.UNPAUSER_ROLE();
    roles[12] = s_poolConfigManager.FEE_COLLECTOR_ROLE();
    return roles;
  }

  function _rolesOf(
    bytes32 firstRole,
    bytes32 secondRole
  ) private pure returns (bytes32[] memory roles) {
    roles = new bytes32[](2);
    roles[0] = firstRole;
    roles[1] = secondRole;
    return roles;
  }

  function _contains(
    bytes32[] memory roles,
    bytes32 role
  ) private pure returns (bool) {
    for (uint256 i = 0; i < roles.length; ++i) {
      if (roles[i] == role) return true;
    }
    return false;
  }
}
