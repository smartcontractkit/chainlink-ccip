// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../../../../interfaces/IPoolV2.sol";

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";

contract FeePolicyStore_getTokenTransferFeeConfigBatch is FeePolicyStoreSetup {
  function test_getTokenTransferFeeConfigBatch() public {
    // s_otherToken gets inline records with other values, so each lane must resolve its own records.
    _setInlineFeeConfig(_getInlineFeeConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, _getOtherFeeConfig()));
    _setInlineDestOverheadConfig(
      _getInlineDestOverheadConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, _getOtherDestOverheadConfig())
    );
    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](3);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: DEST_CHAIN_SELECTOR});
    // s_otherToken has inline records on the dest chain only, so the third lane gets a disabled config.
    lanes[2] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: SOURCE_CHAIN_SELECTOR});

    IPoolV2.TokenTransferFeeConfig[] memory feeConfigs = s_feeStore.getTokenTransferFeeConfigBatch(lanes);

    assertEq(3, feeConfigs.length);
    assertEq(
      abi.encode(_getTokenTransferFeeConfig(_getFeeConfig(), _getDestOverheadConfig())), abi.encode(feeConfigs[0])
    );
    assertEq(
      abi.encode(_getTokenTransferFeeConfig(_getOtherFeeConfig(), _getOtherDestOverheadConfig())),
      abi.encode(feeConfigs[1])
    );
    IPoolV2.TokenTransferFeeConfig memory disabledFeeConfig;
    assertEq(abi.encode(disabledFeeConfig), abi.encode(feeConfigs[2]));
  }

  function test_getTokenTransferFeeConfigBatch_DisabledFeeRecord() public {
    FeePolicyStore.FeeConfig memory disabledFeeRecord;
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, disabledFeeRecord));

    IPoolV2.TokenTransferFeeConfig[] memory feeConfigs =
      s_feeStore.getTokenTransferFeeConfigBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));

    IPoolV2.TokenTransferFeeConfig memory disabledFeeConfig;
    assertEq(abi.encode(disabledFeeConfig), abi.encode(feeConfigs[0]));
  }

  function test_getTokenTransferFeeConfigBatch_NoFeeRecord() public {
    _setTokenFeePolicy(s_token, bytes32(0));

    IPoolV2.TokenTransferFeeConfig[] memory feeConfigs =
      s_feeStore.getTokenTransferFeeConfigBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));

    IPoolV2.TokenTransferFeeConfig memory disabledFeeConfig;
    assertEq(abi.encode(disabledFeeConfig), abi.encode(feeConfigs[0]));
  }

  function test_getTokenTransferFeeConfigBatch_NoDestOverheadRecord() public {
    _setTokenDestOverheadPolicy(s_token, bytes32(0));

    IPoolV2.TokenTransferFeeConfig[] memory feeConfigs =
      s_feeStore.getTokenTransferFeeConfigBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));

    IPoolV2.TokenTransferFeeConfig memory disabledFeeConfig;
    assertEq(abi.encode(disabledFeeConfig), abi.encode(feeConfigs[0]));
  }

  function _getTokenTransferFeeConfig(
    FeePolicyStore.FeeConfig memory feeConfig,
    FeePolicyStore.DestOverheadConfig memory destOverheadConfig
  ) private pure returns (IPoolV2.TokenTransferFeeConfig memory) {
    return IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: destOverheadConfig.destGasOverhead,
      destBytesOverhead: destOverheadConfig.destBytesOverhead,
      finalityFeeUSDCents: feeConfig.finalityFeeUSDCents,
      fastFinalityFeeUSDCents: feeConfig.fastFinalityFeeUSDCents,
      finalityTransferFeeBps: feeConfig.finalityTransferFeeBps,
      fastFinalityTransferFeeBps: feeConfig.fastFinalityTransferFeeBps,
      isEnabled: true
    });
  }
}
