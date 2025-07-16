// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {Test} from "forge-std/Test.sol";

contract ERC20LockBox_constructor is Test {
  function test_Constructor_Success() public {
    BurnMintERC20 token = new BurnMintERC20("LINK", "LNK", 18, 0, 0);
    ERC20LockBox lockBox = new ERC20LockBox(address(token));
    assertEq(lockBox.i_token(), address(token));
  }

  function test_Constructor_RevertWhen_TokenAddressIsZero() public {
    vm.expectRevert(ERC20LockBox.TokenAddressCannotBeZero.selector);
    new ERC20LockBox(address(0));
  }
}
