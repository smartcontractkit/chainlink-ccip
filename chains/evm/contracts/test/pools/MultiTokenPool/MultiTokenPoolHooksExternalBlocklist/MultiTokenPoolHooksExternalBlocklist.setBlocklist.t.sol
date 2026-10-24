// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAddressBlocklist} from "../../../../interfaces/IAddressBlocklist.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {
  MultiTokenPoolHooksExternalBlocklist
} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooksExternalBlocklist.sol";
import {MultiTokenPoolHooksExternalBlocklistSetup} from "./MultiTokenPoolHooksExternalBlocklistSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract MultiTokenPoolHooksExternalBlocklist_setBlocklist is MultiTokenPoolHooksExternalBlocklistSetup {
  IAddressBlocklist internal s_newBlocklist = IAddressBlocklist(makeAddr("newBlocklist"));

  function setUp() public virtual override {
    super.setUp();
    _deployBlocklist(s_newBlocklist);
  }

  function test_setBlocklist() public {
    vm.expectEmit();
    emit MultiTokenPoolHooksExternalBlocklist.BlocklistSet(s_blocklist, s_newBlocklist);
    s_hooks.setBlocklist(s_newBlocklist);

    assertEq(address(s_newBlocklist), address(s_hooks.getBlocklist()));

    // Only the new blocklist is called from now on.
    vm.expectCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), 0);
    vm.expectCall(address(s_newBlocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), 1);
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_setBlocklist_ToZero() public {
    vm.expectEmit();
    emit MultiTokenPoolHooksExternalBlocklist.BlocklistSet(s_blocklist, IAddressBlocklist(address(0)));
    s_hooks.setBlocklist(IAddressBlocklist(address(0)));

    assertEq(address(0), address(s_hooks.getBlocklist()));
  }

  function test_setBlocklist_SameValue() public {
    vm.recordLogs();
    s_hooks.setBlocklist(s_blocklist);

    assertEq(0, vm.getRecordedLogs().length);
    assertEq(address(s_blocklist), address(s_hooks.getBlocklist()));
  }

  // Reverts

  function test_setBlocklist_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_hooks.setBlocklist(s_newBlocklist);
  }

  function test_setBlocklist_RevertWhen_InvalidBlocklist() public {
    IAddressBlocklist blocklistWithoutCode = IAddressBlocklist(makeAddr("blocklistWithoutCode"));

    vm.expectRevert(
      abi.encodeWithSelector(MultiTokenPoolHooksExternalBlocklist.InvalidBlocklist.selector, blocklistWithoutCode)
    );
    s_hooks.setBlocklist(blocklistWithoutCode);
  }
}
