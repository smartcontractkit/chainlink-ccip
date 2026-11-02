// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";

contract FeePolicyStore_getFeeBatch is FeePolicyStoreSetup {
  function test_getFeeBatch() public {
    // s_otherToken gets inline records with other values, so each lane must resolve its own records.
    _setInlineFeeConfig(_getInlineFeeConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, _getOtherFeeConfig()));
    _setInlineDestOverheadConfig(
      _getInlineDestOverheadConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, _getOtherDestOverheadConfig())
    );
    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](3);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: DEST_CHAIN_SELECTOR});
    // A missing lane must not stop the batch or shift later results.
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: SOURCE_CHAIN_SELECTOR});
    lanes[2] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: DEST_CHAIN_SELECTOR});

    IPoolConfigManager.LaneFee[] memory fees = s_feeStore.getFeeBatch(lanes, FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    assertEq(3, fees.length);
    assertEq(abi.encode(_getLaneFee(_getFeeConfig(), _getDestOverheadConfig())), abi.encode(fees[0]));
    IPoolConfigManager.LaneFee memory disabledFee;
    assertEq(abi.encode(disabledFee), abi.encode(fees[1]));
    assertEq(abi.encode(_getLaneFee(_getOtherFeeConfig(), _getOtherDestOverheadConfig())), abi.encode(fees[2]));
  }

  function test_getFeeBatch_PolicyChainRecords() public {
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, DEST_CHAIN_SELECTOR, _getOtherFeeConfig()));
    _setDestOverheadPolicyConfig(
      _getDestOverheadPolicyConfigArgs(DEST_OVERHEAD_POLICY, DEST_CHAIN_SELECTOR, _getOtherDestOverheadConfig())
    );

    IPoolConfigManager.LaneFee[] memory fees =
      s_feeStore.getFeeBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR), FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    // The policy chain records serve the chain instead of the policy default records.
    assertEq(abi.encode(_getLaneFee(_getOtherFeeConfig(), _getOtherDestOverheadConfig())), abi.encode(fees[0]));
  }

  function test_getFeeBatch_FastFinality() public view {
    IPoolConfigManager.LaneFee[] memory fees =
      s_feeStore.getFeeBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR), FinalityCodec._encodeBlockDepth(10));

    assertEq(_getFeeConfig().fastFinalityFeeUSDCents, fees[0].feeUSDCents);
    assertEq(_getFeeConfig().fastFinalityTransferFeeBps, fees[0].tokenFeeBps);
    assertTrue(fees[0].isEnabled);
  }

  function test_getFeeBatch_DisabledFeeRecord() public {
    FeePolicyStore.FeeConfig memory disabledFeeConfig;
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, disabledFeeConfig));

    IPoolConfigManager.LaneFee[] memory fees =
      s_feeStore.getFeeBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR), FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    IPoolConfigManager.LaneFee memory disabledFee;
    assertEq(abi.encode(disabledFee), abi.encode(fees[0]));
  }

  function test_getFeeBatch_DisabledInlineRecordOverridesPolicy() public {
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, DEST_CHAIN_SELECTOR, _getOtherFeeConfig()));
    FeePolicyStore.FeeConfig memory disabledFeeConfig;
    _setInlineFeeConfig(_getInlineFeeConfigArgs(s_token, DEST_CHAIN_SELECTOR, disabledFeeConfig));
    IPoolConfigManager.LaneKey[] memory lanes = _getLaneKeys(s_token, DEST_CHAIN_SELECTOR);

    IPoolConfigManager.LaneFee[] memory fees = s_feeStore.getFeeBatch(lanes, FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    IPoolConfigManager.LaneFee memory disabledFee;
    assertEq(abi.encode(disabledFee), abi.encode(fees[0]));

    fees = s_feeStore.getFeeBatch(lanes, FinalityCodec._encodeBlockDepth(10));

    assertEq(abi.encode(disabledFee), abi.encode(fees[0]));
  }

  function test_getFeeBatch_EnabledZeroFees() public {
    FeePolicyStore.FeeConfig memory feeConfig;
    feeConfig.isEnabled = true;
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, feeConfig));
    IPoolConfigManager.LaneFee memory expectedFee = _getLaneFee(feeConfig, _getDestOverheadConfig());

    IPoolConfigManager.LaneFee[] memory fees =
      s_feeStore.getFeeBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR), FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    assertEq(abi.encode(expectedFee), abi.encode(fees[0]));

    fees = s_feeStore.getFeeBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR), FinalityCodec._encodeBlockDepth(10));

    assertEq(abi.encode(expectedFee), abi.encode(fees[0]));
  }

  function test_getFeeBatch_NoFeeRecord() public {
    _setTokenFeePolicy(s_token, bytes32(0));

    IPoolConfigManager.LaneFee[] memory fees =
      s_feeStore.getFeeBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR), FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    IPoolConfigManager.LaneFee memory disabledFee;
    assertEq(abi.encode(disabledFee), abi.encode(fees[0]));
  }

  function test_getFeeBatch_RemovedInlineFeeRecord() public {
    _setInlineFeeConfig(_getInlineFeeConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, _getFeeConfig()));
    _setInlineDestOverheadConfig(
      _getInlineDestOverheadConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, _getDestOverheadConfig())
    );
    s_feeStore.applyInlineFeeConfigUpdates(
      _getLaneKeys(s_otherToken, DEST_CHAIN_SELECTOR), new FeePolicyStore.InlineFeeConfigArgs[](0)
    );

    IPoolConfigManager.LaneFee[] memory fees =
      s_feeStore.getFeeBatch(_getLaneKeys(s_otherToken, DEST_CHAIN_SELECTOR), FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    IPoolConfigManager.LaneFee memory disabledFee;
    assertEq(abi.encode(disabledFee), abi.encode(fees[0]));
  }

  function test_getFeeBatch_NoDestOverheadRecord() public {
    _setTokenDestOverheadPolicy(s_token, bytes32(0));

    IPoolConfigManager.LaneFee[] memory fees =
      s_feeStore.getFeeBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR), FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    IPoolConfigManager.LaneFee memory disabledFee;
    assertEq(abi.encode(disabledFee), abi.encode(fees[0]));
  }

  /// @dev The fee of a lane with default finality.
  function _getLaneFee(
    FeePolicyStore.FeeConfig memory feeConfig,
    FeePolicyStore.DestOverheadConfig memory destOverheadConfig
  ) private pure returns (IPoolConfigManager.LaneFee memory) {
    return IPoolConfigManager.LaneFee({
      feeUSDCents: feeConfig.finalityFeeUSDCents,
      destGasOverhead: destOverheadConfig.destGasOverhead,
      destBytesOverhead: destOverheadConfig.destBytesOverhead,
      tokenFeeBps: feeConfig.finalityTransferFeeBps,
      isEnabled: true
    });
  }
}
