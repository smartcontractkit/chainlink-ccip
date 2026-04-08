// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {TokenPool} from "../../../pools/TokenPool.sol";
import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";

contract TokenPool_getTokenTransferFeeConfig is AdvancedPoolHooksSetup {
  function test_getTokenTransferFeeConfig() public {
    // Set up a fee config first.
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: 32,
      finalityFeeUSDCents: 100, // $1.00
      fastFinalityFeeUSDCents: 150, // $1.50
      finalityTransferFeeBps: 123,
      fastFinalityTransferFeeBps: 456,
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    // Test getting the config
    IPoolV2.TokenTransferFeeConfig memory returnedFeeConfig =
      s_tokenPool.getTokenTransferFeeConfig(address(s_token), DEST_CHAIN_SELECTOR, 0, "");

    assertEq(returnedFeeConfig.destGasOverhead, feeConfig.destGasOverhead);
    assertEq(returnedFeeConfig.destBytesOverhead, feeConfig.destBytesOverhead);
    assertEq(returnedFeeConfig.finalityFeeUSDCents, feeConfig.finalityFeeUSDCents);
    assertEq(returnedFeeConfig.fastFinalityFeeUSDCents, feeConfig.fastFinalityFeeUSDCents);
    assertEq(returnedFeeConfig.finalityTransferFeeBps, feeConfig.finalityTransferFeeBps);
    assertEq(returnedFeeConfig.fastFinalityTransferFeeBps, feeConfig.fastFinalityTransferFeeBps);
    assertEq(returnedFeeConfig.isEnabled, feeConfig.isEnabled);
  }

  function test_getTokenTransferFeeConfig_DeleteConfig() public {
    uint64[] memory toDelete = new uint64[](1);
    toDelete[0] = DEST_CHAIN_SELECTOR;
    vm.expectEmit();
    emit TokenPool.TokenTransferFeeConfigDeleted(DEST_CHAIN_SELECTOR);
    s_tokenPool.applyTokenTransferFeeConfigUpdates(new TokenPool.TokenTransferFeeConfigArgs[](0), toDelete);

    // Test getting the deleted config
    IPoolV2.TokenTransferFeeConfig memory tokenTransferFeeConfig =
      s_tokenPool.getTokenTransferFeeConfig(address(s_token), DEST_CHAIN_SELECTOR, 0, "");

    // assert default values are returned
    assertEq(tokenTransferFeeConfig.destGasOverhead, 0);
    assertEq(tokenTransferFeeConfig.destBytesOverhead, 0);
    assertEq(tokenTransferFeeConfig.finalityFeeUSDCents, 0);
    assertEq(tokenTransferFeeConfig.fastFinalityFeeUSDCents, 0);
    assertEq(tokenTransferFeeConfig.finalityTransferFeeBps, 0);
    assertEq(tokenTransferFeeConfig.fastFinalityTransferFeeBps, 0);
    assertEq(tokenTransferFeeConfig.isEnabled, false);
  }
}
