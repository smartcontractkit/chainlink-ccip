// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

contract FastTransferTokenPool_fastFill_Test is FastTransferTokenPoolSetup {
  bytes32 public constant FILL_REQUEST_ID = bytes32("fillRequestId");
  address public constant RECEIVER = address(0x5);

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

  function test_FastFill_WhitelistDisabled() public {
    vm.stopPrank();
    vm.prank(OWNER);
    // Disable whitelist
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastTransferBpsFee: 100,
      fillerAllowlistEnabled: false,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: 1000 ether,
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    address nonWhitelistedFiller = address(0x6);

    bytes32 fillId = s_pool.computeFillId(FILL_REQUEST_ID, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);

    // Mint tokens to non-whitelisted filler
    deal(nonWhitelistedFiller, 1000 ether);
    deal(address(s_token), nonWhitelistedFiller, 1000 ether);
    vm.startPrank(nonWhitelistedFiller);
    s_token.approve(address(s_pool), type(uint256).max);

    // Should succeed even though filler is not whitelisted
    s_pool.fastFill(FILL_REQUEST_ID, fillId, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);

    // Verify token balances
    assertEq(s_token.balanceOf(nonWhitelistedFiller), 1000 ether - SOURCE_AMOUNT);
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

  function test_FastFill_RevertWhen_FillerNotWhitelisted() public {
    vm.stopPrank();

    address nonWhitelistedFiller = address(0x6);
    vm.startPrank(nonWhitelistedFiller);

    vm.expectRevert(
      abi.encodeWithSelector(
        FastTransferTokenPoolAbstract.FillerNotAllowlisted.selector, DEST_CHAIN_SELECTOR, nonWhitelistedFiller
      )
    );
    s_pool.fastFill(FILL_REQUEST_ID, bytes32(0x0), DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, RECEIVER);
  }
}
