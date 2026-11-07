// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdditionalPoolDataProvider} from "../../../interfaces/IAdditionalPoolDataProvider.sol";
import {ILockBox} from "../../../interfaces/ILockBox.sol";
import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";
import {IRemoteTokenAddressProvider} from "../../../interfaces/IRemoteTokenAddressProvider.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../libraries/PolicyConfig.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {IERC20Metadata} from "@openzeppelin/contracts@5.3.0/token/ERC20/extensions/IERC20Metadata.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Stores token profiles, connections, trusted remote pools, and pause state for a PoolConfigManager.
/// Builds outbound pool data and validates inbound pool data for multi-token pools.
/// @dev Only authorized pool config managers can update the configuration.
/// Connection records determine which remote chains a token can use. Transfers also require a remote pool record.
/// Inline connections specify the remote token directly; policy connections use an IRemoteTokenAddressProvider.
/// A connection policy can enable a chain for all assigned tokens or restrict it to a subset.
/// Inline records override the corresponding policy records. Connection and remote pool policies have no defaults.
contract ConnectionPolicyStore is ITypeAndVersion, AuthorizedCallers {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;
  using EnumerableSet for EnumerableSet.UintSet;

  error ZeroAddressInvalid();
  error InvalidDecimalArgs(address localToken, uint8 expected, uint8 actual);
  error InvalidLockOrBurnAddress(
    address localToken, IPoolConfigManager.LockOrBurnFlavor lockOrBurnFlavor, address lockOrBurnAddress
  );
  error InvalidRemoteToken(address localToken, uint64 remoteChainSelector);
  error RemotePoolAlreadyAdded(uint64 remoteChainSelector, bytes remotePoolAddress);
  error PolicyRecordHasNoPools(bytes32 policyName, uint64 remoteChainSelector);
  error TokenHasConfig(address localToken);
  error ProviderHasNoCode(address provider);
  error PolicyRemoteTokenMismatch(address localToken, uint64 remoteChainSelector);
  error TokenNotInPolicy(address localToken, bytes32 policyName);
  error SubsetTokenAlreadyAdded(bytes32 policyName, uint64 remoteChainSelector, address localToken);
  error SubsetTokenNotFound(bytes32 policyName, uint64 remoteChainSelector, address localToken);
  error GloballyPaused();
  error TokenIsPaused(address localToken, PauseDirection direction);
  error ChainIsPaused(uint64 remoteChainSelector, PauseDirection direction);
  error ConnectionRecordNotFound(address localToken, uint64 remoteChainSelector);
  error RemotePoolRecordNotFound(address localToken, uint64 remoteChainSelector);
  error InvalidSourcePoolAddress(address localToken, uint64 remoteChainSelector, bytes sourcePoolAddress);
  error InvalidSourcePoolData(bytes sourcePoolData);
  error LegacyPoolDataNotAllowed(address localToken, uint64 remoteChainSelector);
  error LegacyRemotePoolRecord(address localToken, uint64 remoteChainSelector);
  error SourceTokenMismatch(
    address localToken, uint64 remoteChainSelector, bytes expectedSourceToken, bytes actualSourceToken
  );
  error InvalidAdditionalPoolData(address localToken, uint64 remoteChainSelector, bytes additionalPoolData);
  error OverflowDetected(uint8 remoteDecimals, uint8 localDecimals, uint256 remoteAmount);

  event InlineConnectionConfigSet(address indexed localToken, uint64 indexed remoteChainSelector, bytes remoteToken);
  event InlineConnectionConfigRemoved(address indexed localToken, uint64 indexed remoteChainSelector);
  event ConnectionPolicyConfigSet(
    bytes32 indexed policyName, uint64 indexed remoteChainSelector, bool onlySubsetTokens
  );
  /// @dev Removing a record also clears its subset tokens.
  event ConnectionPolicyConfigRemoved(bytes32 indexed policyName, uint64 indexed remoteChainSelector);
  /// @dev Removing a policy also clears its remote token address provider.
  event ConnectionPolicyRemoved(bytes32 indexed policyName);
  event RemoteTokenAddressProviderSet(bytes32 indexed policyName, address remoteTokenAddressProvider);
  event SubsetTokenAdded(bytes32 indexed policyName, uint64 indexed remoteChainSelector, address indexed localToken);
  event SubsetTokenRemoved(bytes32 indexed policyName, uint64 indexed remoteChainSelector, address indexed localToken);
  /// @dev Setting a profile for a new token also registers it.
  event TokenProfileSet(
    address indexed localToken,
    uint8 localDecimals,
    IPoolConfigManager.LockOrBurnFlavor lockOrBurnFlavor,
    address lockOrBurnAddress,
    bytes4 allowedFinalityConfig,
    address additionalPoolDataProvider
  );
  event TokenRemoved(address indexed localToken);
  event TokenConnectionPolicySet(
    address indexed localToken, bytes32 indexed oldPolicyName, bytes32 indexed newPolicyName
  );
  event TokenRemotePoolPolicySet(
    address indexed localToken, bytes32 indexed oldPolicyName, bytes32 indexed newPolicyName
  );
  event InlineRemotePoolConfigSet(
    address indexed localToken,
    uint64 indexed remoteChainSelector,
    bool trustPolicyRemotePools,
    RemotePoolConfig remotePoolConfig
  );
  event InlineRemotePoolConfigRemoved(address indexed localToken, uint64 indexed remoteChainSelector);
  event RemotePoolPolicyConfigSet(
    bytes32 indexed policyName, uint64 indexed remoteChainSelector, bytes[] remotePoolAddresses
  );
  event RemotePoolPolicyConfigRemoved(bytes32 indexed policyName, uint64 indexed remoteChainSelector);
  event RemotePoolPolicyRemoved(bytes32 indexed policyName);
  event GlobalPauseSet(bool isPaused);
  event TokenPauseSet(address indexed localToken, bool isOutboundPaused, bool isInboundPaused);
  event ChainPauseSet(uint64 indexed remoteChainSelector, bool isOutboundPaused, bool isInboundPaused);

  /// @dev Whether to pause outbound transfers, inbound transfers, or both.
  enum PauseDirection {
    Outbound,
    Inbound,
    Both
  }

  /// @notice Connection policy configuration for a remote chain.
  struct ConnectionPolicyConfigArgs {
    bytes32 policyName; // The connection policy name. Must be nonzero.
    uint64 remoteChainSelector; // ─╮ The remote chain selector. Must be nonzero.
    bool onlySubsetTokens; // ──────╯ Whether to restrict this chain to the tokens in the subset list.
  }

  /// @notice Connection configuration for a single token and remote chain.
  struct InlineConnectionConfigArgs {
    address localToken; // ─────────╮ The registered local token.
    uint64 remoteChainSelector; // ─╯ The remote chain selector. Must be nonzero.
    bytes remoteToken; //             Remote token address, ABI encoded for EVM chains. Cannot be empty or all zero.
  }

  /// @notice Identifies a token in a connection policy's subset list for a remote chain.
  struct SubsetTokenArgs {
    bytes32 policyName; // The connection policy name.
    uint64 remoteChainSelector; // ─╮ The remote chain selector.
    address localToken; // ─────────╯ A token assigned to the policy.
  }

  /// @notice Remote token address provider configuration for a connection policy.
  struct RemoteTokenAddressProviderArgs {
    bytes32 policyName; // The name of an existing connection policy.
    address remoteTokenAddressProvider; // The IRemoteTokenAddressProvider address. Zero removes the provider.
  }

  /// @notice Trusted remote pools and the outbound pool data format for a token and chain.
  /// @dev Enable useLegacyPoolDataOnSend while the destination uses a legacy TokenPool, which expects only the token
  /// decimals. Disable it once the destination supports the multi-token pool data format.
  /// Allow legacy pool data on receive from a legacy source pool, or from a multi-token pool that sent legacy data
  /// before this chain migrated. Keep it allowed until all inflight messages with legacy pool data have executed.
  struct RemotePoolConfig {
    bool useLegacyPoolDataOnSend; // Whether to send only the token decimals in the pool data.
    IPoolConfigManager.RemotePool[] remotePools; // Trusted remote pools. Addresses must be nonempty and unique.
  }

  /// @notice Trusted remote pools for a policy and remote chain.
  /// @dev Policy records do not allow legacy pool data.
  struct RemotePoolPolicyConfigArgs {
    bytes32 policyName; // The remote pool policy name. Must be nonzero.
    uint64 remoteChainSelector; // The remote chain selector. Must be nonzero.
    bytes[] remotePoolAddresses; // At least one trusted remote pool. Addresses must be nonempty and unique.
  }

  /// @notice Remote pool configuration for a single token and chain.
  /// @dev An inline record overrides the token's remote pool policy. Enable trustPolicyRemotePools to also trust policy
  /// pools while keeping inline permissions for inflight legacy messages. For pools in both records, the inline
  /// allowLegacyPoolDataOnReceive flag applies.
  struct InlineRemotePoolConfigArgs {
    address localToken; // ──────────╮ The registered local token.
    uint64 remoteChainSelector; //   │ The remote chain selector. Must be nonzero.
    bool trustPolicyRemotePools; // ─╯ Whether to also trust pools from the token's policy for this chain.
    RemotePoolConfig remotePoolConfig; // The trusted pools and outbound pool data format.
  }

  /// @notice Stored token profile, registration status, and pause flags.
  // solhint-disable-next-line gas-struct-packing
  struct TokenProfile {
    bool isRegistered; // ───────────────────────────────────╮ True if the token is registered.
    uint8 localDecimals; //                                  │ The decimals of the local token.
    IPoolConfigManager.LockOrBurnFlavor lockOrBurnFlavor; // │ The token operation used for outbound transfers.
    address lockOrBurnAddress; //                            │ The BurnToAddress target or LockBox address.
    bytes4 allowedFinalityConfig; //                         │ Allowed finality.
    bool isOutboundPaused; //                                │ Whether outbound transfers are paused for this token.
    bool isInboundPaused; // ────────────────────────────────╯ Whether inbound transfers are paused for this token.
    address additionalPoolDataProvider; // The token's IAdditionalPoolDataProvider. Zero means no provider.
  }

  /// @notice Token profile, pause flags, and assigned policies.
  /// @dev Only the pause functions change the pause flags.
  struct TokenState {
    address lockOrBurnAddress; // ───────────────────────────╮ The BurnToAddress target or LockBox address.
    uint8 localDecimals; //                                  │ The decimals of the local token.
    IPoolConfigManager.LockOrBurnFlavor lockOrBurnFlavor; // │ The token operation used for outbound transfers.
    bytes4 allowedFinalityConfig; //                         │ Allowed finality.
    bool isOutboundPaused; //                                │ Whether outbound transfers are paused on all chains.
    bool isInboundPaused; // ────────────────────────────────╯ Whether inbound transfers are paused on all chains.
    address additionalPoolDataProvider; // The token's IAdditionalPoolDataProvider. Zero means no provider.
    bytes32 connectionPolicyName; // The assigned connection policy. Zero means no policy.
    bytes32 remotePoolPolicyName; // The assigned remote pool policy. Zero means no policy.
  }

  /// @notice Controls which tokens assigned to a connection policy can use a remote chain.
  struct ConnectionPolicyRecord {
    bool onlySubsetTokens; // Whether to restrict the chain to subsetTokens. False allows all tokens in the policy.
    EnumerableSet.AddressSet subsetTokens; // Allowed tokens when onlySubsetTokens is true. Can be edited while false.
  }

  /// @notice Trusted remote pools and compatibility settings.
  /// @dev All flags are false for policy records.
  struct RemotePoolRecord {
    bool useLegacyPoolDataOnSend; // ─╮ Whether to send only the token decimals in the pool data.
    bool trustPolicyRemotePools; // ──╯ Whether to also trust pools from the token's policy for this chain.
    EnumerableSet.Bytes32Set remotePoolHashes; // The hashes of the trusted remote pools.
    mapping(bytes32 remotePoolHash => bool isAllowed) allowLegacyPoolDataOnReceive; // Legacy receive flag per pool.
  }

  /// @notice The pause flags of a remote chain.
  struct ChainPause {
    bool isOutboundPaused; // ─╮ Whether outbound transfers to this chain are paused for all tokens.
    bool isInboundPaused; // ──╯ Whether inbound transfers from this chain are paused for all tokens.
  }

  /// @dev Length of the ABI-encoded token decimals sent by legacy TokenPools.
  uint256 private constant POOL_DATA_DECIMALS_LENGTH = 32;
  /// @dev Minimum ABI-encoded length of token decimals, a nonempty source token, and empty additional pool data.
  uint256 private constant POOL_DATA_WITH_SOURCE_TOKEN_MIN_LENGTH = 192;

  /// @dev Each policy store tracks registered tokens independently.
  EnumerableSet.AddressSet private s_tokens;
  /// @dev Tokens can be paused before registration. Profiles and pause flags are retained after deregistration.
  mapping(address localToken => TokenState tokenState) private s_tokenStates;
  /// @dev Pool addresses are stored once and shared by records that reference their hash.
  mapping(bytes32 remotePoolHash => bytes remotePoolAddress) private s_remotePoolAddresses;

  // ================================================================
  // │                       Inline storage                         │
  // ================================================================

  /// @dev Remote chains with an inline connection record, indexed by local token.
  mapping(address localToken => EnumerableSet.UintSet remoteChainSelectors) private s_inlineConnectionChains;
  mapping(address localToken => mapping(uint64 remoteChainSelector => bytes remoteToken)) private s_inlineRemoteTokens;

  /// @dev Remote chains with an inline remote pool record, indexed by local token.
  mapping(address localToken => EnumerableSet.UintSet remoteChainSelectors) private s_inlineRemotePoolChains;
  mapping(address localToken => mapping(uint64 remoteChainSelector => RemotePoolRecord record)) private
    s_inlineRemotePoolRecords;

  // ================================================================
  // │                       Policy storage                         │
  // ================================================================

  PolicyConfig.Policies private s_connectionPolicies;
  mapping(bytes32 policyName => mapping(uint64 remoteChainSelector => ConnectionPolicyRecord record)) private
    s_policyConnectionRecords;
  /// @dev Provider used to resolve remote token addresses for each connection policy. Zero means no provider.
  mapping(bytes32 policyName => address remoteTokenAddressProvider) private s_remoteTokenAddressProviders;

  PolicyConfig.Policies private s_remotePoolPolicies;
  mapping(bytes32 policyName => mapping(uint64 remoteChainSelector => RemotePoolRecord record)) private
    s_policyRemotePoolRecords;

  // ================================================================
  // │                        Pause storage                         │
  // ================================================================

  mapping(uint64 remoteChainSelector => ChainPause chainPause) private s_chainPauses;
  /// @dev Unpausing globally does not clear token or chain pauses.
  bool private s_isGloballyPaused;

  /// @param authorizedCallers The pool config managers that can write to this store.
  constructor(
    address[] memory authorizedCallers
  ) AuthorizedCallers(authorizedCallers) {}

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "ConnectionPolicyStore 2.1.0-dev";
  }

  // ================================================================
  // │                     Transfer validation                      │
  // ================================================================

  /// @notice Validates the outbound lane and builds the pool data for lockOrBurn.
  /// @dev Checks pauses, the connection and remote pool records, and the requested finality. Reverts if a provider
  /// call fails.
  /// @param localToken The local token.
  /// @param remoteChainSelector The destination chain selector.
  /// @param requestedFinalityConfig The requested finality.
  /// @param tokenArgs Sender-supplied arguments passed to the additional pool data provider.
  /// @return lane The resolved connection, pool data, and token profile. The fee amount and hook are left unset.
  function resolveOutboundLane(
    address localToken,
    uint64 remoteChainSelector,
    bytes4 requestedFinalityConfig,
    bytes calldata tokenArgs
  ) external view returns (IPoolConfigManager.OutboundLane memory lane) {
    TokenState storage tokenState = s_tokenStates[localToken];
    if (s_isGloballyPaused) revert GloballyPaused();
    if (tokenState.isOutboundPaused) revert TokenIsPaused(localToken, PauseDirection.Outbound);
    if (s_chainPauses[remoteChainSelector].isOutboundPaused) {
      revert ChainIsPaused(remoteChainSelector, PauseDirection.Outbound);
    }

    (PolicyConfig.RecordLevel connectionLevel, bool isInline, RemotePoolRecord storage record) =
      _resolveLane(localToken, remoteChainSelector);
    FinalityCodec._ensureRequestedFinalityAllowed(requestedFinalityConfig, tokenState.allowedFinalityConfig);

    lane.remoteToken = _resolveRemoteToken(localToken, remoteChainSelector, connectionLevel);
    lane.destPoolData = _encodePoolData(
      localToken,
      remoteChainSelector,
      requestedFinalityConfig,
      tokenArgs,
      tokenState,
      isInline && record.useLegacyPoolDataOnSend
    );
    lane.localDecimals = tokenState.localDecimals;
    lane.lockOrBurnFlavor = tokenState.lockOrBurnFlavor;
    lane.lockOrBurnAddress = tokenState.lockOrBurnAddress;
    return lane;
  }

  /// @notice Validates the inbound lane and pool data. Converts the amount to local token units.
  /// @dev Checks pauses and requires both connection and remote pool records. As in TokenPool, allowedFinalityConfig
  /// applies only to outbound transfers; requestedFinalityConfig is passed to the additional pool data provider.
  /// @param localToken The local token.
  /// @param remoteChainSelector The source chain selector.
  /// @param sourcePoolAddress The source pool address, ABI encoded for EVM chains.
  /// @param sourcePoolData The pool data from the source pool.
  /// @param sourceDenominatedAmount The amount in source token units.
  /// @param requestedFinalityConfig The finality requested by the sender.
  /// @return transfer The amount in local token units and the token's lock or burn configuration. The hook is unset.
  function resolveInboundLane(
    address localToken,
    uint64 remoteChainSelector,
    bytes calldata sourcePoolAddress,
    bytes calldata sourcePoolData,
    uint256 sourceDenominatedAmount,
    bytes4 requestedFinalityConfig
  ) external view returns (IPoolConfigManager.ResolvedInboundTransfer memory transfer) {
    TokenState storage tokenState = s_tokenStates[localToken];
    if (s_isGloballyPaused) revert GloballyPaused();
    if (tokenState.isInboundPaused) revert TokenIsPaused(localToken, PauseDirection.Inbound);
    if (s_chainPauses[remoteChainSelector].isInboundPaused) {
      revert ChainIsPaused(remoteChainSelector, PauseDirection.Inbound);
    }

    // Resolve the connection level inside _checkPoolData to avoid a stack-too-deep error here.
    (, bool isInline, RemotePoolRecord storage record) = _resolveLane(localToken, remoteChainSelector);
    bool allowLegacyPoolDataOnReceive =
      _checkSourcePool(localToken, remoteChainSelector, isInline, record, sourcePoolAddress);
    uint8 remoteDecimals = _checkPoolData(
      localToken, remoteChainSelector, requestedFinalityConfig, sourcePoolData, tokenState, allowLegacyPoolDataOnReceive
    );

    transfer.localAmount = _calculateLocalAmount(sourceDenominatedAmount, remoteDecimals, tokenState.localDecimals);
    transfer.lockOrBurnFlavor = tokenState.lockOrBurnFlavor;
    transfer.lockOrBurnAddress = tokenState.lockOrBurnAddress;
    return transfer;
  }

  /// @dev Resolves the connection level and remote pool record. Reverts if either is missing.
  /// isInline identifies whether the remote pool record is inline and can allow legacy pool data.
  function _resolveLane(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (PolicyConfig.RecordLevel connectionLevel, bool isInline, RemotePoolRecord storage record) {
    connectionLevel = _resolveConnectionLevel(localToken, remoteChainSelector);
    if (connectionLevel == PolicyConfig.RecordLevel.None) {
      // If the chain is configured, the token must have been excluded by the policy's subset restriction.
      bytes32 policyName = s_tokenStates[localToken].connectionPolicyName;
      if (s_connectionPolicies.policies[policyName].remoteChainSelectors.contains(remoteChainSelector)) {
        revert SubsetTokenNotFound(policyName, remoteChainSelector, localToken);
      }
      revert ConnectionRecordNotFound(localToken, remoteChainSelector);
    }

    PolicyConfig.RecordLevel remotePoolLevel;
    (remotePoolLevel, record) = _resolveRemotePoolRecord(localToken, remoteChainSelector);
    if (remotePoolLevel == PolicyConfig.RecordLevel.None) {
      revert RemotePoolRecordNotFound(localToken, remoteChainSelector);
    }
    return (connectionLevel, remotePoolLevel == PolicyConfig.RecordLevel.Inline, record);
  }

  /// @dev Checks that the source pool is trusted and returns whether it may send legacy pool data.
  function _checkSourcePool(
    address localToken,
    uint64 remoteChainSelector,
    bool isInline,
    RemotePoolRecord storage record,
    bytes calldata sourcePoolAddress
  ) private view returns (bool) {
    bytes32 sourcePoolHash = keccak256(sourcePoolAddress);
    // Check the resolved record first, so inline settings take precedence for pools listed in both records.
    if (record.remotePoolHashes.contains(sourcePoolHash)) {
      return isInline && record.allowLegacyPoolDataOnReceive[sourcePoolHash];
    }

    if (isInline && record.trustPolicyRemotePools) {
      (PolicyConfig.RecordLevel policyLevel, RemotePoolRecord storage policyRecord) =
        _resolvePolicyRemotePoolRecord(localToken, remoteChainSelector);

      if (policyLevel != PolicyConfig.RecordLevel.None && policyRecord.remotePoolHashes.contains(sourcePoolHash)) {
        // Pools trusted only through the policy cannot send legacy pool data.
        return false;
      }
    }
    revert InvalidSourcePoolAddress(localToken, remoteChainSelector, sourcePoolAddress);
  }

  /// @dev Returns the inline remote token or queries the connection policy's provider.
  /// Reverts if the provider call fails or the remote token is empty or all zero.
  function _resolveRemoteToken(
    address localToken,
    uint64 remoteChainSelector,
    PolicyConfig.RecordLevel connectionLevel
  ) private view returns (bytes memory remoteToken) {
    if (connectionLevel == PolicyConfig.RecordLevel.Inline) {
      remoteToken = s_inlineRemoteTokens[localToken][remoteChainSelector];
    } else if (connectionLevel == PolicyConfig.RecordLevel.PolicyChain) {
      bytes32 policyName = s_tokenStates[localToken].connectionPolicyName;
      address remoteTokenAddressProvider = s_remoteTokenAddressProviders[policyName];

      if (remoteTokenAddressProvider != address(0)) {
        remoteToken = IRemoteTokenAddressProvider(remoteTokenAddressProvider)
          .getRemoteTokenAddress(localToken, remoteChainSelector);
      }
    }

    if (!_isValidRemoteToken(remoteToken)) revert InvalidRemoteToken(localToken, remoteChainSelector);
    return remoteToken;
  }

  /// @dev Returns the policy's remote token. Returns empty bytes if the provider is missing, reverts, or returns an
  /// invalid address. Malformed return data still reverts because try/catch does not catch ABI decoding errors.
  function _tryGetPolicyRemoteToken(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (bytes memory remoteToken) {
    address remoteTokenAddressProvider = s_remoteTokenAddressProviders[s_tokenStates[localToken].connectionPolicyName];
    // Avoid a decoding failure when the provider is unset or has no code.
    if (remoteTokenAddressProvider.code.length == 0) {
      return remoteToken;
    }

    try IRemoteTokenAddressProvider(remoteTokenAddressProvider)
      .getRemoteTokenAddress(localToken, remoteChainSelector) returns (
      bytes memory providerRemoteToken
    ) {
      if (_isValidRemoteToken(providerRemoteToken)) remoteToken = providerRemoteToken;
    } catch {}
    return remoteToken;
  }

  /// @dev Returns false for empty or all-zero remote token addresses.
  function _isValidRemoteToken(
    bytes memory remoteToken
  ) private pure returns (bool) {
    return keccak256(remoteToken) != keccak256(new bytes(remoteToken.length));
  }

  /// @dev Inline connections take precedence over the token's connection policy.
  function _resolveConnectionLevel(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (PolicyConfig.RecordLevel) {
    if (s_inlineConnectionChains[localToken].contains(remoteChainSelector)) return PolicyConfig.RecordLevel.Inline;

    if (_hasPolicyLane(localToken, remoteChainSelector)) return PolicyConfig.RecordLevel.PolicyChain;

    return PolicyConfig.RecordLevel.None;
  }

  /// @dev Returns whether the token's connection policy enables this chain for the token, including the subset list.
  function _hasPolicyLane(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (bool) {
    bytes32 policyName = s_tokenStates[localToken].connectionPolicyName;

    if (!s_connectionPolicies.policies[policyName].remoteChainSelectors.contains(remoteChainSelector)) {
      return false;
    }

    ConnectionPolicyRecord storage policyRecord = s_policyConnectionRecords[policyName][remoteChainSelector];
    return !policyRecord.onlySubsetTokens || policyRecord.subsetTokens.contains(localToken);
  }

  /// @dev Returns the inline remote pool record if present, otherwise the token's policy record for this chain.
  function _resolveRemotePoolRecord(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (PolicyConfig.RecordLevel level, RemotePoolRecord storage record) {
    if (s_inlineRemotePoolChains[localToken].contains(remoteChainSelector)) {
      return (PolicyConfig.RecordLevel.Inline, s_inlineRemotePoolRecords[localToken][remoteChainSelector]);
    }

    return _resolvePolicyRemotePoolRecord(localToken, remoteChainSelector);
  }

  /// @dev Returns the token's remote pool policy record for this chain, or an empty record with level None if there is
  /// none.
  function _resolvePolicyRemotePoolRecord(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (PolicyConfig.RecordLevel level, RemotePoolRecord storage record) {
    bytes32 policyName = s_tokenStates[localToken].remotePoolPolicyName;

    if (s_remotePoolPolicies.policies[policyName].remoteChainSelectors.contains(remoteChainSelector)) {
      return (PolicyConfig.RecordLevel.PolicyChain, s_policyRemotePoolRecords[policyName][remoteChainSelector]);
    }

    // The zero policy name cannot be configured, so its records are always empty.
    return (PolicyConfig.RecordLevel.None, s_policyRemotePoolRecords[bytes32(0)][remoteChainSelector]);
  }

  // ================================================================
  // │                          Pool data                           │
  // ================================================================

  /// @dev Encodes only token decimals for a legacy destination pool. Otherwise includes the source token and any
  /// additional pool data from the token's provider.
  function _encodePoolData(
    address localToken,
    uint64 remoteChainSelector,
    bytes4 requestedFinalityConfig,
    bytes calldata tokenArgs,
    TokenState storage tokenState,
    bool useLegacyPoolDataOnSend
  ) private view returns (bytes memory) {
    if (useLegacyPoolDataOnSend) {
      return abi.encode(uint256(tokenState.localDecimals));
    }

    bytes memory additionalPoolData;
    if (tokenState.additionalPoolDataProvider != address(0)) {
      additionalPoolData = IAdditionalPoolDataProvider(tokenState.additionalPoolDataProvider)
        .getAdditionalPoolData(localToken, remoteChainSelector, requestedFinalityConfig, tokenArgs);
    }
    // Use the same address encoding that the destination stores for an EVM remote token.
    return abi.encode(uint256(tokenState.localDecimals), abi.encode(localToken), additionalPoolData);
  }

  /// @dev Validates the source token and additional pool data. Returns the source token decimals.
  /// Pool data without a source token requires allowLegacyPoolDataOnReceive for the source pool.
  function _checkPoolData(
    address localToken,
    uint64 remoteChainSelector,
    bytes4 requestedFinalityConfig,
    bytes calldata sourcePoolData,
    TokenState storage tokenState,
    bool allowLegacyPoolDataOnReceive
  ) private view returns (uint8 remoteDecimals) {
    bytes memory sourceToken;
    bytes memory additionalPoolData;
    (remoteDecimals, sourceToken, additionalPoolData) = _decodePoolData(sourcePoolData, tokenState.localDecimals);

    // A trusted multi-token pool can burn different tokens, so its address alone does not identify the source token.
    if (sourceToken.length == 0) {
      if (!allowLegacyPoolDataOnReceive) revert LegacyPoolDataNotAllowed(localToken, remoteChainSelector);
    } else {
      // Legacy messages have no source token, so only this branch resolves the remote token. This keeps legacy
      // messages working without the remote token address provider.
      PolicyConfig.RecordLevel connectionLevel = _resolveConnectionLevel(localToken, remoteChainSelector);
      bytes memory remoteToken = _resolveRemoteToken(localToken, remoteChainSelector, connectionLevel);
      if (keccak256(sourceToken) != keccak256(remoteToken)) {
        revert SourceTokenMismatch(localToken, remoteChainSelector, remoteToken, sourceToken);
      }
    }

    // Validate empty additional data too, so the provider can reject transfers that omit required data.
    address additionalPoolDataProvider = tokenState.additionalPoolDataProvider;
    if (
      additionalPoolDataProvider != address(0)
        && !IAdditionalPoolDataProvider(additionalPoolDataProvider)
          .isAdditionalPoolDataValid(localToken, remoteChainSelector, requestedFinalityConfig, additionalPoolData)
    ) {
      revert InvalidAdditionalPoolData(localToken, remoteChainSelector, additionalPoolData);
    }

    return remoteDecimals;
  }

  /// @dev Decodes legacy pool data or the multi-token format containing decimals, source token, and additional data.
  function _decodePoolData(
    bytes calldata sourcePoolData,
    uint8 localDecimals
  ) private pure returns (uint8 remoteDecimals, bytes memory sourceToken, bytes memory additionalPoolData) {
    // Older pools send no decimals. As in TokenPool, assume the source and destination use the same decimals.
    if (sourcePoolData.length == 0) {
      return (localDecimals, sourceToken, additionalPoolData);
    }

    uint256 decodedDecimals;
    if (sourcePoolData.length == POOL_DATA_DECIMALS_LENGTH) {
      decodedDecimals = abi.decode(sourcePoolData, (uint256));
    } else {
      // Reject truncated data with a custom error before decoding.
      if (sourcePoolData.length < POOL_DATA_WITH_SOURCE_TOKEN_MIN_LENGTH) revert InvalidSourcePoolData(sourcePoolData);
      (decodedDecimals, sourceToken, additionalPoolData) = abi.decode(sourcePoolData, (uint256, bytes, bytes));
      // Require a source token and canonical encoding. abi.decode alone accepts overlapping offsets and trailing data.
      if (
        sourceToken.length == 0
          || keccak256(abi.encode(decodedDecimals, sourceToken, additionalPoolData)) != keccak256(sourcePoolData)
      ) {
        revert InvalidSourcePoolData(sourcePoolData);
      }
    }

    if (decodedDecimals > type(uint8).max) revert InvalidSourcePoolData(sourcePoolData);
    return (uint8(decodedDecimals), sourceToken, additionalPoolData);
  }

  /// @dev Converts the remote amount to local token units using the same scaling and overflow checks as TokenPool.
  function _calculateLocalAmount(
    uint256 remoteAmount,
    uint8 remoteDecimals,
    uint8 localDecimals
  ) private pure returns (uint256) {
    if (remoteDecimals == localDecimals) {
      return remoteAmount;
    }
    if (remoteDecimals > localDecimals) {
      uint8 decimalsDiff = remoteDecimals - localDecimals;
      // A power of ten above 10**77 does not fit in a uint256.
      if (decimalsDiff > 77) {
        revert OverflowDetected(remoteDecimals, localDecimals, remoteAmount);
      }
      // Round down when the local token has fewer decimals.
      return remoteAmount / (10 ** decimalsDiff);
    }

    uint8 diffDecimals = localDecimals - remoteDecimals;
    // Check both exponentiation and multiplication for overflow.
    if (diffDecimals > 77 || remoteAmount > type(uint256).max / (10 ** diffDecimals)) {
      revert OverflowDetected(remoteDecimals, localDecimals, remoteAmount);
    }
    return remoteAmount * (10 ** diffDecimals);
  }

  // ================================================================
  // │                       Resolved views                         │
  // ================================================================

  /// @notice Returns the resolved connection, trusted remote pools, and pause state for each lane.
  /// @dev Returns only pause state for missing lanes. The remote token is empty if its provider is unset, has no code,
  /// reverts, or returns an invalid address. Malformed return data still reverts.
  /// When trustPolicyRemotePools is enabled, inline pools come first, followed by policy pools not already listed.
  /// The inline legacy receive permission applies to pools listed in both records.
  /// @param lanes The lanes to query.
  /// @return connections The connection per lane, aligned with lanes by index.
  function getLaneConnectionBatch(
    IPoolConfigManager.LaneKey[] calldata lanes
  ) external view returns (IPoolConfigManager.LaneConnection[] memory connections) {
    connections = new IPoolConfigManager.LaneConnection[](lanes.length);

    for (uint256 i = 0; i < lanes.length; ++i) {
      address localToken = lanes[i].localToken;
      uint64 remoteChainSelector = lanes[i].remoteChainSelector;
      IPoolConfigManager.LaneConnection memory connection = connections[i];
      TokenState storage tokenState = s_tokenStates[localToken];
      ChainPause storage chainPause = s_chainPauses[remoteChainSelector];

      connection.isGloballyPaused = s_isGloballyPaused;
      connection.isTokenOutboundPaused = tokenState.isOutboundPaused;
      connection.isTokenInboundPaused = tokenState.isInboundPaused;
      connection.isChainOutboundPaused = chainPause.isOutboundPaused;
      connection.isChainInboundPaused = chainPause.isInboundPaused;
      connection.isOutboundPaused =
        connection.isGloballyPaused || connection.isTokenOutboundPaused || connection.isChainOutboundPaused;
      connection.isInboundPaused =
        connection.isGloballyPaused || connection.isTokenInboundPaused || connection.isChainInboundPaused;

      PolicyConfig.RecordLevel level = _resolveConnectionLevel(localToken, remoteChainSelector);
      if (level == PolicyConfig.RecordLevel.None) continue;

      connection.hasConnection = true;
      // Keep the rest of the lane visible if the provider reverts or returns an invalid address.
      connection.remoteToken = level == PolicyConfig.RecordLevel.Inline
        ? s_inlineRemoteTokens[localToken][remoteChainSelector]
        : _tryGetPolicyRemoteToken(localToken, remoteChainSelector);
      (, RemotePoolRecord storage record) = _resolveRemotePoolRecord(localToken, remoteChainSelector);
      connection.useLegacyPoolDataOnSend = record.useLegacyPoolDataOnSend;
      connection.remotePools = _getRemotePools(record);
      if (record.trustPolicyRemotePools) {
        (PolicyConfig.RecordLevel policyLevel, RemotePoolRecord storage policyRecord) =
          _resolvePolicyRemotePoolRecord(localToken, remoteChainSelector);
        if (policyLevel != PolicyConfig.RecordLevel.None) {
          connection.remotePools = _addPolicyRemotePools(connection.remotePools, record, policyRecord);
        }
      }
    }
    return connections;
  }

  /// @notice Returns token profiles, supported chains, and assigned connection and remote pool policies.
  /// @dev Unregistered tokens return zero values. Rate limit, fee, destination overhead, and security policy names are
  /// left unset for the PoolConfigManager to fill from the other stores.
  /// @param localTokens The local tokens.
  /// @return configs The token configs, aligned with localTokens by index.
  function getTokenConfigBatch(
    address[] calldata localTokens
  ) external view returns (IPoolConfigManager.TokenConfig[] memory configs) {
    configs = new IPoolConfigManager.TokenConfig[](localTokens.length);

    for (uint256 i = 0; i < localTokens.length; ++i) {
      address localToken = localTokens[i];
      if (!s_tokens.contains(localToken)) continue;

      TokenState storage tokenState = s_tokenStates[localToken];
      IPoolConfigManager.TokenConfig memory config = configs[i];
      config.isRegistered = true;
      config.localDecimals = tokenState.localDecimals;
      config.lockOrBurnFlavor = tokenState.lockOrBurnFlavor;
      config.allowedFinalityConfig = tokenState.allowedFinalityConfig;
      config.lockOrBurnAddress = tokenState.lockOrBurnAddress;
      config.additionalPoolDataProvider = tokenState.additionalPoolDataProvider;
      config.supportedChains = _getSupportedChains(localToken);
      config.connectionPolicyName = tokenState.connectionPolicyName;
      config.remotePoolPolicyName = tokenState.remotePoolPolicyName;
    }
    return configs;
  }

  /// @notice Checks that every token allows the requested finality.
  /// @dev Finality is checked even for missing lanes, as in TokenPool. Unregistered tokens allow only default finality.
  /// @param lanes The lanes to query.
  /// @param requestedFinalityConfig The requested finality.
  function checkRequestedFinalityBatch(
    IPoolConfigManager.LaneKey[] calldata lanes,
    bytes4 requestedFinalityConfig
  ) external view {
    for (uint256 i = 0; i < lanes.length; ++i) {
      address localToken = lanes[i].localToken;
      bytes4 allowedFinalityConfig = bytes4(0);
      if (s_tokens.contains(localToken)) allowedFinalityConfig = s_tokenStates[localToken].allowedFinalityConfig;

      FinalityCodec._ensureRequestedFinalityAllowed(requestedFinalityConfig, allowedFinalityConfig);
    }
  }

  /// @notice Returns whether each token has a connection to its remote chain, regardless of pause state.
  /// @dev The PoolConfigManager uses this result to clear other stores' view results for missing lanes.
  /// @param lanes The lanes to query.
  /// @return hasLanes Whether each lane has an applicable connection record, in the same order as lanes.
  function hasLaneBatch(
    IPoolConfigManager.LaneKey[] calldata lanes
  ) external view returns (bool[] memory hasLanes) {
    hasLanes = new bool[](lanes.length);

    for (uint256 i = 0; i < lanes.length; ++i) {
      hasLanes[i] =
        _resolveConnectionLevel(lanes[i].localToken, lanes[i].remoteChainSelector) != PolicyConfig.RecordLevel.None;
    }
    return hasLanes;
  }

  /// @notice Reverts if the token has no remote pool record for the chain.
  /// @dev Used by the PoolConfigManager to validate remote pool configuration when enabling a lane.
  /// @param localToken The local token.
  /// @param remoteChainSelector The remote chain.
  function checkLaneRecords(
    address localToken,
    uint64 remoteChainSelector
  ) external view {
    (PolicyConfig.RecordLevel level,) = _resolveRemotePoolRecord(localToken, remoteChainSelector);

    if (level == PolicyConfig.RecordLevel.None) {
      revert RemotePoolRecordNotFound(localToken, remoteChainSelector);
    }
  }

  /// @notice Returns the connection and remote pool record levels for each token and chain.
  /// @dev Resolves the two record types independently, so configuration can be checked before enabling a lane.
  /// Neither policy type supports default records.
  /// @param lanes The lanes to query.
  /// @return connectionLevels The connection record level per lane, aligned with lanes by index.
  /// @return remotePoolLevels The remote pool record level per lane, aligned with lanes by index.
  function getLaneRecordLevelBatch(
    IPoolConfigManager.LaneKey[] calldata lanes
  )
    external
    view
    returns (PolicyConfig.RecordLevel[] memory connectionLevels, PolicyConfig.RecordLevel[] memory remotePoolLevels)
  {
    connectionLevels = new PolicyConfig.RecordLevel[](lanes.length);
    remotePoolLevels = new PolicyConfig.RecordLevel[](lanes.length);

    for (uint256 i = 0; i < lanes.length; ++i) {
      connectionLevels[i] = _resolveConnectionLevel(lanes[i].localToken, lanes[i].remoteChainSelector);
      (remotePoolLevels[i],) = _resolveRemotePoolRecord(lanes[i].localToken, lanes[i].remoteChainSelector);
    }
    return (connectionLevels, remotePoolLevels);
  }

  /// @dev Returns inline chains first, then policy chains enabled for the token. Each chain appears once.
  function _getSupportedChains(
    address localToken
  ) private view returns (uint64[] memory remoteChainSelectors) {
    uint256[] memory inlineChains = s_inlineConnectionChains[localToken].values();
    uint256[] memory policyChains =
      s_connectionPolicies.policies[s_tokenStates[localToken].connectionPolicyName].remoteChainSelectors.values();

    // Move the policy-only chains to the front. This array is a memory copy, so storage does not change.
    uint256 policyOnlyChainCount = 0;
    for (uint256 i = 0; i < policyChains.length; ++i) {
      if (_resolveConnectionLevel(localToken, uint64(policyChains[i])) == PolicyConfig.RecordLevel.PolicyChain) {
        policyChains[policyOnlyChainCount++] = policyChains[i];
      }
    }

    remoteChainSelectors = new uint64[](inlineChains.length + policyOnlyChainCount);
    for (uint256 i = 0; i < inlineChains.length; ++i) {
      remoteChainSelectors[i] = uint64(inlineChains[i]);
    }

    for (uint256 i = 0; i < policyOnlyChainCount; ++i) {
      remoteChainSelectors[inlineChains.length + i] = uint64(policyChains[i]);
    }
    return remoteChainSelectors;
  }

  // ================================================================
  // │                        Stored views                          │
  // ================================================================

  /// @notice Returns a page of registered tokens.
  /// @dev Deregistration can change token ordering. Query pages at the same block to avoid skipping tokens.
  /// @param startIndex The position of the first token to return.
  /// @param maxCount The maximum number of tokens to return.
  /// @return tokens Up to maxCount registered tokens starting at startIndex.
  function getAllConfiguredTokens(
    uint64 startIndex,
    uint64 maxCount
  ) external view returns (address[] memory tokens) {
    uint256 length = s_tokens.length();
    if (startIndex >= length) {
      return tokens;
    }

    uint256 count = maxCount;
    if (count + startIndex > length) {
      count = length - startIndex;
    }

    tokens = new address[](count);
    for (uint256 i = 0; i < count; ++i) {
      tokens[i] = s_tokens.at(startIndex + i);
    }
    return tokens;
  }

  /// @notice Returns stored token profiles and pause flags, including those retained after deregistration.
  /// @param localTokens The local tokens.
  /// @return profiles The profiles, aligned with localTokens by index.
  function getTokenProfiles(
    address[] calldata localTokens
  ) external view returns (TokenProfile[] memory profiles) {
    profiles = new TokenProfile[](localTokens.length);

    for (uint256 i = 0; i < localTokens.length; ++i) {
      TokenState storage tokenState = s_tokenStates[localTokens[i]];
      profiles[i] = TokenProfile({
        isRegistered: s_tokens.contains(localTokens[i]),
        localDecimals: tokenState.localDecimals,
        lockOrBurnFlavor: tokenState.lockOrBurnFlavor,
        lockOrBurnAddress: tokenState.lockOrBurnAddress,
        allowedFinalityConfig: tokenState.allowedFinalityConfig,
        isOutboundPaused: tokenState.isOutboundPaused,
        isInboundPaused: tokenState.isInboundPaused,
        additionalPoolDataProvider: tokenState.additionalPoolDataProvider
      });
    }
    return profiles;
  }

  /// @notice Returns the token's registration status and assigned policies.
  /// @param localToken The local token.
  /// @return isRegistered True if the token is registered in this store.
  /// @return connectionPolicyName The connection policy of the token. Zero if it has none.
  /// @return remotePoolPolicyName The remote pool policy of the token. Zero if it has none.
  function getTokenPolicies(
    address localToken
  ) external view returns (bool isRegistered, bytes32 connectionPolicyName, bytes32 remotePoolPolicyName) {
    TokenState storage tokenState = s_tokenStates[localToken];
    return (s_tokens.contains(localToken), tokenState.connectionPolicyName, tokenState.remotePoolPolicyName);
  }

  /// @notice Returns the inline connection records of a token.
  /// @param localToken The local token.
  /// @return inlineConfigs The inline connection records of the token.
  function getInlineConnectionConfigs(
    address localToken
  ) external view returns (InlineConnectionConfigArgs[] memory inlineConfigs) {
    uint256[] memory remoteChainSelectors = s_inlineConnectionChains[localToken].values();
    inlineConfigs = new InlineConnectionConfigArgs[](remoteChainSelectors.length);

    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChainSelectors[i]);
      inlineConfigs[i] = InlineConnectionConfigArgs({
        localToken: localToken,
        remoteChainSelector: remoteChainSelector,
        remoteToken: s_inlineRemoteTokens[localToken][remoteChainSelector]
      });
    }
    return inlineConfigs;
  }

  /// @notice Returns the inline remote pool records of a token.
  /// @param localToken The local token.
  /// @return inlineConfigs The inline remote pool records of the token.
  function getInlineRemotePoolConfigs(
    address localToken
  ) external view returns (InlineRemotePoolConfigArgs[] memory inlineConfigs) {
    uint256[] memory remoteChainSelectors = s_inlineRemotePoolChains[localToken].values();
    inlineConfigs = new InlineRemotePoolConfigArgs[](remoteChainSelectors.length);

    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChainSelectors[i]);
      RemotePoolRecord storage record = s_inlineRemotePoolRecords[localToken][remoteChainSelector];
      inlineConfigs[i] = InlineRemotePoolConfigArgs({
        localToken: localToken,
        remoteChainSelector: remoteChainSelector,
        trustPolicyRemotePools: record.trustPolicyRemotePools,
        remotePoolConfig: RemotePoolConfig({
          useLegacyPoolDataOnSend: record.useLegacyPoolDataOnSend, remotePools: _getRemotePools(record)
        })
      });
    }
    return inlineConfigs;
  }

  /// @notice Returns the token count, provider, and records of a connection policy.
  /// @param policyName The connection policy.
  /// @return tokenCount The number of tokens that use the policy.
  /// @return remoteTokenAddressProvider The IRemoteTokenAddressProvider of the policy. Zero if it has none.
  /// @return policyConfigs The records of the policy. Empty if the policy does not exist or has no records.
  function getConnectionPolicy(
    bytes32 policyName
  )
    external
    view
    returns (uint32 tokenCount, address remoteTokenAddressProvider, ConnectionPolicyConfigArgs[] memory policyConfigs)
  {
    PolicyConfig.Policy storage policy = s_connectionPolicies.policies[policyName];
    uint256[] memory remoteChainSelectors = policy.remoteChainSelectors.values();
    policyConfigs = new ConnectionPolicyConfigArgs[](remoteChainSelectors.length);

    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChainSelectors[i]);
      policyConfigs[i] = ConnectionPolicyConfigArgs({
        policyName: policyName,
        remoteChainSelector: remoteChainSelector,
        onlySubsetTokens: s_policyConnectionRecords[policyName][remoteChainSelector].onlySubsetTokens
      });
    }
    return (policy.tokenCount, s_remoteTokenAddressProviders[policyName], policyConfigs);
  }

  /// @notice Returns the subset restriction and one page of tokens for a connection policy on a remote chain.
  /// @dev Subset lists can grow with the token count, so this getter uses pages.
  /// Removing a token can change the order. Query pages at the same block to avoid skipping tokens.
  /// @param policyName The connection policy.
  /// @param remoteChainSelector The chain of the policy record.
  /// @param startIndex The position of the first token to return.
  /// @param maxCount The maximum number of tokens to return.
  /// @return onlySubsetTokens Whether the policy restricts this chain to its subset list. False if no record exists.
  /// @return subsetTokens Up to maxCount subset tokens starting at startIndex. Retained but ignored when
  /// onlySubsetTokens is false.
  function getSubsetTokens(
    bytes32 policyName,
    uint64 remoteChainSelector,
    uint64 startIndex,
    uint64 maxCount
  ) external view returns (bool onlySubsetTokens, address[] memory subsetTokens) {
    ConnectionPolicyRecord storage policyRecord = s_policyConnectionRecords[policyName][remoteChainSelector];
    onlySubsetTokens = policyRecord.onlySubsetTokens;
    uint256 length = policyRecord.subsetTokens.length();
    if (startIndex >= length) {
      return (onlySubsetTokens, subsetTokens);
    }

    uint256 count = maxCount;
    if (count + startIndex > length) {
      count = length - startIndex;
    }

    subsetTokens = new address[](count);
    for (uint256 i = 0; i < count; ++i) {
      subsetTokens[i] = policyRecord.subsetTokens.at(startIndex + i);
    }
    return (onlySubsetTokens, subsetTokens);
  }

  /// @notice Returns the token count and the records of a remote pool policy.
  /// @param policyName The remote pool policy.
  /// @return tokenCount The number of tokens that use the policy.
  /// @return policyConfigs The records of the policy. Empty if the policy does not exist or has no records.
  function getRemotePoolPolicy(
    bytes32 policyName
  ) external view returns (uint32 tokenCount, RemotePoolPolicyConfigArgs[] memory policyConfigs) {
    PolicyConfig.Policy storage policy = s_remotePoolPolicies.policies[policyName];
    uint256[] memory remoteChainSelectors = policy.remoteChainSelectors.values();
    policyConfigs = new RemotePoolPolicyConfigArgs[](remoteChainSelectors.length);

    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChainSelectors[i]);
      policyConfigs[i] = RemotePoolPolicyConfigArgs({
        policyName: policyName,
        remoteChainSelector: remoteChainSelector,
        remotePoolAddresses: _getRemotePoolAddresses(s_policyRemotePoolRecords[policyName][remoteChainSelector])
      });
    }
    return (policy.tokenCount, policyConfigs);
  }

  /// @notice Returns the names of all connection policies and remote pool policies.
  /// @return connectionPolicyNames The names of the connection policies.
  /// @return remotePoolPolicyNames The names of the remote pool policies.
  function getPolicyNames()
    external
    view
    returns (bytes32[] memory connectionPolicyNames, bytes32[] memory remotePoolPolicyNames)
  {
    return (s_connectionPolicies.policyNames.values(), s_remotePoolPolicies.policyNames.values());
  }

  function _getRemotePoolAddresses(
    RemotePoolRecord storage record
  ) private view returns (bytes[] memory remotePoolAddresses) {
    bytes32[] memory remotePoolHashes = record.remotePoolHashes.values();
    remotePoolAddresses = new bytes[](remotePoolHashes.length);

    for (uint256 i = 0; i < remotePoolHashes.length; ++i) {
      remotePoolAddresses[i] = s_remotePoolAddresses[remotePoolHashes[i]];
    }
    return remotePoolAddresses;
  }

  function _addPolicyRemotePools(
    IPoolConfigManager.RemotePool[] memory inlineRemotePools,
    RemotePoolRecord storage inlineRecord,
    RemotePoolRecord storage policyRecord
  ) private view returns (IPoolConfigManager.RemotePool[] memory remotePools) {
    bytes32[] memory policyRemotePoolHashes = policyRecord.remotePoolHashes.values();

    // Move the policy-only pools to the front. This array is a memory copy, so storage does not change.
    uint256 policyOnlyPoolCount = 0;
    for (uint256 i = 0; i < policyRemotePoolHashes.length; ++i) {
      if (!inlineRecord.remotePoolHashes.contains(policyRemotePoolHashes[i])) {
        policyRemotePoolHashes[policyOnlyPoolCount++] = policyRemotePoolHashes[i];
      }
    }

    remotePools = new IPoolConfigManager.RemotePool[](inlineRemotePools.length + policyOnlyPoolCount);
    for (uint256 i = 0; i < inlineRemotePools.length; ++i) {
      remotePools[i] = inlineRemotePools[i];
    }

    for (uint256 i = 0; i < policyOnlyPoolCount; ++i) {
      bytes32 remotePoolHash = policyRemotePoolHashes[i];

      remotePools[inlineRemotePools.length + i] = IPoolConfigManager.RemotePool({
        remotePoolAddress: s_remotePoolAddresses[remotePoolHash],
        allowLegacyPoolDataOnReceive: policyRecord.allowLegacyPoolDataOnReceive[remotePoolHash]
      });
    }
    return remotePools;
  }

  function _getRemotePools(
    RemotePoolRecord storage record
  ) private view returns (IPoolConfigManager.RemotePool[] memory remotePools) {
    bytes32[] memory remotePoolHashes = record.remotePoolHashes.values();
    remotePools = new IPoolConfigManager.RemotePool[](remotePoolHashes.length);

    for (uint256 i = 0; i < remotePoolHashes.length; ++i) {
      remotePools[i] = IPoolConfigManager.RemotePool({
        remotePoolAddress: s_remotePoolAddresses[remotePoolHashes[i]],
        allowLegacyPoolDataOnReceive: record.allowLegacyPoolDataOnReceive[remotePoolHashes[i]]
      });
    }
    return remotePools;
  }

  // ================================================================
  // │                            Tokens                            │
  // ================================================================

  /// @notice Deregisters tokens, then registers or updates token profiles.
  /// @dev A registered token's decimals cannot change. Other profile updates preserve its policies and pause flags.
  /// Deregistration requires all assigned policies and inline records to be removed first. The token can then be
  /// registered again with different decimals. The PoolConfigManager applies these updates to every policy store.
  /// @param tokensToRemove The tokens to deregister.
  /// @param profilesToSet The token profiles to add or update.
  function applyTokenProfileUpdates(
    address[] calldata tokensToRemove,
    PolicyConfig.TokenProfileArgs[] calldata profilesToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < tokensToRemove.length; ++i) {
      address localToken = tokensToRemove[i];

      if (!s_tokens.contains(localToken)) revert PolicyConfig.LocalTokenNotRegistered(localToken);
      TokenState storage tokenState = s_tokenStates[localToken];

      if (
        tokenState.connectionPolicyName != bytes32(0) || tokenState.remotePoolPolicyName != bytes32(0)
          || s_inlineConnectionChains[localToken].length() != 0 || s_inlineRemotePoolChains[localToken].length() != 0
      ) {
        revert TokenHasConfig(localToken);
      }

      s_tokens.remove(localToken);
      emit TokenRemoved(localToken);
    }

    for (uint256 i = 0; i < profilesToSet.length; ++i) {
      PolicyConfig.TokenProfileArgs calldata profile = profilesToSet[i];
      address localToken = profile.localToken;
      if (localToken == address(0)) revert ZeroAddressInvalid();
      TokenState storage tokenState = s_tokenStates[localToken];

      if (!s_tokens.contains(localToken)) {
        _validateDecimals(localToken, profile.localDecimals);
      } else if (tokenState.localDecimals != profile.localDecimals) {
        revert InvalidDecimalArgs(localToken, tokenState.localDecimals, profile.localDecimals);
      }
      _validateLockOrBurnAddress(localToken, profile.lockOrBurnFlavor, profile.lockOrBurnAddress);
      _validateProvider(profile.additionalPoolDataProvider);

      s_tokens.add(localToken);
      tokenState.localDecimals = profile.localDecimals;
      tokenState.lockOrBurnFlavor = profile.lockOrBurnFlavor;
      tokenState.lockOrBurnAddress = profile.lockOrBurnAddress;
      // Match TokenPool.setAllowedFinalityConfig, which stores the value without validation.
      tokenState.allowedFinalityConfig = profile.allowedFinalityConfig;
      tokenState.additionalPoolDataProvider = profile.additionalPoolDataProvider;

      emit TokenProfileSet(
        localToken,
        profile.localDecimals,
        profile.lockOrBurnFlavor,
        profile.lockOrBurnAddress,
        profile.allowedFinalityConfig,
        profile.additionalPoolDataProvider
      );
    }
  }

  /// @notice Assigns connection policies to tokens without changing their inline records.
  /// @dev Policy assignment does not validate remote pool configuration. Remove a token from all subset lists before
  /// changing its policy.
  /// @param tokenPoliciesToSet The policy assignments to apply in order. A zero policy name removes the assignment.
  function applyTokenConnectionPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata tokenPoliciesToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < tokenPoliciesToSet.length; ++i) {
      address localToken = tokenPoliciesToSet[i].localToken;
      bytes32 policyName = tokenPoliciesToSet[i].policyName;

      if (!s_tokens.contains(localToken)) revert PolicyConfig.LocalTokenNotRegistered(localToken);
      TokenState storage tokenState = s_tokenStates[localToken];
      bytes32 oldPolicyName = tokenState.connectionPolicyName;
      // Otherwise the old subset entries would come back if the token rejoins the old policy.
      if (policyName != oldPolicyName && _isSubsetToken(localToken, oldPolicyName)) revert TokenHasConfig(localToken);

      PolicyConfig._moveToken(s_connectionPolicies, oldPolicyName, policyName);
      tokenState.connectionPolicyName = policyName;
      emit TokenConnectionPolicySet(localToken, oldPolicyName, policyName);
    }
  }

  /// @notice Assigns remote pool policies to tokens without changing their inline records.
  /// @dev Inline records with trustPolicyRemotePools enabled will trust pools from the new policy.
  /// @param tokenPoliciesToSet The policy assignments to apply in order. A zero policy name removes the assignment.
  function applyTokenRemotePoolPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata tokenPoliciesToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < tokenPoliciesToSet.length; ++i) {
      address localToken = tokenPoliciesToSet[i].localToken;
      bytes32 policyName = tokenPoliciesToSet[i].policyName;

      if (!s_tokens.contains(localToken)) revert PolicyConfig.LocalTokenNotRegistered(localToken);
      TokenState storage tokenState = s_tokenStates[localToken];
      bytes32 oldPolicyName = tokenState.remotePoolPolicyName;

      PolicyConfig._moveToken(s_remotePoolPolicies, oldPolicyName, policyName);
      tokenState.remotePoolPolicyName = policyName;
      emit TokenRemotePoolPolicySet(localToken, oldPolicyName, policyName);
    }
  }

  /// @dev Returns whether the token belongs to any subset list in the connection policy.
  function _isSubsetToken(
    address localToken,
    bytes32 policyName
  ) private view returns (bool) {
    uint256[] memory remoteChainSelectors = s_connectionPolicies.policies[policyName].remoteChainSelectors.values();
    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      if (s_policyConnectionRecords[policyName][uint64(remoteChainSelectors[i])].subsetTokens.contains(localToken)) {
        return true;
      }
    }
    return false;
  }

  /// @dev Checks decimals using the same rules as TokenPool. Since decimals is optional in ERC20, use the supplied
  /// value if the call reverts. Calls to addresses without code and malformed return data still revert.
  function _validateDecimals(
    address localToken,
    uint8 localDecimals
  ) private view {
    try IERC20Metadata(localToken).decimals() returns (uint8 tokenDecimals) {
      if (tokenDecimals != localDecimals) revert InvalidDecimalArgs(localToken, tokenDecimals, localDecimals);
    } catch {}
  }

  /// @dev An unset provider is allowed. Configured providers must have code.
  function _validateProvider(
    address provider
  ) private view {
    if (provider != address(0) && provider.code.length == 0) revert ProviderHasNoCode(provider);
  }

  /// @dev Lock requires a LockBox that supports the token. BurnToAddress accepts any address, matching
  /// BurnToAddressMintTokenPool. All other flavors require a zero lockOrBurnAddress address.
  function _validateLockOrBurnAddress(
    address localToken,
    IPoolConfigManager.LockOrBurnFlavor lockOrBurnFlavor,
    address lockOrBurnAddress
  ) private view {
    if (lockOrBurnFlavor == IPoolConfigManager.LockOrBurnFlavor.Lock) {
      if (lockOrBurnAddress == address(0) || !ILockBox(lockOrBurnAddress).isTokenSupported(localToken)) {
        revert InvalidLockOrBurnAddress(localToken, lockOrBurnFlavor, lockOrBurnAddress);
      }
    } else if (lockOrBurnFlavor != IPoolConfigManager.LockOrBurnFlavor.BurnToAddress && lockOrBurnAddress != address(0))
    {
      revert InvalidLockOrBurnAddress(localToken, lockOrBurnFlavor, lockOrBurnAddress);
    }
  }

  // ================================================================
  // │                       Connection config                      │
  // ================================================================

  /// @notice Removes inline connection records, then adds or updates them.
  /// @dev When the token's connection policy also enables the lane, removal reverts if the policy resolves a different
  /// remote token. An unset or failing provider does not block removal, as outbound transfers would revert anyway, but
  /// malformed return data does. This does not stop later policy or provider updates from changing the remote token.
  /// @param inlineConfigsToRemove The inline records to remove.
  /// @param inlineConfigsToSet The inline records to set.
  function applyInlineConnectionConfigUpdates(
    IPoolConfigManager.LaneKey[] calldata inlineConfigsToRemove,
    InlineConnectionConfigArgs[] calldata inlineConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < inlineConfigsToRemove.length; ++i) {
      address localToken = inlineConfigsToRemove[i].localToken;
      uint64 remoteChainSelector = inlineConfigsToRemove[i].remoteChainSelector;

      // Report a missing inline record before checking the policy's remote token.
      if (!s_inlineConnectionChains[localToken].contains(remoteChainSelector)) {
        revert PolicyConfig.InlineRecordNotFound(localToken, remoteChainSelector);
      }
      // After removal the lane falls back to the policy, which must not point to a different remote token.
      if (_hasPolicyLane(localToken, remoteChainSelector)) {
        bytes memory policyRemoteToken = _tryGetPolicyRemoteToken(localToken, remoteChainSelector);
        if (
          policyRemoteToken.length != 0
            && keccak256(policyRemoteToken) != keccak256(s_inlineRemoteTokens[localToken][remoteChainSelector])
        ) {
          revert PolicyRemoteTokenMismatch(localToken, remoteChainSelector);
        }
      }

      s_inlineConnectionChains[localToken].remove(remoteChainSelector);
      delete s_inlineRemoteTokens[localToken][remoteChainSelector];
      emit InlineConnectionConfigRemoved(localToken, remoteChainSelector);
    }

    for (uint256 i = 0; i < inlineConfigsToSet.length; ++i) {
      address localToken = inlineConfigsToSet[i].localToken;
      uint64 remoteChainSelector = inlineConfigsToSet[i].remoteChainSelector;
      bytes calldata remoteToken = inlineConfigsToSet[i].remoteToken;
      if (!_isValidRemoteToken(remoteToken)) revert InvalidRemoteToken(localToken, remoteChainSelector);

      PolicyConfig._setInlineRecord(s_tokens, s_inlineConnectionChains[localToken], localToken, remoteChainSelector);
      s_inlineRemoteTokens[localToken][remoteChainSelector] = remoteToken;
      emit InlineConnectionConfigSet(localToken, remoteChainSelector, remoteToken);
    }
  }

  /// @notice Removes connection records and unused policies, then adds or updates connection records.
  /// @dev Adding the first record creates a policy. Removing a policy requires no assigned tokens or remaining
  /// records and also clears its provider. Updating a record preserves its subset list; removing it clears the list.
  /// The removal clears the list in one loop, so its gas grows with the list size. Empty a large subset list in
  /// batches with applySubsetTokenUpdates before removing the record.
  /// @param policyConfigsToRemove The records to remove.
  /// @param policiesToRemove The policies to remove.
  /// @param policyConfigsToSet The records to set.
  function applyConnectionPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata policyConfigsToRemove,
    bytes32[] calldata policiesToRemove,
    ConnectionPolicyConfigArgs[] calldata policyConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < policyConfigsToRemove.length; ++i) {
      bytes32 policyName = policyConfigsToRemove[i].policyName;
      uint64 remoteChainSelector = policyConfigsToRemove[i].remoteChainSelector;

      PolicyConfig._removePolicyRecord(s_connectionPolicies, policyName, remoteChainSelector);
      ConnectionPolicyRecord storage policyRecord = s_policyConnectionRecords[policyName][remoteChainSelector];
      // Remove each set entry before deleting the struct, since delete does not clear the set's internal mapping.
      address[] memory subsetTokens = policyRecord.subsetTokens.values();
      for (uint256 j = 0; j < subsetTokens.length; ++j) {
        policyRecord.subsetTokens.remove(subsetTokens[j]);
      }
      delete s_policyConnectionRecords[policyName][remoteChainSelector];
      emit ConnectionPolicyConfigRemoved(policyName, remoteChainSelector);
    }

    for (uint256 i = 0; i < policiesToRemove.length; ++i) {
      bytes32 policyName = policiesToRemove[i];
      PolicyConfig._removePolicy(s_connectionPolicies, policyName);
      delete s_remoteTokenAddressProviders[policyName];
      emit ConnectionPolicyRemoved(policyName);
    }

    for (uint256 i = 0; i < policyConfigsToSet.length; ++i) {
      bytes32 policyName = policyConfigsToSet[i].policyName;
      uint64 remoteChainSelector = policyConfigsToSet[i].remoteChainSelector;
      bool onlySubsetTokens = policyConfigsToSet[i].onlySubsetTokens;

      // Require explicit chain configuration to avoid enabling connections to every chain supported by the Router.
      if (remoteChainSelector == PolicyConfig.DEFAULT_CHAIN_SELECTOR) {
        revert PolicyConfig.ZeroChainSelectorNotAllowed();
      }

      PolicyConfig._setPolicyRecord(s_connectionPolicies, policyName, remoteChainSelector);
      s_policyConnectionRecords[policyName][remoteChainSelector].onlySubsetTokens = onlySubsetTokens;
      emit ConnectionPolicyConfigSet(policyName, remoteChainSelector, onlySubsetTokens);
    }
  }

  /// @notice Removes tokens from connection policy subsets, then adds tokens to them.
  /// @dev Additions require the policy record to exist and the token to be assigned to the policy. Tokens can be added
  /// before onlySubsetTokens is enabled, so the list can be filled before the restriction takes effect.
  /// A token cannot leave its policy until it has been removed from every subset list.
  /// @param tokensToRemove The subset tokens to remove.
  /// @param tokensToAdd The subset tokens to add.
  function applySubsetTokenUpdates(
    SubsetTokenArgs[] calldata tokensToRemove,
    SubsetTokenArgs[] calldata tokensToAdd
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < tokensToRemove.length; ++i) {
      bytes32 policyName = tokensToRemove[i].policyName;
      uint64 remoteChainSelector = tokensToRemove[i].remoteChainSelector;
      address localToken = tokensToRemove[i].localToken;

      if (!s_policyConnectionRecords[policyName][remoteChainSelector].subsetTokens.remove(localToken)) {
        revert SubsetTokenNotFound(policyName, remoteChainSelector, localToken);
      }
      emit SubsetTokenRemoved(policyName, remoteChainSelector, localToken);
    }

    for (uint256 i = 0; i < tokensToAdd.length; ++i) {
      bytes32 policyName = tokensToAdd[i].policyName;
      uint64 remoteChainSelector = tokensToAdd[i].remoteChainSelector;
      address localToken = tokensToAdd[i].localToken;

      if (!s_connectionPolicies.policies[policyName].remoteChainSelectors.contains(remoteChainSelector)) {
        revert PolicyConfig.PolicyRecordNotFound(policyName, remoteChainSelector);
      }
      ConnectionPolicyRecord storage policyRecord = s_policyConnectionRecords[policyName][remoteChainSelector];
      if (s_tokenStates[localToken].connectionPolicyName != policyName) {
        revert TokenNotInPolicy(localToken, policyName);
      }

      if (!policyRecord.subsetTokens.add(localToken)) {
        revert SubsetTokenAlreadyAdded(policyName, remoteChainSelector, localToken);
      }
      emit SubsetTokenAdded(policyName, remoteChainSelector, localToken);
    }
  }

  /// @notice Sets the remote token address provider for each connection policy.
  /// @dev Updates affect all connections that use the policy, except those overridden by inline records.
  /// @param providersToSet The provider updates to apply in order. A zero address removes the provider.
  function applyRemoteTokenAddressProviderUpdates(
    RemoteTokenAddressProviderArgs[] calldata providersToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < providersToSet.length; ++i) {
      bytes32 policyName = providersToSet[i].policyName;
      address remoteTokenAddressProvider = providersToSet[i].remoteTokenAddressProvider;

      if (!s_connectionPolicies.policyNames.contains(policyName)) revert PolicyConfig.PolicyNotFound(policyName);
      _validateProvider(remoteTokenAddressProvider);

      s_remoteTokenAddressProviders[policyName] = remoteTokenAddressProvider;
      emit RemoteTokenAddressProviderSet(policyName, remoteTokenAddressProvider);
    }
  }

  // ================================================================
  // │                      Remote pool config                      │
  // ================================================================

  // Updates set the full list of pools instead of adding and removing them one by one like TokenPool.addRemotePool and
  // removeRemotePool, to stay within the contract size limit. Wrongly removing a pool is an availability risk, while
  // wrongly adding one is a security risk either way. Setting the full list slightly increases the availability risk
  // during normal operations, which we accept.

  /// @notice Removes inline remote pool records, then adds or replaces them.
  /// @dev Keep old remote pools trusted until all their inflight messages have executed.
  /// An empty inline pool list still allows outbound transfers. Inbound transfers then require trustPolicyRemotePools
  /// and a trusted policy pool.
  /// Removing a record with useLegacyPoolDataOnSend set reverts, clear it first. After removal, legacy pool data from
  /// the record's pools is no longer accepted, so wait for inflight legacy messages to execute before removing it.
  /// @param inlineConfigsToRemove The inline records to remove.
  /// @param inlineConfigsToSet The inline records to set.
  function applyInlineRemotePoolConfigUpdates(
    IPoolConfigManager.LaneKey[] calldata inlineConfigsToRemove,
    InlineRemotePoolConfigArgs[] calldata inlineConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < inlineConfigsToRemove.length; ++i) {
      address localToken = inlineConfigsToRemove[i].localToken;
      uint64 remoteChainSelector = inlineConfigsToRemove[i].remoteChainSelector;

      RemotePoolRecord storage record = s_inlineRemotePoolRecords[localToken][remoteChainSelector];
      PolicyConfig._removeInlineRecord(s_inlineRemotePoolChains[localToken], localToken, remoteChainSelector);
      // Policy records don't support legacy pool data. Falling back to the policy while the remote pool is still legacy
      // would burn or lock tokens for messages it can't process. Clearing the flag first confirms the remote pool has
      // been upgraded.
      if (record.useLegacyPoolDataOnSend) revert LegacyRemotePoolRecord(localToken, remoteChainSelector);
      _clearRemotePools(record);
      delete s_inlineRemotePoolRecords[localToken][remoteChainSelector];
      emit InlineRemotePoolConfigRemoved(localToken, remoteChainSelector);
    }

    for (uint256 i = 0; i < inlineConfigsToSet.length; ++i) {
      address localToken = inlineConfigsToSet[i].localToken;
      uint64 remoteChainSelector = inlineConfigsToSet[i].remoteChainSelector;
      bool trustPolicyRemotePools = inlineConfigsToSet[i].trustPolicyRemotePools;
      RemotePoolConfig calldata remotePoolConfig = inlineConfigsToSet[i].remotePoolConfig;

      _validateRemotePools(remoteChainSelector, remotePoolConfig.remotePools);

      PolicyConfig._setInlineRecord(s_tokens, s_inlineRemotePoolChains[localToken], localToken, remoteChainSelector);
      _setInlineRemotePoolRecord(
        s_inlineRemotePoolRecords[localToken][remoteChainSelector], remotePoolConfig, trustPolicyRemotePools
      );
      emit InlineRemotePoolConfigSet(localToken, remoteChainSelector, trustPolicyRemotePools, remotePoolConfig);
    }
  }

  /// @notice Removes remote pool records and unused policies, then adds or replaces remote pool records.
  /// @dev Keep old remote pools trusted until all their inflight messages have executed. Adding the first record
  /// creates a policy. Removing a policy requires no assigned tokens or remaining records.
  /// @param policyConfigsToRemove The records to remove.
  /// @param policiesToRemove The policies to remove.
  /// @param policyConfigsToSet The records to set.
  function applyRemotePoolPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata policyConfigsToRemove,
    bytes32[] calldata policiesToRemove,
    RemotePoolPolicyConfigArgs[] calldata policyConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < policyConfigsToRemove.length; ++i) {
      bytes32 policyName = policyConfigsToRemove[i].policyName;
      uint64 remoteChainSelector = policyConfigsToRemove[i].remoteChainSelector;

      PolicyConfig._removePolicyRecord(s_remotePoolPolicies, policyName, remoteChainSelector);
      _clearRemotePools(s_policyRemotePoolRecords[policyName][remoteChainSelector]);
      delete s_policyRemotePoolRecords[policyName][remoteChainSelector];
      emit RemotePoolPolicyConfigRemoved(policyName, remoteChainSelector);
    }

    for (uint256 i = 0; i < policiesToRemove.length; ++i) {
      bytes32 policyName = policiesToRemove[i];
      PolicyConfig._removePolicy(s_remotePoolPolicies, policyName);
      emit RemotePoolPolicyRemoved(policyName);
    }

    for (uint256 i = 0; i < policyConfigsToSet.length; ++i) {
      bytes32 policyName = policyConfigsToSet[i].policyName;
      uint64 remoteChainSelector = policyConfigsToSet[i].remoteChainSelector;
      bytes[] calldata remotePoolAddresses = policyConfigsToSet[i].remotePoolAddresses;

      // Pool trust must be configured per chain, including chains added later.
      if (remoteChainSelector == PolicyConfig.DEFAULT_CHAIN_SELECTOR) {
        revert PolicyConfig.ZeroChainSelectorNotAllowed();
      }
      // Require at least one pool so the policy does not disable inbound transfers for tokens using this record.
      if (remotePoolAddresses.length == 0) {
        revert PolicyRecordHasNoPools(policyName, remoteChainSelector);
      }
      _validateRemotePoolAddresses(remoteChainSelector, remotePoolAddresses);

      PolicyConfig._setPolicyRecord(s_remotePoolPolicies, policyName, remoteChainSelector);
      _setPolicyRemotePools(s_policyRemotePoolRecords[policyName][remoteChainSelector], remotePoolAddresses);
      emit RemotePoolPolicyConfigSet(policyName, remoteChainSelector, remotePoolAddresses);
    }
  }

  function _validateRemotePools(
    uint64 remoteChainSelector,
    IPoolConfigManager.RemotePool[] calldata remotePools
  ) private pure {
    bytes32[] memory remotePoolHashes = new bytes32[](remotePools.length);
    for (uint256 i = 0; i < remotePools.length; ++i) {
      _validateRemotePool(remoteChainSelector, remotePools[i].remotePoolAddress, remotePoolHashes, i);
    }
  }

  function _validateRemotePoolAddresses(
    uint64 remoteChainSelector,
    bytes[] calldata remotePoolAddresses
  ) private pure {
    bytes32[] memory remotePoolHashes = new bytes32[](remotePoolAddresses.length);
    for (uint256 i = 0; i < remotePoolAddresses.length; ++i) {
      _validateRemotePool(remoteChainSelector, remotePoolAddresses[i], remotePoolHashes, i);
    }
  }

  /// @dev Rejects empty or duplicate pool addresses. Saves the pool's hash for later duplicate checks.
  function _validateRemotePool(
    uint64 remoteChainSelector,
    bytes calldata remotePoolAddress,
    bytes32[] memory remotePoolHashes,
    uint256 index
  ) private pure {
    if (remotePoolAddress.length == 0) revert ZeroAddressInvalid();
    bytes32 remotePoolHash = keccak256(remotePoolAddress);
    for (uint256 i = 0; i < index; ++i) {
      if (remotePoolHashes[i] == remotePoolHash) revert RemotePoolAlreadyAdded(remoteChainSelector, remotePoolAddress);
    }
    remotePoolHashes[index] = remotePoolHash;
  }

  /// @dev The caller must validate the addresses first.
  function _setPolicyRemotePools(
    RemotePoolRecord storage record,
    bytes[] calldata remotePoolAddresses
  ) private {
    _clearRemotePools(record);

    for (uint256 i = 0; i < remotePoolAddresses.length; ++i) {
      _addRemotePool(record, remotePoolAddresses[i]);
    }
  }

  /// @dev The caller must validate the addresses first.
  function _setInlineRemotePoolRecord(
    RemotePoolRecord storage record,
    RemotePoolConfig calldata remotePoolConfig,
    bool trustPolicyRemotePools
  ) private {
    _clearRemotePools(record);
    record.useLegacyPoolDataOnSend = remotePoolConfig.useLegacyPoolDataOnSend;
    record.trustPolicyRemotePools = trustPolicyRemotePools;

    for (uint256 i = 0; i < remotePoolConfig.remotePools.length; ++i) {
      IPoolConfigManager.RemotePool calldata pool = remotePoolConfig.remotePools[i];
      bytes32 remotePoolHash = _addRemotePool(record, pool.remotePoolAddress);
      record.allowLegacyPoolDataOnReceive[remotePoolHash] = pool.allowLegacyPoolDataOnReceive;
    }
  }

  function _addRemotePool(
    RemotePoolRecord storage record,
    bytes calldata remotePoolAddress
  ) private returns (bytes32 remotePoolHash) {
    remotePoolHash = keccak256(remotePoolAddress);
    record.remotePoolHashes.add(remotePoolHash);
    s_remotePoolAddresses[remotePoolHash] = remotePoolAddress;
    return remotePoolHash;
  }

  /// @dev Deleting the struct does not clear the set or the mapping, so each entry is removed individually.
  function _clearRemotePools(
    RemotePoolRecord storage record
  ) private {
    bytes32[] memory remotePoolHashes = record.remotePoolHashes.values();
    for (uint256 i = 0; i < remotePoolHashes.length; ++i) {
      record.remotePoolHashes.remove(remotePoolHashes[i]);
      delete record.allowLegacyPoolDataOnReceive[remotePoolHashes[i]];
    }
  }

  // ================================================================
  // │                            Pause                             │
  // ================================================================

  /// @notice Pauses or unpauses all transfers without changing token or chain pause flags.
  /// @dev Emits an event only when the pause state changes.
  /// @param isPaused True to pause, false to unpause.
  function setGlobalPause(
    bool isPaused
  ) external onlyAuthorizedCallers {
    if (s_isGloballyPaused == isPaused) return;

    s_isGloballyPaused = isPaused;
    emit GlobalPauseSet(isPaused);
  }

  /// @notice Pauses or unpauses token transfers in the specified direction on all chains.
  /// @dev Tokens can be paused before registration. Emits an event only for tokens whose pause state changes.
  /// @param localTokens The tokens.
  /// @param direction Outbound, inbound, or both.
  /// @param isPaused True to pause, false to unpause.
  function setTokenPause(
    address[] calldata localTokens,
    PauseDirection direction,
    bool isPaused
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < localTokens.length; ++i) {
      TokenState storage tokenState = s_tokenStates[localTokens[i]];
      (bool isOutboundPaused, bool isInboundPaused) =
        _resolvePauseFlags(tokenState.isOutboundPaused, tokenState.isInboundPaused, direction, isPaused);
      if (isOutboundPaused == tokenState.isOutboundPaused && isInboundPaused == tokenState.isInboundPaused) continue;

      tokenState.isOutboundPaused = isOutboundPaused;
      tokenState.isInboundPaused = isInboundPaused;
      emit TokenPauseSet(localTokens[i], isOutboundPaused, isInboundPaused);
    }
  }

  /// @notice Pauses or unpauses transfers in the specified direction for all tokens on the given remote chains.
  /// @dev Chains can be paused before any lanes are configured. Only emits an event for chains whose pause state
  /// changes.
  /// @param remoteChainSelectors The remote chains.
  /// @param direction Outbound, inbound, or both.
  /// @param isPaused True to pause, false to unpause.
  function setChainPause(
    uint64[] calldata remoteChainSelectors,
    PauseDirection direction,
    bool isPaused
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      ChainPause storage chainPause = s_chainPauses[remoteChainSelectors[i]];
      (bool isOutboundPaused, bool isInboundPaused) =
        _resolvePauseFlags(chainPause.isOutboundPaused, chainPause.isInboundPaused, direction, isPaused);
      if (isOutboundPaused == chainPause.isOutboundPaused && isInboundPaused == chainPause.isInboundPaused) continue;

      chainPause.isOutboundPaused = isOutboundPaused;
      chainPause.isInboundPaused = isInboundPaused;
      emit ChainPauseSet(remoteChainSelectors[i], isOutboundPaused, isInboundPaused);
    }
  }

  function _resolvePauseFlags(
    bool isOutboundPaused,
    bool isInboundPaused,
    PauseDirection direction,
    bool isPaused
  ) private pure returns (bool newOutboundPaused, bool newInboundPaused) {
    newOutboundPaused = direction == PauseDirection.Inbound ? isOutboundPaused : isPaused;
    newInboundPaused = direction == PauseDirection.Outbound ? isInboundPaused : isPaused;
    return (newOutboundPaused, newInboundPaused);
  }
}
