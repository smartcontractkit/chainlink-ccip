// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../interfaces/IPoolV2.sol";

import {FeeQuoter} from "../../FeeQuoter.sol";
import {IPoolConfigManager} from "../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../libraries/FinalityCodec.sol";
import {Internal} from "../../libraries/Internal.sol";
import {PolicyConfig} from "../../libraries/PolicyConfig.sol";
import {BurnMintTokenPool} from "../../pools/BurnMintTokenPool.sol";
import {ConnectionPolicyStore} from "../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {FeePolicyStore} from "../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {e2e_mttpSetup} from "./e2e.mttpSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

/// @notice Migration behavior: moving from normal single-token pools to MTTP in stages. Covers retained transitional
/// messages, the legacy pool data formats, remote trust repair, retry with realistic gas, and policy adoption after
/// migration. The source and destination use a legacy BurnMintTokenPool next to the MTTP.
contract e2e_mttpMigration is e2e_mttpSetup {
  struct RetainedLegacyTransfer {
    address sourcePool;
    bytes poolData;
    uint256 amount;
  }

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

    _approvePool(s_a, s_a.tokens[0], 100e18);
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
    RetainedLegacyTransfer memory fromA;
    RetainedLegacyTransfer memory fromB;
    vm.startPrank(OWNER);
    fromA.sourcePool = _deployLegacyPoolOn(s_a, s_b, 0);
    fromB.sourcePool = _deployLegacyPoolOn(s_b, s_a, 0);
    BurnMintTokenPool(fromA.sourcePool).addRemotePool(s_b.localSelector, abi.encode(fromB.sourcePool));
    BurnMintTokenPool(fromB.sourcePool).addRemotePool(s_a.localSelector, abi.encode(fromA.sourcePool));
    vm.stopPrank();

    // Both normal-pool transfers are created while their source TARs still select those pools.
    (fromA.poolData, fromA.amount) = _legacyLockOrBurn(fromA.sourcePool, s_a, s_b, 0, 1e18);
    (fromB.poolData, fromB.amount) = _legacyLockOrBurn(fromB.sourcePool, s_b, s_a, 0, 1e18);

    vm.startPrank(OWNER);
    _switchPool(s_a, 0, address(s_a.pool));
    _revokeMintAndBurnRoles(s_a.tokens[0], fromA.sourcePool);
    _setMigrationRemotePool(s_a, 0, s_b.localSelector, fromB.sourcePool, true, true);
    _setInlineDestOverheadConfig(s_a.pcm, s_a.tokens[0], s_b.localSelector, DEFAULT_TOKEN_DEST_GAS_OVERHEAD, 32);
    vm.stopPrank();

    // Retain A's transitional MTTP message while B is still on its normal pool.
    (bytes memory retainedData, uint256 retainedAmount,) =
      _send(s_a, s_b, 0, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    assertEq(retainedData.length, LEGACY_POOL_DATA_LENGTH, "retained legacy pool data");

    vm.startPrank(OWNER);
    _switchPool(s_b, 0, address(s_b.pool));
    _revokeMintAndBurnRoles(s_b.tokens[0], fromB.sourcePool);
    IPoolConfigManager.RemotePool[] memory pools = _legacyThenMttpPools(fromA.sourcePool, address(s_a.pool));
    pools[1].allowLegacyPoolDataOnReceive = true;
    _setMigrationRemotePools(s_b, 0, s_a.localSelector, false, pools);
    _setMigrationRemotePools(s_a, 0, s_b.localSelector, true, _legacyThenMttpPools(fromB.sourcePool, address(s_b.pool)));
    _setInlineDestOverheadConfig(
      s_b.pcm, s_b.tokens[0], s_a.localSelector, DEFAULT_TOKEN_DEST_GAS_OVERHEAD, MIN_DEST_BYTES_OVERHEAD
    );
    vm.stopPrank();

    assertEq(
      _deliver(s_b, s_a, s_b.tokens[0], retainedAmount, retainedData, FinalityCodec.WAIT_FOR_FINALITY_FLAG),
      100e18,
      "retained short message settles after B switches"
    );
    (bytes memory bData, uint256 bAmount,) =
      _send(s_b, s_a, 0, 50e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    assertEq(bData.length, LONG_POOL_DATA_LENGTH, "B sends long format at its switch");
    assertEq(
      _deliver(s_a, s_b, s_a.tokens[0], bAmount, bData, FinalityCodec.WAIT_FOR_FINALITY_FLAG),
      50e18,
      "B->A long send succeeds"
    );

    vm.startPrank(OWNER);
    // Raise the byte allowance before enabling long sends.
    _setInlineDestOverheadConfig(
      s_a.pcm, s_a.tokens[0], s_b.localSelector, DEFAULT_TOKEN_DEST_GAS_OVERHEAD, MIN_DEST_BYTES_OVERHEAD
    );
    _setMigrationRemotePools(
      s_a, 0, s_b.localSelector, false, _legacyThenMttpPools(fromB.sourcePool, address(s_b.pool))
    );
    // The transitional cohort is drained. Keep normal-pool history while removing only the MTTP exception.
    _setMigrationRemotePools(
      s_b, 0, s_a.localSelector, false, _legacyThenMttpPools(fromA.sourcePool, address(s_a.pool))
    );
    vm.stopPrank();

    assertEq(_transfer(s_a, s_b, 0, 50e18), 50e18, "A->B long send succeeds after cleanup");

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.LegacyPoolDataNotAllowed.selector, s_b.tokens[0], s_a.localSelector)
    );
    _deliver(s_b, s_a, s_b.tokens[0], retainedAmount, retainedData, FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    // These messages were genuinely in flight before cutover. Retired pools already lost their token roles.
    assertEq(
      _deliverFromPool(
        s_a,
        s_b.localSelector,
        fromB.sourcePool,
        s_a.tokens[0],
        fromB.amount,
        fromB.poolData,
        FinalityCodec.WAIT_FOR_FINALITY_FLAG
      ),
      1e18,
      "historical B->A normal-pool message settles after role revocation"
    );
    assertEq(
      _deliverFromPool(
        s_b,
        s_a.localSelector,
        fromA.sourcePool,
        s_b.tokens[0],
        fromA.amount,
        fromA.poolData,
        FinalityCodec.WAIT_FOR_FINALITY_FLAG
      ),
      1e18,
      "historical A->B normal-pool message survives MTTP exception cleanup"
    );
  }

  /// @dev Adopts policies after migration while keeping an inline remote pool record that still trusts the old
  /// legacy source. Removing the inline rate limit resets the bucket to the policy configuration.
  function test_mttp_adoptPoliciesAfterMigration() public {
    address legacyT1B;
    vm.startPrank(OWNER);
    legacyT1B = _deployLegacyPoolOn(s_b, s_a, 0);
    vm.stopPrank();

    // A message from T1's old normal source, produced and saved while it is still B's registered pool.
    (bytes memory retainedData, uint256 retainedAmount) = _legacyLockOrBurn(legacyT1B, s_b, s_a, 0, 40e18);

    vm.startPrank(OWNER);
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

  /// @dev Uses the active legacy pool's FeeQuoter fallback, captured before the transfer, for both gas budgets.
  /// This is pool-call gas recovery; OffRamp execution state and replay protection are covered separately.
  function test_mttp_retryWithRealisticGas() public {
    _assertRetryWithLegacyQuote(false);
  }

  function test_mttp_retryWithCustomLegacyPoolQuote() public {
    _assertRetryWithLegacyQuote(true);
  }

  function _assertRetryWithLegacyQuote(
    bool customPoolFee
  ) internal {
    vm.startPrank(OWNER);
    address legacyT1B = _deployLegacyPoolOn(s_b, s_a, 0);
    address legacyT1A = _deployLegacyPoolOn(s_a, s_b, 0);
    _setMigrationRemotePool(s_a, 0, s_b.localSelector, legacyT1B, false, true);
    FeeQuoter feeQuoter = _deployLegacyFeeQuoter();
    if (customPoolFee) {
      TokenPool.TokenTransferFeeConfigArgs[] memory configs = new TokenPool.TokenTransferFeeConfigArgs[](1);
      configs[0] = TokenPool.TokenTransferFeeConfigArgs({
        destChainSelector: s_a.localSelector,
        tokenTransferFeeConfig: IPoolV2.TokenTransferFeeConfig({
          finalityFeeUSDCents: 1,
          fastFinalityFeeUSDCents: 1,
          finalityTransferFeeBps: 0,
          fastFinalityTransferFeeBps: 0,
          destGasOverhead: 85_000,
          destBytesOverhead: 32,
          isEnabled: true
        })
      });
      BurnMintTokenPool(legacyT1B).applyTokenTransferFeeConfigUpdates(configs, new uint64[](0));
    }
    vm.stopPrank();

    // OnRamp resolves the TAR pool first, then uses its custom fee or the FeeQuoter fallback.
    assertEq(s_b.tar.getPool(s_b.tokens[0]), legacyT1B, "legacy source is still registered");
    uint32 quotedGas = _quotedDestGasOverhead(s_b, 0, 100e18, feeQuoter);
    assertEq(quotedGas, customPoolFee ? 85_000 : DEFAULT_TOKEN_DEST_GAS_OVERHEAD, "active legacy lane quote");

    // An MTTP-only config change must not alter a quote while the legacy pool is registered.
    vm.prank(OWNER);
    _setInlineDestOverheadConfig(s_b.pcm, s_b.tokens[0], s_a.localSelector, 1_000_000, MIN_DEST_BYTES_OVERHEAD);
    assertEq(_quotedDestGasOverhead(s_b, 0, 100e18, feeQuoter), quotedGas, "inactive MTTP quote is ignored");

    (bytes memory poolData, uint256 postFee) = _legacyLockOrBurn(legacyT1B, s_b, s_a, 0, 100e18);
    vm.startPrank(OWNER);
    _switchPool(s_a, 0, address(s_a.pool));
    _revokeMintAndBurnRoles(s_a.tokens[0], legacyT1A);
    vm.stopPrank();
    uint256 receiverBefore = _balance(s_a.tokens[0], s_receiver);
    uint256 supplyBefore = IERC20(s_a.tokens[0]).totalSupply();
    bool success = _tryDeliver(
      s_a,
      s_b.localSelector,
      legacyT1B,
      s_a.tokens[0],
      postFee,
      poolData,
      FinalityCodec.WAIT_FOR_FINALITY_FLAG,
      quotedGas
    );
    assertFalse(success, "the legacy quoted overhead alone cannot execute the MTTP receive");
    assertEq(_balance(s_a.tokens[0], s_receiver), receiverBefore, "no payout on failed attempt");
    assertEq(IERC20(s_a.tokens[0]).totalSupply(), supplyBefore, "no mint on failed attempt");

    uint256 gasBefore = gasleft();
    success = _tryDeliver(
      s_a,
      s_b.localSelector,
      legacyT1B,
      s_a.tokens[0],
      postFee,
      poolData,
      FinalityCodec.WAIT_FOR_FINALITY_FLAG,
      uint256(quotedGas) * 2
    );
    uint256 gasUsed = gasBefore - gasleft();
    assertTrue(success, "retry with twice the captured quote succeeds");
    assertEq(_balance(s_a.tokens[0], s_receiver), receiverBefore + postFee, "one successful payout");
    assertEq(IERC20(s_a.tokens[0]).totalSupply(), supplyBefore + postFee, "one successful mint");
    emit log_named_uint("measured receive gas including test harness overhead", gasUsed);
  }

  /// @dev A real FeeQuoter with B->A lane defaults; no token-specific override.
  function _deployLegacyFeeQuoter() internal returns (FeeQuoter) {
    FeeQuoter.DestChainConfigArgs[] memory lanes = new FeeQuoter.DestChainConfigArgs[](1);
    FeeQuoter.DestChainConfig memory config;
    config.isEnabled = true;
    config.defaultTokenDestGasOverhead = DEFAULT_TOKEN_DEST_GAS_OVERHEAD;
    config.defaultTxGasLimit = GAS_LIMIT;
    config.maxPerMsgGasLimit = 3_000_000;
    config.linkFeeMultiplierPercent = 100;
    config.chainFamilySelector = Internal.CHAIN_FAMILY_SELECTOR_EVM;
    lanes[0] = FeeQuoter.DestChainConfigArgs({destChainSelector: s_a.localSelector, destChainConfig: config});
    return new FeeQuoter(
      FeeQuoter.StaticConfig({maxFeeJuelsPerMsg: 1e18, linkToken: s_b.tokens[0]}),
      new address[](0),
      new FeeQuoter.TokenTransferFeeConfigArgs[](0),
      lanes
    );
  }

  /// @dev Mirrors OnRamp's pool-first, FeeQuoter-second selection for these IPoolV2 fixtures.
  function _quotedDestGasOverhead(
    MttpChain storage src,
    uint256 tokenIndex,
    uint256 amount,
    FeeQuoter feeQuoter
  ) internal view returns (uint32 overhead) {
    address token = src.tokens[tokenIndex];
    bool hasCustomFee;
    (, overhead,,, hasCustomFee) = IPoolV2(src.tar.getPool(token))
      .getFee(token, src.remoteSelector, amount, address(0), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
    if (!hasCustomFee) {
      (, overhead,) = feeQuoter.getTokenTransferFee(src.remoteSelector, token);
    }
  }
}
