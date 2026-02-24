// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BurnMintTokenPool} from "../../../pools/BurnMintTokenPool.sol";
import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../../tmp/CrossChainToken.sol";
import {RegistryModuleOwnerCustom} from "../../../tokenAdminRegistry/RegistryModuleOwnerCustom.sol";
import {TokenPoolFactory} from "../../../tokenAdminRegistry/TokenPoolFactory/TokenPoolFactory.sol";
import {TokenAdminRegistrySetup} from "../TokenAdminRegistry/TokenAdminRegistrySetup.t.sol";

contract TokenPoolFactorySetup is TokenAdminRegistrySetup {
  TokenPoolFactory internal s_tokenPoolFactory;
  RegistryModuleOwnerCustom internal s_registryModuleOwnerCustom;
  address internal s_rmnProxy = address(0x1234);

  bytes internal constant POOL_INIT_CODE = type(BurnMintTokenPool).creationCode;
  uint256 public constant PREMINT_AMOUNT = 100 ether;
  bytes32 internal constant FAKE_SALT = keccak256(abi.encode("FAKE_SALT"));

  bytes internal s_tokenInitCode;
  bytes internal s_poolInitArgs;

  function setUp() public virtual override {
    super.setUp();

    s_registryModuleOwnerCustom = new RegistryModuleOwnerCustom(address(s_tokenAdminRegistry));
    s_tokenAdminRegistry.addRegistryModule(address(s_registryModuleOwnerCustom));

    s_tokenPoolFactory =
      new TokenPoolFactory(s_tokenAdminRegistry, s_registryModuleOwnerCustom, s_rmnProxy, address(s_sourceRouter));

    s_tokenInitCode = _buildTokenInitCode(address(s_tokenPoolFactory));

    s_poolInitArgs = abi.encode(address(0), address(0x1234), s_sourceRouter);
  }

  function _buildTokenInitCode(
    address factory
  ) internal pure returns (bytes memory) {
    bytes memory tokenCreationParams = abi.encode(
      BaseERC20.ConstructorParams({
        name: "TestToken",
        symbol: "TT",
        decimals: 18,
        maxSupply: type(uint256).max,
        preMint: PREMINT_AMOUNT,
        ccipAdmin: factory
      }),
      factory,
      OWNER
    );
    return abi.encodePacked(type(CrossChainToken).creationCode, tokenCreationParams);
  }
}
