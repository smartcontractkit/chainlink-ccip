// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract OnRamp_setDynamicConfig is OnRampSetup {
  function test_SetDynamicConfig() public {
    address newFeeQuoter = makeAddr("newFeeQuoter");
    address newFeeAggregator = makeAddr("newFeeAggregator");

    OnRamp.DynamicConfig memory newConfig =
      OnRamp.DynamicConfig({feeQuoter: newFeeQuoter, reentrancyGuardEntered: false, feeAggregator: newFeeAggregator});

    vm.expectEmit();
    emit OnRamp.ConfigSet(s_onRamp.getStaticConfig(), newConfig);

    s_onRamp.setDynamicConfig(newConfig);

    OnRamp.DynamicConfig memory retrievedConfig = s_onRamp.getDynamicConfig();
    assertEq(retrievedConfig.feeQuoter, newFeeQuoter);
    assertEq(retrievedConfig.feeAggregator, newFeeAggregator);
    assertEq(retrievedConfig.reentrancyGuardEntered, false);
  }

  function test_SetDynamicConfig_MultipleUpdates() public {
    address feeQuoter1 = makeAddr("feeQuoter1");
    address feeAggregator1 = makeAddr("feeAggregator1");
    address feeQuoter2 = makeAddr("feeQuoter2");
    address feeAggregator2 = makeAddr("feeAggregator2");

    OnRamp.DynamicConfig memory config1 =
      OnRamp.DynamicConfig({feeQuoter: feeQuoter1, reentrancyGuardEntered: false, feeAggregator: feeAggregator1});

    s_onRamp.setDynamicConfig(config1);

    OnRamp.DynamicConfig memory retrievedConfig1 = s_onRamp.getDynamicConfig();
    assertEq(retrievedConfig1.feeQuoter, feeQuoter1);
    assertEq(retrievedConfig1.feeAggregator, feeAggregator1);

    OnRamp.DynamicConfig memory config2 =
      OnRamp.DynamicConfig({feeQuoter: feeQuoter2, reentrancyGuardEntered: false, feeAggregator: feeAggregator2});

    s_onRamp.setDynamicConfig(config2);

    OnRamp.DynamicConfig memory retrievedConfig2 = s_onRamp.getDynamicConfig();
    assertEq(retrievedConfig2.feeQuoter, feeQuoter2);
    assertEq(retrievedConfig2.feeAggregator, feeAggregator2);
  }

  // Reverts

  function test_SetDynamicConfig_RevertWhen_OnlyCallableByOwner() public {
    OnRamp.DynamicConfig memory newConfig = OnRamp.DynamicConfig({
      feeQuoter: makeAddr("feeQuoter"),
      reentrancyGuardEntered: false,
      feeAggregator: makeAddr("feeAggregator")
    });

    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_onRamp.setDynamicConfig(newConfig);
  }

  function test_SetDynamicConfig_RevertWhen_InvalidConfig_ZeroFeeQuoter() public {
    OnRamp.DynamicConfig memory newConfig = OnRamp.DynamicConfig({
      feeQuoter: address(0),
      reentrancyGuardEntered: false,
      feeAggregator: makeAddr("feeAggregator")
    });

    vm.expectRevert(OnRamp.InvalidConfig.selector);
    s_onRamp.setDynamicConfig(newConfig);
  }

  function test_SetDynamicConfig_RevertWhen_InvalidConfig_ZeroFeeAggregator() public {
    OnRamp.DynamicConfig memory newConfig =
      OnRamp.DynamicConfig({feeQuoter: makeAddr("feeQuoter"), reentrancyGuardEntered: false, feeAggregator: address(0)});

    vm.expectRevert(OnRamp.InvalidConfig.selector);
    s_onRamp.setDynamicConfig(newConfig);
  }

  function test_SetDynamicConfig_RevertWhen_InvalidConfig_ReentrancyGuardSet() public {
    OnRamp.DynamicConfig memory newConfig = OnRamp.DynamicConfig({
      feeQuoter: makeAddr("feeQuoter"),
      reentrancyGuardEntered: true,
      feeAggregator: makeAddr("feeAggregator")
    });

    vm.expectRevert(OnRamp.InvalidConfig.selector);
    s_onRamp.setDynamicConfig(newConfig);
  }
}
