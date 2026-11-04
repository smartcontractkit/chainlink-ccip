// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Token B has no lane to the other chain.
contract PoolConfigManager_getInlineFeeConfigs is PoolConfigManagerSetup {
  /// @notice Returns the inline fee records from the fee store, including a record on a chain without a lane.
  function test_getInlineFeeConfigs() public {
    _setInlineFeeConfig(s_tokenB, OTHER_CHAIN_SELECTOR, _getFeeConfig(100, 10));

    FeePolicyStore.InlineFeeConfigArgs[] memory storeConfigs = s_feeStore.getInlineFeeConfigs(s_tokenB);
    FeePolicyStore.InlineFeeConfigArgs[] memory configs = s_poolConfigManager.getInlineFeeConfigs(s_tokenB);
    assertEq(abi.encode(storeConfigs), abi.encode(configs));
  }
}
