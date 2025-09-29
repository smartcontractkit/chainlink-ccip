// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {RegistryModuleOwnerCustom} from "../../../tokenAdminRegistry/RegistryModuleOwnerCustom.sol";
import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {Test} from "forge-std/Test.sol";

contract RegistryModuleOwnerCustomSetup is Test {
  address internal constant OWNER = 0x00007e64E1fB0C487F25dd6D3601ff6aF8d32e4e;

  RegistryModuleOwnerCustom internal s_registryModuleOwnerCustom;
  TokenAdminRegistry internal s_tokenAdminRegistry;
  address internal s_token;

  function setUp() public virtual {
    vm.startPrank(OWNER);

    s_tokenAdminRegistry = new TokenAdminRegistry();
    s_token = address(new BurnMintERC20("Test", "TST", 18, 0, 0));
    s_registryModuleOwnerCustom = new RegistryModuleOwnerCustom(address(s_tokenAdminRegistry));
    s_tokenAdminRegistry.addRegistryModule(address(s_registryModuleOwnerCustom));
  }
}
