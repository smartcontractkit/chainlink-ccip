// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT security policy, which has a policy default record with no CCVs and no hook.
contract PoolConfigManager_getSecurityPolicy is PoolConfigManagerSetup {
  /// @notice Returns the token count and records from the security store.
  function test_getSecurityPolicy() public {
    address[] memory outboundCCVs = _arrayOf(makeAddr("outboundCCV"));
    address[] memory inboundCCVs = _arrayOf(makeAddr("inboundCCV"));
    address hook = makeAddr("hook");
    _setSecurityPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, outboundCCVs, inboundCCVs, hook);

    (uint32 storeTokenCount, SecurityPolicyStore.SecurityPolicyConfigArgs[] memory storeConfigs) =
      s_securityStore.getSecurityPolicy(DEFAULT_POLICY);
    (uint32 tokenCount, SecurityPolicyStore.SecurityPolicyConfigArgs[] memory configs) =
      s_poolConfigManager.getSecurityPolicy(DEFAULT_POLICY);
    assertEq(abi.encode(storeTokenCount, storeConfigs), abi.encode(tokenCount, configs));
  }
}
