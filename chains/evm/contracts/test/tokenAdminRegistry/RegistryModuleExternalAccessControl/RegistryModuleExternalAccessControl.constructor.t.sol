// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RegistryModuleExternalAccessControl} from "../../../tokenAdminRegistry/RegistryModuleExternalAccessControl.sol";
import {RegistryModuleExternalAccessControlSetup} from "./RegistryModuleExternalAccessControlSetup.t.sol";

contract RegistryModuleExternalAccessControl_constructor is RegistryModuleExternalAccessControlSetup {
  function test_constructor() public {
    RegistryModuleExternalAccessControl module = new RegistryModuleExternalAccessControl(address(s_tokenAdminRegistry));
    s_tokenAdminRegistry.addRegistryModule(address(module));
    module.registerExternalAccessControlDefaultAdmin(_arrayOf(s_token));
    assertEq(OWNER, s_tokenAdminRegistry.getTokenConfig(s_token).pendingAdministrator);
  }

  // Reverts

  function test_constructor_RevertWhen_AddressZero() public {
    vm.expectRevert(abi.encodeWithSelector(RegistryModuleExternalAccessControl.AddressZero.selector));

    new RegistryModuleExternalAccessControl(address(0));
  }
}
