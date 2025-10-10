// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RouterSetup} from "./RouterSetup.t.sol";

contract Router_setWrappedNative is RouterSetup {
  function testFuzz_SetWrappedNative_Success(
    address wrappedNative
  ) public {
    s_sourceRouter.setWrappedNative(wrappedNative);
    assertEq(wrappedNative, s_sourceRouter.getWrappedNative());
  }

  // Reverts
  function test_RevertWhen_OnlyOwner() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_sourceRouter.setWrappedNative(address(1));
  }
}
