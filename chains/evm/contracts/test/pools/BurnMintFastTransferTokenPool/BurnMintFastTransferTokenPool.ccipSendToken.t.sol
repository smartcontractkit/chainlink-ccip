// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";
import {IRouterClient} from "../../../interfaces/IRouterClient.sol";

import {BurnMintFastTransferTokenPoolSetup} from "./BurnMintFastTransferTokenPoolSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract BurnMintFastTransferTokenPool_ccipSendToken is BurnMintFastTransferTokenPoolSetup {
  uint256 internal constant CCIP_SEND_FEE = 1 ether;
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

  function test_ccipSendToken_Success() public {
    uint256 maxFastTransferFee = (TRANSFER_AMOUNT * FAST_FEE_FILLER_BPS) / 10_000;
    uint256 expectedFastTransferFee = maxFastTransferFee;
    uint256 expectedFillerFee = expectedFastTransferFee; // All fee goes to filler in basic tests
    uint256 expectedPoolFee = 0; // No pool fee in basic tests
    uint256 expectedAmountNetFee = TRANSFER_AMOUNT - expectedFastTransferFee;

    uint256 balanceBefore = s_token.balanceOf(OWNER);
    uint256 poolBalanceBefore = s_token.balanceOf(address(s_pool));

    bytes32 expectedFillId = s_pool.computeFillId(
      MESSAGE_ID, SOURCE_CHAIN_SELECTOR, expectedAmountNetFee, SOURCE_DECIMALS, abi.encode(RECEIVER)
    );

    vm.expectEmit();
    emit IERC20.Transfer(OWNER, address(s_pool), TRANSFER_AMOUNT);

    vm.expectEmit();
    emit IERC20.Transfer(address(s_pool), address(0), TRANSFER_AMOUNT);

    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested({
      destinationChainSelector: DEST_CHAIN_SELECTOR,
      fillId: expectedFillId,
      settlementId: MESSAGE_ID,
      sourceAmountNetFee: expectedAmountNetFee,
      sourceDecimals: SOURCE_DECIMALS,
      fillerFee: expectedFillerFee,
      poolFee: expectedPoolFee,
      destinationPool: abi.encode(s_remoteBurnMintPool),
      receiver: abi.encode(RECEIVER)
    });

    bytes32 settlementId = s_pool.ccipSendToken{value: CCIP_SEND_FEE}(
      DEST_CHAIN_SELECTOR,
      TRANSFER_AMOUNT,
      maxFastTransferFee,
      abi.encode(RECEIVER),
      address(0), // native fee token
      ""
    );

    assertEq(settlementId, MESSAGE_ID);
    assertEq(s_token.balanceOf(OWNER), balanceBefore - TRANSFER_AMOUNT);
    // Pool should have 0 balance because tokens were burned
    assertEq(s_token.balanceOf(address(s_pool)), poolBalanceBefore);
  }

  function test_ccipSendToken_WithERC20FeeToken() public {
    address feeToken = address(s_token);
    uint256 maxFastTransferFee = (TRANSFER_AMOUNT * FAST_FEE_FILLER_BPS) / 10_000;

    uint256 balanceBefore = s_token.balanceOf(OWNER);
    uint256 poolBalanceBefore = s_token.balanceOf(address(s_pool));

    IFastTransferPool.Quote memory quote =
      s_pool.getCcipSendTokenFee(DEST_CHAIN_SELECTOR, TRANSFER_AMOUNT, abi.encode(RECEIVER), feeToken, "");

    uint256 expectedAmountNetFee = TRANSFER_AMOUNT - quote.fastTransferFee;
    uint256 expectedFillerFee = quote.fastTransferFee; // All fee goes to filler in basic tests
    uint256 expectedPoolFee = 0; // No pool fee in basic tests
    bytes32 expectedFillId = s_pool.computeFillId(
      MESSAGE_ID, SOURCE_CHAIN_SELECTOR, expectedAmountNetFee, SOURCE_DECIMALS, abi.encode(RECEIVER)
    );

    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested({
      destinationChainSelector: DEST_CHAIN_SELECTOR,
      fillId: expectedFillId,
      settlementId: MESSAGE_ID,
      sourceAmountNetFee: expectedAmountNetFee,
      sourceDecimals: SOURCE_DECIMALS,
      fillerFee: expectedFillerFee,
      poolFee: expectedPoolFee,
      destinationPool: abi.encode(s_remoteBurnMintPool),
      receiver: abi.encode(RECEIVER)
    });

    bytes32 settlementId =
      s_pool.ccipSendToken(DEST_CHAIN_SELECTOR, TRANSFER_AMOUNT, maxFastTransferFee, abi.encode(RECEIVER), feeToken, "");

    assertEq(settlementId, MESSAGE_ID);
    assertEq(s_token.balanceOf(OWNER), balanceBefore - TRANSFER_AMOUNT - quote.ccipSettlementFee);
    // Pool should have the settlement fee but not the transfer amount (since it was burned)
    assertEq(s_token.balanceOf(address(s_pool)), poolBalanceBefore + quote.ccipSettlementFee);
  }
}
