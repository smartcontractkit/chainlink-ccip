// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Token B has no lane to the other chain.
contract PoolConfigManager_getInlineDestOverheadConfigs is PoolConfigManagerSetup {
  /// @notice Returns the inline destination overhead records from the fee store, including a record on a chain
  /// without a lane.
  function test_getInlineDestOverheadConfigs() public {
    _setInlineDestOverheadConfig(
      s_tokenB, OTHER_CHAIN_SELECTOR, DEFAULT_TOKEN_DEST_GAS_OVERHEAD, MIN_DEST_BYTES_OVERHEAD
    );

    FeePolicyStore.InlineDestOverheadConfigArgs[] memory storeConfigs =
      s_feeStore.getInlineDestOverheadConfigs(s_tokenB);
    FeePolicyStore.InlineDestOverheadConfigArgs[] memory configs =
      s_poolConfigManager.getInlineDestOverheadConfigs(s_tokenB);
    assertEq(abi.encode(storeConfigs), abi.encode(configs));
  }
}
