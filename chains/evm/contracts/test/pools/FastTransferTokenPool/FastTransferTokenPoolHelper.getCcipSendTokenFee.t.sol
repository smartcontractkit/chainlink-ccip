// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.10;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";
import {IRouterClient} from "../../../interfaces/IRouterClient.sol";

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";

import {FastTransferTokenPoolHelperSetup} from "./FastTransferTokenPoolHelperSetup.t.sol";

contract FastTransferTokenPoolHelper_getCcipSendTokenFee_Test is FastTransferTokenPoolHelperSetup {
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
    assertEq(quote.sendTokenFee, feeQuoterQuote);
  }

  function test_GetCcipSendTokenFee_RevertWhen_LaneDisabled() public {
    FastTransferTokenPoolAbstract.LaneConfigArgs memory laneConfigArgs = FastTransferTokenPoolAbstract.LaneConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      bpsFastFee: 100,
      enabled: false,
      fillerAllowlistEnabled: true,
      destinationPool: address(0x4),
      fillAmountMaxPerRequest: 1000 ether,
      addFillers: new address[](0),
      removeFillers: new address[](0)
    });
    s_tokenPool.updateLaneConfig(laneConfigArgs);

    bytes memory receiver = abi.encodePacked(address(0x5));
    bytes memory extraArgs = "";

    vm.expectRevert(IFastTransferPool.LaneDisabled.selector);
    s_tokenPool.getCcipSendTokenFee(address(s_token), DEST_CHAIN_SELECTOR, 100 ether, receiver, extraArgs);
  }
}
