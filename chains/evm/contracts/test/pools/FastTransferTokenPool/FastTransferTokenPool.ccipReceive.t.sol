// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {CCIPReceiver} from "../../../applications/CCIPReceiver.sol";

import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";

import {TokenPool} from "../../../pools/TokenPool.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";

contract FastTransferTokenPool_ccipReceive_Test is FastTransferTokenPoolSetup {
  bytes32 public constant MESSAGE_ID = bytes32("messageId");
  address public constant SOURCE_POOL = address(0x123);

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

    Client.Any2EVMMessage memory message =
      _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, FAST_FEE_FILLER_BPS, 0);
    uint256 fastTransferFee = (SOURCE_AMOUNT * FAST_FEE_FILLER_BPS) / 10_000;
    bytes32 fillId = s_pool.computeFillId(
      message.messageId,
      message.sourceChainSelector,
      SOURCE_AMOUNT - fastTransferFee,
      SOURCE_DECIMALS,
      abi.encode(RECEIVER)
    );

    vm.expectEmit();
    emit TokenPool.InboundRateLimitConsumed(SOURCE_CHAIN_SELECTOR, address(s_token), SOURCE_AMOUNT);
    vm.expectEmit();
    emit IFastTransferPool.FastTransferSettled(fillId, MESSAGE_ID, 0, 0, IFastTransferPool.FillState.NOT_FILLED);

    s_pool.ccipReceive(message);

    // Verify receiver got the full amount (transfer + fee)
    assertEq(s_token.balanceOf(RECEIVER), expectedAmount);
  }

  function test_ccipReceive_FastFill_Settlement() public {
    uint256 amountToFill = SOURCE_AMOUNT - (SOURCE_AMOUNT * FAST_FEE_FILLER_BPS / 10_000);
    bytes32 fillId =
      s_pool.computeFillId(MESSAGE_ID, SOURCE_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, abi.encode(RECEIVER));

    vm.stopPrank();
    vm.prank(s_filler);
    s_pool.fastFill(fillId, MESSAGE_ID, SOURCE_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, RECEIVER);

    uint256 fillerBalanceBefore = s_token.balanceOf(s_filler);
    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);

    // Settlement
    Client.Any2EVMMessage memory message =
      _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, FAST_FEE_FILLER_BPS, 0);

    vm.expectEmit();
    emit TokenPool.InboundRateLimitConsumed(SOURCE_CHAIN_SELECTOR, address(s_token), SOURCE_AMOUNT);
    vm.expectEmit();
    emit IFastTransferPool.FastTransferSettled(fillId, MESSAGE_ID, SOURCE_AMOUNT, 0, IFastTransferPool.FillState.FILLED);

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
    assertTrue(sourceDecimals != destDecimals, "Test requires different source and destination decimals");

    uint256 sourceAmount = 100 * 10 ** sourceDecimals; // 100 tokens in source decimals
    uint256 expectedLocalAmount = sourceAmount * 10 ** 18 / 10 ** sourceDecimals; // Should be scaled to 18 decimals

    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);

    Client.Any2EVMMessage memory message =
      _generateMintMessage(RECEIVER, sourceAmount, sourceDecimals, FAST_FEE_FILLER_BPS, 0);

    vm.expectEmit();
    emit TokenPool.InboundRateLimitConsumed(SOURCE_CHAIN_SELECTOR, address(s_token), expectedLocalAmount);
    s_pool.ccipReceive(message);

    // Verify receiver got the scaled amount
    assertEq(s_token.balanceOf(RECEIVER), receiverBalanceBefore + expectedLocalAmount);
  }

  function test_ccipReceive_ZeroFastTransferFeeBps() public {
    uint16 zeroFee = 0;
    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);

    // Prepare CCIP message with zero fast transfer fee
    Client.Any2EVMMessage memory message =
      _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, zeroFee, zeroFee);

    bytes32 fillId = s_pool.computeFillId(
      message.messageId, message.sourceChainSelector, SOURCE_AMOUNT - 0, SOURCE_DECIMALS, abi.encode(RECEIVER)
    );

    vm.expectEmit();
    emit TokenPool.InboundRateLimitConsumed(SOURCE_CHAIN_SELECTOR, address(s_token), SOURCE_AMOUNT);
    vm.expectEmit();
    emit IFastTransferPool.FastTransferSettled(fillId, MESSAGE_ID, 0, 0, IFastTransferPool.FillState.NOT_FILLED);

    s_pool.ccipReceive(message);

    // Verify receiver got only the transfer amount (no fee)
    assertEq(s_token.balanceOf(RECEIVER), receiverBalanceBefore + SOURCE_AMOUNT);
  }

  function test_ccipReceive_SlowFill_WithPoolFee() public {
    uint16 fillerFeeBps = 50; // 0.5%
    uint16 poolFeeBps = 50; // 0.5%

    // Update config to include pool fee
    _updateConfigWithPoolFee(fillerFeeBps, poolFeeBps);

    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);
    uint256 poolBalanceBefore = s_token.balanceOf(address(s_pool));
    uint256 accumulatedFeesBefore = s_pool.getAccumulatedPoolFees();

    uint256 expectedReceiverAmount = SOURCE_AMOUNT;

    Client.Any2EVMMessage memory message =
      _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, fillerFeeBps, poolFeeBps);

    // Compute fillId for the slow fill scenario
    uint256 fillerFeeAmount = (SOURCE_AMOUNT * fillerFeeBps) / 10_000;
    uint256 poolFeeAmount = (SOURCE_AMOUNT * poolFeeBps) / 10_000;
    uint256 amountAfterFees = SOURCE_AMOUNT - fillerFeeAmount - poolFeeAmount;
    bytes32 fillId = s_pool.computeFillId(
      message.messageId, message.sourceChainSelector, amountAfterFees, SOURCE_DECIMALS, abi.encode(RECEIVER)
    );

    vm.expectEmit();
    emit TokenPool.InboundRateLimitConsumed(SOURCE_CHAIN_SELECTOR, address(s_token), SOURCE_AMOUNT);
    vm.expectEmit();
    emit IFastTransferPool.FastTransferSettled(fillId, MESSAGE_ID, 0, 0, IFastTransferPool.FillState.NOT_FILLED);
    s_pool.ccipReceive(message);

    // Verify receiver gets the full amount (helper transfers, doesn't mint)
    assertEq(s_token.balanceOf(RECEIVER), receiverBalanceBefore + expectedReceiverAmount, "t1");
    // Pool balance should decrease by amount transferred to receiver (helper uses transfer)
    assertEq(s_token.balanceOf(address(s_pool)), poolBalanceBefore - expectedReceiverAmount, "t2");
    // Pool should NOT accumulate any fee for slow fills (no fast fill service provided)
    assertEq(s_pool.getAccumulatedPoolFees(), accumulatedFeesBefore, "t3");

    // For slow fills, ccipReceive should mark the fill as SETTLED directly
    FastTransferTokenPoolAbstract.FillInfo memory fillInfo = s_pool.getFillInfo(fillId);
    assertTrue(fillInfo.state == IFastTransferPool.FillState.SETTLED, "Fill state should be SETTLED after slow fill");
    assertEq(fillInfo.filler, address(0), "Filler should be address(0) for slow fills");
  }

  function test_ccipReceive_FastFill_Settlement_WithPoolFee() public {
    uint16 fillerFeeBps = 75; // 0.75%
    uint16 poolFeeBps = 25; // 0.25%

    // Update config to include pool fee
    _updateConfigWithPoolFee(fillerFeeBps, poolFeeBps);

    // Amount filler provides (original amount minus filler fee)
    uint256 fillerFeeAmount = (SOURCE_AMOUNT * fillerFeeBps) / 10_000;
    uint256 poolFeeAmount = (SOURCE_AMOUNT * poolFeeBps) / 10_000;
    uint256 amountToFill = SOURCE_AMOUNT - fillerFeeAmount - poolFeeAmount;

    bytes32 fillId =
      s_pool.computeFillId(MESSAGE_ID, SOURCE_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, abi.encode(RECEIVER));

    // Fast fill first
    vm.stopPrank();
    vm.prank(s_filler);
    s_pool.fastFill(fillId, MESSAGE_ID, SOURCE_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, RECEIVER);

    uint256 fillerBalanceBefore = s_token.balanceOf(s_filler);
    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);
    uint256 poolBalanceBefore = s_token.balanceOf(address(s_pool));
    uint256 accumulatedFeesBefore = s_pool.getAccumulatedPoolFees();

    // Settlement
    Client.Any2EVMMessage memory message =
      _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, fillerFeeBps, poolFeeBps);

    // Expected filler reimbursement = SOURCE_AMOUNT - poolFeeAmount
    uint256 expectedFillerReimbursement = SOURCE_AMOUNT - poolFeeAmount;
    vm.expectEmit();
    emit TokenPool.InboundRateLimitConsumed(SOURCE_CHAIN_SELECTOR, address(s_token), SOURCE_AMOUNT);
    vm.expectEmit();
    emit IFastTransferPool.FastTransferSettled(
      fillId, MESSAGE_ID, expectedFillerReimbursement, poolFeeAmount, IFastTransferPool.FillState.FILLED
    );

    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);

    // Verify filler gets reimbursed: what they provided + their fee
    assertEq(s_token.balanceOf(s_filler), fillerBalanceBefore + amountToFill + fillerFeeAmount, "t1");
    // Receiver balance shouldn't change (already got tokens from fast fill)
    assertEq(s_token.balanceOf(RECEIVER), receiverBalanceBefore, "t2");
    // Pool token balance should decrease by filler reimbursement (helper uses transfer, not mint)
    assertEq(s_token.balanceOf(address(s_pool)), poolBalanceBefore - (amountToFill + fillerFeeAmount), "t3");
    // Pool should have accumulated the pool fee
    assertEq(s_pool.getAccumulatedPoolFees(), accumulatedFeesBefore + poolFeeAmount, "t4");
  }

  function test_ccipReceive_FastFill_Settlement_WithDifferentDecimals() public {
    uint8 sourceDecimals = 6; // USDC-like decimals
    uint8 destDecimals = IERC20Metadata(address(s_token)).decimals();
    assertTrue(sourceDecimals != destDecimals, "Test requires different source and destination decimals");

    uint256 sourceAmount = 100 * 10 ** sourceDecimals; // 100 tokens in source decimals
    uint256 expectedLocalAmount = sourceAmount * 10 ** destDecimals / 10 ** sourceDecimals; // Scale to dest decimals

    // Calculate amounts with fee deduction in source decimals
    uint256 fillerFeeAmount = (sourceAmount * FAST_FEE_FILLER_BPS) / 10_000;
    uint256 sourceAmountAfterFee = sourceAmount - fillerFeeAmount;

    bytes32 fillId = s_pool.computeFillId(
      MESSAGE_ID, SOURCE_CHAIN_SELECTOR, sourceAmountAfterFee, sourceDecimals, abi.encode(RECEIVER)
    );

    // Fast fill first
    vm.stopPrank();
    vm.prank(s_filler);
    s_pool.fastFill(fillId, MESSAGE_ID, SOURCE_CHAIN_SELECTOR, sourceAmountAfterFee, sourceDecimals, RECEIVER);

    uint256 fillerBalanceBefore = s_token.balanceOf(s_filler);
    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);

    // Settlement with different source decimals
    Client.Any2EVMMessage memory message =
      _generateMintMessage(RECEIVER, sourceAmount, sourceDecimals, FAST_FEE_FILLER_BPS, 0);

    // Expected scaled amounts for settlement
    uint256 expectedFillerReimbursement = expectedLocalAmount; // Full amount scaled to dest decimals

    vm.expectEmit();
    emit TokenPool.InboundRateLimitConsumed(SOURCE_CHAIN_SELECTOR, address(s_token), expectedLocalAmount);
    vm.expectEmit();
    emit IFastTransferPool.FastTransferSettled(
      fillId, MESSAGE_ID, expectedFillerReimbursement, 0, IFastTransferPool.FillState.FILLED
    );

    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);

    // Verify filler gets reimbursed with properly scaled amount (transfer amount + fee in dest decimals)
    assertEq(s_token.balanceOf(s_filler), fillerBalanceBefore + expectedLocalAmount);
    // Verify receiver balance didn't change (already received from fast fill)
    assertEq(s_token.balanceOf(RECEIVER), receiverBalanceBefore);
  }

  function test_ccipReceive_RevertWhen_AlreadySettled() public {
    Client.Any2EVMMessage memory message =
      _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, FAST_FEE_FILLER_BPS, 0);

    // First settlement
    s_pool.ccipReceive(message);

    uint256 amountToFill = SOURCE_AMOUNT - (SOURCE_AMOUNT * FAST_FEE_FILLER_BPS / 10_000);
    bytes32 fillId =
      s_pool.computeFillId(MESSAGE_ID, SOURCE_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, abi.encode(RECEIVER));

    // Try to settle again - should revert
    vm.expectRevert(abi.encodeWithSelector(IFastTransferPool.AlreadySettled.selector, fillId));
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
    Client.Any2EVMMessage memory message =
      _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, FAST_FEE_FILLER_BPS, 0);

    address notRouter = makeAddr("notRouter");
    // Call from non-router address should revert
    vm.stopPrank();
    vm.prank(notRouter);

    vm.expectRevert(abi.encodeWithSelector(CCIPReceiver.InvalidRouter.selector, notRouter));
    s_pool.ccipReceive(message);
  }

  function _updateConfigWithPoolFee(uint16 fillerFeeBps, uint16 poolFeeBps) internal {
    vm.stopPrank();
    vm.startPrank(OWNER);

    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: fillerFeeBps,
      fastTransferPoolFeeBps: poolFeeBps,
      fillerAllowlistEnabled: true,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: MAX_FILL_AMOUNT_PER_REQUEST,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    address[] memory fillersToAdd = new address[](1);
    fillersToAdd[0] = s_filler;
    s_pool.updateFillerAllowList(fillersToAdd, new address[](0));

    vm.stopPrank();
    vm.startPrank(address(s_sourceRouter));
  }

  function _generateMintMessage(
    address receiver,
    uint256 sourceAmount,
    uint8 sourceDecimals,
    uint16 fastTransferFillerFeeBps,
    uint16 fastTransferPoolFeeBps
  ) internal pure returns (Client.Any2EVMMessage memory) {
    return Client.Any2EVMMessage({
      messageId: MESSAGE_ID,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      sender: abi.encode(SOURCE_POOL),
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          receiver: abi.encode(receiver),
          sourceAmount: sourceAmount,
          fastTransferFillerFeeBps: fastTransferFillerFeeBps,
          fastTransferPoolFeeBps: fastTransferPoolFeeBps,
          sourceDecimals: sourceDecimals
        })
      ),
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });
  }
}
