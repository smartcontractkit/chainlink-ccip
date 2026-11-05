// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyInlineSecurityConfigUpdates is PoolConfigManagerRolesSetup {
  address internal s_ccv = makeAddr("ccv");

  /// @notice Removes one inline record and sets another in the same call. The security store emits the events.
  function test_applyInlineSecurityConfigUpdates() public {
    address[] memory ccvs = _arrayOf(s_ccv);
    _setInlineSecurityConfig(s_tokenB, OTHER_CHAIN_SELECTOR, ccvs, ccvs, address(0));
    SecurityPolicyStore.InlineSecurityConfigArgs[] memory inlineConfigs =
      new SecurityPolicyStore.InlineSecurityConfigArgs[](1);
    inlineConfigs[0] = SecurityPolicyStore.InlineSecurityConfigArgs({
      localToken: s_tokenA,
      remoteChainSelector: OTHER_CHAIN_SELECTOR,
      securityConfig: SecurityPolicyStore.SecurityConfig({outboundCCVs: ccvs, inboundCCVs: ccvs, hook: address(0)})
    });

    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.InlineSecurityConfigRemoved(s_tokenB, OTHER_CHAIN_SELECTOR);
    vm.expectEmit(address(s_securityStore));
    emit SecurityPolicyStore.InlineSecurityConfigSet(
      s_tokenA,
      OTHER_CHAIN_SELECTOR,
      SecurityPolicyStore.SecurityConfig({outboundCCVs: ccvs, inboundCCVs: ccvs, hook: address(0)})
    );
    s_poolConfigManager.applyInlineSecurityConfigUpdates(_lanesOf(s_tokenB, OTHER_CHAIN_SELECTOR), inlineConfigs);

    assertEq(0, s_poolConfigManager.getInlineSecurityConfigs(s_tokenB).length);
    assertEq(1, s_poolConfigManager.getInlineSecurityConfigs(s_tokenA).length);
  }

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyInlineSecurityConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyInlineSecurityConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new SecurityPolicyStore.InlineSecurityConfigArgs[](0))
      ),
      _getTokenWriteRoles(s_poolConfigManager.SECURITY_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applyInlineSecurityConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyInlineSecurityConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new SecurityPolicyStore.InlineSecurityConfigArgs[](0))
      ),
      _getTokenWriteRoles(s_poolConfigManager.SECURITY_ROLE())
    );
  }
}
