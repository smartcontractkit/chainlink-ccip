// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {ERC20LockBoxSetup} from "./ERC20LockBoxSetup.t.sol";

contract ERC20LockBox_configureAllowedCallers is ERC20LockBoxSetup {
  function test_ConfigureAllowedCallers_AddSingleCaller() public {
    address newCaller = makeAddr("new_caller");
    address token = address(s_token);

    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: newCaller, allowed: true});

    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(token, newCaller, true);

    s_erc20LockBox.configureAllowedCallers(configArgs);

    assertTrue(s_erc20LockBox.isAllowedCaller(token, newCaller));
  }

  function test_ConfigureAllowedCallers_AddMultipleCallers() public {
    address caller1 = makeAddr("caller1");
    address caller2 = makeAddr("caller2");
    address caller3 = makeAddr("caller3");
    address token = address(s_token);

    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](3);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: caller1, allowed: true});
    configArgs[1] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: caller2, allowed: true});
    configArgs[2] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: caller3, allowed: true});

    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(token, caller1, true);
    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(token, caller2, true);
    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(token, caller3, true);

    s_erc20LockBox.configureAllowedCallers(configArgs);

    assertTrue(s_erc20LockBox.isAllowedCaller(token, caller1));
    assertTrue(s_erc20LockBox.isAllowedCaller(token, caller2));
    assertTrue(s_erc20LockBox.isAllowedCaller(token, caller3));
  }

  function test_ConfigureAllowedCallers_RemoveSingleCaller() public {
    address callerToRemove = makeAddr("caller_to_remove");
    address token = address(s_token);

    // First add a caller
    ERC20LockBox.AllowedCallerConfigArgs[] memory addConfig = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    addConfig[0] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: callerToRemove, allowed: true});
    s_erc20LockBox.configureAllowedCallers(addConfig);

    assertTrue(s_erc20LockBox.isAllowedCaller(token, callerToRemove));

    // Now remove the caller
    ERC20LockBox.AllowedCallerConfigArgs[] memory removeConfig = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    removeConfig[0] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: callerToRemove, allowed: false});

    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(token, callerToRemove, false);

    s_erc20LockBox.configureAllowedCallers(removeConfig);

    assertFalse(s_erc20LockBox.isAllowedCaller(token, callerToRemove));
  }

  function test_ConfigureAllowedCallers_AddAndRemoveInSameCall() public {
    address callerToAdd = makeAddr("caller_to_add");
    address callerToRemove = s_allowedCaller; // Use the existing allowed caller
    address token = address(s_token);

    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](2);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: callerToAdd, allowed: true});
    configArgs[1] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: callerToRemove, allowed: false});

    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(token, callerToAdd, true);
    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(token, callerToRemove, false);

    s_erc20LockBox.configureAllowedCallers(configArgs);

    assertTrue(s_erc20LockBox.isAllowedCaller(token, callerToAdd));
    assertFalse(s_erc20LockBox.isAllowedCaller(token, callerToRemove));
  }

  function test_ConfigureAllowedCallers_AddAlreadyAllowedCaller() public {
    address token = address(s_token);

    // Try to add a caller that's already allowed
    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: s_allowedCaller, allowed: true});

    // Should not emit event since caller is already allowed
    s_erc20LockBox.configureAllowedCallers(configArgs);

    assertTrue(s_erc20LockBox.isAllowedCaller(token, s_allowedCaller));
  }

  function test_ConfigureAllowedCallers_RemoveNonExistentCaller() public {
    address nonExistentCaller = makeAddr("non_existent_caller");
    address token = address(s_token);

    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: nonExistentCaller, allowed: false});

    // Should not emit event since caller was not allowed
    s_erc20LockBox.configureAllowedCallers(configArgs);

    assertFalse(s_erc20LockBox.isAllowedCaller(token, nonExistentCaller));
  }

  function test_ConfigureAllowedCallers_RemoveAllCallers() public {
    address caller1 = makeAddr("caller1");
    address caller2 = makeAddr("caller2");
    address token = address(s_token);

    // Add two new callers
    ERC20LockBox.AllowedCallerConfigArgs[] memory addConfig = new ERC20LockBox.AllowedCallerConfigArgs[](2);
    addConfig[0] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: caller1, allowed: true});
    addConfig[1] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: caller2, allowed: true});
    s_erc20LockBox.configureAllowedCallers(addConfig);

    // Now remove all callers including the original one
    ERC20LockBox.AllowedCallerConfigArgs[] memory removeConfig = new ERC20LockBox.AllowedCallerConfigArgs[](3);
    removeConfig[0] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: s_allowedCaller, allowed: false});
    removeConfig[1] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: caller1, allowed: false});
    removeConfig[2] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: caller2, allowed: false});

    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(token, s_allowedCaller, false);
    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(token, caller1, false);
    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(token, caller2, false);

    s_erc20LockBox.configureAllowedCallers(removeConfig);

    // Verify all callers are removed
    assertFalse(s_erc20LockBox.isAllowedCaller(token, s_allowedCaller));
    assertFalse(s_erc20LockBox.isAllowedCaller(token, caller1));
    assertFalse(s_erc20LockBox.isAllowedCaller(token, caller2));
  }

  function test_ConfigureAllowedCallers_ReAddRemovedCaller() public {
    address caller = makeAddr("caller");
    address token = address(s_token);

    // Add caller
    ERC20LockBox.AllowedCallerConfigArgs[] memory addConfig = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    addConfig[0] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: caller, allowed: true});
    s_erc20LockBox.configureAllowedCallers(addConfig);
    assertTrue(s_erc20LockBox.isAllowedCaller(token, caller));

    // Remove caller
    ERC20LockBox.AllowedCallerConfigArgs[] memory removeConfig = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    removeConfig[0] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: caller, allowed: false});
    s_erc20LockBox.configureAllowedCallers(removeConfig);
    assertFalse(s_erc20LockBox.isAllowedCaller(token, caller));

    // Re-add caller
    ERC20LockBox.AllowedCallerConfigArgs[] memory reAddConfig = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    reAddConfig[0] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: caller, allowed: true});

    vm.expectEmit();
    emit ERC20LockBox.AllowedCallerUpdated(token, caller, true);

    s_erc20LockBox.configureAllowedCallers(reAddConfig);
    assertTrue(s_erc20LockBox.isAllowedCaller(token, caller));
  }

  function test_ConfigureAllowedCallers_ZeroAddress() public {
    address zeroAddress = address(0);
    address token = address(s_token);

    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: zeroAddress, allowed: true});

    // Should not revert - zero address can be added as allowed caller
    s_erc20LockBox.configureAllowedCallers(configArgs);

    assertTrue(s_erc20LockBox.isAllowedCaller(token, zeroAddress));
  }

  // ================================================================
  // │                        Revert Tests                          │
  // ================================================================

  function test_RevertWhen_Unauthorized() public {
    address newCaller = makeAddr("new_caller");
    address token = address(s_token);

    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({token: token, caller: newCaller, allowed: true});

    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(ERC20LockBox.Unauthorized.selector, STRANGER));

    s_erc20LockBox.configureAllowedCallers(configArgs);
  }

  function test_ConfigureAllowedCallers_ZeroTokenAddress() public {
    address newCaller = makeAddr("new_caller");
    address zeroToken = address(0);

    ERC20LockBox.AllowedCallerConfigArgs[] memory configArgs = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    configArgs[0] = ERC20LockBox.AllowedCallerConfigArgs({token: zeroToken, caller: newCaller, allowed: true});

    vm.expectRevert(ERC20LockBox.TokenAddressCannotBeZero.selector);
    s_erc20LockBox.configureAllowedCallers(configArgs);
  }
}
