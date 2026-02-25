// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "../../../interfaces/IBurnMintERC20.sol";
import {BurnMintWithLockReleaseFlagTokenPool} from "../../../pools/USDC/BurnMintWithLockReleaseFlagTokenPool.sol";
import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../../tmp/CrossChainToken.sol";
import {TokenPoolSetup} from "../TokenPool/TokenPoolSetup.t.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract BurnMintWithLockReleaseFlagTokenPoolSetup is TokenPoolSetup {
  BurnMintWithLockReleaseFlagTokenPool internal s_pool;

  function setUp() public virtual override {
    super.setUp();

    // To simulate USDC we need to override the decimals to 6
    s_token = IERC20(
      address(
        new CrossChainToken(
          BaseERC20.ConstructorParams({
            name: "Chainlink Token", symbol: "LINK", decimals: 6, maxSupply: 0, preMint: 0, ccipAdmin: OWNER
          }),
          OWNER,
          OWNER
        )
      )
    );

    s_pool = new BurnMintWithLockReleaseFlagTokenPool(
      IBurnMintERC20(address(s_token)), 6, address(0), address(s_mockRMNRemote), address(s_sourceRouter)
    );
    CrossChainToken(address(s_token)).grantMintAndBurnRoles(address(s_pool));

    _applyChainUpdates(address(s_pool));
  }
}
