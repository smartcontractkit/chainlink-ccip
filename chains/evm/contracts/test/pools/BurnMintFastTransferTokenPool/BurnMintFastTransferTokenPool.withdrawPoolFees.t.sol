// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {BurnMintFastTransferTokenPoolSetup} from "./BurnMintFastTransferTokenPoolSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract BurnMintFastTransferTokenPool_withdrawPoolFees_Test is BurnMintFastTransferTokenPoolSetup {
  bytes32 public constant MESSAGE_ID = bytes32("messageId");

  function setUp() public override {
    super.setUp();

    vm.stopPrank();

    // Ensure filler has enough balance - minted by the token which already has mint permissions
    vm.prank(address(s_pool)); // Use pool's permissions to mint
    s_token.mint(s_filler, TRANSFER_AMOUNT * 2);

    // Approve the pool to transfer tokens on behalf of the filler
    vm.prank(s_filler);
    s_token.approve(address(s_pool), type(uint256).max);
  }

  function test_withdrawPoolFees() public {
    uint16 fillerFeeBps = 75; // 0.75%
    uint16 poolFeeBps = 25; // 0.25%

    // Update config to include pool fee
    _updateConfigWithPoolFee(fillerFeeBps, poolFeeBps);

    uint256 poolFeeAmount = (TRANSFER_AMOUNT * poolFeeBps) / 10_000;
    uint256 amountToFill = TRANSFER_AMOUNT - (TRANSFER_AMOUNT * (fillerFeeBps + poolFeeBps)) / 10_000;

    bytes32 fillId =
      s_pool.computeFillId(MESSAGE_ID, DEST_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, abi.encode(RECEIVER));

    // Get initial pool balance
    uint256 poolBalanceBefore = s_token.balanceOf(address(s_pool));

    // Fast fill and settlement to accumulate fees
    vm.prank(s_filler);
    s_pool.fastFill(fillId, MESSAGE_ID, DEST_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, RECEIVER);

    Client.Any2EVMMessage memory message =
      _generateMintMessage(RECEIVER, TRANSFER_AMOUNT, SOURCE_DECIMALS, fillerFeeBps, poolFeeBps);
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);

    // For burn/mint pools: getAccumulatedPoolFees() returns the pool balance (which includes accumulated fees)
    assertEq(s_pool.getAccumulatedPoolFees(), poolFeeAmount, "Burn/mint pools track fees via pool balance");

    // Verify pool fee was minted to the pool contract
    uint256 poolBalanceAfter = s_token.balanceOf(address(s_pool));
    assertEq(poolBalanceAfter, poolBalanceBefore + poolFeeAmount, "Pool fee should be minted to pool");

    // For burn/mint pools, withdrawPoolFees should transfer the entire pool balance to recipient
    address feeRecipient = makeAddr("feeRecipient");
    uint256 recipientBalanceBefore = s_token.balanceOf(feeRecipient);

    // Expect event for the pool balance withdrawal
    vm.expectEmit();
    emit IFastTransferPool.PoolFeeWithdrawn(feeRecipient, poolBalanceAfter);

    vm.prank(OWNER);
    s_pool.withdrawPoolFees(feeRecipient);

    // Verify recipient receives the entire pool balance (which includes both accumulated fees)
    assertEq(
      s_token.balanceOf(feeRecipient),
      recipientBalanceBefore + poolBalanceAfter,
      "Recipient should receive entire pool balance"
    );
    // Pool balance should be 0 after withdrawal
    assertEq(s_token.balanceOf(address(s_pool)), 0, "Pool balance should be 0 after withdrawal");
    // Accumulated fees should now be 0 after withdrawal
    assertEq(s_pool.getAccumulatedPoolFees(), 0, "Accumulated fees should be 0 after withdrawal");
  }

  function test_withdrawPoolFees_ZeroFees() public {
    address feeRecipient = makeAddr("feeRecipient");
    uint256 recipientBalanceBefore = s_token.balanceOf(feeRecipient);

    // Verify no fees accumulated
    assertEq(s_pool.getAccumulatedPoolFees(), 0);

    // Withdraw fees (should be no-op)
    vm.prank(OWNER);
    s_pool.withdrawPoolFees(feeRecipient);

    // Verify no change
    assertEq(s_token.balanceOf(feeRecipient), recipientBalanceBefore);
    assertEq(s_pool.getAccumulatedPoolFees(), 0);
  }

  function test_withdrawPoolFees_BothGasAnd() public {
    uint16 fillerFeeBps = 50; // 0.5%
    uint16 poolFeeBps = 50; // 0.5%

    // Update config to include pool fee
    _updateConfigWithPoolFee(fillerFeeBps, poolFeeBps);

    uint256 poolFeeAmount = (TRANSFER_AMOUNT * poolFeeBps) / 10_000;
    uint256 amountToFill = TRANSFER_AMOUNT - (TRANSFER_AMOUNT * (fillerFeeBps + poolFeeBps)) / 10_000;

    // Get initial pool balance
    uint256 poolBalanceBefore = s_token.balanceOf(address(s_pool));

    // First fast fill round
    bytes32 fillId1 =
      s_pool.computeFillId(MESSAGE_ID, DEST_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, abi.encode(RECEIVER));
    vm.prank(s_filler);
    s_pool.fastFill(fillId1, MESSAGE_ID, DEST_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, RECEIVER);

    Client.Any2EVMMessage memory message1 =
      _generateMintMessage(RECEIVER, TRANSFER_AMOUNT, SOURCE_DECIMALS, fillerFeeBps, poolFeeBps);
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message1);

    // Second fast fill round with different parameters
    bytes32 messageId2 = bytes32("messageId2");
    bytes32 fillId2 =
      s_pool.computeFillId(messageId2, DEST_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, abi.encode(RECEIVER));
    vm.prank(s_filler);
    s_pool.fastFill(fillId2, messageId2, DEST_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, RECEIVER);

    Client.Any2EVMMessage memory message2 = Client.Any2EVMMessage({
      messageId: messageId2,
      sourceChainSelector: DEST_CHAIN_SELECTOR,
      sender: abi.encode(s_remoteBurnMintPool),
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          receiver: abi.encode(RECEIVER),
          sourceAmount: TRANSFER_AMOUNT,
          fastTransferFillerFeeBps: fillerFeeBps,
          fastTransferPoolFeeBps: poolFeeBps,
          sourceDecimals: SOURCE_DECIMALS
        })
      ),
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message2);

    assertEq(s_pool.getAccumulatedPoolFees(), poolFeeAmount * 2, "Burn/mint pools track fees via pool balance");

    // Verify both pool fees were minted to the pool contract
    uint256 poolBalanceAfter = s_token.balanceOf(address(s_pool));
    assertEq(poolBalanceAfter, poolBalanceBefore + (poolFeeAmount * 2), "Both pool fees should be minted to pool");

    // For burn/mint pools, withdrawPoolFees should transfer the entire pool balance to recipient
    address feeRecipient = makeAddr("feeRecipient");
    uint256 recipientBalanceBefore = s_token.balanceOf(feeRecipient);

    // Expect event for the pool balance withdrawal
    vm.expectEmit();
    emit IFastTransferPool.PoolFeeWithdrawn(feeRecipient, poolBalanceAfter);

    vm.prank(OWNER);
    s_pool.withdrawPoolFees(feeRecipient);

    // Verify recipient receives the entire pool balance (which includes both accumulated fees)
    assertEq(
      s_token.balanceOf(feeRecipient),
      recipientBalanceBefore + poolBalanceAfter,
      "Recipient should receive entire pool balance"
    );

    assertEq(s_token.balanceOf(address(s_pool)), 0, "Pool balance should be 0 after withdrawal");

    assertEq(s_pool.getAccumulatedPoolFees(), 0, "Accumulated fees should be 0 after withdrawal");
  }

  function test_withdrawPoolFees_RevertWhen_NotOwner() public {
    address notOwner = makeAddr("notOwner");
    address feeRecipient = makeAddr("feeRecipient");

    vm.prank(notOwner);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_pool.withdrawPoolFees(feeRecipient);
  }

  function _updateConfigWithPoolFee(uint16 fillerFeeBps, uint16 poolFeeBps) internal {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: fillerFeeBps,
      fastTransferPoolFeeBps: poolFeeBps,
      fillerAllowlistEnabled: true,
      destinationPool: abi.encode(s_remoteBurnMintPool),
      maxFillAmountPerRequest: FILL_AMOUNT_MAX,
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
    uint256 amountToFill = TRANSFER_AMOUNT - (TRANSFER_AMOUNT * (fillerFeeBps + poolFeeBps)) / 10_000;
    bytes32 fillId =
      s_pool.computeFillId(MESSAGE_ID, DEST_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, abi.encode(RECEIVER));

    vm.prank(s_filler);
    s_pool.fastFill(fillId, MESSAGE_ID, DEST_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, RECEIVER);

    Client.Any2EVMMessage memory message =
      _generateMintMessage(RECEIVER, TRANSFER_AMOUNT, SOURCE_DECIMALS, fillerFeeBps, poolFeeBps);
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);
  }

  function _generateMintMessage(
    address receiver,
    uint256 sourceAmount,
    uint8 sourceDecimals,
    uint16 fastTransferFillerFeeBps,
    uint16 fastTransferPoolFeeBps
  ) internal view returns (Client.Any2EVMMessage memory) {
    return Client.Any2EVMMessage({
      messageId: MESSAGE_ID,
      sourceChainSelector: DEST_CHAIN_SELECTOR,
      sender: abi.encode(s_remoteBurnMintPool),
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
