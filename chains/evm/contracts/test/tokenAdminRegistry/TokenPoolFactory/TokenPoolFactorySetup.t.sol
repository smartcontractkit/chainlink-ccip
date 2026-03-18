// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BurnMintTokenPool} from "../../../pools/BurnMintTokenPool.sol";
import {RegistryModuleOwnerCustom} from "../../../tokenAdminRegistry/RegistryModuleOwnerCustom.sol";
import {FactoryBurnMintERC20} from "../../../tokenAdminRegistry/TokenPoolFactory/FactoryBurnMintERC20.sol";
import {TokenPoolFactory} from "../../../tokenAdminRegistry/TokenPoolFactory/TokenPoolFactory.sol";
import {TokenAdminRegistrySetup} from "../TokenAdminRegistry/TokenAdminRegistrySetup.t.sol";
import {Create2} from "@openzeppelin/contracts@5.3.0/utils/Create2.sol";

contract TokenPoolFactorySetup is TokenAdminRegistrySetup {
  using Create2 for bytes32;

  TokenPoolFactory internal s_tokenPoolFactory;
  RegistryModuleOwnerCustom internal s_registryModuleOwnerCustom;

  bytes32 internal constant FAKE_SALT = keccak256(abi.encode("FAKE_SALT"));
  uint256 public constant PREMINT_AMOUNT = 100 ether;

  address internal s_rmnProxy = address(0x1234);

  bytes internal constant TOKEN_INIT_CODE = abi.encodePacked(
    type(FactoryBurnMintERC20).creationCode, abi.encode("TestToken", "TT", 18, type(uint256).max, PREMINT_AMOUNT, OWNER)
  );
  bytes internal constant POOL_INIT_CODE = type(BurnMintTokenPool).creationCode;

  function setUp() public virtual override {
    TokenAdminRegistrySetup.setUp();

    s_registryModuleOwnerCustom = new RegistryModuleOwnerCustom(address(s_tokenAdminRegistry));
    s_tokenAdminRegistry.addRegistryModule(address(s_registryModuleOwnerCustom));

    s_tokenPoolFactory =
      new TokenPoolFactory(s_tokenAdminRegistry, s_registryModuleOwnerCustom, s_rmnProxy, address(s_sourceRouter));
  }
}
