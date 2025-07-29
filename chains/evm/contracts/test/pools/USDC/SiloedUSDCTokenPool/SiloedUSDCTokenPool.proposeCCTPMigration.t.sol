// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SiloedLockReleaseTokenPool} from "../../../../pools/SiloedLockReleaseTokenPool.sol";
import {SiloedUSDCTokenPool} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

contract SiloedUSDCTokenPool_proposeCCTPMigration is SiloedUSDCTokenPoolSetup {
  function test_proposeCCTPMigration_Success() public {
    // Arrange: No existing migration proposal
    vm.startPrank(OWNER);
    
    // Act & Assert: Expect the CCTPMigrationProposed event
    vm.expectEmit();
    emit SiloedUSDCTokenPool.CCTPMigrationProposed(DEST_CHAIN_SELECTOR);
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);
    vm.stopPrank();
  }

  function test_proposeCCTPMigration_AfterCancellation() public {
    // Arrange: Propose and then cancel a migration
    vm.startPrank(OWNER);
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);
    s_usdcTokenPool.cancelExistingCCTPMigrationProposal();
    
    // Act: Propose a new migration after cancellation
    s_usdcTokenPool.proposeCCTPMigration(SOURCE_CHAIN_SELECTOR);
    vm.stopPrank();

    // Assert: Verify the new proposal was set
    assertEq(s_usdcTokenPool.getCurrentProposedCCTPChainMigration(), SOURCE_CHAIN_SELECTOR);
  }

  function test_proposeCCTPMigration_ZeroChainSelector() public {
    // Arrange: No existing migration proposal
    vm.startPrank(OWNER);
    
    // Act: Propose migration with zero chain selector
    s_usdcTokenPool.proposeCCTPMigration(0);
    vm.stopPrank();

    // Assert: Verify the proposal was set (zero is valid)
    assertEq(s_usdcTokenPool.getCurrentProposedCCTPChainMigration(), 0);
  }

  // Reverts

    function test_proposeCCTPMigration_RevertWhen_NotOwner() public {
    // Arrange: Non-owner caller
    vm.startPrank(STRANGER);
    
    // Act & Assert: Expect revert when non-owner tries to propose migration
    vm.expectRevert();
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);
    vm.stopPrank();
  }

  function test_proposeCCTPMigration_RevertWhen_ExistingProposal() public {
    // Arrange: Set up an existing migration proposal
    vm.startPrank(OWNER);
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);
    
    // Act & Assert: Try to propose another migration while one is pending
    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.ExistingMigrationProposal.selector));
    s_usdcTokenPool.proposeCCTPMigration(SOURCE_CHAIN_SELECTOR);
    vm.stopPrank();
  }

  function test_proposeCCTPMigration_RevertWhen_ChainAlreadyMigrated() public {
    // Arrange: Set up silo designation for the test chain
    vm.startPrank(OWNER);
    uint64[] memory removes = new uint64[](0);
    SiloedLockReleaseTokenPool.SiloConfigUpdate[] memory adds = new SiloedLockReleaseTokenPool.SiloConfigUpdate[](1);
    adds[0] =
      SiloedLockReleaseTokenPool.SiloConfigUpdate({remoteChainSelector: DEST_CHAIN_SELECTOR, rebalancer: OWNER});
    s_usdcTokenPool.updateSiloDesignations(removes, adds);
    vm.stopPrank();

    // Arrange: Set up a chain that has already been migrated
    vm.startPrank(OWNER);
    
    // First propose and execute a migration to mark the chain as migrated
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);
    
    // Set up the circle migrator address
    address circleMigrator = makeAddr("circleMigrator");
    s_usdcTokenPool.setCircleMigratorAddress(circleMigrator);
    
    // Provide some liquidity and exclude tokens to allow burning
    s_USDCToken.approve(address(s_usdcTokenPool), type(uint256).max);
    s_usdcTokenPool.provideSiloedLiquidity(DEST_CHAIN_SELECTOR, 1000e6);
    s_usdcTokenPool.excludeTokensFromBurn(DEST_CHAIN_SELECTOR, 100e6);
    
    // Execute the migration
    vm.startPrank(circleMigrator);
    s_usdcTokenPool.burnLockedUSDC();
    vm.stopPrank();
    
    // Act & Assert: Try to propose migration for the already migrated chain
    vm.startPrank(OWNER);
    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.ChainAlreadyMigrated.selector, DEST_CHAIN_SELECTOR));
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);
    vm.stopPrank();
  }
} 