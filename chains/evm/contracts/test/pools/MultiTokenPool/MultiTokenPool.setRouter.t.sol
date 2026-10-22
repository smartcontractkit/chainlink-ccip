// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MultiTokenPool} from "../../../pools/MultiTokenPool/MultiTokenPool.sol";
import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract MultiTokenPool_setRouter is MultiTokenPoolSetup {
  function test_setRouter() public {
    address newRouter = makeAddr("newRouter");

    vm.expectEmit();
    emit MultiTokenPool.RouterSet(address(s_sourceRouter), newRouter);
    s_pool.setRouter(newRouter);

    assertEq(newRouter, s_pool.getRouter());
  }

  function test_setRouter_RevertWhen_ZeroAddressInvalid() public {
    vm.expectRevert(MultiTokenPool.ZeroAddressInvalid.selector);
    s_pool.setRouter(address(0));
  }

  function test_setRouter_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_pool.setRouter(makeAddr("newRouter"));
  }
}
