// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract FeePolicyStore_applyInlineDestOverheadConfigUpdates is FeePolicyStoreSetup {
  uint32 internal constant INLINE_DEST_GAS_OVERHEAD = 150_000;

  function test_applyInlineDestOverheadConfigUpdates() public {
    _setDestOverheadPolicyConfig(
      _getDestOverheadPolicyConfigArgs(DEST_OVERHEAD_POLICY, DEST_CHAIN_SELECTOR, _getDestOverheadConfig())
    );
    FeePolicyStore.DestOverheadConfig memory inlineDestOverheadConfig = _getDestOverheadConfig();
    inlineDestOverheadConfig.destGasOverhead = INLINE_DEST_GAS_OVERHEAD;
    FeePolicyStore.InlineDestOverheadConfigArgs[] memory inlineConfigs =
      new FeePolicyStore.InlineDestOverheadConfigArgs[](2);
    inlineConfigs[0] = _getInlineDestOverheadConfigArgs(s_token, DEST_CHAIN_SELECTOR, inlineDestOverheadConfig)[0];
    inlineConfigs[1] = _getInlineDestOverheadConfigArgs(s_token, SOURCE_CHAIN_SELECTOR, _getDestOverheadConfig())[0];

    vm.expectEmit();
    emit FeePolicyStore.InlineDestOverheadConfigSet(s_token, DEST_CHAIN_SELECTOR, inlineDestOverheadConfig);
    vm.expectEmit();
    emit FeePolicyStore.InlineDestOverheadConfigSet(s_token, SOURCE_CHAIN_SELECTOR, _getDestOverheadConfig());
    s_feeStore.applyInlineDestOverheadConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);

    assertEq(abi.encode(inlineConfigs), abi.encode(s_feeStore.getInlineDestOverheadConfigs(s_token)));
    // The inline record overrides the policy chain record of the token.
    IPoolConfigManager.LaneFee[] memory fees = s_feeStore.getFeeBatch(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR), FinalityCodec.WAIT_FOR_FINALITY_FLAG, _getHasLanes(true)
    );
    assertEq(INLINE_DEST_GAS_OVERHEAD, fees[0].destGasOverhead);
  }

  function test_applyInlineDestOverheadConfigUpdates_RemoveInlineRecord() public {
    _setInlineDestOverheadConfig(
      _getInlineDestOverheadConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getDestOverheadConfig())
    );

    vm.expectEmit();
    emit FeePolicyStore.InlineDestOverheadConfigRemoved(s_token, DEST_CHAIN_SELECTOR);
    s_feeStore.applyInlineDestOverheadConfigUpdates(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR), new FeePolicyStore.InlineDestOverheadConfigArgs[](0)
    );

    assertEq(0, s_feeStore.getInlineDestOverheadConfigs(s_token).length);
    // The token falls back to its policy at once.
    (, PolicyConfig.RecordLevel[] memory destOverheadLevels) =
      s_feeStore.getLaneRecordLevelBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));
    assertEq(uint8(PolicyConfig.RecordLevel.PolicyDefault), uint8(destOverheadLevels[0]));
  }

  function test_applyInlineDestOverheadConfigUpdates_RemovalsRunFirst() public {
    _setInlineDestOverheadConfig(
      _getInlineDestOverheadConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getDestOverheadConfig())
    );
    FeePolicyStore.DestOverheadConfig memory newDestOverheadConfig = _getDestOverheadConfig();
    newDestOverheadConfig.destGasOverhead = INLINE_DEST_GAS_OVERHEAD;
    FeePolicyStore.InlineDestOverheadConfigArgs[] memory inlineConfigs =
      _getInlineDestOverheadConfigArgs(s_token, DEST_CHAIN_SELECTOR, newDestOverheadConfig);

    s_feeStore.applyInlineDestOverheadConfigUpdates(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR), inlineConfigs);

    assertEq(abi.encode(inlineConfigs), abi.encode(s_feeStore.getInlineDestOverheadConfigs(s_token)));
  }

  // Reverts

  function test_applyInlineDestOverheadConfigUpdates_RevertWhen_InvalidDestOverheads() public {
    FeePolicyStore.DestOverheadConfig memory destOverheadConfig = _getDestOverheadConfig();
    destOverheadConfig.destGasOverhead = 0;
    FeePolicyStore.InlineDestOverheadConfigArgs[] memory inlineConfigs =
      _getInlineDestOverheadConfigArgs(s_token, DEST_CHAIN_SELECTOR, destOverheadConfig);

    vm.expectRevert(
      abi.encodeWithSelector(
        FeePolicyStore.InvalidDestOverheads.selector, DEST_CHAIN_SELECTOR, 0, MIN_DEST_BYTES_OVERHEAD
      )
    );
    s_feeStore.applyInlineDestOverheadConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);

    destOverheadConfig = _getDestOverheadConfig();
    destOverheadConfig.destBytesOverhead = MIN_DEST_BYTES_OVERHEAD - 1;
    inlineConfigs = _getInlineDestOverheadConfigArgs(s_token, DEST_CHAIN_SELECTOR, destOverheadConfig);

    vm.expectRevert(
      abi.encodeWithSelector(
        FeePolicyStore.InvalidDestOverheads.selector,
        DEST_CHAIN_SELECTOR,
        DEFAULT_TOKEN_DEST_GAS_OVERHEAD,
        MIN_DEST_BYTES_OVERHEAD - 1
      )
    );
    s_feeStore.applyInlineDestOverheadConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function test_applyInlineDestOverheadConfigUpdates_RevertWhen_ZeroChainSelectorNotAllowed() public {
    FeePolicyStore.InlineDestOverheadConfigArgs[] memory inlineConfigs =
      _getInlineDestOverheadConfigArgs(s_token, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _getDestOverheadConfig());

    vm.expectRevert(PolicyConfig.ZeroChainSelectorNotAllowed.selector);
    s_feeStore.applyInlineDestOverheadConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function test_applyInlineDestOverheadConfigUpdates_RevertWhen_LocalTokenNotRegistered() public {
    address unknownToken = makeAddr("unknownToken");
    FeePolicyStore.InlineDestOverheadConfigArgs[] memory inlineConfigs =
      _getInlineDestOverheadConfigArgs(unknownToken, DEST_CHAIN_SELECTOR, _getDestOverheadConfig());

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, unknownToken));
    s_feeStore.applyInlineDestOverheadConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function test_applyInlineDestOverheadConfigUpdates_RevertWhen_InlineRecordNotFound() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.InlineRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR));
    s_feeStore.applyInlineDestOverheadConfigUpdates(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR), new FeePolicyStore.InlineDestOverheadConfigArgs[](0)
    );
  }

  function test_applyInlineDestOverheadConfigUpdates_RevertWhen_UnauthorizedCaller() public {
    FeePolicyStore.InlineDestOverheadConfigArgs[] memory inlineConfigs =
      _getInlineDestOverheadConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getDestOverheadConfig());
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_feeStore.applyInlineDestOverheadConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }
}
