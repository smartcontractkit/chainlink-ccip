// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../../tmp/CrossChainToken.sol";
import {BaseTest} from "../../BaseTest.t.sol";

contract CrossChainTokenSetup is BaseTest {
  CrossChainToken internal s_crossChainToken;

  uint256 internal constant MAX_SUPPLY = 1_000_000e18;
  uint256 internal constant PRE_MINT = 100_000e18;

  function setUp() public virtual override {
    super.setUp();

    s_crossChainToken = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "CrossChain Token",
        symbol: "CCT",
        decimals: DEFAULT_TOKEN_DECIMALS,
        maxSupply: MAX_SUPPLY,
        preMint: PRE_MINT,
        ccipAdmin: OWNER
      }),
      address(0),
      OWNER
    );
  }
}
