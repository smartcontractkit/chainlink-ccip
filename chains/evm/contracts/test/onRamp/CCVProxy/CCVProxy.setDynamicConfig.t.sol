// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVProxy} from "../../../onRamp/CCVProxy.sol";
import {CCVProxySetup} from "./CCVProxySetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CCVProxy_setDynamicConfig is CCVProxySetup {
  function test_SetDynamicConfig() public {
    address newFeeQuoter = makeAddr("newFeeQuoter");
    address newFeeAggregator = makeAddr("newFeeAggregator");

    CCVProxy.DynamicConfig memory newConfig =
      CCVProxy.DynamicConfig({feeQuoter: newFeeQuoter, reentrancyGuardEntered: false, feeAggregator: newFeeAggregator});

    vm.expectEmit();
    emit CCVProxy.ConfigSet(s_ccvProxy.getStaticConfig(), newConfig);

    s_ccvProxy.setDynamicConfig(newConfig);

    CCVProxy.DynamicConfig memory retrievedConfig = s_ccvProxy.getDynamicConfig();
    assertEq(retrievedConfig.feeQuoter, newFeeQuoter);
    assertEq(retrievedConfig.feeAggregator, newFeeAggregator);
    assertEq(retrievedConfig.reentrancyGuardEntered, false);
  }

  function test_SetDynamicConfig_MultipleUpdates() public {
    address feeQuoter1 = makeAddr("feeQuoter1");
    address feeAggregator1 = makeAddr("feeAggregator1");
    address feeQuoter2 = makeAddr("feeQuoter2");
    address feeAggregator2 = makeAddr("feeAggregator2");

    CCVProxy.DynamicConfig memory config1 =
      CCVProxy.DynamicConfig({feeQuoter: feeQuoter1, reentrancyGuardEntered: false, feeAggregator: feeAggregator1});

    s_ccvProxy.setDynamicConfig(config1);

    CCVProxy.DynamicConfig memory retrievedConfig1 = s_ccvProxy.getDynamicConfig();
    assertEq(retrievedConfig1.feeQuoter, feeQuoter1);
    assertEq(retrievedConfig1.feeAggregator, feeAggregator1);

    CCVProxy.DynamicConfig memory config2 =
      CCVProxy.DynamicConfig({feeQuoter: feeQuoter2, reentrancyGuardEntered: false, feeAggregator: feeAggregator2});

    s_ccvProxy.setDynamicConfig(config2);

    CCVProxy.DynamicConfig memory retrievedConfig2 = s_ccvProxy.getDynamicConfig();
    assertEq(retrievedConfig2.feeQuoter, feeQuoter2);
    assertEq(retrievedConfig2.feeAggregator, feeAggregator2);
  }

  // Reverts

  function test_SetDynamicConfig_RevertWhen_OnlyCallableByOwner() public {
    CCVProxy.DynamicConfig memory newConfig = CCVProxy.DynamicConfig({
      feeQuoter: makeAddr("feeQuoter"),
      reentrancyGuardEntered: false,
      feeAggregator: makeAddr("feeAggregator")
    });

    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_ccvProxy.setDynamicConfig(newConfig);
  }

  function test_SetDynamicConfig_RevertWhen_InvalidConfig_ZeroFeeQuoter() public {
    CCVProxy.DynamicConfig memory newConfig = CCVProxy.DynamicConfig({
      feeQuoter: address(0),
      reentrancyGuardEntered: false,
      feeAggregator: makeAddr("feeAggregator")
    });

    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    s_ccvProxy.setDynamicConfig(newConfig);
  }

  function test_SetDynamicConfig_RevertWhen_InvalidConfig_ZeroFeeAggregator() public {
    CCVProxy.DynamicConfig memory newConfig = CCVProxy.DynamicConfig({
      feeQuoter: makeAddr("feeQuoter"),
      reentrancyGuardEntered: false,
      feeAggregator: address(0)
    });

    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    s_ccvProxy.setDynamicConfig(newConfig);
  }

  function test_SetDynamicConfig_RevertWhen_InvalidConfig_ReentrancyGuardSet() public {
    CCVProxy.DynamicConfig memory newConfig = CCVProxy.DynamicConfig({
      feeQuoter: makeAddr("feeQuoter"),
      reentrancyGuardEntered: true,
      feeAggregator: makeAddr("feeAggregator")
    });

    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    s_ccvProxy.setDynamicConfig(newConfig);
  }
}
