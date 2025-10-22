// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_validateLockOrBurn is TokenPoolV2Setup {
  function test_validateLockOrBurn() public {
    Pool.LockOrBurnInV1 memory lockOrBurnIn = _buildLockOrBurnIn(1000e18);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed(DEST_CHAIN_SELECTOR, address(s_token), lockOrBurnIn.amount);

    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(lockOrBurnIn, 0);
  }

  function test_validateLockOrBurn_WithFastFinality() public {
    uint16 minBlockConfirmation = 8;
    RateLimiter.Config memory outboundFastConfig = RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24});
    RateLimiter.Config memory inboundFastConfig = RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24});
    TokenPool.CustomFinalityRateLimitConfigArgs[] memory rateLimitArgs =
      new TokenPool.CustomFinalityRateLimitConfigArgs[](1);
    rateLimitArgs[0] = TokenPool.CustomFinalityRateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundRateLimiterConfig: outboundFastConfig,
      inboundRateLimiterConfig: inboundFastConfig
    });
    vm.startPrank(OWNER);
    s_tokenPool.applyFinalityConfigUpdates(minBlockConfirmation, rateLimitArgs);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _buildLockOrBurnIn(1000e18);

    vm.expectEmit();
    emit TokenPool.CustomFinalityOutboundRateLimitConsumed(DEST_CHAIN_SELECTOR, address(s_token), lockOrBurnIn.amount);

    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(lockOrBurnIn, minBlockConfirmation);

    RateLimiter.TokenBucket memory bucket = s_tokenPool.getFastOutboundBucket(DEST_CHAIN_SELECTOR);
    assertEq(bucket.tokens, outboundFastConfig.capacity - lockOrBurnIn.amount);
  }

  function test_validateLockOrBurn_RevertWhen_InvalidMinBlockConfirmation() public {
    uint16 minBlockConfirmation = 5;
    _applyCustomFinalityConfig(minBlockConfirmation);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _buildLockOrBurnIn(1000e18);

    vm.expectRevert(
      abi.encodeWithSelector(
        TokenPool.InvalidMinBlockConfirmation.selector, minBlockConfirmation - 1, minBlockConfirmation
      )
    );
    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(lockOrBurnIn, minBlockConfirmation - 1);
  }

  function _applyCustomFinalityConfig(
    uint16 minBlockConfirmation
  ) internal {
    TokenPool.CustomFinalityRateLimitConfigArgs[] memory rateLimitArgs =
      new TokenPool.CustomFinalityRateLimitConfigArgs[](1);
    rateLimitArgs[0] = TokenPool.CustomFinalityRateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24}),
      inboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24})
    });
    vm.startPrank(OWNER);
    s_tokenPool.applyFinalityConfigUpdates(minBlockConfirmation, rateLimitArgs);
  }

  function _buildLockOrBurnIn(
    uint256 amount
  ) internal view returns (Pool.LockOrBurnInV1 memory lockOrBurnIn) {
    return lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: s_sender,
      receiver: s_receiver,
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });
  }
}
