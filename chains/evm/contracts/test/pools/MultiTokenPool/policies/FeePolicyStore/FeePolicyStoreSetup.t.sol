// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {BaseTest} from "../../../../BaseTest.t.sol";

/// @notice s_poolConfigManager is the only authorized caller and sends every call. s_token uses FEE_POLICY and
/// DEST_OVERHEAD_POLICY, which both have a policy default record. s_otherToken is registered with no policy.
contract FeePolicyStoreSetup is BaseTest {
  bytes32 internal constant FEE_POLICY = "FEE_POLICY";
  bytes32 internal constant DEST_OVERHEAD_POLICY = "DEST_OVERHEAD_POLICY";
  bytes32 internal constant OTHER_POLICY = "OTHER_POLICY";
  uint256 internal constant BPS_DIVIDER = 10_000;
  uint32 internal constant MIN_DEST_BYTES_OVERHEAD = 192;

  FeePolicyStore internal s_feeStore;
  address internal s_poolConfigManager = makeAddr("poolConfigManager");
  address internal s_token = makeAddr("token");
  address internal s_otherToken = makeAddr("otherToken");

  function setUp() public virtual override {
    super.setUp();
    s_feeStore = new FeePolicyStore(_arrayOf(s_poolConfigManager));
    vm.startPrank(s_poolConfigManager);

    _registerToken(s_token);
    _registerToken(s_otherToken);
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _getFeeConfig()));
    _setDestOverheadPolicyConfig(
      _getDestOverheadPolicyConfigArgs(
        DEST_OVERHEAD_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _getDestOverheadConfig()
      )
    );
    _setTokenFeePolicy(s_token, FEE_POLICY);
    _setTokenDestOverheadPolicy(s_token, DEST_OVERHEAD_POLICY);
  }

  // ================================================================
  // │                          Arguments                           │
  // ================================================================

  /// @dev Each field has a distinct value, so tests can tell the fields apart.
  function _getFeeConfig() internal pure returns (FeePolicyStore.FeeConfig memory) {
    return FeePolicyStore.FeeConfig({
      finalityFeeUSDCents: 100,
      fastFinalityFeeUSDCents: 200,
      finalityTransferFeeBps: 10,
      fastFinalityTransferFeeBps: 30,
      isEnabled: true
    });
  }

  /// @dev Each field differs from _getFeeConfig.
  function _getOtherFeeConfig() internal pure returns (FeePolicyStore.FeeConfig memory) {
    return FeePolicyStore.FeeConfig({
      finalityFeeUSDCents: 300,
      fastFinalityFeeUSDCents: 400,
      finalityTransferFeeBps: 50,
      fastFinalityTransferFeeBps: 70,
      isEnabled: true
    });
  }

  /// @dev The smallest valid bytes overhead.
  function _getDestOverheadConfig() internal pure returns (FeePolicyStore.DestOverheadConfig memory) {
    return FeePolicyStore.DestOverheadConfig({
      destGasOverhead: DEFAULT_TOKEN_DEST_GAS_OVERHEAD, destBytesOverhead: MIN_DEST_BYTES_OVERHEAD
    });
  }

  /// @dev Each field differs from _getDestOverheadConfig.
  function _getOtherDestOverheadConfig() internal pure returns (FeePolicyStore.DestOverheadConfig memory) {
    return FeePolicyStore.DestOverheadConfig({destGasOverhead: 150_000, destBytesOverhead: 256});
  }

  function _getFeePolicyConfigArgs(
    bytes32 policyName,
    uint64 remoteChainSelector,
    FeePolicyStore.FeeConfig memory feeConfig
  ) internal pure returns (FeePolicyStore.FeePolicyConfigArgs[] memory policyConfigs) {
    policyConfigs = new FeePolicyStore.FeePolicyConfigArgs[](1);
    policyConfigs[0] = FeePolicyStore.FeePolicyConfigArgs({
      policyName: policyName, remoteChainSelector: remoteChainSelector, feeConfig: feeConfig
    });
    return policyConfigs;
  }

  function _getInlineFeeConfigArgs(
    address localToken,
    uint64 remoteChainSelector,
    FeePolicyStore.FeeConfig memory feeConfig
  ) internal pure returns (FeePolicyStore.InlineFeeConfigArgs[] memory inlineConfigs) {
    inlineConfigs = new FeePolicyStore.InlineFeeConfigArgs[](1);
    inlineConfigs[0] = FeePolicyStore.InlineFeeConfigArgs({
      localToken: localToken, remoteChainSelector: remoteChainSelector, feeConfig: feeConfig
    });
    return inlineConfigs;
  }

  function _getDestOverheadPolicyConfigArgs(
    bytes32 policyName,
    uint64 remoteChainSelector,
    FeePolicyStore.DestOverheadConfig memory destOverheadConfig
  ) internal pure returns (FeePolicyStore.DestOverheadPolicyConfigArgs[] memory policyConfigs) {
    policyConfigs = new FeePolicyStore.DestOverheadPolicyConfigArgs[](1);
    policyConfigs[0] = FeePolicyStore.DestOverheadPolicyConfigArgs({
      policyName: policyName, remoteChainSelector: remoteChainSelector, destOverheadConfig: destOverheadConfig
    });
    return policyConfigs;
  }

  function _getInlineDestOverheadConfigArgs(
    address localToken,
    uint64 remoteChainSelector,
    FeePolicyStore.DestOverheadConfig memory destOverheadConfig
  ) internal pure returns (FeePolicyStore.InlineDestOverheadConfigArgs[] memory inlineConfigs) {
    inlineConfigs = new FeePolicyStore.InlineDestOverheadConfigArgs[](1);
    inlineConfigs[0] = FeePolicyStore.InlineDestOverheadConfigArgs({
      localToken: localToken, remoteChainSelector: remoteChainSelector, destOverheadConfig: destOverheadConfig
    });
    return inlineConfigs;
  }

  /// @dev The store only reads the token of a profile.
  function _getTokenProfileArgs(
    address localToken
  ) internal pure returns (PolicyConfig.TokenProfileArgs[] memory profiles) {
    profiles = new PolicyConfig.TokenProfileArgs[](1);
    profiles[0].localToken = localToken;
    return profiles;
  }

  function _getTokenPolicyArgs(
    address localToken,
    bytes32 policyName
  ) internal pure returns (PolicyConfig.TokenPolicyArgs[] memory tokenPolicies) {
    tokenPolicies = new PolicyConfig.TokenPolicyArgs[](1);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: localToken, policyName: policyName});
    return tokenPolicies;
  }

  function _getLaneKeys(
    address localToken,
    uint64 remoteChainSelector
  ) internal pure returns (IPoolConfigManager.LaneKey[] memory lanes) {
    lanes = new IPoolConfigManager.LaneKey[](1);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: localToken, remoteChainSelector: remoteChainSelector});
    return lanes;
  }

  // ================================================================
  // │                         Config writes                        │
  // ================================================================

  function _registerToken(
    address localToken
  ) internal {
    s_feeStore.applyTokenProfileUpdates(new address[](0), _getTokenProfileArgs(localToken));
  }

  function _setTokenFeePolicy(
    address localToken,
    bytes32 policyName
  ) internal {
    s_feeStore.applyTokenFeePolicyUpdates(_getTokenPolicyArgs(localToken, policyName));
  }

  function _setTokenDestOverheadPolicy(
    address localToken,
    bytes32 policyName
  ) internal {
    s_feeStore.applyTokenDestOverheadPolicyUpdates(_getTokenPolicyArgs(localToken, policyName));
  }

  function _setFeePolicyConfig(
    FeePolicyStore.FeePolicyConfigArgs[] memory policyConfigs
  ) internal {
    s_feeStore.applyFeePolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs);
  }

  function _setInlineFeeConfig(
    FeePolicyStore.InlineFeeConfigArgs[] memory inlineConfigs
  ) internal {
    s_feeStore.applyInlineFeeConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function _setDestOverheadPolicyConfig(
    FeePolicyStore.DestOverheadPolicyConfigArgs[] memory policyConfigs
  ) internal {
    s_feeStore.applyDestOverheadPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  function _setInlineDestOverheadConfig(
    FeePolicyStore.InlineDestOverheadConfigArgs[] memory inlineConfigs
  ) internal {
    s_feeStore.applyInlineDestOverheadConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }
}
