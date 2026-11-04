// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_unpauseTokens is PoolConfigManagerRolesSetup {
  function setUp() public virtual override {
    super.setUp();
    s_poolConfigManager.pauseTokens(_arrayOf(s_tokenB), ConnectionPolicyStore.PauseDirection.Outbound);
  }

  /// @notice The default admin unpauses a token without UNPAUSER_ROLE and without CONFIG_ROLE.
  function test_unpauseTokens() public {
    s_poolConfigManager.revokeRole(s_poolConfigManager.CONFIG_ROLE(), OWNER);

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.TokenPauseSet(s_tokenB, false, false);
    s_poolConfigManager.unpauseTokens(_arrayOf(s_tokenB), ConnectionPolicyStore.PauseDirection.Outbound);

    assertFalse(_getLaneConnection(s_tokenB, DEST_CHAIN_SELECTOR).isTokenOutboundPaused);
  }

  /// @notice Either DEFAULT_ADMIN_ROLE or UNPAUSER_ROLE authorizes the call on its own.
  function test_unpauseTokens_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.unpauseTokens, (_arrayOf(s_tokenB), ConnectionPolicyStore.PauseDirection.Outbound)
      ),
      _getUnpauseRoles()
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_unpauseTokens_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.unpauseTokens, (_arrayOf(s_tokenB), ConnectionPolicyStore.PauseDirection.Outbound)
      ),
      _getUnpauseRoles()
    );
  }
}
