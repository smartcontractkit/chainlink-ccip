// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SiloedLockReleaseTokenPool} from "../../../../pools/SiloedLockReleaseTokenPool.sol";
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

  function test_proposeCCTPMigration_RevertWhen_ZeroChainSelector() public {
    vm.expectRevert(abi.encodeWithSelector(SiloedLockReleaseTokenPool.InvalidChainSelector.selector, 0));
    s_usdcTokenPool.proposeCCTPMigration(0);
  }

  function test_proposeCCTPMigration_RevertWhen_ExistingProposal() public {
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.ExistingMigrationProposal.selector));
    s_usdcTokenPool.proposeCCTPMigration(SOURCE_CHAIN_SELECTOR);
  }

  function test_proposeCCTPMigration_RevertWhen_ChainAlreadyMigrated() public {
    uint64[] memory removes = new uint64[](0);
    SiloedLockReleaseTokenPool.SiloConfigUpdate[] memory adds = new SiloedLockReleaseTokenPool.SiloConfigUpdate[](1);
    adds[0] = SiloedLockReleaseTokenPool.SiloConfigUpdate({remoteChainSelector: DEST_CHAIN_SELECTOR, rebalancer: OWNER});
    s_usdcTokenPool.updateSiloDesignations(removes, adds);

    // First propose and execute a migration to mark the chain as migrated
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    // Set up the circle migrator address
    address circleMigrator = makeAddr("circleMigrator");
    s_usdcTokenPool.setCircleMigratorAddress(circleMigrator);

    // Provide some liquidity and exclude tokens to allow burning
    s_USDCToken.approve(address(s_usdcTokenPool), type(uint256).max);
    s_usdcTokenPool.provideSiloedLiquidity(DEST_CHAIN_SELECTOR, 1000e6);
    s_usdcTokenPool.excludeTokensFromBurn(DEST_CHAIN_SELECTOR, 100e6);

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
