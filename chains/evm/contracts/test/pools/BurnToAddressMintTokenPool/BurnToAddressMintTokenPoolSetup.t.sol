// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BurnToAddressMintTokenPool} from "../../../pools/BurnToAddressMintTokenPool.sol";
import {TokenPoolSetup} from "../TokenPool/TokenPoolSetup.t.sol";

contract BurnToAddressMintTokenPoolSetup is TokenPoolSetup {
  BurnToAddressMintTokenPool internal s_pool;

  address public constant BURN_ADDRESS = address(0xdead);

  function setUp() public virtual override {
    super.setUp();

    s_pool = new BurnToAddressMintTokenPool(
      s_token, DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), address(s_sourceRouter), BURN_ADDRESS
    );

    s_token.grantMintAndBurnRoles(address(s_pool));

    _applyChainUpdates(address(s_pool));
  }
}
