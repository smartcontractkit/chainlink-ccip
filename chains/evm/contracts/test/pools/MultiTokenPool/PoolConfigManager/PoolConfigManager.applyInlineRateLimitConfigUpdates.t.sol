// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {RateLimiter} from "../../../../libraries/RateLimiter.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyInlineRateLimitConfigUpdates is PoolConfigManagerRolesSetup {
  /// @notice Removes one inline record and sets another in the same call. The security store emits the events.
  function test_applyInlineRateLimitConfigUpdates() public {
    RateLimiter.Config memory config = _getRateLimitConfig(CAPACITY);
    _setInlineRateLimitConfig(s_tokenB, OTHER_CHAIN_SELECTOR, config);

    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.InlineRateLimitConfigRemoved(s_tokenB, OTHER_CHAIN_SELECTOR);
    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.InlineRateLimitConfigSet(
      s_tokenA,
      OTHER_CHAIN_SELECTOR,
      SecurityPolicyStore.RateLimitConfig({
        outbound: config, inbound: config, fastFinalityOutbound: config, fastFinalityInbound: config
      })
    );
    s_poolConfigManager.applyInlineRateLimitConfigUpdates(
      _lanesOf(s_tokenB, OTHER_CHAIN_SELECTOR), _getInlineRateLimitConfigs(s_tokenA, config)
    );

    assertEq(0, s_poolConfigManager.getInlineRateLimitConfigs(s_tokenB).length);
    assertEq(1, s_poolConfigManager.getInlineRateLimitConfigs(s_tokenA).length);
  }

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyInlineRateLimitConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyInlineRateLimitConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new SecurityPolicyStore.InlineRateLimitConfigArgs[](0))
      ),
      _getTokenWriteRoles(s_poolConfigManager.RATE_LIMIT_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applyInlineRateLimitConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyInlineRateLimitConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new SecurityPolicyStore.InlineRateLimitConfigArgs[](0))
      ),
      _getTokenWriteRoles(s_poolConfigManager.RATE_LIMIT_ROLE())
    );
  }

  /// @dev An inline record of the token for the other chain, with the same limit in both directions for both finality
  /// classes.
  function _getInlineRateLimitConfigs(
    address localToken,
    RateLimiter.Config memory config
  ) internal pure returns (SecurityPolicyStore.InlineRateLimitConfigArgs[] memory inlineConfigs) {
    inlineConfigs = new SecurityPolicyStore.InlineRateLimitConfigArgs[](1);
    inlineConfigs[0] = SecurityPolicyStore.InlineRateLimitConfigArgs({
      localToken: localToken,
      remoteChainSelector: OTHER_CHAIN_SELECTOR,
      rateLimitConfig: SecurityPolicyStore.RateLimitConfig({
        outbound: config, inbound: config, fastFinalityOutbound: config, fastFinalityInbound: config
      })
    });
    return inlineConfigs;
  }
}
