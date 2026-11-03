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

/// @notice Stores the token profiles, the pause state, the connection records and the remote pool records of a
/// PoolConfigManager. It builds the pool data for lockOrBurn, and checks the pool data for releaseOrMint.
/// @dev Only authorized pool config managers write to this store.
/// A token has a lane to a chain exactly when a connection record serves it there. A transfer on the lane also needs a
/// remote pool record. An inline connection record holds the remote token. A connection policy record opens its chain
/// for every token of the policy, or only for its subset tokens. It takes the remote token from the
/// IRemoteTokenAddressProvider of the policy. A connection policy and a remote pool policy have no policy default
/// record.
/// A remote pool record holds the trusted remote pools and the flags for legacy pool data.
contract ConnectionPolicyStore is ITypeAndVersion, AuthorizedCallers {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;
  using EnumerableSet for EnumerableSet.UintSet;

  error ZeroAddressInvalid();
  error InvalidDecimalArgs(uint8 expected, uint8 actual);
  error InvalidLockOrBurnAddress(
    address localToken, IPoolConfigManager.LockOrBurnFlavor lockOrBurnFlavor, address lockOrBurnAddress
  );
  error InvalidRemoteToken(address localToken, uint64 remoteChainSelector);
  error RemotePoolAlreadyAdded(uint64 remoteChainSelector, bytes remotePoolAddress);
  error PolicyRecordHasNoPools(bytes32 policyName, uint64 remoteChainSelector);
  error TokenHasConfig(address localToken);
  error ProviderHasNoCode(address provider);
  error PolicyRemoteTokenMismatch(address localToken, uint64 remoteChainSelector);
  error OnlySubsetTokensNotSet(bytes32 policyName, uint64 remoteChainSelector);
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
  /// @dev Also stands for the removal of the subset tokens of the record.
  event ConnectionPolicyConfigRemoved(bytes32 indexed policyName, uint64 indexed remoteChainSelector);
  /// @dev Also stands for the removal of the remote token address provider of the policy.
  event ConnectionPolicyRemoved(bytes32 indexed policyName);
  event RemoteTokenAddressProviderSet(bytes32 indexed policyName, address remoteTokenAddressProvider);
  event SubsetTokenAdded(bytes32 indexed policyName, uint64 indexed remoteChainSelector, address indexed localToken);
  event SubsetTokenRemoved(bytes32 indexed policyName, uint64 indexed remoteChainSelector, address indexed localToken);
  /// @dev A set of a new token also registers it.
  event TokenProfileSet(
    address indexed localToken,
    uint8 localDecimals,
    IPoolConfigManager.LockOrBurnFlavor lockOrBurnFlavor,
    address lockOrBurnAddress,
    bytes4 allowedFinalityConfig,
    address additionalPoolDataProvider
  );
  event TokenRemoved(address indexed localToken);
  event TokenConnectionPolicySet(address indexed localToken, bytes32 indexed policyName);
  event TokenRemotePoolPolicySet(address indexed localToken, bytes32 indexed policyName);
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

  /// @dev The direction of a pause. Both covers outbound and inbound.
  enum PauseDirection {
    Outbound,
    Inbound,
    Both
  }

  /// @dev A connection policy record with its key.
  struct ConnectionPolicyConfigArgs {
    bytes32 policyName; // The connection policy. Not zero.
    uint64 remoteChainSelector; // ─╮ The remote chain. Not zero.
    bool onlySubsetTokens; // ──────╯ True opens the chain only for the subset tokens of the record.
  }

  /// @dev An inline connection record with its token and chain.
  struct InlineConnectionConfigArgs {
    address localToken; // ─────────╮ A registered token.
    uint64 remoteChainSelector; // ─╯ The remote chain. Not zero.
    bytes remoteToken; //             The token on the remote chain, abi encoded for EVM chains. Not all zero.
  }

  /// @dev A subset token of a connection policy record.
  struct SubsetTokenArgs {
    bytes32 policyName; // The connection policy.
    uint64 remoteChainSelector; // ─╮ The chain of the policy record.
    address localToken; // ─────────╯ A token of the policy.
  }

  /// @dev The remote token address provider of a connection policy.
  struct RemoteTokenAddressProviderArgs {
    bytes32 policyName; // The connection policy. It must exist.
    address remoteTokenAddressProvider; // The IRemoteTokenAddressProvider of the policy. Zero means none.
  }

  /// @dev The remote pools and the send flag of an inline remote pool record. Legacy pool data is pool data of 0 or 32
  /// bytes, as a TokenPool sends it. It holds no source token and no additional pool data, so it proves no token. A
  /// legacy pool accepts only legacy pool data, so a lane to a remote legacy pool needs useLegacyPoolDataOnSend on. A
  /// remote legacy pool needs allowLegacyPoolDataOnReceive on forever, because its messages never expire. A remote
  /// multi-token pool needs the flag only for the messages that it sent with useLegacyPoolDataOnSend on.
  struct RemotePoolConfig {
    bool useLegacyPoolDataOnSend; // True sends legacy pool data.
    IPoolConfigManager.RemotePool[] remotePools; // The trusted remote pools. No empty or duplicate address.
  }

  /// @dev The remote pools of a policy record, with its policy record key.
  struct RemotePoolPolicyConfigArgs {
    bytes32 policyName; // The remote pool policy. Not zero.
    uint64 remoteChainSelector; // The remote chain. Not zero.
    bytes[] remotePoolAddresses; // The trusted remote pools. At least one. No empty or duplicate address.
  }

  /// @dev RemotePoolConfig with its token and chain.
  struct InlineRemotePoolConfigArgs {
    address localToken; // ──────────╮ A registered token.
    uint64 remoteChainSelector; //   │ The remote chain. Not zero.
    bool trustPolicyRemotePools; // ─╯ True also trusts the pools of the policy record that would serve the token.
    RemotePoolConfig remotePoolConfig; // The values of the record.
  }

  /// @dev The stored profile and pause flags of an address.
  // solhint-disable-next-line gas-struct-packing
  struct TokenProfile {
    bool isRegistered; // ───────────────────────────────────╮ True if the token is registered.
    uint8 localDecimals; //                                  │ The decimals of the local token.
    IPoolConfigManager.LockOrBurnFlavor lockOrBurnFlavor; // │ How the pool moves the token out.
    address lockOrBurnAddress; //                            │ The BurnToAddress target or the Lock box.
    bytes4 allowedFinalityConfig; //                         │ The allowed FTF finality. Zero allows default only.
    bool isOutboundPaused; //                                │ True if the token is paused outbound.
    bool isInboundPaused; // ────────────────────────────────╯ True if the token is paused inbound.
    address additionalPoolDataProvider; // The IAdditionalPoolDataProvider of the token. Zero means none.
  }

  /// @dev The profile, the pause flags and the policies of a token. Zero means no policy. Only the pause functions
  /// change the pause flags, so a pause holds through every other write.
  struct TokenState {
    address lockOrBurnAddress; // ───────────────────────────╮ The BurnToAddress target or the Lock box.
    uint8 localDecimals; //                                  │ The decimals of the local token.
    IPoolConfigManager.LockOrBurnFlavor lockOrBurnFlavor; // │ How the pool moves the token out.
    bytes4 allowedFinalityConfig; //                         │ The allowed FTF finality. Zero allows default only.
    bool isOutboundPaused; //                                │ True stops lockOrBurn of the token to every chain.
    bool isInboundPaused; // ────────────────────────────────╯ True stops releaseOrMint of the token from every chain.
    address additionalPoolDataProvider; // The IAdditionalPoolDataProvider of the token. Zero means none.
    bytes32 connectionPolicyName; // The connection policy of the token.
    bytes32 remotePoolPolicyName; // The remote pool policy of the token.
  }

  /// @dev The values of a connection policy record.
  struct ConnectionPolicyRecord {
    bool onlySubsetTokens; // True serves only the subset tokens. False serves every token of the policy.
    EnumerableSet.AddressSet subsetTokens; // The tokens that the record serves while onlySubsetTokens is true.
  }

  /// @dev The stored values of a remote pool record. Only _setInlineRemotePoolRecord turns a flag on, so the flags stay
  /// false on a policy record.
  struct RemotePoolRecord {
    bool useLegacyPoolDataOnSend; // ─╮ True sends legacy pool data.
    bool trustPolicyRemotePools; // ──╯ As in InlineRemotePoolConfigArgs.
    EnumerableSet.Bytes32Set remotePoolHashes; // The hashes of the trusted remote pools.
    mapping(bytes32 remotePoolHash => bool isAllowed) allowLegacyPoolDataOnReceive; // For each remote pool.
  }

  /// @dev The pause flags of a remote chain.
  struct ChainPause {
    bool isOutboundPaused; // ─╮ True stops lockOrBurn of every token to the chain.
    bool isInboundPaused; // ──╯ True stops releaseOrMint of every token from the chain.
  }

  /// @dev The length of pool data that holds only the decimals. Every TokenPool since 1.5.1 sends and parses it.
  uint256 private constant POOL_DATA_DECIMALS_LENGTH = 32;

  /// @dev Each store keeps its own list of registered tokens, as the stores never call each other.
  EnumerableSet.AddressSet private s_tokens;
  /// @dev Any address has pause flags, also before it is registered. A deregistered token keeps its profile.
  mapping(address localToken => TokenState tokenState) private s_tokenStates;
  /// @dev A remote pool record keeps only the hash of each remote pool. Every record shares this mapping.
  mapping(bytes32 remotePoolHash => bytes remotePoolAddress) private s_remotePoolAddresses;

  // ================================================================
  // │                       Inline storage                         │
  // ================================================================

  /// @dev A token has an inline connection record for a chain exactly when the chain is in the set of the token.
  mapping(address localToken => EnumerableSet.UintSet remoteChainSelectors) private s_inlineConnectionChains;
  mapping(address localToken => mapping(uint64 remoteChainSelector => bytes remoteToken)) private s_inlineRemoteTokens;

  /// @dev A token has an inline remote pool record for a chain exactly when the chain is in the set of the token.
  mapping(address localToken => EnumerableSet.UintSet remoteChainSelectors) private s_inlineRemotePoolChains;
  mapping(address localToken => mapping(uint64 remoteChainSelector => RemotePoolRecord record)) private
    s_inlineRemotePoolRecords;

  // ================================================================
  // │                       Policy storage                         │
  // ================================================================

  PolicyConfig.Policies private s_connectionPolicies;
  mapping(bytes32 policyName => mapping(uint64 remoteChainSelector => ConnectionPolicyRecord record)) private
    s_policyConnectionRecords;
  /// @dev Zero means that the policy has no provider, so its records give no remote token.
  mapping(bytes32 policyName => address remoteTokenAddressProvider) private s_remoteTokenAddressProviders;

  PolicyConfig.Policies private s_remotePoolPolicies;
  mapping(bytes32 policyName => mapping(uint64 remoteChainSelector => RemotePoolRecord record)) private
    s_policyRemotePoolRecords;

  // ================================================================
  // │                        Pause storage                         │
  // ================================================================

  mapping(uint64 remoteChainSelector => ChainPause chainPause) private s_chainPauses;
  /// @dev Separate from the token and chain pause flags, so they keep their state when the global pause ends.
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

  /// @notice Resolves the lane of a lockOrBurn and builds its pool data.
  /// @dev Reverts if a pause stops the transfer, if the token has no connection record or no remote pool record for
  /// the chain, or if the token does not allow the requested finality. A failure of the remote token address provider
  /// or of the additional pool data provider also stops the transfer.
  /// @param localToken The local token.
  /// @param remoteChainSelector The destination chain.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @param tokenArgs The token args of the message. Only the additional pool data provider gets them.
  /// @return lane The remote token, the pool data and the profile values of the token. The fee store and the security
  /// store give the token fee and the hook.
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

    (bool isInline, RemotePoolRecord storage record) = _resolveLane(localToken, remoteChainSelector);
    FinalityCodec._ensureRequestedFinalityAllowed(requestedFinalityConfig, tokenState.allowedFinalityConfig);

    lane.remoteToken = _resolveRemoteToken(localToken, remoteChainSelector);
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

  /// @notice Resolves the lane of a releaseOrMint, checks the source pool and the pool data, and converts the amount
  /// to local decimals. The finality is not checked, because the allowed finality is the config of this chain as a
  /// source, the same as TokenPool.
  /// @dev Reverts if a pause stops the transfer, if the token has no connection record or no remote pool record for
  /// the chain, or if the remote pool record does not trust the source pool. The pool data must pass _checkPoolData.
  /// @param localToken The local token.
  /// @param remoteChainSelector The source chain.
  /// @param sourcePoolAddress The source pool, abi encoded for EVM chains.
  /// @param sourcePoolData The pool data from the source pool.
  /// @param sourceDenominatedAmount The amount in the decimals of the source chain.
  /// @param requestedFinalityConfig The finality that the sender requested. Only the additional pool data provider
  /// gets it.
  /// @return transfer The amount in local decimals and the profile values of the token. The security store gives the
  /// hook.
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

    (bool isInline, RemotePoolRecord storage record) = _resolveLane(localToken, remoteChainSelector);
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

  /// @dev Returns the remote pool record of the lane, and true if it is inline. Only an inline record gives the flags
  /// for legacy pool data to a transfer. Reverts if the token has no connection record or no remote pool record for
  /// the chain.
  function _resolveLane(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (bool isInline, RemotePoolRecord storage record) {
    if (_resolveConnectionLevel(localToken, remoteChainSelector) == PolicyConfig.RecordLevel.None) {
      revert ConnectionRecordNotFound(localToken, remoteChainSelector);
    }

    PolicyConfig.RecordLevel level;
    (level, record) = _resolveRemotePoolRecord(localToken, remoteChainSelector);
    if (level == PolicyConfig.RecordLevel.None) revert RemotePoolRecordNotFound(localToken, remoteChainSelector);
    return (level == PolicyConfig.RecordLevel.Inline, record);
  }

  /// @dev Reverts if the lane does not trust the source pool. Returns the allowLegacyPoolDataOnReceive flag of the
  /// source pool on an inline record, and false for a pool of a policy record.
  function _checkSourcePool(
    address localToken,
    uint64 remoteChainSelector,
    bool isInline,
    RemotePoolRecord storage record,
    bytes calldata sourcePoolAddress
  ) private view returns (bool) {
    bytes32 sourcePoolHash = keccak256(sourcePoolAddress);
    if (record.remotePoolHashes.contains(sourcePoolHash)) {
      return isInline && record.allowLegacyPoolDataOnReceive[sourcePoolHash];
    }

    if (isInline && record.trustPolicyRemotePools) {
      (PolicyConfig.RecordLevel policyLevel, RemotePoolRecord storage policyRecord) =
        _resolvePolicyRemotePoolRecord(localToken, remoteChainSelector);
      if (policyLevel != PolicyConfig.RecordLevel.None && policyRecord.remotePoolHashes.contains(sourcePoolHash)) {
        return false;
      }
    }
    revert InvalidSourcePoolAddress(localToken, remoteChainSelector, sourcePoolAddress);
  }

  /// @dev Returns the remote token of the connection record of the token. A connection policy record asks the
  /// IRemoteTokenAddressProvider of the policy, and a revert of the provider stops the transfer. Reverts if the remote
  /// token is empty or all zero.
  function _resolveRemoteToken(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (bytes memory remoteToken) {
    PolicyConfig.RecordLevel level = _resolveConnectionLevel(localToken, remoteChainSelector);
    if (level == PolicyConfig.RecordLevel.Inline) {
      remoteToken = s_inlineRemoteTokens[localToken][remoteChainSelector];
    } else if (level == PolicyConfig.RecordLevel.PolicyChain) {
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

  /// @dev Returns the answer of the IRemoteTokenAddressProvider of the connection policy of the token, with the cases
  /// of getLaneConnectionBatch.
  function _tryGetPolicyRemoteToken(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (bytes memory remoteToken) {
    address remoteTokenAddressProvider = s_remoteTokenAddressProviders[s_tokenStates[localToken].connectionPolicyName];
    // An address without code returns no data, which fails the decode. Zero has no code.
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

  /// @dev Returns false if the remote token is empty or all zero. No chain has a token at an all-zero address, so a
  /// send to it loses its amount. Zero bytes of the same length have the same hash exactly when the remote token is
  /// empty or all zero.
  function _isValidRemoteToken(
    bytes memory remoteToken
  ) private pure returns (bool) {
    return keccak256(remoteToken) != keccak256(new bytes(remoteToken.length));
  }

  /// @dev An inline connection record overrides the connection policy of the token.
  function _resolveConnectionLevel(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (PolicyConfig.RecordLevel) {
    if (s_inlineConnectionChains[localToken].contains(remoteChainSelector)) return PolicyConfig.RecordLevel.Inline;
    if (_hasPolicyLane(localToken, remoteChainSelector)) return PolicyConfig.RecordLevel.PolicyChain;
    return PolicyConfig.RecordLevel.None;
  }

  /// @dev True if a connection policy record of the policy of the token serves the token on the chain. While
  /// onlySubsetTokens is on, the record serves only its subset tokens.
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

  /// @dev Returns the remote pool record of the token for the chain, and its level. An inline record overrides the
  /// policy. Otherwise the result is that of _resolvePolicyRemotePoolRecord.
  function _resolveRemotePoolRecord(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (PolicyConfig.RecordLevel level, RemotePoolRecord storage record) {
    if (s_inlineRemotePoolChains[localToken].contains(remoteChainSelector)) {
      return (PolicyConfig.RecordLevel.Inline, s_inlineRemotePoolRecords[localToken][remoteChainSelector]);
    }
    return _resolvePolicyRemotePoolRecord(localToken, remoteChainSelector);
  }

  /// @dev Returns the remote pool record of the policy of the token for the chain, and its level. For level None, the
  /// record is empty.
  function _resolvePolicyRemotePoolRecord(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (PolicyConfig.RecordLevel level, RemotePoolRecord storage record) {
    bytes32 policyName = s_tokenStates[localToken].remotePoolPolicyName;
    if (s_remotePoolPolicies.policies[policyName].remoteChainSelectors.contains(remoteChainSelector)) {
      return (PolicyConfig.RecordLevel.PolicyChain, s_policyRemotePoolRecords[policyName][remoteChainSelector]);
    }
    // No write reaches the zero policy name, so this record is empty.
    return (PolicyConfig.RecordLevel.None, s_policyRemotePoolRecords[bytes32(0)][remoteChainSelector]);
  }

  // ================================================================
  // │                          Pool data                           │
  // ================================================================

  /// @dev With useLegacyPoolDataOnSend on, the pool data holds only the decimals, which a legacy pool parses.
  /// Otherwise it also holds the source token and the additional pool data, which is empty for a token without a
  /// provider.
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
    // The source token has the encoding of a remote token in CCIP, which is abi.encode(address) for an EVM chain.
    return abi.encode(uint256(tokenState.localDecimals), abi.encode(localToken), additionalPoolData);
  }

  /// @dev Decodes the pool data, checks the source token and the additional pool data, and returns the remote
  /// decimals. A source token must equal the remote token of the lane. Pool data without the source token needs
  /// allowLegacyPoolDataOnReceive on the source pool.
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

    // A source multi-token pool serves many tokens, so a trusted source pool does not prove which token it burned.
    if (sourceToken.length == 0) {
      if (!allowLegacyPoolDataOnReceive) revert LegacyPoolDataNotAllowed(localToken, remoteChainSelector);
    } else {
      // Only this check needs the remote token, so a source that omits the source token does not depend on the
      // remote token address provider.
      bytes memory remoteToken = _resolveRemoteToken(localToken, remoteChainSelector);
      if (keccak256(sourceToken) != keccak256(remoteToken)) {
        revert SourceTokenMismatch(localToken, remoteChainSelector, remoteToken, sourceToken);
      }
    }

    // Empty data also goes to the provider, so the issuer decides if a source without additional pool data is valid.
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

  /// @dev Decodes the three pool data formats: empty, the decimals only, and the decimals with the source token and
  /// the additional pool data.
  function _decodePoolData(
    bytes calldata sourcePoolData,
    uint8 localDecimals
  ) private pure returns (uint8 remoteDecimals, bytes memory sourceToken, bytes memory additionalPoolData) {
    // Pools older than 1.5.1 send no pool data. This is the same fallback as TokenPool.
    if (sourcePoolData.length == 0) {
      return (localDecimals, sourceToken, additionalPoolData);
    }

    uint256 decodedDecimals;
    if (sourcePoolData.length == POOL_DATA_DECIMALS_LENGTH) {
      decodedDecimals = abi.decode(sourcePoolData, (uint256));
    } else {
      (decodedDecimals, sourceToken, additionalPoolData) = abi.decode(sourcePoolData, (uint256, bytes, bytes));
      // abi.decode also accepts encodings that abi.encode never makes, such as offsets that overlap. Only the exact
      // bytes of abi.encode are valid, with a source token that is not empty.
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

  /// @dev Same as TokenPool, with the local decimals as a parameter. An amount that fails an overflow check cannot
  /// exist on this chain, so the transfer reverts.
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
      // Solidity rounds down, so the pool never mints more than the remote chain sent.
      return remoteAmount / (10 ** decimalsDiff);
    }

    uint8 diffDecimals = localDecimals - remoteDecimals;
    // The second check stops an overflow of the product.
    if (diffDecimals > 77 || remoteAmount > type(uint256).max / (10 ** diffDecimals)) {
      revert OverflowDetected(remoteDecimals, localDecimals, remoteAmount);
    }
    return remoteAmount * (10 ** diffDecimals);
  }

  // ================================================================
  // │                       Resolved views                         │
  // ================================================================

  /// @notice Returns the connection, the remote pool record and the pause state of each lane, for
  /// IPoolConfigManager.getLaneConnectionBatch.
  /// @dev A token without a lane gets only its pause state. A lane of a connection policy record shows the answer of
  /// the IRemoteTokenAddressProvider of the policy. The remote token is empty if the policy has no provider, if the
  /// provider has no code or reverts, or if the answer is all zero. This view reverts if a provider answer does not
  /// decode as bytes, because try does not catch a failed decode. It can also revert if a provider uses all the gas or
  /// returns a very large answer. An inline record with trustPolicyRemotePools on shows its own remote pools first,
  /// then the other remote pools of the policy record. Each remote pool shows once, with the flag that a transfer uses.
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
      // A provider failure leaves the remote token empty, so ops still see the other values of the lane.
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

  /// @notice Returns the config of each token, for IPoolConfigManager.getTokenConfigBatch.
  /// @dev A token that is not registered gets zero values. The fee store and the security store give the policies of
  /// their types, so those names are zero here.
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

  /// @notice Reverts for the first lane whose token does not allow the requested finality, for
  /// IPoolConfigManager.getFeeBatch.
  /// @dev Checks a token without a lane too, the same as TokenPool. A token that is not registered allows only default
  /// finality.
  /// @param lanes The lanes to query.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  function checkRequestedFinalityBatch(
    IPoolConfigManager.LaneKey[] calldata lanes,
    bytes4 requestedFinalityConfig
  ) external view {
    for (uint256 i = 0; i < lanes.length; ++i) {
      address localToken = lanes[i].localToken;
      bytes4 allowedFinalityConfig =
        s_tokens.contains(localToken) ? s_tokenStates[localToken].allowedFinalityConfig : bytes4(0);
      FinalityCodec._ensureRequestedFinalityAllowed(requestedFinalityConfig, allowedFinalityConfig);
    }
  }

  /// @notice Returns whether each token has a lane to its remote chain. Ignores the pause state.
  /// @dev The PoolConfigManager zeros the lane views of the other stores for a lane that does not exist.
  /// @param lanes The lanes to query.
  /// @return hasLanes True if a connection record serves the token on the chain, aligned with lanes by index.
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
  /// @dev The PoolConfigManager calls it after a write that opens a lane.
  /// @param localToken The local token.
  /// @param remoteChainSelector The remote chain.
  function checkLaneRecords(
    address localToken,
    uint64 remoteChainSelector
  ) external view {
    (PolicyConfig.RecordLevel level,) = _resolveRemotePoolRecord(localToken, remoteChainSelector);
    if (level == PolicyConfig.RecordLevel.None) revert RemotePoolRecordNotFound(localToken, remoteChainSelector);
  }

  /// @notice Returns the record levels that each lane would use, also for a lane that does not exist yet.
  /// @dev Lets tooling check the records before a write opens a lane. The connection level shows whether the lane
  /// exists. Neither level is ever PolicyDefault.
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

  /// @dev The chains of the inline connection records of the token, then the chains where only its connection policy
  /// gives it a lane.
  function _getSupportedChains(
    address localToken
  ) private view returns (uint64[] memory remoteChainSelectors) {
    uint256[] memory inlineChains = s_inlineConnectionChains[localToken].values();
    uint256[] memory policyChains =
      s_connectionPolicies.policies[s_tokenStates[localToken].connectionPolicyName].remoteChainSelectors.values();

    uint256 count = inlineChains.length;
    for (uint256 i = 0; i < policyChains.length; ++i) {
      if (_resolveConnectionLevel(localToken, uint64(policyChains[i])) == PolicyConfig.RecordLevel.PolicyChain) {
        ++count;
      }
    }

    remoteChainSelectors = new uint64[](count);
    for (uint256 i = 0; i < inlineChains.length; ++i) {
      remoteChainSelectors[i] = uint64(inlineChains[i]);
    }
    count = inlineChains.length;
    for (uint256 i = 0; i < policyChains.length; ++i) {
      if (_resolveConnectionLevel(localToken, uint64(policyChains[i])) == PolicyConfig.RecordLevel.PolicyChain) {
        remoteChainSelectors[count++] = uint64(policyChains[i]);
      }
    }
    return remoteChainSelectors;
  }

  // ================================================================
  // │                        Stored views                          │
  // ================================================================

  /// @notice Returns one page of the registered tokens, with the paging rules of
  /// TokenAdminRegistry.getAllConfiguredTokens.
  /// @dev A deregistration moves the last token into the place of the removed token. So a reader can miss a token if a
  /// deregistration happens between two pages.
  /// @param startIndex The position of the first token to return.
  /// @param maxCount The maximum number of tokens to return.
  /// @return tokens The tokens. Fewer than maxCount only at the end of the list.
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

  /// @notice Returns the stored profile and the pause flags of each token. A deregistered token shows the profile that
  /// it had at its deregistration.
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

  /// @notice Returns whether the token is registered, and its policies.
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

  /// @notice Returns the token count, the provider and the records of a connection policy.
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

  /// @notice Returns the flag and the subset tokens of a connection policy record.
  /// @param policyName The connection policy.
  /// @param remoteChainSelector The chain of the policy record.
  /// @return onlySubsetTokens True if the record serves only its subset tokens. False if the record does not exist.
  /// @return subsetTokens The subset tokens. They have a lane only while onlySubsetTokens is true.
  function getSubsetTokens(
    bytes32 policyName,
    uint64 remoteChainSelector
  ) external view returns (bool onlySubsetTokens, address[] memory subsetTokens) {
    ConnectionPolicyRecord storage policyRecord = s_policyConnectionRecords[policyName][remoteChainSelector];
    return (policyRecord.onlySubsetTokens, policyRecord.subsetTokens.values());
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
    uint256 count = inlineRemotePools.length;
    for (uint256 i = 0; i < policyRemotePoolHashes.length; ++i) {
      if (!inlineRecord.remotePoolHashes.contains(policyRemotePoolHashes[i])) ++count;
    }

    remotePools = new IPoolConfigManager.RemotePool[](count);
    for (uint256 i = 0; i < inlineRemotePools.length; ++i) {
      remotePools[i] = inlineRemotePools[i];
    }
    count = inlineRemotePools.length;
    for (uint256 i = 0; i < policyRemotePoolHashes.length; ++i) {
      bytes32 remotePoolHash = policyRemotePoolHashes[i];
      if (inlineRecord.remotePoolHashes.contains(remotePoolHash)) continue;
      remotePools[count++] = IPoolConfigManager.RemotePool({
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

  /// @notice Deregisters tokens, then registers or updates tokens. The PoolConfigManager sends the same call to every
  /// store.
  /// @dev A set of a registered token replaces its profile and keeps its policies. Its decimals never change, because
  /// wrong decimals mint a wrong amount. A deregistration needs a token without policies and inline records, so a
  /// wrong address cannot drop live config. A deregistered token can be registered again, also with other decimals.
  /// @param tokensToRemove The tokens to deregister.
  /// @param profilesToSet The profiles of the tokens to register or to update.
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
        revert InvalidDecimalArgs(tokenState.localDecimals, profile.localDecimals);
      }
      _validateLockOrBurnAddress(localToken, profile.lockOrBurnFlavor, profile.lockOrBurnAddress);
      _validateProvider(profile.additionalPoolDataProvider);

      s_tokens.add(localToken);
      tokenState.localDecimals = profile.localDecimals;
      tokenState.lockOrBurnFlavor = profile.lockOrBurnFlavor;
      tokenState.lockOrBurnAddress = profile.lockOrBurnAddress;
      // Any value is accepted, the same as TokenPool.setAllowedFinalityConfig.
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

  /// @notice Sets the connection policy of tokens. Inline records are kept.
  /// @dev A token that joins a policy gets a lane on each chain where the policy serves it, with no lane check. A token
  /// cannot leave its policy while it is a subset token of a record of that policy, so a later join gives it no lane
  /// that a first join would not give.
  /// @param tokenPoliciesToSet The connection policy of each token, applied in order. Zero removes the token from its
  /// policy.
  function applyTokenConnectionPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata tokenPoliciesToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < tokenPoliciesToSet.length; ++i) {
      address localToken = tokenPoliciesToSet[i].localToken;
      bytes32 policyName = tokenPoliciesToSet[i].policyName;

      if (!s_tokens.contains(localToken)) revert PolicyConfig.LocalTokenNotRegistered(localToken);
      TokenState storage tokenState = s_tokenStates[localToken];
      bytes32 oldPolicyName = tokenState.connectionPolicyName;
      if (policyName != oldPolicyName && _isSubsetToken(localToken, oldPolicyName)) revert TokenHasConfig(localToken);

      PolicyConfig._moveToken(s_connectionPolicies, oldPolicyName, policyName);
      tokenState.connectionPolicyName = policyName;
      emit TokenConnectionPolicySet(localToken, policyName);
    }
  }

  /// @notice Sets the remote pool policy of tokens. Inline records are kept. An inline record with
  /// trustPolicyRemotePools on trusts the remote pools of the new policy.
  /// @param tokenPoliciesToSet The remote pool policy of each token, applied in order. Zero removes the token from its
  /// policy.
  function applyTokenRemotePoolPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata tokenPoliciesToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < tokenPoliciesToSet.length; ++i) {
      address localToken = tokenPoliciesToSet[i].localToken;
      bytes32 policyName = tokenPoliciesToSet[i].policyName;

      if (!s_tokens.contains(localToken)) revert PolicyConfig.LocalTokenNotRegistered(localToken);
      TokenState storage tokenState = s_tokenStates[localToken];

      PolicyConfig._moveToken(s_remotePoolPolicies, tokenState.remotePoolPolicyName, policyName);
      tokenState.remotePoolPolicyName = policyName;
      emit TokenRemotePoolPolicySet(localToken, policyName);
    }
  }

  /// @dev True if the token is a subset token of a record of the connection policy.
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

  /// @dev The same check as the TokenPool constructor. The decimals function is optional in ERC20, so a token whose
  /// call reverts keeps the given decimals. An address without code, or an answer that does not decode, reverts with
  /// no error data, because try does not catch a failed decode.
  function _validateDecimals(
    address localToken,
    uint8 localDecimals
  ) private view {
    try IERC20Metadata(localToken).decimals() returns (uint8 tokenDecimals) {
      if (tokenDecimals != localDecimals) revert InvalidDecimalArgs(tokenDecimals, localDecimals);
    } catch {}
  }

  /// @dev Zero means no provider. Any other provider needs code, because a call to an address without code reverts the
  /// transfer.
  function _validateProvider(
    address provider
  ) private view {
    if (provider != address(0) && provider.code.length == 0) revert ProviderHasNoCode(provider);
  }

  /// @dev BurnToAddress accepts any address, the same as BurnToAddressMintTokenPool. Lock needs a lock box that
  /// supports the token, the same as LockReleaseTokenPool. The other flavors need zero.
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

  /// @notice Removes inline connection records, then sets inline connection records.
  /// @dev After a removal, the token uses its connection policy on that chain, if the policy serves it there. So a
  /// removal reverts if the policy gives the token other remote token bytes, because a send with a wrong remote token
  /// loses its amount. A policy without a provider, a provider that reverts or has no code, and an all-zero answer do
  /// not block the removal, because such a lane stops every send. An answer that does not decode as bytes reverts the
  /// removal with no error data. The check reads the policy records and the provider at the time of the call, so it
  /// covers only this call. A later policy or provider change can still change the remote token of the lane.
  /// @param inlineConfigsToRemove The inline records to remove.
  /// @param inlineConfigsToSet The inline records to set.
  function applyInlineConnectionConfigUpdates(
    IPoolConfigManager.LaneKey[] calldata inlineConfigsToRemove,
    InlineConnectionConfigArgs[] calldata inlineConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < inlineConfigsToRemove.length; ++i) {
      address localToken = inlineConfigsToRemove[i].localToken;
      uint64 remoteChainSelector = inlineConfigsToRemove[i].remoteChainSelector;

      // A missing record reverts here, before the policy check can report a mismatch.
      if (!s_inlineConnectionChains[localToken].contains(remoteChainSelector)) {
        revert PolicyConfig.InlineRecordNotFound(localToken, remoteChainSelector);
      }
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

  /// @notice Removes connection policy records, then removes connection policies, then sets connection policy records.
  /// @dev The first record set under a new name creates the policy. Only a policy without tokens and records can be
  /// removed, and its removal also removes its provider. A record set keeps the subset tokens of the record. A record
  /// removal deletes them, so a record that is set again starts with none.
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
      // Delete does not clear a set inside a struct, so each subset token is removed first.
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

      // A policy default record would open a lane to every chain that the Router serves.
      if (remoteChainSelector == PolicyConfig.DEFAULT_CHAIN_SELECTOR) {
        revert PolicyConfig.ZeroChainSelectorNotAllowed();
      }

      PolicyConfig._setPolicyRecord(s_connectionPolicies, policyName, remoteChainSelector);
      s_policyConnectionRecords[policyName][remoteChainSelector].onlySubsetTokens = onlySubsetTokens;
      emit ConnectionPolicyConfigSet(policyName, remoteChainSelector, onlySubsetTokens);
    }
  }

  /// @notice Removes subset tokens, then adds subset tokens. A subset token has a lane on the chain of its record while
  /// it uses the policy and the record has onlySubsetTokens on.
  /// @dev An add needs a policy record with onlySubsetTokens on, and a token of the policy. A subset token cannot leave
  /// the policy until its subset entries are removed.
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
      if (!policyRecord.onlySubsetTokens) revert OnlySubsetTokensNotSet(policyName, remoteChainSelector);
      if (s_tokenStates[localToken].connectionPolicyName != policyName) {
        revert TokenNotInPolicy(localToken, policyName);
      }

      if (!policyRecord.subsetTokens.add(localToken)) {
        revert SubsetTokenAlreadyAdded(policyName, remoteChainSelector, localToken);
      }
      emit SubsetTokenAdded(policyName, remoteChainSelector, localToken);
    }
  }

  /// @notice Sets the remote token address provider of connection policies. A new provider gives the remote token of
  /// every lane of the policy at once.
  /// @dev A provider needs code.
  /// @param providersToSet The provider of each policy, applied in order. Zero removes the provider of the policy.
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

  /// @notice Removes inline remote pool records, then sets inline remote pool records. Keep an old remote pool until
  /// every message that it sent is executed.
  /// @dev An inline record that trusts no remote pool works outbound only, the same as TokenPool.
  /// @param inlineConfigsToRemove The inline records to remove.
  /// @param inlineConfigsToSet The inline records to set.
  function applyInlineRemotePoolConfigUpdates(
    IPoolConfigManager.LaneKey[] calldata inlineConfigsToRemove,
    InlineRemotePoolConfigArgs[] calldata inlineConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < inlineConfigsToRemove.length; ++i) {
      address localToken = inlineConfigsToRemove[i].localToken;
      uint64 remoteChainSelector = inlineConfigsToRemove[i].remoteChainSelector;

      PolicyConfig._removeInlineRecord(s_inlineRemotePoolChains[localToken], localToken, remoteChainSelector);
      _clearRemotePools(s_inlineRemotePoolRecords[localToken][remoteChainSelector]);
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

  /// @notice Removes remote pool policy records, then removes remote pool policies, then sets remote pool policy
  /// records. Keep an old remote pool until every message that it sent is executed.
  /// @dev The first record set under a new name creates the policy. Only a policy without tokens and records can be
  /// removed.
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

      // A policy default record would trust its remote pools on every chain, also on a chain added later.
      if (remoteChainSelector == PolicyConfig.DEFAULT_CHAIN_SELECTOR) {
        revert PolicyConfig.ZeroChainSelectorNotAllowed();
      }
      // A policy record serves many tokens, so it must trust at least one remote pool.
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

  /// @dev An empty or duplicate remote pool address reverts, the same as TokenPool. remotePoolHashes holds the hashes
  /// of the remote pools before index, and gets the hash of this remote pool at index.
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

  /// @dev Replaces the remote pools of a policy record. The caller runs _validateRemotePoolAddresses first, so each
  /// remote pool is not empty and is new to the cleared set.
  function _setPolicyRemotePools(
    RemotePoolRecord storage record,
    bytes[] calldata remotePoolAddresses
  ) private {
    _clearRemotePools(record);

    for (uint256 i = 0; i < remotePoolAddresses.length; ++i) {
      _addRemotePool(record, remotePoolAddresses[i]);
    }
  }

  /// @dev Replaces every value of an inline record. The caller runs _validateRemotePools first, so each remote pool is
  /// not empty and is new to the cleared set.
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

  /// @dev Delete does not clear a set or a mapping inside a struct, so each remote pool and its flag are removed first.
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

  /// @notice Turns the global pause on or off. The global pause stops every token on every chain in both directions.
  /// The token and chain pause flags do not change.
  /// @dev A call that changes nothing emits nothing.
  /// @param isPaused True to pause, false to unpause.
  function setGlobalPause(
    bool isPaused
  ) external onlyAuthorizedCallers {
    if (s_isGloballyPaused == isPaused) return;

    s_isGloballyPaused = isPaused;
    emit GlobalPauseSet(isPaused);
  }

  /// @notice Pauses or unpauses tokens on every chain. The other direction does not change.
  /// @dev Any address is accepted, also before the token is registered. A token that does not change emits nothing.
  /// @param localTokens The tokens.
  /// @param direction The direction to change.
  /// @param isPaused True to pause, false to unpause.
  function setTokenPause(
    address[] calldata localTokens,
    PauseDirection direction,
    bool isPaused
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < localTokens.length; ++i) {
      TokenState storage tokenState = s_tokenStates[localTokens[i]];
      (bool isOutboundPaused, bool isInboundPaused) =
        _getPauseFlags(tokenState.isOutboundPaused, tokenState.isInboundPaused, direction, isPaused);
      if (isOutboundPaused == tokenState.isOutboundPaused && isInboundPaused == tokenState.isInboundPaused) continue;

      tokenState.isOutboundPaused = isOutboundPaused;
      tokenState.isInboundPaused = isInboundPaused;
      emit TokenPauseSet(localTokens[i], isOutboundPaused, isInboundPaused);
    }
  }

  /// @notice Pauses or unpauses remote chains for every token. The other direction does not change.
  /// @dev Any chain is accepted, also one without lanes. A chain that does not change emits nothing.
  /// @param remoteChainSelectors The remote chains.
  /// @param direction The direction to change.
  /// @param isPaused True to pause, false to unpause.
  function setChainPause(
    uint64[] calldata remoteChainSelectors,
    PauseDirection direction,
    bool isPaused
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      ChainPause storage chainPause = s_chainPauses[remoteChainSelectors[i]];
      (bool isOutboundPaused, bool isInboundPaused) =
        _getPauseFlags(chainPause.isOutboundPaused, chainPause.isInboundPaused, direction, isPaused);
      if (isOutboundPaused == chainPause.isOutboundPaused && isInboundPaused == chainPause.isInboundPaused) continue;

      chainPause.isOutboundPaused = isOutboundPaused;
      chainPause.isInboundPaused = isInboundPaused;
      emit ChainPauseSet(remoteChainSelectors[i], isOutboundPaused, isInboundPaused);
    }
  }

  /// @dev Returns the pause flags after a change of the given direction.
  function _getPauseFlags(
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
