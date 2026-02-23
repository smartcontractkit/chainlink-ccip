// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTTokenPoolSetup} from "./CCTTokenPoolSetup.t.sol";

import {IGetCCIPAdmin} from "../../../interfaces/IGetCCIPAdmin.sol";
import {IPoolV1} from "../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract CCTTokenPool_constructor is CCTTokenPoolSetup {
  function test_constructor() public view {
    assertEq("CCT Token", s_cctPool.name());
    assertEq("CCT", s_cctPool.symbol());
    assertEq(DEFAULT_TOKEN_DECIMALS, s_cctPool.decimals());
    assertEq(MAX_SUPPLY, s_cctPool.maxSupply());
    assertEq(PRE_MINT, s_cctPool.totalSupply());
    assertEq(PRE_MINT, IERC20(address(s_cctPool)).balanceOf(OWNER));
    assertEq(OWNER, s_cctPool.getCCIPAdmin());
  }

  function test_constructor_PoolTokenIsSelf() public view {
    assertEq(address(s_cctPool), address(s_cctPool.getToken()));
  }

  function test_constructor_SupportsInterface() public view {
    assertTrue(s_cctPool.supportsInterface(type(IERC20).interfaceId));
    assertTrue(s_cctPool.supportsInterface(type(IGetCCIPAdmin).interfaceId));
    assertTrue(s_cctPool.supportsInterface(type(IPoolV1).interfaceId));
    assertTrue(s_cctPool.supportsInterface(type(IPoolV2).interfaceId));
  }

  function test_typeAndVersion() public view {
    assertEq("CCTTokenPool 2.0.0-dev", s_cctPool.typeAndVersion());
  }
}
