// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";
import {IRouterClient} from "../../../interfaces/IRouterClient.sol";

import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolHelperSetup.t.sol";

contract FastTransferTokenPoolHelper_getCcipSendTokenFee_Test is FastTransferTokenPoolSetup {
  address public s_receiver;

  function setUp() public override {
    super.setUp();
    s_receiver = vm.randomAddress();
  }

  function test_GetCcipSendTokenFee() public {
    uint256 amount = 100 ether;
    bytes memory receiver = abi.encodePacked(s_receiver);
    bytes memory extraArgs = "";
    uint256 fastFee = amount * FAST_FEE_BPS / 10000;
    uint256 feeQuoterQuote = 1 ether;
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(feeQuoterQuote)
    );
    IFastTransferPool.Quote memory quote =
      s_tokenPool.getCcipSendTokenFee(address(s_token), DEST_CHAIN_SELECTOR, amount, receiver, extraArgs);

    // Fast fee should be 1% of amount (100 bps)
    assertEq(quote.fastTransferFee, fastFee);
    // CCIP fee should be non-zero
    assertEq(quote.ccipSettlementFee, feeQuoterQuote);
  }
}
