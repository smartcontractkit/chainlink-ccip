// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BurnMintTokenPoolV2} from "../../../poolsV2/BurnMintTokenPoolV2.sol";
import {TokenPoolV2Setup} from "../TokenPoolV2/TokenPoolV2Setup.t.sol";

contract BurnMintTokenPoolV2Setup is TokenPoolV2Setup {
  BurnMintTokenPoolV2 internal s_pool;

  function setUp() public virtual override {
    super.setUp();

    s_pool = new BurnMintTokenPoolV2(
      s_token, DEFAULT_TOKEN_DECIMALS, new address[](0), address(s_mockRMNRemote), address(s_sourceRouter)
    );
    s_token.grantMintAndBurnRoles(address(s_pool));

    _applyChainUpdates(address(s_pool));
  }
}
