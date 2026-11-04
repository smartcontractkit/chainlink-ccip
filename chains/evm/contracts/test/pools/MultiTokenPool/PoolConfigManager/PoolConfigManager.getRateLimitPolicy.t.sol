// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT rate limit policy, which has a policy default record with disabled limits.
contract PoolConfigManager_getRateLimitPolicy is PoolConfigManagerSetup {
  /// @notice Returns the token count, decimals, and records from the security store.
  function test_getRateLimitPolicy() public {
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));

    (
      uint32 storeTokenCount,
      uint8 storeLocalDecimals,
      SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory storeConfigs
    ) = s_securityStore.getRateLimitPolicy(DEFAULT_POLICY);
    (uint32 tokenCount, uint8 localDecimals, SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory configs) =
      s_poolConfigManager.getRateLimitPolicy(DEFAULT_POLICY);
    assertEq(
      abi.encode(storeTokenCount, storeLocalDecimals, storeConfigs), abi.encode(tokenCount, localDecimals, configs)
    );
  }
}
