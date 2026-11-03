// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyInlineFeeConfigUpdates is PoolConfigManagerRolesSetup {
  /// @notice CONFIG_ROLE removes an inline record and sets another in one call. The fee store applies the call
  /// and emits the events.
  function test_applyInlineFeeConfigUpdates() public {
    FeePolicyStore.FeeConfig memory feeConfig = _getFeeConfig(10, 5);
    _setInlineFeeConfig(s_tokenB, OTHER_CHAIN_SELECTOR, feeConfig);
    FeePolicyStore.InlineFeeConfigArgs[] memory inlineConfigs = new FeePolicyStore.InlineFeeConfigArgs[](1);
    inlineConfigs[0] = FeePolicyStore.InlineFeeConfigArgs({
      localToken: s_tokenA, remoteChainSelector: OTHER_CHAIN_SELECTOR, feeConfig: feeConfig
    });

    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.InlineFeeConfigRemoved(s_tokenB, OTHER_CHAIN_SELECTOR);
    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.InlineFeeConfigSet(s_tokenA, OTHER_CHAIN_SELECTOR, feeConfig);
    s_poolConfigManager.applyInlineFeeConfigUpdates(_lanesOf(s_tokenB, OTHER_CHAIN_SELECTOR), inlineConfigs);

    assertEq(0, s_poolConfigManager.getInlineFeeConfigs(s_tokenB).length);
    assertEq(1, s_poolConfigManager.getInlineFeeConfigs(s_tokenA).length);
  }

  /// @notice CONFIG_ROLE and TOKEN_MANAGER_ROLE each pass alone.
  function test_applyInlineFeeConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyInlineFeeConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new FeePolicyStore.InlineFeeConfigArgs[](0))
      ),
      _getTokenWriteRoles()
    );
  }

  // Reverts

  /// @notice Each other PCM role and a caller with no role revert. The error names the second role of the check.
  function test_applyInlineFeeConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyInlineFeeConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new FeePolicyStore.InlineFeeConfigArgs[](0))
      ),
      _getTokenWriteRoles()
    );
  }
}
