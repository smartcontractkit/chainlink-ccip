// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../../libraries/Pool.sol";
import {SiloedLockReleaseTokenPool} from "../../../../pools/SiloedLockReleaseTokenPool.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {SiloedUSDCTokenPool} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

contract SiloedUSDCTokenPool_burnLockedUSDC is SiloedUSDCTokenPoolSetup {
  address public CIRCLE = makeAddr("CIRCLE CCTP Migrator");

  function setUp() public override {
    super.setUp();

    // Mark DEST_CHAIN_SELECTOR as siloed with OWNER as the rebalancer
    uint64[] memory removes = new uint64[](0);
    SiloedLockReleaseTokenPool.SiloConfigUpdate[] memory adds = new SiloedLockReleaseTokenPool.SiloConfigUpdate[](1);
    adds[0] = SiloedLockReleaseTokenPool.SiloConfigUpdate({remoteChainSelector: DEST_CHAIN_SELECTOR, rebalancer: OWNER});
    s_usdcTokenPool.updateSiloDesignations(removes, adds);
  }

  function test_burnLockedUSDC() public {
    uint256 amount = 1e6;

    deal(address(s_USDCToken), address(s_usdcTokenPool), amount);

    vm.startPrank(s_routerAllowedOnRamp);

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_USDCToken),
      sender: address(s_routerAllowedOnRamp),
      amount: amount
    });

    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(STRANGER),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_USDCToken)
      })
    );

    // Ensure that the tokens are properly locked
    assertEq(s_USDCToken.balanceOf(address(s_lockBox)), amount, "Incorrect token amount in the tokenPool");

    vm.startPrank(OWNER);

    vm.expectEmit();
    emit SiloedUSDCTokenPool.CircleMigratorAddressSet(CIRCLE);

    s_usdcTokenPool.setCircleMigratorAddress(CIRCLE);

    vm.expectEmit();
    emit SiloedUSDCTokenPool.CCTPMigrationProposed(DEST_CHAIN_SELECTOR);

    // Propose the migration to CCTP
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    assertEq(
      s_usdcTokenPool.getCurrentProposedCCTPChainMigration(),
      DEST_CHAIN_SELECTOR,
      "Current proposed chain migration does not match expected for DEST_CHAIN_SELECTOR"
    );

    // Impersonate the set circle address and execute the proposal
    vm.startPrank(CIRCLE);

    vm.expectEmit();
    emit SiloedUSDCTokenPool.CCTPMigrationExecuted(DEST_CHAIN_SELECTOR, amount);

    // Ensure the call to the burn function is properly
    vm.expectCall(address(s_USDCToken), abi.encodeWithSelector(bytes4(keccak256("burn(uint256)")), amount));
    s_usdcTokenPool.burnLockedUSDC();

    // Assert that the tokens were actually burned
    assertEq(s_USDCToken.balanceOf(address(s_usdcTokenPool)), 0, "Tokens were not burned out of the tokenPool");

    // Ensure the proposal slot was cleared and there's no tokens locked for the destination chain anymore
    assertEq(s_usdcTokenPool.getCurrentProposedCCTPChainMigration(), 0);
    assertEq(
      s_usdcTokenPool.getAvailableTokens(DEST_CHAIN_SELECTOR),
      0,
      "No tokens should be locked for DEST_CHAIN_SELECTOR after CCTP-approved burn"
    );
  }

  // Reverts

  function test_burnLockedUSDC_RevertWhen_InvalidPermissions() public {
    // Deal some tokens to the token pool
    uint256 amount = 1000e6;
    deal(address(s_USDCToken), address(s_usdcTokenPool), amount);

    // Lock or burn those tokens for the destination chain
    vm.startPrank(s_routerAllowedOnRamp);
    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(STRANGER),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_USDCToken)
      })
    );

    vm.startPrank(OWNER);
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    // Set the circle migrator address for later, but don't start pranking as it yet
    s_usdcTokenPool.setCircleMigratorAddress(CIRCLE);

    // Should fail because only Circle can call this function
    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.OnlyCircle.selector));
    s_usdcTokenPool.burnLockedUSDC();
  }

  function test_burnLockedUSDC_RevertWhen_NoMigrationProposalPending() public {
    vm.startPrank(OWNER);
    s_usdcTokenPool.setCircleMigratorAddress(CIRCLE);
    vm.stopPrank();

    vm.startPrank(CIRCLE);

    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.NoMigrationProposalPending.selector));
    s_usdcTokenPool.burnLockedUSDC();
  }

  function test_burnLockedUSDC_RevertWhen_TokenLockingNotAllowedAfterMigration() public {
    // Deal some tokens to the token pool
    uint256 amount = 1000e6;
    deal(address(s_USDCToken), address(s_usdcTokenPool), amount);

    // Lock or burn those tokens for the destination chain
    vm.startPrank(s_routerAllowedOnRamp);
    s_usdcTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: abi.encodePacked(STRANGER),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_USDCToken)
      })
    );

    vm.startPrank(OWNER);
    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    // Set the circle migrator address
    s_usdcTokenPool.setCircleMigratorAddress(CIRCLE);

    // Execute the migration
    vm.startPrank(CIRCLE);
    s_usdcTokenPool.burnLockedUSDC();
    vm.stopPrank();

    // Try to provide liquidity after migration and expect revert
    vm.startPrank(OWNER);
    s_USDCToken.approve(address(s_usdcTokenPool), type(uint256).max);
    vm.expectRevert(
      abi.encodeWithSelector(SiloedUSDCTokenPool.TokenLockingNotAllowedAfterMigration.selector, DEST_CHAIN_SELECTOR)
    );
    s_usdcTokenPool.provideSiloedLiquidity(DEST_CHAIN_SELECTOR, 500e6);
    vm.stopPrank();
  }
}
