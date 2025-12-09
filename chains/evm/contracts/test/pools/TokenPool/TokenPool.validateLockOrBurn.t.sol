// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";

contract TokenPoolV2_validateLockOrBurn is AdvancedPoolHooksSetup {
  function test_validateLockOrBurn() public {
    Pool.LockOrBurnInV1 memory lockOrBurnIn = _buildLockOrBurnIn(1000e18);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed(DEST_CHAIN_SELECTOR, address(s_token), lockOrBurnIn.amount);

    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(lockOrBurnIn, 0, "");
  }

  function test_validateLockOrBurn_WithFastFinality() public {
    uint16 minBlockConfirmation = 5;
    RateLimiter.Config memory outboundFastConfig = RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24});
    RateLimiter.Config memory inboundFastConfig = RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24});
    TokenPool.RateLimitConfigArgs[] memory rateLimitArgs = new TokenPool.RateLimitConfigArgs[](1);
    rateLimitArgs[0] = TokenPool.RateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      customBlockConfirmation: true,
      outboundRateLimiterConfig: outboundFastConfig,
      inboundRateLimiterConfig: inboundFastConfig
    });
    s_tokenPool.setMinBlockConfirmation(minBlockConfirmation);
    s_tokenPool.setRateLimitConfig(rateLimitArgs);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _buildLockOrBurnIn(1000e18);

    vm.expectEmit();
    emit TokenPool.CustomBlockConfirmationOutboundRateLimitConsumed(
      DEST_CHAIN_SELECTOR, address(s_token), lockOrBurnIn.amount
    );

    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(lockOrBurnIn, type(uint16).max, "");

    (RateLimiter.TokenBucket memory outboundBucket,) = s_tokenPool.getCurrentRateLimiterState(DEST_CHAIN_SELECTOR, true);
    assertEq(outboundBucket.tokens, outboundFastConfig.capacity - lockOrBurnIn.amount);
  }

  function test_validateLockOrBurn_RevertWhen_InvalidMinBlockConfirmation() public {
    uint16 minBlockConfirmation = 5;
    s_tokenPool.setMinBlockConfirmation(minBlockConfirmation);
    vm.startPrank(s_allowedOnRamp);

    vm.expectRevert(
      abi.encodeWithSelector(
        TokenPool.InvalidMinBlockConfirmation.selector, minBlockConfirmation - 1, minBlockConfirmation
      )
    );
    s_tokenPool.validateLockOrBurn(_buildLockOrBurnIn(1000e18), minBlockConfirmation - 1, "");
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
