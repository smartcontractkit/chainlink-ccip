// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {Test} from "forge-std/Test.sol";

contract ERC20LockBox_constructor is Test {
  function test_Constructor_Success() public {
    ERC20LockBox lockBox = new ERC20LockBox(address(3), 0);
    assertEq(address(lockBox.getToken()), address(3));
    assertEq(lockBox.i_remoteChainSelector(), 0);
    assertEq(lockBox.typeAndVersion(), "ERC20LockBox 1.7.0-dev");
  }

  function test_Constructor_RevertWhen_TokenIsZeroAddress() public {
    vm.expectRevert(ERC20LockBox.ZeroAddressNotAllowed.selector);
    new ERC20LockBox(address(0), 0);
  }
}
