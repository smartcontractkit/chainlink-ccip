// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenAdminRegistry} from "../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {HybridLockReleaseUSDCTokenPool} from "../../pools/USDC/HybridLockReleaseUSDCTokenPool.sol";
import {ERC20LockBox} from "../../pools/ERC20LockBox.sol";
import {RateLimiter} from "../../libraries/RateLimiter.sol";
import {MCMSForkTest} from "./MCMSForkTest.t.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract USDCForkTest is MCMSForkTest {
  address private s_usdc_token_pool_address_sepolia;

  address private s_usdc_token_address_sepolia;

  address private s_timelock_address_sepolia;

  address private s_lockbox_address_adi_testnet;

  address private s_lockbox_address_jovay_testnet;

  address private s_lockbox_address_ink_testnet;

  uint256 private s_adi_withdraw_amount;
  uint256 private s_jovay_withdraw_amount;
  uint256 private s_ink_withdraw_amount;

  uint256 private s_forkId;
  bytes[] private s_payloads;

  function setUp() public {
    // Skip test if RPC_URL is not set (e.g., in CI without .env)
    string memory rpcUrl = vm.envOr("RPC_URL", string(""));
    vm.skip(bytes(rpcUrl).length == 0);

    s_forkId = vm.createFork(rpcUrl);

    // Load payloads dynamically based on PAYLOAD_COUNT
    uint256 payloadCount = vm.envUint("PAYLOAD_COUNT");
    s_payloads = new bytes[](payloadCount);
    for (uint256 i = 0; i < payloadCount; i++) {
      s_payloads[i] = vm.envBytes(string.concat("PAYLOAD_", vm.toString(i + 1)));
    }

    s_usdc_token_pool_address_sepolia = vm.envAddress("USDC_TOKEN_POOL_ADDRESS_SEPOLIA");
    s_usdc_token_address_sepolia = vm.envAddress("USDC_TOKEN_ADDRESS_SEPOLIA");
    s_timelock_address_sepolia = vm.envAddress("TIMELOCK_ADDRESS_SEPOLIA");
    s_lockbox_address_adi_testnet = vm.envAddress("LOCKBOX_ADDRESS_ADI_TESTNET");
    s_lockbox_address_jovay_testnet = vm.envAddress("LOCKBOX_ADDRESS_JOVAY_TESTNET");
    s_lockbox_address_ink_testnet = vm.envAddress("LOCKBOX_ADDRESS_INK_TESTNET");
    s_adi_withdraw_amount = vm.envUint("ADI_WITHDRAW_AMOUNT");
    s_jovay_withdraw_amount = vm.envUint("JOVAY_WITHDRAW_AMOUNT");
    s_ink_withdraw_amount = vm.envUint("INK_WITHDRAW_AMOUNT");
  }

  function testFork_Migration() public {
    vm.selectFork(s_forkId);

    // Check locked tokens for remote chain selectors (ADI/Jovay/Ink)
    // Check balanceOf the timelock address on the USDC token contract
    // Check balanceOf each LockBox address on the USDC token contract
    // Apply first paylod (withdraw liquidity)
    // Re-check all the above
    // Apply the second payload (approve lockbox to spend timelock's USDC)
    // Apply the third payload (deposit USDC into the lockbox)
    // Re-check all the above

    uint64 remoteChainSelectorADI = 9418205736192840573; // ADI Testnet
    uint64 remoteChainSelectorJovay = 945045181441419236; // Jovay Testnet
    uint64 remoteChainSelectorInk = 9763904284804119144; // Ink Testnet

    uint256 adiTestnetLockedTokens = HybridLockReleaseUSDCTokenPool(s_usdc_token_pool_address_sepolia).getLockedTokensForChain(remoteChainSelectorADI);
    uint256 jovayTestnetLockedTokens = HybridLockReleaseUSDCTokenPool(s_usdc_token_pool_address_sepolia).getLockedTokensForChain(remoteChainSelectorJovay);
    uint256 inkTestnetLockedTokens = HybridLockReleaseUSDCTokenPool(s_usdc_token_pool_address_sepolia).getLockedTokensForChain(remoteChainSelectorInk);
    
    uint256 timelockBalance = IERC20(s_usdc_token_address_sepolia).balanceOf(s_timelock_address_sepolia);
    
    uint256 adiTestnetLockBoxBalance = IERC20(s_usdc_token_address_sepolia).balanceOf(s_lockbox_address_adi_testnet);
    uint256 jovayTestnetLockBoxBalance = IERC20(s_usdc_token_address_sepolia).balanceOf(s_lockbox_address_jovay_testnet);
    uint256 inkTestnetLockBoxBalance = IERC20(s_usdc_token_address_sepolia).balanceOf(s_lockbox_address_ink_testnet);

    // Locked Tokens should be >= withdraw amounts
    assertEq(adiTestnetLockedTokens, s_adi_withdraw_amount, "ADI Testnet locked tokens should be >= withdraw amount");
    assertEq(jovayTestnetLockedTokens, s_jovay_withdraw_amount, "Jovay Testnet locked tokens should be >= withdraw amount");
    assertEq(inkTestnetLockedTokens, s_ink_withdraw_amount, "Ink Testnet locked tokens should be >= withdraw amount");

    // Apply first payload (withdraw liquidity)
    _applyPayload(s_timelock_address_sepolia, s_payloads[0]);

    // Timelock should have the withdraw amounts
    assertEq(timelockBalance, s_adi_withdraw_amount + s_jovay_withdraw_amount + s_ink_withdraw_amount, "Timelock balance should be >= withdraw amounts");

    // LockBoxes should have 0 balance still (no deposit yet)
    assertEq(adiTestnetLockBoxBalance, 0, "ADI Testnet lock box balance should be 0");
    assertEq(jovayTestnetLockBoxBalance, 0, "Jovay Testnet lock box balance should be 0");
    assertEq(inkTestnetLockBoxBalance, 0, "Ink Testnet lock box balance should be 0");

    //Apply the second payload (approve lockbox to spend timelock's USDC)
    _applyPayload(s_timelock_address_sepolia, s_payloads[1]);

    // Apply the third payload (deposit USDC into the lockbox)
    _applyPayload(s_timelock_address_sepolia, s_payloads[2]);

    // Re-check Locked Tokens
    uint256 newadiTestnetLockedTokens = HybridLockReleaseUSDCTokenPool(s_usdc_token_pool_address_sepolia).getLockedTokensForChain(remoteChainSelectorADI);
    uint256 newjovayTestnetLockedTokens = HybridLockReleaseUSDCTokenPool(s_usdc_token_pool_address_sepolia).getLockedTokensForChain(remoteChainSelectorJovay);
    uint256 newinkTestnetLockedTokens = HybridLockReleaseUSDCTokenPool(s_usdc_token_pool_address_sepolia).getLockedTokensForChain(remoteChainSelectorInk);

    // Locked Tokens should be adiTestnetLockedTokens - s_adi_withdraw_amount
    // Jovay Testnet locked tokens should be jovayTestnetLockedTokens - s_jovay_withdraw_amount
    // Ink Testnet locked tokens should be inkTestnetLockedTokens - s_ink_withdraw_amount
    assertEq(newadiTestnetLockedTokens, adiTestnetLockedTokens - s_adi_withdraw_amount, "ADI Testnet locked tokens should be adiTestnetLockedTokens - s_adi_withdraw_amount");
    assertEq(newjovayTestnetLockedTokens, jovayTestnetLockedTokens - s_jovay_withdraw_amount, "Jovay Testnet locked tokens should be jovayTestnetLockedTokens - s_jovay_withdraw_amount");
    assertEq(newinkTestnetLockedTokens, inkTestnetLockedTokens - s_ink_withdraw_amount, "Ink Testnet locked tokens should be inkTestnetLockedTokens - s_ink_withdraw_amount");

    // Re-check Timelock balance
    uint256 newtimelockBalance = IERC20(s_usdc_token_address_sepolia).balanceOf(s_timelock_address_sepolia);

    // Timelock balance should be 0
    assertEq(newtimelockBalance, 0, "Timelock balance should be 0");

    // Re-check LockBox balances
    uint256 newadiTestnetLockBoxBalance = IERC20(s_usdc_token_address_sepolia).balanceOf(s_lockbox_address_adi_testnet);
    uint256 newjovayTestnetLockBoxBalance = IERC20(s_usdc_token_address_sepolia).balanceOf(s_lockbox_address_jovay_testnet);
    uint256 newinkTestnetLockBoxBalance = IERC20(s_usdc_token_address_sepolia).balanceOf(s_lockbox_address_ink_testnet);

    // LockBox balances should be s_adi_withdraw_amount
    assertEq(newadiTestnetLockBoxBalance, s_adi_withdraw_amount, "ADI Testnet lock box balance should be s_adi_withdraw_amount");
    assertEq(newjovayTestnetLockBoxBalance, s_jovay_withdraw_amount, "Jovay Testnet lock box balance should be s_jovay_withdraw_amount");
    assertEq(newinkTestnetLockBoxBalance, s_ink_withdraw_amount, "Ink Testnet lock box balance should be s_ink_withdraw_amount");
  }
}