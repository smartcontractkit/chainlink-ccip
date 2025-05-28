// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

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

    bytes32 fillId = keccak256(abi.encodePacked(FILL_REQUEST_ID, FILL_AMOUNT, RECEIVER));

    vm.expectEmit();
    emit IFastTransferPool.FastFill(FILL_REQUEST_ID, fillId, s_filler, FILL_AMOUNT, RECEIVER);

    vm.prank(s_filler);
    s_pool.fastFill(FILL_REQUEST_ID, DEST_CHAIN_SELECTOR, FILL_AMOUNT, SRC_DECIMALS, RECEIVER);

    assertEq(s_burnMintERC20.balanceOf(s_filler), fillerBalanceBefore - FILL_AMOUNT);
    assertEq(s_burnMintERC20.balanceOf(RECEIVER), receiverBalanceBefore + FILL_AMOUNT);
  }

  function test_FastFill_WithDifferentDecimals() public {
    uint8 srcDecimals = 6; // USDC-like decimals
    uint256 srcAmount = 100e6; // 100 tokens with 6 decimals
    uint256 expectedLocalAmount = 100 ether; // Should be scaled to 18 decimals

    uint256 fillerBalanceBefore = s_burnMintERC20.balanceOf(s_filler);
    uint256 receiverBalanceBefore = s_burnMintERC20.balanceOf(RECEIVER);

    vm.prank(s_filler);
    s_pool.fastFill(FILL_REQUEST_ID, DEST_CHAIN_SELECTOR, srcAmount, srcDecimals, RECEIVER);

    assertEq(s_burnMintERC20.balanceOf(s_filler), fillerBalanceBefore - expectedLocalAmount);
    assertEq(s_burnMintERC20.balanceOf(RECEIVER), receiverBalanceBefore + expectedLocalAmount);
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
        FastTransferTokenPoolAbstract.FillerNotWhitelisted.selector, DEST_CHAIN_SELECTOR, unauthorizedFiller
      )
    );
    vm.prank(unauthorizedFiller);
    s_pool.fastFill(FILL_REQUEST_ID, DEST_CHAIN_SELECTOR, FILL_AMOUNT, SRC_DECIMALS, RECEIVER);
  }

  function test_FastFill_WithWhitelistDisabled() public {
    // Disable whitelist
    FastTransferTokenPoolAbstract.LaneConfigArgs memory laneConfigArgs = FastTransferTokenPoolAbstract.LaneConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      bpsFastFee: FAST_FEE_BPS,
      fillerAllowlistEnabled: false, // disabled
      destinationPool: abi.encode(s_remoteBurnMintPool),
      fillAmountMaxPerRequest: FILL_AMOUNT_MAX,
      addFillers: new address[](0),
      removeFillers: new address[](0)
    });
    vm.prank(OWNER);
    s_pool.updateLaneConfig(laneConfigArgs);

    address anyFiller = makeAddr("anyFiller");
    deal(address(s_burnMintERC20), anyFiller, type(uint256).max);
    vm.prank(anyFiller);
    s_burnMintERC20.approve(address(s_pool), type(uint256).max);

    uint256 receiverBalanceBefore = s_burnMintERC20.balanceOf(RECEIVER);

    vm.prank(anyFiller);
    s_pool.fastFill(FILL_REQUEST_ID, DEST_CHAIN_SELECTOR, FILL_AMOUNT, SRC_DECIMALS, RECEIVER);

    assertEq(s_burnMintERC20.balanceOf(RECEIVER), receiverBalanceBefore + FILL_AMOUNT);
  }

  function test_FastFill_MultipleFillers() public {
    address filler2 = makeAddr("filler2");

    // Add second filler to whitelist
    address[] memory addFillers = new address[](1);
    addFillers[0] = filler2;
    vm.prank(OWNER);
    s_pool.updatefillerAllowList(DEST_CHAIN_SELECTOR, addFillers, new address[](0));

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
