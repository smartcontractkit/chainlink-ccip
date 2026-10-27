// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IMultiTokenPool} from "../../../interfaces/IMultiTokenPool.sol";
import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";

import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_isSupportedChain is MultiTokenPoolSetup {
  function test_isSupportedChain() public view {
    assertTrue(s_pool.isSupportedChain(address(s_token), DEST_CHAIN_SELECTOR));
  }

  function test_isSupportedChain_FalseWhenNoLane() public {
    IPoolConfigManager.LaneConnection memory connection = _defaultLaneConnection();
    connection.hasConnection = false;
    connection.remoteToken = "";
    s_poolConfigManager.setLaneConnection(connection);

    assertFalse(s_pool.isSupportedChain(address(s_token), DEST_CHAIN_SELECTOR));
  }

  function test_isSupportedChain_RevertWhen_LocalTokenRequired() public {
    vm.expectRevert(IMultiTokenPool.LocalTokenRequired.selector);
    s_pool.isSupportedChain(DEST_CHAIN_SELECTOR);
  }
}
