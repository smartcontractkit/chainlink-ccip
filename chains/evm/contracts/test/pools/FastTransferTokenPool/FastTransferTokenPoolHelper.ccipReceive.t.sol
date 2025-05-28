// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";
import {Client} from "../../../libraries/Client.sol";

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";

import {FastTransferTokenPoolHelperSetup} from "./FastTransferTokenPoolHelperSetup.t.sol";

contract FastTransferTokenPoolHelper_ccipReceive_Test is FastTransferTokenPoolHelperSetup {
  bytes32 public messageId;
  address public sourcePool;
  uint64 public sourceChainSelector;
  uint256 public srcAmount;
  uint8 public srcDecimals;
  uint256 public fastTransferFee;
  address public receiver;

  function setUp() public override {
    super.setUp();
    vm.stopPrank();
    messageId = bytes32("messageId");
    sourcePool = address(0x123);
    sourceChainSelector = 1;
    srcAmount = 100 ether;
    srcDecimals = 18;
    fastTransferFee = srcAmount * FAST_FEE_BPS / 10000; // 1% fast fee
    receiver = address(0x5);
    deal(address(s_token), address(s_tokenPool), srcAmount * 2); // Ensure pool has enough balance
    deal(address(s_token), s_filler, srcAmount * 2); // Ensure filler has enough balance
    vm.prank(s_filler);
    s_token.approve(address(s_tokenPool), type(uint256).max);
  }

  function test_CcipReceive_NotFastFilled() public {
    uint256 receiverBalanceBefore = s_token.balanceOf(receiver);
    uint256 expectedAmount = srcAmount + fastTransferFee;

    // Prepare CCIP message
    bytes memory data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        receiver: abi.encode(receiver),
        srcAmountToTransfer: srcAmount,
        srcDecimals: srcDecimals,
        fastTransferFee: fastTransferFee
      })
    );
    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: messageId,
      sourceChainSelector: sourceChainSelector,
      sender: abi.encode(sourcePool),
      data: data,
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });

    // Mock router call
    vm.expectEmit(true, false, false, true);
    emit IFastTransferPool.FastFillSettled(messageId);
    vm.prank(address(s_sourceRouter));
    s_tokenPool.ccipReceive(message);

    // Verify receiver got the full amount (transfer + fee)
    assertEq(s_token.balanceOf(receiver), receiverBalanceBefore + expectedAmount);
  }

  function test_CcipReceive_FastFilled() public {
    // First, fast fill the request
    vm.prank(s_filler);
    s_tokenPool.fastFill(messageId, sourceChainSelector, srcAmount, srcDecimals, receiver);

    uint256 fillerBalanceBefore = s_token.balanceOf(s_filler);
    uint256 receiverBalanceBefore = s_token.balanceOf(receiver);
    uint256 expectedReimbursement = srcAmount + fastTransferFee;

    // Prepare CCIP message
    bytes memory data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        receiver: abi.encode(receiver),
        srcAmountToTransfer: srcAmount,
        srcDecimals: srcDecimals,
        fastTransferFee: fastTransferFee
      })
    );
    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: messageId,
      sourceChainSelector: sourceChainSelector,
      sender: abi.encode(sourcePool),
      data: data,
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });

    // Mock router call
    vm.expectEmit(true, false, false, true);
    emit IFastTransferPool.FastFillSettled(messageId);
    vm.prank(address(s_sourceRouter));
    s_tokenPool.ccipReceive(message);

    // Verify filler was reimbursed (transfer amount + fee)
    assertEq(s_token.balanceOf(s_filler), fillerBalanceBefore + expectedReimbursement);
    // Verify receiver balance didn't change (already received from fast fill)
    assertEq(s_token.balanceOf(receiver), receiverBalanceBefore);
  }

  function test_CcipReceive_WithDifferentDecimals() public {
    srcDecimals = 6; // USDC-like decimals
    srcAmount = 100e6; // 100 tokens with 6 decimals
    uint256 srcFee = srcAmount * FAST_FEE_BPS / 10000; // 1% fast fee
    uint256 expectedLocalAmount = 100 ether; // Should be scaled to 18 decimals
    uint256 expectedLocalFee = 1 ether; // Should be scaled to 18 decimals

    uint256 receiverBalanceBefore = s_token.balanceOf(receiver);

    // Prepare CCIP message with different decimals
    bytes memory data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        receiver: abi.encode(receiver),
        srcAmountToTransfer: srcAmount,
        srcDecimals: srcDecimals,
        fastTransferFee: srcFee
      })
    );
    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: messageId,
      sourceChainSelector: sourceChainSelector,
      sender: abi.encode(sourcePool),
      data: data,
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });

    vm.prank(address(s_sourceRouter));
    s_tokenPool.ccipReceive(message);

    // Verify receiver got the scaled amount
    assertEq(s_token.balanceOf(receiver), receiverBalanceBefore + expectedLocalAmount + expectedLocalFee);
  }

  function test_CcipReceive_ZeroFastTransferFee() public {
    uint256 zeroFee = 0;
    uint256 receiverBalanceBefore = s_token.balanceOf(receiver);

    // Prepare CCIP message with zero fast transfer fee
    bytes memory data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        receiver: abi.encode(receiver),
        srcAmountToTransfer: srcAmount,
        srcDecimals: srcDecimals,
        fastTransferFee: zeroFee
      })
    );
    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: messageId,
      sourceChainSelector: sourceChainSelector,
      sender: abi.encode(sourcePool),
      data: data,
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });

    vm.expectEmit(true, false, false, true);
    emit IFastTransferPool.FastFillSettled(messageId);
    vm.prank(address(s_sourceRouter));
    s_tokenPool.ccipReceive(message);

    // Verify receiver got only the transfer amount (no fee)
    assertEq(s_token.balanceOf(receiver), receiverBalanceBefore + srcAmount);
  }

  function test_RevertWhen_MessageAlreadySettled() public {
    // Prepare CCIP message
    bytes memory data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        receiver: abi.encode(receiver),
        srcAmountToTransfer: srcAmount,
        srcDecimals: srcDecimals,
        fastTransferFee: fastTransferFee
      })
    );
    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: messageId,
      sourceChainSelector: sourceChainSelector,
      sender: abi.encode(sourcePool),
      data: data,
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });

    // First settlement
    vm.prank(address(s_sourceRouter));
    s_tokenPool.ccipReceive(message);

    // Try to settle again - should revert
    vm.expectRevert(abi.encodeWithSelector(IFastTransferPool.MessageAlreadySettled.selector, messageId));
    vm.prank(address(s_sourceRouter));
    s_tokenPool.ccipReceive(message);
  }

  function test_RevertWhen_InvalidData() public {
    // Prepare CCIP message with invalid data
    bytes memory invalidData = abi.encode(srcAmount, srcDecimals); // Missing fastTransferFee and receiver
    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: messageId,
      sourceChainSelector: sourceChainSelector,
      sender: abi.encode(sourcePool),
      data: invalidData,
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });

    // Mock router call
    vm.prank(address(s_sourceRouter));
    vm.expectRevert(); // Should revert due to invalid data format
    s_tokenPool.ccipReceive(message);
  }

  function test_RevertWhen_NotRouter() public {
    // Prepare CCIP message
    bytes memory data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        receiver: abi.encode(receiver),
        srcAmountToTransfer: srcAmount,
        srcDecimals: srcDecimals,
        fastTransferFee: fastTransferFee
      })
    );
    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: messageId,
      sourceChainSelector: sourceChainSelector,
      sender: abi.encode(sourcePool),
      data: data,
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });

    // Call from non-router address should revert
    vm.expectRevert();
    vm.prank(makeAddr("notRouter"));
    s_tokenPool.ccipReceive(message);
  }

  function test_CcipReceive_FastFilledThenSettled_Integration() public {
    // Step 1: Fast fill
    uint256 fillerBalanceBefore = s_token.balanceOf(s_filler);
    uint256 receiverBalanceBefore = s_token.balanceOf(receiver);

    vm.prank(s_filler);
    s_tokenPool.fastFill(messageId, sourceChainSelector, srcAmount, srcDecimals, receiver);

    // Verify fast fill worked
    assertEq(s_token.balanceOf(s_filler), fillerBalanceBefore - srcAmount);
    assertEq(s_token.balanceOf(receiver), receiverBalanceBefore + srcAmount);

    // Step 2: Settlement
    uint256 fillerBalanceAfterFill = s_token.balanceOf(s_filler);
    uint256 receiverBalanceAfterFill = s_token.balanceOf(receiver);

    bytes memory data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        receiver: abi.encode(receiver),
        srcAmountToTransfer: srcAmount,
        srcDecimals: srcDecimals,
        fastTransferFee: fastTransferFee
      })
    );
    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: messageId,
      sourceChainSelector: sourceChainSelector,
      sender: abi.encode(sourcePool),
      data: data,
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });

    vm.prank(address(s_sourceRouter));
    s_tokenPool.ccipReceive(message);

    // Verify filler was reimbursed (original amount + fee)
    assertEq(s_token.balanceOf(s_filler), fillerBalanceAfterFill + srcAmount + fastTransferFee);
    // Verify receiver balance didn't change (already got tokens from fast fill)
    assertEq(s_token.balanceOf(receiver), receiverBalanceAfterFill);
  }
}
