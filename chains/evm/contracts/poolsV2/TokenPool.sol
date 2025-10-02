// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../interfaces/IPoolV2.sol";
import {IRMN} from "../interfaces/IRMN.sol";

import {Client} from "../libraries/Client.sol";
import {Pool} from "../libraries/Pool.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";
import {TokenPool as TokenPoolV1} from "../pools/TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

abstract contract TokenPool is IPoolV2, TokenPoolV1 {
  using SafeERC20 for IERC20;
  using RateLimiter for RateLimiter.TokenBucket;

  error DuplicateCCV(address ccv);
  error InvalidDestBytesOverhead(uint32 destBytesOverhead);
  error InvalidFinality(uint16 requested, uint16 finalityThreshold);
  error AmountExceedsMaxPerRequest(uint256 requested, uint256 maximum);

  event CCVConfigUpdated(uint64 indexed remoteChainSelector, address[] outboundCCVs, address[] inboundCCVs);
  event FinalityConfigUpdated(uint16 finalityConfig, uint16 fastTransferFeeBps, uint256 maxAmountPerRequest);
  event TokenTransferFeeConfigUpdated(uint64 indexed destChainSelector, TokenTransferFeeConfig tokenTransferFeeConfig);
  event TokenTransferFeeConfigDeleted(uint64 indexed destChainSelector);
  /// @notice Emitted when pool fees are withdrawn.
  event PoolFeeWithdrawn(address indexed recipient, uint256 amount);
  event FastTransferOutboundRateLimitConsumed(uint64 indexed remoteChainSelector, address token, uint256 amount);
  event FastTransferInboundRateLimitConsumed(uint64 indexed remoteChainSelector, address token, uint256 amount);

  struct FastFinalityConfig {
    uint16 finalityThreshold; // ──╮ Maximum block depth required for token transfers.
    uint16 fastTransferFeeBps; // ─╯ Fee in basis points for fast transfers [0-10_000].
    uint256 maxAmountPerRequest; // Maximum amount allowed per transfer request.
    mapping(uint64 remoteChainSelector => RateLimiter.TokenBucket tokenBucketOutbound) outboundRateLimiterConfig;
    mapping(uint64 remoteChainSelector => RateLimiter.TokenBucket tokenBucketInbound) inboundRateLimiterConfig;
  }

  struct FastTransferRateLimitConfigArgs {
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

  struct TokenTransferFeeConfig {
    uint32 destGasOverhead; // ──╮ Gas charged to execute the token transfer on the destination chain.
    uint32 destBytesOverhead; // │ Data availability bytes.
    uint32 feeUSDCents; //       │ Fee to charge per token transfer, multiples of 0.01 USD.
    bool isEnabled; // ──────────╯ Whether this token has custom transfer fees.
  }

  /// @dev Struct with args for setting the token transfer fee configurations for a destination chain and a set of tokens.
  struct TokenTransferFeeConfigArgs {
    uint64 destChainSelector; // Destination chain selector.
    TokenTransferFeeConfig tokenTransferFeeConfig; // Token transfer fee configuration.
  }

  /// @notice The division factor for basis points (BPS). This also represents the maximum BPS fee for fast transfer.
  uint256 internal constant BPS_DIVIDER = 10_000;

  FastFinalityConfig internal s_finalityConfig;
  mapping(uint64 remoteChainSelector => CCVConfig ccvConfig) internal s_verifierConfig;
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

  function lockOrBurn(
    Pool.LockOrBurnInV1 memory lockOrBurnIn,
    uint16 finality,
    bytes calldata // tokenArgs
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory, uint256 destTokenAmount) {
    uint256 amountAfterValidation = _validateLockOrBurn(lockOrBurnIn, finality);

    _lockOrBurn(amountAfterValidation);

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: amountAfterValidation
    });

    return (
      Pool.LockOrBurnOutV1({
        destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
        destPoolData: _encodeLocalDecimals()
      }),
      amountAfterValidation
    );
  }

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
  /// @notice Validates the lock or burn request, applies any fast-finality fee, and enforces rate limits.
  /// @dev The validation covers token support, RMN curse status, allowlist membership, onRamp access, and
  /// rate limiting for both standard and fast-transfer lanes.
  /// @param lockOrBurnIn The input to validate. Must reference a supported token, onRamp, and remote chain.
  /// @param finality The finality depth requested by the message. A value of zero uses the standard lane.
  /// @return amountAfterFee The amount that should be locked or burned after fees and validations are applied.
  function _validateLockOrBurn(
    Pool.LockOrBurnInV1 memory lockOrBurnIn,
    uint16 finality
  ) internal returns (uint256 amountAfterFee) {
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

      amount -= (amount * finalityConfig.fastTransferFeeBps) / BPS_DIVIDER;

      finalityConfig.outboundRateLimiterConfig[lockOrBurnIn.remoteChainSelector]._consume(
        amount, lockOrBurnIn.localToken
      );
      emit FastTransferOutboundRateLimitConsumed(lockOrBurnIn.remoteChainSelector, lockOrBurnIn.localToken, amount);
    } else {
      _consumeOutboundRateLimit(lockOrBurnIn.remoteChainSelector, amount);
    }

    return amount;
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

    FastFinalityConfig storage finalityConfig = s_finalityConfig;
    if (finality != 0) {
      finalityConfig.inboundRateLimiterConfig[releaseOrMintIn.remoteChainSelector]._consume(
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
  // │                          Finality                             │
  // ================================================================
  /// @notice Updates the finality configuration for token transfers.
  function applyFinalityConfigUpdates(
    uint16 finalityThreshold,
    uint16 fastTransferFeeBps,
    uint256 maxAmountPerRequest,
    FastTransferRateLimitConfigArgs[] calldata rateLimitConfigArgs
  ) external virtual onlyOwner {
    FastFinalityConfig storage finalityConfig = s_finalityConfig;
    finalityConfig.finalityThreshold = finalityThreshold;
    finalityConfig.fastTransferFeeBps = fastTransferFeeBps;
    finalityConfig.maxAmountPerRequest = maxAmountPerRequest;
    _setFastTransferRateLimitConfig(rateLimitConfigArgs);
    emit FinalityConfigUpdated(finalityThreshold, fastTransferFeeBps, maxAmountPerRequest);
  }

  function setFastTransferRateLimitConfig(
    FastTransferRateLimitConfigArgs[] calldata rateLimitConfigArgs
  ) external virtual onlyOwner {
    _setFastTransferRateLimitConfig(rateLimitConfigArgs);
  }

  function _setFastTransferRateLimitConfig(
    FastTransferRateLimitConfigArgs[] calldata rateLimitConfigArgs
  ) internal {
    FastFinalityConfig storage finalityConfig = s_finalityConfig;
    for (uint256 i = 0; i < rateLimitConfigArgs.length; ++i) {
      FastTransferRateLimitConfigArgs calldata configArgs = rateLimitConfigArgs[i];
      if (!isSupportedChain(configArgs.remoteChainSelector)) revert NonExistentChain(configArgs.remoteChainSelector);
      RateLimiter._validateTokenBucketConfig(configArgs.outboundRateLimiterConfig);
      _initializeFastBucketIfNeeded(
        finalityConfig.outboundRateLimiterConfig[configArgs.remoteChainSelector], configArgs.outboundRateLimiterConfig
      );
      finalityConfig.outboundRateLimiterConfig[configArgs.remoteChainSelector]._setTokenBucketConfig(
        configArgs.outboundRateLimiterConfig
      );
      RateLimiter._validateTokenBucketConfig(configArgs.inboundRateLimiterConfig);
      _initializeFastBucketIfNeeded(
        finalityConfig.inboundRateLimiterConfig[configArgs.remoteChainSelector], configArgs.inboundRateLimiterConfig
      );
      finalityConfig.inboundRateLimiterConfig[configArgs.remoteChainSelector]._setTokenBucketConfig(
        configArgs.inboundRateLimiterConfig
      );
    }
  }

  function _initializeFastBucketIfNeeded(
    RateLimiter.TokenBucket storage bucket,
    RateLimiter.Config memory config
  ) private {
    if (config.isEnabled && !bucket.isEnabled) {
      bucket.tokens = config.capacity;
      bucket.lastUpdated = uint32(block.timestamp);
    }
  }

  // ================================================================
  // │                          CCV                                  │
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
  // │                          Fee                                  │
  // ================================================================
  function applyTokenTransferFeeConfigUpdates(
    TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs,
    uint64[] calldata destToUseDefaultFeeConfigs
  ) external virtual onlyOwner {
    for (uint256 i = 0; i < tokenTransferFeeConfigArgs.length; ++i) {
      uint64 destChainSelector = tokenTransferFeeConfigArgs[i].destChainSelector;
      TokenTransferFeeConfig memory tokenTransferFeeConfig = tokenTransferFeeConfigArgs[i].tokenTransferFeeConfig;
      s_tokenTransferFeeConfig[destChainSelector] = tokenTransferFeeConfig;
      emit TokenTransferFeeConfigUpdated(destChainSelector, tokenTransferFeeConfig);
    }

    for (uint256 i = 0; i < destToUseDefaultFeeConfigs.length; ++i) {
      uint64 destChainSelector = destToUseDefaultFeeConfigs[i];
      delete s_tokenTransferFeeConfig[destChainSelector];
      emit TokenTransferFeeConfigDeleted(destChainSelector);
    }
  }

  function getTokenTransferFeeConfig(
    address, // localToken
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata, // message
    uint16, // finality
    bytes calldata // tokenArgs
  )
    external
    view
    virtual
    returns (bool isEnabled, uint32 destGasOverhead, uint32 destBytesOverhead, uint32 feeUSDCents)
  {
    TokenTransferFeeConfig memory feeConfig = s_tokenTransferFeeConfig[destChainSelector];
    return (feeConfig.isEnabled, feeConfig.destGasOverhead, feeConfig.destBytesOverhead, feeConfig.feeUSDCents);
  }

  // @inheritdoc IPoolV2
  function withdrawFees(
    address recipient
  ) external virtual onlyOwner {
    uint256 amount = getAccumulatedFees();
    if (amount > 0) {
      getToken().safeTransfer(recipient, amount);
      emit PoolFeeWithdrawn(recipient, amount);
    }
  }

  // @inheritdoc IPoolV2
  // Default implementation returns the entire token balance of the pool.
  // Lock/release pools should override this function with their own accounting mechanism.
  function getAccumulatedFees() public view virtual returns (uint256) {
    return getToken().balanceOf(address(this));
  }

  /// @notice Signals which version of the pool interface is supported.
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(TokenPoolV1, IERC165) returns (bool) {
    return interfaceId == Pool.CCIP_POOL_V2 || interfaceId == type(IPoolV2).interfaceId
      || super.supportsInterface(interfaceId);
  }
}
