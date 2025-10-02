// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {RateLimiter} from "../../../libraries/RateLimiter.sol";

import {TokenPool as TokenPoolV1} from "../../../pools/TokenPool.sol";
import {TokenPool} from "../../../poolsV2/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_releaseOrMint is TokenPoolV2Setup {
  function test_releaseOrMint() public {
    uint256 amount = 100e18;
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(amount);
    vm.expectEmit();
    emit TokenPoolV1.ReleasedOrMinted({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: s_allowedOffRamp,
      recipient: releaseOrMintIn.receiver,
      amount: amount
    });

    vm.startPrank(s_allowedOffRamp);
    Pool.ReleaseOrMintOutV1 memory result = s_tokenPool.releaseOrMint(releaseOrMintIn, 0);
    assertEq(result.destinationAmount, amount);
  }

  function test_releaseOrMint_WithFastFinality() public {
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

    uint256 amount = 100e18;
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(amount);

    vm.startPrank(s_allowedOffRamp);
    Pool.ReleaseOrMintOutV1 memory result = s_tokenPool.releaseOrMint(releaseOrMintIn, finalityThreshold);
    assertEq(result.destinationAmount, amount);

    RateLimiter.TokenBucket memory inboundBucket = s_tokenPool.getFastInboundBucket(DEST_CHAIN_SELECTOR);
    assertEq(inboundBucket.capacity, inboundConfig.capacity);
    assertEq(inboundBucket.tokens, inboundConfig.capacity - amount);
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
