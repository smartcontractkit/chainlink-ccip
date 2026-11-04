// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_getTokenTransferFeeConfig is MultiTokenPoolSetup {
  function test_getTokenTransferFeeConfig() public {
    IPoolV2.TokenTransferFeeConfig[] memory feeConfigs = new IPoolV2.TokenTransferFeeConfig[](1);
    feeConfigs[0] = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 200_000,
      destBytesOverhead: 256,
      finalityFeeUSDCents: 100,
      fastFinalityFeeUSDCents: 200,
      finalityTransferFeeBps: 10,
      fastFinalityTransferFeeBps: 20,
      isEnabled: true
    });
    // The mock matches only the lane that the pool must forward.
    vm.mockCall(
      s_poolConfigManager,
      abi.encodeCall(IPoolConfigManager.getTokenTransferFeeConfigBatch, (_laneKeys())),
      abi.encode(feeConfigs)
    );

    IPoolV2.TokenTransferFeeConfig memory feeConfig =
      s_pool.getTokenTransferFeeConfig(address(s_token), DEST_CHAIN_SELECTOR, bytes4(0), "");

    assertEq(abi.encode(feeConfigs[0]), abi.encode(feeConfig));
  }
}
