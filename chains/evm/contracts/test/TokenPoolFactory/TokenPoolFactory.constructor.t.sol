// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenAdminRegistry} from "../../interfaces/ITokenAdminRegistry.sol";

import {TokenPoolFactory} from "../../TokenPoolFactory.sol";
import {RegistryModuleOwnerCustom} from "../../tokenAdminRegistry/RegistryModuleOwnerCustom.sol";
import {TokenPoolFactorySetup} from "./TokenPoolFactorySetup.t.sol";

contract TokenPoolFactory_constructor is TokenPoolFactorySetup {
  function test_constructor_getStaticConfig() public view {
    (address rmnProxy, address tokenAdminRegistry, address registryModuleOwnerCustom, address ccipRouter) =
      s_tokenPoolFactory.getStaticConfig();

    assertEq(rmnProxy, s_rmnProxy);
    assertEq(tokenAdminRegistry, address(s_tokenAdminRegistry));
    assertEq(registryModuleOwnerCustom, address(s_registryModuleOwnerCustom));
    assertEq(ccipRouter, address(s_sourceRouter));
  }

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
