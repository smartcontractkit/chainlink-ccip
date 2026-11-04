// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @dev The ConnectionPolicyStore has no remote pool policy default record, so the remote pool level is never
/// PolicyDefault.
contract PoolConfigManager_getLaneRecordLevelBatch is PoolConfigManagerSetup {
  PolicyConfig.RecordLevel internal constant NONE = PolicyConfig.RecordLevel.None;
  PolicyConfig.RecordLevel internal constant INLINE = PolicyConfig.RecordLevel.Inline;
  PolicyConfig.RecordLevel internal constant POLICY_CHAIN = PolicyConfig.RecordLevel.PolicyChain;
  PolicyConfig.RecordLevel internal constant POLICY_DEFAULT = PolicyConfig.RecordLevel.PolicyDefault;

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
