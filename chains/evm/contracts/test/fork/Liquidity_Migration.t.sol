// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenAdminRegistry} from "../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {MCMSForkTest} from "./MCMSForkTest.t.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface LockReleaseTokenPool {
    function getRebalancer() external view returns (address);
}

contract LiquidityMigration is MCMSForkTest {
    address private oldPoolAddress;
    address private newPoolAddress;
    address private tokenAddress;
    address private tokenAdminRegistryAddress;
    address private rebalancerAddress;
    address private timelockAddress;
    uint256 private forkId;
    bytes[] private payloads;
    uint256 private transferAmount;

    function setUp() public {
        forkId = vm.createFork(vm.envString("RPC_URL"));
        
        // Load payloads dynamically based on PAYLOAD_COUNT
        uint256 payloadCount = vm.envUint("PAYLOAD_COUNT");
        payloads = new bytes[](payloadCount);
        for (uint256 i = 0; i < payloadCount; i++) {
            payloads[i] = vm.envBytes(string.concat("PAYLOAD_", vm.toString(i + 1)));
        }
        
        transferAmount = vm.envUint("TRANSFER_AMOUNT");
        oldPoolAddress = vm.envAddress("OLD_POOL_ADDRESS");
        newPoolAddress = vm.envAddress("NEW_POOL_ADDRESS");
        tokenAddress = vm.envAddress("TOKEN_ADDRESS");
        tokenAdminRegistryAddress = vm.envAddress("TOKEN_ADMIN_REGISTRY_ADDRESS");
        rebalancerAddress = vm.envAddress("REBALANCER_ADDRESS");
        timelockAddress = vm.envAddress("TIMELOCK_ADDRESS");
    }

    function testMigration() public {
        vm.selectFork(forkId);

        // Get the token balance for the old pool
        uint256 oldPoolBalance = IERC20(tokenAddress).balanceOf(oldPoolAddress);

        // Get the token balance for the new pool
        uint256 newPoolBalance = IERC20(tokenAddress).balanceOf(newPoolAddress);

        // Apply the liquidity migration payloads
        for (uint256 i = 0; i < payloads.length; i++) {
            applyPayload(timelockAddress, payloads[i]);
        }

        // Get the token config for the token
        TokenAdminRegistry.TokenConfig memory cfg = TokenAdminRegistry(tokenAdminRegistryAddress).getTokenConfig(tokenAddress);
        assertEq(cfg.tokenPool, newPoolAddress, "Registry should have the new Token Pool");

        // Get the token balance for the old pool
        uint256 updatedOldPoolBalance = IERC20(tokenAddress).balanceOf(oldPoolAddress);

        // Get the token balance for the new pool
        uint256 updatedNewPoolBalance = IERC20(tokenAddress).balanceOf(newPoolAddress);

        // Ensure that the old pool balance has decreased by the expected amounts
        assertEq(oldPoolBalance - updatedOldPoolBalance, transferAmount, "TOKEN balance should have decreased by the expected amount");

        // Ensure that the new pool balance has increased by the expected amounts
        assertEq(updatedNewPoolBalance - newPoolBalance, transferAmount, "TOKEN balance should have increased by the expected amount");

        // Ensure that the rebalancer is set to the expected address
        assertEq(LockReleaseTokenPool(newPoolAddress).getRebalancer(), rebalancerAddress, "Rebalancer should be set to the expected address");
    }
}
