// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Token B uses the DEFAULT connection policy, which has a destination chain record and a remote token
/// address provider.
contract PoolConfigManager_getConnectionPolicy is PoolConfigManagerSetup {
  /// @notice Returns the token count, remote token address provider, and records from the connection store.
  function test_getConnectionPolicy() public view {
    (
      uint32 storeTokenCount,
      address storeProvider,
      ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory storeConfigs
    ) = s_connectionStore.getConnectionPolicy(DEFAULT_POLICY);
    (uint32 tokenCount, address provider, ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory configs) =
      s_poolConfigManager.getConnectionPolicy(DEFAULT_POLICY);
    assertEq(abi.encode(storeTokenCount, storeProvider, storeConfigs), abi.encode(tokenCount, provider, configs));
  }
}
