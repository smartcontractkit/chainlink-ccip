// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

contract SecurityPolicyStore_views is SecurityPolicyStoreSetup {
  /// @notice getTokenPolicies and getTokenPolicyNamesBatch report the same token state.
  function test_getTokenPolicies() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setSecurityPolicyConfig(OTHER_POLICY, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);
    _setTokenSecurityPolicy(s_tokenA, OTHER_POLICY);

    (bool isRegistered, uint8 localDecimals, bytes32 rateLimitPolicyName, bytes32 securityPolicyName) =
      s_store.getTokenPolicies(s_tokenA);
    assertTrue(isRegistered);
    assertEq(DEFAULT_TOKEN_DECIMALS, localDecimals);
    assertEq(DEFAULT_POLICY, rateLimitPolicyName);
    assertEq(OTHER_POLICY, securityPolicyName);

    (bytes32[] memory rateLimitPolicyNames, bytes32[] memory securityPolicyNames) =
      s_store.getTokenPolicyNamesBatch(_arrayOf(s_tokenA));
    assertEq(DEFAULT_POLICY, rateLimitPolicyNames[0]);
    assertEq(OTHER_POLICY, securityPolicyNames[0]);

    (isRegistered, localDecimals, rateLimitPolicyName, securityPolicyName) = s_store.getTokenPolicies(s_tokenB);
    assertTrue(isRegistered);
    assertEq(DEFAULT_TOKEN_DECIMALS, localDecimals);
    assertEq(bytes32(0), rateLimitPolicyName);
    assertEq(bytes32(0), securityPolicyName);
  }

  /// @notice getPolicyNames reports both policy types.
  function test_getPolicyNames() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setSecurityPolicyConfig(OTHER_POLICY, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);

    (bytes32[] memory rateLimitPolicyNames, bytes32[] memory securityPolicyNames) = s_store.getPolicyNames();
    assertEq(1, rateLimitPolicyNames.length);
    assertEq(DEFAULT_POLICY, rateLimitPolicyNames[0]);
    assertEq(1, securityPolicyNames.length);
    assertEq(OTHER_POLICY, securityPolicyNames[0]);
  }

  /// @notice getLaneVerificationBatch returns the hook and the two CCV lists of a lane, and empty for a lane that does
  /// not exist.
  function test_getLaneVerificationBatch() public {
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv2), _arrayOf(s_ccv2), s_hook);

    IPoolConfigManager.LaneKey[] memory lanes = _lanesOf(s_tokenA, DEST_CHAIN_SELECTOR);
    bool[] memory hasLanes = _trueArrayOf();
    (address[] memory hooks, address[][] memory outboundCCVs, address[][] memory inboundCCVs) =
      s_store.getLaneVerificationBatch(lanes, hasLanes);

    assertEq(s_hook, hooks[0]);
    assertEq(abi.encode(_ccvsOf(s_ccv1, s_ccv2)), abi.encode(outboundCCVs[0]));
    assertEq(abi.encode(_arrayOf(s_ccv2)), abi.encode(inboundCCVs[0]));

    // A lane that does not exist gives an empty entry.
    bool[] memory noLanes = new bool[](1);
    (hooks, outboundCCVs, inboundCCVs) = s_store.getLaneVerificationBatch(lanes, noLanes);
    assertEq(address(0), hooks[0]);
    assertEq(0, outboundCCVs[0].length);
    assertEq(0, inboundCCVs[0].length);
  }

  /// @notice getLaneRecordLevelBatch follows the resolution of a transfer: the inline record, then the policy chain
  /// record, then the policy default record, then None.
  function test_getLaneRecordLevelBatch() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setSecurityPolicyConfig(OTHER_POLICY, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);
    _setSecurityPolicyConfig(
      OTHER_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook
    );
    _setTokenSecurityPolicy(s_tokenA, OTHER_POLICY);

    // Token A has an inline rate limit record and a policy chain security record.
    (PolicyConfig.RecordLevel[] memory rateLimitLevels, PolicyConfig.RecordLevel[] memory securityLevels) =
      s_store.getLaneRecordLevelBatch(_lanesOf(s_tokenA, DEST_CHAIN_SELECTOR));
    assertEq(uint256(PolicyConfig.RecordLevel.Inline), uint256(rateLimitLevels[0]));
    assertEq(uint256(PolicyConfig.RecordLevel.PolicyChain), uint256(securityLevels[0]));

    // Token A falls back to the policy default security record on a chain with no chain record.
    (, securityLevels) = s_store.getLaneRecordLevelBatch(_lanesOf(s_tokenA, OTHER_CHAIN_SELECTOR));
    assertEq(uint256(PolicyConfig.RecordLevel.PolicyDefault), uint256(securityLevels[0]));

    // Token B has no record of either type.
    (rateLimitLevels, securityLevels) = s_store.getLaneRecordLevelBatch(_lanesOf(s_tokenB, DEST_CHAIN_SELECTOR));
    assertEq(uint256(PolicyConfig.RecordLevel.None), uint256(rateLimitLevels[0]));
    assertEq(uint256(PolicyConfig.RecordLevel.None), uint256(securityLevels[0]));
  }

  /// @notice checkLaneRecords reverts for the first record type that the token misses.
  function test_checkLaneRecords() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);
    s_store.checkLaneRecords(s_tokenA, DEST_CHAIN_SELECTOR);

    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.RateLimitRecordNotFound.selector, s_tokenB, DEST_CHAIN_SELECTOR)
    );
    s_store.checkLaneRecords(s_tokenB, DEST_CHAIN_SELECTOR);

    _setInlineRateLimitConfig(s_tokenB, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.SecurityRecordNotFound.selector, s_tokenB, DEST_CHAIN_SELECTOR)
    );
    s_store.checkLaneRecords(s_tokenB, DEST_CHAIN_SELECTOR);
  }

  /// @notice getCurrentRateLimiterStateBatch shows a bucket that waits for a policy copy as a full bucket with the
  /// config of its record.
  function test_getCurrentRateLimiterStateBatch() public {
    _setRateLimitPolicyConfigPerSlot(
      DEFAULT_POLICY,
      DEST_CHAIN_SELECTOR,
      _getRateLimitConfig(CAPACITY),
      _getRateLimitConfig(CAPACITY + 1),
      _getRateLimitConfig(CAPACITY + 2),
      _getRateLimitConfig(CAPACITY + 3)
    );
    _setTokenRateLimitPolicy(s_tokenA, DEFAULT_POLICY);

    IPoolConfigManager.LaneRateLimiterState[] memory states =
      s_store.getCurrentRateLimiterStateBatch(_lanesOf(s_tokenA, DEST_CHAIN_SELECTOR), _trueArrayOf());
    _assertFullBucket(_getRateLimitConfig(CAPACITY), states[0].outbound);
    _assertFullBucket(_getRateLimitConfig(CAPACITY + 1), states[0].inbound);
    _assertFullBucket(_getRateLimitConfig(CAPACITY + 2), states[0].fastFinalityOutbound);
    _assertFullBucket(_getRateLimitConfig(CAPACITY + 3), states[0].fastFinalityInbound);

    // A lane that does not exist gives zero values.
    states = s_store.getCurrentRateLimiterStateBatch(_lanesOf(s_tokenA, DEST_CHAIN_SELECTOR), new bool[](1));
    assertEq(0, states[0].outbound.tokens);
    assertEq(0, states[0].outbound.capacity);
  }
}
