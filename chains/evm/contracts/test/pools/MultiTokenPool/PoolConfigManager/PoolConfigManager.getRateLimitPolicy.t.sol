// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {RateLimiter} from "../../../../libraries/RateLimiter.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT rate limit policy, which has a policy default record and the decimals of the
/// test tokens.
contract PoolConfigManager_getRateLimitPolicy is PoolConfigManagerSetup {
  /// @dev Use different decimals to check that the view returns policy decimals rather than token decimals.
  uint8 internal constant OTHER_DECIMALS = 6;

  /// @notice Returns the token count, decimals, and records for each policy. Chain selector zero identifies the
  /// policy default record.
  function test_getRateLimitPolicy() public {
    RateLimiter.Config memory rateLimitConfig = _getRateLimitConfig(CAPACITY);
    _setRateLimitPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, rateLimitConfig);
    _setRateLimitPolicyRecord(OTHER_POLICY, OTHER_CHAIN_SELECTOR, OTHER_DECIMALS);

    (uint32 tokenCount, uint8 localDecimals, SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory policyConfigs) =
      s_poolConfigManager.getRateLimitPolicy(DEFAULT_POLICY);
    assertEq(2, tokenCount);
    assertEq(DEFAULT_TOKEN_DECIMALS, localDecimals);
    assertEq(2, policyConfigs.length);
    assertEq(PolicyConfig.DEFAULT_CHAIN_SELECTOR, policyConfigs[0].remoteChainSelector);
    assertEq(
      abi.encode(
        SecurityPolicyStore.RateLimitPolicyConfigArgs({
          policyName: DEFAULT_POLICY,
          remoteChainSelector: DEST_CHAIN_SELECTOR,
          localDecimals: DEFAULT_TOKEN_DECIMALS,
          rateLimitConfig: SecurityPolicyStore.RateLimitConfig({
            outbound: rateLimitConfig,
            inbound: rateLimitConfig,
            fastFinalityOutbound: rateLimitConfig,
            fastFinalityInbound: rateLimitConfig
          })
        })
      ),
      abi.encode(policyConfigs[1])
    );

    (tokenCount, localDecimals, policyConfigs) = s_poolConfigManager.getRateLimitPolicy(OTHER_POLICY);
    assertEq(0, tokenCount);
    assertEq(OTHER_DECIMALS, localDecimals);
    assertEq(1, policyConfigs.length);
    assertEq(OTHER_CHAIN_SELECTOR, policyConfigs[0].remoteChainSelector);
  }

  /// @dev Sets a policy record with disabled rate limits and the given decimals.
  function _setRateLimitPolicyRecord(
    bytes32 policyName,
    uint64 remoteChainSelector,
    uint8 localDecimals
  ) internal {
    RateLimiter.Config memory disabledConfig;
    SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory policyConfigs =
      new SecurityPolicyStore.RateLimitPolicyConfigArgs[](1);
    policyConfigs[0] = SecurityPolicyStore.RateLimitPolicyConfigArgs({
      policyName: policyName,
      remoteChainSelector: remoteChainSelector,
      localDecimals: localDecimals,
      rateLimitConfig: SecurityPolicyStore.RateLimitConfig({
        outbound: disabledConfig,
        inbound: disabledConfig,
        fastFinalityOutbound: disabledConfig,
        fastFinalityInbound: disabledConfig
      })
    });
    s_poolConfigManager.applyRateLimitPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }
}
