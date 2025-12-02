// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolSetup} from "./TokenPoolSetup.t.sol";

contract TokenPool_setChainRateLimiterConfigs is TokenPoolSetup {
  uint64 internal s_remoteChainSelector;

  function setUp() public virtual override {
    TokenPoolSetup.setUp();

    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(address(2));

    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
    s_remoteChainSelector = 123124;
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: s_remoteChainSelector,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(3)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(new uint64[](0), chainUpdates);
  }

  function testFuzz_setRateLimitConfig(
    uint128 capacity,
    uint128 rate,
    uint32 newTime,
    bool customBlockConfirmations
  ) public {
    rate = uint128(bound(rate, 0, capacity));
    newTime = uint32(bound(newTime, block.timestamp + 1, type(uint32).max));
    vm.warp(newTime);

    RateLimiter.Config memory newOutboundConfig = RateLimiter.Config({isEnabled: true, capacity: capacity, rate: rate});
    RateLimiter.Config memory newInboundConfig =
      RateLimiter.Config({isEnabled: true, capacity: capacity / 2, rate: rate / 2});

    TokenPool.RateLimitConfigArgs[] memory rateLimitConfigArgs = new TokenPool.RateLimitConfigArgs[](1);
    rateLimitConfigArgs[0] = TokenPool.RateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      customBlockConfirmation: customBlockConfirmations,
      outboundRateLimiterConfig: newOutboundConfig,
      inboundRateLimiterConfig: newInboundConfig
    });

    vm.expectEmit();
    emit TokenPool.RateLimitConfigured(
      DEST_CHAIN_SELECTOR, customBlockConfirmations, newOutboundConfig, newInboundConfig
    );

    s_tokenPool.setRateLimitConfig(rateLimitConfigArgs);

    (RateLimiter.TokenBucket memory outboundAfter, RateLimiter.TokenBucket memory inboundAfter) =
      s_tokenPool.getCurrentRateLimiterState(DEST_CHAIN_SELECTOR, customBlockConfirmations);

    _assertRateLimiterState(outboundAfter, newOutboundConfig);
    assertEq(outboundAfter.lastUpdated, newTime);

    _assertRateLimiterState(inboundAfter, newInboundConfig);
    assertEq(inboundAfter.lastUpdated, newTime);
  }

  function test_setRateLimitConfig_customBlockConfs() public {
    RateLimiter.Config memory outboundConfig = RateLimiter.Config({isEnabled: true, capacity: 1e21, rate: 5e20});
    RateLimiter.Config memory inboundConfig = RateLimiter.Config({isEnabled: true, capacity: 2e21, rate: 1e21});
    TokenPool.RateLimitConfigArgs[] memory args = new TokenPool.RateLimitConfigArgs[](1);
    args[0] = TokenPool.RateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      customBlockConfirmation: true,
      outboundRateLimiterConfig: outboundConfig,
      inboundRateLimiterConfig: inboundConfig
    });

    vm.expectEmit();
    emit TokenPool.RateLimitConfigured(DEST_CHAIN_SELECTOR, true, outboundConfig, inboundConfig);

    s_tokenPool.setRateLimitConfig(args);

    (RateLimiter.TokenBucket memory outboundBucket, RateLimiter.TokenBucket memory inboundBucket) =
      s_tokenPool.getCurrentRateLimiterState(DEST_CHAIN_SELECTOR, true);

    assertTrue(outboundBucket.isEnabled);
    _assertRateLimiterState(outboundBucket, outboundConfig);
    assertTrue(inboundBucket.isEnabled);
    _assertRateLimiterState(inboundBucket, inboundConfig);
  }

  function test_setRateLimitConfig_RevertWhen_InvalidRateLimitRate_Outbound() public {
    TokenPool.RateLimitConfigArgs[] memory args = new TokenPool.RateLimitConfigArgs[](1);
    args[0] = TokenPool.RateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      customBlockConfirmation: true,
      outboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1, rate: 2}),
      inboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1, rate: 1})
    });

    vm.expectRevert(
      abi.encodeWithSelector(RateLimiter.InvalidRateLimitRate.selector, args[0].outboundRateLimiterConfig)
    );
    s_tokenPool.setRateLimitConfig(args);
  }

  function _assertRateLimiterState(
    RateLimiter.TokenBucket memory bucket,
    RateLimiter.Config memory config
  ) internal pure {
    assertEq(bucket.capacity, config.capacity);
    assertEq(bucket.rate, config.rate);
    assertEq(bucket.tokens, config.capacity);
  }

  // Reverts

  function test_setRateLimitConfig_RevertWhen_Unauthorized_OnlyOwnerOrRateLimitAdmin() public {
    TokenPool.RateLimitConfigArgs[] memory rateLimitConfigArgs = new TokenPool.RateLimitConfigArgs[](1);
    rateLimitConfigArgs[0] = TokenPool.RateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      customBlockConfirmation: false,
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.Unauthorized.selector, STRANGER));
    s_tokenPool.setRateLimitConfig(rateLimitConfigArgs);
  }

  function test_setRateLimitConfig_RevertWhen_NonExistentChain() public {
    uint64 wrongChainSelector = 9084102894;

    TokenPool.RateLimitConfigArgs[] memory rateLimitConfigArgs = new TokenPool.RateLimitConfigArgs[](1);
    rateLimitConfigArgs[0] = TokenPool.RateLimitConfigArgs({
      remoteChainSelector: wrongChainSelector,
      customBlockConfirmation: false,
      outboundRateLimiterConfig: RateLimiter.Config({isEnabled: false, capacity: 0, rate: 0}),
      inboundRateLimiterConfig: RateLimiter.Config({isEnabled: false, capacity: 0, rate: 0})
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPool.NonExistentChain.selector, wrongChainSelector));
    s_tokenPool.setRateLimitConfig(rateLimitConfigArgs);
  }
}
