// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../interfaces/IPoolV2.sol";
import {IRMN} from "../interfaces/IRMN.sol";

import {Client} from "../libraries/Client.sol";
import {Pool} from "../libraries/Pool.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";
import {TokenPool as TokenPoolV1} from "../pools/TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

abstract contract TokenPool is IPoolV2, TokenPoolV1 {
  using SafeERC20 for IERC20;
  using RateLimiter for RateLimiter.TokenBucket;

  error DuplicateCCV(address ccv);
  error InvalidDestBytesOverhead(uint32 destBytesOverhead);
  error InvalidFinality(uint16 requested, uint16 finalityThreshold);
  error AmountExceedsMaxPerRequest(uint256 requested, uint256 maximum);
  error TokenTransferFeeConfigNotEnabled(uint64 destChainSelector);
  error InvalidFastTransferFeeBps();
  error InvalidFinalityConfig();

  event CCVConfigUpdated(uint64 indexed remoteChainSelector, address[] outboundCCVs, address[] inboundCCVs);
  event FinalityConfigUpdated(uint16 finalityConfig, uint16 fastTransferFeeBps, uint256 maxAmountPerRequest);
  event TokenTransferFeeConfigUpdated(uint64 indexed destChainSelector, TokenTransferFeeConfig tokenTransferFeeConfig);
  event TokenTransferFeeConfigDeleted(uint64 indexed destChainSelector);
  /// @notice Emitted when pool fees are withdrawn.
  event PoolFeeWithdrawn(address indexed recipient, uint256 amount);
  event FastTransferOutboundRateLimitConsumed(uint64 indexed remoteChainSelector, address token, uint256 amount);
  event FastTransferInboundRateLimitConsumed(uint64 indexed remoteChainSelector, address token, uint256 amount);

  struct FastFinalityConfig {
    uint16 finalityThreshold; // ──╮ Minimum block depth on the source chain that token issuers consider sufficiently secure.
    //                             | 0 means the default finality.
    uint16 fastTransferFeeBps; // ─╯ Fee in basis points for fast transfers [0-10_000].
    uint256 maxAmountPerRequest; // Maximum amount allowed per transfer request.
    // Separate buckets isolate fast-finality limits so these transfers cannot deplete the primary pool rate limits.
    mapping(uint64 remoteChainSelector => RateLimiter.TokenBucket tokenBucketOutbound) outboundRateLimiterConfig;
    mapping(uint64 remoteChainSelector => RateLimiter.TokenBucket tokenBucketInbound) inboundRateLimiterConfig;
  }

  struct FastFinalityRateLimitConfigArgs {
    uint64 remoteChainSelector; // Remote chain selector.
    RateLimiter.Config outboundRateLimiterConfig; // Outbound rate limiter configuration.
    RateLimiter.Config inboundRateLimiterConfig; // Inbound rate limiter configuration.
  }

  struct CCVConfig {
    address[] outboundCCVs; // CCVs required for outgoing messages to the remote chain.
    address[] inboundCCVs; // CCVs required for incoming messages from the remote chain.
  }

  struct CCVConfigArg {
    uint64 remoteChainSelector;
    address[] outboundCCVs;
    address[] inboundCCVs;
  }

  /// @dev Struct with args for setting the token transfer fee configurations for a destination chain and a set of tokens.
  struct TokenTransferFeeConfigArgs {
    uint64 destChainSelector; // Destination chain selector.
    TokenTransferFeeConfig tokenTransferFeeConfig; // Token transfer fee configuration.
  }

  /// @notice The division factor for basis points (BPS). This also represents the maximum BPS fee for fast transfer.
  uint256 internal constant BPS_DIVIDER = 10_000;

  // Tracks fast-finality parameters and per-lane rate limit buckets for fast transfers.
  FastFinalityConfig internal s_finalityConfig;
  // Stores verifier (CCV) requirements keyed by remote chain selector.
  mapping(uint64 remoteChainSelector => CCVConfig ccvConfig) internal s_verifierConfig;
  // Optional token-transfer fee overrides keyed by destination chain selector.
  mapping(uint64 destChainSelector => TokenTransferFeeConfig tokenTransferFeeConfig) internal s_tokenTransferFeeConfig;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) TokenPoolV1(token, localTokenDecimals, allowlist, rmnProxy, router) {}

  // ================================================================
  // │                        Lock or Burn                          │
  // ================================================================

  /// @inheritdoc IPoolV2
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 finality,
    bytes calldata // tokenArgs
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory, uint256 destTokenAmount) {
    _validateLockOrBurn(lockOrBurnIn, finality);
    destTokenAmount = _applyFee(lockOrBurnIn, finality);
    _lockOrBurn(destTokenAmount);

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: destTokenAmount
    });

    return (
      Pool.LockOrBurnOutV1({
        destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
        destPoolData: _encodeLocalDecimals()
      }),
      destTokenAmount
    );
  }

  // ================================================================
  // │                      Release or Mint                         │
  // ================================================================

  /// @inheritdoc IPoolV2
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint16 finality
  ) public virtual override(IPoolV2) returns (Pool.ReleaseOrMintOutV1 memory) {
    uint256 localAmount = _calculateLocalAmount(
      releaseOrMintIn.sourceDenominatedAmount, _parseRemoteDecimals(releaseOrMintIn.sourcePoolData)
    );

    _validateReleaseOrMint(releaseOrMintIn, localAmount, finality);

    _releaseOrMint(releaseOrMintIn.receiver, localAmount);

    emit ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      recipient: releaseOrMintIn.receiver,
      amount: localAmount
    });

    return Pool.ReleaseOrMintOutV1({destinationAmount: localAmount});
  }

  // ================================================================
  // │                         Validation                           │
  // ================================================================

  /// @notice Validates the lock or burn request, and enforces rate limits.
  /// @dev The validation covers token support, RMN curse status, allowlist membership, onRamp access, and
  /// rate limiting for both standard and fast-transfer lanes.
  /// @param lockOrBurnIn The input to validate. Must reference a supported token, onRamp, and remote chain.
  /// @param finality The finality depth requested by the message. A value of zero uses the standard lane.
  function _validateLockOrBurn(Pool.LockOrBurnInV1 calldata lockOrBurnIn, uint16 finality) internal {
    if (!isSupportedToken(lockOrBurnIn.localToken)) revert InvalidToken(lockOrBurnIn.localToken);
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(lockOrBurnIn.remoteChainSelector)))) revert CursedByRMN();
    _checkAllowList(lockOrBurnIn.originalSender);

    _onlyOnRamp(lockOrBurnIn.remoteChainSelector);
    FastFinalityConfig storage finalityConfig = s_finalityConfig;
    uint256 amount = lockOrBurnIn.amount;
    if (finality != 0 && finalityConfig.finalityThreshold != 0) {
      if (finality < finalityConfig.finalityThreshold) {
        revert InvalidFinality(finality, finalityConfig.finalityThreshold);
      }
      if (amount > finalityConfig.maxAmountPerRequest) {
        revert AmountExceedsMaxPerRequest(amount, finalityConfig.maxAmountPerRequest);
      }

      finalityConfig.outboundRateLimiterConfig[lockOrBurnIn.remoteChainSelector]._consume(
        amount, lockOrBurnIn.localToken
      );
      emit FastTransferOutboundRateLimitConsumed(lockOrBurnIn.remoteChainSelector, lockOrBurnIn.localToken, amount);
    } else {
      _consumeOutboundRateLimit(lockOrBurnIn.remoteChainSelector, amount);
    }
  }

  /// @notice Validates a release or mint request and enforces the appropriate inbound rate limits.
  /// @dev The validation checks token support, RMN curse status, offRamp access, remote pool configuration,
  /// finality requirements, and consumes either the fast-transfer inbound bucket or the standard bucket.
  /// @param releaseOrMintIn The input to validate. The remote chain, pool, and token must all be configured.
  /// @param localAmount The amount to release or mint on the local chain after any decimal conversion.
  /// @param finality The finality depth requested by the message. A value of zero uses the standard lane.
  function _validateReleaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint256 localAmount,
    uint16 finality
  ) internal {
    if (!isSupportedToken(releaseOrMintIn.localToken)) revert InvalidToken(releaseOrMintIn.localToken);
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(releaseOrMintIn.remoteChainSelector)))) revert CursedByRMN();
    _onlyOffRamp(releaseOrMintIn.remoteChainSelector);

    // Validates that the source pool address is configured on this pool.
    if (!isRemotePool(releaseOrMintIn.remoteChainSelector, releaseOrMintIn.sourcePoolAddress)) {
      revert InvalidSourcePoolAddress(releaseOrMintIn.sourcePoolAddress);
    }

    if (finality != 0) {
      s_finalityConfig.inboundRateLimiterConfig[releaseOrMintIn.remoteChainSelector]._consume(
        localAmount, releaseOrMintIn.localToken
      );
      emit FastTransferInboundRateLimitConsumed(
        releaseOrMintIn.remoteChainSelector, releaseOrMintIn.localToken, localAmount
      );
    } else {
      _consumeInboundRateLimit(releaseOrMintIn.remoteChainSelector, localAmount);
    }
  }

  // ================================================================
  // │                     Finality                                 │
  // ================================================================

  /// @notice Updates the finality configuration for token transfers.
  function applyFinalityConfigUpdates(
    uint16 finalityThreshold,
    uint16 fastTransferFeeBps,
    uint256 maxAmountPerRequest,
    FastFinalityRateLimitConfigArgs[] calldata rateLimitConfigArgs
  ) external virtual onlyOwner {
    FastFinalityConfig storage finalityConfig = s_finalityConfig;
    finalityConfig.finalityThreshold = finalityThreshold;
    if (fastTransferFeeBps >= BPS_DIVIDER) {
      revert InvalidFastTransferFeeBps();
    }
    finalityConfig.fastTransferFeeBps = fastTransferFeeBps;
    finalityConfig.maxAmountPerRequest = maxAmountPerRequest;
    _setFastFinalityRateLimitConfig(rateLimitConfigArgs);
    emit FinalityConfigUpdated(finalityThreshold, fastTransferFeeBps, maxAmountPerRequest);
  }

  /// @notice Sets the fast finality based rate limit configurations for specified remote chains.
  /// @param rateLimitConfigArgs Array of structs containing remote chain selectors and their rate limiter configs.
  function setFastFinalityRateLimitConfig(
    FastFinalityRateLimitConfigArgs[] calldata rateLimitConfigArgs
  ) external virtual onlyOwner {
    _setFastFinalityRateLimitConfig(rateLimitConfigArgs);
  }

  function _setFastFinalityRateLimitConfig(
    FastFinalityRateLimitConfigArgs[] calldata rateLimitConfigArgs
  ) internal {
    FastFinalityConfig storage finalityConfig = s_finalityConfig;
    for (uint256 i = 0; i < rateLimitConfigArgs.length; ++i) {
      FastFinalityRateLimitConfigArgs calldata configArgs = rateLimitConfigArgs[i];
      uint64 remoteChainSelector = configArgs.remoteChainSelector;
      if (!isSupportedChain(remoteChainSelector)) revert NonExistentChain(remoteChainSelector);

      RateLimiter._validateTokenBucketConfig(configArgs.outboundRateLimiterConfig);
      RateLimiter.TokenBucket storage outboundBucket = finalityConfig.outboundRateLimiterConfig[remoteChainSelector];
      bool outboundUninitialized = outboundBucket.lastUpdated == 0 && outboundBucket.capacity == 0
        && outboundBucket.rate == 0 && outboundBucket.tokens == 0 && !outboundBucket.isEnabled;
      if (outboundUninitialized && configArgs.outboundRateLimiterConfig.isEnabled) {
        outboundBucket.tokens = configArgs.outboundRateLimiterConfig.capacity;
        outboundBucket.lastUpdated = uint32(block.timestamp);
      }
      outboundBucket._setTokenBucketConfig(configArgs.outboundRateLimiterConfig);

      RateLimiter._validateTokenBucketConfig(configArgs.inboundRateLimiterConfig);
      RateLimiter.TokenBucket storage inboundBucket = finalityConfig.inboundRateLimiterConfig[remoteChainSelector];
      bool inboundUninitialized = inboundBucket.lastUpdated == 0 && inboundBucket.capacity == 0
        && inboundBucket.rate == 0 && inboundBucket.tokens == 0 && !inboundBucket.isEnabled;
      if (inboundUninitialized && configArgs.inboundRateLimiterConfig.isEnabled) {
        inboundBucket.tokens = configArgs.inboundRateLimiterConfig.capacity;
        inboundBucket.lastUpdated = uint32(block.timestamp);
      }
      inboundBucket._setTokenBucketConfig(configArgs.inboundRateLimiterConfig);
    }
  }

  // ================================================================
  // │                          CCV                                 │
  // ================================================================

  /// @notice Updates the CCV configuration for specified remote chains.
  /// If the array includes address(0), it indicates that the default CCV should be used alongside any other specified CCVs.
  function applyCCVConfigUpdates(
    CCVConfigArg[] calldata ccvConfigArgs
  ) external virtual onlyOwner {
    for (uint256 i = 0; i < ccvConfigArgs.length; ++i) {
      uint64 remoteChainSelector = ccvConfigArgs[i].remoteChainSelector;
      address[] calldata outboundCCVs = ccvConfigArgs[i].outboundCCVs;
      address[] calldata inboundCCVs = ccvConfigArgs[i].inboundCCVs;

      // check for duplicates in outbound CCVs.
      _checkNoDuplicateAddresses(outboundCCVs);

      // check for duplicates in inbound CCVs.
      _checkNoDuplicateAddresses(inboundCCVs);

      CCVConfig memory ccvConfig = CCVConfig({outboundCCVs: outboundCCVs, inboundCCVs: inboundCCVs});
      emit CCVConfigUpdated(remoteChainSelector, outboundCCVs, inboundCCVs);
      s_verifierConfig[remoteChainSelector] = ccvConfig;
    }
  }

  /// @notice Returns the set of required CCVs for incoming messages from a source chain.
  /// @param sourceChainSelector The source chain selector for incoming messages.
  /// This implementation assumes the same set of CCVs are used for all transfers on a lane.
  /// Implementers can override this function to define custom logic based on these params.
  /// @return requiredCCVs Set of required CCV addresses.
  function getRequiredInboundCCVs(
    address, // localToken
    uint64 sourceChainSelector,
    uint256, // amount,
    uint16, // finality
    bytes calldata // sourcePoolData
  ) external view virtual returns (address[] memory requiredCCVs) {
    return s_verifierConfig[sourceChainSelector].inboundCCVs;
  }

  /// @notice Returns the set of required CCVs for outgoing messages to a destination chain.
  /// @param destChainSelector The destination chain selector for outgoing messages.
  /// This implementation assumes the same set of CCVs are used for all transfers on a lane.
  /// Implementers can override this function to define custom logic based on these params.
  /// @return requiredCCVs Set of required CCV addresses.
  function getRequiredOutboundCCVs(
    address, // localToken
    uint64 destChainSelector,
    uint256, // amount
    uint16, // finality
    bytes calldata // tokenArgs
  ) external view virtual returns (address[] memory requiredCCVs) {
    return s_verifierConfig[destChainSelector].outboundCCVs;
  }

  /// @notice Checks a CCV address array for duplicate entries.
  /// @param ccvs The array of CCV addresses to check for duplicates.
  function _checkNoDuplicateAddresses(
    address[] calldata ccvs
  ) private pure {
    for (uint256 i = 0; i < ccvs.length; ++i) {
      for (uint256 j = i + 1; j < ccvs.length; ++j) {
        if (ccvs[i] == ccvs[j]) {
          revert DuplicateCCV(ccvs[i]);
        }
      }
    }
  }

  // ================================================================
  // │                          Fee                                 │
  // ================================================================

  /// @notice Updates the token transfer fee configurations for specified destination chains.
  /// @param tokenTransferFeeConfigArgs Array of structs containing destination chain selectors and their fee.
  /// @param destToUseDefaultFeeConfigs Array of destination chain selectors to delete custom fee configs for.
  function applyTokenTransferFeeConfigUpdates(
    TokenTransferFeeConfigArgs[] calldata tokenTransferFeeConfigArgs,
    uint64[] calldata destToUseDefaultFeeConfigs
  ) external virtual onlyOwner {
    for (uint256 i = 0; i < tokenTransferFeeConfigArgs.length; ++i) {
      uint64 destChainSelector = tokenTransferFeeConfigArgs[i].destChainSelector;
      TokenTransferFeeConfig calldata tokenTransferFeeConfig = tokenTransferFeeConfigArgs[i].tokenTransferFeeConfig;
      if (
        tokenTransferFeeConfig.isEnabled
          && tokenTransferFeeConfig.destBytesOverhead < Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES
      ) {
        revert InvalidDestBytesOverhead(tokenTransferFeeConfig.destBytesOverhead);
      }
      s_tokenTransferFeeConfig[destChainSelector] = tokenTransferFeeConfig;
      emit TokenTransferFeeConfigUpdated(destChainSelector, tokenTransferFeeConfig);
    }

    for (uint256 i = 0; i < destToUseDefaultFeeConfigs.length; ++i) {
      uint64 destChainSelector = destToUseDefaultFeeConfigs[i];
      delete s_tokenTransferFeeConfig[destChainSelector];
      emit TokenTransferFeeConfigDeleted(destChainSelector);
    }
  }

  /// @notice Returns the token transfer fee override for a destination chain.
  /// @param destChainSelector The destination chain selector used for lookup.
  /// @return feeConfig The enabled fee configuration for the lane.
  function getTokenTransferFeeConfig(
    address, // localToken
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata, // message
    uint16, // finality
    bytes calldata // tokenArgs
  ) external view virtual returns (TokenTransferFeeConfig memory feeConfig) {
    return s_tokenTransferFeeConfig[destChainSelector];
  }

  /// @notice Withdraws all accumulated pool fees to the specified recipient.
  /// @dev For burn/mint pools, this transfers the entire token balance of the pool contract.
  /// lock/release pools should override this function with their own accounting mechanism.
  /// @param recipient The address to receive the withdrawn fees.
  function withdrawFees(
    address recipient
  ) external virtual onlyOwner {
    uint256 amount = getAccumulatedFees();
    if (amount > 0) {
      getToken().safeTransfer(recipient, amount);
      emit PoolFeeWithdrawn(recipient, amount);
    }
  }

  /// @notice Gets the accumulated pool fees that can be withdrawn.
  /// @dev burn/mint pools should return the contract's token balance since pool fees
  /// are minted directly to the pool contract (e.g., `return getToken().balanceOf(address(this))`).
  /// lock/release pools should implement their own accounting mechanism for pool fees
  /// by adding a storage variable (e.g., `s_accumulatedPoolFees`) since they cannot mint
  /// additional tokens for pool fee rewards.
  /// Note: Fee accounting can be obscured by sending tokens directly to the pool.
  /// This does not introduce security issues but will need to be handled operationally.
  /// @return The amount of accumulated pool fees available for withdrawal.
  function getAccumulatedFees() public view virtual returns (uint256) {
    return getToken().balanceOf(address(this));
  }

  // @notice Applies any applicable fees to the lock or burn amount.
  /// @param lockOrBurnIn The original lock or burn request.
  /// @param finality The finality depth requested by the message. A value of zero
  function _applyFee(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 finality
  ) internal view virtual returns (uint256 destAmount) {
    destAmount = lockOrBurnIn.amount;
    if (finality != 0) {
      // deduct fast transfer fee
      destAmount -= (lockOrBurnIn.amount * s_finalityConfig.fastTransferFeeBps) / BPS_DIVIDER;
    }
    // TODO : normal transfer fee
    return destAmount;
  }

  /// @notice Signals which version of the pool interface is supported.
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(TokenPoolV1, IERC165) returns (bool) {
    return interfaceId == Pool.CCIP_POOL_V2 || interfaceId == type(IPoolV2).interfaceId
      || super.supportsInterface(interfaceId);
  }
}
