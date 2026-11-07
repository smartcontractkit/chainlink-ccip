// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IMultiTokenPool} from "../../interfaces/IMultiTokenPool.sol";
import {IPoolConfigManager} from "../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../interfaces/IPoolV2.sol";
import {ITokenAdminRegistry} from "../../interfaces/ITokenAdminRegistry.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {PolicyConfig} from "../../libraries/PolicyConfig.sol";
import {Pool} from "../../libraries/Pool.sol";
import {ConnectionPolicyStore} from "./policies/ConnectionPolicyStore.sol";
import {FeePolicyStore} from "./policies/FeePolicyStore.sol";
import {SecurityPolicyStore} from "./policies/SecurityPolicyStore.sol";

import {
  AccessControlDefaultAdminRules
} from "@openzeppelin/contracts@5.3.0/access/extensions/AccessControlDefaultAdminRules.sol";
import {Address} from "@openzeppelin/contracts@5.3.0/utils/Address.sol";

/// @notice Configures tokens and resolves token configuration for multi-token pools.
/// @dev Token configuration stays in the policy stores when this manager is replaced.
contract PoolConfigManager is IPoolConfigManager, ITypeAndVersion, AccessControlDefaultAdminRules {
  error ZeroAddressInvalid();
  error CallerIsNotTokenPool(address caller, address localToken);
  error PoolConfigManagerMismatch(address pool, address poolConfigManager);

  event StaticConfigSet(
    ITokenAdminRegistry tokenAdminRegistry,
    ConnectionPolicyStore connectionStore,
    FeePolicyStore feeStore,
    SecurityPolicyStore securityStore
  );

  // ================================================================
  // │                            Roles                             │
  // ================================================================

  // AccessControlDefaultAdminRules manages DEFAULT_ADMIN_ROLE, admin transfers, and role administration.

  /// @notice Can update all token and policy configuration. Cannot pause, unpause, or withdraw fees.
  bytes32 public constant CONFIG_ROLE = keccak256("CONFIG_ROLE");
  /// @notice Can update policies of every type, including their remote token address providers, subset tokens, and
  /// individual bucket configs.
  bytes32 public constant POLICY_ROLE = keccak256("POLICY_ROLE");
  /// @notice Can make every connection write that POLICY_ROLE or TOKEN_MANAGER_ROLE can make.
  bytes32 public constant CONNECTION_ROLE = keccak256("CONNECTION_ROLE");
  /// @notice Can make every remote pool write that POLICY_ROLE or TOKEN_MANAGER_ROLE can make.
  bytes32 public constant REMOTE_POOL_ROLE = keccak256("REMOTE_POOL_ROLE");
  /// @notice Can make every rate limit write that POLICY_ROLE or TOKEN_MANAGER_ROLE can make.
  bytes32 public constant RATE_LIMIT_ROLE = keccak256("RATE_LIMIT_ROLE");
  /// @notice Can make every fee write that POLICY_ROLE or TOKEN_MANAGER_ROLE can make.
  bytes32 public constant FEE_ROLE = keccak256("FEE_ROLE");
  /// @notice Can make every destination overhead write that POLICY_ROLE or TOKEN_MANAGER_ROLE can make.
  bytes32 public constant DEST_OVERHEAD_ROLE = keccak256("DEST_OVERHEAD_ROLE");
  /// @notice Can make every security write that POLICY_ROLE or TOKEN_MANAGER_ROLE can make.
  bytes32 public constant SECURITY_ROLE = keccak256("SECURITY_ROLE");
  /// @notice Registers and deregisters tokens. Sets their profiles, policy assignments, and inline records.
  /// Can also remove inline records and update individual inline bucket configs.
  bytes32 public constant TOKEN_MANAGER_ROLE = keccak256("TOKEN_MANAGER_ROLE");
  /// @notice Sets token, chain, and global pause flags.
  bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
  /// @notice Clears token, chain, and global pause flags.
  bytes32 public constant UNPAUSER_ROLE = keccak256("UNPAUSER_ROLE");
  /// @notice Withdraws token fees from pools.
  bytes32 public constant FEE_COLLECTOR_ROLE = keccak256("FEE_COLLECTOR_ROLE");

  // ================================================================
  // │                        Static config                         │
  // ================================================================

  /// @dev Identifies the pool that can validate transfers for each token.
  ITokenAdminRegistry private immutable i_tokenAdminRegistry;
  /// @dev Stores connection and remote pool config, token profiles, and pause state.
  ConnectionPolicyStore private immutable i_connectionStore;
  /// @dev Stores fee and destination overhead config.
  FeePolicyStore private immutable i_feeStore;
  /// @dev Stores rate limit and security config, including live bucket state.
  SecurityPolicyStore private immutable i_securityStore;

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "PoolConfigManager 2.1.0-dev";
  }

  constructor(
    ITokenAdminRegistry tokenAdminRegistry,
    ConnectionPolicyStore connectionStore,
    FeePolicyStore feeStore,
    SecurityPolicyStore securityStore,
    address admin
  ) AccessControlDefaultAdminRules(0, admin) {
    if (
      address(tokenAdminRegistry) == address(0) || address(connectionStore) == address(0)
        || address(feeStore) == address(0) || address(securityStore) == address(0)
    ) {
      revert ZeroAddressInvalid();
    }
    i_tokenAdminRegistry = tokenAdminRegistry;
    i_connectionStore = connectionStore;
    i_feeStore = feeStore;
    i_securityStore = securityStore;
    emit StaticConfigSet(tokenAdminRegistry, connectionStore, feeStore, securityStore);
  }

  /// @notice Returns the static config of this pool config manager.
  /// @return tokenAdminRegistry The registry that assigns a pool to each token.
  /// @return connectionStore The store for connection and remote pool config, token profiles, and pause state.
  /// @return feeStore The store for fee and destination overhead config.
  /// @return securityStore The store for rate limit and security config, including live bucket state.
  function getStaticConfig()
    external
    view
    returns (
      ITokenAdminRegistry tokenAdminRegistry,
      ConnectionPolicyStore connectionStore,
      FeePolicyStore feeStore,
      SecurityPolicyStore securityStore
    )
  {
    return (i_tokenAdminRegistry, i_connectionStore, i_feeStore, i_securityStore);
  }

  // ================================================================
  // │                     Transfer validation                      │
  // ================================================================

  /// @inheritdoc IPoolConfigManager
  function validateLockOrBurn(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount,
    bytes4 requestedFinalityConfig,
    bytes calldata tokenArgs,
    bool chargeTokenFee
  ) external override returns (OutboundLane memory lane) {
    _onlyTokenPool(localToken);

    lane = i_connectionStore.resolveOutboundLane(localToken, remoteChainSelector, requestedFinalityConfig, tokenArgs);
    // The fee store checks the pool data length against destBytesOverhead, so it needs the pool data first.
    lane.feeAmount = i_feeStore.validateLockOrBurn(
      localToken, remoteChainSelector, amount, requestedFinalityConfig, chargeTokenFee, lane.destPoolData.length
    );
    // The rate limit applies to the amount after the token fee.
    lane.hook = i_securityStore.validateLockOrBurn(
      localToken, remoteChainSelector, amount - lane.feeAmount, requestedFinalityConfig
    );
    return lane;
  }

  /// @inheritdoc IPoolConfigManager
  function validateReleaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    bytes4 requestedFinalityConfig
  ) external override returns (ResolvedInboundTransfer memory transfer) {
    address localToken = releaseOrMintIn.localToken;
    uint64 remoteChainSelector = releaseOrMintIn.remoteChainSelector;

    _onlyTokenPool(localToken);
    transfer = i_connectionStore.resolveInboundLane(
      localToken,
      remoteChainSelector,
      releaseOrMintIn.sourcePoolAddress,
      releaseOrMintIn.sourcePoolData,
      releaseOrMintIn.sourceDenominatedAmount,
      requestedFinalityConfig
    );
    // The rate limit applies to the amount in local decimals.
    transfer.hook = i_securityStore.validateReleaseOrMint(
      localToken, remoteChainSelector, transfer.localAmount, requestedFinalityConfig
    );

    return transfer;
  }

  // ================================================================
  // │                            Views                             │
  // ================================================================

  // Batch getters do not need pages. Callers can split token and lane inputs into smaller batches.
  // Only the connection store knows whether a lane exists. Getters that read the other stores zero out the entries
  // for lanes that don't exist.

  /// @inheritdoc IPoolConfigManager
  function getFeeBatch(
    LaneKey[] calldata lanes,
    bytes4 requestedFinalityConfig
  ) external view override returns (LaneFee[] memory fees) {
    i_connectionStore.checkRequestedFinalityBatch(lanes, requestedFinalityConfig);
    bool[] memory hasLanes = i_connectionStore.hasLaneBatch(lanes);
    fees = i_feeStore.getFeeBatch(lanes, requestedFinalityConfig);

    for (uint256 i = 0; i < lanes.length; ++i) {
      if (!hasLanes[i]) {
        delete fees[i]; // Sets isEnabled to false.
      }
    }

    return fees;
  }

  /// @inheritdoc IPoolConfigManager
  function getTokenTransferFeeConfigBatch(
    LaneKey[] calldata lanes
  ) external view override returns (IPoolV2.TokenTransferFeeConfig[] memory feeConfigs) {
    bool[] memory hasLanes = i_connectionStore.hasLaneBatch(lanes);
    feeConfigs = i_feeStore.getTokenTransferFeeConfigBatch(lanes);

    for (uint256 i = 0; i < lanes.length; ++i) {
      if (!hasLanes[i]) {
        delete feeConfigs[i];
      }
    }

    return feeConfigs;
  }

  /// @inheritdoc IPoolConfigManager
  function getCurrentRateLimiterStateBatch(
    LaneKey[] calldata lanes
  ) external view override returns (LaneRateLimiterState[] memory states) {
    bool[] memory hasLanes = i_connectionStore.hasLaneBatch(lanes);
    states = i_securityStore.getCurrentRateLimiterStateBatch(lanes);

    for (uint256 i = 0; i < lanes.length; ++i) {
      if (!hasLanes[i]) {
        delete states[i];
      }
    }

    return states;
  }

  /// @inheritdoc IPoolConfigManager
  function getLaneConnectionBatch(
    LaneKey[] calldata lanes
  ) external view override returns (LaneConnection[] memory connections) {
    return i_connectionStore.getLaneConnectionBatch(lanes);
  }

  /// @inheritdoc IPoolConfigManager
  /// @dev The requested finality is ignored, as the hook and CCVs are the same for all finalities.
  function getLaneVerificationBatch(
    LaneKey[] calldata lanes,
    bytes4 // requestedFinalityConfig
  )
    external
    view
    override
    returns (address[] memory hooks, address[][] memory outboundCCVs, address[][] memory inboundCCVs)
  {
    bool[] memory hasLanes = i_connectionStore.hasLaneBatch(lanes);
    (hooks, outboundCCVs, inboundCCVs) = i_securityStore.getLaneVerificationBatch(lanes);

    for (uint256 i = 0; i < lanes.length; ++i) {
      if (!hasLanes[i]) {
        delete hooks[i];
        delete outboundCCVs[i];
        delete inboundCCVs[i];
      }
    }

    return (hooks, outboundCCVs, inboundCCVs);
  }

  /// @inheritdoc IPoolConfigManager
  function getTokenConfigBatch(
    address[] calldata localTokens
  ) external view override returns (TokenConfig[] memory configs) {
    // The connection store returns the profile, lanes, and connection and remote pool policy names. The policy names
    // of the other stores are added below.
    configs = i_connectionStore.getTokenConfigBatch(localTokens);

    (bytes32[] memory feePolicyNames, bytes32[] memory destOverheadPolicyNames) =
      i_feeStore.getTokenPolicyNamesBatch(localTokens);

    (bytes32[] memory rateLimitPolicyNames, bytes32[] memory securityPolicyNames) =
      i_securityStore.getTokenPolicyNamesBatch(localTokens);

    for (uint256 i = 0; i < localTokens.length; ++i) {
      configs[i].rateLimitPolicyName = rateLimitPolicyNames[i];
      configs[i].feePolicyName = feePolicyNames[i];
      configs[i].destOverheadPolicyName = destOverheadPolicyNames[i];
      configs[i].securityPolicyName = securityPolicyNames[i];
    }
    return configs;
  }

  /// @inheritdoc IPoolConfigManager
  function getAllConfiguredTokens(
    uint64 startIndex,
    uint64 maxCount
  ) external view override returns (address[] memory tokens) {
    return i_connectionStore.getAllConfiguredTokens(startIndex, maxCount);
  }

  /// @inheritdoc IPoolConfigManager
  function getTokenAdminRegistry() external view override returns (address tokenAdminRegistry) {
    return address(i_tokenAdminRegistry);
  }

  /// @notice Returns the record level that each config type would use to resolve a transfer on each lane.
  /// @dev Each config type is resolved on its own, including for lanes that do not exist yet, so tooling can check the
  /// records before opening a lane. The connection level shows whether the lane exists. Connection and remote pool
  /// policies have no default record, so their level is never PolicyDefault. An inline remote pool record reports
  /// Inline, even if it also trusts policy pools. Chain selector zero returns None for every type. All returned arrays
  /// are in the same order as lanes.
  /// @param lanes The token and remote chain lanes to query.
  /// @return connectionLevels The level of the connection record.
  /// @return remotePoolLevels The level of the remote pool record.
  /// @return rateLimitLevels The level of the rate limit record.
  /// @return feeLevels The level of the fee record.
  /// @return destOverheadLevels The level of the destination overhead record.
  /// @return securityLevels The level of the security record.
  function getLaneRecordLevelBatch(
    LaneKey[] calldata lanes
  )
    external
    view
    returns (
      PolicyConfig.RecordLevel[] memory connectionLevels,
      PolicyConfig.RecordLevel[] memory remotePoolLevels,
      PolicyConfig.RecordLevel[] memory rateLimitLevels,
      PolicyConfig.RecordLevel[] memory feeLevels,
      PolicyConfig.RecordLevel[] memory destOverheadLevels,
      PolicyConfig.RecordLevel[] memory securityLevels
    )
  {
    (connectionLevels, remotePoolLevels) = i_connectionStore.getLaneRecordLevelBatch(lanes);
    (feeLevels, destOverheadLevels) = i_feeStore.getLaneRecordLevelBatch(lanes);
    (rateLimitLevels, securityLevels) = i_securityStore.getLaneRecordLevelBatch(lanes);
    return (connectionLevels, remotePoolLevels, rateLimitLevels, feeLevels, destOverheadLevels, securityLevels);
  }

  // Forwarded views use the same parameter and return types as the store functions. Names are omitted because
  // _forwardView passes the call data and return data through unchanged. See each store function for return details.

  /// @notice Returns the subset restriction and one page of tokens for a connection policy record.
  /// @dev See ConnectionPolicyStore.getSubsetTokens.
  function getSubsetTokens(
    bytes32, // policyName
    uint64, // remoteChainSelector
    uint64, // startIndex
    uint64 // maxCount
  ) external view returns (bool, address[] memory) {
    _forwardView(address(i_connectionStore));
  }

  // Policy and inline config getters do not use pages. Each policy or token is expected to cover a small
  // number of chains. Remote pool and CCV lists per chain are also expected to be small.

  /// @notice Returns the token count, remote token address provider, and records of a connection policy.
  /// @dev See ConnectionPolicyStore.getConnectionPolicy.
  function getConnectionPolicy(
    bytes32 // policyName
  ) external view returns (uint32, address, ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory) {
    _forwardView(address(i_connectionStore));
  }

  /// @notice Returns the token count and records of a remote pool policy.
  /// @dev See ConnectionPolicyStore.getRemotePoolPolicy.
  function getRemotePoolPolicy(
    bytes32 // policyName
  ) external view returns (uint32, ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory) {
    _forwardView(address(i_connectionStore));
  }

  /// @notice Returns the token count, decimals, and records of a rate limit policy.
  /// @dev See SecurityPolicyStore.getRateLimitPolicy.
  function getRateLimitPolicy(
    bytes32 // policyName
  ) external view returns (uint32, uint8, SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory) {
    _forwardView(address(i_securityStore));
  }

  /// @notice Returns the token count and records of a fee policy.
  /// @dev See FeePolicyStore.getFeePolicy.
  function getFeePolicy(
    bytes32 // policyName
  ) external view returns (uint32, FeePolicyStore.FeePolicyConfigArgs[] memory) {
    _forwardView(address(i_feeStore));
  }

  /// @notice Returns the token count and records of a destination overhead policy.
  /// @dev See FeePolicyStore.getDestOverheadPolicy.
  function getDestOverheadPolicy(
    bytes32 // policyName
  ) external view returns (uint32, FeePolicyStore.DestOverheadPolicyConfigArgs[] memory) {
    _forwardView(address(i_feeStore));
  }

  /// @notice Returns the token count and records of a security policy.
  /// @dev See SecurityPolicyStore.getSecurityPolicy.
  function getSecurityPolicy(
    bytes32 // policyName
  ) external view returns (uint32, SecurityPolicyStore.SecurityPolicyConfigArgs[] memory) {
    _forwardView(address(i_securityStore));
  }

  // Inline record views include records on chains where the token has no lane.

  /// @notice Returns all inline connection records for a token.
  /// @dev See ConnectionPolicyStore.getInlineConnectionConfigs.
  function getInlineConnectionConfigs(
    address // localToken
  ) external view returns (ConnectionPolicyStore.InlineConnectionConfigArgs[] memory) {
    _forwardView(address(i_connectionStore));
  }

  /// @notice Returns all inline remote pool records for a token.
  /// @dev See ConnectionPolicyStore.getInlineRemotePoolConfigs.
  function getInlineRemotePoolConfigs(
    address // localToken
  ) external view returns (ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory) {
    _forwardView(address(i_connectionStore));
  }

  /// @notice Returns all inline rate limit records for a token, including each bucket's config.
  /// @dev See SecurityPolicyStore.getInlineRateLimitConfigs.
  function getInlineRateLimitConfigs(
    address // localToken
  ) external view returns (SecurityPolicyStore.InlineRateLimitConfigArgs[] memory) {
    _forwardView(address(i_securityStore));
  }

  /// @notice Returns all inline fee records for a token.
  /// @dev See FeePolicyStore.getInlineFeeConfigs.
  function getInlineFeeConfigs(
    address // localToken
  ) external view returns (FeePolicyStore.InlineFeeConfigArgs[] memory) {
    _forwardView(address(i_feeStore));
  }

  /// @notice Returns all inline destination overhead records for a token.
  /// @dev See FeePolicyStore.getInlineDestOverheadConfigs.
  function getInlineDestOverheadConfigs(
    address // localToken
  ) external view returns (FeePolicyStore.InlineDestOverheadConfigArgs[] memory) {
    _forwardView(address(i_feeStore));
  }

  /// @notice Returns all inline security records for a token.
  /// @dev See SecurityPolicyStore.getInlineSecurityConfigs.
  function getInlineSecurityConfigs(
    address // localToken
  ) external view returns (SecurityPolicyStore.InlineSecurityConfigArgs[] memory) {
    _forwardView(address(i_securityStore));
  }

  /// @notice Returns the names of all policies, grouped by type.
  /// @dev This getter does not use pages. Policies are expected to be shared by many tokens.
  /// Use inline records for settings that apply to one token. This keeps the policy list small.
  /// Includes policies with no assigned tokens or records.
  /// @return connectionPolicyNames The names of the connection policies.
  /// @return remotePoolPolicyNames The names of the remote pool policies.
  /// @return rateLimitPolicyNames The names of the rate limit policies.
  /// @return feePolicyNames The names of the fee policies.
  /// @return destOverheadPolicyNames The names of the destination overhead policies.
  /// @return securityPolicyNames The names of the security policies.
  function getPolicyNames()
    external
    view
    returns (
      bytes32[] memory connectionPolicyNames,
      bytes32[] memory remotePoolPolicyNames,
      bytes32[] memory rateLimitPolicyNames,
      bytes32[] memory feePolicyNames,
      bytes32[] memory destOverheadPolicyNames,
      bytes32[] memory securityPolicyNames
    )
  {
    (connectionPolicyNames, remotePoolPolicyNames) = i_connectionStore.getPolicyNames();
    (feePolicyNames, destOverheadPolicyNames) = i_feeStore.getPolicyNames();
    (rateLimitPolicyNames, securityPolicyNames) = i_securityStore.getPolicyNames();
    return (
      connectionPolicyNames,
      remotePoolPolicyNames,
      rateLimitPolicyNames,
      feePolicyNames,
      destOverheadPolicyNames,
      securityPolicyNames
    );
  }

  // ================================================================
  // │                         Config writes                        │
  // ================================================================

  // Config writes use the same signatures as the store functions so they can forward msg.data unchanged.
  // Address.functionCall propagates revert data. An empty revert becomes FailedCall.

  /// @notice Deregisters tokens, then sets token profiles in all three stores. See
  /// ConnectionPolicyStore.applyTokenProfileUpdates, FeePolicyStore.applyTokenProfileUpdates, and
  /// SecurityPolicyStore.applyTokenProfileUpdates.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE. Stores are called in connection, fee, and security order.
  /// A revert from any store reverts the whole call.
  function applyTokenProfileUpdates(
    address[] calldata, // tokensToRemove
    PolicyConfig.TokenProfileArgs[] calldata // profilesToSet
  ) external {
    _checkRoleOr(CONFIG_ROLE, TOKEN_MANAGER_ROLE);
    bytes memory callData = msg.data;
    Address.functionCall(address(i_connectionStore), callData);
    Address.functionCall(address(i_feeStore), callData);
    Address.functionCall(address(i_securityStore), callData);
  }

  /// @notice Assigns connection policies to tokens. See ConnectionPolicyStore.applyTokenConnectionPolicyUpdates.
  /// @dev Only callable by CONFIG_ROLE, TOKEN_MANAGER_ROLE, or CONNECTION_ROLE.
  /// Does not check the lanes opened by a new policy assignment.
  function applyTokenConnectionPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata // tokenPoliciesToSet
  ) external {
    _checkTokenRoles(CONNECTION_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
  }

  /// @notice Assigns remote pool policies to tokens. See ConnectionPolicyStore.applyTokenRemotePoolPolicyUpdates.
  /// @dev Only callable by CONFIG_ROLE, TOKEN_MANAGER_ROLE, or REMOTE_POOL_ROLE.
  function applyTokenRemotePoolPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata // tokenPoliciesToSet
  ) external {
    _checkTokenRoles(REMOTE_POOL_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
  }

  /// @notice Assigns rate limit policies to tokens. See SecurityPolicyStore.applyTokenRateLimitPolicyUpdates.
  /// @dev Only callable by CONFIG_ROLE, TOKEN_MANAGER_ROLE, or RATE_LIMIT_ROLE.
  function applyTokenRateLimitPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata // tokenPoliciesToSet
  ) external {
    _checkTokenRoles(RATE_LIMIT_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  /// @notice Assigns fee policies to tokens. See FeePolicyStore.applyTokenFeePolicyUpdates.
  /// @dev Only callable by CONFIG_ROLE, TOKEN_MANAGER_ROLE, or FEE_ROLE.
  function applyTokenFeePolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata // tokenPoliciesToSet
  ) external {
    _checkTokenRoles(FEE_ROLE);
    Address.functionCall(address(i_feeStore), msg.data);
  }

  /// @notice Assigns destination overhead policies to tokens.
  /// See FeePolicyStore.applyTokenDestOverheadPolicyUpdates.
  /// @dev Only callable by CONFIG_ROLE, TOKEN_MANAGER_ROLE, or DEST_OVERHEAD_ROLE.
  function applyTokenDestOverheadPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata // tokenPoliciesToSet
  ) external {
    _checkTokenRoles(DEST_OVERHEAD_ROLE);
    Address.functionCall(address(i_feeStore), msg.data);
  }

  /// @notice Assigns security policies to tokens. See SecurityPolicyStore.applyTokenSecurityPolicyUpdates.
  /// @dev Only callable by CONFIG_ROLE, TOKEN_MANAGER_ROLE, or SECURITY_ROLE.
  function applyTokenSecurityPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata // tokenPoliciesToSet
  ) external {
    _checkTokenRoles(SECURITY_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  /// @notice Removes and sets connection policy records, and removes connection policies. See
  /// ConnectionPolicyStore.applyConnectionPolicyConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE, or CONNECTION_ROLE. Does not check lanes opened by policy
  /// records because the stores do not track the list of tokens assigned to each policy.
  function applyConnectionPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata, // policyConfigsToRemove
    bytes32[] calldata, // policiesToRemove
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] calldata // policyConfigsToSet
  ) external {
    _checkPolicyRoles(CONNECTION_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
  }

  /// @notice Removes and sets inline connection records. See ConnectionPolicyStore.applyInlineConnectionConfigUpdates.
  /// Checks that each configured lane has a record of every other config type.
  /// @dev Only callable by CONFIG_ROLE, TOKEN_MANAGER_ROLE, or CONNECTION_ROLE. Reverts on the first missing record.
  /// @param inlineConfigsToSet The inline records to set.
  function applyInlineConnectionConfigUpdates(
    LaneKey[] calldata, // inlineConfigsToRemove
    ConnectionPolicyStore.InlineConnectionConfigArgs[] calldata inlineConfigsToSet
  ) external {
    _checkTokenRoles(CONNECTION_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
    for (uint256 i = 0; i < inlineConfigsToSet.length; ++i) {
      _checkLane(inlineConfigsToSet[i].localToken, inlineConfigsToSet[i].remoteChainSelector);
    }
  }

  /// @notice Removes and adds subset tokens. See ConnectionPolicyStore.applySubsetTokenUpdates.
  /// Checks that each added token has a record of every other config type for the chain.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE, or CONNECTION_ROLE. Reverts on the first missing record.
  /// The check runs even while onlySubsetTokens is false, so the lane is ready when the restriction is enabled.
  /// @param tokensToAdd The subset tokens to add.
  function applySubsetTokenUpdates(
    ConnectionPolicyStore.SubsetTokenArgs[] calldata, // tokensToRemove
    ConnectionPolicyStore.SubsetTokenArgs[] calldata tokensToAdd
  ) external {
    _checkPolicyRoles(CONNECTION_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
    for (uint256 i = 0; i < tokensToAdd.length; ++i) {
      _checkLane(tokensToAdd[i].localToken, tokensToAdd[i].remoteChainSelector);
    }
  }

  /// @notice Sets the remote token address providers of connection policies. See
  /// ConnectionPolicyStore.applyRemoteTokenAddressProviderUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE, or CONNECTION_ROLE. Changing a provider does not open a lane, so
  /// there is no lane check.
  function applyRemoteTokenAddressProviderUpdates(
    ConnectionPolicyStore.RemoteTokenAddressProviderArgs[] calldata // providersToSet
  ) external {
    _checkPolicyRoles(CONNECTION_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
  }

  /// @notice Removes and sets remote pool policy records, and removes remote pool policies. See
  /// ConnectionPolicyStore.applyRemotePoolPolicyConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE, or REMOTE_POOL_ROLE.
  function applyRemotePoolPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata, // policyConfigsToRemove
    bytes32[] calldata, // policiesToRemove
    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] calldata // policyConfigsToSet
  ) external {
    _checkPolicyRoles(REMOTE_POOL_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
  }

  /// @notice Removes and sets inline remote pool records. See ConnectionPolicyStore.applyInlineRemotePoolConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, TOKEN_MANAGER_ROLE, or REMOTE_POOL_ROLE.
  function applyInlineRemotePoolConfigUpdates(
    LaneKey[] calldata, // inlineConfigsToRemove
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] calldata // inlineConfigsToSet
  ) external {
    _checkTokenRoles(REMOTE_POOL_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
  }

  /// @notice Removes and sets rate limit policy records, and removes rate limit policies. See
  /// SecurityPolicyStore.applyRateLimitPolicyConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE, or RATE_LIMIT_ROLE.
  function applyRateLimitPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata, // policyConfigsToRemove
    bytes32[] calldata, // policiesToRemove
    SecurityPolicyStore.RateLimitPolicyConfigArgs[] calldata // policyConfigsToSet
  ) external {
    _checkPolicyRoles(RATE_LIMIT_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  /// @notice Updates individual bucket configs in existing rate limit policy records. See
  /// SecurityPolicyStore.applyRateLimitPolicyBucketUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE, or RATE_LIMIT_ROLE.
  function applyRateLimitPolicyBucketUpdates(
    SecurityPolicyStore.RateLimitPolicyBucketArgs[] calldata // bucketConfigsToSet
  ) external {
    _checkPolicyRoles(RATE_LIMIT_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  /// @notice Removes and sets inline rate limit records. See SecurityPolicyStore.applyInlineRateLimitConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, TOKEN_MANAGER_ROLE, or RATE_LIMIT_ROLE.
  function applyInlineRateLimitConfigUpdates(
    LaneKey[] calldata, // inlineConfigsToRemove
    SecurityPolicyStore.InlineRateLimitConfigArgs[] calldata // inlineConfigsToSet
  ) external {
    _checkTokenRoles(RATE_LIMIT_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  /// @notice Updates individual bucket configs in existing inline rate limit records. See
  /// SecurityPolicyStore.applyInlineRateLimitBucketUpdates.
  /// @dev Only callable by CONFIG_ROLE, TOKEN_MANAGER_ROLE, or RATE_LIMIT_ROLE.
  function applyInlineRateLimitBucketUpdates(
    SecurityPolicyStore.InlineRateLimitBucketArgs[] calldata // bucketConfigsToSet
  ) external {
    _checkTokenRoles(RATE_LIMIT_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  /// @notice Removes and sets fee policy records, and removes fee policies. See
  /// FeePolicyStore.applyFeePolicyConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE, or FEE_ROLE.
  function applyFeePolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata, // policyConfigsToRemove
    bytes32[] calldata, // policiesToRemove
    FeePolicyStore.FeePolicyConfigArgs[] calldata // policyConfigsToSet
  ) external {
    _checkPolicyRoles(FEE_ROLE);
    Address.functionCall(address(i_feeStore), msg.data);
  }

  /// @notice Removes and sets inline fee records. See FeePolicyStore.applyInlineFeeConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, TOKEN_MANAGER_ROLE, or FEE_ROLE.
  function applyInlineFeeConfigUpdates(
    LaneKey[] calldata, // inlineConfigsToRemove
    FeePolicyStore.InlineFeeConfigArgs[] calldata // inlineConfigsToSet
  ) external {
    _checkTokenRoles(FEE_ROLE);
    Address.functionCall(address(i_feeStore), msg.data);
  }

  /// @notice Removes and sets destination overhead policy records, and removes destination overhead policies. See
  /// FeePolicyStore.applyDestOverheadPolicyConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE, or DEST_OVERHEAD_ROLE.
  function applyDestOverheadPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata, // policyConfigsToRemove
    bytes32[] calldata, // policiesToRemove
    FeePolicyStore.DestOverheadPolicyConfigArgs[] calldata // policyConfigsToSet
  ) external {
    _checkPolicyRoles(DEST_OVERHEAD_ROLE);
    Address.functionCall(address(i_feeStore), msg.data);
  }

  /// @notice Removes and sets inline destination overhead records. See
  /// FeePolicyStore.applyInlineDestOverheadConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, TOKEN_MANAGER_ROLE, or DEST_OVERHEAD_ROLE.
  function applyInlineDestOverheadConfigUpdates(
    LaneKey[] calldata, // inlineConfigsToRemove
    FeePolicyStore.InlineDestOverheadConfigArgs[] calldata // inlineConfigsToSet
  ) external {
    _checkTokenRoles(DEST_OVERHEAD_ROLE);
    Address.functionCall(address(i_feeStore), msg.data);
  }

  /// @notice Removes and sets security policy records, and removes security policies. See
  /// SecurityPolicyStore.applySecurityPolicyConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE, or SECURITY_ROLE.
  function applySecurityPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata, // policyConfigsToRemove
    bytes32[] calldata, // policiesToRemove
    SecurityPolicyStore.SecurityPolicyConfigArgs[] calldata // policyConfigsToSet
  ) external {
    _checkPolicyRoles(SECURITY_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  /// @notice Removes and sets inline security records. See SecurityPolicyStore.applyInlineSecurityConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, TOKEN_MANAGER_ROLE, or SECURITY_ROLE.
  function applyInlineSecurityConfigUpdates(
    LaneKey[] calldata, // inlineConfigsToRemove
    SecurityPolicyStore.InlineSecurityConfigArgs[] calldata // inlineConfigsToSet
  ) external {
    _checkTokenRoles(SECURITY_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  // ================================================================
  // │                            Pause                             │
  // ================================================================

  /// @notice Turns on the global pause. It stops every token on every chain in both directions. The pause flags of
  /// tokens and chains do not change.
  /// @dev Only callable by DEFAULT_ADMIN_ROLE or PAUSER_ROLE. The default admin can stop transfers in an emergency
  /// without first granting itself PAUSER_ROLE.
  function pauseGlobally() external {
    _checkRoleOr(DEFAULT_ADMIN_ROLE, PAUSER_ROLE);
    i_connectionStore.setGlobalPause(true);
  }

  /// @notice Turns off the global pause. Paused tokens and chains stay paused.
  /// @dev Only callable by DEFAULT_ADMIN_ROLE or UNPAUSER_ROLE.
  function unpauseGlobally() external {
    _checkRoleOr(DEFAULT_ADMIN_ROLE, UNPAUSER_ROLE);
    i_connectionStore.setGlobalPause(false);
  }

  /// @notice Pauses tokens in one or both directions. Any address can be paused, even before the token is
  /// registered.
  /// @dev Only callable by DEFAULT_ADMIN_ROLE or PAUSER_ROLE.
  /// @param localTokens The tokens to pause.
  /// @param direction The direction to pause, or Both.
  function pauseTokens(
    address[] calldata localTokens,
    ConnectionPolicyStore.PauseDirection direction
  ) external {
    _checkRoleOr(DEFAULT_ADMIN_ROLE, PAUSER_ROLE);
    i_connectionStore.setTokenPause(localTokens, direction, true);
  }

  /// @notice Unpauses tokens in one or both directions.
  /// @dev Only callable by DEFAULT_ADMIN_ROLE or UNPAUSER_ROLE.
  /// @param localTokens The tokens to unpause.
  /// @param direction The direction to unpause, or Both.
  function unpauseTokens(
    address[] calldata localTokens,
    ConnectionPolicyStore.PauseDirection direction
  ) external {
    _checkRoleOr(DEFAULT_ADMIN_ROLE, UNPAUSER_ROLE);
    i_connectionStore.setTokenPause(localTokens, direction, false);
  }

  /// @notice Pauses remote chains for every token, in one or both directions.
  /// @dev Only callable by DEFAULT_ADMIN_ROLE or PAUSER_ROLE.
  /// @param remoteChainSelectors The remote chains to pause.
  /// @param direction The direction to pause, or Both.
  function pauseChains(
    uint64[] calldata remoteChainSelectors,
    ConnectionPolicyStore.PauseDirection direction
  ) external {
    _checkRoleOr(DEFAULT_ADMIN_ROLE, PAUSER_ROLE);
    i_connectionStore.setChainPause(remoteChainSelectors, direction, true);
  }

  /// @notice Unpauses remote chains in one or both directions.
  /// @dev Only callable by DEFAULT_ADMIN_ROLE or UNPAUSER_ROLE.
  /// @param remoteChainSelectors The remote chains to unpause.
  /// @param direction The direction to unpause, or Both.
  function unpauseChains(
    uint64[] calldata remoteChainSelectors,
    ConnectionPolicyStore.PauseDirection direction
  ) external {
    _checkRoleOr(DEFAULT_ADMIN_ROLE, UNPAUSER_ROLE);
    i_connectionStore.setChainPause(remoteChainSelectors, direction, false);
  }

  // ================================================================
  // │                            Fees                              │
  // ================================================================

  /// @notice Withdraws the pool's full balance of each fee token to the recipient.
  /// @dev Only callable by FEE_COLLECTOR_ROLE. Reverts if the pool does not use this contract as its pool config
  /// manager.
  /// @param pool The pool that holds the fees.
  /// @param feeTokens The tokens to withdraw.
  /// @param recipient The fee recipient. Must not be the zero address.
  function withdrawFeeTokens(
    IMultiTokenPool pool,
    address[] calldata feeTokens,
    address recipient
  ) external onlyRole(FEE_COLLECTOR_ROLE) {
    // The pool would also reject this call, but checking here gives a clearer error.
    address poolConfigManager = pool.getPoolConfigManager();
    if (poolConfigManager != address(this)) revert PoolConfigManagerMismatch(address(pool), poolConfigManager);
    pool.withdrawFeeTokens(feeTokens, recipient);
  }

  // ================================================================
  // │                           Helpers                            │
  // ================================================================

  /// @dev Only the token's registered pool can consume its rate limits.
  function _onlyTokenPool(
    address localToken
  ) private view {
    if (msg.sender != i_tokenAdminRegistry.getPool(localToken)) revert CallerIsNotTokenPool(msg.sender, localToken);
  }

  /// @dev Checks all three stores after setting an inline connection record or adding a subset token.
  /// Other writes, including connection policy changes, require checks in off-chain tooling.
  function _checkLane(
    address localToken,
    uint64 remoteChainSelector
  ) private view {
    i_connectionStore.checkLaneRecords(localToken, remoteChainSelector);
    i_feeStore.checkLaneRecords(localToken, remoteChainSelector);
    i_securityStore.checkLaneRecords(localToken, remoteChainSelector);
  }

  /// @dev Ends the external call with the store's raw return data. The manager and store functions must have matching
  /// signatures and return types.
  function _forwardView(
    address store
  ) private view {
    bytes memory returnData = Address.functionStaticCall(store, msg.data);
    // Return the raw data to avoid decoding and encoding the same values again.
    assembly ("memory-safe") {
      return(add(returnData, 32), mload(returnData))
    }
  }

  /// @dev Requires CONFIG_ROLE, POLICY_ROLE, or the role of the type.
  function _checkPolicyRoles(
    bytes32 typeRole
  ) private view {
    if (!hasRole(CONFIG_ROLE, msg.sender)) _checkRoleOr(POLICY_ROLE, typeRole);
  }

  /// @dev Requires CONFIG_ROLE, TOKEN_MANAGER_ROLE, or the role of the type.
  function _checkTokenRoles(
    bytes32 typeRole
  ) private view {
    if (!hasRole(CONFIG_ROLE, msg.sender)) _checkRoleOr(TOKEN_MANAGER_ROLE, typeRole);
  }

  /// @dev Requires either role. The revert error identifies otherRole.
  function _checkRoleOr(
    bytes32 role,
    bytes32 otherRole
  ) private view {
    if (!hasRole(role, msg.sender)) _checkRole(otherRole);
  }
}
