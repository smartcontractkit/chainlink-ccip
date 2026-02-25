// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../../tmp/CrossChainToken.sol";
import {RegistryModuleOwnerCustom} from "../../../tokenAdminRegistry/RegistryModuleOwnerCustom.sol";
import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";

import {Test} from "forge-std/Test.sol";

contract RegistryModuleOwnerCustomSetup is Test {
  address internal constant OWNER = 0x00007e64E1fB0C487F25dd6D3601ff6aF8d32e4e;

  RegistryModuleOwnerCustom internal s_registryModuleOwnerCustom;
  TokenAdminRegistry internal s_tokenAdminRegistry;
  address internal s_token;

  function setUp() public virtual {
    vm.startPrank(OWNER);

    s_tokenAdminRegistry = new TokenAdminRegistry();
    s_token = address(
      new CrossChainToken(
        BaseERC20.ConstructorParams({
          name: "Test", symbol: "TST", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER
        }),
        OWNER,
        OWNER
      )
    );
    s_registryModuleOwnerCustom = new RegistryModuleOwnerCustom(address(s_tokenAdminRegistry));
    s_tokenAdminRegistry.addRegistryModule(address(s_registryModuleOwnerCustom));
  }
}
