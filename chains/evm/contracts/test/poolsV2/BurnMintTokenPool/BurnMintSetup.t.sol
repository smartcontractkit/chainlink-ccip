// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BurnMintTokenPool} from "../../../poolsV2/BurnMintTokenPool.sol";
import {TokenPoolV2Setup} from "../TokenPool/TokenPoolV2Setup.t.sol";

contract BurnMintSetup is TokenPoolV2Setup {
  BurnMintTokenPool internal s_pool;

  function setUp() public virtual override {
    super.setUp();

    s_pool = new BurnMintTokenPool(
      s_token, DEFAULT_TOKEN_DECIMALS, new address[](0), address(s_mockRMNRemote), address(s_sourceRouter)
    );
    s_token.grantMintAndBurnRoles(address(s_pool));

    _applyChainUpdates(address(s_pool));
  }
}
