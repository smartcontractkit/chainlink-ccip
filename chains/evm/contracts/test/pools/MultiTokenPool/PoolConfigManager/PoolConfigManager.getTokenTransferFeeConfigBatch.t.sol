// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";

import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

contract PoolConfigManager_getTokenTransferFeeConfigBatch is PoolConfigManagerSetup {
  /// @notice Returns resolved fees and destination overheads. Missing lanes return a disabled config with zero
  /// values even when fee records exist.
  function test_getTokenTransferFeeConfigBatch() public {
    FeePolicyStore.FeeConfig memory feeConfig = _getFeeConfig(100, 10);
    _setFeePolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, feeConfig);
    _setFeePolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, feeConfig);

    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](2);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_tokenB, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_tokenB, remoteChainSelector: OTHER_CHAIN_SELECTOR});
    IPoolV2.TokenTransferFeeConfig[] memory feeConfigs = s_poolConfigManager.getTokenTransferFeeConfigBatch(lanes);

    assertEq(DEFAULT_TOKEN_DEST_GAS_OVERHEAD, feeConfigs[0].destGasOverhead);
    assertEq(MIN_DEST_BYTES_OVERHEAD, feeConfigs[0].destBytesOverhead);
    assertEq(feeConfig.finalityFeeUSDCents, feeConfigs[0].finalityFeeUSDCents);
    assertEq(feeConfig.fastFinalityFeeUSDCents, feeConfigs[0].fastFinalityFeeUSDCents);
    assertEq(feeConfig.finalityTransferFeeBps, feeConfigs[0].finalityTransferFeeBps);
    assertEq(feeConfig.fastFinalityTransferFeeBps, feeConfigs[0].fastFinalityTransferFeeBps);
    assertTrue(feeConfigs[0].isEnabled);
    assertEq(abi.encode(IPoolV2.TokenTransferFeeConfig(0, 0, 0, 0, 0, 0, false)), abi.encode(feeConfigs[1]));
  }
}
