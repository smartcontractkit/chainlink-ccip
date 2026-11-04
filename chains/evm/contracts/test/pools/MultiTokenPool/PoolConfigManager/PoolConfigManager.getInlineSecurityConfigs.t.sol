// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Token B has no lane to the other chain.
contract PoolConfigManager_getInlineSecurityConfigs is PoolConfigManagerSetup {
  /// @notice Returns the inline security records from the security store, including a record on a chain without a
  /// lane.
  function test_getInlineSecurityConfigs() public {
    address[] memory outboundCCVs = _arrayOf(makeAddr("outboundCCV"));
    address[] memory inboundCCVs = _arrayOf(makeAddr("inboundCCV"));
    address hook = makeAddr("hook");
    _setInlineSecurityConfig(s_tokenB, OTHER_CHAIN_SELECTOR, outboundCCVs, inboundCCVs, hook);

    SecurityPolicyStore.InlineSecurityConfigArgs[] memory storeConfigs =
      s_securityStore.getInlineSecurityConfigs(s_tokenB);
    SecurityPolicyStore.InlineSecurityConfigArgs[] memory configs =
      s_poolConfigManager.getInlineSecurityConfigs(s_tokenB);
    assertEq(abi.encode(storeConfigs), abi.encode(configs));
  }
}
