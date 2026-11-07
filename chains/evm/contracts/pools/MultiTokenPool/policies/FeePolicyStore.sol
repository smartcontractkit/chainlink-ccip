// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../libraries/PolicyConfig.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice Stores the FeePolicy and the DestinationOverheadPolicy for a PoolConfigManager.
/// @dev Only authorized pool config managers write to this store.
/// Fee records define the USD and token fees charged for outbound transfers. Zero fees are valid. Destination
/// overhead records define the gas and pool data overheads charged for execution on the destination chain.
contract FeePolicyStore is ITypeAndVersion, AuthorizedCallers {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;
  using EnumerableSet for EnumerableSet.UintSet;

  error ZeroAddressInvalid();
  error InvalidInlineFeeConfig(address localToken, uint64 remoteChainSelector, FeeConfig feeConfig);
  error InvalidFeePolicyConfig(bytes32 policyName, uint64 remoteChainSelector, FeeConfig feeConfig);
  error InvalidInlineDestOverheadConfig(
    address localToken, uint64 remoteChainSelector, DestOverheadConfig destOverheadConfig
  );
  error InvalidDestOverheadPolicyConfig(
    bytes32 policyName, uint64 remoteChainSelector, DestOverheadConfig destOverheadConfig
  );
  error TokenHasConfig(address localToken);
  error FeeRecordNotFound(address localToken, uint64 remoteChainSelector);
  error DestOverheadRecordNotFound(address localToken, uint64 remoteChainSelector);
  error PoolDataTooLarge(
    address localToken, uint64 remoteChainSelector, uint256 destPoolDataLength, uint32 destBytesOverhead
  );

  event InlineFeeConfigSet(address indexed localToken, uint64 indexed remoteChainSelector, FeeConfig feeConfig);
  event InlineFeeConfigRemoved(address indexed localToken, uint64 indexed remoteChainSelector);
  event FeePolicyConfigSet(bytes32 indexed policyName, uint64 indexed remoteChainSelector, FeeConfig feeConfig);
  event FeePolicyConfigRemoved(bytes32 indexed policyName, uint64 indexed remoteChainSelector);
  event FeePolicyRemoved(bytes32 indexed policyName);
  /// @dev Setting a profile for a new token also registers it.
  event TokenProfileSet(address indexed localToken);
  event TokenRemoved(address indexed localToken);
  event TokenFeePolicySet(address indexed localToken, bytes32 indexed oldPolicyName, bytes32 indexed newPolicyName);
  event TokenDestOverheadPolicySet(
    address indexed localToken, bytes32 indexed oldPolicyName, bytes32 indexed newPolicyName
  );
  event InlineDestOverheadConfigSet(
    address indexed localToken, uint64 indexed remoteChainSelector, DestOverheadConfig destOverheadConfig
  );
  event InlineDestOverheadConfigRemoved(address indexed localToken, uint64 indexed remoteChainSelector);
  event DestOverheadPolicyConfigSet(
    bytes32 indexed policyName, uint64 indexed remoteChainSelector, DestOverheadConfig destOverheadConfig
  );
  event DestOverheadPolicyConfigRemoved(bytes32 indexed policyName, uint64 indexed remoteChainSelector);
  event DestOverheadPolicyRemoved(bytes32 indexed policyName);

  /// @notice The values of a fee record.
  struct FeeConfig {
    uint32 finalityFeeUSDCents; // ───────╮ Fee for a transfer with default finality, in multiples of 0.01 USD.
    uint32 fastFinalityFeeUSDCents; //    │ Fee for an FTF transfer, in multiples of 0.01 USD.
    uint16 finalityTransferFeeBps; //     │ Token fee for a transfer with default finality, in basis points.
    uint16 fastFinalityTransferFeeBps; // ╯ Token fee for an FTF transfer, in basis points.
  }

  /// @notice The values of a dest overhead record.
  struct DestOverheadConfig {
    uint32 destGasOverhead; // ──╮ Gas for releaseOrMint on the destination chain. Not zero.
    uint32 destBytesOverhead; // ╯ Bytes charged for data availability and maximum pool data length. At least 32.
  }

  /// @notice FeeConfig with its policy record key.
  struct FeePolicyConfigArgs {
    bytes32 policyName; // The fee policy. Not zero.
    uint64 remoteChainSelector; // The remote chain. Zero names the policy default record.
    FeeConfig feeConfig; // The values of the record.
  }

  /// @notice FeeConfig with its token and chain.
  struct InlineFeeConfigArgs {
    address localToken; // ─────────╮ A registered token.
    uint64 remoteChainSelector; // ─╯ The remote chain. Not zero.
    FeeConfig feeConfig; //           The values of the record.
  }

  /// @notice DestOverheadConfig with its policy record key.
  struct DestOverheadPolicyConfigArgs {
    bytes32 policyName; // The dest overhead policy. Not zero.
    uint64 remoteChainSelector; // The remote chain. Zero names the policy default record.
    DestOverheadConfig destOverheadConfig; // The values of the record.
  }

  /// @notice DestOverheadConfig with its token and chain.
  struct InlineDestOverheadConfigArgs {
    address localToken; // ─────────╮ A registered token.
    uint64 remoteChainSelector; // ─╯ The remote chain. Not zero.
    DestOverheadConfig destOverheadConfig; // The values of the record.
  }

  /// @notice The policies of a token. Zero means no policy.
  struct TokenState {
    bytes32 feePolicyName; // The fee policy of the token.
    bytes32 destOverheadPolicyName; // The dest overhead policy of the token.
  }

  /// @dev The division factor for bps. Each bps fee must be below it.
  uint256 private constant BPS_DIVIDER = 10_000;

  /// @dev Each store keeps its own list of registered tokens, as the stores never call each other.
  EnumerableSet.AddressSet private s_tokens;
  /// @dev Only registered tokens have policies.
  mapping(address localToken => TokenState tokenState) private s_tokenStates;

  // ================================================================
  // │                       Inline storage                         │
  // ================================================================

  /// @dev A token has an inline fee record for a chain if the chain is in its set.
  mapping(address localToken => EnumerableSet.UintSet remoteChainSelectors) private s_inlineFeeChains;
  mapping(address localToken => mapping(uint64 remoteChainSelector => FeeConfig feeConfig)) private s_inlineFeeConfigs;

  /// @dev A token has an inline dest overhead record for a chain if the chain is in its set.
  mapping(address localToken => EnumerableSet.UintSet remoteChainSelectors) private s_inlineDestOverheadChains;
  mapping(address localToken => mapping(uint64 remoteChainSelector => DestOverheadConfig destOverheadConfig)) private
    s_inlineDestOverheadConfigs;

  // ================================================================
  // │                       Policy storage                         │
  // ================================================================

  PolicyConfig.Policies private s_feePolicies;
  /// @dev Fee policy record details. Chain selector zero holds the policy default.
  mapping(bytes32 policyName => mapping(uint64 remoteChainSelector => FeeConfig feeConfig)) private s_policyFeeConfigs;

  PolicyConfig.Policies private s_destOverheadPolicies;
  /// @dev Dest overhead policy record details. Chain selector zero holds the policy default.
  mapping(bytes32 policyName => mapping(uint64 remoteChainSelector => DestOverheadConfig destOverheadConfig)) private
    s_policyDestOverheadConfigs;

  /// @param authorizedCallers The pool config managers that can write to this store.
  constructor(
    address[] memory authorizedCallers
  ) AuthorizedCallers(authorizedCallers) {}

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "FeePolicyStore 2.1.0-dev";
  }

  // ================================================================
  // │                         Transfer fee                         │
  // ================================================================

  /// @notice Validates the outbound fee configuration and pool data length. Returns the token fee.
  /// @dev Reverts if the token has no fee or dest overhead record for the chain. The pool data must fit within
  /// destBytesOverhead, including on the V1 path where no token fee is charged. Any requested finality other than
  /// WAIT_FOR_FINALITY_FLAG is charged the FTF token fee.
  /// @param localToken The local token.
  /// @param remoteChainSelector The destination chain.
  /// @param amount The transfer amount before deducting the token fee.
  /// @param requestedFinalityConfig The requested finality.
  /// @param chargeTokenFee False for the V1 path, which charges no token fee.
  /// @param destPoolDataLength The length in bytes of the pool data returned by lockOrBurn.
  /// @return feeAmount The token fee in local token units, rounded down. Zero if chargeTokenFee is false.
  function validateLockOrBurn(
    address localToken,
    uint64 remoteChainSelector,
    uint256 amount,
    bytes4 requestedFinalityConfig,
    bool chargeTokenFee,
    uint256 destPoolDataLength
  ) external view returns (uint256 feeAmount) {
    (FeeConfig storage feeConfig, DestOverheadConfig storage destOverheadConfig) =
      _getLaneRecords(localToken, remoteChainSelector);

    if (destPoolDataLength > destOverheadConfig.destBytesOverhead) {
      revert PoolDataTooLarge(localToken, remoteChainSelector, destPoolDataLength, destOverheadConfig.destBytesOverhead);
    }
    if (!chargeTokenFee) {
      return 0;
    }

    if (requestedFinalityConfig != FinalityCodec.WAIT_FOR_FINALITY_FLAG) {
      return (amount * feeConfig.fastFinalityTransferFeeBps) / BPS_DIVIDER;
    }
    return (amount * feeConfig.finalityTransferFeeBps) / BPS_DIVIDER;
  }

  /// @dev Returns the fee and destination overhead records. Reverts if either is missing.
  function _getLaneRecords(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (FeeConfig storage feeConfig, DestOverheadConfig storage destOverheadConfig) {
    PolicyConfig.RecordLevel level;

    (level, feeConfig) = _resolveFeeConfig(localToken, remoteChainSelector);
    if (level == PolicyConfig.RecordLevel.None) revert FeeRecordNotFound(localToken, remoteChainSelector);

    (level, destOverheadConfig) = _resolveDestOverheadConfig(localToken, remoteChainSelector);
    if (level == PolicyConfig.RecordLevel.None) revert DestOverheadRecordNotFound(localToken, remoteChainSelector);

    return (feeConfig, destOverheadConfig);
  }

  /// @dev Returns the fee record of the token for the chain, and its level. For level None, the record is the empty
  /// inline slot, so callers must check the level first.
  function _resolveFeeConfig(
    address localToken,
    uint64 remoteChainSelector
  ) private view returns (PolicyConfig.RecordLevel level, FeeConfig storage feeConfig) {
    // Inline records override the policy.
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
  // │                       Resolved views                         │
  // ================================================================

  /// @notice Returns the fee of each lane for the requested finality.
  /// @dev Does not check whether a connection exists. The PoolConfigManager clears results for missing lanes.
  /// @param lanes The lanes to query.
  /// @param requestedFinalityConfig The requested finality.
  /// @return fees Fee parameters in the same order as lanes. Returns zero values with isEnabled false if either the
  /// fee or destination overhead record is missing.
  function getFeeBatch(
    IPoolConfigManager.LaneKey[] calldata lanes,
    bytes4 requestedFinalityConfig
  ) external view returns (IPoolConfigManager.LaneFee[] memory fees) {
    bool isFastFinality = requestedFinalityConfig != FinalityCodec.WAIT_FOR_FINALITY_FLAG;
    fees = new IPoolConfigManager.LaneFee[](lanes.length);

    for (uint256 i = 0; i < lanes.length; ++i) {
      (bool hasRecords, FeeConfig storage feeConfig, DestOverheadConfig storage destOverheadConfig) =
        _resolveLaneRecords(lanes[i]);
      if (!hasRecords) continue;

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
  /// @dev Does not check whether a connection exists. The PoolConfigManager clears results for missing lanes.
  /// @param lanes The lanes to query.
  /// @return feeConfigs Fee configurations in the same order as lanes. Returns zero values with isEnabled false if
  /// either the fee or destination overhead record is missing.
  function getTokenTransferFeeConfigBatch(
    IPoolConfigManager.LaneKey[] calldata lanes
  ) external view returns (IPoolV2.TokenTransferFeeConfig[] memory feeConfigs) {
    feeConfigs = new IPoolV2.TokenTransferFeeConfig[](lanes.length);

    for (uint256 i = 0; i < lanes.length; ++i) {
      (bool hasRecords, FeeConfig storage feeConfig, DestOverheadConfig storage destOverheadConfig) =
        _resolveLaneRecords(lanes[i]);
      if (!hasRecords) continue;

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

  /// @notice Reverts if the token has no fee or dest overhead record for the chain.
  /// @dev The PoolConfigManager calls it after it sets an inline connection record or adds a subset token.
  /// @param localToken The local token.
  /// @param remoteChainSelector The remote chain.
  function checkLaneRecords(
    address localToken,
    uint64 remoteChainSelector
  ) external view {
    _getLaneRecords(localToken, remoteChainSelector);
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

  /// @dev Returns the fee and destination overhead records, and whether both exist.
  function _resolveLaneRecords(
    IPoolConfigManager.LaneKey calldata lane
  ) private view returns (bool hasRecords, FeeConfig storage feeConfig, DestOverheadConfig storage destOverheadConfig) {
    PolicyConfig.RecordLevel feeLevel;
    (feeLevel, feeConfig) = _resolveFeeConfig(lane.localToken, lane.remoteChainSelector);

    PolicyConfig.RecordLevel destOverheadLevel;
    (destOverheadLevel, destOverheadConfig) = _resolveDestOverheadConfig(lane.localToken, lane.remoteChainSelector);

    hasRecords = feeLevel != PolicyConfig.RecordLevel.None && destOverheadLevel != PolicyConfig.RecordLevel.None;
    return (hasRecords, feeConfig, destOverheadConfig);
  }

  // ================================================================
  // │                        Stored views                          │
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
  /// @dev The token must be registered and the policy must exist.
  /// @param tokenPoliciesToSet The fee policy of each token, applied in order. Zero removes the token from its policy.
  function applyTokenFeePolicyUpdates(
    PolicyConfig.TokenPolicyArgs[] calldata tokenPoliciesToSet
  ) external onlyAuthorizedCallers {
    for (uint256 i = 0; i < tokenPoliciesToSet.length; ++i) {
      address localToken = tokenPoliciesToSet[i].localToken;
      bytes32 policyName = tokenPoliciesToSet[i].policyName;

      if (!s_tokens.contains(localToken)) revert PolicyConfig.LocalTokenNotRegistered(localToken);
      TokenState storage tokenState = s_tokenStates[localToken];
      bytes32 oldPolicyName = tokenState.feePolicyName;

      PolicyConfig._moveToken(s_feePolicies, oldPolicyName, policyName);
      tokenState.feePolicyName = policyName;
      emit TokenFeePolicySet(localToken, oldPolicyName, policyName);
    }
  }

  /// @notice Sets the dest overhead policy of tokens. Inline records are kept.
  /// @dev The token must be registered and the policy must exist.
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
      bytes32 oldPolicyName = tokenState.destOverheadPolicyName;

      PolicyConfig._moveToken(s_destOverheadPolicies, oldPolicyName, policyName);
      tokenState.destOverheadPolicyName = policyName;
      emit TokenDestOverheadPolicySet(localToken, oldPolicyName, policyName);
    }
  }

  // ================================================================
  // │                          Fee config                          │
  // ================================================================

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
      address localToken = inlineConfigsToSet[i].localToken;
      uint64 remoteChainSelector = inlineConfigsToSet[i].remoteChainSelector;
      FeeConfig calldata feeConfig = inlineConfigsToSet[i].feeConfig;

      if (!_isValidFeeConfig(feeConfig)) revert InvalidInlineFeeConfig(localToken, remoteChainSelector, feeConfig);

      PolicyConfig._setInlineRecord(s_tokens, s_inlineFeeChains[localToken], localToken, remoteChainSelector);
      s_inlineFeeConfigs[localToken][remoteChainSelector] = feeConfig;
      emit InlineFeeConfigSet(localToken, remoteChainSelector, feeConfig);
    }
  }

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
      bytes32 policyName = policiesToRemove[i];
      PolicyConfig._removePolicy(s_feePolicies, policyName);
      emit FeePolicyRemoved(policyName);
    }

    for (uint256 i = 0; i < policyConfigsToSet.length; ++i) {
      bytes32 policyName = policyConfigsToSet[i].policyName;
      uint64 remoteChainSelector = policyConfigsToSet[i].remoteChainSelector;
      FeeConfig calldata feeConfig = policyConfigsToSet[i].feeConfig;

      if (!_isValidFeeConfig(feeConfig)) revert InvalidFeePolicyConfig(policyName, remoteChainSelector, feeConfig);

      PolicyConfig._setPolicyRecord(s_feePolicies, policyName, remoteChainSelector);
      s_policyFeeConfigs[policyName][remoteChainSelector] = feeConfig;
      emit FeePolicyConfigSet(policyName, remoteChainSelector, feeConfig);
    }
  }

  /// @dev Token fees must be less than 100% of the transfer amount, as in TokenPool.
  function _isValidFeeConfig(
    FeeConfig calldata feeConfig
  ) private pure returns (bool) {
    return feeConfig.finalityTransferFeeBps < BPS_DIVIDER && feeConfig.fastFinalityTransferFeeBps < BPS_DIVIDER;
  }

  // ================================================================
  // │                     Dest overhead config                     │
  // ================================================================

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
      address localToken = inlineConfigsToSet[i].localToken;
      uint64 remoteChainSelector = inlineConfigsToSet[i].remoteChainSelector;
      DestOverheadConfig calldata destOverheadConfig = inlineConfigsToSet[i].destOverheadConfig;

      if (!_isValidDestOverheadConfig(destOverheadConfig)) {
        revert InvalidInlineDestOverheadConfig(localToken, remoteChainSelector, destOverheadConfig);
      }

      PolicyConfig._setInlineRecord(s_tokens, s_inlineDestOverheadChains[localToken], localToken, remoteChainSelector);
      s_inlineDestOverheadConfigs[localToken][remoteChainSelector] = destOverheadConfig;
      emit InlineDestOverheadConfigSet(localToken, remoteChainSelector, destOverheadConfig);
    }
  }

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
      bytes32 policyName = policiesToRemove[i];
      PolicyConfig._removePolicy(s_destOverheadPolicies, policyName);
      emit DestOverheadPolicyRemoved(policyName);
    }

    for (uint256 i = 0; i < policyConfigsToSet.length; ++i) {
      bytes32 policyName = policyConfigsToSet[i].policyName;
      uint64 remoteChainSelector = policyConfigsToSet[i].remoteChainSelector;
      DestOverheadConfig calldata destOverheadConfig = policyConfigsToSet[i].destOverheadConfig;

      if (!_isValidDestOverheadConfig(destOverheadConfig)) {
        revert InvalidDestOverheadPolicyConfig(policyName, remoteChainSelector, destOverheadConfig);
      }

      PolicyConfig._setPolicyRecord(s_destOverheadPolicies, policyName, remoteChainSelector);
      s_policyDestOverheadConfigs[policyName][remoteChainSelector] = destOverheadConfig;
      emit DestOverheadPolicyConfigSet(policyName, remoteChainSelector, destOverheadConfig);
    }
  }

  /// @dev Matches TokenPool and FeeQuoter: the gas overhead must be nonzero and the bytes overhead must fit legacy pool
  /// data. Larger pool data needs a higher destBytesOverhead, validateLockOrBurn checks the actual length.
  function _isValidDestOverheadConfig(
    DestOverheadConfig calldata destOverheadConfig
  ) private pure returns (bool) {
    return destOverheadConfig.destGasOverhead != 0
      && destOverheadConfig.destBytesOverhead >= Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES;
  }
}
