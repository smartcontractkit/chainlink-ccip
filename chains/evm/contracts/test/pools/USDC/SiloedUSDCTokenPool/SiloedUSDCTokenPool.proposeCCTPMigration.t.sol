// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SiloedUSDCTokenPool} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

contract SiloedUSDCTokenPool_proposeCCTPMigration is SiloedUSDCTokenPoolSetup {
  function test_proposeCCTPMigration() public {
    vm.expectEmit();
    emit SiloedUSDCTokenPool.CCTPMigrationProposed(DEST_CHAIN_SELECTOR);
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);
  }

  function test_proposeCCTPMigration_AfterCancellation() public {
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);
    s_usdcTokenPool.cancelExistingCCTPMigrationProposal();

    s_usdcTokenPool.proposeCCTPMigration(SOURCE_CHAIN_SELECTOR);

    assertEq(s_usdcTokenPool.getCurrentProposedCCTPChainMigration(), SOURCE_CHAIN_SELECTOR);
  }

  // Reverts

  function test_proposeCCTPMigration_RevertWhen_ExistingProposal() public {
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.ExistingMigrationProposal.selector));
    s_usdcTokenPool.proposeCCTPMigration(SOURCE_CHAIN_SELECTOR);
  }

  function test_proposeCCTPMigration_RevertWhen_ChainAlreadyMigrated() public {
    // First propose and execute a migration to mark the chain as migrated
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    // Set up the circle migrator address
    address circleMigrator = makeAddr("circleMigrator");
    s_usdcTokenPool.setCircleMigratorAddress(circleMigrator);

    // Provide some liquidity to the lockbox and exclude tokens to allow burning
    deal(address(s_USDCToken), address(s_destLockBox), 1000e6);
    s_usdcTokenPool.excludeTokensFromBurn(DEST_CHAIN_SELECTOR, 100e6);
    s_usdcTokenPool.setLockedUSDCToBurn(DEST_CHAIN_SELECTOR, 1000e6);

    // Stop the OWNER prank so that the circle migrator can be set
    vm.stopPrank();

    // Execute the migration
    vm.prank(circleMigrator);
    s_usdcTokenPool.burnLockedUSDC();

    vm.startPrank(OWNER);
    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.ChainAlreadyMigrated.selector, DEST_CHAIN_SELECTOR));
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);
  }
}
