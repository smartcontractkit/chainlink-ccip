// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT destination overhead policy, which has a policy default record with the
/// default gas and the minimum bytes.
contract PoolConfigManager_getDestOverheadPolicy is PoolConfigManagerSetup {
  /// @notice Returns the token count and records from the fee store.
  function test_getDestOverheadPolicy() public view {
    (uint32 storeTokenCount, FeePolicyStore.DestOverheadPolicyConfigArgs[] memory storeConfigs) =
      s_feeStore.getDestOverheadPolicy(DEFAULT_POLICY);
    (uint32 tokenCount, FeePolicyStore.DestOverheadPolicyConfigArgs[] memory configs) =
      s_poolConfigManager.getDestOverheadPolicy(DEFAULT_POLICY);
    assertEq(abi.encode(storeTokenCount, storeConfigs), abi.encode(tokenCount, configs));
  }
}
