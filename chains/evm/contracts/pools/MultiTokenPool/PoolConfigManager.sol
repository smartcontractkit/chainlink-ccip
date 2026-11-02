// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../interfaces/IPoolV2.sol";
import {ITokenAdminRegistry} from "../../interfaces/ITokenAdminRegistry.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Pool} from "../../libraries/Pool.sol";
import {FeePolicyStore} from "./policies/FeePolicyStore.sol";
import {SecurityPolicyStore} from "./policies/SecurityPolicyStore.sol";
import {ConnectionPolicyStore} from "./tmp/ConnectionPolicyStore.stub.sol";

import {
  AccessControlDefaultAdminRules
} from "@openzeppelin/contracts@5.3.0/access/extensions/AccessControlDefaultAdminRules.sol";

/// @notice Entry point to configure the tokens of multi-token pools and to read their config.
/// @dev This contract holds roles, not config, so it can be replaced without copying state.
contract PoolConfigManager is IPoolConfigManager, ITypeAndVersion, AccessControlDefaultAdminRules {
  error ZeroAddressInvalid();
  error CallerIsNotTokenPool(address caller, address localToken);

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
  /// @dev The connection store resolves the lane, the fee store computes the token fee, and the security store
  /// consumes the outbound rate limit of the amount after the fee. Only the security store changes state.
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
    lane.feeAmount =
      i_feeStore.resolveTokenFee(localToken, remoteChainSelector, amount, requestedFinalityConfig, chargeTokenFee);
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

  // ================================================================
  // │                           Helpers                            │
  // ================================================================

  /// @dev The validate functions consume rate limits. The token admin decides which pool may do that.
  function _onlyTokenPool(
    address localToken
  ) private view {
    if (msg.sender != i_tokenAdminRegistry.getPool(localToken)) revert CallerIsNotTokenPool(msg.sender, localToken);
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
