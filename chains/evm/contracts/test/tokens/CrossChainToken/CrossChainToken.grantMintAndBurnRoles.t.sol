// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CrossChainTokenSetup} from "./CrossChainTokenSetup.t.sol";

contract CrossChainToken_grantMintAndBurnRoles is CrossChainTokenSetup {
  function setUp() public virtual override {
    super.setUp();

    s_crossChainToken.grantRole(s_crossChainToken.BURN_MINT_ADMIN_ROLE(), OWNER);
  }

  function test_grantMintAndBurnRoles() public {
    address burnAndMinter = makeAddr("burnAndMinter");

    s_crossChainToken.grantMintAndBurnRoles(burnAndMinter);

    assertTrue(s_crossChainToken.hasRole(s_crossChainToken.MINTER_ROLE(), burnAndMinter));
    assertTrue(s_crossChainToken.hasRole(s_crossChainToken.BURNER_ROLE(), burnAndMinter));
  }
}
