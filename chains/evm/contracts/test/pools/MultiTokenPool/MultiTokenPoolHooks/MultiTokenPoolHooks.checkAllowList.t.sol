// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MultiTokenPoolHooks} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooks.sol";
import {MultiTokenPoolHooksSetup} from "./MultiTokenPoolHooksSetup.t.sol";

contract MultiTokenPoolHooks_checkAllowList is MultiTokenPoolHooksSetup {
  function test_checkAllowList() public view {
    s_hooksWithAllowList.checkAllowList(s_sender);
  }

  function test_checkAllowList_AllowListDisabled() public view {
    s_hooks.checkAllowList(STRANGER);
    s_hooks.checkAllowList(address(0));
  }

  // Reverts

  function test_checkAllowList_RevertWhen_SenderNotAllowed() public {
    vm.expectRevert(abi.encodeWithSelector(MultiTokenPoolHooks.SenderNotAllowed.selector, STRANGER));
    s_hooksWithAllowList.checkAllowList(STRANGER);
  }
}
