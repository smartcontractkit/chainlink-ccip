// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {CCVConfigValidation} from "../../../libraries/CCVConfigValidation.sol";
import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../libraries/PolicyConfig.sol";
import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Registry of reusable policies (rate limit, security) and the bindings that attach them to tokens.
/// @dev A token is bound to one named policy of each type through its TokenState, and can additionally carry
/// per-remote-chain inline records that apply only to that (token, chain) pair. Only an authorized pool config manager
/// writes to this store.
///
/// Each token has its own buckets for each chain. Tokens of one policy share the config, not the state.
/// A security record holds the outbound and inbound CCVs and the IMultiTokenPoolHooks address.
contract SecurityPolicyStore is ITypeAndVersion, AuthorizedCallers {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;
  using EnumerableSet for EnumerableSet.UintSet;
  using RateLimiter for RateLimiter.TokenBucket;

  error ZeroAddressInvalid();
  error InvalidDecimalArgs(address localToken, uint8 expected, uint8 actual);
  error InvalidRateLimitPolicyDecimals(bytes32 policyName, uint8 expected, uint8 actual);
  error TokenHasConfig(address localToken);
  error RateLimitRecordNotFound(address localToken, uint64 remoteChainSelector);
  error SecurityRecordNotFound(address localToken, uint64 remoteChainSelector);
  error InvalidInlineRateLimitConfig(address localToken, uint64 remoteChainSelector, RateLimitConfig rateLimitConfig);
  error InvalidRateLimitPolicyConfig(bytes32 policyName, uint64 remoteChainSelector, RateLimitConfig rateLimitConfig);
  error InvalidInlineRateLimitBucketConfig(
    address localToken,
    uint64 remoteChainSelector,
    IPoolV2.MessageDirection direction,
    bool fastFinality,
    RateLimiter.Config config
  );
  error InvalidRateLimitPolicyBucketConfig(
    bytes32 policyName,
    uint64 remoteChainSelector,
    IPoolV2.MessageDirection direction,
    bool fastFinality,
    RateLimiter.Config config
  );

  event InlineRateLimitConfigSet(
    address indexed localToken, uint64 indexed remoteChainSelector, RateLimitConfig rateLimitConfig
  );
  event InlineRateLimitConfigRemoved(address indexed localToken, uint64 indexed remoteChainSelector);
  event InlineRateLimitBucketSet(
    address indexed localToken,
    uint64 indexed remoteChainSelector,
    IPoolV2.MessageDirection direction,
    bool fastFinality,
    RateLimiter.Config config
  );
  event RateLimitPolicyConfigSet(
    bytes32 indexed policyName, uint64 indexed remoteChainSelector, uint8 localDecimals, RateLimitConfig rateLimitConfig
  );
  event RateLimitPolicyConfigRemoved(bytes32 indexed policyName, uint64 indexed remoteChainSelector);
  /// @dev Also stands for the removal of the decimals of the policy.
  event RateLimitPolicyRemoved(bytes32 indexed policyName);
  event RateLimitPolicyBucketSet(
    bytes32 indexed policyName,
    uint64 indexed remoteChainSelector,
    uint8 localDecimals,
    IPoolV2.MessageDirection direction,
    bool fastFinality,
    RateLimiter.Config config
  );
  event TokenProfileSet(address indexed localToken, uint8 localDecimals);
  event TokenRemoved(address indexed localToken);
  event TokenRateLimitPolicySet(
    address indexed localToken, bytes32 indexed oldPolicyName, bytes32 indexed newPolicyName
  );
  event TokenSecurityPolicySet(
    address indexed localToken, bytes32 indexed oldPolicyName, bytes32 indexed newPolicyName
  );
  event OutboundRateLimitConsumed(uint64 indexed remoteChainSelector, address token, uint256 amount);
  event InboundRateLimitConsumed(uint64 indexed remoteChainSelector, address token, uint256 amount);
  event FastFinalityOutboundRateLimitConsumed(uint64 indexed remoteChainSelector, address token, uint256 amount);
  event FastFinalityInboundRateLimitConsumed(uint64 indexed remoteChainSelector, address token, uint256 amount);
  event InlineSecurityConfigSet(
    address indexed localToken, uint64 indexed remoteChainSelector, SecurityConfig securityConfig
  );
  event InlineSecurityConfigRemoved(address indexed localToken, uint64 indexed remoteChainSelector);
  event SecurityPolicyConfigSet(
    bytes32 indexed policyName, uint64 indexed remoteChainSelector, SecurityConfig securityConfig
  );
  event SecurityPolicyConfigRemoved(bytes32 indexed policyName, uint64 indexed remoteChainSelector);
  event SecurityPolicyRemoved(bytes32 indexed policyName);

  /// @notice The state of one token in this store.
  struct TokenState {
    uint8 localDecimals; // Decimals of the local token.
    bytes32 rateLimitPolicyName; // Identifier of the rate limit policy attached to the token, or zero.
    bytes32 securityPolicyName; // Identifier of the security policy attached to the token, or zero.
  }

  /// @notice The values of a rate limit record. Amounts are in local token units.
  /// @dev While an FTF config is disabled, an FTF transfer in its direction uses the default finality bucket of the
  /// same record.
  struct RateLimitConfig {
    RateLimiter.Config outbound; // The outbound limit for default finality, meaning the rate limits for all of the onRamps for the given chain.
    RateLimiter.Config inbound; // The inbound limit for default finality, meaning the rate limits for all of the offRamps for the given chain.
    RateLimiter.Config fastFinalityOutbound; // The outbound limit for FTF.
    RateLimiter.Config fastFinalityInbound; // The inbound limit for FTF.
  }

  /// @notice RateLimitConfig with its policy record key.
  struct RateLimitPolicyConfigArgs {
    bytes32 policyName; // The rate limit policy.
    uint64 remoteChainSelector; // ─╮ The remote chain. Zero names the policy default record.
    uint8 localDecimals; // ────────╯ The decimals of the policy. The first record of the policy fixes them.
    RateLimitConfig rateLimitConfig; // The values of the record.
  }

  /// @notice RateLimitConfig with its token and chain.
  struct InlineRateLimitConfigArgs {
    address localToken; // ─────────╮ The local token address.
    uint64 remoteChainSelector; // ─╯ The remote chain.
    RateLimitConfig rateLimitConfig; // The values of the record.
  }

  /// @notice One bucket of an inline rate limit record to update.
  // solhint-disable-next-line gas-struct-packing
  struct InlineRateLimitBucketArgs {
    address localToken; // ────────────────╮ The local token address.
    uint64 remoteChainSelector; //         │ The remote chain.
    IPoolV2.MessageDirection direction; // │ The direction of the bucket.
    bool fastFinality; // ─────────────────╯ True for the fast finality bucket.
    RateLimiter.Config config; // The limit of the bucket.
  }

  /// @notice One bucket of a rate limit policy record to update.
  // solhint-disable-next-line gas-struct-packing
  struct RateLimitPolicyBucketArgs {
    bytes32 policyName; // The rate limit policy.
    uint64 remoteChainSelector; // ────────╮ The remote chain. Zero names the policy default record.
    IPoolV2.MessageDirection direction; // │ The direction of the bucket.
    bool fastFinality; // ─────────────────╯ True for the fast finality bucket.
    RateLimiter.Config config; // The limit of the bucket.
  }

  /// @notice Security config describing the verification and hook requirements for a lane.
  /// @dev An address(0) entry stands for the default CCVs of the ramps. A list has at most one entry per CCV.
  struct SecurityConfig {
    address[] outboundCCVs; // CCVs required for outgoing messages to the remote chain.
    address[] inboundCCVs; // CCVs required for incoming messages from the remote chain.
    address hook; // The IMultiTokenPoolHooks of the lane. Zero means no hook.
  }

  /// @notice SecurityConfig with its policy record key.
  struct SecurityPolicyConfigArgs {
    bytes32 policyName; // The security policy.
    uint64 remoteChainSelector; // The remote chain. Zero names the policy default record.
    SecurityConfig securityConfig; // The values of the record.
  }

  /// @notice SecurityConfig with its token and chain.
  struct InlineSecurityConfigArgs {
    address localToken; // ─────────╮ The local token address.
    uint64 remoteChainSelector; // ─╯ The remote chain.
    SecurityConfig securityConfig; // The values of the record.
  }

  /// @notice The registered tokens. A token is registered exactly when it is in this set.
  EnumerableSet.AddressSet private s_tokens;

  mapping(address localToken => TokenState tokenState) private s_tokenStates;

  // ================================================================
  // │                       Inline storage                         │
  // ================================================================

  /// @notice The remote chains where a token has an inline rate limit record.
  /// @dev An inline record exists exactly when its chain is in the set of its token. The set is the only record of
  /// which inline records a token has, because the bucket mapping holds a slot for every chain.
  mapping(address localToken => EnumerableSet.UintSet remoteChainSelectors) private s_inlineRateLimitChains;

  /// @notice The live rate limit buckets of each token, per remote chain, finality class and direction.
  /// @dev A bucket holds the live state (tokens, lastUpdated) and the config that applies to it (isEnabled, capacity,
  /// rate). An inline record keeps its config here, so its four buckets are its four configs. A policy record keeps
  /// its config in the policy, and the store copies it into the bucket on use. A bucket is keyed by the remote chain
  /// of the transfer, also while the token uses a policy record.
  /// Separate buckets give fast finality transfers isolated rate limits, as their risk profile differs from default
  /// transfers.
  mapping(
    address localToken
      => mapping(
      uint64 remoteChainSelector
        => mapping(bool fastFinality => mapping(IPoolV2.MessageDirection direction => RateLimiter.TokenBucket bucket))
    )
  ) private s_tokenBuckets;

  /// @notice The remote chains where a token has an inline security record.
  /// @dev An inline record exists exactly when its chain is in the set of its token. The set is the only record of
  /// which inline records a token has, because the config mapping holds a slot for every chain.
  mapping(address localToken => EnumerableSet.UintSet remoteChainSelectors) private s_inlineSecurityChains;

  /// @notice The inline security config of each token, per remote chain.
  /// @dev An inline record exists exactly when its chain is in s_inlineSecurityChains of its token. A security record
  /// has no live state, so its config stands alone here, unlike a rate limit record.
  mapping(address localToken => mapping(uint64 remoteChainSelector => SecurityConfig securityConfig)) private
    s_inlineSecurityConfigs;

  // ================================================================
  // │                       Policy storage                         │
  // ================================================================

  /// @notice The rate limit policies. Each policy is a named set of rate limit records.
  /// @dev A policy name is never zero, so a token without a rate limit policy never resolves to a policy record. A
  /// policy exists from the set of its first record until its removal, and it can exist without records.
  PolicyConfig.Policies private s_rateLimitPolicies;

  /// @notice The decimals of each rate limit policy. Every token of the policy has these decimals.
  /// @dev The first record of the policy fixes them, and they stay fixed until the policy is removed. Amounts are in
  /// local token units, so a shared rate limit means the same amount for every token of the policy.
  mapping(bytes32 policyName => uint8 localDecimals) private s_rateLimitPolicyDecimals;

  /// @notice The rate limit config of each policy record, per remote chain, finality class and direction.
  /// @dev A policy record keeps its config here, unlike an inline record, which keeps it in the buckets of its token.
  /// A transfer copies the config of a policy record into the bucket of its token on use, when the bucket holds
  /// another config, which refills the bucket.
  mapping(
    bytes32 policyName
      => mapping(
      uint64 remoteChainSelector
        => mapping(bool fastFinality => mapping(IPoolV2.MessageDirection direction => RateLimiter.Config config))
    )
  ) private s_policyRateLimitConfigs;

  /// @notice The security policies. Each policy is a named set of security records.
  /// @dev A policy name is never zero, so a token without a security policy never resolves to a policy record. A
  /// policy exists from the set of its first record until its removal, and it can exist without records.
  PolicyConfig.Policies private s_securityPolicies;

  /// @notice The security record of each policy, per remote chain. Zero names the policy default record.
  mapping(bytes32 policyName => mapping(uint64 remoteChainSelector => SecurityConfig securityConfig)) private
    s_policySecurityConfigs;

  constructor(
    address[] memory authorizedCallers
  ) AuthorizedCallers(authorizedCallers) {}

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "SecurityPolicyStore 2.1.0-dev";
  }

  // ================================================================
  // │                     Transfer validation                      │
  // ================================================================

  // The functions below run on every transfer. Each one consumes the rate limit of the transfer
  // and returns the hook of the security record. A transfer reverts if the token has no record for the chain.

  /// @notice Consumes the outbound rate limit of an outbound transfer, and returns the hook of the lane.
  /// @dev Reverts if the token has no security or rate limit record for the chain. Any requested finality other than
  /// WAIT_FOR_FINALITY_FLAG uses the FTF bucket.
  /// @param localToken The local token.
  /// @param remoteChainSelector The destination chain.
  /// @param amount The amount after the token fee, in local decimals.
  /// @param requestedFinalityConfig The requested finality.
  /// @return hook The hook of the security record. Zero means no hook.
  function validateLockOrBurn(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount,
    bytes4 requestedFinalityConfig
  ) external onlyAuthorizedCallers returns (address hook) {
    (PolicyConfig.RecordLevel level, SecurityConfig storage securityConfig) =
      _resolveSecurityConfig(localToken, remoteChainSelector);

    if (level == PolicyConfig.RecordLevel.None) revert SecurityRecordNotFound(localToken, remoteChainSelector);

    _consumeRateLimit(
      localToken,
      remoteChainSelector,
      IPoolV2.MessageDirection.Outbound,
      requestedFinalityConfig != FinalityCodec.WAIT_FOR_FINALITY_FLAG,
      amount
    );
    return securityConfig.hook;
  }

  /// @notice Consumes the inbound rate limit, and returns the hook of the lane.
  /// @dev Reverts if the token has no rate limit or security record for the chain. Any requested finality other than
  /// WAIT_FOR_FINALITY_FLAG uses the FTF bucket.
  /// @param localToken The local token.
  /// @param remoteChainSelector The source chain.
  /// @param localAmount The amount in local decimals.
  /// @param requestedFinalityConfig The requested finality.
  /// @return hook The hook of the security record. Zero means no hook.
  function validateReleaseOrMint(
    address localToken,
    uint64 remoteChainSelector,
    uint256 localAmount,
    bytes4 requestedFinalityConfig
  ) external onlyAuthorizedCallers returns (address hook) {
    _consumeRateLimit(
      localToken,
      remoteChainSelector,
      IPoolV2.MessageDirection.Inbound,
      requestedFinalityConfig != FinalityCodec.WAIT_FOR_FINALITY_FLAG,
      localAmount
    );

    (PolicyConfig.RecordLevel level, SecurityConfig storage securityConfig) =
      _resolveSecurityConfig(localToken, remoteChainSelector);

    if (level == PolicyConfig.RecordLevel.None) revert SecurityRecordNotFound(localToken, remoteChainSelector);

    return securityConfig.hook;
  }

  /// @dev Reverts if the token has no rate limit record for the chain.
  function _consumeRateLimit(
    address localToken,
    uint64 remoteChainSelector,
    IPoolV2.MessageDirection direction,
    bool fastFinality,
    uint256 amount
  ) private {
    (PolicyConfig.RecordLevel level, bytes32 policyName, uint64 recordChainSelector) =
      _resolveRateLimitRecord(localToken, remoteChainSelector);

    if (level == PolicyConfig.RecordLevel.None) revert RateLimitRecordNotFound(localToken, remoteChainSelector);

    // If the FTF config of the record is not enabled, an FTF transfer uses the default finality bucket, the same as
    // TokenPool. A bucket gets the config of its policy record only on use, so this check reads the record.
    if (fastFinality && !_isFastFinalityRateLimitEnabled(localToken, policyName, recordChainSelector, direction)) {
      fastFinality = false;
    }
    _syncBucket(localToken, remoteChainSelector, policyName, recordChainSelector, direction, fastFinality)
      ._consume(amount, localToken);

    if (direction == IPoolV2.MessageDirection.Outbound) {
      if (fastFinality) {
        emit FastFinalityOutboundRateLimitConsumed(remoteChainSelector, localToken, amount);
      } else {
        emit OutboundRateLimitConsumed(remoteChainSelector, localToken, amount);
      }
    } else if (fastFinality) {
      emit FastFinalityInboundRateLimitConsumed(remoteChainSelector, localToken, amount);
    } else {
      emit InboundRateLimitConsumed(remoteChainSelector, localToken, amount);
    }
  }

  function _isFastFinalityRateLimitEnabled(
    address localToken,
    bytes32 policyName,
    uint64 recordChainSelector,
    IPoolV2.MessageDirection direction
  ) private view returns (bool) {
    if (policyName == bytes32(0)) {
      return s_tokenBuckets[localToken][recordChainSelector][true][direction].isEnabled;
    }
    return s_policyRateLimitConfigs[policyName][recordChainSelector][true][direction].isEnabled;
  }

  /// @dev Lazily syncs a token's bucket with its policy record. If the bucket's config differs from the record's, the
  /// record's config is copied in and the bucket is refilled to capacity. This happens on the first use after either:
  /// - the record's config changed, or
  /// - the token moved to a record with a different config.
  /// Otherwise the bucket keeps its current state. Because the sync is lazy, a bucket never holds more tokens than it
  /// would if every config change or record move had refilled it immediately.
  function _syncBucket(
    address localToken,
    uint64 remoteChainSelector,
    bytes32 policyName,
    uint64 recordChainSelector,
    IPoolV2.MessageDirection direction,
    bool fastFinality
  ) private returns (RateLimiter.TokenBucket storage bucket) {
    bucket = s_tokenBuckets[localToken][remoteChainSelector][fastFinality][direction];
    // An inline record keeps its config in the bucket, so there is nothing to sync.
    if (policyName == bytes32(0)) {
      return bucket;
    }

    RateLimiter.Config storage config =
      s_policyRateLimitConfigs[policyName][recordChainSelector][fastFinality][direction];
    if (!_hasConfig(bucket, config)) {
      bucket._setTokenBucketConfig(config);
    }
    return bucket;
  }

  /// @dev A bucket holds the config of its record exactly when the two agree on every field.
  function _hasConfig(
    RateLimiter.TokenBucket storage bucket,
    RateLimiter.Config storage config
  ) private view returns (bool) {
    return bucket.isEnabled == config.isEnabled && bucket.capacity == config.capacity && bucket.rate == config.rate;
  }

  /// @dev Inline records keep their config in the token's buckets, so they return a zero policy name.
  /// recordChainSelector is zero for a policy default record and the remote chain otherwise. Returns (None, 0, 0) if
  /// the token has no record.
  function _resolveRateLimitRecord(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (PolicyConfig.RecordLevel level, bytes32 policyName, uint64 recordChainSelector) {
    // Inline records override the policy.
    if (s_inlineRateLimitChains[localToken].contains(remoteChainSelector)) {
      return (PolicyConfig.RecordLevel.Inline, bytes32(0), remoteChainSelector);
    }

    policyName = s_tokenStates[localToken].rateLimitPolicyName;
    level = PolicyConfig._resolvePolicyRecordLevel(
      s_rateLimitPolicies.policies[policyName].remoteChainSelectors, remoteChainSelector
    );

    if (level == PolicyConfig.RecordLevel.PolicyChain) {
      return (level, policyName, remoteChainSelector);
    }
    if (level == PolicyConfig.RecordLevel.PolicyDefault) {
      return (level, policyName, PolicyConfig.DEFAULT_CHAIN_SELECTOR);
    }
    return (level, bytes32(0), 0);
  }

  /// @dev Resolves in the same order as _resolveRateLimitRecord. Returns an empty config if the token has no record.
  function _resolveSecurityConfig(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (PolicyConfig.RecordLevel level, SecurityConfig storage securityConfig) {
    if (s_inlineSecurityChains[localToken].contains(remoteChainSelector)) {
      return (PolicyConfig.RecordLevel.Inline, s_inlineSecurityConfigs[localToken][remoteChainSelector]);
    }

    bytes32 policyName = s_tokenStates[localToken].securityPolicyName;
    level = PolicyConfig._resolvePolicyRecordLevel(
      s_securityPolicies.policies[policyName].remoteChainSelectors, remoteChainSelector
    );

    if (level == PolicyConfig.RecordLevel.PolicyChain) {
      return (level, s_policySecurityConfigs[policyName][remoteChainSelector]);
    }
    if (level == PolicyConfig.RecordLevel.PolicyDefault) {
      return (level, s_policySecurityConfigs[policyName][PolicyConfig.DEFAULT_CHAIN_SELECTOR]);
    }
    // Without any record this slot is empty, as removing an inline record deletes its values.
    return (level, s_inlineSecurityConfigs[localToken][remoteChainSelector]);
  }

  // ================================================================
  // │                       Resolved views                         │
  // ================================================================

  /// @notice Returns the hook and the CCV lists of each lane, for IPoolConfigManager.getLaneVerificationBatch.
  /// @dev Does not check that the lane exists. The pool config manager zeros a lane that does not exist.
  /// @param lanes The token and remote chain lanes to query.
  /// @return hooks The hook per lane, aligned with lanes by index. Zero if no hook is configured or if the token has no
  /// security record.
  /// @return outboundCCVs The outbound CCVs per lane, aligned by index. Empty if the token has no security record.
  /// @return inboundCCVs The inbound CCVs per lane, aligned by index. Empty if the token has no security record.
  function getLaneVerificationBatch(
    IPoolConfigManager.LaneKey[] calldata lanes
  ) external view returns (address[] memory hooks, address[][] memory outboundCCVs, address[][] memory inboundCCVs) {
    hooks = new address[](lanes.length);
    outboundCCVs = new address[][](lanes.length);
    inboundCCVs = new address[][](lanes.length);

    for (uint256 i = 0; i < lanes.length; ++i) {
      (, SecurityConfig storage securityConfig) =
        _resolveSecurityConfig(lanes[i].localToken, lanes[i].remoteChainSelector);
      hooks[i] = securityConfig.hook;
      outboundCCVs[i] = securityConfig.outboundCCVs;
      inboundCCVs[i] = securityConfig.inboundCCVs;
    }
    return (hooks, outboundCCVs, inboundCCVs);
  }

  /// @notice Returns the four buckets of each lane as they are at this block, for
  /// IPoolConfigManager.getCurrentRateLimiterStateBatch.
  /// @dev Does not check that the lane exists. The pool config manager zeros a lane that does not exist. A bucket that
  /// waits for a copy shows as a full bucket with the config of its policy record.
  /// @param lanes The token and remote chain lanes to query.
  /// @return states The buckets per lane, aligned with lanes by index. All zero if the token has no rate limit record.
  function getCurrentRateLimiterStateBatch(
    IPoolConfigManager.LaneKey[] calldata lanes
  ) external view returns (IPoolConfigManager.LaneRateLimiterState[] memory states) {
    states = new IPoolConfigManager.LaneRateLimiterState[](lanes.length);

    for (uint256 i = 0; i < lanes.length; ++i) {
      states[i] = _getCurrentRateLimiterState(lanes[i].localToken, lanes[i].remoteChainSelector);
    }
    return states;
  }

  /// @notice Checks that the token has a rate limit record and a security record for the remote chain. The lane check
  /// of the pool config manager calls it after FeePolicyStore.checkLaneRecords.
  /// @dev Reverts for the first type that the token misses, in the order rate limit, security.
  /// @param localToken The local token.
  /// @param remoteChainSelector The remote chain.
  function checkLaneRecords(
    address localToken,
    uint64 remoteChainSelector
  ) external view {
    (PolicyConfig.RecordLevel level,,) = _resolveRateLimitRecord(localToken, remoteChainSelector);
    if (level == PolicyConfig.RecordLevel.None) revert RateLimitRecordNotFound(localToken, remoteChainSelector);

    (level,) = _resolveSecurityConfig(localToken, remoteChainSelector);
    if (level == PolicyConfig.RecordLevel.None) revert SecurityRecordNotFound(localToken, remoteChainSelector);
  }

  /// @notice Returns the level of the rate limit and security records that each token would use on its remote chain.
  /// Never reverts.
  /// @dev Each level follows the resolution of a transfer, also where the token has no lane. Chain selector zero gives
  /// None.
  /// @param lanes The token and remote chain lanes to query.
  /// @return rateLimitLevels The level of the rate limit record per lane, aligned with lanes by index.
  /// @return securityLevels The level of the security record per lane, aligned with lanes by index.
  function getLaneRecordLevelBatch(
    IPoolConfigManager.LaneKey[] calldata lanes
  )
    external
    view
    returns (PolicyConfig.RecordLevel[] memory rateLimitLevels, PolicyConfig.RecordLevel[] memory securityLevels)
  {
    rateLimitLevels = new PolicyConfig.RecordLevel[](lanes.length);
    securityLevels = new PolicyConfig.RecordLevel[](lanes.length);

    for (uint256 i = 0; i < lanes.length; ++i) {
      address localToken = lanes[i].localToken;
      uint64 remoteChainSelector = lanes[i].remoteChainSelector;

      (rateLimitLevels[i],,) = _resolveRateLimitRecord(localToken, remoteChainSelector);
      (securityLevels[i],) = _resolveSecurityConfig(localToken, remoteChainSelector);
    }
    return (rateLimitLevels, securityLevels);
  }

  /// @dev All zero for a token without a rate limit record.
  function _getCurrentRateLimiterState(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (IPoolConfigManager.LaneRateLimiterState memory state) {
    (PolicyConfig.RecordLevel level, bytes32 policyName, uint64 recordChainSelector) =
      _resolveRateLimitRecord(localToken, remoteChainSelector);

    if (level == PolicyConfig.RecordLevel.None) {
      return state;
    }

    return IPoolConfigManager.LaneRateLimiterState({
      outbound: _getCurrentTokenBucketState(
        localToken, remoteChainSelector, policyName, recordChainSelector, false, IPoolV2.MessageDirection.Outbound
      ),
      inbound: _getCurrentTokenBucketState(
        localToken, remoteChainSelector, policyName, recordChainSelector, false, IPoolV2.MessageDirection.Inbound
      ),
      fastFinalityOutbound: _getCurrentTokenBucketState(
        localToken, remoteChainSelector, policyName, recordChainSelector, true, IPoolV2.MessageDirection.Outbound
      ),
      fastFinalityInbound: _getCurrentTokenBucketState(
        localToken, remoteChainSelector, policyName, recordChainSelector, true, IPoolV2.MessageDirection.Inbound
      )
    });
  }

  /// @dev The view form of _syncBucket.
  function _getCurrentTokenBucketState(
    address localToken,
    uint64 remoteChainSelector,
    bytes32 policyName,
    uint64 recordChainSelector,
    bool fastFinality,
    IPoolV2.MessageDirection direction
  ) private view returns (IPoolConfigManager.TokenBucket memory) {
    RateLimiter.TokenBucket storage bucket = s_tokenBuckets[localToken][remoteChainSelector][fastFinality][direction];

    if (policyName != bytes32(0)) {
      RateLimiter.Config storage config =
        s_policyRateLimitConfigs[policyName][recordChainSelector][fastFinality][direction];

      if (!_hasConfig(bucket, config)) {
        return IPoolConfigManager.TokenBucket({
          tokens: config.capacity,
          lastUpdated: uint32(block.timestamp),
          isEnabled: config.isEnabled,
          capacity: config.capacity,
          rate: config.rate
        });
      }
    }
    RateLimiter.TokenBucket memory current = bucket._currentTokenBucketState();
    return IPoolConfigManager.TokenBucket({
      tokens: current.tokens,
      lastUpdated: current.lastUpdated,
      isEnabled: current.isEnabled,
      capacity: current.capacity,
      rate: current.rate
    });
  }

  // ================================================================
  // │                        Stored views                          │
  // ================================================================

  /// @notice Returns the state of a token in this store. Never reverts.
  /// @param localToken The local token.
  /// @return isRegistered True if the token is registered in this store.
  /// @return localDecimals The decimals of the token. They stay after a deregistration.
  /// @return rateLimitPolicyName The rate limit policy of the token. Zero if it has none.
  /// @return securityPolicyName The security policy of the token. Zero if it has none.
  function getTokenPolicies(
    address localToken
  )
    external
    view
    returns (bool isRegistered, uint8 localDecimals, bytes32 rateLimitPolicyName, bytes32 securityPolicyName)
  {
    TokenState memory tokenState = s_tokenStates[localToken];
    return (
      s_tokens.contains(localToken),
      tokenState.localDecimals,
      tokenState.rateLimitPolicyName,
      tokenState.securityPolicyName
    );
  }

  /// @notice Returns the rate limit and security policies of each token. Never reverts. The pool config manager
  /// adds them to IPoolConfigManager.getTokenConfigBatch.
  /// @dev A token that is not registered has no policy, because a deregistration needs a token without a policy, and a
  /// token policy set needs a registered token.
  /// @param localTokens The local tokens.
  /// @return rateLimitPolicyNames The rate limit policy per token, aligned with localTokens by index. Zero if it has
  /// none.
  /// @return securityPolicyNames The security policy per token, aligned with localTokens by index. Zero if it has none.
  function getTokenPolicyNamesBatch(
    address[] calldata localTokens
  ) external view returns (bytes32[] memory rateLimitPolicyNames, bytes32[] memory securityPolicyNames) {
    rateLimitPolicyNames = new bytes32[](localTokens.length);
    securityPolicyNames = new bytes32[](localTokens.length);

    for (uint256 i = 0; i < localTokens.length; ++i) {
      TokenState storage tokenState = s_tokenStates[localTokens[i]];
      rateLimitPolicyNames[i] = tokenState.rateLimitPolicyName;
      securityPolicyNames[i] = tokenState.securityPolicyName;
    }
    return (rateLimitPolicyNames, securityPolicyNames);
  }

  /// @notice Returns the inline rate limit records of a token. Never reverts.
  /// @dev A token without an inline record on a chain has no entry. An inline record keeps its config in the buckets of
  /// the token, so this reads the buckets.
  /// @param localToken The local token.
  /// @return inlineConfigs The inline rate limit records of the token.
  function getInlineRateLimitConfigs(
    address localToken
  ) external view returns (InlineRateLimitConfigArgs[] memory inlineConfigs) {
    uint256[] memory remoteChainSelectors = s_inlineRateLimitChains[localToken].values();
    inlineConfigs = new InlineRateLimitConfigArgs[](remoteChainSelectors.length);

    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChainSelectors[i]);
      inlineConfigs[i] = InlineRateLimitConfigArgs({
        localToken: localToken,
        remoteChainSelector: remoteChainSelector,
        rateLimitConfig: RateLimitConfig({
          outbound: _getInlineRateLimitConfig(
            localToken, remoteChainSelector, false, IPoolV2.MessageDirection.Outbound
          ),
          inbound: _getInlineRateLimitConfig(localToken, remoteChainSelector, false, IPoolV2.MessageDirection.Inbound),
          fastFinalityOutbound: _getInlineRateLimitConfig(
            localToken, remoteChainSelector, true, IPoolV2.MessageDirection.Outbound
          ),
          fastFinalityInbound: _getInlineRateLimitConfig(
            localToken, remoteChainSelector, true, IPoolV2.MessageDirection.Inbound
          )
        })
      });
    }
    return inlineConfigs;
  }

  /// @notice Returns the inline security records of a token.
  /// @dev A token without an inline record on a chain has no entry.
  /// @param localToken The local token.
  /// @return inlineConfigs The inline security records of the token.
  function getInlineSecurityConfigs(
    address localToken
  ) external view returns (InlineSecurityConfigArgs[] memory inlineConfigs) {
    uint256[] memory remoteChainSelectors = s_inlineSecurityChains[localToken].values();
    inlineConfigs = new InlineSecurityConfigArgs[](remoteChainSelectors.length);

    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChainSelectors[i]);
      SecurityConfig storage securityConfig = s_inlineSecurityConfigs[localToken][remoteChainSelector];
      inlineConfigs[i] = InlineSecurityConfigArgs({
        localToken: localToken, remoteChainSelector: remoteChainSelector, securityConfig: securityConfig
      });
    }
    return inlineConfigs;
  }

  /// @notice Returns the records of a rate limit policy, in the shape of applyRateLimitPolicyConfigUpdates. Never
  /// reverts.
  /// @param policyName The rate limit policy.
  /// @return tokenCount The number of tokens that use the policy.
  /// @return localDecimals The decimals of the policy. Every token of the policy has them. Zero if the policy does not
  /// exist. getPolicyNames shows whether the policy exists.
  /// @return policyConfigs The records of the policy. Empty if the policy does not exist or has no records.
  function getRateLimitPolicy(
    bytes32 policyName
  ) external view returns (uint32 tokenCount, uint8 localDecimals, RateLimitPolicyConfigArgs[] memory policyConfigs) {
    PolicyConfig.Policy storage policy = s_rateLimitPolicies.policies[policyName];
    uint256[] memory remoteChainSelectors = policy.remoteChainSelectors.values();
    policyConfigs = new RateLimitPolicyConfigArgs[](remoteChainSelectors.length);

    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChainSelectors[i]);
      policyConfigs[i] = RateLimitPolicyConfigArgs({
        policyName: policyName,
        remoteChainSelector: remoteChainSelector,
        localDecimals: s_rateLimitPolicyDecimals[policyName],
        rateLimitConfig: RateLimitConfig({
          outbound: _getRateLimitPolicyConfig(
            policyName, remoteChainSelector, false, IPoolV2.MessageDirection.Outbound
          ),
          inbound: _getRateLimitPolicyConfig(policyName, remoteChainSelector, false, IPoolV2.MessageDirection.Inbound),
          fastFinalityOutbound: _getRateLimitPolicyConfig(
            policyName, remoteChainSelector, true, IPoolV2.MessageDirection.Outbound
          ),
          fastFinalityInbound: _getRateLimitPolicyConfig(
            policyName, remoteChainSelector, true, IPoolV2.MessageDirection.Inbound
          )
        })
      });
    }
    return (policy.tokenCount, s_rateLimitPolicyDecimals[policyName], policyConfigs);
  }

  /// @notice Returns the records of a security policy, in the shape of applySecurityPolicyConfigUpdates. Never reverts.
  /// @param policyName The security policy.
  /// @return tokenCount The number of tokens that use the policy.
  /// @return policyConfigs The records of the policy. Empty if the policy does not exist or has no records.
  function getSecurityPolicy(
    bytes32 policyName
  ) external view returns (uint32 tokenCount, SecurityPolicyConfigArgs[] memory policyConfigs) {
    PolicyConfig.Policy storage policy = s_securityPolicies.policies[policyName];
    uint256[] memory remoteChainSelectors = policy.remoteChainSelectors.values();
    policyConfigs = new SecurityPolicyConfigArgs[](remoteChainSelectors.length);

    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChainSelectors[i]);
      SecurityConfig storage securityConfig = s_policySecurityConfigs[policyName][remoteChainSelector];
      policyConfigs[i] = SecurityPolicyConfigArgs({
        policyName: policyName, remoteChainSelector: remoteChainSelector, securityConfig: securityConfig
      });
    }
    return (policy.tokenCount, policyConfigs);
  }

  /// @notice Returns the name of every rate limit policy and every security policy. Never reverts.
  /// @return rateLimitPolicyNames The names of the rate limit policies.
  /// @return securityPolicyNames The names of the security policies.
  function getPolicyNames()
    external
    view
    returns (bytes32[] memory rateLimitPolicyNames, bytes32[] memory securityPolicyNames)
  {
    return (s_rateLimitPolicies.policyNames.values(), s_securityPolicies.policyNames.values());
  }

  /// @dev The config of an inline rate limit record lives in the bucket of the token, so the getter reads the bucket.
  function _getInlineRateLimitConfig(
    address localToken,
    uint64 remoteChainSelector,
    bool fastFinality,
    IPoolV2.MessageDirection direction
  ) private view returns (RateLimiter.Config memory) {
    RateLimiter.TokenBucket storage bucket = s_tokenBuckets[localToken][remoteChainSelector][fastFinality][direction];
    return RateLimiter.Config({isEnabled: bucket.isEnabled, capacity: bucket.capacity, rate: bucket.rate});
  }

  /// @dev The config of a rate limit policy record lives in the policy, so the getter reads the policy.
  function _getRateLimitPolicyConfig(
    bytes32 policyName,
    uint64 remoteChainSelector,
    bool fastFinality,
    IPoolV2.MessageDirection direction
  ) private view returns (RateLimiter.Config memory) {
    return s_policyRateLimitConfigs[policyName][remoteChainSelector][fastFinality][direction];
  }

  // ================================================================
  // │                           Tokens                             │
  // ================================================================

  /// @notice Deregisters tokens, then sets the profile of tokens. A set of a new token registers it with no policy.
  /// All three stores apply the same call in one transaction.
  /// @dev This store only keeps the token decimals. Decimals cannot change while a token is registered, same as in the
  /// connection store. A token can only be deregistered once it has no policy and no inline records in this store. It
  /// can be registered again later, with different decimals if needed.
  /// @param tokensToRemove The registered tokens to deregister.
  /// @param profilesToSet The tokens to register or to update.
  function applyTokenProfileUpdates(
    address[] calldata tokensToRemove,
    PolicyConfig.TokenProfileArgs[] calldata profilesToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < tokensToRemove.length; ++i) {
      address localToken = tokensToRemove[i];

      if (!s_tokens.contains(localToken)) revert PolicyConfig.LocalTokenNotRegistered(localToken);

      TokenState storage tokenState = s_tokenStates[localToken];
      if (
        tokenState.rateLimitPolicyName != bytes32(0) || tokenState.securityPolicyName != bytes32(0)
          || s_inlineRateLimitChains[localToken].length() != 0 || s_inlineSecurityChains[localToken].length() != 0
      ) {
        revert TokenHasConfig(localToken);
      }

      s_tokens.remove(localToken);
      emit TokenRemoved(localToken);
    }

    for (uint256 i = 0; i < profilesToSet.length; ++i) {
      PolicyConfig.TokenProfileArgs calldata args = profilesToSet[i];
      address localToken = args.localToken;

      if (localToken == address(0)) revert ZeroAddressInvalid();

      TokenState storage tokenState = s_tokenStates[localToken];
      if (s_tokens.contains(localToken) && tokenState.localDecimals != args.localDecimals) {
        revert InvalidDecimalArgs(localToken, tokenState.localDecimals, args.localDecimals);
      }

      // A PoolConfigManager is the only expected caller of this function.
      // Here we expect it to have validated the decimals already.
      s_tokens.add(localToken);
      tokenState.localDecimals = args.localDecimals;
      emit TokenProfileSet(localToken, args.localDecimals);
    }
  }

  /// @notice Sets the rate limit policy of tokens. A token uses the records of its new policy at once.
  /// @dev A set needs a registered token and a policy that exists. The policy must have the decimals of the token,
  /// because a policy rate limit is a raw amount in local token units.
  /// @param tokenPoliciesToSet The rate limit policy of each token. Zero removes the token from its policy.
  function applyTokenRateLimitPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata tokenPoliciesToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < tokenPoliciesToSet.length; ++i) {
      address localToken = tokenPoliciesToSet[i].localToken;
      bytes32 policyName = tokenPoliciesToSet[i].policyName;

      if (!s_tokens.contains(localToken)) revert PolicyConfig.LocalTokenNotRegistered(localToken);

      TokenState storage tokenState = s_tokenStates[localToken];
      // The existence check runs first, because a policy that does not exist has zero decimals.
      if (policyName != bytes32(0)) {
        if (!s_rateLimitPolicies.policyNames.contains(policyName)) revert PolicyConfig.PolicyNotFound(policyName);

        uint8 policyDecimals = s_rateLimitPolicyDecimals[policyName];
        if (policyDecimals != tokenState.localDecimals) {
          revert InvalidRateLimitPolicyDecimals(policyName, policyDecimals, tokenState.localDecimals);
        }
      }

      bytes32 oldPolicyName = tokenState.rateLimitPolicyName;
      PolicyConfig._moveToken(s_rateLimitPolicies, oldPolicyName, policyName);
      tokenState.rateLimitPolicyName = policyName;
      emit TokenRateLimitPolicySet(localToken, oldPolicyName, policyName);
    }
  }

  /// @notice Sets the security policy of tokens. A token uses the records of its new policy at once.
  /// @dev A set needs a registered token and a policy that exists.
  /// @param tokenPoliciesToSet The security policy of each token. Zero removes the token from its policy.
  function applyTokenSecurityPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata tokenPoliciesToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < tokenPoliciesToSet.length; ++i) {
      address localToken = tokenPoliciesToSet[i].localToken;
      bytes32 policyName = tokenPoliciesToSet[i].policyName;

      if (!s_tokens.contains(localToken)) revert PolicyConfig.LocalTokenNotRegistered(localToken);

      TokenState storage tokenState = s_tokenStates[localToken];
      bytes32 oldPolicyName = tokenState.securityPolicyName;

      PolicyConfig._moveToken(s_securityPolicies, oldPolicyName, policyName);
      tokenState.securityPolicyName = policyName;
      emit TokenSecurityPolicySet(localToken, oldPolicyName, policyName);
    }
  }

  // ================================================================
  // │                        Rate limiting                         │
  // ================================================================

  /// @notice Removes inline rate limit records, and creates or replaces them. A set writes all four configs of the
  /// record.
  /// @dev Removals are applied before sets. A set creates the record if it does not exist. Configs are validated with
  /// the same checks as RateLimiter. A bucket is only written, and refilled, when its config changes, so setting the
  /// same config keeps the current bucket state. Removing a record deletes its buckets. To refill a bucket without
  /// changing its config, remove and set the record in the same call.
  /// @param inlineConfigsToRemove The inline records to remove.
  /// @param inlineConfigsToSet The inline records to set.
  function applyInlineRateLimitConfigUpdates(
    IPoolConfigManager.LaneKey[] calldata inlineConfigsToRemove,
    InlineRateLimitConfigArgs[] calldata inlineConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < inlineConfigsToRemove.length; ++i) {
      address localToken = inlineConfigsToRemove[i].localToken;
      uint64 remoteChainSelector = inlineConfigsToRemove[i].remoteChainSelector;

      PolicyConfig._removeInlineRecord(s_inlineRateLimitChains[localToken], localToken, remoteChainSelector);

      mapping(bool fastFinality => mapping(IPoolV2.MessageDirection direction => RateLimiter.TokenBucket bucket))
        storage buckets = s_tokenBuckets[localToken][remoteChainSelector];
      delete buckets[false][IPoolV2.MessageDirection.Outbound];
      delete buckets[false][IPoolV2.MessageDirection.Inbound];
      delete buckets[true][IPoolV2.MessageDirection.Outbound];
      delete buckets[true][IPoolV2.MessageDirection.Inbound];
      emit InlineRateLimitConfigRemoved(localToken, remoteChainSelector);
    }

    for (uint256 i = 0; i < inlineConfigsToSet.length; ++i) {
      InlineRateLimitConfigArgs calldata args = inlineConfigsToSet[i];
      RateLimitConfig calldata rateLimitConfig = args.rateLimitConfig;

      if (
        !_isValidRateLimitConfig(rateLimitConfig.outbound) || !_isValidRateLimitConfig(rateLimitConfig.inbound)
          || !_isValidRateLimitConfig(rateLimitConfig.fastFinalityOutbound)
          || !_isValidRateLimitConfig(rateLimitConfig.fastFinalityInbound)
      ) {
        revert InvalidInlineRateLimitConfig(args.localToken, args.remoteChainSelector, rateLimitConfig);
      }

      PolicyConfig._setInlineRecord(
        s_tokens, s_inlineRateLimitChains[args.localToken], args.localToken, args.remoteChainSelector
      );

      mapping(bool fastFinality => mapping(IPoolV2.MessageDirection direction => RateLimiter.TokenBucket bucket))
        storage buckets = s_tokenBuckets[args.localToken][args.remoteChainSelector];
      _setBucketIfChanged(buckets[false][IPoolV2.MessageDirection.Outbound], rateLimitConfig.outbound);
      _setBucketIfChanged(buckets[false][IPoolV2.MessageDirection.Inbound], rateLimitConfig.inbound);
      _setBucketIfChanged(buckets[true][IPoolV2.MessageDirection.Outbound], rateLimitConfig.fastFinalityOutbound);
      _setBucketIfChanged(buckets[true][IPoolV2.MessageDirection.Inbound], rateLimitConfig.fastFinalityInbound);
      emit InlineRateLimitConfigSet(args.localToken, args.remoteChainSelector, rateLimitConfig);
    }
  }

  /// @notice Updates individual buckets of inline rate limit records, leaving the other buckets unchanged.
  /// @dev The record must already exist, see applyInlineRateLimitConfigUpdates. Configs are validated with the same
  /// checks as RateLimiter. A bucket is only written, and refilled, when its config changes.
  /// @param bucketConfigsToSet The bucket records to set.
  function applyInlineRateLimitBucketUpdates(
    InlineRateLimitBucketArgs[] calldata bucketConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < bucketConfigsToSet.length; ++i) {
      InlineRateLimitBucketArgs calldata args = bucketConfigsToSet[i];

      if (!s_inlineRateLimitChains[args.localToken].contains(args.remoteChainSelector)) {
        revert RateLimitRecordNotFound(args.localToken, args.remoteChainSelector);
      }

      if (!_isValidRateLimitConfig(args.config)) {
        revert InvalidInlineRateLimitBucketConfig(
          args.localToken, args.remoteChainSelector, args.direction, args.fastFinality, args.config
        );
      }

      mapping(bool fastFinality => mapping(IPoolV2.MessageDirection direction => RateLimiter.TokenBucket bucket))
        storage buckets = s_tokenBuckets[args.localToken][args.remoteChainSelector];
      _setBucketIfChanged(buckets[args.fastFinality][args.direction], args.config);
      emit InlineRateLimitBucketSet(
        args.localToken, args.remoteChainSelector, args.direction, args.fastFinality, args.config
      );
    }
  }

  /// @notice Removes and sets rate limit policy records, and removes rate limit policies. A set writes all four
  /// configs of the record.
  /// @dev Record removals are applied first, then policy removals, then record sets. The first record set under a new
  /// name creates the policy and fixes its decimals until the policy is removed. Removing a record keeps the policy. A
  /// policy can only be removed once it has no tokens and no records. Chain selector zero is the policy default record.
  /// Configs are validated with the same checks as RateLimiter. Sets do not write any buckets: a token's bucket picks
  /// up the new config on its next use, and is refilled only if the config changed. After a policy record is removed,
  /// the buckets of its tokens keep their state.
  /// @param policyConfigsToRemove The policy records to remove.
  /// @param policiesToRemove The policies to remove.
  /// @param policyConfigsToSet The policy records to set.
  function applyRateLimitPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata policyConfigsToRemove,
    bytes32[] calldata policiesToRemove,
    RateLimitPolicyConfigArgs[] calldata policyConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < policyConfigsToRemove.length; ++i) {
      bytes32 policyName = policyConfigsToRemove[i].policyName;
      uint64 remoteChainSelector = policyConfigsToRemove[i].remoteChainSelector;

      PolicyConfig._removePolicyRecord(s_rateLimitPolicies, policyName, remoteChainSelector);

      mapping(bool fastFinality => mapping(IPoolV2.MessageDirection direction => RateLimiter.Config config)) storage
        recordConfigs = s_policyRateLimitConfigs[policyName][remoteChainSelector];
      delete recordConfigs[false][IPoolV2.MessageDirection.Outbound];
      delete recordConfigs[false][IPoolV2.MessageDirection.Inbound];
      delete recordConfigs[true][IPoolV2.MessageDirection.Outbound];
      delete recordConfigs[true][IPoolV2.MessageDirection.Inbound];
      emit RateLimitPolicyConfigRemoved(policyName, remoteChainSelector);
    }

    for (uint256 i = 0; i < policiesToRemove.length; ++i) {
      bytes32 policyName = policiesToRemove[i];

      PolicyConfig._removePolicy(s_rateLimitPolicies, policyName);
      delete s_rateLimitPolicyDecimals[policyName];
      emit RateLimitPolicyRemoved(policyName);
    }

    for (uint256 i = 0; i < policyConfigsToSet.length; ++i) {
      RateLimitPolicyConfigArgs calldata args = policyConfigsToSet[i];
      bool isNewPolicy = !s_rateLimitPolicies.policyNames.contains(args.policyName);

      if (!isNewPolicy && s_rateLimitPolicyDecimals[args.policyName] != args.localDecimals) {
        revert InvalidRateLimitPolicyDecimals(
          args.policyName, s_rateLimitPolicyDecimals[args.policyName], args.localDecimals
        );
      }

      RateLimitConfig calldata rateLimitConfig = args.rateLimitConfig;
      if (
        !_isValidRateLimitConfig(rateLimitConfig.outbound) || !_isValidRateLimitConfig(rateLimitConfig.inbound)
          || !_isValidRateLimitConfig(rateLimitConfig.fastFinalityOutbound)
          || !_isValidRateLimitConfig(rateLimitConfig.fastFinalityInbound)
      ) {
        revert InvalidRateLimitPolicyConfig(args.policyName, args.remoteChainSelector, rateLimitConfig);
      }

      PolicyConfig._setPolicyRecord(s_rateLimitPolicies, args.policyName, args.remoteChainSelector);
      if (isNewPolicy) {
        s_rateLimitPolicyDecimals[args.policyName] = args.localDecimals;
      }
      mapping(bool fastFinality => mapping(IPoolV2.MessageDirection direction => RateLimiter.Config config)) storage
        recordConfigs = s_policyRateLimitConfigs[args.policyName][args.remoteChainSelector];
      recordConfigs[false][IPoolV2.MessageDirection.Outbound] = rateLimitConfig.outbound;
      recordConfigs[false][IPoolV2.MessageDirection.Inbound] = rateLimitConfig.inbound;
      recordConfigs[true][IPoolV2.MessageDirection.Outbound] = rateLimitConfig.fastFinalityOutbound;
      recordConfigs[true][IPoolV2.MessageDirection.Inbound] = rateLimitConfig.fastFinalityInbound;
      emit RateLimitPolicyConfigSet(args.policyName, args.remoteChainSelector, args.localDecimals, rateLimitConfig);
    }
  }

  /// @notice Updates individual buckets of rate limit policy records, leaving the other buckets unchanged.
  /// @dev The record must already exist, see applyRateLimitPolicyConfigUpdates. Chain selector zero is the policy
  /// default record. Configs are validated with the same checks as RateLimiter. Sets do not write any buckets: a
  /// token's bucket picks up the new config on its next use, and is refilled only if the config changed.
  /// @param bucketConfigsToSet The bucket records to set.
  function applyRateLimitPolicyBucketUpdates(
    RateLimitPolicyBucketArgs[] calldata bucketConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < bucketConfigsToSet.length; ++i) {
      RateLimitPolicyBucketArgs calldata args = bucketConfigsToSet[i];

      // A record exists exactly when its chain is in the record set of its policy. Zero names the default record.
      if (!s_rateLimitPolicies.policies[args.policyName].remoteChainSelectors.contains(args.remoteChainSelector)) {
        revert PolicyConfig.PolicyRecordNotFound(args.policyName, args.remoteChainSelector);
      }

      if (!_isValidRateLimitConfig(args.config)) {
        revert InvalidRateLimitPolicyBucketConfig(
          args.policyName, args.remoteChainSelector, args.direction, args.fastFinality, args.config
        );
      }

      s_policyRateLimitConfigs[args.policyName][args.remoteChainSelector][args.fastFinality][args.direction] =
      args.config;
      emit RateLimitPolicyBucketSet(
        args.policyName,
        args.remoteChainSelector,
        s_rateLimitPolicyDecimals[args.policyName],
        args.direction,
        args.fastFinality,
        args.config
      );
    }
  }

  /// @dev The checks of RateLimiter._setTokenBucketConfig. A set runs them for every config before it writes. A policy
  /// record stores a config without a bucket, so these are its only checks. The callers revert with the policy or token
  /// and chain, so the error identifies the failing record in a batch.
  function _isValidRateLimitConfig(
    RateLimiter.Config calldata config
  ) private pure returns (bool) {
    if (config.isEnabled) {
      return config.rate <= config.capacity;
    }
    return config.rate == 0 && config.capacity == 0;
  }

  /// @dev Writes a config into a bucket unless the bucket already holds it. A write refills the bucket, so a set of the
  /// same config keeps the state of the bucket, the same as a policy bucket that a transfer syncs.
  function _setBucketIfChanged(
    RateLimiter.TokenBucket storage bucket,
    RateLimiter.Config calldata config
  ) private {
    if (bucket.isEnabled == config.isEnabled && bucket.capacity == config.capacity && bucket.rate == config.rate) {
      return;
    }
    bucket._setTokenBucketConfig(config);
  }

  // ================================================================
  // │                           Security                           │
  // ================================================================

  /// @notice Removes and sets inline security records. A set writes every value of the record, so it replaces the
  /// record.
  /// @dev Removals are applied before sets. A set creates the record if it does not exist. CCV lists must not contain
  /// duplicates.
  /// @param inlineConfigsToRemove The inline records to remove.
  /// @param inlineConfigsToSet The inline records to set.
  function applyInlineSecurityConfigUpdates(
    IPoolConfigManager.LaneKey[] calldata inlineConfigsToRemove,
    InlineSecurityConfigArgs[] calldata inlineConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < inlineConfigsToRemove.length; ++i) {
      address localToken = inlineConfigsToRemove[i].localToken;
      uint64 remoteChainSelector = inlineConfigsToRemove[i].remoteChainSelector;

      PolicyConfig._removeInlineRecord(s_inlineSecurityChains[localToken], localToken, remoteChainSelector);
      delete s_inlineSecurityConfigs[localToken][remoteChainSelector];
      emit InlineSecurityConfigRemoved(localToken, remoteChainSelector);
    }

    for (uint256 i = 0; i < inlineConfigsToSet.length; ++i) {
      InlineSecurityConfigArgs calldata args = inlineConfigsToSet[i];
      SecurityConfig calldata securityConfig = args.securityConfig;

      _validateSecurityConfig(securityConfig.outboundCCVs, securityConfig.inboundCCVs);

      PolicyConfig._setInlineRecord(
        s_tokens, s_inlineSecurityChains[args.localToken], args.localToken, args.remoteChainSelector
      );
      _setSecurityConfig(s_inlineSecurityConfigs[args.localToken][args.remoteChainSelector], securityConfig);
      emit InlineSecurityConfigSet(args.localToken, args.remoteChainSelector, securityConfig);
    }
  }

  /// @notice Removes and sets security policy records, and removes security policies. A set writes every value of the
  /// record, so it replaces the record.
  /// @dev The record removals run first, then the policy removals, then the record sets. The first record set under a
  /// new name creates the policy. A record removal keeps the policy. A policy removal needs a policy without tokens
  /// and without records. Chain selector zero names the policy default record. CCV lists must not contain duplicates.
  /// @param policyConfigsToRemove The policy records to remove.
  /// @param policiesToRemove The policies to remove.
  /// @param policyConfigsToSet The policy records to set.
  function applySecurityPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata policyConfigsToRemove,
    bytes32[] calldata policiesToRemove,
    SecurityPolicyConfigArgs[] calldata policyConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < policyConfigsToRemove.length; ++i) {
      bytes32 policyName = policyConfigsToRemove[i].policyName;
      uint64 remoteChainSelector = policyConfigsToRemove[i].remoteChainSelector;

      PolicyConfig._removePolicyRecord(s_securityPolicies, policyName, remoteChainSelector);
      delete s_policySecurityConfigs[policyName][remoteChainSelector];
      emit SecurityPolicyConfigRemoved(policyName, remoteChainSelector);
    }

    for (uint256 i = 0; i < policiesToRemove.length; ++i) {
      bytes32 policyName = policiesToRemove[i];

      PolicyConfig._removePolicy(s_securityPolicies, policyName);
      emit SecurityPolicyRemoved(policyName);
    }

    for (uint256 i = 0; i < policyConfigsToSet.length; ++i) {
      SecurityPolicyConfigArgs calldata args = policyConfigsToSet[i];
      SecurityConfig calldata securityConfig = args.securityConfig;

      _validateSecurityConfig(securityConfig.outboundCCVs, securityConfig.inboundCCVs);

      PolicyConfig._setPolicyRecord(s_securityPolicies, args.policyName, args.remoteChainSelector);
      _setSecurityConfig(s_policySecurityConfigs[args.policyName][args.remoteChainSelector], securityConfig);
      emit SecurityPolicyConfigSet(args.policyName, args.remoteChainSelector, securityConfig);
    }
  }

  /// @dev Rejects duplicate CCVs. address(0) is allowed and means the default CCVs.
  function _validateSecurityConfig(
    address[] calldata outboundCCVs,
    address[] calldata inboundCCVs
  ) private pure {
    CCVConfigValidation._assertNoDuplicates(outboundCCVs);
    CCVConfigValidation._assertNoDuplicates(inboundCCVs);
  }

  /// @dev Any address is accepted for the hook. A MultiTokenPoolHooks must list every pool that serves a token of its
  /// lanes as an authorized caller.
  function _setSecurityConfig(
    SecurityConfig storage storedConfig,
    SecurityConfig calldata config
  ) private {
    storedConfig.outboundCCVs = config.outboundCCVs;
    storedConfig.inboundCCVs = config.inboundCCVs;
    storedConfig.hook = config.hook;
  }
}
