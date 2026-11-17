// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../libraries/PolicyConfig.sol";
import {MultiTokenPoolHooks} from "../../pools/MultiTokenPool/MultiTokenPoolHooks.sol";
import {ConnectionPolicyStore} from "../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {e2e_mttpSetup} from "./e2e.mttpSetup.t.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

/// @notice Core MTTP behavior: transfers, policy profiles, shared-policy isolation, token-scoped pauses, subset
/// onboarding and record precedence. Every test runs through the pools, pool config managers and stores on both
/// chains.
contract e2e_mttpCore is e2e_mttpSetup {
  bytes32 internal constant FEE_POLICY_T2 = "FEE_T2";

  /// @dev The baseline lane has no fees, long-format pool data and no hook. Every token has to transfer in both
  /// directions with exact balances.
  function test_mttp_transfersEveryTokenBothWays() public {
    for (uint256 i = 0; i < 3; ++i) {
      uint256 amount = (i + 1) * 100e18;

      uint256 receiverBefore = _balance(s_b.tokens[i], s_receiver);
      uint256 received = _transfer(s_a, s_b, i, amount);
      assertEq(received, amount, "A->B receipt");
      assertEq(_balance(s_b.tokens[i], s_receiver), receiverBefore + amount, "A->B receiver balance");
      assertEq(_balance(s_a.tokens[i], address(s_a.pool)), 0, "A->B no fee retained at baseline");

      receiverBefore = _balance(s_a.tokens[i], s_receiver);
      received = _transfer(s_b, s_a, i, amount);
      assertEq(received, amount, "B->A receipt");
      assertEq(_balance(s_a.tokens[i], s_receiver), receiverBefore + amount, "B->A receiver balance");
      assertEq(_balance(s_b.tokens[i], address(s_b.pool)), 0, "B->A no fee retained at baseline");
    }
  }

  /// @dev The source burn and the destination mint are exact, and the pool data is the long format an MTTP sends
  /// to another MTTP.
  function test_mttp_usesLongFormatPoolData() public {
    (bytes memory poolData, uint256 destAmount,) =
      _send(s_a, s_b, 0, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    assertEq(poolData.length, LONG_POOL_DATA_LENGTH, "long pool data length");
    assertEq(destAmount, 100e18, "post-fee amount");
  }

  /// @dev Transferring one asset must not move the balances of the other two.
  function test_mttp_unrelatedAssetsUnchanged() public {
    uint256 t2Before = _balance(s_a.tokens[1], address(s_a.pool));
    uint256 t3Before = _balance(s_a.tokens[2], address(s_a.pool));
    uint256 t2ReceiverBefore = _balance(s_b.tokens[1], s_receiver);
    uint256 t3ReceiverBefore = _balance(s_b.tokens[2], s_receiver);

    _transfer(s_a, s_b, 0, 100e18);

    assertEq(_balance(s_a.tokens[1], address(s_a.pool)), t2Before, "T2 pool balance unchanged");
    assertEq(_balance(s_a.tokens[2], address(s_a.pool)), t3Before, "T3 pool balance unchanged");
    assertEq(_balance(s_b.tokens[1], s_receiver), t2ReceiverBefore, "T2 receiver unchanged");
    assertEq(_balance(s_b.tokens[2], s_receiver), t3ReceiverBefore, "T3 receiver unchanged");
  }

  /// @dev The three policy profiles: T1 uses the fee policy default, T2 uses a chain-specific record on its own
  /// fee policy, and T3 uses an inline record. Each transfer keeps an exact, distinct fee.
  function test_mttp_policyProfilesChargeDistinctFees() public {
    vm.startPrank(OWNER);
    // T1: policy default record, 100 bps.
    _setFeePolicyConfig(s_a.pcm, FEE_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _feeConfig(100));
    // T2: separate policy with a chain-specific record for B, 200 bps.
    _setFeePolicyConfig(s_a.pcm, FEE_POLICY_T2, s_a.remoteSelector, _feeConfig(200));
    s_a.pcm.applyTokenFeePolicyUpdates(_tokenPolicyOf(s_a.tokens[1], FEE_POLICY_T2));
    // T3: inline record, 300 bps, overriding its assigned policy.
    _setInlineFeeConfig(s_a.pcm, s_a.tokens[2], s_a.remoteSelector, _feeConfig(300));
    vm.stopPrank();

    uint256 amount = 1000e18;

    (uint256 received1, uint256 fee1) = _sendAndDeliver(s_a, s_b, 0, amount);
    assertEq(fee1, amount * 100 / 10_000, "T1 fee");
    assertEq(received1, amount - fee1, "T1 receipt");

    (uint256 received2, uint256 fee2) = _sendAndDeliver(s_a, s_b, 1, amount);
    assertEq(fee2, amount * 200 / 10_000, "T2 fee");
    assertEq(received2, amount - fee2, "T2 receipt");

    (uint256 received3, uint256 fee3) = _sendAndDeliver(s_a, s_b, 2, amount);
    assertEq(fee3, amount * 300 / 10_000, "T3 fee");
    assertEq(received3, amount - fee3, "T3 receipt");
  }

  /// @dev The lane hook runs the sender allowlist in preflight. A listed sender goes through, and an unlisted
  /// sender reverts before any burn or mint.
  function test_mttp_hookAllowsListedSenderAndRejectsOthers() public {
    vm.startPrank(OWNER);
    MultiTokenPoolHooks hooks =
      new MultiTokenPoolHooks(_arrayOf(s_allowedSender), IPolicyEngine(address(0)), _arrayOf(address(s_a.pool)));
    _setInlineSecurityConfig(
      s_a.pcm, s_a.tokens[0], s_a.remoteSelector, new address[](0), new address[](0), address(hooks)
    );
    vm.stopPrank();

    // Allowed sender passes the preflight hook.
    uint256 poolBefore = _balance(s_a.tokens[0], address(s_a.pool));
    (bytes memory poolData, uint256 postFee,) =
      _send(s_a, s_b, 0, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    assertEq(postFee, 100e18, "allowed sender burns");
    assertEq(_balance(s_a.tokens[0], address(s_a.pool)) - poolBefore, 0, "no fee");

    // Disallowed sender reverts on the preflight hook, and no burn is committed.
    _fundPool(s_a, s_a.tokens[0], 100e18);
    uint256 poolBeforeReject = _balance(s_a.tokens[0], address(s_a.pool));
    vm.expectRevert(abi.encodeWithSelector(MultiTokenPoolHooks.SenderNotAllowed.selector, s_disallowedSender));
    _lockOrBurn(s_a, s_b, 0, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_disallowedSender);
    assertEq(_balance(s_a.tokens[0], address(s_a.pool)), poolBeforeReject, "no burn on hook rejection");

    // allowed message still delivers.
    uint256 received = _deliver(s_b, s_a, s_b.tokens[0], postFee, poolData, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
    assertEq(received, 100e18, "allowed message delivered");
  }

  /// @dev T1 and T2 share a rate limit policy. Draining T1's outbound bucket must not touch T2's, because tokens
  /// share the policy config but each has its own bucket.
  function test_mttp_sharedRateLimitPolicyHasIndependentBuckets() public {
    vm.startPrank(OWNER);
    // T1 and T2 share RATE_LIMIT_POLICY with a small capacity. T3 keeps a larger inline limit.
    _setRateLimitPolicyConfig(s_a.pcm, RATE_LIMIT_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _rateLimit(100e18));
    _setInlineRateLimitConfig(s_a.pcm, s_a.tokens[2], s_a.remoteSelector, _rateLimit(1000e18));
    vm.stopPrank();

    // T1 exhausts its outbound allowance.
    (uint256 received,) = _sendAndDeliver(s_a, s_b, 0, 100e18);
    assertEq(received, 100e18, "T1 first send");

    // T1 cannot send again.
    _fundPool(s_a, s_a.tokens[0], 1e18);
    vm.expectRevert();
    _lockOrBurn(s_a, s_b, 0, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);

    // T2 still sends: the config is shared, the bucket is not.
    (received,) = _sendAndDeliver(s_a, s_b, 1, 100e18);
    assertEq(received, 100e18, "T2 independent bucket");
  }

  /// @dev An inbound pause is scoped to one token. T1's inbound is paused, T2 and T3 still execute, and the
  /// failed T1 message succeeds after unpausing.
  function test_mttp_inboundPauseIsTokenScopedAndRetrySucceeds() public {
    vm.prank(OWNER);
    s_b.pcm.pauseTokens(_arrayOf(s_b.tokens[0]), ConnectionPolicyStore.PauseDirection.Inbound);

    // T1: the source burn commits, the destination mint is blocked.
    (bytes memory poolData, uint256 postFee,) =
      _send(s_a, s_b, 0, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    uint256 receiverBefore = _balance(s_b.tokens[0], s_receiver);
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.TokenIsPaused.selector, s_b.tokens[0], ConnectionPolicyStore.PauseDirection.Inbound
      )
    );
    _deliver(s_b, s_a, s_b.tokens[0], postFee, poolData, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
    assertEq(_balance(s_b.tokens[0], s_receiver), receiverBefore, "no mint while paused");

    // T2 and T3 are unaffected by T1's pause.
    (uint256 received2,) = _sendAndDeliver(s_a, s_b, 1, 100e18);
    assertEq(received2, 100e18, "T2 executes while T1 paused");
    (uint256 received3,) = _sendAndDeliver(s_a, s_b, 2, 100e18);
    assertEq(received3, 100e18, "T3 executes while T1 paused");

    // Unpause and retry the same message.
    vm.prank(OWNER);
    s_b.pcm.unpauseTokens(_arrayOf(s_b.tokens[0]), ConnectionPolicyStore.PauseDirection.Inbound);
    uint256 received = _deliver(s_b, s_a, s_b.tokens[0], postFee, poolData, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
    assertEq(received, 100e18, "T1 retry succeeds once");
    assertEq(_balance(s_b.tokens[0], s_receiver), receiverBefore + 100e18, "T1 retry payout once");
  }

  /// @dev A connection policy restricted to a subset lets T1 transfer while T2 and T3 are rejected. Adding T2 to
  /// the subset activates it and T3 stays excluded. Rejected sends leave balances and fees unchanged.
  function test_mttp_subsetOnboardingActivatesTokenByToken() public {
    vm.startPrank(OWNER);
    // Restrict the connection policy to a subset and include only T1.
    _setConnectionPolicyConfig(s_a.pcm, CONNECTION_POLICY, s_a.remoteSelector, true);
    _addSubsetTokens(s_a.pcm, CONNECTION_POLICY, s_a.remoteSelector, _arrayOf(s_a.tokens[0]));
    vm.stopPrank();

    // T1 transfers.
    assertEq(_transfer(s_a, s_b, 0, 100e18), 100e18, "T1 in subset transfers");

    // T2 and T3 cannot send. Rejected sends leave balances and fees unchanged.
    for (uint256 i = 1; i < 3; ++i) {
      uint256 poolBefore = _balance(s_a.tokens[i], address(s_a.pool));
      uint256 receiverBefore = _balance(s_b.tokens[i], s_receiver);
      _fundPool(s_a, s_a.tokens[i], 100e18);
      vm.expectRevert(
        abi.encodeWithSelector(
          ConnectionPolicyStore.SubsetTokenNotFound.selector, CONNECTION_POLICY, s_a.remoteSelector, s_a.tokens[i]
        )
      );
      _lockOrBurn(s_a, s_b, i, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
      // Only the funding transfer is on the pool; nothing was burned.
      assertEq(_balance(s_a.tokens[i], address(s_a.pool)), poolBefore + 100e18, "no burn on rejected send");
      assertEq(_balance(s_b.tokens[i], s_receiver), receiverBefore, "no mint on rejected send");
    }

    // Add T2. Its remote mapping and remote trust are already in place.
    vm.prank(OWNER);
    _addSubsetTokens(s_a.pcm, CONNECTION_POLICY, s_a.remoteSelector, _arrayOf(s_a.tokens[1]));
    assertEq(_transfer(s_a, s_b, 1, 100e18), 100e18, "T2 transfers after activation");

    // T3 remains excluded.
    _fundPool(s_a, s_a.tokens[2], 100e18);
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.SubsetTokenNotFound.selector, CONNECTION_POLICY, s_a.remoteSelector, s_a.tokens[2]
      )
    );
    _lockOrBurn(s_a, s_b, 2, 100e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
  }
}
