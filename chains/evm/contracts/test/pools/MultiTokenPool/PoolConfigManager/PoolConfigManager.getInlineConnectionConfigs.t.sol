// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Token A has an inline connection record to the destination chain. An inline connection record always gives
/// its token a lane, so no inline connection record exists on a chain without a lane.
contract PoolConfigManager_getInlineConnectionConfigs is PoolConfigManagerSetup {
  /// @notice Returns the inline connection records from the connection store.
  function test_getInlineConnectionConfigs() public view {
    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory storeConfigs =
      s_connectionStore.getInlineConnectionConfigs(s_tokenA);
    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory configs =
      s_poolConfigManager.getInlineConnectionConfigs(s_tokenA);
    assertEq(abi.encode(storeConfigs), abi.encode(configs));
  }
}
