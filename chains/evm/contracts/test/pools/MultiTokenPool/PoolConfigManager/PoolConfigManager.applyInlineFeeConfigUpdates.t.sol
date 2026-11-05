// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyInlineFeeConfigUpdates is PoolConfigManagerRolesSetup {
  /// @notice Removes one inline record and sets another in the same call. The fee store emits the events.
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

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyInlineFeeConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyInlineFeeConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new FeePolicyStore.InlineFeeConfigArgs[](0))
      ),
      _getTokenWriteRoles(s_poolConfigManager.FEE_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applyInlineFeeConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyInlineFeeConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new FeePolicyStore.InlineFeeConfigArgs[](0))
      ),
      _getTokenWriteRoles(s_poolConfigManager.FEE_ROLE())
    );
  }
}
