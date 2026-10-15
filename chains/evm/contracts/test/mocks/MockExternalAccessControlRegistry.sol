// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {AccessControl} from "@openzeppelin/contracts@5.3.0/access/AccessControl.sol";

contract MockExternalAccessControlRegistry is AccessControl {
  constructor(
    address admin
  ) {
    _grantRole(DEFAULT_ADMIN_ROLE, admin);
  }
}
