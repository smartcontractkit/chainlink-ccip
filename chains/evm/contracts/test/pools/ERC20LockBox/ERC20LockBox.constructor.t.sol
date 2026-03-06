// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {Test} from "forge-std/Test.sol";

contract ERC20LockBox_constructor is Test {
  function test_constructor() public {
    address token = makeAddr("TOKEN");
    ERC20LockBox lockBox = new ERC20LockBox(token);
    assertTrue(lockBox.isTokenSupported(token));
    assertEq(address(lockBox.getToken()), token);
  }

  function test_constructor_RevertWhen_TokenIsZeroAddress() public {
    vm.expectRevert(AuthorizedCallers.ZeroAddressNotAllowed.selector);
    new ERC20LockBox(address(0));
  }
}
