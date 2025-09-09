pragma solidity ^0.8.24;

import {BurnMintWithLockReleaseFlagTokenPool} from "../../../pools/USDC/BurnMintWithLockReleaseFlagTokenPool.sol";
import {BurnMintSetup} from "../BurnMintTokenPool/BurnMintSetup.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract BurnMintWithLockReleaseFlagTokenPoolSetup is BurnMintSetup {
  BurnMintWithLockReleaseFlagTokenPool internal s_pool;

  function setUp() public virtual override {
    super.setUp();

    // To simulate USDC we need to override the decimals to 6
    s_token = new BurnMintERC20("Chainlink Token", "LINK", 6, 0, 0);

    s_pool = new BurnMintWithLockReleaseFlagTokenPool(
      s_token, 6, new address[](0), address(s_mockRMNRemote), address(s_sourceRouter)
    );
    s_token.grantMintAndBurnRoles(address(s_pool));

    _applyChainUpdates(address(s_pool));
  }
}
