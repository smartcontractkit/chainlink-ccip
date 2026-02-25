// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SiloedLockReleaseTokenPool} from "../../../pools/SiloedLockReleaseTokenPool.sol";
import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../../tmp/CrossChainToken.sol";
import {BaseTest} from "../../BaseTest.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract SiloedLockReleaseTokenPool_constructor is BaseTest {
  function test_constructor() public {
    CrossChainToken token = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "TKN", symbol: "T", decimals: DEFAULT_TOKEN_DECIMALS, maxSupply: 0, preMint: 0, ccipAdmin: OWNER
      }),
      OWNER,
      OWNER
    );

    SiloedLockReleaseTokenPool pool = new SiloedLockReleaseTokenPool(
      IERC20(address(token)), DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), address(s_sourceRouter)
    );

    assertEq(address(pool.getToken()), address(token));
    assertEq(pool.typeAndVersion(), "SiloedLockReleaseTokenPool 2.0.0-dev");
  }
}
