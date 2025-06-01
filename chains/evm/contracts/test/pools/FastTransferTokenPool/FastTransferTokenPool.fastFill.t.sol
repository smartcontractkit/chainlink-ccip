// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

contract FastTransferTokenPool_fastFill_Test is FastTransferTokenPoolSetup {
  bytes32 public constant FILL_REQUEST_ID = bytes32("fillRequestId");

  function setUp() public override {
    super.setUp();
    vm.stopPrank();
    deal(address(s_token), s_filler, 1000 ether);
    vm.startPrank(s_filler);
    s_token.approve(address(s_pool), type(uint256).max);
  }

  function test_FastFill() public {
    uint256 balanceBefore = s_token.balanceOf(s_filler);

    bytes32 fillId = s_pool.computeFillId(FILL_REQUEST_ID, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);

    s_pool.fastFill(FILL_REQUEST_ID, fillId, DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);

    // Verify token balances
    assertEq(s_token.balanceOf(s_filler), balanceBefore - SOURCE_AMOUNT);
    assertEq(s_token.balanceOf(RECEIVER), SOURCE_AMOUNT);
  }

  function test_FastFill_AllowlistDisabled() public {
    vm.stopPrank();
    vm.prank(OWNER);
    // Disable allowlist
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastTransferBpsFee: 100,
      fillerAllowlistEnabled: false,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: 1000 ether,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    address nonAllowlistedFiller = address(0x6);

    bytes32 fillId = s_pool.computeFillId(FILL_REQUEST_ID, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);

    // Mint tokens to non-allowlisted filler
    deal(nonAllowlistedFiller, 1000 ether);
    deal(address(s_token), nonAllowlistedFiller, 1000 ether);
    vm.startPrank(nonAllowlistedFiller);
    s_token.approve(address(s_pool), type(uint256).max);

    // Should succeed even though filler is not allowlisted
    s_pool.fastFill(FILL_REQUEST_ID, fillId, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);

    // Verify token balances
    assertEq(s_token.balanceOf(nonAllowlistedFiller), 1000 ether - SOURCE_AMOUNT);
    assertEq(s_token.balanceOf(RECEIVER), SOURCE_AMOUNT);
  }

  function test_FastFill_RevertWhen_AlreadyFilled() public {
    bytes32 fillId = s_pool.computeFillId(FILL_REQUEST_ID, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);

    // First fill
    s_pool.fastFill(FILL_REQUEST_ID, fillId, DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);

    // Attempt second fill
    vm.expectRevert(abi.encodeWithSelector(IFastTransferPool.AlreadyFilled.selector, FILL_REQUEST_ID));
    s_pool.fastFill(FILL_REQUEST_ID, fillId, DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);
  }

  function test_FastFill_RevertWhen_FillerNotAllowlisted() public {
    vm.stopPrank();

    address nonAllowlistedFiller = address(0x6);
    vm.startPrank(nonAllowlistedFiller);

    vm.expectRevert(
      abi.encodeWithSelector(
        FastTransferTokenPoolAbstract.FillerNotAllowlisted.selector, DEST_CHAIN_SELECTOR, nonAllowlistedFiller
      )
    );
    s_pool.fastFill(FILL_REQUEST_ID, bytes32(0x0), DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);
  }
}
