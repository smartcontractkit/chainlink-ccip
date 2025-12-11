// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {ERC20LockBoxSetup} from "./ERC20LockBoxSetup.t.sol";

contract ERC20LockBox_configureAllowedCallers is ERC20LockBoxSetup {
  function test_ConfigureAllowedCallers_AddSingleCaller() public {
    address newCaller = makeAddr("new_caller");

    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({caller: newCaller, allowed: true});

    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(address(s_token), newCaller, true);

    s_erc20LockBox.configureAllowedCallers(configArgs);

    assertTrue(s_erc20LockBox.isAllowedCaller(newCaller));
  }

  function test_ConfigureAllowedCallers_AddMultipleCallers() public {
    address caller1 = makeAddr("caller1");
    address caller2 = makeAddr("caller2");
    address caller3 = makeAddr("caller3");

    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](3);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({caller: caller1, allowed: true});
    configArgs[1] = ERC20LockBox.AllowedCallerConfigArgs({caller: caller2, allowed: true});
    configArgs[2] = ERC20LockBox.AllowedCallerConfigArgs({caller: caller3, allowed: true});

    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(address(s_token), caller1, true);
    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(address(s_token), caller2, true);
    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(address(s_token), caller3, true);

    s_erc20LockBox.configureAllowedCallers(configArgs);

    assertTrue(s_erc20LockBox.isAllowedCaller(caller1));
    assertTrue(s_erc20LockBox.isAllowedCaller(caller2));
    assertTrue(s_erc20LockBox.isAllowedCaller(caller3));
  }

  function test_ConfigureAllowedCallers_RemoveSingleCaller() public {
    address callerToRemove = makeAddr("caller_to_remove");

    // First add a caller
    ERC20LockBox.AllowedCallerConfigArgs[] memory addConfig = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    addConfig[0] = ERC20LockBox.AllowedCallerConfigArgs({caller: callerToRemove, allowed: true});
    s_erc20LockBox.configureAllowedCallers(addConfig);

    assertTrue(s_erc20LockBox.isAllowedCaller(callerToRemove));

    // Now remove the caller
    ERC20LockBox.AllowedCallerConfigArgs[] memory removeConfig = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    removeConfig[0] = ERC20LockBox.AllowedCallerConfigArgs({caller: callerToRemove, allowed: false});

    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(address(s_token), callerToRemove, false);

    s_erc20LockBox.configureAllowedCallers(removeConfig);

    assertFalse(s_erc20LockBox.isAllowedCaller(callerToRemove));
  }

  function test_ConfigureAllowedCallers_AddAndRemoveInSameCall() public {
    address callerToAdd = makeAddr("caller_to_add");
    address callerToRemove = s_allowedCaller; // Use the existing allowed caller

    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](2);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({caller: callerToAdd, allowed: true});
    configArgs[1] = ERC20LockBox.AllowedCallerConfigArgs({caller: callerToRemove, allowed: false});

    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(address(s_token), callerToAdd, true);
    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(address(s_token), callerToRemove, false);

    s_erc20LockBox.configureAllowedCallers(configArgs);

    assertTrue(s_erc20LockBox.isAllowedCaller(callerToAdd));
    assertFalse(s_erc20LockBox.isAllowedCaller(callerToRemove));
  }

  function test_ConfigureAllowedCallers_AddAlreadyAllowedCaller() public {
    // Try to add a caller that's already allowed
    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({caller: s_allowedCaller, allowed: true});

    // Should not emit event since caller is already allowed
    s_erc20LockBox.configureAllowedCallers(configArgs);

    assertTrue(s_erc20LockBox.isAllowedCaller(s_allowedCaller));
  }

  function test_ConfigureAllowedCallers_RemoveNonExistentCaller() public {
    address nonExistentCaller = makeAddr("non_existent_caller");

    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({caller: nonExistentCaller, allowed: false});

    // Should not emit event since caller was not allowed
    s_erc20LockBox.configureAllowedCallers(configArgs);

    assertFalse(s_erc20LockBox.isAllowedCaller(nonExistentCaller));
  }

  function test_ConfigureAllowedCallers_RemoveAllCallers() public {
    address caller1 = makeAddr("caller1");
    address caller2 = makeAddr("caller2");

    // Add two new callers
    ERC20LockBox.AllowedCallerConfigArgs[] memory addConfig = new ERC20LockBox.AllowedCallerConfigArgs[](2);
    addConfig[0] = ERC20LockBox.AllowedCallerConfigArgs({caller: caller1, allowed: true});
    addConfig[1] = ERC20LockBox.AllowedCallerConfigArgs({caller: caller2, allowed: true});
    s_erc20LockBox.configureAllowedCallers(addConfig);

    // Now remove all callers including the original one
    ERC20LockBox.AllowedCallerConfigArgs[] memory removeConfig = new ERC20LockBox.AllowedCallerConfigArgs[](3);
    removeConfig[0] = ERC20LockBox.AllowedCallerConfigArgs({caller: s_allowedCaller, allowed: false});
    removeConfig[1] = ERC20LockBox.AllowedCallerConfigArgs({caller: caller1, allowed: false});
    removeConfig[2] = ERC20LockBox.AllowedCallerConfigArgs({caller: caller2, allowed: false});

    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(address(s_token), s_allowedCaller, false);
    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(address(s_token), caller1, false);
    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(address(s_token), caller2, false);

    s_erc20LockBox.configureAllowedCallers(removeConfig);

    // Verify all callers are removed
    assertFalse(s_erc20LockBox.isAllowedCaller(s_allowedCaller));
    assertFalse(s_erc20LockBox.isAllowedCaller(caller1));
    assertFalse(s_erc20LockBox.isAllowedCaller(caller2));
  }

  function test_ConfigureAllowedCallers_ReAddRemovedCaller() public {
    address caller = makeAddr("caller");

    // Add caller
    ERC20LockBox.AllowedCallerConfigArgs[] memory addConfig = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    addConfig[0] = ERC20LockBox.AllowedCallerConfigArgs({caller: caller, allowed: true});
    s_erc20LockBox.configureAllowedCallers(addConfig);
    assertTrue(s_erc20LockBox.isAllowedCaller(caller));

    // Remove caller
    ERC20LockBox.AllowedCallerConfigArgs[] memory removeConfig = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    removeConfig[0] = ERC20LockBox.AllowedCallerConfigArgs({caller: caller, allowed: false});
    s_erc20LockBox.configureAllowedCallers(removeConfig);
    assertFalse(s_erc20LockBox.isAllowedCaller(caller));

    // Re-add caller
    ERC20LockBox.AllowedCallerConfigArgs[] memory reAddConfig = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    reAddConfig[0] = ERC20LockBox.AllowedCallerConfigArgs({caller: caller, allowed: true});

    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(address(s_token), caller, true);

    s_erc20LockBox.configureAllowedCallers(reAddConfig);
    assertTrue(s_erc20LockBox.isAllowedCaller(caller));
  }

  // ================================================================
  // │                        Revert Tests                          │
  // ================================================================

  function test_RevertWhen_Unauthorized() public {
    address newCaller = makeAddr("new_caller");

    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({caller: newCaller, allowed: true});

    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(ERC20LockBox.Unauthorized.selector, STRANGER));

    s_erc20LockBox.configureAllowedCallers(configArgs);
  }

  function test_ConfigureAllowedCallers_ZeroCaller() public {
    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({caller: address(0), allowed: true});

    vm.expectRevert(ERC20LockBox.ZeroAddressNotAllowed.selector);
    s_erc20LockBox.configureAllowedCallers(configArgs);
  }
}
