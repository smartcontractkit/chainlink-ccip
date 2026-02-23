// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "../../../interfaces/IBurnMintERC20.sol";
import {IGetCCIPAdmin} from "../../../interfaces/IGetCCIPAdmin.sol";
import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../../tmp/CrossChainToken.sol";
import {CrossChainTokenSetup} from "./CrossChainTokenSetup.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@5.3.0/access/IAccessControl.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract CrossChainToken_constructor is CrossChainTokenSetup {
  function test_constructor() public view {
    assertEq("CrossChain Token", s_crossChainToken.name());
    assertEq("CCT", s_crossChainToken.symbol());
    assertEq(DEFAULT_TOKEN_DECIMALS, s_crossChainToken.decimals());
    assertEq(MAX_SUPPLY, s_crossChainToken.maxSupply());
    assertEq(PRE_MINT, s_crossChainToken.totalSupply());
    assertEq(PRE_MINT, s_crossChainToken.balanceOf(OWNER));
    assertEq(OWNER, s_crossChainToken.getCCIPAdmin());
  }

  function test_constructor_OwnerDefaultsToMsgSender() public {
    CrossChainToken token = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "Test", symbol: "T", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER
      }),
      address(0),
      address(0)
    );

    assertEq(OWNER, token.defaultAdmin());
  }

  function test_constructor_SupportsInterface() public view {
    assertTrue(s_crossChainToken.supportsInterface(type(IERC20).interfaceId));
    assertTrue(s_crossChainToken.supportsInterface(type(IGetCCIPAdmin).interfaceId));
    assertTrue(s_crossChainToken.supportsInterface(type(IBurnMintERC20).interfaceId));
    assertTrue(s_crossChainToken.supportsInterface(type(IAccessControl).interfaceId));
  }

  function test_typeAndVersion() public view {
    assertEq("CrossChainToken 2.0.0-dev", s_crossChainToken.typeAndVersion());
  }

  function test_constructor_Roles() public view {
    bytes32 minterRole = s_crossChainToken.MINTER_ROLE();
    bytes32 burnerRole = s_crossChainToken.BURNER_ROLE();
    bytes32 burnMintAdminRole = s_crossChainToken.BURN_MINT_ADMIN_ROLE();

    assertEq(burnMintAdminRole, s_crossChainToken.getRoleAdmin(minterRole));
    assertEq(burnMintAdminRole, s_crossChainToken.getRoleAdmin(burnerRole));
  }
}
