// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CrossChainPoolToken} from "../../../pools/CrossChainPoolToken.sol";
import {BaseERC20} from "../../../tokens/BaseERC20.sol";
import {TokenPoolSetup} from "../TokenPool/TokenPoolSetup.t.sol";

contract CrossChainPoolTokenSetup is TokenPoolSetup {
  CrossChainPoolToken internal s_cctPool;

  uint256 internal constant MAX_SUPPLY = 1_000_000e18;
  uint256 internal constant PRE_MINT = 100_000e18;

  function setUp() public virtual override {
    super.setUp();

    s_cctPool = new CrossChainPoolToken(
      BaseERC20.ConstructorParams({
        name: "CCT Token",
        symbol: "CCT",
        decimals: DEFAULT_TOKEN_DECIMALS,
        maxSupply: MAX_SUPPLY,
        preMint: PRE_MINT,
        ccipAdmin: OWNER
      }),
      address(0),
      address(s_mockRMNRemote),
      address(s_sourceRouter)
    );

    _applyChainUpdates(address(s_cctPool));
  }
}
