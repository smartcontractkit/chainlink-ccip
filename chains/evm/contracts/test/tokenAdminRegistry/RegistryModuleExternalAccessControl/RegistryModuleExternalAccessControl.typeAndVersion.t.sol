// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RegistryModuleExternalAccessControlSetup} from "./RegistryModuleExternalAccessControlSetup.t.sol";

contract RegistryModuleExternalAccessControl_typeAndVersion is RegistryModuleExternalAccessControlSetup {
  function test_typeAndVersion() public view {
    assertEq("RegistryModuleExternalAccessControl 2.1.0-dev", s_registryModuleExternalAccessControl.typeAndVersion());
  }
}
