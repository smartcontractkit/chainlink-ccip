// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {RateLimiter} from "../../../libraries/RateLimiter.sol";

import {TokenPool as TokenPoolV1} from "../../../pools/TokenPool.sol";
import {TokenPool} from "../../../poolsV2/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_validateReleaseOrMint is TokenPoolV2Setup {
  function test_validateReleaseOrMint() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(100e18);

    vm.expectEmit();
    emit TokenPoolV1.InboundRateLimitConsumed(DEST_CHAIN_SELECTOR, address(s_token), 100e18);

    vm.startPrank(s_allowedOffRamp);
    uint256 localAmount = s_tokenPool.validateReleaseOrMint(releaseOrMintIn, 0);

    assertEq(localAmount, 100e18);
  }

  function test_validateReleaseOrMint_WithFastFinality() public {
    uint16 finalityThreshold = 8;
    uint16 fastTransferFeeBps = 500;
    uint256 maxAmountPerRequest = 1000e18;
    RateLimiter.Config memory outboundConfig = RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24});
    RateLimiter.Config memory inboundConfig = RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24});
    TokenPool.FastTransferRateLimitConfigArgs[] memory rateLimitArgs =
      new TokenPool.FastTransferRateLimitConfigArgs[](1);
    rateLimitArgs[0] = TokenPool.FastTransferRateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundRateLimiterConfig: outboundConfig,
      inboundRateLimiterConfig: inboundConfig
    });
    s_tokenPool.applyFinalityConfigUpdates(finalityThreshold, fastTransferFeeBps, maxAmountPerRequest, rateLimitArgs);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(100e18);

    vm.expectEmit();
    emit TokenPool.FastTransferInboundRateLimitConsumed(DEST_CHAIN_SELECTOR, address(s_token), 100e18);

    vm.startPrank(s_allowedOffRamp);
    uint256 localAmount = s_tokenPool.validateReleaseOrMint(releaseOrMintIn, finalityThreshold);

    assertEq(localAmount, 100e18);

    RateLimiter.TokenBucket memory inboundBucket = s_tokenPool.getFastInboundBucket(DEST_CHAIN_SELECTOR);
    assertEq(inboundBucket.tokens, inboundConfig.capacity - 100e18);
  }

  function test_validateReleaseOrMint_RevertWhen_InvalidFinality() public {
    uint16 finalityThreshold = 5;
    uint16 fastTransferFeeBps = 500;
    uint256 maxAmountPerRequest = 1000e18;
    _applyFastFinalityConfig(finalityThreshold, fastTransferFeeBps, maxAmountPerRequest);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(100e18);

    vm.expectRevert(
      abi.encodeWithSelector(TokenPool.InvalidFinality.selector, finalityThreshold - 1, finalityThreshold)
    );
    vm.startPrank(s_allowedOffRamp);
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, finalityThreshold - 1);
  }

  function test_validateReleaseOrMint_RevertWhen_AmountExceedsMaxPerRequest() public {
    uint16 finalityThreshold = 8;
    uint16 fastTransferFeeBps = 500;
    uint256 maxAmountPerRequest = 50e18;
    _applyFastFinalityConfig(finalityThreshold, fastTransferFeeBps, maxAmountPerRequest);

    uint256 amount = maxAmountPerRequest + 1;
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(amount);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.AmountExceedsMaxPerRequest.selector, amount, maxAmountPerRequest));
    vm.startPrank(s_allowedOffRamp);
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, finalityThreshold);
  }

  function _applyFastFinalityConfig(
    uint16 finalityThreshold,
    uint16 fastTransferFeeBps,
    uint256 maxAmountPerRequest
  ) internal {
    TokenPool.FastTransferRateLimitConfigArgs[] memory rateLimitArgs =
      new TokenPool.FastTransferRateLimitConfigArgs[](1);
    rateLimitArgs[0] = TokenPool.FastTransferRateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24}),
      inboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24})
    });
    s_tokenPool.applyFinalityConfigUpdates(finalityThreshold, fastTransferFeeBps, maxAmountPerRequest, rateLimitArgs);
  }

  function _buildReleaseOrMintIn(
    uint256 amount
  ) internal view returns (Pool.ReleaseOrMintInV1 memory) {
    return Pool.ReleaseOrMintInV1({
      originalSender: abi.encode(OWNER),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      receiver: OWNER,
      sourceDenominatedAmount: amount,
      localToken: address(s_token),
      sourcePoolAddress: abi.encode(s_initialRemotePool),
      sourcePoolData: abi.encode(uint256(s_token.decimals())),
      offchainTokenData: ""
    });
  }
}
