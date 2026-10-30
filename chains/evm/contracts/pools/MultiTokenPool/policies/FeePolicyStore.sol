// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../libraries/PolicyConfig.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Holds the fee type and the destination overhead type of the tokens of a pool config manager. With an enabled
/// fee record, the OnRamp prices a send with both records.
/// @dev Only authorized pool config managers write to this store. A transfer only reads it.
///
/// A fee record holds the USD fees and the token fees of a send. A disabled fee record makes the OnRamp use the
/// FeeQuoter. A destination overhead record holds the gas and the bytes that a send needs on the destination chain. A
/// send needs a destination overhead record also while its fee record is disabled.
contract FeePolicyStore is ITypeAndVersion, AuthorizedCallers {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;
  using EnumerableSet for EnumerableSet.UintSet;

  error ZeroAddressInvalid();
  error TokenHasConfig(address localToken);
  error InvalidTransferFeeBps(uint256 bps);
  error DisabledNonZeroFeeConfig(FeeConfig feeConfig);
  error InvalidDestOverheads(uint64 remoteChainSelector, uint32 destGasOverhead, uint32 destBytesOverhead);
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

  /// @notice A fee policy record to set, and the shape of a fee policy record in the views.
  struct FeePolicyConfigArgs {
    bytes32 policyName; // The fee policy. Not zero.
    uint64 remoteChainSelector; // ───────╮ The remote chain. Zero names the policy default record.
    uint32 finalityFeeUSDCents; //        │ Fee for a transfer with default finality, in multiples of 0.01 USD.
    uint32 fastFinalityFeeUSDCents; //    │ Fee for an FTF transfer, in multiples of 0.01 USD.
    uint16 finalityTransferFeeBps; //     │ Token fee for a transfer with default finality, in basis points.
    uint16 fastFinalityTransferFeeBps; // │ Token fee for an FTF transfer, in basis points.
    bool isEnabled; // ───────────────────╯ False makes the OnRamp use the FeeQuoter.
  }

  /// @notice An inline fee record to set, and the shape of an inline fee record in the views.
  struct InlineFeeConfigArgs {
    address localToken; // ───────────────╮ A registered token.
    uint64 remoteChainSelector; //        │ The remote chain. Not zero.
    uint32 finalityFeeUSDCents; // ───────╯ Fee for a transfer with default finality, in multiples of 0.01 USD.
    uint32 fastFinalityFeeUSDCents; // ───╮ Fee for an FTF transfer, in multiples of 0.01 USD.
    uint16 finalityTransferFeeBps; //     │ Token fee for a transfer with default finality, in basis points.
    uint16 fastFinalityTransferFeeBps; // │ Token fee for an FTF transfer, in basis points.
    bool isEnabled; // ───────────────────╯ False makes the OnRamp use the FeeQuoter.
  }

  /// @notice A destination overhead policy record to set, and the shape of such a record in the views.
  struct DestOverheadPolicyConfigArgs {
    bytes32 policyName; // The destination overhead policy. Not zero.
    uint64 remoteChainSelector; // ─╮ The remote chain. Zero names the policy default record.
    uint32 destGasOverhead; //      │ Gas to execute the token transfer on the destination chain. Not zero.
    uint32 destBytesOverhead; // ───╯ Data availability bytes of the pool data. At least 192.
  }

  /// @notice An inline destination overhead record to set, and the shape of such a record in the views.
  struct InlineDestOverheadConfigArgs {
    address localToken; // ─────────╮ A registered token.
    uint64 remoteChainSelector; //  │ The remote chain. Not zero.
    uint32 destGasOverhead; // ─────╯ Gas to execute the token transfer on the destination chain. Not zero.
    uint32 destBytesOverhead; // Data availability bytes of the pool data. At least 192.
  }

  /// @notice The values of a fee record.
  struct FeeConfig {
    uint32 finalityFeeUSDCents; // ───────╮ Fee for a transfer with default finality, in multiples of 0.01 USD.
    uint32 fastFinalityFeeUSDCents; //    │ Fee for an FTF transfer, in multiples of 0.01 USD.
    uint16 finalityTransferFeeBps; //     │ Token fee for a transfer with default finality, in basis points.
    uint16 fastFinalityTransferFeeBps; // │ Token fee for an FTF transfer, in basis points.
    bool isEnabled; // ───────────────────╯ False makes the OnRamp use the FeeQuoter.
  }

  /// @notice The values of a destination overhead record.
  struct DestOverheadConfig {
    uint32 destGasOverhead; // ──╮ Gas to execute the token transfer on the destination chain.
    uint32 destBytesOverhead; // ╯ Data availability bytes of the pool data.
  }

  /// @notice The policies of one token in this store. Zero means that the token has no policy of the type.
  struct TokenState {
    bytes32 feePolicyName; // The fee policy of the token.
    bytes32 destOverheadPolicyName; // The destination overhead policy of the token.
  }

  /// @dev The divisor of a fee in basis points. Each fee in basis points must be below it.
  uint256 private constant BPS_DIVIDER = 10_000;
  /// @dev The length of pool data that carries an EVM source token and empty additional pool data. With an enabled fee
  /// record, the OnRamp rejects pool data that is longer than destBytesOverhead. So a lane with an enabled fee record
  /// can always send the source token with empty additional pool data.
  uint32 private constant MIN_DEST_BYTES_OVERHEAD = 192;

  /// @dev The registered tokens. This store keeps its own list, because the stores never call each other.
  EnumerableSet.AddressSet private s_tokens;
  /// @dev The policies of each token. A token that is not registered has no policy, because a token policy set needs a
  /// registered token, and a deregistration needs a token without a policy.
  mapping(address localToken => TokenState tokenState) private s_tokenStates;

  /// @dev The fee policies, with the chains of their records and their token counts.
  PolicyConfig.Policies private s_feePolicies;
  /// @dev The values of each fee policy record. Chain selector zero holds the policy default record.
  mapping(bytes32 policyName => mapping(uint64 remoteChainSelector => FeeConfig feeConfig)) private s_policyFeeConfigs;
  /// @dev The chains of the inline fee records of each token. An inline record exists exactly when its chain is in
  /// the set of its token.
  mapping(address localToken => EnumerableSet.UintSet remoteChainSelectors) private s_inlineFeeChains;
  /// @dev The values of each inline fee record.
  mapping(address localToken => mapping(uint64 remoteChainSelector => FeeConfig feeConfig)) private s_inlineFeeConfigs;

  /// @dev The destination overhead policies, with the chains of their records and their token counts.
  PolicyConfig.Policies private s_destOverheadPolicies;
  /// @dev The values of each destination overhead policy record. Chain selector zero holds the policy default record.
  mapping(bytes32 policyName => mapping(uint64 remoteChainSelector => DestOverheadConfig destOverheadConfig)) private
    s_policyDestOverheadConfigs;
  /// @dev The chains of the inline destination overhead records of each token. An inline record exists exactly when
  /// its chain is in the set of its token.
  mapping(address localToken => EnumerableSet.UintSet remoteChainSelectors) private s_inlineDestOverheadChains;
  /// @dev The values of each inline destination overhead record.
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

  /// @notice Returns the token fee of an outbound transfer.
  /// @dev Reverts if the token has no fee record or no destination overhead record for the chain, in that order. The
  /// connection store checks the requested finality before this call.
  /// @param localToken The local token.
  /// @param remoteChainSelector The destination chain.
  /// @param amount The amount before the token fee.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @param chargeTokenFee False for the V1 path, which charges no token fee.
  /// @return feeAmount The token fee, in local token units. Zero while the fee record is disabled.
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

  /// @notice Reverts if the token has no fee record or no destination overhead record for the chain, in that order.
  /// @dev The lane check of the pool config manager calls it, after a write that opens a lane.
  /// @param localToken The local token.
  /// @param remoteChainSelector The remote chain.
  function checkLaneRecords(
    address localToken,
    uint64 remoteChainSelector
  ) external view {
    _getLaneFeeConfig(localToken, remoteChainSelector);
  }

  /// @dev Returns the fee record of the token for the chain, after it checks that the token has both records.
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

  /// @notice Returns the fee parameters of each lane for the requested finality. The pool config manager serves
  /// IPoolConfigManager.getFeeBatch with it.
  /// @dev The connection store checks the requested finality and gives hasLanes before this call.
  /// @param lanes The token and remote chain lanes to query.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @param hasLanes True for each lane that exists, aligned with lanes by index.
  /// @return fees The fee parameters per lane, aligned with lanes by index. Zeros with isEnabled false for a lane that
  /// does not exist, has no enabled fee record or has no destination overhead record.
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

  /// @notice Returns the fee config of each lane for both finality classes. The pool config manager serves
  /// IPoolConfigManager.getTokenTransferFeeConfigBatch with it.
  /// @dev The connection store gives hasLanes before this call.
  /// @param lanes The token and remote chain lanes to query.
  /// @param hasLanes True for each lane that exists, aligned with lanes by index.
  /// @return feeConfigs The fee config per lane, aligned with lanes by index. A disabled zero config for a lane that
  /// does not exist, has no enabled fee record or has no destination overhead record.
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

  /// @dev Returns true if the token has an enabled fee record and a destination overhead record for the chain, and
  /// returns both records.
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

  /// @notice Returns the level of the fee record and of the destination overhead record that each token would use on
  /// its remote chain. Never reverts.
  /// @dev Each level follows the resolution of a transfer, also for a token without a lane, so tooling can check the
  /// records before a write opens the lane. Chain selector zero gives None.
  /// @param lanes The token and remote chain lanes to query.
  /// @return feeLevels The level of the fee record per lane, aligned with lanes by index.
  /// @return destOverheadLevels The level of the destination overhead record per lane, aligned with lanes by index.
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

  /// @dev Returns the fee record that serves the token on the chain, and its level. An inline record overrides the
  /// policy of the token, so the inline record is checked first. For level None, it returns the inline slot of the
  /// token for the chain, which holds no record. The caller checks the level before it uses the record.
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

  /// @dev Returns the destination overhead record that serves the token on the chain, and its level. It follows the
  /// same order as _resolveFeeConfig, and also returns the inline slot for level None.
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
  // │                         Stored views                         │
  // ================================================================

  /// @notice Returns the registration and the policies of a token in this store. Never reverts.
  /// @param localToken The local token.
  /// @return isRegistered True if the token is registered in this store.
  /// @return feePolicyName The fee policy of the token. Zero if it has none.
  /// @return destOverheadPolicyName The destination overhead policy of the token. Zero if it has none.
  function getTokenPolicies(
    address localToken
  ) external view returns (bool isRegistered, bytes32 feePolicyName, bytes32 destOverheadPolicyName) {
    TokenState storage tokenState = s_tokenStates[localToken];
    return (s_tokens.contains(localToken), tokenState.feePolicyName, tokenState.destOverheadPolicyName);
  }

  /// @notice Returns the policies of each token. The pool config manager adds them to
  /// IPoolConfigManager.getTokenConfigBatch. Never reverts.
  /// @param localTokens The local tokens.
  /// @return feePolicyNames The fee policy per token, aligned with localTokens by index. Zero if it has none.
  /// @return destOverheadPolicyNames The destination overhead policy per token, aligned with localTokens by index.
  /// Zero if it has none.
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

  /// @notice Returns the name of every fee policy and every destination overhead policy. Never reverts.
  /// @return feePolicyNames The names of the fee policies.
  /// @return destOverheadPolicyNames The names of the destination overhead policies.
  function getPolicyNames()
    external
    view
    returns (bytes32[] memory feePolicyNames, bytes32[] memory destOverheadPolicyNames)
  {
    return (s_feePolicies.policyNames.values(), s_destOverheadPolicies.policyNames.values());
  }

  /// @notice Returns the records of a fee policy, in the shape of applyFeePolicyConfigUpdates. Never reverts.
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
      FeeConfig storage feeConfig = s_policyFeeConfigs[policyName][remoteChainSelector];
      policyConfigs[i] = FeePolicyConfigArgs({
        policyName: policyName,
        remoteChainSelector: remoteChainSelector,
        finalityFeeUSDCents: feeConfig.finalityFeeUSDCents,
        fastFinalityFeeUSDCents: feeConfig.fastFinalityFeeUSDCents,
        finalityTransferFeeBps: feeConfig.finalityTransferFeeBps,
        fastFinalityTransferFeeBps: feeConfig.fastFinalityTransferFeeBps,
        isEnabled: feeConfig.isEnabled
      });
    }
    return (policy.tokenCount, policyConfigs);
  }

  /// @notice Returns the inline fee records of a token, in the shape of applyInlineFeeConfigUpdates. Never reverts.
  /// @param localToken The local token.
  /// @return inlineConfigs The inline fee records of the token.
  function getInlineFeeConfigs(
    address localToken
  ) external view returns (InlineFeeConfigArgs[] memory inlineConfigs) {
    uint256[] memory remoteChainSelectors = s_inlineFeeChains[localToken].values();
    inlineConfigs = new InlineFeeConfigArgs[](remoteChainSelectors.length);
    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChainSelectors[i]);
      FeeConfig storage feeConfig = s_inlineFeeConfigs[localToken][remoteChainSelector];
      inlineConfigs[i] = InlineFeeConfigArgs({
        localToken: localToken,
        remoteChainSelector: remoteChainSelector,
        finalityFeeUSDCents: feeConfig.finalityFeeUSDCents,
        fastFinalityFeeUSDCents: feeConfig.fastFinalityFeeUSDCents,
        finalityTransferFeeBps: feeConfig.finalityTransferFeeBps,
        fastFinalityTransferFeeBps: feeConfig.fastFinalityTransferFeeBps,
        isEnabled: feeConfig.isEnabled
      });
    }
    return inlineConfigs;
  }

  /// @notice Returns the records of a destination overhead policy, in the shape of
  /// applyDestOverheadPolicyConfigUpdates. Never reverts.
  /// @param policyName The destination overhead policy.
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
      DestOverheadConfig storage destOverheadConfig = s_policyDestOverheadConfigs[policyName][remoteChainSelector];
      policyConfigs[i] = DestOverheadPolicyConfigArgs({
        policyName: policyName,
        remoteChainSelector: remoteChainSelector,
        destGasOverhead: destOverheadConfig.destGasOverhead,
        destBytesOverhead: destOverheadConfig.destBytesOverhead
      });
    }
    return (policy.tokenCount, policyConfigs);
  }

  /// @notice Returns the inline destination overhead records of a token, in the shape of
  /// applyInlineDestOverheadConfigUpdates. Never reverts.
  /// @param localToken The local token.
  /// @return inlineConfigs The inline destination overhead records of the token.
  function getInlineDestOverheadConfigs(
    address localToken
  ) external view returns (InlineDestOverheadConfigArgs[] memory inlineConfigs) {
    uint256[] memory remoteChainSelectors = s_inlineDestOverheadChains[localToken].values();
    inlineConfigs = new InlineDestOverheadConfigArgs[](remoteChainSelectors.length);
    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChainSelectors[i]);
      DestOverheadConfig storage destOverheadConfig = s_inlineDestOverheadConfigs[localToken][remoteChainSelector];
      inlineConfigs[i] = InlineDestOverheadConfigArgs({
        localToken: localToken,
        remoteChainSelector: remoteChainSelector,
        destGasOverhead: destOverheadConfig.destGasOverhead,
        destBytesOverhead: destOverheadConfig.destBytesOverhead
      });
    }
    return inlineConfigs;
  }

  // ================================================================
  // │                            Tokens                            │
  // ================================================================

  /// @notice Deregisters tokens, then registers or updates tokens. The pool config manager makes the same call to every
  /// store in one transaction.
  /// @dev Only an authorized pool config manager can call this function. This store keeps no value of the profile, so
  /// it reads only the token of each profile. A set of a new token registers it with no policy. A set of a registered
  /// token changes nothing in this store. A deregistration needs a token without a policy and without an inline record
  /// in this store, so one wrong address cannot drop live config.
  /// @param tokensToRemove The registered tokens to deregister.
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

  /// @notice Sets the fee policy of tokens. A token uses the records of its new policy at once. Its inline records
  /// stay.
  /// @dev Only an authorized pool config manager can call this function. Each item needs a registered token and a
  /// policy that exists.
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

  /// @notice Sets the destination overhead policy of tokens. A token uses the records of its new policy at once. Its
  /// inline records stay.
  /// @dev Only an authorized pool config manager can call this function. Each item needs a registered token and a
  /// policy that exists.
  /// @param tokenPoliciesToSet The destination overhead policy of each token, applied in order. Zero removes the token
  /// from its policy.
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
  // │                           Fee type                           │
  // ================================================================

  /// @notice Removes fee policy records, removes fee policies, then sets fee policy records.
  /// @dev Only an authorized pool config manager can call this function. The first record set under a new name
  /// creates the policy. A record removal keeps the policy, also when it has no records left. A policy removal needs a
  /// policy without tokens and without records.
  /// @param policyConfigsToRemove The policy records to remove. Chain selector zero names the policy default record.
  /// @param policiesToRemove The policies to remove.
  /// @param policyConfigsToSet The policy records to set. Each one replaces every value of its record.
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
      FeeConfig memory feeConfig = FeeConfig({
        finalityFeeUSDCents: args.finalityFeeUSDCents,
        fastFinalityFeeUSDCents: args.fastFinalityFeeUSDCents,
        finalityTransferFeeBps: args.finalityTransferFeeBps,
        fastFinalityTransferFeeBps: args.fastFinalityTransferFeeBps,
        isEnabled: args.isEnabled
      });
      _validateFeeConfig(feeConfig);

      PolicyConfig._setPolicyRecord(s_feePolicies, args.policyName, args.remoteChainSelector);
      s_policyFeeConfigs[args.policyName][args.remoteChainSelector] = feeConfig;
      emit FeePolicyConfigSet(args.policyName, args.remoteChainSelector, feeConfig);
    }
  }

  /// @notice Removes inline fee records, then sets inline fee records.
  /// @dev Only an authorized pool config manager can call this function.
  /// @param inlineConfigsToRemove The inline records to remove.
  /// @param inlineConfigsToSet The inline records to set. Each one replaces every value of its record.
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
      FeeConfig memory feeConfig = FeeConfig({
        finalityFeeUSDCents: args.finalityFeeUSDCents,
        fastFinalityFeeUSDCents: args.fastFinalityFeeUSDCents,
        finalityTransferFeeBps: args.finalityTransferFeeBps,
        fastFinalityTransferFeeBps: args.fastFinalityTransferFeeBps,
        isEnabled: args.isEnabled
      });
      _validateFeeConfig(feeConfig);

      PolicyConfig._setInlineRecord(
        s_tokens, s_inlineFeeChains[args.localToken], args.localToken, args.remoteChainSelector
      );
      s_inlineFeeConfigs[args.localToken][args.remoteChainSelector] = feeConfig;
      emit InlineFeeConfigSet(args.localToken, args.remoteChainSelector, feeConfig);
    }
  }

  /// @dev An enabled config gets the basis point checks of TokenPool.applyTokenTransferFeeConfigUpdates. A disabled
  /// config must be zero in every field, so that no fee value stays hidden in storage.
  function _validateFeeConfig(
    FeeConfig memory feeConfig
  ) private pure {
    if (!feeConfig.isEnabled) {
      if (
        feeConfig.finalityFeeUSDCents != 0 || feeConfig.fastFinalityFeeUSDCents != 0
          || feeConfig.finalityTransferFeeBps != 0 || feeConfig.fastFinalityTransferFeeBps != 0
      ) {
        revert DisabledNonZeroFeeConfig(feeConfig);
      }
      return;
    }
    if (feeConfig.finalityTransferFeeBps >= BPS_DIVIDER) {
      revert InvalidTransferFeeBps(feeConfig.finalityTransferFeeBps);
    }
    if (feeConfig.fastFinalityTransferFeeBps >= BPS_DIVIDER) {
      revert InvalidTransferFeeBps(feeConfig.fastFinalityTransferFeeBps);
    }
  }

  // ================================================================
  // │                   Destination overhead type                  │
  // ================================================================

  /// @notice Removes destination overhead policy records, removes destination overhead policies, then sets destination
  /// overhead policy records.
  /// @dev Only an authorized pool config manager can call this function. The first record set under a new name
  /// creates the policy. A record removal keeps the policy, also when it has no records left. A policy removal needs a
  /// policy without tokens and without records.
  /// @param policyConfigsToRemove The policy records to remove. Chain selector zero names the policy default record.
  /// @param policiesToRemove The policies to remove.
  /// @param policyConfigsToSet The policy records to set. Each one replaces every value of its record.
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
      DestOverheadConfig memory destOverheadConfig =
        DestOverheadConfig({destGasOverhead: args.destGasOverhead, destBytesOverhead: args.destBytesOverhead});
      _validateDestOverheadConfig(args.remoteChainSelector, destOverheadConfig);

      PolicyConfig._setPolicyRecord(s_destOverheadPolicies, args.policyName, args.remoteChainSelector);
      s_policyDestOverheadConfigs[args.policyName][args.remoteChainSelector] = destOverheadConfig;
      emit DestOverheadPolicyConfigSet(args.policyName, args.remoteChainSelector, destOverheadConfig);
    }
  }

  /// @notice Removes inline destination overhead records, then sets inline destination overhead records.
  /// @dev Only an authorized pool config manager can call this function.
  /// @param inlineConfigsToRemove The inline records to remove.
  /// @param inlineConfigsToSet The inline records to set. Each one replaces every value of its record.
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
      DestOverheadConfig memory destOverheadConfig =
        DestOverheadConfig({destGasOverhead: args.destGasOverhead, destBytesOverhead: args.destBytesOverhead});
      _validateDestOverheadConfig(args.remoteChainSelector, destOverheadConfig);

      PolicyConfig._setInlineRecord(
        s_tokens, s_inlineDestOverheadChains[args.localToken], args.localToken, args.remoteChainSelector
      );
      s_inlineDestOverheadConfigs[args.localToken][args.remoteChainSelector] = destOverheadConfig;
      emit InlineDestOverheadConfigSet(args.localToken, args.remoteChainSelector, destOverheadConfig);
    }
  }

  /// @dev The gas overhead must be above zero, the same as in TokenPool.applyTokenTransferFeeConfigUpdates. The bytes
  /// overhead must be at least MIN_DEST_BYTES_OVERHEAD.
  function _validateDestOverheadConfig(
    uint64 remoteChainSelector,
    DestOverheadConfig memory destOverheadConfig
  ) private pure {
    if (destOverheadConfig.destGasOverhead == 0 || destOverheadConfig.destBytesOverhead < MIN_DEST_BYTES_OVERHEAD) {
      revert InvalidDestOverheads(
        remoteChainSelector, destOverheadConfig.destGasOverhead, destOverheadConfig.destBytesOverhead
      );
    }
  }
}
