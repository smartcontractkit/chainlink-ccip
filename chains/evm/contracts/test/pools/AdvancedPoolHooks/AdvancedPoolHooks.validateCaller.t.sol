// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract AdvancedPoolHooks_validateCaller is AdvancedPoolHooksSetup {
  address internal s_authorizedCaller = makeAddr("authorizedCaller");
  address internal s_unauthorizedCaller = makeAddr("unauthorizedCaller");

  AdvancedPoolHooks internal s_hooksWithAuthorizedCallers;

  function setUp() public virtual override {
    super.setUp();

    address[] memory authorizedCallers = new address[](1);
    authorizedCallers[0] = s_authorizedCaller;
    s_hooksWithAuthorizedCallers = new AdvancedPoolHooks(new address[](0), 0, address(0), authorizedCallers);

    vm.stopPrank();
  }

  function test_getAuthorizedCallersEnabled() public view {
    assertFalse(s_advancedPoolHooks.getAuthorizedCallersEnabled());
    assertTrue(s_hooksWithAuthorizedCallers.getAuthorizedCallersEnabled());
  }

  function test_validateCaller_authorizedCallersDisabled() public {
    assertFalse(s_advancedPoolHooks.getAuthorizedCallersEnabled());

    vm.prank(s_unauthorizedCaller);
    s_advancedPoolHooks.validateCaller();
  }

  function test_validateCaller_callerIsAuthorized() public {
    assertTrue(s_hooksWithAuthorizedCallers.getAuthorizedCallersEnabled());

    vm.prank(s_authorizedCaller);
    s_hooksWithAuthorizedCallers.validateCaller();
  }

  function test_validateCaller_RevertWhen_UnauthorizedCaller() public {
    assertTrue(s_hooksWithAuthorizedCallers.getAuthorizedCallersEnabled());

    vm.prank(s_unauthorizedCaller);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, s_unauthorizedCaller));
    s_hooksWithAuthorizedCallers.validateCaller();
  }
}
