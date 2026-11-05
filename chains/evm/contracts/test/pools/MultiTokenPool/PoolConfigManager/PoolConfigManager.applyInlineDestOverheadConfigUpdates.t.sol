// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyInlineDestOverheadConfigUpdates is PoolConfigManagerRolesSetup {
  /// @notice Removes one inline record and sets another in the same call. The fee store emits the events.
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

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyInlineDestOverheadConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyInlineDestOverheadConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new FeePolicyStore.InlineDestOverheadConfigArgs[](0))
      ),
      _getTokenWriteRoles(s_poolConfigManager.DEST_OVERHEAD_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applyInlineDestOverheadConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyInlineDestOverheadConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new FeePolicyStore.InlineDestOverheadConfigArgs[](0))
      ),
      _getTokenWriteRoles(s_poolConfigManager.DEST_OVERHEAD_ROLE())
    );
  }
}
