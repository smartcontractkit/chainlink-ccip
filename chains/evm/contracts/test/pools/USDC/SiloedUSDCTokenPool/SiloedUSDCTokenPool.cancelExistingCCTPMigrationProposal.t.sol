// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SiloedLockReleaseTokenPool} from "../../../../pools/SiloedLockReleaseTokenPool.sol";
import {SiloedUSDCTokenPool} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

contract SiloedUSDCTokenPool_cancelExistingCCTPMigrationProposal is SiloedUSDCTokenPoolSetup {
  function test_cancelExistingCCTPMigrationProposal() public {
    // Propose a migration first
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    // Cancel the migration proposal
    s_usdcTokenPool.cancelExistingCCTPMigrationProposal();

    // Verify the proposal was cancelled
    assertEq(s_usdcTokenPool.getCurrentProposedCCTPChainMigration(), 0);
  }

  // Reverts

  function test_cancelExistingCCTPMigrationProposal_RevertWhen_NoProposalPending() public {
    // Expect revert when trying to cancel without a proposal
    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.NoMigrationProposalPending.selector));

    s_usdcTokenPool.cancelExistingCCTPMigrationProposal();
  }

  function test_cancelExistingCCTPMigrationProposal_ResetsExcludedTokens() public {
    // Set up silo designation for the test chain
    uint64[] memory removes = new uint64[](0);
    SiloedLockReleaseTokenPool.SiloConfigUpdate[] memory adds = new SiloedLockReleaseTokenPool.SiloConfigUpdate[](1);
    adds[0] = SiloedLockReleaseTokenPool.SiloConfigUpdate({remoteChainSelector: DEST_CHAIN_SELECTOR, rebalancer: OWNER});
    s_usdcTokenPool.updateSiloDesignations(removes, adds);

    // Provide some amount of liquidity to DEST_CHAIN_SELECTOR and use that amount for exclusion
    uint256 amount = 100e6;
    s_USDCToken.approve(address(s_usdcTokenPool), amount);
    s_usdcTokenPool.provideSiloedLiquidity(DEST_CHAIN_SELECTOR, amount);

    // Propose migration and exclude some tokens
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    // Exclude some tokens
    s_usdcTokenPool.excludeTokensFromBurn(DEST_CHAIN_SELECTOR, amount);

    // Verify tokens were excluded
    assertEq(s_usdcTokenPool.getExcludedTokensByChain(DEST_CHAIN_SELECTOR), amount);

    // Cancel the migration proposal
    s_usdcTokenPool.cancelExistingCCTPMigrationProposal();

    // Verify excluded tokens were reset
    assertEq(s_usdcTokenPool.getExcludedTokensByChain(DEST_CHAIN_SELECTOR), 0);
  }

  function test_cancelExistingCCTPMigrationProposal_EmitsEvent() public {
    // Propose a migration
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    // Expect the CCTPMigrationCancelled event
    vm.expectEmit();
    emit SiloedUSDCTokenPool.CCTPMigrationCancelled(DEST_CHAIN_SELECTOR);
    s_usdcTokenPool.cancelExistingCCTPMigrationProposal();
  }
}
