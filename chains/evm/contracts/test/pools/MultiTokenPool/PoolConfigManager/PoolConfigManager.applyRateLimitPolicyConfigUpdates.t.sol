// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {RateLimiter} from "../../../../libraries/RateLimiter.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyRateLimitPolicyConfigUpdates is PoolConfigManagerRolesSetup {
  /// @notice Removes the last record and its policy, then sets a record for another policy. The security store
  /// emits the events.
  function test_applyRateLimitPolicyConfigUpdates() public {
    RateLimiter.Config memory config = _getRateLimitConfig(CAPACITY);
    _setRateLimitPolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR, config);

    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.RateLimitPolicyConfigRemoved(OTHER_POLICY, OTHER_CHAIN_SELECTOR);
    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.RateLimitPolicyRemoved(OTHER_POLICY);
    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.RateLimitPolicyConfigSet(
      DEFAULT_POLICY,
      OTHER_CHAIN_SELECTOR,
      DEFAULT_TOKEN_DECIMALS,
      SecurityPolicyStore.RateLimitConfig({
        outbound: config, inbound: config, fastFinalityOutbound: config, fastFinalityInbound: config
      })
    );
    s_poolConfigManager.applyRateLimitPolicyConfigUpdates(
      _policyRecordKeysOf(OTHER_POLICY, OTHER_CHAIN_SELECTOR),
      _policyNamesOf(OTHER_POLICY),
      _getRateLimitPolicyConfigs(DEFAULT_POLICY)
    );

    (,, bytes32[] memory rateLimitPolicyNames,,,) = s_poolConfigManager.getPolicyNames();
    assertEq(_policyNamesOf(DEFAULT_POLICY), rateLimitPolicyNames);
  }

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyRateLimitPolicyConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyRateLimitPolicyConfigUpdates,
        (
          new PolicyConfig.PolicyRecordKey[](0),
          new bytes32[](0),
          new SecurityPolicyStore.RateLimitPolicyConfigArgs[](0)
        )
      ),
      _getPolicyWriteRoles(s_poolConfigManager.RATE_LIMIT_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applyRateLimitPolicyConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyRateLimitPolicyConfigUpdates,
        (
          new PolicyConfig.PolicyRecordKey[](0),
          new bytes32[](0),
          new SecurityPolicyStore.RateLimitPolicyConfigArgs[](0)
        )
      ),
      _getPolicyWriteRoles(s_poolConfigManager.RATE_LIMIT_ROLE())
    );
  }

  /// @dev A policy record for the other chain, with the same enabled limit in both directions for both finality
  /// classes. The policy has the decimals of the test tokens.
  function _getRateLimitPolicyConfigs(
    bytes32 policyName
  ) internal pure returns (SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory policyConfigs) {
    RateLimiter.Config memory config = _getRateLimitConfig(CAPACITY);
    policyConfigs = new SecurityPolicyStore.RateLimitPolicyConfigArgs[](1);
    policyConfigs[0] = SecurityPolicyStore.RateLimitPolicyConfigArgs({
      policyName: policyName,
      remoteChainSelector: OTHER_CHAIN_SELECTOR,
      localDecimals: DEFAULT_TOKEN_DECIMALS,
      rateLimitConfig: SecurityPolicyStore.RateLimitConfig({
        outbound: config, inbound: config, fastFinalityOutbound: config, fastFinalityInbound: config
      })
    });
    return policyConfigs;
  }
}
