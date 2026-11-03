// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyInlineDestOverheadConfigUpdates is PoolConfigManagerRolesSetup {
  /// @notice CONFIG_ROLE removes an inline record and sets another in one call. The fee store applies the call and
  /// emits the events.
  function test_applyInlineDestOverheadConfigUpdates() public {
    _setInlineDestOverheadConfig(s_tokenB, OTHER_CHAIN_SELECTOR, 300_000, 320);
    FeePolicyStore.InlineDestOverheadConfigArgs[] memory inlineConfigs =
      new FeePolicyStore.InlineDestOverheadConfigArgs[](1);
    inlineConfigs[0] = FeePolicyStore.InlineDestOverheadConfigArgs({
      localToken: s_tokenA,
      remoteChainSelector: OTHER_CHAIN_SELECTOR,
      destOverheadConfig: FeePolicyStore.DestOverheadConfig({destGasOverhead: 400_000, destBytesOverhead: 352})
    });

    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.InlineDestOverheadConfigRemoved(s_tokenB, OTHER_CHAIN_SELECTOR);
    vm.expectEmit(address(s_feeStore));
    emit FeePolicyStore.InlineDestOverheadConfigSet(
      s_tokenA,
      OTHER_CHAIN_SELECTOR,
      FeePolicyStore.DestOverheadConfig({destGasOverhead: 400_000, destBytesOverhead: 352})
    );
    s_poolConfigManager.applyInlineDestOverheadConfigUpdates(_lanesOf(s_tokenB, OTHER_CHAIN_SELECTOR), inlineConfigs);

    assertEq(0, s_poolConfigManager.getInlineDestOverheadConfigs(s_tokenB).length);
    assertEq(1, s_poolConfigManager.getInlineDestOverheadConfigs(s_tokenA).length);
  }

  /// @notice CONFIG_ROLE and TOKEN_MANAGER_ROLE each pass alone.
  function test_applyInlineDestOverheadConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyInlineDestOverheadConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new FeePolicyStore.InlineDestOverheadConfigArgs[](0))
      ),
      _getTokenWriteRoles()
    );
  }

  // Reverts

  /// @notice Each other PCM role and a caller with no role revert. The error names the second role of the check.
  function test_applyInlineDestOverheadConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyInlineDestOverheadConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new FeePolicyStore.InlineDestOverheadConfigArgs[](0))
      ),
      _getTokenWriteRoles()
    );
  }
}
