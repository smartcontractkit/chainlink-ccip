// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT security policy, which has a policy default record with no CCVs and no hook.
/// Token B has a lane to the destination chain and no lane to the other chain.
contract PoolConfigManager_getInlineSecurityConfigs is PoolConfigManagerSetup {
  address internal s_hook = makeAddr("hook");

  /// @notice Returns both inline security records, including the record for a chain without a lane.
  function test_getInlineSecurityConfigs() public {
    SecurityPolicyStore.InlineSecurityConfigArgs[] memory expectedConfigs =
      new SecurityPolicyStore.InlineSecurityConfigArgs[](2);
    expectedConfigs[0] = _getInlineSecurityConfigArgs(s_tokenB, DEST_CHAIN_SELECTOR);
    expectedConfigs[1] = SecurityPolicyStore.InlineSecurityConfigArgs({
      localToken: s_tokenB,
      remoteChainSelector: OTHER_CHAIN_SELECTOR,
      securityConfig: SecurityPolicyStore.SecurityConfig({
        outboundCCVs: new address[](0), inboundCCVs: new address[](0), hook: address(0)
      })
    });
    _applyInlineSecurityConfigs(expectedConfigs);

    assertEq(abi.encode(expectedConfigs), abi.encode(s_poolConfigManager.getInlineSecurityConfigs(s_tokenB)));
  }

  /// @dev Two outbound CCVs, the placeholder for default inbound CCVs, and a hook.
  function _getInlineSecurityConfigArgs(
    address localToken,
    uint64 remoteChainSelector
  ) internal returns (SecurityPolicyStore.InlineSecurityConfigArgs memory) {
    address[] memory outboundCCVs = new address[](2);
    outboundCCVs[0] = makeAddr("outboundCCV1");
    outboundCCVs[1] = makeAddr("outboundCCV2");
    address[] memory inboundCCVs = new address[](1);
    inboundCCVs[0] = address(0);
    return SecurityPolicyStore.InlineSecurityConfigArgs({
      localToken: localToken,
      remoteChainSelector: remoteChainSelector,
      securityConfig: SecurityPolicyStore.SecurityConfig({
        outboundCCVs: outboundCCVs, inboundCCVs: inboundCCVs, hook: s_hook
      })
    });
  }

  function _applyInlineSecurityConfigs(
    SecurityPolicyStore.InlineSecurityConfigArgs[] memory inlineConfigs
  ) internal {
    s_poolConfigManager.applyInlineSecurityConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }
}
