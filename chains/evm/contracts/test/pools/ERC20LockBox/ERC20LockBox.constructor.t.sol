// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {Test} from "forge-std/Test.sol";

contract ERC20LockBox_constructor is Test {
  function test_Constructor_Success() public {
    ERC20LockBox lockBox = new ERC20LockBox(address(3));
    assertEq(address(lockBox.i_tokenAdminRegistry()), address(3));
  }

  function test_Constructor_RevertWhen_TokenAdminRegistryIsZeroAddress() public {
    vm.expectRevert(ERC20LockBox.ZeroAddressNotAllowed.selector);
    new ERC20LockBox(address(0));
  }
}
