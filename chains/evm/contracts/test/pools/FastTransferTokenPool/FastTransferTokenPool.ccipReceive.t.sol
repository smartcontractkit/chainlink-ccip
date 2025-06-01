// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {CCIPReceiver} from "../../../applications/CCIPReceiver.sol";
import {Client} from "../../../libraries/Client.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

import {IERC20Metadata} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/extensions/IERC20Metadata.sol";

contract FastTransferTokenPool_ccipReceive_Test is FastTransferTokenPoolSetup {
  bytes32 public constant MESSAGE_ID = bytes32("messageId");
  address public constant SOURCE_POOL = address(0x123);
  address public constant RECEIVER = address(0x5);

  function setUp() public override {
    super.setUp();

    vm.stopPrank();

    deal(address(s_token), address(s_pool), SOURCE_AMOUNT * 2); // Ensure pool has enough balance
    deal(address(s_token), s_filler, SOURCE_AMOUNT * 2); // Ensure filler has enough balance

    // Approve the pool to transfer tokens on behalf of the filler
    vm.prank(s_filler);
    s_token.approve(address(s_pool), type(uint256).max);

    vm.startPrank(address(s_sourceRouter));
  }

  function test_ccipReceive_SlowFill() public {
    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);
    uint256 expectedAmount = receiverBalanceBefore + SOURCE_AMOUNT;

    Client.Any2EVMMessage memory message = _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, FAST_FEE_BPS);

    vm.expectEmit();
    emit IFastTransferPool.FastTransferSettled(MESSAGE_ID);

    s_pool.ccipReceive(message);

    // Verify receiver got the full amount (transfer + fee)
    assertEq(s_token.balanceOf(RECEIVER), expectedAmount);
  }

  function test_ccipReceive_FastFill_Settlement() public {
    uint256 amountToFill = SOURCE_AMOUNT - (SOURCE_AMOUNT * FAST_FEE_BPS / 10_000);
    bytes32 fillId = s_pool.computeFillId(MESSAGE_ID, amountToFill, SOURCE_DECIMALS, RECEIVER);

    vm.stopPrank();
    vm.prank(s_filler);
    s_pool.fastFill(MESSAGE_ID, fillId, SOURCE_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, RECEIVER);

    uint256 fillerBalanceBefore = s_token.balanceOf(s_filler);
    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);

    // Settlement
    Client.Any2EVMMessage memory message = _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, FAST_FEE_BPS);

    vm.expectEmit();
    emit IFastTransferPool.FastTransferSettled(MESSAGE_ID);

    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);

    // Verify filler was reimbursed (transfer amount + fee)
    assertEq(s_token.balanceOf(s_filler), fillerBalanceBefore + SOURCE_AMOUNT);
    // Verify receiver balance didn't change (already received from fast fill)
    assertEq(s_token.balanceOf(RECEIVER), receiverBalanceBefore);
  }

  function test_ccipReceive_WithDifferentDecimals() public {
    uint8 sourceDecimals = 6; // USDC-like decimals
    uint8 destDecimals = IERC20Metadata(address(s_token)).decimals();
    require(sourceDecimals != destDecimals, "Test requires different source and destination decimals");

    uint256 sourceAmount = 100 * 10 ** sourceDecimals; // 100 tokens in source decimals
    uint256 expectedLocalAmount = sourceAmount * 10 ** 18 / 10 ** sourceDecimals; // Should be scaled to 18 decimals

    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);

    Client.Any2EVMMessage memory message = _generateMintMessage(RECEIVER, sourceAmount, sourceDecimals, FAST_FEE_BPS);

    s_pool.ccipReceive(message);

    // Verify receiver got the scaled amount
    assertEq(s_token.balanceOf(RECEIVER), receiverBalanceBefore + expectedLocalAmount);
  }

  function test_ccipReceive_ZeroFastTransferFeeBps() public {
    uint16 zeroFee = 0;
    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);

    // Prepare CCIP message with zero fast transfer fee
    Client.Any2EVMMessage memory message = _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, zeroFee);

    vm.expectEmit();
    emit IFastTransferPool.FastTransferSettled(MESSAGE_ID);

    s_pool.ccipReceive(message);

    // Verify receiver got only the transfer amount (no fee)
    assertEq(s_token.balanceOf(RECEIVER), receiverBalanceBefore + SOURCE_AMOUNT);
  }

  function test_ccipReceive_RevertWhen_AlreadySettled() public {
    Client.Any2EVMMessage memory message = _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, FAST_FEE_BPS);

    // First settlement
    s_pool.ccipReceive(message);

    // Try to settle again - should revert
    vm.expectRevert(abi.encodeWithSelector(IFastTransferPool.AlreadySettled.selector, MESSAGE_ID));
    s_pool.ccipReceive(message);
  }

  function test_ccipReceive_RevertWhen_InvalidData() public {
    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: MESSAGE_ID,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      sender: abi.encode(SOURCE_POOL),
      data: abi.encode(SOURCE_AMOUNT, SOURCE_DECIMALS),
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });

    vm.expectRevert(); // Should revert due to invalid data format
    s_pool.ccipReceive(message);
  }

  function test_ccipReceive_RevertWhen_NotRouter() public {
    Client.Any2EVMMessage memory message = _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, FAST_FEE_BPS);

    address notRouter = makeAddr("notRouter");
    // Call from non-router address should revert
    vm.stopPrank();
    vm.prank(notRouter);

    vm.expectRevert(abi.encodeWithSelector(CCIPReceiver.InvalidRouter.selector, notRouter));
    s_pool.ccipReceive(message);
  }

  function _generateMintMessage(
    address receiver,
    uint256 sourceAmount,
    uint8 sourceDecimals,
    uint16 fastTransferFeeBps
  ) internal pure returns (Client.Any2EVMMessage memory) {
    return Client.Any2EVMMessage({
      messageId: MESSAGE_ID,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      sender: abi.encode(SOURCE_POOL),
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          receiver: abi.encode(receiver),
          sourceAmount: sourceAmount,
          sourceDecimals: sourceDecimals,
          fastTransferFeeBps: fastTransferFeeBps
        })
      ),
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });
  }
}
