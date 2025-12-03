// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenAdminRegistry} from "../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {MCMSForkTest} from "./MCMSForkTest.t.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

interface ILockReleaseTokenPool {
  function getRebalancer() external view returns (address);
}

contract LiquidityMigration is MCMSForkTest {
  address private s_oldPoolAddress;
  address private s_newPoolAddress;
  address private s_tokenAddress;
  address private s_tokenAdminRegistryAddress;
  address private s_rebalancerAddress;
  address private s_timelockAddress;
  uint256 private s_forkId;
  bytes[] private s_payloads;
  uint256 private s_transferAmount;

  function setUp() public {
    // Skip test if required env vars are not set (e.g., in CI without .env)
    string memory rpcUrl = vm.envOr("RPC_URL", string(""));
    uint256 payloadCount = vm.envOr("PAYLOAD_COUNT", uint256(0));
    s_transferAmount = vm.envOr("TRANSFER_AMOUNT", uint256(0));
    s_oldPoolAddress = vm.envOr("OLD_POOL_ADDRESS", address(0));
    s_newPoolAddress = vm.envOr("NEW_POOL_ADDRESS", address(0));
    s_tokenAddress = vm.envOr("TOKEN_ADDRESS", address(0));
    s_tokenAdminRegistryAddress = vm.envOr("TOKEN_ADMIN_REGISTRY_ADDRESS", address(0));
    s_rebalancerAddress = vm.envOr("REBALANCER_ADDRESS", address(0));
    s_timelockAddress = vm.envOr("TIMELOCK_ADDRESS", address(0));

    // Skip if any required env var is missing
    bool shouldSkip = bytes(rpcUrl).length == 0 || payloadCount == 0 || s_transferAmount == 0
      || s_oldPoolAddress == address(0) || s_newPoolAddress == address(0) || s_tokenAddress == address(0)
      || s_tokenAdminRegistryAddress == address(0) || s_rebalancerAddress == address(0) || s_timelockAddress == address(0);
    vm.skip(shouldSkip);

    s_forkId = vm.createFork(rpcUrl);

    // Load payloads dynamically based on PAYLOAD_COUNT
    s_payloads = new bytes[](payloadCount);
    for (uint256 i = 0; i < payloadCount; i++) {
      s_payloads[i] = vm.envBytes(string.concat("PAYLOAD_", vm.toString(i + 1)));
    }
  }

  function testMigration() public {
    vm.selectFork(s_forkId);

    // Get the token balance for the old pool
    uint256 oldPoolBalance = IERC20(s_tokenAddress).balanceOf(s_oldPoolAddress);

    // Get the token balance for the new pool
    uint256 newPoolBalance = IERC20(s_tokenAddress).balanceOf(s_newPoolAddress);

    // Apply the liquidity migration payloads
    for (uint256 i = 0; i < s_payloads.length; i++) {
      _applyPayload(s_timelockAddress, s_payloads[i]);
    }

    // Get the token config for the token
    TokenAdminRegistry.TokenConfig memory cfg =
      TokenAdminRegistry(s_tokenAdminRegistryAddress).getTokenConfig(s_tokenAddress);
    assertEq(cfg.tokenPool, s_newPoolAddress, "Registry should have the new Token Pool");

    // Get the token balance for the old pool
    uint256 updatedOldPoolBalance = IERC20(s_tokenAddress).balanceOf(s_oldPoolAddress);

    // Get the token balance for the new pool
    uint256 updatedNewPoolBalance = IERC20(s_tokenAddress).balanceOf(s_newPoolAddress);

    // Ensure that the old pool balance has decreased by the expected amounts
    assertEq(
      oldPoolBalance - updatedOldPoolBalance,
      s_transferAmount,
      "TOKEN balance should have decreased by the expected amount"
    );

    // Ensure that the new pool balance has increased by the expected amounts
    assertEq(
      updatedNewPoolBalance - newPoolBalance,
      s_transferAmount,
      "TOKEN balance should have increased by the expected amount"
    );

    // Ensure that the rebalancer is set to the expected address
    assertEq(
      ILockReleaseTokenPool(s_newPoolAddress).getRebalancer(),
      s_rebalancerAddress,
      "Rebalancer should be set to the expected address"
    );
  }
}
