// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BurnMintMultiTokenPool} from "../../../pools/MultiTokenPool/BurnMintMultiTokenPool.sol";
import {MultiTokenPool} from "../../../pools/MultiTokenPool/MultiTokenPool.sol";
import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_constructor is MultiTokenPoolSetup {
  function test_constructor() public {
    address poolConfigManager = makeAddr("poolConfigManager");

    vm.expectEmit();
    emit MultiTokenPool.RouterSet(address(0), address(s_sourceRouter));
    vm.expectEmit();
    emit MultiTokenPool.PoolConfigManagerSet(address(0), poolConfigManager);

    BurnMintMultiTokenPool pool =
      new BurnMintMultiTokenPool(address(s_mockRMNRemote), address(s_sourceRouter), poolConfigManager);

    assertEq(address(s_mockRMNRemote), pool.getRmnProxy());
    assertEq(address(s_sourceRouter), pool.getRouter());
    assertEq(poolConfigManager, pool.getPoolConfigManager());
  }

  function test_constructor_RevertWhen_ZeroAddressInvalid() public {
    vm.expectRevert(MultiTokenPool.ZeroAddressInvalid.selector);
    new BurnMintMultiTokenPool(address(0), address(s_sourceRouter), address(s_poolConfigManager));

    vm.expectRevert(MultiTokenPool.ZeroAddressInvalid.selector);
    new BurnMintMultiTokenPool(address(s_mockRMNRemote), address(0), address(s_poolConfigManager));

    vm.expectRevert(MultiTokenPool.ZeroAddressInvalid.selector);
    new BurnMintMultiTokenPool(address(s_mockRMNRemote), address(s_sourceRouter), address(0));
  }
}
