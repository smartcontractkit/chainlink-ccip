// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";

import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_pauseTokens is PoolConfigManagerRolesSetup {
  /// @notice The default admin pauses a token without PAUSER_ROLE and without CONFIG_ROLE.
  function test_pauseTokens() public {
    s_poolConfigManager.revokeRole(s_poolConfigManager.CONFIG_ROLE(), OWNER);

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.TokenPauseSet(s_tokenB, true, false);
    s_poolConfigManager.pauseTokens(_arrayOf(s_tokenB), ConnectionPolicyStore.PauseDirection.Outbound);

    assertTrue(_getLaneConnection(s_tokenB, DEST_CHAIN_SELECTOR).isTokenOutboundPaused);
  }

  /// @notice Either DEFAULT_ADMIN_ROLE or PAUSER_ROLE authorizes the call on its own.
  function test_pauseTokens_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.pauseTokens, (_arrayOf(s_tokenB), ConnectionPolicyStore.PauseDirection.Outbound)
      ),
      _getPauseRoles()
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_pauseTokens_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.pauseTokens, (_arrayOf(s_tokenB), ConnectionPolicyStore.PauseDirection.Outbound)
      ),
      _getPauseRoles()
    );
  }
}
