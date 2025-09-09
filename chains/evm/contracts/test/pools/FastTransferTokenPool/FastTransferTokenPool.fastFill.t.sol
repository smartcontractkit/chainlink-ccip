// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

contract FastTransferTokenPool_fastFill_Test is FastTransferTokenPoolSetup {
  bytes32 public constant SETTLEMENT_ID = bytes32("settlementId");

  function setUp() public override {
    super.setUp();
    vm.stopPrank();
    deal(address(s_token), s_filler, 1000 ether);
    vm.startPrank(s_filler);
    s_token.approve(address(s_pool), type(uint256).max);
  }

  function test_FastFill() public {
    uint256 balanceBefore = s_token.balanceOf(s_filler);
    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);

    bytes32 fillId =
      s_pool.computeFillId(SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, abi.encode(RECEIVER));

    vm.expectEmit();
    emit IFastTransferPool.FastTransferFilled(fillId, SETTLEMENT_ID, s_filler, SOURCE_AMOUNT, RECEIVER);

    s_pool.fastFill(fillId, SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);

    // Verify token balances
    assertEq(s_token.balanceOf(s_filler), balanceBefore - SOURCE_AMOUNT);
    assertEq(s_token.balanceOf(RECEIVER), receiverBalanceBefore + SOURCE_AMOUNT);

    FastTransferTokenPoolAbstract.FillInfo memory fillInfo = s_pool.getFillInfo(fillId);
    assertTrue(fillInfo.state == IFastTransferPool.FillState.FILLED);
    assertEq(fillInfo.filler, s_filler);
  }

  function test_FastFill_WithDifferentDecimals() public {
    uint8 sourceDecimals = 6; // USDC-like decimals
    uint256 srcAmountToFill = 100e6; // 100 tokens with 6 decimals
    uint256 expectedLocalAmount = 100 ether; // Should be scaled to 18 decimals

    uint256 fillerBalanceBefore = s_token.balanceOf(s_filler);
    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);

    bytes32 fillId =
      s_pool.computeFillId(SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, srcAmountToFill, sourceDecimals, abi.encode(RECEIVER));

    vm.expectEmit();
    emit IFastTransferPool.FastTransferFilled(fillId, SETTLEMENT_ID, s_filler, expectedLocalAmount, RECEIVER);

    s_pool.fastFill(fillId, SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, srcAmountToFill, sourceDecimals, RECEIVER);

    assertEq(s_token.balanceOf(s_filler), fillerBalanceBefore - expectedLocalAmount);
    assertEq(s_token.balanceOf(RECEIVER), receiverBalanceBefore + expectedLocalAmount);
    FastTransferTokenPoolAbstract.FillInfo memory fillInfo = s_pool.getFillInfo(
      s_pool.computeFillId(SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, srcAmountToFill, sourceDecimals, abi.encode(RECEIVER))
    );
    assertTrue(fillInfo.state == IFastTransferPool.FillState.FILLED);
    assertEq(fillInfo.filler, s_filler);
  }

  function test_FastFill_AllowlistDisabled() public {
    vm.stopPrank();
    vm.prank(OWNER);
    // Disable allowlist
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: 100,
      fastTransferPoolFeeBps: 0, // No pool fee for this test
      fillerAllowlistEnabled: false,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: 1000 ether,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    vm.expectEmit();
    emit FastTransferTokenPoolAbstract.DestChainConfigUpdated(
      SOURCE_CHAIN_SELECTOR,
      100,
      0,
      1000 ether,
      destPoolAddress,
      Internal.CHAIN_FAMILY_SELECTOR_EVM,
      SETTLEMENT_GAS_OVERHEAD,
      false
    );

    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    address nonAllowlistedFiller = address(0x6);

    bytes32 fillId =
      s_pool.computeFillId(SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, abi.encode(RECEIVER));

    // Mint tokens to non-allowlisted filler
    deal(nonAllowlistedFiller, 1000 ether);
    deal(address(s_token), nonAllowlistedFiller, 1000 ether);
    vm.startPrank(nonAllowlistedFiller);
    s_token.approve(address(s_pool), type(uint256).max);

    vm.expectEmit();
    emit IFastTransferPool.FastTransferFilled(fillId, SETTLEMENT_ID, nonAllowlistedFiller, SOURCE_AMOUNT, RECEIVER);

    // Should succeed even though filler is not allowlisted
    s_pool.fastFill(fillId, SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);

    // Verify token balances
    assertEq(s_token.balanceOf(nonAllowlistedFiller), 1000 ether - SOURCE_AMOUNT);
    assertEq(s_token.balanceOf(RECEIVER), SOURCE_AMOUNT);
  }

  function test_FastFill_MultipleFillers() public {
    address filler2 = makeAddr("filler2");

    // Add second filler to allowlist
    address[] memory addFillers = new address[](1);
    addFillers[0] = filler2;
    vm.stopPrank();

    vm.expectEmit();
    emit FastTransferTokenPoolAbstract.FillerAllowListUpdated(addFillers, new address[](0));

    vm.prank(OWNER);
    s_pool.updateFillerAllowList(addFillers, new address[](0));

    deal(address(s_token), filler2, type(uint256).max);
    vm.prank(filler2);
    s_token.approve(address(s_pool), type(uint256).max);

    bytes32 settlementId2 = keccak256("settlementId2");

    bytes32 fillId1 =
      s_pool.computeFillId(SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, abi.encode(RECEIVER));
    bytes32 fillId2 =
      s_pool.computeFillId(settlementId2, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, abi.encode(RECEIVER));

    vm.expectEmit();
    emit IFastTransferPool.FastTransferFilled(fillId1, SETTLEMENT_ID, s_filler, SOURCE_AMOUNT, RECEIVER);

    // Both fillers can fill different requests
    vm.prank(s_filler);
    s_pool.fastFill(fillId1, SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);

    vm.expectEmit();
    emit IFastTransferPool.FastTransferFilled(fillId2, settlementId2, filler2, SOURCE_AMOUNT, RECEIVER);

    vm.prank(filler2);
    s_pool.fastFill(fillId2, settlementId2, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);

    assertEq(s_token.balanceOf(RECEIVER), SOURCE_AMOUNT * 2);
  }

  function test_FastFill_RevertWhen_AlreadyFilledOrSettled() public {
    bytes32 fillId =
      s_pool.computeFillId(SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, abi.encode(RECEIVER));

    // First fill
    s_pool.fastFill(fillId, SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);

    // Attempt second fill
    vm.expectRevert(abi.encodeWithSelector(IFastTransferPool.AlreadyFilledOrSettled.selector, fillId));
    s_pool.fastFill(fillId, SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);
  }

  function test_FastFill_RevertWhen_FillerNotAllowed() public {
    vm.stopPrank();

    address nonAllowlistedFiller = address(0x6);
    vm.startPrank(nonAllowlistedFiller);

    vm.expectRevert(
      abi.encodeWithSelector(
        FastTransferTokenPoolAbstract.FillerNotAllowlisted.selector, DEST_CHAIN_SELECTOR, nonAllowlistedFiller
      )
    );
    s_pool.fastFill(SETTLEMENT_ID, bytes32(0x0), DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);
  }

  function test_RevertWhen_InvalidFillId() public {
    // Use an incorrect fillId (different from what would be computed)
    bytes32 incorrectFillId = keccak256("incorrect_fill_id");

    vm.expectRevert(abi.encodeWithSelector(FastTransferTokenPoolAbstract.InvalidFillId.selector, incorrectFillId));
    s_pool.fastFill(incorrectFillId, SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);
  }

  function test_RevertWhen_InvalidFillId_WrongAmount() public {
    // Create fillId with different amount
    uint256 wrongAmount = SOURCE_AMOUNT + 1 ether;
    bytes32 fillIdWithWrongAmount =
      s_pool.computeFillId(SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, wrongAmount, SOURCE_DECIMALS, abi.encode(RECEIVER));

    vm.expectRevert(abi.encodeWithSelector(FastTransferTokenPoolAbstract.InvalidFillId.selector, fillIdWithWrongAmount));
    s_pool.fastFill(
      fillIdWithWrongAmount, SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER
    );
  }

  function test_RevertWhen_InvalidFillId_WrongDecimals() public {
    // Create fillId with different decimals
    uint8 wrongDecimals = 6;
    bytes32 fillIdWithWrongDecimals =
      s_pool.computeFillId(SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, wrongDecimals, abi.encode(RECEIVER));

    vm.expectRevert(
      abi.encodeWithSelector(FastTransferTokenPoolAbstract.InvalidFillId.selector, fillIdWithWrongDecimals)
    );
    s_pool.fastFill(
      fillIdWithWrongDecimals, SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER
    );
  }

  function test_RevertWhen_InvalidFillId_WrongReceiver() public {
    // Create fillId with different receiver
    address wrongReceiver = address(0x5678);
    bytes32 fillIdWithWrongReceiver = s_pool.computeFillId(
      SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, abi.encode(wrongReceiver)
    );

    vm.expectRevert(
      abi.encodeWithSelector(FastTransferTokenPoolAbstract.InvalidFillId.selector, fillIdWithWrongReceiver)
    );
    s_pool.fastFill(
      fillIdWithWrongReceiver, SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER
    );
  }

  function test_RevertWhen_InvalidFillId_WrongChainSelector() public {
    // Create fillId with different chain selector
    uint32 wrongChainSelector = uint32(uint256(keccak256("WRONG_CHAIN_SELECTOR")));
    bytes32 fillIdWithWrongChainSelector =
      s_pool.computeFillId(SETTLEMENT_ID, wrongChainSelector, SOURCE_AMOUNT, SOURCE_DECIMALS, abi.encode(RECEIVER));

    vm.expectRevert(
      abi.encodeWithSelector(FastTransferTokenPoolAbstract.InvalidFillId.selector, fillIdWithWrongChainSelector)
    );
    s_pool.fastFill(
      fillIdWithWrongChainSelector, SETTLEMENT_ID, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER
    );
  }
}
