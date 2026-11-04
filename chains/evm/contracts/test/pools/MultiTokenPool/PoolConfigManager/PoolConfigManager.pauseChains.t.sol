// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";

import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_pauseChains is PoolConfigManagerRolesSetup {
  uint64[] internal s_remoteChainSelectors = [DEST_CHAIN_SELECTOR];

  /// @notice The default admin pauses a chain without PAUSER_ROLE and without CONFIG_ROLE.
  function test_pauseChains() public {
    s_poolConfigManager.revokeRole(s_poolConfigManager.CONFIG_ROLE(), OWNER);

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.ChainPauseSet(DEST_CHAIN_SELECTOR, false, true);
    s_poolConfigManager.pauseChains(s_remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Inbound);

    assertTrue(_getLaneConnection(s_tokenB, DEST_CHAIN_SELECTOR).isChainInboundPaused);
  }

  /// @notice Either DEFAULT_ADMIN_ROLE or PAUSER_ROLE authorizes the call on its own.
  function test_pauseChains_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.pauseChains, (s_remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Outbound)
      ),
      _getPauseRoles()
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_pauseChains_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.pauseChains, (s_remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Outbound)
      ),
      _getPauseRoles()
    );
  }
}
