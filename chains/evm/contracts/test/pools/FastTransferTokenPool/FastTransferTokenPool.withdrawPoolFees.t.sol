// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract FastTransferTokenPool_withdrawPoolFees_Test is FastTransferTokenPoolSetup {
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
  }

  function test_withdrawPoolFees() public {
    uint16 fillerFeeBps = 75; // 0.75%
    uint16 poolFeeBps = 25; // 0.25%

    // Update config to include pool fee
    _updateConfigWithPoolFee(fillerFeeBps, poolFeeBps);

    uint256 poolFeeAmount = (SOURCE_AMOUNT * poolFeeBps) / 10_000;
    uint256 amountToFill = SOURCE_AMOUNT - (SOURCE_AMOUNT * (fillerFeeBps + poolFeeBps)) / 10_000;

    bytes32 fillId =
      s_pool.computeFillId(MESSAGE_ID, SOURCE_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, abi.encode(RECEIVER));

    // Fast fill and settlement to accumulate fees
    vm.prank(s_filler);
    s_pool.fastFill(fillId, MESSAGE_ID, SOURCE_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, RECEIVER);

    Client.Any2EVMMessage memory message =
      _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, fillerFeeBps, poolFeeBps);
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);

    // Verify fee accumulated
    assertEq(s_pool.getAccumulatedPoolFees(), poolFeeAmount);

    // Withdraw fees as pool owner
    address feeRecipient = makeAddr("feeRecipient");
    uint256 recipientBalanceBefore = s_token.balanceOf(feeRecipient);

    vm.expectEmit();
    emit IFastTransferPool.PoolFeeWithdrawn(feeRecipient, poolFeeAmount);

    vm.prank(OWNER);
    s_pool.withdrawPoolFees(feeRecipient);

    // Verify withdrawal
    assertEq(s_token.balanceOf(feeRecipient), recipientBalanceBefore + poolFeeAmount);
    assertEq(s_pool.getAccumulatedPoolFees(), 0);
  }

  function test_withdrawPoolFees_ZeroFees() public {
    address feeRecipient = makeAddr("feeRecipient");
    uint256 recipientBalanceBefore = s_token.balanceOf(feeRecipient);

    // Verify no fees accumulated
    assertEq(s_pool.getAccumulatedPoolFees(), 0);

    // No event should be emitted for zero fee withdrawals (gas optimization)
    // vm.expectEmit(); // Removed - no event emitted for zero amounts

    // Withdraw fees (should be no-op)
    vm.prank(OWNER);
    s_pool.withdrawPoolFees(feeRecipient);

    // Verify no change
    assertEq(s_token.balanceOf(feeRecipient), recipientBalanceBefore);
    assertEq(s_pool.getAccumulatedPoolFees(), 0);
  }

  function test_withdrawPoolFees_MultipleFees() public {
    uint16 fillerFeeBps = 50; // 0.5%
    uint16 poolFeeBps = 50; // 0.5%

    // Update config to include pool fee
    _updateConfigWithPoolFee(fillerFeeBps, poolFeeBps);

    uint256 poolFeeAmount = (SOURCE_AMOUNT * poolFeeBps) / 10_000;
    uint256 amountToFill = SOURCE_AMOUNT - (SOURCE_AMOUNT * (fillerFeeBps + poolFeeBps)) / 10_000;

    // First fast fill round
    bytes32 fillId1 =
      s_pool.computeFillId(MESSAGE_ID, SOURCE_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, abi.encode(RECEIVER));
    vm.prank(s_filler);
    s_pool.fastFill(fillId1, MESSAGE_ID, SOURCE_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, RECEIVER);

    Client.Any2EVMMessage memory message1 =
      _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, fillerFeeBps, poolFeeBps);
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message1);

    // Second fast fill round with different parameters
    bytes32 messageId2 = bytes32("messageId2");
    bytes32 fillId2 =
      s_pool.computeFillId(messageId2, SOURCE_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, abi.encode(RECEIVER));
    vm.prank(s_filler);
    s_pool.fastFill(fillId2, messageId2, SOURCE_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, RECEIVER);

    Client.Any2EVMMessage memory message2 = Client.Any2EVMMessage({
      messageId: messageId2,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      sender: abi.encode(SOURCE_POOL),
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          receiver: abi.encode(RECEIVER),
          sourceAmount: SOURCE_AMOUNT,
          fastTransferFillerFeeBps: fillerFeeBps,
          fastTransferPoolFeeBps: poolFeeBps,
          sourceDecimals: SOURCE_DECIMALS
        })
      ),
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message2);

    // Verify accumulated fees (should be 2x poolFeeAmount)
    assertEq(s_pool.getAccumulatedPoolFees(), poolFeeAmount * 2);

    // Withdraw all fees
    address feeRecipient = makeAddr("feeRecipient");
    uint256 recipientBalanceBefore = s_token.balanceOf(feeRecipient);

    vm.expectEmit();
    emit IFastTransferPool.PoolFeeWithdrawn(feeRecipient, poolFeeAmount * 2);

    vm.prank(OWNER);
    s_pool.withdrawPoolFees(feeRecipient);

    // Verify withdrawal
    assertEq(s_token.balanceOf(feeRecipient), recipientBalanceBefore + (poolFeeAmount * 2));
    assertEq(s_pool.getAccumulatedPoolFees(), 0);
  }

  function test_withdrawPoolFees_RevertWhen_NotOwner() public {
    address notOwner = makeAddr("notOwner");
    address feeRecipient = makeAddr("feeRecipient");

    vm.prank(notOwner);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_pool.withdrawPoolFees(feeRecipient);
  }

  function test_withdrawPoolFees_WithZeroAddress() public {
    uint16 fillerFeeBps = 100; // 1%
    uint16 poolFeeBps = 100; // 1%

    // Update config and accumulate some fees
    _updateConfigWithPoolFee(fillerFeeBps, poolFeeBps);
    _accumulatePoolFees(fillerFeeBps, poolFeeBps);

    // Verify we have fees to withdraw
    uint256 expectedFeeAmount = (SOURCE_AMOUNT * poolFeeBps) / 10_000;
    assertEq(s_pool.getAccumulatedPoolFees(), expectedFeeAmount);

    // Withdraw to zero address should revert
    vm.prank(OWNER);
    vm.expectRevert("ERC20: transfer to the zero address");
    s_pool.withdrawPoolFees(address(0));
  }

  function _updateConfigWithPoolFee(uint16 fillerFeeBps, uint16 poolFeeBps) internal {
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

    vm.prank(OWNER);
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    address[] memory fillersToAdd = new address[](1);
    fillersToAdd[0] = s_filler;
    vm.prank(OWNER);
    s_pool.updateFillerAllowList(fillersToAdd, new address[](0));
  }

  function _accumulatePoolFees(uint16 fillerFeeBps, uint16 poolFeeBps) internal {
    uint256 amountToFill = SOURCE_AMOUNT - (SOURCE_AMOUNT * (fillerFeeBps + poolFeeBps)) / 10_000;
    bytes32 fillId =
      s_pool.computeFillId(MESSAGE_ID, SOURCE_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, abi.encode(RECEIVER));

    vm.prank(s_filler);
    s_pool.fastFill(fillId, MESSAGE_ID, SOURCE_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, RECEIVER);

    Client.Any2EVMMessage memory message =
      _generateMintMessage(RECEIVER, SOURCE_AMOUNT, SOURCE_DECIMALS, fillerFeeBps, poolFeeBps);
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);
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
