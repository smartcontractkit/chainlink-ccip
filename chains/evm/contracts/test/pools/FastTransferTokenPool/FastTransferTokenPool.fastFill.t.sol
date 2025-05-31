// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

contract FastTransferTokenPool_fastFill_Test is FastTransferTokenPoolSetup {
  bytes32 public fillRequestId;
  address public receiver;

  function setUp() public override {
    super.setUp();
    vm.stopPrank();
    deal(address(s_token), s_filler, 1000 ether);
    vm.prank(s_filler);
    s_token.approve(address(s_tokenPool), type(uint256).max);
    receiver = address(0x5);
  }

  function test_FastFill() public {
    uint256 srcAmount = 100 ether;
    uint8 sourceDecimals = 18;

    vm.prank(s_filler);
    s_tokenPool.fastFill(fillRequestId, DEST_CHAIN_SELECTOR, srcAmount, sourceDecimals, receiver);

    // Verify token balances
    assertEq(s_token.balanceOf(s_filler), 1000 ether - srcAmount);
    assertEq(s_token.balanceOf(receiver), srcAmount);
  }

  function test_FastFill_RevertWhen_AlreadyFilled() public {
    uint256 srcAmount = 100 ether;
    uint8 sourceDecimals = 18;

    // First fill
    vm.prank(s_filler);
    s_tokenPool.fastFill(fillRequestId, DEST_CHAIN_SELECTOR, srcAmount, sourceDecimals, receiver);

    // Attempt second fill
    vm.expectRevert(abi.encodeWithSelector(IFastTransferPool.AlreadyFilled.selector, fillRequestId));
    vm.prank(s_filler);
    s_tokenPool.fastFill(fillRequestId, DEST_CHAIN_SELECTOR, srcAmount, sourceDecimals, receiver);
  }

  function test_FastFill_RevertWhen_FillerNotWhitelisted() public {
    uint256 srcAmount = 100 ether;
    uint8 sourceDecimals = 18;
    address nonWhitelistedFiller = address(0x6);
    deal(address(s_token), nonWhitelistedFiller, 1000 ether);
    vm.startPrank(nonWhitelistedFiller);
    s_token.approve(address(s_tokenPool), type(uint256).max);
    vm.expectRevert(
      abi.encodeWithSelector(
        FastTransferTokenPoolAbstract.FillerNotAllowlisted.selector, DEST_CHAIN_SELECTOR, nonWhitelistedFiller
      )
    );
    s_tokenPool.fastFill(fillRequestId, DEST_CHAIN_SELECTOR, srcAmount, sourceDecimals, receiver);
    vm.stopPrank();
  }

  function test_FastFill_WhitelistDisabled() public {
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
      evmToAnyMessageExtraArgsBytes: ""
    });
    s_tokenPool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    uint256 srcAmount = 100 ether;
    uint8 sourceDecimals = 18;
    address nonWhitelistedFiller = address(0x6);

    // Mint tokens to non-whitelisted filler
    deal(nonWhitelistedFiller, 1000 ether);
    deal(address(s_token), nonWhitelistedFiller, 1000 ether);
    vm.startPrank(nonWhitelistedFiller);
    s_token.approve(address(s_tokenPool), type(uint256).max);

    // Should succeed even though filler is not whitelisted
    s_tokenPool.fastFill(fillRequestId, SOURCE_CHAIN_SELECTOR, srcAmount, sourceDecimals, receiver);
    vm.stopPrank();

    // Verify token balances
    assertEq(s_token.balanceOf(nonWhitelistedFiller), 1000 ether - srcAmount);
    assertEq(s_token.balanceOf(receiver), srcAmount);
  }
}
