// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {AccessControl} from "@openzeppelin/contracts@5.3.0/access/AccessControl.sol";

/// @notice AccessControl mock with a default admin set at deployment.
contract MockExternalAccessControlRegistry is AccessControl {
  /// @param admin The address to grant the DEFAULT_ADMIN_ROLE.
  constructor(
    address admin
  ) {
    _grantRole(DEFAULT_ADMIN_ROLE, admin);
  }
}
