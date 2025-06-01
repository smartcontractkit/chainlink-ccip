// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

contract FastTransferTokenPool_fastFill_Test is FastTransferTokenPoolSetup {
  bytes32 public fillRequestId;
  address public receiver = address(0x5);

  function setUp() public override {
    super.setUp();
    vm.stopPrank();
    deal(address(s_token), s_filler, 1000 ether);
    vm.startPrank(s_filler);
    s_token.approve(address(s_pool), type(uint256).max);
  }

  function test_FastFill() public {
    bytes32 fillId = s_pool.computeFillId(fillRequestId, SOURCE_AMOUNT, SOURCE_DECIMALS, receiver);

    s_pool.fastFill(fillRequestId, fillId, DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, receiver);

    // Verify token balances
    assertEq(s_token.balanceOf(s_filler), 1000 ether - SOURCE_AMOUNT);
    assertEq(s_token.balanceOf(receiver), SOURCE_AMOUNT);
  }

  function test_FastFill_RevertWhen_AlreadyFilled() public {
    bytes32 fillId = s_pool.computeFillId(fillRequestId, SOURCE_AMOUNT, SOURCE_DECIMALS, receiver);

    // First fill
    s_pool.fastFill(fillRequestId, fillId, DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, receiver);

    // Attempt second fill
    vm.expectRevert(abi.encodeWithSelector(IFastTransferPool.AlreadyFilled.selector, fillRequestId));
    s_pool.fastFill(fillRequestId, fillId, DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, receiver);
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
    s_pool.fastFill(fillRequestId, bytes32(0x0), DEST_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, receiver);
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

    bytes32 fillId = s_pool.computeFillId(fillRequestId, SOURCE_AMOUNT, SOURCE_DECIMALS, receiver);

    // Mint tokens to non-whitelisted filler
    deal(nonWhitelistedFiller, 1000 ether);
    deal(address(s_token), nonWhitelistedFiller, 1000 ether);
    vm.startPrank(nonWhitelistedFiller);
    s_token.approve(address(s_pool), type(uint256).max);

    // Should succeed even though filler is not whitelisted
    s_pool.fastFill(fillRequestId, fillId, SOURCE_CHAIN_SELECTOR, SOURCE_AMOUNT, SOURCE_DECIMALS, receiver);

    // Verify token balances
    assertEq(s_token.balanceOf(nonWhitelistedFiller), 1000 ether - SOURCE_AMOUNT);
    assertEq(s_token.balanceOf(receiver), SOURCE_AMOUNT);
  }
}
