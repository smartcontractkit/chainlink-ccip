// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";

import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_unpauseGlobally is PoolConfigManagerRolesSetup {
  function setUp() public virtual override {
    super.setUp();
    s_poolConfigManager.pauseGlobally();
  }

  /// @notice The default admin turns off the global pause without UNPAUSER_ROLE and without CONFIG_ROLE.
  function test_unpauseGlobally() public {
    s_poolConfigManager.revokeRole(s_poolConfigManager.CONFIG_ROLE(), OWNER);

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.GlobalPauseSet(false);
    s_poolConfigManager.unpauseGlobally();

    assertFalse(_getLaneConnection(s_tokenB, DEST_CHAIN_SELECTOR).isGloballyPaused);
  }

  /// @notice The default admin and UNPAUSER_ROLE each pass alone.
  function test_unpauseGlobally_AllowedRoles() public {
    _assertRolesPass(abi.encodeCall(PoolConfigManager.unpauseGlobally, ()), _getUnpauseRoles());
  }

  // Reverts

  /// @notice Each other PCM role and a caller with no role revert. The error names the second role of the check.
  function test_unpauseGlobally_RevertWhen_MissingRole() public {
    _assertRolesRevert(abi.encodeCall(PoolConfigManager.unpauseGlobally, ()), _getUnpauseRoles());
  }
}
