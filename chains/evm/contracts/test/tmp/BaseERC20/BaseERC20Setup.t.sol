// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {BaseTest} from "../../BaseTest.t.sol";

contract BaseERC20Setup is BaseTest {
  BaseERC20 internal s_baseERC20;

  uint256 internal constant MAX_SUPPLY = 1_000_000e18;
  uint256 internal constant PRE_MINT = 100_000e18;

  function setUp() public virtual override {
    super.setUp();

    s_baseERC20 = new BaseERC20(
      BaseERC20.ConstructorParams({
        name: "Base Token",
        symbol: "BASE",
        decimals: DEFAULT_TOKEN_DECIMALS,
        maxSupply: MAX_SUPPLY,
        preMint: PRE_MINT,
        ccipAdmin: OWNER
      })
    );
  }
}
