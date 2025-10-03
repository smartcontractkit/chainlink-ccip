// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {TokenPool as TokenPoolV1} from "../../../pools/TokenPool.sol";
import {TokenPool} from "../../../poolsV2/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_validateReleaseOrMint is TokenPoolV2Setup {
  uint256 internal constant AMOUNT = 100e18;

  function test_validateReleaseOrMint() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(AMOUNT);

    vm.expectEmit();
    emit TokenPoolV1.InboundRateLimitConsumed(DEST_CHAIN_SELECTOR, address(s_token), AMOUNT);

    vm.startPrank(s_allowedOffRamp);
    uint256 localAmount = s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT, 0);

    assertEq(localAmount, AMOUNT);
  }

  function test_validateReleaseOrMint_WithFastFinality() public {
    uint16 finalityThreshold = 8;
    uint16 fastTransferFeeBps = 500;
    uint256 maxAmountPerRequest = 1000e18;
    RateLimiter.Config memory outboundConfig = RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24});
    RateLimiter.Config memory inboundConfig = RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24});
    TokenPool.FastFinalityRateLimitConfigArgs[] memory rateLimitArgs =
      new TokenPool.FastFinalityRateLimitConfigArgs[](1);
    rateLimitArgs[0] = TokenPool.FastFinalityRateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundRateLimiterConfig: outboundConfig,
      inboundRateLimiterConfig: inboundConfig
    });
    vm.startPrank(OWNER);
    s_tokenPool.applyFinalityConfigUpdates(finalityThreshold, fastTransferFeeBps, maxAmountPerRequest, rateLimitArgs);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(AMOUNT);

    vm.expectEmit();
    emit TokenPool.FastTransferInboundRateLimitConsumed(DEST_CHAIN_SELECTOR, address(s_token), AMOUNT);

    vm.startPrank(s_allowedOffRamp);
    uint256 localAmount = s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT, finalityThreshold);

    assertEq(localAmount, AMOUNT);

    RateLimiter.TokenBucket memory inboundBucket = s_tokenPool.getFastInboundBucket(DEST_CHAIN_SELECTOR);
    assertEq(inboundBucket.tokens, inboundConfig.capacity - AMOUNT);
  }

  function test_validateReleaseOrMint_RateLimitLocalAmount() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(AMOUNT);

    // Pretend the local amount is 10x the source amount and assert the rate limit is applied on the local amount.
    uint256 localAmount = AMOUNT * 10;

    vm.expectEmit();
    emit TokenPoolV1.InboundRateLimitConsumed({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: releaseOrMintIn.localToken,
      amount: localAmount
    });

    vm.startPrank(s_allowedOffRamp);
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, localAmount, 0);
  }

  function test_validateReleaseOrMint_InvalidToken() public {
    address wrongToken = address(0x456);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(AMOUNT);
    releaseOrMintIn.localToken = wrongToken; // Invalid token address.

    vm.expectRevert(abi.encodeWithSelector(TokenPoolV1.InvalidToken.selector, wrongToken));
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT, 0);
  }

  function test_validateReleaseOrMint_CursedByRMN() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(AMOUNT);

    // Mock RMN to be cursed
    vm.mockCall(
      address(s_mockRMNRemote),
      abi.encodeWithSignature("isCursed(bytes16)", bytes16(uint128(DEST_CHAIN_SELECTOR))),
      abi.encode(true)
    );

    vm.expectRevert(TokenPoolV1.CursedByRMN.selector);
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT, 0);
  }

  function test_validateReleaseOrMint_InvalidOffRamp() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(AMOUNT);

    // Mock router to return false for isOffRamp
    vm.mockCall(
      address(s_sourceRouter),
      abi.encodeWithSelector(IRouter.isOffRamp.selector, DEST_CHAIN_SELECTOR, s_allowedOffRamp),
      abi.encode(false)
    );

    vm.expectRevert(abi.encodeWithSelector(TokenPoolV1.CallerIsNotARampOnRouter.selector, s_allowedOffRamp));
    vm.startPrank(s_allowedOffRamp);
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT, 0);
  }

  function test_validateReleaseOrMint_InvalidSourcePool() public {
    address invalidPool = address(0x789);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _buildReleaseOrMintIn(AMOUNT);
    releaseOrMintIn.sourcePoolAddress = abi.encode(invalidPool);

    vm.expectRevert(abi.encodeWithSelector(TokenPoolV1.InvalidSourcePoolAddress.selector, abi.encode(invalidPool)));
    vm.startPrank(s_allowedOffRamp);
    s_tokenPool.validateReleaseOrMint(releaseOrMintIn, AMOUNT, 0);
  }

  function _applyFastFinalityConfig(
    uint16 finalityThreshold,
    uint16 fastTransferFeeBps,
    uint256 maxAmountPerRequest
  ) internal {
    TokenPool.FastFinalityRateLimitConfigArgs[] memory rateLimitArgs =
      new TokenPool.FastFinalityRateLimitConfigArgs[](1);
    rateLimitArgs[0] = TokenPool.FastFinalityRateLimitConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24}),
      inboundRateLimiterConfig: RateLimiter.Config({isEnabled: true, capacity: 1e24, rate: 1e24})
    });
    vm.startPrank(OWNER);
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
