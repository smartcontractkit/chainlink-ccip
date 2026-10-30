// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract FeePolicyStore_applyInlineFeeConfigUpdates is FeePolicyStoreSetup {
  uint16 internal constant INLINE_TRANSFER_FEE_BPS = 50;

  function test_applyInlineFeeConfigUpdates() public {
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, DEST_CHAIN_SELECTOR, _getFeeConfig()));
    FeePolicyStore.FeeConfig memory inlineFeeConfig = _getFeeConfig();
    inlineFeeConfig.finalityTransferFeeBps = INLINE_TRANSFER_FEE_BPS;
    FeePolicyStore.InlineFeeConfigArgs[] memory inlineConfigs = new FeePolicyStore.InlineFeeConfigArgs[](2);
    inlineConfigs[0] = _getInlineFeeConfigArgs(s_token, DEST_CHAIN_SELECTOR, inlineFeeConfig)[0];
    inlineConfigs[1] = _getInlineFeeConfigArgs(s_token, SOURCE_CHAIN_SELECTOR, _getFeeConfig())[0];

    vm.expectEmit();
    emit FeePolicyStore.InlineFeeConfigSet(s_token, DEST_CHAIN_SELECTOR, inlineFeeConfig);
    vm.expectEmit();
    emit FeePolicyStore.InlineFeeConfigSet(s_token, SOURCE_CHAIN_SELECTOR, _getFeeConfig());
    s_feeStore.applyInlineFeeConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);

    assertEq(abi.encode(inlineConfigs), abi.encode(s_feeStore.getInlineFeeConfigs(s_token)));
    // The inline record overrides the policy chain record of the token.
    assertEq(
      (1e18 * uint256(INLINE_TRANSFER_FEE_BPS)) / BPS_DIVIDER,
      s_feeStore.resolveTokenFee(s_token, DEST_CHAIN_SELECTOR, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, true)
    );
  }

  function test_applyInlineFeeConfigUpdates_RemoveInlineRecord() public {
    _setInlineFeeConfig(_getInlineFeeConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getFeeConfig()));

    vm.expectEmit();
    emit FeePolicyStore.InlineFeeConfigRemoved(s_token, DEST_CHAIN_SELECTOR);
    s_feeStore.applyInlineFeeConfigUpdates(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR), new FeePolicyStore.InlineFeeConfigArgs[](0)
    );

    assertEq(0, s_feeStore.getInlineFeeConfigs(s_token).length);
    // The token falls back to its policy at once.
    (PolicyConfig.RecordLevel[] memory feeLevels,) =
      s_feeStore.getLaneRecordLevelBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));
    assertEq(uint8(PolicyConfig.RecordLevel.PolicyDefault), uint8(feeLevels[0]));
  }

  function test_applyInlineFeeConfigUpdates_RemovalsRunFirst() public {
    _setInlineFeeConfig(_getInlineFeeConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getFeeConfig()));
    FeePolicyStore.FeeConfig memory newFeeConfig = _getFeeConfig();
    newFeeConfig.finalityFeeUSDCents = 500;
    FeePolicyStore.InlineFeeConfigArgs[] memory inlineConfigs =
      _getInlineFeeConfigArgs(s_token, DEST_CHAIN_SELECTOR, newFeeConfig);

    s_feeStore.applyInlineFeeConfigUpdates(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR), inlineConfigs);

    assertEq(abi.encode(inlineConfigs), abi.encode(s_feeStore.getInlineFeeConfigs(s_token)));
  }

  // Reverts

  function test_applyInlineFeeConfigUpdates_RevertWhen_InvalidInlineFeeConfig() public {
    // A disabled config with fees, and a bps fee at BPS_DIVIDER.
    FeePolicyStore.FeeConfig[] memory invalidFeeConfigs = new FeePolicyStore.FeeConfig[](2);
    invalidFeeConfigs[0] = _getFeeConfig();
    invalidFeeConfigs[0].isEnabled = false;
    invalidFeeConfigs[1] = _getFeeConfig();
    invalidFeeConfigs[1].finalityTransferFeeBps = uint16(BPS_DIVIDER);

    for (uint256 i = 0; i < invalidFeeConfigs.length; ++i) {
      FeePolicyStore.InlineFeeConfigArgs[] memory inlineConfigs =
        _getInlineFeeConfigArgs(s_token, DEST_CHAIN_SELECTOR, invalidFeeConfigs[i]);

      vm.expectRevert(
        abi.encodeWithSelector(
          FeePolicyStore.InvalidInlineFeeConfig.selector, s_token, DEST_CHAIN_SELECTOR, invalidFeeConfigs[i]
        )
      );
      s_feeStore.applyInlineFeeConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
    }
  }

  function test_applyInlineFeeConfigUpdates_RevertWhen_ZeroChainSelectorNotAllowed() public {
    FeePolicyStore.InlineFeeConfigArgs[] memory inlineConfigs =
      _getInlineFeeConfigArgs(s_token, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _getFeeConfig());

    vm.expectRevert(PolicyConfig.ZeroChainSelectorNotAllowed.selector);
    s_feeStore.applyInlineFeeConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function test_applyInlineFeeConfigUpdates_RevertWhen_LocalTokenNotRegistered() public {
    address unknownToken = makeAddr("unknownToken");
    FeePolicyStore.InlineFeeConfigArgs[] memory inlineConfigs =
      _getInlineFeeConfigArgs(unknownToken, DEST_CHAIN_SELECTOR, _getFeeConfig());

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, unknownToken));
    s_feeStore.applyInlineFeeConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function test_applyInlineFeeConfigUpdates_RevertWhen_InlineRecordNotFound() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.InlineRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR));
    s_feeStore.applyInlineFeeConfigUpdates(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR), new FeePolicyStore.InlineFeeConfigArgs[](0)
    );
  }

  function test_applyInlineFeeConfigUpdates_RevertWhen_UnauthorizedCaller() public {
    FeePolicyStore.InlineFeeConfigArgs[] memory inlineConfigs =
      _getInlineFeeConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getFeeConfig());
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_feeStore.applyInlineFeeConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }
}
