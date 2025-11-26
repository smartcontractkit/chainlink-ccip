// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenAdminRegistry} from "../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {MCMSForkTest} from "./MCMSForkTest.t.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface LockReleaseTokenPool {
    function getRebalancer() external view returns (address);
}

contract WSTLINKMigration is MCMSForkTest {
    address private constant OLD_WSTLINK_POOL = 0x21377fe476Fb8587CbAFd47155093597Fa4df45E;
    address private constant NEW_WSTLINK_POOL = 0xF6403CF6E954a43699097322e0867C63d653C2D0;
    address private constant WSTLINK_TOKEN = 0x911D86C72155c33993d594B0Ec7E6206B4C803da;
    address private constant TOKEN_ADMIN_REGISTRY = 0xb22764f98dD05c789929716D677382Df22C05Cb6;

    address private constant ETHEREUM_REBALANCER = 0xB351EC0FEaF4B99FdFD36b484d9EC90D0422493D; 

    address private constant TIMELOCK = 0x44835bBBA9D40DEDa9b64858095EcFB2693c9449;

    uint256 private ethereumForkId;
    bytes private ethereumPayload;
    bytes private ethereumPayload2;
    uint256 private ethereumTransferAmount;

    function setUp() public {
        ethereumForkId = vm.createFork(vm.envString("ETHEREUM_RPC_URL"));
        ethereumPayload = vm.envBytes("ETHEREUM_PAYLOAD");
        ethereumPayload2 = vm.envBytes("ETHEREUM_PAYLOAD_2");
        ethereumTransferAmount = vm.envUint("ETHEREUM_TRANSFER_AMOUNT");
    }

    function testMigration() public {
        vm.selectFork(ethereumForkId);

        // Get the token balance for the old pool
        uint256 oldPoolBalance = IERC20(WSTLINK_TOKEN).balanceOf(OLD_WSTLINK_POOL);

        // Get the token balance for the new pool
        uint256 newPoolBalance = IERC20(WSTLINK_TOKEN).balanceOf(NEW_WSTLINK_POOL);

        // Apply the liquidity migration
        applyPayload(TIMELOCK, ethereumPayload);
        applyPayload(TIMELOCK, ethereumPayload2);

        // Get the token config for the WSTLINK token
        TokenAdminRegistry.TokenConfig memory cfg = TokenAdminRegistry(TOKEN_ADMIN_REGISTRY).getTokenConfig(WSTLINK_TOKEN);
        assertEq(cfg.tokenPool, NEW_WSTLINK_POOL, "Registry should have the new WSTLINK pool");

        // Get the token balance for the old pool
        uint256 updatedOldPoolBalance = IERC20(WSTLINK_TOKEN).balanceOf(OLD_WSTLINK_POOL);

        // Get the token balance for the new pool
        uint256 updatedNewPoolBalance = IERC20(WSTLINK_TOKEN).balanceOf(NEW_WSTLINK_POOL);

        // Ensure that the old pool balance has decreased by the expected amounts
        assertEq(oldPoolBalance - updatedOldPoolBalance, ethereumTransferAmount, "WSTLINK balance should have decreased by the expected amount");

        // Ensure that the new pool balance has increased by the expected amounts
        assertEq(updatedNewPoolBalance - newPoolBalance, ethereumTransferAmount, "WSTLINK balance should have increased by the expected amount");

        // Ensure that the rebalancer is set to the expected address
        assertEq(LockReleaseTokenPool(NEW_WSTLINK_POOL).getRebalancer(), ETHEREUM_REBALANCER, "Rebalancer should be set to the expected address");
    }
}
