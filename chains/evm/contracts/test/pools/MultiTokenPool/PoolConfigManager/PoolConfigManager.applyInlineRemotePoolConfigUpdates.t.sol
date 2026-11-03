// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyInlineRemotePoolConfigUpdates is PoolConfigManagerRolesSetup {
  /// @notice CONFIG_ROLE removes an inline record and sets another in one call. The connection store applies the call
  /// and emits the events.
  function test_applyInlineRemotePoolConfigUpdates() public {
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    _setInlineRemotePoolConfig(s_tokenB, OTHER_CHAIN_SELECTOR, remotePoolConfig);
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      new ConnectionPolicyStore.InlineRemotePoolConfigArgs[](1);
    inlineConfigs[0] = ConnectionPolicyStore.InlineRemotePoolConfigArgs({
      localToken: s_tokenA,
      remoteChainSelector: OTHER_CHAIN_SELECTOR,
      trustPolicyRemotePools: false,
      remotePoolConfig: remotePoolConfig
    });

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.InlineRemotePoolConfigRemoved(s_tokenB, OTHER_CHAIN_SELECTOR);
    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.InlineRemotePoolConfigSet(s_tokenA, OTHER_CHAIN_SELECTOR, false, remotePoolConfig);
    s_poolConfigManager.applyInlineRemotePoolConfigUpdates(_lanesOf(s_tokenB, OTHER_CHAIN_SELECTOR), inlineConfigs);

    (, PolicyConfig.RecordLevel[] memory remotePoolLevels,,,,) =
      s_poolConfigManager.getLaneRecordLevelBatch(_lanesOf(s_tokenB, OTHER_CHAIN_SELECTOR));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(remotePoolLevels[0]));
    (, remotePoolLevels,,,,) = s_poolConfigManager.getLaneRecordLevelBatch(_lanesOf(s_tokenA, OTHER_CHAIN_SELECTOR));
    assertEq(uint8(PolicyConfig.RecordLevel.Inline), uint8(remotePoolLevels[0]));
  }

  /// @notice CONFIG_ROLE and TOKEN_MANAGER_ROLE each pass alone.
  function test_applyInlineRemotePoolConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyInlineRemotePoolConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new ConnectionPolicyStore.InlineRemotePoolConfigArgs[](0))
      ),
      _getTokenWriteRoles()
    );
  }

  // Reverts

  /// @notice Each other PCM role and a caller with no role revert. The error names the second role of the check.
  function test_applyInlineRemotePoolConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyInlineRemotePoolConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new ConnectionPolicyStore.InlineRemotePoolConfigArgs[](0))
      ),
      _getTokenWriteRoles()
    );
  }
}
