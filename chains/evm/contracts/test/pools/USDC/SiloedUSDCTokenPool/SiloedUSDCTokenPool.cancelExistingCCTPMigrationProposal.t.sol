// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SiloedLockReleaseTokenPool} from "../../../../pools/SiloedLockReleaseTokenPool.sol";
import {SiloedUSDCTokenPool} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

contract SiloedUSDCTokenPool_cancelExistingCCTPMigrationProposal is SiloedUSDCTokenPoolSetup {
  function test_cancelExistingCCTPMigrationProposal_Success() public {
    // Arrange: Propose a migration first
    vm.startPrank(OWNER);
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    // Act: Cancel the migration proposal
    s_usdcTokenPool.cancelExistingCCTPMigrationProposal();
    vm.stopPrank();

    // Assert: Verify the proposal was cancelled
    assertEq(s_usdcTokenPool.getCurrentProposedCCTPChainMigration(), 0);
  }

  function test_cancelExistingCCTPMigrationProposal_RevertWhen_NoProposalPending() public {
    // Arrange: No migration proposal is set
    vm.startPrank(OWNER);

    // Act & Assert: Expect revert when trying to cancel without a proposal
    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.NoMigrationProposalPending.selector));
    s_usdcTokenPool.cancelExistingCCTPMigrationProposal();
    vm.stopPrank();
  }

  function test_cancelExistingCCTPMigrationProposal_RevertWhen_NotOwner() public {
    // Arrange: Set up a migration proposal
    vm.startPrank(OWNER);
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);
    vm.stopPrank();

    // Act & Assert: Expect revert when non-owner tries to cancel
    vm.startPrank(STRANGER);
    vm.expectRevert();
    s_usdcTokenPool.cancelExistingCCTPMigrationProposal();
    vm.stopPrank();
  }

  function test_cancelExistingCCTPMigrationProposal_ResetsExcludedTokens() public {
    // Arrange: Set up silo designation for the test chain
    vm.startPrank(OWNER);
    uint64[] memory removes = new uint64[](0);
    SiloedLockReleaseTokenPool.SiloConfigUpdate[] memory adds = new SiloedLockReleaseTokenPool.SiloConfigUpdate[](1);
    adds[0] = SiloedLockReleaseTokenPool.SiloConfigUpdate({remoteChainSelector: DEST_CHAIN_SELECTOR, rebalancer: OWNER});
    s_usdcTokenPool.updateSiloDesignations(removes, adds);

    // Provide some amount of liquidity to DEST_CHAIN_SELECTOR and use that amount for exclusion
    uint256 amount = 100e6;
    s_USDCToken.approve(address(s_usdcTokenPool), amount);
    s_usdcTokenPool.provideSiloedLiquidity(DEST_CHAIN_SELECTOR, amount);

    // Arrange: Propose migration and exclude some tokens
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    // Exclude some tokens
    s_usdcTokenPool.excludeTokensFromBurn(DEST_CHAIN_SELECTOR, amount);

    // Verify tokens were excluded
    assertEq(s_usdcTokenPool.getExcludedTokensByChain(DEST_CHAIN_SELECTOR), amount);

    // Act: Cancel the migration proposal
    s_usdcTokenPool.cancelExistingCCTPMigrationProposal();
    vm.stopPrank();

    // Assert: Verify excluded tokens were reset
    assertEq(s_usdcTokenPool.getExcludedTokensByChain(DEST_CHAIN_SELECTOR), 0);
  }

  function test_cancelExistingCCTPMigrationProposal_EmitsEvent() public {
    // Arrange: Propose a migration
    vm.startPrank(OWNER);
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    // Act & Assert: Expect the CCTPMigrationCancelled event
    vm.expectEmit();
    emit SiloedUSDCTokenPool.CCTPMigrationCancelled(DEST_CHAIN_SELECTOR);
    s_usdcTokenPool.cancelExistingCCTPMigrationProposal();
    vm.stopPrank();
  }
}
