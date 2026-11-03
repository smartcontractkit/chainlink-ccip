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

/// @notice Entry point to configure the tokens of multi-token pools and to read their config.
/// @dev This contract holds roles, not config, so it can be replaced without copying state.
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

  // AccessControlDefaultAdminRules provides DEFAULT_ADMIN_ROLE, the delayed default-admin transfer and the role
  // administration functions. This contract only adds its own roles and the checks that gate each function.

  /// @notice Can make every config write that POLICY_ROLE or TOKEN_MANAGER_ROLE can make. It cannot pause, unpause or
  /// withdraw fees.
  bytes32 public constant CONFIG_ROLE = keccak256("CONFIG_ROLE");
  /// @notice Can make every config write that the policy role of any type can make.
  bytes32 public constant POLICY_ROLE = keccak256("POLICY_ROLE");
  /// @notice Sets and removes connection policy records, removes connection policies, sets the remote token address
  /// providers of connection policies, and adds and removes subset tokens.
  bytes32 public constant CONNECTION_POLICY_ROLE = keccak256("CONNECTION_POLICY_ROLE");
  /// @notice Sets and removes remote pool policy records, and removes remote pool policies.
  bytes32 public constant REMOTE_POOL_POLICY_ROLE = keccak256("REMOTE_POOL_POLICY_ROLE");
  /// @notice Sets and removes rate limit policy records, updates their configs, and removes rate limit policies.
  bytes32 public constant RATE_LIMIT_POLICY_ROLE = keccak256("RATE_LIMIT_POLICY_ROLE");
  /// @notice Sets and removes fee policy records, and removes fee policies.
  bytes32 public constant FEE_POLICY_ROLE = keccak256("FEE_POLICY_ROLE");
  /// @notice Sets and removes destination overhead policy records, and removes destination overhead policies.
  bytes32 public constant DEST_OVERHEAD_POLICY_ROLE = keccak256("DEST_OVERHEAD_POLICY_ROLE");
  /// @notice Sets and removes security policy records, and removes security policies.
  bytes32 public constant SECURITY_POLICY_ROLE = keccak256("SECURITY_POLICY_ROLE");
  /// @notice Registers and deregisters tokens, sets their profiles and their policies, sets and removes their inline
  /// records of every type, and updates the configs of their inline rate limit records.
  bytes32 public constant TOKEN_MANAGER_ROLE = keccak256("TOKEN_MANAGER_ROLE");
  /// @notice Pauses tokens and chains, and turns on the global pause.
  bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
  /// @notice Unpauses tokens and chains, and turns off the global pause. An unpause lets transfers start again, so this
  /// role is separate from pausing.
  bytes32 public constant UNPAUSER_ROLE = keccak256("UNPAUSER_ROLE");
  /// @notice Withdraws the token fees that a pool holds.
  bytes32 public constant FEE_COLLECTOR_ROLE = keccak256("FEE_COLLECTOR_ROLE");

  // ================================================================
  // │                        Static config                         │
  // ================================================================

  /// @dev The registry decides which pool can call the validate functions for a token.
  ITokenAdminRegistry private immutable i_tokenAdminRegistry;
  /// @dev Holds the connection and remote pool types, the token profiles and the pause state.
  ConnectionPolicyStore private immutable i_connectionStore;
  /// @dev Holds the fee and destination overhead types.
  FeePolicyStore private immutable i_feeStore;
  /// @dev Holds the rate limit and security types and the live buckets.
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
  /// @return tokenAdminRegistry The TokenAdminRegistry that names the pool of each token.
  /// @return connectionStore The store of the connection and remote pool types, the token profiles and the pause
  /// state.
  /// @return feeStore The store of the fee and destination overhead types.
  /// @return securityStore The store of the rate limit and security types and the live buckets.
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
  /// @dev The connection store resolves the lane and builds the pool data. The fee store checks the pool data size and
  /// computes the token fee. The security store consumes the outbound rate limit of the amount after the fee. Only the
  /// security store changes state.
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
    lane.feeAmount = i_feeStore.validateLockOrBurn(
      localToken, remoteChainSelector, amount, requestedFinalityConfig, chargeTokenFee, lane.destPoolData.length
    );
    lane.hook = i_securityStore.validateLockOrBurn(
      localToken, remoteChainSelector, amount - lane.feeAmount, requestedFinalityConfig
    );
    return lane;
  }

  /// @inheritdoc IPoolConfigManager
  /// @dev The connection store resolves the lane, checks the pool data and converts the amount. The security store
  /// then consumes the inbound rate limit of the amount in local decimals. Only the security store changes state.
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
    transfer.hook = i_securityStore.validateReleaseOrMint(
      localToken, remoteChainSelector, transfer.localAmount, requestedFinalityConfig
    );
    return transfer;
  }

  // ================================================================
  // │                            Views                             │
  // ================================================================

  /// @inheritdoc IPoolConfigManager
  /// @dev The connection store checks the requested finality, then the fee store builds the fees. A lane that does
  /// not exist is zeroed.
  function getFeeBatch(
    LaneKey[] calldata lanes,
    bytes4 requestedFinalityConfig
  ) external view override returns (LaneFee[] memory fees) {
    i_connectionStore.checkRequestedFinalityBatch(lanes, requestedFinalityConfig);
    bool[] memory hasLanes = i_connectionStore.hasLaneBatch(lanes);
    fees = i_feeStore.getFeeBatch(lanes, requestedFinalityConfig);
    for (uint256 i = 0; i < lanes.length; ++i) {
      if (!hasLanes[i]) {
        delete fees[i];
      }
    }
    return fees;
  }

  /// @inheritdoc IPoolConfigManager
  /// @dev The connection store identifies existing lanes, then the fee store builds the fee configs. A lane that does
  /// not exist is zeroed.
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
  /// @dev Ignores the pause state. The connection store identifies existing lanes, then the security store gives the
  /// buckets. A lane that does not exist gets zero buckets. SecurityPolicyStore documents the bucket semantics.
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
  /// @dev Forwards to ConnectionPolicyStore.getLaneConnectionBatch, which documents the rules, also for a token
  /// without a lane.
  function getLaneConnectionBatch(
    LaneKey[] calldata lanes
  ) external view override returns (LaneConnection[] memory connections) {
    return i_connectionStore.getLaneConnectionBatch(lanes);
  }

  /// @inheritdoc IPoolConfigManager
  /// @dev The requested finality is unused. The connection store identifies existing lanes, then the security store
  /// gives the hook and the CCV lists. A lane that does not exist gets zero values.
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
  /// @dev The connection store gives the profile and the lanes; the fee and security stores add their policy names. A
  /// token that is not registered gets zero values.
  function getTokenConfigBatch(
    address[] calldata localTokens
  ) external view override returns (TokenConfig[] memory configs) {
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

  /// @notice Returns the level of the record that each type would use for each lane, following the resolution of a
  /// transfer. Never reverts.
  /// @dev The connection level shows whether the lane exists. Connection and remote pool policies have no default
  /// records, so their levels are never PolicyDefault. An inline remote pool record may also trust policy pools; its
  /// level stays Inline. Chain selector zero gives None for every type. Each returned array is aligned with lanes by
  /// index.
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

  // These views forward the call with _forwardView, so each one has the parameters and returns of its store function
  // and no names, because the forward does not read them. The store function documents the result.

  /// @notice Returns the flag and the subset tokens of a connection policy record. Forwards to
  /// ConnectionPolicyStore.getSubsetTokens. Never reverts.
  function getSubsetTokens(
    bytes32, // policyName
    uint64 // remoteChainSelector
  ) external view returns (bool, address[] memory) {
    _forwardView(address(i_connectionStore));
  }

  /// @notice Returns the records of a connection policy, in the shape of applyConnectionPolicyConfigUpdates. Forwards
  /// to ConnectionPolicyStore.getConnectionPolicy. Never reverts.
  function getConnectionPolicy(
    bytes32 // policyName
  ) external view returns (uint32, address, ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory) {
    _forwardView(address(i_connectionStore));
  }

  /// @notice Returns the records of a remote pool policy, in the shape of applyRemotePoolPolicyConfigUpdates.
  /// Forwards to ConnectionPolicyStore.getRemotePoolPolicy. Never reverts.
  function getRemotePoolPolicy(
    bytes32 // policyName
  ) external view returns (uint32, ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory) {
    _forwardView(address(i_connectionStore));
  }

  /// @notice Returns the records of a rate limit policy, in the shape of applyRateLimitPolicyConfigUpdates.
  /// Forwards to SecurityPolicyStore.getRateLimitPolicy. Never reverts.
  function getRateLimitPolicy(
    bytes32 // policyName
  ) external view returns (uint32, uint8, SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory) {
    _forwardView(address(i_securityStore));
  }

  /// @notice Returns the records of a fee policy, in the shape of applyFeePolicyConfigUpdates. Forwards to
  /// FeePolicyStore.getFeePolicy. Never reverts.
  function getFeePolicy(
    bytes32 // policyName
  ) external view returns (uint32, FeePolicyStore.FeePolicyConfigArgs[] memory) {
    _forwardView(address(i_feeStore));
  }

  /// @notice Returns the records of a destination overhead policy, in the shape of
  /// applyDestOverheadPolicyConfigUpdates. Forwards to FeePolicyStore.getDestOverheadPolicy. Never reverts.
  function getDestOverheadPolicy(
    bytes32 // policyName
  ) external view returns (uint32, FeePolicyStore.DestOverheadPolicyConfigArgs[] memory) {
    _forwardView(address(i_feeStore));
  }

  /// @notice Returns the records of a security policy, in the shape of applySecurityPolicyConfigUpdates. Forwards to
  /// SecurityPolicyStore.getSecurityPolicy. Never reverts.
  function getSecurityPolicy(
    bytes32 // policyName
  ) external view returns (uint32, SecurityPolicyStore.SecurityPolicyConfigArgs[] memory) {
    _forwardView(address(i_securityStore));
  }

  // Each inline record view lists every inline record of its type for a token, also on a chain where the token has no
  // lane. A store view with the same signature builds the list.

  /// @notice Returns the inline connection records of a token, in the shape of applyInlineConnectionConfigUpdates.
  /// Forwards to ConnectionPolicyStore.getInlineConnectionConfigs. Never reverts.
  function getInlineConnectionConfigs(
    address // localToken
  ) external view returns (ConnectionPolicyStore.InlineConnectionConfigArgs[] memory) {
    _forwardView(address(i_connectionStore));
  }

  /// @notice Returns the inline remote pool records of a token, in the shape of applyInlineRemotePoolConfigUpdates.
  /// Forwards to ConnectionPolicyStore.getInlineRemotePoolConfigs. Never reverts.
  function getInlineRemotePoolConfigs(
    address // localToken
  ) external view returns (ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory) {
    _forwardView(address(i_connectionStore));
  }

  /// @notice Returns the inline rate limit records of a token, in the shape of applyInlineRateLimitConfigUpdates.
  /// Forwards to SecurityPolicyStore.getInlineRateLimitConfigs. Never reverts. Each record shows the config of its
  /// buckets.
  function getInlineRateLimitConfigs(
    address // localToken
  ) external view returns (SecurityPolicyStore.InlineRateLimitConfigArgs[] memory) {
    _forwardView(address(i_securityStore));
  }

  /// @notice Returns the inline fee records of a token, in the shape of applyInlineFeeConfigUpdates. Forwards to
  /// FeePolicyStore.getInlineFeeConfigs. Never reverts.
  function getInlineFeeConfigs(
    address // localToken
  ) external view returns (FeePolicyStore.InlineFeeConfigArgs[] memory) {
    _forwardView(address(i_feeStore));
  }

  /// @notice Returns the inline destination overhead records of a token, in the shape of
  /// applyInlineDestOverheadConfigUpdates. Forwards to FeePolicyStore.getInlineDestOverheadConfigs. Never reverts.
  function getInlineDestOverheadConfigs(
    address // localToken
  ) external view returns (FeePolicyStore.InlineDestOverheadConfigArgs[] memory) {
    _forwardView(address(i_feeStore));
  }

  /// @notice Returns the inline security records of a token, in the shape of applyInlineSecurityConfigUpdates.
  /// Forwards to SecurityPolicyStore.getInlineSecurityConfigs. Never reverts.
  function getInlineSecurityConfigs(
    address // localToken
  ) external view returns (SecurityPolicyStore.InlineSecurityConfigArgs[] memory) {
    _forwardView(address(i_securityStore));
  }

  /// @notice Returns the name of every policy of each type. Never reverts.
  /// @dev The lists include policies that no token uses and policies with no records.
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

  // Each config write has a store function with the same signature, so its call data is also a valid call to that
  // store function. Address.functionCall returns each store revert that has data unchanged. A revert without data
  // becomes FailedCall.

  /// @notice Deregisters tokens, then sets the profile of tokens. All three stores apply the call. See
  /// ConnectionPolicyStore.applyTokenProfileUpdates, FeePolicyStore.applyTokenProfileUpdates and
  /// SecurityPolicyStore.applyTokenProfileUpdates.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE. The connection store runs first, so a token that fails
  /// in more than one store reverts with the error of the first of them.
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

  /// @notice Sets the connection policy of tokens. See ConnectionPolicyStore.applyTokenConnectionPolicyUpdates.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE. The lanes that the new policy opens are not checked, the
  /// same as for a connection policy record.
  function applyTokenConnectionPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata // tokenPoliciesToSet
  ) external {
    _checkRoleOr(CONFIG_ROLE, TOKEN_MANAGER_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
  }

  /// @notice Sets the remote pool policy of tokens. See ConnectionPolicyStore.applyTokenRemotePoolPolicyUpdates.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE.
  function applyTokenRemotePoolPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata // tokenPoliciesToSet
  ) external {
    _checkRoleOr(CONFIG_ROLE, TOKEN_MANAGER_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
  }

  /// @notice Sets the rate limit policy of tokens. See SecurityPolicyStore.applyTokenRateLimitPolicyUpdates.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE.
  function applyTokenRateLimitPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata // tokenPoliciesToSet
  ) external {
    _checkRoleOr(CONFIG_ROLE, TOKEN_MANAGER_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  /// @notice Sets the fee policy of tokens. See FeePolicyStore.applyTokenFeePolicyUpdates.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE.
  function applyTokenFeePolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata // tokenPoliciesToSet
  ) external {
    _checkRoleOr(CONFIG_ROLE, TOKEN_MANAGER_ROLE);
    Address.functionCall(address(i_feeStore), msg.data);
  }

  /// @notice Sets the destination overhead policy of tokens. See FeePolicyStore.applyTokenDestOverheadPolicyUpdates.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE.
  function applyTokenDestOverheadPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata // tokenPoliciesToSet
  ) external {
    _checkRoleOr(CONFIG_ROLE, TOKEN_MANAGER_ROLE);
    Address.functionCall(address(i_feeStore), msg.data);
  }

  /// @notice Sets the security policy of tokens. See SecurityPolicyStore.applyTokenSecurityPolicyUpdates.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE.
  function applyTokenSecurityPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata // tokenPoliciesToSet
  ) external {
    _checkRoleOr(CONFIG_ROLE, TOKEN_MANAGER_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  /// @notice Removes and sets connection policy records, and removes connection policies. See
  /// ConnectionPolicyStore.applyConnectionPolicyConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE or CONNECTION_POLICY_ROLE. A policy record is not checked, because
  /// no store lists the tokens that it serves.
  function applyConnectionPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata, // policyConfigsToRemove
    bytes32[] calldata, // policiesToRemove
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] calldata // policyConfigsToSet
  ) external {
    _checkPolicyRoles(CONNECTION_POLICY_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
  }

  /// @notice Removes and sets inline connection records. See ConnectionPolicyStore.applyInlineConnectionConfigUpdates.
  /// Each inline record that is set opens a lane, so the token must have a record of each other type for the chain
  /// after the call.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE. Reverts for the first type that a token misses.
  /// @param inlineConfigsToSet The inline records to set.
  function applyInlineConnectionConfigUpdates(
    LaneKey[] calldata, // inlineConfigsToRemove
    ConnectionPolicyStore.InlineConnectionConfigArgs[] calldata inlineConfigsToSet
  ) external {
    _checkRoleOr(CONFIG_ROLE, TOKEN_MANAGER_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
    for (uint256 i = 0; i < inlineConfigsToSet.length; ++i) {
      _checkLane(inlineConfigsToSet[i].localToken, inlineConfigsToSet[i].remoteChainSelector);
    }
  }

  /// @notice Removes and adds subset tokens. See ConnectionPolicyStore.applySubsetTokenUpdates. Each subset token that
  /// is added opens a lane, so the token must have a record of each other type for the chain after the call.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE or CONNECTION_POLICY_ROLE. Reverts for the first type that a token
  /// misses.
  /// @param tokensToAdd The subset tokens to add.
  function applySubsetTokenUpdates(
    ConnectionPolicyStore.SubsetTokenArgs[] calldata, // tokensToRemove
    ConnectionPolicyStore.SubsetTokenArgs[] calldata tokensToAdd
  ) external {
    _checkPolicyRoles(CONNECTION_POLICY_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
    for (uint256 i = 0; i < tokensToAdd.length; ++i) {
      _checkLane(tokensToAdd[i].localToken, tokensToAdd[i].remoteChainSelector);
    }
  }

  /// @notice Sets the remote token address providers of connection policies. See
  /// ConnectionPolicyStore.applyRemoteTokenAddressProviderUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE or CONNECTION_POLICY_ROLE. A provider set opens no lane, so the
  /// lane check does not run.
  function applyRemoteTokenAddressProviderUpdates(
    ConnectionPolicyStore.RemoteTokenAddressProviderArgs[] calldata // providersToSet
  ) external {
    _checkPolicyRoles(CONNECTION_POLICY_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
  }

  /// @notice Removes and sets remote pool policy records, and removes remote pool policies. See
  /// ConnectionPolicyStore.applyRemotePoolPolicyConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE or REMOTE_POOL_POLICY_ROLE.
  function applyRemotePoolPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata, // policyConfigsToRemove
    bytes32[] calldata, // policiesToRemove
    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] calldata // policyConfigsToSet
  ) external {
    _checkPolicyRoles(REMOTE_POOL_POLICY_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
  }

  /// @notice Removes and sets inline remote pool records. See ConnectionPolicyStore.applyInlineRemotePoolConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE.
  function applyInlineRemotePoolConfigUpdates(
    LaneKey[] calldata, // inlineConfigsToRemove
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] calldata // inlineConfigsToSet
  ) external {
    _checkRoleOr(CONFIG_ROLE, TOKEN_MANAGER_ROLE);
    Address.functionCall(address(i_connectionStore), msg.data);
  }

  /// @notice Removes and sets rate limit policy records, and removes rate limit policies. See
  /// SecurityPolicyStore.applyRateLimitPolicyConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE or RATE_LIMIT_POLICY_ROLE.
  function applyRateLimitPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata, // policyConfigsToRemove
    bytes32[] calldata, // policiesToRemove
    SecurityPolicyStore.RateLimitPolicyConfigArgs[] calldata // policyConfigsToSet
  ) external {
    _checkPolicyRoles(RATE_LIMIT_POLICY_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  /// @notice Updates one config of existing rate limit policy records. See
  /// SecurityPolicyStore.applyRateLimitPolicyBucketUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE or RATE_LIMIT_POLICY_ROLE.
  function applyRateLimitPolicyBucketUpdates(
    SecurityPolicyStore.RateLimitPolicyBucketArgs[] calldata // bucketConfigsToSet
  ) external {
    _checkPolicyRoles(RATE_LIMIT_POLICY_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  /// @notice Removes and sets inline rate limit records. See SecurityPolicyStore.applyInlineRateLimitConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE.
  function applyInlineRateLimitConfigUpdates(
    LaneKey[] calldata, // inlineConfigsToRemove
    SecurityPolicyStore.InlineRateLimitConfigArgs[] calldata // inlineConfigsToSet
  ) external {
    _checkRoleOr(CONFIG_ROLE, TOKEN_MANAGER_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  /// @notice Updates one config of existing inline rate limit records. See
  /// SecurityPolicyStore.applyInlineRateLimitBucketUpdates.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE.
  function applyInlineRateLimitBucketUpdates(
    SecurityPolicyStore.InlineRateLimitBucketArgs[] calldata // bucketConfigsToSet
  ) external {
    _checkRoleOr(CONFIG_ROLE, TOKEN_MANAGER_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  /// @notice Removes and sets fee policy records, and removes fee policies. See
  /// FeePolicyStore.applyFeePolicyConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE or FEE_POLICY_ROLE.
  function applyFeePolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata, // policyConfigsToRemove
    bytes32[] calldata, // policiesToRemove
    FeePolicyStore.FeePolicyConfigArgs[] calldata // policyConfigsToSet
  ) external {
    _checkPolicyRoles(FEE_POLICY_ROLE);
    Address.functionCall(address(i_feeStore), msg.data);
  }

  /// @notice Removes and sets inline fee records. See FeePolicyStore.applyInlineFeeConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE.
  function applyInlineFeeConfigUpdates(
    LaneKey[] calldata, // inlineConfigsToRemove
    FeePolicyStore.InlineFeeConfigArgs[] calldata // inlineConfigsToSet
  ) external {
    _checkRoleOr(CONFIG_ROLE, TOKEN_MANAGER_ROLE);
    Address.functionCall(address(i_feeStore), msg.data);
  }

  /// @notice Removes and sets destination overhead policy records, and removes destination overhead policies. See
  /// FeePolicyStore.applyDestOverheadPolicyConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE or DEST_OVERHEAD_POLICY_ROLE.
  function applyDestOverheadPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata, // policyConfigsToRemove
    bytes32[] calldata, // policiesToRemove
    FeePolicyStore.DestOverheadPolicyConfigArgs[] calldata // policyConfigsToSet
  ) external {
    _checkPolicyRoles(DEST_OVERHEAD_POLICY_ROLE);
    Address.functionCall(address(i_feeStore), msg.data);
  }

  /// @notice Removes and sets inline destination overhead records. See
  /// FeePolicyStore.applyInlineDestOverheadConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE.
  function applyInlineDestOverheadConfigUpdates(
    LaneKey[] calldata, // inlineConfigsToRemove
    FeePolicyStore.InlineDestOverheadConfigArgs[] calldata // inlineConfigsToSet
  ) external {
    _checkRoleOr(CONFIG_ROLE, TOKEN_MANAGER_ROLE);
    Address.functionCall(address(i_feeStore), msg.data);
  }

  /// @notice Removes and sets security policy records, and removes security policies. See
  /// SecurityPolicyStore.applySecurityPolicyConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE, POLICY_ROLE or SECURITY_POLICY_ROLE.
  function applySecurityPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata, // policyConfigsToRemove
    bytes32[] calldata, // policiesToRemove
    SecurityPolicyStore.SecurityPolicyConfigArgs[] calldata // policyConfigsToSet
  ) external {
    _checkPolicyRoles(SECURITY_POLICY_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  /// @notice Removes and sets inline security records. See SecurityPolicyStore.applyInlineSecurityConfigUpdates.
  /// @dev Only callable by CONFIG_ROLE or TOKEN_MANAGER_ROLE.
  function applyInlineSecurityConfigUpdates(
    LaneKey[] calldata, // inlineConfigsToRemove
    SecurityPolicyStore.InlineSecurityConfigArgs[] calldata // inlineConfigsToSet
  ) external {
    _checkRoleOr(CONFIG_ROLE, TOKEN_MANAGER_ROLE);
    Address.functionCall(address(i_securityStore), msg.data);
  }

  // ================================================================
  // │                            Pause                             │
  // ================================================================

  /// @notice Turns on the global pause. It stops every token on every chain in both directions. The pause flags of
  /// tokens and chains do not change.
  /// @dev Only callable by DEFAULT_ADMIN_ROLE or PAUSER_ROLE. So the default admin can stop all transfers in an
  /// emergency without first granting itself a role.
  function pauseGlobally() external {
    _checkRoleOr(DEFAULT_ADMIN_ROLE, PAUSER_ROLE);
    i_connectionStore.setGlobalPause(true);
  }

  /// @notice Turns off the global pause. Paused tokens and chains stay paused.
  /// @dev Only callable by DEFAULT_ADMIN_ROLE or UNPAUSER_ROLE.
  function unpauseGlobally() external {
    _checkRoleOr(DEFAULT_ADMIN_ROLE, PAUSER_ROLE);
    i_connectionStore.setGlobalPause(false);
  }

  /// @notice Pauses tokens in one direction or in both directions. Any address can be paused, also before the token is
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

  /// @notice Unpauses tokens in one direction or in both directions.
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

  /// @notice Pauses remote chains for every token, in one direction or in both directions.
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

  /// @notice Unpauses remote chains in one direction or in both directions.
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

  /// @notice Sends the token fees that a pool holds to a recipient. The pool sends its whole balance of each token.
  /// @dev Only callable by FEE_COLLECTOR_ROLE. The pool must name this contract as its pool config manager, so this
  /// call reaches only the pools that this contract serves. The pool also accepts this call only from its pool config
  /// manager.
  /// @param pool The pool that holds the fees.
  /// @param feeTokens The tokens to withdraw.
  /// @param recipient The receiver of the fees. Must not be zero.
  function withdrawFeeTokens(
    IMultiTokenPool pool,
    address[] calldata feeTokens,
    address recipient
  ) external onlyRole(FEE_COLLECTOR_ROLE) {
    address poolConfigManager = pool.getPoolConfigManager();
    if (poolConfigManager != address(this)) revert PoolConfigManagerMismatch(address(pool), poolConfigManager);
    pool.withdrawFeeTokens(feeTokens, recipient);
  }

  // ================================================================
  // │                           Helpers                            │
  // ================================================================

  /// @dev The validate functions consume rate limits. The token admin decides which pool may do that.
  function _onlyTokenPool(
    address localToken
  ) private view {
    if (msg.sender != i_tokenAdminRegistry.getPool(localToken)) revert CallerIsNotTokenPool(msg.sender, localToken);
  }

  /// @dev The one check that reads all three stores. It runs after each write that opens a lane: an inline connection
  /// record set and a subset token add. Off-chain tooling checks the other writes, such as a connection policy record
  /// set. The connection store checks the remote pool record, then the fee store checks its records, then the security
  /// store checks its records.
  function _checkLane(
    address localToken,
    uint64 remoteChainSelector
  ) private view {
    i_connectionStore.checkLaneRecords(localToken, remoteChainSelector);
    i_feeStore.checkLaneRecords(localToken, remoteChainSelector);
    i_securityStore.checkLaneRecords(localToken, remoteChainSelector);
  }

  /// @dev Ends the call with the unchanged return data of the store view that has the signature of the called view.
  /// Both views return the same types, so the data is also valid here and this contract needs no ABI coder for those
  /// types. Only assembly can return raw bytes from a function with typed returns.
  function _forwardView(
    address store
  ) private view {
    bytes memory returnData = Address.functionStaticCall(store, msg.data);
    assembly ("memory-safe") {
      return(add(returnData, 32), mload(returnData))
    }
  }

  /// @dev Passes for CONFIG_ROLE, POLICY_ROLE or the policy role of the type.
  function _checkPolicyRoles(
    bytes32 typePolicyRole
  ) private view {
    if (!hasRole(CONFIG_ROLE, msg.sender)) _checkRoleOr(POLICY_ROLE, typePolicyRole);
  }

  /// @dev Passes for either role. A failed check names the second role.
  function _checkRoleOr(
    bytes32 role,
    bytes32 otherRole
  ) private view {
    if (!hasRole(role, msg.sender)) _checkRole(otherRole);
  }
}
