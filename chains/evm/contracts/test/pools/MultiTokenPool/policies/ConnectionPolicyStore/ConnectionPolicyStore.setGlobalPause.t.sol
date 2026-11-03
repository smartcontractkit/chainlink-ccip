// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract ConnectionPolicyStore_setGlobalPause is ConnectionPolicyStoreSetup {
  function test_setGlobalPause() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.GlobalPauseSet(true);
    s_connectionStore.setGlobalPause(true);

    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_token, DEST_CHAIN_SELECTOR);
    assertTrue(connection.isGloballyPaused);
    assertTrue(connection.isOutboundPaused);
    assertTrue(connection.isInboundPaused);
    vm.expectRevert(ConnectionPolicyStore.GloballyPaused.selector);
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");

    vm.expectEmit();
    emit ConnectionPolicyStore.GlobalPauseSet(false);
    s_connectionStore.setGlobalPause(false);

    connection = _getLaneConnection(s_token, DEST_CHAIN_SELECTOR);
    assertFalse(connection.isGloballyPaused);
    assertFalse(connection.isOutboundPaused);
    assertFalse(connection.isInboundPaused);
  }

  function test_setGlobalPause_KeepsTokenAndChainPauses() public {
    uint64[] memory remoteChainSelectors = _getRemoteChainSelectors(DEST_CHAIN_SELECTOR);
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Outbound, true);
    s_connectionStore.setChainPause(remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Inbound, true);

    s_connectionStore.setGlobalPause(true);
    s_connectionStore.setGlobalPause(false);

    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_token, DEST_CHAIN_SELECTOR);
    assertFalse(connection.isGloballyPaused);
    assertTrue(connection.isTokenOutboundPaused);
    assertFalse(connection.isTokenInboundPaused);
    assertFalse(connection.isChainOutboundPaused);
    assertTrue(connection.isChainInboundPaused);
    assertTrue(connection.isOutboundPaused);
    assertTrue(connection.isInboundPaused);
  }

  function test_setGlobalPause_NoChange() public {
    vm.recordLogs();
    s_connectionStore.setGlobalPause(false);
    assertEq(0, vm.getRecordedLogs().length);

    s_connectionStore.setGlobalPause(true);
    vm.recordLogs();
    s_connectionStore.setGlobalPause(true);
    assertEq(0, vm.getRecordedLogs().length);
  }

  // Reverts

  function test_setGlobalPause_RevertWhen_UnauthorizedCaller() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_connectionStore.setGlobalPause(true);
  }
}
