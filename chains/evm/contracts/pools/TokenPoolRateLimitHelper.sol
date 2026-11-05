// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Ownable, Ownable2Step} from "@openzeppelin/contracts@5.3.0/access/Ownable2Step.sol";

/// @notice Minimal slice of the CCIP TokenPool v2.0 interface used by the helper.
/// @dev Struct layouts mirror RateLimiter.TokenBucket, RateLimiter.Config and TokenPool.RateLimitConfigArgs exactly,
/// so the ABI encoding (and therefore the function selectors) match the real pool. Declared locally so this contract
/// has no dependency on the BUSL-licensed CCIP sources.
interface ITokenPoolV2RateLimits {
  struct TokenBucket {
    uint128 tokens;
    uint32 lastUpdated;
    bool isEnabled;
    uint128 capacity;
    uint128 rate;
  }

  struct Config {
    bool isEnabled;
    uint128 capacity;
    uint128 rate;
  }

  struct RateLimitConfigArgs {
    uint64 remoteChainSelector;
    bool fastFinality;
    Config outboundRateLimiterConfig;
    Config inboundRateLimiterConfig;
  }

  function getCurrentRateLimiterState(
    uint64 remoteChainSelector,
    bool fastFinality
  ) external view returns (TokenBucket memory outboundRateLimiterState, TokenBucket memory inboundRateLimiterState);

  function setRateLimitConfig(
    RateLimitConfigArgs[] calldata rateLimitConfigArgs
  ) external;

  function getDynamicConfig() external view returns (address router, address rateLimitAdmin, address feeAdmin);
}

/// @title TokenPoolRateLimitHelper
/// @notice Set as the `rateLimitAdmin` of a CCIP TokenPool v2.0 (e.g. LockReleaseTokenPool or
/// SiloedLockReleaseTokenPool) to update ONLY the outbound or ONLY the inbound rate limiter for a lane, leaving the
/// other direction's config (isEnabled / capacity / rate) as it was.
/// @dev The pool's `setRateLimitConfig` always writes both directions, and RateLimiter._setTokenBucketConfig resets
/// `tokens = capacity` on every write. The helper therefore has to re-write the direction you are NOT changing, and
/// does so as follows:
///
/// - Untouched bucket is DISABLED: written back as (false, 0, 0). No effect.
/// - Untouched bucket has rate == 0 (quota mode): written back as (true, capacity = CURRENT TOKENS, 0). Because a
///   zero-rate bucket never refills, the amount that can still be transferred is exactly `tokens`, so this preserves
///   the remaining quota precisely. The trade-off is that the stored `capacity` number for that direction drops to
///   the remaining amount (the original cap stays visible in this contract's events).
/// - Untouched bucket has rate > 0: written back with its existing (capacity, rate). The config is unchanged, but the
///   bucket is REFILLED TO FULL CAPACITY. This is unavoidable through the pool's public API.
///
/// The direction you DO change is always reset to `tokens = newConfig.capacity` by the pool. With rate == 0 that
/// means the new capacity is the new remaining quota, not a new cap on top of past usage.
contract TokenPoolRateLimitHelper is Ownable2Step {
  enum Direction {
    Outbound,
    Inbound
  }

  struct DirectionalUpdate {
    uint64 remoteChainSelector;
    bool fastFinality; // true = the separate fast-finality bucket pair; false = the standard bucket pair.
    Direction direction;
    ITokenPoolV2RateLimits.Config config;
  }

  error ZeroAddressNotAllowed();
  error Unauthorized(address caller);
  error ConfigMismatchAfterUpdate(uint64 remoteChainSelector, bool fastFinality);

  event ManagerSet(address indexed manager, bool allowed);
  event DirectionalRateLimitUpdated(
    uint64 indexed remoteChainSelector,
    bool fastFinality,
    Direction direction,
    ITokenPoolV2RateLimits.Config previousConfig,
    ITokenPoolV2RateLimits.Config newConfig,
    ITokenPoolV2RateLimits.Config untouchedPreviousConfig,
    ITokenPoolV2RateLimits.Config untouchedWrittenBack
  );

  /// @notice The token pool this helper manages. Fixed at deployment so a manager can only ever touch this pool.
  ITokenPoolV2RateLimits public immutable i_pool;

  /// @notice Addresses (besides the owner) allowed to push rate-limit updates through this helper.
  mapping(address manager => bool allowed) internal s_managers;

  modifier onlyOwnerOrManager() {
    if (msg.sender != owner() && !s_managers[msg.sender]) revert Unauthorized(msg.sender);
    _;
  }

  constructor(
    address pool,
    address initialOwner
  ) Ownable(initialOwner) {
    if (pool == address(0)) revert ZeroAddressNotAllowed();
    i_pool = ITokenPoolV2RateLimits(pool);
  }

  // ================================================================
  // │                        Access control                        │
  // ================================================================

  /// @notice Grants or revokes manager rights.
  function setManager(
    address manager,
    bool allowed
  ) external onlyOwner {
    if (manager == address(0)) revert ZeroAddressNotAllowed();
    s_managers[manager] = allowed;
    emit ManagerSet(manager, allowed);
  }

  function isManager(
    address account
  ) external view returns (bool) {
    return s_managers[account];
  }

  // ================================================================
  // │                       Rate limit updates                     │
  // ================================================================

  /// @notice Updates only the outbound limiter for a lane; the inbound config is re-written unchanged.
  function setOutboundRateLimit(
    uint64 remoteChainSelector,
    bool fastFinality,
    ITokenPoolV2RateLimits.Config calldata config
  ) external onlyOwnerOrManager {
    _applyDirectional(remoteChainSelector, fastFinality, Direction.Outbound, config);
  }

  /// @notice Updates only the inbound limiter for a lane; the outbound config is re-written unchanged.
  function setInboundRateLimit(
    uint64 remoteChainSelector,
    bool fastFinality,
    ITokenPoolV2RateLimits.Config calldata config
  ) external onlyOwnerOrManager {
    _applyDirectional(remoteChainSelector, fastFinality, Direction.Inbound, config);
  }

  /// @notice Applies several single-direction updates, in order, atomically.
  /// @dev Each entry reads fresh state from the pool, so two entries for the same lane (one outbound, one inbound)
  /// compose correctly instead of the second overwriting the first.
  function applyDirectionalUpdates(
    DirectionalUpdate[] calldata updates
  ) external onlyOwnerOrManager {
    for (uint256 i = 0; i < updates.length; ++i) {
      DirectionalUpdate calldata u = updates[i];
      _applyDirectional(u.remoteChainSelector, u.fastFinality, u.direction, u.config);
    }
  }

  // ================================================================
  // │                            Views                             │
  // ================================================================

  /// @notice Returns the current (outbound, inbound) configs for a lane, as they would be written back.
  function getRateLimitConfigs(
    uint64 remoteChainSelector,
    bool fastFinality
  ) public view returns (ITokenPoolV2RateLimits.Config memory outbound, ITokenPoolV2RateLimits.Config memory inbound) {
    (ITokenPoolV2RateLimits.TokenBucket memory out, ITokenPoolV2RateLimits.TokenBucket memory inb) =
      i_pool.getCurrentRateLimiterState(remoteChainSelector, fastFinality);
    return (_toConfig(out), _toConfig(inb));
  }

  /// @notice True if the pool currently has this contract set as its rateLimitAdmin.
  function isActiveRateLimitAdmin() external view returns (bool) {
    (, address rateLimitAdmin,) = i_pool.getDynamicConfig();
    return rateLimitAdmin == address(this);
  }

  // ================================================================
  // │                           Internal                           │
  // ================================================================

  function _applyDirectional(
    uint64 remoteChainSelector,
    bool fastFinality,
    Direction direction,
    ITokenPoolV2RateLimits.Config memory newConfig
  ) internal {
    bool isOutbound = direction == Direction.Outbound;

    (ITokenPoolV2RateLimits.TokenBucket memory outBucket, ITokenPoolV2RateLimits.TokenBucket memory inBucket) =
      i_pool.getCurrentRateLimiterState(remoteChainSelector, fastFinality);
    ITokenPoolV2RateLimits.TokenBucket memory targetBucket = isOutbound ? outBucket : inBucket;
    ITokenPoolV2RateLimits.TokenBucket memory untouchedBucket = isOutbound ? inBucket : outBucket;

    ITokenPoolV2RateLimits.Config memory untouchedWriteBack = _writeBackConfig(untouchedBucket);

    ITokenPoolV2RateLimits.RateLimitConfigArgs[] memory args = new ITokenPoolV2RateLimits.RateLimitConfigArgs[](1);
    args[0] = ITokenPoolV2RateLimits.RateLimitConfigArgs({
      remoteChainSelector: remoteChainSelector,
      fastFinality: fastFinality,
      outboundRateLimiterConfig: isOutbound ? newConfig : untouchedWriteBack,
      inboundRateLimiterConfig: isOutbound ? untouchedWriteBack : newConfig
    });

    // The pool validates the config, checks the chain is supported, and checks msg.sender is its rateLimitAdmin.
    i_pool.setRateLimitConfig(args);

    // Defensive post-condition: both directions must now hold exactly what we intended to write.
    (ITokenPoolV2RateLimits.Config memory afterOut, ITokenPoolV2RateLimits.Config memory afterIn) =
      getRateLimitConfigs(remoteChainSelector, fastFinality);
    if (!_eq(afterOut, args[0].outboundRateLimiterConfig) || !_eq(afterIn, args[0].inboundRateLimiterConfig)) {
      revert ConfigMismatchAfterUpdate(remoteChainSelector, fastFinality);
    }

    emit DirectionalRateLimitUpdated(
      remoteChainSelector,
      fastFinality,
      direction,
      _toConfig(targetBucket),
      newConfig,
      _toConfig(untouchedBucket),
      untouchedWriteBack
    );
  }

  /// @dev How the direction that is NOT being changed gets re-written. See the contract-level docs.
  function _writeBackConfig(
    ITokenPoolV2RateLimits.TokenBucket memory bucket
  ) internal pure returns (ITokenPoolV2RateLimits.Config memory) {
    if (!bucket.isEnabled) return ITokenPoolV2RateLimits.Config({isEnabled: false, capacity: 0, rate: 0});
    // Zero-rate bucket: `tokens` is the entire remaining allowance, so write it back as the capacity to preserve it.
    // (rate 0 <= any capacity, including 0, so this always passes the pool's validation.)
    if (bucket.rate == 0) return ITokenPoolV2RateLimits.Config({isEnabled: true, capacity: bucket.tokens, rate: 0});
    return ITokenPoolV2RateLimits.Config({isEnabled: true, capacity: bucket.capacity, rate: bucket.rate});
  }

  /// @dev A disabled bucket is normalised to (false, 0, 0); the pool rejects disabled configs with non-zero values.
  function _toConfig(
    ITokenPoolV2RateLimits.TokenBucket memory bucket
  ) internal pure returns (ITokenPoolV2RateLimits.Config memory) {
    if (!bucket.isEnabled) return ITokenPoolV2RateLimits.Config({isEnabled: false, capacity: 0, rate: 0});
    return ITokenPoolV2RateLimits.Config({isEnabled: true, capacity: bucket.capacity, rate: bucket.rate});
  }

  function _eq(
    ITokenPoolV2RateLimits.Config memory a,
    ITokenPoolV2RateLimits.Config memory b
  ) internal pure returns (bool) {
    return a.isEnabled == b.isEnabled && a.capacity == b.capacity && a.rate == b.rate;
  }
}
