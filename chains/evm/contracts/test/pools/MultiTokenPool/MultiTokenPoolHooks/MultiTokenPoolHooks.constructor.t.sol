// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {MultiTokenPoolHooks} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooks.sol";
import {MultiTokenPoolHooksSetup} from "./MultiTokenPoolHooksSetup.t.sol";

contract MultiTokenPoolHooks_constructor is MultiTokenPoolHooksSetup {
  function test_constructor() public {
    address[] memory allowlist = new address[](2);
    allowlist[0] = s_sender;
    allowlist[1] = STRANGER;

    vm.expectEmit();
    emit MultiTokenPoolHooks.AllowListAdd(s_sender);
    vm.expectEmit();
    emit MultiTokenPoolHooks.AllowListAdd(STRANGER);
    vm.expectEmit();
    emit MultiTokenPoolHooks.PolicyEngineAttached(s_policyEngine);
    MultiTokenPoolHooks hooks = new MultiTokenPoolHooks(allowlist, s_policyEngine, _arrayOf(OWNER));

    assertEq("MultiTokenPoolHooks 2.1.0-dev", hooks.typeAndVersion());
    assertTrue(hooks.getAllowListEnabled());
    assertEq(allowlist, hooks.getAllowList());
    assertEq(address(s_policyEngine), address(hooks.getPolicyEngine()));
    assertTrue(s_policyEngine.isAttached(address(hooks)));
    assertEq(_arrayOf(OWNER), hooks.getAllAuthorizedCallers());
  }

  function test_constructor_NoAllowListAndNoPolicyEngine() public {
    vm.recordLogs();
    MultiTokenPoolHooks hooks = new MultiTokenPoolHooks(new address[](0), IPolicyEngine(address(0)), new address[](0));

    assertEq(0, vm.getRecordedLogs().length);
    assertFalse(hooks.getAllowListEnabled());
    assertEq(0, hooks.getAllowList().length);
    assertEq(address(0), address(hooks.getPolicyEngine()));
    assertEq(0, hooks.getAllAuthorizedCallers().length);
  }

  function test_constructor_AllowListOfZeroAddress() public {
    MultiTokenPoolHooks hooks =
      new MultiTokenPoolHooks(_arrayOf(address(0)), IPolicyEngine(address(0)), new address[](0));

    // A list with only the zero address still enables the allowlist, but leaves it empty, so nobody can send.
    assertTrue(hooks.getAllowListEnabled());
    assertEq(0, hooks.getAllowList().length);
  }
}
