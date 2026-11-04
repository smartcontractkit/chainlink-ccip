// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Token B has no lane to the other chain.
contract PoolConfigManager_getInlineRateLimitConfigs is PoolConfigManagerSetup {
  /// @notice Returns the inline rate limit records from the security store, including a record on a chain without a
  /// lane.
  function test_getInlineRateLimitConfigs() public {
    _setInlineRateLimitConfig(s_tokenB, OTHER_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));

    SecurityPolicyStore.InlineRateLimitConfigArgs[] memory storeConfigs =
      s_securityStore.getInlineRateLimitConfigs(s_tokenB);
    SecurityPolicyStore.InlineRateLimitConfigArgs[] memory configs =
      s_poolConfigManager.getInlineRateLimitConfigs(s_tokenB);
    assertEq(abi.encode(storeConfigs), abi.encode(configs));
  }
}
