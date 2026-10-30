// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../libraries/PolicyConfig.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Stores the FeePolicy and the DestinationOverheadPolicy for a PoolConfigManager.
/// @dev Only authorized pool config managers write to this store.
/// A fee record holds the USD fees and the token fees to be charged at lockOrBurn. A disabled fee record makes the
/// OnRamp use the FeeQuoter. A destination overhead record holds the gas and the bytes needed for releaseOrMint on the
/// destination chain.
contract FeePolicyStore is ITypeAndVersion, AuthorizedCallers {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;
  using EnumerableSet for EnumerableSet.UintSet;

  error ZeroAddressInvalid();
  error TokenHasConfig(address localToken);
  error InvalidFeePolicyConfig(bytes32 policyName, uint64 remoteChainSelector, FeeConfig feeConfig);
  error InvalidInlineFeeConfig(address localToken, uint64 remoteChainSelector, FeeConfig feeConfig);
  error InvalidDestOverheadPolicyConfig(
    bytes32 policyName, uint64 remoteChainSelector, DestOverheadConfig destOverheadConfig
  );
  error InvalidInlineDestOverheadConfig(
    address localToken, uint64 remoteChainSelector, DestOverheadConfig destOverheadConfig
  );
  error FeeRecordNotFound(address localToken, uint64 remoteChainSelector);
  error DestOverheadRecordNotFound(address localToken, uint64 remoteChainSelector);
  error HasLanesLengthMismatch(uint256 lanesLength, uint256 hasLanesLength);

  /// @dev A set of a new token also registers it.
  event TokenProfileSet(address indexed localToken);
  event TokenRemoved(address indexed localToken);
  event TokenFeePolicySet(address indexed localToken, bytes32 indexed policyName);
  event TokenDestOverheadPolicySet(address indexed localToken, bytes32 indexed policyName);
  event FeePolicyConfigSet(bytes32 indexed policyName, uint64 indexed remoteChainSelector, FeeConfig feeConfig);
  event FeePolicyConfigRemoved(bytes32 indexed policyName, uint64 indexed remoteChainSelector);
  event FeePolicyRemoved(bytes32 indexed policyName);
  event InlineFeeConfigSet(address indexed localToken, uint64 indexed remoteChainSelector, FeeConfig feeConfig);
  event InlineFeeConfigRemoved(address indexed localToken, uint64 indexed remoteChainSelector);
  event DestOverheadPolicyConfigSet(
    bytes32 indexed policyName, uint64 indexed remoteChainSelector, DestOverheadConfig destOverheadConfig
  );
  event DestOverheadPolicyConfigRemoved(bytes32 indexed policyName, uint64 indexed remoteChainSelector);
  event DestOverheadPolicyRemoved(bytes32 indexed policyName);
  event InlineDestOverheadConfigSet(
    address indexed localToken, uint64 indexed remoteChainSelector, DestOverheadConfig destOverheadConfig
  );
  event InlineDestOverheadConfigRemoved(address indexed localToken, uint64 indexed remoteChainSelector);

  /// @dev The values of a fee record.
  struct FeeConfig {
    uint32 finalityFeeUSDCents; // ───────╮ Fee for a transfer with default finality, in multiples of 0.01 USD.
    uint32 fastFinalityFeeUSDCents; //    │ Fee for an FTF transfer, in multiples of 0.01 USD.
    uint16 finalityTransferFeeBps; //     │ Token fee for a transfer with default finality, in basis points.
    uint16 fastFinalityTransferFeeBps; // │ Token fee for an FTF transfer, in basis points.
    bool isEnabled; // ───────────────────╯ False makes the OnRamp use the FeeQuoter.
  }

  /// @dev The values of a dest overhead record.
  struct DestOverheadConfig {
    uint32 destGasOverhead; // ──╮ Gas for releaseOrMint on the destination chain. Not zero.
    uint32 destBytesOverhead; // ╯ Data availability bytes of the pool data. At least 192.
  }

  /// @dev FeeConfig with its policy record key.
  struct FeePolicyConfigArgs {
    bytes32 policyName; // The fee policy. Not zero.
    uint64 remoteChainSelector; // The remote chain. Zero names the policy default record.
    FeeConfig feeConfig; // The values of the record.
  }

  /// @dev FeeConfig with its token and chain.
  struct InlineFeeConfigArgs {
    address localToken; // ─────────╮ A registered token.
    uint64 remoteChainSelector; // ─╯ The remote chain. Not zero.
    FeeConfig feeConfig; //           The values of the record.
  }

  /// @dev DestOverheadConfig with its policy record key.
  struct DestOverheadPolicyConfigArgs {
    bytes32 policyName; // The dest overhead policy. Not zero.
    uint64 remoteChainSelector; // The remote chain. Zero names the policy default record.
    DestOverheadConfig destOverheadConfig; // The values of the record.
  }

  /// @dev DestOverheadConfig with its token and chain.
  struct InlineDestOverheadConfigArgs {
    address localToken; // ─────────╮ A registered token.
    uint64 remoteChainSelector; // ─╯ The remote chain. Not zero.
    DestOverheadConfig destOverheadConfig; // The values of the record.
  }

  /// @dev The policies of a token. Zero means no policy.
  struct TokenState {
    bytes32 feePolicyName; // The fee policy of the token.
    bytes32 destOverheadPolicyName; // The dest overhead policy of the token.
  }

  /// @dev The division factor for bps. Each bps fee must be below it.
  uint256 private constant BPS_DIVIDER = 10_000;
  /// @dev The pool data length of an EVM source token with empty additional pool data, so that every lane with an
  /// enabled fee record can send the source token.
  uint32 private constant MIN_DEST_BYTES_OVERHEAD = 192;

  /// @dev Each store keeps its own list of registered tokens, as the stores never call each other.
  EnumerableSet.AddressSet private s_tokens;
  /// @dev Only registered tokens have policies.
  mapping(address localToken => TokenState tokenState) private s_tokenStates;

  PolicyConfig.Policies private s_feePolicies;
  /// @dev Fee policy record details. Chain selector zero holds the policy default.
  mapping(bytes32 policyName => mapping(uint64 remoteChainSelector => FeeConfig feeConfig)) private s_policyFeeConfigs;

  /// @dev An inline fee record exists when a token x chain is in the set.
  mapping(address localToken => EnumerableSet.UintSet remoteChainSelectors) private s_inlineFeeChains;
  mapping(address localToken => mapping(uint64 remoteChainSelector => FeeConfig feeConfig)) private s_inlineFeeConfigs;

  PolicyConfig.Policies private s_destOverheadPolicies;
  /// @dev Overheads policy record details. Chain selector zero holds the policy default.
  mapping(bytes32 policyName => mapping(uint64 remoteChainSelector => DestOverheadConfig destOverheadConfig)) private
    s_policyDestOverheadConfigs;
  /// @dev An inline overhead record exists when a token x chain is in the set.
  mapping(address localToken => EnumerableSet.UintSet remoteChainSelectors) private s_inlineDestOverheadChains;
  mapping(address localToken => mapping(uint64 remoteChainSelector => DestOverheadConfig destOverheadConfig)) private
    s_inlineDestOverheadConfigs;

  /// @param authorizedCallers The pool config managers that can write to this store.
  constructor(
    address[] memory authorizedCallers
  ) AuthorizedCallers(authorizedCallers) {}

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "FeePolicyStore 2.1.0-dev";
  }

  // ================================================================
  // │                       Fee resolution                         │
  // ================================================================

  /// @notice Returns the token fee of a lockOrBurn.
  /// @dev Reverts if the token has no fee record or no dest overhead record for the chain.
  /// @param localToken The local token.
  /// @param remoteChainSelector The destination chain.
  /// @param amount The amount before the token fee.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @param chargeTokenFee False for the V1 path, which charges no token fee.
  /// @return feeAmount The token fee in local token units. Zero if the fee record is disabled.
  function resolveTokenFee(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount,
    bytes4 requestedFinalityConfig,
    bool chargeTokenFee
  ) external view returns (uint256 feeAmount) {
    FeeConfig storage feeConfig = _getLaneFeeConfig(localToken, remoteChainSelector);
    if (!chargeTokenFee || !feeConfig.isEnabled) {
      return 0;
    }

    uint16 transferFeeBps = requestedFinalityConfig == FinalityCodec.WAIT_FOR_FINALITY_FLAG
      ? feeConfig.finalityTransferFeeBps
      : feeConfig.fastFinalityTransferFeeBps;
    return (amount * transferFeeBps) / BPS_DIVIDER;
  }

  /// @notice Reverts if the token has no fee record or no dest overhead record for the chain.
  /// @dev The PoolConfigManager calls it after a write that opens a lane.
  /// @param localToken The local token.
  /// @param remoteChainSelector The remote chain.
  function checkLaneRecords(
    address localToken,
    uint64 remoteChainSelector
  ) external view {
    _getLaneFeeConfig(localToken, remoteChainSelector);
  }

  /// @dev Returns the fee record. Reverts if the fee record or the dest overhead record is missing.
  function _getLaneFeeConfig(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (FeeConfig storage feeConfig) {
    PolicyConfig.RecordLevel level;
    (level, feeConfig) = _resolveFeeConfig(localToken, remoteChainSelector);
    if (level == PolicyConfig.RecordLevel.None) revert FeeRecordNotFound(localToken, remoteChainSelector);
    (level,) = _resolveDestOverheadConfig(localToken, remoteChainSelector);
    if (level == PolicyConfig.RecordLevel.None) revert DestOverheadRecordNotFound(localToken, remoteChainSelector);
    return feeConfig;
  }

  /// @notice Returns the fee of each lane for the requested finality, for IPoolConfigManager.getFeeBatch.
  /// @dev The connection store checks the requested finality and gives hasLanes before this call.
  /// @param lanes The lanes to query.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @param hasLanes Whether each lane exists, aligned with lanes by index.
  /// @return fees The fee per lane, aligned with lanes by index. Zero with isEnabled false for a lane that does not
  /// exist, has no enabled fee record or has no dest overhead record.
  function getFeeBatch(
    IPoolConfigManager.LaneKey[] calldata lanes,
    bytes4 requestedFinalityConfig,
    bool[] calldata hasLanes
  ) external view returns (IPoolConfigManager.LaneFee[] memory fees) {
    if (lanes.length != hasLanes.length) revert HasLanesLengthMismatch(lanes.length, hasLanes.length);

    bool isFastFinality = requestedFinalityConfig != FinalityCodec.WAIT_FOR_FINALITY_FLAG;
    fees = new IPoolConfigManager.LaneFee[](lanes.length);
    for (uint256 i = 0; i < lanes.length; ++i) {
      if (!hasLanes[i]) continue;
      (bool isEnabled, FeeConfig storage feeConfig, DestOverheadConfig storage destOverheadConfig) =
        _resolveEnabledFee(lanes[i]);
      if (!isEnabled) continue;

      fees[i] = IPoolConfigManager.LaneFee({
        feeUSDCents: isFastFinality ? feeConfig.fastFinalityFeeUSDCents : feeConfig.finalityFeeUSDCents,
        destGasOverhead: destOverheadConfig.destGasOverhead,
        destBytesOverhead: destOverheadConfig.destBytesOverhead,
        tokenFeeBps: isFastFinality ? feeConfig.fastFinalityTransferFeeBps : feeConfig.finalityTransferFeeBps,
        isEnabled: true
      });
    }
    return fees;
  }

  /// @notice Returns the fee config of each lane, for IPoolConfigManager.getTokenTransferFeeConfigBatch.
  /// @dev The connection store gives hasLanes before this call.
  /// @param lanes The lanes to query.
  /// @param hasLanes Whether each lane exists, aligned with lanes by index.
  /// @return feeConfigs The fee config per lane, aligned with lanes by index. A disabled zero config for a lane that
  /// does not exist, has no enabled fee record or has no dest overhead record.
  function getTokenTransferFeeConfigBatch(
    IPoolConfigManager.LaneKey[] calldata lanes,
    bool[] calldata hasLanes
  ) external view returns (IPoolV2.TokenTransferFeeConfig[] memory feeConfigs) {
    if (lanes.length != hasLanes.length) revert HasLanesLengthMismatch(lanes.length, hasLanes.length);

    feeConfigs = new IPoolV2.TokenTransferFeeConfig[](lanes.length);
    for (uint256 i = 0; i < lanes.length; ++i) {
      if (!hasLanes[i]) continue;
      (bool isEnabled, FeeConfig storage feeConfig, DestOverheadConfig storage destOverheadConfig) =
        _resolveEnabledFee(lanes[i]);
      if (!isEnabled) continue;

      feeConfigs[i] = IPoolV2.TokenTransferFeeConfig({
        destGasOverhead: destOverheadConfig.destGasOverhead,
        destBytesOverhead: destOverheadConfig.destBytesOverhead,
        finalityFeeUSDCents: feeConfig.finalityFeeUSDCents,
        fastFinalityFeeUSDCents: feeConfig.fastFinalityFeeUSDCents,
        finalityTransferFeeBps: feeConfig.finalityTransferFeeBps,
        fastFinalityTransferFeeBps: feeConfig.fastFinalityTransferFeeBps,
        isEnabled: true
      });
    }
    return feeConfigs;
  }

  /// @dev Returns both records, and whether the lane has an enabled fee record and a dest overhead record.
  function _resolveEnabledFee(
    IPoolConfigManager.LaneKey calldata lane
  ) private view returns (bool isEnabled, FeeConfig storage feeConfig, DestOverheadConfig storage destOverheadConfig) {
    PolicyConfig.RecordLevel feeLevel;
    (feeLevel, feeConfig) = _resolveFeeConfig(lane.localToken, lane.remoteChainSelector);
    PolicyConfig.RecordLevel destOverheadLevel;
    (destOverheadLevel, destOverheadConfig) = _resolveDestOverheadConfig(lane.localToken, lane.remoteChainSelector);
    isEnabled = feeLevel != PolicyConfig.RecordLevel.None && feeConfig.isEnabled
      && destOverheadLevel != PolicyConfig.RecordLevel.None;
    return (isEnabled, feeConfig, destOverheadConfig);
  }

  /// @notice Returns the record levels that each lane would use, also for a lane that does not exist yet.
  /// @dev Lets tooling check the records before a write opens a lane.
  /// @param lanes The lanes to query.
  /// @return feeLevels The fee record level per lane, aligned with lanes by index.
  /// @return destOverheadLevels The dest overhead record level per lane, aligned with lanes by index.
  function getLaneRecordLevelBatch(
    IPoolConfigManager.LaneKey[] calldata lanes
  )
    external
    view
    returns (PolicyConfig.RecordLevel[] memory feeLevels, PolicyConfig.RecordLevel[] memory destOverheadLevels)
  {
    feeLevels = new PolicyConfig.RecordLevel[](lanes.length);
    destOverheadLevels = new PolicyConfig.RecordLevel[](lanes.length);
    for (uint256 i = 0; i < lanes.length; ++i) {
      (feeLevels[i],) = _resolveFeeConfig(lanes[i].localToken, lanes[i].remoteChainSelector);
      (destOverheadLevels[i],) = _resolveDestOverheadConfig(lanes[i].localToken, lanes[i].remoteChainSelector);
    }
    return (feeLevels, destOverheadLevels);
  }

  /// @dev Returns the fee record of the token for the chain, and its level. An inline record overrides the policy. For
  /// level None, the record is the empty inline slot, so callers must check the level first.
  function _resolveFeeConfig(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (PolicyConfig.RecordLevel level, FeeConfig storage feeConfig) {
    if (s_inlineFeeChains[localToken].contains(remoteChainSelector)) {
      return (PolicyConfig.RecordLevel.Inline, s_inlineFeeConfigs[localToken][remoteChainSelector]);
    }
    bytes32 policyName = s_tokenStates[localToken].feePolicyName;
    level = PolicyConfig._resolvePolicyRecordLevel(
      s_feePolicies.policies[policyName].remoteChainSelectors, remoteChainSelector
    );
    if (level == PolicyConfig.RecordLevel.PolicyChain) {
      return (level, s_policyFeeConfigs[policyName][remoteChainSelector]);
    }
    if (level == PolicyConfig.RecordLevel.PolicyDefault) {
      return (level, s_policyFeeConfigs[policyName][PolicyConfig.DEFAULT_CHAIN_SELECTOR]);
    }
    return (level, s_inlineFeeConfigs[localToken][remoteChainSelector]);
  }

  /// @dev Same as _resolveFeeConfig, for the dest overhead record.
  function _resolveDestOverheadConfig(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (PolicyConfig.RecordLevel level, DestOverheadConfig storage destOverheadConfig) {
    if (s_inlineDestOverheadChains[localToken].contains(remoteChainSelector)) {
      return (PolicyConfig.RecordLevel.Inline, s_inlineDestOverheadConfigs[localToken][remoteChainSelector]);
    }
    bytes32 policyName = s_tokenStates[localToken].destOverheadPolicyName;
    level = PolicyConfig._resolvePolicyRecordLevel(
      s_destOverheadPolicies.policies[policyName].remoteChainSelectors, remoteChainSelector
    );
    if (level == PolicyConfig.RecordLevel.PolicyChain) {
      return (level, s_policyDestOverheadConfigs[policyName][remoteChainSelector]);
    }
    if (level == PolicyConfig.RecordLevel.PolicyDefault) {
      return (level, s_policyDestOverheadConfigs[policyName][PolicyConfig.DEFAULT_CHAIN_SELECTOR]);
    }
    return (level, s_inlineDestOverheadConfigs[localToken][remoteChainSelector]);
  }

  // ================================================================
  // │                         Config views                         │
  // ================================================================

  /// @notice Returns whether the token is registered, and its policies.
  /// @param localToken The local token.
  /// @return isRegistered True if the token is registered in this store.
  /// @return feePolicyName The fee policy of the token. Zero if it has none.
  /// @return destOverheadPolicyName The dest overhead policy of the token. Zero if it has none.
  function getTokenPolicies(
    address localToken
  ) external view returns (bool isRegistered, bytes32 feePolicyName, bytes32 destOverheadPolicyName) {
    TokenState storage tokenState = s_tokenStates[localToken];
    return (s_tokens.contains(localToken), tokenState.feePolicyName, tokenState.destOverheadPolicyName);
  }

  /// @notice Returns the policies of each token, for IPoolConfigManager.getTokenConfigBatch.
  /// @param localTokens The local tokens.
  /// @return feePolicyNames The fee policy per token, aligned with localTokens by index. Zero if it has none.
  /// @return destOverheadPolicyNames The dest overhead policy per token, aligned with localTokens by index. Zero if it
  /// has none.
  function getTokenPolicyNamesBatch(
    address[] calldata localTokens
  ) external view returns (bytes32[] memory feePolicyNames, bytes32[] memory destOverheadPolicyNames) {
    feePolicyNames = new bytes32[](localTokens.length);
    destOverheadPolicyNames = new bytes32[](localTokens.length);
    for (uint256 i = 0; i < localTokens.length; ++i) {
      TokenState storage tokenState = s_tokenStates[localTokens[i]];
      feePolicyNames[i] = tokenState.feePolicyName;
      destOverheadPolicyNames[i] = tokenState.destOverheadPolicyName;
    }
    return (feePolicyNames, destOverheadPolicyNames);
  }

  /// @notice Returns the names of all fee policies and dest overhead policies.
  /// @return feePolicyNames The names of the fee policies.
  /// @return destOverheadPolicyNames The names of the dest overhead policies.
  function getPolicyNames()
    external
    view
    returns (bytes32[] memory feePolicyNames, bytes32[] memory destOverheadPolicyNames)
  {
    return (s_feePolicies.policyNames.values(), s_destOverheadPolicies.policyNames.values());
  }

  /// @notice Returns the token count and the records of a fee policy.
  /// @param policyName The fee policy.
  /// @return tokenCount The number of tokens that use the policy.
  /// @return policyConfigs The records of the policy. Empty if the policy does not exist or has no records.
  function getFeePolicy(
    bytes32 policyName
  ) external view returns (uint32 tokenCount, FeePolicyConfigArgs[] memory policyConfigs) {
    PolicyConfig.Policy storage policy = s_feePolicies.policies[policyName];
    uint256[] memory remoteChainSelectors = policy.remoteChainSelectors.values();
    policyConfigs = new FeePolicyConfigArgs[](remoteChainSelectors.length);
    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChainSelectors[i]);
      policyConfigs[i] = FeePolicyConfigArgs({
        policyName: policyName,
        remoteChainSelector: remoteChainSelector,
        feeConfig: s_policyFeeConfigs[policyName][remoteChainSelector]
      });
    }
    return (policy.tokenCount, policyConfigs);
  }

  /// @notice Returns the inline fee records of a token.
  /// @param localToken The local token.
  /// @return inlineConfigs The inline fee records of the token.
  function getInlineFeeConfigs(
    address localToken
  ) external view returns (InlineFeeConfigArgs[] memory inlineConfigs) {
    uint256[] memory remoteChainSelectors = s_inlineFeeChains[localToken].values();
    inlineConfigs = new InlineFeeConfigArgs[](remoteChainSelectors.length);
    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChainSelectors[i]);
      inlineConfigs[i] = InlineFeeConfigArgs({
        localToken: localToken,
        remoteChainSelector: remoteChainSelector,
        feeConfig: s_inlineFeeConfigs[localToken][remoteChainSelector]
      });
    }
    return inlineConfigs;
  }

  /// @notice Returns the token count and the records of a dest overhead policy.
  /// @param policyName The dest overhead policy.
  /// @return tokenCount The number of tokens that use the policy.
  /// @return policyConfigs The records of the policy. Empty if the policy does not exist or has no records.
  function getDestOverheadPolicy(
    bytes32 policyName
  ) external view returns (uint32 tokenCount, DestOverheadPolicyConfigArgs[] memory policyConfigs) {
    PolicyConfig.Policy storage policy = s_destOverheadPolicies.policies[policyName];
    uint256[] memory remoteChainSelectors = policy.remoteChainSelectors.values();
    policyConfigs = new DestOverheadPolicyConfigArgs[](remoteChainSelectors.length);
    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChainSelectors[i]);
      policyConfigs[i] = DestOverheadPolicyConfigArgs({
        policyName: policyName,
        remoteChainSelector: remoteChainSelector,
        destOverheadConfig: s_policyDestOverheadConfigs[policyName][remoteChainSelector]
      });
    }
    return (policy.tokenCount, policyConfigs);
  }

  /// @notice Returns the inline dest overhead records of a token.
  /// @param localToken The local token.
  /// @return inlineConfigs The inline dest overhead records of the token.
  function getInlineDestOverheadConfigs(
    address localToken
  ) external view returns (InlineDestOverheadConfigArgs[] memory inlineConfigs) {
    uint256[] memory remoteChainSelectors = s_inlineDestOverheadChains[localToken].values();
    inlineConfigs = new InlineDestOverheadConfigArgs[](remoteChainSelectors.length);
    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChainSelectors[i]);
      inlineConfigs[i] = InlineDestOverheadConfigArgs({
        localToken: localToken,
        remoteChainSelector: remoteChainSelector,
        destOverheadConfig: s_inlineDestOverheadConfigs[localToken][remoteChainSelector]
      });
    }
    return inlineConfigs;
  }

  // ================================================================
  // │                            Tokens                            │
  // ================================================================

  /// @notice Deregisters tokens, then registers tokens. The PoolConfigManager sends the same call to every store.
  /// @dev This store keeps no profile values, so it only reads the token of each profile. A deregistration needs a
  /// token without policies and inline records, so a wrong address cannot drop live config.
  /// @param tokensToRemove The tokens to deregister.
  /// @param profilesToSet The profiles of the tokens to register. A registered token stays as it is.
  function applyTokenProfileUpdates(
    address[] calldata tokensToRemove,
    PolicyConfig.TokenProfileArgs[] calldata profilesToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < tokensToRemove.length; ++i) {
      address localToken = tokensToRemove[i];
      if (!s_tokens.contains(localToken)) revert PolicyConfig.LocalTokenNotRegistered(localToken);
      TokenState storage tokenState = s_tokenStates[localToken];
      if (
        tokenState.feePolicyName != bytes32(0) || tokenState.destOverheadPolicyName != bytes32(0)
          || s_inlineFeeChains[localToken].length() != 0 || s_inlineDestOverheadChains[localToken].length() != 0
      ) {
        revert TokenHasConfig(localToken);
      }

      s_tokens.remove(localToken);
      emit TokenRemoved(localToken);
    }

    for (uint256 i = 0; i < profilesToSet.length; ++i) {
      address localToken = profilesToSet[i].localToken;
      if (localToken == address(0)) revert ZeroAddressInvalid();

      s_tokens.add(localToken);
      emit TokenProfileSet(localToken);
    }
  }

  /// @notice Sets the fee policy of tokens. Inline records are kept.
  /// @param tokenPoliciesToSet The fee policy of each token, applied in order. Zero removes the token from its policy.
  function applyTokenFeePolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata tokenPoliciesToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < tokenPoliciesToSet.length; ++i) {
      address localToken = tokenPoliciesToSet[i].localToken;
      bytes32 policyName = tokenPoliciesToSet[i].policyName;
      if (!s_tokens.contains(localToken)) revert PolicyConfig.LocalTokenNotRegistered(localToken);
      TokenState storage tokenState = s_tokenStates[localToken];

      PolicyConfig._moveToken(s_feePolicies, tokenState.feePolicyName, policyName);
      tokenState.feePolicyName = policyName;
      emit TokenFeePolicySet(localToken, policyName);
    }
  }

  /// @notice Sets the dest overhead policy of tokens. Inline records are kept.
  /// @param tokenPoliciesToSet The dest overhead policy of each token, applied in order. Zero removes the token from
  /// its policy.
  function applyTokenDestOverheadPolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata tokenPoliciesToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < tokenPoliciesToSet.length; ++i) {
      address localToken = tokenPoliciesToSet[i].localToken;
      bytes32 policyName = tokenPoliciesToSet[i].policyName;
      if (!s_tokens.contains(localToken)) revert PolicyConfig.LocalTokenNotRegistered(localToken);
      TokenState storage tokenState = s_tokenStates[localToken];

      PolicyConfig._moveToken(s_destOverheadPolicies, tokenState.destOverheadPolicyName, policyName);
      tokenState.destOverheadPolicyName = policyName;
      emit TokenDestOverheadPolicySet(localToken, policyName);
    }
  }

  // ================================================================
  // │                          Fee config                          │
  // ================================================================

  /// @notice Removes fee policy records, then removes fee policies, then sets fee policy records.
  /// @dev The first record set under a new name creates the policy. Only a policy without tokens and records can be
  /// removed.
  /// @param policyConfigsToRemove The records to remove.
  /// @param policiesToRemove The policies to remove.
  /// @param policyConfigsToSet The records to set.
  function applyFeePolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata policyConfigsToRemove,
    bytes32[] calldata policiesToRemove,
    FeePolicyConfigArgs[] calldata policyConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < policyConfigsToRemove.length; ++i) {
      bytes32 policyName = policyConfigsToRemove[i].policyName;
      uint64 remoteChainSelector = policyConfigsToRemove[i].remoteChainSelector;

      PolicyConfig._removePolicyRecord(s_feePolicies, policyName, remoteChainSelector);
      delete s_policyFeeConfigs[policyName][remoteChainSelector];
      emit FeePolicyConfigRemoved(policyName, remoteChainSelector);
    }

    for (uint256 i = 0; i < policiesToRemove.length; ++i) {
      PolicyConfig._removePolicy(s_feePolicies, policiesToRemove[i]);
      emit FeePolicyRemoved(policiesToRemove[i]);
    }

    for (uint256 i = 0; i < policyConfigsToSet.length; ++i) {
      FeePolicyConfigArgs calldata args = policyConfigsToSet[i];
      if (!_isValidFeeConfig(args.feeConfig)) {
        revert InvalidFeePolicyConfig(args.policyName, args.remoteChainSelector, args.feeConfig);
      }

      PolicyConfig._setPolicyRecord(s_feePolicies, args.policyName, args.remoteChainSelector);
      s_policyFeeConfigs[args.policyName][args.remoteChainSelector] = args.feeConfig;
      emit FeePolicyConfigSet(args.policyName, args.remoteChainSelector, args.feeConfig);
    }
  }

  /// @notice Removes inline fee records, then sets inline fee records.
  /// @param inlineConfigsToRemove The inline records to remove.
  /// @param inlineConfigsToSet The inline records to set.
  function applyInlineFeeConfigUpdates(
    IPoolConfigManager.LaneKey[] calldata inlineConfigsToRemove,
    InlineFeeConfigArgs[] calldata inlineConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < inlineConfigsToRemove.length; ++i) {
      address localToken = inlineConfigsToRemove[i].localToken;
      uint64 remoteChainSelector = inlineConfigsToRemove[i].remoteChainSelector;

      PolicyConfig._removeInlineRecord(s_inlineFeeChains[localToken], localToken, remoteChainSelector);
      delete s_inlineFeeConfigs[localToken][remoteChainSelector];
      emit InlineFeeConfigRemoved(localToken, remoteChainSelector);
    }

    for (uint256 i = 0; i < inlineConfigsToSet.length; ++i) {
      InlineFeeConfigArgs calldata args = inlineConfigsToSet[i];
      if (!_isValidFeeConfig(args.feeConfig)) {
        revert InvalidInlineFeeConfig(args.localToken, args.remoteChainSelector, args.feeConfig);
      }

      PolicyConfig._setInlineRecord(
        s_tokens, s_inlineFeeChains[args.localToken], args.localToken, args.remoteChainSelector
      );
      s_inlineFeeConfigs[args.localToken][args.remoteChainSelector] = args.feeConfig;
      emit InlineFeeConfigSet(args.localToken, args.remoteChainSelector, args.feeConfig);
    }
  }

  /// @dev Each bps fee must be below BPS_DIVIDER, the same as TokenPool. A disabled config must be all zero, so no fee
  /// value stays hidden in storage.
  function _isValidFeeConfig(
    FeeConfig calldata feeConfig
  ) private pure returns (bool) {
    if (!feeConfig.isEnabled) {
      return feeConfig.finalityFeeUSDCents == 0 && feeConfig.fastFinalityFeeUSDCents == 0
        && feeConfig.finalityTransferFeeBps == 0 && feeConfig.fastFinalityTransferFeeBps == 0;
    }
    return feeConfig.finalityTransferFeeBps < BPS_DIVIDER && feeConfig.fastFinalityTransferFeeBps < BPS_DIVIDER;
  }

  // ================================================================
  // │                     Dest overhead config                     │
  // ================================================================

  /// @notice Removes dest overhead policy records, then removes dest overhead policies, then sets dest overhead policy
  /// records.
  /// @dev The first record set under a new name creates the policy. Only a policy without tokens and records can be
  /// removed.
  /// @param policyConfigsToRemove The records to remove.
  /// @param policiesToRemove The policies to remove.
  /// @param policyConfigsToSet The records to set.
  function applyDestOverheadPolicyConfigUpdates(
    PolicyConfig.PolicyRecordKey[] calldata policyConfigsToRemove,
    bytes32[] calldata policiesToRemove,
    DestOverheadPolicyConfigArgs[] calldata policyConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < policyConfigsToRemove.length; ++i) {
      bytes32 policyName = policyConfigsToRemove[i].policyName;
      uint64 remoteChainSelector = policyConfigsToRemove[i].remoteChainSelector;

      PolicyConfig._removePolicyRecord(s_destOverheadPolicies, policyName, remoteChainSelector);
      delete s_policyDestOverheadConfigs[policyName][remoteChainSelector];
      emit DestOverheadPolicyConfigRemoved(policyName, remoteChainSelector);
    }

    for (uint256 i = 0; i < policiesToRemove.length; ++i) {
      PolicyConfig._removePolicy(s_destOverheadPolicies, policiesToRemove[i]);
      emit DestOverheadPolicyRemoved(policiesToRemove[i]);
    }

    for (uint256 i = 0; i < policyConfigsToSet.length; ++i) {
      DestOverheadPolicyConfigArgs calldata args = policyConfigsToSet[i];
      if (!_isValidDestOverheadConfig(args.destOverheadConfig)) {
        revert InvalidDestOverheadPolicyConfig(args.policyName, args.remoteChainSelector, args.destOverheadConfig);
      }

      PolicyConfig._setPolicyRecord(s_destOverheadPolicies, args.policyName, args.remoteChainSelector);
      s_policyDestOverheadConfigs[args.policyName][args.remoteChainSelector] = args.destOverheadConfig;
      emit DestOverheadPolicyConfigSet(args.policyName, args.remoteChainSelector, args.destOverheadConfig);
    }
  }

  /// @notice Removes inline dest overhead records, then sets inline dest overhead records.
  /// @param inlineConfigsToRemove The inline records to remove.
  /// @param inlineConfigsToSet The inline records to set.
  function applyInlineDestOverheadConfigUpdates(
    IPoolConfigManager.LaneKey[] calldata inlineConfigsToRemove,
    InlineDestOverheadConfigArgs[] calldata inlineConfigsToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < inlineConfigsToRemove.length; ++i) {
      address localToken = inlineConfigsToRemove[i].localToken;
      uint64 remoteChainSelector = inlineConfigsToRemove[i].remoteChainSelector;

      PolicyConfig._removeInlineRecord(s_inlineDestOverheadChains[localToken], localToken, remoteChainSelector);
      delete s_inlineDestOverheadConfigs[localToken][remoteChainSelector];
      emit InlineDestOverheadConfigRemoved(localToken, remoteChainSelector);
    }

    for (uint256 i = 0; i < inlineConfigsToSet.length; ++i) {
      InlineDestOverheadConfigArgs calldata args = inlineConfigsToSet[i];
      if (!_isValidDestOverheadConfig(args.destOverheadConfig)) {
        revert InvalidInlineDestOverheadConfig(args.localToken, args.remoteChainSelector, args.destOverheadConfig);
      }

      PolicyConfig._setInlineRecord(
        s_tokens, s_inlineDestOverheadChains[args.localToken], args.localToken, args.remoteChainSelector
      );
      s_inlineDestOverheadConfigs[args.localToken][args.remoteChainSelector] = args.destOverheadConfig;
      emit InlineDestOverheadConfigSet(args.localToken, args.remoteChainSelector, args.destOverheadConfig);
    }
  }

  /// @dev The gas overhead must be non-zero, the same as TokenPool. The bytes overhead must be at least
  /// MIN_DEST_BYTES_OVERHEAD.
  function _isValidDestOverheadConfig(
    DestOverheadConfig calldata destOverheadConfig
  ) private pure returns (bool) {
    return destOverheadConfig.destGasOverhead != 0 && destOverheadConfig.destBytesOverhead >= MIN_DEST_BYTES_OVERHEAD;
  }
}
