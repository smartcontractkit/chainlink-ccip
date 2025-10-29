// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Client} from "../../../libraries/Client.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_getTokenTransferFeeConfig is TokenPoolV2Setup {
  function test_getTokenTransferFeeConfig() public {
    // Set up a fee config first.
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: 32,
      defaultBlockConfirmationFeeUSDCents: 100, // $1.00
      customBlockConfirmationFeeUSDCents: 150, // $1.50
      defaultBlockConfirmationTransferFeeBps: 123,
      customBlockConfirmationTransferFeeBps: 456
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    // Test getting the config
    Client.EVM2AnyMessage memory message;
    IPoolV2.TokenTransferFeeConfig memory returnedFeeConfig =
      s_tokenPool.getTokenTransferFeeConfig(address(s_token), DEST_CHAIN_SELECTOR, message, 0, "");

    assertEq(returnedFeeConfig.destGasOverhead, feeConfig.destGasOverhead);
    assertEq(returnedFeeConfig.destBytesOverhead, feeConfig.destBytesOverhead);
    assertEq(returnedFeeConfig.defaultBlockConfirmationFeeUSDCents, feeConfig.defaultBlockConfirmationFeeUSDCents);
    assertEq(returnedFeeConfig.customBlockConfirmationFeeUSDCents, feeConfig.customBlockConfirmationFeeUSDCents);
    assertEq(returnedFeeConfig.defaultBlockConfirmationTransferFeeBps, feeConfig.defaultBlockConfirmationTransferFeeBps);
    assertEq(returnedFeeConfig.customBlockConfirmationTransferFeeBps, feeConfig.customBlockConfirmationTransferFeeBps);
  }

  function test_getTokenTransferFeeConfig_DeleteConfig() public {
    uint64[] memory toDelete = new uint64[](1);
    toDelete[0] = DEST_CHAIN_SELECTOR;
    vm.expectEmit();
    emit TokenPool.TokenTransferFeeConfigDeleted(DEST_CHAIN_SELECTOR);
    s_tokenPool.applyTokenTransferFeeConfigUpdates(new TokenPool.TokenTransferFeeConfigArgs[](0), toDelete);

    // Test getting the deleted config
    Client.EVM2AnyMessage memory message;
    IPoolV2.TokenTransferFeeConfig memory tokenTransferFeeConfig =
      s_tokenPool.getTokenTransferFeeConfig(address(s_token), DEST_CHAIN_SELECTOR, message, 0, "");

    // assert default values are returned
    assertEq(tokenTransferFeeConfig.destGasOverhead, 0);
    assertEq(tokenTransferFeeConfig.destBytesOverhead, 0);
    assertEq(tokenTransferFeeConfig.defaultBlockConfirmationFeeUSDCents, 0);
    assertEq(tokenTransferFeeConfig.customBlockConfirmationFeeUSDCents, 0);
    assertEq(tokenTransferFeeConfig.defaultBlockConfirmationTransferFeeBps, 0);
    assertEq(tokenTransferFeeConfig.customBlockConfirmationTransferFeeBps, 0);
  }
}
