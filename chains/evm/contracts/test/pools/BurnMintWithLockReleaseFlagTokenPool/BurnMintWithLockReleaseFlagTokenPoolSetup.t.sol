// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "../../../interfaces/IBurnMintERC20.sol";
import {BurnMintWithLockReleaseFlagTokenPool} from "../../../pools/USDC/BurnMintWithLockReleaseFlagTokenPool.sol";
import {BurnMintSetup} from "../BurnMintTokenPool/BurnMintSetup.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract BurnMintWithLockReleaseFlagTokenPoolSetup is BurnMintSetup {
  BurnMintWithLockReleaseFlagTokenPool internal s_pool;

  function setUp() public virtual override {
    super.setUp();

    // To simulate USDC we need to override the decimals to 6
    s_token = IERC20(address(new BurnMintERC20("Chainlink Token", "LINK", 6, 0, 0)));

    s_pool = new BurnMintWithLockReleaseFlagTokenPool(
      IBurnMintERC20(address(s_token)), 6, address(0), address(s_mockRMNRemote), address(s_sourceRouter)
    );
    BurnMintERC20(address(s_token)).grantMintAndBurnRoles(address(s_pool));

    _applyChainUpdates(address(s_pool));
  }
}
