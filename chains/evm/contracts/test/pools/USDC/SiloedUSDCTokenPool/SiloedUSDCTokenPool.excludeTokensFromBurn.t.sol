// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SiloedLockReleaseTokenPool} from "../../../../pools/SiloedLockReleaseTokenPool.sol";
import {SiloedUSDCTokenPool} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

contract SiloedUSDCTokenPool_excludeTokensFromBurn is SiloedUSDCTokenPoolSetup {
  function test_excludeTokensFromBurn() public {
    // Provide some liquidity to the lockbox
    uint256 amount = 1000e6;
    deal(address(s_USDCToken), address(s_destLockBox), amount);

    // Propose migration
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    // Exclude tokens
    uint256 excludeAmount = 500e6;
    vm.expectEmit();
    emit SiloedUSDCTokenPool.TokensExcludedFromBurn(DEST_CHAIN_SELECTOR, excludeAmount, amount - excludeAmount);
    s_usdcTokenPool.excludeTokensFromBurn(DEST_CHAIN_SELECTOR, excludeAmount);

    // Verify tokens were excluded
    assertEq(s_usdcTokenPool.getExcludedTokensByChain(DEST_CHAIN_SELECTOR), excludeAmount);
  }

  // Reverts

  function test_excludeTokensFromBurn_RevertWhen_NoMigrationProposalPending() public {
    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.NoMigrationProposalPending.selector));
    s_usdcTokenPool.excludeTokensFromBurn(SOURCE_CHAIN_SELECTOR, 1e6);
  }
}
