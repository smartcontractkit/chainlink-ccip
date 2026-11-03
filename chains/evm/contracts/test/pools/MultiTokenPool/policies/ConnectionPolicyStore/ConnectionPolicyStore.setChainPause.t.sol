// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {Vm} from "forge-std/Vm.sol";

contract ConnectionPolicyStore_setChainPause is ConnectionPolicyStoreSetup {
  function test_setChainPause() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.ChainPauseSet(DEST_CHAIN_SELECTOR, true, false);
    s_connectionStore.setChainPause(
      _getRemoteChainSelectors(DEST_CHAIN_SELECTOR), ConnectionPolicyStore.PauseDirection.Outbound, true
    );

    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_token, DEST_CHAIN_SELECTOR);
    assertTrue(connection.isChainOutboundPaused);
    assertFalse(connection.isChainInboundPaused);
    assertTrue(connection.isOutboundPaused);
    assertFalse(connection.isInboundPaused);
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.ChainIsPaused.selector, DEST_CHAIN_SELECTOR, ConnectionPolicyStore.PauseDirection.Outbound
      )
    );
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
  }

  function test_setChainPause_Inbound() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.ChainPauseSet(DEST_CHAIN_SELECTOR, false, true);
    s_connectionStore.setChainPause(
      _getRemoteChainSelectors(DEST_CHAIN_SELECTOR), ConnectionPolicyStore.PauseDirection.Inbound, true
    );

    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_token, DEST_CHAIN_SELECTOR);
    assertFalse(connection.isChainOutboundPaused);
    assertTrue(connection.isChainInboundPaused);
    assertFalse(connection.isOutboundPaused);
    assertTrue(connection.isInboundPaused);
  }

  function test_setChainPause_Both() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.ChainPauseSet(DEST_CHAIN_SELECTOR, true, true);
    s_connectionStore.setChainPause(
      _getRemoteChainSelectors(DEST_CHAIN_SELECTOR), ConnectionPolicyStore.PauseDirection.Both, true
    );

    vm.expectEmit();
    emit ConnectionPolicyStore.ChainPauseSet(DEST_CHAIN_SELECTOR, false, false);
    s_connectionStore.setChainPause(
      _getRemoteChainSelectors(DEST_CHAIN_SELECTOR), ConnectionPolicyStore.PauseDirection.Both, false
    );

    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_token, DEST_CHAIN_SELECTOR);
    assertFalse(connection.isChainOutboundPaused);
    assertFalse(connection.isChainInboundPaused);
  }

  function test_setChainPause_OtherDirectionStays() public {
    s_connectionStore.setChainPause(
      _getRemoteChainSelectors(DEST_CHAIN_SELECTOR), ConnectionPolicyStore.PauseDirection.Outbound, true
    );

    vm.expectEmit();
    emit ConnectionPolicyStore.ChainPauseSet(DEST_CHAIN_SELECTOR, true, true);
    s_connectionStore.setChainPause(
      _getRemoteChainSelectors(DEST_CHAIN_SELECTOR), ConnectionPolicyStore.PauseDirection.Inbound, true
    );

    vm.expectEmit();
    emit ConnectionPolicyStore.ChainPauseSet(DEST_CHAIN_SELECTOR, false, true);
    s_connectionStore.setChainPause(
      _getRemoteChainSelectors(DEST_CHAIN_SELECTOR), ConnectionPolicyStore.PauseDirection.Outbound, false
    );

    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_token, DEST_CHAIN_SELECTOR);
    assertFalse(connection.isChainOutboundPaused);
    assertTrue(connection.isChainInboundPaused);
  }

  function test_setChainPause_NoChange() public {
    s_connectionStore.setChainPause(
      _getRemoteChainSelectors(DEST_CHAIN_SELECTOR), ConnectionPolicyStore.PauseDirection.Outbound, true
    );
    uint64[] memory remoteChainSelectors = new uint64[](2);
    remoteChainSelectors[0] = DEST_CHAIN_SELECTOR;
    remoteChainSelectors[1] = SOURCE_CHAIN_SELECTOR;

    // Only the chain whose pause state changes should emit an event.
    vm.recordLogs();
    s_connectionStore.setChainPause(remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Outbound, true);

    Vm.Log[] memory logs = vm.getRecordedLogs();
    assertEq(1, logs.length);
    assertEq(ConnectionPolicyStore.ChainPauseSet.selector, logs[0].topics[0]);
    assertEq(bytes32(uint256(SOURCE_CHAIN_SELECTOR)), logs[0].topics[1]);

    vm.recordLogs();
    s_connectionStore.setChainPause(
      _getRemoteChainSelectors(DEST_CHAIN_SELECTOR), ConnectionPolicyStore.PauseDirection.Inbound, false
    );
    assertEq(0, vm.getRecordedLogs().length);
  }

  function test_setChainPause_ChainWithoutLane() public {
    // Chain selector zero cannot have a connection, so pausing it cannot block transfers.
    uint64[] memory remoteChainSelectors = new uint64[](2);
    remoteChainSelectors[0] = SOURCE_CHAIN_SELECTOR;
    remoteChainSelectors[1] = PolicyConfig.DEFAULT_CHAIN_SELECTOR;

    vm.expectEmit();
    emit ConnectionPolicyStore.ChainPauseSet(SOURCE_CHAIN_SELECTOR, true, true);
    vm.expectEmit();
    emit ConnectionPolicyStore.ChainPauseSet(PolicyConfig.DEFAULT_CHAIN_SELECTOR, true, true);
    s_connectionStore.setChainPause(remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Both, true);

    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_token, DEST_CHAIN_SELECTOR);
    assertFalse(connection.isOutboundPaused);
    assertFalse(connection.isInboundPaused);
  }

  // Reverts

  function test_setChainPause_RevertWhen_UnauthorizedCaller() public {
    uint64[] memory remoteChainSelectors = _getRemoteChainSelectors(DEST_CHAIN_SELECTOR);
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_connectionStore.setChainPause(remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Both, true);
  }
}
