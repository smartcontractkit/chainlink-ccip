// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

contract SecurityPolicyStore_getLaneRecordLevelBatch is SecurityPolicyStoreSetup {
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
}
