// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT security policy, which has a policy default record.
contract PoolConfigManager_getSecurityPolicy is PoolConfigManagerSetup {
  /// @notice Returns the token count and records for each policy. Chain selector zero identifies the policy default
  /// record.
  function test_getSecurityPolicy() public {
    address[] memory outboundCCVs = _arrayOf(makeAddr("outboundCCV"));
    address[] memory inboundCCVs = _arrayOf(makeAddr("inboundCCV"));
    address hook = makeAddr("hook");
    _setSecurityPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, outboundCCVs, inboundCCVs, hook);
    _setEmptySecurityPolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR);

    (uint32 tokenCount, SecurityPolicyStore.SecurityPolicyConfigArgs[] memory policyConfigs) =
      s_poolConfigManager.getSecurityPolicy(DEFAULT_POLICY);
    assertEq(2, tokenCount);
    assertEq(2, policyConfigs.length);
    assertEq(PolicyConfig.DEFAULT_CHAIN_SELECTOR, policyConfigs[0].remoteChainSelector);
    assertEq(
      abi.encode(
        SecurityPolicyStore.SecurityPolicyConfigArgs({
          policyName: DEFAULT_POLICY,
          remoteChainSelector: DEST_CHAIN_SELECTOR,
          securityConfig: SecurityPolicyStore.SecurityConfig({
            outboundCCVs: outboundCCVs, inboundCCVs: inboundCCVs, hook: hook
          })
        })
      ),
      abi.encode(policyConfigs[1])
    );

    (tokenCount, policyConfigs) = s_poolConfigManager.getSecurityPolicy(OTHER_POLICY);
    assertEq(0, tokenCount);
    assertEq(1, policyConfigs.length);
    assertEq(OTHER_CHAIN_SELECTOR, policyConfigs[0].remoteChainSelector);
  }

  /// @dev A policy record with no CCVs and no hook.
  function _setEmptySecurityPolicyConfig(
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal {
    _setSecurityPolicyConfig(policyName, remoteChainSelector, new address[](0), new address[](0), address(0));
  }
}
