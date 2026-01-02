// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";

contract TokenPoolV2_validateLockOrBurn is AdvancedPoolHooksSetup {
  function test_validateLockOrBurn() public {
    Pool.LockOrBurnInV1 memory lockOrBurnIn = _buildLockOrBurnIn(1000e18);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed(DEST_CHAIN_SELECTOR, address(s_token), lockOrBurnIn.amount);

    uint256 fee = s_tokenPool.getFee(lockOrBurnIn, 0);
    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(lockOrBurnIn, 0, "", fee);
  }

  function test_validateLockOrBurn_RevertWhen_InvalidToken() public {
    address wrongToken = address(0x456);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _buildLockOrBurnIn(1000e18);
    lockOrBurnIn.localToken = wrongToken; // Invalid token address.

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidToken.selector, wrongToken));
    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(lockOrBurnIn, 0, "", 0);
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

    uint256 fee = s_tokenPool.getFee(lockOrBurnIn, type(uint16).max);
    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(lockOrBurnIn, type(uint16).max, "", fee);

    (RateLimiter.TokenBucket memory outboundBucket,) = s_tokenPool.getCurrentRateLimiterState(DEST_CHAIN_SELECTOR, true);
    assertEq(outboundBucket.tokens, outboundFastConfig.capacity - lockOrBurnIn.amount);
  }

  function test_validateLockOrBurn_WithFastFinality_ConsumesAfterFee() public {
    RateLimiter.Config memory outboundFastConfig = RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24});
    RateLimiter.Config memory inboundFastConfig = RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24});
    TokenPool.RateLimitConfigArgs[] memory rateLimitArgs = new TokenPool.RateLimitConfigArgs[](1);
    rateLimitArgs[0] = TokenPool.RateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      customBlockConfirmation: true,
      outboundRateLimiterConfig: outboundFastConfig,
      inboundRateLimiterConfig: inboundFastConfig
    });

    vm.startPrank(OWNER);
    s_tokenPool.setDynamicConfig(address(s_sourceRouter), address(0), address(0));
    // Enable custom block confirmation handling so consumption emits.
    s_tokenPool.setMinBlockConfirmation(1);
    s_tokenPool.setRateLimitConfig(rateLimitArgs);

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] = TokenPool.TokenTransferFeeConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      tokenTransferFeeConfig: IPoolV2.TokenTransferFeeConfig({
        destGasOverhead: 50_000,
        destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
        defaultBlockConfirmationFeeUSDCents: 0,
        customBlockConfirmationFeeUSDCents: 0,
        defaultBlockConfirmationTransferFeeBps: 0,
        customBlockConfirmationTransferFeeBps: 250, // 2.5%
        isEnabled: true
      })
    });
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));
    vm.stopPrank();

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _buildLockOrBurnIn(1_000e18);
    uint256 fee = s_tokenPool.getFee(lockOrBurnIn, type(uint16).max);
    uint256 expectedAmount = lockOrBurnIn.amount - fee;

    vm.expectEmit();
    emit TokenPool.CustomBlockConfirmationOutboundRateLimitConsumed(
      DEST_CHAIN_SELECTOR, address(s_token), expectedAmount
    );

    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(lockOrBurnIn, type(uint16).max, "", fee);

    (RateLimiter.TokenBucket memory outboundBucket,) = s_tokenPool.getCurrentRateLimiterState(DEST_CHAIN_SELECTOR, true);
    assertEq(outboundBucket.tokens, outboundFastConfig.capacity - expectedAmount);
  }

  function test_validateLockOrBurn_ConsumesRateLimitAfterFee() public {
    uint16 defaultFeeBps = 250; // 2.5%
    vm.startPrank(OWNER);
    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] = TokenPool.TokenTransferFeeConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      tokenTransferFeeConfig: IPoolV2.TokenTransferFeeConfig({
        destGasOverhead: 50_000,
        destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
        defaultBlockConfirmationFeeUSDCents: 0,
        customBlockConfirmationFeeUSDCents: 0,
        defaultBlockConfirmationTransferFeeBps: defaultFeeBps,
        customBlockConfirmationTransferFeeBps: 0,
        isEnabled: true
      })
    });
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));
    vm.stopPrank();

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _buildLockOrBurnIn(1_000e18);
    uint256 fee = s_tokenPool.getFee(lockOrBurnIn, 0);
    uint256 expectedAmount = lockOrBurnIn.amount - fee;

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed(DEST_CHAIN_SELECTOR, address(s_token), expectedAmount);

    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.validateLockOrBurn(lockOrBurnIn, 0, "", fee);

    (RateLimiter.TokenBucket memory outboundBucket,) =
      s_tokenPool.getCurrentRateLimiterState(DEST_CHAIN_SELECTOR, false);
    assertEq(outboundBucket.tokens, _getOutboundRateLimiterConfig().capacity - expectedAmount);
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
    s_tokenPool.validateLockOrBurn(_buildLockOrBurnIn(1000e18), minBlockConfirmation - 1, "", 0);
  }

  function test_validateLockOrBurn_RevertWhen_CustomBlockConfirmationsNotEnabled() public {
    vm.startPrank(OWNER);
    s_tokenPool.setMinBlockConfirmation(0);

    vm.startPrank(s_allowedOnRamp);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.CustomBlockConfirmationsNotEnabled.selector));
    s_tokenPool.validateLockOrBurn(_buildLockOrBurnIn(1e18), 1, "", 0);
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
