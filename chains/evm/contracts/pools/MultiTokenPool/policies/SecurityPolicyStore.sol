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

  error ZeroAddressInvalid();
  error InvalidDecimalArgs(uint8 expected, uint8 actual);
  error InvalidRateLimitPolicyDecimals(bytes32 policyName, uint8 expected, uint8 actual);
  error TokenHasConfig(address localToken);
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
  event RateLimitPolicyConfigSet(
    bytes32 indexed policyName,
    uint64 indexed remoteChainSelector,
    uint8 localDecimals,
    RateLimiter.Config outboundRateLimiterConfig,
    RateLimiter.Config inboundRateLimiterConfig,
    RateLimiter.Config fastFinalityOutboundRateLimiterConfig,
    RateLimiter.Config fastFinalityInboundRateLimiterConfig
  );
  event RateLimitPolicyConfigRemoved(bytes32 indexed policyName, uint64 indexed remoteChainSelector);
  /// @dev Also stands for the removal of the decimals of the policy.
  event RateLimitPolicyRemoved(bytes32 indexed policyName);
  event RateLimitPolicyDirectionSet(
    bytes32 indexed policyName,
    uint64 indexed remoteChainSelector,
    uint8 localDecimals,
    IPoolV2.MessageDirection direction,
    RateLimiter.Config defaultFinalityConfig,
    RateLimiter.Config fastFinalityConfig
  );
  /// @dev A set of a new token also registers it.
  event TokenProfileSet(address indexed localToken, uint8 localDecimals);
  event TokenRemoved(address indexed localToken);
  event TokenRateLimitPolicySet(address indexed localToken, bytes32 indexed policyName);
  event TokenSecurityPolicySet(address indexed localToken, bytes32 indexed policyName);
  event InlineSecurityConfigSet(
    address indexed localToken,
    uint64 indexed remoteChainSelector,
    address[] outboundCCVs,
    address[] inboundCCVs,
    address hook
  );
  event InlineSecurityConfigRemoved(address indexed localToken, uint64 indexed remoteChainSelector);
  event SecurityPolicyConfigSet(
    bytes32 indexed policyName,
    uint64 indexed remoteChainSelector,
    address[] outboundCCVs,
    address[] inboundCCVs,
    address hook
  );
  event SecurityPolicyConfigRemoved(bytes32 indexed policyName, uint64 indexed remoteChainSelector);
  event SecurityPolicyRemoved(bytes32 indexed policyName);

  /// @notice The state of one token in this store.
  struct TokenState {
    uint8 localDecimals; // Decimals of the local token.
    bytes32 rateLimitPolicyName; // Identifier of the rate limit policy attached to the token, or zero.
    bytes32 securityPolicyName; // Identifier of the security policy attached to the token, or zero.
  }

  /// @notice An inline rate limit record to set and to read. Amounts are in local token units. While an FTF config is
  /// disabled, an FTF transfer in its direction uses the default finality bucket of the same record.
  struct InlineRateLimitConfigArgs {
    address localToken; // ────────╮ A registered token.
    uint64 remoteChainSelector; // ╯ The remote chain selector.
    RateLimiter.Config outboundRateLimiterConfig; // The outbound limit for default finality.
    RateLimiter.Config inboundRateLimiterConfig; // The inbound limit for default finality.
    RateLimiter.Config fastFinalityOutboundRateLimiterConfig; // The outbound limit for FTF.
    RateLimiter.Config fastFinalityInboundRateLimiterConfig; // The inbound limit for FTF.
  }

  /// @notice One direction of an inline rate limit record to update. Both finality classes of the direction are set
  /// together, so the direction is never left with a stale FTF config.
  struct InlineRateLimitDirectionArgs {
    address localToken; // ────────╮ A registered token.
    uint64 remoteChainSelector; // ╯ The remote chain. Not zero.
    IPoolV2.MessageDirection direction; // The direction to update.
    RateLimiter.Config defaultFinalityConfig; // The limit for default finality.
    RateLimiter.Config fastFinalityConfig; // The limit for FTF. Disabled means FTF uses the default bucket.
  }

  /// @notice A rate limit policy record to set and to read. Amounts are in local token units. While an FTF config is
  /// disabled, an FTF transfer in its direction uses the default finality bucket of the same record.
  struct RateLimitPolicyConfigArgs {
    bytes32 policyName; // The rate limit policy. Not zero.
    uint64 remoteChainSelector; // ─╮ The remote chain. Zero names the policy default record.
    uint8 localDecimals; // ────────╯ The decimals of the policy. The first record of the policy fixes them.
    RateLimiter.Config outboundRateLimiterConfig; // The outbound limit for default finality.
    RateLimiter.Config inboundRateLimiterConfig; // The inbound limit for default finality.
    RateLimiter.Config fastFinalityOutboundRateLimiterConfig; // The outbound limit for FTF.
    RateLimiter.Config fastFinalityInboundRateLimiterConfig; // The inbound limit for FTF.
  }

  /// @notice One direction of a rate limit policy record to update. Both finality classes of the direction are set
  /// together, so the direction is never left with a stale FTF config.
  struct RateLimitPolicyDirectionArgs {
    bytes32 policyName; // The rate limit policy. Not zero.
    uint64 remoteChainSelector; // ─────────╮ The remote chain. Zero names the policy default record.
    IPoolV2.MessageDirection direction; // ─╯ The direction to update.
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
    address localToken; // ────────╮ A registered token.
    uint64 remoteChainSelector; // ╯ The remote chain. Not zero.
    address[] outboundCCVs; // CCVs required for outgoing messages to the remote chain. No duplicate.
    address[] inboundCCVs; // CCVs required for incoming messages from the remote chain. No duplicate.
    address hook; // Address of the IMultiTokenPoolHooks contract of the lane. Zero means no hook.
  }

  /// @notice A security policy record to set.
  struct SecurityPolicyConfigArgs {
    bytes32 policyName; // The security policy. Not zero.
    uint64 remoteChainSelector; // The remote chain. Zero names the policy default record.
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
  // │                           Tokens                             │
  // ================================================================

  /// @notice Deregisters tokens, then sets the profile of tokens. A set of a new token registers it with no policy.
  /// All three stores apply the same call in one transaction.
  /// @dev This store keeps only the decimals of the profile. The connection store checks the decimals of a new token
  /// against the token in the same transaction, so this store copies the decimals of a new token without a check. A set
  /// of a registered token needs the stored decimals, the same rule as in the connection store. So a caller that writes
  /// only this store cannot change the decimals of a registered token. A deregistration needs a token without a policy
  /// and without an inline record of a type of this store. A deregistered token can be registered again, also with other
  /// decimals.
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
        revert InvalidDecimalArgs(tokenState.localDecimals, args.localDecimals);
      }

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
      PolicyConfig._moveToken(s_rateLimitPolicies, tokenState.rateLimitPolicyName, policyName);
      tokenState.rateLimitPolicyName = policyName;
      emit TokenRateLimitPolicySet(localToken, policyName);
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
      PolicyConfig._moveToken(s_securityPolicies, tokenState.securityPolicyName, policyName);
      tokenState.securityPolicyName = policyName;
      emit TokenSecurityPolicySet(localToken, policyName);
    }
  }

  // ================================================================
  // │                        Rate limiting                         │
  // ================================================================

  /// @notice Removes inline rate limit records, and creates or replaces them. A set writes all four configs of the
  /// record.
  /// @dev The removals run first, then the sets. A set creates the record of its lane if it does not exist, so a
  /// removal and a set of the same lane in one call replaces the record. Each config follows the checks of
  /// RateLimiter. A set writes a bucket only when its config changes, so a set of the same config keeps the state of
  /// the bucket. A removal deletes the buckets of the record. To refill a bucket without a config change, remove the
  /// record and set it again in one call, or raise its capacity.
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
      _validateRateLimitConfig(args.outboundRateLimiterConfig);
      _validateRateLimitConfig(args.inboundRateLimiterConfig);
      _validateRateLimitConfig(args.fastFinalityOutboundRateLimiterConfig);
      _validateRateLimitConfig(args.fastFinalityInboundRateLimiterConfig);

      PolicyConfig._setInlineRecord(
        s_tokens, s_inlineRateLimitChains[args.localToken], args.localToken, args.remoteChainSelector
      );
      mapping(bool fastFinality => mapping(IPoolV2.MessageDirection direction => RateLimiter.TokenBucket bucket))
        storage buckets = s_tokenBuckets[args.localToken][args.remoteChainSelector];
      _setBucketIfChanged(buckets[false][IPoolV2.MessageDirection.Outbound], args.outboundRateLimiterConfig);
      _setBucketIfChanged(buckets[false][IPoolV2.MessageDirection.Inbound], args.inboundRateLimiterConfig);
      _setBucketIfChanged(buckets[true][IPoolV2.MessageDirection.Outbound], args.fastFinalityOutboundRateLimiterConfig);
      _setBucketIfChanged(buckets[true][IPoolV2.MessageDirection.Inbound], args.fastFinalityInboundRateLimiterConfig);
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

  /// @notice Sets one direction of inline rate limit records. Both finality classes of the direction are set together,
  /// so the direction is never left with a stale FTF config. This writes only the two buckets of the direction, so a
  /// rate limit manager raises one direction without touching the other, and a set of the same config keeps the state
  /// of the bucket.
  /// @dev The record must exist, so a caller creates it with applyInlineRateLimitConfigUpdates first. Each config
  /// follows the checks of RateLimiter.
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
      _setBucketIfChanged(buckets[false][args.direction], args.defaultFinalityConfig);
      _setBucketIfChanged(buckets[true][args.direction], args.fastFinalityConfig);
      emit InlineRateLimitDirectionSet(
        args.localToken, args.remoteChainSelector, args.direction, args.defaultFinalityConfig, args.fastFinalityConfig
      );
    }
  }

  /// @notice Removes and sets rate limit policy records, and removes rate limit policies. A set writes all four
  /// configs of the record.
  /// @dev The record removals run first, then the policy removals, then the record sets. The first record set under a
  /// new name creates the policy and fixes its decimals, which stay fixed until the policy is removed. A record removal
  /// keeps the policy. A policy removal needs a policy without tokens and without records. Chain selector zero names
  /// the policy default record. Each config follows the checks of RateLimiter. A set changes no bucket. A bucket of a
  /// token that uses a policy record copies the config of its record on its next use, when the config of the bucket
  /// differs, which refills the bucket. So a set of the same config refills no bucket. After a removal of a policy
  /// record, the buckets of its tokens keep their state. To refill a bucket without a config change, raise its
  /// capacity.
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
      // isNewPolicy is read before _setPolicyRecord, which adds the name of a new policy.
      bool isNewPolicy = !s_rateLimitPolicies.policyNames.contains(args.policyName);
      if (!isNewPolicy && s_rateLimitPolicyDecimals[args.policyName] != args.localDecimals) {
        revert InvalidRateLimitPolicyDecimals(
          args.policyName, s_rateLimitPolicyDecimals[args.policyName], args.localDecimals
        );
      }
      _validateRateLimitConfig(args.outboundRateLimiterConfig);
      _validateRateLimitConfig(args.inboundRateLimiterConfig);
      _validateRateLimitConfig(args.fastFinalityOutboundRateLimiterConfig);
      _validateRateLimitConfig(args.fastFinalityInboundRateLimiterConfig);

      PolicyConfig._setPolicyRecord(s_rateLimitPolicies, args.policyName, args.remoteChainSelector);
      if (isNewPolicy) {
        s_rateLimitPolicyDecimals[args.policyName] = args.localDecimals;
      }
      mapping(bool fastFinality => mapping(IPoolV2.MessageDirection direction => RateLimiter.Config config)) storage
        recordConfigs = s_policyRateLimitConfigs[args.policyName][args.remoteChainSelector];
      recordConfigs[false][IPoolV2.MessageDirection.Outbound] = args.outboundRateLimiterConfig;
      recordConfigs[false][IPoolV2.MessageDirection.Inbound] = args.inboundRateLimiterConfig;
      recordConfigs[true][IPoolV2.MessageDirection.Outbound] = args.fastFinalityOutboundRateLimiterConfig;
      recordConfigs[true][IPoolV2.MessageDirection.Inbound] = args.fastFinalityInboundRateLimiterConfig;
      emit RateLimitPolicyConfigSet(
        args.policyName,
        args.remoteChainSelector,
        args.localDecimals,
        args.outboundRateLimiterConfig,
        args.inboundRateLimiterConfig,
        args.fastFinalityOutboundRateLimiterConfig,
        args.fastFinalityInboundRateLimiterConfig
      );
    }
  }

  /// @notice Sets one direction of rate limit policy records. Both finality classes of the direction are set together,
  /// so the direction is never left with a stale FTF config.
  /// @dev The record must exist, so a caller creates it with applyRateLimitPolicyConfigUpdates first. Chain selector
  /// zero names the policy default record. Each config follows the checks of RateLimiter. A set changes no bucket. A
  /// bucket of a token that uses a policy record copies the config of its record on its next use, when the config of
  /// the bucket differs, which refills the bucket. So a set of the same config refills no bucket.
  /// @param directionConfigsToSet The direction records to set.
  function applyRateLimitPolicyDirectionUpdates(
    RateLimitPolicyDirectionArgs[] calldata directionConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < directionConfigsToSet.length; ++i) {
      RateLimitPolicyDirectionArgs calldata args = directionConfigsToSet[i];
      // A record exists exactly when its chain is in the record set of its policy. Zero names the default record.
      if (!s_rateLimitPolicies.policies[args.policyName].remoteChainSelectors.contains(args.remoteChainSelector)) {
        revert PolicyConfig.PolicyRecordNotFound(args.policyName, args.remoteChainSelector);
      }
      _validateRateLimitConfig(args.defaultFinalityConfig);
      _validateRateLimitConfig(args.fastFinalityConfig);

      mapping(bool fastFinality => mapping(IPoolV2.MessageDirection direction => RateLimiter.Config config)) storage
        recordConfigs = s_policyRateLimitConfigs[args.policyName][args.remoteChainSelector];
      recordConfigs[false][args.direction] = args.defaultFinalityConfig;
      recordConfigs[true][args.direction] = args.fastFinalityConfig;
      emit RateLimitPolicyDirectionSet(
        args.policyName,
        args.remoteChainSelector,
        s_rateLimitPolicyDecimals[args.policyName],
        args.direction,
        args.defaultFinalityConfig,
        args.fastFinalityConfig
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
  /// @dev The removals run first, then the sets. Each set runs the checks of CCVConfigValidation on both CCV lists.
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

  /// @notice Removes and sets security policy records, and removes security policies. A set writes every value of the
  /// record, so it replaces the record.
  /// @dev The record removals run first, then the policy removals, then the record sets. The first record set under a
  /// new name creates the policy. A record removal keeps the policy. A policy removal needs a policy without tokens
  /// and without records. Chain selector zero names the policy default record. Each set runs the checks of
  /// CCVConfigValidation on both CCV lists.
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
      _validateSecurityConfig(args.outboundCCVs, args.inboundCCVs);
      PolicyConfig._setPolicyRecord(s_securityPolicies, args.policyName, args.remoteChainSelector);
      _setSecurityConfig(
        s_policySecurityConfigs[args.policyName][args.remoteChainSelector],
        args.outboundCCVs,
        args.inboundCCVs,
        args.hook
      );
      emit SecurityPolicyConfigSet(
        args.policyName, args.remoteChainSelector, args.outboundCCVs, args.inboundCCVs, args.hook
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
