// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolHelper} from "../../helpers/TokenPoolHelper.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

import {TokenPoolWithAllowListSetup} from "./TokenPoolWithAllowListSetup.t.sol";

contract TokenPoolWithAllowList_applyAllowListUpdates is TokenPoolWithAllowListSetup {
  function test_SetAllowList() public {
    address[] memory newAddresses = new address[](2);
    newAddresses[0] = address(1);
    newAddresses[1] = address(2);

    for (uint256 i = 0; i < 2; ++i) {
      vm.expectEmit();
      emit AdvancedPoolHooks.AllowListAdd(newAddresses[i]);
    }

    s_advancedPoolHooks.applyAllowListUpdates(new address[](0), newAddresses);
    address[] memory setAddresses = s_advancedPoolHooks.getAllowList();

    assertEq(s_allowedSenders[0], setAddresses[0]);
    assertEq(s_allowedSenders[1], setAddresses[1]);
    assertEq(address(1), setAddresses[2]);
    assertEq(address(2), setAddresses[3]);

    // address(2) exists noop, add address(3), remove address(1)
    newAddresses = new address[](2);
    newAddresses[0] = address(2);
    newAddresses[1] = address(3);

    address[] memory removeAddresses = new address[](1);
    removeAddresses[0] = address(1);

    vm.expectEmit();
    emit AdvancedPoolHooks.AllowListRemove(address(1));

    vm.expectEmit();
    emit AdvancedPoolHooks.AllowListAdd(address(3));

    s_advancedPoolHooks.applyAllowListUpdates(removeAddresses, newAddresses);
    setAddresses = s_advancedPoolHooks.getAllowList();

    assertEq(s_allowedSenders[0], setAddresses[0]);
    assertEq(s_allowedSenders[1], setAddresses[1]);
    assertEq(address(2), setAddresses[2]);
    assertEq(address(3), setAddresses[3]);

    // remove all from allowlist
    for (uint256 i = 0; i < setAddresses.length; ++i) {
      vm.expectEmit();
      emit AdvancedPoolHooks.AllowListRemove(setAddresses[i]);
    }

    s_advancedPoolHooks.applyAllowListUpdates(setAddresses, new address[](0));
    setAddresses = s_advancedPoolHooks.getAllowList();

    assertEq(0, setAddresses.length);
  }

  function test_SetAllowListSkipsZero() public {
    uint256 setAddressesLength = s_advancedPoolHooks.getAllowList().length;

    address[] memory newAddresses = new address[](1);
    newAddresses[0] = address(0);

    s_advancedPoolHooks.applyAllowListUpdates(new address[](0), newAddresses);
    address[] memory setAddresses = s_advancedPoolHooks.getAllowList();

    assertEq(setAddresses.length, setAddressesLength);
  }

  // Reverts

  function test_RevertWhen_OnlyOwner() public {
    vm.stopPrank();
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    address[] memory newAddresses = new address[](2);
    s_advancedPoolHooks.applyAllowListUpdates(new address[](0), newAddresses);
  }

  function test_RevertWhen_AllowListNotEnabled() public {
    AdvancedPoolHooks hooksWithoutAllowList = new AdvancedPoolHooks(new address[](0), 0);

    vm.expectRevert(AdvancedPoolHooks.AllowListNotEnabled.selector);

    hooksWithoutAllowList.applyAllowListUpdates(new address[](0), new address[](2));
  }
}
