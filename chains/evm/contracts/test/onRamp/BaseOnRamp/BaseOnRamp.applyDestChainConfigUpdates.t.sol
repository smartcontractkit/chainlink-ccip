// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {BaseOnRampSetup} from "./BaseOnRampSetup.t.sol";

contract BaseOnRamp_applyDestChainConfigUpdates is BaseOnRampSetup {
  function test_applyDestChainConfigUpdates() public {
    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);

    destChainConfigs[0] = _getDestChainConfig(makeAddr("ccvProxy1"), 12345, true);

    vm.expectEmit();
    emit BaseOnRamp.DestChainConfigSet(
      destChainConfigs[0].destChainSelector, destChainConfigs[0].ccvProxy, destChainConfigs[0].allowlistEnabled
    );

    s_baseOnRamp.applyDestChainConfigUpdates(destChainConfigs);

    // Verify config was set.
    (bool allowlistEnabled, address ccvProxy,) = s_baseOnRamp.getDestChainConfig(destChainConfigs[0].destChainSelector);
    assertEq(ccvProxy, destChainConfigs[0].ccvProxy);
    assertEq(allowlistEnabled, destChainConfigs[0].allowlistEnabled);
  }

  // Reverts

  function test_applyDestChainConfigUpdates_RevertWhen_InvalidDestChainConfig_ZeroDestChainSelector() public {
    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = _getDestChainConfig(makeAddr("ccvProxy"), 0, false);

    vm.expectRevert(abi.encodeWithSelector(BaseOnRamp.InvalidDestChainConfig.selector, 0));
    s_baseOnRamp.applyDestChainConfigUpdates(destChainConfigs);
  }
}
