// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAddressBlocklist} from "../../../../interfaces/IAddressBlocklist.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {MultiTokenPoolHooks} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooks.sol";
import {
  MultiTokenPoolHooksExternalBlocklist
} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooksExternalBlocklist.sol";
import {MockPolicyEngine} from "../../../mocks/MockPolicyEngine.sol";
import {MultiTokenPoolHooksExternalBlocklistSetup} from "./MultiTokenPoolHooksExternalBlocklistSetup.t.sol";

contract MultiTokenPoolHooksExternalBlocklist_constructor is MultiTokenPoolHooksExternalBlocklistSetup {
  function test_constructor() public {
    MockPolicyEngine policyEngine = new MockPolicyEngine();

    vm.expectEmit();
    emit MultiTokenPoolHooks.AllowListAdd(s_sender);
    vm.expectEmit();
    emit MultiTokenPoolHooks.PolicyEngineAttached(policyEngine);
    vm.expectEmit();
    emit MultiTokenPoolHooksExternalBlocklist.BlocklistSet(IAddressBlocklist(address(0)), s_blocklist);
    MultiTokenPoolHooksExternalBlocklist hooks =
      new MultiTokenPoolHooksExternalBlocklist(_arrayOf(s_sender), policyEngine, s_blocklist, _arrayOf(OWNER));

    assertEq("MultiTokenPoolHooksExternalBlocklist 2.1.0-dev", hooks.typeAndVersion());
    assertEq(address(s_blocklist), address(hooks.getBlocklist()));
    assertEq(_arrayOf(s_sender), hooks.getAllowList());
    assertEq(address(policyEngine), address(hooks.getPolicyEngine()));
    assertEq(_arrayOf(OWNER), hooks.getAllAuthorizedCallers());
  }

  function test_constructor_NoBlocklist() public {
    vm.recordLogs();
    MultiTokenPoolHooksExternalBlocklist hooks = new MultiTokenPoolHooksExternalBlocklist(
      new address[](0), IPolicyEngine(address(0)), IAddressBlocklist(address(0)), new address[](0)
    );

    assertEq(0, vm.getRecordedLogs().length);
    assertEq(address(0), address(hooks.getBlocklist()));
  }

  // Reverts

  function test_constructor_RevertWhen_InvalidBlocklist() public {
    IAddressBlocklist blocklistWithoutCode = IAddressBlocklist(makeAddr("blocklistWithoutCode"));

    vm.expectRevert(
      abi.encodeWithSelector(MultiTokenPoolHooksExternalBlocklist.InvalidBlocklist.selector, blocklistWithoutCode)
    );
    new MultiTokenPoolHooksExternalBlocklist(
      new address[](0), IPolicyEngine(address(0)), blocklistWithoutCode, _arrayOf(OWNER)
    );
  }
}
