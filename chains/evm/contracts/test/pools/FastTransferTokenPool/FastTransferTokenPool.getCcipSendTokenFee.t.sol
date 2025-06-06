// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";
import {IRouterClient} from "../../../interfaces/IRouterClient.sol";

import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

contract FastTransferTokenPool_getCcipSendTokenFee_Test is FastTransferTokenPoolSetup {
  function test_GetCcipSendTokenFee() public {
    uint256 fastFee = SOURCE_AMOUNT * FAST_FEE_FILLER_BPS / 10000;
    uint256 settlementQuote = 1 ether;

    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(settlementQuote)
    );
    IFastTransferPool.Quote memory quote =
      s_pool.getCcipSendTokenFee(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, abi.encodePacked(RECEIVER), address(s_token), "");

    assertEq(quote.fastTransferFee, fastFee);
    assertEq(quote.ccipSettlementFee, settlementQuote);
  }

  function test_GetCcipSendTokenFee_WithNativeFeeToken() public {
    uint256 fastFee = SOURCE_AMOUNT * FAST_FEE_FILLER_BPS / 10000;
    uint256 settlementQuote = 1 ether;

    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(settlementQuote)
    );

    // Test with native fee token (address(0))
    IFastTransferPool.Quote memory quote =
      s_pool.getCcipSendTokenFee(DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, abi.encodePacked(RECEIVER), address(0), "");

    assertEq(quote.fastTransferFee, fastFee);
    assertEq(quote.ccipSettlementFee, settlementQuote);
  }
}
