// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BurnFromMintTokenPool} from "../../../pools/BurnFromMintTokenPool.sol";
import {TokenPoolSetup} from "../TokenPool/TokenPoolSetup.t.sol";

contract BurnFromMintTokenPoolSetup is TokenPoolSetup {
  BurnFromMintTokenPool internal s_pool;

  function setUp() public virtual override {
    super.setUp();

    s_pool = new BurnFromMintTokenPool(
      s_token, DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), address(s_sourceRouter)
    );
    s_token.grantMintAndBurnRoles(address(s_pool));

    _applyChainUpdates(address(s_pool));
  }
}
