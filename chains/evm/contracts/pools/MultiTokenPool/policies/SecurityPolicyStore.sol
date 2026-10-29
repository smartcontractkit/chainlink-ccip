// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {CCVConfigValidation} from "../../../libraries/CCVConfigValidation.sol";
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
/// A security record holds the outbound and inbound CCVs and the MultiTokenPoolHook.
contract SecurityPolicyStore is ITypeAndVersion, AuthorizedCallers {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;
  using EnumerableSet for EnumerableSet.UintSet;
  using RateLimiter for RateLimiter.TokenBucket;

  error InlineRateLimitRecordExists(address localToken, uint64 remoteChainSelector);
  error RateLimitRecordNotFound(address localToken, uint64 remoteChainSelector);

  event InlineRateLimitConfigSet(
    address indexed localToken,
    uint64 indexed remoteChainSelector,
    RateLimiter.Config outboundRateLimiterConfig,
    RateLimiter.Config inboundRateLimiterConfig,
    RateLimiter.Config fastFinalityOutboundRateLimiterConfig,
    RateLimiter.Config fastFinalityInboundRateLimiterConfig
  );
  event InlineRateLimitConfigRemoved(address indexed localToken, uint64 indexed remoteChainSelector);
  event InlineRateLimitDirectionSet(
    address indexed localToken,
    uint64 indexed remoteChainSelector,
    IPoolV2.MessageDirection direction,
    RateLimiter.Config defaultFinalityConfig,
    RateLimiter.Config fastFinalityConfig
  );
  event InlineSecurityConfigSet(
    address indexed localToken,
    uint64 indexed remoteChainSelector,
    address[] outboundCCVs,
    address[] inboundCCVs,
    address hook
  );
  event InlineSecurityConfigRemoved(address indexed localToken, uint64 indexed remoteChainSelector);

  /// @notice The state of one token in this store.
  struct TokenState {
    uint8 localDecimals; // Decimals of the local token.
    bytes32 rateLimitPolicyName; // Identifier of the rate limit policy attached to the token, or zero.
    bytes32 securityPolicyName; // Identifier of the security policy attached to the token, or zero.
  }

  /// @notice An inline rate limit record to set and to read. Amounts are in local token units. While an FTF config is
  /// disabled, an FTF transfer in its direction uses the default finality bucket of the same record.
  struct InlineRateLimitConfigArgs {
    address localToken; // A registered token.
    uint64 remoteChainSelector; // Remote chain selector.
    RateLimiter.Config outboundRateLimiterConfig; // The outbound limit for default finality, meaning the rate limits for all of the onRamps for the given chain.
    RateLimiter.Config inboundRateLimiterConfig; // The inbound limit for default finality, meaning the rate limits for all of the offRamps for the given chain.
    RateLimiter.Config fastFinalityOutboundRateLimiterConfig; // The outbound limit for FTF.
    RateLimiter.Config fastFinalityInboundRateLimiterConfig; // The inbound limit for FTF.
  }

  /// @notice One direction of an inline rate limit record. Both finality classes of the direction are set together, so
  /// the direction is never left with a stale FTF config.
  struct InlineRateLimitDirectionArgs {
    address localToken; // A registered token.
    uint64 remoteChainSelector; // Remote chain selector.
    IPoolV2.MessageDirection direction; // The direction to set.
    RateLimiter.Config defaultFinalityConfig; // The limit for default finality.
    RateLimiter.Config fastFinalityConfig; // The limit for FTF. Disabled means FTF uses the default bucket.
  }

  /// @notice Security config describing the verification and hook requirements for a lane.
  /// @dev An address(0) entry stands for the default CCVs of the ramps. A list has at most one entry per CCV.
  struct SecurityConfig {
    address[] outboundCCVs; // CCVs required for outgoing messages to the remote chain.
    address[] inboundCCVs; // CCVs required for incoming messages from the remote chain.
    address hook; // Address of the IMultiTokenPoolHooks contract of the lane. Zero means no hook.
  }

  /// @notice An inline security record to set and to read.
  struct InlineSecurityConfigArgs {
    address localToken; // A registered token.
    uint64 remoteChainSelector; // Remote chain selector.
    address[] outboundCCVs; // CCVs required for outgoing messages to the remote chain. No duplicate.
    address[] inboundCCVs; // CCVs required for incoming messages from the remote chain. No duplicate.
    address hook; // Address of the IMultiTokenPoolHooks contract of the lane. Zero means no hook.
  }

  /// @notice The registered tokens. A token is registered exactly when it is in this set.
  /// @dev Each store keeps its own list of the registered tokens. The pool config manager registers and deregisters a
  /// token in every store in one transaction.
  EnumerableSet.AddressSet private s_tokens;

  /// @notice The state of each token: the policies it is bound to, and its decimals.
  /// @dev The decimals stay after a deregistration.
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
  /// transfers. When the FTF buckets are not configured, the default bucket of the same direction is used for all
  /// transfers regardless of the finality requirements.
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
  // │                        Stored views                          │
  // ================================================================

  /// @notice Returns the inline rate limit records of a token, in the shape of the setter. Never reverts.
  /// @dev A token without an inline record on a chain has no entry. An inline record keeps its config in the buckets of
  /// the token, so this reads the buckets.
  /// @param localToken The local token.
  /// @return inlineConfigs The inline rate limit records of the token.
  function getInlineRateLimitConfigs(
    address localToken
  ) external view returns (InlineRateLimitConfigArgs[] memory inlineConfigs) {
    uint256[] memory chains = s_inlineRateLimitChains[localToken].values();
    inlineConfigs = new InlineRateLimitConfigArgs[](chains.length);
    for (uint256 i = 0; i < chains.length; ++i) {
      uint64 remoteChainSelector = uint64(chains[i]);
      inlineConfigs[i] = InlineRateLimitConfigArgs({
        localToken: localToken,
        remoteChainSelector: remoteChainSelector,
        outboundRateLimiterConfig: _getInlineRateLimitConfig(
          localToken, remoteChainSelector, false, IPoolV2.MessageDirection.Outbound
        ),
        inboundRateLimiterConfig: _getInlineRateLimitConfig(
          localToken, remoteChainSelector, false, IPoolV2.MessageDirection.Inbound
        ),
        fastFinalityOutboundRateLimiterConfig: _getInlineRateLimitConfig(
          localToken, remoteChainSelector, true, IPoolV2.MessageDirection.Outbound
        ),
        fastFinalityInboundRateLimiterConfig: _getInlineRateLimitConfig(
          localToken, remoteChainSelector, true, IPoolV2.MessageDirection.Inbound
        )
      });
    }
  }

  /// @notice Returns the inline security records of a token.
  /// @dev A token without an inline record on a chain has no entry.
  /// @param localToken The local token.
  /// @return inlineConfigs The inline security records of the token.
  function getInlineSecurityConfigs(
    address localToken
  ) external view returns (InlineSecurityConfigArgs[] memory inlineConfigs) {
    uint256[] memory chains = s_inlineSecurityChains[localToken].values();
    inlineConfigs = new InlineSecurityConfigArgs[](chains.length);
    for (uint256 i = 0; i < chains.length; ++i) {
      uint64 remoteChainSelector = uint64(chains[i]);
      SecurityConfig storage securityConfig = s_inlineSecurityConfigs[localToken][remoteChainSelector];
      inlineConfigs[i] = InlineSecurityConfigArgs({
        localToken: localToken,
        remoteChainSelector: remoteChainSelector,
        outboundCCVs: securityConfig.outboundCCVs,
        inboundCCVs: securityConfig.inboundCCVs,
        hook: securityConfig.hook
      });
    }
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

  // ================================================================
  // │                        Rate limiting                         │
  // ================================================================

  /// @notice Removes inline rate limit records and creates new ones. Each set writes all four configs of a new record.
  /// To change one direction of an existing record, use applyInlineRateLimitDirectionUpdates. A set refills every
  /// bucket of the record, so it cannot change one direction: the pass-through write would refill the other direction
  /// and re-grant the tokens that transfers consumed since the caller read it.
  /// @dev Only callable by an authorized pool config manager. The removals run first, then the sets. A set needs a
  /// record that does not exist, so a removal and a set of the same lane in one call replaces the record. Each config
  /// follows the checks of RateLimiter. A removal deletes the buckets of the record.
  /// @param inlineConfigsToRemove The inline records to remove.
  /// @param inlineConfigsToSet The inline records to create.
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
      if (s_inlineRateLimitChains[args.localToken].contains(args.remoteChainSelector)) {
        revert InlineRateLimitRecordExists(args.localToken, args.remoteChainSelector);
      }
      _validateRateLimitConfig(args.outboundRateLimiterConfig);
      _validateRateLimitConfig(args.inboundRateLimiterConfig);
      _validateRateLimitConfig(args.fastFinalityOutboundRateLimiterConfig);
      _validateRateLimitConfig(args.fastFinalityInboundRateLimiterConfig);

      PolicyConfig._setInlineRecord(
        s_tokens, s_inlineRateLimitChains[args.localToken], args.localToken, args.remoteChainSelector
      );
      mapping(bool fastFinality => mapping(IPoolV2.MessageDirection direction => RateLimiter.TokenBucket bucket))
        storage buckets = s_tokenBuckets[args.localToken][args.remoteChainSelector];
      buckets[false][IPoolV2.MessageDirection.Outbound]._setTokenBucketConfig(args.outboundRateLimiterConfig);
      buckets[false][IPoolV2.MessageDirection.Inbound]._setTokenBucketConfig(args.inboundRateLimiterConfig);
      buckets[true][IPoolV2.MessageDirection.Outbound]._setTokenBucketConfig(args.fastFinalityOutboundRateLimiterConfig);
      buckets[true][IPoolV2.MessageDirection.Inbound]._setTokenBucketConfig(args.fastFinalityInboundRateLimiterConfig);
      emit InlineRateLimitConfigSet(
        args.localToken,
        args.remoteChainSelector,
        args.outboundRateLimiterConfig,
        args.inboundRateLimiterConfig,
        args.fastFinalityOutboundRateLimiterConfig,
        args.fastFinalityInboundRateLimiterConfig
      );
    }
  }

  /// @notice Sets the rate limit of one direction of each listed token and chain. Both finality classes of the
  /// direction are set together, so the direction is never left with a stale FTF config. This writes only the two
  /// buckets of the direction, so a rate limit manager can raise one direction without refilling the other, which
  /// would re-grant the tokens that transfers consumed since it read the other direction.
  /// @dev Only callable by an authorized pool config manager. The record must exist, so a caller creates it with
  /// applyInlineRateLimitConfigUpdates first. Each set refills the two buckets of the direction, also with the same
  /// values. Each config follows the checks of RateLimiter.
  /// @param directionConfigsToSet The direction records to set.
  function applyInlineRateLimitDirectionUpdates(
    InlineRateLimitDirectionArgs[] calldata directionConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < directionConfigsToSet.length; ++i) {
      InlineRateLimitDirectionArgs calldata args = directionConfigsToSet[i];
      if (!s_inlineRateLimitChains[args.localToken].contains(args.remoteChainSelector)) {
        revert RateLimitRecordNotFound(args.localToken, args.remoteChainSelector);
      }
      _validateRateLimitConfig(args.defaultFinalityConfig);
      _validateRateLimitConfig(args.fastFinalityConfig);

      mapping(bool fastFinality => mapping(IPoolV2.MessageDirection direction => RateLimiter.TokenBucket bucket))
        storage buckets = s_tokenBuckets[args.localToken][args.remoteChainSelector];
      buckets[false][args.direction]._setTokenBucketConfig(args.defaultFinalityConfig);
      buckets[true][args.direction]._setTokenBucketConfig(args.fastFinalityConfig);
      emit InlineRateLimitDirectionSet(
        args.localToken, args.remoteChainSelector, args.direction, args.defaultFinalityConfig, args.fastFinalityConfig
      );
    }
  }

  /// @dev The checks of RateLimiter._setTokenBucketConfig. A set runs them for every config before it writes. A policy
  /// record stores a config without a bucket, so these are its only checks.
  function _validateRateLimitConfig(
    RateLimiter.Config calldata config
  ) private pure {
    if (config.isEnabled) {
      if (config.rate > config.capacity) revert RateLimiter.InvalidRateLimitRate(config);
    } else if (config.rate != 0 || config.capacity != 0) {
      revert RateLimiter.DisabledNonZeroRateLimit(config);
    }
  }

  // ================================================================
  // │                           Security                           │
  // ================================================================

  /// @notice Removes and sets inline security records. A set writes every value of the record, so it replaces the
  /// record.
  /// @dev Only callable by an authorized pool config manager. The removals run first, then the sets. Each set runs the
  /// checks of CCVConfigValidation on both CCV lists.
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
      _validateSecurityConfig(args.outboundCCVs, args.inboundCCVs);
      PolicyConfig._setInlineRecord(
        s_tokens, s_inlineSecurityChains[args.localToken], args.localToken, args.remoteChainSelector
      );
      _setSecurityConfig(
        s_inlineSecurityConfigs[args.localToken][args.remoteChainSelector],
        args.outboundCCVs,
        args.inboundCCVs,
        args.hook
      );
      emit InlineSecurityConfigSet(
        args.localToken, args.remoteChainSelector, args.outboundCCVs, args.inboundCCVs, args.hook
      );
    }
  }

  /// @dev A CCV list must not have a duplicate.
  function _validateSecurityConfig(
    address[] calldata outboundCCVs,
    address[] calldata inboundCCVs
  ) private pure {
    CCVConfigValidation._assertNoDuplicates(outboundCCVs);
    CCVConfigValidation._assertNoDuplicates(inboundCCVs);
  }

  /// @dev Any address is accepted for the hook. A MultiTokenPoolHooks must list every pool
  /// that serves a token of its lanes as an authorized caller.
  function _setSecurityConfig(
    SecurityConfig storage securityConfig,
    address[] calldata outboundCCVs,
    address[] calldata inboundCCVs,
    address hook
  ) private {
    securityConfig.outboundCCVs = outboundCCVs;
    securityConfig.inboundCCVs = inboundCCVs;
    securityConfig.hook = hook;
  }
}
