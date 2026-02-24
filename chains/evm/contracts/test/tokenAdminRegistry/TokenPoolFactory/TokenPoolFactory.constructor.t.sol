// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenAdminRegistry} from "../../../interfaces/ITokenAdminRegistry.sol";

import {RegistryModuleOwnerCustom} from "../../../tokenAdminRegistry/RegistryModuleOwnerCustom.sol";
import {TokenPoolFactory} from "../../../tokenAdminRegistry/TokenPoolFactory/TokenPoolFactory.sol";
import {TokenPoolFactorySetup} from "./TokenPoolFactorySetup.t.sol";

contract TokenPoolFactory_constructor is TokenPoolFactorySetup {
  function test_constructor_RevertWhen_InvalidZeroAddress() public {
    // Revert cause the tokenAdminRegistry is address(0).
    vm.expectRevert(TokenPoolFactory.InvalidZeroAddress.selector);
    new TokenPoolFactory(ITokenAdminRegistry(address(0)), RegistryModuleOwnerCustom(address(0)), address(0), address(0));

    new TokenPoolFactory(
      ITokenAdminRegistry(makeAddr("TOKEN_ADMIN_REGISTRY")),
      RegistryModuleOwnerCustom(makeAddr("REGISTRY_MODULE_OWNER_CUSTOM")),
      makeAddr("RMN_PROXY"),
      makeAddr("ROUTER")
    );
  }
}
