// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

contract FastTransferTokenPool_updateFillerAllowlist is FastTransferTokenPoolSetup {
  uint64 internal constant NEW_CHAIN_SELECTOR = 12345;

  function test_updateFillerAllowList_Success() public {
    address[] memory addFillers = new address[](2);
    addFillers[0] = makeAddr("newFiller1");
    addFillers[1] = makeAddr("newFiller2");

    address[] memory removeFillers = new address[](1);
    removeFillers[0] = s_filler;

    vm.expectEmit();
    emit FastTransferTokenPoolAbstract.FillerAllowListUpdated(DEST_CHAIN_SELECTOR, addFillers, removeFillers);

    s_pool.updateFillerAllowList(DEST_CHAIN_SELECTOR, addFillers, removeFillers);

    // Check changes
    assertFalse(s_pool.isAllowedFiller(DEST_CHAIN_SELECTOR, s_filler));
    assertTrue(s_pool.isAllowedFiller(DEST_CHAIN_SELECTOR, addFillers[0]));
    assertTrue(s_pool.isAllowedFiller(DEST_CHAIN_SELECTOR, addFillers[1]));
  }

  function test_updateFillerAllowList_RevertWhen_NotOwner() public {
    vm.stopPrank();

    vm.expectRevert(); // TODO revert message
    s_pool.updateFillerAllowList(DEST_CHAIN_SELECTOR, new address[](0), new address[](0));
  }

  function test_updateFillerAllowList_GetAllowListedFillers() public {
    // Add multiple fillers
    address[] memory addFillers = new address[](3);
    addFillers[0] = makeAddr("filler1");
    addFillers[1] = makeAddr("filler2");
    addFillers[2] = makeAddr("filler3");

    s_pool.updateFillerAllowList(NEW_CHAIN_SELECTOR, addFillers, new address[](0));

    // Get all allowlisted fillers
    address[] memory allowlistedFillers = s_pool.getAllowedFillers(NEW_CHAIN_SELECTOR);

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

  function test_getAllowedFillers_AfterRemoval() public {
    // First add fillers
    address[] memory addFillers = new address[](2);
    addFillers[0] = makeAddr("filler1");
    addFillers[1] = makeAddr("filler2");

    s_pool.updateFillerAllowList(NEW_CHAIN_SELECTOR, addFillers, new address[](0));

    // Then remove one filler
    address[] memory removeFillers = new address[](1);
    removeFillers[0] = addFillers[0];

    s_pool.updateFillerAllowList(NEW_CHAIN_SELECTOR, new address[](0), removeFillers);

    // Verify only one filler remains
    address[] memory allowlistedFillers = s_pool.getAllowedFillers(NEW_CHAIN_SELECTOR);
    assertEq(allowlistedFillers.length, 1);
    assertEq(allowlistedFillers[0], addFillers[1]);
  }
}
