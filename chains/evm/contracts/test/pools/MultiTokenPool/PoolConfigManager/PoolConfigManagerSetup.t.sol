// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";
import {ITokenAdminRegistry} from "../../../../interfaces/ITokenAdminRegistry.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {RateLimiter} from "../../../../libraries/RateLimiter.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {
  SameAddressRemoteTokenAddressProvider
} from "../../../../pools/MultiTokenPool/SameAddressRemoteTokenAddressProvider.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {TokenFixture} from "../../../TokenFixture.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

/// @notice Sets up the manager and three policy stores without a pool. Each policy type has a DEFAULT policy. The
/// connection and remote pool policies have records for the destination chain. The other policy types use default
/// records. The connection policy uses SameAddressRemoteTokenAddressProvider.
/// @dev Rate limits are disabled, fees are zero, and security records have no CCVs or hook. Destination overheads
/// use the default gas and minimum bytes. Token A uses an inline connection and the other DEFAULT policies. Token B
/// uses every DEFAULT policy. Both tokens allow only default finality. The test contract owner holds CONFIG_ROLE.
contract PoolConfigManagerSetup is TokenFixture {
  bytes32 internal constant DEFAULT_POLICY = "DEFAULT";
  bytes32 internal constant OTHER_POLICY = "OTHER";
  // Each policy type has a different name so swapped fields fail the assertions.
  bytes32 internal constant CONNECTION_POLICY = "CONNECTION";
  bytes32 internal constant REMOTE_POOL_POLICY = "REMOTE_POOL";
  bytes32 internal constant RATE_LIMIT_POLICY = "RATE_LIMIT";
  bytes32 internal constant FEE_POLICY = "FEE";
  bytes32 internal constant DEST_OVERHEAD_POLICY = "DEST_OVERHEAD";
  bytes32 internal constant SECURITY_POLICY = "SECURITY";
  /// @dev A third chain, Arbitrum One, for records that the setup does not add.
  uint64 internal constant OTHER_CHAIN_SELECTOR = 4949039107694359620;
  uint32 internal constant MIN_DEST_BYTES_OVERHEAD = 192;
  uint128 internal constant CAPACITY = 1000e18;

  ConnectionPolicyStore internal s_connectionStore;
  FeePolicyStore internal s_feeStore;
  SecurityPolicyStore internal s_securityStore;
  PoolConfigManager internal s_poolConfigManager;

  address internal s_tokenA;
  address internal s_tokenB;

  address internal s_remotePool = makeAddr("remotePool");
  address internal s_remoteTokenA = makeAddr("remoteTokenA");

  function setUp() public virtual override {
    super.setUp();

    s_connectionStore = new ConnectionPolicyStore(new address[](0));
    s_feeStore = new FeePolicyStore(new address[](0));
    s_securityStore = new SecurityPolicyStore(new address[](0));
    s_poolConfigManager = new PoolConfigManager(
      ITokenAdminRegistry(address(s_tokenAdminRegistry)), s_connectionStore, s_feeStore, s_securityStore, OWNER
    );
    AuthorizedCallers.AuthorizedCallerArgs memory callerArgs = AuthorizedCallers.AuthorizedCallerArgs({
      addedCallers: _arrayOf(address(s_poolConfigManager)), removedCallers: new address[](0)
    });
    s_connectionStore.applyAuthorizedCallerUpdates(callerArgs);
    s_feeStore.applyAuthorizedCallerUpdates(callerArgs);
    s_securityStore.applyAuthorizedCallerUpdates(callerArgs);
    s_poolConfigManager.grantRole(s_poolConfigManager.CONFIG_ROLE(), OWNER);

    s_tokenA = _deployToken("A");
    s_tokenB = _deployToken("B");

    _setConnectionPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, false);
    _setRemoteTokenAddressProvider(DEFAULT_POLICY, address(new SameAddressRemoteTokenAddressProvider()));
    _addPolicyDefaultRecords(DEFAULT_POLICY);

    _setTokenProfile(_getTokenProfileArgs(s_tokenA));
    _setTokenRemotePoolPolicy(s_tokenA, DEFAULT_POLICY);
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setTokenFeePolicy(s_tokenA, DEFAULT_POLICY);
    _setTokenDestOverheadPolicy(s_tokenA, DEFAULT_POLICY);
    _setTokenSecurityPolicy(s_tokenA, DEFAULT_POLICY);
    _setInlineConnectionConfig(s_tokenA, DEST_CHAIN_SELECTOR, abi.encode(s_remoteTokenA));
    _registerToken(s_tokenB);
  }

  // ================================================================
  // │                         Arguments                            │
  // ================================================================

  /// @dev A Burn token with the default decimals, default finality only, and no additional pool data provider.
  function _getTokenProfileArgs(
    address localToken
  ) internal pure returns (PolicyConfig.TokenProfileArgs memory) {
    return PolicyConfig.TokenProfileArgs({
      localToken: localToken,
      localDecimals: DEFAULT_TOKEN_DECIMALS,
      lockOrBurnFlavor: IPoolConfigManager.LockOrBurnFlavor.Burn,
      allowedFinalityConfig: FinalityCodec.WAIT_FOR_FINALITY_FLAG,
      lockOrBurnAddress: address(0),
      additionalPoolDataProvider: address(0)
    });
  }

  /// @dev A remote pool record that trusts the remote pool and sends extended pool data.
  function _getRemotePoolConfig() internal view returns (ConnectionPolicyStore.RemotePoolConfig memory config) {
    IPoolConfigManager.RemotePool[] memory remotePools = new IPoolConfigManager.RemotePool[](1);
    remotePools[0] =
      IPoolConfigManager.RemotePool({remotePoolAddress: abi.encode(s_remotePool), allowLegacyPoolDataOnReceive: false});
    config.remotePools = remotePools;
    return config;
  }

  function _getRemotePoolAddresses(
    bytes memory remotePool
  ) internal pure returns (bytes[] memory remotePoolAddresses) {
    remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = remotePool;
    return remotePoolAddresses;
  }

  /// @dev An enabled limit with a rate of one token unit per second.
  function _getRateLimitConfig(
    uint128 capacity
  ) internal pure returns (RateLimiter.Config memory) {
    return RateLimiter.Config({isEnabled: true, capacity: capacity, rate: 1});
  }

  function _getFeeConfig(
    uint32 finalityFeeUSDCents,
    uint16 finalityTransferFeeBps
  ) internal pure returns (FeePolicyStore.FeeConfig memory) {
    return FeePolicyStore.FeeConfig({
      finalityFeeUSDCents: finalityFeeUSDCents,
      fastFinalityFeeUSDCents: finalityFeeUSDCents * 2,
      finalityTransferFeeBps: finalityTransferFeeBps,
      fastFinalityTransferFeeBps: finalityTransferFeeBps * 2
    });
  }

  function _getFastFinalityConfig() internal pure returns (bytes4) {
    return FinalityCodec._encodeBlockDepth(10);
  }

  // ================================================================
  // │                       Config helpers                         │
  // ================================================================

  function _deployToken(
    string memory tokenName
  ) internal returns (address) {
    return address(_deployCrossChainToken(tokenName, DEFAULT_TOKEN_DECIMALS));
  }

  /// @dev Adds the policy records of the remote pool, rate limit, fee, destination overhead, and security types. The
  /// remote pool policy has no default record, so its record is on the destination chain. The rate limits are disabled,
  /// the fees are zero, the overheads are the default gas and the minimum bytes, and the security record is empty.
  /// The first record creates each policy.
  function _addPolicyDefaultRecords(
    bytes32 policyName
  ) internal {
    uint64 defaultChainSelector = PolicyConfig.DEFAULT_CHAIN_SELECTOR;
    _setRemotePoolPolicyConfig(policyName, DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(abi.encode(s_remotePool)));
    RateLimiter.Config memory disabled;
    _setRateLimitPolicyConfig(policyName, defaultChainSelector, disabled);
    FeePolicyStore.FeeConfig memory zeroFee;
    _setFeePolicyConfig(policyName, defaultChainSelector, zeroFee);
    _setDestOverheadPolicyConfig(
      policyName, defaultChainSelector, DEFAULT_TOKEN_DEST_GAS_OVERHEAD, MIN_DEST_BYTES_OVERHEAD
    );
    _setSecurityPolicyConfig(policyName, defaultChainSelector, new address[](0), new address[](0), address(0));
  }

  /// @dev Creates one policy of each type, named after its type, with no tokens. The connection policy has a record
  /// for the other chain. The other policies have the records of the DEFAULT policies.
  function _addPolicyOfEachType() internal {
    RateLimiter.Config memory disabledRateLimitConfig;
    FeePolicyStore.FeeConfig memory zeroFeeConfig;
    uint64 defaultChainSelector = PolicyConfig.DEFAULT_CHAIN_SELECTOR;
    _setConnectionPolicyConfig(CONNECTION_POLICY, OTHER_CHAIN_SELECTOR, false);
    _setRemotePoolPolicyConfig(
      REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(abi.encode(s_remotePool))
    );
    _setRateLimitPolicyConfig(RATE_LIMIT_POLICY, defaultChainSelector, disabledRateLimitConfig);
    _setFeePolicyConfig(FEE_POLICY, defaultChainSelector, zeroFeeConfig);
    _setDestOverheadPolicyConfig(
      DEST_OVERHEAD_POLICY, defaultChainSelector, DEFAULT_TOKEN_DEST_GAS_OVERHEAD, MIN_DEST_BYTES_OVERHEAD
    );
    _setSecurityPolicyConfig(SECURITY_POLICY, defaultChainSelector, new address[](0), new address[](0), address(0));
  }

  /// @dev Sets the profile of one token.
  function _setTokenProfile(
    PolicyConfig.TokenProfileArgs memory profile
  ) internal {
    PolicyConfig.TokenProfileArgs[] memory profiles = new PolicyConfig.TokenProfileArgs[](1);
    profiles[0] = profile;
    s_poolConfigManager.applyTokenProfileUpdates(new address[](0), profiles);
  }

  /// @dev Registers a token with the default profile and policies. Assigns the connection policy last.
  function _registerToken(
    address localToken
  ) internal {
    _setTokenProfile(_getTokenProfileArgs(localToken));
    _setTokenRemotePoolPolicy(localToken, DEFAULT_POLICY);
    _setTokenRateLimitPolicy(localToken, DEFAULT_POLICY);
    _setTokenFeePolicy(localToken, DEFAULT_POLICY);
    _setTokenDestOverheadPolicy(localToken, DEFAULT_POLICY);
    _setTokenSecurityPolicy(localToken, DEFAULT_POLICY);
    _setTokenConnectionPolicy(localToken, DEFAULT_POLICY);
  }

  function _setTokenConnectionPolicy(
    address localToken,
    bytes32 policyName
  ) internal {
    s_poolConfigManager.applyTokenConnectionPolicyUpdates(_tokenPoliciesOf(localToken, policyName));
  }

  function _setTokenRemotePoolPolicy(
    address localToken,
    bytes32 policyName
  ) internal {
    s_poolConfigManager.applyTokenRemotePoolPolicyUpdates(_tokenPoliciesOf(localToken, policyName));
  }

  function _setTokenRateLimitPolicy(
    address localToken,
    bytes32 policyName
  ) internal {
    s_poolConfigManager.applyTokenRateLimitPolicyUpdates(_tokenPoliciesOf(localToken, policyName));
  }

  function _setTokenFeePolicy(
    address localToken,
    bytes32 policyName
  ) internal {
    s_poolConfigManager.applyTokenFeePolicyUpdates(_tokenPoliciesOf(localToken, policyName));
  }

  function _setTokenDestOverheadPolicy(
    address localToken,
    bytes32 policyName
  ) internal {
    s_poolConfigManager.applyTokenDestOverheadPolicyUpdates(_tokenPoliciesOf(localToken, policyName));
  }

  function _setTokenSecurityPolicy(
    address localToken,
    bytes32 policyName
  ) internal {
    s_poolConfigManager.applyTokenSecurityPolicyUpdates(_tokenPoliciesOf(localToken, policyName));
  }

  /// @dev Returns the registered token profile as arguments for applyTokenProfileUpdates.
  function _getStoredTokenProfile(
    address localToken
  ) internal view returns (PolicyConfig.TokenProfileArgs memory tokenProfile) {
    ConnectionPolicyStore.TokenProfile memory profile = s_connectionStore.getTokenProfiles(_arrayOf(localToken))[0];
    tokenProfile.localToken = localToken;
    tokenProfile.localDecimals = profile.localDecimals;
    tokenProfile.lockOrBurnFlavor = profile.lockOrBurnFlavor;
    tokenProfile.allowedFinalityConfig = profile.allowedFinalityConfig;
    tokenProfile.lockOrBurnAddress = profile.lockOrBurnAddress;
    tokenProfile.additionalPoolDataProvider = profile.additionalPoolDataProvider;
    return tokenProfile;
  }

  /// @dev Sets the allowed finality and keeps the other values of the token profile.
  function _setAllowedFinalityConfig(
    address localToken,
    bytes4 allowedFinalityConfig
  ) internal {
    PolicyConfig.TokenProfileArgs memory tokenProfile = _getStoredTokenProfile(localToken);
    tokenProfile.allowedFinalityConfig = allowedFinalityConfig;
    _setTokenProfile(tokenProfile);
  }

  // ================================================================
  // │                       Record helpers                         │
  // ================================================================

  // Each helper makes one write through the pool config manager, with empty lists for the other arguments.

  function _setConnectionPolicyConfig(
    bytes32 policyName,
    uint64 remoteChainSelector,
    bool onlySubsetTokens
  ) internal {
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs =
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](1);
    policyConfigs[0] = ConnectionPolicyStore.ConnectionPolicyConfigArgs({
      policyName: policyName, remoteChainSelector: remoteChainSelector, onlySubsetTokens: onlySubsetTokens
    });
    s_poolConfigManager.applyConnectionPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  /// @dev Sets the remote token address provider of an existing connection policy. Zero removes it.
  function _setRemoteTokenAddressProvider(
    bytes32 policyName,
    address remoteTokenAddressProvider
  ) internal {
    ConnectionPolicyStore.RemoteTokenAddressProviderArgs[] memory providers =
      new ConnectionPolicyStore.RemoteTokenAddressProviderArgs[](1);
    providers[0] = ConnectionPolicyStore.RemoteTokenAddressProviderArgs({
      policyName: policyName, remoteTokenAddressProvider: remoteTokenAddressProvider
    });
    s_poolConfigManager.applyRemoteTokenAddressProviderUpdates(providers);
  }

  function _setInlineConnectionConfig(
    address localToken,
    uint64 remoteChainSelector,
    bytes memory remoteToken
  ) internal {
    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory inlineConfigs =
      new ConnectionPolicyStore.InlineConnectionConfigArgs[](1);
    inlineConfigs[0] = ConnectionPolicyStore.InlineConnectionConfigArgs({
      localToken: localToken, remoteChainSelector: remoteChainSelector, remoteToken: remoteToken
    });
    s_poolConfigManager.applyInlineConnectionConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function _setRemotePoolPolicyConfig(
    bytes32 policyName,
    uint64 remoteChainSelector,
    bytes[] memory remotePoolAddresses
  ) internal {
    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory policyConfigs =
      new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](1);
    policyConfigs[0] = ConnectionPolicyStore.RemotePoolPolicyConfigArgs({
      policyName: policyName, remoteChainSelector: remoteChainSelector, remotePoolAddresses: remotePoolAddresses
    });
    s_poolConfigManager.applyRemotePoolPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  function _setInlineRemotePoolConfig(
    address localToken,
    uint64 remoteChainSelector,
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig
  ) internal {
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      new ConnectionPolicyStore.InlineRemotePoolConfigArgs[](1);
    inlineConfigs[0] = ConnectionPolicyStore.InlineRemotePoolConfigArgs({
      localToken: localToken,
      remoteChainSelector: remoteChainSelector,
      trustPolicyRemotePools: false,
      remotePoolConfig: remotePoolConfig
    });
    s_poolConfigManager.applyInlineRemotePoolConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  /// @dev Sets the same limit in both directions for both finality classes. The policy has the decimals of the test
  /// tokens.
  function _setRateLimitPolicyConfig(
    bytes32 policyName,
    uint64 remoteChainSelector,
    RateLimiter.Config memory config
  ) internal {
    SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory policyConfigs =
      new SecurityPolicyStore.RateLimitPolicyConfigArgs[](1);
    policyConfigs[0] = SecurityPolicyStore.RateLimitPolicyConfigArgs({
      policyName: policyName,
      remoteChainSelector: remoteChainSelector,
      localDecimals: DEFAULT_TOKEN_DECIMALS,
      rateLimitConfig: SecurityPolicyStore.RateLimitConfig({
        outbound: config, inbound: config, fastFinalityOutbound: config, fastFinalityInbound: config
      })
    });
    s_poolConfigManager.applyRateLimitPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  /// @dev Sets the same limit in both directions for both finality classes.
  function _setInlineRateLimitConfig(
    address localToken,
    uint64 remoteChainSelector,
    RateLimiter.Config memory config
  ) internal {
    SecurityPolicyStore.InlineRateLimitConfigArgs[] memory inlineConfigs =
      new SecurityPolicyStore.InlineRateLimitConfigArgs[](1);
    inlineConfigs[0] = SecurityPolicyStore.InlineRateLimitConfigArgs({
      localToken: localToken,
      remoteChainSelector: remoteChainSelector,
      rateLimitConfig: SecurityPolicyStore.RateLimitConfig({
        outbound: config, inbound: config, fastFinalityOutbound: config, fastFinalityInbound: config
      })
    });
    s_poolConfigManager.applyInlineRateLimitConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function _setFeePolicyConfig(
    bytes32 policyName,
    uint64 remoteChainSelector,
    FeePolicyStore.FeeConfig memory feeConfig
  ) internal {
    FeePolicyStore.FeePolicyConfigArgs[] memory policyConfigs = new FeePolicyStore.FeePolicyConfigArgs[](1);
    policyConfigs[0] = FeePolicyStore.FeePolicyConfigArgs({
      policyName: policyName, remoteChainSelector: remoteChainSelector, feeConfig: feeConfig
    });
    s_poolConfigManager.applyFeePolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  function _setInlineFeeConfig(
    address localToken,
    uint64 remoteChainSelector,
    FeePolicyStore.FeeConfig memory feeConfig
  ) internal {
    FeePolicyStore.InlineFeeConfigArgs[] memory inlineConfigs = new FeePolicyStore.InlineFeeConfigArgs[](1);
    inlineConfigs[0] = FeePolicyStore.InlineFeeConfigArgs({
      localToken: localToken, remoteChainSelector: remoteChainSelector, feeConfig: feeConfig
    });
    s_poolConfigManager.applyInlineFeeConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function _setDestOverheadPolicyConfig(
    bytes32 policyName,
    uint64 remoteChainSelector,
    uint32 destGasOverhead,
    uint32 destBytesOverhead
  ) internal {
    FeePolicyStore.DestOverheadPolicyConfigArgs[] memory policyConfigs =
      new FeePolicyStore.DestOverheadPolicyConfigArgs[](1);
    policyConfigs[0] = FeePolicyStore.DestOverheadPolicyConfigArgs({
      policyName: policyName,
      remoteChainSelector: remoteChainSelector,
      destOverheadConfig: FeePolicyStore.DestOverheadConfig({
        destGasOverhead: destGasOverhead, destBytesOverhead: destBytesOverhead
      })
    });
    s_poolConfigManager.applyDestOverheadPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  function _setInlineDestOverheadConfig(
    address localToken,
    uint64 remoteChainSelector,
    uint32 destGasOverhead,
    uint32 destBytesOverhead
  ) internal {
    FeePolicyStore.InlineDestOverheadConfigArgs[] memory inlineConfigs =
      new FeePolicyStore.InlineDestOverheadConfigArgs[](1);
    inlineConfigs[0] = FeePolicyStore.InlineDestOverheadConfigArgs({
      localToken: localToken,
      remoteChainSelector: remoteChainSelector,
      destOverheadConfig: FeePolicyStore.DestOverheadConfig({
        destGasOverhead: destGasOverhead, destBytesOverhead: destBytesOverhead
      })
    });
    s_poolConfigManager.applyInlineDestOverheadConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function _setSecurityPolicyConfig(
    bytes32 policyName,
    uint64 remoteChainSelector,
    address[] memory outboundCCVs,
    address[] memory inboundCCVs,
    address hook
  ) internal {
    SecurityPolicyStore.SecurityPolicyConfigArgs[] memory policyConfigs =
      new SecurityPolicyStore.SecurityPolicyConfigArgs[](1);
    policyConfigs[0] = SecurityPolicyStore.SecurityPolicyConfigArgs({
      policyName: policyName,
      remoteChainSelector: remoteChainSelector,
      securityConfig: SecurityPolicyStore.SecurityConfig({
        outboundCCVs: outboundCCVs, inboundCCVs: inboundCCVs, hook: hook
      })
    });
    s_poolConfigManager.applySecurityPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  function _setInlineSecurityConfig(
    address localToken,
    uint64 remoteChainSelector,
    address[] memory outboundCCVs,
    address[] memory inboundCCVs,
    address hook
  ) internal {
    SecurityPolicyStore.InlineSecurityConfigArgs[] memory inlineConfigs =
      new SecurityPolicyStore.InlineSecurityConfigArgs[](1);
    inlineConfigs[0] = SecurityPolicyStore.InlineSecurityConfigArgs({
      localToken: localToken,
      remoteChainSelector: remoteChainSelector,
      securityConfig: SecurityPolicyStore.SecurityConfig({
        outboundCCVs: outboundCCVs, inboundCCVs: inboundCCVs, hook: hook
      })
    });
    s_poolConfigManager.applyInlineSecurityConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function _policyRecordKeysOf(
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal pure returns (PolicyConfig.PolicyRecordKey[] memory keys) {
    keys = new PolicyConfig.PolicyRecordKey[](1);
    keys[0] = PolicyConfig.PolicyRecordKey({policyName: policyName, remoteChainSelector: remoteChainSelector});
    return keys;
  }

  function _policyNamesOf(
    bytes32 policyName
  ) internal pure returns (bytes32[] memory policyNames) {
    policyNames = new bytes32[](1);
    policyNames[0] = policyName;
    return policyNames;
  }

  function _tokenPoliciesOf(
    address localToken,
    bytes32 policyName
  ) internal pure returns (PolicyConfig.TokenPolicyArgs[] memory tokenPolicies) {
    tokenPolicies = new PolicyConfig.TokenPolicyArgs[](1);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: localToken, policyName: policyName});
    return tokenPolicies;
  }

  // ================================================================
  // │                         View helpers                         │
  // ================================================================

  /// @dev The pool config manager views take a batch. Most tests read one lane or one token.
  function _getLaneConnection(
    address localToken,
    uint64 remoteChainSelector
  ) internal view returns (IPoolConfigManager.LaneConnection memory) {
    return s_poolConfigManager.getLaneConnectionBatch(_lanesOf(localToken, remoteChainSelector))[0];
  }

  function _getTokenConfig(
    address localToken
  ) internal view returns (IPoolConfigManager.TokenConfig memory) {
    return s_poolConfigManager.getTokenConfigBatch(_arrayOf(localToken))[0];
  }

  function _lanesOf(
    address localToken,
    uint64 remoteChainSelector
  ) internal pure returns (IPoolConfigManager.LaneKey[] memory lanes) {
    lanes = new IPoolConfigManager.LaneKey[](1);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: localToken, remoteChainSelector: remoteChainSelector});
    return lanes;
  }

  /// @dev The four buckets of the token on the chain. All zero for a token without a rate limit record.
  function _getRateLimiterState(
    address localToken,
    uint64 remoteChainSelector
  ) internal view returns (IPoolConfigManager.LaneRateLimiterState memory) {
    return s_securityStore.getCurrentRateLimiterStateBatch(_lanesOf(localToken, remoteChainSelector))[0];
  }

  function _setCaller(
    address caller
  ) internal {
    vm.stopPrank();
    vm.startPrank(caller);
  }
}
