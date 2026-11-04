// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";
import {IRemoteTokenAddressProvider} from "../../../../../interfaces/IRemoteTokenAddressProvider.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

import {Vm} from "forge-std/Vm.sol";

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
    // A missing lane must not shift later results.
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

    // Pause state is returned even when no connection exists.
    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_otherToken, DEST_CHAIN_SELECTOR);

    assertFalse(connection.hasConnection);
    assertTrue(connection.isGloballyPaused);
    assertTrue(connection.isTokenOutboundPaused);
    assertFalse(connection.isTokenInboundPaused);
    assertFalse(connection.isChainOutboundPaused);
    assertTrue(connection.isChainInboundPaused);
    assertTrue(connection.isOutboundPaused);
    assertTrue(connection.isInboundPaused);

    // Clearing the global pause leaves the token's outbound pause and the chain's inbound pause active.
    s_connectionStore.setGlobalPause(false);
    connection = _getLaneConnection(s_otherToken, DEST_CHAIN_SELECTOR);

    assertFalse(connection.isGloballyPaused);
    assertTrue(connection.isOutboundPaused);
    assertTrue(connection.isInboundPaused);
  }

  /// @notice Provider failures leave the remote token empty without hiding the rest of the lane configuration.
  function test_getLaneConnectionBatch_ProviderGivesNoRemoteToken() public {
    address remoteTokenAddressProvider = address(s_remoteTokenAddressProvider);
    bytes memory getRemoteTokenAddressCall =
      abi.encodeCall(IRemoteTokenAddressProvider.getRemoteTokenAddress, (s_token, DEST_CHAIN_SELECTOR));
    IPoolConfigManager.LaneConnection memory expectedConnection;
    expectedConnection.hasConnection = true;
    expectedConnection.remotePools = _getRemotePoolConfig().remotePools;

    vm.mockCallRevert(remoteTokenAddressProvider, getRemoteTokenAddressCall, "");
    _assertLaneConnection(expectedConnection);

    vm.mockCall(remoteTokenAddressProvider, getRemoteTokenAddressCall, abi.encode(bytes("")));
    _assertLaneConnection(expectedConnection);

    vm.mockCall(remoteTokenAddressProvider, getRemoteTokenAddressCall, abi.encode(abi.encode(address(0))));
    _assertLaneConnection(expectedConnection);

    // The provider had code when configured, but no longer does.
    vm.clearMockedCalls();
    vm.etch(remoteTokenAddressProvider, "");
    _assertLaneConnection(expectedConnection);

    _setRemoteTokenAddressProvider(CONNECTION_POLICY, address(0));
    _assertLaneConnection(expectedConnection);
  }

  function test_getLaneConnectionBatch_SubsetToken() public {
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, OTHER_CHAIN_SELECTOR, true));
    _addSubsetToken(CONNECTION_POLICY, OTHER_CHAIN_SELECTOR, s_token);
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(REMOTE_POOL_POLICY, OTHER_CHAIN_SELECTOR, _getRemotePoolAddresses(s_remotePool))
    );

    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_token, OTHER_CHAIN_SELECTOR);

    assertTrue(connection.hasConnection);
    assertEq(abi.encode(s_token), connection.remoteToken);
    assertEq(abi.encode(_getRemotePoolConfig().remotePools), abi.encode(connection.remotePools));
  }

  function test_getLaneConnectionBatch_TrustPolicyRemotePools() public {
    _setInlineRemotePoolConfigTrustingPolicy(s_token, s_otherRemotePool, true);
    // Opting in without an assigned policy still returns only the inline pool.
    _setInlineConnectionConfig(
      _getInlineConnectionConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, abi.encode(s_otherToken))
    );
    _setInlineRemotePoolConfigTrustingPolicy(s_otherToken, s_otherRemotePool, true);

    // Inline pools precede policy pools.
    IPoolConfigManager.RemotePool[] memory expectedRemotePools = new IPoolConfigManager.RemotePool[](2);
    expectedRemotePools[0] =
      IPoolConfigManager.RemotePool({remotePoolAddress: s_otherRemotePool, allowLegacyPoolDataOnReceive: true});
    expectedRemotePools[1] =
      IPoolConfigManager.RemotePool({remotePoolAddress: s_remotePool, allowLegacyPoolDataOnReceive: false});
    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_token, DEST_CHAIN_SELECTOR);
    assertEq(abi.encode(expectedRemotePools), abi.encode(connection.remotePools));

    connection = _getLaneConnection(s_otherToken, DEST_CHAIN_SELECTOR);
    assertEq(1, connection.remotePools.length);
    assertEq(abi.encode(expectedRemotePools[0]), abi.encode(connection.remotePools[0]));

    // A pool listed in both records appears once, with its inline legacy receive permission.
    _setInlineRemotePoolConfigTrustingPolicy(s_token, s_remotePool, true);
    connection = _getLaneConnection(s_token, DEST_CHAIN_SELECTOR);
    assertEq(1, connection.remotePools.length);
    assertEq(s_remotePool, connection.remotePools[0].remotePoolAddress);
    assertTrue(connection.remotePools[0].allowLegacyPoolDataOnReceive);
  }

  function test_getLaneConnectionBatch_TrustPolicyRemotePools_PolicyRecordChange() public {
    _setInlineRemotePoolConfigTrustingPolicy(s_token, s_otherRemotePool, true);
    bytes memory newRemotePool = abi.encode(makeAddr("newRemotePool"));

    // Updating the policy changes the resolved view without emitting a token-specific event.
    vm.recordLogs();
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(newRemotePool))
    );
    Vm.Log[] memory logs = vm.getRecordedLogs();
    assertEq(1, logs.length);
    assertEq(ConnectionPolicyStore.RemotePoolPolicyConfigSet.selector, logs[0].topics[0]);

    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_token, DEST_CHAIN_SELECTOR);
    assertEq(2, connection.remotePools.length);
    assertEq(s_otherRemotePool, connection.remotePools[0].remotePoolAddress);
    assertEq(newRemotePool, connection.remotePools[1].remotePoolAddress);
  }

  function test_getLaneConnectionBatch_TrustPolicyRemotePools_PolicyChainRecord() public {
    // Configure different pools for the two chains to detect a lookup on the wrong chain.
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(
        REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(s_otherRemotePool)
      )
    );
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(REMOTE_POOL_POLICY, SOURCE_CHAIN_SELECTOR, _getRemotePoolAddresses(s_remotePool))
    );
    bytes memory legacyPool = abi.encode(makeAddr("legacyPool"));
    _setInlineRemotePoolConfigTrustingPolicy(s_token, legacyPool, true);

    // The view and receive path must use the same chain's policy pools.
    IPoolConfigManager.RemotePool[] memory expectedRemotePools = new IPoolConfigManager.RemotePool[](2);
    expectedRemotePools[0] =
      IPoolConfigManager.RemotePool({remotePoolAddress: legacyPool, allowLegacyPoolDataOnReceive: true});
    expectedRemotePools[1] =
      IPoolConfigManager.RemotePool({remotePoolAddress: s_otherRemotePool, allowLegacyPoolDataOnReceive: false});
    assertEq(abi.encode(expectedRemotePools), abi.encode(_getLaneConnection(s_token, DEST_CHAIN_SELECTOR).remotePools));
    _resolveInboundLane(s_otherRemotePool);
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_remotePool
      )
    );
    _resolveInboundLane(s_remotePool);
  }

  function test_getLaneConnectionBatch_InlineRecordReplacesPolicyRecord() public {
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePools[0].remotePoolAddress = s_otherRemotePool;
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig));

    // Without the opt-in, neither the view nor the receive path includes policy pools.
    assertEq(
      abi.encode(remotePoolConfig.remotePools), abi.encode(_getLaneConnection(s_token, DEST_CHAIN_SELECTOR).remotePools)
    );
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_remotePool
      )
    );
    _resolveInboundLane(s_remotePool);

    // Policy updates must not change the pools trusted by an overriding inline record.
    bytes memory chainRemotePool = abi.encode(makeAddr("chainRemotePool"));
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(chainRemotePool))
    );
    assertEq(
      abi.encode(remotePoolConfig.remotePools), abi.encode(_getLaneConnection(s_token, DEST_CHAIN_SELECTOR).remotePools)
    );
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, chainRemotePool
      )
    );
    _resolveInboundLane(chainRemotePool);
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

    // ABI decoding errors are not caught by try/catch.
    vm.expectRevert(bytes(""));
    s_connectionStore.getLaneConnectionBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));
  }

  function _assertLaneConnection(
    IPoolConfigManager.LaneConnection memory expectedConnection
  ) private view {
    assertEq(abi.encode(expectedConnection), abi.encode(_getLaneConnection(s_token, DEST_CHAIN_SELECTOR)));
  }

  /// @dev Validates an inbound transfer using pool data that includes the source token.
  function _resolveInboundLane(
    bytes memory sourcePoolAddress
  ) private view {
    s_connectionStore.resolveInboundLane(
      s_token,
      DEST_CHAIN_SELECTOR,
      sourcePoolAddress,
      _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""),
      1e18,
      FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
  }
}
