// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {IRouterClient} from "../../../interfaces/IRouterClient.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BurnMintFastTransferTokenPoolSetup} from "./BurnMintFastTransferTokenPoolSetup.t.sol";

contract BurnMintFastTransferTokenPool_ccipSendToken is BurnMintFastTransferTokenPoolSetup {
  uint256 internal constant TRANSFER_AMOUNT = 100 ether;
  address internal constant RECEIVER = address(0x1234);
  uint256 internal constant CCIP_SEND_FEE = 1 ether; // Mocked fee for sending tokens via CCIP
  bytes32 internal constant MESSAGE_ID = keccak256("messageId");

  function setUp() public virtual override {
    super.setUp();
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(CCIP_SEND_FEE)
    );
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector), abi.encode(MESSAGE_ID)
    );
    deal(address(s_token), OWNER, TRANSFER_AMOUNT * 10);
    s_token.approve(address(s_pool), type(uint256).max);
  }

  function test_CcipSendToken() public {
    uint256 balanceBefore = s_token.balanceOf(OWNER);

    IFastTransferPool.Quote memory quote = s_pool.getCcipSendTokenFee(
      address(0), // native fee token
      DEST_CHAIN_SELECTOR,
      TRANSFER_AMOUNT,
      abi.encode(RECEIVER),
      ""
    );

    uint256 expectedFastFee = (TRANSFER_AMOUNT * FAST_FEE_BPS) / 10_000;
    assertEq(quote.fastTransferFee, expectedFastFee);
    assertEq(quote.ccipSettlementFee, CCIP_SEND_FEE);

    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested(
      MESSAGE_ID, DEST_CHAIN_SELECTOR, TRANSFER_AMOUNT, expectedFastFee, abi.encode(RECEIVER)
    );

    bytes32 fillRequestId = s_pool.ccipSendToken{value: quote.ccipSettlementFee}(
      address(0), // native fee token
      DEST_CHAIN_SELECTOR,
      TRANSFER_AMOUNT,
      abi.encode(RECEIVER),
      ""
    );

    assertEq(fillRequestId, MESSAGE_ID);
    assertEq(s_token.balanceOf(OWNER), balanceBefore - TRANSFER_AMOUNT);
  }

  function test_CcipSendToken_WithERC20FeeToken() public {
    // Setup fee token
    address feeToken = address(s_token);
    uint256 balanceBefore = s_token.balanceOf(OWNER);

    IFastTransferPool.Quote memory quote =
      s_pool.getCcipSendTokenFee(feeToken, DEST_CHAIN_SELECTOR, TRANSFER_AMOUNT, abi.encode(RECEIVER), "");

    bytes32 fillRequestId =
      s_pool.ccipSendToken(feeToken, DEST_CHAIN_SELECTOR, TRANSFER_AMOUNT, abi.encode(RECEIVER), "");

    assertTrue(fillRequestId != bytes32(0));
    assertEq(s_token.balanceOf(OWNER), balanceBefore - TRANSFER_AMOUNT - quote.ccipSettlementFee);
  }

  function test_CcipSendToken_ToSVM() public {}

  function test_RevertWhen_CursedByRMN() public {
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));

    vm.expectRevert(TokenPool.CursedByRMN.selector);
    s_pool.ccipSendToken{value: 1 ether}(address(0), DEST_CHAIN_SELECTOR, TRANSFER_AMOUNT, abi.encode(RECEIVER), "");
  }
}
