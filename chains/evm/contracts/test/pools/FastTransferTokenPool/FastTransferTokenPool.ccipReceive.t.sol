// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";
import {Client} from "../../../libraries/Client.sol";

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";

import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

contract FastTransferTokenPool_ccipReceive_Test is FastTransferTokenPoolSetup {
  bytes32 public messageId;
  address public sourcePool;
  uint64 public sourceChainSelector;
  uint256 public srcAmount;
  uint8 public sourceDecimals;
  address public receiver;

  function setUp() public override {
    super.setUp();
    vm.stopPrank();
    messageId = bytes32("messageId");
    sourcePool = address(0x123);
    sourceChainSelector = 1;
    srcAmount = 100 ether;
    sourceDecimals = 18;
    receiver = address(0x5);
    deal(address(s_token), address(s_pool), srcAmount * 2); // Ensure pool has enough balance
    deal(address(s_token), s_filler, srcAmount * 2); // Ensure filler has enough balance
    vm.prank(s_filler);
    s_token.approve(address(s_pool), type(uint256).max);
  }

  function test_CcipReceive_NotFastFilled() public {
    uint256 receiverBalanceBefore = s_token.balanceOf(receiver);
    uint256 expectedAmount = srcAmount;

    // Prepare CCIP message
    bytes memory data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        receiver: abi.encode(receiver),
        sourceAmount: srcAmount,
        sourceDecimals: sourceDecimals,
        fastTransferFeeBps: FAST_FEE_BPS
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
    emit IFastTransferPool.FastTransferSettled(messageId);
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);

    // Verify receiver got the full amount (transfer + fee)
    assertEq(s_token.balanceOf(receiver), receiverBalanceBefore + expectedAmount);
  }

  function test_CcipReceive_FastFilled() public {
    // First, fast fill the request
    uint256 amountToFill = srcAmount - (srcAmount * FAST_FEE_BPS / 10_000);
    bytes32 fillId = s_pool.computeFillId(messageId, amountToFill, sourceDecimals, receiver);
    
    vm.prank(s_filler);
    s_pool.fastFill(
      messageId, fillId, sourceChainSelector, amountToFill, sourceDecimals, receiver
    );

    uint256 fillerBalanceBefore = s_token.balanceOf(s_filler);
    uint256 receiverBalanceBefore = s_token.balanceOf(receiver);

    // Prepare CCIP message
    bytes memory data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        receiver: abi.encode(receiver),
        sourceAmount: srcAmount,
        sourceDecimals: sourceDecimals,
        fastTransferFeeBps: FAST_FEE_BPS
      })
    );
    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: messageId,
      sourceChainSelector: sourceChainSelector,
      sender: abi.encode(sourcePool),
      data: data,
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });

    vm.expectEmit();
    emit IFastTransferPool.FastTransferSettled(messageId);
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);

    // Verify filler was reimbursed (transfer amount + fee)
    assertEq(s_token.balanceOf(s_filler), fillerBalanceBefore + srcAmount);
    // Verify receiver balance didn't change (already received from fast fill)
    assertEq(s_token.balanceOf(receiver), receiverBalanceBefore);
  }

  function test_CcipReceive_WithDifferentDecimals() public {
    sourceDecimals = 6; // USDC-like decimals
    srcAmount = 100e6; // 100 tokens with 6 decimals
    uint256 expectedLocalAmount = 100 ether; // Should be scaled to 18 decimals

    uint256 receiverBalanceBefore = s_token.balanceOf(receiver);

    // Prepare CCIP message with different decimals
    bytes memory data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        receiver: abi.encode(receiver),
        sourceAmount: srcAmount,
        sourceDecimals: sourceDecimals,
        fastTransferFeeBps: FAST_FEE_BPS
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
    s_pool.ccipReceive(message);

    // Verify receiver got the scaled amount
    assertEq(s_token.balanceOf(receiver), receiverBalanceBefore + expectedLocalAmount);
  }

  function test_CcipReceive_ZeroFastTransferFee() public {
    uint16 zeroFee = 0;
    uint256 receiverBalanceBefore = s_token.balanceOf(receiver);

    // Prepare CCIP message with zero fast transfer fee
    bytes memory data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        receiver: abi.encode(receiver),
        sourceAmount: srcAmount,
        sourceDecimals: sourceDecimals,
        fastTransferFeeBps: zeroFee
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
    emit IFastTransferPool.FastTransferSettled(messageId);
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);

    // Verify receiver got only the transfer amount (no fee)
    assertEq(s_token.balanceOf(receiver), receiverBalanceBefore + srcAmount);
  }

  function test_RevertWhen_AlreadySettled() public {
    // Prepare CCIP message
    bytes memory data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        receiver: abi.encode(receiver),
        sourceAmount: srcAmount,
        sourceDecimals: sourceDecimals,
        fastTransferFeeBps: FAST_FEE_BPS
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
    s_pool.ccipReceive(message);

    // Try to settle again - should revert
    vm.expectRevert(abi.encodeWithSelector(IFastTransferPool.AlreadySettled.selector, messageId));
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);
  }

  function test_RevertWhen_InvalidData() public {
    // Prepare CCIP message with invalid data
    bytes memory invalidData = abi.encode(srcAmount, sourceDecimals); // Missing fastTransferFee and receiver
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
    s_pool.ccipReceive(message);
  }

  function test_RevertWhen_NotRouter() public {
    // Prepare CCIP message
    bytes memory data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        receiver: abi.encode(receiver),
        sourceAmount: srcAmount,
        sourceDecimals: sourceDecimals,
        fastTransferFeeBps: FAST_FEE_BPS
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
    s_pool.ccipReceive(message);
  }
}
