// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract ConnectionPolicyStore_applyInlineRemotePoolConfigUpdates is ConnectionPolicyStoreSetup {
  function test_applyInlineRemotePoolConfigUpdates() public {
    ConnectionPolicyStore.RemotePoolConfig memory destRemotePoolConfig = _getRemotePoolConfig();
    destRemotePoolConfig.remotePools[0].remotePoolAddress = s_otherRemotePool;
    destRemotePoolConfig.remotePools[0].allowLegacyPoolDataOnReceive = true;
    destRemotePoolConfig.useLegacyPoolDataOnSend = true;
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      new ConnectionPolicyStore.InlineRemotePoolConfigArgs[](2);
    inlineConfigs[0] = _getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, destRemotePoolConfig)[0];
    inlineConfigs[1] = _getInlineRemotePoolConfigArgs(s_token, SOURCE_CHAIN_SELECTOR, _getRemotePoolConfig())[0];

    vm.expectEmit();
    emit ConnectionPolicyStore.InlineRemotePoolConfigSet(s_token, DEST_CHAIN_SELECTOR, false, destRemotePoolConfig);
    vm.expectEmit();
    emit ConnectionPolicyStore.InlineRemotePoolConfigSet(s_token, SOURCE_CHAIN_SELECTOR, false, _getRemotePoolConfig());
    s_connectionStore.applyInlineRemotePoolConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);

    assertEq(abi.encode(inlineConfigs), abi.encode(s_connectionStore.getInlineRemotePoolConfigs(s_token)));
    (, PolicyConfig.RecordLevel[] memory remotePoolLevels) =
      s_connectionStore.getLaneRecordLevelBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));
    assertEq(uint8(PolicyConfig.RecordLevel.Inline), uint8(remotePoolLevels[0]));
    // Each inline record overrides the policy and trusts only its configured pools.
    s_connectionStore.resolveInboundLane(
      s_token, DEST_CHAIN_SELECTOR, s_otherRemotePool, "", 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_remotePool
      )
    );
    s_connectionStore.resolveInboundLane(
      s_token, DEST_CHAIN_SELECTOR, s_remotePool, "", 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
  }

  function test_applyInlineRemotePoolConfigUpdates_NoRemotePools() public {
    ConnectionPolicyStore.RemotePoolConfig memory emptyRemotePoolConfig;
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      _getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, emptyRemotePoolConfig);

    vm.expectEmit();
    emit ConnectionPolicyStore.InlineRemotePoolConfigSet(s_token, DEST_CHAIN_SELECTOR, false, emptyRemotePoolConfig);
    s_connectionStore.applyInlineRemotePoolConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);

    // An empty pool list still allows outbound transfers.
    IPoolConfigManager.OutboundLane memory lane =
      s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
    assertEq(_getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""), lane.destPoolData);
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_remotePool
      )
    );
    s_connectionStore.resolveInboundLane(
      s_token, DEST_CHAIN_SELECTOR, s_remotePool, sourcePoolData, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
  }

  function test_applyInlineRemotePoolConfigUpdates_ReplaceRecord() public {
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getRemotePoolConfig()));
    ConnectionPolicyStore.RemotePoolConfig memory newRemotePoolConfig = _getRemotePoolConfig();
    newRemotePoolConfig.remotePools[0].remotePoolAddress = s_otherRemotePool;
    newRemotePoolConfig.remotePools[0].allowLegacyPoolDataOnReceive = true;
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      _getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, newRemotePoolConfig);

    vm.expectEmit();
    emit ConnectionPolicyStore.InlineRemotePoolConfigSet(s_token, DEST_CHAIN_SELECTOR, false, newRemotePoolConfig);
    s_connectionStore.applyInlineRemotePoolConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);

    assertEq(abi.encode(inlineConfigs), abi.encode(s_connectionStore.getInlineRemotePoolConfigs(s_token)));
    s_connectionStore.resolveInboundLane(
      s_token, DEST_CHAIN_SELECTOR, s_otherRemotePool, "", 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_remotePool
      )
    );
    s_connectionStore.resolveInboundLane(
      s_token, DEST_CHAIN_SELECTOR, s_remotePool, "", 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );

    // Updating a pool that remains in the list must replace its legacy receive permission.
    newRemotePoolConfig.remotePools[0].allowLegacyPoolDataOnReceive = false;
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, newRemotePoolConfig));

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.LegacyPoolDataNotAllowed.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_connectionStore.resolveInboundLane(
      s_token, DEST_CHAIN_SELECTOR, s_otherRemotePool, "", 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
  }

  function test_applyInlineRemotePoolConfigUpdates_RemoveInlineRecord() public {
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePools[0].remotePoolAddress = s_otherRemotePool;
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig));

    vm.expectEmit();
    emit ConnectionPolicyStore.InlineRemotePoolConfigRemoved(s_token, DEST_CHAIN_SELECTOR);
    s_connectionStore.applyInlineRemotePoolConfigUpdates(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR), new ConnectionPolicyStore.InlineRemotePoolConfigArgs[](0)
    );

    assertEq(0, s_connectionStore.getInlineRemotePoolConfigs(s_token).length);
    // Removing the inline record restores the policy configuration.
    (, PolicyConfig.RecordLevel[] memory remotePoolLevels) =
      s_connectionStore.getLaneRecordLevelBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));
    assertEq(uint8(PolicyConfig.RecordLevel.PolicyChain), uint8(remotePoolLevels[0]));
    IPoolConfigManager.LaneConnection[] memory connections =
      s_connectionStore.getLaneConnectionBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));
    assertEq(1, connections[0].remotePools.length);
    assertEq(s_remotePool, connections[0].remotePools[0].remotePoolAddress);
  }

  function test_applyInlineRemotePoolConfigUpdates_RemoveDeletesValues() public {
    // With no policy to fall back to, removal leaves no remote pool configuration.
    _setInlineConnectionConfig(
      _getInlineConnectionConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, abi.encode(s_otherToken))
    );
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.useLegacyPoolDataOnSend = true;
    remotePoolConfig.remotePools[0].allowLegacyPoolDataOnReceive = true;
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, remotePoolConfig));
    // Removal requires the legacy flags to be cleared first.
    remotePoolConfig.useLegacyPoolDataOnSend = false;
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, remotePoolConfig));

    s_connectionStore.applyInlineRemotePoolConfigUpdates(
      _getLaneKeys(s_otherToken, DEST_CHAIN_SELECTOR), new ConnectionPolicyStore.InlineRemotePoolConfigArgs[](0)
    );

    IPoolConfigManager.LaneConnection[] memory connections =
      s_connectionStore.getLaneConnectionBatch(_getLaneKeys(s_otherToken, DEST_CHAIN_SELECTOR));
    assertTrue(connections[0].hasConnection);
    assertFalse(connections[0].useLegacyPoolDataOnSend);
    assertEq(0, connections[0].remotePools.length);
  }

  function test_applyInlineRemotePoolConfigUpdates_RemoveRecordTrustingPolicy() public {
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePools[0].remotePoolAddress = s_otherRemotePool;
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      _getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig);
    inlineConfigs[0].trustPolicyRemotePools = true;
    _setInlineRemotePoolConfig(inlineConfigs);
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");
    s_connectionStore.resolveInboundLane(
      s_token, DEST_CHAIN_SELECTOR, s_otherRemotePool, sourcePoolData, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );

    s_connectionStore.applyInlineRemotePoolConfigUpdates(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR), new ConnectionPolicyStore.InlineRemotePoolConfigArgs[](0)
    );

    // Falling back to the policy drops trust in the inline-only pool.
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_otherRemotePool
      )
    );
    s_connectionStore.resolveInboundLane(
      s_token, DEST_CHAIN_SELECTOR, s_otherRemotePool, sourcePoolData, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
  }

  function test_applyInlineRemotePoolConfigUpdates_RemovalsRunFirst() public {
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getRemotePoolConfig()));
    ConnectionPolicyStore.RemotePoolConfig memory newRemotePoolConfig = _getRemotePoolConfig();
    newRemotePoolConfig.remotePools[0].remotePoolAddress = s_otherRemotePool;
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      _getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, newRemotePoolConfig);

    s_connectionStore.applyInlineRemotePoolConfigUpdates(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR), inlineConfigs);

    assertEq(abi.encode(inlineConfigs), abi.encode(s_connectionStore.getInlineRemotePoolConfigs(s_token)));
  }

  function test_applyInlineRemotePoolConfigUpdates_TrustPolicyRemotePools() public {
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePools[0].remotePoolAddress = s_otherRemotePool;
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      _getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig);
    inlineConfigs[0].trustPolicyRemotePools = true;
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");

    vm.expectEmit();
    emit ConnectionPolicyStore.InlineRemotePoolConfigSet(s_token, DEST_CHAIN_SELECTOR, true, remotePoolConfig);
    s_connectionStore.applyInlineRemotePoolConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);

    assertEq(abi.encode(inlineConfigs), abi.encode(s_connectionStore.getInlineRemotePoolConfigs(s_token)));
    // Opting in trusts both the inline and policy pools.
    s_connectionStore.resolveInboundLane(
      s_token, DEST_CHAIN_SELECTOR, s_otherRemotePool, sourcePoolData, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    s_connectionStore.resolveInboundLane(
      s_token, DEST_CHAIN_SELECTOR, s_remotePool, sourcePoolData, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );

    // Disabling trustPolicyRemotePools removes trust in pools listed only in the policy.
    inlineConfigs[0].trustPolicyRemotePools = false;
    s_connectionStore.applyInlineRemotePoolConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);

    assertEq(abi.encode(inlineConfigs), abi.encode(s_connectionStore.getInlineRemotePoolConfigs(s_token)));
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_remotePool
      )
    );
    s_connectionStore.resolveInboundLane(
      s_token, DEST_CHAIN_SELECTOR, s_remotePool, sourcePoolData, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
  }

  // Reverts

  function test_applyInlineRemotePoolConfigUpdates_RevertWhen_ZeroAddressInvalid() public {
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePools[0].remotePoolAddress = "";
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      _getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig);

    vm.expectRevert(ConnectionPolicyStore.ZeroAddressInvalid.selector);
    s_connectionStore.applyInlineRemotePoolConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function test_applyInlineRemotePoolConfigUpdates_RevertWhen_RemotePoolAlreadyAdded() public {
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePools = new IPoolConfigManager.RemotePool[](3);
    remotePoolConfig.remotePools[0].remotePoolAddress = s_remotePool;
    remotePoolConfig.remotePools[1].remotePoolAddress = s_otherRemotePool;
    // Different legacy flags do not make duplicate pool addresses valid.
    remotePoolConfig.remotePools[2].remotePoolAddress = s_remotePool;
    remotePoolConfig.remotePools[2].allowLegacyPoolDataOnReceive = true;
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      _getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig);

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.RemotePoolAlreadyAdded.selector, DEST_CHAIN_SELECTOR, s_remotePool)
    );
    s_connectionStore.applyInlineRemotePoolConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function test_applyInlineRemotePoolConfigUpdates_RevertWhen_ZeroChainSelectorNotAllowed() public {
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      _getInlineRemotePoolConfigArgs(s_token, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _getRemotePoolConfig());

    vm.expectRevert(PolicyConfig.ZeroChainSelectorNotAllowed.selector);
    s_connectionStore.applyInlineRemotePoolConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function test_applyInlineRemotePoolConfigUpdates_RevertWhen_LocalTokenNotRegistered() public {
    address unknownToken = makeAddr("unknownToken");
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      _getInlineRemotePoolConfigArgs(unknownToken, DEST_CHAIN_SELECTOR, _getRemotePoolConfig());

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, unknownToken));
    s_connectionStore.applyInlineRemotePoolConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function test_applyInlineRemotePoolConfigUpdates_RevertWhen_InlineRecordNotFound() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.InlineRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR));
    s_connectionStore.applyInlineRemotePoolConfigUpdates(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR), new ConnectionPolicyStore.InlineRemotePoolConfigArgs[](0)
    );
  }

  /// @dev Policy records have no legacy flags, so removing a legacy record would switch the lane to long pool data.
  function test_applyInlineRemotePoolConfigUpdates_RevertWhen_LegacyRemotePoolRecord() public {
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePools = new IPoolConfigManager.RemotePool[](2);
    remotePoolConfig.remotePools[0].remotePoolAddress = s_remotePool;
    remotePoolConfig.remotePools[1].remotePoolAddress = s_otherRemotePool;
    remotePoolConfig.remotePools[1].allowLegacyPoolDataOnReceive = true;
    remotePoolConfig.useLegacyPoolDataOnSend = true;
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig));

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.LegacyRemotePoolRecord.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_connectionStore.applyInlineRemotePoolConfigUpdates(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR), new ConnectionPolicyStore.InlineRemotePoolConfigArgs[](0)
    );

    // Clearing the send flag allows removal, even with a receive flag still set.
    remotePoolConfig.useLegacyPoolDataOnSend = false;
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig));
    s_connectionStore.applyInlineRemotePoolConfigUpdates(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR), new ConnectionPolicyStore.InlineRemotePoolConfigArgs[](0)
    );
    assertEq(0, s_connectionStore.getInlineRemotePoolConfigs(s_token).length);
  }

  function test_applyInlineRemotePoolConfigUpdates_RevertWhen_UnauthorizedCaller() public {
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      _getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getRemotePoolConfig());
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_connectionStore.applyInlineRemotePoolConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }
}
