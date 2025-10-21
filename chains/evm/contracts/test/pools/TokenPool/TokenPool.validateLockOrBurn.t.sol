// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

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
    (, uint16 bps) = s_tokenPool.getCustomFinalityConfig();
  }

  function test_validateLockOrBurn_WithFastFinality() public {
    uint16 finalityThreshold = 8;
    uint16 customFinalityTransferFeeBps = 500; // 5%
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
    s_tokenPool.applyFinalityConfigUpdates(finalityThreshold, customFinalityTransferFeeBps, rateLimitArgs);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _buildLockOrBurnIn(1000e18);

    vm.expectEmit();
    emit TokenPool.CustomFinalityOutboundRateLimitConsumed(DEST_CHAIN_SELECTOR, address(s_token), lockOrBurnIn.amount);

    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(lockOrBurnIn, finalityThreshold);

    RateLimiter.TokenBucket memory bucket =
      s_tokenPool.getCurrentCustomFinalityRateLimiterState(DEST_CHAIN_SELECTOR, IPoolV2.MessageDirection.Outbound);
    assertEq(bucket.tokens, outboundFastConfig.capacity - lockOrBurnIn.amount);
  }

  function test_validateLockOrBurn_RevertWhen_InvalidFinality() public {
    uint16 finalityThreshold = 5;
    uint16 customFinalityTransferFeeBps = 500;
    _applyCustomFinalityConfig(finalityThreshold, customFinalityTransferFeeBps);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _buildLockOrBurnIn(1000e18);

    vm.expectRevert(
      abi.encodeWithSelector(TokenPool.InvalidFinality.selector, finalityThreshold - 1, finalityThreshold)
    );
    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(lockOrBurnIn, finalityThreshold - 1);
  }

  function _applyCustomFinalityConfig(uint16 finalityThreshold, uint16 customFinalityTransferFeeBps) internal {
    TokenPool.CustomFinalityRateLimitConfigArgs[] memory rateLimitArgs =
      new TokenPool.CustomFinalityRateLimitConfigArgs[](1);
    rateLimitArgs[0] = TokenPool.CustomFinalityRateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24}),
      inboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24})
    });
    vm.startPrank(OWNER);
    s_tokenPool.applyFinalityConfigUpdates(finalityThreshold, customFinalityTransferFeeBps, rateLimitArgs);
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
