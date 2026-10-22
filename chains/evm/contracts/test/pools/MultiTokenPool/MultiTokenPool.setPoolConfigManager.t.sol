// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MultiTokenPool} from "../../../pools/MultiTokenPool/MultiTokenPool.sol";
import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract MultiTokenPool_setPoolConfigManager is MultiTokenPoolSetup {
  function test_setPoolConfigManager() public {
    address newPoolConfigManager = makeAddr("newPoolConfigManager");

    vm.expectEmit();
    emit MultiTokenPool.PoolConfigManagerSet(address(s_poolConfigManager), newPoolConfigManager);
    s_pool.setPoolConfigManager(newPoolConfigManager);

    assertEq(newPoolConfigManager, s_pool.getPoolConfigManager());
  }

  function test_setPoolConfigManager_RevertWhen_ZeroAddressInvalid() public {
    vm.expectRevert(MultiTokenPool.ZeroAddressInvalid.selector);
    s_pool.setPoolConfigManager(address(0));
  }

  function test_setPoolConfigManager_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_pool.setPoolConfigManager(makeAddr("newPoolConfigManager"));
  }
}
