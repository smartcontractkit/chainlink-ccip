// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";
import {IRouterClient} from "../../../interfaces/IRouterClient.sol";

import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

contract FastTransferTokenPool_getCcipSendTokenFee_Test is FastTransferTokenPoolSetup {
  function test_GetCcipSendTokenFee() public {
    uint256 fastFee = SOURCE_AMOUNT * FAST_FEE_BPS / 10000;
    uint256 settlementQuote = 1 ether;

    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(settlementQuote)
    );
    IFastTransferPool.Quote memory quote =
      s_pool.getCcipSendTokenFee(address(s_token), DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, abi.encodePacked(RECEIVER), "");

    assertEq(quote.fastTransferFee, fastFee);
    assertEq(quote.ccipSettlementFee, settlementQuote);
  }
}
