// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {RateLimiter} from "../../../../../libraries/RateLimiter.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {BaseTest} from "../../../../BaseTest.t.sol";

/// @notice One SecurityPolicyStore and two registered tokens, A and B, that have no policy. A pool config manager is the
/// only authorized caller, so each test pranks it. The store keeps only the config, so the tests read the buckets and
/// the records back from the store.
contract SecurityPolicyStoreSetup is BaseTest {
  bytes32 internal constant DEFAULT_POLICY = "DEFAULT";
  bytes32 internal constant OTHER_POLICY = "OTHER";
  /// @dev A third chain, Arbitrum One, for records that the setup does not add.
  uint64 internal constant OTHER_CHAIN_SELECTOR = 4949039107694359620;
  uint128 internal constant CAPACITY = 1000e18;

  SecurityPolicyStore internal s_store;

  address internal s_tokenA;
  address internal s_tokenB;
  address internal s_poolConfigManager = makeAddr("poolConfigManager");
  address internal s_hook = makeAddr("hook");
  address internal s_ccv1 = makeAddr("ccv1");
  address internal s_ccv2 = makeAddr("ccv2");

  function setUp() public virtual override {
    super.setUp();
    // BaseTest leaves an ongoing OWNER prank. The store writes come from the pool config manager, so end it.
    vm.stopPrank();

    s_store = new SecurityPolicyStore(_arrayOf(s_poolConfigManager));
    s_tokenA = makeAddr("tokenA");
    s_tokenB = makeAddr("tokenB");

    PolicyConfig.TokenProfileArgs[] memory profiles = new PolicyConfig.TokenProfileArgs[](2);
    profiles[0] = _getTokenProfile(s_tokenA);
    profiles[1] = _getTokenProfile(s_tokenB);
    _applyTokenProfileUpdates(new address[](0), profiles);
  }

  // ================================================================
  // │                         Test config                          │
  // ================================================================

  function _getTokenProfile(
    address localToken
  ) internal pure returns (PolicyConfig.TokenProfileArgs memory) {
    return PolicyConfig.TokenProfileArgs({
      localToken: localToken,
      localDecimals: DEFAULT_TOKEN_DECIMALS,
      lockOrBurnFlavor: IPoolConfigManager.LockOrBurnFlavor.Burn,
      allowedFinalityConfig: bytes4(0),
      lockOrBurnAddress: address(0),
      additionalPoolDataProvider: address(0)
    });
  }

  /// @dev An enabled limit with a rate of one token unit per second.
  function _getRateLimitConfig(
    uint128 capacity
  ) internal pure returns (RateLimiter.Config memory) {
    return RateLimiter.Config({isEnabled: true, capacity: capacity, rate: 1});
  }

  function _getFastFinalityConfig() internal pure returns (bytes4) {
    return FinalityCodec._encodeBlockDepth(10);
  }

  function _lanesOf(
    address localToken,
    uint64 remoteChainSelector
  ) internal pure returns (IPoolConfigManager.LaneKey[] memory lanes) {
    lanes = new IPoolConfigManager.LaneKey[](1);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: localToken, remoteChainSelector: remoteChainSelector});
    return lanes;
  }

  function _trueArrayOf() internal pure returns (bool[] memory flags) {
    flags = new bool[](1);
    flags[0] = true;
    return flags;
  }

  function _ccvsOf(
    address firstCCV,
    address secondCCV
  ) internal pure returns (address[] memory ccvs) {
    ccvs = new address[](2);
    ccvs[0] = firstCCV;
    ccvs[1] = secondCCV;
    return ccvs;
  }

  // ================================================================
  // │                       Store wrappers                         │
  // ================================================================

  function _applyTokenProfileUpdates(
    address[] memory tokensToRemove,
    PolicyConfig.TokenProfileArgs[] memory profilesToSet
  ) internal {
    vm.prank(s_poolConfigManager);
    s_store.applyTokenProfileUpdates(tokensToRemove, profilesToSet);
  }

  function _setTokenRateLimitPolicy(
    address localToken,
    bytes32 policyName
  ) internal {
    PolicyConfig.TokenPolicyArgs[] memory tokenPolicies = new PolicyConfig.TokenPolicyArgs[](1);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: localToken, policyName: policyName});
    vm.prank(s_poolConfigManager);
    s_store.applyTokenRateLimitPolicyUpdates(tokenPolicies);
  }

  function _setTokenSecurityPolicy(
    address localToken,
    bytes32 policyName
  ) internal {
    PolicyConfig.TokenPolicyArgs[] memory tokenPolicies = new PolicyConfig.TokenPolicyArgs[](1);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: localToken, policyName: policyName});
    vm.prank(s_poolConfigManager);
    s_store.applyTokenSecurityPolicyUpdates(tokenPolicies);
  }

  /// @dev Sets the same limit in both directions for both finality classes.
  function _setInlineRateLimitConfig(
    address localToken,
    uint64 remoteChainSelector,
    RateLimiter.Config memory config
  ) internal {
    _setInlineRateLimitConfigPerSlot(localToken, remoteChainSelector, config, config, config, config);
  }

  /// @dev Gives each of the four configs its own value, so a read-back finds a config in a wrong slot.
  function _setInlineRateLimitConfigPerSlot(
    address localToken,
    uint64 remoteChainSelector,
    RateLimiter.Config memory outbound,
    RateLimiter.Config memory inbound,
    RateLimiter.Config memory fastFinalityOutbound,
    RateLimiter.Config memory fastFinalityInbound
  ) internal {
    SecurityPolicyStore.InlineRateLimitConfigArgs[] memory inlineConfigs =
      new SecurityPolicyStore.InlineRateLimitConfigArgs[](1);
    inlineConfigs[0] = SecurityPolicyStore.InlineRateLimitConfigArgs({
      localToken: localToken,
      remoteChainSelector: remoteChainSelector,
      rateLimitConfig: SecurityPolicyStore.RateLimitConfig({
        outbound: outbound,
        inbound: inbound,
        fastFinalityOutbound: fastFinalityOutbound,
        fastFinalityInbound: fastFinalityInbound
      })
    });
    _setInlineRateLimitRecords(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function _setInlineRateLimitRecords(
    IPoolConfigManager.LaneKey[] memory inlineConfigsToRemove,
    SecurityPolicyStore.InlineRateLimitConfigArgs[] memory inlineConfigsToSet
  ) internal {
    vm.prank(s_poolConfigManager);
    s_store.applyInlineRateLimitConfigUpdates(inlineConfigsToRemove, inlineConfigsToSet);
  }

  /// @dev Sets the same limit in both directions for both finality classes.
  function _setRateLimitPolicyConfig(
    bytes32 policyName,
    uint64 remoteChainSelector,
    RateLimiter.Config memory config
  ) internal {
    _setRateLimitPolicyConfigPerSlot(policyName, remoteChainSelector, config, config, config, config);
  }

  function _setRateLimitPolicyConfigPerSlot(
    bytes32 policyName,
    uint64 remoteChainSelector,
    RateLimiter.Config memory outbound,
    RateLimiter.Config memory inbound,
    RateLimiter.Config memory fastFinalityOutbound,
    RateLimiter.Config memory fastFinalityInbound
  ) internal {
    SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory policyConfigs =
      new SecurityPolicyStore.RateLimitPolicyConfigArgs[](1);
    policyConfigs[0] = SecurityPolicyStore.RateLimitPolicyConfigArgs({
      policyName: policyName,
      remoteChainSelector: remoteChainSelector,
      localDecimals: DEFAULT_TOKEN_DECIMALS,
      rateLimitConfig: SecurityPolicyStore.RateLimitConfig({
        outbound: outbound,
        inbound: inbound,
        fastFinalityOutbound: fastFinalityOutbound,
        fastFinalityInbound: fastFinalityInbound
      })
    });
    vm.prank(s_poolConfigManager);
    s_store.applyRateLimitPolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs);
  }

  function _setInlineSecurityConfig(
    address localToken,
    uint64 remoteChainSelector,
    address[] memory outboundCCVs,
    address[] memory inboundCCVs,
    address hook
  ) internal {
    SecurityPolicyStore.InlineSecurityConfigArgs[] memory configs =
      new SecurityPolicyStore.InlineSecurityConfigArgs[](1);
    configs[0] = SecurityPolicyStore.InlineSecurityConfigArgs({
      localToken: localToken,
      remoteChainSelector: remoteChainSelector,
      securityConfig: SecurityPolicyStore.SecurityConfig({
        outboundCCVs: outboundCCVs, inboundCCVs: inboundCCVs, hook: hook
      })
    });
    vm.prank(s_poolConfigManager);
    s_store.applyInlineSecurityConfigUpdates(new IPoolConfigManager.LaneKey[](0), configs);
  }

  function _setSecurityPolicyConfig(
    bytes32 policyName,
    uint64 remoteChainSelector,
    address[] memory outboundCCVs,
    address[] memory inboundCCVs,
    address hook
  ) internal {
    SecurityPolicyStore.SecurityPolicyConfigArgs[] memory configs =
      new SecurityPolicyStore.SecurityPolicyConfigArgs[](1);
    configs[0] = SecurityPolicyStore.SecurityPolicyConfigArgs({
      policyName: policyName,
      remoteChainSelector: remoteChainSelector,
      securityConfig: SecurityPolicyStore.SecurityConfig({
        outboundCCVs: outboundCCVs, inboundCCVs: inboundCCVs, hook: hook
      })
    });
    vm.prank(s_poolConfigManager);
    s_store.applySecurityPolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), configs);
  }

  // ================================================================
  // │                          Read back                           │
  // ================================================================

  function _getRateLimiterState(
    address localToken,
    uint64 remoteChainSelector
  ) internal view returns (IPoolConfigManager.LaneRateLimiterState memory) {
    IPoolConfigManager.LaneRateLimiterState[] memory states =
      s_store.getCurrentRateLimiterStateBatch(_lanesOf(localToken, remoteChainSelector), _trueArrayOf());
    return states[0];
  }

  function _getInlineRateLimitConfig(
    address localToken,
    uint64 remoteChainSelector
  ) internal view returns (SecurityPolicyStore.InlineRateLimitConfigArgs memory) {
    SecurityPolicyStore.InlineRateLimitConfigArgs[] memory configs = s_store.getInlineRateLimitConfigs(localToken);
    for (uint256 i = 0; i < configs.length; ++i) {
      if (configs[i].remoteChainSelector == remoteChainSelector) return configs[i];
    }
    revert("inline record not found");
  }

  /// @dev Checks that a bucket is full and holds the config.
  function _assertFullBucket(
    RateLimiter.Config memory expectedConfig,
    IPoolConfigManager.TokenBucket memory bucket
  ) internal pure {
    assertEq(expectedConfig.capacity, bucket.tokens);
    assertEq(expectedConfig.capacity, bucket.capacity);
    assertEq(expectedConfig.rate, bucket.rate);
    assertEq(expectedConfig.isEnabled, bucket.isEnabled);
  }
}
