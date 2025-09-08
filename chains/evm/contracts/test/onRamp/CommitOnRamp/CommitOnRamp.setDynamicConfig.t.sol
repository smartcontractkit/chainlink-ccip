// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitOnRamp} from "../../../onRamp/CommitOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CommitOnRamp_setDynamicConfig is CommitOnRampSetup {
  function test_setDynamicConfig() public {
    CommitOnRamp.DynamicConfig memory newConfig =
      _getDynamicConfig(makeAddr("newFeeQuoter"), makeAddr("newFeeAggregator"), makeAddr("newAllowlistAdmin"));

    vm.expectEmit();
    emit CommitOnRamp.ConfigSet(
      CommitOnRamp.StaticConfig({rmnRemote: address(s_mockRMNRemote), nonceManager: address(s_nonceManager)}), newConfig
    );

    s_commitOnRamp.setDynamicConfig(newConfig);

    CommitOnRamp.DynamicConfig memory retrievedConfig = s_commitOnRamp.getDynamicConfig();
    assertEq(retrievedConfig.feeQuoter, newConfig.feeQuoter);
    assertEq(retrievedConfig.feeAggregator, newConfig.feeAggregator);
    assertEq(retrievedConfig.allowlistAdmin, newConfig.allowlistAdmin);
  }

  function test_setDynamicConfig_MultipleTimes() public {
    // First update
    CommitOnRamp.DynamicConfig memory firstConfig =
      _getDynamicConfig(makeAddr("firstFeeQuoter"), makeAddr("firstFeeAggregator"), makeAddr("firstAllowlistAdmin"));

    s_commitOnRamp.setDynamicConfig(firstConfig);

    CommitOnRamp.DynamicConfig memory retrievedConfig = s_commitOnRamp.getDynamicConfig();
    assertEq(retrievedConfig.feeQuoter, firstConfig.feeQuoter);
    assertEq(retrievedConfig.feeAggregator, firstConfig.feeAggregator);
    assertEq(retrievedConfig.allowlistAdmin, firstConfig.allowlistAdmin);

    // Second update
    CommitOnRamp.DynamicConfig memory secondConfig =
      _getDynamicConfig(makeAddr("secondFeeQuoter"), makeAddr("secondFeeAggregator"), makeAddr("secondAllowlistAdmin"));

    vm.expectEmit();
    emit CommitOnRamp.ConfigSet(
      CommitOnRamp.StaticConfig({rmnRemote: address(s_mockRMNRemote), nonceManager: address(s_nonceManager)}),
      secondConfig
    );

    s_commitOnRamp.setDynamicConfig(secondConfig);

    retrievedConfig = s_commitOnRamp.getDynamicConfig();
    assertEq(retrievedConfig.feeQuoter, secondConfig.feeQuoter);
    assertEq(retrievedConfig.feeAggregator, secondConfig.feeAggregator);
    assertEq(retrievedConfig.allowlistAdmin, secondConfig.allowlistAdmin);
  }

  function test_setDynamicConfig_WithZeroAllowlistAdmin() public {
    CommitOnRamp.DynamicConfig memory newConfig = _getDynamicConfig(
      makeAddr("newFeeQuoter"),
      makeAddr("newFeeAggregator"),
      address(0) // Zero allowlist admin should be allowed
    );

    s_commitOnRamp.setDynamicConfig(newConfig);

    CommitOnRamp.DynamicConfig memory retrievedConfig = s_commitOnRamp.getDynamicConfig();
    assertEq(retrievedConfig.allowlistAdmin, address(0));
  }

  // Reverts

  function test_setDynamicConfig_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);

    CommitOnRamp.DynamicConfig memory newConfig =
      _getDynamicConfig(makeAddr("newFeeQuoter"), makeAddr("newFeeAggregator"), makeAddr("newAllowlistAdmin"));

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_commitOnRamp.setDynamicConfig(newConfig);
  }

  function test_setDynamicConfig_RevertWhen_FeeQuoterZeroAddress() public {
    CommitOnRamp.DynamicConfig memory newConfig = _getDynamicConfig(
      address(0), // Zero fee quoter
      makeAddr("newFeeAggregator"),
      makeAddr("newAllowlistAdmin")
    );

    vm.expectRevert(CommitOnRamp.InvalidConfig.selector);
    s_commitOnRamp.setDynamicConfig(newConfig);
  }

  function test_setDynamicConfig_RevertWhen_FeeAggregatorZeroAddress() public {
    CommitOnRamp.DynamicConfig memory newConfig = _getDynamicConfig(
      makeAddr("newFeeQuoter"),
      address(0), // Zero fee aggregator
      makeAddr("newAllowlistAdmin")
    );

    vm.expectRevert(CommitOnRamp.InvalidConfig.selector);
    s_commitOnRamp.setDynamicConfig(newConfig);
  }

  function test_setDynamicConfig_RevertWhen_BothFeeQuoterAndFeeAggregatorZero() public {
    CommitOnRamp.DynamicConfig memory newConfig = _getDynamicConfig(
      address(0), // Zero fee quoter
      address(0), // Zero fee aggregator
      makeAddr("newAllowlistAdmin")
    );

    vm.expectRevert(CommitOnRamp.InvalidConfig.selector);
    s_commitOnRamp.setDynamicConfig(newConfig);
  }
}
