// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ILiquidityContainer} from "../../../../interfaces/ILiquidityContainer.sol";

import {TokenPool} from "../../../../pools/TokenPool.sol";
import {HybridLockReleaseUSDCTokenPool} from "../../../../pools/USDC/HybridLockReleaseUSDCTokenPool.sol";
import {USDCBridgeMigratorSetup} from "./USDCBridgeMigratorSetup.t.sol";

contract USDCBridgeMigrator_withdrawLiquidity is USDCBridgeMigratorSetup {
  uint256 public constant LIQUIDITY_AMOUNT = 1e12;

  function setUp() public override {
    super.setUp();

    vm.startPrank(OWNER);

    // Designate the SOURCE_CHAIN as not using native-USDC, and so the L/R mechanism must be used instead
    uint64[] memory destChainAdds = new uint64[](1);
    destChainAdds[0] = DEST_CHAIN_SELECTOR;

    s_usdcTokenPool.updateChainSelectorMechanisms(new uint64[](0), destChainAdds);

    assertTrue(
      s_usdcTokenPool.shouldUseLockRelease(DEST_CHAIN_SELECTOR),
      "Lock/Release mech not configured for incoming message from DEST_CHAIN_SELECTOR"
    );

    vm.startPrank(OWNER);
    s_usdcTokenPool.setLiquidityProvider(DEST_CHAIN_SELECTOR, OWNER);

    // Add 1e12 liquidity so that there's enough to release
    vm.startPrank(s_usdcTokenPool.getLiquidityProvider(DEST_CHAIN_SELECTOR));

    s_token.approve(address(s_usdcTokenPool), type(uint256).max);

    s_usdcTokenPool.provideLiquidity(DEST_CHAIN_SELECTOR, LIQUIDITY_AMOUNT);
  }

  function test_withdrawLiquidity_Success() public {
    vm.startPrank(OWNER);

    vm.expectEmit();
    emit ILiquidityContainer.LiquidityRemoved(OWNER, LIQUIDITY_AMOUNT);

    s_usdcTokenPool.withdrawLiquidity(DEST_CHAIN_SELECTOR, LIQUIDITY_AMOUNT);

    assertEq(s_usdcTokenPool.getLockedTokensForChain(DEST_CHAIN_SELECTOR), 0);
  }

  // Reverts
  function test_RevertWhen_LanePausedForCCTPMigration() public {
    vm.startPrank(OWNER);

    s_usdcTokenPool.proposeCCTPMigration(DEST_CHAIN_SELECTOR);

    vm.expectRevert(
      abi.encodeWithSelector(HybridLockReleaseUSDCTokenPool.LanePausedForCCTPMigration.selector, DEST_CHAIN_SELECTOR)
    );

    s_usdcTokenPool.withdrawLiquidity(DEST_CHAIN_SELECTOR, LIQUIDITY_AMOUNT);
  }
}
