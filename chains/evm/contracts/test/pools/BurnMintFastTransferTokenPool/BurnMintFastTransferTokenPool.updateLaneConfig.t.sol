// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
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
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      evmToAnyMessageExtraArgsBytes: "",
      addFillers: addFillers,
      removeFillers: new address[](0)
    });

    vm.expectEmit();
    emit FastTransferTokenPoolAbstract.DestChainConfigUpdated(
      NEW_CHAIN_SELECTOR,
      NEW_FAST_FEE_BPS,
      NEW_FILL_AMOUNT_MAX,
      NEW_DESTINATION_POOL,
      addFillers,
      new address[](0),
      Internal.CHAIN_FAMILY_SELECTOR_EVM,
      200_000,
      false
    );

    s_pool.updateDestChainConfig(laneConfigArgs);

    FastTransferTokenPoolAbstract.DestChainConfigView memory config = s_pool.getDestChainConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.fastTransferBpsFee, NEW_FAST_FEE_BPS);
    assertFalse(config.fillerAllowlistEnabled);
    assertEq(config.destinationPool, NEW_DESTINATION_POOL);
    assertEq(config.maxFillAmountPerRequest, NEW_FILL_AMOUNT_MAX);

    // Check fillers are added
    assertTrue(s_pool.isFillerAllowListed(NEW_CHAIN_SELECTOR, addFillers[0]));
    assertTrue(s_pool.isFillerAllowListed(NEW_CHAIN_SELECTOR, addFillers[1]));
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
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      evmToAnyMessageExtraArgsBytes: "",
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
    assertFalse(s_pool.isFillerAllowListed(DEST_CHAIN_SELECTOR, s_filler));
    assertTrue(s_pool.isFillerAllowListed(DEST_CHAIN_SELECTOR, addFillers[0]));
  }

  function test_RevertWhen_InvalidFastFeeBps() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferBpsFee: 10_001, // > 10_000 (100%)
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      evmToAnyMessageExtraArgsBytes: "",
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
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      evmToAnyMessageExtraArgsBytes: "",
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

    s_pool.updateFillerAllowList(DEST_CHAIN_SELECTOR, addFillers, removeFillers);

    // Check changes
    assertFalse(s_pool.isFillerAllowListed(DEST_CHAIN_SELECTOR, s_filler));
    assertTrue(s_pool.isFillerAllowListed(DEST_CHAIN_SELECTOR, addFillers[0]));
    assertTrue(s_pool.isFillerAllowListed(DEST_CHAIN_SELECTOR, addFillers[1]));
  }

  function test_RevertWhen_UpdateFillerAllowList_NotOwner() public {
    address[] memory addFillers = new address[](1);
    addFillers[0] = makeAddr("newFiller");
    vm.stopPrank();
    vm.expectRevert();
    vm.prank(makeAddr("notOwner"));
    s_pool.updateFillerAllowList(DEST_CHAIN_SELECTOR, addFillers, new address[](0));
  }

  function test_UpdateDestChainConfig_ZeroFastFee() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferBpsFee: 0, // No fast fee
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      evmToAnyMessageExtraArgsBytes: "",
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
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      evmToAnyMessageExtraArgsBytes: "",
      addFillers: new address[](0),
      removeFillers: new address[](0)
    });

    s_pool.updateDestChainConfig(laneConfigArgs);

    FastTransferTokenPoolAbstract.DestChainConfigView memory config = s_pool.getDestChainConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.fastTransferBpsFee, 10_000);
  }

  function test_GetAllowListedFillers() public {
    // Add multiple fillers
    address[] memory addFillers = new address[](3);
    addFillers[0] = makeAddr("filler1");
    addFillers[1] = makeAddr("filler2");
    addFillers[2] = makeAddr("filler3");

    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferBpsFee: NEW_FAST_FEE_BPS,
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      evmToAnyMessageExtraArgsBytes: "",
      addFillers: addFillers,
      removeFillers: new address[](0)
    });

    s_pool.updateDestChainConfig(laneConfigArgs);

    // Get all allowlisted fillers
    address[] memory allowlistedFillers = s_pool.getAllowListedFillers(NEW_CHAIN_SELECTOR);

    // Verify all fillers are returned
    assertEq(allowlistedFillers.length, 3);

    // Verify each filler is in the list (order may vary due to EnumerableSet)
    bool foundFiller1 = false;
    bool foundFiller2 = false;
    bool foundFiller3 = false;

    for (uint256 i = 0; i < allowlistedFillers.length; i++) {
      if (allowlistedFillers[i] == addFillers[0]) foundFiller1 = true;
      if (allowlistedFillers[i] == addFillers[1]) foundFiller2 = true;
      if (allowlistedFillers[i] == addFillers[2]) foundFiller3 = true;
    }

    assertTrue(foundFiller1);
    assertTrue(foundFiller2);
    assertTrue(foundFiller3);
  }

  function test_GetAllowListedFillers_EmptyList() public view {
    // Test with no fillers added
    address[] memory allowlistedFillers = s_pool.getAllowListedFillers(NEW_CHAIN_SELECTOR);
    assertEq(allowlistedFillers.length, 0);
  }

  function test_GetAllowListedFillers_AfterRemoval() public {
    // First add fillers
    address[] memory addFillers = new address[](2);
    addFillers[0] = makeAddr("filler1");
    addFillers[1] = makeAddr("filler2");

    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferBpsFee: NEW_FAST_FEE_BPS,
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      evmToAnyMessageExtraArgsBytes: "",
      addFillers: addFillers,
      removeFillers: new address[](0)
    });

    s_pool.updateDestChainConfig(laneConfigArgs);

    // Then remove one filler
    address[] memory removeFillers = new address[](1);
    removeFillers[0] = addFillers[0];

    s_pool.updateFillerAllowList(NEW_CHAIN_SELECTOR, new address[](0), removeFillers);

    // Verify only one filler remains
    address[] memory allowlistedFillers = s_pool.getAllowListedFillers(NEW_CHAIN_SELECTOR);
    assertEq(allowlistedFillers.length, 1);
    assertEq(allowlistedFillers[0], addFillers[1]);
  }
}
