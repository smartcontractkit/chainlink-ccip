// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SiloedLockReleaseTokenPool} from "../../../pools/SiloedLockReleaseTokenPool.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract SiloedLockReleaseTokenPool_constructor is BaseTest {
  function test_constructor() public {
    BurnMintERC20 token = new BurnMintERC20("TKN", "T", DEFAULT_TOKEN_DECIMALS, 0, 0);

    SiloedLockReleaseTokenPool pool = new SiloedLockReleaseTokenPool(
      IERC20(address(token)), DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), address(s_sourceRouter)
    );

    assertEq(address(pool.getToken()), address(token));
    assertEq(pool.typeAndVersion(), "SiloedLockReleaseTokenPool 1.7.0-dev");
  }
}
