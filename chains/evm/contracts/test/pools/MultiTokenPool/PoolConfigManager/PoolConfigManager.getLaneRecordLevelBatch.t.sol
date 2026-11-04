// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerTransferSetup} from "./PoolConfigManagerTransferSetup.t.sol";

/// @notice Token C uses the DEFAULT policy of every type, and the pool serves it.
/// @dev The ConnectionPolicyStore has no remote pool policy default record, so the remote pool level is never
/// PolicyDefault.
contract PoolConfigManager_getLaneRecordLevelBatch is PoolConfigManagerTransferSetup {
  uint256 internal constant AMOUNT = 1e18;
  PolicyConfig.RecordLevel internal constant NONE = PolicyConfig.RecordLevel.None;
  PolicyConfig.RecordLevel internal constant INLINE = PolicyConfig.RecordLevel.Inline;
  PolicyConfig.RecordLevel internal constant POLICY_CHAIN = PolicyConfig.RecordLevel.PolicyChain;
  PolicyConfig.RecordLevel internal constant POLICY_DEFAULT = PolicyConfig.RecordLevel.PolicyDefault;

  address internal s_tokenC;

  function setUp() public virtual override {
    super.setUp();
    s_tokenC = _deployToken("C");
    _registerToken(s_tokenC);
    _registerPool(s_tokenC);
  }

  /// @notice Each lane has a distinct combination of record levels. Swapping returned arrays must fail the
  /// assertions.
  function test_getLaneRecordLevelBatch() public {
    _setInlineConnectionConfig(s_tokenB, DEST_CHAIN_SELECTOR, abi.encode(makeAddr("remoteTokenB")));
    _setInlineRemotePoolConfig(s_tokenB, OTHER_CHAIN_SELECTOR, _getRemotePoolConfig());
    _setInlineRateLimitConfig(s_tokenB, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setFeePolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getFeeConfig(100, 10));
    _setInlineDestOverheadConfig(s_tokenB, OTHER_CHAIN_SELECTOR, 300_000, 320);
    _setSecurityPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, new address[](0), new address[](0), address(0));
    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](2);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_tokenB, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_tokenB, remoteChainSelector: OTHER_CHAIN_SELECTOR});

    (
      PolicyConfig.RecordLevel[] memory connectionLevels,
      PolicyConfig.RecordLevel[] memory remotePoolLevels,
      PolicyConfig.RecordLevel[] memory rateLimitLevels,
      PolicyConfig.RecordLevel[] memory feeLevels,
      PolicyConfig.RecordLevel[] memory destOverheadLevels,
      PolicyConfig.RecordLevel[] memory securityLevels
    ) = s_poolConfigManager.getLaneRecordLevelBatch(lanes);

    assertEq(abi.encode(_levelsOf(INLINE, NONE)), abi.encode(connectionLevels), "connection");
    assertEq(abi.encode(_levelsOf(POLICY_CHAIN, INLINE)), abi.encode(remotePoolLevels), "remote pool");
    assertEq(abi.encode(_levelsOf(INLINE, POLICY_DEFAULT)), abi.encode(rateLimitLevels), "rate limit");
    assertEq(abi.encode(_levelsOf(POLICY_CHAIN, POLICY_DEFAULT)), abi.encode(feeLevels), "fee");
    assertEq(abi.encode(_levelsOf(POLICY_DEFAULT, INLINE)), abi.encode(destOverheadLevels), "destination overhead");
    assertEq(abi.encode(_levelsOf(POLICY_DEFAULT, POLICY_CHAIN)), abi.encode(securityLevels), "security");
  }

  /// @notice A missing record returns None and causes transfer validation to revert. Assigning the missing policy
  /// updates the reported level immediately. Validation then reaches the next missing record: remote pool, fee,
  /// destination overhead, security, and rate limit.
  function test_getLaneRecordLevelBatch_NoneMatchesTransferRevert() public {
    _setTokenRemotePoolPolicy(s_tokenC, bytes32(0));
    _setTokenRateLimitPolicy(s_tokenC, bytes32(0));
    _setTokenFeePolicy(s_tokenC, bytes32(0));
    _setTokenDestOverheadPolicy(s_tokenC, bytes32(0));
    _setTokenSecurityPolicy(s_tokenC, bytes32(0));

    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](2);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_tokenC, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_tokenC, remoteChainSelector: OTHER_CHAIN_SELECTOR});
    _assertLevels(
      lanes, 0, PolicyConfig.RecordLevel.PolicyChain, PolicyConfig.RecordLevel.None, PolicyConfig.RecordLevel.None
    );
    _assertLevels(lanes, 1, PolicyConfig.RecordLevel.None, PolicyConfig.RecordLevel.None, PolicyConfig.RecordLevel.None);
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.ConnectionRecordNotFound.selector, s_tokenC, OTHER_CHAIN_SELECTOR)
    );
    _validateLockOrBurn(s_tokenC, OTHER_CHAIN_SELECTOR);
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.RemotePoolRecordNotFound.selector, s_tokenC, DEST_CHAIN_SELECTOR)
    );
    _validateLockOrBurn(s_tokenC, DEST_CHAIN_SELECTOR);

    _setTokenRemotePoolPolicy(s_tokenC, DEFAULT_POLICY);
    (, PolicyConfig.RecordLevel[] memory remotePoolLevels,, PolicyConfig.RecordLevel[] memory feeLevels,,) =
      s_poolConfigManager.getLaneRecordLevelBatch(lanes);
    assertEq(uint8(PolicyConfig.RecordLevel.PolicyChain), uint8(remotePoolLevels[0]));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(feeLevels[0]));
    vm.expectRevert(abi.encodeWithSelector(FeePolicyStore.FeeRecordNotFound.selector, s_tokenC, DEST_CHAIN_SELECTOR));
    _validateLockOrBurn(s_tokenC, DEST_CHAIN_SELECTOR);

    _setTokenFeePolicy(s_tokenC, DEFAULT_POLICY);
    PolicyConfig.RecordLevel[] memory destOverheadLevels;
    (,,, feeLevels, destOverheadLevels,) = s_poolConfigManager.getLaneRecordLevelBatch(lanes);
    assertEq(uint8(PolicyConfig.RecordLevel.PolicyDefault), uint8(feeLevels[0]));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(destOverheadLevels[0]));
    vm.expectRevert(
      abi.encodeWithSelector(FeePolicyStore.DestOverheadRecordNotFound.selector, s_tokenC, DEST_CHAIN_SELECTOR)
    );
    _validateLockOrBurn(s_tokenC, DEST_CHAIN_SELECTOR);

    _setTokenDestOverheadPolicy(s_tokenC, DEFAULT_POLICY);
    PolicyConfig.RecordLevel[] memory securityLevels;
    (,,,, destOverheadLevels, securityLevels) = s_poolConfigManager.getLaneRecordLevelBatch(lanes);
    assertEq(uint8(PolicyConfig.RecordLevel.PolicyDefault), uint8(destOverheadLevels[0]));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(securityLevels[0]));
    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.SecurityRecordNotFound.selector, s_tokenC, DEST_CHAIN_SELECTOR)
    );
    _validateLockOrBurn(s_tokenC, DEST_CHAIN_SELECTOR);

    _setTokenSecurityPolicy(s_tokenC, DEFAULT_POLICY);
    PolicyConfig.RecordLevel[] memory rateLimitLevels;
    (,, rateLimitLevels,,, securityLevels) = s_poolConfigManager.getLaneRecordLevelBatch(lanes);
    assertEq(uint8(PolicyConfig.RecordLevel.PolicyDefault), uint8(securityLevels[0]));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(rateLimitLevels[0]));
    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.RateLimitRecordNotFound.selector, s_tokenC, DEST_CHAIN_SELECTOR)
    );
    _validateLockOrBurn(s_tokenC, DEST_CHAIN_SELECTOR);

    _setTokenRateLimitPolicy(s_tokenC, DEFAULT_POLICY);
    _assertLevels(
      lanes,
      0,
      PolicyConfig.RecordLevel.PolicyChain,
      PolicyConfig.RecordLevel.PolicyChain,
      PolicyConfig.RecordLevel.PolicyDefault
    );
    _validateLockOrBurn(s_tokenC, DEST_CHAIN_SELECTOR);
  }

  // ================================================================
  // │                           Helpers                            │
  // ================================================================

  /// @dev Checks the levels of one lane of the batch. The types other than the connection and remote pool types share
  /// one expected level, because the tests set their records together.
  function _assertLevels(
    IPoolConfigManager.LaneKey[] memory lanes,
    uint256 index,
    PolicyConfig.RecordLevel connectionLevel,
    PolicyConfig.RecordLevel remotePoolLevel,
    PolicyConfig.RecordLevel otherLevel
  ) internal view {
    (
      PolicyConfig.RecordLevel[] memory connectionLevels,
      PolicyConfig.RecordLevel[] memory remotePoolLevels,
      PolicyConfig.RecordLevel[] memory rateLimitLevels,
      PolicyConfig.RecordLevel[] memory feeLevels,
      PolicyConfig.RecordLevel[] memory destOverheadLevels,
      PolicyConfig.RecordLevel[] memory securityLevels
    ) = s_poolConfigManager.getLaneRecordLevelBatch(lanes);
    assertEq(lanes.length, connectionLevels.length);
    assertEq(uint8(connectionLevel), uint8(connectionLevels[index]), "connection");
    assertEq(uint8(remotePoolLevel), uint8(remotePoolLevels[index]), "remote pool");
    assertEq(uint8(otherLevel), uint8(rateLimitLevels[index]), "rate limit");
    assertEq(uint8(otherLevel), uint8(feeLevels[index]), "fee");
    assertEq(uint8(otherLevel), uint8(destOverheadLevels[index]), "destination overhead");
    assertEq(uint8(otherLevel), uint8(securityLevels[index]), "security");
  }

  /// @dev Calls validateLockOrBurn as the pool, for default finality and with the token fee.
  function _validateLockOrBurn(
    address localToken,
    uint64 remoteChainSelector
  ) internal returns (IPoolConfigManager.OutboundLane memory lane) {
    _setCaller(address(s_pool));
    lane = s_poolConfigManager.validateLockOrBurn(
      localToken, remoteChainSelector, AMOUNT, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", true
    );
    _setCaller(OWNER);
    return lane;
  }

  function _levelsOf(
    PolicyConfig.RecordLevel firstLevel,
    PolicyConfig.RecordLevel secondLevel
  ) internal pure returns (PolicyConfig.RecordLevel[] memory levels) {
    levels = new PolicyConfig.RecordLevel[](2);
    levels[0] = firstLevel;
    levels[1] = secondLevel;
    return levels;
  }
}
