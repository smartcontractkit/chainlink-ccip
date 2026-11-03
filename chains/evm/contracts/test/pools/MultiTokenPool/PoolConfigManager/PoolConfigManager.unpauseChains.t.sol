// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_unpauseChains is PoolConfigManagerRolesSetup {
  uint64[] internal s_remoteChainSelectors = [DEST_CHAIN_SELECTOR];

  function setUp() public virtual override {
    super.setUp();
    s_poolConfigManager.pauseChains(s_remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Inbound);
  }

  /// @notice The default admin unpauses a chain without UNPAUSER_ROLE and without CONFIG_ROLE.
  function test_unpauseChains() public {
    s_poolConfigManager.revokeRole(s_poolConfigManager.CONFIG_ROLE(), OWNER);

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.ChainPauseSet(DEST_CHAIN_SELECTOR, false, false);
    s_poolConfigManager.unpauseChains(s_remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Inbound);

    assertFalse(_getLaneConnection(s_tokenB, DEST_CHAIN_SELECTOR).isChainInboundPaused);
  }

  /// @notice The default admin and UNPAUSER_ROLE each pass alone.
  function test_unpauseChains_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.unpauseChains, (s_remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Outbound)
      ),
      _getUnpauseRoles()
    );
  }

  // Reverts

  /// @notice Each other PCM role and a caller with no role revert. The error names the second role of the check.
  function test_unpauseChains_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.unpauseChains, (s_remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Outbound)
      ),
      _getUnpauseRoles()
    );
  }
}
