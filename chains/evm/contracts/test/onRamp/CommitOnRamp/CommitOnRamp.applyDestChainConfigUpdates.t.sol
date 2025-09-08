// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CommitOnRamp_applyDestChainConfigUpdates is CommitOnRampSetup {
  function test_applyDestChainConfigUpdates() public {
    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](2);

    destChainConfigs[0] = _getDestChainConfig(makeAddr("ccvProxy1"), 12345, true);
    destChainConfigs[1] = _getDestChainConfig(makeAddr("ccvProxy2"), 67890, false);

    // Expect events for both configurations
    for (uint256 i = 0; i < destChainConfigs.length; ++i) {
      vm.expectEmit();
      emit BaseOnRamp.DestChainConfigSet(
        destChainConfigs[i].destChainSelector, destChainConfigs[i].ccvProxy, destChainConfigs[i].allowlistEnabled
      );
    }

    s_commitOnRamp.applyDestChainConfigUpdates(destChainConfigs);

    // Verify configs were set
    for (uint256 i = 0; i < destChainConfigs.length; ++i) {
      (bool allowlistEnabled, address ccvProxy,) =
        s_commitOnRamp.getDestChainConfig(destChainConfigs[i].destChainSelector);
      assertEq(ccvProxy, destChainConfigs[i].ccvProxy);
      assertEq(allowlistEnabled, destChainConfigs[i].allowlistEnabled);
    }
  }

  function test_applyDestChainConfigUpdates_UpdateExisting() public {
    // First, set initial config
    BaseOnRamp.DestChainConfigArgs[] memory initialConfig = new BaseOnRamp.DestChainConfigArgs[](1);
    initialConfig[0] = _getDestChainConfig(makeAddr("initialCcvProxy"), DEST_CHAIN_SELECTOR, false);

    s_commitOnRamp.applyDestChainConfigUpdates(initialConfig);

    // Now update the same destination chain
    BaseOnRamp.DestChainConfigArgs[] memory updatedConfig = new BaseOnRamp.DestChainConfigArgs[](1);
    updatedConfig[0] = _getDestChainConfig(makeAddr("updatedCcvProxy"), DEST_CHAIN_SELECTOR, true);

    vm.expectEmit();
    emit BaseOnRamp.DestChainConfigSet(
      DEST_CHAIN_SELECTOR, updatedConfig[0].ccvProxy, updatedConfig[0].allowlistEnabled
    );

    s_commitOnRamp.applyDestChainConfigUpdates(updatedConfig);

    // Verify config was updated
    (bool allowlistEnabled, address ccvProxy,) = s_commitOnRamp.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(ccvProxy, updatedConfig[0].ccvProxy);
    assertEq(allowlistEnabled, updatedConfig[0].allowlistEnabled);
  }

  function test_applyDestChainConfigUpdates_WithZeroCcvProxy() public {
    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = _getDestChainConfig(address(0), 12345, false);

    vm.expectEmit();
    emit BaseOnRamp.DestChainConfigSet(destChainConfigs[0].destChainSelector, address(0), false);

    s_commitOnRamp.applyDestChainConfigUpdates(destChainConfigs);

    (, address ccvProxy,) = s_commitOnRamp.getDestChainConfig(12345);
    assertEq(ccvProxy, address(0));
  }

  function test_applyDestChainConfigUpdates_EmptyArray() public {
    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](0);

    // Should succeed with no changes
    s_commitOnRamp.applyDestChainConfigUpdates(destChainConfigs);
  }

  function test_applyDestChainConfigUpdates_MultipleUpdatesAndAdds() public {
    BaseOnRamp.DestChainConfigArgs[] memory configs = new BaseOnRamp.DestChainConfigArgs[](3);

    configs[0] = _getDestChainConfig(makeAddr("updatedCcvProxy"), DEST_CHAIN_SELECTOR, true);
    configs[1] = _getDestChainConfig(makeAddr("newCcvProxy1"), 11111, false);
    configs[2] = _getDestChainConfig(makeAddr("newCcvProxy2"), 22222, true);

    for (uint256 i = 0; i < configs.length; ++i) {
      vm.expectEmit();
      emit BaseOnRamp.DestChainConfigSet(configs[i].destChainSelector, configs[i].ccvProxy, configs[i].allowlistEnabled);
    }

    s_commitOnRamp.applyDestChainConfigUpdates(configs);

    // Verify all configs
    for (uint256 i = 0; i < configs.length; ++i) {
      (bool allowlistEnabled, address ccvProxy,) = s_commitOnRamp.getDestChainConfig(configs[i].destChainSelector);
      assertEq(ccvProxy, configs[i].ccvProxy);
      assertEq(allowlistEnabled, configs[i].allowlistEnabled);
    }
  }

  // Reverts

  function test_applyDestChainConfigUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);

    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = _getDestChainConfig(makeAddr("ccvProxy"), 12345, false);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_commitOnRamp.applyDestChainConfigUpdates(destChainConfigs);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_ZeroDestChainSelector() public {
    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = _getDestChainConfig(makeAddr("ccvProxy"), 0, false);

    vm.expectRevert(abi.encodeWithSelector(BaseOnRamp.InvalidDestChainConfig.selector, 0));
    s_commitOnRamp.applyDestChainConfigUpdates(destChainConfigs);
  }
}
