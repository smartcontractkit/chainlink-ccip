// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../interfaces/IPool.sol";
import {IPoolV2} from "../interfaces/IPoolV2.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

import {IRMN} from "../interfaces/IRMN.sol";
import {IRouter} from "../interfaces/IRouter.sol";

import {Client} from "../libraries/Client.sol";
import {Pool} from "../libraries/Pool.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.0.2/utils/structs/EnumerableSet.sol";
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

abstract contract TokenPool is IPoolV2, Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.Bytes32Set;
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.UintSet;
  using RateLimiter for RateLimiter.TokenBucket;
  using SafeERC20 for IERC20;

  error DuplicateCCV(address ccv);
  error InvalidDestBytesOverhead(uint32 destBytesOverhead);
  error InvalidFinality(uint16 requested, uint16 finalityThreshold);
  error AmountExceedsMaxPerRequest(uint256 requested, uint256 maximum);
  error TokenTransferFeeConfigNotEnabled(uint64 destChainSelector);
  error InvalidFastTransferFeeBps();
  error InvalidFinalityConfig();
  error CallerIsNotARampOnRouter(address caller);
  error ZeroAddressInvalid();
  error SenderNotAllowed(address sender);
  error AllowListNotEnabled();
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
  event ChainConfigured(
    uint64 remoteChainSelector,
    RateLimiter.Config outboundRateLimiterConfig,
    RateLimiter.Config inboundRateLimiterConfig
  );
  event ChainRemoved(uint64 remoteChainSelector);
  event RemotePoolAdded(uint64 indexed remoteChainSelector, bytes remotePoolAddress);
  event RemotePoolRemoved(uint64 indexed remoteChainSelector, bytes remotePoolAddress);
  event AllowListAdd(address sender);
  event AllowListRemove(address sender);
  event RouterUpdated(address oldRouter, address newRouter);
  event RateLimitAdminSet(address rateLimitAdmin);
  event OutboundRateLimitConsumed(uint64 indexed remoteChainSelector, address token, uint256 amount);
  event InboundRateLimitConsumed(uint64 indexed remoteChainSelector, address token, uint256 amount);
  event CCVConfigUpdated(uint64 indexed remoteChainSelector, address[] outboundCCVs, address[] inboundCCVs);
  event FinalityConfigUpdated(uint16 finalityConfig, uint16 fastTransferFeeBps, uint256 maxAmountPerRequest);
  event TokenTransferFeeConfigUpdated(uint64 indexed destChainSelector, TokenTransferFeeConfig tokenTransferFeeConfig);
  event TokenTransferFeeConfigDeleted(uint64 indexed destChainSelector);
  /// @notice Emitted when pool fees are withdrawn.
  event PoolFeeWithdrawn(address indexed recipient, uint256 amount);
  event FastTransferOutboundRateLimitConsumed(uint64 indexed remoteChainSelector, address token, uint256 amount);
  event FastTransferInboundRateLimitConsumed(uint64 indexed remoteChainSelector, address token, uint256 amount);

  struct ChainUpdate {
    uint64 remoteChainSelector; // Remote chain selector
    bytes[] remotePoolAddresses; // Address of the remote pool, ABI encoded in the case of a remote EVM chain.
    bytes remoteTokenAddress; // Address of the remote token, ABI encoded in the case of a remote EVM chain.
    RateLimiter.Config outboundRateLimiterConfig; // Outbound rate limited config, meaning the rate limits for all of the onRamps for the given chain
    RateLimiter.Config inboundRateLimiterConfig; // Inbound rate limited config, meaning the rate limits for all of the offRamps for the given chain
  }

  struct RemoteChainConfig {
    RateLimiter.TokenBucket outboundRateLimiterConfig; // Outbound rate limited config, meaning the rate limits for all of the onRamps for the given chain
    RateLimiter.TokenBucket inboundRateLimiterConfig; // Inbound rate limited config, meaning the rate limits for all of the offRamps for the given chain
    bytes remoteTokenAddress; // Address of the remote token, ABI encoded in the case of a remote EVM chain.
    EnumerableSet.Bytes32Set remotePools; // Set of remote pool hashes, ABI encoded in the case of a remote EVM chain.
  }

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
  /// @dev The bridgeable token that is managed by this pool. Pools could support multiple tokens at the same time if
  /// required, but this implementation only supports one token.
  IERC20 internal immutable i_token;
  /// @dev The number of decimals of the token managed by this pool.
  uint8 internal immutable i_tokenDecimals;
  /// @dev The address of the RMN proxy
  address internal immutable i_rmnProxy;
  /// @dev The immutable flag that indicates if the pool is access-controlled.
  bool internal immutable i_allowlistEnabled;
  /// @dev A set of addresses allowed to trigger lockOrBurn as original senders.
  /// Only takes effect if i_allowlistEnabled is true.
  /// This can be used to ensure only token-issuer specified addresses can move tokens.
  EnumerableSet.AddressSet internal s_allowlist;
  /// @dev The address of the router
  IRouter internal s_router;
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
  // Tracks fast-finality parameters and per-lane rate limit buckets for fast transfers.
  FastFinalityConfig internal s_finalityConfig;
  // Stores verifier (CCV) requirements keyed by remote chain selector.
  mapping(uint64 remoteChainSelector => CCVConfig ccvConfig) internal s_verifierConfig;
  // Optional token-transfer fee overrides keyed by destination chain selector.
  mapping(uint64 destChainSelector => TokenTransferFeeConfig tokenTransferFeeConfig) internal s_tokenTransferFeeConfig;

  constructor(IERC20 token, uint8 localTokenDecimals, address[] memory allowlist, address rmnProxy, address router) {
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

    s_router = IRouter(router);

    // Pool can be set as permissioned or permissionless at deployment time only to save hot-path gas.
    i_allowlistEnabled = allowlist.length > 0;
    if (i_allowlistEnabled) {
      _applyAllowListUpdates(new address[](0), allowlist);
    }
  }

  /// @inheritdoc IPoolV1
  function isSupportedToken(
    address token
  ) public view virtual returns (bool) {
    return token == address(i_token);
  }

  /// @notice Gets the IERC20 token that this pool can lock or burn.
  /// @return token The IERC20 token representation.
  function getToken() public view returns (IERC20 token) {
    return i_token;
  }

  /// @notice Get RMN proxy address
  /// @return rmnProxy Address of RMN proxy
  function getRmnProxy() public view returns (address rmnProxy) {
    return i_rmnProxy;
  }

  /// @notice Gets the pool's Router
  /// @return router The pool's Router
  function getRouter() public view virtual returns (address router) {
    return address(s_router);
  }

  /// @notice Sets the pool's Router
  /// @param newRouter The new Router
  function setRouter(
    address newRouter
  ) public onlyOwner {
    if (newRouter == address(0)) revert ZeroAddressInvalid();
    address oldRouter = address(s_router);
    s_router = IRouter(newRouter);

    emit RouterUpdated(oldRouter, newRouter);
  }

  /// @notice Signals which version of the pool interface is supported.
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override returns (bool) {
    return interfaceId == Pool.CCIP_POOL_V2 || interfaceId == Pool.CCIP_POOL_V1
      || interfaceId == type(IPoolV2).interfaceId || interfaceId == type(IPoolV1).interfaceId
      || interfaceId == type(IERC165).interfaceId;
  }

  // ================================================================
  // │                        Lock or Burn                          │
  // ================================================================

  /// @inheritdoc IPoolV2
  /// @dev The _validateLockOrBurn check is an essential security check
  /// @dev The _applyFee function deducts the fee from the amount and returns the amount after fee deduction.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 finality,
    bytes memory // tokenArgs
  ) public virtual returns (Pool.LockOrBurnOutV1 memory, uint256 destTokenAmount) {
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

  /// @notice Burn the token in the pool
  /// @dev The _validateLockOrBurn check is an essential security check
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual returns (Pool.LockOrBurnOutV1 memory lockOrBurnOutV1) {
    (lockOrBurnOutV1,) = lockOrBurn(lockOrBurnIn, uint16(0), "");
    return lockOrBurnOutV1;
  }

  /// @notice Contains the specific lock or burn token logic for a pool.
  /// @dev overriding this method allows us to create pools with different lock/burn signatures
  /// without duplicating the underlying logic.
  function _lockOrBurn(
    uint256 amount
  ) internal virtual {}

  // ================================================================
  // │                      Release or Mint                         │
  // ================================================================

  /// @inheritdoc IPoolV2
  /// @dev The _validateReleaseOrMint check is an essential security check
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

  /// @notice Mint tokens from the pool to the recipient
  /// @dev The _validateReleaseOrMint check is an essential security check
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    return releaseOrMint(releaseOrMintIn, 0);
  }

  /// @notice Contains the specific release or mint token logic for a pool.
  /// @dev overriding this method allows us to create pools with different release/mint signatures
  /// without duplicating the underlying logic.
  function _releaseOrMint(address receiver, uint256 amount) internal virtual {}

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
  function _calculateLocalAmount(uint256 remoteAmount, uint8 remoteDecimals) internal view virtual returns (uint256) {
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
  ) public view returns (bytes[] memory) {
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
  function isRemotePool(uint64 remoteChainSelector, bytes memory remotePoolAddress) public view returns (bool) {
    return s_remoteChainConfigs[remoteChainSelector].remotePools.contains(keccak256(remotePoolAddress));
  }

  /// @notice Gets the token address on the remote chain.
  /// @param remoteChainSelector Remote chain selector.
  /// @dev To support non-evm chains, this value is encoded into bytes
  function getRemoteToken(
    uint64 remoteChainSelector
  ) public view returns (bytes memory) {
    return s_remoteChainConfigs[remoteChainSelector].remoteTokenAddress;
  }

  /// @notice Adds a remote pool for a given chain selector. This could be due to a pool being upgraded on the remote
  /// chain. We don't simply want to replace the old pool as there could still be valid inflight messages from the old
  /// pool. This function allows for multiple pools to be added for a single chain selector.
  /// @param remoteChainSelector The remote chain selector for which the remote pool address is being added.
  /// @param remotePoolAddress The address of the new remote pool.
  function addRemotePool(uint64 remoteChainSelector, bytes calldata remotePoolAddress) external onlyOwner {
    if (!isSupportedChain(remoteChainSelector)) revert NonExistentChain(remoteChainSelector);

    _setRemotePool(remoteChainSelector, remotePoolAddress);
  }

  /// @notice Removes the remote pool address for a given chain selector.
  /// @dev All inflight txs from the remote pool will be rejected after it is removed. To ensure no loss of funds, there
  /// should be no inflight txs from the given pool.
  function removeRemotePool(uint64 remoteChainSelector, bytes calldata remotePoolAddress) external onlyOwner {
    if (!isSupportedChain(remoteChainSelector)) revert NonExistentChain(remoteChainSelector);

    if (!s_remoteChainConfigs[remoteChainSelector].remotePools.remove(keccak256(remotePoolAddress))) {
      revert InvalidRemotePoolForChain(remoteChainSelector, remotePoolAddress);
    }

    emit RemotePoolRemoved(remoteChainSelector, remotePoolAddress);
  }

  /// @inheritdoc IPoolV1
  function isSupportedChain(
    uint64 remoteChainSelector
  ) public view returns (bool) {
    return s_remoteChainSelectors.contains(remoteChainSelector);
  }

  /// @notice Get list of allowed chains
  /// @return list of chains.
  function getSupportedChains() public view returns (uint64[] memory) {
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
      // If the chain doesn't exist, revert
      if (!s_remoteChainSelectors.remove(remoteChainSelectorToRemove)) {
        revert NonExistentChain(remoteChainSelectorToRemove);
      }

      // Remove all remote pool hashes for the chain
      bytes32[] memory remotePools = s_remoteChainConfigs[remoteChainSelectorToRemove].remotePools.values();
      for (uint256 j = 0; j < remotePools.length; ++j) {
        s_remoteChainConfigs[remoteChainSelectorToRemove].remotePools.remove(remotePools[j]);
      }

      delete s_remoteChainConfigs[remoteChainSelectorToRemove];

      emit ChainRemoved(remoteChainSelectorToRemove);
    }

    for (uint256 i = 0; i < chainsToAdd.length; ++i) {
      ChainUpdate memory newChain = chainsToAdd[i];
      RateLimiter._validateTokenBucketConfig(newChain.outboundRateLimiterConfig);
      RateLimiter._validateTokenBucketConfig(newChain.inboundRateLimiterConfig);

      if (newChain.remoteTokenAddress.length == 0) {
        revert ZeroAddressInvalid();
      }

      // If the chain already exists, revert
      if (!s_remoteChainSelectors.add(newChain.remoteChainSelector)) {
        revert ChainAlreadyExists(newChain.remoteChainSelector);
      }

      RemoteChainConfig storage remoteChainConfig = s_remoteChainConfigs[newChain.remoteChainSelector];

      remoteChainConfig.outboundRateLimiterConfig = RateLimiter.TokenBucket({
        rate: newChain.outboundRateLimiterConfig.rate,
        capacity: newChain.outboundRateLimiterConfig.capacity,
        tokens: newChain.outboundRateLimiterConfig.capacity,
        lastUpdated: uint32(block.timestamp),
        isEnabled: newChain.outboundRateLimiterConfig.isEnabled
      });
      remoteChainConfig.inboundRateLimiterConfig = RateLimiter.TokenBucket({
        rate: newChain.inboundRateLimiterConfig.rate,
        capacity: newChain.inboundRateLimiterConfig.capacity,
        tokens: newChain.inboundRateLimiterConfig.capacity,
        lastUpdated: uint32(block.timestamp),
        isEnabled: newChain.inboundRateLimiterConfig.isEnabled
      });
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
  function _setRemotePool(uint64 remoteChainSelector, bytes memory remotePoolAddress) internal {
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

  /// @notice Sets the rate limiter admin address.
  /// @dev Only callable by the owner.
  /// @param rateLimitAdmin The new rate limiter admin address.
  function setRateLimitAdmin(
    address rateLimitAdmin
  ) external onlyOwner {
    s_rateLimitAdmin = rateLimitAdmin;
    emit RateLimitAdminSet(rateLimitAdmin);
  }

  /// @notice Gets the rate limiter admin address.
  function getRateLimitAdmin() external view returns (address) {
    return s_rateLimitAdmin;
  }

  /// @notice Consumes outbound rate limiting capacity in this pool
  function _consumeOutboundRateLimit(uint64 remoteChainSelector, uint256 amount) internal {
    s_remoteChainConfigs[remoteChainSelector].outboundRateLimiterConfig._consume(amount, address(i_token));

    emit OutboundRateLimitConsumed({token: address(i_token), remoteChainSelector: remoteChainSelector, amount: amount});
  }

  /// @notice Consumes inbound rate limiting capacity in this pool
  function _consumeInboundRateLimit(uint64 remoteChainSelector, uint256 amount) internal {
    s_remoteChainConfigs[remoteChainSelector].inboundRateLimiterConfig._consume(amount, address(i_token));

    emit InboundRateLimitConsumed({token: address(i_token), remoteChainSelector: remoteChainSelector, amount: amount});
  }

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @return The token bucket.
  function getCurrentOutboundRateLimiterState(
    uint64 remoteChainSelector
  ) external view returns (RateLimiter.TokenBucket memory) {
    return s_remoteChainConfigs[remoteChainSelector].outboundRateLimiterConfig._currentTokenBucketState();
  }

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @return The token bucket.
  function getCurrentInboundRateLimiterState(
    uint64 remoteChainSelector
  ) external view returns (RateLimiter.TokenBucket memory) {
    return s_remoteChainConfigs[remoteChainSelector].inboundRateLimiterConfig._currentTokenBucketState();
  }

  /// @notice Sets multiple chain rate limiter configs.
  /// @param remoteChainSelectors The remote chain selector for which the rate limits apply.
  /// @param outboundConfigs The new outbound rate limiter config, meaning the onRamp rate limits for the given chain.
  /// @param inboundConfigs The new inbound rate limiter config, meaning the offRamp rate limits for the given chain.
  function setChainRateLimiterConfigs(
    uint64[] calldata remoteChainSelectors,
    RateLimiter.Config[] calldata outboundConfigs,
    RateLimiter.Config[] calldata inboundConfigs
  ) external {
    if (msg.sender != s_rateLimitAdmin && msg.sender != owner()) revert Unauthorized(msg.sender);
    if (remoteChainSelectors.length != outboundConfigs.length || remoteChainSelectors.length != inboundConfigs.length) {
      revert MismatchedArrayLengths();
    }

    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      _setRateLimitConfig(remoteChainSelectors[i], outboundConfigs[i], inboundConfigs[i]);
    }
  }

  /// @notice Sets the chain rate limiter config.
  /// @param remoteChainSelector The remote chain selector for which the rate limits apply.
  /// @param outboundConfig The new outbound rate limiter config, meaning the onRamp rate limits for the given chain.
  /// @param inboundConfig The new inbound rate limiter config, meaning the offRamp rate limits for the given chain.
  function setChainRateLimiterConfig(
    uint64 remoteChainSelector,
    RateLimiter.Config memory outboundConfig,
    RateLimiter.Config memory inboundConfig
  ) external {
    if (msg.sender != s_rateLimitAdmin && msg.sender != owner()) revert Unauthorized(msg.sender);

    _setRateLimitConfig(remoteChainSelector, outboundConfig, inboundConfig);
  }

  function _setRateLimitConfig(
    uint64 remoteChainSelector,
    RateLimiter.Config memory outboundConfig,
    RateLimiter.Config memory inboundConfig
  ) internal {
    if (!isSupportedChain(remoteChainSelector)) revert NonExistentChain(remoteChainSelector);
    RateLimiter._validateTokenBucketConfig(outboundConfig);
    s_remoteChainConfigs[remoteChainSelector].outboundRateLimiterConfig._setTokenBucketConfig(outboundConfig);
    RateLimiter._validateTokenBucketConfig(inboundConfig);
    s_remoteChainConfigs[remoteChainSelector].inboundRateLimiterConfig._setTokenBucketConfig(inboundConfig);
    emit ChainConfigured(remoteChainSelector, outboundConfig, inboundConfig);
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
  function _onlyOffRamp(
    uint64 remoteChainSelector
  ) internal view virtual {
    if (!isSupportedChain(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);
    if (!s_router.isOffRamp(remoteChainSelector, msg.sender)) revert CallerIsNotARampOnRouter(msg.sender);
  }

  // ================================================================
  // │                          Allowlist                           │
  // ================================================================

  function _checkAllowList(
    address sender
  ) internal view {
    if (i_allowlistEnabled) {
      if (!s_allowlist.contains(sender)) {
        revert SenderNotAllowed(sender);
      }
    }
  }

  /// @notice Gets whether the allowlist functionality is enabled.
  /// @return true is enabled, false if not.
  function getAllowListEnabled() external view returns (bool) {
    return i_allowlistEnabled;
  }

  /// @notice Gets the allowed addresses.
  /// @return The allowed addresses.
  function getAllowList() external view returns (address[] memory) {
    return s_allowlist.values();
  }

  /// @notice Apply updates to the allow list.
  /// @param removes The addresses to be removed.
  /// @param adds The addresses to be added.
  function applyAllowListUpdates(address[] calldata removes, address[] calldata adds) external onlyOwner {
    _applyAllowListUpdates(removes, adds);
  }

  /// @notice Internal version of applyAllowListUpdates to allow for reuse in the constructor.
  function _applyAllowListUpdates(address[] memory removes, address[] memory adds) internal {
    if (!i_allowlistEnabled) revert AllowListNotEnabled();

    for (uint256 i = 0; i < removes.length; ++i) {
      address toRemove = removes[i];
      if (s_allowlist.remove(toRemove)) {
        emit AllowListRemove(toRemove);
      }
    }
    for (uint256 i = 0; i < adds.length; ++i) {
      address toAdd = adds[i];
      if (toAdd == address(0)) {
        continue;
      }
      if (s_allowlist.add(toAdd)) {
        emit AllowListAdd(toAdd);
      }
    }
  }

  // ================================================================
  // │                         Finality                             │
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
}
