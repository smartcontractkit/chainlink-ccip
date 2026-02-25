// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../interfaces/IAdvancedPoolHooks.sol";
import {IPoolV1} from "../interfaces/IPool.sol";
import {IPoolV1V2} from "../interfaces/IPoolV1V2.sol";
import {IPoolV2} from "../interfaces/IPoolV2.sol";
import {IRMN} from "../interfaces/IRMN.sol";
import {IRouter} from "../interfaces/IRouter.sol";

import {FeeTokenHandler} from "../libraries/FeeTokenHandler.sol";
import {Pool} from "../libraries/Pool.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {IERC20Metadata} from "@openzeppelin/contracts@5.3.0/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/utils/SafeERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Base abstract class with common functions for all token pools.
/// A token pool serves as isolated place for holding tokens and token specific logic
/// that may execute as tokens move across the bridge.
/// @dev This pool supports different decimals on different chains but using this feature could impact the total number
/// of tokens in circulation. Since all of the tokens are locked/burned on the source, and a rounded amount is
/// minted/released on the destination, the number of tokens minted/released could be less than the number of tokens
/// burned/locked. This is because the source chain does not know about the destination token decimals. This is not a
/// problem if the decimals are the same on both chains.
///
/// Example:
/// Assume there is a token with 6 decimals on chain A and 3 decimals on chain B.
/// - 1.234567 tokens are burned on chain A.
/// - 1.234    tokens are minted on chain B.
/// When sending the 1.234 tokens back to chain A, you will receive 1.234000 tokens on chain A, effectively losing
/// 0.000567 tokens.
/// In the case of a burnMint pool on chain A, these funds are burned in the pool on chain A.
/// In the case of a lockRelease pool on chain A, these funds accumulate in the pool on chain A.
abstract contract TokenPool is IPoolV1V2, Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.Bytes32Set;
  using EnumerableSet for EnumerableSet.UintSet;
  using RateLimiter for RateLimiter.TokenBucket;
  using SafeERC20 for IERC20;

  error InvalidMinBlockConfirmations(uint16 requested, uint16 minBlockConfirmations);
  error CustomBlockConfirmationsNotEnabled();
  error InvalidTransferFeeBps(uint256 bps);
  error InvalidTokenTransferFeeConfig(uint64 destChainSelector);
  error CallerIsNotARampOnRouter(address caller);
  error ZeroAddressInvalid();
  error NonExistentChain(uint64 remoteChainSelector);
  error ChainNotAllowed(uint64 remoteChainSelector);
  error CursedByRMN();
  error ChainAlreadyExists(uint64 chainSelector);
  error InvalidSourcePoolAddress(bytes sourcePoolAddress);
  error InvalidToken(address token);
  error Unauthorized(address caller);
  error PoolAlreadyAdded(uint64 remoteChainSelector, bytes remotePoolAddress);
  error InvalidRemotePoolForChain(uint64 remoteChainSelector, bytes remotePoolAddress);
  error InvalidRemoteChainDecimals(bytes sourcePoolData);
  error MismatchedArrayLengths();
  error OverflowDetected(uint8 remoteDecimals, uint8 localDecimals, uint256 remoteAmount);
  error InvalidDecimalArgs(uint8 expected, uint8 actual);
  error CallerIsNotOwnerOrFeeAdmin(address caller);

  event LockedOrBurned(uint64 indexed remoteChainSelector, address token, address sender, uint256 amount);
  event ReleasedOrMinted(
    uint64 indexed remoteChainSelector, address token, address sender, address recipient, uint256 amount
  );
  event ChainAdded(
    uint64 remoteChainSelector,
    bytes remoteToken,
    RateLimiter.Config outboundRateLimiterConfig,
    RateLimiter.Config inboundRateLimiterConfig
  );
  event ChainRemoved(uint64 remoteChainSelector);
  event RemotePoolAdded(uint64 indexed remoteChainSelector, bytes remotePoolAddress);
  event RemotePoolRemoved(uint64 indexed remoteChainSelector, bytes remotePoolAddress);
  event DynamicConfigSet(address router, address rateLimitAdmin, address feeAdmin);
  event OutboundRateLimitConsumed(uint64 indexed remoteChainSelector, address token, uint256 amount);
  event InboundRateLimitConsumed(uint64 indexed remoteChainSelector, address token, uint256 amount);
  event TokenTransferFeeConfigUpdated(uint64 indexed destChainSelector, TokenTransferFeeConfig tokenTransferFeeConfig);
  event TokenTransferFeeConfigDeleted(uint64 indexed destChainSelector);
  event CustomBlockConfirmationsOutboundRateLimitConsumed(
    uint64 indexed remoteChainSelector, address token, uint256 amount
  );
  event CustomBlockConfirmationsInboundRateLimitConsumed(
    uint64 indexed remoteChainSelector, address token, uint256 amount
  );
  event RateLimitConfigured(
    uint64 indexed remoteChainSelector,
    bool customBlockConfirmations,
    RateLimiter.Config outboundRateLimiterConfig,
    RateLimiter.Config inboundRateLimiterConfig
  );
  event MinBlockConfirmationsSet(uint16 minBlockConfirmations);
  event AdvancedPoolHooksUpdated(IAdvancedPoolHooks oldHook, IAdvancedPoolHooks newHook);

  struct ChainUpdate {
    uint64 remoteChainSelector; // Remote chain selector.
    bytes[] remotePoolAddresses; // Address of the remote pool, ABI encoded in the case of a remote EVM chain.
    bytes remoteTokenAddress; // Address of the remote token, ABI encoded in the case of a remote EVM chain.
    RateLimiter.Config outboundRateLimiterConfig; // Outbound rate limited config, meaning the rate limits for all of the onRamps for the given chain.
    RateLimiter.Config inboundRateLimiterConfig; // Inbound rate limited config, meaning the rate limits for all of the offRamps for the given chain.
  }

  struct RemoteChainConfig {
    RateLimiter.TokenBucket outboundRateLimiterConfig; // Outbound rate limited config, meaning the rate limits for all of the onRamps for the given chain.
    RateLimiter.TokenBucket inboundRateLimiterConfig; // Inbound rate limited config, meaning the rate limits for all of the offRamps for the given chain.
    bytes remoteTokenAddress; // Address of the remote token, ABI encoded in the case of a remote EVM chain.
    EnumerableSet.Bytes32Set remotePools; // Set of remote pool hashes, ABI encoded in the case of a remote EVM chain.
  }

  struct RateLimitConfigArgs {
    uint64 remoteChainSelector; // Remote chain selector.
    bool customBlockConfirmations; // Whether the rate limit config is for custom block confirmations transfers.
    RateLimiter.Config outboundRateLimiterConfig; // Outbound rate limiter configuration.
    RateLimiter.Config inboundRateLimiterConfig; // Inbound rate limiter configuration.
  }

  /// @dev Struct with args for setting the token transfer fee configurations for a destination chain and a set of tokens.
  struct TokenTransferFeeConfigArgs {
    uint64 destChainSelector; // Destination chain selector.
    TokenTransferFeeConfig tokenTransferFeeConfig; // Token transfer fee configuration.
  }

  /// @notice The division factor for bps. This also represents the maximum bps fee.
  uint256 internal constant BPS_DIVIDER = 10_000;
  /// @dev Constant representing the default finality.
  uint16 internal constant WAIT_FOR_FINALITY = 0;
  /// @dev The bridgeable token that is managed by this pool. Pools could support multiple tokens at the same time if
  /// required, but this implementation only supports one token.
  IERC20 internal immutable i_token;
  /// @dev The number of decimals of the token managed by this pool.
  uint8 internal immutable i_tokenDecimals;
  /// @dev The address of the RMN proxy.
  address internal immutable i_rmnProxy;

  /// @dev The address of the router.
  IRouter internal s_router;
  /// @dev Minimum block confirmations on the source chain, 0 means the default finality.
  uint16 internal s_minBlockConfirmations;
  /// @dev Optional advanced pool hooks contract for additional features like allowlists and CCV management.
  IAdvancedPoolHooks internal s_advancedPoolHooks;
  // Separate buckets provide isolated rate limits for transfers with custom block confirmations, as their risk profiles differ from default transfers.
  mapping(uint64 remoteChainSelector => RateLimiter.TokenBucket tokenBucketOutbound) internal
    s_customBlockConfirmationsOutboundRateLimiterConfig;
  mapping(uint64 remoteChainSelector => RateLimiter.TokenBucket tokenBucketInbound) internal
    s_customBlockConfirmationsInboundRateLimiterConfig;
  /// @dev A set of allowed chain selectors. We want the allowlist to be enumerable to
  /// be able to quickly determine (without parsing logs) who can access the pool.
  /// @dev The chain selectors are in uint256 format because of the EnumerableSet implementation.
  EnumerableSet.UintSet internal s_remoteChainSelectors;
  mapping(uint64 remoteChainSelector => RemoteChainConfig) internal s_remoteChainConfigs;
  /// @notice A mapping of hashed pool addresses to their unhashed form. This is used to be able to find the actually
  /// configured pools and not just their hashed versions.
  mapping(bytes32 poolAddressHash => bytes poolAddress) internal s_remotePoolAddresses;
  /// @notice The address of the rate limiter admin.
  /// @dev Can be address(0) if none is configured.
  address internal s_rateLimitAdmin;
  /// @dev Optional token-transfer fee overrides keyed by destination chain selector.
  mapping(uint64 destChainSelector => TokenTransferFeeConfig tokenTransferFeeConfig) internal s_tokenTransferFeeConfig;
  /// @notice The address of the fee admin.
  /// @dev Constructor does not set this value so it is opt in only.
  address internal s_feeAdmin;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address advancedPoolHooks,
    address rmnProxy,
    address router
  ) {
    if (address(token) == address(0) || router == address(0) || rmnProxy == address(0)) {
      revert ZeroAddressInvalid();
    }
    i_token = token;
    i_rmnProxy = rmnProxy;

    try IERC20Metadata(address(token)).decimals() returns (uint8 actualTokenDecimals) {
      if (localTokenDecimals != actualTokenDecimals) {
        revert InvalidDecimalArgs(localTokenDecimals, actualTokenDecimals);
      }
    } catch {
      // The decimals function doesn't exist, which is possible since it's optional in the ERC20 spec. We skip the check and
      // assume the supplied token decimals are correct.
    }
    i_tokenDecimals = localTokenDecimals;
    s_advancedPoolHooks = IAdvancedPoolHooks(advancedPoolHooks);

    s_router = IRouter(router);
  }

  /// @inheritdoc IPoolV1
  /// @param token The token address to check.
  function isSupportedToken(
    address token
  ) public view virtual returns (bool) {
    return token == address(i_token);
  }

  /// @notice Gets the IERC20 token that this pool can lock or burn.
  /// @return token The IERC20 token representation.
  function getToken() public view virtual returns (IERC20 token) {
    return i_token;
  }

  /// @notice Get RMN proxy address.
  /// @return rmnProxy Address of RMN proxy.
  function getRmnProxy() public view virtual returns (address rmnProxy) {
    return i_rmnProxy;
  }

  /// @notice Gets the pools dynamic configuration.
  function getDynamicConfig() public view virtual returns (address router, address rateLimitAdmin, address feeAdmin) {
    return (address(s_router), s_rateLimitAdmin, s_feeAdmin);
  }

  /// @notice Gets the minimum block confirmations required for custom finality transfers.
  function getMinBlockConfirmations() public view virtual returns (uint16 minBlockConfirmations) {
    return s_minBlockConfirmations;
  }

  /// @notice Gets the advanced pool hook contract address used by this pool.
  function getAdvancedPoolHooks() public view virtual returns (IAdvancedPoolHooks advancedPoolHook) {
    return s_advancedPoolHooks;
  }

  /// @notice Sets the dynamic configuration for the pool.
  /// @param router The address of the router contract.
  /// @param rateLimitAdmin The address of the rate limiter admin.
  /// @param feeAdmin An additional address that can withdraw fees from this contract.
  /// @dev FeeTokenHandler will revert if feeAdmin is zero when withdrawing fees.
  /// @dev If only the owner can withdraw fees, set feeAdmin to address(0).
  function setDynamicConfig(
    address router,
    address rateLimitAdmin,
    address feeAdmin
  ) public virtual onlyOwner {
    if (router == address(0)) revert ZeroAddressInvalid();
    s_router = IRouter(router);
    s_rateLimitAdmin = rateLimitAdmin;
    s_feeAdmin = feeAdmin;
    emit DynamicConfigSet(router, rateLimitAdmin, feeAdmin);
  }

  /// @notice Sets the minimum block confirmations required for custom finality transfers.
  /// @param minBlockConfirmations The minimum block confirmations required for custom finality transfers.
  function setMinBlockConfirmations(
    uint16 minBlockConfirmations
  ) public virtual onlyOwner {
    // Since 0 means default finality it is a valid value.
    s_minBlockConfirmations = minBlockConfirmations;
    emit MinBlockConfirmationsSet(minBlockConfirmations);
  }

  /// @notice Updates the advanced pool hook.
  /// @param newHook The new advanced pool hooks contract.
  function updateAdvancedPoolHooks(
    IAdvancedPoolHooks newHook
  ) public virtual onlyOwner {
    emit AdvancedPoolHooksUpdated(s_advancedPoolHooks, newHook);
    s_advancedPoolHooks = newHook;
  }

  /// @notice Signals which version of the pool interface is supported.
  /// @param interfaceId The interface identifier, as specified in ERC-165.
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override returns (bool) {
    return interfaceId == Pool.CCIP_POOL_V1 || interfaceId == type(IPoolV2).interfaceId
      || interfaceId == type(IPoolV1).interfaceId || interfaceId == type(IERC165).interfaceId;
  }

  // ================================================================
  // │                        Lock or Burn                          │
  // ================================================================

  /// @inheritdoc IPoolV2
  /// @dev The _validateLockOrBurn check is an essential security check.
  /// @dev The _getFee function deducts the fee from the amount and returns the amount after fee deduction.
  /// @param lockOrBurnIn Encoded data fields for the processing of tokens on the source chain.
  /// @param blockConfirmationsRequested Requested block confirmations.
  /// @param tokenArgs Additional token arguments.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationsRequested,
    bytes calldata tokenArgs
  ) public virtual returns (Pool.LockOrBurnOutV1 memory, uint256 destTokenAmount) {
    uint256 feeAmount = _getFee(lockOrBurnIn, blockConfirmationsRequested);
    _validateLockOrBurn(lockOrBurnIn, blockConfirmationsRequested, tokenArgs, feeAmount);
    destTokenAmount = lockOrBurnIn.amount - feeAmount;
    _lockOrBurn(lockOrBurnIn.remoteChainSelector, destTokenAmount);

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: lockOrBurnIn.localToken,
      sender: msg.sender,
      amount: destTokenAmount
    });

    return (
      Pool.LockOrBurnOutV1({
        destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector), destPoolData: _encodeLocalDecimals()
      }),
      destTokenAmount
    );
  }

  /// @inheritdoc IPoolV1
  /// @dev The _validateLockOrBurn check is an essential security check.
  /// @dev _getFee is not called in this legacy method, so the full amount is locked or burned.
  /// @param lockOrBurnIn Encoded data fields for the processing of tokens on the source chain.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual returns (Pool.LockOrBurnOutV1 memory lockOrBurnOutV1) {
    _validateLockOrBurn(lockOrBurnIn, WAIT_FOR_FINALITY, "", 0); // feeAmount is zero
    _lockOrBurn(lockOrBurnIn.remoteChainSelector, lockOrBurnIn.amount);

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: lockOrBurnIn.localToken,
      sender: msg.sender,
      amount: lockOrBurnIn.amount
    });

    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector), destPoolData: _encodeLocalDecimals()
    });
  }

  /// @notice Contains the specific lock or burn token logic for a pool.
  /// @dev overriding this method allows us to create pools with different lock/burn signatures
  /// without duplicating the underlying logic.
  /// @param remoteChainSelector The selector of the remote chain.
  /// @param amount The amount of tokens to lock or burn.
  function _lockOrBurn(
    uint64 remoteChainSelector,
    uint256 amount
  ) internal virtual {}

  // ================================================================
  // │                      Release or Mint                         │
  // ================================================================

  /// @inheritdoc IPoolV2
  /// @dev The _validateReleaseOrMint check is an essential security check.
  /// @param releaseOrMintIn Encoded data fields for the processing of tokens on the destination chain.
  /// @param blockConfirmationsRequested Requested block confirmations.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint16 blockConfirmationsRequested
  ) public virtual override(IPoolV2) returns (Pool.ReleaseOrMintOutV1 memory) {
    uint256 localAmount = _calculateLocalAmount(
      releaseOrMintIn.sourceDenominatedAmount, _parseRemoteDecimals(releaseOrMintIn.sourcePoolData)
    );

    _validateReleaseOrMint(releaseOrMintIn, localAmount, blockConfirmationsRequested);

    _releaseOrMint(releaseOrMintIn.receiver, localAmount, releaseOrMintIn.remoteChainSelector);

    emit ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: releaseOrMintIn.localToken,
      sender: msg.sender,
      recipient: releaseOrMintIn.receiver,
      amount: localAmount
    });

    return Pool.ReleaseOrMintOutV1({destinationAmount: localAmount});
  }

  /// @inheritdoc IPoolV1
  /// @dev calls IPoolV2.releaseOrMint with default finality.
  /// @param releaseOrMintIn Encoded data fields for the processing of tokens on the destination chain.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    return releaseOrMint(releaseOrMintIn, WAIT_FOR_FINALITY);
  }

  /// @notice Contains the specific release or mint token logic for a pool.
  /// @dev overriding this method allows us to create pools with different release/mint signatures
  /// without duplicating the underlying logic.
  /// @param receiver The address to receive the tokens.
  /// @param amount The amount of tokens to release or mint.
  /// @param remoteChainSelector The selector of the remote chain.
  function _releaseOrMint(
    address receiver,
    uint256 amount,
    uint64 remoteChainSelector
  ) internal virtual {}

  // ================================================================
  // │                         Validation                           │
  // ================================================================

  /// @notice Validates the lock or burn input for correctness on
  /// - token to be locked or burned
  /// - RMN curse status
  /// - if the sender is a valid onRamp
  /// - rate limiting for either default or custom block confirmations transfer messages.
  /// - preflight checks hooks (if enabled)
  /// @param lockOrBurnIn The input to validate.
  /// @param blockConfirmationsRequested The minimum block confirmations requested by the message. A value of zero is used for default finality.
  /// @param tokenArgs Additional token arguments passed in by the sender of the message.
  /// @param feeAmount The fee amount deducted from the transfer amount.
  /// @dev This function should always be called before executing a lock or burn. Not doing so would allow
  /// for various exploits.
  function _validateLockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationsRequested,
    bytes memory tokenArgs,
    uint256 feeAmount
  ) internal virtual {
    if (!isSupportedToken(lockOrBurnIn.localToken)) {
      revert InvalidToken(lockOrBurnIn.localToken);
    }
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(lockOrBurnIn.remoteChainSelector)))) revert CursedByRMN();

    _onlyOnRamp(lockOrBurnIn.remoteChainSelector);

    uint256 amount = lockOrBurnIn.amount - feeAmount;
    // If custom block confirmations are requested, validate against the minimum and apply the custom rate limit.
    if (blockConfirmationsRequested != WAIT_FOR_FINALITY) {
      uint16 minBlockConfirmationsConfigured = s_minBlockConfirmations;
      if (minBlockConfirmationsConfigured == WAIT_FOR_FINALITY) {
        revert CustomBlockConfirmationsNotEnabled();
      }
      if (blockConfirmationsRequested < minBlockConfirmationsConfigured) {
        revert InvalidMinBlockConfirmations(blockConfirmationsRequested, minBlockConfirmationsConfigured);
      }
      _consumeCustomBlockConfirmationsOutboundRateLimit(
        lockOrBurnIn.localToken, lockOrBurnIn.remoteChainSelector, amount
      );
    } else {
      _consumeOutboundRateLimit(lockOrBurnIn.localToken, lockOrBurnIn.remoteChainSelector, amount);
    }

    _preflightCheck(lockOrBurnIn, blockConfirmationsRequested, tokenArgs, amount);
  }

  /// @notice Hook for pre-flight checks on lock or burn.
  /// @dev These hooks are optional but take up a lot of space in the contracts bytecode. To avoid this overhead when
  /// not needed, you can override this function in the derived contract with an empty implementation. This will result
  /// in the compiler removing the function and all related code, saving close to 1KB.
  /// @param lockOrBurnIn The input to validate.
  /// @param blockConfirmationsRequested The minimum block confirmations requested by the message.
  /// @param tokenArgs Additional token arguments passed in by the sender of the message.
  /// @param amountPostFee The amount after token pool bps-based fees have been deducted.
  function _preflightCheck(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationsRequested,
    bytes memory tokenArgs,
    uint256 amountPostFee
  ) internal virtual {
    if (address(s_advancedPoolHooks) != address(0)) {
      s_advancedPoolHooks.preflightCheck(lockOrBurnIn, blockConfirmationsRequested, tokenArgs, amountPostFee);
    }
  }

  /// @notice Validates the release or mint input for correctness on
  /// - token to be released or minted
  /// - RMN curse status
  /// - if the sender is a valid offRamp
  /// - if the source pool is configured for the remote chain
  /// - rate limiting for either default or custom block confirmations transfer messages.
  /// @param releaseOrMintIn The input to validate.
  /// @param localAmount The local amount to be released or minted.
  /// @param blockConfirmationsRequested The minimum block confirmations requested by the message. A value of zero is used for default finality.
  /// @dev This function should always be called before executing a release or mint. Not doing so would allow
  /// for various exploits.
  function _validateReleaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint256 localAmount,
    uint16 blockConfirmationsRequested
  ) internal virtual {
    if (!isSupportedToken(releaseOrMintIn.localToken)) {
      revert InvalidToken(releaseOrMintIn.localToken);
    }
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(releaseOrMintIn.remoteChainSelector)))) revert CursedByRMN();
    _onlyOffRamp(releaseOrMintIn.remoteChainSelector);

    // Validates that the source pool address is configured on this pool.
    if (!isRemotePool(releaseOrMintIn.remoteChainSelector, releaseOrMintIn.sourcePoolAddress)) {
      revert InvalidSourcePoolAddress(releaseOrMintIn.sourcePoolAddress);
    }
    if (blockConfirmationsRequested != WAIT_FOR_FINALITY) {
      _consumeCustomBlockConfirmationsInboundRateLimit(
        releaseOrMintIn.localToken, releaseOrMintIn.remoteChainSelector, localAmount
      );
    } else {
      _consumeInboundRateLimit(releaseOrMintIn.localToken, releaseOrMintIn.remoteChainSelector, localAmount);
    }

    _postflightCheck(releaseOrMintIn, localAmount, blockConfirmationsRequested);
  }

  /// @notice Hook for post-flight checks on release or mint.
  /// @dev These hooks are optional but take up a lot of space in the contracts bytecode. To avoid this overhead when
  /// not needed, you can override this function in the derived contract with an empty implementation. This will result
  /// in the compiler removing the function and all related code, saving close to 1KB.
  /// @param releaseOrMintIn The input to validate.
  /// @param localAmount The local amount to be released or minted.
  /// @param blockConfirmationsRequested The minimum block confirmations requested by the message.
  function _postflightCheck(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint256 localAmount,
    uint16 blockConfirmationsRequested
  ) internal virtual {
    if (address(s_advancedPoolHooks) != address(0)) {
      s_advancedPoolHooks.postflightCheck(releaseOrMintIn, localAmount, blockConfirmationsRequested);
    }
  }

  // ================================================================
  // │                      Token decimals                          │
  // ================================================================

  /// @notice Gets the IERC20 token decimals on the local chain.
  function getTokenDecimals() public view virtual returns (uint8 decimals) {
    return i_tokenDecimals;
  }

  function _encodeLocalDecimals() internal view virtual returns (bytes memory) {
    return abi.encode(i_tokenDecimals);
  }

  function _parseRemoteDecimals(
    bytes memory sourcePoolData
  ) internal view virtual returns (uint8) {
    // Fallback to the local token decimals if the source pool data is empty. This allows for backwards compatibility.
    if (sourcePoolData.length == 0) {
      return i_tokenDecimals;
    }
    if (sourcePoolData.length != 32) {
      revert InvalidRemoteChainDecimals(sourcePoolData);
    }
    uint256 remoteDecimals = abi.decode(sourcePoolData, (uint256));
    if (remoteDecimals > type(uint8).max) {
      revert InvalidRemoteChainDecimals(sourcePoolData);
    }
    return uint8(remoteDecimals);
  }

  /// @notice Calculates the local amount based on the remote amount and decimals.
  /// @param remoteAmount The amount on the remote chain.
  /// @param remoteDecimals The decimals of the token on the remote chain.
  /// @return The local amount.
  /// @dev This function protects against overflows. If there is a transaction that hits the overflow check, it is
  /// probably incorrect as that means the amount cannot be represented on this chain. If the local decimals have been
  /// wrongly configured, the token issuer could redeploy the pool with the correct decimals and manually re-execute the
  /// CCIP tx to fix the issue.
  function _calculateLocalAmount(
    uint256 remoteAmount,
    uint8 remoteDecimals
  ) internal view virtual returns (uint256) {
    if (remoteDecimals == i_tokenDecimals) {
      return remoteAmount;
    }
    if (remoteDecimals > i_tokenDecimals) {
      uint8 decimalsDiff = remoteDecimals - i_tokenDecimals;
      if (decimalsDiff > 77) {
        // This is a safety check to prevent overflow in the next calculation.
        revert OverflowDetected(remoteDecimals, i_tokenDecimals, remoteAmount);
      }
      // Solidity rounds down so there is no risk of minting more tokens than the remote chain sent.
      return remoteAmount / (10 ** decimalsDiff);
    }

    // This is a safety check to prevent overflow in the next calculation.
    // More than 77 would never fit in a uint256 and would cause an overflow. We also check if the resulting amount
    // would overflow.
    uint8 diffDecimals = i_tokenDecimals - remoteDecimals;
    if (diffDecimals > 77 || remoteAmount > type(uint256).max / (10 ** diffDecimals)) {
      revert OverflowDetected(remoteDecimals, i_tokenDecimals, remoteAmount);
    }

    return remoteAmount * (10 ** diffDecimals);
  }

  // ================================================================
  // │                     Chain permissions                        │
  // ================================================================

  /// @notice Gets the pool address on the remote chain.
  /// @param remoteChainSelector Remote chain selector.
  /// @dev To support non-evm chains, this value is encoded into bytes
  function getRemotePools(
    uint64 remoteChainSelector
  ) public view virtual returns (bytes[] memory) {
    bytes32[] memory remotePoolHashes = s_remoteChainConfigs[remoteChainSelector].remotePools.values();

    bytes[] memory remotePools = new bytes[](remotePoolHashes.length);
    for (uint256 i = 0; i < remotePoolHashes.length; ++i) {
      remotePools[i] = s_remotePoolAddresses[remotePoolHashes[i]];
    }

    return remotePools;
  }

  /// @notice Checks if the pool address is configured on the remote chain.
  /// @param remoteChainSelector Remote chain selector.
  /// @param remotePoolAddress The address of the remote pool.
  function isRemotePool(
    uint64 remoteChainSelector,
    bytes memory remotePoolAddress
  ) public view virtual returns (bool) {
    return s_remoteChainConfigs[remoteChainSelector].remotePools.contains(keccak256(remotePoolAddress));
  }

  /// @inheritdoc IPoolV2
  /// @param remoteChainSelector Remote chain selector.
  function getRemoteToken(
    uint64 remoteChainSelector
  ) public view virtual returns (bytes memory) {
    return s_remoteChainConfigs[remoteChainSelector].remoteTokenAddress;
  }

  /// @notice Adds a remote pool for a given chain selector. This could be due to a pool being upgraded on the remote
  /// chain. We don't simply want to replace the old pool as there could still be valid inflight messages from the old
  /// pool. This function allows for multiple pools to be added for a single chain selector.
  /// @param remoteChainSelector The remote chain selector for which the remote pool address is being added.
  /// @param remotePoolAddress The address of the new remote pool.
  function addRemotePool(
    uint64 remoteChainSelector,
    bytes calldata remotePoolAddress
  ) external virtual onlyOwner {
    if (!isSupportedChain(remoteChainSelector)) revert NonExistentChain(remoteChainSelector);

    _setRemotePool(remoteChainSelector, remotePoolAddress);
  }

  /// @notice Removes the remote pool address for a given chain selector.
  /// @dev All inflight txs from the remote pool will be rejected after it is removed. To ensure no loss of funds, there
  /// should be no inflight txs from the given pool.
  /// @param remoteChainSelector The remote chain selector.
  /// @param remotePoolAddress The remote pool address to remove.
  function removeRemotePool(
    uint64 remoteChainSelector,
    bytes calldata remotePoolAddress
  ) external virtual onlyOwner {
    if (!isSupportedChain(remoteChainSelector)) revert NonExistentChain(remoteChainSelector);

    if (!s_remoteChainConfigs[remoteChainSelector].remotePools.remove(keccak256(remotePoolAddress))) {
      revert InvalidRemotePoolForChain(remoteChainSelector, remotePoolAddress);
    }

    emit RemotePoolRemoved(remoteChainSelector, remotePoolAddress);
  }

  /// @inheritdoc IPoolV1
  /// @param remoteChainSelector The remote chain selector to check.
  function isSupportedChain(
    uint64 remoteChainSelector
  ) public view virtual returns (bool) {
    return s_remoteChainSelectors.contains(remoteChainSelector);
  }

  /// @notice Get list of allowed chains
  /// @return list of chains.
  function getSupportedChains() public view virtual returns (uint64[] memory) {
    uint256[] memory uint256ChainSelectors = s_remoteChainSelectors.values();
    uint64[] memory chainSelectors = new uint64[](uint256ChainSelectors.length);
    for (uint256 i = 0; i < uint256ChainSelectors.length; ++i) {
      chainSelectors[i] = uint64(uint256ChainSelectors[i]);
    }

    return chainSelectors;
  }

  /// @notice Sets the permissions for a list of chains selectors. Actual senders for these chains
  /// need to be allowed on the Router to interact with this pool.
  /// @param remoteChainSelectorsToRemove A list of chain selectors to remove.
  /// @param chainsToAdd A list of chains and their new permission status & rate limits. Rate limits
  /// are only used when the chain is being added through `allowed` being true.
  /// @dev Only callable by the owner
  function applyChainUpdates(
    uint64[] calldata remoteChainSelectorsToRemove,
    ChainUpdate[] calldata chainsToAdd
  ) external virtual onlyOwner {
    for (uint256 i = 0; i < remoteChainSelectorsToRemove.length; ++i) {
      uint64 remoteChainSelectorToRemove = remoteChainSelectorsToRemove[i];
      // If the chain doesn't exist, revert.
      if (!s_remoteChainSelectors.remove(remoteChainSelectorToRemove)) {
        revert NonExistentChain(remoteChainSelectorToRemove);
      }

      // Remove all remote pool hashes for the chain.
      bytes32[] memory remotePools = s_remoteChainConfigs[remoteChainSelectorToRemove].remotePools.values();
      for (uint256 j = 0; j < remotePools.length; ++j) {
        s_remoteChainConfigs[remoteChainSelectorToRemove].remotePools.remove(remotePools[j]);
      }

      delete s_remoteChainConfigs[remoteChainSelectorToRemove];

      emit ChainRemoved(remoteChainSelectorToRemove);
    }

    for (uint256 i = 0; i < chainsToAdd.length; ++i) {
      ChainUpdate memory newChain = chainsToAdd[i];
      if (newChain.remoteTokenAddress.length == 0) {
        revert ZeroAddressInvalid();
      }

      // If the chain already exists, revert
      if (!s_remoteChainSelectors.add(newChain.remoteChainSelector)) {
        revert ChainAlreadyExists(newChain.remoteChainSelector);
      }

      RemoteChainConfig storage remoteChainConfig = s_remoteChainConfigs[newChain.remoteChainSelector];
      remoteChainConfig.outboundRateLimiterConfig._setTokenBucketConfig(newChain.outboundRateLimiterConfig);
      remoteChainConfig.inboundRateLimiterConfig._setTokenBucketConfig(newChain.inboundRateLimiterConfig);

      remoteChainConfig.remoteTokenAddress = newChain.remoteTokenAddress;

      for (uint256 j = 0; j < newChain.remotePoolAddresses.length; ++j) {
        _setRemotePool(newChain.remoteChainSelector, newChain.remotePoolAddresses[j]);
      }

      emit ChainAdded(
        newChain.remoteChainSelector,
        newChain.remoteTokenAddress,
        newChain.outboundRateLimiterConfig,
        newChain.inboundRateLimiterConfig
      );
    }
  }

  /// @notice Adds a pool address to the allowed remote token pools for a particular chain.
  /// @param remoteChainSelector The remote chain selector for which the remote pool address is being added.
  /// @param remotePoolAddress The address of the new remote pool.
  function _setRemotePool(
    uint64 remoteChainSelector,
    bytes memory remotePoolAddress
  ) internal virtual {
    if (remotePoolAddress.length == 0) {
      revert ZeroAddressInvalid();
    }

    bytes32 poolHash = keccak256(remotePoolAddress);

    // Check if the pool already exists.
    if (!s_remoteChainConfigs[remoteChainSelector].remotePools.add(poolHash)) {
      revert PoolAlreadyAdded(remoteChainSelector, remotePoolAddress);
    }

    // Add the pool to the mapping to be able to un-hash it later.
    s_remotePoolAddresses[poolHash] = remotePoolAddress;

    emit RemotePoolAdded(remoteChainSelector, remotePoolAddress);
  }

  // ================================================================
  // │                        Rate limiting                         │
  // ================================================================

  /// @dev The inbound rate limits should be slightly higher than the outbound rate limits. This is because many chains
  /// finalize blocks in batches. CCIP also commits messages in batches: the commit plugin bundles multiple messages in
  /// a single merkle root.
  /// Imagine the following scenario.
  /// - Chain A has an inbound and outbound rate limit of 100 tokens capacity and 1 token per second refill rate.
  /// - Chain B has an inbound and outbound rate limit of 100 tokens capacity and 1 token per second refill rate.
  ///
  /// At time 0:
  /// - Chain A sends 100 tokens to Chain B.
  /// At time 5:
  /// - Chain A sends 5 tokens to Chain B.
  /// At time 6:
  /// The epoch that contains blocks [0-5] is finalized.
  /// Both transactions will be included in the same merkle root and become executable at the same time. This means
  /// the token pool on chain B requires a capacity of 105 to successfully execute both messages at the same time.
  /// The exact additional capacity required depends on the refill rate and the size of the source chain epochs and the
  /// CCIP round time. For simplicity, a 5-10% buffer should be sufficient in most cases.

  /// @notice Consumes outbound rate limiting capacity in this pool.
  /// @param remoteChainSelector The remote chain selector.
  /// @param amount The amount of tokens consumed.
  function _consumeOutboundRateLimit(
    address token,
    uint64 remoteChainSelector,
    uint256 amount
  ) internal virtual {
    s_remoteChainConfigs[remoteChainSelector].outboundRateLimiterConfig._consume(amount, token);

    emit OutboundRateLimitConsumed({token: token, remoteChainSelector: remoteChainSelector, amount: amount});
  }

  /// @notice Consumes inbound rate limiting capacity in this pool.
  /// @param remoteChainSelector The remote chain selector.
  /// @param amount The amount of tokens consumed.
  function _consumeInboundRateLimit(
    address token,
    uint64 remoteChainSelector,
    uint256 amount
  ) internal virtual {
    s_remoteChainConfigs[remoteChainSelector].inboundRateLimiterConfig._consume(amount, token);

    emit InboundRateLimitConsumed({token: token, remoteChainSelector: remoteChainSelector, amount: amount});
  }

  /// @notice Consumes custom block confirmations outbound rate limiting capacity in this pool.
  /// @param remoteChainSelector The remote chain selector.
  /// @param amount The amount of tokens consumed.
  function _consumeCustomBlockConfirmationsOutboundRateLimit(
    address token,
    uint64 remoteChainSelector,
    uint256 amount
  ) internal virtual {
    s_customBlockConfirmationsOutboundRateLimiterConfig[remoteChainSelector]._consume(amount, token);

    emit CustomBlockConfirmationsOutboundRateLimitConsumed({
      token: token, remoteChainSelector: remoteChainSelector, amount: amount
    });
  }

  /// @notice Consumes custom block confirmations inbound rate limiting capacity in this pool.
  /// @param remoteChainSelector The remote chain selector.
  /// @param amount The amount of tokens consumed.
  function _consumeCustomBlockConfirmationsInboundRateLimit(
    address token,
    uint64 remoteChainSelector,
    uint256 amount
  ) internal virtual {
    s_customBlockConfirmationsInboundRateLimiterConfig[remoteChainSelector]._consume(amount, token);

    emit CustomBlockConfirmationsInboundRateLimitConsumed({
      token: token, remoteChainSelector: remoteChainSelector, amount: amount
    });
  }

  /// @notice Returns the outbound and inbound rate limiter state for the given remote chain at the time of the call.
  /// @param remoteChainSelector The remote chain selector.
  /// @param customBlockConfirmations Whether to get the custom block confirmations rate limiter state.
  /// @return outboundRateLimiterState The outbound token bucket.
  /// @return inboundRateLimiterState The inbound token bucket.
  function getCurrentRateLimiterState(
    uint64 remoteChainSelector,
    bool customBlockConfirmations
  )
    external
    view
    virtual
    returns (
      RateLimiter.TokenBucket memory outboundRateLimiterState,
      RateLimiter.TokenBucket memory inboundRateLimiterState
    )
  {
    if (customBlockConfirmations) {
      return (
        s_customBlockConfirmationsOutboundRateLimiterConfig[remoteChainSelector]._currentTokenBucketState(),
        s_customBlockConfirmationsInboundRateLimiterConfig[remoteChainSelector]._currentTokenBucketState()
      );
    }
    RemoteChainConfig storage config = s_remoteChainConfigs[remoteChainSelector];
    return (
      config.outboundRateLimiterConfig._currentTokenBucketState(),
      config.inboundRateLimiterConfig._currentTokenBucketState()
    );
  }

  /// @notice Sets the rate limit configurations for specified remote chains.
  /// @param rateLimitConfigArgs Array of structs containing remote chain selectors and their rate limiter configs.
  function setRateLimitConfig(
    RateLimitConfigArgs[] calldata rateLimitConfigArgs
  ) external virtual {
    _onlyOwnerOrRateLimitAdmin();

    for (uint256 i = 0; i < rateLimitConfigArgs.length; ++i) {
      RateLimitConfigArgs calldata configArgs = rateLimitConfigArgs[i];

      uint64 remoteChainSelector = configArgs.remoteChainSelector;
      if (!isSupportedChain(remoteChainSelector)) revert NonExistentChain(remoteChainSelector);

      if (configArgs.customBlockConfirmations) {
        s_customBlockConfirmationsOutboundRateLimiterConfig[remoteChainSelector]._setTokenBucketConfig(
          configArgs.outboundRateLimiterConfig
        );
        s_customBlockConfirmationsInboundRateLimiterConfig[remoteChainSelector]._setTokenBucketConfig(
          configArgs.inboundRateLimiterConfig
        );
      } else {
        s_remoteChainConfigs[remoteChainSelector].outboundRateLimiterConfig
          ._setTokenBucketConfig(configArgs.outboundRateLimiterConfig);
        s_remoteChainConfigs[remoteChainSelector].inboundRateLimiterConfig
          ._setTokenBucketConfig(configArgs.inboundRateLimiterConfig);
      }

      emit RateLimitConfigured(
        remoteChainSelector,
        configArgs.customBlockConfirmations,
        configArgs.outboundRateLimiterConfig,
        configArgs.inboundRateLimiterConfig
      );
    }
  }

  // ================================================================
  // │                           Access                             │
  // ================================================================

  /// @notice Checks whether remote chain selector is configured on this contract, and if the msg.sender
  /// is a permissioned onRamp for the given chain on the Router.
  /// @dev This function is marked virtual as other token pools may inherit from this contract, but do
  /// not receive calls from the ramps directly, instead receiving them from a proxy contract. In that
  /// situation this function must be overridden and the ramp-check removed and replaced with a different
  /// access-control scheme.
  /// @param remoteChainSelector The remote chain selector.
  function _onlyOnRamp(
    uint64 remoteChainSelector
  ) internal view virtual {
    if (!isSupportedChain(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);
    if (!(msg.sender == s_router.getOnRamp(remoteChainSelector))) revert CallerIsNotARampOnRouter(msg.sender);
  }

  /// @notice Checks whether remote chain selector is configured on this contract, and if the msg.sender
  /// is a permissioned offRamp for the given chain on the Router.
  /// @dev This function is marked virtual as other token pools may inherit from this contract, but do
  /// not receive calls from the ramps directly, instead receiving them from a proxy contract. In that
  /// situation this function must be overridden and the ramp-check removed and replaced with a different
  /// access-control scheme.
  /// @param remoteChainSelector The remote chain selector.
  function _onlyOffRamp(
    uint64 remoteChainSelector
  ) internal view virtual {
    if (!isSupportedChain(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);
    if (!s_router.isOffRamp(remoteChainSelector, msg.sender)) revert CallerIsNotARampOnRouter(msg.sender);
  }

  /// @notice Checks whether the msg.sender is either the owner or the rate limit admin.
  function _onlyOwnerOrRateLimitAdmin() internal view virtual {
    if (msg.sender != s_rateLimitAdmin && msg.sender != owner()) {
      revert Unauthorized(msg.sender);
    }
  }

  /// @notice Returns the set of required CCVs for transfers in a specific direction.
  /// @dev This function delegates to AdvancedPoolHooks if configured, otherwise returns an empty array.
  /// @param localToken The address of the local token.
  /// @param remoteChainSelector The remote chain selector for this transfer.
  /// @param amount The amount being transferred.
  /// @param blockConfirmationsRequested Requested block confirmations.
  /// @param extraData Direction-specific payload forwarded by the caller (e.g. token args or source pool data).
  /// @param direction The direction of the transfer (Inbound or Outbound).
  /// @return requiredCCVs Set of required CCV addresses.
  function getRequiredCCVs(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount,
    uint16 blockConfirmationsRequested,
    bytes calldata extraData,
    IPoolV2.MessageDirection direction
  ) external view virtual returns (address[] memory requiredCCVs) {
    if (address(s_advancedPoolHooks) == address(0)) {
      return new address[](0);
    }

    // The source fee amount is not classified as transferred value, meaning we have to subtract it from the amount
    // before passing it into the hook. The inbound amount is already post-fee so we only need to do this for outbound
    // transfers.
    if (direction == IPoolV2.MessageDirection.Outbound) {
      TokenTransferFeeConfig memory feeConfig = s_tokenTransferFeeConfig[remoteChainSelector];
      if (feeConfig.isEnabled) {
        if (blockConfirmationsRequested != WAIT_FOR_FINALITY) {
          amount -= (amount * feeConfig.customBlockConfirmationsTransferFeeBps) / BPS_DIVIDER;
        } else {
          amount -= (amount * feeConfig.defaultBlockConfirmationsTransferFeeBps) / BPS_DIVIDER;
        }
      }
    }

    return s_advancedPoolHooks.getRequiredCCVs(
      localToken, remoteChainSelector, amount, blockConfirmationsRequested, extraData, direction
    );
  }

  // ================================================================
  // │                          Fee                                 │
  // ================================================================

  /// @notice Updates the token transfer fee configurations for specified destination chains.
  /// @param tokenTransferFeeConfigArgs Array of structs containing destination chain selectors and their fee configs.
  /// @param disableTokenTransferFeeConfigs Array of destination chain selectors to disable custom fee configs for.
  function applyTokenTransferFeeConfigUpdates(
    TokenTransferFeeConfigArgs[] calldata tokenTransferFeeConfigArgs,
    uint64[] calldata disableTokenTransferFeeConfigs
  ) external virtual onlyOwner {
    for (uint256 i = 0; i < tokenTransferFeeConfigArgs.length; ++i) {
      uint64 destChainSelector = tokenTransferFeeConfigArgs[i].destChainSelector;
      TokenTransferFeeConfig calldata tokenTransferFeeConfig = tokenTransferFeeConfigArgs[i].tokenTransferFeeConfig;

      // Reject configs with isEnabled: false - use disableTokenTransferFeeConfigs parameter instead.
      if (!tokenTransferFeeConfig.isEnabled) {
        revert InvalidTokenTransferFeeConfig(destChainSelector);
      }

      if (tokenTransferFeeConfig.defaultBlockConfirmationsTransferFeeBps >= BPS_DIVIDER) {
        revert InvalidTransferFeeBps(tokenTransferFeeConfig.defaultBlockConfirmationsTransferFeeBps);
      }
      if (tokenTransferFeeConfig.customBlockConfirmationsTransferFeeBps >= BPS_DIVIDER) {
        revert InvalidTransferFeeBps(tokenTransferFeeConfig.customBlockConfirmationsTransferFeeBps);
      }
      // Gas overhead must be non-zero for proper fee accounting.
      if (tokenTransferFeeConfig.destGasOverhead == 0) {
        revert InvalidTokenTransferFeeConfig(destChainSelector);
      }

      s_tokenTransferFeeConfig[destChainSelector] = tokenTransferFeeConfig;
      emit TokenTransferFeeConfigUpdated(destChainSelector, tokenTransferFeeConfig);
    }

    for (uint256 i = 0; i < disableTokenTransferFeeConfigs.length; ++i) {
      uint64 destChainSelector = disableTokenTransferFeeConfigs[i];
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
    uint16, // blockConfirmationsRequested
    bytes calldata // tokenArgs
  ) external view virtual returns (TokenTransferFeeConfig memory feeConfig) {
    return s_tokenTransferFeeConfig[destChainSelector];
  }

  /// @inheritdoc IPoolV2
  /// @notice Returns the pool fee parameters that will apply to a transfer.
  /// @param destChainSelector The destination lane selector.
  /// @param blockConfirmationsRequested Requested block confirmations.
  function getFee(
    address, // localToken
    uint64 destChainSelector,
    uint256, // amount
    address, // feeToken
    uint16 blockConfirmationsRequested,
    bytes calldata // tokenArgs
  )
    external
    view
    virtual
    returns (uint256 feeUSDCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps, bool isEnabled)
  {
    uint16 minBlockConfirmationsConfigured = s_minBlockConfirmations;
    if (blockConfirmationsRequested != WAIT_FOR_FINALITY && minBlockConfirmationsConfigured == 0) {
      revert CustomBlockConfirmationsNotEnabled();
    }
    TokenTransferFeeConfig memory feeConfig = s_tokenTransferFeeConfig[destChainSelector];

    // If config is disabled, return zeros with isEnabled=false to signal OnRamp to use FeeQuoter defaults.
    if (!feeConfig.isEnabled) {
      return (0, 0, 0, 0, false);
    }

    if (blockConfirmationsRequested != WAIT_FOR_FINALITY) {
      if (blockConfirmationsRequested < minBlockConfirmationsConfigured) {
        revert InvalidMinBlockConfirmations(blockConfirmationsRequested, minBlockConfirmationsConfigured);
      }
      return (
        feeConfig.customBlockConfirmationsFeeUSDCents,
        feeConfig.destGasOverhead,
        feeConfig.destBytesOverhead,
        feeConfig.customBlockConfirmationsTransferFeeBps,
        true
      );
    }
    return (
      feeConfig.defaultBlockConfirmationsFeeUSDCents,
      feeConfig.destGasOverhead,
      feeConfig.destBytesOverhead,
      feeConfig.defaultBlockConfirmationsTransferFeeBps,
      true
    );
  }

  /// @dev Calculates the fee based on the transferred amount, and the configured basis points.
  /// @param lockOrBurnIn The original lock or burn request.
  /// @param blockConfirmationsRequested The minimum block confirmations requested by the message.
  /// A value of zero (WAIT_FOR_FINALITY) applies default finality fees.
  /// Returns the fee amount.
  function _getFee(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationsRequested
  ) internal view virtual returns (uint256) {
    TokenTransferFeeConfig storage feeConfig = s_tokenTransferFeeConfig[lockOrBurnIn.remoteChainSelector];

    // Determine which fee basis points to apply based on finality type.
    if (blockConfirmationsRequested != WAIT_FOR_FINALITY) {
      return (lockOrBurnIn.amount * feeConfig.customBlockConfirmationsTransferFeeBps) / BPS_DIVIDER;
    } else {
      return (lockOrBurnIn.amount * feeConfig.defaultBlockConfirmationsTransferFeeBps) / BPS_DIVIDER;
    }
  }

  /// @notice Withdraws accrued fee token balances to the provided `recipient`.
  /// @dev Only callable by the owner or the fee admin.
  /// @dev FeeTokenHandler will revert if `recipient` is zero address.
  /// @dev Pools accrue fees directly on this contract. Lock/release pools send bridge liquidity to their ERC20 lockbox
  /// during the lock flow, which means any balance left on this contract represents fees that have accrued to the pool.
  /// Because user liquidity never resides on `address(this)` for lock/release pools, transferring the full contract balance is safe
  /// and clears only accrued fees.
  /// @param feeTokens The token addresses to withdraw, including the pool token when applicable.
  /// @param recipient The address to withdraw the fee tokens to.
  function withdrawFeeTokens(
    address[] calldata feeTokens,
    address recipient
  ) external virtual {
    if (msg.sender != owner() && msg.sender != s_feeAdmin) {
      revert CallerIsNotOwnerOrFeeAdmin(msg.sender);
    }
    FeeTokenHandler._withdrawFeeTokens(feeTokens, recipient);
  }
}
