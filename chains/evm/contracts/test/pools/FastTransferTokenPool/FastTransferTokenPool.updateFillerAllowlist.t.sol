// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract FastTransferTokenPool_updateFillerAllowlist is FastTransferTokenPoolSetup {
  function test_updateFillerAllowList_Success() public {
    address[] memory addFillers = new address[](2);
    addFillers[0] = makeAddr("newFiller1");
    addFillers[1] = makeAddr("newFiller2");

    address[] memory removeFillers = new address[](1);
    removeFillers[0] = s_filler;

    vm.expectEmit();
    emit FastTransferTokenPoolAbstract.FillerAllowListUpdated(addFillers, removeFillers);

    s_pool.updateFillerAllowList(addFillers, removeFillers);

    // Check changes
    assertFalse(s_pool.isAllowedFiller(s_filler));
    assertTrue(s_pool.isAllowedFiller(addFillers[0]));
    assertTrue(s_pool.isAllowedFiller(addFillers[1]));
  }

  function test_updateFillerAllowList_RevertWhen_NotOwner() public {
    vm.stopPrank();

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_pool.updateFillerAllowList(new address[](0), new address[](0));
  }

  function test_updateFillerAllowList_GetAllowListedFillers() public {
    uint256 initialFillerCount = s_pool.getAllowedFillers().length;

    // Add multiple fillers
    address[] memory addFillers = new address[](3);
    addFillers[0] = makeAddr("filler1");
    addFillers[1] = makeAddr("filler2");
    addFillers[2] = makeAddr("filler3");

    s_pool.updateFillerAllowList(addFillers, new address[](0));

    // Get all allowlisted fillers
    address[] memory allowlistedFillers = s_pool.getAllowedFillers();

    // Verify all fillers are returned
    assertEq(allowlistedFillers.length, initialFillerCount + 3);

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
    uint256 initialFillerCount = s_pool.getAllowedFillers().length;

    // First add fillers
    address[] memory addFillers = new address[](2);
    addFillers[0] = makeAddr("filler1");
    addFillers[1] = makeAddr("filler2");

    vm.expectEmit();
    emit FastTransferTokenPoolAbstract.FillerAllowListUpdated(addFillers, new address[](0));

    s_pool.updateFillerAllowList(addFillers, new address[](0));

    // Verify both fillers are added
    address[] memory allowlistedFillers = s_pool.getAllowedFillers();
    assertEq(allowlistedFillers.length, initialFillerCount + 2);
    assertTrue(s_pool.isAllowedFiller(addFillers[0]));
    assertTrue(s_pool.isAllowedFiller(addFillers[1]));

    // Then remove one filler
    address[] memory removeFillers = new address[](1);
    removeFillers[0] = addFillers[0];

    vm.expectEmit();
    emit FastTransferTokenPoolAbstract.FillerAllowListUpdated(new address[](0), removeFillers);

    s_pool.updateFillerAllowList(new address[](0), removeFillers);

    // Verify only one filler remains
    allowlistedFillers = s_pool.getAllowedFillers();
    assertEq(allowlistedFillers.length, initialFillerCount + 1);
    assertFalse(s_pool.isAllowedFiller(addFillers[0]));
  }
}
