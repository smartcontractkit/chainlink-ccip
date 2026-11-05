// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPoolRateLimitHelper} from "../../../pools/TokenPoolRateLimitHelper.sol";
import {TokenPoolRateLimitHelperSetup} from "./TokenPoolRateLimitHelperSetup.t.sol";

import {Ownable} from "@openzeppelin/contracts@5.3.0/access/Ownable.sol";

contract TokenPoolRateLimitHelper_setManager is TokenPoolRateLimitHelperSetup {
  function test_setManager() public {
    assertFalse(s_helper.isManager(STRANGER));

    vm.expectEmit();
    emit TokenPoolRateLimitHelper.ManagerSet(STRANGER, true);
    s_helper.setManager(STRANGER, true);
    assertTrue(s_helper.isManager(STRANGER));

    vm.expectEmit();
    emit TokenPoolRateLimitHelper.ManagerSet(STRANGER, false);
    s_helper.setManager(STRANGER, false);
    assertFalse(s_helper.isManager(STRANGER));
  }

  function test_RevertWhen_CallerIsNotOwner() public {
    vm.startPrank(s_manager);
    vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, s_manager));
    s_helper.setManager(STRANGER, true);
  }

  function test_RevertWhen_ZeroAddress() public {
    vm.expectRevert(TokenPoolRateLimitHelper.ZeroAddressNotAllowed.selector);
    s_helper.setManager(address(0), true);
  }
}
