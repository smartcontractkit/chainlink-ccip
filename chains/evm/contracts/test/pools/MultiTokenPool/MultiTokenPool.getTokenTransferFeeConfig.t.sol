// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_getTokenTransferFeeConfig is MultiTokenPoolSetup {
  function test_getTokenTransferFeeConfig() public {
    s_poolConfigManager.setTokenTransferFeeConfig(
      IPoolV2.TokenTransferFeeConfig({
        destGasOverhead: 200_000,
        destBytesOverhead: 256,
        finalityFeeUSDCents: 100,
        fastFinalityFeeUSDCents: 200,
        finalityTransferFeeBps: 10,
        fastFinalityTransferFeeBps: 20,
        isEnabled: true
      })
    );

    IPoolV2.TokenTransferFeeConfig memory feeConfig =
      s_pool.getTokenTransferFeeConfig(address(s_token), DEST_CHAIN_SELECTOR, bytes4(0), "");

    assertEq(feeConfig.destGasOverhead, 200_000);
    assertEq(feeConfig.destBytesOverhead, 256);
    assertEq(feeConfig.finalityFeeUSDCents, 100);
    assertEq(feeConfig.fastFinalityFeeUSDCents, 200);
    assertEq(feeConfig.finalityTransferFeeBps, 10);
    assertEq(feeConfig.fastFinalityTransferFeeBps, 20);
    assertTrue(feeConfig.isEnabled);
  }
}
