// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenAdminRegistry} from "../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {HybridLockReleaseUSDCTokenPool} from "../../pools/USDC/HybridLockReleaseUSDCTokenPool.sol";
import {RateLimiter} from "../../libraries/RateLimiter.sol";
import {MCMSForkTest} from "./MCMSForkTest.t.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract HybridLockReleaseUSDCTokenPoolSetup is MCMSForkTest {
  address private s_poolAddress;
  address private s_jovayPoolAddress;
  address private s_pharosPoolAddress;
  address private s_tokenAddress;
  address private s_jovayTokenAddress;
  address private s_pharosTokenAddress;
  address private s_tokenAdminRegistryAddress;
  address private s_timelockAddress;
  bool private s_shouldUseLockRelease;
  uint64 private s_jovayChainSelector;
  uint64 private s_pharosChainSelector;
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

    s_jovayChainSelector = uint64(vm.envUint("JOVAY_CHAIN_SELECTOR"));
    s_pharosChainSelector = uint64(vm.envUint("PHAROS_CHAIN_SELECTOR"));
    s_poolAddress = vm.envAddress("POOL_ADDRESS");
    s_jovayPoolAddress = vm.envAddress("JOVAY_POOL_ADDRESS");
    s_pharosPoolAddress = vm.envAddress("PHAROS_POOL_ADDRESS");
    s_tokenAddress = vm.envAddress("TOKEN_ADDRESS");
    s_tokenAdminRegistryAddress = vm.envAddress("TOKEN_ADMIN_REGISTRY_ADDRESS");
    s_timelockAddress = vm.envAddress("TIMELOCK_ADDRESS");
    s_jovayTokenAddress = vm.envAddress("JOVAY_TOKEN_ADDRESS");
    s_pharosTokenAddress = vm.envAddress("PHAROS_TOKEN_ADDRESS");
  }

  function testFork_Migration() public {
    vm.selectFork(s_forkId);

    // Apply the payloads
    for (uint256 i = 0; i < s_payloads.length; i++) {
      _applyPayload(s_timelockAddress, s_payloads[i]);
    }

    // Get the Liquidity Provider
    address liquidityProviderJovay = HybridLockReleaseUSDCTokenPool(s_poolAddress).getLiquidityProvider(s_jovayChainSelector);
    address liquidityProviderPharos = HybridLockReleaseUSDCTokenPool(s_poolAddress).getLiquidityProvider(s_pharosChainSelector);

    // Check the shouldUseLockRelease function
    bool shouldUseLockReleaseJovay = HybridLockReleaseUSDCTokenPool(s_poolAddress).shouldUseLockRelease(s_jovayChainSelector);
    bool shouldUseLockReleasePharos = HybridLockReleaseUSDCTokenPool(s_poolAddress).shouldUseLockRelease(s_pharosChainSelector);

    // Check the new supported chains
    bool isSupportedChainJovay = HybridLockReleaseUSDCTokenPool(s_poolAddress).isSupportedChain(s_jovayChainSelector);
    bool isSupportedChainPharos = HybridLockReleaseUSDCTokenPool(s_poolAddress).isSupportedChain(s_pharosChainSelector);

    // Check the new remote pools
    bytes[] memory remotePoolsJovay = HybridLockReleaseUSDCTokenPool(s_poolAddress).getRemotePools(s_jovayChainSelector);
    bytes[] memory remotePoolsPharos = HybridLockReleaseUSDCTokenPool(s_poolAddress).getRemotePools(s_pharosChainSelector);

    // Check the new remote token
    bytes memory remoteTokenJovay = HybridLockReleaseUSDCTokenPool(s_poolAddress).getRemoteToken(s_jovayChainSelector);
    bytes memory remoteTokenPharos = HybridLockReleaseUSDCTokenPool(s_poolAddress).getRemoteToken(s_pharosChainSelector);

    // Check the in and outbound rate limits
    RateLimiter.TokenBucket memory inboundRateLimitJovay = HybridLockReleaseUSDCTokenPool(s_poolAddress).getCurrentInboundRateLimiterState(s_jovayChainSelector);
    RateLimiter.TokenBucket memory outboundRateLimitJovay = HybridLockReleaseUSDCTokenPool(s_poolAddress).getCurrentOutboundRateLimiterState(s_jovayChainSelector);
    RateLimiter.TokenBucket memory inboundRateLimitPharos = HybridLockReleaseUSDCTokenPool(s_poolAddress).getCurrentInboundRateLimiterState(s_pharosChainSelector);
    RateLimiter.TokenBucket memory outboundRateLimitPharos = HybridLockReleaseUSDCTokenPool(s_poolAddress).getCurrentOutboundRateLimiterState(s_pharosChainSelector);

    // Ensure that the liquidiy provider matches the timelock address -- Jovay
    assertEq(
      liquidityProviderJovay,
      s_timelockAddress,
      "Liquidity provider should be set to the timelock address"
    );

    // Ensure that the liquidiy provider matches the timelock address -- Pharos
    assertEq(
      liquidityProviderPharos,
      s_timelockAddress,
      "Liquidity provider should be set to the timelock address"
    );  

    // Ensure that the shouldUseLockRelease function matches the expected value -- Jovay
    assertEq(shouldUseLockReleaseJovay, true, "shouldUseLockRelease should be true");

    // Ensure that the shouldUseLockRelease function matches the expected value -- Pharos
    assertEq(shouldUseLockReleasePharos, true, "shouldUseLockRelease should be true");

    // Ensure that the isSupportedChain function matches the expected value -- Jovay
    assertEq(isSupportedChainJovay, true, "isSupportedChain should be true");

    // Ensure that the isSupportedChain function matches the expected value -- Pharos
    assertEq(isSupportedChainPharos, true, "isSupportedChain should be true");

    // Ensure that the remote pools array is not empty -- Jovay
    assertTrue(remotePoolsJovay.length > 0, "Remote pools array should not be empty for Jovay");
    assertEq(abi.decode(remotePoolsJovay[0], (address)), s_jovayPoolAddress, "First remote pool should be set to the pool address for Jovay");

    // Ensure that the remote pools array is not empty -- Pharos
    assertTrue(remotePoolsPharos.length > 0, "Remote pools array should not be empty for Pharos");
    assertEq(abi.decode(remotePoolsPharos[0], (address)), s_pharosPoolAddress, "First remote pool should be set to the pool address for Pharos");

    // Ensure that the remote token matches the expected value -- Jovay
    assertEq(abi.decode(remoteTokenJovay, (address)), s_jovayTokenAddress, "Remote token should be set to the token address");

    // Ensure that the remote token matches the expected value -- Pharos
    assertEq(abi.decode(remoteTokenPharos, (address)), s_pharosTokenAddress, "Remote token should be set to the token address");

    // Ensure that the inbound rate limit matches the expected value -- Jovay
    assertEq(inboundRateLimitJovay.capacity, 0, "Inbound rate limit should be 0");
    assertEq(inboundRateLimitJovay.rate, 0, "Inbound rate limit should be 0");

    // Ensure that the outbound rate limit matches the expected value -- Jovay
    assertEq(outboundRateLimitJovay.capacity, 0, "Outbound rate limit should be 0");
    assertEq(outboundRateLimitJovay.rate, 0, "Outbound rate limit should be 0");

    // Ensure that the inbound rate limit matches the expected value -- Pharos
    assertEq(inboundRateLimitPharos.capacity, 0, "Inbound rate limit should be 0");
    assertEq(inboundRateLimitPharos.rate, 0, "Inbound rate limit should be 0");

    // Ensure that the outbound rate limit matches the expected value -- Pharos
    assertEq(outboundRateLimitPharos.capacity, 0, "Outbound rate limit should be 0");
    assertEq(outboundRateLimitPharos.rate, 0, "Outbound rate limit should be 0");

  }
}
