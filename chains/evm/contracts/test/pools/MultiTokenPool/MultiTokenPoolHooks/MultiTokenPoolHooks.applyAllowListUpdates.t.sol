// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MultiTokenPoolHooks} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooks.sol";
import {MultiTokenPoolHooksSetup} from "./MultiTokenPoolHooksSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract MultiTokenPoolHooks_applyAllowListUpdates is MultiTokenPoolHooksSetup {
  address internal s_newSender1 = makeAddr("newSender1");
  address internal s_newSender2 = makeAddr("newSender2");

  function setUp() public virtual override {
    super.setUp();
    s_hooks = _deployHooksWithAllowList();
  }

  function test_applyAllowListUpdates() public {
    address[] memory adds = new address[](2);
    adds[0] = s_newSender1;
    adds[1] = s_newSender2;

    vm.expectEmit();
    emit MultiTokenPoolHooks.AllowListAdd(s_newSender1);
    vm.expectEmit();
    emit MultiTokenPoolHooks.AllowListAdd(s_newSender2);
    s_hooks.applyAllowListUpdates(new address[](0), adds);

    address[] memory expected = new address[](3);
    expected[0] = s_sender;
    expected[1] = s_newSender1;
    expected[2] = s_newSender2;
    assertEq(expected, s_hooks.getAllowList());
  }

  function test_applyAllowListUpdates_RemovesBeforeAdds() public {
    address[] memory removes = new address[](2);
    removes[0] = s_sender;
    removes[1] = s_newSender1;
    address[] memory adds = new address[](2);
    adds[0] = s_newSender1;
    adds[1] = s_sender;

    vm.expectEmit();
    emit MultiTokenPoolHooks.AllowListRemove(s_sender);
    vm.expectEmit();
    emit MultiTokenPoolHooks.AllowListAdd(s_newSender1);
    vm.expectEmit();
    emit MultiTokenPoolHooks.AllowListAdd(s_sender);
    s_hooks.applyAllowListUpdates(removes, adds);

    address[] memory expected = new address[](2);
    expected[0] = s_newSender1;
    expected[1] = s_sender;
    assertEq(expected, s_hooks.getAllowList());
  }

  function test_applyAllowListUpdates_SkipsZeroDuplicatesAndMissing() public {
    address[] memory removes = new address[](1);
    removes[0] = s_newSender2;
    address[] memory adds = new address[](4);
    adds[0] = address(0);
    adds[1] = s_sender;
    adds[2] = s_newSender1;
    adds[3] = s_newSender1;

    vm.recordLogs();
    s_hooks.applyAllowListUpdates(removes, adds);

    // The zero address, the existing sender and the duplicate are skipped, so only one address is added.
    assertEq(1, vm.getRecordedLogs().length);
    address[] memory expected = new address[](2);
    expected[0] = s_sender;
    expected[1] = s_newSender1;
    assertEq(expected, s_hooks.getAllowList());
  }

  function test_applyAllowListUpdates_EmptyLists() public {
    vm.recordLogs();
    s_hooks.applyAllowListUpdates(new address[](0), new address[](0));

    assertEq(0, vm.getRecordedLogs().length);
    assertEq(_arrayOf(s_sender), s_hooks.getAllowList());
  }

  function test_applyAllowListUpdates_RemoveAll() public {
    vm.expectEmit();
    emit MultiTokenPoolHooks.AllowListRemove(s_sender);
    s_hooks.applyAllowListUpdates(_arrayOf(s_sender), new address[](0));

    assertEq(0, s_hooks.getAllowList().length);
    // Removing every sender does not turn the allowlist off, so nobody can send until a sender is added again.
    assertTrue(s_hooks.getAllowListEnabled());
  }

  // Reverts

  function test_applyAllowListUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_hooks.applyAllowListUpdates(new address[](0), _arrayOf(s_newSender1));
  }

  function test_applyAllowListUpdates_RevertWhen_AllowListNotEnabled() public {
    MultiTokenPoolHooks hooks = new MultiTokenPoolHooks(new address[](0), s_policyEngine, _arrayOf(OWNER));

    vm.expectRevert(MultiTokenPoolHooks.AllowListNotEnabled.selector);
    hooks.applyAllowListUpdates(new address[](0), _arrayOf(s_newSender1));
  }
}
