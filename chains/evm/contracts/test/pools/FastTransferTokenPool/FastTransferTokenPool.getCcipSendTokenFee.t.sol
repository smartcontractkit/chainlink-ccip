// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";
import {IRouterClient} from "../../../interfaces/IRouterClient.sol";

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

contract FastTransferTokenPool_getCcipSendTokenFee_Test is FastTransferTokenPoolSetup {
  function test_getCcipSendTokenFee() public {
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

  function test_getCcipSendTokenFee_WithNativeFeeToken() public {
    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: 9999,
      remotePoolAddresses: new bytes[](0),
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_pool.applyChainUpdates(new uint64[](0), chainUpdates);

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
