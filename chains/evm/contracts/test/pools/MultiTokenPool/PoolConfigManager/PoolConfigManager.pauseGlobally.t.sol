// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";

import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_pauseGlobally is PoolConfigManagerRolesSetup {
  /// @notice The default admin turns on the global pause without PAUSER_ROLE and without CONFIG_ROLE.
  function test_pauseGlobally() public {
    s_poolConfigManager.revokeRole(s_poolConfigManager.CONFIG_ROLE(), OWNER);

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.GlobalPauseSet(true);
    s_poolConfigManager.pauseGlobally();

    assertTrue(_getLaneConnection(s_tokenB, DEST_CHAIN_SELECTOR).isGloballyPaused);
  }

  /// @notice Either DEFAULT_ADMIN_ROLE or PAUSER_ROLE authorizes the call on its own.
  function test_pauseGlobally_AllowedRoles() public {
    _assertRolesPass(abi.encodeCall(PoolConfigManager.pauseGlobally, ()), _getPauseRoles());
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_pauseGlobally_RevertWhen_MissingRole() public {
    _assertRolesRevert(abi.encodeCall(PoolConfigManager.pauseGlobally, ()), _getPauseRoles());
  }
}
