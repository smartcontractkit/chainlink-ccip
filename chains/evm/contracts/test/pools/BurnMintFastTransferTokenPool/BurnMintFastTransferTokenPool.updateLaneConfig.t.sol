// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {BurnMintFastTransferTokenPoolSetup} from "./BurnMintFastTransferTokenPoolSetup.t.sol";

contract BurnMintFastTransferTokenPool_updateDestChainConfig is BurnMintFastTransferTokenPoolSetup {
  uint64 internal constant NEW_CHAIN_SELECTOR = 12345;
  bytes internal constant NEW_DESTINATION_POOL = abi.encode(address(0x5678));
  uint16 internal constant NEW_FAST_FEE_BPS = 200; // 2%
  uint256 internal constant NEW_FILL_AMOUNT_MAX = 2000 ether;

  function test_UpdateDestChainConfig() public {
    address[] memory addFillers = new address[](2);
    addFillers[0] = makeAddr("newFiller1");
    addFillers[1] = makeAddr("newFiller2");

    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferBpsFee: NEW_FAST_FEE_BPS,
      fillerAllowlistEnabled: false,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      addFillers: addFillers,
      removeFillers: new address[](0)
    });

    vm.expectEmit();
    emit FastTransferTokenPoolAbstract.LaneUpdated(
      NEW_CHAIN_SELECTOR, NEW_FAST_FEE_BPS, NEW_FILL_AMOUNT_MAX, NEW_DESTINATION_POOL, addFillers, new address[](0)
    );

    s_pool.updateDestChainConfig(laneConfigArgs);

    FastTransferTokenPoolAbstract.DestChainConfigView memory config = s_pool.getDestChainConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.fastTransferBpsFee, NEW_FAST_FEE_BPS);
    assertFalse(config.fillerAllowlistEnabled);
    assertEq(config.destinationPool, NEW_DESTINATION_POOL);
    assertEq(config.maxFillAmountPerRequest, NEW_FILL_AMOUNT_MAX);

    // Check fillers are added
    assertTrue(s_pool.isfillerAllowListed(NEW_CHAIN_SELECTOR, addFillers[0]));
    assertTrue(s_pool.isfillerAllowListed(NEW_CHAIN_SELECTOR, addFillers[1]));
  }

  function test_UpdateDestChainConfig_ModifyExisting() public {
    // Modify existing lane config
    address[] memory removeFillers = new address[](1);
    removeFillers[0] = s_filler;

    address[] memory addFillers = new address[](1);
    addFillers[0] = makeAddr("replacementFiller");

    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastTransferBpsFee: NEW_FAST_FEE_BPS,
      fillerAllowlistEnabled: false, // disable whitelist
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      addFillers: addFillers,
      removeFillers: removeFillers
    });

    s_pool.updateDestChainConfig(laneConfigArgs);

    FastTransferTokenPoolAbstract.DestChainConfigView memory config = s_pool.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(config.fastTransferBpsFee, NEW_FAST_FEE_BPS);
    assertFalse(config.fillerAllowlistEnabled);
    assertEq(config.destinationPool, NEW_DESTINATION_POOL);
    assertEq(config.maxFillAmountPerRequest, NEW_FILL_AMOUNT_MAX);

    // Check filler changes
    assertFalse(s_pool.isfillerAllowListed(DEST_CHAIN_SELECTOR, s_filler));
    assertTrue(s_pool.isfillerAllowListed(DEST_CHAIN_SELECTOR, addFillers[0]));
  }

  function test_RevertWhen_InvalidFastFeeBps() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferBpsFee: 10_001, // > 10_000 (100%)
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      addFillers: new address[](0),
      removeFillers: new address[](0)
    });

    vm.expectRevert(FastTransferTokenPoolAbstract.InvalidDestChainConfig.selector);
    s_pool.updateDestChainConfig(laneConfigArgs);
  }

  function test_RevertWhen_NotOwners() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferBpsFee: NEW_FAST_FEE_BPS,
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      addFillers: new address[](0),
      removeFillers: new address[](0)
    });
    vm.stopPrank();
    vm.expectRevert();
    vm.prank(makeAddr("notOwner"));
    s_pool.updateDestChainConfig(laneConfigArgs);
  }

  function test_UpdateFillerAllowList_Success() public {
    address[] memory addFillers = new address[](2);
    addFillers[0] = makeAddr("newFiller1");
    addFillers[1] = makeAddr("newFiller2");

    address[] memory removeFillers = new address[](1);
    removeFillers[0] = s_filler;

    vm.expectEmit();
    emit FastTransferTokenPoolAbstract.FillerAllowListUpdated(DEST_CHAIN_SELECTOR, addFillers, removeFillers);

    s_pool.updatefillerAllowList(DEST_CHAIN_SELECTOR, addFillers, removeFillers);

    // Check changes
    assertFalse(s_pool.isfillerAllowListed(DEST_CHAIN_SELECTOR, s_filler));
    assertTrue(s_pool.isfillerAllowListed(DEST_CHAIN_SELECTOR, addFillers[0]));
    assertTrue(s_pool.isfillerAllowListed(DEST_CHAIN_SELECTOR, addFillers[1]));
  }

  function test_RevertWhen_UpdateFillerAllowList_NotOwner() public {
    address[] memory addFillers = new address[](1);
    addFillers[0] = makeAddr("newFiller");
    vm.stopPrank();
    vm.expectRevert();
    vm.prank(makeAddr("notOwner"));
    s_pool.updatefillerAllowList(DEST_CHAIN_SELECTOR, addFillers, new address[](0));
  }

  function test_UpdateDestChainConfig_ZeroFastFee() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferBpsFee: 0, // No fast fee
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      addFillers: new address[](0),
      removeFillers: new address[](0)
    });

    s_pool.updateDestChainConfig(laneConfigArgs);

    FastTransferTokenPoolAbstract.DestChainConfigView memory config = s_pool.getDestChainConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.fastTransferBpsFee, 0);
  }

  function test_UpdateDestChainConfig_MaxFastFee() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferBpsFee: 10_000, // 100% fee
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      addFillers: new address[](0),
      removeFillers: new address[](0)
    });

    s_pool.updateDestChainConfig(laneConfigArgs);

    FastTransferTokenPoolAbstract.DestChainConfigView memory config = s_pool.getDestChainConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.fastTransferBpsFee, 10_000);
  }
}
