  // SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.10;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";
import {IRouterClient} from "../../../interfaces/IRouterClient.sol";

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";

import {FastTransferTokenPoolHelperSetup} from "./FastTransferTokenPoolHelperSetup.t.sol";

contract FastTransferTokenPoolHelper_ccipSendToken_Test is FastTransferTokenPoolHelperSetup {
  function setUp() public override {
    super.setUp();
  }

  function test_CcipSendToken_NativeFee() public {
    uint256 amount = 100 ether;
    uint256 balanceBefore = s_token.balanceOf(OWNER);
    bytes memory receiver = abi.encodePacked(address(0x5));
    bytes memory extraArgs = "";
    bytes32 mockMessageId = keccak256("mockMessageId");
    uint256 fastFeeExpected = amount * FAST_FEE_BPS / 10000;
    vm.mockCall(address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(1 ether));
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector), abi.encode(mockMessageId)
    );
    vm.expectEmit();
    emit IFastTransferPool.FastFillRequest(mockMessageId, DEST_CHAIN_SELECTOR, amount, fastFeeExpected, receiver);
    bytes32 fillRequestId =
      s_tokenPool.ccipSendToken{value: 1 ether}(address(0), DEST_CHAIN_SELECTOR, amount, receiver, extraArgs);

    // Verify fillRequestId is non-zero
    assertEq(fillRequestId, mockMessageId);

    // Verify token balances
    assertEq(s_token.balanceOf(OWNER), balanceBefore - amount - fastFeeExpected);
    assertEq(s_token.balanceOf(address(s_tokenPool)), amount + fastFeeExpected);
  }

  function test_CcipSendToken_RevertWhen_LaneDisabled() public {
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

    vm.expectRevert(abi.encodeWithSelector(IFastTransferPool.LaneDisabled.selector));
    s_tokenPool.ccipSendToken(address(s_token), DEST_CHAIN_SELECTOR, 100 ether, receiver, extraArgs);
  }
}
