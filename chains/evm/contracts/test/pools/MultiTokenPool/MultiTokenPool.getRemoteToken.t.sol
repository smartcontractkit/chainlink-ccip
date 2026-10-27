// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IMultiTokenPool} from "../../../interfaces/IMultiTokenPool.sol";
import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";

import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_getRemoteToken is MultiTokenPoolSetup {
  function test_getRemoteToken() public view {
    assertEq(abi.encode(s_remoteToken), s_pool.getRemoteToken(address(s_token), DEST_CHAIN_SELECTOR));
  }

  function test_getRemoteToken_EmptyWhenNoLane() public {
    IPoolConfigManager.LaneConnection memory connection = _defaultLaneConnection();
    connection.hasConnection = false;
    connection.remoteToken = "";
    s_poolConfigManager.setLaneConnection(connection);

    assertEq("", s_pool.getRemoteToken(address(s_token), DEST_CHAIN_SELECTOR));
  }

  function test_getRemoteToken_RevertWhen_LocalTokenRequired() public {
    vm.expectRevert(IMultiTokenPool.LocalTokenRequired.selector);
    s_pool.getRemoteToken(DEST_CHAIN_SELECTOR);
  }
}
