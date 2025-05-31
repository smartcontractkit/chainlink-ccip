// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {BurnMintFastTransferTokenPoolSetup} from "./BurnMintFastTransferTokenPoolSetup.t.sol";

contract BurnMintFastTransferTokenPool_fastFill is BurnMintFastTransferTokenPoolSetup {
  uint256 internal constant FILL_AMOUNT = 100 ether;
  address internal constant RECEIVER = address(0x1234);
  bytes32 internal constant FILL_REQUEST_ID = keccak256("fillRequestId");
  uint8 internal constant SRC_DECIMALS = 18;

  function setUp() public virtual override {
    super.setUp();
    vm.stopPrank();
    // Give filler tokens to fill with
    deal(address(s_burnMintERC20), s_filler, type(uint256).max);
    vm.prank(s_filler);
    s_burnMintERC20.approve(address(s_pool), type(uint256).max);
  }

  function test_FastFill() public {
    uint256 fillerBalanceBefore = s_burnMintERC20.balanceOf(s_filler);
    uint256 receiverBalanceBefore = s_burnMintERC20.balanceOf(RECEIVER);

    bytes32 fillId = s_pool.computeFillId(FILL_REQUEST_ID, FILL_AMOUNT, RECEIVER);

    vm.expectEmit();
    emit IFastTransferPool.FastTransferFilled(FILL_REQUEST_ID, fillId, s_filler, FILL_AMOUNT, RECEIVER);

    vm.prank(s_filler);
    s_pool.fastFill(FILL_REQUEST_ID, DEST_CHAIN_SELECTOR, FILL_AMOUNT, SRC_DECIMALS, RECEIVER);

    assertEq(s_burnMintERC20.balanceOf(s_filler), fillerBalanceBefore - FILL_AMOUNT);
    assertEq(s_burnMintERC20.balanceOf(RECEIVER), receiverBalanceBefore + FILL_AMOUNT);

    s_pool.getFillInfo(fillId);
  }

  function test_FastFill_WithDifferentDecimals() public {
    uint8 sourceDecimals = 6; // USDC-like decimals
    uint256 srcAmountToFill = 100e6; // 100 tokens with 6 decimals
    uint256 expectedLocalAmount = 100 ether; // Should be scaled to 18 decimals

    uint256 fillerBalanceBefore = s_burnMintERC20.balanceOf(s_filler);
    uint256 receiverBalanceBefore = s_burnMintERC20.balanceOf(RECEIVER);

    vm.prank(s_filler);
    s_pool.fastFill(FILL_REQUEST_ID, DEST_CHAIN_SELECTOR, srcAmountToFill, sourceDecimals, RECEIVER);

    assertEq(s_burnMintERC20.balanceOf(s_filler), fillerBalanceBefore - expectedLocalAmount);
    assertEq(s_burnMintERC20.balanceOf(RECEIVER), receiverBalanceBefore + expectedLocalAmount);
    FastTransferTokenPoolAbstract.FillInfo memory fillInfo =
      s_pool.getFillInfo(s_pool.computeFillId(FILL_REQUEST_ID, expectedLocalAmount, RECEIVER));
    assertTrue(fillInfo.state == FastTransferTokenPoolAbstract.FillState.FILLED);
    assertEq(fillInfo.filler, s_filler);
  }

  function test_RevertWhen_AlreadyFilled() public {
    // First fill
    vm.prank(s_filler);
    s_pool.fastFill(FILL_REQUEST_ID, DEST_CHAIN_SELECTOR, FILL_AMOUNT, SRC_DECIMALS, RECEIVER);

    // Try to fill again
    vm.expectRevert(abi.encodeWithSelector(IFastTransferPool.AlreadyFilled.selector, FILL_REQUEST_ID));
    vm.prank(s_filler);
    s_pool.fastFill(FILL_REQUEST_ID, DEST_CHAIN_SELECTOR, FILL_AMOUNT, SRC_DECIMALS, RECEIVER);
  }

  function test_RevertWhen_FillerNotWhitelisted() public {
    address unauthorizedFiller = makeAddr("unauthorizedFiller");
    deal(address(s_burnMintERC20), unauthorizedFiller, type(uint256).max);
    vm.prank(unauthorizedFiller);
    s_burnMintERC20.approve(address(s_pool), type(uint256).max);

    vm.expectRevert(
      abi.encodeWithSelector(
        FastTransferTokenPoolAbstract.FillerNotAllowlisted.selector, DEST_CHAIN_SELECTOR, unauthorizedFiller
      )
    );
    vm.prank(unauthorizedFiller);
    s_pool.fastFill(FILL_REQUEST_ID, DEST_CHAIN_SELECTOR, FILL_AMOUNT, SRC_DECIMALS, RECEIVER);
  }

  function test_FastFill_WithWhitelistDisabled() public {
    // Disable whitelist
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastTransferBpsFee: FAST_FEE_BPS,
      fillerAllowlistEnabled: false, // disabled
      destinationPool: abi.encode(s_remoteBurnMintPool),
      maxFillAmountPerRequest: FILL_AMOUNT_MAX,
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      evmToAnyMessageExtraArgsBytes: ""
    });
    vm.prank(OWNER);
    s_pool.updateDestChainConfig(laneConfigArgs);

    address anyFiller = makeAddr("anyFiller");
    deal(address(s_burnMintERC20), anyFiller, type(uint256).max);
    vm.prank(anyFiller);
    s_burnMintERC20.approve(address(s_pool), type(uint256).max);

    uint256 receiverBalanceBefore = s_burnMintERC20.balanceOf(RECEIVER);

    vm.prank(anyFiller);
    s_pool.fastFill(FILL_REQUEST_ID, DEST_CHAIN_SELECTOR, FILL_AMOUNT, SRC_DECIMALS, RECEIVER);

    assertEq(s_burnMintERC20.balanceOf(RECEIVER), receiverBalanceBefore + FILL_AMOUNT);

    FastTransferTokenPoolAbstract.FillInfo memory fillInfo =
      s_pool.getFillInfo(s_pool.computeFillId(FILL_REQUEST_ID, FILL_AMOUNT, RECEIVER));
    assertTrue(fillInfo.state == FastTransferTokenPoolAbstract.FillState.FILLED);
    assertEq(fillInfo.filler, anyFiller);
  }

  function test_FastFill_MultipleFillers() public {
    address filler2 = makeAddr("filler2");

    // Add second filler to whitelist
    address[] memory addFillers = new address[](1);
    addFillers[0] = filler2;
    vm.prank(OWNER);
    s_pool.updateFillerAllowList(DEST_CHAIN_SELECTOR, addFillers, new address[](0));

    deal(address(s_burnMintERC20), filler2, type(uint256).max);
    vm.prank(filler2);
    s_burnMintERC20.approve(address(s_pool), type(uint256).max);

    bytes32 fillRequestId2 = keccak256("fillRequestId2");

    // Both fillers can fill different requests
    vm.prank(s_filler);
    s_pool.fastFill(FILL_REQUEST_ID, DEST_CHAIN_SELECTOR, FILL_AMOUNT, SRC_DECIMALS, RECEIVER);

    vm.prank(filler2);
    s_pool.fastFill(fillRequestId2, DEST_CHAIN_SELECTOR, FILL_AMOUNT, SRC_DECIMALS, RECEIVER);

    assertEq(s_burnMintERC20.balanceOf(RECEIVER), FILL_AMOUNT * 2);
  }
}
