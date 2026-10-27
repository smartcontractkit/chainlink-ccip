// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenWithUIDFactory} from "../../../../interfaces/ITokenWithUIDFactory.sol";

import {
  FactoryTokenWithUIDPoolDataProvider
} from "../../../../pools/MultiTokenPool/FactoryTokenWithUIDPoolDataProvider.sol";
import {FactoryTokenWithUIDPoolDataProviderSetup} from "./FactoryTokenWithUIDPoolDataProviderSetup.t.sol";

contract FactoryTokenWithUIDPoolDataProvider_constructor is FactoryTokenWithUIDPoolDataProviderSetup {
  function test_constructor() public {
    FactoryTokenWithUIDPoolDataProvider provider = new FactoryTokenWithUIDPoolDataProvider(s_tokenFactory);

    assertEq("FactoryTokenWithUIDPoolDataProvider 2.1.0-dev", provider.typeAndVersion());
    assertEq(address(s_tokenFactory), address(provider.getTokenFactory()));
  }

  // Reverts

  function test_constructor_RevertWhen_InvalidTokenFactory() public {
    ITokenWithUIDFactory factoryWithoutCode = ITokenWithUIDFactory(makeAddr("factoryWithoutCode"));

    vm.expectRevert(
      abi.encodeWithSelector(FactoryTokenWithUIDPoolDataProvider.InvalidTokenFactory.selector, factoryWithoutCode)
    );
    new FactoryTokenWithUIDPoolDataProvider(factoryWithoutCode);
  }

  function test_constructor_RevertWhen_InvalidTokenFactory_ZeroAddress() public {
    vm.expectRevert(
      abi.encodeWithSelector(FactoryTokenWithUIDPoolDataProvider.InvalidTokenFactory.selector, address(0))
    );
    new FactoryTokenWithUIDPoolDataProvider(ITokenWithUIDFactory(address(0)));
  }
}
