// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "../../../interfaces/IBurnMintERC20.sol";
import {BurnFromMintTokenPool} from "../../../pools/BurnFromMintTokenPool.sol";
import {BurnMintSetup} from "../BurnMintTokenPool/BurnMintSetup.t.sol";

contract BurnFromMintTokenPoolSetup is BurnMintSetup {
  BurnFromMintTokenPool internal s_pool;

  function setUp() public virtual override {
    super.setUp();

    s_pool = new BurnFromMintTokenPool(
      IBurnMintERC20(address(s_token)),
      DEFAULT_TOKEN_DECIMALS,
      address(0),
      address(s_mockRMNRemote),
      address(s_sourceRouter)
    );
    s_token.grantMintAndBurnRoles(address(s_pool));

    _applyChainUpdates(address(s_pool));
  }
}
