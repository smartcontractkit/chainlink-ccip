// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";
import {IRemoteTokenAddressProvider} from "../../../../../interfaces/IRemoteTokenAddressProvider.sol";

import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_getLaneConnectionBatch is ConnectionPolicyStoreSetup {
  function test_getLaneConnectionBatch() public {
    bytes memory inlineRemoteToken = abi.encode(makeAddr("inlineRemoteToken"));
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, inlineRemoteToken));
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.useLegacyPoolDataOnSend = true;
    remotePoolConfig.remotePools = new IPoolConfigManager.RemotePool[](2);
    remotePoolConfig.remotePools[0].remotePoolAddress = s_remotePool;
    remotePoolConfig.remotePools[1].remotePoolAddress = abi.encode(makeAddr("legacyPool"));
    remotePoolConfig.remotePools[1].allowLegacyPoolDataOnReceive = true;
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, remotePoolConfig));

    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](3);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: DEST_CHAIN_SELECTOR});
    // A lane that does not exist keeps the later lanes in place.
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: SOURCE_CHAIN_SELECTOR});
    lanes[2] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: DEST_CHAIN_SELECTOR});

    IPoolConfigManager.LaneConnection[] memory connections = s_connectionStore.getLaneConnectionBatch(lanes);

    assertEq(3, connections.length);
    IPoolConfigManager.LaneConnection memory expectedConnection;
    expectedConnection.hasConnection = true;
    expectedConnection.remoteToken = abi.encode(s_token);
    expectedConnection.remotePools = _getRemotePoolConfig().remotePools;
    assertEq(abi.encode(expectedConnection), abi.encode(connections[0]));

    IPoolConfigManager.LaneConnection memory noConnection;
    assertEq(abi.encode(noConnection), abi.encode(connections[1]));

    expectedConnection.remoteToken = inlineRemoteToken;
    expectedConnection.useLegacyPoolDataOnSend = true;
    expectedConnection.remotePools = remotePoolConfig.remotePools;
    assertEq(abi.encode(expectedConnection), abi.encode(connections[2]));
  }

  function test_getLaneConnectionBatch_PauseState() public {
    s_connectionStore.setGlobalPause(true);
    s_connectionStore.setTokenPause(_arrayOf(s_otherToken), ConnectionPolicyStore.PauseDirection.Outbound, true);
    uint64[] memory remoteChainSelectors = _getRemoteChainSelectors(DEST_CHAIN_SELECTOR);
    s_connectionStore.setChainPause(remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Inbound, true);

    // s_otherToken has no lane, and its pause state still shows.
    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_otherToken, DEST_CHAIN_SELECTOR);

    assertFalse(connection.hasConnection);
    assertTrue(connection.isGloballyPaused);
    assertTrue(connection.isTokenOutboundPaused);
    assertFalse(connection.isTokenInboundPaused);
    assertFalse(connection.isChainOutboundPaused);
    assertTrue(connection.isChainInboundPaused);
    assertTrue(connection.isOutboundPaused);
    assertTrue(connection.isInboundPaused);

    // Each flag alone sets the effective pause of its direction.
    s_connectionStore.setGlobalPause(false);
    connection = _getLaneConnection(s_otherToken, DEST_CHAIN_SELECTOR);

    assertFalse(connection.isGloballyPaused);
    assertTrue(connection.isOutboundPaused);
    assertTrue(connection.isInboundPaused);
  }

  /// @notice A provider failure leaves the remote token empty, and the other values of the lane still show.
  function test_getLaneConnectionBatch_ProviderGivesNoRemoteToken() public {
    address remoteTokenAddressProvider = address(s_remoteTokenAddressProvider);
    bytes memory getRemoteTokenAddressCall =
      abi.encodeCall(IRemoteTokenAddressProvider.getRemoteTokenAddress, (s_token, DEST_CHAIN_SELECTOR));
    IPoolConfigManager.LaneConnection memory expectedConnection;
    expectedConnection.hasConnection = true;
    expectedConnection.remotePools = _getRemotePoolConfig().remotePools;

    // The provider reverts.
    vm.mockCallRevert(remoteTokenAddressProvider, getRemoteTokenAddressCall, "");
    _assertLaneConnection(expectedConnection);

    // The provider gives empty bytes.
    vm.mockCall(remoteTokenAddressProvider, getRemoteTokenAddressCall, abi.encode(bytes("")));
    _assertLaneConnection(expectedConnection);

    // The provider gives an all-zero answer.
    vm.mockCall(remoteTokenAddressProvider, getRemoteTokenAddressCall, abi.encode(abi.encode(address(0))));
    _assertLaneConnection(expectedConnection);

    // The provider has no code.
    _setRemoteTokenAddressProvider(CONNECTION_POLICY, makeAddr("providerWithoutCode"));
    _assertLaneConnection(expectedConnection);

    // The policy has no provider.
    _setRemoteTokenAddressProvider(CONNECTION_POLICY, address(0));
    _assertLaneConnection(expectedConnection);
  }

  function test_getLaneConnectionBatch_SubsetToken() public {
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, OTHER_CHAIN_SELECTOR, true));
    _addSubsetToken(CONNECTION_POLICY, OTHER_CHAIN_SELECTOR, s_token);

    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_token, OTHER_CHAIN_SELECTOR);

    assertTrue(connection.hasConnection);
    assertEq(abi.encode(s_token), connection.remoteToken);
    assertEq(abi.encode(_getRemotePoolConfig().remotePools), abi.encode(connection.remotePools));
  }

  function test_getLaneConnectionBatch_NoRemotePoolRecord() public {
    _setTokenRemotePoolPolicy(s_token, bytes32(0));

    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_token, DEST_CHAIN_SELECTOR);

    assertTrue(connection.hasConnection);
    assertEq(abi.encode(s_token), connection.remoteToken);
    assertEq(0, connection.remotePools.length);
  }

  // Reverts

  function test_getLaneConnectionBatch_RevertWhen_ProviderAnswerDoesNotDecode() public {
    vm.mockCall(
      address(s_remoteTokenAddressProvider),
      abi.encodeCall(IRemoteTokenAddressProvider.getRemoteTokenAddress, (s_token, DEST_CHAIN_SELECTOR)),
      hex"01"
    );

    // try does not catch a failed decode of the answer.
    vm.expectRevert(bytes(""));
    s_connectionStore.getLaneConnectionBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));
  }

  function _assertLaneConnection(
    IPoolConfigManager.LaneConnection memory expectedConnection
  ) private view {
    assertEq(abi.encode(expectedConnection), abi.encode(_getLaneConnection(s_token, DEST_CHAIN_SELECTOR)));
  }
}
