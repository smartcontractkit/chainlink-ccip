// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";

contract FeePolicyStore_getLaneRecordLevelBatch is FeePolicyStoreSetup {
  /// @dev Arbitrum One, a third chain for records that the setup does not add.
  uint64 internal constant OTHER_CHAIN_SELECTOR = 4949039107694359620;

  function test_getLaneRecordLevelBatch() public {
    // Both policies get a policy chain record on the source chain and on the other chain. The token also gets an inline
    // destination overhead record on the source chain, and an inline fee record on the other chain. An inline record
    // overrides a policy chain record.
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, SOURCE_CHAIN_SELECTOR, _getFeeConfig()));
    _setFeePolicyConfig(_getFeePolicyConfigArgs(FEE_POLICY, OTHER_CHAIN_SELECTOR, _getFeeConfig()));
    _setDestOverheadPolicyConfig(
      _getDestOverheadPolicyConfigArgs(DEST_OVERHEAD_POLICY, SOURCE_CHAIN_SELECTOR, _getDestOverheadConfig())
    );
    _setDestOverheadPolicyConfig(
      _getDestOverheadPolicyConfigArgs(DEST_OVERHEAD_POLICY, OTHER_CHAIN_SELECTOR, _getDestOverheadConfig())
    );
    _setInlineDestOverheadConfig(
      _getInlineDestOverheadConfigArgs(s_token, SOURCE_CHAIN_SELECTOR, _getDestOverheadConfig())
    );
    _setInlineFeeConfig(_getInlineFeeConfigArgs(s_token, OTHER_CHAIN_SELECTOR, _getFeeConfig()));

    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](5);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: SOURCE_CHAIN_SELECTOR});
    lanes[2] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: OTHER_CHAIN_SELECTOR});
    lanes[3] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[4] =
      IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: PolicyConfig.DEFAULT_CHAIN_SELECTOR});

    (PolicyConfig.RecordLevel[] memory feeLevels, PolicyConfig.RecordLevel[] memory destOverheadLevels) =
      s_feeStore.getLaneRecordLevelBatch(lanes);

    assertEq(uint8(PolicyConfig.RecordLevel.PolicyDefault), uint8(feeLevels[0]));
    assertEq(uint8(PolicyConfig.RecordLevel.PolicyDefault), uint8(destOverheadLevels[0]));
    assertEq(uint8(PolicyConfig.RecordLevel.PolicyChain), uint8(feeLevels[1]));
    assertEq(uint8(PolicyConfig.RecordLevel.Inline), uint8(destOverheadLevels[1]));
    assertEq(uint8(PolicyConfig.RecordLevel.Inline), uint8(feeLevels[2]));
    assertEq(uint8(PolicyConfig.RecordLevel.PolicyChain), uint8(destOverheadLevels[2]));
    // s_otherToken has no policy and no inline record.
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(feeLevels[3]));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(destOverheadLevels[3]));
    // Chain selector zero names the policy default record, so it never resolves to a record.
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(feeLevels[4]));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(destOverheadLevels[4]));
  }
}
