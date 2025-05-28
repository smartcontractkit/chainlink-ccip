// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.10;

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {BurnMintFastTransferTokenPoolSetup} from "./BurnMintFastTransferTokenPoolSetup.t.sol";

contract BurnMintFastTransferTokenPool_updateLaneConfig is BurnMintFastTransferTokenPoolSetup {
  uint64 internal constant NEW_CHAIN_SELECTOR = 12345;
  address internal constant NEW_DESTINATION_POOL = address(0x5678);
  uint16 internal constant NEW_FAST_FEE_BPS = 200; // 2%
  uint256 internal constant NEW_FILL_AMOUNT_MAX = 2000 ether;

  function test_UpdateLaneConfig() public {
    address[] memory addFillers = new address[](2);
    addFillers[0] = makeAddr("newFiller1");
    addFillers[1] = makeAddr("newFiller2");

    FastTransferTokenPoolAbstract.LaneConfigArgs memory laneConfigArgs = FastTransferTokenPoolAbstract.LaneConfigArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      bpsFastFee: NEW_FAST_FEE_BPS,
      enabled: true,
      fillerAllowlistEnabled: false,
      destinationPool: NEW_DESTINATION_POOL,
      fillAmountMaxPerRequest: NEW_FILL_AMOUNT_MAX,
      addFillers: addFillers,
      removeFillers: new address[](0)
    });

    vm.expectEmit();
    emit FastTransferTokenPoolAbstract.LaneUpdated(
      NEW_CHAIN_SELECTOR,
      NEW_FAST_FEE_BPS,
      true,
      NEW_FILL_AMOUNT_MAX,
      NEW_DESTINATION_POOL,
      addFillers,
      new address[](0)
    );

    s_pool.updateLaneConfig(laneConfigArgs);

    FastTransferTokenPoolAbstract.LaneConfigView memory config = s_pool.getLaneConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.bpsFastFee, NEW_FAST_FEE_BPS);
    assertTrue(config.enabled);
    assertFalse(config.fillerAllowlistEnabled);
    assertEq(config.destinationPool, NEW_DESTINATION_POOL);
    assertEq(config.fillAmountMaxPerRequest, NEW_FILL_AMOUNT_MAX);

    // Check fillers are added
    assertTrue(s_pool.isfillerAllowListed(NEW_CHAIN_SELECTOR, addFillers[0]));
    assertTrue(s_pool.isfillerAllowListed(NEW_CHAIN_SELECTOR, addFillers[1]));
  }

  function test_UpdateLaneConfig_ModifyExisting() public {
    // Modify existing lane config
    address[] memory removeFillers = new address[](1);
    removeFillers[0] = s_filler;

    address[] memory addFillers = new address[](1);
    addFillers[0] = makeAddr("replacementFiller");

    FastTransferTokenPoolAbstract.LaneConfigArgs memory laneConfigArgs = FastTransferTokenPoolAbstract.LaneConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      bpsFastFee: NEW_FAST_FEE_BPS,
      enabled: false, // disable
      fillerAllowlistEnabled: false, // disable whitelist
      destinationPool: NEW_DESTINATION_POOL,
      fillAmountMaxPerRequest: NEW_FILL_AMOUNT_MAX,
      addFillers: addFillers,
      removeFillers: removeFillers
    });

    s_pool.updateLaneConfig(laneConfigArgs);

    FastTransferTokenPoolAbstract.LaneConfigView memory config = s_pool.getLaneConfig(DEST_CHAIN_SELECTOR);
    assertEq(config.bpsFastFee, NEW_FAST_FEE_BPS);
    assertFalse(config.enabled);
    assertFalse(config.fillerAllowlistEnabled);
    assertEq(config.destinationPool, NEW_DESTINATION_POOL);
    assertEq(config.fillAmountMaxPerRequest, NEW_FILL_AMOUNT_MAX);

    // Check filler changes
    assertFalse(s_pool.isfillerAllowListed(DEST_CHAIN_SELECTOR, s_filler));
    assertTrue(s_pool.isfillerAllowListed(DEST_CHAIN_SELECTOR, addFillers[0]));
  }

  function test_RevertWhen_InvalidFastFeeBps() public {
    FastTransferTokenPoolAbstract.LaneConfigArgs memory laneConfigArgs = FastTransferTokenPoolAbstract.LaneConfigArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      bpsFastFee: 10_001, // > 10_000 (100%)
      enabled: true,
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      fillAmountMaxPerRequest: NEW_FILL_AMOUNT_MAX,
      addFillers: new address[](0),
      removeFillers: new address[](0)
    });

    vm.expectRevert(FastTransferTokenPoolAbstract.InvalidLaneConfig.selector);
    s_pool.updateLaneConfig(laneConfigArgs);
  }

  function test_RevertWhen_NotOwners() public {
    FastTransferTokenPoolAbstract.LaneConfigArgs memory laneConfigArgs = FastTransferTokenPoolAbstract.LaneConfigArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      bpsFastFee: NEW_FAST_FEE_BPS,
      enabled: true,
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      fillAmountMaxPerRequest: NEW_FILL_AMOUNT_MAX,
      addFillers: new address[](0),
      removeFillers: new address[](0)
    });
    vm.stopPrank();
    vm.expectRevert();
    vm.prank(makeAddr("notOwner"));
    s_pool.updateLaneConfig(laneConfigArgs);
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

  function test_UpdateLaneConfig_ZeroFastFee() public {
    FastTransferTokenPoolAbstract.LaneConfigArgs memory laneConfigArgs = FastTransferTokenPoolAbstract.LaneConfigArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      bpsFastFee: 0, // No fast fee
      enabled: true,
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      fillAmountMaxPerRequest: NEW_FILL_AMOUNT_MAX,
      addFillers: new address[](0),
      removeFillers: new address[](0)
    });

    s_pool.updateLaneConfig(laneConfigArgs);

    FastTransferTokenPoolAbstract.LaneConfigView memory config = s_pool.getLaneConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.bpsFastFee, 0);
  }

  function test_UpdateLaneConfig_MaxFastFee() public {
    FastTransferTokenPoolAbstract.LaneConfigArgs memory laneConfigArgs = FastTransferTokenPoolAbstract.LaneConfigArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      bpsFastFee: 10_000, // 100% fee
      enabled: true,
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      fillAmountMaxPerRequest: NEW_FILL_AMOUNT_MAX,
      addFillers: new address[](0),
      removeFillers: new address[](0)
    });

    s_pool.updateLaneConfig(laneConfigArgs);

    FastTransferTokenPoolAbstract.LaneConfigView memory config = s_pool.getLaneConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.bpsFastFee, 10_000);
  }
}
