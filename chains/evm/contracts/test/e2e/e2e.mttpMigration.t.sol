// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../libraries/PolicyConfig.sol";
import {BurnMintTokenPool} from "../../pools/BurnMintTokenPool.sol";
import {ConnectionPolicyStore} from "../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {FeePolicyStore} from "../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {e2e_mttpSetup} from "./e2e.mttpSetup.t.sol";

/// @notice Migration behavior: moving from normal single-token pools to MTTP in stages. Covers retained transitional
/// messages, the legacy pool data formats, remote trust repair, retry with realistic gas, and policy adoption after
/// migration. The source and destination use a legacy BurnMintTokenPool next to the MTTP.
contract e2e_mttpMigration is e2e_mttpSetup {
  /// @dev Normal source to migrated destination. A message produced before A switches to the MTTP and one
  /// produced after both execute against whatever pool A's token admin registry points at.
  function test_mttp_normalSourceToMigratedDestination() public {
    address legacyT1B;
    vm.startPrank(OWNER);
    legacyT1B = _deployLegacyPoolOn(s_b, s_a, 0);
    // A's MTTP trusts the legacy source pool and accepts its legacy pool data.
    _setMigrationRemotePool(s_a, 0, s_b.localSelector, legacyT1B, false, true);
    // A's T1 is still served by a normal pool until the switch.
    _deployLegacyPoolOn(s_a, s_b, 0);
    vm.stopPrank();

    // A message produced before A's switch, deferred.
    (bytes memory retainedData, uint256 retainedAmount) = _legacyLockOrBurn(legacyT1B, s_b, s_a, 0, 100e18);
    assertEq(retainedData.length, LEGACY_POOL_DATA_LENGTH, "legacy pool data");

    // Switch A to the MTTP.
    vm.prank(OWNER);
    _switchPool(s_a, 0, address(s_a.pool));

    // The retained message executes against A's current TAR pool.
    uint256 received = _deliverFromPool(
      s_a,
      s_b.localSelector,
      legacyT1B,
      s_a.tokens[0],
      retainedAmount,
      retainedData,
      FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    assertEq(received, 100e18, "retained message succeeds after switch");

    // A message produced after the switch also succeeds.
    (bytes memory newData, uint256 newAmount) = _legacyLockOrBurn(legacyT1B, s_b, s_a, 0, 50e18);
    received = _deliverFromPool(
      s_a, s_b.localSelector, legacyT1B, s_a.tokens[0], newAmount, newData, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    assertEq(received, 50e18, "new message succeeds after switch");

    // T2/T3 traffic is unchanged.
    assertEq(_transfer(s_a, s_b, 1, 10e18), 10e18, "T2 unchanged");
    assertEq(_transfer(s_a, s_b, 2, 10e18), 10e18, "T3 unchanged");
  }

  /// @dev Migrated source to normal destination. A's MTTP sends the 32-byte legacy pool data to a legacy
  /// pool on B. Removing remote trust blocks the receive, restoring it lets the same message through, and a
  /// long-format send under a 32-byte cap reverts before any burn.
  function test_mttp_migratedSourceToNormalDestination() public {
    address legacyT1B;
    vm.startPrank(OWNER);
    legacyT1B = _deployLegacyPoolOn(s_b, s_a, 0);
    // A's MTTP sends legacy data and accepts legacy data from B's legacy pool.
    _setMigrationRemotePool(s_a, 0, s_b.localSelector, legacyT1B, true, true);
    vm.stopPrank();

    // A MTTP -> B legacy pool. The pool data is the 32-byte legacy format.
    (bytes memory poolData, uint256 postFee,) =
      _send(s_a, s_b, 0, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    assertEq(poolData.length, LEGACY_POOL_DATA_LENGTH, "legacy 32-byte pool data");
    uint256 received = _legacyReleaseOrMint(legacyT1B, s_b, s_a.localSelector, address(s_a.pool), 0, postFee, poolData);
    assertEq(received, 100e18, "normal destination receives correct amount");

    // Remove remote trust: the receive fails.
    vm.prank(OWNER);
    BurnMintTokenPool(legacyT1B).removeRemotePool(s_a.localSelector, abi.encode(address(s_a.pool)));
    (poolData, postFee,) = _send(s_a, s_b, 0, 10e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    vm.expectRevert();
    _legacyReleaseOrMint(legacyT1B, s_b, s_a.localSelector, address(s_a.pool), 0, postFee, poolData);

    // Restore trust and retry the same message.
    vm.prank(OWNER);
    BurnMintTokenPool(legacyT1B).addRemotePool(s_a.localSelector, abi.encode(address(s_a.pool)));
    received = _legacyReleaseOrMint(legacyT1B, s_b, s_a.localSelector, address(s_a.pool), 0, postFee, poolData);
    assertEq(received, 10e18, "retry after trust repair");

    // Attempt a long send while retaining the 32-byte cap. It reverts before a committed burn or fee.
    vm.startPrank(OWNER);
    _setMigrationRemotePool(s_a, 0, s_b.localSelector, legacyT1B, false, true);
    _setInlineDestOverheadConfig(s_a.pcm, s_a.tokens[0], s_b.localSelector, DEFAULT_TOKEN_DEST_GAS_OVERHEAD, 32);
    vm.stopPrank();

    _fundPool(s_a, s_a.tokens[0], 100e18);
    uint256 poolBefore = _balance(s_a.tokens[0], address(s_a.pool));
    vm.expectRevert(
      abi.encodeWithSelector(
        FeePolicyStore.PoolDataTooLarge.selector, s_a.tokens[0], s_b.localSelector, LONG_POOL_DATA_LENGTH, 32
      )
    );
    _lockOrBurn(s_a, s_b, 0, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    assertEq(_balance(s_a.tokens[0], address(s_a.pool)), poolBefore, "no burn or fee on capped send");
  }

  /// @dev Finishes the migration. A retained legacy message executes once after B switches, both long directions
  /// succeed, and the temporary legacy receive permission is only turned off after the short-format cohort drains.
  function test_mttp_finishMigrationAndDrainTransitionalData() public {
    address legacyT1B;
    vm.startPrank(OWNER);
    legacyT1B = _deployLegacyPoolOn(s_b, s_a, 0);
    // A's MTTP sends legacy to B's legacy pool and accepts legacy from it.
    _setMigrationRemotePool(s_a, 0, s_b.localSelector, legacyT1B, true, true);
    vm.stopPrank();

    // Retain a legacy-format message from A's MTTP before B switches.
    (bytes memory retainedData, uint256 retainedAmount,) =
      _send(s_a, s_b, 0, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    assertEq(retainedData.length, LEGACY_POOL_DATA_LENGTH, "retained legacy pool data");

    // Switch B to the MTTP. B sends long and still accepts the retained short data from A's MTTP.
    vm.startPrank(OWNER);
    _switchPool(s_b, 0, address(s_b.pool));
    _setMigrationRemotePool(s_b, 0, s_a.localSelector, address(s_a.pool), false, true);
    // A now trusts B's MTTP for B's long sends, and keeps sending legacy until it switches.
    _setMigrationRemotePool(s_a, 0, s_b.localSelector, address(s_b.pool), true, true);
    vm.stopPrank();

    // The retained A->B short message executes once against B's MTTP.
    uint256 received = _deliverFromPool(
      s_b,
      s_a.localSelector,
      address(s_a.pool),
      s_b.tokens[0],
      retainedAmount,
      retainedData,
      FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    assertEq(received, 100e18, "retained short message succeeds once");

    // B->A long send succeeds at B's switch.
    (bytes memory bData, uint256 bAmount,) =
      _send(s_b, s_a, 0, 50e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    assertEq(bData.length, LONG_POOL_DATA_LENGTH, "B sends long format");
    received = _deliverFromPool(
      s_a, s_b.localSelector, address(s_b.pool), s_a.tokens[0], bAmount, bData, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    assertEq(received, 50e18, "B->A long send succeeds");

    // Switch A to long format.
    vm.prank(OWNER);
    _setMigrationRemotePool(s_a, 0, s_b.localSelector, address(s_b.pool), false, true);

    (bytes memory aData, uint256 aAmount,) =
      _send(s_a, s_b, 0, 50e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    assertEq(aData.length, LONG_POOL_DATA_LENGTH, "A sends long format");
    received = _deliverFromPool(
      s_b, s_a.localSelector, address(s_a.pool), s_b.tokens[0], aAmount, aData, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    assertEq(received, 50e18, "A->B long send succeeds");

    // Drain the short-format cohort, then disable B's legacy receive for A's MTTP.
    vm.prank(OWNER);
    _setMigrationRemotePool(s_b, 0, s_a.localSelector, address(s_a.pool), false, false);

    // Produce one more legacy message by temporarily reverting A to legacy sends, then disable again.
    vm.prank(OWNER);
    _setMigrationRemotePool(s_a, 0, s_b.localSelector, address(s_b.pool), true, true);
    (bytes memory leftoverData, uint256 leftoverAmount,) =
      _send(s_a, s_b, 0, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    vm.prank(OWNER);
    _setMigrationRemotePool(s_a, 0, s_b.localSelector, address(s_b.pool), false, true);

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.LegacyPoolDataNotAllowed.selector, s_b.tokens[0], s_a.localSelector)
    );
    _deliverFromPool(
      s_b,
      s_a.localSelector,
      address(s_a.pool),
      s_b.tokens[0],
      leftoverAmount,
      leftoverData,
      FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );

    // Historical normal-pool receive entries are retained: B's legacy pool still accepts A's MTTP.
    uint256 legacyReceived =
      _legacyReleaseOrMint(legacyT1B, s_b, s_a.localSelector, address(s_a.pool), 0, leftoverAmount, leftoverData);
    assertEq(legacyReceived, 1e18, "legacy pool retains historical receive");
  }

  /// @dev Adopts policies after migration while keeping an inline remote pool record that still trusts the old
  /// legacy source. Removing the inline rate limit resets the bucket to the policy configuration.
  function test_mttp_adoptPoliciesAfterMigration() public {
    address legacyT1B;
    vm.startPrank(OWNER);
    legacyT1B = _deployLegacyPoolOn(s_b, s_a, 0);
    // Legacy pool is the old source; B's live T1 pool is the MTTP here.
    _switchPool(s_b, 0, address(s_b.pool));

    // T1 keeps an inline remote pool record that trusts the old legacy pool and also trusts the policy pools.
    IPoolConfigManager.RemotePool[] memory pools = new IPoolConfigManager.RemotePool[](1);
    pools[0] =
      IPoolConfigManager.RemotePool({remotePoolAddress: abi.encode(legacyT1B), allowLegacyPoolDataOnReceive: true});
    ConnectionPolicyStore.RemotePoolConfig memory config =
      ConnectionPolicyStore.RemotePoolConfig({useLegacyPoolDataOnSend: false, remotePools: pools});
    _setInlineRemotePoolConfig(s_a.pcm, s_a.tokens[0], s_b.localSelector, true, config);

    // Assign policies before removing overrides. T1 has an inline rate limit and a policy rate limit to fall back to.
    _setRateLimitPolicyConfig(s_a.pcm, RATE_LIMIT_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _rateLimit(100e18));
    _setInlineRateLimitConfig(s_a.pcm, s_a.tokens[0], s_b.localSelector, _rateLimit(50e18));
    vm.stopPrank();

    // A retained message from T1's old normal source, produced before adoption.
    (bytes memory retainedData, uint256 retainedAmount) = _legacyLockOrBurn(legacyT1B, s_b, s_a, 0, 40e18);

    // Transfer T1 during adoption, consuming the inline bucket.
    (uint256 received,) = _sendAndDeliver(s_a, s_b, 0, 40e18);
    assertEq(received, 40e18, "T1 transfers during adoption");

    // Drop the inline rate limit; bucket resets to the policy config.
    vm.prank(OWNER);
    _removeInlineRateLimit(s_a.pcm, s_a.tokens[0], s_b.localSelector);
    IPoolConfigManager.LaneRateLimiterState memory state =
      _getRateLimiterState(s_a.pcm, s_a.tokens[0], s_b.localSelector);
    assertEq(state.outbound.capacity, 100e18, "bucket reset to policy capacity");
    assertEq(state.outbound.tokens, 100e18, "bucket refilled to policy capacity");

    // T2 and T3 are controls.
    assertEq(_transfer(s_a, s_b, 1, 10e18), 10e18, "T2 control transfers");
    assertEq(_transfer(s_a, s_b, 2, 10e18), 10e18, "T3 control transfers");

    // Retained message from the old source still runs through the inline legacy entry.
    received = _deliverFromPool(
      s_a,
      s_b.localSelector,
      legacyT1B,
      s_a.tokens[0],
      retainedAmount,
      retainedData,
      FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    assertEq(received, 40e18, "historical legacy receive works after adoption");

    // A new message from the policy-trusted remote MTTP works.
    (bytes memory newData, uint256 newAmount,) =
      _send(s_b, s_a, 0, 10e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    received = _deliverFromPool(
      s_a, s_b.localSelector, address(s_b.pool), s_a.tokens[0], newAmount, newData, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    assertEq(received, 10e18, "policy-trusted remote MTTP works");
  }

  /// @dev Retries a retained message with realistic gas. The message fails under a low gas limit, then succeeds
  /// on retry and pays out once. The configured destination overhead is a price input, not a cap on the pool call.
  function test_mttp_retryWithRealisticGas() public {
    address legacyT1B;
    vm.startPrank(OWNER);
    legacyT1B = _deployLegacyPoolOn(s_b, s_a, 0);
    _setMigrationRemotePool(s_a, 0, s_b.localSelector, legacyT1B, false, true);
    vm.stopPrank();

    // A source message produced before the cutover, deferred.
    (bytes memory poolData, uint256 postFee) = _legacyLockOrBurn(legacyT1B, s_b, s_a, 0, 100e18);

    // A low-gas execution attempt fails and commits no mint.
    uint256 receiverBefore = _balance(s_a.tokens[0], s_receiver);
    bool success = _tryDeliver(
      s_a, s_b.localSelector, legacyT1B, s_a.tokens[0], postFee, poolData, FinalityCodec.WAIT_FOR_FINALITY_FLAG, 30_000
    );
    assertFalse(success, "low gas attempt fails");
    assertEq(_balance(s_a.tokens[0], s_receiver), receiverBefore, "no mint on low gas");

    // Retry the same message with sufficient gas.
    uint256 gasBefore = gasleft();
    uint256 received = _deliverFromPool(
      s_a, s_b.localSelector, legacyT1B, s_a.tokens[0], postFee, poolData, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    uint256 gasUsed = gasBefore - gasleft();
    assertEq(received, 100e18, "retry succeeds");
    assertEq(_balance(s_a.tokens[0], s_receiver), receiverBefore + 100e18, "one payout only");
    assertGt(gasUsed, 0, "receive gas measured");
    emit log_named_uint("measured receive gas", gasUsed);

    // The configured overhead is a price input, not a cap: the lane still charges the configured value.
    IPoolConfigManager.LaneFee memory fee = _fee(s_a.pcm, s_a.tokens[0], s_b.localSelector);
    assertEq(fee.destGasOverhead, DEFAULT_TOKEN_DEST_GAS_OVERHEAD, "configured overhead is a price input");
  }
}
