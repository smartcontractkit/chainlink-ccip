// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenAdminRegistry} from "../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {MCMSForkTest} from "./MCMSForkTest.t.sol";

interface HybridLockReleaseUSDCTokenPool {
    function getLockedTokensForChain(uint64 chainSelector) external view returns (uint256);
    function getLiquidityProvider(uint64 chainSelector) external view returns (address);
}

contract USDCLiquidityMigration is MCMSForkTest {
    address private constant OLD_HYBRID_USDC_POOL = 0xc2e3A3C18ccb634622B57fF119a1C8C7f12e8C0c;
    address private constant NEW_HYBRID_USDC_POOL = 0x03D19033AdA17750D5BC2d8E325337D0748F9FEF;
    address private constant USDC_TOKEN = 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48;
    address private constant TOKEN_ADMIN_REGISTRY = 0xb22764f98dD05c789929716D677382Df22C05Cb6;

    address private constant BITLAYER_REBALANCER = 0x2728df4D22253004C017675bd609962cD641D797;
    uint64 private constant BITLAYER_SELECTOR = 7937294810946806131;

    address private constant BOB_REBALANCER = 0x450D55a4B4136805B0e5A6BB59377c71FC4FaCBb;
    uint64 private constant BOB_SELECTOR = 3849287863852499584;

    address private constant RONIN_REBALANCER = 0x0000000000000000000000000000000000000000;
    uint64 private constant RONIN_SELECTOR = 6916147374840168594;

    address private constant WEMIX_REBALANCER = 0x0000000000000000000000000000000000000000;
    uint64 private constant WEMIX_SELECTOR = 5142893604156789321;

    uint256 private ethereumForkId;
    bytes private ethereumPayload1;
    bytes private ethereumPayload2;
    bytes private ethereumPayload3;
    bytes private ethereumPayload4;
    bytes private ethereumPayload5;
    uint256 private bitlayerTransferAmount;
    uint256 private bobTransferAmount;
    uint256 private roninTransferAmount;
    uint256 private wemixTransferAmount;

    function setUp() public {
        ethereumForkId = vm.createFork(vm.envString("ETHEREUM_RPC_URL"));
        ethereumPayload1 = vm.envBytes("ETHEREUM_PAYLOAD_1");
        ethereumPayload2 = vm.envBytes("ETHEREUM_PAYLOAD_2");
        ethereumPayload3 = vm.envBytes("ETHEREUM_PAYLOAD_3");
        ethereumPayload4 = vm.envBytes("ETHEREUM_PAYLOAD_4");
        ethereumPayload5 = vm.envBytes("ETHEREUM_PAYLOAD_5");
        bitlayerTransferAmount = vm.envUint("BITLAYER_TRANSFER_AMOUNT");
        bobTransferAmount = vm.envUint("BOB_TRANSFER_AMOUNT");
        roninTransferAmount = vm.envUint("RONIN_TRANSFER_AMOUNT");
        wemixTransferAmount = vm.envUint("WEMIX_TRANSFER_AMOUNT");
    }

    function testMigration() public {
        vm.selectFork(ethereumForkId);
        
        // Get the balances of each silo on the old hybrid USDC pool prior to liquidity migration
        uint256 bitlayerBalanceOnOld = HybridLockReleaseUSDCTokenPool(OLD_HYBRID_USDC_POOL).getLockedTokensForChain(BITLAYER_SELECTOR);
        uint256 bobBalanceOnOld = HybridLockReleaseUSDCTokenPool(OLD_HYBRID_USDC_POOL).getLockedTokensForChain(BOB_SELECTOR);
        uint256 roninBalanceOnOld = HybridLockReleaseUSDCTokenPool(OLD_HYBRID_USDC_POOL).getLockedTokensForChain(RONIN_SELECTOR);
        uint256 wemixBalanceOnOld = HybridLockReleaseUSDCTokenPool(OLD_HYBRID_USDC_POOL).getLockedTokensForChain(WEMIX_SELECTOR);

        // Get the balances of each silo on the new hybrid USDC pool prior to liquidity migration
        uint256 bitlayerBalanceOnNew = HybridLockReleaseUSDCTokenPool(NEW_HYBRID_USDC_POOL).getLockedTokensForChain(BITLAYER_SELECTOR);
        uint256 bobBalanceOnNew = HybridLockReleaseUSDCTokenPool(NEW_HYBRID_USDC_POOL).getLockedTokensForChain(BOB_SELECTOR);
        uint256 roninBalanceOnNew = HybridLockReleaseUSDCTokenPool(NEW_HYBRID_USDC_POOL).getLockedTokensForChain(RONIN_SELECTOR);
        uint256 wemixBalanceOnNew = HybridLockReleaseUSDCTokenPool(NEW_HYBRID_USDC_POOL).getLockedTokensForChain(WEMIX_SELECTOR);

        // Apply the liquidity migration
        applyPayload(ethereumPayload1);
        applyPayload(ethereumPayload2);
        applyPayload(ethereumPayload3);
        applyPayload(ethereumPayload4);
        applyPayload(ethereumPayload5);

        TokenAdminRegistry.TokenConfig memory cfg = TokenAdminRegistry(TOKEN_ADMIN_REGISTRY).getTokenConfig(USDC_TOKEN);
        // assertEq(cfg.tokenPool, NEW_HYBRID_USDC_POOL, "Registry should have the new hybrid USDC pool");

        // Check the rebalancers on the old hybrid USDC pool
        assertEq(HybridLockReleaseUSDCTokenPool(OLD_HYBRID_USDC_POOL).getLiquidityProvider(BITLAYER_SELECTOR), BITLAYER_REBALANCER, "BitLayer rebalancer should match expected");
        assertEq(HybridLockReleaseUSDCTokenPool(OLD_HYBRID_USDC_POOL).getLiquidityProvider(BOB_SELECTOR), BOB_REBALANCER, "Bob rebalancer should match expected");
        assertEq(HybridLockReleaseUSDCTokenPool(OLD_HYBRID_USDC_POOL).getLiquidityProvider(RONIN_SELECTOR), RONIN_REBALANCER, "Ronin rebalancer should match expected");
        assertEq(HybridLockReleaseUSDCTokenPool(OLD_HYBRID_USDC_POOL).getLiquidityProvider(WEMIX_SELECTOR), WEMIX_REBALANCER, "WEMIX rebalancer should match expected");

        // Check rebalancers on the new hybrid USDC pool
        assertEq(HybridLockReleaseUSDCTokenPool(NEW_HYBRID_USDC_POOL).getLiquidityProvider(BITLAYER_SELECTOR), BITLAYER_REBALANCER, "BitLayer rebalancer should match expected");
        assertEq(HybridLockReleaseUSDCTokenPool(NEW_HYBRID_USDC_POOL).getLiquidityProvider(BOB_SELECTOR), BOB_REBALANCER, "Bob rebalancer should match expected");
        assertEq(HybridLockReleaseUSDCTokenPool(NEW_HYBRID_USDC_POOL).getLiquidityProvider(RONIN_SELECTOR), RONIN_REBALANCER, "Ronin rebalancer should match expected");
        assertEq(HybridLockReleaseUSDCTokenPool(NEW_HYBRID_USDC_POOL).getLiquidityProvider(WEMIX_SELECTOR), WEMIX_REBALANCER, "WEMIX rebalancer should match expected");

        // Ensure that the old hybrid USDC pool balance has decreased by the expected amounts
        assertEq(bitlayerBalanceOnOld - HybridLockReleaseUSDCTokenPool(OLD_HYBRID_USDC_POOL).getLockedTokensForChain(BITLAYER_SELECTOR), bitlayerTransferAmount, "BitLayer balance should have decreased by the expected amount");
        assertEq(bobBalanceOnOld - HybridLockReleaseUSDCTokenPool(OLD_HYBRID_USDC_POOL).getLockedTokensForChain(BOB_SELECTOR), bobTransferAmount, "Bob balance should have decreased by the expected amount");
        assertEq(roninBalanceOnOld - HybridLockReleaseUSDCTokenPool(OLD_HYBRID_USDC_POOL).getLockedTokensForChain(RONIN_SELECTOR), roninTransferAmount, "Ronin balance should have decreased by the expected amount");
        assertEq(wemixBalanceOnOld - HybridLockReleaseUSDCTokenPool(OLD_HYBRID_USDC_POOL).getLockedTokensForChain(WEMIX_SELECTOR), wemixTransferAmount, "WEMIX balance should have decreased by the expected amount");

        // Ensure that the new hybrid USDC pool balance has increased by the expected amounts
        assertEq(HybridLockReleaseUSDCTokenPool(NEW_HYBRID_USDC_POOL).getLockedTokensForChain(BITLAYER_SELECTOR), bitlayerBalanceOnNew + bitlayerTransferAmount, "BitLayer balance should have increased by the expected amount");
        assertEq(HybridLockReleaseUSDCTokenPool(NEW_HYBRID_USDC_POOL).getLockedTokensForChain(BOB_SELECTOR), bobBalanceOnNew + bobTransferAmount, "Bob balance should have increased by the expected amount");
        assertEq(HybridLockReleaseUSDCTokenPool(NEW_HYBRID_USDC_POOL).getLockedTokensForChain(RONIN_SELECTOR), roninBalanceOnNew + roninTransferAmount, "Ronin balance should have increased by the expected amount");
        assertEq(HybridLockReleaseUSDCTokenPool(NEW_HYBRID_USDC_POOL).getLockedTokensForChain(WEMIX_SELECTOR), wemixBalanceOnNew + wemixTransferAmount, "WEMIX balance should have increased by the expected amount");
    }
}
