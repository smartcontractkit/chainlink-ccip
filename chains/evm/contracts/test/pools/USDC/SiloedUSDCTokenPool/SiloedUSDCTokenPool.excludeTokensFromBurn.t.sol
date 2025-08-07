// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SiloedLockReleaseTokenPool} from "../../../../pools/SiloedLockReleaseTokenPool.sol";
import {SiloedUSDCTokenPool} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

contract SiloedUSDCTokenPool_excludeTokensFromBurn is SiloedUSDCTokenPoolSetup {
  function test_excludeTokensFromBurn() public {
    // Set up silo designation for the test chain
    uint64[] memory removes = new uint64[](0);
    SiloedLockReleaseTokenPool.SiloConfigUpdate[] memory adds = new SiloedLockReleaseTokenPool.SiloConfigUpdate[](1);
    adds[0] = SiloedLockReleaseTokenPool.SiloConfigUpdate({remoteChainSelector: DEST_CHAIN_SELECTOR, rebalancer: OWNER});
    s_usdcTokenPool.updateSiloDesignations(removes, adds);

    // Provide some liquidity
    uint256 amount = 1000e6;
    s_USDCToken.approve(address(s_usdcTokenPool), amount);
    s_usdcTokenPool.provideSiloedLiquidity(DEST_CHAIN_SELECTOR, amount);

    // Propose migration
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    // Exclude tokens
    uint256 excludeAmount = 500e6;
    s_usdcTokenPool.excludeTokensFromBurn(DEST_CHAIN_SELECTOR, excludeAmount);

    // Verify tokens were excluded
    assertEq(s_usdcTokenPool.getExcludedTokensByChain(DEST_CHAIN_SELECTOR), excludeAmount);
  }

  function test_excludeTokensFromBurn_EmitsEvent() public {
    // Set up silo designation for the test chain
    uint64[] memory removes = new uint64[](0);
    SiloedLockReleaseTokenPool.SiloConfigUpdate[] memory adds = new SiloedLockReleaseTokenPool.SiloConfigUpdate[](1);
    adds[0] = SiloedLockReleaseTokenPool.SiloConfigUpdate({remoteChainSelector: DEST_CHAIN_SELECTOR, rebalancer: OWNER});
    s_usdcTokenPool.updateSiloDesignations(removes, adds);

    // Provide some liquidity
    uint256 amount = 1000e6;
    s_USDCToken.approve(address(s_usdcTokenPool), amount);
    s_usdcTokenPool.provideSiloedLiquidity(DEST_CHAIN_SELECTOR, amount);

    // Propose migration
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    // Expect the TokensExcludedFromBurn event
    uint256 excludeAmount = 500e6;
    vm.expectEmit();
    emit SiloedUSDCTokenPool.TokensExcludedFromBurn(DEST_CHAIN_SELECTOR, excludeAmount, amount - excludeAmount);
    s_usdcTokenPool.excludeTokensFromBurn(DEST_CHAIN_SELECTOR, excludeAmount);
  }

  // Reverts

  function test_excludeTokensFromBurn_RevertWhen_NoMigrationProposalPending() public {
    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.NoMigrationProposalPending.selector));
    s_usdcTokenPool.excludeTokensFromBurn(SOURCE_CHAIN_SELECTOR, 1e6);
  }
}
