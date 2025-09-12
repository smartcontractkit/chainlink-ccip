// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitOnRamp} from "../../../onRamp/CommitOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";

contract CommitOnRamp_constructor is CommitOnRampSetup {
  function test_constructor() public {
    address expectedFeeQuoter = address(s_feeQuoter);
    address expectedFeeAggregator = FEE_AGGREGATOR;
    address expectedAllowlistAdmin = ALLOWLIST_ADMIN;

    // Expect ConfigSet event for the deployment below.
    vm.expectEmit();
    emit CommitOnRamp.ConfigSet(
      _createDynamicConfigArgs(expectedFeeQuoter, expectedFeeAggregator, expectedAllowlistAdmin)
    );

    CommitOnRamp newOnRamp =
      new CommitOnRamp(_createDynamicConfigArgs(expectedFeeQuoter, expectedFeeAggregator, expectedAllowlistAdmin));

    // Verify dynamic config.
    CommitOnRamp.DynamicConfig memory dynamicConfig = newOnRamp.getDynamicConfig();
    assertEq(dynamicConfig.feeQuoter, expectedFeeQuoter);
    assertEq(dynamicConfig.feeAggregator, expectedFeeAggregator);
    assertEq(dynamicConfig.allowlistAdmin, expectedAllowlistAdmin);
  }

  // Reverts

  function test_constructor_RevertWhen_InvalidConfig_FeeQuoterZeroAddress() public {
    vm.expectRevert(CommitOnRamp.InvalidConfig.selector);
    new CommitOnRamp(
      _createDynamicConfigArgs(address(0), FEE_AGGREGATOR, ALLOWLIST_ADMIN) // Zero fee quoter address
    );
  }

  function test_constructor_RevertWhen_InvalidConfig_FeeAggregatorZeroAddress() public {
    vm.expectRevert(CommitOnRamp.InvalidConfig.selector);
    new CommitOnRamp(
      _createDynamicConfigArgs(address(s_feeQuoter), address(0), ALLOWLIST_ADMIN) // Zero fee aggregator address
    );
  }
}
